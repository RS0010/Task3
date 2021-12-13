package database

const (
	Undo = iota
	Done
	All
)

func (user User) TodoAdd(data Data) {
	data.UserId = user.Id
	db.Create(&data)
}

func (user User) TodoDo(done bool, id ...int) {
	for _, id := range id {
		TodoDB.Where("userid = ? AND id = ?", user.Id, id).Update("done", done)
	}
}

func (user User) TodoDoAll(done bool) {
	TodoDB.Where("userid = ?", user.Id).Update("done", done)
}

func (user User) TodoList(page int, pageSize int, kind int, search ...string) (data []Data) {
	median := TodoDB.Limit(pageSize).Offset(page*pageSize).Where("userid = ?", user.Id)
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
	TodoDB.Where("userid = ? and id = ?", user.Id, id).Find(&data)
	if data.Id == 0 {
		return data, false
	}
	return data, true
}

func (user User) TodoDeleteId(id ...int) {
	TodoDB.Where("userid = ?", user.Id).Delete(&Data{}, id)
}

func (user User) TodoDeleteAll(kind int) {
	median := TodoDB.Where("userid = ?", user.Id)
	if kind != All {
		median = median.Where("done = ?", kind == Done)
	}
	median.Delete(&Data{})
}
