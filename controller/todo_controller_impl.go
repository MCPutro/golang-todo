package controller

import (
	"fmt"
	"github.com/MCPutro/golang-todo/helper"
	"github.com/MCPutro/golang-todo/model"
	"github.com/MCPutro/golang-todo/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"strings"
)

type todoControllerImpl struct {
	service service.TodoService
}

func NewTodoController(service service.TodoService) TodoController {
	return &todoControllerImpl{service: service}
}

func (t *todoControllerImpl) GetAll(c *fiber.Ctx) error {
	filter1 := c.Query("activity_group_id", "")

	var todos []*model.Todo
	var err error

	if filter1 == "" {
		//call repo
		todos, err = t.service.FindAll(c.UserContext())

		//return negative
		if err != nil {
			return helper.PrintResponse(c, fiber.StatusInternalServerError, "Internal Server Error", "Failed to get users data", nil)
		}
	} else {
		actId, err := strconv.Atoi(filter1)
		if err != nil {
			return helper.PrintResponse(c, fiber.StatusBadRequest, "Invalid ID Parameter", fmt.Sprintf("ID %s not valid", filter1), nil)
		}

		//call repo
		todos, err = t.service.FindByActivityID(c.UserContext(), actId)

		//return negative
		if err != nil {
			if err.Error() == helper.NO_DATA_FOUND {
				return helper.PrintResponse(c, fiber.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %s Not Found", filter1), nil)
			} else {
				log.Println(err.Error())
				return helper.PrintResponse(c, fiber.StatusInternalServerError, "Internal Server Error", fmt.Sprintf("Failed get Activity with ID %s", actId), nil)
			}
		}
	}

	// return success
	return helper.PrintResponse(c, fiber.StatusOK, "Success", "Success", todos)
}

func (t *todoControllerImpl) GetOne(c *fiber.Ctx) error {
	//get id parameter
	todoId := c.Params("TodoId", "-1")
	todoIdInt, err := strconv.Atoi(todoId)
	if err != nil {
		return helper.PrintResponse(c, fiber.StatusBadRequest, "Invalid ID Parameter", fmt.Sprintf("ID %s not valid", todoId), nil)
	}

	todo, err := t.service.FindById(c.UserContext(), todoIdInt)

	//return resp
	if err != nil {
		if err.Error() == helper.NO_DATA_FOUND {
			return helper.PrintResponse(c, fiber.StatusNotFound, "Not Found", fmt.Sprintf("Todo with ID %s Not Found", todoId), nil)
		} else {
			log.Println(err.Error())
			return helper.PrintResponse(c, fiber.StatusInternalServerError, "Internal Server Error", fmt.Sprintf("Failed get Todo with ID %s", todoId), nil)
		}
	}
	return helper.PrintResponse(c, fiber.StatusOK, "Success", "Success", todo)
}

func (t *todoControllerImpl) Create(c *fiber.Ctx) error {
	//create temp variable
	body := new(model.Todo)

	//parse string json to struct
	if err := c.BodyParser(body); err != nil {
		return helper.PrintResponse(c, fiber.StatusBadRequest, "Bad Request", "Invalid request body", nil)
	}

	//call service
	create, err := t.service.Create(c.UserContext(), body)

	//return negative
	if err != nil {
		if err.Error() == helper.NO_DATA_FOUND {
			return helper.PrintResponse(c, fiber.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", body.Activity_group_id), nil)
		} else if strings.Contains(err.Error(), helper.NO_EXISTS) {
			return helper.PrintResponse(c, fiber.StatusNotFound, "Not Found", err.Error(), nil)
		} else {
			log.Println(err.Error())
			return helper.PrintResponse(c, fiber.StatusInternalServerError, "Internal Server Error", "Failed save new Todo", nil)
		}
	}

	//return positive
	return helper.PrintResponse(c, fiber.StatusOK, "Success", "Success", create)

}

func (t *todoControllerImpl) Update(c *fiber.Ctx) error {
	//get id parameter
	todoId := c.Params("TodoId", "-1")
	todoIdInt, err := strconv.Atoi(todoId)
	if err != nil {
		return helper.PrintResponse(c, fiber.StatusBadRequest, "Invalid ID Parameter", fmt.Sprintf("Todo ID %s not valid", todoId), nil)
	}

	var body = new(model.Todo)
	//parse string json to struct
	if err := c.BodyParser(body); err != nil {
		return helper.PrintResponse(c, fiber.StatusBadRequest, "Bad Request", "Invalid request body", nil)
	}
	body.Todo_id = todoIdInt

	update, err := t.service.Update(c.UserContext(), body)
	if err != nil {
		if err.Error() == helper.NO_DATA_FOUND {
			return helper.PrintResponse(c, fiber.StatusNotFound, "Not Found", fmt.Sprintf("Todo with ID %s Not Found", todoId), nil)
		} else {
			log.Println(err)
			return helper.PrintResponse(c, fiber.StatusInternalServerError, "Internal Server Error", fmt.Sprintf("Failed update todo with ID %s", todoId), nil)
			//return c.SendString("ada yg error : " + err.Error())
		}
	}

	//return positive
	return helper.PrintResponse(c, fiber.StatusOK, "Success", "Success", update)

}

func (t *todoControllerImpl) Delete(c *fiber.Ctx) error {
	//get id parameter
	todoId := c.Params("TodoId", "-1")
	todoIdInt, err := strconv.Atoi(todoId)
	if err != nil {
		return helper.PrintResponse(c, fiber.StatusBadRequest, "Invalid ID Parameter", fmt.Sprintf("Todo ID %s not valid", todoId), nil)
	}

	err = t.service.Delete(c.UserContext(), todoIdInt)
	if err != nil {
		if err.Error() == helper.NO_DATA_FOUND {
			return helper.PrintResponse(c, fiber.StatusNotFound, "Not Found", fmt.Sprintf("Todo with ID %s Not Found", todoId), nil)
		} else {
			log.Println(err.Error())
			return helper.PrintResponse(c, fiber.StatusInternalServerError, "Internal Server Error", fmt.Sprintf("Failed delete Todo with ID %s", todoId), nil)
		}
	}
	return helper.PrintResponse(c, fiber.StatusOK, "Success", fmt.Sprintf("Todo with ID %s has been deleted", todoId), nil)
}
