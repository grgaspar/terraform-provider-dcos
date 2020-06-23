output "setId" {
 value = "${dcos_ham_app.my-server.id}"
}

output "pokemon-name" {
 value = "${dcos_ham_app.my-server.gsil_pokemon_name}"
}

output "injector-token" {
 value = "${dcos_ham_app.my-server.gsil_injector_token}"
}

output "intent-manager-url" {
 value = "${dcos_ham_app.my-server.gsil_intent_manager_url}"
}

output "register-key-response" {
  value = "${dcos_ham_app.my-server.gsil_register_key_response}"
}