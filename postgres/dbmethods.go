package postgres

import "fmt"

// GetUsersByName is called within our user query for graphql
func (d *DatabaseInit) GetUsersByName(name string) []User {
	// Create slice of Users for our response
	users := []User{}
	// Prepare query, takes a name argument, protects from sql injection
	stmt := d.Raw("SELECT * FROM users WHERE name=$1", name).Find(&users)
	if stmt.Error != nil {
		fmt.Println("GetUserByName Preperation Err: ", stmt.Error)
	}

	return users
}
