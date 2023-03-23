-- +migrate Up

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION trigger_update_timestamp ()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = NOW();
RETURN NEW;
END;$$
LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate Down

DROP FUNCTION IF EXISTS trigger_update_timestamp;
