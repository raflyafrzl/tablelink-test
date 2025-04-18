CREATE TABLE Users(
id SERIAL PRIMARY KEY ,
xid varchar(10) NOT NULL,
email varchar(50) NOT NULL,
username varchar(12) NOT NULL,
password varchar(25) NOT NULL
)