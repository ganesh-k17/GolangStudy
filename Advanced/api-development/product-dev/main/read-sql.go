package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	id   int
	name string
}

func main() {
	DbUser := "root"
	DbPassword := "ganesh"
	DBName := "learning"

	// connectionString := fmt.Sprintf("%v,%v@tcp(127.0.0.1:3306)/%v", DbUser, DbPassword, DBName)
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DbUser, DbPassword, DBName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	/*** Insert Data ***/
	result, err := db.Exec("insert into data values(5, 'xyz')")

	if err != nil {
		fmt.Println("Error inserting data", err.Error())
		return
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		fmt.Println("Error fetching last insert id", err.Error())
		return
	}

	fmt.Println("Last inserted Id ", lastInsertedId)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error fetching rows affected:", err.Error())
		return
	}
	fmt.Println("Rows Affected ", rowsAffected)

	/*** Select Data ***/
	rows, err := db.Query("select * from data")

	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		if err != nil {
			fmt.Println("Error")
			return
		}
		fmt.Println(data)
	}

	fmt.Println(connectionString)
}
