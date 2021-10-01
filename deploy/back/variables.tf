// Mandatory argument to set the "app_name"
variable "app_name" {
  description = "Name of the Heroku app"
  type = string
  default = "test_app"
}

// Mandatory
variable "discord_token" {
  description = "Discord token"
  type = string
}

// Mandatory
variable "server_addr_front" {
  description = "Address of the front-end"
  type = string
}

variable "go_version" {
  description = "Version of golang"
  type = string
  default = 1.17
}

variable "environments" {
  description = "App environment"
  type = list(string)
  default = ["dev", "prod"]
}

