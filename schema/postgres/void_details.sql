CREATE TABLE IF NOT EXISTS void_detail
(
    id           BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    invoice      VARCHAR(15)    NOT NULL,
    user_id      VARCHAR(30)    NOT NULL,
    plu          VARCHAR(30)    NOT NULL,
    name         VARCHAR(255)   NOT NULL,
    unit_name    VARCHAR(30)    NOT NULL,
    barcode      VARCHAR(30)    NOT NULL,
    ppn          BOOLEAN        NOT NULL DEFAULT FALSE,
    qty          BIGINT         NOT NULL DEFAULT 0,
    price        NUMERIC(10, 2) NOT NULL DEFAULT 0,
    purchase     NUMERIC(10, 2) NOT NULL DEFAULT 0,
    discount     NUMERIC(10, 2) NOT NULL DEFAULT 0,
    member_id    BIGINT         NOT NULL DEFAULT 0,
    inventory_id BIGINT         NOT NULL DEFAULT 0,
    create_time  TIMESTAMP      NOT NULL DEFAULT NOW(),
    update_time  TIMESTAMP
);

CREATE INDEX IF NOT EXISTS void_detail_plu_idx ON void_detail (plu);
CREATE INDEX IF NOT EXISTS void_detail_barcode_idx ON void_detail (barcode);
CREATE INDEX IF NOT EXISTS invoice_user_id_void_det_idx ON void_detail (invoice, user_id);
CREATE INDEX IF NOT EXISTS void_detail_create_time_idx ON void_detail ((create_time::DATE));