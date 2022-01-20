package controller

import (
	"Task3/database"
	"Task3/middleware/Authenticator"
	. "Task3/route"
	"Task3/schemas"
	"Task3/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	group := Version[V1].Group("/todo").Use(Authenticator.Authenticator(tools.AuthenticationToken))
	group.POST(todoAdd)
	group.GET(todoGet)
	group.DELETE(todoDelete)
	subgroup := group.Group("/done")
	subgroup.POST(todoDone)
	subgroup.DELETE(todoUndo)
}

func todoAdd(context *gin.Context) {
	var todo schemas.TodoAdding
	var reply schemas.Reply

	if err := context.ShouldBindJSON(&todo); err != nil {
		reply.Status = http.StatusBadRequest
		reply.Error = err.Error()
		context.JSON(reply.Status, reply)
		return
	}

	user := context.MustGet("user").(database.User)
	data := user.TodoAdd(database.Data{
		Title:     todo.Title,
		Content:   todo.Content,
		StartTime: todo.StartTime,
		Deadline:  todo.Deadline,
	})

	reply.Status = http.StatusOK
	reply.Data = gin.H{
		"item":  data,
		"total": 1,
	}
	context.JSON(reply.Status, reply)
}

func todoDelete(context *gin.Context) {
	var reply schemas.Reply
	var deletes schemas.TodoDeleting

	if err := context.ShouldBindQuery(&deletes); err != nil {
		reply.Status = http.StatusBadRequest
		reply.Error = err.Error()
		context.JSON(reply.Status, reply)
		return
	}

	user := context.MustGet("user").(database.User)
	if deletes.Id == 0 {
		user.TodoDeleteAll(deletes.Type)
		if deletes.Type == database.All {
			reply.Data = "all"
		} else if deletes.Type == database.Done {
			reply.Data = "done"
		} else {
			reply.Data = "undo"
		}
	} else {
		user.TodoDeleteId(deletes.Id)
		reply.Data = deletes.Id
	}
	reply.Status = http.StatusOK
	context.JSON(reply.Status, reply)
}

func todoGet(context *gin.Context) {
	var reply schemas.Reply
	var gets schemas.TodoGetting

	if err := context.ShouldBindQuery(&gets); err != nil {
		reply.Status = http.StatusBadRequest
		reply.Error = err.Error()
		context.JSON(reply.Status, reply)
		return
	}

	user := context.MustGet("user").(database.User)
	if gets.Id != 0 {
		todo, ok := user.TodoListId(gets.Id)
		if ok {
			reply.Data = gin.H{
				"item":  []database.Data{todo},
				"total": 1,
			}
		} else {
			reply.Data = gin.H{
				"item":  []database.Data{},
				"total": 0,
			}
		}
	} else {
		todos := user.TodoList(gets.Page-1, 20, gets.Type, gets.Keyword)
		reply.Data = gin.H{
			"item":  todos,
			"total": len(todos),
		}
	}
	reply.Status = http.StatusOK
	context.JSON(reply.Status, reply)
}

func todoDone(context *gin.Context) {
	var reply schemas.Reply
	var done schemas.TodoDone

	if err := context.ShouldBindQuery(&done); err != nil {
		reply.Status = http.StatusBadRequest
		reply.Error = err.Error()
		context.JSON(reply.Status, reply)
		return
	}

	user := context.MustGet("user").(database.User)
	if done.Id == 0 {
		user.TodoDoAll(true)
		reply.Data = "all"
	} else {
		user.TodoDo(true, done.Id)
		reply.Data = done.Id
	}
	reply.Status = http.StatusOK
	context.JSON(reply.Status, reply)
}

func todoUndo(context *gin.Context) {
	var reply schemas.Reply
	var undo schemas.TodoDone

	if err := context.ShouldBindQuery(&undo); err != nil {
		reply.Status = http.StatusBadRequest
		reply.Error = err.Error()
		context.JSON(reply.Status, reply)
		return
	}

	user := context.MustGet("user").(database.User)
	if undo.Id == 0 {
		user.TodoDoAll(false)
		reply.Data = "all"
	} else {
		user.TodoDo(false, undo.Id)
		reply.Data = undo.Id
	}
	reply.Status = http.StatusOK
	context.JSON(reply.Status, reply)
}
