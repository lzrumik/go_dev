

GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '123456' WITH GRANT OPTION;
flush privileges;

create database gomysql;
use gomysql;

create table person(
user_id int(11) primary key auto_increment,
username varchar(260),
sex varchar(10),
email varchar(260)
);


create table place(
country varchar (200),
city varchar (200),
telcode int
);