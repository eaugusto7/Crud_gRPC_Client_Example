package routes

import (
	"log"
	"net/http"

	"users/grpc/client/controllers"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/handlers"
)

func HandleRequest() {
	r := gin.Default()

	r.Static("/assets", "./assets")

	//CRUD Clientes
	r.POST("/api/v1/users", controllers.InsertUser)
	r.GET("/api/v1/users", controllers.GetAllUsers)
	r.GET("/api/v1/users/:id", controllers.GetUserById)
	r.PATCH("/api/v1/users/:id", controllers.UpdateUser)
	r.DELETE("/api/v1/users/:id", controllers.DeleteUser)

	log.Fatal(http.ListenAndServe(":8090", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
