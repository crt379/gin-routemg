package ginroutemg_test

import (
	"testing"

	ginroutemg "github.com/crt379/gin-routemg"

	"github.com/gin-gonic/gin"
)

func H(c *gin.Context) {}

func M(c *gin.Context) {}

func H1(c *gin.Context) {}

func M1(c *gin.Context) {}

func Test(t *testing.T) {
	var rsinfo gin.RoutesInfo
	r := gin.Default()
	rmg := ginroutemg.NewRouteMG().SetRouter(r)

	rmg.AppendRoute(
		&ginroutemg.Route{
			Path:        "/api/as",
			Middlewares: []gin.HandlerFunc{M},
		},
	)
	rmg.RegisterRouter()
	rsinfo = r.Routes()
	if len(rsinfo) != 0 {
		t.Log(rsinfo)
		t.Fail()
	}

	rmg.AppendRoute(
		&ginroutemg.Route{
			GroupPath:   "/api/bs",
			Path:        "/api/bs",
			Middlewares: []gin.HandlerFunc{},
			Handlers:    []gin.HandlerFunc{H},
			MethodFunc:  ginroutemg.DefRouteMethod.GET,
		},
	)
	rmg.RegisterRouter()
	rsinfo = r.Routes()
	if len(rsinfo) != 1 {
		t.Log(rsinfo)
		t.Fail()
	}

	rmg.AppendRoute(
		&ginroutemg.Route{
			Path:        "/api/cs",
			Middlewares: []gin.HandlerFunc{M},
			Handlers:    []gin.HandlerFunc{H},
			MethodFunc:  ginroutemg.DefRouteMethod.GET,
		},
	)
	rmg.RegisterRouter()
	rsinfo = r.Routes()
	if len(rsinfo) != 2 {
		t.Log(rsinfo)
		t.Fail()
	}

	rmg.AppendRoute(
		&ginroutemg.Route{
			GroupPath:   "/api/ds",
			Path:        "/api/ds",
			Middlewares: []gin.HandlerFunc{},
			Handlers:    []gin.HandlerFunc{H},
			MethodFunc:  ginroutemg.DefRouteMethod.GET,
		},
	)
	rmg.RegisterRouter()
	rsinfo = r.Routes()
	if len(rsinfo) != 3 {
		t.Log(rsinfo)
		t.Fail()
	}

	rmg.AppendRoute(
		&ginroutemg.Route{
			GroupPath:  "/api/ds",
			Path:       "/api/ds/:id",
			Handlers:   []gin.HandlerFunc{H},
			MethodFunc: ginroutemg.DefRouteMethod.GET,
		},
	)
	rmg.RegisterRouter()
	rsinfo = r.Routes()
	if len(rsinfo) != 4 {
		t.Log(rsinfo)
		t.Fail()
	}

	rmg.AppendRoute(
		&ginroutemg.Route{
			GroupPath:  "/api/es",
			Path:       "/api/es/:id",
			Handlers:   []gin.HandlerFunc{H},
			MethodFunc: ginroutemg.DefRouteMethod.GET,
		},
		&ginroutemg.Route{
			GroupPath:   "/api/es",
			Path:        "/api/es",
			Middlewares: []gin.HandlerFunc{},
		},
		&ginroutemg.Route{
			GroupPath:  "/api/es",
			Path:       "/api/es",
			Handlers:   []gin.HandlerFunc{H},
			MethodFunc: ginroutemg.DefRouteMethod.GET,
		},
	)
	rmg.RegisterRouter()
	rsinfo = r.Routes()
	if len(rsinfo) != 6 {
		t.Log(rsinfo)
		t.Fail()
	}

	rmg.AppendRoute(
		&ginroutemg.Route{
			GroupPath:        "/api/fs",
			Path:             "/api/fs",
			Middlewares:      []gin.HandlerFunc{},
			Handlers:         []gin.HandlerFunc{H},
			MethodFunc:       ginroutemg.DefRouteMethod.GET,
			MulMethodHandler: &ginroutemg.MethodsHandlersBox{},
		},
	)
	rmg.RegisterRouter()
	rsinfo = r.Routes()
	if len(rsinfo) != 7 {
		t.Log(rsinfo)
		t.Fail()
	}

	rmg.AppendRoute(
		&ginroutemg.Route{
			GroupPath:   "/api/gs",
			Path:        "/api/gs",
			Middlewares: []gin.HandlerFunc{},
			Handlers:    []gin.HandlerFunc{H},
			MethodFunc:  ginroutemg.DefRouteMethod.GET,
			MulMethodHandler: &ginroutemg.MethodsHandlersBox{
				Methods: []ginroutemg.MethodFunc{
					ginroutemg.DefRouteMethod.POST,
					ginroutemg.DefRouteMethod.DELETE,
				},
				HandlersList: [][]gin.HandlerFunc{
					{H},
					{H},
				},
			},
		},
	)
	rmg.RegisterRouter()
	rsinfo = r.Routes()
	if len(rsinfo) != 10 {
		t.Log(rsinfo)
		t.Fail()
	}
}

func Test1(t *testing.T) {
	var rsinfo gin.RoutesInfo
	rmg := ginroutemg.NewRouteMG()

	func() {
		rmg.AppendRoute(
			&ginroutemg.Route{
				Path:        "/api/as",
				Middlewares: []gin.HandlerFunc{M1},
			},
			&ginroutemg.Route{
				GroupPath:  "/api/as",
				Path:       "/api/as/:id",
				Handlers:   []gin.HandlerFunc{H1},
				MethodFunc: ginroutemg.DefRouteMethod.POST,
			},
		)
	}()

	func() {
		rmg.AppendRoute(
			&ginroutemg.Route{
				GroupPath:  "/api/as",
				Path:       "/api/as/:id/bs",
				Handlers:   []gin.HandlerFunc{H},
				MethodFunc: ginroutemg.DefRouteMethod.POST,
			},
			&ginroutemg.Route{
				GroupPath:  "/api/as",
				Path:       "/api/as/:id/bs",
				Handlers:   []gin.HandlerFunc{H},
				MethodFunc: ginroutemg.DefRouteMethod.GET,
				MulMethodHandler: &ginroutemg.MethodsHandlersBox{
					Methods: []ginroutemg.MethodFunc{
						ginroutemg.DefRouteMethod.DELETE,
						ginroutemg.DefRouteMethod.PATCH,
					},
					HandlersList: [][]gin.HandlerFunc{
						{H},
						{H},
					},
				},
			},
		)
	}()

	r := gin.Default()
	rmg.SetRouter(r)
	rmg.RegisterRouter()
	rsinfo = r.Routes()
	if len(rsinfo) != 5 {
		t.Log(rsinfo)
		t.Fail()
	}
}

func Test2(t *testing.T) {
	var rsinfo gin.RoutesInfo
	rmg := ginroutemg.NewRouteMG()

	rmg.AppendRoute(
		&ginroutemg.Route{
			GroupPath:       "/api/applications",
			Path:            "/api/applications/:aid/services",
			Middlewares:     []gin.HandlerFunc{H},
			MiddlewaresFunc: ginroutemg.DefRouteMethod.Use,
		},
		&ginroutemg.Route{
			GroupPath: "/api/applications/:aid/services",
			Path:      "/api/applications/:aid/services",
			MulMethodHandler: &ginroutemg.MethodsHandlersBox{
				Methods: []ginroutemg.MethodFunc{
					ginroutemg.DefRouteMethod.GET,
					ginroutemg.DefRouteMethod.POST,
				},
				HandlersList: [][]gin.HandlerFunc{
					{H},
					{H1},
				},
			},
		},
	)

	r := gin.Default()
	rmg.SetRouter(r)
	rmg.RegisterRouter()
	rsinfo = r.Routes()
	if len(rsinfo) != 0 {
		t.Log(rsinfo)
		t.Fail()
	}
}
