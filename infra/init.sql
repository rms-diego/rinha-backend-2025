CREATE TABLE payments (
  "correlationId" UUID PRIMARY KEY,
  "amount" DECIMAL NOT NULL,
  "requestedAt" TIMESTAMP NOT NULL UNIQUE
);
