
-- +migrate Up
CREATE TABLE army_lists (
    id serial PRIMARY KEY NOT NULL,
    list_name VARCHAR(50), 
    faction VARCHAR(50),
    list text,
    user_id VARCHAR,
    created_at timestamp not null,
    updated_at timestamp not null);
-- +migrate Down
DROP TABLE army_lists;
