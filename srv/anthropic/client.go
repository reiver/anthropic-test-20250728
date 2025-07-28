package anthropicsrv

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"

	"github.com/reiver/anthropic-test-20250728/cfg"
)

var Client anthropic.Client = anthropic.NewClient(
	option.WithAPIKey(cfg.ApiKey),
)
