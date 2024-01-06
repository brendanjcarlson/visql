CURRENT_TIMESTAMP := $(shell date +%Y_%m_%d_%H:%M:%S)
SERVER_DIR := server
SERVER_TEST_RESULT_DIR := test_result
CLIENT_DIR := client

.PHONY: devpg_up
devpg_up:
	@docker compose up -d --build devpg

.PHONY: devpg_down
devpg_down:
	@docker compose down devpg

.PHONY: devpg_reset
devpg_reset:
	@docker compose down devpg
	@docker compose up -d --build devpg

.PHONY: devpg_logs
devpg_logs:
	@docker compose logs -f -t -n 100 -f devpg

.PHONY: devpg_psql
devpg_psql:
	@docker compose exec devpg psql -U postgres

.PHONY: run_server
run_server:
	@go -C $(SERVER_DIR) run ./src/cmd/main.go

.PHONY: build_server
build_server:
	@go -C $(SERVER_DIR) build -o ./bin/server ./src/cmd/main.go

.PHONY: test_server
test_server:
	@go -C $(SERVER_DIR) test ./src/... -v > ./$(SERVER_DIR)/$(SERVER_TEST_RESULT_DIR)/$(CURRENT_TIMESTAMP).log
	@cat ./$(SERVER_DIR)/$(SERVER_TEST_RESULT_DIR)/$(CURRENT_TIMESTAMP).log

.PHONY: run_client
run_client:
	@cd $(CLIENT_DIR) && npm run dev

.PHONY: build_client
build_client:
	@cd $(CLIENT_DIR) && npm run build

.PHONY: test_client
test_client:
	@cd $(CLIENT_DIR) && npm run test
