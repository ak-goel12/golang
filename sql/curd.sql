create database firstapp;
use firstapp;
create table test(
Sno int,
Name varchar(50),
Age int,
Gender varchar(10),
Phone varchar(10));
insert into demo values(1,"raju",20,'Male',789456123);
insert into demo values(2,"ramesh",28,'Female',889456123);
insert into demo values(3,"rahul",20,'Female',897456123);
insert into demo values(4,"aman",15,'Male',7789456123);
insert into demo values(5,"ajay",86,'Female',7979456123);

select* from test
set gender='Male';
update demo
set age='28'
where Name='ramesh';

select *from test;
delete from test where Name='ajay';

select* from test;

drop database firstapp;

