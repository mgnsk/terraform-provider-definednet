variable "definednet_token" {
  description = "Defined.net HTTP API token"
  sensitive   = true
}

provider "definednet" {
  token = var.definednet_token
}

resource "definednet_host" "example" {
  name       = "example.defined.test"
  network_id = "network-7P81MCS2TVAY9XJWQTNJ3PWYPD"
  role_id    = "role-WSG78880Z655TQJVQFL5CZ405B"
  tags       = ["service:app"]

  metrics {
    listen               = "127.0.0.1:8080"
    path                 = "/metrics"
    namespace            = "nebula"
    subsystem            = "host"
    enable_extra_metrics = true
  }
}
