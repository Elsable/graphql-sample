-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE user (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255),
  password varchar(255),
  created_at datetime,
  updated_at datetime,
  PRIMARY KEY (id),
  UNIQUE KEY `name_UNIQUE` (`name`)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE user;
