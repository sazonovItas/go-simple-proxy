INSERT INTO proxy_users (email, login, password_hash, user_role, verified) 
  VALUES ('admin@mail.com', 'admin', '$2y$10$RwSvWCoHH2.6pYs8wMIKSugU2KhAOZEGWPgwtiWee7WYDkIHv8FKm', 'admin', true) 
  ON CONFLICT DO NOTHING;
