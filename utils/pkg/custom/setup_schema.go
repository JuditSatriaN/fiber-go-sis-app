package custom

import (
	"embed"

	"github.com/fiber-go-sis-app/utils/pkg/databases/postgres"
)

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
