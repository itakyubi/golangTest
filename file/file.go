package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func AppendDataToCSV(context []string, dir string, fileName string) error {
	path := dir + "/" + fileName

	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	var file *os.File
	_, err = os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create(path)
	} else {
		file, err = os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0666)
	}

	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	err = w.Write(context)
	w.Flush()

	return err
}

func CountBySecond(path string) {
	file, _ := os.Open(path)
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	flag := uint64(0)
	count := 0
	index := 0
	totalSeconds := 0
	for _, row := range content {
		ts, _ := strconv.ParseUint(row[0], 0, 64)
		sd := ts / 1e3

		if index == 0 {
			flag = sd
		}

		if flag == sd {
			count++
		} else {
			fmt.Println("Second:", flag, ", FPS:", count)
			count = 1
			flag = sd
			totalSeconds++
		}
		index++
	}
	fmt.Println("Total seconds:", totalSeconds)
}

func CountBySecond2(path string) {
	file, _ := os.Open(path)
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	flag := uint64(0)
	count := 0
	index := 0
	totalSeconds := 0
	total := 0
	for _, row := range content {
		ts, _ := strconv.ParseUint(row[0], 0, 64)
		sd := ts / 1e3

		if index == 0 {
			flag = sd
		}

		if flag == sd {
			count++
		} else {
			if sd-flag > 1 {
				total++
				tm := time.Unix(int64(sd), 0)
				fmt.Println("Time:", tm.Format("2006-01-02 15:04:05"), ",Second:", sd, ", interval:", sd-flag)
			}
			count = 1
			flag = sd
			totalSeconds++
		}
		index++
	}
	fmt.Println("Total:", total)
}

func CountBySecond3(path string) {
	file, _ := os.Open(path)
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	flag := uint64(0)
	count := 0
	index := 0
	totalSeconds := 0
	for _, row := range content {
		ts, _ := strconv.ParseUint(row[0], 0, 64)
		sd := ts / 1e3

		if index == 0 {
			flag = sd
		}

		if flag == sd {
			count++
		} else {
			if count < 8 {
				fmt.Println("Second:", flag, ", FPS:", count)
			}
			count = 1
			flag = sd
			totalSeconds++
		}
		index++
	}
	fmt.Println("Total seconds:", totalSeconds)
}
