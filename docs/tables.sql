CREATE TABLE IF NOT EXISTS template (
  id VARCHAR(20) PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  org_id VARCHAR(255) NoT NULL,
  version INT,
  is_simple BOOLEAN,
  request TEXT,
  response TEXT,
  retries INT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
) ENGINE=MYISAM CHARACTER SET utf8mb4;

CREATE TABLE IF NOT EXISTS job (
  id VARCHAR(20) PRIMARY KEY NOT NULL,
  template_id VARCHAR(20) NOT NULL,
  template_version INT,
  url_variables TEXT,
  body_variables TEXT,
  extra_query_params TEXT,
  extra_headers TEXT
) ENGINE=MYISAM CHARACTER SET utf8mb4; 
