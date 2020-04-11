// db
// Perform all db related operation here.
// Let's not store data in DB for now, instead i will be using global variable for time being.
package db

import (
	"fmt"
	model "mycircleapi/models"

	"time"
)

type DB struct{}

var messages = make(map[string][]model.Message)

//AddMessage Add new User
func (this *DB) AddMessage(user string, m *model.Message) error {
	if m == nil {
		return fmt.Errorf("Found nil message.")
	}
	m.LastUpdated = time.Now().UTC()
	if msgs, k := messages[user]; k {
		msgs = append(msgs, *m)
		messages[user] = msgs
	} else {
		messages[user] = []model.Message{*m}
	}

	return nil
}

//GetMessage Add new User
func (this *DB) GetMessage(user string) (msgs []model.Message, err error) {
	var k bool
	if msgs, k = messages[user]; !k {
		return msgs, fmt.Errorf("Messages not found")
	}
	return msgs, nil
}
