package routes

import (
	"context"
	"fmt"

	"github.com/abhinavv9/codee/internal/container"
	"github.com/abhinavv9/codee/types"
	"github.com/docker/docker/client"
	"github.com/gofiber/fiber/v2"
)

func JobHandler(ctx context.Context, cli *client.Client, jobCh <-chan types.Job, c *fiber.Ctx)  {
	job := <-jobCh
	r, e := container.SpinContainer(ctx, cli, job)
    c.SendString(r)

fmt.Println(e)	
}

