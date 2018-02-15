-- +migrate Up
CREATE TABLE IF NOT EXISTS people (id int);
INSERT INTO people VALUES(4);
INSERT INTO people VALUES(5);


-- +migrate Down
DROP TABLE people;
