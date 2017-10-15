package leaniot
import (
    "log"
 
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


func writeToTidb() error {
    db, err := sql.Open("mysql", "root@tcp(localhost:4000)/test?charset=utf8")
    stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
    if err != nil {
        log.Fatal(err)
        return err
    }


    res, err := stmt.Exec("astaxie", "R&D", "2012-12-09")
    if err != nil {
        log.Fatal(err)
        return err
    }



    id, err := res.LastInsertId()
    if err != nil {
        log.Fatal(err)
        return err
    }

    log.Print("laster insert id ", id)
    return nil

}

