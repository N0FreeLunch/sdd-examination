resource "aws_s3_bucket" "backup_bucket" {
  bucket_prefix = "sdd-exam-backup-"
}

resource "aws_s3_bucket_versioning" "backup_bucket_ver" {
  bucket = aws_s3_bucket.backup_bucket.id
  versioning_configuration {
    status = "Enabled"
  }
}

# IAM User for Application (Litestream)
resource "aws_iam_user" "app_user" {
  name = "sdd-exam-app-user"
}

resource "aws_iam_access_key" "app_user_key" {
  user = aws_iam_user.app_user.name
}

resource "aws_iam_user_policy" "app_user_policy" {
  name = "sdd-exam-app-policy"
  user = aws_iam_user.app_user.name

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:PutObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:DeleteObject"
        ]
        Effect   = "Allow"
        Resource = [
          aws_s3_bucket.backup_bucket.arn,
          "${aws_s3_bucket.backup_bucket.arn}/*"
        ]
      }
    ]
  })
}

output "app_user_access_key" {
  value = aws_iam_access_key.app_user_key.id
}

output "app_user_secret_key" {
  value     = aws_iam_access_key.app_user_key.secret
  sensitive = true
}

output "backup_bucket_name" {
  value = aws_s3_bucket.backup_bucket.bucket
}
