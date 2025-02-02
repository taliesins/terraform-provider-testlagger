terraform {
  required_providers {
    testlagger = {
      source  = "taliesins/testlagger"
      version = "1.0.0"
    }
  }
}

resource "testlagger_lag" "iter1" {
  count        = 1
  create_delay = 1000
  read_delay   = 1000
  update_delay = 1000
  delete_delay = 1000
  input        = "hello"
}

resource "testlagger_lag" "iter2" {
  count        = var.enabled ? 1 : 0
  create_delay = 10000
  read_delay   = 10000
  update_delay = 10000
  delete_delay = 10000
  input        = "hello"
}

resource "testlagger_lag" "iter3" {
  count        = var.enabled ? 1 : 0
  create_delay = 1000
  read_delay   = 1000
  update_delay = 1000
  delete_delay = 1000
  input        = "hello"
}

resource "testlagger_lag" "iter4" {
  count        = var.enabled ? 1 : 0
  create_delay = 1000
  read_delay   = 1000
  update_delay = 1000
  delete_delay = 1000
  input        = "hello"
}

resource "testlagger_lag" "iter5" {
  count        = var.enabled ? 1 : 0
  create_delay = 1000
  read_delay   = 1000
  update_delay = 1000
  delete_delay = 1000
  input        = "hello"
}

resource "testlagger_lag" "iter6" {
  count        = var.enabled ? 1 : 0
  create_delay = 1000
  read_delay   = 1000
  update_delay = 1000
  delete_delay = 1000
  input        = "hello"
}

resource "testlagger_lag" "iter7" {
  count        = var.enabled ? 1 : 0
  create_delay = 1000
  read_delay   = 1000
  update_delay = 1000
  delete_delay = 1000
  input        = "hello"
}

variable "enabled" {
  type = bool
  default = false
}
