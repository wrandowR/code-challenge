BEGIN;

CREATE TABLE transactions (
    amount int NOT NULL,
    date varchar(100) NOT NULL,
);

COMMIT;