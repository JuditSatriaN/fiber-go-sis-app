package main

import (
	"embed"
	"log"
	"net/http"
	"time"

	goccyJson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"

	serviceRoutes "github.com/fiber-go-sis-app/routes/services"
	webRoutes "github.com/fiber-go-sis-app/routes/web"

	customPkg "github.com/fiber-go-sis-app/utils/pkg/custom"
	postgresPkg "github.com/fiber-go-sis-app/utils/pkg/databases/postgres"

	jwtMiddleware "github.com/fiber-go-sis-app/utils/pkg/middleware/jwt"
)

// Embed a template directory
//go:embed templates/*
var embedDirTemplate embed.FS

// Embed a static directory
//go:embed static/*
var embedDirStatic embed.FS

// Embed a schemes' directory
//go:embed utils/schemes/postgres/*
var embedSchemaFiles embed.FS

// Embed a private pem files
//go:embed utils/schemes/pem/private.pem
var embedPrivatePEMFile []byte

func main() {
	// Initialization App Config
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 3 * time.Second,
		JSONEncoder:  goccyJson.Marshal,
		JSONDecoder:  goccyJson.Unmarshal,
		AppName:      constantsEntity.AppName,
		ErrorHandler: webRoutes.CustomErrorHandler,
		Views:        html.NewFileSystem(http.FS(embedDirTemplate), ".html"),
	})

	// Load Environment
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Setting basic configuration
	app.Use(logger.New(), recover.New())

	// Setting static files in .static folder
	app.Use(constantsEntity.StaticUrl, filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "static",
		Browse:     true,
	}))

	// Setting key token to encrypt cookie
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: constantsEntity.CookiesKeyToken,
	}))

	// Setting JWT RS256
	if err := customPkg.GenerateJWT(embedPrivatePEMFile); err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	// Open Postgres Connection
	if err := postgresPkg.OpenConnection(); err != nil {
		panic(err)
	}

	// Setup schema if table postgres / index elasticsearch not exists
	if err := customPkg.SetupPostgresTable(embedSchemaFiles); err != nil {
		panic(err)
	}

	// Web handler login
	webRoutes.BuildLoginRoutes(app)
	webRoutes.BuildError404NotFound(app)

	// Web handler - SIS
	sisGroup := app.Group("/sis")
	sisGroup.Use(jwtMiddleware.AccessTokenMiddleware(constantsEntity.WebSource))
	sisGroup.Use(jwtMiddleware.RefreshTokenMiddleware(constantsEntity.WebSource))
	webRoutes.BuildSISRoutes(sisGroup)

	// Service Group
	svcGroup := app.Group("/svc")
	serviceRoutes.BuildUserRoutes(svcGroup)
	serviceRoutes.BuildLoginRoutes(svcGroup)
	serviceRoutes.BuildMemberRoutes(svcGroup)
	serviceRoutes.BuildProductRoutes(svcGroup)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
