create type user_role as enum ('moderator', 'user');

create table user (
  id uuid primary key, 
  name varchar(255) not null,
  email varchar(255) unique not null,
  password varchar(255) not null,
  role user_role not null
);
