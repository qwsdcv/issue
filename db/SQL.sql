use my;

CREATE TABLE IF NOT EXISTS articles(
    id int auto_increment primary key,
    parent_id int,
    title varchar(256),
    create_date datetime,
    type varchar(32),
    content text,
    visits int
);

CREATE TABLE IF NOT EXISTS comments(
    id int auto_increment primary key,
    articleid int,
    nick_name varchar(32),
    ip int,
    content text,
    visits int,
    create_date datetime,
    foreign key(id) references articles(id)
);