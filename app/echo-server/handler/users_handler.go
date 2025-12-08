package handler

import (
	"log"
	"net/http"
	"socialmedia/domain"
	"socialmedia/service/users_service"
	"strings"

	"github.com/AMFarhan21/fres"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	UsersHandler struct {
		usersService users_service.Service
		validate     *validator.Validate
	}

	RegisterInput struct {
		FirstName string `json:"first_name" bson:"first_name" validate:"required,min=4"`
		LastName  string `json:"last_name" bson:"last_name"`
		Address   string `json:"address" bson:"address" validate:"required,min=4"`
		Email     string `json:"email" bson:"email" validate:"required,email"`
		Username  string `json:"username" bson:"username" validate:"required,min=4"`
		Password  string `json:"password" bson:"password" validate:"required,min=4"`
		Age       int    `json:"age" bson:"age" validate:"required,min=12"`
	}

	LoginInput struct {
		Email    string `json:"email" bson:"email" validate:"required"`
		Password string `json:"password" bson:"password" validate:"required"`
	}

	GetAllInput struct {
		Page  int `json:"page" bson:"page"`
		Limit int `json:"limit" bson:"limit"`
	}
)

func NewUsersHandler(usersService users_service.Service) *UsersHandler {
	return &UsersHandler{
		usersService: usersService,
		validate:     validator.New(),
	}
}

func (h UsersHandler) RegisterNewUser(e echo.Context) error {
	var request RegisterInput

	if err := e.Bind(&request); err != nil {
		log.Printf("Error on RegisterNewUser request body %s", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Print("Error on RegisterNewUser validation", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	user, err := h.usersService.Register(domain.Users{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Address:   request.Address,
		Email:     request.Email,
		Username:  request.Username,
		Password:  request.Password,
		Age:       request.Age,
	})
	if err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			log.Printf("Error on RegisterNewUser request: %s", err.Error())
			return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
		}
		log.Printf("Error on RegisterNewUser internal: %s", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusCreated, fres.Response.StatusCreated(user))
}

func (h UsersHandler) LoginUser(e echo.Context) error {
	var request LoginInput

	if err := e.Bind(&request); err != nil {
		log.Printf("Error on LoginUser request body %s", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Print("Error on LoginUser validation", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	signedToken, err := h.usersService.Login(request.Email, request.Password)
	if err != nil {
		log.Printf("Error on LoginUser internal: %s", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK(signedToken))
}

func (h UsersHandler) GetAllUsers(e echo.Context) error {
	var request GetAllInput

	if err := e.Bind(&request); err != nil {
		log.Printf("Error on GetAllUsers request body: %s", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	users, err := h.usersService.GetAllUsers(request.Page, request.Limit)
	if err != nil {
		log.Printf("Error on GetAllUsers internal: %s", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK(users))
}

func (h UsersHandler) GetMe(e echo.Context) error {
	user_id := e.Get("id").(string)

	user, err := h.usersService.GetUserByID(user_id)
	if err != nil {
		log.Printf("Error on GetUserByID internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK(user))
}

func (h UsersHandler) GetUserByID(e echo.Context) error {
	user_id := e.Param("id")

	user, err := h.usersService.GetUserByID(user_id)
	if err != nil {
		log.Printf("Error on GetUserByID internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK(user))
}
