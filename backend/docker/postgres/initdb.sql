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

create table movies (
    id serial primary key,
    imdb_id varchar(16) unique,
    title varchar(128) not null,
    year smallint not null,
    poster_url varchar(128) not null
);

create table movies_words (
    movie_id int references movies (id) on delete cascade,
    word_id char(24) not null, /* mongodb dictionary._id */
    primary key (movie_id, word_id)
);

CREATE INDEX idx_movies_words_movie_id ON movies_words(movie_id);

create table movies_users (
    movie_id int references movies (id) on delete cascade,
    user_id int references users (id) on delete cascade,
    primary key (movie_id, user_id)
);

create table vocabulary (
    user_id int references users (id) on delete cascade not null,
    word_id char(24) not null, /* mongodb dictionary._id */

    already_learned boolean not null default false,
    learning_step int not null default 0,

    next_challenge int,

    count int not null,
    primary key (user_id, word_id)
);

CREATE INDEX idx_vocabulary_user_id ON vocabulary(user_id);
