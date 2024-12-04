CREATE TABLE companies (
    id UUID PRIMARY KEY,
    name VARCHAR(15) NOT NULL UNIQUE,
    description TEXT,
    num_employees INT NOT NULL,
    is_registered BOOLEAN NOT NULL,
    type VARCHAR(20) NOT NULL
);
