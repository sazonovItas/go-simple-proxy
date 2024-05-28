CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE e_user_role AS ENUM ('user', 'admin');

CREATE TABLE IF NOT EXISTS proxy_users (
  id            uuid          NOT NULL DEFAULT uuid_generate_v4(),
  email         varchar(320)  NOT NULL,
  login         varchar(40)   NOT NULL,
  password_hash varchar(72)   NOT NULL,
  user_role     e_user_role   NOT NULL DEFAULT 'user',
  reset_token   varchar(256)  DEFAULT NULL,
  verify_token  varchar(256)  DEFAULT NULL,
  verified      boolean       DEFAULT FALSE,
  created_at    timestamptz   NOT NULL DEFAULT NOW(),
  updated_at    timestamptz   NOT NULL DEFAULT NOW(),
  deleted_at    timestamptz   DEFAULT NULL, 
  PRIMARY KEY (id, email, login)
);

CREATE INDEX IF NOT EXISTS idx_proxy_users_id ON proxy_users (id);
CREATE INDEX IF NOT EXISTS idx_proxy_email ON proxy_users (email);
CREATE INDEX IF NOT EXISTS idx_proxy_login ON proxy_users (login);
