package framework

/*
import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-adodb"
)

type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         SA
}

type SA struct {
	user   string
	passwd string
}

func (m *Mssql) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.dataSource)
	if m.windows {
		// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器(需要在安装sqlserver时候设置)，
		// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
		conf = append(conf, "integrated security=SSPI")
	}
	conf = append(conf, "Initial Catalog="+m.database)
	conf = append(conf, "user id="+m.sa.user)
	conf = append(conf, "password="+m.sa.passwd)

	m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db := Mssql{
		dataSource: "BAIYUN-MOBL1",
		database:   "SdxOneStopDB",
		// windwos: true 为windows身份验证，false 必须设置sa账号和密码
		windows: true,
		sa: SA{
			user:   "sdxonestop",
			passwd: "baiyun+=1992",
		},
	}
	// 连接数据库
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return
	}
	defer db.Close()
	// 执行SQL语句
	rows, err := db.Query("select * from Book")
	if err != nil {
		fmt.Println("query: ", err)
		return
	}
	//fmt.Printf(rows.ColumnTypes())
	for rows.Next() {
		var ID int
		var Isbn int
		var Title string
		var Firstname string
		var Lastname string
		rows.Scan(&ID, &Isbn, &Title, &Firstname, &Lastname)
		fmt.Println(strconv.Itoa(ID), strconv.Itoa(Isbn), Title, Firstname, Lastname)
	}
}
*/
