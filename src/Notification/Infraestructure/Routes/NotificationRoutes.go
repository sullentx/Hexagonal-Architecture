package routes

import (
	controller "tienda/src/Notification/Infraestructure/Controller"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, postNotification *controller.PostNotificationHandler, deleteNotification *controller.DeleteNotificationHandler, searchNotification *controller.SearchNotificationHandler, modifyNotification *controller.ModifyNotificationHandler, getAllNotification *controller.GetAllNotificationHandler) {
	router.POST("/notification", postNotification.HandlePost)
	router.DELETE("/notification/:id", deleteNotification.HandleDelete)
	router.GET("/notification/:id", searchNotification.HandleSearch)
	router.PUT("/notification/:id", modifyNotification.HandleModify)
	router.GET("/notification", getAllNotification.HandleGetAll)
}
