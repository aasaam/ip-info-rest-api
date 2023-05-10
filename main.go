package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/urfave/cli/v2"
)

func runServer(c *cli.Context) error {
	gp, gpErr := newGeoParser(c.String("mmdb-city-path"), c.String("mmdb-asn-path"))

	if gpErr != nil {
		return gpErr
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               false,
		UnescapePath:          true,
		CaseSensitive:         true,
		StrictRouting:         true,
		BodyLimit:             c.Int("body-limit-size"),

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			ctx.Status(code).SendString(err.Error())

			return nil
		},
	})

	if c.String("basic-auth-username") != "" && c.String("basic-auth-password") != "" {

		app.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				c.String("basic-auth-username"): c.String("basic-auth-password"),
			},
		}))
	}

	app.Get("/info/:ip", func(c *fiber.Ctx) error {
		gr := gp.newResultFromIP(c.Params("ip"))

		jB, _ := json.Marshal(gr)

		c.Type("application/json", "utf8")
		return c.Send(jB)
	})

	return app.Listen(c.String("listen"))
}

func main() {

	app := cli.NewApp()
	app.Usage = "IP information rest api"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:   "run",
			Usage:  "Run server",
			Action: runServer,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "listen",
					Usage:    "Application listen http ip:port address",
					Value:    "0.0.0.0:4000",
					Required: false,
					EnvVars:  []string{"ASM_IP_INFO_LISTEN_ADDRESS"},
				},
				&cli.IntFlag{
					Name:     "body-limit-size",
					Usage:    "Request limit size",
					Value:    1 * 1024 * 1024,
					Required: false,
					EnvVars:  []string{"ASM_IP_INFO_BODY_LIMIT_SIZE"},
				},

				&cli.StringFlag{
					Name:     "basic-auth-username",
					Usage:    "Basic authentication username",
					Value:    "",
					Required: false,
					EnvVars:  []string{"ASM_IP_INFO_BASIC_AUTH_USERNAME"},
				},

				&cli.StringFlag{
					Name:     "basic-auth-password",
					Usage:    "Basic authentication password",
					Value:    "",
					Required: false,
					EnvVars:  []string{"ASM_IP_INFO_BASIC_AUTH_PASSWORD"},
				},

				&cli.StringFlag{
					Name:     "mmdb-city-path",
					Usage:    "MMDB city database path",
					Value:    "/tmp/GeoLite2-City.mmdb",
					Required: false,
					EnvVars:  []string{"ASM_IP_INFO_COLLECTOR_MMDB_CITY_PATH"},
				},

				&cli.StringFlag{
					Name:     "mmdb-asn-path",
					Usage:    "MMDB asn database path",
					Value:    "/tmp/GeoLite2-ASN.mmdb",
					Required: false,
					EnvVars:  []string{"ASM_IP_INFO_COLLECTOR_MMDB_ASN_PATH"},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
