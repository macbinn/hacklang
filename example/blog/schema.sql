create table blogs
(
	id integer not null primary key autoincrement,
	title text default '' not null,
	content text default '' not null,
	author_id integer default 0 not null,
	created_at timestamp default current_timestamp not null
);

create index blogs_author_index on blogs (author_id);

create table users
(
	id integer not null primary key autoincrement,
	name varchar(64) default '' not null,
	password_hash varchar(64) default '' not null,
	created_at timestamp default current_timestamp not null
);

create unique index users_name_uindex on users (name);
