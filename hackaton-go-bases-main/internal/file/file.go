package file

import (
	"encoding/csv"
	"hackaton-go-bases-main/internal/service"
	"os"
	"strconv"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {

	var arrayTickets []service.Ticket
	tickets, err := os.Open(f.Path)
	if err != nil {
		panic("No se pudo abrir el file")
	}
	defer tickets.Close()

	csvReader := csv.NewReader(tickets)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic("No se pudo leer el File")
	}

	for i := 0; i < len(records); i++ {
		id, _ := strconv.Atoi(records[i][0])
		precio, _ := strconv.Atoi(records[i][5])
		arrayTickets = append(arrayTickets, service.Ticket{id, records[i][1], records[i][2], records[i][3], records[i][4], precio})
	}

	return arrayTickets, nil

}

func (f *File) Write(service.Ticket) error {
	return nil
}
