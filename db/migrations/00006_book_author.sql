-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE book_author (
  id int NOT NULL AUTO_INCREMENT,
  book_id int NOT NULL,
  author_id int NOT NULL,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE book_author;
