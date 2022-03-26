package tests

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/wuruipeng404/scaffold/orm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gdb *gorm.DB

func init() {
	initDB()
}

func initDB() {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "mysql", "localhost", 8306, "rumple",
	)

	if gdb, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                     dsn,
		DefaultStringSize:       256,
		DontSupportRenameIndex:  true,
		DontSupportRenameColumn: true,
	}), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
	}); err != nil {
		log.Fatalf("create orm connect error:%s", err)
	}

	sqlDb, err := gdb.DB()
	if err != nil {
		log.Fatalf("get sql error:%s", err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	log.Println("connect mysql success")
}

type Slice []string

type RumpleTest struct {
	orm.UModel
	Name string `json:"name"`
	Age  string `json:"age"`
	Sli  Slice  `json:"sli"`
}

func (s *Slice) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal json value: %v", value)
	}

	var result []string
	err := json.Unmarshal(bytes, &result)

	*s = result
	return err
}

func (s Slice) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}

	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

func (Slice) GormDataType() string {
	return "text"
}

func TestOrm(t *testing.T) {
	gdb.AutoMigrate(new(RumpleTest))

	gdb.Create(&RumpleTest{
		Name: "rumple111",
		Age:  "111122222",
		Sli:  Slice{"1", "2", "3"},
	})
}
