package database

import (
	"context"
	"fmt"
	"os"

	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
	"github.com/rhul-compsoc/compsoc-api-go/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbMaker func() *gorm.DB

const (
	MembersTable  = "members"
	UsersTable    = "users"
	EventsTable   = "events"
	AdminsTable   = "admins"
	StudentsTable = "students"
)

// Stores a *gorm.DB and mutex.
type Store struct {
	db *gorm.DB
}

// Create a new pointer to a Store.
func New(m DbMaker) Store {
	return Store{
		db: m(),
	}
}

// AutoMigrate models to tables
func (s *Store) AutoMigrate() error {
	err := s.db.Table(MembersTable).AutoMigrate(models.MemberModel{})
	if err != nil {
		return err
	}

	err = s.db.Table(UsersTable).AutoMigrate(models.UserModel{})
	if err != nil {
		return err
	}

	err = s.db.Table(EventsTable).AutoMigrate(models.EventModel{})
	if err != nil {
		return err
	}

	err = s.db.Table(AdminsTable).AutoMigrate(models.Admin{})
	if err != nil {
		return err
	}

	err = s.db.Table(StudentsTable).AutoMigrate(models.StudentModel{})
	if err != nil {
		return err
	}

	return err
}

// Ping the database.
func (s *Store) Ping() error {
	ctx := context.Background()
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	return err
}

// Close the database.
func (s *Store) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return nil
	}

	err = db.Close()
	return err
}

// Creates a pointer to a gorm db.
//
// This uses environmental variables for the dsn.
//
// A connection is then opened, checked for errors and returned.
//
// Keys for environmental variables:
//   - DB_ADDR : stores the address (IP)
//   - DB_PORT : stores the port
//   - DB_USER : stores the username
//   - DB_PASS : stores the password
//   - DB_NAME : stores the database name
func PostgresDB() DbMaker {
	return func() *gorm.DB {
		dsn := fmt.Sprintf(`
		host=%s 
		user=%s 
		password=%s 
		dbname=%s 
		port=%s 
		sslmode=disable 
		TimeZone=Europe/London`,
			os.Getenv("DB_ADDR"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		util.ErrOut(err)
		return db
	}
}

// Creates a pointer to a gorm db.
//
// This uses environmental variables for the SQLite file path.
//
// A connection is then opened, checked for errors and returned.
//
// Keys for environmental variables:
//   - SQLIET_DB : stores location of the sqlite db file.
func SQLiteDB() DbMaker {
	return func() *gorm.DB {
		db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB")), &gorm.Config{})
		util.ErrOut(err)
		return db
	}
}
