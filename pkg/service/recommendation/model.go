//go:generate easyjson -lower_camel_case $GOFILE
package recommendation

//easyjson:json
type Response struct {
	Variety     string `json:"variety"`
	Description string `json:"description"`
}
