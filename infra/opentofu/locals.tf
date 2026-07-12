locals {
  default_tags = {
    Environment = var.environment
    ManagedBy   = "opentofu"
    Project     = var.project
  }
}
