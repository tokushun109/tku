output "state_bucket_name" {
  description = "Use this value in each environment's S3 backend configuration."
  value       = aws_s3_bucket.state.bucket
}
