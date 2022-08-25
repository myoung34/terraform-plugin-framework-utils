terraform {
  required_version = "~> 1.1.7"

  required_providers {
    test = {
      source  = "local/local/test"
      version = "1.0.0"
    }
  }
}
