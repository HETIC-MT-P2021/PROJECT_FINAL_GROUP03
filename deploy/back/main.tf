// Set requirements for provider
terraform {
  required_providers {
    heroku = {
      source  = "heroku/heroku"
      version = "~> 4.6"
    }
  }
}

// Global config
resource "heroku_config" "global" {
  vars = {
    LOG_LEVEL = "info"
  }

  sensitive_vars = {
    DISCORD_TOKEN = var.discord_token
    SERVER_ADDR_FRONT = var.server_addr_front
  }
}

// Heroku app creation
resource "heroku_app" "final_project_group3" {
  for_each = toset(var.environments)
  name   = format("%s-%s", var.app_name, each.value)
  config_vars = {
    GOVERSION: var.go_version
    APP_BASE: "back"
    PROJECT_PATH: "back"
    SERVER_ADDR_FRONT: heroku_config.global.sensitive_vars.SERVER_ADDR_FRONT
    DISCORD_TOKEN: heroku_config.global.sensitive_vars.DISCORD_TOKEN
  }
  region = "eu"
}

// App build
resource "heroku_build" "api_build" {
  for_each = heroku_app.final_project_group3
  app        = each.value.name
  buildpacks = ["https://github.com/timanovsky/subdir-heroku-buildpack", "https://github.com/heroku/heroku-buildpack-go"]

  source {
    url = "https://github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/archive/refs/tags/v0.0.1.tar.gz"
    version = "v0.0.1"
  }
  depends_on = [heroku_app.final_project_group3]
}

# Launch the app's web process by scaling-up
resource "heroku_formation" "api_formation" {
  for_each = heroku_app.final_project_group3
  app        = each.value.name
  type       = "web"
  quantity   = 1
  size       = "free"
  depends_on = [heroku_build.api_build]
}

// Database Addon
resource "heroku_addon" "api_pg" {
  for_each = heroku_app.final_project_group3
  app  = each.value.name
  plan = "heroku-postgresql:hobby-dev"
}

