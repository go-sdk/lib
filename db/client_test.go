package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-sdk/lib/crypto"
	"github.com/go-sdk/lib/token"
)

// For Postgres:
//     db:
//      type: postgres
//      dsn: host=localhost user=postgres password=p1ssw0rd dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai statement_cache_mode=describe
// For Sqlite:
//      db:
//       type: sqlite
//       dsn: file:sqlite?mode=memory&cache=shared

type User struct {
	Meta
	Name   string `gorm:"size:64"`
	Email  string `gorm:"size:128"`
	Mobile string `gorm:"size:20"`
}

func (User) TableName() string {
	return "user1"
}

type UserD struct {
	MetaD
	Name   string `gorm:"size:64;uniqueIndex:uk_deleted_at"`
	Email  string `gorm:"size:128"`
	Mobile string `gorm:"size:20"`
}

func (UserD) TableName() string {
	return "user2"
}

var ctx = token.New("*", "123456", 0).WithContext()

func user(names ...string) *User {
	name := fmt.Sprintf("%s-%s-%d", "admin", crypto.RandString(6), time.Now().UnixNano())
	if len(names) > 0 && names[0] != "" {
		name = names[0]
	}
	return &User{
		Name:   name,
		Email:  "admin@example.com",
		Mobile: "12312341234",
	}
}

func userD(names ...string) *UserD {
	name := fmt.Sprintf("%s-%s-%d", "admin", crypto.RandString(6), time.Now().UnixNano())
	if len(names) > 0 && names[0] != "" {
		name = names[0]
	}
	return &UserD{
		Name:   name,
		Email:  "admin@example.com",
		Mobile: "12312341234",
	}
}

func TestDefaultWithMeta(t *testing.T) {
	assert.NoError(t, Default().AutoMigrate(&User{}))

	res := Default().Create(user())
	assert.NoError(t, res.Error)
	assert.Equal(t, int64(1), res.RowsAffected)
	assert.NoError(t, Default().WithContext(ctx).Create(user()).Error)
	assert.NoError(t, Default().Clauses(OnConflict{DoNothing: true}).Create(user("1")).Error)
	assert.NoError(t, Default().Clauses(OnConflict{Columns: []Column{{Name: "id"}}, DoUpdates: Assignments(map[string]interface{}{"name": "user01"})}).Create(user("1")).Error)

	res = Default().First(&User{})
	assert.NoError(t, res.Error)
	assert.Equal(t, int64(1), res.RowsAffected)
	assert.Equal(t, ErrRecordNotFound, Default().Take(&User{}, Eq{Column: "id", Value: "2"}).Error)

	assert.NoError(t, Default().WithContext(ctx).Model(&User{}).Where(Eq{Column: "id", Value: "1"}).Update("email", "user01@example.com").Error)

	assert.NoError(t, Default().WithContext(ctx).Delete(&User{}, Eq{Column: "id", Value: "1"}).Error)
	assert.NoError(t, Default().Unscoped().Find(&User{}).Error)
}

func TestDefaultWithMetaD(t *testing.T) {
	assert.NoError(t, Default().AutoMigrate(&UserD{}))

	res := Default().Create(userD())
	assert.NoError(t, res.Error)
	assert.Equal(t, int64(1), res.RowsAffected)
	assert.NoError(t, Default().WithContext(ctx).Create(userD()).Error)
	assert.NoError(t, Default().Clauses(OnConflict{DoNothing: true}).Create(userD("1")).Error)
	assert.NoError(t, Default().Clauses(OnConflict{Columns: []Column{{Name: "name"}, {Name: "deleted_at"}}, DoUpdates: Assignments(map[string]interface{}{"name": "user01"})}).Create(userD("1")).Error)

	res = Default().First(&UserD{})
	assert.NoError(t, res.Error)
	assert.Equal(t, int64(1), res.RowsAffected)
	assert.Equal(t, ErrRecordNotFound, Default().Take(&UserD{}, Eq{Column: "id", Value: "2"}).Error)

	assert.NoError(t, Default().WithContext(ctx).Model(&UserD{}).Where(Eq{Column: "id", Value: "1"}).Update("email", "user01@example.com").Error)

	assert.NoError(t, Default().WithContext(ctx).Delete(&UserD{}, Eq{Column: "id", Value: "1"}).Error)
	assert.NoError(t, Default().Unscoped().Find(&UserD{}).Error)
}
