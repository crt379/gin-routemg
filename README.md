# gin-routemg

gin-routemg 是方便注入 route 到 gin 中一个管理包

## Getting started

### Getting gin-routemg

- import
```
import "github.com/crt379/gin-routemg"
```

- get
```
go get -u github.com/crt379/gin-routemg
```

### Use gin-routemg

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

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)