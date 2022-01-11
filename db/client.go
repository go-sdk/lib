package db

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/go-sdk/lib/conf"
	"github.com/go-sdk/lib/log"
)

type Client struct {
	*gorm.DB
}

var (
	mu sync.RWMutex
	db *Client
)

func init() {
	typ := strings.ToLower(conf.Get("db.type").StringD("sqlite"))
	dsn := conf.Get("db.dsn").String()
	if dsn == "" {
		return
	}
	x, err := New(typ, dsn)
	if err != nil {
		panic(err)
	}
	SetDefaultDB(x)
}

func New(typ, dsn string, def ...bool) (*Client, error) {
	var dial gorm.Dialector

	switch typ {
	case "sqlite":
		dial = sqlite.Open(dsn)
	case "postgres":
		dial = postgres.Open(dsn)
	case "mysql":
		dial = mysql.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported db driver: [%s]", typ)
	}

	cfg := &gorm.Config{}
	cfg.Logger = logger.Discard
	if !conf.Get("db.log.disable").Bool() {
		cfg.Logger = &gormLogger{
			e:                 log.DefaultLogger().WithField("span", "db"),
			UseInfoSQL:        conf.Get("db.log.use_info_sql").Bool(),
			ShowNotFoundError: conf.Get("db.log.show_not_found_error").Bool(),
			SlowThreshold:     conf.Get("db.log.slow_threshold").DurationD(500 * time.Second),
			LogLevel:          gormLogInfo,
		}
	}
	cfg.DisableForeignKeyConstraintWhenMigrating = true
	cfg.AllowGlobalUpdate = true
	cfg.QueryFields = true
	cfg.CreateBatchSize = 100

	d, err := gorm.Open(dial, cfg)
	if err != nil {
		return nil, err
	}

	x := &Client{DB: d}

	if len(def) == 0 || !def[0] {
		return x, nil
	}

	SetDefaultDB(x)

	return db, nil
}

func Default() *Client {
	mu.RLock()
	defer mu.RUnlock()
	return db
}

func SetDefaultDB(x *Client) {
	mu.Lock()
	defer mu.Unlock()
	db = x
}
