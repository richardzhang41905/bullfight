drop table posts;
drop table threads;
drop table sessions;
drop table users;


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

create table games(
	id          serial primary key,
	uuid1       varchar(64) not null,
	uuid2       varchar(64),
	user_id1    integer references users(id),
	user_id2    integer references users(id),
	created_at  datetime not null default CURRENT_TIMESTAMP (),
	join_at     datetime null,
	closed_at   datetime null,
	left_cards  text,
	user1_cards text,
	user2_cards text,
	status integer,
	result integer
);