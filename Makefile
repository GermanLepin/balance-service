# Запуск проекта
start-project: docker.build docker.start

#Запуск интеграционных тестов
test.integration:
	go test -tags=integration ./integration_tests -v 

# Запуск UNIT тестов 
test.unit:
	go test ./...

# Билд проекта 
docker.build:
	docker compose build

# Старт докер
docker.start:
	docker-compose up -d 
	docker-compose up -d

# Остановка докера
docker.stop:
	docker-compose down

create-database:
	@docker-compose exec pgdb psql -U postgres -c 'CREATE DATABASE avito_db;'
	
create-users-table:
	@docker-compose exec pgdb psql  avito_db -U postgres  -c 'CREATE TABLE IF NOT EXISTS users (id INTEGER NOT NULL PRIMARY KEY, balance NUMERIC(10,2));'

create-description-table:
	@docker-compose exec pgdb psql  avito_db -U postgres  -c 'CREATE TABLE IF NOT EXISTS descriptions (id_description SERIAL PRIMARY KEY, sender_receiver VARCHAR(100), amount NUMERIC(10,2), description VARCHAR(255), balance_at_moment NUMERIC(10,2), user_id INTEGER, FOREIGN KEY(user_id) REFERENCES users (id), created_at TIMESTAMP NOT NULL, refill VARCHAR(100) NOT NULL);'

# Просмотр данных в таблицах
check-users-table:
	@docker-compose exec pgdb psql -U postgres -d avito_db -c 'SELECT * FROM users';

check-descriptions-table:
	@docker-compose exec pgdb psql -U postgres -d avito_db -c 'SELECT * FROM descriptions';

# Очистить содержимое таблиц
truncate-descriptions-table:
	@docker-compose exec pgdb psql -U postgres -d avito_db -c 'truncate descriptions';

truncate-users-table:
	@docker-compose exec pgdb psql -U postgres -d avito_db -c 'truncate users cascade';

# Удаление базы данных и таблиц
drop-users-table:
	@docker-compose exec pgdb psql -U postgres -d avito_db -c 'DROP TABLE users';

drop-descriptions-table:
	@docker-compose exec pgdb psql -U postgres -d avito_db -c 'DROP TABLE descriptions';

drop-database:
	@docker-compose exec pgdb psql -U postgres -d avito_db -c 'DROP DATABASE avito_db';
