CREATE TABLE IF NOT EXISTS proxy_requests (
  id              uuid          NOT NULL,
  user_id         uuid          NOT NULL,
  proxy_id        uuid          NOT NULL,
  remote_ip       varchar(39)   NOT NULL,
  host            varchar(255)  NOT NULL,
  upload          bigint        NOT NULL,
  download        bigint        NOT NULL,
  created_at      timestamptz   NOT NULL  DEFAULT NOW(),
  PRIMARY KEY (id, created_at)
);

CREATE INDEX IF NOT EXISTS idx_proxy_requests_id ON proxy_requests (id);
CREATE INDEX IF NOT EXISTS idx_proxy_requests_created_at ON proxy_requests (created_at DESC NULLS LAST);
