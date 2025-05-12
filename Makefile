ifeq (revision,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for run
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ... and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

.PHONY:

start:
	go run src/main.go


db:
	docker run --name=api_db \
			-e SSL_MODE='disable'\
		-e POSTGRES_USER=$$PG_USER\
		-e POSTGRES_PASSWORD=$$PG_PASSWORD\
		-e POSTGRES_DB=$$PG_DB_NAME\
		-e TZ=GMT-3\
		-p $$PG_PORT:5432 -d --rm postgres:alpine

stop_db:
	docker stop api_db
