create extension if not exists "uuid-ossp";

create table if not exists clients (
    id uuid primary key default uuid_generate_v4(),
    first_name varchar(40),
    last_name varchar(40),
    password varchar (255),
    phone_number varchar(12) unique,

    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp default null
);

CREATE UNIQUE INDEX unique_phone_number_client_active ON clients (phone_number) WHERE deleted_at IS NULL;

create table if not exists staffs (
    id uuid primary key default uuid_generate_v4(),
    first_name varchar(40),
    last_name varchar(40),
    phone_number varchar(12) unique,
    password varchar(255),

    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp default null
);

CREATE UNIQUE INDEX unique_phone_number_staff_active ON staffs (phone_number) WHERE deleted_at IS NULL;