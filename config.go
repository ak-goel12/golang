package main

import(
	"database/sql"
	"fmt"


	_"github.com/go-sql-driver/mysql" //import the pacakage
)

const (  
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "go_sql"
)

//return connection string
func dsn(dbName string) string {  
    return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main(){
	fmt.Println("Go connecting to mysql")

	
	db,err := sql.Open("mysql",dsn(dbname)) //connect to the mysql 
 
	if err != nil{ //if any error then print it
		panic(err.Error())	
		return
	}
	fmt.Println("Connected")
	defer db.Close() //close the connection
	
// insert query
	insert,err := db.Query("INSERT INTO user_info (name,mobile,address) VALUES ('AKSHIT','9818715981','435A/2')") 
	

	if err!=nil{
		panic(err.Error())
	}

	fmt.Println("Successfully inserted the data")

	defer insert.Close()

// update query
	update,err := db.Query("Update user_info SET name='Ashish', mobile ='1234567890' where id = 2")

	if err!=nil{
		panic(err.Error())
	}

	defer update.Close()

//delete query
	delete,err := db.Query("DELETE FROM user_info WHERE id = 3")

	if err!=nil{
		panic(err.Error())
	}

	defer delete.Close()

//read query
	results,err := db.Query("Select * from user_info")

	if err!=nil{
		panic(err.Error())
	}

	for results.Next() {
		var id int
		var name string
		var address string
		var mobile string

		err = results.Scan(&id,&name,&address,&mobile)
		if err!=nil{
			panic(err.Error())
		}
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(address)
		fmt.Println(mobile)	
	}
	defer results.Close()


}
