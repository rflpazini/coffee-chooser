//go:generate easyjson -lower_camel_case $GOFILE

package auth

//easyjson:json
type Response struct {
	SessionID string `json:"session_id"`
	Token     string `json:"session_token"`
}
