output "people_service_url" {
  description = "URL for People Service"
  value       = google_cloud_run_service.people_service.status.url
}

output "contract_service_url" {
  description = "URL for Contract Service"
  value       = google_cloud_run_service.contract_service.status.url
}

output "property_service_url" {
  description = "URL for Property Service"
  value       = google_cloud_run_service.property_service.status.url
}

output "load_balancer_url" {
  description = "Load balancer URL for all services"
  value       = google_compute_global_forwarding_rule.default.self_link
}
