CREATE TYPE statusTr AS ENUM ('debet', 'credit');

CREATE TABLE public.users (
id varchar unique NOT null primary key,
username varchar not NULL,
fullname varchar not NULL,
"password" varchar not NULL,
email varchar not NULL,
birthdate date not NULL,
domicile varchar not NULL,
occupation varchar not NULL,
created_at timestamp default now(),
updated_at timestamp,
updated_by integer
);

CREATE TABLE public.account (
id varchar unique null primary key,
balance integer NULL,
created_at timestamp default now(),
updated_at timestamp NULL,
created_by integer null,
user_id varchar,
constraint fk_user_id FOREIGN key(user_id) REFERENCES public.users(id)
);

CREATE TABLE public.transaction_log (
id varchar unique null primary key,
account_id varchar not NULL,
"date" timestamp not NULL,
"type" statusTr not NULL,
description varchar NULL,
ammount integer NULL,
balance integer NULL,
created_at timestamp default now(),
updated_by timestamp NULL,
created_by integer null,
updated_at integer null,
constraint id foreign key(account_id) references public.account(id)
);