create table if not exists products (
    id uuid primary key default uuid_generate_v4(),

    name varchar(255) not null,
    description text,
    price numeric not null,

    category_id uuid default null,
    brand_id uuid default null,
    image_id uuid default null,
    video_id uuid default null,

    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default null,

    foreign key (category_id) references categories(id) on delete set null,
    foreign key (brand_id) references brands(id) on delete set null,
    foreign key (image_id) references files(id) on delete set null,
    foreign key (video_id) references files(id) on delete set null
);

create table if not exists product_files (
    product_id uuid not null,
    file_id uuid not null,

    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp default null,
    foreign key (product_id) references products(id) on delete cascade,
    foreign key (file_id) references files(id) on delete cascade,

    unique (product_id, file_id)
);

create table if not exists product_installments (
    product_id uuid not null,
    price numeric not null,
    period int not null,

    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp default null,

    foreign key (product_id) references products(id) on delete cascade,
    unique (product_id, period)
);