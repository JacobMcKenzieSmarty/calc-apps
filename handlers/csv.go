package handler

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/JacobMcKenzieSmarty/calc-lib/calc"
)

type CSVHandler struct {
	reader *csv.Reader
	writer *csv.Writer
	logger *log.Logger
}

func NewCSVHandler(reader io.Reader, writer io.Writer, logger *log.Logger) *CSVHandler {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1 //ignore field number mismatch

	return &CSVHandler{
		reader: csvReader,
		writer: csv.NewWriter(writer),
		logger: logger,
	}
}

func (this *CSVHandler) Handle() error {
	for {
		record, err := this.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read error: %w", err)
		}

		op1, err := strconv.Atoi(record[0])
		if err != nil {
			this.logger.Println(err)
			continue
		}
		operator := record[1]
		op2, err := strconv.Atoi(record[2])
		if err != nil {
			this.logger.Println(err)
			continue
		}
		myCalc := calculators[operator]
		if myCalc == nil {
			this.logger.Println("Unknown operator")
			continue
		}
		result := myCalc.Calculate(op1, op2)

		err = this.writer.Write(append(record, strconv.Itoa(result)))
		if err != nil {
			this.logger.Println(err)
			break //really difficult to cover in unit test because you'd have to be exactly 4096
		}

	}

	this.writer.Flush()
	return this.writer.Error()
}

var calculators = map[string]calc.Calculator{
	"+": calc.Addition{},
	"-": calc.Subtraction{},
	"*": calc.Multiplication{},
	"/": calc.Division{},
}
