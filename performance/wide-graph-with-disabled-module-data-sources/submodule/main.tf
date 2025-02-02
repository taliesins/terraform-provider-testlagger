terraform {
  required_providers {
    testlagger = {
      source  = "taliesins/testlagger"
      version = "1.0.0"
    }
  }
}

data "testlagger_lag" "test" {
  read_delay = 1000
  input      = "hello"
}
