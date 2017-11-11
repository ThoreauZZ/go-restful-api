package model

type User struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Ext        int    `json:"ext"`
	Privileged bool   `json:"privileged"`
}
