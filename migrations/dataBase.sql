-- +migrate Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(100) NOT NULL,
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Stores (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL
);

CREATE TABLE Cassettes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    genre VARCHAR(50),
    year_of_release INT
);

CREATE TABLE cassetteAvailability (
    cassette_id INT REFERENCES Cassettes(id) ON DELETE CASCADE,
    store_id INT REFERENCES Stores(id) ON DELETE CASCADE,
    total_count INT NOT NULL,
    rented_count INT DEFAULT 0,
    PRIMARY KEY (cassette_id, store_id)
);

CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    cassette_id INT REFERENCES Cassettes(id) ON DELETE CASCADE,
    store_id INT REFERENCES Stores(id) ON DELETE CASCADE,
    reservation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    receipt_date TIMESTAMP,
    return_date TIMESTAMP,
    status VARCHAR(20) CHECK (status IN ('active', 'returned', 'cancelled'))
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    store_id INT REFERENCES Stores(id) ON DELETE CASCADE,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    cassette_id INT REFERENCES Cassettes(id) -- Связь с кассетами
);

CREATE TABLE order_items (
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    cassette_id INT REFERENCES Cassettes(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    price_per_item DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (order_id, cassette_id)
);

-- Таблица для хранения JWT-токенов
CREATE TABLE IF NOT EXISTS jwt (
    user_id INT UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    refresh_token VARCHAR(255),
    expiresAt TIMESTAMP WITH TIME ZONE
);

CREATE TABLE reserve_pool (
    id SERIAL PRIMARY KEY,
    cassette_id INT REFERENCES Cassettes(id) NOT NULL,
    user_id INT REFERENCES users(id) NOT NULL,
    reservation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT unique_cassette_user_pair UNIQUE (cassette_id, user_id)
);
