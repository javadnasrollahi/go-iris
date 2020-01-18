package mysql

import (
	"database/sql"
	// mysql golang client
	_ "github.com/go-sql-driver/mysql"
)

//NewClient : create a new client for mysql data base
func NewClient() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:pass@/dbname")
	return db, err
}

/*
EXAMPLES:

defer db.Close()

    * perform a db.Query insert
    insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")


 * Tag... - a very simple struct

 type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}


    * Execute the query
    results, err := db.Query("SELECT id, name FROM tags")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for results.Next() {
        var tag Tag
        * for each row, scan the result into our tag composite object
        err = results.Scan(&tag.ID, &tag.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                * and then print out the tag's Name attribute
        log.Printf(tag.Name)
    }

	var tag Tag

	@ Querying a Single ROW
* Execute the query
err = db.QueryRow("SELECT id, name FROM tags where id = ?", 2).Scan(&tag.ID, &tag.Name)
if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
}

log.Println(tag.ID)
log.Println(tag.Name)

*/
