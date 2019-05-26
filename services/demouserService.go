package services

import (
	"github.com/astaxie/beego/orm"
	m "github.com/xlgwr/demoApigo/models"
)

type DemoUserServices struct {
}

func init() {
	orm.RegisterModel(new(m.DemoUser))
}

//Add func ..
//add a demouser
//return id
func (s *DemoUserServices) Add(a m.DemoUser) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&a)
	if err != nil {
		return 0, err
	}
	return id, err
}

//Get user by id
//ret demouser
func (s *DemoUserServices) Get(id int) (m.DemoUser, error) {
	o := orm.NewOrm()
	u := m.DemoUser{ID: int64(id)}
	err := o.Read(&u)
	return u, err
}
