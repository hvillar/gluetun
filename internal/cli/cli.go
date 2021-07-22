// Package cli defines an interface CLI to run command line operations.
package cli

import (
	"context"

	"github.com/qdm12/golibs/logging"
	"github.com/qdm12/golibs/os"
	"github.com/qdm12/golibs/params"
)

type CLI interface {
	ClientKey(args []string, openFile os.OpenFileFunc) error
	HealthCheck(ctx context.Context, env params.Env, os os.OS, logger logging.Logger) error
	OpenvpnConfig(os os.OS, logger logging.Logger) error
	Update(ctx context.Context, args []string, os os.OS, logger logging.Logger) error
}

type cli struct{}

func New() CLI {
	return &cli{}
}
