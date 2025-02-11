---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zabbix_template Resource - terraform-provider-zabbix"
subcategory: ""
description: |-
  
---

# zabbix_template (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `groups` (Set of String) Host Group IDs
- `host` (String) Template hostname (internal name)

### Optional

- `description` (String) Template description
- `macro` (Block List) (see [below for nested schema](#nestedblock--macro))
- `name` (String) Template Display Name (defaults to host)
- `templates` (Set of String) linked templates

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--macro"></a>
### Nested Schema for `macro`

Required:

- `name` (String) Macro Name (key)
- `value` (String) Macro Value

Read-Only:

- `id` (String)
