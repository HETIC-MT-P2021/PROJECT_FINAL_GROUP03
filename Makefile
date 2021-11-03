init:
	make copy-files && make start

start:
	docker-compose up --build -d

stop:
	docker-compose stop

log:
	docker-compose logs -f

copy-files:
	cp env/domainApi/.env.dist domainApi/.env
	cp env/front/.env.dist front/.env
	cp env/frontApi/.env.dist frontApi/.env
	cp env/global/.env .env
	cat env/global/.env >> domainApi/.env
	cat env/global/.env >> frontApi/.env
	cat env/global/.env >> front/.env
