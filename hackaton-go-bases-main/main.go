package main

import (
	"fmt"
	"hackaton-go-bases-main/internal/file"
	"hackaton-go-bases-main/internal/service"
)

var filepath = file.File{Path: "./tickets.csv"}

func main() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Printf("Ha ocurrido un error. Error: %v\n", err)
		}
	}()

	//filepath := file.File{Path: "./tickets.csv"}
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	tickets, err := filepath.Read()
	if err != nil {
		panic(err)
	}
	//fmt.Println(tickets)
	b := service.NewBookings(tickets)

	ticketAdd, err := b.Create(
		service.Ticket{
			Id:          1001,
			Names:       "Martin Perez",
			Email:       "martinperez@gmail.com",
			Destination: "Montevideo",
			Date:        "20:23",
			Price:       350})
	if err != nil {
		panic(err)
	}
	fmt.Println(ticketAdd)

	ticket, err := b.Read(999)
	if err != nil {
		panic(err)
	}
	fmt.Println(ticket)

	ticketUpd, err := b.Update(999,
		service.Ticket{
			Names:       "Pedro Lopez",
			Email:       "pedrolopez@gmail.com",
			Destination: "Buenos Aires",
			Date:        "15:23",
			Price:       450})
	if err != nil {
		panic(err)
	}
	fmt.Println(ticketUpd)

	ticketDel, err := b.Delete(3)
	if err != nil {
		panic(err)
	}
	fmt.Println(ticketDel)

	//fmt.Println(b)

	// ticketAdd, err := b.Create(service.Ticket{1001, "Martin Perez", "martinp@gmail.com", "Montevideo", "18:19", 1400})
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(ticketAdd)
	// fmt.Println(b)

	// records := readCsvFile("tickets.csv")
	// fmt.Println(records[1])
	// for i := 0; i < len(records); i++ {
	// 	id, _ := strconv.Atoi(records[i][0])
	// 	precio, _ := strconv.Atoi(records[i][5])
	// 	testito = append(testito, service.Ticket{id, records[i][1], records[i][2], records[i][3], records[i][4], precio})
	// 	fmt.Println(testito)
	// }

	// tickets := strings.Split(string(records), "\n")
	// for i := 0; i < len(tickets)-1; i++ {
	// 	ticket := strings.Split(records[i], ",")
	// 	fmt.Printf("%s \t%s \t%s \t%s \t%s\n", ticket[0], ticket[1], ticket[2], ticket[3], ticket[4])
	// }

	// for i := 0; i < 1; i++ {
	// 	for j := 0; j < len(records[i]); j++{
	// 		fmt.Println(records[i][j])
	// 	}
	// 	// stringcito := records[i]
	// 	// fmt.Println(stringcito)
	// 	// testito = append(testito, service.Ticket{record, "Martin Perez", "martinp@gmail.com", "Montevideo", "18:19", 1400})
	// }
	// fmt.Println(valores)
	// for _, valor := range valores {
	// // 	//testito = append(testito, service.Ticket{strconv.ParseInt(valor[0], 6, 12), strconv.Itoa(valor[1]), valor[2], valor[3], valor[4], valor[5]})
	// // 	fmt.Println(valor[1])
	// }
	//fmt.Println(valores)
	//testito = append(testito, service.Ticket{valor[0], "Martin Perez", "martinp@gmail.com", "Montevideo", "18:19", 1400})

	// for i := 0; i < len(records); i++ {
	// 	testito = append(testito, service.Ticket{record, "Martin Perez", "martinp@gmail.com", "Montevideo", "18:19", 1400})
	// }
	// tickets := strings.Split(string(records), "\n")
	// for i := 0; i < len(tickets)-1; i++ {
	// 	ticket := strings.Split(records[i], ",")
	// 	fmt.Printf("%s \t%s \t%s \t%s \t%s\n", ticket[0], ticket[1], ticket[2], ticket[3], ticket[4])
	// }

	//fmt.Println(records)

	// 	records := []Employee{
	// 		{"E01", 25},
	// 		{"E02", 26},
	// 		{"E03", 24},
	// 		{"E04", 26},
	// 	}

	// 	//file, err := os.OpenFile("records.csv", os.O_APPEND|os.O_CREATE, 0666)
	// 	file, err := os.Open("records.csv")
	// 	//defer file.Close()
	// 	if err != nil {
	// 		log.Fatalln("failed to open file", err)
	// 	}
	// 	w := csv.NewWriter(file)
	// 	defer w.Flush()
	// 	// Using Write
	// 	for _, record := range records {
	// 		row := []string{record.ID, strconv.Itoa(record.Age)}
	// 		if err := w.Write(row); err != nil {
	// 			log.Fatalln("error writing record to file", err)
	// 		}
	// 	}

	// 	// Using WriteAll
	// 	var data [][]string
	// 	for _, record := range records {
	// 		row := []string{record.ID, strconv.Itoa(record.Age)}
	// 		data = append(data, row)
	// 	}
	// 	w.WriteAll(data)
	// }

	// package main

	// import (
	// 	"encoding/csv"
	// 	"fmt"
	// 	"hackaton-go-bases-main/internal/service"
	// 	"log"
	// 	"os"
	// 	"strconv"
	// )

	// func main() {
	// 	//var choice string
	// 	var tickets []service.Ticket
	// 	// Funcion para obtener tickets del archivo csv
	// 	//service.NewBookings(tickets)
	// 	fmt.Println(tickets)

	// 	//file, err := os.OpenFile("./test.csv", os.O_APPEND|os.O_CREATE, os.ModePerm)
	// 	file, err := os.Create("./test2.csv")

	// 	defer file.Close()

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	tickets = append(tickets, service.Ticket{1001, "Martin Perez", "martinp@gmail.com", "Montevideo", "18:19", 1400})

	// 	//fmt.Println(tickets)
	// 	df := csv.NewWriter(file)
	// 	df.Flush()

	// 	// Using Write
	// 	for _, record := range tickets {
	// 		row := []string{strconv.Itoa(record.Id), record.Names, record.Email, record.Destination, record.Date, strconv.Itoa(record.Price)}
	// 		fmt.Println(row)
	// 		if err := df.Write(row); err != nil {
	// 			log.Fatalln("error writing record to file", err)
	// 		}
	// 	}

	// 	// for _, record := range tickets {
	// 	// 	err := df.Write([]string{string(rune(record.Id)), record.Names, record.Email, record.Destination, record.Date, string(rune(record.Price))})
	// 	// 	if err != nil {
	// 	// 		panic(err)
	// 	// 	}
	// 	// }
	// 	//df.Flush()
	// 	fmt.Println("Record inserted in CSV file")

	// 	// forlabel:
	// 	// 	for{
	// 	// 		fmt.Println("Do you want to write data in CSV file")
	// 	// 		fmt.Scanln(&choice)
	// 	// 		switch choice {
	// 	// 		case "yes":
	// 	// 			var Id int
	// 	// 			var Names, Email, Destination, Date string
	// 	// 			var Price int

	// 	// 		case "no":
	// 	// 			break forlabel
	// 	// 		}
	// 	// 	}

	// 	// empData := [][]string{
	// 	// 	{"Juan", "Uruguay", "Python"},
	// 	// }

	// 	// csvFile, err := os.Create("employee.csv")

	// 	// if err != nil {
	// 	// 	log.Fatalf("failed creating file: %s", err)
	// 	// }

	// 	// csvwriter := csv.NewWriter(csvFile)

	// 	// for _, empRow := range empData {
	// 	// 	_ = csvwriter.Write(empRow)
	// 	// }
	// 	// csvwriter.Flush()
	// 	// csvFile.Close()

	// 	// contenido, err := os.ReadFile("./tickets.csv")
	// 	// if err != nil {
	// 	// 	defer func() {
	// 	// 		fmt.Println("ejecución finalizada")
	// 	// 	}()
	// 	// 	panic("el archivo indicado no fue encontrado o está dañado")
	// 	// }

	// 	// ticketss := strings.Split(string(contenido), "\n")
	// 	// for i := 0; i < len(ticketss)-1; i++ {
	// 	// 	ticket := strings.Split(ticketss[i], ",")
	// 	// 	fmt.Printf("%s \t%s \t%s \t%s \t%s\n", ticket[0], ticket[1], ticket[2], ticket[3], ticket[4])
	// 	// 	//precio, _ := strconv.ParseFloat(empleado[1], 64)
	// 	// }

}
