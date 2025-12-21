


create table users (
    id varchar(60) primary key,
    first_name varchar(120) not null,
    last_name varchar(120) not null,
    address varchar(255) not null,
    email varchar(255) not null unique,
    username varchar(255) not null,
    password varchar(255) not null,
    age int not null
);

create table posts (
    id serial primary key,
    user_id  varchar(255) not null references users(id) on delete cascade,
    content text not null,
    image_url text,
    created_at timestamp set default now(),
    updated_at timestamp set default now()
);

create table comments (
    id serial primary key,
    user_id varchar(255) not null references users(id) on delete cascade,
    post_id int not null  references posts(id) on delete cascade,
    content text not null,
    image_url text
    created_at timestamp set default now(),
    updated_at timestamp set default now()
);

create table likes (
    id serial primary key,
    user_id varchar(255) not null references users(id) on delete cascade,
    post_id int references posts(id) on delete cascade,
    comment_id int references comments(id) on delete cascade
);

