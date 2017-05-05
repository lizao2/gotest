package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var (
    dbhostsip  = "127.0.0.1:3306"//IP地址
    dbusername = "root"//用户名
    dbpassword = "root"//密码
    dbname     = "Test"//表名
)

func main() {
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Test?charset=utf8")
    checkErr(err)

    //插入数据
    stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
    checkErr(err)

    res, err := stmt.Exec("码农", "研发部门", "2016-03-06")
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println("insertid:%d", id)
    //更新数据
    stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    checkErr(err)

    res, err = stmt.Exec("码农二代", id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println("affected:%d", affect)

    //查询数据
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println("uid:%d", uid)
        fmt.Println("username:", username)
        fmt.Println(department)
        fmt.Println(created)
        fmt.Println("...................................")
    }

    //删除数据
//    stmt, err = db.Prepare("delete from userinfo where uid=?")
 //   checkErr(err)

  //  res, err = stmt.Exec(id)
   // checkErr(err)

    //affect, err = res.RowsAffected()
    //checkErr(err)

   // fmt.Println("delete:%d", affect)

    db.Close()

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

