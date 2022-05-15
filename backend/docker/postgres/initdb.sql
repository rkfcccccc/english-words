create table users (
    id serial primary key,
    email varchar(64) not null,
    password char(60) not null,
    register_date timestamp
);