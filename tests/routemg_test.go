package ginroutemg_test

import (
	"testing"

	ginroutemg "github.com/crt379/gin-routemg"

	"github.com/gin-gonic/gin"
)

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
			GroupPath:  "/api/bs",
			Path:       "/api/bs",
			Handlers:   []gin.HandlerFunc{H},
			MethodFunc: ginroutemg.DefRouteMethod.GET,
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
			GroupPath:  "/api/ds",
			Path:       "/api/ds",
			Handlers:   []gin.HandlerFunc{H},
			MethodFunc: ginroutemg.DefRouteMethod.GET,
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
			GroupPath: "/api/es",
			Path:      "/api/es",
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
			GroupPath:  "/api/gs",
			Path:       "/api/gs",
			Handlers:   []gin.HandlerFunc{H},
			MethodFunc: ginroutemg.DefRouteMethod.GET,
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
				Middlewares: []gin.HandlerFunc{AM},
			},
			&ginroutemg.Route{
				GroupPath:  "/api/as",
				Path:       "/api/as/:id",
				Handlers:   []gin.HandlerFunc{AH},
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

func H(c *gin.Context) {}

func M(c *gin.Context) {}

func AH(c *gin.Context) {}

func AM(c *gin.Context) {}
