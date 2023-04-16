package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/mybookshelf/auth"
	"github.com/mujahxd/mybookshelf/helper"
)

type Handler interface {
	RegisterUser() echo.HandlerFunc
	Login() echo.HandlerFunc
	// DeleteActiveUser() echo.HandlerFunc
	// GetProfileUser() echo.HandlerFunc
	// UpdateProfileUser() echo.HandlerFunc
	// // GetAllUserBooks() echo.HandlerFunc
}

type handler struct {
	userService Service
	authService auth.Service
}

func NewHandler(userService Service, authService auth.Service) *handler {
	return &handler{userService, authService}
}

func (h *handler) RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input RegisterUserInput
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error(err.Error())
			response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", err.Error())
			return c.JSON(http.StatusUnprocessableEntity, response)
		}
		// validas inputan kosong
		validate := validator.New()
		err = validate.Struct(input)
		if err != nil {
			c.Logger().Error(err.Error())
			errors := helper.FormatValidationError(err)
			errorMessage := echo.Map{"errors": errors}
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", errorMessage)
			return c.JSON(http.StatusBadRequest, response)
		}

		newUser, err := h.userService.RegisterUser(input)
		if err != nil {
			c.Logger().Error(err.Error())
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
			return c.JSON(http.StatusBadRequest, response)

		}

		formatter := FormatRegisterUser(newUser)
		response := helper.APIResponse("Account has been registered", http.StatusCreated, "success", formatter)
		return c.JSON(http.StatusCreated, response)
	}
}

func (h *handler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginInput

		err := c.Bind(&input)
		if err != nil {
			response := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", err.Error())
			return c.JSON(http.StatusUnprocessableEntity, response)
		}
		// validas inputan kosong
		validate := validator.New()
		err = validate.Struct(input)
		if err != nil {
			c.Logger().Error(err.Error())
			errors := helper.FormatValidationError(err)
			errorMessage := echo.Map{"errors": errors}
			response := helper.APIResponse("login failed", http.StatusBadRequest, "error", errorMessage)
			return c.JSON(http.StatusBadRequest, response)
		}

		loggedinUser, err := h.userService.Login(input)
		if err != nil {
			errorMessage := echo.Map{"errors": err.Error()}

			response := helper.APIResponse("login failed", http.StatusBadRequest, "error", errorMessage)
			return c.JSON(http.StatusBadRequest, response)

		}
		token, err := h.authService.GenerateToken(loggedinUser.ID)
		if err != nil {
			response := helper.APIResponse("login failed", http.StatusBadRequest, "error", err.Error())
			return c.JSON(http.StatusBadRequest, response)

		}

		formatter := FormatLoginUser(loggedinUser, token)
		response := helper.APIResponse("successfully loggedin", http.StatusOK, "success", formatter)
		return c.JSON(http.StatusOK, response)
	}
}
