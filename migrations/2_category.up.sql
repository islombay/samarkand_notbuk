create table if not exists categories (
    id uuid primary key default uuid_generate_v4(),
    name_uz varchar(30) not null,
    parent_id uuid,

    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp default null,
    CONSTRAINT fk_parent FOREIGN KEY (parent_id) REFERENCES categories(id)
);

create unique index unique_name_uz_category on categories (name_uz) where deleted_at is null;