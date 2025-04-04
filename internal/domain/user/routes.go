package user

import (
	"net/http"

	"github.com/thiagocprado/golang-api-structure/pkg/router"
)

func GetUserRoutes(controller *Handler) []router.Route {
	apiBaseUrlv1 := "/api/v1/users"

	return []router.Route{
		{
			URI:                   apiBaseUrlv1,
			Method:                http.MethodPost,
			Function:              controller.Create,
			RequestAuthentication: true,
		},
		{
			URI:                   apiBaseUrlv1,
			Method:                http.MethodGet,
			Function:              controller.FindAll,
			RequestAuthentication: true,
		},
	}
}
