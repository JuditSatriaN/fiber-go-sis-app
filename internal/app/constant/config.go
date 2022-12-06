package constant

var LinkPageList = map[string]string{
	// Login URL a
	"WebLoginURL": WebLoginURL,

	// SIS URL
	"WebSISHome":        WebSISHomeURL,
	"WebSISUser":        WebSISUserURL,
	"WebSISUnit":        WebSISUnitURL,
	"WebSISMember":      WebSISMemberURL,
	"WebSISProduct":     WebSISProductURL,
	"WebSISInventory":   WebSISInventoryURL,
	"WebSISUpdateStock": WebSISUpdateStockURL,
}

type WebData struct {
	Title        string
	BaseURL      string
	CurrentURL   string
	StaticUrl    string
	LinkPageList map[string]string
}
