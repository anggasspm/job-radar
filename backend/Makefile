server : 
	APP_ENV=dev air

DB_CONTAINER=backend-postgres-1
DB_USER=root
DB_NAME=job-radar

schema:
	docker exec $(DB_CONTAINER) \
		pg_dump \
		-U $(DB_USER) \
		-d $(DB_NAME) \
		-s \
		--no-owner \
		--no-privileges > database/schema.sql