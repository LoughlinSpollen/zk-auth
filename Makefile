.PHONY: run-servers stop-servers docker-build docker-clean compose-up compose-down install-dev

run-servers:
	@scripts/run-servers.sh

stop-servers:
	@scripts/stop-servers.sh

docker-build:
	@scripts/docker-build.sh

docker-clean:
	@scripts/docker-clean.sh

compose-up:
	@scripts/docker-compose.sh

compose-down:
	docker-compose -f docker-compose.yml -f down -v
	docker-compose down --volumes

install-dev:
	@scripts/install-dev.sh