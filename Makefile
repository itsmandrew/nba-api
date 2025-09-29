
# Configurable API URL and API key
API_URL ?= http://localhost:8080/v1
API_KEY ?= NBA_API_KEY
PRIVATE_ENDPOINT ?= players/random
TOKEN_LIFETIME ?= 5 # seconds

.PHONY: test-expiration call-private

# Generate token and call a private endpoint immediately
call-private:
	@echo "Generating JWT and calling private endpoint..."
	@TOKEN=$$(curl -s -X POST $(API_URL)/generate-token \
		-H "X-API-KEY: $(API_KEY)" \
	| jq -r .token) && \
	echo "Token: $$TOKEN" && \
	curl -s $(API_URL)/$(PRIVATE_ENDPOINT) \
		-H "Authorization: Bearer $$TOKEN" | jq .

# Test token expiration workflow
test-expiration:
	@echo "Generating short-lived JWT for expiration test..."
	@TOKEN=$$(curl -s -X POST $(API_URL)/generate-token \
		-H "X-API-KEY: $(API_KEY)" \
	| jq -r .token) && \
	echo "Token: $$TOKEN" && \
	echo "Calling private endpoint immediately..." && \
	curl -s $(API_URL)/$(PRIVATE_ENDPOINT) \
		-H "Authorization: Bearer $$TOKEN" | jq . && \
	echo "Sleeping $(TOKEN_LIFETIME) seconds for token to expire..." && \
	sleep $(TOKEN_LIFETIME) && \
	echo "Calling private endpoint after expiration..." && \
	curl -s $(API_URL)/$(PRIVATE_ENDPOINT) \
		-H "Authorization: Bearer $$TOKEN" | jq .

