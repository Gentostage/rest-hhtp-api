CREATE TABLE users (
    id bigserial not null PRIMARY KEY,
    email VARCHAR not NULL UNIQUE,
    encrypted_password VARCHAR not NULL
);