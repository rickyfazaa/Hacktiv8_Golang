CREATE TABLE orders (
    order_id int primary key,
    customer_name varchar(100),
    ordered_at timestamp
)

CREATE TABLE items (
    item_id int primary key,
    item_code varchar(11) not null,
    description varchar(100),
    quantity int,
    order_id int
)