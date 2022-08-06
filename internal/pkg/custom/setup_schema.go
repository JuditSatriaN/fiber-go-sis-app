package custom

import (
	"embed"

	"github.com/fiber-go-sis-app/internal/pkg/database/postgres"
)

//SetupPostgresTable : Function to set up postgres table
func SetupPostgresTable(embedSchemaFiles embed.FS) error {
	// Initialize Connection
	db := postgres.GetPgConn()

	contents, err := GetAllContentFiles(&embedSchemaFiles, "")
	if err != nil {
		return err
	}

	for _, content := range contents {
		_, err := db.Exec(content)
		if err != nil {
			return err
		}
	}

	return nil
}
