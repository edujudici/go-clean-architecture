package route

import (
	"time"

	"github.com/edujudici/go-backend-clean-architecture/api/controller"
	"github.com/edujudici/go-backend-clean-architecture/bootstrap"
	"github.com/edujudici/go-backend-clean-architecture/domain"
	"github.com/edujudici/go-backend-clean-architecture/mongo"
	"github.com/edujudici/go-backend-clean-architecture/repository"
	"github.com/edujudici/go-backend-clean-architecture/usecase"
	"github.com/gin-gonic/gin"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
