
CREATE TABLE IF NOT EXISTS articles(
    id int auto_increment primary key,
    title varchar(256),
    content text,
    create_date datetime,
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