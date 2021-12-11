package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type (
	group struct {
		path       string
		apis       []api
		middleware []gin.HandlerFunc
		subGroups  []group
	}

	api struct {
		path    string
		POST    gin.HandlerFunc
		GET     gin.HandlerFunc
		DELETE  gin.HandlerFunc
		PATCH   gin.HandlerFunc
		PUT     gin.HandlerFunc
		OPTIONS gin.HandlerFunc
		HEAD    gin.HandlerFunc
		CONNECT gin.HandlerFunc
		TRACE   gin.HandlerFunc
	}
)

const apiVersionNumber = 1
const (
	v1 = iota
)

var version = make([]group, apiVersionNumber)

func Init(router *gin.Engine) {
	for i := 0; i < apiVersionNumber; i++ {
		version[i].path = "/v" + strconv.Itoa(i+1)
		version[i].handle(&router.RouterGroup)
	}
}

func (g group) handle(group *gin.RouterGroup) {
	root := group.Group(g.path)
	root.Use(g.middleware...)
	for _, api := range g.apis {
		api.handle(http.MethodPost, api.POST, root)
		api.handle(http.MethodGet, api.GET, root)
		api.handle(http.MethodDelete, api.DELETE, root)
		api.handle(http.MethodPatch, api.PATCH, root)
		api.handle(http.MethodPut, api.PUT, root)
		api.handle(http.MethodOptions, api.OPTIONS, root)
		api.handle(http.MethodHead, api.HEAD, root)
		api.handle(http.MethodConnect, api.CONNECT, root)
		api.handle(http.MethodTrace, api.TRACE, root)
	}
	for _, group := range g.subGroups {
		group.handle(root)
	}
}

func (g *group) addApi(a api) {
	g.apis = append(g.apis, a)
}

func (g *group) addGroup(group2 group) {
	g.subGroups = append(g.subGroups, group2)
}

func (g *group) Use(handlers ...gin.HandlerFunc) *group {
	g.middleware = append(g.middleware, handlers...)
	return g
}

func (g *group) Group(relativePath string, handlers ...gin.HandlerFunc) *group {
	group := group{
		path: relativePath,
	}
	group.Use(handlers...)
	return &group
}

func (a api) handle(httpMethod string, function gin.HandlerFunc, group *gin.RouterGroup) {
	if function != nil {
		group.Handle(httpMethod, a.path, function)
	}
}
