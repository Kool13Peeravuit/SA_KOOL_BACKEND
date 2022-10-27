package main

import (
	"github.com/Kool13Peeravuit/sa-project/controller"
	"github.com/Kool13Peeravuit/sa-project/entity"
	"github.com/Kool13Peeravuit/sa-project/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// User Routes
			router.GET("/users", controller.ListUser)
			router.GET("/users/:id", controller.GetUser)
			router.PATCH("/users", controller.UpdateUser)
			router.DELETE("/users/:id", controller.DeleteUser)

			router.GET("/computers", controller.ListComputers)
			router.GET("/computers/:id", controller.GetComputer)
			router.POST("/computers", controller.CreateComputer)
			router.PATCH("/computers", controller.UpdateComputer)
			router.DELETE("/computers/:id", controller.DeleteComputer)

			router.GET("/place_classes", controller.ListPlace_Class)
			router.GET("/place_classes/:id", controller.GetPlace_Class)
			router.POST("/place_classes", controller.CreatePlace_Class)
			router.PATCH("/place_classes", controller.UpdatePlace_Class)
			router.DELETE("/place_classes/:id", controller.DeletePlace_Class)

			router.GET("/problems", controller.ListProblems)
			router.GET("/problems/:id", controller.GetProblem)
			router.POST("/problems", controller.CreateProblem)
			router.PATCH("/problems", controller.UpdateProblem)
			router.DELETE("/problems/:id", controller.DeleteProblem)

			router.GET("/problemreports", controller.ListProblemReport)
			router.GET("/problemreports/:id", controller.GetProblemReport)
			router.POST("/problemreports", controller.CreateProblemReport)
			router.PATCH("/problemreports", controller.UpdateProblemReport)
			router.DELETE("/problemreports/:id", controller.DeleteProblemReport)

			router.GET("/readingzones", controller.ListReadingZone)
			router.GET("/readingzones/:id", controller.GetReadingZone)
			router.POST("/readingzones", controller.CreateReadingZone)
			router.PATCH("/readingzones", controller.UpdateReadingZone)
			router.DELETE("/readingzones/:id", controller.DeleteReadingZone)

			// router.GET("/realtions", controller.ListRelations)
			// router.GET("/realtions/:id", controller.GetRelation)
			// router.POST("/realtions", controller.CreateRelation)
			// router.PATCH("/realtions", controller.UpdateRelation)
			// router.DELETE("/realtions/:id", controller.DeleteRelation)

			// Research_Room Routes
			router.GET("/researchrooms", controller.ListResearchRooms)
			router.GET("/researchroom/:id", controller.GetResearchRoom)
			router.POST("/researchrooms", controller.CreateResearchRoom)
			router.PATCH("/researchrooms", controller.UpdateResearchRoom)
			router.DELETE("/researchrooms/:id", controller.DeleteResearchRoom)

			router.GET("/roomtypes", controller.ListRoomTypes)
			router.GET("/roomtype/:id", controller.GetRoomType)
			router.POST("/roomtypes", controller.CreateRoomType)
			router.PATCH("/roomtypes", controller.UpdateRoomType)
			router.DELETE("/roomtypes/:id", controller.DeleteRoomType)

			router.GET("/toilets", controller.ListToilet)
			router.GET("/toilets/:id", controller.GetToilet)
			router.POST("/toilets", controller.CreateToilet)
			router.PATCH("/toilets", controller.UpdateToilet)
			router.DELETE("/toilets/:id", controller.DeleteToilet)

		}
	}
	// User Routes
	r.POST("/users", controller.CreateUser)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
