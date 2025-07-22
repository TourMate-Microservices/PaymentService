package repository

import (
	"fmt"
	"math"
)

// Caculate the offset number of records from a table in database
func getOffSetAmount(limitAmount, pageNumber int) int {
	return (pageNumber - 1) * limitAmount
}

// Caculate total pages
func caculateTotalPages(records, limitAmount int) int {
	return int(math.Ceil(float64(records) / float64(limitAmount)))
}

// Generate query to count the number of all records from a table from database
func generateCountTotalRecordsQuery(table, condition string) string {
	return "SELECT COUNT(*) FROM " + table + condition
}

// Generate query based on demand
func generateRetrieveQuery(table, condition string, limitAmount, pageNumber int, isGetCount bool) string {
	if isGetCount {
		return "SELECT COUNT(*) FROM " + table + condition
	}

	return fmt.Sprintf("SELECT * FROM "+table+" LIMIT %d + OFFSET %d"+condition, limitAmount, getOffSetAmount(limitAmount, pageNumber))
}

// Generate order based on demand
func generateOrderCondition(filterProb, order string) string {
	return " ORDER BY " + filterProb + " " + order
}
