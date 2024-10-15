CREATE TABLE Clients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    password VARCHAR(100) NOT NULL,
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Stores (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255)
);

CREATE TABLE Cassettes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    genre VARCHAR(50),
    year_of_release INT
);

CREATE TABLE CassetteAvailability (
    cassette_id INT REFERENCES Cassettes(id),
    store_id INT REFERENCES Stores(id),
    total_count INT NOT NULL,
    rented_count INT DEFAULT 0,
    PRIMARY KEY (cassette_id, store_id)
);

CREATE TABLE Reservations (
    id SERIAL PRIMARY KEY,
    client_id INT REFERENCES Clients(id),
    cassette_id INT REFERENCES Cassettes(id),
    store_id INT REFERENCES Stores(id),
    reservation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    receipt_date TIMESTAMP,
    return_date TIMESTAMP,
    status VARCHAR(20)
);

CREATE TABLE s (
    id SERIAL PRIMARY KEY,
    reservaton_id INT REFERENCES Reservations(id),
    delivery_date TIMESTAMP,
    isDeliver bool
);

-- CREATE TABLE Orders (
--     id SERIAL PRIMARY KEY,
--     client_id INT REFERENCES Clients(id),
--     cassette_id INT REFERENCES Cassettes(id),
--     store_id INT REFERENCES Stores(id),
--     order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     status VARCHAR(20)
-- );


-- CREATE TABLE OrderExecutions (
--     order_id INT REFERENCES Orders(id) PRIMARY KEY,
--     issue_date TIMESTAMP,
--     return_date TIMESTAMP
-- );