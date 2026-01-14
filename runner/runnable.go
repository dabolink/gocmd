package runner

import (
	"context"
)

type Runnable[TProvider any] interface {
	Run(ctx context.Context, provider TProvider) error
}
