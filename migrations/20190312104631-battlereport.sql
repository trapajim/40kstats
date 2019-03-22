
-- +migrate Up
CREATE TABLE battlereports (
    id serial PRIMARY KEY NOT NULL,
    user_id VARCHAR,
    user_faction VARCHAR(50), 
    list_id int REFERENCES army_lists(id) on delete set null, 
    enemy_faction VARCHAR(50),
    enemy_list text,
    game_mode int,
    win boolean,
    player_score int,
    enemy_score int,
    created_at timestamp not null,
    updated_at timestamp not null);
-- +migrate Down
DROP TABLE battlereports;