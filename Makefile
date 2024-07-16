local_deployment_file=./deployments/local/docker-compose.yml

local-up: $(local_deployment_file)
	docker compose -f $(local_deployment_file) up -d

local-down: $(local_deployment_file)
	docker compose -f $(local_deployment_file) down

test-unit: auth-test-unit

test-e2e: auth-test-e2e

test: test-unit test-e2e

auth-test-unit:
	go test -short chat/internal/app/auth/...

auth-test-e2e:
	go test chat/internal/app/auth/...