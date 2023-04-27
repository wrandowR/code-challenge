BEGIN;

CREATE TABLE transactions (
    id varchar(100),
    amount DOUBLE PRECISION NOT NULL,
    date varchar(100) NOT NULL
);

COMMIT;