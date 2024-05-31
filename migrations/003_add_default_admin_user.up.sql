INSERT INTO proxy_users (email, login, password_hash, user_role, verified) 
  VALUES ('admin@example.com', 'admin', '$2y$10$UMSAWYRGmxdnQJyWBnakguyUtlv2OrWMedNm5uWsL6mgr289zFWmq', 'admin', true) 
  ON CONFLICT DO NOTHING;
