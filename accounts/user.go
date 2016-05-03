package accounts

import (
	"time"
	"fmt"
	"errors"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int  `orm:"auto"`
	Name string `orm:"size(100);unique"`
	Salt string
	Age  int
	Password string `orm:"size(200)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

const user_table = "users"

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableEngine() string {
	return "INNODB"
}

func (u *User) TableName() string {
	return user_table
}

func CreateUser(user *User) (*User, error) {
	o := orm.NewOrm()

	id, err := o.Insert(user)
	if err != nil {
		fmt.Println(err)
		fmt.Println(id)
		return user, err
	}

	return user, nil
}

func GetUser(id string) (*User, error) {
	o := orm.NewOrm()

	uuid, _ :=  strconv.Atoi(id)
	user := User{ Id: uuid }

	err := o.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println(errors.New("not"))
		return &user, errors.New("404")
	}

	return &user, nil
}

func UpdateUser(user *User) (*User, error) {
	o := orm.NewOrm()

	id, err := o.Update(user)
	if err != nil {
		fmt.Println(err)
		fmt.Println(id)
		return user, err
	}

	return user, nil
}

func DeleteUser(id string) (bool, error) {
	o := orm.NewOrm()

	uuid, _ :=  strconv.Atoi(id)
	user := &User{ Id: uuid }

	_, err := o.Delete(user)
	if err != nil {
		return false, err
	}

	return true, nil
}


func GetAllUsers() []*User {
	o := orm.NewOrm()

	var users []*User
	qs := o.QueryTable(user_table)
	_, err := qs.All(&users)

	if (err != nil) {
		fmt.Println(err)
	}

	return users
}




