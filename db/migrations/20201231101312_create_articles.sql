-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table articles (
  id int auto_increment,
  title varchar(100),
  primary key(id)
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table articles;
