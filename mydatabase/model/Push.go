package model

import (
	"database/sql"
	"fmt"
)

type Push struct {
	Id int
	Type int
	Status int
	BoxId int
	TimeStamp int
	Time string
	CreateUserId int
	CreateTime string
	UpdateUserId int
	UpdateTime string
}


func GetPush(Db *sql.DB, id int) Push {
	stmt,_ := Db.Prepare("SELECT * FROM `push` AS `p` WHERE `p`.`id` = ? LIMIT 1")
	defer stmt.Close()

	data := Push{}

	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("SELECT ERROR %s\n", err.Error())
		return data
	}

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Type, &data.Status, &data.BoxId, &data.TimeStamp, &data.Time, &data.CreateUserId, &data.CreateTime, &data.UpdateUserId, &data.UpdateTime)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		fmt.Println("get data, id:", data.Id, ", box_id:", data.BoxId, "\n")
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}

	return data
}