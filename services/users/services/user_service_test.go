package services

import (
	"fmt"
	"time"
)

type User struct {
	ID   int
	Name string
	// Otros campos relevantes del usuario
}

func getUserByID(userID int) (*User, error) {
	// Simula la búsqueda del usuario por ID en algún servicio o base de datos
	// En este ejemplo, simplemente simulamos una búsqueda con un retraso.
	time.Sleep(time.Millisecond * 500)

	// Retorna un usuario ficticio
	return &User{
		ID:   userID,
		Name: fmt.Sprintf("User%d", userID),
	}, nil
}

func main() {
	// Número de solicitudes concurrentes
	numSolicitudes := 10

	// Canal para recibir los resultados de las solicitudes
	resultados := make(chan *User)

	for i := 1; i <= numSolicitudes; i++ {
		// Lanza una goroutine para cada solicitud concurrente
		go func(userID int) {
			user, err := getUserByID(userID)
			if err != nil {
				fmt.Printf("Error al obtener usuario con ID %d: %v\n", userID, err)
				return
			}
			// Envía el resultado a través del canal
			resultados <- user
		}(i)
	}

	// Esp
	close(resultados)

	// Recorre el canal para obtener los resultados de las solicitudes
	for user := range resultados {
		fmt.Printf("Usuario obtenido: %s\n", user.Name)
	}
}
