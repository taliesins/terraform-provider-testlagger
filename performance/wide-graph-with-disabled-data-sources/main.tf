terraform {
  required_providers {
    testlagger = {
      source  = "taliesins/testlagger"
      version = "1.0.0"
    }
  }
}

provider "testlagger" {
  client_initialize_delay     = 1000
  datasource_configure_delay  = 1000
  resource_configure_delay    = 1000
  resource_import_state_delay = 1000
}

module "generated-1" {
  source = "./generated"
}

module "generated-2" {
  source = "./generated"
}

module "generated-3" {
  source = "./generated"
}

module "generated-4" {
  source = "./generated"
}

module "generated-5" {
  source = "./generated"
}

module "generated-6" {
  source = "./generated"
}

module "generated-7" {
  source = "./generated"
}

module "generated-8" {
  source = "./generated"
}

module "generated-9" {
  source = "./generated"
}

module "generated-10" {
  source = "./generated"
}

