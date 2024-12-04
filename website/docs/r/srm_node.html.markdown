---
layout: "vmc"

page_title: "VMC: vmc_srm_node"
sidebar_current: "docs-vmc_srm_node"

description: |-
  Provides a resource to add an instance to SDDC after site recovery has been activated.
---

# vmc_srm_node

 Provides a resource to add an instance to SDDC after site recovery has been activated.
~> **Note:** SRM node resource depends on site recovery resource creation. Site recovery must be activated to add SRM node instance. For details on how to activate site recovery refer to the site recovery resource [vmc_site_recovery](https://www.terraform.io/docs/providers/vmc/r/site_recovery.html).

## Example Usage

```hcl

provider "vmc" {
  refresh_token = var.api_token
  org_id = var.org_id
}

resource "vmc_srm_node" "srm_node_1"{
  sddc_id = vmc_sddc.sddc_1.id
  srm_node_extension_key_suffix = var.srm_node_srm_extension_key_suffix
  depends_on = [vmc_site_recovery.site_recovery_1]
}

```

## Argument Reference

The following arguments are supported for vmc_srm_node resource:

* `sddc_id` - (Required) SDDC identifier.

* `srm_node_extension_key_suffix` - (Required) Custom extension key suffix for SRM. If not specified, default extension key will be used. 
The custom extension suffix must contain 13 characters or fewer, be composed of letters, numbers, ., - characters. 
The extension suffix must begin and end with a letter or number. The suffix is appended to com.vmware.vcDr- to form the full extension key.

## Attributes Reference

In addition to arguments listed above, the following attributes are exported after site recovery activation:

* `site_recovery_state` - Site recovery state. Possible values are: ACTIVATED, ACTIVATING, CANCELED, DEACTIVATED, DEACTIVATING, DELETED, FAILED.

* `srm_node` - Site recovery node information.

* `vr_node` - VR node information.

## Import

SRM node resource can be imported using the `id` and `sddc_id` , e.g.

`$ terraform import vmc_srm_node.srm_node_1 id,sddc_id`

- id = SRM Node Identifier
- sddc_id = SDDC Identifier


`$ terraform import vmc_srm_node.srm_node_1 7aad97e9-9a4f-4e43-8817-5c8d8c0e87a5,afe7a0fd-3f0a-48b2-9ddb-0489c22732ae`