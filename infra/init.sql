CREATE TABLE payments (
  correlationId UUID PRIMARY KEY,
  amount DECIMAL NOT NULL,
  requested_at TIMESTAMP NOT NULL UNIQUE
);