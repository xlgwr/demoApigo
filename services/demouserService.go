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
func (s *DemoUserServices) Add(a m.DemoUser) (userid int64, err error) {
	o := orm.NewOrm()
	id, err := o.Insert(&a)
	if err != nil {
		return 0, err
	}
	return id, err
}
