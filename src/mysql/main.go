package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "cdb_outerroot:hongan666@tcp(56c1f4575346c.gz.cdb.myqcloud.com:12818)/himall")
	checkError(err)

	//insert
	stmt, err := db.Prepare("INSERT azt_login_ticket SET UserId=?,Ticket=?,CreateTime=?, `Desc`=?")
	checkError(err)

	res, err := stmt.Exec(9, "bbb", "2019-05-11 15:24:50", "ddddd")
	checkError(err)

	newId, err := res.LastInsertId()
	checkError(err)

	fmt.Println("new id=", newId)

	//update
	stmt, err = db.Prepare("update azt_login_ticket set Ticket=? where Id = ?")
	checkError(err)

	res, err = stmt.Exec("tttttt", 12210)
	checkError(err)

	aff, err := res.RowsAffected()
	checkError(err)

	fmt.Print("aff=", aff)

	//select
	rows, err := db.Query("select * from azt_login_ticket where id>=12210")
	checkError(err)

	for rows.Next() {
		var id int64
		var userId int64
		var ticket string
		var createTime string
		var desc string
		rows.Scan(&id, &userId, &ticket, &createTime, &desc)
		fmt.Println(id)
		fmt.Println(userId)
		fmt.Println(ticket)
		fmt.Println(createTime)
		fmt.Println(desc)
	}

	//delete
	stmt, err = db.Prepare("delete from azt_login_ticket where id=?")
	checkError(err)

	res, err = stmt.Exec(12215)
	checkError(err)

	aff, err = res.RowsAffected()
	checkError(err)

	fmt.Print("deleted, aff=", aff)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
