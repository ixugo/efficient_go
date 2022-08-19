package sqlite3

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	// "gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"not null;size:10"`
}

func TestSQLite(t *testing.T) {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	require.NoError(t, err)
	defer func() {
		if err != nil {
			os.Remove("./gorm.db?cache=shared&mode=rwc")
		}
	}()

	err = db.Set("gorm:table_options", "").AutoMigrate(new(User))
	require.NoError(t, err)

	dbPool, err := db.DB()
	dbPool.SetMaxIdleConns(1)
	dbPool.SetMaxOpenConns(1)
	dbPool.SetConnMaxLifetime(time.Hour)
	require.NoError(t, err)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			u := User{Name: fmt.Sprintf("hello_%d", i)}
			err := db.Create(&u).Error
			require.NoError(t, err)
		}(i)
	}
	wg.Wait()

	var user []User
	err = db.Find(&user).Error
	require.NoError(t, err)
	for _, v := range user {
		fmt.Printf("%+v\n", v)
	}

	os.Remove("./gorm.db")
}
