# Specify the provider and access details
provider "aws" {
  region = "us-east-1"
}

resource "aws_security_group" "security_mongodb" {
  name        = "terraform-${var.region}-${var.env}-mongodb"
  description = "Testing"

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "ALLOW http ALL"
  }

  # outbound internet access
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow all"
  }
  tags {
    Name =  "terraform-${var.region}-${var.env}-mongodb"
    Terraform = "true"
  }
}
