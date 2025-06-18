CREATE DATABASE rssfeeds;
CREATE USER rssuser WITH PASSWORD 'password';

ALTER ROLE rssuser SET client_encoding TO 'utf8';
ALTER ROLE rssuser SET default_transaction_isolation TO 'read committed';
ALTER ROLE rssuser SET timezone TO 'UTC';


GRANT ALL PRIVILEGES ON DATABASE rssfeeds TO rssuser;

CREATE TABLE feeds(
    title text NOT NULL ,
    description text,
    pubdate varchar,
    link varchar PRIMARY KEY
);



