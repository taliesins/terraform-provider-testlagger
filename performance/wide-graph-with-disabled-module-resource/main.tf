variable "enabled" {
  type = bool
  default = false
}

module "generated-1" {
  count  = var.enabled ? 1 : 0
  source = "./generated"
}

module "generated-2" {
  count  = var.enabled ? 1 : 0
  source = "./generated"
}

module "generated-3" {
  count  = var.enabled ? 1 : 0
  source = "./generated"
}

module "generated-4" {
  count  = var.enabled ? 1 : 0
  source = "./generated"
}

module "generated-5" {
  count  = var.enabled ? 1 : 0
  source = "./generated"
}

module "generated-6" {
  count  = var.enabled ? 1 : 0
  source = "./generated"
}

module "generated-7" {
  count  = var.enabled ? 1 : 0
  source = "./generated"
}

module "generated-8" {
  count  = var.enabled ? 1 : 0
  source = "./generated"
}

module "generated-9" {
  count  = var.enabled ? 1 : 0
  source = "./generated"
}

module "generated-10" {
  count  = var.enabled ? 1 : 0
  source = "./generated"
}
