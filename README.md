terraform-provider-dozens
=========================

Terraform provider plugin to access [Dozens](https://dozens.jp/) DNS API.

Usage
-------------------------

```
provider "dozens" {
  user = "YOUR_DOZENS_USERNAME"
  key = "YOUR_DOZENS_API_KEY"
}

resource "dozens_domain" "example_org" {
  name = "example.org"
  mail = "admin@example.org"
}

resource "dozens_record" "test_example_org" {
  depends_on = "dozens_domain.example_org"
  domain = "example.org"
  name = "test"
  address = "127.0.0.1"
  type = "A"
  ttl = "7200"
  priority = "10"
}
```

Installation
-------------------------

```
go get -u github.com/takebayashi/terraform-provider-dozens
```


License
-------------------------

[MIT License](./LICENSE)
