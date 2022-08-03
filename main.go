package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/ramrodo/stori-technical-challenge/models"
)

var db models.MongoDB

type transactionPerMonth struct {
	Month  string
	Number int
}

type templateData struct {
	TotalBalance         float64
	TransactionsPerMonth []transactionPerMonth
	AverageDebit         float64
	AverageCredit        float64
}

func mapCSV() ([]models.Transaction, error) {
	csvFile, err := os.Open("txns.csv")

	if err != nil {
		return []models.Transaction{}, err
	}

	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	data, err := reader.ReadAll()

	if err != nil {
		return []models.Transaction{}, err
	}

	trs := []models.Transaction{}

	for i, row := range data {
		if i == 0 {
			continue
		}

		trs = append(trs, models.Transaction{
			Id:          row[0],
			Date:        row[1],
			Transaction: row[2],
		})
	}

	return trs, nil
}

func calculateBalance(trs []models.Transaction) (float64, error) {

	totalBalance := 0.0

	for _, t := range trs {
		number, err := strconv.ParseFloat(t.Transaction, 64)

		if err != nil {
			return 0.0, err
		}

		totalBalance += number
	}

	return totalBalance, nil
}

func calculateTransactionsPerMonth(trs []models.Transaction) (map[string]int, error) {

	transactionsPerMonth := make(map[string]int)

	for _, t := range trs {
		month := strings.Split(t.Date, "/")[0]
		transactionsPerMonth[month] += 1
	}

	return transactionsPerMonth, nil
}

func calculateAverages(trs []models.Transaction) ([]float64, error) {

	creditAmounts := []float64{}
	debitAmounts := []float64{}

	for _, t := range trs {
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

func makeEmail(balance float64, trs map[string]int, averages []float64) {
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
	totalBalance := math.Round(balance*100) / 100
	averageDebit := math.Round(averages[1]*100) / 100
	averageCredit := math.Round(averages[0]*100) / 100

	fmt.Println("Total balance is:", totalBalance)
	for month, trs := range trs {
		fmt.Printf("Number of transactions in %s: %d\n", months[month], trs)
	}
	fmt.Println("Average debit amount:", averageDebit)
	fmt.Println("Average credit amount:", averageCredit)

	tmpData := templateData{}
	tmpData.TotalBalance = totalBalance
	tmpData.AverageDebit = averageDebit
	tmpData.AverageCredit = averageCredit

	for month, trs := range trs {
		tr := transactionPerMonth{
			Month:  months[month],
			Number: trs,
		}
		tmpData.TransactionsPerMonth = append(tmpData.TransactionsPerMonth, tr)
	}

	t := template.Must(template.ParseFiles("email-template.html"))

	f, err := os.Create("/template/email.html")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = t.Execute(f, tmpData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email template created:", f.Name())
}

func insertTransactions(trs []models.Transaction) error {
	for _, tr := range trs {
		transactionObjectId, err := db.InsertTransaction(tr)
		if err != nil {
			return err
		}
		fmt.Println("Transaction added:", transactionObjectId)
	}

	return nil
}

func main() {

	var mongoDB models.MongoDB
	db = mongoDB.ConnectDB()
	defer db.CloseDB()

	// Get transactions from CSV file and insert into MongoDB
	trsCSV, err := mapCSV()
	if err != nil {
		log.Fatalln(err)
	}
	err = insertTransactions(trsCSV)
	if err != nil {
		log.Fatalln(err)
	}

	trs, err := db.GetAllTransactions()
	if err != nil {
		log.Fatalln(err)
	}

	totalBalance, err := calculateBalance(trs)
	if err != nil {
		log.Fatalln(err)
	}

	trsPerMonth, err := calculateTransactionsPerMonth(trs)
	if err != nil {
		log.Fatalln(err)
	}

	averages, err := calculateAverages(trs)
	if err != nil {
		log.Fatalln(err)
	}

	makeEmail(totalBalance, trsPerMonth, averages)
}
