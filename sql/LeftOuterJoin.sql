select customers.id,name ,phone,items,total from customers LEFT OUTER JOIN orders ON customers.id=orders.id;
