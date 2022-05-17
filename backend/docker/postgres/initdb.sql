CREATE extension IF NOT EXISTS "uuid-ossp";

create table users (
    id serial primary key,
    email varchar(64) not null,
    password char(60) not null,
    register_date timestamp not null
);

create table verifications (
    id uuid primary key default uuid_generate_v4(),
    email varchar(64) not null,

    /* there can be different types: registration (0), password recovery (1), password change, etc... */
    type_id smallint not null,

    /* smallint is suitable for 4-digit codes [****] */
    /* maybe there is some reason to store its hash */
    code smallint not null,

    attempts smallint not null,
    expire_time timestamp not null
);