create table if not exists sellers (
    id uuid primary key default uuid_generate_v4(),
    first_name varchar(40),
    last_name varchar(40),
    phone_number varchar(12) unique,

    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp default null
);

create unique index unique_phone_number_seller on sellers (phone_number) where deleted_at is null;