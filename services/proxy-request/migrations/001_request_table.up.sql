CREATE TABLE IF NOT EXISTS proxy_requests (
  id              UUID,
  proxy_id        UUID,
  proxy_name      varchar(255),
  proxy_user_id   UUID,
  proxy_user_ip   varchar(39),
  proxy_user_name varchar(255),
  host            varchar(255),
  upload          bigint,
  download        bigint,
  created_at      timestamptz DEFAULT NOW(),
  PRIMARY KEY (id, created_at)
) PARTITION BY RANGE (created_at);

CREATE INDEX IF NOT EXISTS proxy_requests_id_idx ON proxy_requests USING btree (id);
CREATE INDEX IF NOT EXISTS proxy_requests_created_at_idx ON proxy_requests (created_at DESC NULLS LAST);
