package dao

import (
	"SqlRestDemo/framework"
	"fmt"

	//Only need init package
	_ "github.com/mattn/go-adodb"
)

//ConnectDB is to Init and connect DB
type ConnectDB struct {
}

//InitDB is to initialize DB
func (*ConnectDB) InitDB() framework.Mssql {

	var user = framework.SA{
		Username: "sdxonestop",
		Password: "baiyun+=1992",
	}

	db := framework.Mssql{
		DataSource: "BAIYUN-MOBL1",
		Database:   "SdxOneStopDB",
		// windwos: true 为windows身份验证，false 必须设置sa账号和密码
		Windows: true,
		Sa:      user,
	}
	// 连接数据库
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return db
	} else {
		fmt.Println("Succeed to open DB...")
	}
	//defer db.Close()

	return db
}
