package defs

// request
type UserCredential struct {
	UserName string `json:"user_name"`
	Pws string 	`json:"pwd"`
}

//response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

// session
type SimpleSession struct {
	Username string
	TTl int64
}
