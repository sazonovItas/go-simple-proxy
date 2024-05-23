CREATE TABLE IF NOT EXISTS proxy_requests (
  id              uuid          NOT NULL,
  proxy_id        uuid          NOT NULL,
  proxy_name      varchar(255)  NOT NULL,
  proxy_user_id   uuid          NOT NULL,
  proxy_user_ip   varchar(39)   NOT NULL,
  proxy_user_name varchar(255)  NOT NULL,
  host            varchar(255)  NOT NULL,
  upload          bigint        NOT NULL,
  download        bigint        NOT NULL,
  created_at      timestamptz   NOT NULL  DEFAULT NOW(),
  PRIMARY KEY (id, created_at)
);

CREATE INDEX IF NOT EXISTS proxy_requests_id_idx ON proxy_requests USING btree (id);
CREATE INDEX IF NOT EXISTS proxy_requests_created_at_idx ON proxy_requests (created_at DESC NULLS LAST);
