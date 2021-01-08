start:
	reflex -d none -r '(\.go$|go\.mod)' -s -- go run ./cmd/genius/

stop:
	docker-compose down --remove-orphans
	docker-compose rm -s -f

debug: ##@application Starts the API app in debug mode, available at 0.0.0.0:2345
	dlv debug ./cmd/genius -l 0.0.0.0:2345 --headless=true --log=true --api-version=2 ./cmd/genius

logs:
	docker-compose logs -f --tail 100 genius
