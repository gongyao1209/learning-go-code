package main

import (
	"../mydatabase/model"
	"../mydatabase/db"
	"fmt"
)

func main() {

	Db := db.Get()
	defer Db.Close()

	push := model.GetPush(Db,32006)
	
	fmt.Printf("%v\n", push)

	//dbw.changeStatus(32006, 1)


	//var name int
	//err = dbw.Db.QueryRow("SELECT `box_id` FROM `push` WHERE `id` = 32001").Scan(&name)
	//
	//fmt.Println(name)
	/**
	'r_database' => array(
        'adapter' => 'Mysql',
        'host' => '127.0.0.1',
        'username' => 'root',
        'password' => 'S12p_w99Q',
//        'dbname' => 'phalcon_test',
        'dbname' => 'xiangxin0828',
    ),
	 */
}

//type DbWorker struct {
//	Dsn	string
//	Db	*sql.DB
//	Push push
//}
//
//type push struct {
//	Id int
//	Type int
//	Status int
//	BoxId int
//	TimeStamp int
//	Time string
//	CreateUserId int
//	CreateTime string
//	UpdateUserId int
//	UpdateTime string
//}
//
//
//func (dbw *DbWorker) insertData() {
//	//stmt, _ = dbw.Db.Prepare("INSERT INTO ")
//}
//
//func (dbw *DbWorker) queryData(id int) {
//
//	stmt,_ := dbw.Db.Prepare("SELECT * FROM `push` AS `p` WHERE `p`.`id` = ? LIMIT 1")
//	defer stmt.Close()
//
//	dbw.Push = push{}
//
//	rows, err := stmt.Query(id)
//	if err != nil {
//		fmt.Printf("SELECT ERROR %s\n", err.Error())
//		return
//	}
//
//	for rows.Next() {
//		err = rows.Scan(&dbw.Push.Id, &dbw.Push.Type, &dbw.Push.Status, &dbw.Push.BoxId, &dbw.Push.TimeStamp, &dbw.Push.Time, &dbw.Push.CreateUserId, &dbw.Push.CreateTime, &dbw.Push.UpdateUserId, &dbw.Push.UpdateTime)
//		if err != nil {
//			fmt.Printf(err.Error())
//			continue
//		}
//
//		fmt.Println("get data, id:", dbw.Push.Id, ", box_id:", dbw.Push.BoxId, "\n")
//	}
//
//	err = rows.Err()
//	if err != nil {
//		fmt.Printf(err.Error())
//	}
//}
//
//func (dbw *DbWorker) changeStatus(id int, status int) {
//	stmt, _ := dbw.Db.Prepare("UPDATE `push` SET `status` = ? WHERE `id` = ?")
//	defer stmt.Close()
//
//	ret, err := stmt.Exec(status, id)
//	if err != nil {
//		fmt.Printf("Update data error: %v\n", err.Error())
//		return
//	}
//
//	if affected, err := ret.RowsAffected(); err == nil {
//		fmt.Println("RowsAffected:", affected)
//	}
//}
