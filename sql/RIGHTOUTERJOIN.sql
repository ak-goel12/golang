select customers.id,name ,phone,items,total from customers RIGHT OUTER JOIN orders ON customers.id=orders.id;
