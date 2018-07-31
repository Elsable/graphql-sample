-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE book (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255),
  created_at datetime,
  updated_at datetime,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE book;
