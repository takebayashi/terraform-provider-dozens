terraform-provider-dozens
=========================

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
```

Installation
-------------------------
```
go get -u github.com/takebayashi/terraform-provider-dozens
```
