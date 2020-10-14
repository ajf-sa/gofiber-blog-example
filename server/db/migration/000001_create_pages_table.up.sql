CREATE TABLE IF NOT EXISTS pages(
    id serial primary key,
    parent_id int not null default 0,
    title VARCHAR not null default '',
    slug VARCHAR not null unique,
    body text not null default '',
    create_at TIMESTAMPTZ not null default(now()),
    foreign key (parent_id) references pages(id)

);