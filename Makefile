
test:
	go test -v -cover	./...
run:
	docker-compose up 
setupdb:
	python3 setup/data.py
delivery:
	python3 setup/delivery.py
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc
