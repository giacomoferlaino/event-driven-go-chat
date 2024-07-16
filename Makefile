# Variables
local_deployment_file=./deployments/local/docker-compose.yml
auth_app_path=app/auth

# Functions
define get_unit_tests
	
endef

# Commands
local-up: $(local_deployment_file)
	docker compose -f $(local_deployment_file) up -d

local-down: $(local_deployment_file)
	docker compose -f $(local_deployment_file) down

codegen: auth-codegen

test-unit: auth-test-unit

test-e2e: auth-test-e2e

test: test-unit test-e2e

auth-test-unit:
	go test chat/$(auth_app_path)/...

auth-test-e2e:
	go test chat/$(auth_app_path)/_e2e

auth-codegen:
	@echo "Generating auth module code"
	cd $(auth_app_path); go run github.com/99designs/gqlgen generate; go run github.com/Khan/genqlient