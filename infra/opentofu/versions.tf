terraform {
  required_version = ">= 1.8.0"

  backend "s3" {}

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }

    railway = {
      source  = "terraform-community-providers/railway"
      version = "~> 0.6"
    }
  }
}
