variable "tags" {
  description = "[Required unless var.context used] Default tags."
  type        = "map"
  default     = {}
}

variable "namespace" {
  description = "[Required unless var.context used] The namespace or group used for all AWS resource names."
  default     = ""
}

variable "name" {
  description = "[Required unless var.context used] The name of this project."
  default     = ""
}

variable "environment" {
  description = "[Required unless var.context used] The tenant environment. One of production, uat or development."
  default     = ""
}

variable "attributes" {
  type        = "list"
  description = "[Required unless var.context used] Any extra attributes for tagging or defining these resources."
  default     = []
}

variable "context" {
  type        = "map"
  description = "[Required unless var.context used] Used to pass in other label module context."
  default     = {}
}

module "label" {
  source      = "git::https://github.com/cloudposse/terraform-null-label.git?ref=master"
  attributes  = ["${var.attributes}"]
  namespace   = "${var.namespace}"
  environment = "${var.environment}"
  name        = "${var.name}"
  context     = "${var.context}"
  tags        = "${var.tags}"
}
