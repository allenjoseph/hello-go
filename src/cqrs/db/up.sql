DROP TABLE IF EXISTS woofs;
CREATE TABLE woofs (
  id VARCHAR(36) PRIMARY KEY,
  body TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL
);