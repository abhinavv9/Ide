package types

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/gofiber/fiber/v2"
)

type Job struct {
	UserID string
	Code   string
	Lang   string
	Image  string
}
l
