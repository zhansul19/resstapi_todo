CREATE TABLE IF NOT EXISTS  users(
    id              serial       unique not null ,
    name            varchar(255) not null,
    username        varchar(255) unique not null ,
    password_hash   varchar(255) not null 
);

CREATE TABLE IF NOT EXISTS  todo_lists
(
    id          serial      not null unique,
    title       varchar(255)not null,
    description varchar(255)
);
CREATE TABLE  IF NOT EXISTS  user_lists
(
    id serial                                                 unique not null ,
    user_id int references users (id)       on delete cascade not null,
    list_id int references todo_lists (id)  on delete cascade not null
);
CREATE TABLE IF NOT EXISTS  todo_item
(
    id          serial      unique not null ,
    title       varchar(255)not null,
    description varchar(255),
    done        boolean     not null default false  
);
CREATE TABLE IF NOT EXISTS  lists_items
(
    id serial                                         unique not null ,
    items_id int references todo_item (id)       on delete cascade not null,
    list_id int references todo_lists (id)  on delete cascade not null
);