provider "google" {
  project = var.project_id
  region  = var.region
}

# Create a VPC Network for Cloud SQL
resource "google_compute_network" "vpc_network" {
  name = "debitor-case-network"
}

resource "google_compute_subnetwork" "subnetwork" {
  name          = "debitor-case-subnetwork"
  ip_cidr_range = "10.0.0.0/16"
  network       = google_compute_network.vpc_network.name
}

# Cloud SQL Instances and Cloud Run Services

## Cloud SQL for People-Service
resource "google_sql_database_instance" "people_db_instance" {
  name             = "people-db-instance"
  database_version = "POSTGRES_12"
  region           = var.region
  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled    = true
      private_network = google_compute_network.vpc_network.id
    }
  }
}

resource "google_sql_database" "people_db" {
  name     = "people_db"
  instance = google_sql_database_instance.people_db_instance.name
}

resource "google_sql_user" "people_db_user" {
  instance = google_sql_database_instance.people_db_instance.name
  name     = "people_user"
  password = "people_password"
}

resource "google_cloud_run_service" "people_service" {
  name     = "people-service"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/people-service:latest"
        ports {
          container_port = 8081
        }
        env {
          name  = "DB_HOST"
          value = google_sql_database_instance.people_db_instance.connection_name
        }
        env {
          name  = "DB_USER"
          value = google_sql_user.people_db_user.name
        }
        env {
          name  = "DB_PASSWORD"
          value = google_sql_user.people_db_user.password
        }
        env {
          name  = "DB_NAME"
          value = google_sql_database.people_db.name
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

## Cloud SQL for Contract-Service
resource "google_sql_database_instance" "contract_db_instance" {
  name             = "contract-db-instance"
  database_version = "POSTGRES_12"
  region           = var.region
  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled    = true
      private_network = google_compute_network.vpc_network.id
    }
  }
}

resource "google_sql_database" "contract_db" {
  name     = "contract_db"
  instance = google_sql_database_instance.contract_db_instance.name
}

resource "google_sql_user" "contract_db_user" {
  instance = google_sql_database_instance.contract_db_instance.name
  name     = "contract_user"
  password = "contract_password"
}

resource "google_cloud_run_service" "contract_service" {
  name     = "contract-service"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/contract-service:latest"
        ports {
          container_port = 8082
        }
        env {
          name  = "DB_HOST"
          value = google_sql_database_instance.contract_db_instance.connection_name
        }
        env {
          name  = "DB_USER"
          value = google_sql_user.contract_db_user.name
        }
        env {
          name  = "DB_PASSWORD"
          value = google_sql_user.contract_db_user.password
        }
        env {
          name  = "DB_NAME"
          value = google_sql_database.contract_db.name
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

## Cloud SQL for Property-Service
resource "google_sql_database_instance" "property_db_instance" {
  name             = "property-db-instance"
  database_version = "POSTGRES_12"
  region           = var.region
  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled    = true
      private_network = google_compute_network.vpc_network.id
    }
  }
}

resource "google_sql_database" "property_db" {
  name     = "property_db"
  instance = google_sql_database_instance.property_db_instance.name
}

resource "google_sql_user" "property_db_user" {
  instance = google_sql_database_instance.property_db_instance.name
  name     = "property_user"
  password = "property_password"
}

resource "google_cloud_run_service" "property_service" {
  name     = "property-service"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/property-service:latest"
        ports {
          container_port = 8083
        }
        env {
          name  = "DB_HOST"
          value = google_sql_database_instance.property_db_instance.connection_name
        }
        env {
          name  = "DB_USER"
          value = google_sql_user.property_db_user.name
        }
        env {
          name  = "DB_PASSWORD"
          value = google_sql_user.property_db_user.password
        }
        env {
          name  = "DB_NAME"
          value = google_sql_database.property_db.name
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

# Backend Service for Cloud Run Services
resource "google_compute_backend_service" "cloud_run_backend" {
  name                  = "debitor-case-backend"
  load_balancing_scheme = "EXTERNAL"

  dynamic "backend" {
    for_each = {
      "people-service"    = google_cloud_run_service.people_service.status.url
      "contract-service"  = google_cloud_run_service.contract_service.status.url
      "property-service"  = google_cloud_run_service.property_service.status.url
    }
    content {
      group = backend.value
    }
  }

  health_checks = [google_compute_health_check.http_health_check.self_link]
}

# HTTP(S) Load Balancer
resource "google_compute_url_map" "url_map" {
  name = "debitor-case-url-map"

  default_service = google_compute_backend_service.cloud_run_backend.self_link

  host_rule {
    hosts        = ["*"]
    path_matcher = "all-paths"
  }

  path_matcher {
    name            = "all-paths"
    default_service = google_compute_backend_service.cloud_run_backend.self_link

    path_rule {
      paths   = ["/people/*"]
      service = google_cloud_run_service.people_service.status.url
    }

    path_rule {
      paths   = ["/contracts/*"]
      service = google_cloud_run_service.contract_service.status.url
    }

    path_rule {
      paths   = ["/properties/*"]
      service = google_cloud_run_service.property_service.status.url
    }
  }
}

# Global Forwarding Rule
resource "google_compute_global_forwarding_rule" "default" {
  name       = "debitor-case-forwarding-rule"
  ip_protocol = "TCP"
  port_range  = "80"

  target = google_compute_url_map.url_map.self_link
}
