-- +migrate Up
CREATE TABLE IF NOT EXISTS people (id int);
INSERT INTO people VALUES(1);
INSERT INTO people VALUES(2);
INSERT INTO people VALUES(3);


-- +migrate Down
DROP TABLE people;
