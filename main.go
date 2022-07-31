package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	Id          string `json:"Id"`
	Date        string `json:"Date"`
	Transaction string `json:"Transaction"`
}

func mapCSV() ([]Transaction, error) {
	csvFile, err := os.Open("txns.csv")

	if err != nil {
		return []Transaction{}, err
	}

	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	data, err := reader.ReadAll()

	if err != nil {
		return []Transaction{}, err
	}

	transactions := []Transaction{}

	for i, row := range data {
		if i == 0 {
			continue
		}

		transactions = append(transactions, Transaction{
			Id:          row[0],
			Date:        row[1],
			Transaction: row[2],
		})
	}

	return transactions, nil
}

func calculateBalance(transactions []Transaction) (float64, error) {

	totalBalance := 0.0

	for _, t := range transactions {
		number, err := strconv.ParseFloat(t.Transaction, 64)

		if err != nil {
			return 0.0, err
		}

		totalBalance += number
	}

	return totalBalance, nil
}

func calculateTransactionsPerMonth(transactions []Transaction) (map[string]int, error) {

	transactionsPerMonth := make(map[string]int)

	for _, t := range transactions {
		month := strings.Split(t.Date, "/")[0]
		transactionsPerMonth[month] += 1
	}

	return transactionsPerMonth, nil
}

func calculateAverageDebit(transactions []Transaction) ([]float64, error) {

	creditAmounts := []float64{}
	debitAmounts := []float64{}

	for _, t := range transactions {
		number, err := strconv.ParseFloat(t.Transaction, 64)

		if err != nil {
			return []float64{}, err
		}

		if number > 0 { // POSITIVE => Credit
			creditAmounts = append(creditAmounts, number)
		} else { // NEGATIVE => Debit
			debitAmounts = append(debitAmounts, number)
		}
	}

	totalCredit := 0.0
	for _, amount := range creditAmounts {
		totalCredit = totalCredit + amount
	}
	averageCreditAmounts := totalCredit / float64(len(creditAmounts))

	totalDebit := 0.0
	for _, amount := range debitAmounts {
		totalDebit = totalDebit + amount
	}
	averageDebitAmounts := totalDebit / float64(len(debitAmounts))

	average := make([]float64, 2)
	average[0] = averageCreditAmounts
	average[1] = averageDebitAmounts

	return average, nil
}

func makeEmail(totalBalance float64, transactionsPerMonth map[string]int, averages []float64) {
	months := map[string]string{
		"1":  "January",
		"2":  "February",
		"3":  "March",
		"4":  "April",
		"5":  "May",
		"6":  "June",
		"7":  "July",
		"8":  "August",
		"9":  "September",
		"10": "October",
		"11": "November",
		"12": "December",
	}

	// Email summary
	fmt.Println("Total balance is:", math.Round(totalBalance*100)/100)
	for month, trs := range transactionsPerMonth {
		fmt.Printf("Number of transactions in %s: %d \n", months[month], trs)
	}
	fmt.Println("Average debit amount:", averages[1])
	fmt.Println("Average credit amount:", averages[0])
}

func main() {

	transactions, err := mapCSV()

	if err != nil {
		log.Fatalln(err)
	}

	totalBalance, err := calculateBalance(transactions)

	if err != nil {
		log.Fatalln(err)
	}

	transactionsPerMonth, err := calculateTransactionsPerMonth(transactions)

	if err != nil {
		log.Fatalln(err)
	}

	averages, err := calculateAverageDebit(transactions)

	if err != nil {
		log.Fatalln(err)
	}

	makeEmail(totalBalance, transactionsPerMonth, averages)
}

// TODO: Only one for cycle to do all the calculations
// TODO: Concurrencia al leer el archivo
// TODO: Concurrencia al procesar los datos
