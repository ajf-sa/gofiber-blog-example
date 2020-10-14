CREATE TABLE IF NOT EXISTS pages(
    id serial primary key,
    title VARCHAR not null default '',
    slug VARCHAR UNIQUE not null,
    body text not null default '',
    is_active boolean not null DEFAULT false,
    create_at TIMESTAMPTZ not null default(now())


);