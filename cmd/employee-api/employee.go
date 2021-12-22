package main

import (
	"net/http"
	"quero2pay/pkg/business"
	"quero2pay/pkg/employee"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Response type represents a common API response.
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error error       `json:"error,omitempty"`
}

// Service type stores this API used services.
type Service struct {
	Employee employee.Service
	Business business.Service
}

// API represents this application resources.
type API struct {
	App     *fiber.App
	Service *Service
}

// Router defines the URL routing properties.
func (a *API) Router() {
	a.App.Get("/", Accepts)
	a.App.Post("/employee", a.create)
	a.App.Get("/employee", a.readAll)
	a.App.Get("/employee/:id", a.readOne)
	a.App.Put("/employee", a.update)
	a.App.Delete("/employee/:id", a.delete)
}

// add through API service request.
func (a *API) create(c *fiber.Ctx) error {
	data := &employee.Employee{}

	// Parses request body employee data
	if err := c.BodyParser(&data); err != nil {
		_log.Err(err).Msg("on employee.create body parser")
		return c.Status(http.StatusBadRequest).JSON(Response{
			Data:  data,
			Error: err,
		})
	}

	// Creates and saves this employee data
	if err := a.Service.Business.ReadOne(data.Business); err != nil {
		_log.Err(err).Send()
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}

	// Creates and saves this employee data
	if err := a.Service.Employee.Create(data); err != nil {
		_log.Err(err).Send()
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	LogEmployee(data, "created employee")

	return c.Status(http.StatusOK).JSON(
		Response{Data: data},
	)
}

// getAll through API client request.
func (a *API) readAll(c *fiber.Ctx) error {
	data := &[]employee.Employee{}

	// Get all saved employee
	if err := a.Service.Employee.ReadAll(data); err != nil {
		_log.Err(err).Msg("read all employee")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	_log.Info().Msg("read all employee")

	return c.Status(http.StatusOK).JSON(
		Response{Data: data},
	)
}

// getOneWithID through API client request.
func (a *API) readOne(c *fiber.Ctx) error {
	data := &employee.Employee{}
	data.ID, _ = strconv.Atoi(c.Params("id"))

	// Get this data with id
	if err := a.Service.Employee.ReadOne(data); err != nil {
		_log.Err(err).Msg("read one employee")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	LogEmployee(data, "read data")

	return c.Status(http.StatusOK).JSON(
		Response{Data: data},
	)
}

// update through API service request.
func (a *API) update(c *fiber.Ctx) error {
	data := &employee.Employee{}

	// Parses request body employee data
	if err := c.BodyParser(&data); err != nil {
		_log.Err(err).Msg("on employee.update body parser")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	find := *data

	// Get this data by its id
	if err := a.Service.Employee.ReadOne(&find); err != nil {
		_log.Err(err).Msg("read one employee")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}

	if find.Business.Name != data.Business.Name {
		// Creates and saves this employee data
		if err := a.Service.Business.ReadOne(data.Business); err != nil {
			_log.Err(err).Send()
			return c.Status(http.StatusUnprocessableEntity).JSON(Response{
				Data:  data,
				Error: err,
			})
		}
	}

	// Update saved data
	if err := a.Service.Employee.Update(data); err != nil {
		_log.Err(err).Msg("employee update")
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	LogEmployee(data, "updated employee")

	return c.Status(200).JSON(Response{
		Data:  data,
		Error: nil,
	})
}

// remove through API client request.
func (a *API) delete(c *fiber.Ctx) error {
	data := &employee.Employee{}
	data.ID, _ = strconv.Atoi(c.Params("id"))

	// Get this data by its id
	if err := a.Service.Employee.ReadOne(data); err != nil {
		_log.Err(err).Msg("read one employee")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}

	// Delete this data by its id
	if err := a.Service.Employee.Delete(data); err != nil {
		_log.Err(err).Msg("employee delete")
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	LogEmployee(data, "deleted employee")

	return c.Status(http.StatusOK).JSON(
		Response{Data: data},
	)
}
