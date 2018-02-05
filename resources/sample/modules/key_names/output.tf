output "key_name" {
   value = "${var.aws_region == "eu-west-1" ? local.key_eu :  local.key_other}"
}
