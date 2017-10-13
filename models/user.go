package models

import (
	"github.com/astaxie/beego/orm"
	"time"

)

var (
	UserList map[int]*User
)

func init() {
	orm.RegisterModel(new(User))
	/*
	o := orm.NewOrm()
	o.Using("default")
	profile := new(Profile)
	profile.Age = 24
	profile.Address = "zhongguo"
	profile.Gender = "male"
	profile.Email = "124296140@qq.com"

	user := new(User)
	user.Username = "zhuyunsong"
	user.Password = "123213"
	user.Profile = profile

	o.Insert(profile)
	o.Insert(user)
	*/
	/*
	UserList = make(map[int]*User)
	u := User{1, "astaxie", "11111", Profile{1,"male", 20, "Singapore", "astaxie@gmail.com"}}
	UserList[1] = &u
	*/
}

type User struct {
	Id       		int
	CardNum 		string
	Username 		string `orm:"size(64)" json:"username"`
	CreateDate 		time.Time          `orm:"auto_now_add;type(datetime);null" json:"create_date"`
	Yueee	 		int
	Guoqi   		time.Time		`orm:"auto_now_add;type(datetime);null" json:"guoqi"`
	Status    		string
	UserType  		string
	Phone     		string
	Password 		string
	Gender   		string
	Touxiang  		string
	Jiankangzhishu 	string
	Shengao  		string
	Tizhong  		string
	Age      		int
	Address  		string
	Email    		string
}

/*
func AddUser(u User) int {
	u.Id = 10
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid int) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[int]*User {
	return UserList
}

func UpdateUser(uid int, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid int) {
	delete(UserList, uid)
}
*/