-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE user_book (
  id int NOT NULL AUTO_INCREMENT,
  book_id int NOT NULL,
  user_id int NOT NULL,
  created_at datetime,
  updated_at datetime,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE user_book;
