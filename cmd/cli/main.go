package main

import (
	"fmt"
	"os"

	"github-analyzer/internal/database"
	"github-analyzer/internal/services"
)

func main() {

	database.ConnectMongo()

	if len(os.Args) < 2 {

		fmt.Println("Comandos disponibles:")
		fmt.Println("users")

		return
	}

	command := os.Args[1]

	switch command {

	case "users":

		users, err := services.GetAllUsers()

		if err != nil {

			fmt.Println("Error:", err)

			return
		}

		fmt.Println("\n=== Usuarios Registrados ===\n")

		for i, user := range users {

			fmt.Printf(
				"%d. %s %s\n",
				i+1,
				user.Name,
				user.LastName,
			)

			fmt.Printf(
				"   Email: %s\n",
				user.GithubEmail,
			)

			fmt.Printf(
				"   Fecha Nac: %s\n\n",
				user.BirthDate,
			)
		}

	default:

		fmt.Println("Comando no reconocido")
	}
}
