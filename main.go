package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func encodeHandler(c *cli.Context) error {
	path := c.Args().Get(0)
	if path == "" {
		return cli.NewExitError("must specify file path", 1)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		m := fmt.Sprintf("can't open %s", path)
		return cli.NewExitError(m, 1)
	}

	if c.Bool("uri") {
		mime := http.DetectContentType(data)
		fmt.Printf("data:%s;base64,", mime)
	}

	_, err = base64.NewEncoder(base64.StdEncoding, os.Stdout).Write(data)
	if err != nil {
		return cli.NewExitError("fail to encode base64", 1)
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "cfb64"
	app.Usage = "convert file to/from Base64"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "encode",
			Aliases: []string{"e"},
			Usage:   "Encode file to Base64",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "uri",
					Usage: "format output with data URI scheme",
				},
			},
			Action: encodeHandler,
		},
	}

	app.Run(os.Args)
}
