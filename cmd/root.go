package cmd

import (
	"github.com/tashima42/rock-paper-scissors-server/server"
	"github.com/urfave/cli/v2"
)

func NewRootCommand() *cli.App {
	app := cli.NewApp()

	app.Commands = append(app.Commands, &cli.Command{
		Name:        "server",
		Description: "start the rock paper scissors server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "port",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "jwt-secret",
				EnvVars:  []string{"JWT_SECRET"},
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			port := ctx.String("port")
			jwtSecret := ctx.String("jwt-secret")
			return server.Serve(port, []byte(jwtSecret))
		},
	})
	return app
}
