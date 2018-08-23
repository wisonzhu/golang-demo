package main
import (
    "fmt"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
    if err != nil{
        log.Fatal(err)
    } else {
    fmt.Println("this is test")
}
    defer db.Close()

 _, err = db.Exec("CREATE TABLE IF NOT EXISTS test.hello(world varchar(50))")
    if err != nil{
        log.Fatalln(err)
    }


rs, err := db.Exec("INSERT INTO test.hello(world) VALUES ('hello world')")
    if err != nil{
        log.Fatalln(err)
    }
    rowCount, err := rs.RowsAffected()
    if err != nil{
        log.Fatalln(err)
    }
    log.Printf("inserted %d rows", rowCount)


rows, err := db.Query("SELECT world FROM test.hello")
    if err != nil{
        log.Fatalln(err)
    }

    for rows.Next(){
        var s string
        err = rows.Scan(&s)
        if err !=nil{
            log.Fatalln(err)
        }
        log.Printf("found row containing %q", s)
    }
    rows.Close()

}
