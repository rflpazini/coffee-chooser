package healthcheck

import (
	"net/http"
	"os"
	"runtime/debug"

	"coffee-choose/pkg/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type makeHealthCheckParams struct {
	dig.In

	*config.ServerConfig
}

func makeHealthCheckHandler(p makeHealthCheckParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		var commitHash string
		var goVersion string

		if bi, ok := debug.ReadBuildInfo(); ok {
			goVersion = bi.GoVersion

			for _, kv := range bi.Settings {
				if kv.Key == "vcs.revision" {
					commitHash = kv.Value
					break
				}
			}
		}

		rsp := Response{
			App: App{
				Name:      p.ServerConfig.AppName,
				Version:   p.ServerConfig.AppVersion,
				GoVersion: goVersion,
				Codebase: &Codebase{
					Repository: p.ServerConfig.Repository,
					CommitHash: commitHash,
					Branch:     os.Getenv("BRANCH_NAME"),
				},
				Environment: &Environment{
					Name:       hostName(),
					Region:     os.Getenv("EC2_REGION"),
					InstanceId: os.Getenv("INSTANCE_ID"),
				},
			},
		}

		return c.JSON(http.StatusOK, rsp)
	}
}

func hostName() (name string) {
	name, _ = os.Hostname()
	if os.Getenv("STACK") != "" {
		name = os.Getenv("STACK")
	}

	return
}
