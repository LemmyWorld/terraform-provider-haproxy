---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "haproxy_userlist Data Source - terraform-provider-haproxy"
subcategory: ""
description: |-
  
---

# haproxy_userlist (Data Source)



## Example Usage

```terraform
data "haproxy_userlist" "user1" {
  name = haproxy_userlist.userslist.name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the Userlist. It must be unique

### Read-Only

- `id` (String) The ID of this resource.