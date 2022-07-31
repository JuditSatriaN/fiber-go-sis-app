CREATE TABLE IF NOT EXISTS inventory
(
    id           INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    plu          VARCHAR(30)    NOT NULL DEFAULT 0,
    unit_id      INT            NOT NULL DEFAULT 0,
    multiplier   BIGINT         NOT NULL DEFAULT 0,
    stock        BIGINT         NOT NULL DEFAULT 0,
    price        NUMERIC(10, 2) NOT NULL DEFAULT 0,
    member_price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    purchase     NUMERIC(10, 2) NOT NULL DEFAULT 0,
    discount     NUMERIC(10, 2) NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX IF NOT EXISTS inventory_plu_unit_idx ON inventory (plu, unit_id);