
-- +migrate Up
    ALTER TABLE army_lists ADD COLUMN  pl INT, ADD COLUMN pts INT, ADD COLUMN cp INT;
    
-- +migrate Down
    ALTER TABLE army_lists DROP COLUMN pl, DROP COLUMN pts, DROP COLUMN cp;