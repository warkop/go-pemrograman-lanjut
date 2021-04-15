package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

//User is
type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=80"`
}

//CustomValidator is
type CustomValidator struct {
	validator *validator.Validate
}

//Validate is
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	e.HTTPErrorHandler = func(err error, context echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		errPage := fmt.Sprintf("%d.html", report.Code)
		if err := context.File(errPage); err != nil {
			context.HTML(report.Code, "")
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
				}
				break
			}
		}

		context.Logger().Error(report)
		context.JSON(report.Code, report)
	}

	e.POST("/users", func(context echo.Context) error {
		u := new(User)
		if err := context.Bind(u); err != nil {
			return err
		}
		if err := context.Validate(u); err != nil {
			return err
		}

		return context.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
