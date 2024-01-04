-- migrate:up
create table wikipedia(
    wiki_id varchar(20) not null primary key,
    wiki_word varchar(50) not null,
    wiki_description text not null
);

-- migrate:down
drop table if exists wikipedia;
