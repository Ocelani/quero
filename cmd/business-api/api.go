package main

import (
	"net/http"
	"quero2pay/internal/viacep"
	"quero2pay/pkg/address"
	"quero2pay/pkg/business"
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
	Business business.Service
	Address  address.Service
}

// API represents this application resources.
type API struct {
	App     *fiber.App
	Service *Service
}

// Router defines the URL routing properties.
func (a *API) Router() {
	a.App.Get("/", Accepts)
	a.App.Post("/business", a.create)
	a.App.Get("/business", a.readAll)
	a.App.Get("/business/:id", a.readOne)
	a.App.Put("/business", a.update)
	a.App.Delete("/business/:id", a.delete)
}

// add through API service request.
func (a *API) create(c *fiber.Ctx) error {
	data := &business.Business{}

	// Parses request body business data
	if err := c.BodyParser(&data); err != nil {
		_log.Err(err).Msg("on business.create body parser")
		return c.Status(http.StatusBadRequest).JSON(Response{
			Data:  data,
			Error: err,
		})
	}

	// Finds the full address based on the zip code
	viacep, err := viacep.NewAddressViaCEP(data.Address.ZipCode)
	if err != nil {
		_log.Err(err).Msg("on business address.NewAddressViaCEP")
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	data.Address = viacep.ConvertToAddress()

	// Creates and saves this address data
	if err := a.Service.Address.Create(data.Address); err != nil {
		_log.Err(err).Send()
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	LogAddress(data.Address, "created address")

	// Creates and saves this business data
	if err := a.Service.Business.Create(data); err != nil {
		_log.Err(err).Send()
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	LogBusiness(data, "created business")

	return c.Status(http.StatusOK).JSON(
		Response{Data: data},
	)
}

// getAll through API client request.
func (a *API) readAll(c *fiber.Ctx) error {
	data := &[]business.Business{}

	// Get all saved business
	if err := a.Service.Business.ReadAll(data); err != nil {
		_log.Err(err).Msg("read all business")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	_log.Info().Msg("read all business")

	return c.Status(http.StatusOK).JSON(
		Response{Data: data},
	)
}

// getOneWithID through API client request.
func (a *API) readOne(c *fiber.Ctx) error {
	data := &business.Business{}
	data.ID, _ = strconv.Atoi(c.Params("id"))

	// Get this data with id
	if err := a.Service.Business.ReadOne(data); err != nil {
		_log.Err(err).Msg("read one business")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	LogBusiness(data, "read data")

	return c.Status(http.StatusOK).JSON(
		Response{Data: data},
	)
}

// update through API service request.
func (a *API) update(c *fiber.Ctx) error {
	data := &business.Business{}

	// Parses request body business data
	if err := c.BodyParser(&data); err != nil {
		_log.Err(err).Msg("on business.update body parser")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	find := *data

	// Get this data by its id
	if err := a.Service.Business.ReadOne(&find); err != nil {
		_log.Err(err).Msg("read one business")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}

	// Update saved data
	if err := a.Service.Business.Update(data); err != nil {
		_log.Err(err).Msg("business update")
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	LogBusiness(data, "updated business")

	return c.Status(200).JSON(Response{
		Data:  data,
		Error: nil,
	})
}

// remove through API client request.
func (a *API) delete(c *fiber.Ctx) error {
	data := &business.Business{}
	data.ID, _ = strconv.Atoi(c.Params("id"))

	// Get this data by its id
	if err := a.Service.Business.ReadOne(data); err != nil {
		_log.Err(err).Msg("read one business")
		return c.Status(http.StatusUnprocessableEntity).JSON(Response{
			Data:  data,
			Error: err,
		})
	}

	// Delete this data by its id
	if err := a.Service.Business.Delete(data); err != nil {
		_log.Err(err).Msg("business delete")
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Data:  data,
			Error: err,
		})
	}
	LogBusiness(data, "deleted business")

	return c.Status(http.StatusOK).JSON(
		Response{Data: data},
	)
}
