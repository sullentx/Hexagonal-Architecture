package routes

import (
	controller "tienda/src/Notification/Infraestructure/Controller"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, postNotification *controller.PostNotificationHandler,
	deleteNotification *controller.DeleteNotificationHandler,
	searchNotification *controller.SearchNotificationHandler,
	modifyNotification *controller.ModifyNotificationHandler,
	getAllNotification *controller.GetAllNotificationHandler, getAllNotificationShort *controller.GetNotificationsHandlerShort) {
	router.POST("/notifications", postNotification.HandlePost)
	router.DELETE("/notifications/:id", deleteNotification.HandleDelete)
	router.GET("/notifications/:id", searchNotification.HandleSearch)
	router.PUT("/notifications/:id", modifyNotification.HandleModify)
	router.GET("/notifications", getAllNotification.HandleGetAll)
	router.GET("/notifications/short-polling/:id", getAllNotificationShort.HandleGet)
}
