create_database:
	@docker-compose exec pgdb psql -U postgres -c 'CREATE DATABASE avito_users_db'
	
create-users-table:
	@docker-compose exec pgdb psql  avito_users_db -U postgres  -c 'CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY, balance NUMERIC(10,2));'

create-description-table:
	@docker-compose exec pgdb psql  avito_users_db -U postgres  -c 'CREATE TABLE description (id_description SERIAL PRIMARY KEY, created_at TIMESTAMP NOT NULL, description VARCHAR(255), sender_receiver VARCHAR(100), balance_at_moment NUMERIC(10,2), amount NUMERIC(10,2), refill VARCHAR(100) NOT NULL, userId INTEGER, FOREIGN KEY(userId) REFERENCES users (id));'

show-users-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c '\d+ users;'

show-description-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c '\d+ description;'

check-users-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'SELECT * FROM users';

check-description-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'SELECT * FROM description';

drop-users-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'DROP TABLE users';

drop-description-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'DROP TABLE description';
 
check-users-description-table:
	@docker-compose exec pgdb psql -U postgres -d avito_users_db -c 'SELECT * FROM users us INNER JOIN description de ON us.id=de.userId';
