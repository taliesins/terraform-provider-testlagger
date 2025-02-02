terraform {
  required_providers {
    testlagger = {
      source  = "taliesins/testlagger"
      version = "1.0.0"
    }
  }
}

resource "testlagger_lag" "iter" {
  create_delay = 1000
  read_delay   = 1000
  update_delay = 1000
  delete_delay = 1000
  input        = "hello"
}
