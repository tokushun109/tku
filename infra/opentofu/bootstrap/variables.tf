variable "aws_region" {
  description = "Region used for the OpenTofu state bucket."
  type        = string
  default     = "ap-northeast-1"
}

variable "project" {
  description = "Project name used for the state bucket tags."
  type        = string
  default     = "tku"
}

variable "state_bucket_name" {
  description = "Globally unique S3 bucket name for OpenTofu state."
  type        = string
}
