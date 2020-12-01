package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsName := "root:123456@tcp(47.110.141.191:3306)/sjyb?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", dsName)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(3)
	db.SetConnMaxLifetime(7 * time.Hour)

	fmt.Println(db.Query("select now()"))
	/*
		db.Exec()
		db.ExecContext()
		db.Query()
		db.QueryContext()
		db.QueryRow()
		db.QueryRowContext()
		db.Prepare/PrepareContext() // 创建预编译对象
		db.Begin/BeginTx() // 开启一个事务, 返回的Tx事务对象会被绑定到单个连接
		db.Stats() 返回数据库统计信息
	*/

}
