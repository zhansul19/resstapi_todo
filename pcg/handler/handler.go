package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhansul19/restapi_todo/pcg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		services: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}
	api := router.Group("/api", h.UserIdentity)
	{
		lists := api.Group("lists")
		{
			lists.POST("/", h.CreateLists)
			lists.GET("/", h.GetLists)
			lists.GET("/:id", h.GetListsById)
			lists.DELETE("/:id", h.DeleteLists)
			lists.PUT("/:id", h.UpdateLists)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.CreateItems)
				items.GET("/", h.GetItems)
			}
		}
		items:=api.Group("items")
		{
			items.GET("/:id", h.GetItemsById)
			items.DELETE("/:id", h.DeleteItems)
			items.PUT("/:id", h.UpdateItems)
		}
	}
	return router
}
