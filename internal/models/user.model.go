package models

type User struct {
	UserName string `bson:"username"`
	Logs     []Log  `bson:"logs"`
	Id       int    `bson:"_id, omitempty"`
}
