authorize:
	curl -X GET "http://localhost:8080/oauth/authorize?response_type=code&redirect_uri=http%3A%2F%2Flocalhost%3A8081%2Fcallback&client_id=web_client&scope=openid%20profile%20email"

m2m:
	curl -X POST http://localhost:8080/oauth/token \
  	-H "Content-Type: application/x-www-form-urlencoded" \
  	-d "grant_type=client_credentials" \
  	-d "client_id=web_client" \
  	-d "client_secret=axaa"
