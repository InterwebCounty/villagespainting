package routes

import (
	"github.com/interwebcounty/villagespainting/server/api/todo/routes"
	"github.com/interwebcounty/villagespainting/server/common/static"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	todoroutes.Init(router)
	static.Init(router)
}
