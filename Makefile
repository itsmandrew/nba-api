
# Configurable API URL and API key
API_URL ?= http://localhost:8080/v1
API_KEY ?= API_KEY
PRIVATE_ENDPOINT ?= players/random
TOKEN_LIFETIME ?= 5 # seconds
RATE_LIMIT_TEST_REQUESTS ?= 25
BASE_URL ?= localhost:8080/v1/players

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


# Test rate limiter by making repeated calls
test-ratelimit:
	@echo "Generating JWT for rate limit test..."
	@TOKEN=$$(curl -s -X POST $(API_URL)/generate-token \
		-H "X-API-KEY: $(API_KEY)" | jq -r .token) && \
	echo "Token: $$TOKEN" && \
	echo "Sending $(RATE_LIMIT_TEST_REQUESTS) requests to $(PRIVATE_ENDPOINT)..." && \
	for i in $$(seq 1 $(RATE_LIMIT_TEST_REQUESTS)); do \
		echo "Request $$i:"; \
		curl -s -o /dev/null -w "HTTP %{http_code}\n" $(API_URL)/$(PRIVATE_ENDPOINT) \
			-H "Authorization: Bearer $$TOKEN"; \
	done



test-pagination:
	@echo "Testing pagination"
	@TOKEN=$$(curl -s -X POST $(API_URL)/generate-token \
		-H "X-API-KEY: $(API_KEY)" | jq -r .token) && \
	echo "Token: $$TOKEN" && \
	echo "Fetching page 1..." && \
	curl -s "$(BASE_URL)?limit=10&page=1" -H "Authorization: Bearer $$TOKEN" | jq . && \
	echo "\nFetching page 2..." && \
	curl -s "$(BASE_URL)?limit=10&page=2" -H "Authorization: Bearer $$TOKEN" | jq .
