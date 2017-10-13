package models

import (
	// "github.com/astaxie/beego/orm"
)

var (
	Objects map[string]*Object
)

type Object struct {
	Id int
	ObjectId   string
	Score      int64
	PlayerName string
}

func init() {
	//orm.RegisterModel(new(Object))

	Objects = make(map[string]*Object)
	Objects["hjkhsbnmn123"] = &Object{1,"hjkhsbnmn123", 100, "astaxie"}
	Objects["mjjkxsxsaa23"] = &Object{2,"mjjkxsxsaa23", 101, "someone"}
}



/*
func AddOne(object Object) (ObjectId string) {
	object.ObjectId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Objects[object.ObjectId] = &object
	return object.ObjectId
}

func GetOne(ObjectId string) (object *Object, err error) {
	if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

func GetAll() map[string]*Object {
	return Objects
}


func Update(ObjectId string, Score int64) (err error) {
	if v, ok := Objects[ObjectId]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ObjectId Not Exist")
}

func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}*/

