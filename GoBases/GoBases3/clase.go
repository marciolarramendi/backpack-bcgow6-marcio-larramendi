package main

import (
	"io"
	"os"
)

func main() {
	/*
		//SETEO DE VARIABLE DE ENTORNO
		err := os.Setenv("NAME", "Marcio")
		if err != nil {
			panic(err)
		}

		//MOSTRAR VALOR DE VARIABLE DE ENTORNO
		name := os.Getenv("NAME")
		fmt.Println(name)

		//LookupEnv para chequear si existe la variable de entorno
		directorioRaiz, ok := os.LookupEnv("HOME")
		if ok {
			fmt.Println(userprofile)
		} else {
			fmt.Println("No existe")
		}
		userprofile, ok := os.LookupEnv("USER")
		if ok {
			fmt.Println(userprofile)
		} else {
			fmt.Println("No existe")
		}

	*/

	/*
		//USO de EXIT (hay distintos tipos de codigo de exit, el 0 no imprime nada)
		flag := false
		if flag {
			fmt.Println("flag is true")
		} else {
			fmt.Println("flag is false")
			os.Exit(1)
		}

		fmt.Println("Fin de las instrucciones")
	*/

	/*
		//ReadDir lee lo que haya en el directorio, con . lee el directorio actual
		//con .. lee el directorio anterior
		entries, err := os.ReadDir("..")
		if err != nil {
			fmt.Printf("Collected entries: %v\n", entries)
			panic(err)
		}

		fmt.Printf("Collected entries: %v\n", entries)

		// Con este for revisamos los nombres, sino solo muestra algo como [0x14000186040]
		for _, value := range entries {
			fmt.Println(value.Name())
		}
	*/

	/*
		// Con ReadFile leemos archivos
		fileContentAsBytes, err := os.ReadFile("../go.mod")
		if err != nil {
			panic(err)
		}

		//Asi se ve en bytes
		fmt.Println(fileContentAsBytes)

		//Asi se ve el contenido del file cambiandolo a string
		fmt.Println(string(fileContentAsBytes))
	*/

	/*
		// Con WriteFile escribimos archivos
		message := "Hola, esto es un test de creacion de file :D"

		//0644 es el permiso, hay otros
		err := os.WriteFile("./test.txt", []byte(message), 0644)

		if err != nil {
			panic(err)
		}

		fmt.Println("Archivo creado satisfactoriamente")
	*/

	/*
		//Copiar archivos con io.Copy
		reader := strings.NewReader(
			"Hola mundo",
		)

		_, err := io.Copy(os.Stdout, reader)

		if err != nil {
			panic(err)
		}
	*/

	/*
		//ReadAll
		reader := strings.NewReader(
			"Hola mundo",
		)

		b, err := io.ReadAll(reader)

		if err != nil {
			panic(err)
		}

		println(string(b))
	*/

	//WriteString
	io.WriteString(
		os.Stdout,
		"Hola Mundo",
	)
}
