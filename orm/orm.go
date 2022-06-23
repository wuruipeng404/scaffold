package orm

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
)

type DbType string

const (
	MySQL    DbType = "mysql"
	Postgres DbType = "postgres"
)

var cli *gorm.DB

type InitOption struct {
	Type   DbType
	User   string
	Pass   string
	DbName string
	Host   string
	Port   int

	MaxIdleConn            int
	MaxOpenConn            int
	ConnMaxLifetime        time.Duration
	LogLevel               glog.LogLevel
	SkipDefaultTransaction bool
}

func Init(opt *InitOption) {
	var (
		err  error
		dsn  string
		dia  gorm.Dialector
		conf = &gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: opt.SkipDefaultTransaction,
		}
	)

	if opt == nil {
		log.Fatal("orm options need")
	}

	if opt.MaxIdleConn == 0 {
		opt.MaxIdleConn = 10
	}

	if opt.MaxOpenConn == 0 {
		opt.MaxOpenConn = 100
	}

	if opt.ConnMaxLifetime == 0 {
		opt.ConnMaxLifetime = time.Hour
	}

	if opt.LogLevel != glog.Silent {
		var prefix = "\n"

		if runtime.GOOS == "windows" {
			prefix = "\r\n"
		}

		newLogger := glog.New(
			log.New(os.Stdout, prefix, log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
			glog.Config{
				SlowThreshold:             time.Second,  // 慢 SQL 阈值
				LogLevel:                  opt.LogLevel, // 日志级别
				IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,         // 彩色打印
			},
		)
		conf.Logger = newLogger
	} else {
		conf.Logger = glog.Default.LogMode(glog.Error)
	}

	switch opt.Type {
	case MySQL:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			opt.User, opt.Pass, opt.Host, opt.Port, opt.DbName)

		dia = mysql.New(mysql.Config{
			DSN:                     dsn,
			DefaultStringSize:       256,
			DontSupportRenameIndex:  true,
			DontSupportRenameColumn: true,
		})

	case Postgres:
		dsn = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
			opt.Host, opt.Port, opt.User, opt.Pass, opt.DbName,
		)

		dia = postgres.Open(dsn)
	default:
		log.Fatalf("unsupported db type:%s", opt.Type)
	}

	if cli, err = gorm.Open(dia, conf); err != nil {
		log.Fatalf("gorm open error:%s", err)
	}

	sqlDb, err := cli.DB()
	if err != nil {
		log.Fatalf("get sql db error:%s", err)
	}

	sqlDb.SetMaxIdleConns(opt.MaxIdleConn)
	sqlDb.SetMaxOpenConns(opt.MaxOpenConn)
	sqlDb.SetConnMaxLifetime(opt.ConnMaxLifetime)
}

func C() *gorm.DB {
	return cli
}
