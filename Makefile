create_database:
	@docker-compose exec pgdb psql -U postgres -c 'CREATE DATABASE avito_users_db'
	
create-users-table:
	@docker-compose exec pgdb psql  avito_users_db -U postgres  -c 'CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY, balance NUMERIC(10,2));'

create-description-table:
	@docker-compose exec pgdb psql  avito_users_db -U postgres  -c 'CREATE TABLE description (id_description SERIAL PRIMARY KEY, sender_receiver VARCHAR(100), amount NUMERIC(10,2), description VARCHAR(255), balance_at_moment NUMERIC(10,2), user_id INTEGER, FOREIGN KEY(user_id) REFERENCES users (id), created_at TIMESTAMP NOT NULL, refill VARCHAR(100) NOT NULL);'

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
