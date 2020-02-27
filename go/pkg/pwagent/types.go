package pwagent

import (
	"time"

	"go.uber.org/zap"
	"pathwar.land/v2/go/internal/randstring"
)

type Opts struct {
	DomainSuffix      string
	HostIP            string
	HostPort          string
	ModeratorPassword string
	Salt              string
	ForceRecreate     bool
	NginxDockerImage  string
	Cleanup           bool
	RunOnce           bool
	LoopDelay         time.Duration
	HTTPAPIAddr       string
	Name              string

	Logger *zap.Logger
}

func (opts *Opts) applyDefaults() error {
	if opts.Logger == nil {
		opts.Logger = zap.NewNop()
	}
	if opts.Salt == "" {
		opts.Salt = randstring.RandString(10)
		opts.Logger.Warn("random salt generated", zap.String("salt", opts.Salt))
	}
	if opts.ModeratorPassword == "" {
		opts.ModeratorPassword = randstring.RandString(10)
		opts.Logger.Warn("random moderator password generated", zap.String("password", opts.ModeratorPassword))
	}
	return nil
}