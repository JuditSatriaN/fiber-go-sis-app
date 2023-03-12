package invoice

import (
	"database/sql"
	"github.com/fiber-go-sis-app/internal/app/model"
	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/database/postgres"
	"github.com/gofiber/fiber/v2"
)

const queryGetInvoice = `
	SELECT head_fak, last_counter
	FROM invoices
	WHERE head_fak = $1
`

func GetInvoice(ctx *fiber.Ctx, headFak string) (model.Invoice, bool, error) {
	var invoice model.Invoice

	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &invoice, queryGetInvoice, headFak); err != nil {
		if err == sql.ErrNoRows {
			return invoice, false, nil
		}

		return invoice, false, err
	}

	return invoice, true, nil
}

const queryUpsertInvoice = `
	INSERT INTO invoices (head_fak, last_counter, create_time)
	VALUES (:head_fak, 1, NOW()) 
	ON CONFLICT (head_fak) DO UPDATE 
	SET last_counter = invoices.last_counter + 1,
		update_time = NOW()
`

func UpsertInvoice(ctx *fiber.Ctx, invoice model.Invoice) error {
	db := postgresPkg.GetPgConn()
	_, err := db.NamedQueryContext(ctx.Context(), queryUpsertInvoice, invoice)
	return err
}
