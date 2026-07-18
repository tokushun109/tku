output "state_bucket_name" {
  description = "Use this value in each environment's S3 backend configuration."
  value       = aws_s3_bucket.state.bucket
}

output "github_actions_opentofu_plan_role_arn" {
  description = "Set this ARN as the AWS_PLAN_ROLE_ARN GitHub Actions variable."
  value       = aws_iam_role.github_actions_opentofu_plan.arn
}
