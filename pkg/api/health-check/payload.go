//go:generate easyjson -lower_camel_case $GOFILE
package healthcheck

//easyjson:json
type Response struct {
	App App `json:"app"`
}

type App struct {
	Name      string   `json:"name"`
	Version   string   `json:"version"`
	GoVersion string   `json:"goVersion"`
	Codebase  Codebase `json:"codebase,omitempty"`
}

type Codebase struct {
	CommitHash string `json:"commit,omitempty"`
	Branch     string `json:"branch,omitempty"`
	Repository string `json:"repository,omitempty"`
}
