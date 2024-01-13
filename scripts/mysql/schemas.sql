CREATE TABLE
    IF NOT EXISTS awesomedatabase.users (
        id VARCHAR(36) PRIMARY KEY NOT NULL,
        name VARCHAR(60) NOT NULL,
        birth_date DATE NOT NULL,
        email VARCHAR(60) NOT NULL,
        password VARCHAR(60) NOT NULL,
        address VARCHAR(100) NOT NULL
    );