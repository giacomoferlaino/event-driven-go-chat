local_deployment_file=./deployments/local/docker-compose.yml

local-up: $(local_deployment_file)
	docker compose -f $(local_deployment_file) up -d

local-down: $(local_deployment_file)
	docker compose -f $(local_deployment_file) down