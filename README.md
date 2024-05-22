# gin-routemg

方便全局注入 route 到 gin 中

## 使用示例

```go
// routemg.go
import (
	ginroutemg "github.com/crt379/gin-routemg"

	"github.com/gin-gonic/gin"
)
rmg := ginroutemg.NewRouteMG()

// apia.go
func init() {
    rmg.AppendRoute(
        &ginroutemg.Route{
            Path:        "/api/as",
            Middlewares: []gin.HandlerFunc{AM},
        },
        &ginroutemg.Route{
            GroupPath:  "/api/as",
            Path:        "/api/as/:id",
            Handlers:   []gin.HandlerFunc{AH},
            MethodFunc: ginroutemg.DefRouteMethod.POST,
        },
    )
}

func AH(c *gin.Context) {}

func AM(c *gin.Context) {}

// apib.go
func init() {
    rmg.AppendRoute(
        &ginroutemg.Route{
            GroupPath:  "/api/as", // 会获取到 apia 中的 group, 即会走到 AM
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
}

// main.go
r := gin.Default()
rmg.SetRouter(r)
rmg.RegisterRouter()
```