use my;

CREATE TABLE IF NOT EXISTS articles(
    id int auto_increment primary key,
    parent_id int NOT NULL,
    title varchar(256) NOT NULL,
    create_date datetime NOT NULL,
    type varchar(32) NOT NULL,
    content text NOT NULL DEFAULT '',
    visits int NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS comments(
    id int auto_increment primary key,
    articleid int NOT NULL,
    nick_name varchar(32) NOT NULL,
    ip int NOT NULL,
    content text NOT NULL,
    create_date datetime NOT NULL,
    foreign key(id) references articles(id)
);