package Utilities

import (
	lokihook "github.com/akkuman/logrus-loki-hook"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

type BaseResponse struct {
	Ctx       *fiber.Ctx
	Err       error
	Data      interface{}
	Status    int
	Count     int
	Page      int
	ExtraData fiber.Map
	Log       *logrus.Logger
}

func (r *BaseResponse) LogManagement() {
	// Create a new Loki client
	log := logrus.New()

	lokiHookConfig := &lokihook.Config{
		URL:       "http://admin:114136Iqbal@52.74.15.85:3100/api/prom/push",
		LevelName: "severity",
		Labels: map[string]string{
			"application": "test",
		},
	}
	hook, err := lokihook.NewHook(lokiHookConfig)

	if err != nil {
		log.Error(err)
	} else {
		log.AddHook(hook)
		r.Log = log
	}
	log.Info("test")
}

func (r *BaseResponse) NotFound() error {
	return r.Ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error":       r.Err.Error(),
		"data":        nil,
		"status":      fiber.StatusNotFound,
		"server_time": time.Now().Unix(),
		"meta": fiber.Map{
			"count":      0,
			"page":       0,
			"extra_data": r.ExtraData,
		},
	})
}

func (r *BaseResponse) BadRequest() error {
	return r.Ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":       r.Err.Error(),
		"data":        nil,
		"status":      fiber.StatusBadRequest,
		"server_time": time.Now().Unix(),
		"meta": fiber.Map{
			"count":      0,
			"page":       0,
			"extra_data": r.ExtraData,
		},
	})
}

func (r *BaseResponse) OK() error {
	return r.Ctx.Status(r.Status).JSON(fiber.Map{
		"error":       nil,
		"data":        r.Data,
		"status":      r.Status,
		"server_time": time.Now().Unix(),
		"meta": fiber.Map{
			"count":      r.Count,
			"page":       r.Page,
			"extra_data": r.ExtraData,
		},
	})
}

func NewResponse(config BaseResponse) *BaseResponse {
	response := &config
	response.LogManagement()

	return response
}
