resource "railway_project" "production" {
  name = "attractive-ambition"

  lifecycle {
    prevent_destroy = true
    ignore_changes  = all
  }
}

resource "railway_environment" "production" {
  name       = "production"
  project_id = railway_project.production.id

  lifecycle {
    prevent_destroy = true
    ignore_changes  = all
  }
}

resource "railway_service" "api" {
  name       = "tku"
  project_id = railway_project.production.id

  lifecycle {
    prevent_destroy = true
    ignore_changes  = all
  }
}

resource "railway_service" "mysql" {
  name       = "MySQL-Q0ak"
  project_id = railway_project.production.id

  lifecycle {
    prevent_destroy = true
    ignore_changes  = all
  }
}

resource "railway_service" "migration" {
  name       = "migrate"
  project_id = railway_project.production.id

  lifecycle {
    prevent_destroy = true
    ignore_changes  = all
  }
}

resource "railway_custom_domain" "api" {
  domain         = "api.tocoriri.com"
  environment_id = railway_environment.production.id
  service_id     = railway_service.api.id

  lifecycle {
    prevent_destroy = true
    ignore_changes  = all
  }
}
