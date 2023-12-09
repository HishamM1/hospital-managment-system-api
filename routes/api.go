package routes

import (
	"main/app/controllers"

	"github.com/gin-gonic/gin"
)

func Api() {
	route := gin.Default()

	// patients
	route.GET("/patients", controllers.GetPatients)
	route.GET("/patients/:id", controllers.GetPatient)
	route.POST("/patients", controllers.CreatePatient)
	route.PUT("/patients/:id", controllers.UpdatePatient)
	route.DELETE("/patients/:id", controllers.DeletePatient)

	// doctors
	route.GET("/doctors", controllers.GetDoctors)
	route.GET("/doctors/:id", controllers.GetDoctor)
	route.POST("/doctors", controllers.CreateDoctor)
	route.PUT("/doctors/:id", controllers.UpdateDoctor)
	route.DELETE("/doctors/:id", controllers.DeleteDoctor)

	// nurses
	route.GET("/nurses", controllers.GetNurses)
	route.GET("/nurses/:id", controllers.GetNurse)
	route.POST("/nurses", controllers.CreateNurse)
	route.PUT("/nurses/:id", controllers.UpdateNurse)
	route.DELETE("/nurses/:id", controllers.DeleteNurse)

	// rooms
	route.GET("/rooms", controllers.GetRooms)
	route.GET("/rooms/:id", controllers.GetRoom)
	route.POST("/rooms", controllers.CreateRoom)
	route.PUT("/rooms/:id", controllers.UpdateRoom)
	route.DELETE("/rooms/:id", controllers.DeleteRoom)

	// wardboys
	route.GET("/wardboys", controllers.GetWardboys)
	route.GET("/wardboys/:id", controllers.GetWardboy)
	route.POST("/wardboys", controllers.CreateWardboy)
	route.PUT("/wardboys/:id", controllers.UpdateWardboy)
	route.DELETE("/wardboys/:id", controllers.DeleteWardboy)

	// treatments
	route.GET("/treatments", controllers.GetTreatments)
	route.GET("/treatments/:id", controllers.GetTreatment)
	route.POST("/treatments", controllers.CreateTreatment)
	route.PUT("/treatments/:id", controllers.UpdateTreatment)
	route.DELETE("/treatments/:id", controllers.DeleteTreatment)

	// numbers
	route.GET("/numbers", controllers.GetNumbers)
	route.GET("/numbers/:id", controllers.GetNumber)
	route.POST("/numbers", controllers.CreateNumber)
	route.PUT("/numbers/:id", controllers.UpdateNumber)
	route.DELETE("/numbers/:id", controllers.DeleteNumber)

	// bills
	route.GET("/bills", controllers.GetBills)
	route.GET("/bills/:id", controllers.GetBill)
	route.POST("/bills", controllers.CreateBill)
	route.PUT("/bills/:id", controllers.UpdateBill)
	route.DELETE("/bills/:id", controllers.DeleteBill)

	route.Run("127.0.0.1:8080")
}
