-- +goose Up
create table orders(
   id serial PRIMARY KEY,
   updated_at timestamptz NOT NULL,
   product_id int not null references products(id) on delete cascade,
   user_id int not null references users(id) on delete cascade,
   unique(user_id,product_id)
);

-- +goose Down
drop table orders;