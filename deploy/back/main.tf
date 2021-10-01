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
    PRIVATE_KEY = "some_private_key"
  }
}

// Heroku app creation
resource "heroku_app" "final_project_group3" {
  name   = var.app_name
  config_vars = {
    GOVERSION: var.go_version
    APP_BASE: "back"
    PROJECT_PATH: "back"
    SERVER_ADDR_FRONT: var.server_addr_front
    DISCORD_TOKEN: var.discord_token
  }
  region = "eu"
}

// App build
resource "heroku_build" "api_build" {
  app        = heroku_app.final_project_group3.name
  buildpacks = ["https://github.com/timanovsky/subdir-heroku-buildpack", "https://github.com/heroku/heroku-buildpack-go"]

  source {
    url = "https://github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/archive/refs/tags/v0.0.1.tar.gz"
    version = "v0.0.1"
  }
  depends_on = [heroku_app.final_project_group3]
}

# Launch the app's web process by scaling-up
resource "heroku_formation" "api_formation" {
  app        = heroku_app.final_project_group3.name
  type       = "web"
  quantity   = 1
  size       = "Standard-1x"
  depends_on = [heroku_build.api_build]
}

// Database Addon
resource "heroku_addon" "api_pg" {
  app  = heroku_app.final_project_group3.name
  plan = "heroku-postgresql:hobby-dev"
}

