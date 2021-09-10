package main

import (
	"fmt"
	"log"
	"net"
	"os"

	cli "github.com/urfave/cli/v2"
)

var (
	nsCommand = &cli.Command{
		Name:  "ns",
		Usage: "Looks Up the NameServers for a Particular Host",
		Action: func(c *cli.Context) error {
			host := c.Args().Get(0)
			if host == "" {
				log.Fatal("no valid info input")
			}
			ns, err := net.LookupNS(host)
			if err != nil {
				return err
			}
			fmt.Println("Get ns Host:")
			for i := range ns {
				fmt.Println(ns[i].Host)
			}
			return nil
		},
	}

	cnameCommand = &cli.Command{
		Name:  "cname",
		Usage: "Looks up the CNAME for a particular host",
		Action: func(c *cli.Context) error {
			host := c.Args().Get(0)
			if host == "" {
				log.Fatal("no valid info input")
			}
			cname, err := net.LookupCNAME(host)
			if err != nil {
				return err
			}
			fmt.Println("Get CNAME:")
			fmt.Println(cname)
			return nil
		},
	}

	ipCommand = &cli.Command{
		Name:  "ip",
		Usage: "Looks up the IP addresses for a particular host",
		Action: func(c *cli.Context) error {
			host := c.Args().Get(0)
			if host == "" {
				log.Fatal("no valid info input")
			}
			ips, err := net.LookupIP(host)
			if err != nil {
				return err
			}
			fmt.Println("Get ip:")
			for i := range ips {
				fmt.Println(ips[i].String())
			}
			return nil
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Commands = append(app.Commands, nsCommand, cnameCommand, ipCommand)
	app.Usage = "cli [global options] command [command options] [arguments...]"
	app.Name = "Website Lookup CLI - Let's you query IPs, CNAMEs and Name Servers!"
	app.Version = "0.0.0"
	app.Flags = append(app.Flags, cli.VersionFlag)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
