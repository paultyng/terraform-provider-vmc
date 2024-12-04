---
layout: "vmc"
page_title: "VMC: customer_subnets"
sidebar_current: "docs-vmc-datasource-customer-subnets"
description: A customer subnets data source.
---

# vmc_customer_subnets

The customer subnets data source provides information about customer's compatible subnets for account linking.
## Example Usage

```hcl
data "vmc_customer_subnets" "my_subnets" {
  connected_account_id = data.vmc_connected_accounts.my_accounts.id
  region               = var.sddc_region
}
```

## Argument Reference

* `org_id` - (Computed) Organization identifier.

* `region` - (Required) The AWS specific (e.g. us-west-2) or VMC specific region (e.g. US_WEST_2) of the cloud resources to work in.

* `num_hosts` - (Optional) The number of hosts.

* `connected_account_id` - (Required) The linked connected account identifier.

* `sddc_id` - (Optional) SDDC identifier.

* `force_refresh` - (Optional) When true, forces the mappings for datacenters to be refreshed for the connected account.

* `instance_type` - (Optional) The server instance type to be used.

* `sddc_type` - (Optional) The sddc type to be used. (1NODE, SingleAZ, MultiAZ)

## Attributes Reference

In addition to arguments listed above, the following attributes are exported:

* `customer_available_zones` - A list of AWS availability zones.

* `ids` - A list of AWS subnet IDs to create links to in the customer's account.