.PHONY: run-servers stop-servers test-servers docker compose-up compose-down install-dev docker-clean

run-servers:
	@scripts/run-servers.sh

stop-servers:
	@scripts/stop-servers.sh

test-servers:
	@scripts/run-servers.sh
	@scripts/test-e2e.sh
	@scripts/stop-servers.sh

docker-build:
	@scripts/docker-build.sh

docker-clean:
	@scripts/docker-clean.sh

compose-up:
	@scripts/docker-compose.sh

compose-down:
	docker-compose -f docker-compose.yml -f docker-compose.test.yml down -v

install-dev:
	@scripts/install-dev.sh