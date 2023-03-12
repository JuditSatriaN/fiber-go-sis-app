CREATE TABLE IF NOT EXISTS system_conf (
    id VARCHAR(30) PRIMARY KEY,
    value TEXT
);

INSERT INTO system_conf (id, value) VALUES ('password_void', '123') ON CONFLICT DO NOTHING;