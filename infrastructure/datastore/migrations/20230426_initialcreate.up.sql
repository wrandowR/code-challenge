BEGIN;

CREATE TABLE transactions (
    amount DOUBLE PRECISION NOT NULL,
    date varchar(100) NOT NULL
);

COMMIT;