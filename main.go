package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// func Routes() *chi.Mux {
// 	router.Get("/posts", AllPosts)
//     router.Get("/posts/{id}", DetailPost)
//     return router
// }func Routes() *chi.Mux {
// 	router.Get("/posts", AllPosts)
//     router.Get("/posts/{id}", DetailPost)
//     return router
// }

func main() {

}
func init() {

	condb, errdb := sql.Open("mssql", "server=localhost\\Standard;user id=sa;password=prestige;database=Epcrysystem")

	var pcrID string

	if errdb != nil {
		fmt.Println("Error opening db:", errdb.Error())
	}

	rows, err := condb.Query("select * from Organization")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&pcrID)
		if err != nil {

		}
	}

}
