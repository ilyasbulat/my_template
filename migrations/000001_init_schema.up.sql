BEGIN;

--EXTENSIONS
CREATE EXTENSION IF NOT EXISTS pgcrypto;

--TABLES
CREATE TABLE public.currency (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  symbol TEXT NOT NULL
);

CREATE TABLE public.category (id SERIAL PRIMARY KEY, name TEXT NOT NULL);

CREATE TABLE public.product (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  price BIGINT,
  currency_id INT,
  rating INT,
  category_id INT,
  specification JSONB,
  image UUID,
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ
);

--DATA 
INSERT INTO
  public.currency (name, symbol)
VALUES
  ('USD', '$'),
  ('KZT', '₸');

COMMIT;