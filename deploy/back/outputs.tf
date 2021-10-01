output "app_url" {
  value = {
    for k, v in heroku_app.final_project_group3 : k => format("https://%s.herokuapp.com", v.name)
  }
}
