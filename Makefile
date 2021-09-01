init:
	make copy-files && make start

start:
	docker-compose up --build -d

stop:
	docker-compose stop

log:
	docker-compose logs -f

copy-files:
	cp env/back/.env.dist back/.env
	cp env/front/.env.dist front/.env
	cp env/global/.env .env
	cat env/global/.env >> back/.env
	cat env/global/.env >> front/.env
