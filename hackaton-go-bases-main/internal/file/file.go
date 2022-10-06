package file

import (
	"encoding/csv"
	"encoding/json"
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
		arrayTickets = append(arrayTickets, service.Ticket{
			Id:          id,
			Names:       records[i][1],
			Email:       records[i][2],
			Destination: records[i][3],
			Date:        records[i][4],
			Price:       precio})
	}

	return arrayTickets, nil

}

func (f *File) Write() error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.Path, fileData, 0644)
}

//fmt.Println(newFile)
// //records := service.Ticket
// file, err := os.OpenFile("records.csv", os.O_APPEND|os.O_CREATE, 0666)
// //file, err := os.Open("records.csv")
// defer file.Close()
// if err != nil {
// 	log.Fatalln("failed to open file", err)
// }
// w := csv.NewWriter(file)
// defer w.Flush()

// w := csv.NewWriter(f)
// err = w.WriteAll(records) // calls Flush internally

// // Using Write
// for _, record := range records {
// 	row := []string{record.ID, strconv.Itoa(record.Age)}
// 	if err := w.Write(row); err != nil {
// 		log.Fatalln("error writing record to file", err)
// 	}
// }
