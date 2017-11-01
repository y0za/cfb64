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
	if c.NArg() == 0 {
		return cli.NewExitError("must specify file path", 1)
	}

	args := c.Args()
	for i := 0; i < c.NArg(); i++ {
		path := args.Get(i)

		data, err := ioutil.ReadFile(path)
		if err != nil {
			m := fmt.Sprintf("can't open %s", path)
			return cli.NewExitError(m, 1)
		}

		r := base64.StdEncoding.EncodeToString(data)
		if c.Bool("uri") {
			mime := http.DetectContentType(data)
			r = fmt.Sprintf("data:%s;base64,", mime) + r
		}
		fmt.Println(r)
	}

	return nil
}

func decodeHandler(c *cli.Context) error {
	if c.NArg() == 0 {
		return cli.NewExitError("must specify file path", 1)
	}
	path := c.Args().Get(0)

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return cli.NewExitError("can't read stdin", 1)
	}

	data, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		return cli.NewExitError("fail to decode", 1)
	}

	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		m := fmt.Sprintf("can't write data to %s", path)
		return cli.NewExitError(m, 1)
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "cfb64"
	app.Usage = "convert file to/from Base64"
	app.Version = "0.2.0"
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
		{
			Name:    "decode",
			Aliases: []string{"d"},
			Usage:   "Decode Base64 to file",
			Action:  decodeHandler,
		},
	}

	app.Run(os.Args)
}
