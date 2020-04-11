package model

import (
	"time"
)

type Message struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userid"`
	Message     string    `json:"message"`
	LastUpdated time.Time `json:"lastupdated"`
}

// ContactsList
type Registration struct {
	DeviceId string `json:"deviceid"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
}

type Contact struct {
	DeviceId        string `json:"deviceid"`
	ContactIdentity int    `json:"contactidentity"`
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	AffectedCount   int    `json:"affectedcount"`
	Proximity       string `json:"proximity"`
}

type Affected struct {
	Mobile         string `json:"mobile"`
	AffectedStatus int    `json:"affectedstatus"`
}
