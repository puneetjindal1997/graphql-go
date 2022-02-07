package postgres

import (
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Db is our database struct used for interacting with the database
type DatabaseInit struct {
	*gorm.DB
}

var DbMethods PostGresMethods

// methods interface
type PostGresMethods interface {
	GetUsersByName(string) []User
}

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func New(connString string) (*DatabaseInit, error) {
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("connected db")
	return &DatabaseInit{db}, nil
}

// ConnString returns a connection string based on the parameters it's given
// This would normally also contain the password, however we're not using one
func ConnString(host string, port int, user, dbName, password string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, password,
	)
}

// User shape
type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}
