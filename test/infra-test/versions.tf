terraform {
  required_version = "~> 1.3.5"

  required_providers {
    test = {
      source  = "local/local/test"
      version = "1.0.0"
    }
  }
}
