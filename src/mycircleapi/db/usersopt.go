// db
// Perform all db related operation here.
// Let's not store data in DB for now, instead i will be using global variable for time being.
package db

import (
	"crypto/rand"
	"fmt"
	"log"

	model "mycircleapi/models"
)

var alluser = make(map[string]*model.User)

//RegisterUser Add new User
func (this *DB) RegisterUser(user *model.User) error {
	if user == nil {
		return fmt.Errorf("Found nil user.")
	}
	log.Println("Got User ", user)
	alluser[user.Email] = user
	return nil
}

//Login handling User login
func (this *DB) Login(lr *model.LoginRequest) (*model.LoginResponse, error) {
	if lr == nil {
		return nil, fmt.Errorf("Found nil Login request.")
	}

	// Check User exist
	u, k := alluser[lr.Email]
	if !k {
		return nil, fmt.Errorf("User not found.")
	}

	// Check User credentials
	if u.Password == lr.Password {
		response := new(model.LoginResponse)
		//Create random token
		b := make([]byte, 32)
		rand.Read(b)
		response.Token = fmt.Sprintf("%x", b)
		response.Id = u.Email
		return response, nil
	}

	return nil, fmt.Errorf("Invalid Credentials.")

}

//GetUser Get Specific User
func (this *DB) GetUser(email string) (*model.User, error) {
	// Check User exist
	u, k := alluser[email]
	if !k {
		return nil, fmt.Errorf("User not found.")
	}

	return u, nil
}
