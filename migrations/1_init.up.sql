create table if not exists users (
    id       INTEGER PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    passHash blob not null
);

create index if not exists idx_email ON users (email);


create table if not exists apps (
    id     INTEGER PRIMARY KEY,
    name   text not null unique,
    secret text not null unique
);