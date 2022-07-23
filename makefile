in:
	docker exec -ti simple_db_simple_db_1 /bin/bash
up:
	docker-compose up
down:
	docker-compose down
reboot:
	docker-compose down
	docker-compose up
go_test:
	docker exec -ti simple_db_simple_db_1 bash -c "go test -v ./..."