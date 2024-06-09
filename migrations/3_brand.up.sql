create table if not exists brands (
    id uuid primary key default uuid_generate_v4(),
    name varchar(40),

    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp default null
);

create unique index unique_name_brand on brands (name) where deleted_at is null;