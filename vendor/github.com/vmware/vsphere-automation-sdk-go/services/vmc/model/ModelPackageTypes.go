/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vmc.model.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package model

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"time"
)


type HostInstanceTypes string

const (
	HostInstanceTypes_I3_METAL HostInstanceTypes = "I3_METAL"
	HostInstanceTypes_R5_METAL HostInstanceTypes = "R5_METAL"
	HostInstanceTypes_I3EN_METAL HostInstanceTypes = "I3EN_METAL"
)

func (h HostInstanceTypes) HostInstanceTypes() bool {
	switch h {
	case HostInstanceTypes_I3_METAL:
		return true
	case HostInstanceTypes_R5_METAL:
		return true
	case HostInstanceTypes_I3EN_METAL:
		return true
	default:
		return false
	}
}


type OfferType string

const (
	OfferType_TERM OfferType = "TERM"
	OfferType_ON_DEMAND OfferType = "ON_DEMAND"
)

func (o OfferType) OfferType() bool {
	switch o {
	case OfferType_TERM:
		return true
	case OfferType_ON_DEMAND:
		return true
	default:
		return false
	}
}


type AbstractEntity struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
}

func (s AbstractEntity) GetType__() bindings.BindingType {
	return AbstractEntityBindingType()
}

func (s AbstractEntity) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AbstractEntity._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AccountLinkConfig struct {
    // Boolean flag identifying whether account linking should be delayed or not for the SDDC.
	DelayAccountLink *bool
}

func (s AccountLinkConfig) GetType__() bindings.BindingType {
	return AccountLinkConfigBindingType()
}

func (s AccountLinkConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AccountLinkConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AccountLinkSddcConfig struct {
	CustomerSubnetIds []string
    // The ID of the customer connected account to work with.
	ConnectedAccountId *string
}

func (s AccountLinkSddcConfig) GetType__() bindings.BindingType {
	return AccountLinkSddcConfigBindingType()
}

func (s AccountLinkSddcConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AccountLinkSddcConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Source or Destination for firewall rule. Default is 'any'.
type AddressFWSourceDestination struct {
    // Exclude the specified source or destination.
	Exclude *bool
    // List of string. Can specify single IP address, range of IP address, or in CIDR format. Can define multiple.
	IpAddress []string
    // List of string. Id of cluster, datacenter, distributedPortGroup, legacyPortGroup, VirtualMachine, vApp, resourcePool, logicalSwitch, IPSet, securityGroup. Can define multiple.
	GroupingObjectId []string
    // List of string. Possible values are vnic-index-[1-9], vse, external or internal. Can define multiple.
	VnicGroupId []string
}

func (s AddressFWSourceDestination) GetType__() bindings.BindingType {
	return AddressFWSourceDestinationBindingType()
}

func (s AddressFWSourceDestination) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AddressFWSourceDestination._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Agent struct {
    // The internal IP address of the agent which is provided by the underlying cloud provider
	InternalIp *string
    // The accessible URL of the agent service, it is resolved to public IP address from the Internet and private IP address within SDDC
	AgentUrl *string
    // The internal management IP address of the agent exposed to the SDDC, which might be different from the internal IP
	ManagementIp *string
    // Boolean flag to indicate if the agent is using FQDN in the certificate
	HostnameVerifierEnabled *bool
    // Boolean flag to indicate if the agent is the master, only the master Agent is accessible
	Master *bool
    // Network netmask of the agent
	NetworkNetmask *string
    // Network gateway of the agent
	NetworkGateway *string
    // The cloud provider
	Provider string
    // Boolean flag to indicate if the agent is using CA signed certificate
	CertEnabled *bool
    // Possible values are: 
    //
    // * Agent#Agent_AGENT_STATE_NOT_READY
    // * Agent#Agent_AGENT_STATE_DEPLOYING
    // * Agent#Agent_AGENT_STATE_CUSTOMIZING
    // * Agent#Agent_AGENT_STATE_READY
    // * Agent#Agent_AGENT_STATE_DELETING
    // * Agent#Agent_AGENT_STATE_DELETED
    // * Agent#Agent_AGENT_STATE_FAILED
    //
    //  Agent state
	AgentState *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const Agent__TYPE_IDENTIFIER = "Agent"
const Agent_AGENT_STATE_NOT_READY = "NOT_READY"
const Agent_AGENT_STATE_DEPLOYING = "DEPLOYING"
const Agent_AGENT_STATE_CUSTOMIZING = "CUSTOMIZING"
const Agent_AGENT_STATE_READY = "READY"
const Agent_AGENT_STATE_DELETING = "DELETING"
const Agent_AGENT_STATE_DELETED = "DELETED"
const Agent_AGENT_STATE_FAILED = "FAILED"

func (s Agent) GetType__() bindings.BindingType {
	return AgentBindingType()
}

func (s Agent) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Agent._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// the AmiInfo used for deploying esx of the sddc
type AmiInfo struct {
    // instance type of the esx ami
	InstanceType *string
    // the region of the esx ami
	Region *string
    // the ami id for the esx
	Id *string
    // the name of the esx ami
	Name *string
}

func (s AmiInfo) GetType__() bindings.BindingType {
	return AmiInfoBindingType()
}

func (s AmiInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AmiInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NSX Edge appliance summary.
type AppliancesSummary struct {
    // vCenter MOID of the active NSX Edge appliance's data store.
	DataStoreMoidOfActiveVse *string
    // Value is true if FIPS is enabled on NSX Edge appliance.
	EnableFips *bool
    // Host name of the active NSX Edge appliance.
	HostNameOfActiveVse *string
    // NSX Edge appliance build version.
	VmBuildInfo *string
    // Value is true if NSX Edge appliances are to be deployed.
	DeployAppliances *bool
    // Communication channel used to communicate with NSX Edge appliance.
	CommunicationChannel *string
    // Name of the active NSX Edge appliance.
	VmNameOfActiveVse *string
    // Number of deployed appliances of the NSX Edge. format: int32
	NumberOfDeployedVms *int64
    // vCenter MOID of the active NSX Edge appliance's resource pool/cluster. Can be resource pool ID, e.g. resgroup-15 or cluster ID, e.g. domain-c41.
	ResourcePoolMoidOfActiveVse *string
    // Datastore name of the active NSX Edge appliance.
	DataStoreNameOfActiveVse *string
    // vCenter MOID of the active NSX Edge appliance.
	VmMoidOfActiveVse *string
    // Time stamp value when healthcheck status was last updated for the NSX Edge appliance. format: int64
	StatusFromVseUpdatedOn *int64
    // FQDN of the NSX Edge.
	Fqdn *string
    // NSX Edge appliance size.
	ApplianceSize *string
    // Resource Pool/Cluster name of the active NSX Edge appliance.
	ResourcePoolNameOfActiveVse *string
    // HA index of the active NSX Edge appliance. format: int32
	ActiveVseHaIndex *int64
    // NSX Edge appliance version.
	VmVersion *string
    // vCenter MOID of the active NSX Edge appliance's host.
	HostMoidOfActiveVse *string
}

func (s AppliancesSummary) GetType__() bindings.BindingType {
	return AppliancesSummaryBindingType()
}

func (s AppliancesSummary) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AppliancesSummary._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Application for firewall rule
type Application struct {
    // List of string. Id of service or serviceGroup groupingObject. Can define multiple.
	ApplicationId []string
    // List of protocol and ports. Can define multiple.
	Service []Nsxfirewallservice
}

func (s Application) GetType__() bindings.BindingType {
	return ApplicationBindingType()
}

func (s Application) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Application._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AvailableZoneInfo struct {
    // null
	Subnets []Subnet
    // available zone name
	Name *string
}

func (s AvailableZoneInfo) GetType__() bindings.BindingType {
	return AvailableZoneInfoBindingType()
}

func (s AvailableZoneInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AvailableZoneInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsAgent struct {
	InstanceId *string
	KeyPair *AwsKeyPair
    // The internal IP address of the agent which is provided by the underlying cloud provider
	InternalIp *string
    // The accessible URL of the agent service, it is resolved to public IP address from the Internet and private IP address within SDDC
	AgentUrl *string
    // The internal management IP address of the agent exposed to the SDDC, which might be different from the internal IP
	ManagementIp *string
    // Boolean flag to indicate if the agent is using FQDN in the certificate
	HostnameVerifierEnabled *bool
    // Boolean flag to indicate if the agent is the master, only the master Agent is accessible
	Master *bool
    // Network netmask of the agent
	NetworkNetmask *string
    // Network gateway of the agent
	NetworkGateway *string
    // The cloud provider
	Provider string
    // Boolean flag to indicate if the agent is using CA signed certificate
	CertEnabled *bool
    // Possible values are: 
    //
    // * Agent#Agent_AGENT_STATE_NOT_READY
    // * Agent#Agent_AGENT_STATE_DEPLOYING
    // * Agent#Agent_AGENT_STATE_CUSTOMIZING
    // * Agent#Agent_AGENT_STATE_READY
    // * Agent#Agent_AGENT_STATE_DELETING
    // * Agent#Agent_AGENT_STATE_DELETED
    // * Agent#Agent_AGENT_STATE_FAILED
    //
    //  Agent state
	AgentState *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsAgent__TYPE_IDENTIFIER = "AWS"

func (s AwsAgent) GetType__() bindings.BindingType {
	return AwsAgentBindingType()
}

func (s AwsAgent) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsAgent._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsCloudProvider struct {
	Regions []string
    // Name of the Cloud Provider
	Provider string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsCloudProvider__TYPE_IDENTIFIER = "AWS"

func (s AwsCloudProvider) GetType__() bindings.BindingType {
	return AwsCloudProviderBindingType()
}

func (s AwsCloudProvider) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsCloudProvider._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsCompatibleSubnets struct {
	CustomerAvailableZones []string
	VpcMap map[string]VpcInfoSubnets
}

func (s AwsCompatibleSubnets) GetType__() bindings.BindingType {
	return AwsCompatibleSubnetsBindingType()
}

func (s AwsCompatibleSubnets) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsCompatibleSubnets._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsCustomerConnectedAccount struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
	PolicyPayerArn *string
    // Provides a map of regions to availability zones from the shadow account's perspective
	RegionToAzToShadowMapping map[string]map[string]string
	OrgId *string
	CfStackName *string
    // Possible values are: 
    //
    // * AwsCustomerConnectedAccount#AwsCustomerConnectedAccount_STATE_ACTIVE
    // * AwsCustomerConnectedAccount#AwsCustomerConnectedAccount_STATE_BROKEN
    // * AwsCustomerConnectedAccount#AwsCustomerConnectedAccount_STATE_DELETED
	State *string
	AccountNumber *string
	PolicyServiceArn *string
	PolicyExternalId *string
	PolicyPayerLinkedArn *string
}
const AwsCustomerConnectedAccount_STATE_ACTIVE = "ACTIVE"
const AwsCustomerConnectedAccount_STATE_BROKEN = "BROKEN"
const AwsCustomerConnectedAccount_STATE_DELETED = "DELETED"

func (s AwsCustomerConnectedAccount) GetType__() bindings.BindingType {
	return AwsCustomerConnectedAccountBindingType()
}

func (s AwsCustomerConnectedAccount) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsCustomerConnectedAccount._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsEsxHost struct {
	InternalPublicIpPool []SddcPublicIp
	Name *string
    // Availability zone where the host is provisioned.
	AvailabilityZone *string
	EsxId *string
	Hostname *string
	Provider string
    // Backing cloud provider instance type for host.
	InstanceType *string
	MacAddress *string
	CustomProperties map[string]string
    // Possible values are: 
    //
    // * EsxHost#EsxHost_ESX_STATE_DEPLOYING
    // * EsxHost#EsxHost_ESX_STATE_INITIALIZING
    // * EsxHost#EsxHost_ESX_STATE_PROVISIONED
    // * EsxHost#EsxHost_ESX_STATE_READY
    // * EsxHost#EsxHost_ESX_STATE_DELETING
    // * EsxHost#EsxHost_ESX_STATE_DELETED
    // * EsxHost#EsxHost_ESX_STATE_FAILED
    // * EsxHost#EsxHost_ESX_STATE_ADDING_TO_VCENTER
    // * EsxHost#EsxHost_ESX_STATE_DELETING_FROM_VCENTER
    // * EsxHost#EsxHost_ESX_STATE_PENDING_CLOUD_DELETION
	EsxState *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsEsxHost__TYPE_IDENTIFIER = "AWS"

func (s AwsEsxHost) GetType__() bindings.BindingType {
	return AwsEsxHostBindingType()
}

func (s AwsEsxHost) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsEsxHost._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsKeyPair struct {
	KeyName *string
	KeyFingerprint *string
	KeyMaterial *string
}

func (s AwsKeyPair) GetType__() bindings.BindingType {
	return AwsKeyPairBindingType()
}

func (s AwsKeyPair) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsKeyPair._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsKmsInfo struct {
    // The ARN associated with the customer master key for this cluster.
	AmazonResourceName string
}

func (s AwsKmsInfo) GetType__() bindings.BindingType {
	return AwsKmsInfoBindingType()
}

func (s AwsKmsInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsKmsInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsSddcConfig struct {
	Region string
    // AWS VPC IP range. Only prefix of 16 or 20 is currently supported.
	VpcCidr *string
    // The instance type for the esx hosts in the primary cluster of the SDDC.
	HostInstanceType *HostInstanceTypes
    // skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan *bool
    // VXLAN IP subnet in CIDR for compute gateway
	VxlanSubnet *string
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_SIZE_NSX_SMALL
    // * SddcConfig#SddcConfig_SIZE_MEDIUM
    // * SddcConfig#SddcConfig_SIZE_LARGE
    // * SddcConfig#SddcConfig_SIZE_NSX_LARGE
    //
    //  The size of the vCenter and NSX appliances. \"large\" sddcSize corresponds to a 'large' vCenter appliance and 'large' NSX appliance. 'medium' sddcSize corresponds to 'medium' vCenter appliance and 'medium' NSX appliance. Value defaults to 'medium'.
	Size *string
    // The storage capacity value to be requested for the sddc primary cluster, in GiBs. If provided, instead of using the direct-attached storage, a capacity value amount of seperable storage will be used. format: int64
	StorageCapacity *int64
	Name string
    // A list of the SDDC linking configurations to use.
	AccountLinkSddcConfig []AccountLinkSddcConfig
    // If provided, will be assigned as SDDC id of the provisioned SDDC. format: UUID
	SddcId *string
	NumHosts int64
    // Denotes the sddc type , if the value is null or empty, the type is considered as default.
	SddcType *string
    // The account linking configuration, we will keep this one and remove accountLinkSddcConfig finally.
	AccountLinkConfig *AccountLinkConfig
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_PROVIDER_AWS
    // * SddcConfig#SddcConfig_PROVIDER_ZEROCLOUD
    //
    //  Determines what additional properties are available based on cloud provider.
	Provider string
    // The SSO domain name to use for vSphere users. If not specified, vmc.local will be used.
	SsoDomain *string
    // If provided, configuration from the template will applied to the provisioned SDDC. format: UUID
	SddcTemplateId *string
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_DEPLOYMENT_TYPE_SINGLEAZ
    // * SddcConfig#SddcConfig_DEPLOYMENT_TYPE_MULTIAZ
    //
    //  Denotes if request is for a SingleAZ or a MultiAZ SDDC. Default is SingleAZ.
	DeploymentType *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsSddcConfig__TYPE_IDENTIFIER = "AWS"

func (s AwsSddcConfig) GetType__() bindings.BindingType {
	return AwsSddcConfigBindingType()
}

func (s AwsSddcConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsSddcConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsSddcConnection struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
    // The CIDR block of the customer's subnet this link is in.
	CidrBlockSubnet *string
    // The corresponding connected (customer) account UUID this connection is attached to.
	ConnectedAccountId *string
    // Which group the ENIs belongs to. (deprecated)
	EniGroup *string
    // The ID of the subnet this link is to.
	SubnetId *string
    // Determines whether the CGW is present in this connection set or not. Used for multi-az deployments.
	CgwPresent *bool
    // The org this link belongs to.
	OrgId *string
    // The SDDC this link is used for.
	SddcId *string
    // The CIDR block of the customer's VPC.
	CidrBlockVpc *string
    // The order of the connection
	ConnectionOrder *int64
    // Possible values are: 
    //
    // * AwsSddcConnection#AwsSddcConnection_STATE_ACTIVE
    // * AwsSddcConnection#AwsSddcConnection_STATE_BROKEN
    // * AwsSddcConnection#AwsSddcConnection_STATE_DELETED
    //
    //  The state of the connection.
	State *string
    // Which availability zone is this connection in?
	SubnetAvailabilityZone *string
    // The VPC ID of the subnet this link is to.
	VpcId *string
    // A list of all ENIs used for this connection.
	CustomerEniInfos []CustomerEniInfo
    // The default routing table in the customer's VPC.
	DefaultRouteTable *string
}
const AwsSddcConnection_STATE_ACTIVE = "ACTIVE"
const AwsSddcConnection_STATE_BROKEN = "BROKEN"
const AwsSddcConnection_STATE_DELETED = "DELETED"

func (s AwsSddcConnection) GetType__() bindings.BindingType {
	return AwsSddcConnectionBindingType()
}

func (s AwsSddcConnection) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsSddcConnection._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsSddcResourceConfig struct {
	BackupRestoreBucket *string
	PublicIpPool []SddcPublicIp
	VpcInfo *VpcInfo
	KmsVpcEndpoint *KmsVpcEndpoint
    // maximum number of public IP that user can allocate.
	MaxNumPublicIp *int64
	AccountLinkSddcConfig []SddcLinkConfig
	VsanEncryptionConfig *VsanEncryptionConfig
	VpcInfoPeeredAgent *VpcInfo
    // Name for management appliance network.
	MgmtApplianceNetworkName *string
    // if true, NSX-T UI is enabled.
	Nsxt *bool
    // Management Gateway Id
	MgwId *string
    // URL of the NSX Manager
	NsxMgrUrl *string
    // PSC internal management IP
	PscManagementIp *string
    // URL of the PSC server
	PscUrl *string
	Cgws []string
    // Availability zones over which esx hosts are provisioned. MultiAZ SDDCs will have hosts provisioned over two availability zones while SingleAZ SDDCs will provision over one.
	AvailabilityZones []string
    // The ManagedObjectReference of the management Datastore
	ManagementDs *string
    // nsx api entire base url
	NsxApiPublicEndpointUrl *string
	CustomProperties map[string]string
    // Password for vCenter SDDC administrator
	CloudPassword *string
    // Possible values are: 
    //
    // * SddcResourceConfig#SddcResourceConfig_PROVIDER_AWS
    // * SddcResourceConfig#SddcResourceConfig_PROVIDER_ZEROCLOUD
    //
    //  Discriminator for additional properties
	Provider string
    // List of clusters in the SDDC.
	Clusters []Cluster
    // vCenter internal management IP
	VcManagementIp *string
	SddcNetworks []string
    // Username for vCenter SDDC administrator
	CloudUsername *string
	EsxHosts []AwsEsxHost
    // NSX Manager internal management IP
	NsxMgrManagementIp *string
    // unique id of the vCenter server
	VcInstanceId *string
    // Cluster Id to add ESX workflow
	EsxClusterId *string
    // vCenter public IP
	VcPublicIp *string
    // skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan *bool
    // URL of the vCenter server
	VcUrl *string
	SddcManifest *SddcManifest
    // VXLAN IP subnet
	VxlanSubnet *string
    // Group name for vCenter SDDC administrator
	CloudUserGroup *string
	ManagementRp *string
    // region in which sddc is deployed
	Region *string
    // Availability zone where the witness node is provisioned for a MultiAZ SDDC. This is null for a SingleAZ SDDC.
	WitnessAvailabilityZone *string
    // sddc identifier
	SddcId *string
	PopAgentXeniConnection *PopAgentXeniConnection
	SddcSize *SddcSize
    // List of Controller IPs
	NsxControllerIps []string
    // ESX host subnet
	EsxHostSubnet *string
    // The SSO domain name to use for vSphere users
	SsoDomain *string
    // Possible values are: 
    //
    // * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_SINGLE_AZ
    // * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_MULTI_AZ
    //
    //  Denotes if this is a SingleAZ SDDC or a MultiAZ SDDC.
	DeploymentType *string
	NsxtAddons *NsxtAddons
    // if true, use the private IP addresses to register DNS records for the management VMs
	DnsWithManagementVmPrivateIp *bool
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsSddcResourceConfig__TYPE_IDENTIFIER = "AWS"

func (s AwsSddcResourceConfig) GetType__() bindings.BindingType {
	return AwsSddcResourceConfigBindingType()
}

func (s AwsSddcResourceConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsSddcResourceConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsSubnet struct {
    // The connected account ID this subnet is accessible through. This is an internal ID formatted as a UUID specific to Skyscraper.
	ConnectedAccountId *string
    // The region this subnet is in, usually in the form of country code, general location, and a number (ex. us-west-2).
	RegionName *string
    // The availability zone this subnet is in, which should be the region name plus one extra letter (ex. us-west-2a).
	AvailabilityZone *string
    // The subnet ID in AWS, provided in the form 'subnet-######'.
	SubnetId *string
    // The CIDR block of the subnet, in the form of '#.#.#.#/#'.
	SubnetCidrBlock *string
    // Flag indicating whether this subnet is compatible. If true, this is a valid choice for the customer to deploy a SDDC in.
	IsCompatible *bool
    // The VPC ID the subnet resides in within AWS. Tends to be 'vpc-#######'.
	VpcId *string
    // The CIDR block of the VPC, in the form of '#.#.#.#/#'.
	VpcCidrBlock *string
    // Optional field (may not be provided by AWS), indicates the found name tag for the subnet.
	Name *string
}

func (s AwsSubnet) GetType__() bindings.BindingType {
	return AwsSubnetBindingType()
}

func (s AwsSubnet) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsSubnet._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// CA certificate list. Optional.
type CaCertificates struct {
	CaCertificate []string
}

func (s CaCertificates) GetType__() bindings.BindingType {
	return CaCertificatesBindingType()
}

func (s CaCertificates) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CaCertificates._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Statistics data for each vnic.
type CbmStatistic struct {
    // Vnic index. format: int32
	Vnic *int64
    // Timestamp value. format: int64
	Timestamp *int64
    // Tx rate (Kilobits per second - kbps) format: double
	Out *float64
    // Rx rate (Kilobits per second - kbps) format: double
	In *float64
}

func (s CbmStatistic) GetType__() bindings.BindingType {
	return CbmStatisticBindingType()
}

func (s CbmStatistic) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CbmStatistic._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NSX Edge Interface Statistics.
type CbmStatistics struct {
    // Statistics data.
	DataDto *CbmStatsData
    // Start time, end time and interval details.
	MetaDto *MetaDashboardStats
}

func (s CbmStatistics) GetType__() bindings.BindingType {
	return CbmStatisticsBindingType()
}

func (s CbmStatistics) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CbmStatistics._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Statistics data.
type CbmStatsData struct {
	Vnic9 []CbmStatistic
	Vnic8 []CbmStatistic
	Vnic7 []CbmStatistic
	Vnic6 []CbmStatistic
	Vnic5 []CbmStatistic
	Vnic4 []CbmStatistic
	Vnic3 []CbmStatistic
	Vnic2 []CbmStatistic
	Vnic1 []CbmStatistic
	Vnic0 []CbmStatistic
}

func (s CbmStatsData) GetType__() bindings.BindingType {
	return CbmStatsDataBindingType()
}

func (s CbmStatsData) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CbmStatsData._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type CloudProvider struct {
    // Name of the Cloud Provider
	Provider string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const CloudProvider__TYPE_IDENTIFIER = "CloudProvider"

func (s CloudProvider) GetType__() bindings.BindingType {
	return CloudProviderBindingType()
}

func (s CloudProvider) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CloudProvider._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Cluster struct {
	EsxHostList []AwsEsxHost
    // Possible values are: 
    //
    // * Cluster#Cluster_CLUSTER_STATE_DEPLOYING
    // * Cluster#Cluster_CLUSTER_STATE_ADDING_HOSTS
    // * Cluster#Cluster_CLUSTER_STATE_READY
    // * Cluster#Cluster_CLUSTER_STATE_FAILED
	ClusterState *string
    // AWS Key Management Service information associated with this cluster
	AwsKmsInfo *AwsKmsInfo
    // Information of the hosts added to this cluster
	EsxHostInfo *EsxHostInfo
    // Number of cores enabled on ESX hosts added to this cluster format: int32
	HostCpuCoresCount *int64
    // The capacity of this cluster.
	ClusterCapacity *EntityCapacity
	ClusterId string
	ClusterName *string
}
const Cluster_CLUSTER_STATE_DEPLOYING = "DEPLOYING"
const Cluster_CLUSTER_STATE_ADDING_HOSTS = "ADDING_HOSTS"
const Cluster_CLUSTER_STATE_READY = "READY"
const Cluster_CLUSTER_STATE_FAILED = "FAILED"

func (s Cluster) GetType__() bindings.BindingType {
	return ClusterBindingType()
}

func (s Cluster) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Cluster._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ClusterConfig struct {
    // Customize CPU cores on hosts in a cluster. Specify number of cores to be enabled on hosts in a cluster. format: int32
	HostCpuCoresCount *int64
    // The instance type for the esx hosts added to this cluster.
	HostInstanceType *HostInstanceTypes
    // For EBS-backed instances only, the requested storage capacity in GiB. format: int64
	StorageCapacity *int64
	NumHosts int64
}

func (s ClusterConfig) GetType__() bindings.BindingType {
	return ClusterConfigBindingType()
}

func (s ClusterConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ClusterConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ClusterReconfigureParams struct {
    // The final desired storage capacity after reconfiguring the cluster in GiB. format: int64
	StorageCapacity *int64
    // Bias value as obtained from the storage constraints call.
	Bias *string
    // Number of hosts in the cluster after reconfiguring. format: int32
	NumHosts int64
}

func (s ClusterReconfigureParams) GetType__() bindings.BindingType {
	return ClusterReconfigureParamsBindingType()
}

func (s ClusterReconfigureParams) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ClusterReconfigureParams._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ComputeGatewayTemplate struct {
	PublicIp *SddcPublicIp
	PrimaryDns *string
	SecondaryDns *string
	FirewallRules []FirewallRule
	Vpns []Vpn
	LogicalNetworks []LogicalNetwork
	NatRules []NatRule
	L2Vpn *data.StructValue
}

func (s ComputeGatewayTemplate) GetType__() bindings.BindingType {
	return ComputeGatewayTemplateBindingType()
}

func (s ComputeGatewayTemplate) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ComputeGatewayTemplate._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Represents a configuration spec for any sddc provision operation.
type ConfigSpec struct {
    // Indicates after how many days the sddc should expire
	ExpiryInDays *int64
    // Map of region to instance types available in that region
	Availability map[string][]InstanceTypeConfig
	SddcSizes []string
}

func (s ConfigSpec) GetType__() bindings.BindingType {
	return ConfigSpecBindingType()
}

func (s ConfigSpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConfigSpec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ConnectivityAgentValidation struct {
    // Possible values are: 
    //
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_SOURCE_VCENTER
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_SOURCE_SRM
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_SOURCE_VR
    //
    //  source appliance of connectivity test, i.e. VCENTER, SRM, VR.
	Source *string
    // Possible values are: 
    //
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_TYPE_PING
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_TYPE_TRACEROUTE
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_TYPE_DNS
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_TYPE_CONNECTIVITY
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_TYPE_CURL
    //
    //  type of connectivity test, i.e. PING, TRACEROUTE, DNS, CONNECTIVITY, CURL. For CONNECTIVITY and CURL tests only, please specify the ports to be tested against.
	Type_ *string
    // TCP ports ONLY for CONNECTIVITY and CURL tests.
	Ports []string
    // URL path ONLY for CURL tests.
	Path *string
}
const ConnectivityAgentValidation_SOURCE_VCENTER = "VCENTER"
const ConnectivityAgentValidation_SOURCE_SRM = "SRM"
const ConnectivityAgentValidation_SOURCE_VR = "VR"
const ConnectivityAgentValidation_TYPE_PING = "PING"
const ConnectivityAgentValidation_TYPE_TRACEROUTE = "TRACEROUTE"
const ConnectivityAgentValidation_TYPE_DNS = "DNS"
const ConnectivityAgentValidation_TYPE_CONNECTIVITY = "CONNECTIVITY"
const ConnectivityAgentValidation_TYPE_CURL = "CURL"

func (s ConnectivityAgentValidation) GetType__() bindings.BindingType {
	return ConnectivityAgentValidationBindingType()
}

func (s ConnectivityAgentValidation) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectivityAgentValidation._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ConnectivityValidationGroup struct {
    // Possible values are: 
    //
    // * ConnectivityValidationGroup#ConnectivityValidationGroup_ID_HLM
    // * ConnectivityValidationGroup#ConnectivityValidationGroup_ID_DRAAS
    //
    //  test group id, currently, only HLM.
	Id *string
    // Name of the test group.
	Name *string
    // List of sub groups.
	SubGroups []ConnectivityValidationSubGroup
}
const ConnectivityValidationGroup_ID_HLM = "HLM"
const ConnectivityValidationGroup_ID_DRAAS = "DRAAS"

func (s ConnectivityValidationGroup) GetType__() bindings.BindingType {
	return ConnectivityValidationGroupBindingType()
}

func (s ConnectivityValidationGroup) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectivityValidationGroup._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ConnectivityValidationGroups struct {
    // List of groups.
	Groups []ConnectivityValidationGroup
}

func (s ConnectivityValidationGroups) GetType__() bindings.BindingType {
	return ConnectivityValidationGroupsBindingType()
}

func (s ConnectivityValidationGroups) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectivityValidationGroups._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ConnectivityValidationInput struct {
    // Possible values are: 
    //
    // * ConnectivityValidationInput#ConnectivityValidationInput_ID_HOSTNAME
    // * ConnectivityValidationInput#ConnectivityValidationInput_ID_HOST_IP
    // * ConnectivityValidationInput#ConnectivityValidationInput_ID_HOSTNAME_OR_IP
    //
    //  input value type, i.e. HOSTNAME_OR_IP, HOST_IP, HOSTNAME. Accept FQDN or IP address as input value when id = HOSTNAME_OR_IP, accept FQDN ONLY when id = HOSTNAME, accept IP address ONLY when id = HOST_IP.
	Id *string
    // the FQDN or IP address to run the test against, use \\#primary-dns or \\#secondary-dns as the on-prem primary/secondary DNS server IP.
	Value *string
    // (Optional, for UI display only) input value label.
	Label *string
}
const ConnectivityValidationInput_ID_HOSTNAME = "HOSTNAME"
const ConnectivityValidationInput_ID_HOST_IP = "HOST_IP"
const ConnectivityValidationInput_ID_HOSTNAME_OR_IP = "HOSTNAME_OR_IP"

func (s ConnectivityValidationInput) GetType__() bindings.BindingType {
	return ConnectivityValidationInputBindingType()
}

func (s ConnectivityValidationInput) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectivityValidationInput._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ConnectivityValidationSubGroup struct {
    // List of user inputs for the sub group.
	Inputs []ConnectivityValidationInput
    // List of connectivity tests.
	Tests []ConnectivityAgentValidation
    // Name of the sub-group.
	Label *string
    // Help text.
	Help *string
    // Possible values are: 
    //
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_PRIMARY_DNS
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_SECONDARY_DNS
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_ONPREM_VCENTER
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_ONPREM_PSC
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_ACTIVE_DIRECTORY
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_ONPREM_ESX
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_VCENTER
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_PSC
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_SRM
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_VR
    //
    //  subGroup id, i.e. PRIMARY_DNS, SECONDARY_DNS, ONPREM_VCENTER, ONPREM_PSC, ACTIVE_DIRECTORY, ONPREM_ESX, DRAAS_ONPREM_VCENTER, DRAAS_ONPREM_PSC, DRAAS_ONPREM_SRM and DRAAS_ONPREM_VR.
	Id *string
}
const ConnectivityValidationSubGroup_ID_PRIMARY_DNS = "PRIMARY_DNS"
const ConnectivityValidationSubGroup_ID_SECONDARY_DNS = "SECONDARY_DNS"
const ConnectivityValidationSubGroup_ID_ONPREM_VCENTER = "ONPREM_VCENTER"
const ConnectivityValidationSubGroup_ID_ONPREM_PSC = "ONPREM_PSC"
const ConnectivityValidationSubGroup_ID_ACTIVE_DIRECTORY = "ACTIVE_DIRECTORY"
const ConnectivityValidationSubGroup_ID_ONPREM_ESX = "ONPREM_ESX"
const ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_VCENTER = "DRAAS_ONPREM_VCENTER"
const ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_PSC = "DRAAS_ONPREM_PSC"
const ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_SRM = "DRAAS_ONPREM_SRM"
const ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_VR = "DRAAS_ONPREM_VR"

func (s ConnectivityValidationSubGroup) GetType__() bindings.BindingType {
	return ConnectivityValidationSubGroupBindingType()
}

func (s ConnectivityValidationSubGroup) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectivityValidationSubGroup._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// CRL certificate list. Optional.
type CrlCertificates struct {
	CrlCertificate []string
}

func (s CrlCertificates) GetType__() bindings.BindingType {
	return CrlCertificatesBindingType()
}

func (s CrlCertificates) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CrlCertificates._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Indicates a single cross-account ENI and its characteristics.
type CustomerEniInfo struct {
    // Indicates list of secondary IP created for this ENI.
	SecondaryIpAddresses []string
    // Interface ID on customer account.
	EniId *string
    // Indicates primary address of the ENI.
	PrimaryIpAddress *string
}

func (s CustomerEniInfo) GetType__() bindings.BindingType {
	return CustomerEniInfoBindingType()
}

func (s CustomerEniInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CustomerEniInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Dashboard Statistics data.
type DashboardData struct {
    // NSX Edge Firewall Statistics data.
	Firewall *FirewallDashboardStats
    // NSX Edge SSL VPN Statistics data.
	Sslvpn *SslvpnDashboardStats
    // NSX Edge Interface Statistics data.
	Interfaces *InterfacesDashboardStats
    // NSX Edge Load Balancer Statistics data.
	LoadBalancer *LoadBalancerDashboardStats
    // NSX Edge Ipsec Statistics data.
	Ipsec *IpsecDashboardStats
}

func (s DashboardData) GetType__() bindings.BindingType {
	return DashboardDataBindingType()
}

func (s DashboardData) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DashboardData._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type DashboardStat struct {
	Timestamp *int64
	Value *float64
}

func (s DashboardStat) GetType__() bindings.BindingType {
	return DashboardStatBindingType()
}

func (s DashboardStat) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DashboardStat._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Dashboard Statistics data.
type DashboardStatistics struct {
    // Dashboard Statistics data.
	DataDto *DashboardData
    // Start time, end time and interval details.
	MetaDto *MetaDashboardStats
}

func (s DashboardStatistics) GetType__() bindings.BindingType {
	return DashboardStatisticsBindingType()
}

func (s DashboardStatistics) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DashboardStatistics._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type DataPageEdgeSummary struct {
	PagingInfo *PagingInfo
	Data []EdgeSummary
}

func (s DataPageEdgeSummary) GetType__() bindings.BindingType {
	return DataPageEdgeSummaryBindingType()
}

func (s DataPageEdgeSummary) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DataPageEdgeSummary._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type DataPageSddcNetwork struct {
	PagingInfo *PagingInfo
	Data []SddcNetwork
}

func (s DataPageSddcNetwork) GetType__() bindings.BindingType {
	return DataPageSddcNetworkBindingType()
}

func (s DataPageSddcNetwork) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DataPageSddcNetwork._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type DataPermissions struct {
	SavePermission *bool
	PublishPermission *bool
}

func (s DataPermissions) GetType__() bindings.BindingType {
	return DataPermissionsBindingType()
}

func (s DataPermissions) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DataPermissions._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DHCP lease information.
type DhcpLeaseInfo struct {
    // List of DHCP leases.
	HostLeaseInfoDtos []HostLeaseInfo
}

func (s DhcpLeaseInfo) GetType__() bindings.BindingType {
	return DhcpLeaseInfoBindingType()
}

func (s DhcpLeaseInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DhcpLeaseInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DHCP leases information
type DhcpLeases struct {
    // The timestamp of the DHCP lease. format: int64
	TimeStamp *int64
    // DHCP lease information.
	HostLeaseInfosDto *DhcpLeaseInfo
}

func (s DhcpLeases) GetType__() bindings.BindingType {
	return DhcpLeasesBindingType()
}

func (s DhcpLeases) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DhcpLeases._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DNS configuration
type DnsConfig struct {
	FeatureType *string
    // DNS logging setting.
	Logging *Logging
    // Value is true if feature is enabled. Default value is true. Optional.
	Enabled *bool
    // List of DNS views.
	DnsViews *DnsViews
    // List of DNS listeners.
	Listeners *DnsListeners
    // Version number tracking each configuration change. To avoid problems with overwriting changes, always retrieve and modify the latest configuration to include the current version number in your request. If you provide a version number which is not current, the request is rejected. If you omit the version number, the request is accepted but may overwrite any current changes if your change is not in sync with the latest change. format: int64
	Version *int64
	Template *string
    // The cache size of the DNS service. format: int64
	CacheSize *int64
	DnsServers *IpAddresses
}

func (s DnsConfig) GetType__() bindings.BindingType {
	return DnsConfigBindingType()
}

func (s DnsConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DnsConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DNS forwarders.
type DnsForwarders struct {
    // IP addresses of the DNS servers.
	IpAddress []string
}

func (s DnsForwarders) GetType__() bindings.BindingType {
	return DnsForwardersBindingType()
}

func (s DnsForwarders) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DnsForwarders._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type DnsListeners struct {
    // List of IP addresses.
	IpAddress []string
    // Vnic for DNS listener.
	Vnic []string
	Type_ *string
}

func (s DnsListeners) GetType__() bindings.BindingType {
	return DnsListenersBindingType()
}

func (s DnsListeners) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DnsListeners._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DNS response statistics.
type DnsResponseStats struct {
	Total *int64
	FormErr *int64
	NxDomain *int64
	Success *int64
	ServerFail *int64
	Nxrrset *int64
	Others *int64
}

func (s DnsResponseStats) GetType__() bindings.BindingType {
	return DnsResponseStatsBindingType()
}

func (s DnsResponseStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DnsResponseStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DNS statistics.
type DnsStatusAndStats struct {
	TimeStamp *int64
	Requests *Requests
	Responses *DnsResponseStats
	CachedDBRRSet *int64
}

func (s DnsStatusAndStats) GetType__() bindings.BindingType {
	return DnsStatusAndStatsBindingType()
}

func (s DnsStatusAndStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DnsStatusAndStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DNS View
type DnsView struct {
    // Name of the DNS view.
	Name string
    // Rules that match the DNS query to this view. The rule can be ipAddress, or ipSet. Defaults to ipAddress 'any' and 'any' vnic.
	ViewMatch *DnsViewMatch
    // Recursion enabled on DNS view.
	Recursion *bool
    // Identifier for the DNS view.
	ViewId *string
    // DNS forwarders.
	Forwarders *DnsForwarders
    // DNS view is enabled.
	Enabled *bool
}

func (s DnsView) GetType__() bindings.BindingType {
	return DnsViewBindingType()
}

func (s DnsView) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DnsView._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Dns view match
type DnsViewMatch struct {
	Vnic []string
	IpSet []string
	IpAddress []string
}

func (s DnsViewMatch) GetType__() bindings.BindingType {
	return DnsViewMatchBindingType()
}

func (s DnsViewMatch) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DnsViewMatch._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DNS views.
type DnsViews struct {
    // List of DNS views.
	DnsView []DnsView
}

func (s DnsViews) GetType__() bindings.BindingType {
	return DnsViewsBindingType()
}

func (s DnsViews) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DnsViews._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// information for EBS-backed VSAN configuration
type EbsBackedVsanConfig struct {
    // instance type for EBS-backed VSAN configuration
	InstanceType *string
}

func (s EbsBackedVsanConfig) GetType__() bindings.BindingType {
	return EbsBackedVsanConfigBindingType()
}

func (s EbsBackedVsanConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EbsBackedVsanConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Job status information for the configuration change carried out on NSX Edge.
type EdgeJob struct {
    // Job status.
	Status *string
    // NSX Edge ID.
	EdgeId *string
    // Module information.
	Module *string
    // Job ID.
	JobId *string
    // Error code identifying the failure of the configuration change.
	ErrorCode *string
    // Job result information.
	Result []Result
    // Job start time. format: date-time
	StartTime *time.Time
    // Job message.
	Message *string
    // Job end time. format: date-time
	EndTime *time.Time
}

func (s EdgeJob) GetType__() bindings.BindingType {
	return EdgeJobBindingType()
}

func (s EdgeJob) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EdgeJob._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NSX Edge Appliance status.
type EdgeStatus struct {
    // Value is true if pre rules publish is enabled.
	PreRulesExists *bool
    // Individual feature status.
	FeatureStatuses []FeatureStatus
    // Timestamp value at which the NSX Edge healthcheck was done. format: int64
	Timestamp *int64
    // Status of the latest configuration change for the NSX Edge. Values are APPLIED or PERSISTED (not published to NSX Edge Appliance yet).
	PublishStatus *string
    // Value of the last published pre rules generation number. format: int64
	LastPublishedPreRulesGenerationNumber *int64
    // Version number of the current configuration. format: int64
	Version *int64
    // Detailed status of each of the deployed NSX Edge appliances.
	EdgeVmStatus []EdgeVmStatus
    // Index of the active NSX Edge appliance. Values are 0 and 1. format: int32
	ActiveVseHaIndex *int64
    // System status of the active NSX Edge appliance.
	SystemStatus *string
    // Index of the vnic consumed for NSX Edge HA. format: int32
	HaVnicInUse *int64
    // NSX Edge appliance health status identified by GREY (unknown status), GREEN (health checks are successful), YELLOW (intermittent health check failure), RED (none of the appliances are in serving state). If health check fails for 5 consecutive times for all appliance (2 for HA else 1) then status will turn from YELLOW to RED.
	EdgeStatus *string
}

func (s EdgeStatus) GetType__() bindings.BindingType {
	return EdgeStatusBindingType()
}

func (s EdgeStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EdgeStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NSX Edge summary. Read only.
type EdgeSummary struct {
    // List of Features and their capability details based on Edge appliance form factor.
	FeatureCapabilities *FeatureCapabilities
    // NSX Edge type, whether 'gatewayServices' or 'distributedRouter'.
	EdgeType *string
    // Backing type scope (DistributedVirtualSwitch - VLAN, TransportZone -VXLAN) and its ID for the Distributed Logical Router.
	LogicalRouterScopes *LogicalRouterScopes
    // Job information for the most recent configuration change carried out on the NSX Edge.
	RecentJobInfo *EdgeJob
	HypervisorAssist *bool
    // ID generated by NSX Manager for Distributed Logical Router only. format: int64
	EdgeAssistId *int64
    // NSX Edge appliance health status identified by GREY (unknown status), GREEN (health checks are successful), YELLOW (intermittent health check failure), RED (none of the appliances are in serving state). If health check fails for 5 consecutive times for all appliance (2 for HA else 1) then status will turn from YELLOW to RED.
	EdgeStatus *string
    // Name derived by NSX Manager only for Distributed Logical Router.
	EdgeAssistInstanceName *string
	ObjectId *string
	NodeId *string
    // NSX Edge ID.
	Id *string
    // Datacenter name where the NSX Edge is deployed.
	DatacenterName *string
    // Deployment state of the NSX Edge appliance. Values are 'deployed' when VMs have been deployed, 'undeployed' when no VMs are deployed and 'active' when Edge type is Distributed Logical Router and has no appliance deployed but is serving data path.
	State *string
	ClientHandle *string
	Scope *ScopeInfo
	Type_ *ObjectType
	Revision *int64
	VsmUuid *string
	Description *string
	ExtendedAttributes []ExtendedAttribute
    // Value is true if local egress is enabled for UDLR traffic. Applicable only for Universal Distributed Logical Router.
	LocalEgressEnabled *bool
	UniversalRevision *int64
	AllowedActions []string
	ObjectTypeName *string
    // Value is true if NSX Edge upgrade is available.
	IsUpgradeAvailable *bool
	IsUniversal *bool
	Name *string
    // Distributed Logical Router UUID provided by the NSX Controller.
	LrouterUuid *string
    // NSX Edge appliance summary.
	AppliancesSummary *AppliancesSummary
    // REST API version applicable for the NSX Edge.
	ApiVersion *string
    // Tenant ID for the NSX Edge.
	TenantId *string
    // vCenter MOID of the datacenter where the NSX Edge is deployed.
	DatacenterMoid *string
    // Number of connected vnics that are configured on the NSX Edge. format: int32
	NumberOfConnectedVnics *int64
}

func (s EdgeSummary) GetType__() bindings.BindingType {
	return EdgeSummaryBindingType()
}

func (s EdgeSummary) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EdgeSummary._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Status of each of the deployed NSX Edge appliances.
type EdgeVmStatus struct {
    // High Availability index of the appliance. Values are 0 and 1. format: int32
	Index *int64
    // High Availability state of the appliance. Values are active and standby.
	HaState *string
    // Name of the NSX Edge appliance.
	Name *string
    // vCenter MOID of the NSX Edge appliance.
	Id *string
    // NSX Edge appliance health status identified by GREY (unknown status), GREEN (health checks are successful), YELLOW (intermittent health check failure), RED (appliance not in serving state).
	EdgeVMStatus *string
    // Value of the last published pre rules generation number. format: int64
	PreRulesGenerationNumber *int64
}

func (s EdgeVmStatus) GetType__() bindings.BindingType {
	return EdgeVmStatusBindingType()
}

func (s EdgeVmStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EdgeVmStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Address group configuration of the NSX Edge vnic. An interface can have one primary and multiple secondary IP addresses.
type EdgeVnicAddressGroup struct {
    // Subnet prefix length of the primary IP address.
	SubnetPrefixLength *string
    // Secondary IP addresses of the NSX Edge vnic address group. Optional.
	SecondaryAddresses *SecondaryAddresses
    // Primary IP address of the vnic interface. Required.
	PrimaryAddress *string
	SubnetMask *string
}

func (s EdgeVnicAddressGroup) GetType__() bindings.BindingType {
	return EdgeVnicAddressGroupBindingType()
}

func (s EdgeVnicAddressGroup) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EdgeVnicAddressGroup._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NSX Edge vnic address group configuration details.
type EdgeVnicAddressGroups struct {
    // Address group configuration of the NSX Edge vnic. Vnic can be configured to have more than one address group/subnets.
	AddressGroups []EdgeVnicAddressGroup
}

func (s EdgeVnicAddressGroups) GetType__() bindings.BindingType {
	return EdgeVnicAddressGroupsBindingType()
}

func (s EdgeVnicAddressGroups) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EdgeVnicAddressGroups._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Information of the x-eni created.
type EniInfo struct {
    // Subnet it belongs to.
	SubnetId *string
    // Interface ID.
	Id *string
    // Security Group of Eni.
	SecurityGroupId *string
    // Private ip of eni.
	PrivateIp *string
    // Mac address of nic.
	MacAddress *string
}

func (s EniInfo) GetType__() bindings.BindingType {
	return EniInfoBindingType()
}

func (s EniInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EniInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Decribes the capacity of a given entity.
type EntityCapacity struct {
	StorageCapacityGib *int64
	MemoryCapacityGib *int64
	TotalNumberOfCores *int64
	NumberOfSsds *int64
	CpuCapacityGhz *float64
	NumberOfSockets *int64
}

func (s EntityCapacity) GetType__() bindings.BindingType {
	return EntityCapacityBindingType()
}

func (s EntityCapacity) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EntityCapacity._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ErrorResponse struct {
    // HTTP status code
	Status int64
    // Originating request URI
	Path string
    // If true, client should retry operation
	Retryable bool
    // unique error code
	ErrorCode string
    // localized error messages
	ErrorMessages []string
}

func (s ErrorResponse) GetType__() bindings.BindingType {
	return ErrorResponseBindingType()
}

func (s ErrorResponse) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ErrorResponse._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type EsxConfig struct {
    // Availability zone where the hosts should be provisioned. (Can be specified only for privileged host operations).
	AvailabilityZone *string
	Esxs []string
    // An optional cluster id if the esxs operation has to be on a specific cluster.
	ClusterId *string
	NumHosts int64
}

func (s EsxConfig) GetType__() bindings.BindingType {
	return EsxConfigBindingType()
}

func (s EsxConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EsxConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type EsxHost struct {
	Name *string
    // Availability zone where the host is provisioned.
	AvailabilityZone *string
	EsxId *string
	Hostname *string
	Provider string
    // Backing cloud provider instance type for host.
	InstanceType *string
	MacAddress *string
	CustomProperties map[string]string
    // Possible values are: 
    //
    // * EsxHost#EsxHost_ESX_STATE_DEPLOYING
    // * EsxHost#EsxHost_ESX_STATE_INITIALIZING
    // * EsxHost#EsxHost_ESX_STATE_PROVISIONED
    // * EsxHost#EsxHost_ESX_STATE_READY
    // * EsxHost#EsxHost_ESX_STATE_DELETING
    // * EsxHost#EsxHost_ESX_STATE_DELETED
    // * EsxHost#EsxHost_ESX_STATE_FAILED
    // * EsxHost#EsxHost_ESX_STATE_ADDING_TO_VCENTER
    // * EsxHost#EsxHost_ESX_STATE_DELETING_FROM_VCENTER
    // * EsxHost#EsxHost_ESX_STATE_PENDING_CLOUD_DELETION
	EsxState *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const EsxHost__TYPE_IDENTIFIER = "EsxHost"
const EsxHost_ESX_STATE_DEPLOYING = "DEPLOYING"
const EsxHost_ESX_STATE_INITIALIZING = "INITIALIZING"
const EsxHost_ESX_STATE_PROVISIONED = "PROVISIONED"
const EsxHost_ESX_STATE_READY = "READY"
const EsxHost_ESX_STATE_DELETING = "DELETING"
const EsxHost_ESX_STATE_DELETED = "DELETED"
const EsxHost_ESX_STATE_FAILED = "FAILED"
const EsxHost_ESX_STATE_ADDING_TO_VCENTER = "ADDING_TO_VCENTER"
const EsxHost_ESX_STATE_DELETING_FROM_VCENTER = "DELETING_FROM_VCENTER"
const EsxHost_ESX_STATE_PENDING_CLOUD_DELETION = "PENDING_CLOUD_DELETION"

func (s EsxHost) GetType__() bindings.BindingType {
	return EsxHostBindingType()
}

func (s EsxHost) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EsxHost._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type EsxHostInfo struct {
    // Backing cloud provider instance type for cluster.
	InstanceType *string
}

func (s EsxHostInfo) GetType__() bindings.BindingType {
	return EsxHostInfoBindingType()
}

func (s EsxHostInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EsxHostInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ExtendedAttribute struct {
	Name *string
	Value *string
}

func (s ExtendedAttribute) GetType__() bindings.BindingType {
	return ExtendedAttributeBindingType()
}

func (s ExtendedAttribute) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ExtendedAttribute._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// List of features and their capability details based on NSX Edge appliance form factor.
type FeatureCapabilities struct {
    // Time stamp value at which the feature capabilities were retrieved. format: int64
	Timestamp *int64
    // List of feature capability information.
	FeatureCapabilities []FeatureCapability
}

func (s FeatureCapabilities) GetType__() bindings.BindingType {
	return FeatureCapabilitiesBindingType()
}

func (s FeatureCapabilities) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FeatureCapabilities._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Feature capability information.
type FeatureCapability struct {
    // List of key value pairs describing the feature configuration limits.
	ConfigurationLimits []KeyValueAttributes
    // Value is true if feature is supported on NSX Edge.
	IsSupported *bool
    // Name of the feature or service.
	Service *string
    // Licence and access control information for the feature.
	Permission *LicenceAclPermissions
}

func (s FeatureCapability) GetType__() bindings.BindingType {
	return FeatureCapabilityBindingType()
}

func (s FeatureCapability) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FeatureCapability._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Individual feature status.
type FeatureStatus struct {
    // Status of the feature or service.
	Status *string
    // Value is true if feature is configured.
	Configured *bool
    // Server status of the feature or service. Values are up and down.
	ServerStatus *string
    // Publish status of the feature, whether APPLIED or PERSISTED.
	PublishStatus *string
    // Name of the feature or service.
	Service *string
}

func (s FeatureStatus) GetType__() bindings.BindingType {
	return FeatureStatusBindingType()
}

func (s FeatureStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FeatureStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Firewall Configuration
type FirewallConfig struct {
    // Ordered list of firewall rules.
	FirewallRules *FirewallRules
	FeatureType *string
    // Version number tracking each configuration change. To avoid problems with overwriting changes, always retrieve and modify the latest configuration to include the current version number in your request. If you provide a version number which is not current, the request is rejected. If you omit the version number, the request is accepted but may overwrite any current changes if your change is not in sync with the latest change. format: int64
	Version *int64
	Template *string
    // Global configuration applicable to all rules.
	GlobalConfig *FirewallGlobalConfig
    // Value is true if feature is enabled. Default value is true. Optional.
	Enabled *bool
    // Default Policy.
	DefaultPolicy *FirewallDefaultPolicy
}

func (s FirewallConfig) GetType__() bindings.BindingType {
	return FirewallConfigBindingType()
}

func (s FirewallConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Dashboard Statistics data for Firewall.
type FirewallDashboardStats struct {
    // Number of NSX Edge firewall connections and rules.
	Connections []DashboardStat
}

func (s FirewallDashboardStats) GetType__() bindings.BindingType {
	return FirewallDashboardStatsBindingType()
}

func (s FirewallDashboardStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallDashboardStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Firewall default policy. Default is deny.
type FirewallDefaultPolicy struct {
    // Action. Default is deny. Supported values accept, deny
	Action *string
    // Enable logging for the rule.
	LoggingEnabled *bool
}

func (s FirewallDefaultPolicy) GetType__() bindings.BindingType {
	return FirewallDefaultPolicyBindingType()
}

func (s FirewallDefaultPolicy) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallDefaultPolicy._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Global configuration applicable to all rules.
type FirewallGlobalConfig struct {
    // Allow TCP out of window packets.
	TcpAllowOutOfWindowPackets *bool
    // UDP timeout close. format: int32
	UdpTimeout *int64
    // IP generic timeout. format: int32
	IpGenericTimeout *int64
    // Pick TCP ongoing connections.
	TcpPickOngoingConnections *bool
    // TCP timeout open. format: int32
	TcpTimeoutOpen *int64
    // TCP timeout close. format: int32
	TcpTimeoutClose *int64
    // ICMP6 timeout. format: int32
	Icmp6Timeout *int64
    // Drop icmp replays.
	DropIcmpReplays *bool
    // Log icmp errors.
	LogIcmpErrors *bool
    // Send TCP reset for closed NSX Edge ports.
	TcpSendResetForClosedVsePorts *bool
    // Drop invalid traffic.
	DropInvalidTraffic *bool
    // Protect against SYN flood attacks by detecting bogus TCP connections and terminating them without consuming firewall state tracking resources. Default : false
	EnableSynFloodProtection *bool
    // ICMP timeout. format: int32
	IcmpTimeout *int64
    // TCP timeout established. format: int32
	TcpTimeoutEstablished *int64
    // Log invalid traffic.
	LogInvalidTraffic *bool
}

func (s FirewallGlobalConfig) GetType__() bindings.BindingType {
	return FirewallGlobalConfigBindingType()
}

func (s FirewallGlobalConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallGlobalConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type FirewallRule struct {
    // Possible values are: 
    //
    // * FirewallRule#FirewallRule_RULE_TYPE_USER
    // * FirewallRule#FirewallRule_RULE_TYPE_DEFAULT
	RuleType *string
	ApplicationIds []string
	Name *string
    // Deprecated, left for backwards compatibility. Remove once UI stops using it.
	RuleInterface *string
    // Optional. Possible formats are IP, IP1-IPn, CIDR or comma separated list of those entries. If not specified, defaults to 'any'.
	Destination *string
	Id *string
	DestinationScope *FirewallRuleScope
    // Optional. Possible formats are IP, IP1-IPn, CIDR or comma separated list of those entries. If not specified, defaults to 'any'.
	Source *string
	SourceScope *FirewallRuleScope
    // list of protocols and ports for this firewall rule
	Services []FirewallService
    // Possible values are: 
    //
    // * FirewallRule#FirewallRule_ACTION_ALLOW
    // * FirewallRule#FirewallRule_ACTION_DENY
	Action *string
    // current revision of the list of firewall rules, used to protect against concurrent modification (first writer wins) format: int32
	Revision *int64
}
const FirewallRule_RULE_TYPE_USER = "USER"
const FirewallRule_RULE_TYPE_DEFAULT = "DEFAULT"
const FirewallRule_ACTION_ALLOW = "ALLOW"
const FirewallRule_ACTION_DENY = "DENY"

func (s FirewallRule) GetType__() bindings.BindingType {
	return FirewallRuleBindingType()
}

func (s FirewallRule) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallRule._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Optional for FirewallRule. If not specified, defaults to 'any'.
type FirewallRuleScope struct {
	GroupingObjectIds []string
    // Possible values are: 
    //
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VSE
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_INTERNAL
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_EXTERNAL
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_0
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_1
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_2
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_3
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_4
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_5
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_6
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_7
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_8
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_9
	VnicGroupIds []string
}
const FirewallRuleScope_VNIC_GROUP_IDS_VSE = "vse"
const FirewallRuleScope_VNIC_GROUP_IDS_INTERNAL = "internal"
const FirewallRuleScope_VNIC_GROUP_IDS_EXTERNAL = "external"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_0 = "vnic-index-0"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_1 = "vnic-index-1"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_2 = "vnic-index-2"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_3 = "vnic-index-3"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_4 = "vnic-index-4"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_5 = "vnic-index-5"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_6 = "vnic-index-6"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_7 = "vnic-index-7"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_8 = "vnic-index-8"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_9 = "vnic-index-9"

func (s FirewallRuleScope) GetType__() bindings.BindingType {
	return FirewallRuleScopeBindingType()
}

func (s FirewallRuleScope) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallRuleScope._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Statistics for firewall rule
type FirewallRuleStats struct {
    // Timestamp of statistics collection. format: int64
	Timestamp *int64
    // Connection count. format: int64
	ConnectionCount *int64
    // Byte count. format: int64
	ByteCount *int64
    // Packet count. format: int64
	PacketCount *int64
}

func (s FirewallRuleStats) GetType__() bindings.BindingType {
	return FirewallRuleStatsBindingType()
}

func (s FirewallRuleStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallRuleStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Ordered list of firewall rules.
type FirewallRules struct {
    // Ordered list of firewall rules.
	FirewallRules []Nsxfirewallrule
}

func (s FirewallRules) GetType__() bindings.BindingType {
	return FirewallRulesBindingType()
}

func (s FirewallRules) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallRules._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type FirewallService struct {
    // protocol name, such as 'tcp', 'udp' etc.
	Protocol *string
    // a list of port numbers and port ranges, such as {80, 91-95, 99}. If not specified, defaults to 'any'.
	Ports []string
}

func (s FirewallService) GetType__() bindings.BindingType {
	return FirewallServiceBindingType()
}

func (s FirewallService) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallService._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Describes common properties for MGW and CGW configuration templates
type GatewayTemplate struct {
	PublicIp *SddcPublicIp
	PrimaryDns *string
	SecondaryDns *string
	FirewallRules []FirewallRule
	Vpns []Vpn
}

func (s GatewayTemplate) GetType__() bindings.BindingType {
	return GatewayTemplateBindingType()
}

func (s GatewayTemplate) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for GatewayTemplate._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// the GlcmBundle used for deploying the sddc
type GlcmBundle struct {
    // the glcmbundle's s3 bucket
	S3Bucket *string
    // the glcmbundle's id
	Id *string
}

func (s GlcmBundle) GetType__() bindings.BindingType {
	return GlcmBundleBindingType()
}

func (s GlcmBundle) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for GlcmBundle._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DHCP lease information.
type HostLeaseInfo struct {
    // MAC address of the client.
	MacAddress *string
    // End time of the lease.
	Ends *string
    // Time stamp of when IP address was marked as abandoned.
	Abandoned *string
    // Client Last Transaction Time of the lease info.
	Cltt *string
    // Name of the client.
	ClientHostname *string
    // Start time of the lease.
	Starts *string
    // Lease's binding state.
	BindingState *string
    // The hardware type on which the lease will be used.
	HardwareType *string
    // Time Sent From Partner of the lease info.
	Tsfp *string
    // Uid to identify the DHCP lease.
	Uid *string
    // Indicates what state the lease will move to when the current state expires.
	NextBindingState *string
    // IP address of the client.
	IpAddress *string
    // Time Sent To Partner of the lease info.
	Tstp *string
}

func (s HostLeaseInfo) GetType__() bindings.BindingType {
	return HostLeaseInfoBindingType()
}

func (s HostLeaseInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for HostLeaseInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Represents a structure for instance type config
type InstanceTypeConfig struct {
    // Instance type name.
	InstanceType *string
    // Array of number of hosts allowed for this operation. Range of hosts user can select for sddc provision
	Hosts []int64
    // Display name of instance_type.
	DisplayName *string
    // The capacity of the given instance type.
	EntityCapacity *EntityCapacity
}

func (s InstanceTypeConfig) GetType__() bindings.BindingType {
	return InstanceTypeConfigBindingType()
}

func (s InstanceTypeConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for InstanceTypeConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type InteractionPermissions struct {
	ManagePermission *bool
	ViewPermission *bool
}

func (s InteractionPermissions) GetType__() bindings.BindingType {
	return InteractionPermissionsBindingType()
}

func (s InteractionPermissions) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for InteractionPermissions._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Dashboard Statistics data for Interfaces.
type InterfacesDashboardStats struct {
	Vnic7InPkt []DashboardStat
	Vnic0InByte []DashboardStat
	Vnic8OutPkt []DashboardStat
	Vnic5InByte []DashboardStat
	Vnic2InPkt []DashboardStat
	Vnic3InPkt []DashboardStat
	Vnic6OutByte []DashboardStat
	Vnic3InByte []DashboardStat
	Vnic8InPkt []DashboardStat
	Vnic1InByte []DashboardStat
	Vnic1OutPkt []DashboardStat
	Vnic5OutByte []DashboardStat
	Vnic0OutPkt []DashboardStat
	Vnic0OutByte []DashboardStat
	Vnic6OutPkt []DashboardStat
	Vnic3OutByte []DashboardStat
	Vnic7InByte []DashboardStat
	Vnic1OutByte []DashboardStat
	Vnic9OutPkt []DashboardStat
	Vnic9InPkt []DashboardStat
	Vnic4InByte []DashboardStat
	Vnic5OutPkt []DashboardStat
	Vnic2OutPkt []DashboardStat
	Vnic2InByte []DashboardStat
	Vnic5InPkt []DashboardStat
	Vnic7OutPkt []DashboardStat
	Vnic3OutPkt []DashboardStat
	Vnic4OutPkt []DashboardStat
	Vnic4OutByte []DashboardStat
	Vnic1InPkt []DashboardStat
	Vnic2OutByte []DashboardStat
	Vnic6InByte []DashboardStat
	Vnic0InPkt []DashboardStat
	Vnic9InByte []DashboardStat
	Vnic7OutByte []DashboardStat
	Vnic4InPkt []DashboardStat
	Vnic9OutByte []DashboardStat
	Vnic8OutByte []DashboardStat
	Vnic8InByte []DashboardStat
	Vnic6InPkt []DashboardStat
}

func (s InterfacesDashboardStats) GetType__() bindings.BindingType {
	return InterfacesDashboardStatsBindingType()
}

func (s InterfacesDashboardStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for InterfacesDashboardStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// IP address
type IpAddresses struct {
    // List of IP addresses.
	IpAddress []string
}

func (s IpAddresses) GetType__() bindings.BindingType {
	return IpAddressesBindingType()
}

func (s IpAddresses) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for IpAddresses._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NSX Edge IPsec configuration details.
type Ipsec struct {
	FeatureType *string
    // Configure logging for the feature on NSX Edge appliance. Logging is disabled by default. Optional.
	Logging *Logging
    // IPsec Global configuration details.
	Global *IpsecGlobalConfig
    // Value is true if feature is enabled. Default value is true. Optional.
	Enabled *bool
    // IPsec Site configuration details.
	Sites *IpsecSites
    // Enable/disable event generation on NSX Edge appliance for IPsec.
	DisableEvent *bool
    // Version number tracking each configuration change. To avoid problems with overwriting changes, always retrieve and modify the latest configuration to include the current version number in your request. If you provide a version number which is not current, the request is rejected. If you omit the version number, the request is accepted but may overwrite any current changes if your change is not in sync with the latest change. format: int64
	Version *int64
	Template *string
}

func (s Ipsec) GetType__() bindings.BindingType {
	return IpsecBindingType()
}

func (s Ipsec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Ipsec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Dashboard Statistics data for Ipsec.
type IpsecDashboardStats struct {
    // Tx transmitted bytes.
	IpsecBytesOut []DashboardStat
    // Rx received bytes.
	IpsecBytesIn []DashboardStat
    // Number of Ipsec tunnels.
	IpsecTunnels []DashboardStat
}

func (s IpsecDashboardStats) GetType__() bindings.BindingType {
	return IpsecDashboardStatsBindingType()
}

func (s IpsecDashboardStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for IpsecDashboardStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// IPsec Global configuration details.
type IpsecGlobalConfig struct {
    // IPsec Global Pre Shared Key. Maximum characters is 128. Required when peerIp is configured as 'any' in NSX Edge IPsec Site configuration.
	Psk *string
    // CA certificate list. Optional.
	CaCertificates *CaCertificates
    // Certificate name or identifier. Required when x.509 is selected as the authentication mode.
	ServiceCertificate *string
    // CRL certificate list. Optional.
	CrlCertificates *CrlCertificates
	Extension *string
}

func (s IpsecGlobalConfig) GetType__() bindings.BindingType {
	return IpsecGlobalConfigBindingType()
}

func (s IpsecGlobalConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for IpsecGlobalConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NSX Edge IPsec Site configuration details.
type IpsecSite struct {
    // Pre Shared Key for the IPsec Site. Required if Site peerIp is not 'any'. Global PSK is used when Authentication mode is PSK and Site peerIp is 'any'.
	Psk *string
    // Local ID of the IPsec Site. Defaults to the local IP.
	LocalId *string
    // Enable/disable Perfect Forward Secrecy. Default is true.
	EnablePfs *bool
    // Authentication mode for the IPsec Site. Valid values are psk and x.509, with psk as default.
	AuthenticationMode *string
    // Peer subnets for which IPsec VPN is configured.
	PeerSubnets *Subnets
    // Diffie-Hellman algorithm group. Defaults to DH14 for FIPS enabled NSX Edge. DH2 and DH5 are not supported when FIPS is enabled on NSX Edge. Valid values are DH2, DH5, DH14, DH15, DH16.
	DhGroup *string
    // ID of the IPsec Site configuration provided by NSX Manager.
	SiteId *string
    // Description of the IPsec Site.
	Description *string
    // IP (IPv4) address or FQDN of the Peer. Can also be specified as 'any'. Required.
	PeerIp *string
    // Name of the IPsec Site.
	Name *string
	Certificate *string
    // Local IP of the IPsec Site. Should be one of the IP addresses configured on the uplink interfaces of the NSX Edge. Required.
	LocalIp *string
    // IPsec encryption algorithm with default as aes256. Valid values are 'aes', 'aes256', '3des', 'aes-gcm'.
	EncryptionAlgorithm *string
    // Enable/disable IPsec Site.
	Enabled *bool
    // MTU for the IPsec site. Defaults to the mtu of the NSX Edge vnic specified by the localIp. Optional. format: int32
	Mtu *int64
	Extension *string
    // Peer ID. Should be unique for all IPsec Site's configured for an NSX Edge.
	PeerId *string
    // Local subnets for which IPsec VPN is configured.
	LocalSubnets *Subnets
}

func (s IpsecSite) GetType__() bindings.BindingType {
	return IpsecSiteBindingType()
}

func (s IpsecSite) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for IpsecSite._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type IpsecSiteIKEStatus struct {
	ChannelStatus *string
	ChannelState *string
	PeerIpAddress *string
	LocalIpAddress *string
	PeerSubnets []string
	PeerId *string
	LastInformationalMessage *string
	LocalSubnets []string
}

func (s IpsecSiteIKEStatus) GetType__() bindings.BindingType {
	return IpsecSiteIKEStatusBindingType()
}

func (s IpsecSiteIKEStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for IpsecSiteIKEStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type IpsecSiteStats struct {
	RxBytesOnSite *int64
	TunnelStats []IpsecTunnelStats
	IkeStatus *IpsecSiteIKEStatus
	SiteStatus *string
	TxBytesFromSite *int64
}

func (s IpsecSiteStats) GetType__() bindings.BindingType {
	return IpsecSiteStatsBindingType()
}

func (s IpsecSiteStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for IpsecSiteStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// List of IPsec sites for NSX Edge.
type IpsecSites struct {
	Sites []IpsecSite
}

func (s IpsecSites) GetType__() bindings.BindingType {
	return IpsecSitesBindingType()
}

func (s IpsecSites) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for IpsecSites._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type IpsecStatusAndStats struct {
	TimeStamp *int64
	ServerStatus *string
	SiteStatistics []IpsecSiteStats
}

func (s IpsecStatusAndStats) GetType__() bindings.BindingType {
	return IpsecStatusAndStatsBindingType()
}

func (s IpsecStatusAndStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for IpsecStatusAndStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type IpsecTunnelStats struct {
	TunnelStatus *string
	PeerSPI *string
	RxBytesOnLocalSubnet *int64
	EstablishedDate *string
	PeerSubnet *string
	AuthenticationAlgorithm *string
	TunnelState *string
	TxBytesFromLocalSubnet *int64
	LastInformationalMessage *string
	LocalSPI *string
	EncryptionAlgorithm *string
	LocalSubnet *string
}

func (s IpsecTunnelStats) GetType__() bindings.BindingType {
	return IpsecTunnelStatsBindingType()
}

func (s IpsecTunnelStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for IpsecTunnelStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Key value pair describing the feature configuration limit.
type KeyValueAttributes struct {
    // Value corresponding to the key of the configuration limit parameter.
	Value *string
    // Key name of the configuration limit parameter.
	Key *string
}

func (s KeyValueAttributes) GetType__() bindings.BindingType {
	return KeyValueAttributesBindingType()
}

func (s KeyValueAttributes) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for KeyValueAttributes._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type KmsVpcEndpoint struct {
    // The identifier of the VPC endpoint created to AWS Key Management Service
	VpcEndpointId *string
	NetworkInterfaceIds []string
}

func (s KmsVpcEndpoint) GetType__() bindings.BindingType {
	return KmsVpcEndpointBindingType()
}

func (s KmsVpcEndpoint) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for KmsVpcEndpoint._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Layer 2 extension.
type L2Extension struct {
    // Identifier for layer 2 extension tunnel. Valid range: 1-4093. format: int32
	TunnelId int64
}

func (s L2Extension) GetType__() bindings.BindingType {
	return L2ExtensionBindingType()
}

func (s L2Extension) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for L2Extension._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type L2Vpn struct {
    // Enable (true) or disable (false) L2 VPN.
	Enabled *bool
    // Array of L2 vpn site config.
	Sites []Site
    // Public uplink ip address. IP of external interface on which L2VPN service listens to.
	ListenerIp *string
}

func (s L2Vpn) GetType__() bindings.BindingType {
	return L2VpnBindingType()
}

func (s L2Vpn) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for L2Vpn._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// L2 VPN status and statistics of a single L2 VPN site.
type L2vpnStats struct {
    // Status of the tunnel (UP/DOWN).
	TunnelStatus *string
    // Tunnel established date. format: int64
	EstablishedDate *int64
    // User defined name of the site.
	Name *string
    // Number of received packets dropped.
	DroppedRxPackets *int64
    // Cipher used in encryption.
	EncryptionAlgorithm *string
    // Reason for the tunnel down.
	FailureMessage *string
    // Number of bytes transferred from local subnet.
	TxBytesFromLocalSubnet *int64
    // Number of bytes received on the local subnet.
	RxBytesOnLocalSubnet *int64
    // Number of transferred packets dropped.
	DroppedTxPackets *int64
    // Time stamp of the statistics collection. format: int64
	LastUpdatedTime *int64
}

func (s L2vpnStats) GetType__() bindings.BindingType {
	return L2vpnStatsBindingType()
}

func (s L2vpnStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for L2vpnStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// L2 VPN status and statistics.
type L2vpnStatusAndStats struct {
    // Time stamp of statistics collection. format: int64
	TimeStamp *int64
	ServerStatus *string
    // List of statistics for each Site.
	SiteStats []L2vpnStats
}

func (s L2vpnStatusAndStats) GetType__() bindings.BindingType {
	return L2vpnStatusAndStatsBindingType()
}

func (s L2vpnStatusAndStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for L2vpnStatusAndStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Licence and access control information for the feature.
type LicenceAclPermissions struct {
    // Data access control information for the feature.
	DataPermission *DataPermissions
    // Value is true if feature is licenced.
	IsLicensed *bool
    // Access control information for the feature.
	AccessPermission *InteractionPermissions
}

func (s LicenceAclPermissions) GetType__() bindings.BindingType {
	return LicenceAclPermissionsBindingType()
}

func (s LicenceAclPermissions) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for LicenceAclPermissions._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Dashboard Statistics data for Load Balancer.
type LoadBalancerDashboardStats struct {
    // Number of bytes in.
	LbBpsIn []DashboardStat
    // Number of HTTP requests received by Load Balancer.
	LbHttpReqs []DashboardStat
    // Number of bytes out.
	LbBpsOut []DashboardStat
    // Number of Load Balancer sessions.
	LbSessions []DashboardStat
}

func (s LoadBalancerDashboardStats) GetType__() bindings.BindingType {
	return LoadBalancerDashboardStatsBindingType()
}

func (s LoadBalancerDashboardStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for LoadBalancerDashboardStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// logging.
type Logging struct {
    // Log level. Valid values: emergency, alert, critical, error, warning, notice, info, debug.
	LogLevel *string
    // Logging enabled.
	Enable *bool
}

func (s Logging) GetType__() bindings.BindingType {
	return LoggingBindingType()
}

func (s Logging) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Logging._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type LogicalNetwork struct {
    // the subnet cidr
	SubnetCidr *string
    // name of the network
	Name *string
    // gateway ip of the logical network
	GatewayIp *string
    // if 'true' - enabled; if 'false' - disabled
	DhcpEnabled *string
    // ip range within the subnet mask, range delimiter is '-' (example 10.118.10.130-10.118.10.140)
	DhcpIpRange *string
    // tunnel id of extended network format: int32
	TunnelId *int64
	Id *string
    // Possible values are: 
    //
    // * LogicalNetwork#LogicalNetwork_NETWORK_TYPE_HOSTED
    // * LogicalNetwork#LogicalNetwork_NETWORK_TYPE_ROUTED
    // * LogicalNetwork#LogicalNetwork_NETWORK_TYPE_EXTENDED
	NetworkType *string
}
const LogicalNetwork_NETWORK_TYPE_HOSTED = "HOSTED"
const LogicalNetwork_NETWORK_TYPE_ROUTED = "ROUTED"
const LogicalNetwork_NETWORK_TYPE_EXTENDED = "EXTENDED"

func (s LogicalNetwork) GetType__() bindings.BindingType {
	return LogicalNetworkBindingType()
}

func (s LogicalNetwork) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for LogicalNetwork._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type LogicalRouterScope struct {
	Type_ *string
	Id *string
}

func (s LogicalRouterScope) GetType__() bindings.BindingType {
	return LogicalRouterScopeBindingType()
}

func (s LogicalRouterScope) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for LogicalRouterScope._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type LogicalRouterScopes struct {
	LogicalRouterScope []LogicalRouterScope
}

func (s LogicalRouterScopes) GetType__() bindings.BindingType {
	return LogicalRouterScopesBindingType()
}

func (s LogicalRouterScopes) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for LogicalRouterScopes._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type MacAddress struct {
	EdgeVmHaIndex *int64
	Value *string
}

func (s MacAddress) GetType__() bindings.BindingType {
	return MacAddressBindingType()
}

func (s MacAddress) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MacAddress._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type MaintenanceWindow struct {
    // Possible values are: 
    //
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_SUNDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_MONDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_TUESDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_WEDNESDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_THURSDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_FRIDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_SATURDAY
	DayOfWeek *string
	HourOfDay *int64
}
const MaintenanceWindow_DAY_OF_WEEK_SUNDAY = "SUNDAY"
const MaintenanceWindow_DAY_OF_WEEK_MONDAY = "MONDAY"
const MaintenanceWindow_DAY_OF_WEEK_TUESDAY = "TUESDAY"
const MaintenanceWindow_DAY_OF_WEEK_WEDNESDAY = "WEDNESDAY"
const MaintenanceWindow_DAY_OF_WEEK_THURSDAY = "THURSDAY"
const MaintenanceWindow_DAY_OF_WEEK_FRIDAY = "FRIDAY"
const MaintenanceWindow_DAY_OF_WEEK_SATURDAY = "SATURDAY"

func (s MaintenanceWindow) GetType__() bindings.BindingType {
	return MaintenanceWindowBindingType()
}

func (s MaintenanceWindow) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MaintenanceWindow._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type MaintenanceWindowEntry struct {
    // true if the SDDC is in the defined Mainentance Window
	InMaintenanceWindow *bool
	ReservationSchedule *ReservationSchedule
    // ID for reservation format: uuid
	ReservationId *string
    // true if the SDDC is currently undergoing maintenance
	InMaintenanceMode *bool
    // SDDC ID for this reservation format: uuid
	SddcId *string
}

func (s MaintenanceWindowEntry) GetType__() bindings.BindingType {
	return MaintenanceWindowEntryBindingType()
}

func (s MaintenanceWindowEntry) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MaintenanceWindowEntry._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type MaintenanceWindowGet struct {
    // Possible values are: 
    //
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_SUNDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_MONDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_TUESDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_WEDNESDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_THURSDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_FRIDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_SATURDAY
	DayOfWeek *string
	HourOfDay *int64
	DurationMin *int64
	Version *int64
}
const MaintenanceWindowGet_DAY_OF_WEEK_SUNDAY = "SUNDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_MONDAY = "MONDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_TUESDAY = "TUESDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_WEDNESDAY = "WEDNESDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_THURSDAY = "THURSDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_FRIDAY = "FRIDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_SATURDAY = "SATURDAY"

func (s MaintenanceWindowGet) GetType__() bindings.BindingType {
	return MaintenanceWindowGetBindingType()
}

func (s MaintenanceWindowGet) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MaintenanceWindowGet._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ManagementGatewayTemplate struct {
	PublicIp *SddcPublicIp
	PrimaryDns *string
	SecondaryDns *string
	FirewallRules []FirewallRule
	Vpns []Vpn
    // mgw network subnet cidr
	SubnetCidr *string
}

func (s ManagementGatewayTemplate) GetType__() bindings.BindingType {
	return ManagementGatewayTemplateBindingType()
}

func (s ManagementGatewayTemplate) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ManagementGatewayTemplate._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type MapZonesRequest struct {
    // The connected account ID to remap. This is a standard UUID.
	ConnectedAccountId *string
    // The org ID to remap in. This is a standard UUID.
	OrgId *string
    // A list of Petronas regions to map.
	PetronasRegionsToMap []string
}

func (s MapZonesRequest) GetType__() bindings.BindingType {
	return MapZonesRequestBindingType()
}

func (s MapZonesRequest) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MapZonesRequest._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Start time, end time and interval details.
type MetaDashboardStats struct {
    // Statistics data is collected for these vNICs.
	Vnics []Vnic
    // End time in seconds. format: int64
	EndTime *int64
    // Start time in seconds. format: int64
	StartTime *int64
    // Time interval in seconds. format: int32
	Interval *int64
}

func (s MetaDashboardStats) GetType__() bindings.BindingType {
	return MetaDashboardStatsBindingType()
}

func (s MetaDashboardStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MetaDashboardStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// metadata of the sddc manifest
type Metadata struct {
    // the timestamp for the bundle
	Timestamp *string
    // the cycle id
	CycleId *string
}

func (s Metadata) GetType__() bindings.BindingType {
	return MetadataBindingType()
}

func (s Metadata) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Metadata._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NAT configuration
type Nat struct {
    // Ordered list of NAT rules.
	Rules *NatRules
	FeatureType *string
    // Version number tracking each configuration change. To avoid problems with overwriting changes, always retrieve and modify the latest configuration to include the current version number in your request. If you provide a version number which is not current, the request is rejected. If you omit the version number, the request is accepted but may overwrite any current changes if your change is not in sync with the latest change. format: int64
	Version *int64
    // Value is true if feature is enabled. Default value is true. Optional.
	Enabled *bool
	Template *string
}

func (s Nat) GetType__() bindings.BindingType {
	return NatBindingType()
}

func (s Nat) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Nat._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type NatRule struct {
	RuleType *string
	Protocol *string
	Name *string
	InternalPorts *string
	PublicPorts *string
	PublicIp *string
	InternalIp *string
    // Possible values are: 
    //
    // * NatRule#NatRule_ACTION_DNAT
    // * NatRule#NatRule_ACTION_SNAT
	Action *string
	Id *string
    // current revision of the list of nat rules, used to protect against concurrent modification (first writer wins) format: int32
	Revision *int64
}
const NatRule_ACTION_DNAT = "dnat"
const NatRule_ACTION_SNAT = "snat"

func (s NatRule) GetType__() bindings.BindingType {
	return NatRuleBindingType()
}

func (s NatRule) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NatRule._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Ordered list of NAT rules.
type NatRules struct {
    // Ordered list of NAT rules.
	NatRulesDtos []Nsxnatrule
}

func (s NatRules) GetType__() bindings.BindingType {
	return NatRulesBindingType()
}

func (s NatRules) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NatRules._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type NetworkTemplate struct {
	ManagementGatewayTemplates []ManagementGatewayTemplate
	ComputeGatewayTemplates []ComputeGatewayTemplate
}

func (s NetworkTemplate) GetType__() bindings.BindingType {
	return NetworkTemplateBindingType()
}

func (s NetworkTemplate) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NetworkTemplate._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type NewCredentials struct {
    // Username of the credentials
	Username string
    // Password associated with the credentials
	Password string
    // Name of the credentials
	Name string
}

func (s NewCredentials) GetType__() bindings.BindingType {
	return NewCredentialsBindingType()
}

func (s NewCredentials) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NewCredentials._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Firewall Rule
type Nsxfirewallrule struct {
    // Identifies the type of the rule. internal_high or user.
	RuleType *string
    // Description for the rule
	Description *string
    // Identifier for the rule. format: int64
	RuleId *int64
    // Defines the order of NAT and Firewall pipeline. When false, firewall happens before NAT. Default : false
	MatchTranslated *bool
	InvalidApplication *bool
    // Direction. Possible values in or out. Default is 'any'.
	Direction *string
    // Statistics for the rule
	Statistics *FirewallRuleStats
    // Name for the rule.
	Name *string
	InvalidSource *bool
    // Enable logging for the rule.
	LoggingEnabled *bool
    // List of destinations. Default is any.
	Destination *AddressFWSourceDestination
    // Enable rule.
	Enabled *bool
    // List of applications. Default is any.
	Application *Application
    // List of sources. Default is any.
	Source *AddressFWSourceDestination
    // Action. Values : accept, deny
	Action *string
	InvalidDestination *bool
    // Rule tag. Used to specify user-defined ruleId. If not specified NSX Manager will generate ruleId. format: int64
	RuleTag *int64
}

func (s Nsxfirewallrule) GetType__() bindings.BindingType {
	return NsxfirewallruleBindingType()
}

func (s Nsxfirewallrule) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Nsxfirewallrule._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Application (service) for firewall rule.
type Nsxfirewallservice struct {
    // List of source ports.
	SourcePort []string
    // Protocol.
	Protocol *string
    // List of destination ports.
	Port []string
    // IcmpType. Only supported when protocol is icmp. Default is 'any'.
	IcmpType *string
}

func (s Nsxfirewallservice) GetType__() bindings.BindingType {
	return NsxfirewallserviceBindingType()
}

func (s Nsxfirewallservice) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Nsxfirewallservice._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// L2 VPN server configuration.
type Nsxl2vpn struct {
    // Listener IP addresses.
	ListenerIps []string
    // Enabled state of L2 VPN service.
	Enabled *bool
    // List of L2 VPN sites.
	Sites Sites
}

func (s Nsxl2vpn) GetType__() bindings.BindingType {
	return Nsxl2vpnBindingType()
}

func (s Nsxl2vpn) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Nsxl2vpn._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NAT rule
type Nsxnatrule struct {
    // Interface on which the NAT rule is applied.
	Vnic *string
    // Identifies the type of the rule. internal_high or user.
	RuleType *string
    // Protocol. Default is 'any'
	Protocol *string
    // Description for the rule.
	Description *string
    // Identifier for the rule. format: int64
	RuleId *int64
    // Apply SNAT rule only if traffic has this destination port. Default is 'any'.
	SnatMatchDestinationPort *string
    // Original address or address range. This is the original source address for SNAT rules and the original destination address for DNAT rules.
	OriginalAddress *string
    // Apply DNAT rule only if traffic has this source address. Default is 'any'.
	DnatMatchSourceAddress *string
    // Apply DNAT rule only if traffic has this source port. Default is 'any'.
	DnatMatchSourcePort *string
    // Apply SNAT rule only if traffic has this destination address. Default is 'any'.
	SnatMatchDestinationAddress *string
    // Original port. This is the original source port for SNAT rules, and the original destination port for DNAT rules.
	OriginalPort *string
    // Enable logging for the rule.
	LoggingEnabled *bool
    // Translated address or address range.
	TranslatedAddress *string
    // Enable rule.
	Enabled *bool
    // ICMP type. Only supported when protocol is icmp. Default is 'any'.
	IcmpType *string
    // Translated port. Supported in DNAT rules only.
	TranslatedPort *string
    // Action for the rule. SNAT or DNAT.
	Action *string
    // Rule tag. Used to specify user-defined ruleId. If not specified NSX Manager will generate ruleId. format: int64
	RuleTag *int64
}

func (s Nsxnatrule) GetType__() bindings.BindingType {
	return NsxnatruleBindingType()
}

func (s Nsxnatrule) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Nsxnatrule._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// L2 VPN site.
type Nsxsite struct {
    // Secure L2VPN traffic.
	SecureTraffic *bool
    // Identifier for L2 VPN site.
	SiteId *string
    // Name of L2 VPN site. Length: 1-255 characters.
	Name *string
    // Password for L2 VPN user. Passwords must contain the following: 12-63 characters, a mix of upper case letters, lower case letters, numbers, and at least one special character. Password must not contain the username as a substring. Do not repeat a character 3 or more times.
	Password *string
    // L2 VPN user ID. Valid user names: 1-63 characters, letters and numbers only. No white space or special characters.
	UserId *string
    // Description of L2 VPN site.
	Description *string
}

func (s Nsxsite) GetType__() bindings.BindingType {
	return NsxsiteBindingType()
}

func (s Nsxsite) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Nsxsite._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Details the state of different NSX add-ons.
type NsxtAddons struct {
    // Indicates whether NSX Advanced addon is enabled or disabled.
	EnableNsxAdvancedAddon *bool
}

func (s NsxtAddons) GetType__() bindings.BindingType {
	return NsxtAddonsBindingType()
}

func (s NsxtAddons) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NsxtAddons._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ObjectType struct {
	Name *string
}

func (s ObjectType) GetType__() bindings.BindingType {
	return ObjectTypeBindingType()
}

func (s ObjectType) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ObjectType._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Holder for the offer instances.
type OfferInstancesHolder struct {
	OnDemand OnDemandOfferInstance
	Offers []TermOfferInstance
}

func (s OfferInstancesHolder) GetType__() bindings.BindingType {
	return OfferInstancesHolderBindingType()
}

func (s OfferInstancesHolder) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OfferInstancesHolder._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Holder for the on-demand offer instance.
type OnDemandOfferInstance struct {
	Product string
    // Deprecated. Please use product and type fields instead.
	ProductType *string
	Name string
	Currency string
	Region string
	UnitPrice string
	MonthlyCost string
	Version string
	Type_ string
	Description string
}

func (s OnDemandOfferInstance) GetType__() bindings.BindingType {
	return OnDemandOfferInstanceBindingType()
}

func (s OnDemandOfferInstance) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OnDemandOfferInstance._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type OrgProperties struct {
    // A map of string properties to values.
	Values map[string]string
}

func (s OrgProperties) GetType__() bindings.BindingType {
	return OrgPropertiesBindingType()
}

func (s OrgProperties) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OrgProperties._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Organization struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
    // ORG_TYPE to be associated with the org
	OrgType *string
	DisplayName *string
	Name *string
    // Possible values are: 
    //
    // * Organization#Organization_PROJECT_STATE_CREATED
    // * Organization#Organization_PROJECT_STATE_DELETED
	ProjectState *string
	Properties *OrgProperties
}
const Organization_PROJECT_STATE_CREATED = "CREATED"
const Organization_PROJECT_STATE_DELETED = "DELETED"

func (s Organization) GetType__() bindings.BindingType {
	return OrganizationBindingType()
}

func (s Organization) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Organization._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NSX Edges listed by pages.
type PagedEdgeList struct {
    // Page details with matched records.
	EdgePage *DataPageEdgeSummary
}

func (s PagedEdgeList) GetType__() bindings.BindingType {
	return PagedEdgeListBindingType()
}

func (s PagedEdgeList) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PagedEdgeList._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type PagingInfo struct {
	SortOrderAscending *bool
	TotalCount *int64
	StartIndex *int64
	SortBy *string
	PageSize *int64
}

func (s PagingInfo) GetType__() bindings.BindingType {
	return PagingInfoBindingType()
}

func (s PagingInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PagingInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type PaymentMethodInfo struct {
	Type_ *string
	DefaultFlag *bool
	PaymentMethodId *string
}

func (s PaymentMethodInfo) GetType__() bindings.BindingType {
	return PaymentMethodInfoBindingType()
}

func (s PaymentMethodInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PaymentMethodInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type PopAgentXeniConnection struct {
    // The gateway route ip fo the subnet.
	DefaultSubnetRoute *string
	EniInfo *EniInfo
}

func (s PopAgentXeniConnection) GetType__() bindings.BindingType {
	return PopAgentXeniConnectionBindingType()
}

func (s PopAgentXeniConnection) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PopAgentXeniConnection._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type PopAmiInfo struct {
    // instance type of the esx ami
	InstanceType *string
    // the region of the esx ami
	Region *string
    // the ami id for the esx
	Id *string
    // the name of the esx ami
	Name *string
    // Possible values are: 
    //
    // * PopAmiInfo#PopAmiInfo_TYPE_CENTOS
    // * PopAmiInfo#PopAmiInfo_TYPE_POP
    //
    //  PoP AMI type. CENTOS: a Centos AMI; POP: a PoP AMI.
	Type_ *string
}
const PopAmiInfo_TYPE_CENTOS = "CENTOS"
const PopAmiInfo_TYPE_POP = "POP"

func (s PopAmiInfo) GetType__() bindings.BindingType {
	return PopAmiInfoBindingType()
}

func (s PopAmiInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PopAmiInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Present a SDDC PoP information.
type PopInfo struct {
    // A map of [region name of PoP / PoP-AMI]:[PopAmiInfo].
	AmiInfos map[string]PopAmiInfo
    // The PopInfo (or PoP AMI) created time. Using ISO 8601 date-time pattern. format: date-time
	CreatedAt *time.Time
    // A map of [service type]:[PopServiceInfo]
	ServiceInfos map[string]PopServiceInfo
    // UUID of the PopInfo format: UUID
	Id *string
    // version of the manifest.
	ManifestVersion *string
}

func (s PopInfo) GetType__() bindings.BindingType {
	return PopInfoBindingType()
}

func (s PopInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PopInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type PopServiceInfo struct {
    // The service change set number.
	Cln *string
    // The service API version.
	Version *string
    // The service build number.
	Build *string
    // Possible values are: 
    //
    // * PopServiceInfo#PopServiceInfo_SERVICE_OS
    // * PopServiceInfo#PopServiceInfo_SERVICE_AGENT
    // * PopServiceInfo#PopServiceInfo_SERVICE_GLCM
    // * PopServiceInfo#PopServiceInfo_SERVICE_S3_ADAPTER
    // * PopServiceInfo#PopServiceInfo_SERVICE_JRE
    // * PopServiceInfo#PopServiceInfo_SERVICE_DOCKER
    // * PopServiceInfo#PopServiceInfo_SERVICE_AIDE
    // * PopServiceInfo#PopServiceInfo_SERVICE_RTS
    // * PopServiceInfo#PopServiceInfo_SERVICE_FM_LOG_COLLECTOR
    // * PopServiceInfo#PopServiceInfo_SERVICE_FM_METRICS_COLLECTOR
    // * PopServiceInfo#PopServiceInfo_SERVICE_BRE
    // * PopServiceInfo#PopServiceInfo_SERVICE_BRF
    // * PopServiceInfo#PopServiceInfo_SERVICE_REVERSE_PROXY
    // * PopServiceInfo#PopServiceInfo_SERVICE_FORWARD_PROXY
    // * PopServiceInfo#PopServiceInfo_SERVICE_DNS
    // * PopServiceInfo#PopServiceInfo_SERVICE_NTP
    // * PopServiceInfo#PopServiceInfo_SERVICE_LOGZ_LOG_COLLECTOR
    //
    //  An enum of PoP related services (including os platform and JRE).
	Service *string
}
const PopServiceInfo_SERVICE_OS = "OS"
const PopServiceInfo_SERVICE_AGENT = "AGENT"
const PopServiceInfo_SERVICE_GLCM = "GLCM"
const PopServiceInfo_SERVICE_S3_ADAPTER = "S3_ADAPTER"
const PopServiceInfo_SERVICE_JRE = "JRE"
const PopServiceInfo_SERVICE_DOCKER = "DOCKER"
const PopServiceInfo_SERVICE_AIDE = "AIDE"
const PopServiceInfo_SERVICE_RTS = "RTS"
const PopServiceInfo_SERVICE_FM_LOG_COLLECTOR = "FM_LOG_COLLECTOR"
const PopServiceInfo_SERVICE_FM_METRICS_COLLECTOR = "FM_METRICS_COLLECTOR"
const PopServiceInfo_SERVICE_BRE = "BRE"
const PopServiceInfo_SERVICE_BRF = "BRF"
const PopServiceInfo_SERVICE_REVERSE_PROXY = "REVERSE_PROXY"
const PopServiceInfo_SERVICE_FORWARD_PROXY = "FORWARD_PROXY"
const PopServiceInfo_SERVICE_DNS = "DNS"
const PopServiceInfo_SERVICE_NTP = "NTP"
const PopServiceInfo_SERVICE_LOGZ_LOG_COLLECTOR = "LOGZ_LOG_COLLECTOR"

func (s PopServiceInfo) GetType__() bindings.BindingType {
	return PopServiceInfoBindingType()
}

func (s PopServiceInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PopServiceInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Represents a provisioning spec for a sddc
type ProvisionSpec struct {
    // Map of provider to sddc config spec
	Provider map[string]SddcConfigSpec
}

func (s ProvisionSpec) GetType__() bindings.BindingType {
	return ProvisionSpecBindingType()
}

func (s ProvisionSpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ProvisionSpec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DNS request statistics.
type Requests struct {
	Total *int64
	Queries *int64
}

func (s Requests) GetType__() bindings.BindingType {
	return RequestsBindingType()
}

func (s Requests) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Requests._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Reservation struct {
    // Duration - required for reservation in maintenance window format: int64
	Duration *int64
    // Reservation ID format: uuid
	Rid *string
    // Optional
	CreateTime *string
    // Start time of a reservation format: date-time
	StartTime *time.Time
    // Optional
	Metadata map[string]string
}

func (s Reservation) GetType__() bindings.BindingType {
	return ReservationBindingType()
}

func (s Reservation) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Reservation._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ReservationInMw struct {
    // Reservation ID format: uuid
	Rid *string
    // SUNDAY of the week that maintenance is scheduled, ISO format date
	WeekOf *string
    // Optional format: date-time
	CreateTime *time.Time
    // Optional
	Metadata map[string]string
}

func (s ReservationInMw) GetType__() bindings.BindingType {
	return ReservationInMwBindingType()
}

func (s ReservationInMw) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ReservationInMw._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ReservationSchedule struct {
    // Possible values are: 
    //
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_SUNDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_MONDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_TUESDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_WEDNESDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_THURSDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_FRIDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_SATURDAY
	DayOfWeek *string
	HourOfDay *int64
	DurationMin *int64
	Version *int64
	Reservations []Reservation
	ReservationsMw []ReservationInMw
}
const ReservationSchedule_DAY_OF_WEEK_SUNDAY = "SUNDAY"
const ReservationSchedule_DAY_OF_WEEK_MONDAY = "MONDAY"
const ReservationSchedule_DAY_OF_WEEK_TUESDAY = "TUESDAY"
const ReservationSchedule_DAY_OF_WEEK_WEDNESDAY = "WEDNESDAY"
const ReservationSchedule_DAY_OF_WEEK_THURSDAY = "THURSDAY"
const ReservationSchedule_DAY_OF_WEEK_FRIDAY = "FRIDAY"
const ReservationSchedule_DAY_OF_WEEK_SATURDAY = "SATURDAY"

func (s ReservationSchedule) GetType__() bindings.BindingType {
	return ReservationScheduleBindingType()
}

func (s ReservationSchedule) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ReservationSchedule._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ReservationWindow struct {
    // Possible values are: 
    //
    // * ReservationWindow#ReservationWindow_RESERVATION_STATE_SCHEDULED
    // * ReservationWindow#ReservationWindow_RESERVATION_STATE_RUNNING
    // * ReservationWindow#ReservationWindow_RESERVATION_STATE_CANCELED
    // * ReservationWindow#ReservationWindow_RESERVATION_STATE_COMPLETED
	ReservationState *string
	Emergency *bool
	MaintenanceProperties *ReservationWindowMaintenanceProperties
	ReserveId *string
	StartHour *int64
	SddcId *string
	ManifestId *string
	DurationHours *int64
	StartDate *string
    // Metadata for reservation window, in key-value form
	Metadata map[string]string
}
const ReservationWindow_RESERVATION_STATE_SCHEDULED = "SCHEDULED"
const ReservationWindow_RESERVATION_STATE_RUNNING = "RUNNING"
const ReservationWindow_RESERVATION_STATE_CANCELED = "CANCELED"
const ReservationWindow_RESERVATION_STATE_COMPLETED = "COMPLETED"

func (s ReservationWindow) GetType__() bindings.BindingType {
	return ReservationWindowBindingType()
}

func (s ReservationWindow) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ReservationWindow._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ReservationWindowMaintenanceProperties struct {
    // Status of upgrade, if any
	Status *string
}

func (s ReservationWindowMaintenanceProperties) GetType__() bindings.BindingType {
	return ReservationWindowMaintenancePropertiesBindingType()
}

func (s ReservationWindowMaintenanceProperties) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ReservationWindowMaintenanceProperties._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Job result information for the configuration change carried out on NSX Edge.
type Result struct {
    // Job Result value associated with key ID.
	Value *string
    // Job Result key ID.
	Key *string
}

func (s Result) GetType__() bindings.BindingType {
	return ResultBindingType()
}

func (s Result) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Result._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type RouteTableInfo struct {
    // route table name
	Name *string
    // route table id
	Id *string
}

func (s RouteTableInfo) GetType__() bindings.BindingType {
	return RouteTableInfoBindingType()
}

func (s RouteTableInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for RouteTableInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ScopeInfo struct {
	ObjectTypeName *string
	Id *string
	Name *string
}

func (s ScopeInfo) GetType__() bindings.BindingType {
	return ScopeInfoBindingType()
}

func (s ScopeInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ScopeInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Sddc struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
    // name for SDDC
	Name *string
    // Possible values are: 
    //
    // * Sddc#Sddc_SDDC_STATE_DEPLOYING
    // * Sddc#Sddc_SDDC_STATE_READY
    // * Sddc#Sddc_SDDC_STATE_DELETING
    // * Sddc#Sddc_SDDC_STATE_DELETED
    // * Sddc#Sddc_SDDC_STATE_FAILED
    // * Sddc#Sddc_SDDC_STATE_CANCELED
    // * Sddc#Sddc_SDDC_STATE_READY_FOR_GLCM_BRINGUP
	SddcState *string
    // Expiration date of a sddc in UTC (will be set if its applicable) format: date-time
	ExpirationDate *time.Time
	OrgId *string
    // Type of the sddc
	SddcType *string
    // Possible values are: 
    //
    // * Sddc#Sddc_PROVIDER_AWS
    // * Sddc#Sddc_PROVIDER_ZEROCLOUD
	Provider *string
    // Possible values are: 
    //
    // * Sddc#Sddc_ACCOUNT_LINK_STATE_DELAYED
    // * Sddc#Sddc_ACCOUNT_LINK_STATE_LINKED
    // * Sddc#Sddc_ACCOUNT_LINK_STATE_UNLINKED
    //
    //  Account linking state of the sddc
	AccountLinkState *string
    // Describes the access state of sddc, valid state is DISABLED or ENABLED
	SddcAccessState *string
	ResourceConfig *AwsSddcResourceConfig
}
const Sddc_SDDC_STATE_DEPLOYING = "DEPLOYING"
const Sddc_SDDC_STATE_READY = "READY"
const Sddc_SDDC_STATE_DELETING = "DELETING"
const Sddc_SDDC_STATE_DELETED = "DELETED"
const Sddc_SDDC_STATE_FAILED = "FAILED"
const Sddc_SDDC_STATE_CANCELED = "CANCELED"
const Sddc_SDDC_STATE_READY_FOR_GLCM_BRINGUP = "READY_FOR_GLCM_BRINGUP"
const Sddc_PROVIDER_AWS = "AWS"
const Sddc_PROVIDER_ZEROCLOUD = "ZEROCLOUD"
const Sddc_ACCOUNT_LINK_STATE_DELAYED = "DELAYED"
const Sddc_ACCOUNT_LINK_STATE_LINKED = "LINKED"
const Sddc_ACCOUNT_LINK_STATE_UNLINKED = "UNLINKED"

func (s Sddc) GetType__() bindings.BindingType {
	return SddcBindingType()
}

func (s Sddc) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Sddc._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcAllocatePublicIpSpec struct {
	Count int64
    // List of workload VM private IPs to be assigned the public IP just allocated.
	PrivateIps []string
    // List of names for the workload VM public IP assignment.
	Names []string
}

func (s SddcAllocatePublicIpSpec) GetType__() bindings.BindingType {
	return SddcAllocatePublicIpSpecBindingType()
}

func (s SddcAllocatePublicIpSpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcAllocatePublicIpSpec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcConfig struct {
    // AWS VPC IP range. Only prefix of 16 or 20 is currently supported.
	VpcCidr *string
    // The instance type for the esx hosts in the primary cluster of the SDDC.
	HostInstanceType *HostInstanceTypes
    // skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan *bool
    // VXLAN IP subnet in CIDR for compute gateway
	VxlanSubnet *string
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_SIZE_NSX_SMALL
    // * SddcConfig#SddcConfig_SIZE_MEDIUM
    // * SddcConfig#SddcConfig_SIZE_LARGE
    // * SddcConfig#SddcConfig_SIZE_NSX_LARGE
    //
    //  The size of the vCenter and NSX appliances. \"large\" sddcSize corresponds to a 'large' vCenter appliance and 'large' NSX appliance. 'medium' sddcSize corresponds to 'medium' vCenter appliance and 'medium' NSX appliance. Value defaults to 'medium'.
	Size *string
    // The storage capacity value to be requested for the sddc primary cluster, in GiBs. If provided, instead of using the direct-attached storage, a capacity value amount of seperable storage will be used. format: int64
	StorageCapacity *int64
	Name string
    // A list of the SDDC linking configurations to use.
	AccountLinkSddcConfig []AccountLinkSddcConfig
    // If provided, will be assigned as SDDC id of the provisioned SDDC. format: UUID
	SddcId *string
	NumHosts int64
    // Denotes the sddc type , if the value is null or empty, the type is considered as default.
	SddcType *string
    // The account linking configuration, we will keep this one and remove accountLinkSddcConfig finally.
	AccountLinkConfig *AccountLinkConfig
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_PROVIDER_AWS
    // * SddcConfig#SddcConfig_PROVIDER_ZEROCLOUD
    //
    //  Determines what additional properties are available based on cloud provider.
	Provider string
    // The SSO domain name to use for vSphere users. If not specified, vmc.local will be used.
	SsoDomain *string
    // If provided, configuration from the template will applied to the provisioned SDDC. format: UUID
	SddcTemplateId *string
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_DEPLOYMENT_TYPE_SINGLEAZ
    // * SddcConfig#SddcConfig_DEPLOYMENT_TYPE_MULTIAZ
    //
    //  Denotes if request is for a SingleAZ or a MultiAZ SDDC. Default is SingleAZ.
	DeploymentType *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const SddcConfig__TYPE_IDENTIFIER = "SddcConfig"
const SddcConfig_SIZE_NSX_SMALL = "nsx_small"
const SddcConfig_SIZE_MEDIUM = "medium"
const SddcConfig_SIZE_LARGE = "large"
const SddcConfig_SIZE_NSX_LARGE = "nsx_large"
const SddcConfig_PROVIDER_AWS = "AWS"
const SddcConfig_PROVIDER_ZEROCLOUD = "ZEROCLOUD"
const SddcConfig_DEPLOYMENT_TYPE_SINGLEAZ = "SingleAZ"
const SddcConfig_DEPLOYMENT_TYPE_MULTIAZ = "MultiAZ"

func (s SddcConfig) GetType__() bindings.BindingType {
	return SddcConfigBindingType()
}

func (s SddcConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Represents a configuration spec for a sddc
type SddcConfigSpec struct {
    // Map of sddc type to config spec
	SddcTypeConfigSpec map[string]ConfigSpec
    // The region name to display names mapping
	RegionDisplayNames map[string]string
}

func (s SddcConfigSpec) GetType__() bindings.BindingType {
	return SddcConfigSpecBindingType()
}

func (s SddcConfigSpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcConfigSpec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcId struct {
    // Sddc ID
	SddcId *string
}

func (s SddcId) GetType__() bindings.BindingType {
	return SddcIdBindingType()
}

func (s SddcId) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcId._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcLinkConfig struct {
	CustomerSubnetIds []string
    // Determines which connected customer account to link to
	ConnectedAccountId *string
}

func (s SddcLinkConfig) GetType__() bindings.BindingType {
	return SddcLinkConfigBindingType()
}

func (s SddcLinkConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcLinkConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Describes software components of the installed SDDC
type SddcManifest struct {
    // the vmcVersion of the sddc for display
	VmcVersion *string
	GlcmBundle *GlcmBundle
	PopInfo *PopInfo
    // the vmcInternalVersion of the sddc for internal use
	VmcInternalVersion *string
	EbsBackedVsanConfig *EbsBackedVsanConfig
	VsanWitnessAmi *AmiInfo
	EsxAmi *AmiInfo
	EsxNsxtAmi *AmiInfo
	Metadata *Metadata
}

func (s SddcManifest) GetType__() bindings.BindingType {
	return SddcManifestBindingType()
}

func (s SddcManifest) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcManifest._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Logical network.
type SddcNetwork struct {
    // Network address groups for routed logical networks.
	Subnets *SddcNetworkAddressGroups
    // Name of the compute gateway to which the logical network is attached.
	CgwName *string
    // Name of logical network. Length needs to be between 1-35 characters.
	Name string
    // Layer 2 extension for extended logical networks.
	L2Extension *L2Extension
    // ID of the compute gateway edge to which the logical network is attached.
	CgwId string
    // DHCP configuration for routed logical networks.
	DhcpConfigs *SddcNetworkDhcpConfig
    // ID of logical network.
	Id *string
}

func (s SddcNetwork) GetType__() bindings.BindingType {
	return SddcNetworkBindingType()
}

func (s SddcNetwork) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcNetwork._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Logical Network address group.
type SddcNetworkAddressGroup struct {
    // Prefix length of logical network.
	PrefixLength *string
    // Primary address for logical network.
	PrimaryAddress *string
}

func (s SddcNetworkAddressGroup) GetType__() bindings.BindingType {
	return SddcNetworkAddressGroupBindingType()
}

func (s SddcNetworkAddressGroup) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcNetworkAddressGroup._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Logical network address groups.
type SddcNetworkAddressGroups struct {
    // List of logical network address groups.
	AddressGroups []SddcNetworkAddressGroup
}

func (s SddcNetworkAddressGroups) GetType__() bindings.BindingType {
	return SddcNetworkAddressGroupsBindingType()
}

func (s SddcNetworkAddressGroups) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcNetworkAddressGroups._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DHCP configuration for the logical network.
type SddcNetworkDhcpConfig struct {
    // List of IP pools in DHCP configuration.
	IpPools []SddcNetworkDhcpIpPool
}

func (s SddcNetworkDhcpConfig) GetType__() bindings.BindingType {
	return SddcNetworkDhcpConfigBindingType()
}

func (s SddcNetworkDhcpConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcNetworkDhcpConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// DHCP IP pool for logical network.
type SddcNetworkDhcpIpPool struct {
    // IP range for DHCP IP pool.
	IpRange *string
    // DNS domain name.
	DomainName *string
}

func (s SddcNetworkDhcpIpPool) GetType__() bindings.BindingType {
	return SddcNetworkDhcpIpPoolBindingType()
}

func (s SddcNetworkDhcpIpPool) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcNetworkDhcpIpPool._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Patch request body for SDDC
type SddcPatchRequest struct {
    // The new name of the SDDC to be changed to.
	Name *string
}

func (s SddcPatchRequest) GetType__() bindings.BindingType {
	return SddcPatchRequestBindingType()
}

func (s SddcPatchRequest) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcPatchRequest._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcPublicIp struct {
	PublicIp string
	Name *string
	AllocationId *string
	DnatRuleId *string
	AssociatedPrivateIp *string
	SnatRuleId *string
}

func (s SddcPublicIp) GetType__() bindings.BindingType {
	return SddcPublicIpBindingType()
}

func (s SddcPublicIp) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcPublicIp._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcResourceConfig struct {
    // Name for management appliance network.
	MgmtApplianceNetworkName *string
    // if true, NSX-T UI is enabled.
	Nsxt *bool
    // Management Gateway Id
	MgwId *string
    // URL of the NSX Manager
	NsxMgrUrl *string
    // PSC internal management IP
	PscManagementIp *string
    // URL of the PSC server
	PscUrl *string
	Cgws []string
    // Availability zones over which esx hosts are provisioned. MultiAZ SDDCs will have hosts provisioned over two availability zones while SingleAZ SDDCs will provision over one.
	AvailabilityZones []string
    // The ManagedObjectReference of the management Datastore
	ManagementDs *string
    // nsx api entire base url
	NsxApiPublicEndpointUrl *string
	CustomProperties map[string]string
    // Password for vCenter SDDC administrator
	CloudPassword *string
    // Possible values are: 
    //
    // * SddcResourceConfig#SddcResourceConfig_PROVIDER_AWS
    // * SddcResourceConfig#SddcResourceConfig_PROVIDER_ZEROCLOUD
    //
    //  Discriminator for additional properties
	Provider string
    // List of clusters in the SDDC.
	Clusters []Cluster
    // vCenter internal management IP
	VcManagementIp *string
	SddcNetworks []string
    // Username for vCenter SDDC administrator
	CloudUsername *string
	EsxHosts []AwsEsxHost
    // NSX Manager internal management IP
	NsxMgrManagementIp *string
    // unique id of the vCenter server
	VcInstanceId *string
    // Cluster Id to add ESX workflow
	EsxClusterId *string
    // vCenter public IP
	VcPublicIp *string
    // skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan *bool
    // URL of the vCenter server
	VcUrl *string
	SddcManifest *SddcManifest
    // VXLAN IP subnet
	VxlanSubnet *string
    // Group name for vCenter SDDC administrator
	CloudUserGroup *string
	ManagementRp *string
    // region in which sddc is deployed
	Region *string
    // Availability zone where the witness node is provisioned for a MultiAZ SDDC. This is null for a SingleAZ SDDC.
	WitnessAvailabilityZone *string
    // sddc identifier
	SddcId *string
	PopAgentXeniConnection *PopAgentXeniConnection
	SddcSize *SddcSize
    // List of Controller IPs
	NsxControllerIps []string
    // ESX host subnet
	EsxHostSubnet *string
    // The SSO domain name to use for vSphere users
	SsoDomain *string
    // Possible values are: 
    //
    // * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_SINGLE_AZ
    // * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_MULTI_AZ
    //
    //  Denotes if this is a SingleAZ SDDC or a MultiAZ SDDC.
	DeploymentType *string
	NsxtAddons *NsxtAddons
    // if true, use the private IP addresses to register DNS records for the management VMs
	DnsWithManagementVmPrivateIp *bool
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const SddcResourceConfig__TYPE_IDENTIFIER = "SddcResourceConfig"
const SddcResourceConfig_PROVIDER_AWS = "AWS"
const SddcResourceConfig_PROVIDER_ZEROCLOUD = "ZEROCLOUD"
const SddcResourceConfig_DEPLOYMENT_TYPE_SINGLE_AZ = "SINGLE_AZ"
const SddcResourceConfig_DEPLOYMENT_TYPE_MULTI_AZ = "MULTI_AZ"

func (s SddcResourceConfig) GetType__() bindings.BindingType {
	return SddcResourceConfigBindingType()
}

func (s SddcResourceConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcResourceConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Size of the SDDC
type SddcSize struct {
    // Possible values are: 
    //
    // * SddcSize#SddcSize_VC_SIZE_MEDIUM
    // * SddcSize#SddcSize_VC_SIZE_LARGE
	VcSize *string
    // Possible values are: 
    //
    // * SddcSize#SddcSize_NSX_SIZE_SMALL
    // * SddcSize#SddcSize_NSX_SIZE_MEDIUM
    // * SddcSize#SddcSize_NSX_SIZE_LARGE
	NsxSize *string
    // Possible values are: 
    //
    // * SddcSize#SddcSize_SIZE_NSX_SMALL
    // * SddcSize#SddcSize_SIZE_MEDIUM
    // * SddcSize#SddcSize_SIZE_LARGE
    // * SddcSize#SddcSize_SIZE_NSX_LARGE
	Size *string
}
const SddcSize_VC_SIZE_MEDIUM = "medium"
const SddcSize_VC_SIZE_LARGE = "large"
const SddcSize_NSX_SIZE_SMALL = "small"
const SddcSize_NSX_SIZE_MEDIUM = "medium"
const SddcSize_NSX_SIZE_LARGE = "large"
const SddcSize_SIZE_NSX_SMALL = "NSX_SMALL"
const SddcSize_SIZE_MEDIUM = "MEDIUM"
const SddcSize_SIZE_LARGE = "LARGE"
const SddcSize_SIZE_NSX_LARGE = "NSX_LARGE"

func (s SddcSize) GetType__() bindings.BindingType {
	return SddcSizeBindingType()
}

func (s SddcSize) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcSize._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcStateRequest struct {
	Sddcs []string
    // Possible values are: 
    //
    // * SddcStateRequest#SddcStateRequest_STATES_SCHEDULED
    // * SddcStateRequest#SddcStateRequest_STATES_RUNNING
    // * SddcStateRequest#SddcStateRequest_STATES_CANCELED
    // * SddcStateRequest#SddcStateRequest_STATES_COMPLETED
	States []string
}
const SddcStateRequest_STATES_SCHEDULED = "SCHEDULED"
const SddcStateRequest_STATES_RUNNING = "RUNNING"
const SddcStateRequest_STATES_CANCELED = "CANCELED"
const SddcStateRequest_STATES_COMPLETED = "COMPLETED"

func (s SddcStateRequest) GetType__() bindings.BindingType {
	return SddcStateRequestBindingType()
}

func (s SddcStateRequest) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcStateRequest._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcTemplate struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
    // A list of the SDDC linking configurations to use.
	AccountLinkSddcConfigs []AccountLinkSddcConfig
    // Possible values are: 
    //
    // * SddcTemplate#SddcTemplate_STATE_INITIALIZATION
    // * SddcTemplate#SddcTemplate_STATE_AVAILABLE
    // * SddcTemplate#SddcTemplate_STATE_INUSE
    // * SddcTemplate#SddcTemplate_STATE_APPLIED
    // * SddcTemplate#SddcTemplate_STATE_DELETING
    // * SddcTemplate#SddcTemplate_STATE_DELETED
    // * SddcTemplate#SddcTemplate_STATE_FAILED
	State *string
	NetworkTemplate *NetworkTemplate
    // name for SDDC configuration template
	Name *string
	SourceSddcId *string
	OrgId *string
	Sddc *Sddc
}
const SddcTemplate_STATE_INITIALIZATION = "INITIALIZATION"
const SddcTemplate_STATE_AVAILABLE = "AVAILABLE"
const SddcTemplate_STATE_INUSE = "INUSE"
const SddcTemplate_STATE_APPLIED = "APPLIED"
const SddcTemplate_STATE_DELETING = "DELETING"
const SddcTemplate_STATE_DELETED = "DELETED"
const SddcTemplate_STATE_FAILED = "FAILED"

func (s SddcTemplate) GetType__() bindings.BindingType {
	return SddcTemplateBindingType()
}

func (s SddcTemplate) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcTemplate._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Secondary IP addresses of the NSX Edge vnic address group. These are used for NAT, LB, VPN etc. Optional.
type SecondaryAddresses struct {
	Type_ *string
    // List of IP addresses.
	IpAddress []string
}

func (s SecondaryAddresses) GetType__() bindings.BindingType {
	return SecondaryAddressesBindingType()
}

func (s SecondaryAddresses) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SecondaryAddresses._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Detailed service errors associated with a task.
type ServiceError struct {
    // Error message in English.
	DefaultMessage *string
    // The original service name of the error.
	OriginalService *string
    // The localized message.
	LocalizedMessage *string
    // The original error code of the service.
	OriginalServiceErrorCode *string
}

func (s ServiceError) GetType__() bindings.BindingType {
	return ServiceErrorBindingType()
}

func (s ServiceError) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ServiceError._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Site struct {
    // Site password.
	Password *string
    // Site user id.
	UserId *string
    // Unique name for the site getting configured.
	Name *string
    // Bytes received on local network. format: int64
	RxBytesOnLocalSubnet *int64
    // Enable/disable encription.
	SecureTraffic *bool
    // Date tunnel was established.
	EstablishedDate *string
    // failure message.
	FailureMessage *string
    // Number of transmitted packets dropped.
	DroppedTxPackets *string
    // Number of received packets dropped.
	DroppedRxPackets *string
    // Possible values are: 
    //
    // * Site#Site_TUNNEL_STATUS_CONNECTED
    // * Site#Site_TUNNEL_STATUS_DISCONNECTED
    // * Site#Site_TUNNEL_STATUS_UNKNOWN
    //
    //  Site tunnel status.
	TunnelStatus *string
    // Bytes transmitted from local subnet. format: int64
	TxBytesFromLocalSubnet *int64
}
const Site_TUNNEL_STATUS_CONNECTED = "CONNECTED"
const Site_TUNNEL_STATUS_DISCONNECTED = "DISCONNECTED"
const Site_TUNNEL_STATUS_UNKNOWN = "UNKNOWN"

func (s Site) GetType__() bindings.BindingType {
	return SiteBindingType()
}

func (s Site) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Site._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// L2 VPN sites.
type Sites struct {
	Sites []Nsxsite
}

func (s Sites) GetType__() bindings.BindingType {
	return SitesBindingType()
}

func (s Sites) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Sites._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Dashboard Statistics data for SSL VPN.
type SslvpnDashboardStats struct {
    // Number of active clients.
	ActiveClients []DashboardStat
    // Rx bytes received for SSL VPN.
	SslvpnBytesIn []DashboardStat
    // Number of authentication failures.
	AuthFailures []DashboardStat
    // Number of SSL VPN sessions created.
	SessionsCreated []DashboardStat
    // Tx bytes transmitted for SSL VPN.
	SslvpnBytesOut []DashboardStat
}

func (s SslvpnDashboardStats) GetType__() bindings.BindingType {
	return SslvpnDashboardStatsBindingType()
}

func (s SslvpnDashboardStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SslvpnDashboardStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NSX Edge sub interface configuration details. Sub interfaces are created on a trunk interface.
type SubInterface struct {
    // Index of the sub interface assigned by NSX Manager. Min value is 10 and max value is 4103. format: int32
	Index *int64
    // Valid values for tunnel ID are min 1 to max 4093. Required. format: int32
	TunnelId int64
    // Name of the sub interface. Required.
	Name *string
    // Address group configuration of the sub interface.
	AddressGroups *EdgeVnicAddressGroups
    // VLAN ID of the virtual LAN used by this sub interface. VLAN IDs can range from 0 to 4094. format: int32
	VlanId *int64
    // Sub interface label of format vNic_{index} provided by NSX Manager. Read only.
	Label *string
    // Name of the logical switch connected to this sub interface.
	LogicalSwitchName *string
    // Value is true if the sub interface is connected to a logical switch, standard portgroup or distributed portgroup.
	IsConnected *bool
    // MTU value of the sub interface. This value would be the least mtu for all the trunk interfaces of the NSX Edge. Default is 1500. format: int32
	Mtu *int64
    // ID of the logical switch connected to this sub interface.
	LogicalSwitchId *string
    // Value is true if send redirects is enabled. Enable ICMP redirect to convey routing information to hosts.
	EnableSendRedirects *bool
}

func (s SubInterface) GetType__() bindings.BindingType {
	return SubInterfaceBindingType()
}

func (s SubInterface) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubInterface._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SubInterfaces struct {
    // List of sub interfaces.
	SubInterfaces []SubInterface
}

func (s SubInterfaces) GetType__() bindings.BindingType {
	return SubInterfacesBindingType()
}

func (s SubInterfaces) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubInterfaces._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// (as there's already one SubnetInfo, use Subnet instead)
type Subnet struct {
    // subnet id
	SubnetId *string
    // subnet name
	Name *string
	RouteTables []SubnetRouteTableInfo
}

func (s Subnet) GetType__() bindings.BindingType {
	return SubnetBindingType()
}

func (s Subnet) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Subnet._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SubnetInfo struct {
    // Is this customer subnet compatible with the SDDC?
	Compatible *bool
    // The ID of the connected account this subnet is from.
	ConnectedAccountId *string
    // The region this subnet is from.
	RegionName *string
    // The availability zone (customer-centric) this subnet is in.
	AvailabilityZone *string
    // The ID of the subnet.
	SubnetId *string
    // The availability zone id (customer-centric) this subnet is in.
	AvailabilityZoneId *string
    // The CIDR block of the subnet.
	SubnetCidrBlock *string
    // Why a subnet is marked as not compatible. May be blank if compatible.
	Note *string
    // The ID of the VPC this subnet resides in.
	VpcId *string
    // The CIDR block of the VPC containing this subnet.
	VpcCidrBlock *string
    // The name of the subnet. This is either the tagged name or the default AWS id it was given.
	Name *string
}

func (s SubnetInfo) GetType__() bindings.BindingType {
	return SubnetInfoBindingType()
}

func (s SubnetInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubnetInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SubnetRouteTableInfo struct {
    // subnet id
	SubnetId *string
    // subnet - route table association id
	AssociationId *string
    // route table id
	RoutetableId *string
}

func (s SubnetRouteTableInfo) GetType__() bindings.BindingType {
	return SubnetRouteTableInfoBindingType()
}

func (s SubnetRouteTableInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubnetRouteTableInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Subnets struct {
    // List of subnets for which IPsec VPN is configured. Subnets should be network address specified in CIDR format and can accept '0.0.0.0/0' (any)
	Subnets []string
}

func (s Subnets) GetType__() bindings.BindingType {
	return SubnetsBindingType()
}

func (s Subnets) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Subnets._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// details of a subscription
type SubscriptionDetails struct {
    // Possible values are: 
    //
    // * SubscriptionDetails#SubscriptionDetails_STATUS_CREATED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_ACTIVATED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_FAILED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_CANCELLED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_EXPIRED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_PENDING_PROVISIONING
    // * SubscriptionDetails#SubscriptionDetails_STATUS_ORDER_SUBMITTED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_SUSPENDED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_TERMINATED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_UKNOWN
	Status *string
	AnniversaryBillingDate *string
	EndDate *string
    // The frequency at which the customer is billed. Currently supported values are \"Upfront\" and \"Monthly\"
	BillingFrequency *string
	AutoRenewedAllowed *string
	CommitmentTerm *string
	CspSubscriptionId *string
	BillingSubscriptionId *string
	OfferVersion *string
	OfferType *OfferType
	Description *string
	ProductId *string
	Region *string
	ProductName *string
	OfferName *string
    // unit of measurment for commitment term
	CommitmentTermUom *string
	StartDate *string
	Quantity *string
}
const SubscriptionDetails_STATUS_CREATED = "CREATED"
const SubscriptionDetails_STATUS_ACTIVATED = "ACTIVATED"
const SubscriptionDetails_STATUS_FAILED = "FAILED"
const SubscriptionDetails_STATUS_CANCELLED = "CANCELLED"
const SubscriptionDetails_STATUS_EXPIRED = "EXPIRED"
const SubscriptionDetails_STATUS_PENDING_PROVISIONING = "PENDING_PROVISIONING"
const SubscriptionDetails_STATUS_ORDER_SUBMITTED = "ORDER_SUBMITTED"
const SubscriptionDetails_STATUS_SUSPENDED = "SUSPENDED"
const SubscriptionDetails_STATUS_TERMINATED = "TERMINATED"
const SubscriptionDetails_STATUS_UKNOWN = "UKNOWN"

func (s SubscriptionDetails) GetType__() bindings.BindingType {
	return SubscriptionDetailsBindingType()
}

func (s SubscriptionDetails) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubscriptionDetails._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Details of products that are available for purchase.
type SubscriptionProducts struct {
    // The name of the product
	Product *string
    // A list of different types/version for the product.
	Types []string
}

func (s SubscriptionProducts) GetType__() bindings.BindingType {
	return SubscriptionProductsBindingType()
}

func (s SubscriptionProducts) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubscriptionProducts._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SubscriptionRequest struct {
    // The product for which subscription needs to be created. Refer /vmc/api/orgs/{orgId}/products.
	Product *string
    // Old identifier for product. \*Deprecarted\*. See product and type
	ProductType string
	ProductId *string
    // Frequency of the billing.
	BillingFrequency *string
	Region string
	CommitmentTerm string
	OfferContextId *string
	OfferVersion string
	OfferName string
	Quantity int64
    // The type of the product for which the subscription needs to be created.
	Type_ *string
	Price *int64
	ProductChargeId *string
}

func (s SubscriptionRequest) GetType__() bindings.BindingType {
	return SubscriptionRequestBindingType()
}

func (s SubscriptionRequest) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubscriptionRequest._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SupportWindow struct {
    // Possible values are: 
    //
    // * SupportWindow#SupportWindow_START_DAY_MONDAY
    // * SupportWindow#SupportWindow_START_DAY_TUESDAY
    // * SupportWindow#SupportWindow_START_DAY_WEDNESDAY
    // * SupportWindow#SupportWindow_START_DAY_THURSDAY
    // * SupportWindow#SupportWindow_START_DAY_FRIDAY
    // * SupportWindow#SupportWindow_START_DAY_SATURDAY
    // * SupportWindow#SupportWindow_START_DAY_SUNDAY
	StartDay *string
	Seats *int64
    // SDDCs in this window format: UUID
	Sddcs []string
	DurationHours *int64
	StartHour *int64
	SupportWindowId *string
	Metadata *data.StructValue
}
const SupportWindow_START_DAY_MONDAY = "MONDAY"
const SupportWindow_START_DAY_TUESDAY = "TUESDAY"
const SupportWindow_START_DAY_WEDNESDAY = "WEDNESDAY"
const SupportWindow_START_DAY_THURSDAY = "THURSDAY"
const SupportWindow_START_DAY_FRIDAY = "FRIDAY"
const SupportWindow_START_DAY_SATURDAY = "SATURDAY"
const SupportWindow_START_DAY_SUNDAY = "SUNDAY"

func (s SupportWindow) GetType__() bindings.BindingType {
	return SupportWindowBindingType()
}

func (s SupportWindow) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SupportWindow._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SupportWindowId struct {
    // Support Window ID
	WindowId *string
}

func (s SupportWindowId) GetType__() bindings.BindingType {
	return SupportWindowIdBindingType()
}

func (s SupportWindowId) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SupportWindowId._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Task struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
    // Possible values are: 
    //
    // * Task#Task_STATUS_STARTED
    // * Task#Task_STATUS_CANCELING
    // * Task#Task_STATUS_FINISHED
    // * Task#Task_STATUS_FAILED
    // * Task#Task_STATUS_CANCELED
	Status *string
	LocalizedErrorMessage *string
    // UUID of the resource the task is acting upon
	ResourceId *string
    // If this task was created by another task - this provides the linkage. Mostly for debugging.
	ParentTaskId *string
	TaskVersion *string
    // (Optional) Client provided uniqifier to make task creation idempotent. Be aware not all tasks support this. For tasks that do - supplying the same correlation Id, for the same task type, within a predefined window will ensure the operation happens at most once.
	CorrelationId *string
    // Entity version of the resource at the start of the task. This is only set for some task types. format: int32
	StartResourceEntityVersion *int64
	SubStatus *string
	TaskType *string
	StartTime *time.Time
    // Task progress phases involved in current task execution
	TaskProgressPhases []TaskProgressPhase
	ErrorMessage *string
	OrgId *string
    // Entity version of the resource at the end of the task. This is only set for some task types. format: int32
	EndResourceEntityVersion *int64
    // Service errors returned from SDDC services.
	ServiceErrors []ServiceError
	OrgType *string
    // Estimated remaining time in minute of the task execution, < 0 means no estimation for the task. format: int32
	EstimatedRemainingMinutes *int64
	Params *data.StructValue
    // Estimated progress percentage the task executed format: int32
	ProgressPercent *int64
    // The current in progress phase ID in the task execution, if none in progress, empty string returned.
	PhaseInProgress *string
    // Type of resource being acted upon
	ResourceType *string
	EndTime *time.Time
}
const Task_STATUS_STARTED = "STARTED"
const Task_STATUS_CANCELING = "CANCELING"
const Task_STATUS_FINISHED = "FINISHED"
const Task_STATUS_FAILED = "FAILED"
const Task_STATUS_CANCELED = "CANCELED"

func (s Task) GetType__() bindings.BindingType {
	return TaskBindingType()
}

func (s Task) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Task._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// A task progress can be (but does NOT have to be) divided to more meaningful progress phases.
type TaskProgressPhase struct {
    // The identifier of the task progress phase
	Id string
    // The display name of the task progress phase
	Name string
    // The percentage of the phase that has completed format: int32
	ProgressPercent int64
}

func (s TaskProgressPhase) GetType__() bindings.BindingType {
	return TaskProgressPhaseBindingType()
}

func (s TaskProgressPhase) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TaskProgressPhase._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Holder for term billing options.
type TermBillingOptions struct {
	UnitPrice *string
	BillingFrequency *string
}

func (s TermBillingOptions) GetType__() bindings.BindingType {
	return TermBillingOptionsBindingType()
}

func (s TermBillingOptions) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TermBillingOptions._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Holder for the term offer instances.
type TermOfferInstance struct {
	Description string
	Product string
    // Deprecated. Please use product and type fields instead.
	ProductType *string
	Name string
	Currency string
	Region string
	CommitmentTerm int64
    // (deprecated. unit_price is moved into TermBillingOptions. For backward compatibility, this field reflect \"Prepaid\" price at the offer level.)
	UnitPrice string
	BillingOptions []TermBillingOptions
	Version string
	OfferContextId *string
	ProductChargeId *string
	Type_ string
	ProductId *string
}

func (s TermOfferInstance) GetType__() bindings.BindingType {
	return TermOfferInstanceBindingType()
}

func (s TermOfferInstance) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TermOfferInstance._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type TrafficShapingPolicy struct {
	BurstSize *int64
	AverageBandwidth *int64
	PeakBandwidth *int64
	Enabled *bool
	Inherited *bool
}

func (s TrafficShapingPolicy) GetType__() bindings.BindingType {
	return TrafficShapingPolicyBindingType()
}

func (s TrafficShapingPolicy) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TrafficShapingPolicy._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type UpdateCredentials struct {
    // Username of the credentials
	Username string
    // Password associated with the credentials
	Password string
}

func (s UpdateCredentials) GetType__() bindings.BindingType {
	return UpdateCredentialsBindingType()
}

func (s UpdateCredentials) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for UpdateCredentials._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VmcLocale struct {
    // The locale to be used for translating responses for the session
	Locale *string
}

func (s VmcLocale) GetType__() bindings.BindingType {
	return VmcLocaleBindingType()
}

func (s VmcLocale) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VmcLocale._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// NSX Edge vnic configuration details.
type Vnic struct {
    // List of sub interfaces. Sub interfaces can be created only on a trunk interface.
	SubInterfaces *SubInterfaces
    // Address group configuration of the interface.
	AddressGroups *EdgeVnicAddressGroups
    // Value is true if the vnic is connected to a logical switch, standard portgroup or distributed portgroup.
	IsConnected *bool
    // Value is true if send redirects is enabled. Enable ICMP redirect to convey routing information to hosts.
	EnableSendRedirects *bool
	InShapingPolicy *TrafficShapingPolicy
    // Interface label of format vNic_{vnicIndex} provided by NSX Manager. Read only.
	Label *string
    // Value is true if proxy arp is enabled. Enable proxy ARP if you want to allow the NSX Edge of type 'gatewayServices' to answer ARP requests intended for other machines.
	EnableProxyArp *bool
    // Index of the vnic. Min value is 0 and max value is 9. format: int32
	Index int64
    // Name of the interface. Optional.
	Name *string
    // MTU of the interface, with default as 1500. Min is 68, Max is 9000. Optional. format: int32
	Mtu *int64
	FenceParameters []KeyValueAttributes
    // Distinct MAC addresses configured for the vnic. Optional.
	MacAddresses []MacAddress
	OutShapingPolicy *TrafficShapingPolicy
    // Name of the port group or logical switch.
	PortgroupName *string
    // Value is true if bridge mode is enabled.
	EnableBridgeMode *bool
    // Type of the vnic. Values are uplink, internal, trunk. At least one internal interface must be configured for NSX Edge HA to work.
	Type_ *string
    // Value are port group ID (standard portgroup or distributed portgroup) or virtual wire ID (logical switch). Logical switch cannot be used for a TRUNK vnic. Portgroup cannot be shared among vnics/LIFs. Required when isConnected is specified as true. Example 'network-17' (standard portgroup), 'dvportgroup-34' (distributed portgroup) or 'virtualwire-2' (logical switch).
	PortgroupId *string
}

func (s Vnic) GetType__() bindings.BindingType {
	return VnicBindingType()
}

func (s Vnic) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Vnic._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Ordered list of NSX Edge vnics. Until one connected vnic is configured, none of the configured features will serve the network.
type Vnics struct {
    // Ordered list of NSX Edge vnics.
	Vnics []Vnic
}

func (s Vnics) GetType__() bindings.BindingType {
	return VnicsBindingType()
}

func (s Vnics) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Vnics._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VpcInfo struct {
	VpcCidr *string
	VgwId *string
	EsxPublicSecurityGroupId *string
    // set of virtual interfaces attached to the sddc
	VifIds []string
	VmSecurityGroupId *string
    // Mapping from AZ to a list of IP addresses assigned to TGW ENI that's connected with Vpc
	TgwIps map[string][]string
    // (deprecated)
	RouteTableId *string
    // Id of the NSX edge associated with this VPC (deprecated)
	EdgeSubnetId *string
    // vpc id
	Id *string
    // Id of the association between subnet and route-table (deprecated)
	ApiAssociationId *string
    // Id associated with this VPC (deprecated)
	ApiSubnetId *string
    // (deprecated)
	PrivateSubnetId *string
    // (deprecated)
	PrivateAssociationId *string
	EsxSecurityGroupId *string
    // (deprecated)
	SubnetId *string
	InternetGatewayId *string
	SecurityGroupId *string
    // (deprecated)
	AssociationId *string
    // Route table which contains the route to VGW (deprecated)
	VgwRouteTableId *string
    // Id of the association between edge subnet and route-table (deprecated)
	EdgeAssociationId *string
	Provider *string
    // (deprecated)
	PeeringConnectionId *string
	NetworkType *string
	AvailableZones []AvailableZoneInfo
    // map from routeTableName to routeTableInfo
	Routetables map[string]RouteTableInfo
}

func (s VpcInfo) GetType__() bindings.BindingType {
	return VpcInfoBindingType()
}

func (s VpcInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VpcInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VpcInfoSubnets struct {
    // The ID of the VPC these subnets belong to.
	VpcId *string
    // The overall CIDR block of the VPC. This is the AWS primary CIDR block.
	CidrBlock *string
    // The description of the VPC; usually it's name or id.
	Description *string
	Subnets []SubnetInfo
}

func (s VpcInfoSubnets) GetType__() bindings.BindingType {
	return VpcInfoSubnetsBindingType()
}

func (s VpcInfoSubnets) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VpcInfoSubnets._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Vpn struct {
	Version *string
	OnPremGatewayIp *string
	OnPremNetworkCidr *string
	PfsEnabled *bool
	Id *string
	ChannelStatus *VpnChannelStatus
	OnPremNatIp *string
	Name *string
	InternalNetworkIds []string
	TunnelStatuses []VpnTunnelStatus
    // Possible values are: 
    //
    // * Vpn#Vpn_ENCRYPTION_AES
    // * Vpn#Vpn_ENCRYPTION_AES256
    // * Vpn#Vpn_ENCRYPTION_AES_GCM
    // * Vpn#Vpn_ENCRYPTION_TRIPLE_DES
    // * Vpn#Vpn_ENCRYPTION_UNKNOWN
	Encryption *string
	Enabled *bool
    // Possible values are: 
    //
    // * Vpn#Vpn_STATE_CONNECTED
    // * Vpn#Vpn_STATE_DISCONNECTED
    // * Vpn#Vpn_STATE_PARTIALLY_CONNECTED
    // * Vpn#Vpn_STATE_UNKNOWN
	State *string
    // Possible values are: 
    //
    // * Vpn#Vpn_DH_GROUP_DH2
    // * Vpn#Vpn_DH_GROUP_DH5
    // * Vpn#Vpn_DH_GROUP_DH14
    // * Vpn#Vpn_DH_GROUP_DH15
    // * Vpn#Vpn_DH_GROUP_DH16
    // * Vpn#Vpn_DH_GROUP_UNKNOWN
	DhGroup *string
    // Possible values are: 
    //
    // * Vpn#Vpn_AUTHENTICATION_PSK
    // * Vpn#Vpn_AUTHENTICATION_UNKNOWN
	Authentication *string
	PreSharedKey *string
    // Possible values are: 
    //
    // * Vpn#Vpn_IKE_OPTION_IKEV1
    // * Vpn#Vpn_IKE_OPTION_IKEV2
	IkeOption *string
    // Possible values are: 
    //
    // * Vpn#Vpn_DIGEST_ALGORITHM_SHA1
    // * Vpn#Vpn_DIGEST_ALGORITHM_SHA_256
	DigestAlgorithm *string
}
const Vpn_ENCRYPTION_AES = "AES"
const Vpn_ENCRYPTION_AES256 = "AES256"
const Vpn_ENCRYPTION_AES_GCM = "AES_GCM"
const Vpn_ENCRYPTION_TRIPLE_DES = "TRIPLE_DES"
const Vpn_ENCRYPTION_UNKNOWN = "UNKNOWN"
const Vpn_STATE_CONNECTED = "CONNECTED"
const Vpn_STATE_DISCONNECTED = "DISCONNECTED"
const Vpn_STATE_PARTIALLY_CONNECTED = "PARTIALLY_CONNECTED"
const Vpn_STATE_UNKNOWN = "UNKNOWN"
const Vpn_DH_GROUP_DH2 = "DH2"
const Vpn_DH_GROUP_DH5 = "DH5"
const Vpn_DH_GROUP_DH14 = "DH14"
const Vpn_DH_GROUP_DH15 = "DH15"
const Vpn_DH_GROUP_DH16 = "DH16"
const Vpn_DH_GROUP_UNKNOWN = "UNKNOWN"
const Vpn_AUTHENTICATION_PSK = "PSK"
const Vpn_AUTHENTICATION_UNKNOWN = "UNKNOWN"
const Vpn_IKE_OPTION_IKEV1 = "IKEV1"
const Vpn_IKE_OPTION_IKEV2 = "IKEV2"
const Vpn_DIGEST_ALGORITHM_SHA1 = "SHA1"
const Vpn_DIGEST_ALGORITHM_SHA_256 = "SHA_256"

func (s Vpn) GetType__() bindings.BindingType {
	return VpnBindingType()
}

func (s Vpn) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Vpn._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VpnChannelStatus struct {
    // Possible values are: 
    //
    // * VpnChannelStatus#VpnChannelStatus_CHANNEL_STATUS_CONNECTED
    // * VpnChannelStatus#VpnChannelStatus_CHANNEL_STATUS_DISCONNECTED
    // * VpnChannelStatus#VpnChannelStatus_CHANNEL_STATUS_UNKNOWN
	ChannelStatus *string
	ChannelState *string
	LastInfoMessage *string
	FailureMessage *string
}
const VpnChannelStatus_CHANNEL_STATUS_CONNECTED = "CONNECTED"
const VpnChannelStatus_CHANNEL_STATUS_DISCONNECTED = "DISCONNECTED"
const VpnChannelStatus_CHANNEL_STATUS_UNKNOWN = "UNKNOWN"

func (s VpnChannelStatus) GetType__() bindings.BindingType {
	return VpnChannelStatusBindingType()
}

func (s VpnChannelStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VpnChannelStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VpnTunnelStatus struct {
	OnPremSubnet *string
	TrafficStats *VpnTunnelTrafficStats
	LastInfoMessage *string
	LocalSubnet *string
	TunnelState *string
	FailureMessage *string
    // Possible values are: 
    //
    // * VpnTunnelStatus#VpnTunnelStatus_TUNNEL_STATUS_CONNECTED
    // * VpnTunnelStatus#VpnTunnelStatus_TUNNEL_STATUS_DISCONNECTED
    // * VpnTunnelStatus#VpnTunnelStatus_TUNNEL_STATUS_UNKNOWN
	TunnelStatus *string
}
const VpnTunnelStatus_TUNNEL_STATUS_CONNECTED = "CONNECTED"
const VpnTunnelStatus_TUNNEL_STATUS_DISCONNECTED = "DISCONNECTED"
const VpnTunnelStatus_TUNNEL_STATUS_UNKNOWN = "UNKNOWN"

func (s VpnTunnelStatus) GetType__() bindings.BindingType {
	return VpnTunnelStatusBindingType()
}

func (s VpnTunnelStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VpnTunnelStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VpnTunnelTrafficStats struct {
	PacketsOut *string
	PacketReceivedErrors *string
	RxBytesOnLocalSubnet *string
	ReplayErrors *string
	SequenceNumberOverFlowErrors *string
	EncryptionFailures *string
	IntegrityErrors *string
	PacketSentErrors *string
	DecryptionFailures *string
	PacketsIn *string
	TxBytesFromLocalSubnet *string
}

func (s VpnTunnelTrafficStats) GetType__() bindings.BindingType {
	return VpnTunnelTrafficStatsBindingType()
}

func (s VpnTunnelTrafficStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VpnTunnelTrafficStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Infomation about an available vSAN capacity in a cluster.
type VsanAvailableCapacity struct {
	Cost string
	Quality string
	Size int64
}

func (s VsanAvailableCapacity) GetType__() bindings.BindingType {
	return VsanAvailableCapacityBindingType()
}

func (s VsanAvailableCapacity) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VsanAvailableCapacity._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Bias for reconfiguring vSAN in a cluster.
type VsanClusterReconfigBias struct {
	ShortDescription string
	FullDescription string
	Id string
}

func (s VsanClusterReconfigBias) GetType__() bindings.BindingType {
	return VsanClusterReconfigBiasBindingType()
}

func (s VsanClusterReconfigBias) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VsanClusterReconfigBias._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Storage constraint information for reconfiguring vSAN in existing cluster.
type VsanClusterReconfigConstraints struct {
    // Biases to reconfigure vSAN in an existing cluster.
	ReconfigBiases []VsanClusterReconfigBias
    // A map of VsanClusterReconfigBias id to the list of VsanAvailableCapacity. It gives all of available vSAN capacities for each of reconfiguration biases.
	AvailableCapacities map[string][]VsanAvailableCapacity
    // A map of VsanClusterReconfigBias id to a VsanAvailableCapacity. It gives the default VsanAvailableCapacity for each of reconfiguration biases.
	DefaultCapacities map[string]VsanAvailableCapacity
    // The number of hosts in a cluster for the constraints. format: int32
	Hosts int64
    // The id of default VsanClusterReconfigBias for this constraints.
	DefaultReconfigBiasId string
}

func (s VsanClusterReconfigConstraints) GetType__() bindings.BindingType {
	return VsanClusterReconfigConstraintsBindingType()
}

func (s VsanClusterReconfigConstraints) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VsanClusterReconfigConstraints._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// This describes the possible physical storage capacity choices for use with a given VsanStorageDesigner implementation. These choices are specific to a customer-defined number of hosts per cluster.
type VsanConfigConstraints struct {
    // Maximum capacity supported for cluster (GiB). format: int64
	MaxCapacity int64
    // List of supported capacities which may offer preferable performance (GiB). format: int64
	RecommendedCapacities []int64
    // Increment to be added to min_capacity to result in a supported capacity (GiB). format: int64
	SupportedCapacityIncrement *int64
    // Minimum capacity supported for cluster (GiB). format: int64
	MinCapacity int64
    // Number of hosts in cluster. format: int64
	NumHosts int64
}

func (s VsanConfigConstraints) GetType__() bindings.BindingType {
	return VsanConfigConstraintsBindingType()
}

func (s VsanConfigConstraints) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VsanConfigConstraints._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VsanEncryptionConfig struct {
    // Port to connect to AWS Key Management Service
	Port *int64
    // Public certificate used to connect to AWS Key Management Service
	Certificate *string
}

func (s VsanEncryptionConfig) GetType__() bindings.BindingType {
	return VsanEncryptionConfigBindingType()
}

func (s VsanEncryptionConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VsanEncryptionConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}





func AbstractEntityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.abstract_entity", fields, reflect.TypeOf(AbstractEntity{}), fieldNameMap, validators)
}

func AccountLinkConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["delay_account_link"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["delay_account_link"] = "DelayAccountLink"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.account_link_config", fields, reflect.TypeOf(AccountLinkConfig{}), fieldNameMap, validators)
}

func AccountLinkSddcConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customer_subnet_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["customer_subnet_ids"] = "CustomerSubnetIds"
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.account_link_sddc_config", fields, reflect.TypeOf(AccountLinkSddcConfig{}), fieldNameMap, validators)
}

func AddressFWSourceDestinationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["exclude"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["exclude"] = "Exclude"
	fields["ipAddress"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ipAddress"] = "IpAddress"
	fields["groupingObjectId"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["groupingObjectId"] = "GroupingObjectId"
	fields["vnicGroupId"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vnicGroupId"] = "VnicGroupId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.address_FW_source_destination", fields, reflect.TypeOf(AddressFWSourceDestination{}), fieldNameMap, validators)
}

func AgentBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["internal_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_ip"] = "InternalIp"
	fields["agent_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["agent_url"] = "AgentUrl"
	fields["management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_ip"] = "ManagementIp"
	fields["hostname_verifier_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["hostname_verifier_enabled"] = "HostnameVerifierEnabled"
	fields["master"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["master"] = "Master"
	fields["network_netmask"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_netmask"] = "NetworkNetmask"
	fields["network_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_gateway"] = "NetworkGateway"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["cert_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cert_enabled"] = "CertEnabled"
	fields["agent_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["agent_state"] = "AgentState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.agent", fields, reflect.TypeOf(Agent{}), fieldNameMap, validators)
}

func AmiInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ami_info", fields, reflect.TypeOf(AmiInfo{}), fieldNameMap, validators)
}

func AppliancesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dataStoreMoidOfActiveVse"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dataStoreMoidOfActiveVse"] = "DataStoreMoidOfActiveVse"
	fields["enableFips"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enableFips"] = "EnableFips"
	fields["hostNameOfActiveVse"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["hostNameOfActiveVse"] = "HostNameOfActiveVse"
	fields["vmBuildInfo"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vmBuildInfo"] = "VmBuildInfo"
	fields["deployAppliances"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["deployAppliances"] = "DeployAppliances"
	fields["communicationChannel"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["communicationChannel"] = "CommunicationChannel"
	fields["vmNameOfActiveVse"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vmNameOfActiveVse"] = "VmNameOfActiveVse"
	fields["numberOfDeployedVms"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["numberOfDeployedVms"] = "NumberOfDeployedVms"
	fields["resourcePoolMoidOfActiveVse"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resourcePoolMoidOfActiveVse"] = "ResourcePoolMoidOfActiveVse"
	fields["dataStoreNameOfActiveVse"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dataStoreNameOfActiveVse"] = "DataStoreNameOfActiveVse"
	fields["vmMoidOfActiveVse"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vmMoidOfActiveVse"] = "VmMoidOfActiveVse"
	fields["statusFromVseUpdatedOn"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["statusFromVseUpdatedOn"] = "StatusFromVseUpdatedOn"
	fields["fqdn"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["fqdn"] = "Fqdn"
	fields["applianceSize"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["applianceSize"] = "ApplianceSize"
	fields["resourcePoolNameOfActiveVse"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resourcePoolNameOfActiveVse"] = "ResourcePoolNameOfActiveVse"
	fields["activeVseHaIndex"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["activeVseHaIndex"] = "ActiveVseHaIndex"
	fields["vmVersion"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vmVersion"] = "VmVersion"
	fields["hostMoidOfActiveVse"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["hostMoidOfActiveVse"] = "HostMoidOfActiveVse"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.appliances_summary", fields, reflect.TypeOf(AppliancesSummary{}), fieldNameMap, validators)
}

func ApplicationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["applicationId"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["applicationId"] = "ApplicationId"
	fields["service"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NsxfirewallserviceBindingType), reflect.TypeOf([]Nsxfirewallservice{})))
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.application", fields, reflect.TypeOf(Application{}), fieldNameMap, validators)
}

func AvailableZoneInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SubnetBindingType), reflect.TypeOf([]Subnet{})))
	fieldNameMap["subnets"] = "Subnets"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.available_zone_info", fields, reflect.TypeOf(AvailableZoneInfo{}), fieldNameMap, validators)
}

func AwsAgentBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_id"] = "InstanceId"
	fields["key_pair"] = bindings.NewOptionalType(bindings.NewReferenceType(AwsKeyPairBindingType))
	fieldNameMap["key_pair"] = "KeyPair"
	fields["internal_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_ip"] = "InternalIp"
	fields["agent_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["agent_url"] = "AgentUrl"
	fields["management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_ip"] = "ManagementIp"
	fields["hostname_verifier_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["hostname_verifier_enabled"] = "HostnameVerifierEnabled"
	fields["master"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["master"] = "Master"
	fields["network_netmask"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_netmask"] = "NetworkNetmask"
	fields["network_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_gateway"] = "NetworkGateway"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["cert_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cert_enabled"] = "CertEnabled"
	fields["agent_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["agent_state"] = "AgentState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_agent", fields, reflect.TypeOf(AwsAgent{}), fieldNameMap, validators)
}

func AwsCloudProviderBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["regions"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["regions"] = "Regions"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_cloud_provider", fields, reflect.TypeOf(AwsCloudProvider{}), fieldNameMap, validators)
}

func AwsCompatibleSubnetsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customer_available_zones"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["customer_available_zones"] = "CustomerAvailableZones"
	fields["vpc_map"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(VpcInfoSubnetsBindingType),reflect.TypeOf(map[string]VpcInfoSubnets{})))
	fieldNameMap["vpc_map"] = "VpcMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_compatible_subnets", fields, reflect.TypeOf(AwsCompatibleSubnets{}), fieldNameMap, validators)
}

func AwsCustomerConnectedAccountBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["policy_payer_arn"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["policy_payer_arn"] = "PolicyPayerArn"
	fields["region_to_az_to_shadow_mapping"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})),reflect.TypeOf(map[string]map[string]string{})))
	fieldNameMap["region_to_az_to_shadow_mapping"] = "RegionToAzToShadowMapping"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["cf_stack_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cf_stack_name"] = "CfStackName"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["account_number"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["account_number"] = "AccountNumber"
	fields["policy_service_arn"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["policy_service_arn"] = "PolicyServiceArn"
	fields["policy_external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["policy_external_id"] = "PolicyExternalId"
	fields["policy_payer_linked_arn"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["policy_payer_linked_arn"] = "PolicyPayerLinkedArn"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_customer_connected_account", fields, reflect.TypeOf(AwsCustomerConnectedAccount{}), fieldNameMap, validators)
}

func AwsEsxHostBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["internal_public_ip_pool"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SddcPublicIpBindingType), reflect.TypeOf([]SddcPublicIp{})))
	fieldNameMap["internal_public_ip_pool"] = "InternalPublicIpPool"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["esx_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_id"] = "EsxId"
	fields["hostname"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["custom_properties"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["esx_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_state"] = "EsxState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_esx_host", fields, reflect.TypeOf(AwsEsxHost{}), fieldNameMap, validators)
}

func AwsKeyPairBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["key_name"] = "KeyName"
	fields["key_fingerprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["key_fingerprint"] = "KeyFingerprint"
	fields["key_material"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["key_material"] = "KeyMaterial"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_key_pair", fields, reflect.TypeOf(AwsKeyPair{}), fieldNameMap, validators)
}

func AwsKmsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["amazon_resource_name"] = bindings.NewStringType()
	fieldNameMap["amazon_resource_name"] = "AmazonResourceName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_kms_info", fields, reflect.TypeOf(AwsKmsInfo{}), fieldNameMap, validators)
}

func AwsSddcConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["region"] = bindings.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["vpc_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_cidr"] = "VpcCidr"
	fields["host_instance_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vmc.model.host_instance_types", reflect.TypeOf(HostInstanceTypes(HostInstanceTypes_I3_METAL))))
	fieldNameMap["host_instance_type"] = "HostInstanceType"
	fields["skip_creating_vxlan"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["vxlan_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["size"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["size"] = "Size"
	fields["storage_capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["account_link_sddc_config"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AccountLinkSddcConfigBindingType), reflect.TypeOf([]AccountLinkSddcConfig{})))
	fieldNameMap["account_link_sddc_config"] = "AccountLinkSddcConfig"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	fields["sddc_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_type"] = "SddcType"
	fields["account_link_config"] = bindings.NewOptionalType(bindings.NewReferenceType(AccountLinkConfigBindingType))
	fieldNameMap["account_link_config"] = "AccountLinkConfig"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["sso_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["sddc_template_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_template_id"] = "SddcTemplateId"
	fields["deployment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_sddc_config", fields, reflect.TypeOf(AwsSddcConfig{}), fieldNameMap, validators)
}

func AwsSddcConnectionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["cidr_block_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cidr_block_subnet"] = "CidrBlockSubnet"
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["eni_group"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["eni_group"] = "EniGroup"
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["cgw_present"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cgw_present"] = "CgwPresent"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["cidr_block_vpc"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cidr_block_vpc"] = "CidrBlockVpc"
	fields["connection_order"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["connection_order"] = "ConnectionOrder"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["subnet_availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_availability_zone"] = "SubnetAvailabilityZone"
	fields["vpc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["customer_eni_infos"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CustomerEniInfoBindingType), reflect.TypeOf([]CustomerEniInfo{})))
	fieldNameMap["customer_eni_infos"] = "CustomerEniInfos"
	fields["default_route_table"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["default_route_table"] = "DefaultRouteTable"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_sddc_connection", fields, reflect.TypeOf(AwsSddcConnection{}), fieldNameMap, validators)
}

func AwsSddcResourceConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backup_restore_bucket"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["backup_restore_bucket"] = "BackupRestoreBucket"
	fields["public_ip_pool"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SddcPublicIpBindingType), reflect.TypeOf([]SddcPublicIp{})))
	fieldNameMap["public_ip_pool"] = "PublicIpPool"
	fields["vpc_info"] = bindings.NewOptionalType(bindings.NewReferenceType(VpcInfoBindingType))
	fieldNameMap["vpc_info"] = "VpcInfo"
	fields["kms_vpc_endpoint"] = bindings.NewOptionalType(bindings.NewReferenceType(KmsVpcEndpointBindingType))
	fieldNameMap["kms_vpc_endpoint"] = "KmsVpcEndpoint"
	fields["max_num_public_ip"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["max_num_public_ip"] = "MaxNumPublicIp"
	fields["account_link_sddc_config"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SddcLinkConfigBindingType), reflect.TypeOf([]SddcLinkConfig{})))
	fieldNameMap["account_link_sddc_config"] = "AccountLinkSddcConfig"
	fields["vsan_encryption_config"] = bindings.NewOptionalType(bindings.NewReferenceType(VsanEncryptionConfigBindingType))
	fieldNameMap["vsan_encryption_config"] = "VsanEncryptionConfig"
	fields["vpc_info_peered_agent"] = bindings.NewOptionalType(bindings.NewReferenceType(VpcInfoBindingType))
	fieldNameMap["vpc_info_peered_agent"] = "VpcInfoPeeredAgent"
	fields["mgmt_appliance_network_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mgmt_appliance_network_name"] = "MgmtApplianceNetworkName"
	fields["nsxt"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["nsxt"] = "Nsxt"
	fields["mgw_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mgw_id"] = "MgwId"
	fields["nsx_mgr_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_mgr_url"] = "NsxMgrUrl"
	fields["psc_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["psc_management_ip"] = "PscManagementIp"
	fields["psc_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["psc_url"] = "PscUrl"
	fields["cgws"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["cgws"] = "Cgws"
	fields["availability_zones"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["availability_zones"] = "AvailabilityZones"
	fields["management_ds"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_ds"] = "ManagementDs"
	fields["nsx_api_public_endpoint_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_api_public_endpoint_url"] = "NsxApiPublicEndpointUrl"
	fields["custom_properties"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["cloud_password"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_password"] = "CloudPassword"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["clusters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ClusterBindingType), reflect.TypeOf([]Cluster{})))
	fieldNameMap["clusters"] = "Clusters"
	fields["vc_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_management_ip"] = "VcManagementIp"
	fields["sddc_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_networks"] = "SddcNetworks"
	fields["cloud_username"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_username"] = "CloudUsername"
	fields["esx_hosts"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AwsEsxHostBindingType), reflect.TypeOf([]AwsEsxHost{})))
	fieldNameMap["esx_hosts"] = "EsxHosts"
	fields["nsx_mgr_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_mgr_management_ip"] = "NsxMgrManagementIp"
	fields["vc_instance_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_instance_id"] = "VcInstanceId"
	fields["esx_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_cluster_id"] = "EsxClusterId"
	fields["vc_public_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_public_ip"] = "VcPublicIp"
	fields["skip_creating_vxlan"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["vc_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_url"] = "VcUrl"
	fields["sddc_manifest"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcManifestBindingType))
	fieldNameMap["sddc_manifest"] = "SddcManifest"
	fields["vxlan_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["cloud_user_group"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_user_group"] = "CloudUserGroup"
	fields["management_rp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_rp"] = "ManagementRp"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["witness_availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["witness_availability_zone"] = "WitnessAvailabilityZone"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["pop_agent_xeni_connection"] = bindings.NewOptionalType(bindings.NewReferenceType(PopAgentXeniConnectionBindingType))
	fieldNameMap["pop_agent_xeni_connection"] = "PopAgentXeniConnection"
	fields["sddc_size"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcSizeBindingType))
	fieldNameMap["sddc_size"] = "SddcSize"
	fields["nsx_controller_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["nsx_controller_ips"] = "NsxControllerIps"
	fields["esx_host_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_host_subnet"] = "EsxHostSubnet"
	fields["sso_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["deployment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	fields["nsxt_addons"] = bindings.NewOptionalType(bindings.NewReferenceType(NsxtAddonsBindingType))
	fieldNameMap["nsxt_addons"] = "NsxtAddons"
	fields["dns_with_management_vm_private_ip"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["dns_with_management_vm_private_ip"] = "DnsWithManagementVmPrivateIp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_sddc_resource_config", fields, reflect.TypeOf(AwsSddcResourceConfig{}), fieldNameMap, validators)
}

func AwsSubnetBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["region_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region_name"] = "RegionName"
	fields["availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["subnet_cidr_block"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_cidr_block"] = "SubnetCidrBlock"
	fields["is_compatible"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["is_compatible"] = "IsCompatible"
	fields["vpc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["vpc_cidr_block"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_cidr_block"] = "VpcCidrBlock"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_subnet", fields, reflect.TypeOf(AwsSubnet{}), fieldNameMap, validators)
}

func CaCertificatesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["caCertificate"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["caCertificate"] = "CaCertificate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ca_certificates", fields, reflect.TypeOf(CaCertificates{}), fieldNameMap, validators)
}

func CbmStatisticBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vnic"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["vnic"] = "Vnic"
	fields["timestamp"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["timestamp"] = "Timestamp"
	fields["out"] = bindings.NewOptionalType(bindings.NewDoubleType())
	fieldNameMap["out"] = "Out"
	fields["in"] = bindings.NewOptionalType(bindings.NewDoubleType())
	fieldNameMap["in"] = "In"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.cbm_statistic", fields, reflect.TypeOf(CbmStatistic{}), fieldNameMap, validators)
}

func CbmStatisticsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dataDto"] = bindings.NewOptionalType(bindings.NewReferenceType(CbmStatsDataBindingType))
	fieldNameMap["dataDto"] = "DataDto"
	fields["metaDto"] = bindings.NewOptionalType(bindings.NewReferenceType(MetaDashboardStatsBindingType))
	fieldNameMap["metaDto"] = "MetaDto"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.cbm_statistics", fields, reflect.TypeOf(CbmStatistics{}), fieldNameMap, validators)
}

func CbmStatsDataBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vnic_9"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CbmStatisticBindingType), reflect.TypeOf([]CbmStatistic{})))
	fieldNameMap["vnic_9"] = "Vnic9"
	fields["vnic_8"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CbmStatisticBindingType), reflect.TypeOf([]CbmStatistic{})))
	fieldNameMap["vnic_8"] = "Vnic8"
	fields["vnic_7"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CbmStatisticBindingType), reflect.TypeOf([]CbmStatistic{})))
	fieldNameMap["vnic_7"] = "Vnic7"
	fields["vnic_6"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CbmStatisticBindingType), reflect.TypeOf([]CbmStatistic{})))
	fieldNameMap["vnic_6"] = "Vnic6"
	fields["vnic_5"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CbmStatisticBindingType), reflect.TypeOf([]CbmStatistic{})))
	fieldNameMap["vnic_5"] = "Vnic5"
	fields["vnic_4"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CbmStatisticBindingType), reflect.TypeOf([]CbmStatistic{})))
	fieldNameMap["vnic_4"] = "Vnic4"
	fields["vnic_3"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CbmStatisticBindingType), reflect.TypeOf([]CbmStatistic{})))
	fieldNameMap["vnic_3"] = "Vnic3"
	fields["vnic_2"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CbmStatisticBindingType), reflect.TypeOf([]CbmStatistic{})))
	fieldNameMap["vnic_2"] = "Vnic2"
	fields["vnic_1"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CbmStatisticBindingType), reflect.TypeOf([]CbmStatistic{})))
	fieldNameMap["vnic_1"] = "Vnic1"
	fields["vnic_0"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CbmStatisticBindingType), reflect.TypeOf([]CbmStatistic{})))
	fieldNameMap["vnic_0"] = "Vnic0"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.cbm_stats_data", fields, reflect.TypeOf(CbmStatsData{}), fieldNameMap, validators)
}

func CloudProviderBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.cloud_provider", fields, reflect.TypeOf(CloudProvider{}), fieldNameMap, validators)
}

func ClusterBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["esx_host_list"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AwsEsxHostBindingType), reflect.TypeOf([]AwsEsxHost{})))
	fieldNameMap["esx_host_list"] = "EsxHostList"
	fields["cluster_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster_state"] = "ClusterState"
	fields["aws_kms_info"] = bindings.NewOptionalType(bindings.NewReferenceType(AwsKmsInfoBindingType))
	fieldNameMap["aws_kms_info"] = "AwsKmsInfo"
	fields["esx_host_info"] = bindings.NewOptionalType(bindings.NewReferenceType(EsxHostInfoBindingType))
	fieldNameMap["esx_host_info"] = "EsxHostInfo"
	fields["host_cpu_cores_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["host_cpu_cores_count"] = "HostCpuCoresCount"
	fields["cluster_capacity"] = bindings.NewOptionalType(bindings.NewReferenceType(EntityCapacityBindingType))
	fieldNameMap["cluster_capacity"] = "ClusterCapacity"
	fields["cluster_id"] = bindings.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["cluster_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster_name"] = "ClusterName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.cluster", fields, reflect.TypeOf(Cluster{}), fieldNameMap, validators)
}

func ClusterConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host_cpu_cores_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["host_cpu_cores_count"] = "HostCpuCoresCount"
	fields["host_instance_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vmc.model.host_instance_types", reflect.TypeOf(HostInstanceTypes(HostInstanceTypes_I3_METAL))))
	fieldNameMap["host_instance_type"] = "HostInstanceType"
	fields["storage_capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.cluster_config", fields, reflect.TypeOf(ClusterConfig{}), fieldNameMap, validators)
}

func ClusterReconfigureParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["storage_capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["bias"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["bias"] = "Bias"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.cluster_reconfigure_params", fields, reflect.TypeOf(ClusterReconfigureParams{}), fieldNameMap, validators)
}

func ComputeGatewayTemplateBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcPublicIpBindingType))
	fieldNameMap["public_ip"] = "PublicIp"
	fields["primary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["primary_dns"] = "PrimaryDns"
	fields["secondary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["secondary_dns"] = "SecondaryDns"
	fields["firewall_rules"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(FirewallRuleBindingType), reflect.TypeOf([]FirewallRule{})))
	fieldNameMap["firewall_rules"] = "FirewallRules"
	fields["vpns"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VpnBindingType), reflect.TypeOf([]Vpn{})))
	fieldNameMap["vpns"] = "Vpns"
	fields["logical_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(LogicalNetworkBindingType), reflect.TypeOf([]LogicalNetwork{})))
	fieldNameMap["logical_networks"] = "LogicalNetworks"
	fields["nat_rules"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NatRuleBindingType), reflect.TypeOf([]NatRule{})))
	fieldNameMap["nat_rules"] = "NatRules"
	fields["l2_vpn"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.REST))
	fieldNameMap["l2_vpn"] = "L2Vpn"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.compute_gateway_template", fields, reflect.TypeOf(ComputeGatewayTemplate{}), fieldNameMap, validators)
}

func ConfigSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["expiry_in_days"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["expiry_in_days"] = "ExpiryInDays"
	fields["availability"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewReferenceType(InstanceTypeConfigBindingType), reflect.TypeOf([]InstanceTypeConfig{})),reflect.TypeOf(map[string][]InstanceTypeConfig{})))
	fieldNameMap["availability"] = "Availability"
	fields["sddc_sizes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_sizes"] = "SddcSizes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.config_spec", fields, reflect.TypeOf(ConfigSpec{}), fieldNameMap, validators)
}

func ConnectivityAgentValidationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source"] = "Source"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["ports"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ports"] = "Ports"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.connectivity_agent_validation", fields, reflect.TypeOf(ConnectivityAgentValidation{}), fieldNameMap, validators)
}

func ConnectivityValidationGroupBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["sub_groups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ConnectivityValidationSubGroupBindingType), reflect.TypeOf([]ConnectivityValidationSubGroup{})))
	fieldNameMap["sub_groups"] = "SubGroups"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.connectivity_validation_group", fields, reflect.TypeOf(ConnectivityValidationGroup{}), fieldNameMap, validators)
}

func ConnectivityValidationGroupsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["groups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ConnectivityValidationGroupBindingType), reflect.TypeOf([]ConnectivityValidationGroup{})))
	fieldNameMap["groups"] = "Groups"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.connectivity_validation_groups", fields, reflect.TypeOf(ConnectivityValidationGroups{}), fieldNameMap, validators)
}

func ConnectivityValidationInputBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["value"] = "Value"
	fields["label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["label"] = "Label"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.connectivity_validation_input", fields, reflect.TypeOf(ConnectivityValidationInput{}), fieldNameMap, validators)
}

func ConnectivityValidationSubGroupBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["inputs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ConnectivityValidationInputBindingType), reflect.TypeOf([]ConnectivityValidationInput{})))
	fieldNameMap["inputs"] = "Inputs"
	fields["tests"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ConnectivityAgentValidationBindingType), reflect.TypeOf([]ConnectivityAgentValidation{})))
	fieldNameMap["tests"] = "Tests"
	fields["label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["label"] = "Label"
	fields["help"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["help"] = "Help"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.connectivity_validation_sub_group", fields, reflect.TypeOf(ConnectivityValidationSubGroup{}), fieldNameMap, validators)
}

func CrlCertificatesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["crlCertificate"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["crlCertificate"] = "CrlCertificate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.crl_certificates", fields, reflect.TypeOf(CrlCertificates{}), fieldNameMap, validators)
}

func CustomerEniInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["secondary_ip_addresses"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["secondary_ip_addresses"] = "SecondaryIpAddresses"
	fields["eni_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["eni_id"] = "EniId"
	fields["primary_ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["primary_ip_address"] = "PrimaryIpAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.customer_eni_info", fields, reflect.TypeOf(CustomerEniInfo{}), fieldNameMap, validators)
}

func DashboardDataBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewall"] = bindings.NewOptionalType(bindings.NewReferenceType(FirewallDashboardStatsBindingType))
	fieldNameMap["firewall"] = "Firewall"
	fields["sslvpn"] = bindings.NewOptionalType(bindings.NewReferenceType(SslvpnDashboardStatsBindingType))
	fieldNameMap["sslvpn"] = "Sslvpn"
	fields["interfaces"] = bindings.NewOptionalType(bindings.NewReferenceType(InterfacesDashboardStatsBindingType))
	fieldNameMap["interfaces"] = "Interfaces"
	fields["loadBalancer"] = bindings.NewOptionalType(bindings.NewReferenceType(LoadBalancerDashboardStatsBindingType))
	fieldNameMap["loadBalancer"] = "LoadBalancer"
	fields["ipsec"] = bindings.NewOptionalType(bindings.NewReferenceType(IpsecDashboardStatsBindingType))
	fieldNameMap["ipsec"] = "Ipsec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dashboard_data", fields, reflect.TypeOf(DashboardData{}), fieldNameMap, validators)
}

func DashboardStatBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timestamp"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["timestamp"] = "Timestamp"
	fields["value"] = bindings.NewOptionalType(bindings.NewDoubleType())
	fieldNameMap["value"] = "Value"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dashboard_stat", fields, reflect.TypeOf(DashboardStat{}), fieldNameMap, validators)
}

func DashboardStatisticsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dataDto"] = bindings.NewOptionalType(bindings.NewReferenceType(DashboardDataBindingType))
	fieldNameMap["dataDto"] = "DataDto"
	fields["metaDto"] = bindings.NewOptionalType(bindings.NewReferenceType(MetaDashboardStatsBindingType))
	fieldNameMap["metaDto"] = "MetaDto"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dashboard_statistics", fields, reflect.TypeOf(DashboardStatistics{}), fieldNameMap, validators)
}

func DataPageEdgeSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["pagingInfo"] = bindings.NewOptionalType(bindings.NewReferenceType(PagingInfoBindingType))
	fieldNameMap["pagingInfo"] = "PagingInfo"
	fields["data"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(EdgeSummaryBindingType), reflect.TypeOf([]EdgeSummary{})))
	fieldNameMap["data"] = "Data"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.data_page_edge_summary", fields, reflect.TypeOf(DataPageEdgeSummary{}), fieldNameMap, validators)
}

func DataPageSddcNetworkBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["pagingInfo"] = bindings.NewOptionalType(bindings.NewReferenceType(PagingInfoBindingType))
	fieldNameMap["pagingInfo"] = "PagingInfo"
	fields["data"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SddcNetworkBindingType), reflect.TypeOf([]SddcNetwork{})))
	fieldNameMap["data"] = "Data"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.data_page_sddc_network", fields, reflect.TypeOf(DataPageSddcNetwork{}), fieldNameMap, validators)
}

func DataPermissionsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["savePermission"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["savePermission"] = "SavePermission"
	fields["publishPermission"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["publishPermission"] = "PublishPermission"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.data_permissions", fields, reflect.TypeOf(DataPermissions{}), fieldNameMap, validators)
}

func DhcpLeaseInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hostLeaseInfoDtos"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(HostLeaseInfoBindingType), reflect.TypeOf([]HostLeaseInfo{})))
	fieldNameMap["hostLeaseInfoDtos"] = "HostLeaseInfoDtos"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dhcp_lease_info", fields, reflect.TypeOf(DhcpLeaseInfo{}), fieldNameMap, validators)
}

func DhcpLeasesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timeStamp"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["timeStamp"] = "TimeStamp"
	fields["hostLeaseInfosDto"] = bindings.NewOptionalType(bindings.NewReferenceType(DhcpLeaseInfoBindingType))
	fieldNameMap["hostLeaseInfosDto"] = "HostLeaseInfosDto"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dhcp_leases", fields, reflect.TypeOf(DhcpLeases{}), fieldNameMap, validators)
}

func DnsConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["featureType"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["featureType"] = "FeatureType"
	fields["logging"] = bindings.NewOptionalType(bindings.NewReferenceType(LoggingBindingType))
	fieldNameMap["logging"] = "Logging"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["dnsViews"] = bindings.NewOptionalType(bindings.NewReferenceType(DnsViewsBindingType))
	fieldNameMap["dnsViews"] = "DnsViews"
	fields["listeners"] = bindings.NewOptionalType(bindings.NewReferenceType(DnsListenersBindingType))
	fieldNameMap["listeners"] = "Listeners"
	fields["version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["template"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["template"] = "Template"
	fields["cacheSize"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cacheSize"] = "CacheSize"
	fields["dnsServers"] = bindings.NewOptionalType(bindings.NewReferenceType(IpAddressesBindingType))
	fieldNameMap["dnsServers"] = "DnsServers"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dns_config", fields, reflect.TypeOf(DnsConfig{}), fieldNameMap, validators)
}

func DnsForwardersBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipAddress"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ipAddress"] = "IpAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dns_forwarders", fields, reflect.TypeOf(DnsForwarders{}), fieldNameMap, validators)
}

func DnsListenersBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipAddress"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ipAddress"] = "IpAddress"
	fields["vnic"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vnic"] = "Vnic"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dns_listeners", fields, reflect.TypeOf(DnsListeners{}), fieldNameMap, validators)
}

func DnsResponseStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["total"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["total"] = "Total"
	fields["formErr"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["formErr"] = "FormErr"
	fields["nxDomain"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["nxDomain"] = "NxDomain"
	fields["success"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["success"] = "Success"
	fields["serverFail"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["serverFail"] = "ServerFail"
	fields["nxrrset"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["nxrrset"] = "Nxrrset"
	fields["others"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["others"] = "Others"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dns_response_stats", fields, reflect.TypeOf(DnsResponseStats{}), fieldNameMap, validators)
}

func DnsStatusAndStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timeStamp"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["timeStamp"] = "TimeStamp"
	fields["requests"] = bindings.NewOptionalType(bindings.NewReferenceType(RequestsBindingType))
	fieldNameMap["requests"] = "Requests"
	fields["responses"] = bindings.NewOptionalType(bindings.NewReferenceType(DnsResponseStatsBindingType))
	fieldNameMap["responses"] = "Responses"
	fields["cachedDBRRSet"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cachedDBRRSet"] = "CachedDBRRSet"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dns_status_and_stats", fields, reflect.TypeOf(DnsStatusAndStats{}), fieldNameMap, validators)
}

func DnsViewBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["viewMatch"] = bindings.NewOptionalType(bindings.NewReferenceType(DnsViewMatchBindingType))
	fieldNameMap["viewMatch"] = "ViewMatch"
	fields["recursion"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["recursion"] = "Recursion"
	fields["viewId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["viewId"] = "ViewId"
	fields["forwarders"] = bindings.NewOptionalType(bindings.NewReferenceType(DnsForwardersBindingType))
	fieldNameMap["forwarders"] = "Forwarders"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dns_view", fields, reflect.TypeOf(DnsView{}), fieldNameMap, validators)
}

func DnsViewMatchBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vnic"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vnic"] = "Vnic"
	fields["ipSet"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ipSet"] = "IpSet"
	fields["ipAddress"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ipAddress"] = "IpAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dns_view_match", fields, reflect.TypeOf(DnsViewMatch{}), fieldNameMap, validators)
}

func DnsViewsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dnsView"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DnsViewBindingType), reflect.TypeOf([]DnsView{})))
	fieldNameMap["dnsView"] = "DnsView"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.dns_views", fields, reflect.TypeOf(DnsViews{}), fieldNameMap, validators)
}

func EbsBackedVsanConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ebs_backed_vsan_config", fields, reflect.TypeOf(EbsBackedVsanConfig{}), fieldNameMap, validators)
}

func EdgeJobBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["edgeId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edgeId"] = "EdgeId"
	fields["module"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["module"] = "Module"
	fields["jobId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["jobId"] = "JobId"
	fields["errorCode"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["errorCode"] = "ErrorCode"
	fields["result"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResultBindingType), reflect.TypeOf([]Result{})))
	fieldNameMap["result"] = "Result"
	fields["startTime"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["startTime"] = "StartTime"
	fields["message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["message"] = "Message"
	fields["endTime"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["endTime"] = "EndTime"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.edge_job", fields, reflect.TypeOf(EdgeJob{}), fieldNameMap, validators)
}

func EdgeStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["preRulesExists"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["preRulesExists"] = "PreRulesExists"
	fields["featureStatuses"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(FeatureStatusBindingType), reflect.TypeOf([]FeatureStatus{})))
	fieldNameMap["featureStatuses"] = "FeatureStatuses"
	fields["timestamp"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["timestamp"] = "Timestamp"
	fields["publishStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["publishStatus"] = "PublishStatus"
	fields["lastPublishedPreRulesGenerationNumber"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["lastPublishedPreRulesGenerationNumber"] = "LastPublishedPreRulesGenerationNumber"
	fields["version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["edgeVmStatus"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(EdgeVmStatusBindingType), reflect.TypeOf([]EdgeVmStatus{})))
	fieldNameMap["edgeVmStatus"] = "EdgeVmStatus"
	fields["activeVseHaIndex"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["activeVseHaIndex"] = "ActiveVseHaIndex"
	fields["systemStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["systemStatus"] = "SystemStatus"
	fields["haVnicInUse"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["haVnicInUse"] = "HaVnicInUse"
	fields["edgeStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edgeStatus"] = "EdgeStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.edge_status", fields, reflect.TypeOf(EdgeStatus{}), fieldNameMap, validators)
}

func EdgeSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["featureCapabilities"] = bindings.NewOptionalType(bindings.NewReferenceType(FeatureCapabilitiesBindingType))
	fieldNameMap["featureCapabilities"] = "FeatureCapabilities"
	fields["edgeType"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edgeType"] = "EdgeType"
	fields["logicalRouterScopes"] = bindings.NewOptionalType(bindings.NewReferenceType(LogicalRouterScopesBindingType))
	fieldNameMap["logicalRouterScopes"] = "LogicalRouterScopes"
	fields["recentJobInfo"] = bindings.NewOptionalType(bindings.NewReferenceType(EdgeJobBindingType))
	fieldNameMap["recentJobInfo"] = "RecentJobInfo"
	fields["hypervisorAssist"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["hypervisorAssist"] = "HypervisorAssist"
	fields["edgeAssistId"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["edgeAssistId"] = "EdgeAssistId"
	fields["edgeStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edgeStatus"] = "EdgeStatus"
	fields["edgeAssistInstanceName"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edgeAssistInstanceName"] = "EdgeAssistInstanceName"
	fields["objectId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["objectId"] = "ObjectId"
	fields["nodeId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nodeId"] = "NodeId"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["datacenterName"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["datacenterName"] = "DatacenterName"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["clientHandle"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["clientHandle"] = "ClientHandle"
	fields["scope"] = bindings.NewOptionalType(bindings.NewReferenceType(ScopeInfoBindingType))
	fieldNameMap["scope"] = "Scope"
	fields["type"] = bindings.NewOptionalType(bindings.NewReferenceType(ObjectTypeBindingType))
	fieldNameMap["type"] = "Type_"
	fields["revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["revision"] = "Revision"
	fields["vsmUuid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vsmUuid"] = "VsmUuid"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["extendedAttributes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ExtendedAttributeBindingType), reflect.TypeOf([]ExtendedAttribute{})))
	fieldNameMap["extendedAttributes"] = "ExtendedAttributes"
	fields["localEgressEnabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["localEgressEnabled"] = "LocalEgressEnabled"
	fields["universalRevision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["universalRevision"] = "UniversalRevision"
	fields["allowedActions"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["allowedActions"] = "AllowedActions"
	fields["objectTypeName"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["objectTypeName"] = "ObjectTypeName"
	fields["isUpgradeAvailable"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["isUpgradeAvailable"] = "IsUpgradeAvailable"
	fields["isUniversal"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["isUniversal"] = "IsUniversal"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["lrouterUuid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lrouterUuid"] = "LrouterUuid"
	fields["appliancesSummary"] = bindings.NewOptionalType(bindings.NewReferenceType(AppliancesSummaryBindingType))
	fieldNameMap["appliancesSummary"] = "AppliancesSummary"
	fields["apiVersion"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["apiVersion"] = "ApiVersion"
	fields["tenantId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tenantId"] = "TenantId"
	fields["datacenterMoid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["datacenterMoid"] = "DatacenterMoid"
	fields["numberOfConnectedVnics"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["numberOfConnectedVnics"] = "NumberOfConnectedVnics"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.edge_summary", fields, reflect.TypeOf(EdgeSummary{}), fieldNameMap, validators)
}

func EdgeVmStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["index"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["index"] = "Index"
	fields["haState"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["haState"] = "HaState"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["edgeVMStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edgeVMStatus"] = "EdgeVMStatus"
	fields["preRulesGenerationNumber"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["preRulesGenerationNumber"] = "PreRulesGenerationNumber"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.edge_vm_status", fields, reflect.TypeOf(EdgeVmStatus{}), fieldNameMap, validators)
}

func EdgeVnicAddressGroupBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnetPrefixLength"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnetPrefixLength"] = "SubnetPrefixLength"
	fields["secondaryAddresses"] = bindings.NewOptionalType(bindings.NewReferenceType(SecondaryAddressesBindingType))
	fieldNameMap["secondaryAddresses"] = "SecondaryAddresses"
	fields["primaryAddress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["primaryAddress"] = "PrimaryAddress"
	fields["subnetMask"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnetMask"] = "SubnetMask"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.edge_vnic_address_group", fields, reflect.TypeOf(EdgeVnicAddressGroup{}), fieldNameMap, validators)
}

func EdgeVnicAddressGroupsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["addressGroups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(EdgeVnicAddressGroupBindingType), reflect.TypeOf([]EdgeVnicAddressGroup{})))
	fieldNameMap["addressGroups"] = "AddressGroups"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.edge_vnic_address_groups", fields, reflect.TypeOf(EdgeVnicAddressGroups{}), fieldNameMap, validators)
}

func EniInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["security_group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["security_group_id"] = "SecurityGroupId"
	fields["private_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["private_ip"] = "PrivateIp"
	fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.eni_info", fields, reflect.TypeOf(EniInfo{}), fieldNameMap, validators)
}

func EntityCapacityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["storage_capacity_gib"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["storage_capacity_gib"] = "StorageCapacityGib"
	fields["memory_capacity_gib"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["memory_capacity_gib"] = "MemoryCapacityGib"
	fields["total_number_of_cores"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["total_number_of_cores"] = "TotalNumberOfCores"
	fields["number_of_ssds"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["number_of_ssds"] = "NumberOfSsds"
	fields["cpu_capacity_ghz"] = bindings.NewOptionalType(bindings.NewDoubleType())
	fieldNameMap["cpu_capacity_ghz"] = "CpuCapacityGhz"
	fields["number_of_sockets"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["number_of_sockets"] = "NumberOfSockets"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.entity_capacity", fields, reflect.TypeOf(EntityCapacity{}), fieldNameMap, validators)
}

func ErrorResponseBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewIntegerType()
	fieldNameMap["status"] = "Status"
	fields["path"] = bindings.NewStringType()
	fieldNameMap["path"] = "Path"
	fields["retryable"] = bindings.NewBooleanType()
	fieldNameMap["retryable"] = "Retryable"
	fields["error_code"] = bindings.NewStringType()
	fieldNameMap["error_code"] = "ErrorCode"
	fields["error_messages"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["error_messages"] = "ErrorMessages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.error_response", fields, reflect.TypeOf(ErrorResponse{}), fieldNameMap, validators)
}

func EsxConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["esxs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["esxs"] = "Esxs"
	fields["cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.esx_config", fields, reflect.TypeOf(EsxConfig{}), fieldNameMap, validators)
}

func EsxHostBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["esx_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_id"] = "EsxId"
	fields["hostname"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["custom_properties"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["esx_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_state"] = "EsxState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.esx_host", fields, reflect.TypeOf(EsxHost{}), fieldNameMap, validators)
}

func EsxHostInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.esx_host_info", fields, reflect.TypeOf(EsxHostInfo{}), fieldNameMap, validators)
}

func ExtendedAttributeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["value"] = "Value"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.extended_attribute", fields, reflect.TypeOf(ExtendedAttribute{}), fieldNameMap, validators)
}

func FeatureCapabilitiesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timestamp"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["timestamp"] = "Timestamp"
	fields["featureCapabilities"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(FeatureCapabilityBindingType), reflect.TypeOf([]FeatureCapability{})))
	fieldNameMap["featureCapabilities"] = "FeatureCapabilities"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.feature_capabilities", fields, reflect.TypeOf(FeatureCapabilities{}), fieldNameMap, validators)
}

func FeatureCapabilityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["configurationLimits"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(KeyValueAttributesBindingType), reflect.TypeOf([]KeyValueAttributes{})))
	fieldNameMap["configurationLimits"] = "ConfigurationLimits"
	fields["isSupported"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["isSupported"] = "IsSupported"
	fields["service"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service"] = "Service"
	fields["permission"] = bindings.NewOptionalType(bindings.NewReferenceType(LicenceAclPermissionsBindingType))
	fieldNameMap["permission"] = "Permission"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.feature_capability", fields, reflect.TypeOf(FeatureCapability{}), fieldNameMap, validators)
}

func FeatureStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["configured"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["configured"] = "Configured"
	fields["serverStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["serverStatus"] = "ServerStatus"
	fields["publishStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["publishStatus"] = "PublishStatus"
	fields["service"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.feature_status", fields, reflect.TypeOf(FeatureStatus{}), fieldNameMap, validators)
}

func FirewallConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewallRules"] = bindings.NewOptionalType(bindings.NewReferenceType(FirewallRulesBindingType))
	fieldNameMap["firewallRules"] = "FirewallRules"
	fields["featureType"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["featureType"] = "FeatureType"
	fields["version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["template"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["template"] = "Template"
	fields["globalConfig"] = bindings.NewOptionalType(bindings.NewReferenceType(FirewallGlobalConfigBindingType))
	fieldNameMap["globalConfig"] = "GlobalConfig"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["defaultPolicy"] = bindings.NewOptionalType(bindings.NewReferenceType(FirewallDefaultPolicyBindingType))
	fieldNameMap["defaultPolicy"] = "DefaultPolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_config", fields, reflect.TypeOf(FirewallConfig{}), fieldNameMap, validators)
}

func FirewallDashboardStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connections"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["connections"] = "Connections"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_dashboard_stats", fields, reflect.TypeOf(FirewallDashboardStats{}), fieldNameMap, validators)
}

func FirewallDefaultPolicyBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["loggingEnabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["loggingEnabled"] = "LoggingEnabled"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_default_policy", fields, reflect.TypeOf(FirewallDefaultPolicy{}), fieldNameMap, validators)
}

func FirewallGlobalConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tcpAllowOutOfWindowPackets"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["tcpAllowOutOfWindowPackets"] = "TcpAllowOutOfWindowPackets"
	fields["udpTimeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["udpTimeout"] = "UdpTimeout"
	fields["ipGenericTimeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["ipGenericTimeout"] = "IpGenericTimeout"
	fields["tcpPickOngoingConnections"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["tcpPickOngoingConnections"] = "TcpPickOngoingConnections"
	fields["tcpTimeoutOpen"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["tcpTimeoutOpen"] = "TcpTimeoutOpen"
	fields["tcpTimeoutClose"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["tcpTimeoutClose"] = "TcpTimeoutClose"
	fields["icmp6Timeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["icmp6Timeout"] = "Icmp6Timeout"
	fields["dropIcmpReplays"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["dropIcmpReplays"] = "DropIcmpReplays"
	fields["logIcmpErrors"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["logIcmpErrors"] = "LogIcmpErrors"
	fields["tcpSendResetForClosedVsePorts"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["tcpSendResetForClosedVsePorts"] = "TcpSendResetForClosedVsePorts"
	fields["dropInvalidTraffic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["dropInvalidTraffic"] = "DropInvalidTraffic"
	fields["enableSynFloodProtection"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enableSynFloodProtection"] = "EnableSynFloodProtection"
	fields["icmpTimeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["icmpTimeout"] = "IcmpTimeout"
	fields["tcpTimeoutEstablished"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["tcpTimeoutEstablished"] = "TcpTimeoutEstablished"
	fields["logInvalidTraffic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["logInvalidTraffic"] = "LogInvalidTraffic"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_global_config", fields, reflect.TypeOf(FirewallGlobalConfig{}), fieldNameMap, validators)
}

func FirewallRuleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rule_type"] = "RuleType"
	fields["application_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["application_ids"] = "ApplicationIds"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["rule_interface"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rule_interface"] = "RuleInterface"
	fields["destination"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["destination"] = "Destination"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["destination_scope"] = bindings.NewOptionalType(bindings.NewReferenceType(FirewallRuleScopeBindingType))
	fieldNameMap["destination_scope"] = "DestinationScope"
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source"] = "Source"
	fields["source_scope"] = bindings.NewOptionalType(bindings.NewReferenceType(FirewallRuleScopeBindingType))
	fieldNameMap["source_scope"] = "SourceScope"
	fields["services"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(FirewallServiceBindingType), reflect.TypeOf([]FirewallService{})))
	fieldNameMap["services"] = "Services"
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["revision"] = "Revision"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_rule", fields, reflect.TypeOf(FirewallRule{}), fieldNameMap, validators)
}

func FirewallRuleScopeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["grouping_object_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["grouping_object_ids"] = "GroupingObjectIds"
	fields["vnic_group_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vnic_group_ids"] = "VnicGroupIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_rule_scope", fields, reflect.TypeOf(FirewallRuleScope{}), fieldNameMap, validators)
}

func FirewallRuleStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timestamp"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["timestamp"] = "Timestamp"
	fields["connectionCount"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["connectionCount"] = "ConnectionCount"
	fields["byteCount"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["byteCount"] = "ByteCount"
	fields["packetCount"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["packetCount"] = "PacketCount"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_rule_stats", fields, reflect.TypeOf(FirewallRuleStats{}), fieldNameMap, validators)
}

func FirewallRulesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewallRules"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NsxfirewallruleBindingType), reflect.TypeOf([]Nsxfirewallrule{})))
	fieldNameMap["firewallRules"] = "FirewallRules"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_rules", fields, reflect.TypeOf(FirewallRules{}), fieldNameMap, validators)
}

func FirewallServiceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["protocol"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["protocol"] = "Protocol"
	fields["ports"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ports"] = "Ports"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_service", fields, reflect.TypeOf(FirewallService{}), fieldNameMap, validators)
}

func GatewayTemplateBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcPublicIpBindingType))
	fieldNameMap["public_ip"] = "PublicIp"
	fields["primary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["primary_dns"] = "PrimaryDns"
	fields["secondary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["secondary_dns"] = "SecondaryDns"
	fields["firewall_rules"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(FirewallRuleBindingType), reflect.TypeOf([]FirewallRule{})))
	fieldNameMap["firewall_rules"] = "FirewallRules"
	fields["vpns"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VpnBindingType), reflect.TypeOf([]Vpn{})))
	fieldNameMap["vpns"] = "Vpns"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.gateway_template", fields, reflect.TypeOf(GatewayTemplate{}), fieldNameMap, validators)
}

func GlcmBundleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["s3Bucket"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["s3Bucket"] = "S3Bucket"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.glcm_bundle", fields, reflect.TypeOf(GlcmBundle{}), fieldNameMap, validators)
}

func HostLeaseInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["macAddress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["macAddress"] = "MacAddress"
	fields["ends"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ends"] = "Ends"
	fields["abandoned"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["abandoned"] = "Abandoned"
	fields["cltt"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cltt"] = "Cltt"
	fields["clientHostname"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["clientHostname"] = "ClientHostname"
	fields["starts"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["starts"] = "Starts"
	fields["bindingState"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["bindingState"] = "BindingState"
	fields["hardwareType"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["hardwareType"] = "HardwareType"
	fields["tsfp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tsfp"] = "Tsfp"
	fields["uid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["uid"] = "Uid"
	fields["nextBindingState"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nextBindingState"] = "NextBindingState"
	fields["ipAddress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ipAddress"] = "IpAddress"
	fields["tstp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tstp"] = "Tstp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.host_lease_info", fields, reflect.TypeOf(HostLeaseInfo{}), fieldNameMap, validators)
}

func InstanceTypeConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["hosts"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{})))
	fieldNameMap["hosts"] = "Hosts"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["entity_capacity"] = bindings.NewOptionalType(bindings.NewReferenceType(EntityCapacityBindingType))
	fieldNameMap["entity_capacity"] = "EntityCapacity"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.instance_type_config", fields, reflect.TypeOf(InstanceTypeConfig{}), fieldNameMap, validators)
}

func InteractionPermissionsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["managePermission"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["managePermission"] = "ManagePermission"
	fields["viewPermission"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["viewPermission"] = "ViewPermission"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.interaction_permissions", fields, reflect.TypeOf(InteractionPermissions{}), fieldNameMap, validators)
}

func InterfacesDashboardStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vnic_7_in_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_7_in_pkt"] = "Vnic7InPkt"
	fields["vnic_0_in_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_0_in_byte"] = "Vnic0InByte"
	fields["vnic_8_out_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_8_out_pkt"] = "Vnic8OutPkt"
	fields["vnic_5_in_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_5_in_byte"] = "Vnic5InByte"
	fields["vnic_2_in_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_2_in_pkt"] = "Vnic2InPkt"
	fields["vnic_3_in_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_3_in_pkt"] = "Vnic3InPkt"
	fields["vnic_6_out_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_6_out_byte"] = "Vnic6OutByte"
	fields["vnic_3_in_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_3_in_byte"] = "Vnic3InByte"
	fields["vnic_8_in_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_8_in_pkt"] = "Vnic8InPkt"
	fields["vnic_1_in_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_1_in_byte"] = "Vnic1InByte"
	fields["vnic_1_out_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_1_out_pkt"] = "Vnic1OutPkt"
	fields["vnic_5_out_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_5_out_byte"] = "Vnic5OutByte"
	fields["vnic_0_out_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_0_out_pkt"] = "Vnic0OutPkt"
	fields["vnic_0_out_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_0_out_byte"] = "Vnic0OutByte"
	fields["vnic_6_out_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_6_out_pkt"] = "Vnic6OutPkt"
	fields["vnic_3_out_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_3_out_byte"] = "Vnic3OutByte"
	fields["vnic_7_in_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_7_in_byte"] = "Vnic7InByte"
	fields["vnic_1_out_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_1_out_byte"] = "Vnic1OutByte"
	fields["vnic_9_out_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_9_out_pkt"] = "Vnic9OutPkt"
	fields["vnic_9_in_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_9_in_pkt"] = "Vnic9InPkt"
	fields["vnic_4_in_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_4_in_byte"] = "Vnic4InByte"
	fields["vnic_5_out_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_5_out_pkt"] = "Vnic5OutPkt"
	fields["vnic_2_out_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_2_out_pkt"] = "Vnic2OutPkt"
	fields["vnic_2_in_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_2_in_byte"] = "Vnic2InByte"
	fields["vnic_5_in_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_5_in_pkt"] = "Vnic5InPkt"
	fields["vnic_7_out_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_7_out_pkt"] = "Vnic7OutPkt"
	fields["vnic_3_out_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_3_out_pkt"] = "Vnic3OutPkt"
	fields["vnic_4_out_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_4_out_pkt"] = "Vnic4OutPkt"
	fields["vnic_4_out_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_4_out_byte"] = "Vnic4OutByte"
	fields["vnic_1_in_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_1_in_pkt"] = "Vnic1InPkt"
	fields["vnic_2_out_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_2_out_byte"] = "Vnic2OutByte"
	fields["vnic_6_in_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_6_in_byte"] = "Vnic6InByte"
	fields["vnic_0_in_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_0_in_pkt"] = "Vnic0InPkt"
	fields["vnic_9_in_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_9_in_byte"] = "Vnic9InByte"
	fields["vnic_7_out_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_7_out_byte"] = "Vnic7OutByte"
	fields["vnic_4_in_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_4_in_pkt"] = "Vnic4InPkt"
	fields["vnic_9_out_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_9_out_byte"] = "Vnic9OutByte"
	fields["vnic_8_out_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_8_out_byte"] = "Vnic8OutByte"
	fields["vnic_8_in_byte"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_8_in_byte"] = "Vnic8InByte"
	fields["vnic_6_in_pkt"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["vnic_6_in_pkt"] = "Vnic6InPkt"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.interfaces_dashboard_stats", fields, reflect.TypeOf(InterfacesDashboardStats{}), fieldNameMap, validators)
}

func IpAddressesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipAddress"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ipAddress"] = "IpAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ip_addresses", fields, reflect.TypeOf(IpAddresses{}), fieldNameMap, validators)
}

func IpsecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["featureType"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["featureType"] = "FeatureType"
	fields["logging"] = bindings.NewOptionalType(bindings.NewReferenceType(LoggingBindingType))
	fieldNameMap["logging"] = "Logging"
	fields["global"] = bindings.NewOptionalType(bindings.NewReferenceType(IpsecGlobalConfigBindingType))
	fieldNameMap["global"] = "Global"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["sites"] = bindings.NewOptionalType(bindings.NewReferenceType(IpsecSitesBindingType))
	fieldNameMap["sites"] = "Sites"
	fields["disableEvent"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["disableEvent"] = "DisableEvent"
	fields["version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["template"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["template"] = "Template"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ipsec", fields, reflect.TypeOf(Ipsec{}), fieldNameMap, validators)
}

func IpsecDashboardStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsecBytesOut"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["ipsecBytesOut"] = "IpsecBytesOut"
	fields["ipsecBytesIn"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["ipsecBytesIn"] = "IpsecBytesIn"
	fields["ipsecTunnels"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["ipsecTunnels"] = "IpsecTunnels"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ipsec_dashboard_stats", fields, reflect.TypeOf(IpsecDashboardStats{}), fieldNameMap, validators)
}

func IpsecGlobalConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["psk"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["psk"] = "Psk"
	fields["caCertificates"] = bindings.NewOptionalType(bindings.NewReferenceType(CaCertificatesBindingType))
	fieldNameMap["caCertificates"] = "CaCertificates"
	fields["serviceCertificate"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["serviceCertificate"] = "ServiceCertificate"
	fields["crlCertificates"] = bindings.NewOptionalType(bindings.NewReferenceType(CrlCertificatesBindingType))
	fieldNameMap["crlCertificates"] = "CrlCertificates"
	fields["extension"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["extension"] = "Extension"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ipsec_global_config", fields, reflect.TypeOf(IpsecGlobalConfig{}), fieldNameMap, validators)
}

func IpsecSiteBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["psk"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["psk"] = "Psk"
	fields["localId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["localId"] = "LocalId"
	fields["enablePfs"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enablePfs"] = "EnablePfs"
	fields["authenticationMode"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["authenticationMode"] = "AuthenticationMode"
	fields["peerSubnets"] = bindings.NewOptionalType(bindings.NewReferenceType(SubnetsBindingType))
	fieldNameMap["peerSubnets"] = "PeerSubnets"
	fields["dhGroup"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dhGroup"] = "DhGroup"
	fields["siteId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["siteId"] = "SiteId"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["peerIp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["peerIp"] = "PeerIp"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["certificate"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["certificate"] = "Certificate"
	fields["localIp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["localIp"] = "LocalIp"
	fields["encryptionAlgorithm"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["encryptionAlgorithm"] = "EncryptionAlgorithm"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["mtu"] = "Mtu"
	fields["extension"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["extension"] = "Extension"
	fields["peerId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["peerId"] = "PeerId"
	fields["localSubnets"] = bindings.NewOptionalType(bindings.NewReferenceType(SubnetsBindingType))
	fieldNameMap["localSubnets"] = "LocalSubnets"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ipsec_site", fields, reflect.TypeOf(IpsecSite{}), fieldNameMap, validators)
}

func IpsecSiteIKEStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["channelStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["channelStatus"] = "ChannelStatus"
	fields["channelState"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["channelState"] = "ChannelState"
	fields["peerIpAddress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["peerIpAddress"] = "PeerIpAddress"
	fields["localIpAddress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["localIpAddress"] = "LocalIpAddress"
	fields["peerSubnets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["peerSubnets"] = "PeerSubnets"
	fields["peerId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["peerId"] = "PeerId"
	fields["lastInformationalMessage"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lastInformationalMessage"] = "LastInformationalMessage"
	fields["localSubnets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["localSubnets"] = "LocalSubnets"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ipsec_site_IKE_status", fields, reflect.TypeOf(IpsecSiteIKEStatus{}), fieldNameMap, validators)
}

func IpsecSiteStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rxBytesOnSite"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["rxBytesOnSite"] = "RxBytesOnSite"
	fields["tunnelStats"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(IpsecTunnelStatsBindingType), reflect.TypeOf([]IpsecTunnelStats{})))
	fieldNameMap["tunnelStats"] = "TunnelStats"
	fields["ikeStatus"] = bindings.NewOptionalType(bindings.NewReferenceType(IpsecSiteIKEStatusBindingType))
	fieldNameMap["ikeStatus"] = "IkeStatus"
	fields["siteStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["siteStatus"] = "SiteStatus"
	fields["txBytesFromSite"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["txBytesFromSite"] = "TxBytesFromSite"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ipsec_site_stats", fields, reflect.TypeOf(IpsecSiteStats{}), fieldNameMap, validators)
}

func IpsecSitesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sites"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(IpsecSiteBindingType), reflect.TypeOf([]IpsecSite{})))
	fieldNameMap["sites"] = "Sites"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ipsec_sites", fields, reflect.TypeOf(IpsecSites{}), fieldNameMap, validators)
}

func IpsecStatusAndStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timeStamp"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["timeStamp"] = "TimeStamp"
	fields["serverStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["serverStatus"] = "ServerStatus"
	fields["siteStatistics"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(IpsecSiteStatsBindingType), reflect.TypeOf([]IpsecSiteStats{})))
	fieldNameMap["siteStatistics"] = "SiteStatistics"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ipsec_status_and_stats", fields, reflect.TypeOf(IpsecStatusAndStats{}), fieldNameMap, validators)
}

func IpsecTunnelStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tunnelStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tunnelStatus"] = "TunnelStatus"
	fields["peerSPI"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["peerSPI"] = "PeerSPI"
	fields["rxBytesOnLocalSubnet"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["rxBytesOnLocalSubnet"] = "RxBytesOnLocalSubnet"
	fields["establishedDate"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["establishedDate"] = "EstablishedDate"
	fields["peerSubnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["peerSubnet"] = "PeerSubnet"
	fields["authenticationAlgorithm"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["authenticationAlgorithm"] = "AuthenticationAlgorithm"
	fields["tunnelState"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tunnelState"] = "TunnelState"
	fields["txBytesFromLocalSubnet"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["txBytesFromLocalSubnet"] = "TxBytesFromLocalSubnet"
	fields["lastInformationalMessage"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lastInformationalMessage"] = "LastInformationalMessage"
	fields["localSPI"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["localSPI"] = "LocalSPI"
	fields["encryptionAlgorithm"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["encryptionAlgorithm"] = "EncryptionAlgorithm"
	fields["localSubnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["localSubnet"] = "LocalSubnet"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ipsec_tunnel_stats", fields, reflect.TypeOf(IpsecTunnelStats{}), fieldNameMap, validators)
}

func KeyValueAttributesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["value"] = "Value"
	fields["key"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["key"] = "Key"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.key_value_attributes", fields, reflect.TypeOf(KeyValueAttributes{}), fieldNameMap, validators)
}

func KmsVpcEndpointBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vpc_endpoint_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_endpoint_id"] = "VpcEndpointId"
	fields["network_interface_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["network_interface_ids"] = "NetworkInterfaceIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.kms_vpc_endpoint", fields, reflect.TypeOf(KmsVpcEndpoint{}), fieldNameMap, validators)
}

func L2ExtensionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tunnelId"] = bindings.NewIntegerType()
	fieldNameMap["tunnelId"] = "TunnelId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.l2_extension", fields, reflect.TypeOf(L2Extension{}), fieldNameMap, validators)
}

func L2VpnBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["sites"] = bindings.NewListType(bindings.NewReferenceType(SiteBindingType), reflect.TypeOf([]Site{}))
	fieldNameMap["sites"] = "Sites"
	fields["listener_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["listener_ip"] = "ListenerIp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.l2_vpn", fields, reflect.TypeOf(L2Vpn{}), fieldNameMap, validators)
}

func L2vpnStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tunnelStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tunnelStatus"] = "TunnelStatus"
	fields["establishedDate"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["establishedDate"] = "EstablishedDate"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["droppedRxPackets"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["droppedRxPackets"] = "DroppedRxPackets"
	fields["encryptionAlgorithm"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["encryptionAlgorithm"] = "EncryptionAlgorithm"
	fields["failureMessage"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["failureMessage"] = "FailureMessage"
	fields["txBytesFromLocalSubnet"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["txBytesFromLocalSubnet"] = "TxBytesFromLocalSubnet"
	fields["rxBytesOnLocalSubnet"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["rxBytesOnLocalSubnet"] = "RxBytesOnLocalSubnet"
	fields["droppedTxPackets"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["droppedTxPackets"] = "DroppedTxPackets"
	fields["lastUpdatedTime"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["lastUpdatedTime"] = "LastUpdatedTime"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.l2vpn_stats", fields, reflect.TypeOf(L2vpnStats{}), fieldNameMap, validators)
}

func L2vpnStatusAndStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timeStamp"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["timeStamp"] = "TimeStamp"
	fields["serverStatus"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["serverStatus"] = "ServerStatus"
	fields["siteStats"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(L2vpnStatsBindingType), reflect.TypeOf([]L2vpnStats{})))
	fieldNameMap["siteStats"] = "SiteStats"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.l2vpn_status_and_stats", fields, reflect.TypeOf(L2vpnStatusAndStats{}), fieldNameMap, validators)
}

func LicenceAclPermissionsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dataPermission"] = bindings.NewOptionalType(bindings.NewReferenceType(DataPermissionsBindingType))
	fieldNameMap["dataPermission"] = "DataPermission"
	fields["isLicensed"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["isLicensed"] = "IsLicensed"
	fields["accessPermission"] = bindings.NewOptionalType(bindings.NewReferenceType(InteractionPermissionsBindingType))
	fieldNameMap["accessPermission"] = "AccessPermission"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.licence_acl_permissions", fields, reflect.TypeOf(LicenceAclPermissions{}), fieldNameMap, validators)
}

func LoadBalancerDashboardStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lbBpsIn"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["lbBpsIn"] = "LbBpsIn"
	fields["lbHttpReqs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["lbHttpReqs"] = "LbHttpReqs"
	fields["lbBpsOut"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["lbBpsOut"] = "LbBpsOut"
	fields["lbSessions"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["lbSessions"] = "LbSessions"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.load_balancer_dashboard_stats", fields, reflect.TypeOf(LoadBalancerDashboardStats{}), fieldNameMap, validators)
}

func LoggingBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logLevel"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logLevel"] = "LogLevel"
	fields["enable"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enable"] = "Enable"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.logging", fields, reflect.TypeOf(Logging{}), fieldNameMap, validators)
}

func LogicalNetworkBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_cidr"] = "SubnetCidr"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["gatewayIp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["gatewayIp"] = "GatewayIp"
	fields["dhcp_enabled"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dhcp_enabled"] = "DhcpEnabled"
	fields["dhcp_ip_range"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dhcp_ip_range"] = "DhcpIpRange"
	fields["tunnel_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["tunnel_id"] = "TunnelId"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["network_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_type"] = "NetworkType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.logical_network", fields, reflect.TypeOf(LogicalNetwork{}), fieldNameMap, validators)
}

func LogicalRouterScopeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.logical_router_scope", fields, reflect.TypeOf(LogicalRouterScope{}), fieldNameMap, validators)
}

func LogicalRouterScopesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logicalRouterScope"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(LogicalRouterScopeBindingType), reflect.TypeOf([]LogicalRouterScope{})))
	fieldNameMap["logicalRouterScope"] = "LogicalRouterScope"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.logical_router_scopes", fields, reflect.TypeOf(LogicalRouterScopes{}), fieldNameMap, validators)
}

func MacAddressBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edgeVmHaIndex"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["edgeVmHaIndex"] = "EdgeVmHaIndex"
	fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["value"] = "Value"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.mac_address", fields, reflect.TypeOf(MacAddress{}), fieldNameMap, validators)
}

func MaintenanceWindowBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["day_of_week"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["day_of_week"] = "DayOfWeek"
	fields["hour_of_day"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["hour_of_day"] = "HourOfDay"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.maintenance_window", fields, reflect.TypeOf(MaintenanceWindow{}), fieldNameMap, validators)
}

func MaintenanceWindowEntryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["in_maintenance_window"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["in_maintenance_window"] = "InMaintenanceWindow"
	fields["reservation_schedule"] = bindings.NewOptionalType(bindings.NewReferenceType(ReservationScheduleBindingType))
	fieldNameMap["reservation_schedule"] = "ReservationSchedule"
	fields["reservation_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["reservation_id"] = "ReservationId"
	fields["in_maintenance_mode"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["in_maintenance_mode"] = "InMaintenanceMode"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.maintenance_window_entry", fields, reflect.TypeOf(MaintenanceWindowEntry{}), fieldNameMap, validators)
}

func MaintenanceWindowGetBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["day_of_week"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["day_of_week"] = "DayOfWeek"
	fields["hour_of_day"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["hour_of_day"] = "HourOfDay"
	fields["duration_min"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["duration_min"] = "DurationMin"
	fields["version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.maintenance_window_get", fields, reflect.TypeOf(MaintenanceWindowGet{}), fieldNameMap, validators)
}

func ManagementGatewayTemplateBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcPublicIpBindingType))
	fieldNameMap["public_ip"] = "PublicIp"
	fields["primary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["primary_dns"] = "PrimaryDns"
	fields["secondary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["secondary_dns"] = "SecondaryDns"
	fields["firewall_rules"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(FirewallRuleBindingType), reflect.TypeOf([]FirewallRule{})))
	fieldNameMap["firewall_rules"] = "FirewallRules"
	fields["vpns"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VpnBindingType), reflect.TypeOf([]Vpn{})))
	fieldNameMap["vpns"] = "Vpns"
	fields["subnet_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_cidr"] = "SubnetCidr"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.management_gateway_template", fields, reflect.TypeOf(ManagementGatewayTemplate{}), fieldNameMap, validators)
}

func MapZonesRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["petronas_regions_to_map"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["petronas_regions_to_map"] = "PetronasRegionsToMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.map_zones_request", fields, reflect.TypeOf(MapZonesRequest{}), fieldNameMap, validators)
}

func MetaDashboardStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vnics"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VnicBindingType), reflect.TypeOf([]Vnic{})))
	fieldNameMap["vnics"] = "Vnics"
	fields["endTime"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["endTime"] = "EndTime"
	fields["startTime"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["startTime"] = "StartTime"
	fields["interval"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["interval"] = "Interval"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.meta_dashboard_stats", fields, reflect.TypeOf(MetaDashboardStats{}), fieldNameMap, validators)
}

func MetadataBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timestamp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["timestamp"] = "Timestamp"
	fields["cycle_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cycle_id"] = "CycleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.metadata", fields, reflect.TypeOf(Metadata{}), fieldNameMap, validators)
}

func NatBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rules"] = bindings.NewOptionalType(bindings.NewReferenceType(NatRulesBindingType))
	fieldNameMap["rules"] = "Rules"
	fields["featureType"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["featureType"] = "FeatureType"
	fields["version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["template"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["template"] = "Template"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.nat", fields, reflect.TypeOf(Nat{}), fieldNameMap, validators)
}

func NatRuleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rule_type"] = "RuleType"
	fields["protocol"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["protocol"] = "Protocol"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["internal_ports"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_ports"] = "InternalPorts"
	fields["public_ports"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["public_ports"] = "PublicPorts"
	fields["public_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["public_ip"] = "PublicIp"
	fields["internal_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_ip"] = "InternalIp"
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["revision"] = "Revision"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.nat_rule", fields, reflect.TypeOf(NatRule{}), fieldNameMap, validators)
}

func NatRulesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["natRulesDtos"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NsxnatruleBindingType), reflect.TypeOf([]Nsxnatrule{})))
	fieldNameMap["natRulesDtos"] = "NatRulesDtos"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.nat_rules", fields, reflect.TypeOf(NatRules{}), fieldNameMap, validators)
}

func NetworkTemplateBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["management_gateway_templates"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ManagementGatewayTemplateBindingType), reflect.TypeOf([]ManagementGatewayTemplate{})))
	fieldNameMap["management_gateway_templates"] = "ManagementGatewayTemplates"
	fields["compute_gateway_templates"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ComputeGatewayTemplateBindingType), reflect.TypeOf([]ComputeGatewayTemplate{})))
	fieldNameMap["compute_gateway_templates"] = "ComputeGatewayTemplates"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.network_template", fields, reflect.TypeOf(NetworkTemplate{}), fieldNameMap, validators)
}

func NewCredentialsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewStringType()
	fieldNameMap["password"] = "Password"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.new_credentials", fields, reflect.TypeOf(NewCredentials{}), fieldNameMap, validators)
}

func NsxfirewallruleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ruleType"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ruleType"] = "RuleType"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["ruleId"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["ruleId"] = "RuleId"
	fields["matchTranslated"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["matchTranslated"] = "MatchTranslated"
	fields["invalidApplication"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["invalidApplication"] = "InvalidApplication"
	fields["direction"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["direction"] = "Direction"
	fields["statistics"] = bindings.NewOptionalType(bindings.NewReferenceType(FirewallRuleStatsBindingType))
	fieldNameMap["statistics"] = "Statistics"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["invalidSource"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["invalidSource"] = "InvalidSource"
	fields["loggingEnabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["loggingEnabled"] = "LoggingEnabled"
	fields["destination"] = bindings.NewOptionalType(bindings.NewReferenceType(AddressFWSourceDestinationBindingType))
	fieldNameMap["destination"] = "Destination"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["application"] = bindings.NewOptionalType(bindings.NewReferenceType(ApplicationBindingType))
	fieldNameMap["application"] = "Application"
	fields["source"] = bindings.NewOptionalType(bindings.NewReferenceType(AddressFWSourceDestinationBindingType))
	fieldNameMap["source"] = "Source"
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["invalidDestination"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["invalidDestination"] = "InvalidDestination"
	fields["ruleTag"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["ruleTag"] = "RuleTag"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.nsxfirewallrule", fields, reflect.TypeOf(Nsxfirewallrule{}), fieldNameMap, validators)
}

func NsxfirewallserviceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sourcePort"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sourcePort"] = "SourcePort"
	fields["protocol"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["protocol"] = "Protocol"
	fields["port"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["port"] = "Port"
	fields["icmpType"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["icmpType"] = "IcmpType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.nsxfirewallservice", fields, reflect.TypeOf(Nsxfirewallservice{}), fieldNameMap, validators)
}

func Nsxl2vpnBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["listenerIps"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["listenerIps"] = "ListenerIps"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["sites"] = bindings.NewReferenceType(SitesBindingType)
	fieldNameMap["sites"] = "Sites"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.nsxl2vpn", fields, reflect.TypeOf(Nsxl2vpn{}), fieldNameMap, validators)
}

func NsxnatruleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vnic"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vnic"] = "Vnic"
	fields["ruleType"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ruleType"] = "RuleType"
	fields["protocol"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["protocol"] = "Protocol"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["ruleId"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["ruleId"] = "RuleId"
	fields["snatMatchDestinationPort"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["snatMatchDestinationPort"] = "SnatMatchDestinationPort"
	fields["originalAddress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["originalAddress"] = "OriginalAddress"
	fields["dnatMatchSourceAddress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dnatMatchSourceAddress"] = "DnatMatchSourceAddress"
	fields["dnatMatchSourcePort"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dnatMatchSourcePort"] = "DnatMatchSourcePort"
	fields["snatMatchDestinationAddress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["snatMatchDestinationAddress"] = "SnatMatchDestinationAddress"
	fields["originalPort"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["originalPort"] = "OriginalPort"
	fields["loggingEnabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["loggingEnabled"] = "LoggingEnabled"
	fields["translatedAddress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["translatedAddress"] = "TranslatedAddress"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["icmpType"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["icmpType"] = "IcmpType"
	fields["translatedPort"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["translatedPort"] = "TranslatedPort"
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["ruleTag"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["ruleTag"] = "RuleTag"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.nsxnatrule", fields, reflect.TypeOf(Nsxnatrule{}), fieldNameMap, validators)
}

func NsxsiteBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["secureTraffic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["secureTraffic"] = "SecureTraffic"
	fields["siteId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["siteId"] = "SiteId"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["password"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["password"] = "Password"
	fields["userId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["userId"] = "UserId"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.nsxsite", fields, reflect.TypeOf(Nsxsite{}), fieldNameMap, validators)
}

func NsxtAddonsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enable_nsx_advanced_addon"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enable_nsx_advanced_addon"] = "EnableNsxAdvancedAddon"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.nsxt_addons", fields, reflect.TypeOf(NsxtAddons{}), fieldNameMap, validators)
}

func ObjectTypeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.object_type", fields, reflect.TypeOf(ObjectType{}), fieldNameMap, validators)
}

func OfferInstancesHolderBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["on_demand"] = bindings.NewReferenceType(OnDemandOfferInstanceBindingType)
	fieldNameMap["on_demand"] = "OnDemand"
	fields["offers"] = bindings.NewListType(bindings.NewReferenceType(TermOfferInstanceBindingType), reflect.TypeOf([]TermOfferInstance{}))
	fieldNameMap["offers"] = "Offers"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.offer_instances_holder", fields, reflect.TypeOf(OfferInstancesHolder{}), fieldNameMap, validators)
}

func OnDemandOfferInstanceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = bindings.NewStringType()
	fieldNameMap["product"] = "Product"
	fields["product_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_type"] = "ProductType"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["currency"] = bindings.NewStringType()
	fieldNameMap["currency"] = "Currency"
	fields["region"] = bindings.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["unit_price"] = bindings.NewStringType()
	fieldNameMap["unit_price"] = "UnitPrice"
	fields["monthly_cost"] = bindings.NewStringType()
	fieldNameMap["monthly_cost"] = "MonthlyCost"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.on_demand_offer_instance", fields, reflect.TypeOf(OnDemandOfferInstance{}), fieldNameMap, validators)
}

func OrgPropertiesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["values"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["values"] = "Values"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.org_properties", fields, reflect.TypeOf(OrgProperties{}), fieldNameMap, validators)
}

func OrganizationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["org_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_type"] = "OrgType"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["project_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["project_state"] = "ProjectState"
	fields["properties"] = bindings.NewOptionalType(bindings.NewReferenceType(OrgPropertiesBindingType))
	fieldNameMap["properties"] = "Properties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.organization", fields, reflect.TypeOf(Organization{}), fieldNameMap, validators)
}

func PagedEdgeListBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edgePage"] = bindings.NewOptionalType(bindings.NewReferenceType(DataPageEdgeSummaryBindingType))
	fieldNameMap["edgePage"] = "EdgePage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.paged_edge_list", fields, reflect.TypeOf(PagedEdgeList{}), fieldNameMap, validators)
}

func PagingInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sortOrderAscending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sortOrderAscending"] = "SortOrderAscending"
	fields["totalCount"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["totalCount"] = "TotalCount"
	fields["startIndex"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["startIndex"] = "StartIndex"
	fields["sortBy"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sortBy"] = "SortBy"
	fields["pageSize"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["pageSize"] = "PageSize"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.paging_info", fields, reflect.TypeOf(PagingInfo{}), fieldNameMap, validators)
}

func PaymentMethodInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["default_flag"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["default_flag"] = "DefaultFlag"
	fields["payment_method_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["payment_method_id"] = "PaymentMethodId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.payment_method_info", fields, reflect.TypeOf(PaymentMethodInfo{}), fieldNameMap, validators)
}

func PopAgentXeniConnectionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["default_subnet_route"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["default_subnet_route"] = "DefaultSubnetRoute"
	fields["eni_info"] = bindings.NewOptionalType(bindings.NewReferenceType(EniInfoBindingType))
	fieldNameMap["eni_info"] = "EniInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.pop_agent_xeni_connection", fields, reflect.TypeOf(PopAgentXeniConnection{}), fieldNameMap, validators)
}

func PopAmiInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.pop_ami_info", fields, reflect.TypeOf(PopAmiInfo{}), fieldNameMap, validators)
}

func PopInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ami_infos"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(PopAmiInfoBindingType),reflect.TypeOf(map[string]PopAmiInfo{}))
	fieldNameMap["ami_infos"] = "AmiInfos"
	fields["created_at"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["created_at"] = "CreatedAt"
	fields["service_infos"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(PopServiceInfoBindingType),reflect.TypeOf(map[string]PopServiceInfo{})))
	fieldNameMap["service_infos"] = "ServiceInfos"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["manifest_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["manifest_version"] = "ManifestVersion"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.pop_info", fields, reflect.TypeOf(PopInfo{}), fieldNameMap, validators)
}

func PopServiceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cln"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cln"] = "Cln"
	fields["version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["version"] = "Version"
	fields["build"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["build"] = "Build"
	fields["service"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.pop_service_info", fields, reflect.TypeOf(PopServiceInfo{}), fieldNameMap, validators)
}

func ProvisionSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(SddcConfigSpecBindingType),reflect.TypeOf(map[string]SddcConfigSpec{})))
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.provision_spec", fields, reflect.TypeOf(ProvisionSpec{}), fieldNameMap, validators)
}

func RequestsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["total"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["total"] = "Total"
	fields["queries"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["queries"] = "Queries"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.requests", fields, reflect.TypeOf(Requests{}), fieldNameMap, validators)
}

func ReservationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["duration"] = "Duration"
	fields["rid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rid"] = "Rid"
	fields["create_time"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["create_time"] = "CreateTime"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["metadata"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.reservation", fields, reflect.TypeOf(Reservation{}), fieldNameMap, validators)
}

func ReservationInMwBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rid"] = "Rid"
	fields["week_of"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["week_of"] = "WeekOf"
	fields["create_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["create_time"] = "CreateTime"
	fields["metadata"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.reservation_in_mw", fields, reflect.TypeOf(ReservationInMw{}), fieldNameMap, validators)
}

func ReservationScheduleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["day_of_week"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["day_of_week"] = "DayOfWeek"
	fields["hour_of_day"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["hour_of_day"] = "HourOfDay"
	fields["duration_min"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["duration_min"] = "DurationMin"
	fields["version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["reservations"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ReservationBindingType), reflect.TypeOf([]Reservation{})))
	fieldNameMap["reservations"] = "Reservations"
	fields["reservations_mw"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ReservationInMwBindingType), reflect.TypeOf([]ReservationInMw{})))
	fieldNameMap["reservations_mw"] = "ReservationsMw"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.reservation_schedule", fields, reflect.TypeOf(ReservationSchedule{}), fieldNameMap, validators)
}

func ReservationWindowBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reservation_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["reservation_state"] = "ReservationState"
	fields["emergency"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["emergency"] = "Emergency"
	fields["maintenance_properties"] = bindings.NewOptionalType(bindings.NewReferenceType(ReservationWindowMaintenancePropertiesBindingType))
	fieldNameMap["maintenance_properties"] = "MaintenanceProperties"
	fields["reserve_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["reserve_id"] = "ReserveId"
	fields["start_hour"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["start_hour"] = "StartHour"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["manifest_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["manifest_id"] = "ManifestId"
	fields["duration_hours"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["duration_hours"] = "DurationHours"
	fields["start_date"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["start_date"] = "StartDate"
	fields["metadata"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.reservation_window", fields, reflect.TypeOf(ReservationWindow{}), fieldNameMap, validators)
}

func ReservationWindowMaintenancePropertiesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.reservation_window_maintenance_properties", fields, reflect.TypeOf(ReservationWindowMaintenanceProperties{}), fieldNameMap, validators)
}

func ResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["value"] = "Value"
	fields["key"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["key"] = "Key"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.result", fields, reflect.TypeOf(Result{}), fieldNameMap, validators)
}

func RouteTableInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.route_table_info", fields, reflect.TypeOf(RouteTableInfo{}), fieldNameMap, validators)
}

func ScopeInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["objectTypeName"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["objectTypeName"] = "ObjectTypeName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.scope_info", fields, reflect.TypeOf(ScopeInfo{}), fieldNameMap, validators)
}

func SddcBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["sddc_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_state"] = "SddcState"
	fields["expiration_date"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["expiration_date"] = "ExpirationDate"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["sddc_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_type"] = "SddcType"
	fields["provider"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["provider"] = "Provider"
	fields["account_link_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["account_link_state"] = "AccountLinkState"
	fields["sddc_access_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_access_state"] = "SddcAccessState"
	fields["resource_config"] = bindings.NewOptionalType(bindings.NewReferenceType(AwsSddcResourceConfigBindingType))
	fieldNameMap["resource_config"] = "ResourceConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc", fields, reflect.TypeOf(Sddc{}), fieldNameMap, validators)
}

func SddcAllocatePublicIpSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["count"] = bindings.NewIntegerType()
	fieldNameMap["count"] = "Count"
	fields["private_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["private_ips"] = "PrivateIps"
	fields["names"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["names"] = "Names"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_allocate_public_ip_spec", fields, reflect.TypeOf(SddcAllocatePublicIpSpec{}), fieldNameMap, validators)
}

func SddcConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vpc_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_cidr"] = "VpcCidr"
	fields["host_instance_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vmc.model.host_instance_types", reflect.TypeOf(HostInstanceTypes(HostInstanceTypes_I3_METAL))))
	fieldNameMap["host_instance_type"] = "HostInstanceType"
	fields["skip_creating_vxlan"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["vxlan_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["size"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["size"] = "Size"
	fields["storage_capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["account_link_sddc_config"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AccountLinkSddcConfigBindingType), reflect.TypeOf([]AccountLinkSddcConfig{})))
	fieldNameMap["account_link_sddc_config"] = "AccountLinkSddcConfig"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	fields["sddc_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_type"] = "SddcType"
	fields["account_link_config"] = bindings.NewOptionalType(bindings.NewReferenceType(AccountLinkConfigBindingType))
	fieldNameMap["account_link_config"] = "AccountLinkConfig"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["sso_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["sddc_template_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_template_id"] = "SddcTemplateId"
	fields["deployment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_config", fields, reflect.TypeOf(SddcConfig{}), fieldNameMap, validators)
}

func SddcConfigSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_type_config_spec"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ConfigSpecBindingType),reflect.TypeOf(map[string]ConfigSpec{})))
	fieldNameMap["sddc_type_config_spec"] = "SddcTypeConfigSpec"
	fields["region_display_names"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["region_display_names"] = "RegionDisplayNames"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_config_spec", fields, reflect.TypeOf(SddcConfigSpec{}), fieldNameMap, validators)
}

func SddcIdBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_id", fields, reflect.TypeOf(SddcId{}), fieldNameMap, validators)
}

func SddcLinkConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customer_subnet_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["customer_subnet_ids"] = "CustomerSubnetIds"
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_link_config", fields, reflect.TypeOf(SddcLinkConfig{}), fieldNameMap, validators)
}

func SddcManifestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vmc_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vmc_version"] = "VmcVersion"
	fields["glcm_bundle"] = bindings.NewOptionalType(bindings.NewReferenceType(GlcmBundleBindingType))
	fieldNameMap["glcm_bundle"] = "GlcmBundle"
	fields["pop_info"] = bindings.NewOptionalType(bindings.NewReferenceType(PopInfoBindingType))
	fieldNameMap["pop_info"] = "PopInfo"
	fields["vmc_internal_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vmc_internal_version"] = "VmcInternalVersion"
	fields["ebs_backed_vsan_config"] = bindings.NewOptionalType(bindings.NewReferenceType(EbsBackedVsanConfigBindingType))
	fieldNameMap["ebs_backed_vsan_config"] = "EbsBackedVsanConfig"
	fields["vsan_witness_ami"] = bindings.NewOptionalType(bindings.NewReferenceType(AmiInfoBindingType))
	fieldNameMap["vsan_witness_ami"] = "VsanWitnessAmi"
	fields["esx_ami"] = bindings.NewOptionalType(bindings.NewReferenceType(AmiInfoBindingType))
	fieldNameMap["esx_ami"] = "EsxAmi"
	fields["esx_nsxt_ami"] = bindings.NewOptionalType(bindings.NewReferenceType(AmiInfoBindingType))
	fieldNameMap["esx_nsxt_ami"] = "EsxNsxtAmi"
	fields["metadata"] = bindings.NewOptionalType(bindings.NewReferenceType(MetadataBindingType))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_manifest", fields, reflect.TypeOf(SddcManifest{}), fieldNameMap, validators)
}

func SddcNetworkBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnets"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcNetworkAddressGroupsBindingType))
	fieldNameMap["subnets"] = "Subnets"
	fields["cgwName"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cgwName"] = "CgwName"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["l2Extension"] = bindings.NewOptionalType(bindings.NewReferenceType(L2ExtensionBindingType))
	fieldNameMap["l2Extension"] = "L2Extension"
	fields["cgwId"] = bindings.NewStringType()
	fieldNameMap["cgwId"] = "CgwId"
	fields["dhcpConfigs"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcNetworkDhcpConfigBindingType))
	fieldNameMap["dhcpConfigs"] = "DhcpConfigs"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_network", fields, reflect.TypeOf(SddcNetwork{}), fieldNameMap, validators)
}

func SddcNetworkAddressGroupBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["prefixLength"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["prefixLength"] = "PrefixLength"
	fields["primaryAddress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["primaryAddress"] = "PrimaryAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_network_address_group", fields, reflect.TypeOf(SddcNetworkAddressGroup{}), fieldNameMap, validators)
}

func SddcNetworkAddressGroupsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["addressGroups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SddcNetworkAddressGroupBindingType), reflect.TypeOf([]SddcNetworkAddressGroup{})))
	fieldNameMap["addressGroups"] = "AddressGroups"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_network_address_groups", fields, reflect.TypeOf(SddcNetworkAddressGroups{}), fieldNameMap, validators)
}

func SddcNetworkDhcpConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipPools"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SddcNetworkDhcpIpPoolBindingType), reflect.TypeOf([]SddcNetworkDhcpIpPool{})))
	fieldNameMap["ipPools"] = "IpPools"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_network_dhcp_config", fields, reflect.TypeOf(SddcNetworkDhcpConfig{}), fieldNameMap, validators)
}

func SddcNetworkDhcpIpPoolBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipRange"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ipRange"] = "IpRange"
	fields["domainName"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domainName"] = "DomainName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_network_dhcp_ip_pool", fields, reflect.TypeOf(SddcNetworkDhcpIpPool{}), fieldNameMap, validators)
}

func SddcPatchRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_patch_request", fields, reflect.TypeOf(SddcPatchRequest{}), fieldNameMap, validators)
}

func SddcPublicIpBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = bindings.NewStringType()
	fieldNameMap["public_ip"] = "PublicIp"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["allocation_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["allocation_id"] = "AllocationId"
	fields["dnat_rule_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dnat_rule_id"] = "DnatRuleId"
	fields["associated_private_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["associated_private_ip"] = "AssociatedPrivateIp"
	fields["snat_rule_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["snat_rule_id"] = "SnatRuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_public_ip", fields, reflect.TypeOf(SddcPublicIp{}), fieldNameMap, validators)
}

func SddcResourceConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mgmt_appliance_network_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mgmt_appliance_network_name"] = "MgmtApplianceNetworkName"
	fields["nsxt"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["nsxt"] = "Nsxt"
	fields["mgw_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mgw_id"] = "MgwId"
	fields["nsx_mgr_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_mgr_url"] = "NsxMgrUrl"
	fields["psc_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["psc_management_ip"] = "PscManagementIp"
	fields["psc_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["psc_url"] = "PscUrl"
	fields["cgws"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["cgws"] = "Cgws"
	fields["availability_zones"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["availability_zones"] = "AvailabilityZones"
	fields["management_ds"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_ds"] = "ManagementDs"
	fields["nsx_api_public_endpoint_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_api_public_endpoint_url"] = "NsxApiPublicEndpointUrl"
	fields["custom_properties"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["cloud_password"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_password"] = "CloudPassword"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["clusters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ClusterBindingType), reflect.TypeOf([]Cluster{})))
	fieldNameMap["clusters"] = "Clusters"
	fields["vc_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_management_ip"] = "VcManagementIp"
	fields["sddc_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_networks"] = "SddcNetworks"
	fields["cloud_username"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_username"] = "CloudUsername"
	fields["esx_hosts"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AwsEsxHostBindingType), reflect.TypeOf([]AwsEsxHost{})))
	fieldNameMap["esx_hosts"] = "EsxHosts"
	fields["nsx_mgr_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_mgr_management_ip"] = "NsxMgrManagementIp"
	fields["vc_instance_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_instance_id"] = "VcInstanceId"
	fields["esx_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_cluster_id"] = "EsxClusterId"
	fields["vc_public_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_public_ip"] = "VcPublicIp"
	fields["skip_creating_vxlan"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["vc_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_url"] = "VcUrl"
	fields["sddc_manifest"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcManifestBindingType))
	fieldNameMap["sddc_manifest"] = "SddcManifest"
	fields["vxlan_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["cloud_user_group"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_user_group"] = "CloudUserGroup"
	fields["management_rp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_rp"] = "ManagementRp"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["witness_availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["witness_availability_zone"] = "WitnessAvailabilityZone"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["pop_agent_xeni_connection"] = bindings.NewOptionalType(bindings.NewReferenceType(PopAgentXeniConnectionBindingType))
	fieldNameMap["pop_agent_xeni_connection"] = "PopAgentXeniConnection"
	fields["sddc_size"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcSizeBindingType))
	fieldNameMap["sddc_size"] = "SddcSize"
	fields["nsx_controller_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["nsx_controller_ips"] = "NsxControllerIps"
	fields["esx_host_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_host_subnet"] = "EsxHostSubnet"
	fields["sso_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["deployment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	fields["nsxt_addons"] = bindings.NewOptionalType(bindings.NewReferenceType(NsxtAddonsBindingType))
	fieldNameMap["nsxt_addons"] = "NsxtAddons"
	fields["dns_with_management_vm_private_ip"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["dns_with_management_vm_private_ip"] = "DnsWithManagementVmPrivateIp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_resource_config", fields, reflect.TypeOf(SddcResourceConfig{}), fieldNameMap, validators)
}

func SddcSizeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vc_size"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_size"] = "VcSize"
	fields["nsx_size"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_size"] = "NsxSize"
	fields["size"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["size"] = "Size"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_size", fields, reflect.TypeOf(SddcSize{}), fieldNameMap, validators)
}

func SddcStateRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddcs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddcs"] = "Sddcs"
	fields["states"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["states"] = "States"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_state_request", fields, reflect.TypeOf(SddcStateRequest{}), fieldNameMap, validators)
}

func SddcTemplateBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["account_link_sddc_configs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AccountLinkSddcConfigBindingType), reflect.TypeOf([]AccountLinkSddcConfig{})))
	fieldNameMap["account_link_sddc_configs"] = "AccountLinkSddcConfigs"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["network_template"] = bindings.NewOptionalType(bindings.NewReferenceType(NetworkTemplateBindingType))
	fieldNameMap["network_template"] = "NetworkTemplate"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["source_sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source_sddc_id"] = "SourceSddcId"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["sddc"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcBindingType))
	fieldNameMap["sddc"] = "Sddc"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_template", fields, reflect.TypeOf(SddcTemplate{}), fieldNameMap, validators)
}

func SecondaryAddressesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["ipAddress"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ipAddress"] = "IpAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.secondary_addresses", fields, reflect.TypeOf(SecondaryAddresses{}), fieldNameMap, validators)
}

func ServiceErrorBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["default_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["default_message"] = "DefaultMessage"
	fields["original_service"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["original_service"] = "OriginalService"
	fields["localized_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["localized_message"] = "LocalizedMessage"
	fields["original_service_error_code"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["original_service_error_code"] = "OriginalServiceErrorCode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.service_error", fields, reflect.TypeOf(ServiceError{}), fieldNameMap, validators)
}

func SiteBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["password"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["password"] = "Password"
	fields["user_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["user_id"] = "UserId"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["rx_bytes_on_local_subnet"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["rx_bytes_on_local_subnet"] = "RxBytesOnLocalSubnet"
	fields["secure_traffic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["secure_traffic"] = "SecureTraffic"
	fields["established_date"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["established_date"] = "EstablishedDate"
	fields["failure_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["failure_message"] = "FailureMessage"
	fields["dropped_tx_packets"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dropped_tx_packets"] = "DroppedTxPackets"
	fields["dropped_rx_packets"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dropped_rx_packets"] = "DroppedRxPackets"
	fields["tunnel_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tunnel_status"] = "TunnelStatus"
	fields["tx_bytes_from_local_subnet"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["tx_bytes_from_local_subnet"] = "TxBytesFromLocalSubnet"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.site", fields, reflect.TypeOf(Site{}), fieldNameMap, validators)
}

func SitesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sites"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NsxsiteBindingType), reflect.TypeOf([]Nsxsite{})))
	fieldNameMap["sites"] = "Sites"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sites", fields, reflect.TypeOf(Sites{}), fieldNameMap, validators)
}

func SslvpnDashboardStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["activeClients"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["activeClients"] = "ActiveClients"
	fields["sslvpnBytesIn"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["sslvpnBytesIn"] = "SslvpnBytesIn"
	fields["authFailures"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["authFailures"] = "AuthFailures"
	fields["sessionsCreated"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["sessionsCreated"] = "SessionsCreated"
	fields["sslvpnBytesOut"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DashboardStatBindingType), reflect.TypeOf([]DashboardStat{})))
	fieldNameMap["sslvpnBytesOut"] = "SslvpnBytesOut"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sslvpn_dashboard_stats", fields, reflect.TypeOf(SslvpnDashboardStats{}), fieldNameMap, validators)
}

func SubInterfaceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["index"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["index"] = "Index"
	fields["tunnelId"] = bindings.NewIntegerType()
	fieldNameMap["tunnelId"] = "TunnelId"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["addressGroups"] = bindings.NewOptionalType(bindings.NewReferenceType(EdgeVnicAddressGroupsBindingType))
	fieldNameMap["addressGroups"] = "AddressGroups"
	fields["vlanId"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["vlanId"] = "VlanId"
	fields["label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["label"] = "Label"
	fields["logicalSwitchName"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logicalSwitchName"] = "LogicalSwitchName"
	fields["isConnected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["isConnected"] = "IsConnected"
	fields["mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["mtu"] = "Mtu"
	fields["logicalSwitchId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logicalSwitchId"] = "LogicalSwitchId"
	fields["enableSendRedirects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enableSendRedirects"] = "EnableSendRedirects"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sub_interface", fields, reflect.TypeOf(SubInterface{}), fieldNameMap, validators)
}

func SubInterfacesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subInterfaces"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SubInterfaceBindingType), reflect.TypeOf([]SubInterface{})))
	fieldNameMap["subInterfaces"] = "SubInterfaces"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sub_interfaces", fields, reflect.TypeOf(SubInterfaces{}), fieldNameMap, validators)
}

func SubnetBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["route_tables"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SubnetRouteTableInfoBindingType), reflect.TypeOf([]SubnetRouteTableInfo{})))
	fieldNameMap["route_tables"] = "RouteTables"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subnet", fields, reflect.TypeOf(Subnet{}), fieldNameMap, validators)
}

func SubnetInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compatible"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["compatible"] = "Compatible"
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["region_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region_name"] = "RegionName"
	fields["availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["availability_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone_id"] = "AvailabilityZoneId"
	fields["subnet_cidr_block"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_cidr_block"] = "SubnetCidrBlock"
	fields["note"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["note"] = "Note"
	fields["vpc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["vpc_cidr_block"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_cidr_block"] = "VpcCidrBlock"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subnet_info", fields, reflect.TypeOf(SubnetInfo{}), fieldNameMap, validators)
}

func SubnetRouteTableInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["association_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["association_id"] = "AssociationId"
	fields["routetable_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["routetable_id"] = "RoutetableId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subnet_route_table_info", fields, reflect.TypeOf(SubnetRouteTableInfo{}), fieldNameMap, validators)
}

func SubnetsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["subnets"] = "Subnets"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subnets", fields, reflect.TypeOf(Subnets{}), fieldNameMap, validators)
}

func SubscriptionDetailsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["anniversary_billing_date"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["anniversary_billing_date"] = "AnniversaryBillingDate"
	fields["end_date"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["end_date"] = "EndDate"
	fields["billing_frequency"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["billing_frequency"] = "BillingFrequency"
	fields["auto_renewed_allowed"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["auto_renewed_allowed"] = "AutoRenewedAllowed"
	fields["commitment_term"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["commitment_term"] = "CommitmentTerm"
	fields["csp_subscription_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["csp_subscription_id"] = "CspSubscriptionId"
	fields["billing_subscription_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["billing_subscription_id"] = "BillingSubscriptionId"
	fields["offer_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["offer_version"] = "OfferVersion"
	fields["offer_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vmc.model.offer_type", reflect.TypeOf(OfferType(OfferType_TERM))))
	fieldNameMap["offer_type"] = "OfferType"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["product_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_id"] = "ProductId"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["product_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_name"] = "ProductName"
	fields["offer_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["offer_name"] = "OfferName"
	fields["commitment_term_uom"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["commitment_term_uom"] = "CommitmentTermUom"
	fields["start_date"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["start_date"] = "StartDate"
	fields["quantity"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["quantity"] = "Quantity"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subscription_details", fields, reflect.TypeOf(SubscriptionDetails{}), fieldNameMap, validators)
}

func SubscriptionProductsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product"] = "Product"
	fields["types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["types"] = "Types"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subscription_products", fields, reflect.TypeOf(SubscriptionProducts{}), fieldNameMap, validators)
}

func SubscriptionRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product"] = "Product"
	fields["product_type"] = bindings.NewStringType()
	fieldNameMap["product_type"] = "ProductType"
	fields["product_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_id"] = "ProductId"
	fields["billing_frequency"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["billing_frequency"] = "BillingFrequency"
	fields["region"] = bindings.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["commitment_term"] = bindings.NewStringType()
	fieldNameMap["commitment_term"] = "CommitmentTerm"
	fields["offer_context_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["offer_context_id"] = "OfferContextId"
	fields["offer_version"] = bindings.NewStringType()
	fieldNameMap["offer_version"] = "OfferVersion"
	fields["offer_name"] = bindings.NewStringType()
	fieldNameMap["offer_name"] = "OfferName"
	fields["quantity"] = bindings.NewIntegerType()
	fieldNameMap["quantity"] = "Quantity"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["price"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["price"] = "Price"
	fields["product_charge_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_charge_id"] = "ProductChargeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subscription_request", fields, reflect.TypeOf(SubscriptionRequest{}), fieldNameMap, validators)
}

func SupportWindowBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["start_day"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["start_day"] = "StartDay"
	fields["seats"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["seats"] = "Seats"
	fields["sddcs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddcs"] = "Sddcs"
	fields["duration_hours"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["duration_hours"] = "DurationHours"
	fields["start_hour"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["start_hour"] = "StartHour"
	fields["support_window_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["support_window_id"] = "SupportWindowId"
	fields["metadata"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.REST))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.support_window", fields, reflect.TypeOf(SupportWindow{}), fieldNameMap, validators)
}

func SupportWindowIdBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["window_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["window_id"] = "WindowId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.support_window_id", fields, reflect.TypeOf(SupportWindowId{}), fieldNameMap, validators)
}

func TaskBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["localized_error_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["localized_error_message"] = "LocalizedErrorMessage"
	fields["resource_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["parent_task_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_task_id"] = "ParentTaskId"
	fields["task_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["task_version"] = "TaskVersion"
	fields["correlation_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["correlation_id"] = "CorrelationId"
	fields["start_resource_entity_version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["start_resource_entity_version"] = "StartResourceEntityVersion"
	fields["sub_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sub_status"] = "SubStatus"
	fields["task_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["task_type"] = "TaskType"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["task_progress_phases"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TaskProgressPhaseBindingType), reflect.TypeOf([]TaskProgressPhase{})))
	fieldNameMap["task_progress_phases"] = "TaskProgressPhases"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["end_resource_entity_version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["end_resource_entity_version"] = "EndResourceEntityVersion"
	fields["service_errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ServiceErrorBindingType), reflect.TypeOf([]ServiceError{})))
	fieldNameMap["service_errors"] = "ServiceErrors"
	fields["org_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_type"] = "OrgType"
	fields["estimated_remaining_minutes"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["estimated_remaining_minutes"] = "EstimatedRemainingMinutes"
	fields["params"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.REST))
	fieldNameMap["params"] = "Params"
	fields["progress_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["progress_percent"] = "ProgressPercent"
	fields["phase_in_progress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["phase_in_progress"] = "PhaseInProgress"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.task", fields, reflect.TypeOf(Task{}), fieldNameMap, validators)
}

func TaskProgressPhaseBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["progress_percent"] = bindings.NewIntegerType()
	fieldNameMap["progress_percent"] = "ProgressPercent"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.task_progress_phase", fields, reflect.TypeOf(TaskProgressPhase{}), fieldNameMap, validators)
}

func TermBillingOptionsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["unit_price"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["unit_price"] = "UnitPrice"
	fields["billing_frequency"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["billing_frequency"] = "BillingFrequency"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.term_billing_options", fields, reflect.TypeOf(TermBillingOptions{}), fieldNameMap, validators)
}

func TermOfferInstanceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["product"] = bindings.NewStringType()
	fieldNameMap["product"] = "Product"
	fields["product_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_type"] = "ProductType"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["currency"] = bindings.NewStringType()
	fieldNameMap["currency"] = "Currency"
	fields["region"] = bindings.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["commitment_term"] = bindings.NewIntegerType()
	fieldNameMap["commitment_term"] = "CommitmentTerm"
	fields["unit_price"] = bindings.NewStringType()
	fieldNameMap["unit_price"] = "UnitPrice"
	fields["billing_options"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TermBillingOptionsBindingType), reflect.TypeOf([]TermBillingOptions{})))
	fieldNameMap["billing_options"] = "BillingOptions"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["offer_context_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["offer_context_id"] = "OfferContextId"
	fields["product_charge_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_charge_id"] = "ProductChargeId"
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["product_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_id"] = "ProductId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.term_offer_instance", fields, reflect.TypeOf(TermOfferInstance{}), fieldNameMap, validators)
}

func TrafficShapingPolicyBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["burstSize"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["burstSize"] = "BurstSize"
	fields["averageBandwidth"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["averageBandwidth"] = "AverageBandwidth"
	fields["peakBandwidth"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["peakBandwidth"] = "PeakBandwidth"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["inherited"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["inherited"] = "Inherited"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.traffic_shaping_policy", fields, reflect.TypeOf(TrafficShapingPolicy{}), fieldNameMap, validators)
}

func UpdateCredentialsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewStringType()
	fieldNameMap["password"] = "Password"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.update_credentials", fields, reflect.TypeOf(UpdateCredentials{}), fieldNameMap, validators)
}

func VmcLocaleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["locale"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["locale"] = "Locale"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vmc_locale", fields, reflect.TypeOf(VmcLocale{}), fieldNameMap, validators)
}

func VnicBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subInterfaces"] = bindings.NewOptionalType(bindings.NewReferenceType(SubInterfacesBindingType))
	fieldNameMap["subInterfaces"] = "SubInterfaces"
	fields["addressGroups"] = bindings.NewOptionalType(bindings.NewReferenceType(EdgeVnicAddressGroupsBindingType))
	fieldNameMap["addressGroups"] = "AddressGroups"
	fields["isConnected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["isConnected"] = "IsConnected"
	fields["enableSendRedirects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enableSendRedirects"] = "EnableSendRedirects"
	fields["inShapingPolicy"] = bindings.NewOptionalType(bindings.NewReferenceType(TrafficShapingPolicyBindingType))
	fieldNameMap["inShapingPolicy"] = "InShapingPolicy"
	fields["label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["label"] = "Label"
	fields["enableProxyArp"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enableProxyArp"] = "EnableProxyArp"
	fields["index"] = bindings.NewIntegerType()
	fieldNameMap["index"] = "Index"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["mtu"] = "Mtu"
	fields["fenceParameters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(KeyValueAttributesBindingType), reflect.TypeOf([]KeyValueAttributes{})))
	fieldNameMap["fenceParameters"] = "FenceParameters"
	fields["macAddresses"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(MacAddressBindingType), reflect.TypeOf([]MacAddress{})))
	fieldNameMap["macAddresses"] = "MacAddresses"
	fields["outShapingPolicy"] = bindings.NewOptionalType(bindings.NewReferenceType(TrafficShapingPolicyBindingType))
	fieldNameMap["outShapingPolicy"] = "OutShapingPolicy"
	fields["portgroupName"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["portgroupName"] = "PortgroupName"
	fields["enableBridgeMode"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enableBridgeMode"] = "EnableBridgeMode"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["portgroupId"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["portgroupId"] = "PortgroupId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vnic", fields, reflect.TypeOf(Vnic{}), fieldNameMap, validators)
}

func VnicsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vnics"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VnicBindingType), reflect.TypeOf([]Vnic{})))
	fieldNameMap["vnics"] = "Vnics"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vnics", fields, reflect.TypeOf(Vnics{}), fieldNameMap, validators)
}

func VpcInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vpc_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_cidr"] = "VpcCidr"
	fields["vgw_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vgw_id"] = "VgwId"
	fields["esx_public_security_group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_public_security_group_id"] = "EsxPublicSecurityGroupId"
	fields["vif_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vif_ids"] = "VifIds"
	fields["vm_security_group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vm_security_group_id"] = "VmSecurityGroupId"
	fields["tgwIps"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
	fieldNameMap["tgwIps"] = "TgwIps"
	fields["route_table_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["route_table_id"] = "RouteTableId"
	fields["edge_subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_subnet_id"] = "EdgeSubnetId"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["api_association_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["api_association_id"] = "ApiAssociationId"
	fields["api_subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["api_subnet_id"] = "ApiSubnetId"
	fields["private_subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["private_subnet_id"] = "PrivateSubnetId"
	fields["private_association_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["private_association_id"] = "PrivateAssociationId"
	fields["esx_security_group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_security_group_id"] = "EsxSecurityGroupId"
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["internet_gateway_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internet_gateway_id"] = "InternetGatewayId"
	fields["security_group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["security_group_id"] = "SecurityGroupId"
	fields["association_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["association_id"] = "AssociationId"
	fields["vgw_route_table_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vgw_route_table_id"] = "VgwRouteTableId"
	fields["edge_association_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_association_id"] = "EdgeAssociationId"
	fields["provider"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["provider"] = "Provider"
	fields["peering_connection_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["peering_connection_id"] = "PeeringConnectionId"
	fields["network_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_type"] = "NetworkType"
	fields["available_zones"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AvailableZoneInfoBindingType), reflect.TypeOf([]AvailableZoneInfo{})))
	fieldNameMap["available_zones"] = "AvailableZones"
	fields["routetables"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(RouteTableInfoBindingType),reflect.TypeOf(map[string]RouteTableInfo{})))
	fieldNameMap["routetables"] = "Routetables"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpc_info", fields, reflect.TypeOf(VpcInfo{}), fieldNameMap, validators)
}

func VpcInfoSubnetsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vpc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["cidr_block"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cidr_block"] = "CidrBlock"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["subnets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SubnetInfoBindingType), reflect.TypeOf([]SubnetInfo{})))
	fieldNameMap["subnets"] = "Subnets"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpc_info_subnets", fields, reflect.TypeOf(VpcInfoSubnets{}), fieldNameMap, validators)
}

func VpnBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["version"] = "Version"
	fields["on_prem_gateway_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["on_prem_gateway_ip"] = "OnPremGatewayIp"
	fields["on_prem_network_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["on_prem_network_cidr"] = "OnPremNetworkCidr"
	fields["pfs_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["pfs_enabled"] = "PfsEnabled"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["channel_status"] = bindings.NewOptionalType(bindings.NewReferenceType(VpnChannelStatusBindingType))
	fieldNameMap["channel_status"] = "ChannelStatus"
	fields["on_prem_nat_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["on_prem_nat_ip"] = "OnPremNatIp"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["internal_network_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["internal_network_ids"] = "InternalNetworkIds"
	fields["tunnel_statuses"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VpnTunnelStatusBindingType), reflect.TypeOf([]VpnTunnelStatus{})))
	fieldNameMap["tunnel_statuses"] = "TunnelStatuses"
	fields["encryption"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["encryption"] = "Encryption"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["dh_group"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dh_group"] = "DhGroup"
	fields["authentication"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["authentication"] = "Authentication"
	fields["pre_shared_key"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["pre_shared_key"] = "PreSharedKey"
	fields["ike_option"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ike_option"] = "IkeOption"
	fields["digest_algorithm"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["digest_algorithm"] = "DigestAlgorithm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpn", fields, reflect.TypeOf(Vpn{}), fieldNameMap, validators)
}

func VpnChannelStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["channel_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["channel_status"] = "ChannelStatus"
	fields["channel_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["channel_state"] = "ChannelState"
	fields["last_info_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["last_info_message"] = "LastInfoMessage"
	fields["failure_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["failure_message"] = "FailureMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpn_channel_status", fields, reflect.TypeOf(VpnChannelStatus{}), fieldNameMap, validators)
}

func VpnTunnelStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["on_prem_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["on_prem_subnet"] = "OnPremSubnet"
	fields["traffic_stats"] = bindings.NewOptionalType(bindings.NewReferenceType(VpnTunnelTrafficStatsBindingType))
	fieldNameMap["traffic_stats"] = "TrafficStats"
	fields["last_info_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["last_info_message"] = "LastInfoMessage"
	fields["local_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["local_subnet"] = "LocalSubnet"
	fields["tunnel_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tunnel_state"] = "TunnelState"
	fields["failure_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["failure_message"] = "FailureMessage"
	fields["tunnel_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tunnel_status"] = "TunnelStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpn_tunnel_status", fields, reflect.TypeOf(VpnTunnelStatus{}), fieldNameMap, validators)
}

func VpnTunnelTrafficStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["packets_out"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["packets_out"] = "PacketsOut"
	fields["packet_received_errors"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["packet_received_errors"] = "PacketReceivedErrors"
	fields["rx_bytes_on_local_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rx_bytes_on_local_subnet"] = "RxBytesOnLocalSubnet"
	fields["replay_errors"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["replay_errors"] = "ReplayErrors"
	fields["sequence_number_over_flow_errors"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sequence_number_over_flow_errors"] = "SequenceNumberOverFlowErrors"
	fields["encryption_failures"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["encryption_failures"] = "EncryptionFailures"
	fields["integrity_errors"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["integrity_errors"] = "IntegrityErrors"
	fields["packet_sent_errors"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["packet_sent_errors"] = "PacketSentErrors"
	fields["decryption_failures"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["decryption_failures"] = "DecryptionFailures"
	fields["packets_in"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["packets_in"] = "PacketsIn"
	fields["tx_bytes_from_local_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tx_bytes_from_local_subnet"] = "TxBytesFromLocalSubnet"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpn_tunnel_traffic_stats", fields, reflect.TypeOf(VpnTunnelTrafficStats{}), fieldNameMap, validators)
}

func VsanAvailableCapacityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cost"] = bindings.NewStringType()
	fieldNameMap["cost"] = "Cost"
	fields["quality"] = bindings.NewStringType()
	fieldNameMap["quality"] = "Quality"
	fields["size"] = bindings.NewIntegerType()
	fieldNameMap["size"] = "Size"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vsan_available_capacity", fields, reflect.TypeOf(VsanAvailableCapacity{}), fieldNameMap, validators)
}

func VsanClusterReconfigBiasBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["short_description"] = bindings.NewStringType()
	fieldNameMap["short_description"] = "ShortDescription"
	fields["full_description"] = bindings.NewStringType()
	fieldNameMap["full_description"] = "FullDescription"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vsan_cluster_reconfig_bias", fields, reflect.TypeOf(VsanClusterReconfigBias{}), fieldNameMap, validators)
}

func VsanClusterReconfigConstraintsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reconfig_biases"] = bindings.NewListType(bindings.NewReferenceType(VsanClusterReconfigBiasBindingType), reflect.TypeOf([]VsanClusterReconfigBias{}))
	fieldNameMap["reconfig_biases"] = "ReconfigBiases"
	fields["available_capacities"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewReferenceType(VsanAvailableCapacityBindingType), reflect.TypeOf([]VsanAvailableCapacity{})),reflect.TypeOf(map[string][]VsanAvailableCapacity{}))
	fieldNameMap["available_capacities"] = "AvailableCapacities"
	fields["default_capacities"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(VsanAvailableCapacityBindingType),reflect.TypeOf(map[string]VsanAvailableCapacity{}))
	fieldNameMap["default_capacities"] = "DefaultCapacities"
	fields["hosts"] = bindings.NewIntegerType()
	fieldNameMap["hosts"] = "Hosts"
	fields["default_reconfig_bias_id"] = bindings.NewStringType()
	fieldNameMap["default_reconfig_bias_id"] = "DefaultReconfigBiasId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vsan_cluster_reconfig_constraints", fields, reflect.TypeOf(VsanClusterReconfigConstraints{}), fieldNameMap, validators)
}

func VsanConfigConstraintsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["max_capacity"] = bindings.NewIntegerType()
	fieldNameMap["max_capacity"] = "MaxCapacity"
	fields["recommended_capacities"] = bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{}))
	fieldNameMap["recommended_capacities"] = "RecommendedCapacities"
	fields["supported_capacity_increment"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["supported_capacity_increment"] = "SupportedCapacityIncrement"
	fields["min_capacity"] = bindings.NewIntegerType()
	fieldNameMap["min_capacity"] = "MinCapacity"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vsan_config_constraints", fields, reflect.TypeOf(VsanConfigConstraints{}), fieldNameMap, validators)
}

func VsanEncryptionConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["port"] = "Port"
	fields["certificate"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["certificate"] = "Certificate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vsan_encryption_config", fields, reflect.TypeOf(VsanEncryptionConfig{}), fieldNameMap, validators)
}


