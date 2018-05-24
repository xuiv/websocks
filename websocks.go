package main

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"os/exec"

	"io/ioutil"

	"time"

	"github.com/urfave/cli"
	"github.com/xuiv/websocks/core"
)

func main() {
	app := cli.NewApp()
	app.Name = "WebSocks"
	app.Version = "0.4.0"
	app.Usage = "A secure proxy based on WebSocket."
	app.Description = "See https://github.com/lzjluzijie/websocks"
	app.Author = "Halulu"
	app.Email = "lzjluzijie@gmail.com"

	app.Commands = []cli.Command{
		{
			Name:    "client",
			Aliases: []string{"c"},
			Usage:   "start websocks client",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "l",
					Value: ":10801",
					Usage: "local listening port",
				},
				cli.StringFlag{
					Name:  "s",
					Value: "ws://localhost:23333/websocks",
					Usage: "server url",
				},
				cli.StringFlag{
					Name:  "n",
					Value: "",
					Usage: "fake server name for tls client hello, leave blank to disable",
				},
				cli.BoolFlag{
					Name:  "insecure",
					Usage: "InsecureSkipVerify: true",
				},
			},
			Action: func(c *cli.Context) (err error) {
				listenAddr := c.String("l")
				serverURL := c.String("s")
				serverName := c.String("n")
				insecureCert := false
				if c.Bool("insecure") {
					insecureCert = true
				}

				u, err := url.Parse(serverURL)
				if err != nil {
					return
				}

				lAddr, err := net.ResolveTCPAddr("tcp", listenAddr)
				if err != nil {
					return
				}

				local := core.Client{
					ListenAddr:   lAddr,
					URL:          u,
					ServerName:   serverName,
					InsecureCert: insecureCert,
				}

				err = local.Listen()
				if err != nil {
					return
				}

				return nil
			},
		},
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start websocks server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "l",
					Value: "127.0.0.1:23333",
					Usage: "local listening port",
				},
				cli.StringFlag{
					Name:  "p",
					Value: "/websocks",
					Usage: "server.com/pattern, like password, start with '/'",
				},
				cli.BoolFlag{
					Name:  "tls",
					Usage: "enable built-in tls",
				},
				cli.StringFlag{
					Name:  "cert",
					Value: "websocks.cer",
					Usage: "tls cert path",
				},
				cli.StringFlag{
					Name:  "key",
					Value: "websocks.key",
					Usage: "tls key path",
				},
				cli.StringFlag{
					Name:  "proxy",
					Value: "",
					Usage: "reverse proxy url, leave blank to disable",
				},
			},
			Action: func(c *cli.Context) (err error) {
				listenAddr := c.String("l")
				pattern := c.String("p")
				tls := c.Bool("tls")
				certPath := c.String("cert")
				keyPath := c.String("key")
				proxy := c.String("proxy")

				server := core.Server{
					Pattern:    pattern,
					ListenAddr: listenAddr,
					TLS:        tls,
					CertPath:   certPath,
					KeyPath:    keyPath,
					Proxy:      proxy,
					CreatedAt:  time.Now(),
				}

				err = server.Listen()
				if err != nil {
					return
				}

				return
			},
		},
		{
			Name:    "github",
			Aliases: []string{"github"},
			Usage:   "open official github page",
			Action: func(c *cli.Context) (err error) {
				err = exec.Command("explorer", "https://github.com/lzjluzijie/websocks").Run()
				return
			},
		},
		{
			Name:    "cert",
			Aliases: []string{"cert"},
			Usage:   "generate self signed key and cert(default rsa 2048)",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "ecdsa",
					Usage: "generate ecdsa key and cert(P-256)",
				},
				cli.StringSliceFlag{
					Name:  "hosts",
					Value: nil,
					Usage: "certificate hosts",
				},
			},
			Action: func(c *cli.Context) (err error) {
				ecdsa := c.Bool("ecdsa")
				hosts := c.StringSlice("hosts")

				var key, cert []byte
				if ecdsa {
					key, cert, err = core.GenP256(hosts)
					fmt.Println("Generated ecdsa P-256 key and cert")
				} else {
					key, cert, err = core.GenRSA2048(hosts)
					fmt.Println("Generated rsa 2048 key and cert")
				}

				err = ioutil.WriteFile("websocks.key", key, 0600)
				if err != nil {
					return
				}
				err = ioutil.WriteFile("websocks.cer", cert, 0600)
				if err != nil {
					return
				}
				return
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}

}
