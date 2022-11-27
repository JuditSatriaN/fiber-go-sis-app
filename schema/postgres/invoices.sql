CREATE TABLE IF NOT EXISTS invoices
(
    head_fak     VARCHAR(50) NOT NULL PRIMARY KEY,
    last_counter INT         NOT NULL,
    create_time  TIMESTAMP   NOT NULL DEFAULT NOW(),
    update_time  TIMESTAMP
);