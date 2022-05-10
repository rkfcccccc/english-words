create table users (
    id serial primary key,
    email varchar(64) not null,
    password varchar(32) not null,
    register_date timestamp
);