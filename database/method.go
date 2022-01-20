package database

import (
	"Task3/schemas"
	"Task3/tools"
)

const (
	All = iota
	Done
	Undo
)

func (user User) TodoAdd(data Data) Data {
	data.UserId = user.Id
	todoDb().Create(&data)
	return data
}

func (user User) TodoDo(done bool, id ...int) {
	for _, id := range id {
		todoDb().Where("userid = ? AND id = ?", user.Id, id).Update("done", done)
	}
}

func (user User) TodoDoAll(done bool) {
	todoDb().Where("userid = ?", user.Id).Update("done", done)
}

func (user User) TodoList(page int, pageSize int, kind int, search ...string) (data []Data) {
	median := todoDb().Limit(pageSize).Offset(page*pageSize).Where("userid = ?", user.Id)
	if kind != All {
		median = median.Where("done = ?", kind == Done)
	}
	for _, keyword := range search {
		median = median.Where("title like ?", "%"+keyword+"%")
	}
	median.Find(&data)
	return data
}

func (user User) TodoListId(id int) (Data, bool) {
	var data Data
	todoDb().Where("userid = ? and id = ?", user.Id, id).Find(&data)
	if data.Id == 0 {
		return data, false
	}
	return data, true
}

func (user User) TodoDeleteId(id ...int) {
	todoDb().Where("userid = ?", user.Id).Delete(&Data{}, id)
}

func (user User) TodoDeleteAll(kind int) {
	median := todoDb().Where("userid = ?", user.Id)
	if kind != All {
		median = median.Where("done = ?", kind == Done)
	}
	median.Delete(&Data{})
}

func UserAdd(u schemas.User) (user User) {
	user.Name = u.Username
	user.Code = tools.HashGenerate(u.Passcode)
	userDb().Create(&user)
	return user
}

func UserGet(id int) (user User, _ bool) {
	userDb().Where("id = ?", id).Find(&user)
	if user.Id == 0 {
		return User{}, false
	}
	return user, true
}
