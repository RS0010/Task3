package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type (
	routerGroup struct {
		maker    func(group *gin.RouterGroup, handlers ...gin.HandlerFunc) *gin.RouterGroup
		handlers gin.HandlersChain
		routers  []*router
		groups   []*routerGroup
	}

	router struct {
		maker    func(group *gin.RouterGroup, handlers ...gin.HandlerFunc)
		handlers gin.HandlersChain
	}
)

var Version = make([]routerGroup, apiVersionNumber)

func Init(router *gin.Engine) {
	for i := 0; i < apiVersionNumber; i++ {
		path := "/v" + strconv.Itoa(i+1)
		Version[i].maker = func(group *gin.RouterGroup, handlers ...gin.HandlerFunc) *gin.RouterGroup {
			return group.Group(path, handlers...)
		}
	}
	for _, group := range Version {
		group.builder(&router.RouterGroup)
	}
}

func (r *router) Use(handlers ...gin.HandlerFunc) *router {
	r.handlers = append(r.handlers, handlers...)
	return r
}

func (r routerGroup) builder(group *gin.RouterGroup) {
	root := r.maker(group, r.handlers...)
	for _, router := range r.routers {
		router.maker(root, router.handlers...)
	}
	for _, group := range r.groups {
		group.builder(root)
	}
}

func (r *routerGroup) Group(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	var group routerGroup
	group.maker = func(group *gin.RouterGroup, handlers ...gin.HandlerFunc) *gin.RouterGroup {
		return group.Group(relativePath, handlers...)
	}
	group.handlers = handlers
	r.groups = append(r.groups, &group)
	return &group
}

func (r *routerGroup) Use(handlers ...gin.HandlerFunc) *routerGroup {
	r.handlers = append(r.handlers, handlers...)
	return r
}

func (r *routerGroup) handle(httpMethod string, handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *router {
	var root router
	root.maker = func(group *gin.RouterGroup, handlers ...gin.HandlerFunc) {
		group.Handle(httpMethod, "", handler).Use(handlers...)
	}
	root.handlers = middlewares
	r.routers = append(r.routers, &root)
	return &root
}

func (r *routerGroup) POST(handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *router {
	return r.handle(http.MethodPost, handler, middlewares...)
}

func (r *routerGroup) GET(handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *router {
	return r.handle(http.MethodGet, handler, middlewares...)
}

func (r *routerGroup) DELETE(handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *router {
	return r.handle(http.MethodDelete, handler, middlewares...)
}

func (r *routerGroup) PATCH(handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *router {
	return r.handle(http.MethodPatch, handler, middlewares...)
}

func (r *routerGroup) PUT(handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *router {
	return r.handle(http.MethodPut, handler, middlewares...)
}

func (r *routerGroup) OPTIONS(handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *router {
	return r.handle(http.MethodOptions, handler, middlewares...)
}

func (r *routerGroup) HEAD(handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *router {
	return r.handle(http.MethodHead, handler, middlewares...)
}
