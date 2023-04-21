

CREATE TABLE products ( 
     id            serial       not null unique,
    product_id int not null unique REFERENCES order_details(id),
    card_id int not null unique,
    product_title varchar(255) not null,
    category_id int not null unique REFERENCES category(id),
    price int not null,
    quantity int not null,
    product_description varchar(255) not null,
    image1 bytea not null,

    PRIMARY KEY (id,product_id,category_id))
