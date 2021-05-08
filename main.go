package main

import (
	"fmt"
	"golab20errors/router"
	"net/http"
	"os"
)

var (
	PORT       string        = os.Getenv("APP_PORT")
	httpRouter router.Router = router.NewChiRouter()
)

func main() {
	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "hw..")
	})
	httpRouter.SERVE(PORT)
}
