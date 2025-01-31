package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/alexbrainman/odbc"
	_ "github.com/databricks/databricks-sql-go"

)
