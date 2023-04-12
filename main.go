package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MCPutro/golang-todo/config"
	"github.com/MCPutro/golang-todo/database"
	"github.com/MCPutro/golang-todo/model"
	"github.com/MCPutro/golang-todo/repository"
	"github.com/MCPutro/golang-todo/service"
)

func main() {
	ctx := context.Background()

	db, err := database.GetInstance()
	if err != nil {
		return
	}

	activityRepository := repository.NewActivityRepository()

	activityService := service.NewActivityService(activityRepository, db)

	//test dummy
	activities := model.Activities{
		Activity_id: 12,
		Title:       "makan - ah",
		Email:       "minum222",
	}

	update, err := activityService.Update(ctx, &activities)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(err)
	}

	//
	//create, err := activityService.Create(ctx, &activities)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//

	fmt.Println(update)
	marshal, err := json.Marshal(update)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(marshal))

}

func tes() {
	fmt.Println(config.DbName)
	db, err := database.GetInstance()
	if err != nil {
		return
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}

	activities := model.Activities{
		Title: "haha",
		Email: "hahahihi",
	}

	ctx := context.Background()

	activityRepository := repository.NewActivityRepository()
	save, err := activityRepository.Save(ctx, tx, &activities)

	all, err := activityRepository.FindAll(ctx, tx)

	id, err := activityRepository.FindByID(ctx, tx, 6)

	if err != nil {
		fmt.Println(err)
		fmt.Println("rollback")
		tx.Rollback()
	} else {
		fmt.Println("Commit")
		tx.Commit()
	}

	for i, activities := range all {
		fmt.Println(i, activities)
	}

	fmt.Println("------------")
	fmt.Println(id)
	fmt.Println("------------")
	jsonData, err := json.Marshal(save)

	fmt.Println(">>", string(jsonData))
}
