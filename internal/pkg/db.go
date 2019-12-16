package pkg

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"
	"github.com/go-xorm/xorm"
	"time"
)

const (
	DriverName     = "mysql"
	driverTimezone = "Asia/Shanghai"
)

func NewMysqlEngine(driverName string, sourceName string, prefix string) (*xorm.Engine, error) {
	mysqlEngine, err := xorm.NewEngine(driverName, sourceName)
	if err != nil {
		return nil, err
	}
	mysqlEngine.SetMapper(core.SnakeMapper{})
	mysqlEngine.TZLocation, err = time.LoadLocation(driverTimezone)

	if err != nil {
		return nil, err
	}
	tableMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)
	mysqlEngine.SetTableMapper(tableMapper)
	return mysqlEngine, nil
}

func GenerateMysqlSource(username string, password string, dbhost string, dbport int, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, dbhost, dbport, database)
}
