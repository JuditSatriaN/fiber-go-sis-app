CREATE TABLE IF NOT EXISTS store_stats
(
    store_id      VARCHAR(30) NOT NULL PRIMARY KEY,
    total_product BIGINT      NOT NULL DEFAULT 0,
    create_time   timestamp   NOT NULL DEFAULT now(),
    update_time   timestamp
);

INSERT INTO store_stats (store_id, total_product)
VALUES ('1', 0)
ON CONFLICT DO NOTHING;