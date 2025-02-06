package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/databricks/databricks-sql-go"
	_ "github.com/gofiber/swagger"
	//_ "github.com/swaggo/swag/cmd/swag"

	_ "github.com/go-playground/validator/v10"

)
