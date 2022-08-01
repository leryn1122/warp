package main

import (
	"kreutzer/pkg/cluster"
	"kreutzer/pkg/version"
	"kreutzer/pkg/web"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	kubeconfig string
)

func main() {

	// os.Unsetenv("")
	// os.Unsetenv("")

	app := cli.NewApp()
	app.Version = ""
	app.Usage = "Kreutzer Devops Platform"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "kubeconfig",
			Usage:       "Path to kube config for access to k8s cluster.",
			EnvVar:      "KUBECONFIG",
			Destination: &kubeconfig,
		},
	}

	app.Action = func(c *cli.Context) error {
		return run(c)
	}

	cluster.ListNamespaces()

	app.ExitErrHandler = func(c *cli.Context, err error) {
		logrus.Fatal(err)
	}

	app.Run(os.Args)
}

func run(cli *cli.Context) error {
	logrus.Infof("Kreutzer version %s is running.", version.CurrentVersion())
	logrus.Infof("Kreutzer arguments: ")

	os.Unsetenv("KUBECONFIG")

	web.StartWebServer()
	return nil
}
