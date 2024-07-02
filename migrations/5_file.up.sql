create table if not exists files (
    id uuid primary key default uuid_generate_v4(),
    file_url text not null,

    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp default null
);