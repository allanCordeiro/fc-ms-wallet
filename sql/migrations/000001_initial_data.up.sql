USE wallet;


CREATE TABLE IF NOT EXISTS clients
(id varchar(255), name varchar(255), email varchar(255), created_at timestamp);

CREATE TABLE IF NOT EXISTS accounts
(id varchar(255), client_id varchar(255), balance float, created_at timestamp);

CREATE TABLE IF NOT EXISTS transactions
(id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at timestamp);