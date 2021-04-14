# create cloudexport for AWS
resource "kentik-cloudexport_item" "terraform_aws_export" {
  name           = "test_terraform_aws_export"
  type           = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
  enabled        = true
  description    = "terraform aws cloud export"
  plan_id        = "11467"
  cloud_provider = "aws"
  aws {
    bucket            = "terraform-aws-bucket"
    iam_role_arn      = "arn:aws:iam::003740049406:role/trafficTerraformIngestRole"
    region            = "us-east-2"
    delete_after_read = false
    multiple_buckets  = false
  }
}

output "aws" {
  value = kentik-cloudexport_item.terraform_aws_export
}

# create cloudexport for IBM
resource "kentik-cloudexport_item" "terraform_ibm_export" {
  name           = "test_terraform_ibm_export"
  type           = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
  enabled        = false
  description    = "terraform ibm cloud export"
  plan_id        = "11467"
  cloud_provider = "ibm"
  ibm {
    bucket = "terraform-ibm-bucket"
  }
}

output "ibm" {
  value = kentik-cloudexport_item.terraform_ibm_export
}

# create cloudexport for GCE
resource "kentik-cloudexport_item" "terraform_gce_export" {
  name           = "test_terraform_gce_export"
  type           = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
  enabled        = false
  description    = "terraform gce cloud export"
  plan_id        = "11467"
  cloud_provider = "gce"
  gce {
    project      = "project gce"
    subscription = "subscription gce"
  }
}

output "gce" {
  value = kentik-cloudexport_item.terraform_gce_export
}
