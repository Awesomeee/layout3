package layout

import (
	"github.com/hpt/go-client/app/middleware"
)

func Routes(routerGroup middleware.RouterGroup) {
	checkDuplicateLayout := routerGroup.RouterDefault.Group("/layout")
	checkDuplicateLayout.POST("/checkDuplicateLayoutID", checkDuplicateLayoutID)
}
