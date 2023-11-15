//go:generate easyjson -lower_camel_case $GOFILE
package healthcheck

//easyjson:json
type Response struct {
	App App `json:"app"`
}

type App struct {
	Name        string       `json:"name"`
	Version     string       `json:"version"`
	GoVersion   string       `json:"goVersion"`
	Codebase    *Codebase    `json:"codebase,omitempty"`
	Environment *Environment `json:"environment,omitempty"`
}

type Codebase struct {
	Branch     string `json:"branch,omitempty"`
	CommitHash string `json:"commit,omitempty"`
	Repository string `json:"repository,omitempty"`
}

type Environment struct {
	Name       string `json:"name,omitempty"`
	Region     string `json:"region,omitempty"`
	InstanceId string `json:"instanceId,omitempty"`
}
