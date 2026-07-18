variable "aws_region" {
  description = "AWS resources are operated in this region."
  type        = string
  default     = "ap-northeast-1"
}

variable "environment" {
  description = "Deployment environment name."
  type        = string

  validation {
    condition     = contains(["development", "production"], var.environment)
    error_message = "environment must be either development or production."
  }
}

variable "project" {
  description = "Project name used for shared resource tags."
  type        = string
  default     = "tku"
}
