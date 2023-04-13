package controller

import (
	"fmt"
	"github.com/MCPutro/golang-todo/helper"
	"github.com/MCPutro/golang-todo/model"
	"github.com/MCPutro/golang-todo/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type activityControllerImpl struct {
	service service.ActivityService
}

func NewActivityController(service service.ActivityService) ActivityController {
	return &activityControllerImpl{service: service}
}

func (a *activityControllerImpl) GetAll(c *fiber.Ctx) error {
	activities, err := a.service.FindAll(c.UserContext())
	if err != nil {
		return helper.PrintResponse(c, fiber.StatusInternalServerError, "Internal Server Error", "Failed to get Activities", nil)
	}
	return helper.PrintResponse(c, fiber.StatusOK, "Success", "Success", activities)
}

func (a *activityControllerImpl) GetOne(c *fiber.Ctx) error {
	//get id parameter
	actId := c.Params("ActId", "-1")
	actIdInt, err := strconv.Atoi(actId)
	if err != nil {
		return helper.PrintResponse(c, fiber.StatusBadRequest, "Invalid ID Parameter", fmt.Sprintf("ID %s not valid", actId), nil)
	}

	//call service
	activity, err := a.service.FindById(c.UserContext(), actIdInt)

	//return resp
	if err != nil {
		if err.Error() == helper.NO_DATA_FOUND {
			return helper.PrintResponse(c, fiber.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %s Not Found", actId), nil)
		} else {
			log.Println(err.Error())
			return helper.PrintResponse(c, fiber.StatusInternalServerError, "Internal Server Error", fmt.Sprintf("Failed get Activity with ID %s", actId), nil)
		}
	}
	return helper.PrintResponse(c, fiber.StatusOK, "Success", "Success", activity)

}

func (a *activityControllerImpl) Create(c *fiber.Ctx) error {

	//create temp variable
	body := new(model.Activity)

	//parse string json to struct
	if err := c.BodyParser(body); err != nil {
		return helper.PrintResponse(c, fiber.StatusBadRequest, "Bad Request", "Invalid request body", nil)
	}

	//call service
	activity, err := a.service.Create(c.UserContext(), body)

	//return response
	if err != nil {
		log.Println(err.Error())
		return helper.PrintResponse(c, fiber.StatusInternalServerError, "Internal Server Error", "Failed save new activity", nil)
	}
	return helper.PrintResponse(c, fiber.StatusCreated, "Success", "Success", activity)

}

func (a *activityControllerImpl) Update(c *fiber.Ctx) error {

	//get parameter
	actId := c.Params("ActId", "-1")
	actIdInt, err := strconv.Atoi(actId)
	if err != nil {
		return helper.PrintResponse(c, fiber.StatusBadRequest, "Invalid ID Parameter", fmt.Sprintf("ID %s not valid", actId), nil)
	}

	//get json req
	//create temp variable
	body := new(model.Activity)
	//parse string json to var body
	if err := c.BodyParser(body); err != nil {
		return helper.PrintResponse(c, fiber.StatusBadRequest, "Bad Request", "Invalid request body", nil)
	}
	body.Activity_id = actIdInt

	//call service
	update, err := a.service.Update(c.UserContext(), body)
	if err != nil {
		if err.Error() == helper.NO_DATA_FOUND {
			return helper.PrintResponse(c, fiber.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %s Not Found", actId), nil)
		} else {
			log.Println(err.Error())
			return helper.PrintResponse(c, fiber.StatusInternalServerError, "Internal Server Error", "Failed update activity", nil)
		}
	}
	return helper.PrintResponse(c, fiber.StatusOK, "Success", "Success", update)

}

func (a *activityControllerImpl) Delete(c *fiber.Ctx) error {
	//get id parameter
	actId := c.Params("ActId", "-1")
	actIdInt, err := strconv.Atoi(actId)
	if err != nil {
		return helper.PrintResponse(c, fiber.StatusBadRequest, "Invalid ID Parameter", fmt.Sprintf("ID %s not valid", actId), nil)
	}

	//call service
	err = a.service.Delete(c.UserContext(), actIdInt)
	if err != nil {
		if err.Error() == helper.NO_DATA_FOUND {
			return helper.PrintResponse(c, fiber.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %s Not Found", actId), nil)
		} else {
			log.Println(err.Error())
			return helper.PrintResponse(c, fiber.StatusInternalServerError, "Internal Server Error", fmt.Sprintf("Failed delete Activity with ID %s", actId), nil)
		}
	}
	return helper.PrintResponse(c, fiber.StatusOK, "Success", fmt.Sprintf("Activity with ID %s has been deleted", actId), nil)
}
