package models

type MonitorLog struct {
	Err     error
	Id      string `bson:"_id omitempty"`
	Message string
	Path    string
	Code    int
}
