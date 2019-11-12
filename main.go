package main

import (
	"fmt"
	"os"

	"github.com/docker/go-plugins-helpers/volume"
	"github.com/mikroskeem/docker-zfs-plugin/zfs"
	"github.com/urfave/cli"
)

const (
	version = "1.0.2"
)

func main() {

	app := cli.NewApp()
	app.Name = "docker-zfs-plugin"
	app.Usage = "Docker ZFS Plugin"
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "dataset-name",
			Usage: "Name of the ZFS dataset to be used. It will be created if it doesn't exist.",
		},
	}
	app.Action = Run
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

// Run runs the driver
func Run(ctx *cli.Context) error {
	if ctx.String("dataset-name") == "" {
		return fmt.Errorf("zfs dataset name is a required field")
	}

	d, err := zfsdriver.NewZfsDriver(ctx.StringSlice("dataset-name")...)
	if err != nil {
		return err
	}
	h := volume.NewHandler(d)

	return h.ServeUnix("zfs", 0)
}
