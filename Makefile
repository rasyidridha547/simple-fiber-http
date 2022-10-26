dev:
	@docker-compose --env-file ./.env up --build && docker-compose rm -fsv
local:
	set -a; source .env; set +a && go run .