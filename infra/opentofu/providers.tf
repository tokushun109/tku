provider "aws" {
  region = var.aws_region

  default_tags {
    tags = local.default_tags
  }
}

provider "aws" {
  alias  = "singapore"
  region = "ap-southeast-1"

  default_tags {
    tags = local.default_tags
  }
}

provider "railway" {}
