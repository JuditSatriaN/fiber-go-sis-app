package main

import (
	"embed"
	"log"
	"net/http"
	"time"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"

	apiRouter "github.com/fiber-go-sis-app/api/router"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/databases/postgres"
	webRouter "github.com/fiber-go-sis-app/web/app/router"
	goccyJson "github.com/goccy/go-json"
)

// Embed a template directory
//go:embed web/template/*
var embedDirTemplate embed.FS

// Embed a static directory
//go:embed web/static/*
var embedDirStatic embed.FS

// Embed a schema directory
//go:embed schema/postgres/*
var embedSchemaFiles embed.FS

// Embed a private pem files
//go:embed schema/pem/private.pem
var embedPrivatePEMFile []byte

func main() {
	// Initialization App Config
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 3 * time.Second,
		AppName:      constant.AppName,
		JSONEncoder:  goccyJson.Marshal,
		JSONDecoder:  goccyJson.Unmarshal,
		ErrorHandler: customPkg.CustomErrorHandler,
		Views:        html.NewFileSystem(http.FS(embedDirTemplate), ".html"),
	})

	// Load Environment
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Setting basic configuration
	app.Use(logger.New(), recover.New())

	// Setting static files in .static folder
	app.Use(constant.StaticUrl, filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "web/static",
		Browse:     true,
	}))

	// Setting key token to encrypt cookie
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: constant.CookiesKeyToken,
	}))

	// Setting JWT RS256
	if err := customPkg.GenerateJWT(embedPrivatePEMFile); err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	// Open Postgres Connection
	if err := postgresPkg.OpenConnection(); err != nil {
		panic(err)
	}

	// Setup schema if table postgres not exists
	if err := customPkg.SetupPostgresTable(embedSchemaFiles); err != nil {
		panic(err)
	}

	// Web handler
	webRouter.BuildWebRouter(app)

	// API Handler
	apiRouter.BuildAPIRouter(app)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
