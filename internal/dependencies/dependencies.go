package dependencies

import (
	"github.com/thiagocprado/golang-api-structure/internal/database"
	"github.com/thiagocprado/golang-api-structure/internal/domain/user"
	"github.com/thiagocprado/golang-api-structure/pkg/router"
)

func LoadDependencies(router *router.Router) {
	loadUser(router)
}

func loadUser(router *router.Router) {
	r := user.NewRepository(database.MySqlDB)
	u := user.NewUseCase(r)
	h := user.NewHandler(u)
	router.InjectRoutes(user.GetUserRoutes(h))
}
