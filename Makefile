start:
	docker-compose up --build -d

stop:
	docker-compose stop

log:
	docker-compose logs -f 
