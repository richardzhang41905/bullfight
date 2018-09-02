# bullfight
##这是一个自己写的扑克牌斗牛游戏的演示，能完成自动洗牌，发牌，比大小。
##这个项目主要完成核心逻辑的演示。
##后续可以考虑把这个游戏做到区块链上。


create database bullfight CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
use wallet_web_production
create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null   
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);

create table threads (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  topic      text,
  user_id    integer references users(id),
  created_at timestamp not null       
);

create table posts (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  thread_id  integer references threads(id),
  created_at timestamp not null  
);