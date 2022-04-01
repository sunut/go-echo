package main

import (
	config "fake.com/develop/config"
	errorHandler "fake.com/develop/error"
	model "fake.com/develop/models"
	"fake.com/develop/repo"
	"fake.com/develop/util"
	//error "fake.com/develop/error"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"net/http"
)

func initConfig() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configuration.Database.DBName)
	fmt.Println("Port is\t\t", configuration.Server.Port)
	fmt.Println("EXAMPLE_PATH is\t", configuration.EXAMPLE_PATH)
	fmt.Println("EXAMPLE_VAR is\t", configuration.EXAMPLE_VAR)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model..")
	fmt.Println("Database is\t", viper.GetString("database.dbname"))
	fmt.Println("Port is\t\t", viper.GetInt("server.port"))
	fmt.Println("EXAMPLE_PATH is\t", viper.GetString("EXAMPLE_PATH"))
	fmt.Println("EXAMPLE_VAR is\t", viper.GetString("EXAMPLE_VAR"))
}

func handleRequests() {
	initConfig()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//custom Error handler
	e.HTTPErrorHandler = errorHandler.HttpErrorHandler

	// Routes
	e.GET("/", hello)

	e.GET("/users/finduser", func(c echo.Context) error {
		res, err := repo.GetError()
		if err != nil {
			fmt.Println("has error")
			panic(err)
		}
		return c.JSON(http.StatusOK, res)
		//return c.String(http.StatusOK, "/users/:id")
	})

	e.GET("/users", func(c echo.Context) error {
		res, err := repo.GetAll()
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, res)
		//return c.String(http.StatusOK, "/users/:id")
	})

	e.GET("/errors", func(c echo.Context) error {
		return errorHandler.HTTPError(http.StatusBadRequest, 1999, "InvalidID", "invalid user id")
	})

	e.POST("/users", AddNewUser)

	e.PUT("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, UpdateUser(c))
	})

	e.DELETE("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, DeleteUser(c))
	})

	// Start server
	e.Logger.Fatal(e.Start(viper.GetString("server.port")))
}

func AddNewUser(c echo.Context) error {
	fmt.Println("start post user")
	payload := model.User{} // coul
	if err := c.Bind(&payload); err != nil {
		util.ErrorCheck(err)
		fmt.Println("has error from request")
		//return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println("insert process....")
	repo.AddNewUser(payload)
	return c.String(http.StatusOK, "Success")
}

func UpdateUser(c echo.Context) string {
	payload := model.User{} // could/should be different type of struct so transport layer can be different from business domain
	if err := c.Bind(&payload); err != nil {
		util.ErrorCheck(err)
	}
	repo.UpdateUser(payload)
	return "update success"
}

func DeleteUser(c echo.Context) string {
	payload := model.User{} // could/should be different type of struct so transport layer can be different from business domain
	if err := c.Bind(&payload); err != nil {
		util.ErrorCheck(err)
	}
	repo.DeleteUser(payload)
	return "delete success"
}

func main() {
	//repo.DbConn()
	handleRequests()
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
