package forms

type JWTTokenKey struct {
	AccessToken  string
	RefreshToken string
}

type JWTRequest struct {
	UserID  string
	Name    string
	IsAdmin bool
}
