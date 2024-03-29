CREATE TABLE IF NOT EXISTS sales_head
(
    id                BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    invoice           VARCHAR(15)    NOT NULL,
    user_id           VARCHAR(30)    NOT NULL,
    total_item        INT            NOT NULL DEFAULT 0,
    total_price       NUMERIC(10, 2) NOT NULL DEFAULT 0,
    total_purchase    NUMERIC(10, 2) NOT NULL DEFAULT 0,
    total_tax         NUMERIC(10, 2) NOT NULL DEFAULT 0,
    total_discount    NUMERIC(10, 2) NOT NULL DEFAULT 0,
    total_pay         NUMERIC(10, 2) NOT NULL DEFAULT 0,
    create_time       TIMESTAMP      NOT NULL DEFAULT NOW(),
    update_time       TIMESTAMP,
    value_text_search TSVECTOR
);

CREATE INDEX IF NOT EXISTS create_time_idx ON sales_head ((create_time::date));
CREATE UNIQUE INDEX IF NOT EXISTS invoice_user_id_sales_head_idx ON sales_head (invoice, user_id);
CREATE INDEX IF NOT EXISTS value_text_search_idx ON sales_head USING GIN (value_text_search);

CREATE OR REPLACE FUNCTION sales_head_upsert_search_trigger() RETURNS trigger AS
$$
BEGIN
    new.value_text_search := TO_TSVECTOR(new.invoice || ' ' || new.user_id );
    RETURN new;
END
$$ LANGUAGE plpgsql;

BEGIN;
DROP TRIGGER IF EXISTS sales_head_upsert_search ON sales_head;
CREATE TRIGGER sales_head_upsert_search
    BEFORE INSERT OR UPDATE
    ON sales_head
    FOR EACH ROW
EXECUTE PROCEDURE sales_head_upsert_search_trigger();
COMMIT;

