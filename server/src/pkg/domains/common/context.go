package common

import (
	"context"
	"time"
)

func NewQueryContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
