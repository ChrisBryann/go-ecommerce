CREATE TYPE order_status AS ENUM('pending', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    userId BIGINT NOT NULL REFERENCES users(id),
    status order_status NOT NULL DEFAULT 'pending',
    address TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);