
test:
	go test -v -cover	./...
run:
	docker-compose up -d
stop:
	docker-compose stop 
down:
	docker-compose down 
setupdb:
	python3 setup/data.py
delivery:
	python3 setup/delivery.py
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc
