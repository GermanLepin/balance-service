# Первый запуск проекта
start-project: docker.first-start create.all-tables 

# Запуск интеграционных тестов
docker.test.integration: start-project test-integration docker.stop

test-integration:
	go test -tags=integration ./integration_tests -v 

# Запуск UNIT тестов 
test.unit:
	go test ./...

# Первый запуск докера
docker.first-start:
	docker compose build
	docker-compose up -d
	docker-compose up -d

# Запуск проекта
docker.start:
	docker-compose up -d
	docker-compose up -d

# Остановка докера
docker.stop:
	docker-compose down

# Перезапуск проекта
docker.restart: docker.stop docker.start

# Создание базы данных и таблиц
create.all-tables: create-database create-users-table create-description-table

create-database:
	@docker-compose exec pgdb psql -U postgres -c  "SELECT 1 FROM pg_database WHERE datname = 'avito_users_db'"
	
create-users-table:
	@docker-compose exec pgdb psql  avito_users_db -U postgres  -c 'CREATE TABLE IF NOT EXISTS users (id INTEGER NOT NULL PRIMARY KEY, balance NUMERIC(10,2));'

create-description-table:
	@docker-compose exec pgdb psql  avito_users_db -U postgres  -c 'CREATE TABLE IF NOT EXISTS description (id_description SERIAL PRIMARY KEY, sender_receiver VARCHAR(100), amount NUMERIC(10,2), description VARCHAR(255), balance_at_moment NUMERIC(10,2), user_id INTEGER, FOREIGN KEY(user_id) REFERENCES users (id), created_at TIMESTAMP NOT NULL, refill VARCHAR(100) NOT NULL);'

# Просмотр данных в таблицах
check-users-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'SELECT * FROM users';

check-description-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'SELECT * FROM description';

check-users-description-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'SELECT * FROM users us INNER JOIN description de ON us.id=de.userId';

# Удаление базы данных и таблиц
drop-users-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'DROP TABLE users';

drop-description-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'DROP TABLE description';

drop-database:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'DROP DATABASE avito_users_db';
