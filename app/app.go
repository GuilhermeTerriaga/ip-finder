package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate retorna aplicação
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "ip finder"
	app.Usage = "busca IPs e name servers na internet"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "hostname",
			Usage:  "Busca hostnames de endereços na internet",
			Flags:  flags,
			Action: buscarHostnames,
		},
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarHostnames(c *cli.Context) {
	host := c.String("host")
	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
