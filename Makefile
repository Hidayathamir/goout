# Lint using golangci-lint.
lint:
	golangci-lint run --config .golangci.yml ./...

###################################

# Remove docker image with tag None.
clear-none-docker-images:
	docker images --filter "dangling=true" -q --no-trunc | xargs docker rmi

###################################

# Run postgres and redis container.
compose-up-postgres-redis:
	docker compose up -d goout-db-postgres goout-cache-redis

compose-down-postgres-redis:
	docker compose down goout-db-postgres goout-cache-redis

###################################

# Run deployment.
deploy:
	docker compose up --build

undeploy:
	docker compose down

###################################

migrate-up:
	go run internal/repo/db/migration/migrate_up.go

###################################

.SILENT:godoc
godoc:
	echo "" && \
	echo "Please go to link below to see documentation" && \
	echo http://localhost:7011/pkg/github.com/Hidayathamir/goout/?m=all && \
	echo "" && \
	godoc -http localhost:7011
