


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

create index idx_users_username on users(username)
create index idx_posts_user_id on posts(user_id)
create index idx_posts_created_at on posts(created_at desc)
create index idx_comments_post_id on comments(post_id)
create index idx_comments_user_id on comments(user_id)
create index idx_comments_created_at on comments(created_at desc)
create index idx_likes_user_id on likes(user_id)
create index idx_likes_post_id on likes(post_id)
create index idx_likes_comment_id on likes(comment_id)
