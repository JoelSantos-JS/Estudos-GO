package app


import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

// Gerar vai retornar a linha de comando
func Gerar() *cli.App {
		app := cli.NewApp()
		app.Name  = "Aplicaçao de linha de Comando"
		app.Usage = "Busca Ips e nome de servidor na internet"

		flags := []cli.Flag{
			cli.StringFlag{
				Name: "host",
				Value: "google.com",
			},
		}

		app.Commands = []cli.Command{
			{
				Name: "ip",
				Usage: "Busca ips da net",
				Flags: flags,

				Action: buscarIp,
			},
			{
				Name: "servidores",
				Usage: "Buscas servidores",
				Flags: flags,
				Action: buscarServer,
			},
		}


		return app

}

func buscarIp(c *cli.Context) {
	host := c.String("host")

	ips , erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _ , ip := range ips {
		fmt.Println(ip)
	}
}
func buscarServer(c *cli.Context) {
	host := c.String("host")

	servidores , erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _ , server := range servidores {
		fmt.Println(server)
	}
}