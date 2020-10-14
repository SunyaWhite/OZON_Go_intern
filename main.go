package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// Структура, хранящая информацию для отчета
type ReportRecord struct {
	AllQuantity int // Общее кол-во
	AllCost float64 // Общая стоимость
}

// метод закрытия файла
func closeFile( file *os.File){

	err := file.Close()

	if err != nil{
		log.Fatal("Fatal error occurred ", err)
		panic(err)
	}

	log.Println("File closed")
}

// метод откртыия файла
func openFile(path string) *os.File{

	file, err := os.OpenFile(path, os.O_RDONLY, 744)

	if err != nil{
		log.Fatal("Fatal error occurred ", err)
		panic(err)
	}

	fileExt := filepath.Ext(file.Name())

	if fileExt != ".csv"{
		log.Fatal("Fatal error occurred: file should be csv ")
		panic("Fatal error occurred: file should be csv ")
	}

	log.Println("Opened file")
	return file
}

func generateReport(file *os.File){

	// Словарь, хранящий информацию
	var goods = make(map[string] ReportRecord)

	// Ридер csv-файла
	reader := csv.NewReader(file)
	firstRead := false

	for{
		line, err := reader.Read()

		// Конец файла
		if err == io.EOF{
			log.Println("End of file")
			break
		}
		// Ошибка :(
		if err != nil{
			log.Fatal("Fatal error occurred ", err)
			panic(err)
		}

		// Считываем перввую строку с заголовком
		if !firstRead{
			firstRead = true
			continue
		}

		name := line[2]
		cost, _ := strconv.ParseFloat(line[3], 32)
		quantity, _ := strconv.Atoi(line[4])

		good, exists := goods[name]

		if exists{
			// Вносим изменения
			good.AllCost += cost
			good.AllQuantity += quantity
			// Записываем новый результат
			goods[name] = good
		} else{
			// Создаем новую запись
			goods[name] = ReportRecord{
				AllQuantity: quantity,
				AllCost: cost,
			}
		}
	}

	fmt.Println("Generated report")

	for key , report := range goods{

		fmt.Printf("\n\tName : %s\t\tTotal quantity : %d\t\tTotalCost : %.2f\n", key, report.AllQuantity, report.AllCost)
	}
	fmt.Println()
}

func main() {

	path := ""

	if len(os.Args) < 2{
		path = "test.csv"
	} else{
		path = os.Args[1]
	}

	// Проверка на пустой путь
	if path =="" {
		path = "test.csv"
	}

	file := openFile(path)
	defer closeFile(file)

	generateReport(file)

}
