INSERT INTO proxy_users (email, login, password_hash, user_role) 
  VALUES ('admin@example.com', 'admin', '$2y$10$UMSAWYRGmxdnQJyWBnakguyUtlv2OrWMedNm5uWsL6mgr289zFWmq', 'admin') 
  ON CONFLICT DO NOTHING;
