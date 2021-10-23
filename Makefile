create_database:
	@docker-compose exec pgdb psql -U postgres -c 'CREATE DATABASE avito_users_db'
	
create-users-table:
	@docker-compose exec pgdb psql  avito_users_db -U postgres  -c 'CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY, balance NUMERIC(10,2));'

show-users-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c '\d+ users;'

check-users-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'SELECT * FROM users';

drop-users-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'DROP TABLE users';
 