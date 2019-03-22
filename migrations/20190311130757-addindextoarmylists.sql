
-- +migrate Up
CREATE INDEX user_id_index_on_army_lists ON army_lists(user_id);
-- +migrate Down
DROP INDEX user_id_index_on_army_lists;