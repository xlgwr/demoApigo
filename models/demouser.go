package models

import "time"

// DemoUser ..a
type DemoUser struct {
	ID       int64
	Username string
	Password string
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}
