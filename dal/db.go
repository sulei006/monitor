package dal

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"monitor/model"
)

var (
	Db  *sql.DB
	err error
)

// 获取数据库连接
func DatabaseConnect(host string, port string, user string, pwd string, database string) error {
	//连接字符串获取连接url
	connectionStr := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + database
	Db, err = sql.Open("mysql", connectionStr)
	if err != nil {
		log.Fatalf("get the connection to db fail, error: %s", err)
	}
	return err
}

func GetSteamDataFromDb(lower, upper int) ([]*model.StremData, error) {
	fmt.Println(lower)
	fmt.Println(upper)

	queryStr := "select * from stream where id>= ? and id <= ?"
	//定义放回结果切片
	var streamData []*model.StremData
	//执行
	rows, err := Db.Query(queryStr, lower, upper)
	if err != nil {
		return nil, fmt.Errorf("get data from db error : %s", err)
	}

	for rows.Next() {
		sdata := &model.StremData{}
		err := rows.Scan(&sdata.Id,&sdata.Output,&sdata.Input)
		if err != nil {
			return nil, fmt.Errorf("get the result error : %s", err)
		}
		streamData = append(streamData, sdata)
	}
	return streamData, nil
}
