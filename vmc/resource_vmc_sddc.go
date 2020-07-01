/* Copyright 2019 VMware, Inc.
   SPDX-License-Identifier: MPL-2.0 */

package vmc

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	autoscalercluster "github.com/vmware/vsphere-automation-sdk-go/services/vmc/autoscaler/api/orgs/sddcs/clusters"
	autoscalermodel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/orgs"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/orgs/sddcs"
)

func resourceSddc() *schema.Resource {
	return &schema.Resource{
		Create: resourceSddcCreate,
		Read:   resourceSddcRead,
		Update: resourceSddcUpdate,
		Delete: resourceSddcDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(300 * time.Minute),
			Update: schema.DefaultTimeout(300 * time.Minute),
			Delete: schema.DefaultTimeout(180 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"storage_capacity": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"15TB", "20TB", "25TB", "30TB", "35TB"}, false),
			},
			"sddc_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"account_link_sddc_config": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"customer_subnet_ids": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								// Optional: true,
							},
							Optional: true,
						},
						"connected_account_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				Optional: true,
				ForceNew: true,
			},
			"vpc_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"num_host": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(1),
			},
			"sddc_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"vxlan_subnet": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"delay_account_link": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"provider_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "AWS",
				ValidateFunc: validation.StringInSlice([]string{
					"AWS", "ZEROCLOUD"}, false),
			},
			"skip_creating_vxlan": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},
			"sso_domain": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "vmc.local",
			},
			"sddc_template_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"deployment_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  SingleAvailabilityZone,
				ValidateFunc: validation.StringInSlice([]string{
					SingleAvailabilityZone, MultiAvailabilityZone,
				}, false),
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.All(
					validation.NoZeroValues,
				),
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_instance_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice(
					[]string{HostInstancetypeI3, HostInstancetypeR5, HostInstancetypeI3EN}, false),
			},
			"sddc_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vc_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsxt_reverse_proxy_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_info": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"availability_zones": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"edrs_policy_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  StorageScaleUpPolicyType,
				ValidateFunc: validation.StringInSlice(
					[]string{StorageScaleUpPolicyType, CostPolicyType, PerformancePolicyType, RapidScaleUpPolicyType}, false),
				Description: "The EDRS policy type. This can either be 'cost', 'performance', 'storage-scaleup' or 'rapid-scaleup'. Default value : storage-scaleup ",
			},
			"enable_edrs": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "True if EDRS is enabled",
			},
			"min_hosts": {
				Type:         schema.TypeInt,
				Default:      3,
				Optional:     true,
				ValidateFunc: validation.IntBetween(3, 16),
				Description:  "The minimum number of hosts that the cluster can scale in to.",
			},
			"max_hosts": {
				Type:         schema.TypeInt,
				Default:      16,
				Optional:     true,
				ValidateFunc: validation.IntBetween(3, 16),
				Description:  "The maximum number of hosts that the cluster can scale out to.",
			},
		},
		CustomizeDiff: func(d *schema.ResourceDiff, meta interface{}) error {

			deploymentType := d.Get("deployment_type").(string)
			numHosts := d.Get("num_host").(int)

			if deploymentType == MultiAvailabilityZone && numHosts < MinMultiAZHosts {
				return fmt.Errorf("for MulitAZ deployment type number of hosts must be atleast %d ", MinMultiAZHosts)
			}

			if deploymentType == MultiAvailabilityZone {
				accountLinkSddcConfig := d.Get("account_link_sddc_config").([]interface{})
				for _, config := range accountLinkSddcConfig {
					c := config.(map[string]interface{})
					if len(c["customer_subnet_ids"].([]interface{})) < 2 {
						return fmt.Errorf("deployment type %s requires 2 subnets one in each availability zone ", deploymentType)
					}
				}
			}

			newInstanceType := d.Get("host_instance_type").(string)

			switch newInstanceType {

			case HostInstancetypeI3, HostInstancetypeI3EN:

				if d.Get("storage_capacity").(string) != "" {

					return fmt.Errorf("storage_capacity is not supported for host_instance_type %q", newInstanceType)

				}
			case HostInstancetypeR5:

				if d.Get("storage_capacity").(string) == "" {

					return fmt.Errorf("storage_capacity is required for host_instance_type %q", newInstanceType)

				}

			}
			return nil
		},
	}
}

func resourceSddcCreate(d *schema.ResourceData, m interface{}) error {
	var storageCapacityConverted int64
	connectorWrapper := m.(*ConnectorWrapper)
	sddcClient := orgs.NewDefaultSddcsClient(connectorWrapper)
	orgID := connectorWrapper.OrgID

	storageCapacity := d.Get("storage_capacity").(string)
	if len(strings.TrimSpace(storageCapacity)) > 0 {
		storageCapacityConverted = ConvertStorageCapacitytoInt(storageCapacity)
	}

	sddcName := d.Get("sddc_name").(string)
	vpcCidr := d.Get("vpc_cidr").(string)
	numHost := d.Get("num_host").(int)
	sddcType := d.Get("sddc_type").(string)

	var sddcTypePtr *string
	if sddcType != "" {
		sddcTypePtr = &sddcType
	}
	vxlanSubnet := d.Get("vxlan_subnet").(string)
	delayAccountLink := d.Get("delay_account_link").(bool)
	accountLinkConfig := &model.AccountLinkConfig{
		DelayAccountLink: &delayAccountLink,
	}
	providerType := d.Get("provider_type").(string)
	skipCreatingVxlan := d.Get("skip_creating_vxlan").(bool)
	ssoDomain := d.Get("sso_domain").(string)
	sddcTemplateID := d.Get("sddc_template_id").(string)
	deploymentType := d.Get("deployment_type").(string)
	region := d.Get("region").(string)
	accountLinkSddcConfig := expandAccountLinkSddcConfig(d.Get("account_link_sddc_config").([]interface{}))
	hostInstanceType := model.HostInstanceTypes(d.Get("host_instance_type").(string))

	var awsSddcConfig = &model.AwsSddcConfig{
		StorageCapacity:       &storageCapacityConverted,
		Name:                  sddcName,
		VpcCidr:               &vpcCidr,
		NumHosts:              int64(numHost),
		SddcType:              sddcTypePtr,
		VxlanSubnet:           &vxlanSubnet,
		AccountLinkConfig:     accountLinkConfig,
		Provider:              providerType,
		SkipCreatingVxlan:     &skipCreatingVxlan,
		AccountLinkSddcConfig: accountLinkSddcConfig,
		SsoDomain:             &ssoDomain,
		SddcTemplateId:        &sddcTemplateID,
		DeploymentType:        &deploymentType,
		Region:                region,
		HostInstanceType:      &hostInstanceType,
	}

	// Create a Sddc
	task, err := sddcClient.Create(orgID, *awsSddcConfig, nil)
	if err != nil {
		return HandleCreateError("SDDC", err)
	}

	// Wait until Sddc is created
	sddcID := task.ResourceId
	d.SetId(*sddcID)
	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		tasksClient := orgs.NewDefaultTasksClient(connectorWrapper)
		task, err := tasksClient.Get(orgID, task.Id)
		if err != nil {
			if err.Error() == (errors.Unauthenticated{}.Error()) {
				log.Print("Auth error", err.Error(), errors.Unauthenticated{}.Error())
				err = connectorWrapper.authenticate()
				if err != nil {
					return resource.NonRetryableError(fmt.Errorf("authentication error from Cloud Service Provider : %s", err))
				}
				return resource.RetryableError(fmt.Errorf("instance creation still in progress"))
			}
			return resource.NonRetryableError(fmt.Errorf("error describing instance: %s", err))

		}
		if *task.Status != "FINISHED" {
			return resource.RetryableError(fmt.Errorf("expected instance to be created but was in state %s", *task.Status))
		}
		return resource.NonRetryableError(resourceSddcRead(d, m))
	})
}

func resourceSddcRead(d *schema.ResourceData, m interface{}) error {
	connector := (m.(*ConnectorWrapper)).Connector
	sddcID := d.Id()
	orgID := (m.(*ConnectorWrapper)).OrgID
	sddc, err := GetSDDC(connector, orgID, sddcID)
	if err != nil {
		return HandleReadError(d, "SDDC", sddcID, err)
	}

	if *sddc.SddcState == "DELETED" {
		log.Printf("Unable to retrieve SDDC with ID %s", sddc.Id)
		d.SetId("")
		return nil
	}

	d.SetId(sddc.Id)

	d.Set("sddc_name", sddc.Name)
	d.Set("updated", sddc.Updated)
	d.Set("user_id", sddc.UserId)
	d.Set("updated_by_user_id", sddc.UpdatedByUserId)
	d.Set("created", sddc.Created)
	d.Set("version", sddc.Version)
	d.Set("updated_by_user_name", sddc.UpdatedByUserName)
	d.Set("user_name", sddc.UserName)
	d.Set("org_id", sddc.OrgId)
	d.Set("sddc_type", sddc.SddcType)
	d.Set("provider", sddc.Provider)
	d.Set("account_link_state", sddc.AccountLinkState)
	d.Set("sddc_access_state", sddc.SddcAccessState)
	d.Set("sddc_type", sddc.SddcType)
	d.Set("sddc_state", sddc.SddcState)
	d.Set("num_host", len(sddc.ResourceConfig.EsxHosts))

	if sddc.ResourceConfig != nil {
		d.Set("vc_url", sddc.ResourceConfig.VcUrl)
		d.Set("cloud_username", sddc.ResourceConfig.CloudUsername)
		d.Set("cloud_password", sddc.ResourceConfig.CloudPassword)
		d.Set("nsxt_reverse_proxy_url", sddc.ResourceConfig.NsxApiPublicEndpointUrl)
		d.Set("availability_zones", sddc.ResourceConfig.AvailabilityZones)
	}
	sddcClient := sddcs.NewDefaultPrimaryclusterClient(connector)
	primaryCluster, err := sddcClient.Get(orgID, sddcID)
	if err != nil {
		return HandleReadError(d, "Primary Cluster", sddcID, err)
	}

	d.Set("cluster_id", primaryCluster.ClusterId)
	cluster := map[string]string{}
	cluster["cluster_name"] = *primaryCluster.ClusterName
	cluster["cluster_state"] = *primaryCluster.ClusterState
	cluster["host_instance_type"] = *primaryCluster.EsxHostInfo.InstanceType
	d.Set("cluster_info", cluster)

	edrsPolicyClient := autoscalercluster.NewDefaultEdrsPolicyClient(connector)
	edrsPolicy, err := edrsPolicyClient.Get(orgID, sddcID, primaryCluster.ClusterId)

	d.Set("edrs_policy_type", *edrsPolicy.PolicyType)
	d.Set("enable_edrs", edrsPolicy.EnableEdrs)
	d.Set("max_hosts", *edrsPolicy.MaxHosts)
	d.Set("min_hosts", *edrsPolicy.MinHosts)
	return nil
}

func resourceSddcDelete(d *schema.ResourceData, m interface{}) error {
	connector := (m.(*ConnectorWrapper)).Connector
	sddcClient := orgs.NewDefaultSddcsClient(connector)
	sddcID := d.Id()
	orgID := (m.(*ConnectorWrapper)).OrgID

	task, err := sddcClient.Delete(orgID, sddcID, nil, nil, nil)
	if err != nil {
		return HandleDeleteError("SDDC", sddcID, err)
	}
	tasksClient := orgs.NewDefaultTasksClient(connector)
	return resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		task, err := tasksClient.Get(orgID, task.Id)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("error while deleting SDDC %s: %v", sddcID, err))
		}
		if *task.Status != "FINISHED" {
			return resource.RetryableError(fmt.Errorf("expected instance to be deleted but was in state %s", *task.Status))
		}
		d.SetId("")
		return resource.NonRetryableError(nil)
	})
}

func resourceSddcUpdate(d *schema.ResourceData, m interface{}) error {
	connector := (m.(*ConnectorWrapper)).Connector
	esxsClient := sddcs.NewDefaultEsxsClient(connector)
	sddcID := d.Id()
	orgID := (m.(*ConnectorWrapper)).OrgID

	// Add,remove hosts
	if d.HasChange("num_host") {
		oldTmp, newTmp := d.GetChange("num_host")
		oldNum := oldTmp.(int)
		newNum := newTmp.(int)

		action := "add"
		diffNum := newNum - oldNum

		if newNum < oldNum {
			action = "remove"
			diffNum = oldNum - newNum
		}
		if d.Get("deployment_type").(string) == MultiAvailabilityZone && diffNum%2 != 0 {

			return fmt.Errorf("for multiAZ deployment type, SDDC hosts must be added in pairs across availability zones")
		}
		esxConfig := model.EsxConfig{
			NumHosts: int64(diffNum),
		}

		task, err := esxsClient.Create(orgID, sddcID, esxConfig, &action)

		if err != nil {
			return HandleUpdateError("SDDC", err)
		}
		tasksClient := orgs.NewDefaultTasksClient(connector)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			task, err := tasksClient.Get(orgID, task.Id)
			if err != nil {
				return resource.NonRetryableError(fmt.Errorf("error while waiting for task %s: %v", task.Id, err))
			}
			if *task.Status != "FINISHED" {
				return resource.RetryableError(fmt.Errorf("expected hosts to be updated but were in state %s", *task.Status))
			}
			return resource.NonRetryableError(resourceSddcRead(d, m))
		})
		if err != nil {
			return err
		}
	}
	// Update sddc name
	if d.HasChange("sddc_name") {
		sddcClient := orgs.NewDefaultSddcsClient(connector)
		newSDDCName := d.Get("sddc_name").(string)
		sddcPatchRequest := model.SddcPatchRequest{
			Name: &newSDDCName,
		}
		sddc, err := sddcClient.Patch(orgID, sddcID, sddcPatchRequest)

		if err != nil {
			return HandleUpdateError("SDDC", err)
		}
		d.Set("sddc_name", sddc.Name)
	}

	if d.HasChange("edrs_policy_type") || d.HasChange("enable_edrs") || d.HasChange("min_hosts") || d.HasChange("max_hosts") {
		edrsPolicyClient := autoscalercluster.NewDefaultEdrsPolicyClient(connector)
		clusterID := d.Get("cluster_id").(string)
		minHosts := int64(d.Get("min_hosts").(int))
		maxHosts := int64(d.Get("max_hosts").(int))
		policyType := d.Get("edrs_policy_type").(string)
		enableEDRS := d.Get("enable_edrs").(bool)
		edrsPolicy := &autoscalermodel.EdrsPolicy{
			EnableEdrs: enableEDRS,
			PolicyType: &policyType,
			MinHosts:   &minHosts,
			MaxHosts:   &maxHosts,
		}
		_, err := edrsPolicyClient.Post(orgID, sddcID, clusterID, *edrsPolicy)
		if err != nil {
			return HandleUpdateError("EDRS Policy", err)
		}
		time.Sleep(2 * time.Minute)
	}
	return resourceSddcRead(d, m)
}

func expandAccountLinkSddcConfig(l []interface{}) []model.AccountLinkSddcConfig {

	if len(l) == 0 {
		return nil
	}

	var configs []model.AccountLinkSddcConfig

	for _, config := range l {
		c := config.(map[string]interface{})
		var subnetIds []string
		for _, subnetID := range c["customer_subnet_ids"].([]interface{}) {
			subnetIds = append(subnetIds, subnetID.(string))
		}
		var connectedAccId = c["connected_account_id"].(string)
		con := model.AccountLinkSddcConfig{
			CustomerSubnetIds:  subnetIds,
			ConnectedAccountId: &connectedAccId,
		}

		configs = append(configs, con)
	}
	return configs
}
