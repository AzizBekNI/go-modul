<<<<<<< HEAD
package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func OpenDb(name string){
	db, err := sql.Open("mysql","root:Aa12345$@tcp(127.0.0.1:3306)/" + name)
	if err != nil{
		panic(err)
	}
	defer db.Close()
=======
package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func OpenDb(name string){
	db, err := sql.Open("mysql","root:Aa12345$@tcp(127.0.0.1:3306)/" + name)
	if err != nil{
		panic(err)
	}
	defer db.Close()
>>>>>>> c799814a28fe0b6c5b0331ef28c89a611725ddac
}