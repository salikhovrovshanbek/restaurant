drop table basket;
create table basket (hek_id UUID not null references chek(id),
    product_id UUID not null,
    tip int not null
);