package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fseverino1981-environment-configurator/config"
)

func main() {
    env :=flag.String("env", "dev", "Os ambientes disponíveis são : dev, hml e prd.")
	flag.Parse()

	if (*env != "dev" || *env != "hml" || *env != "prd"){
		fmt.Printf("Ambiente %s inválido. Os ambientes disponíveis são : dev, hml e prd.\n",*env)
		os.Exit(0)
	}

	config := config.GetConfig(*env)

	fmt.Printf("Ambiente: %s\n", *env)
	fmt.Printf("Aplicação: %s\n", config.AppName) 
	fmt.Printf("Configuração do Banco de Dados:\n")
    fmt.Printf("  Host: %s\n", config.Database.Host)
    fmt.Printf("  Port: %d\n", config.Database.Port)
    fmt.Printf("  Usuário: %s\n", config.Database.Username)
}
