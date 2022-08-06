CREATE TABLE IF NOT EXISTS products
(
    plu               VARCHAR(30)  NOT NULL PRIMARY KEY,
    name              VARCHAR(255) NOT NULL,
    barcode           VARCHAR(30)  NOT NULL DEFAULT '0',
    ppn               BOOLEAN      NOT NULL DEFAULT FALSE,
    create_time       TIMESTAMP    NOT NULL DEFAULT NOW(),
    update_time       TIMESTAMP,
    value_text_search TSVECTOR
);


CREATE INDEX IF NOT EXISTS barcode_products_idx ON products (barcode);
CREATE INDEX IF NOT EXISTS value_text_search_idx ON products USING GIN (value_text_search);

CREATE OR REPLACE FUNCTION products_upsert_search_trigger() RETURNS trigger AS
$$
BEGIN
    new.value_text_search := TO_TSVECTOR(new.plu || ' ' || new.name || ' ' || new.barcode);
    RETURN new;
END
$$ LANGUAGE plpgsql;

BEGIN;
DROP TRIGGER IF EXISTS product_upsert_search ON products;
CREATE TRIGGER product_upsert_search
    BEFORE INSERT OR UPDATE
    ON products
    FOR EACH ROW
EXECUTE PROCEDURE products_upsert_search_trigger();
COMMIT;