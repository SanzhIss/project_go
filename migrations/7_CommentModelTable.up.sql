
    


CREATE TABLE comments (
    id            serial       not null unique,
    product_id int not null REFERENCES products(id),
    user_id int not null REFERENCES users (id),
    comment_body varchar(255) not null,
    rating int,
    PRIMARY KEY (id,product_id,user_id)
)
