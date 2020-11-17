use shop;
select* from customers;
select * from orders;
select customers.id,name ,phone,items,total from customers INNER JOIN orders ON customers.id=orders.id;

