# Variables
local_deployment_file=./deployments/local/docker-compose.yml
e2e_deployment_file=./deployments/e2e/docker-compose.yml
auth_app_path=app/auth
chat_app_path=app/chat

# Commands
local-up: $(local_deployment_file)
	docker compose -f $(local_deployment_file) up -d

local-down: $(local_deployment_file)
	docker compose -f $(local_deployment_file) down

e2e-up: $(e2e_deployment_file)
	docker compose -f $(e2e_deployment_file) up -d

e2e-down: $(e2e_deployment_file)
	docker compose -f $(e2e_deployment_file) down

e2e-restart: e2e-down e2e-up

codegen: auth-codegen chat-codegen

test-unit: auth-test-unit chat-test-unit
	go test ./pkg/...

test-e2e: auth-test-e2e chat-test-e2e

test: test-unit test-e2e

auth-test-unit:
	go test chat/$(auth_app_path)/...

auth-test-e2e:
	go test chat/$(auth_app_path)/_e2e

auth-codegen:
	@echo "Generating auth module code"
	cd $(auth_app_path); go run github.com/99designs/gqlgen generate; go run github.com/Khan/genqlient

chat-test-unit:
	go test chat/$(chat_app_path)/...

chat-test-e2e:
	go test chat/$(chat_app_path)/_e2e

chat-codegen:
	@echo "Generating chat module code"
	cd $(chat_app_path); go run github.com/99designs/gqlgen generate; go run github.com/Khan/genqlient