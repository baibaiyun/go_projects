package framework

import (
	"database/sql"
	"strings"

	//Only need init package
	_ "github.com/mattn/go-adodb"
)

//Mssql is to configuration for database initialization
type Mssql struct {
	*sql.DB
	DataSource string
	Database   string
	Windows    bool
	Sa         SA
}

//Open is connect with database
func (m *Mssql) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.DataSource)
	if m.Windows {
		// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器(需要在安装sqlserver时候设置)，
		// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
		conf = append(conf, "integrated security=SSPI")
	}
	conf = append(conf, "Initial Catalog="+m.Database)
	conf = append(conf, "user id="+m.Sa.Username)
	conf = append(conf, "password="+m.Sa.Password)

	m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		return err
	}
	return nil
}
