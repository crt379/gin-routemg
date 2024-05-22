package ginroutemg

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var DefRouteMethod RouteMethod

type RouteMG struct {
	rs     []*Route
	unrs   []*Route
	pir    map[string]gin.IRouter
	count  int32
	once   sync.Once
	lck    sync.Mutex
	router *gin.Engine
}

func NewRouteMG() *RouteMG {
	return &RouteMG{}
}

func (m *RouteMG) AppendRoute(rs ...*Route) {
	m.lck.Lock()
	defer m.lck.Unlock()

	m.rs = append(m.rs, rs...)
}

func (m *RouteMG) SetRouter(router *gin.Engine) *RouteMG {
	m.once.Do(func() {
		m.setrouter(router)
	})

	return m
}

func (m *RouteMG) RegisterRouter() {
	m.lck.Lock()
	defer m.lck.Unlock()

	m.register(m.rs)
	m.rs = m.unrs
	m.unrs = nil
	m.register(m.rs)

	m.rs = nil
}

func (m *RouteMG) setrouter(router *gin.Engine) {
	m.router = router
	m.pir = map[string]gin.IRouter{"": router}
}

func (m *RouteMG) register(routs []*Route) {
	if len(routs) == 0 {
		return
	}

	pir := m.pir
	unrs := routs
	unrslen := len(unrs)
	for len(unrs) > 0 {
		_unrs := make([]*Route, 0)
		for _, route := range unrs {
			debug(route)
			group, ok := pir[route.GroupPath]
			if !ok && !(route.GroupPath != "" && route.Middlewares != nil && route.GroupPath == route.Path) {
				debug("!ok")
				_unrs = append(_unrs, route)
				continue
			}

			m.count += 1
			if !ok { //route.GroupPath != "" && route.GroupPath == route.Path
				debug("!ok group")
				group = m.router.Group(route.GroupPath)
				pir[route.GroupPath] = group
			}

			rpath := route.GetRelativePath()
			path := route.GetPath()

			if route.Middlewares != nil {
				debug("route.Middlewares != nil")
				ir, ok := pir[path]
				if !ok {
					ir = group.Group(rpath)
					route.GroupPath = path
					pir[path] = ir
					group = ir
				}
				if route.MiddlewaresFunc == nil {
					route.MiddlewaresFunc = DefRouteMethod.Use
				}
				route.MiddlewaresFunc(ir.(*gin.RouterGroup), route.Middlewares...)
			}

			rpath = route.GetRelativePath()

			if route.MethodFunc != nil {
				debug("route.MethodFunc != nil")
				route.MethodFunc(group, rpath, route.Handlers...)
			}

			if route.MulMethodHandler != nil {
				debug("route.MulMethodHandler != nil")
				mmh := route.MulMethodHandler
				if len(mmh.Methods) != len(mmh.HandlersList) {
					panic(fmt.Sprintf("MulMethodHandler.HandlersList len(%d) != MulMethodHandler.Methods len(%d) 数量对不上",
						len(mmh.HandlersList), len(mmh.Methods)))
				}
				for i, method := range mmh.Methods {
					method(group, rpath, mmh.HandlersList[i]...)
				}
			}
		}
		if unrslen == len(_unrs) {
			debug("unrslen == len(_unrs)", len(_unrs))
			break
		}
		unrs = _unrs
		unrslen = len(unrs)
	}

	m.unrs = append(m.unrs, unrs...)
	debug("unrslen == len(_unrs)", len(m.unrs))
}

type Route struct {
	GroupPath        string
	RelativePath     string
	Path             string
	Middlewares      []gin.HandlerFunc   `json:"-"`
	Handlers         []gin.HandlerFunc   `json:"-"`
	MiddlewaresFunc  MiddlewaresFunc     `json:"-"`
	MethodFunc       MethodFunc          `json:"-"`
	MulMethodHandler *MethodsHandlersBox `json:"-"`
}

type MethodsHandlersBox struct {
	Methods      []MethodFunc
	HandlersList [][]gin.HandlerFunc
}

type MethodFunc func(i gin.IRouter, p string, hs ...gin.HandlerFunc)
type MiddlewaresFunc func(i gin.IRouter, hs ...gin.HandlerFunc)

func (r *Route) GetRelativePath() string {
	if r.RelativePath == "" {
		if r.GroupPath == "" {
			return r.Path
		} else if r.Path != "" && r.GroupPath == r.Path { // r.GroupPath != "" && r.Path != "" && r.GroupPath == r.Path
			return r.RelativePath
		} else if r.Path != "" && r.GroupPath != r.Path { // r.GroupPath != "" && r.Path != "" && r.GroupPath != r.Path
			if len(r.Path) > len(r.GroupPath) && r.GroupPath == r.Path[:len(r.GroupPath)] {
				return r.Path[len(r.GroupPath):]
			}
		}
	}

	return r.RelativePath
}

func (r *Route) GetPath() string {
	if r.Path == "" {
		return r.GroupPath + r.RelativePath
	}
	return r.Path
}

type RouteMethod struct{}

func (r *RouteMethod) Use(i gin.IRouter, hs ...gin.HandlerFunc) {
	i.Use(hs...)
}

func (r *RouteMethod) Any(i gin.IRouter, p string, hs ...gin.HandlerFunc) {
	i.Any(p, hs...)
}

func (r *RouteMethod) GET(i gin.IRouter, p string, hs ...gin.HandlerFunc) {
	i.GET(p, hs...)
}

func (r *RouteMethod) POST(i gin.IRouter, p string, hs ...gin.HandlerFunc) {
	i.POST(p, hs...)
}

func (r *RouteMethod) DELETE(i gin.IRouter, p string, hs ...gin.HandlerFunc) {
	i.DELETE(p, hs...)
}

func (r *RouteMethod) PATCH(i gin.IRouter, p string, hs ...gin.HandlerFunc) {
	i.PATCH(p, hs...)
}

func (r *RouteMethod) PUT(i gin.IRouter, p string, hs ...gin.HandlerFunc) {
	i.PUT(p, hs...)
}

func (r *RouteMethod) OPTIONS(i gin.IRouter, p string, hs ...gin.HandlerFunc) {
	i.OPTIONS(p, hs...)
}

func (r *RouteMethod) HEAD(i gin.IRouter, p string, hs ...gin.HandlerFunc) {
	i.HEAD(p, hs...)
}
