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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
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
