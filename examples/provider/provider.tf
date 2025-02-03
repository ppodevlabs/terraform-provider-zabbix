provider "zabbix" {
  # Required
  url = "http://example.com/api_jsonrpc.php"
  
  # Optional

  # use one of username/password or token
  username = "<api_user>"
  password = "<api_password>"

  token = "<api_token>"

  # Disable TLS verfication (false by default)
  tls_insecure = true

  # Serialize Zabbix API calls (false by default)
  # Note: race conditions have been observed, enable this if required
  serialize = true
}

