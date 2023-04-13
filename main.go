package main

import (
	"github.com/MCPutro/golang-todo/controller"
	"github.com/MCPutro/golang-todo/database"
	"github.com/MCPutro/golang-todo/repository"
	"github.com/MCPutro/golang-todo/router"
	"github.com/MCPutro/golang-todo/service"
	"log"
)

func main() {
	db, err := database.GetInstance()
	if err != nil {
		log.Fatalln("")
		return
	}
	//prepare database table
	err = database.PrepareDB(db)
	if err != nil {
		log.Fatalln("error prepare DB, error : ", err)
		return
	}

	activityRepository := repository.NewActivityRepository()
	activityService := service.NewActivityService(activityRepository, db)
	activityController := controller.NewActivityController(activityService)

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, activityRepository, db)
	todoController := controller.NewTodoController(todoService)

	newRouter := router.NewRouter(activityController, todoController)

	PORT := "3030"
	log.Println("Running in port", PORT)

	err = newRouter.Listen(":" + PORT)
	if err != nil {
		log.Fatalln(err)
	}
}
