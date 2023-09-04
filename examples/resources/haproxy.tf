resource "haproxy_userlist" "userslist" {
  name = "userslist"
}

resource "haproxy_user" "sepehr" {
  username        = "sepehr"
  userlist        = haproxy_userlist.userslist.name
  password        = "123456999"
  secure_password = true
  depends_on      = [haproxy_userlist.userslist]
}

data "haproxy_user" "sepehr" {
  username = "sepehr"
  userlist = haproxy_userlist.userslist.name
}

output "sepehr" {
  value = data.haproxy_user.sepehr
}
