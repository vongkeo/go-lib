/**
 * Author: Vongkeo KEOSAVANH
 * File: go-lib.go
 */

package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Contains(lists string, fileEx string) bool {
	//  split string to array
	list := strings.Split(lists, ",")
	// check file extension
	for _, v := range list {
		if v == fileEx {
			return true
		}
	}
	return false

}

func IsMobileNo(mobileNo string) bool {
	// mobile no must start with 20 and length 10
	if len(mobileNo) != 10 || !strings.HasPrefix(mobileNo, "20") {
		return false
	}
	return true

}

func GetTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05.000000")

}

func DateStartZone(date string, zone string) time.Time {
	// 2024-01-18 00:00:00
	loc, _ := time.LoadLocation(zone)
	t, _ := time.ParseInLocation("2006-01-02", date, loc)
	return t

}

func DateEndZone(date string, zone string) time.Time {
	// 2024-01-18 23:59:59
	loc, _ := time.LoadLocation(zone)
	t, _ := time.ParseInLocation("2006-01-02", date, loc)
	return t.Add(time.Hour*24 - time.Nanosecond)
}

func GetDate(zone string) time.Time {
	loc, _ := time.LoadLocation(zone)
	now := time.Now().In(loc).Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", now, loc)
	return t

}

func GetDateDiff(start_date string, end_date string) int {
	// start_date and end_date must be between 6 months
	startDate, _ := time.Parse("2006-01-02", start_date)
	endDate, _ := time.Parse("2006-01-02", end_date)
	diff := endDate.Sub(startDate).Hours() / 24
	return int(diff)
}

func GetMonthDiff(start_date string, end_date string) int {
	// start_date and end_date must be between 6 months
	startDate, _ := time.Parse("2006-01-02", start_date)
	endDate, _ := time.Parse("2006-01-02", end_date)
	diff := endDate.Sub(startDate).Hours() / 24 / 30
	return int(diff)
}

func StringToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func StringToFloat(str string) float64 {
	num, _ := strconv.ParseFloat(str, 64)
	return num
}

func GenerateReference(length int) string {
	code := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := ""
	for i := 0; i < length; i++ {
		result += string(code[randomInt(0, len(code)-1)])
	}
	return result
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func InArray(val string, array []string) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}

func StructToJSON(model interface{}) (string, error) {
	// StructToJSON
	jsonData, err := json.Marshal(model)
	if err != nil {
		fmt.Println("Error:", err)
		return "", fmt.Errorf("Error: %s", err)
	}
	// Convert the JSON byte slice to a string for printing
	jsonString := string(jsonData)
	return jsonString, nil
}

// LogWithType logs a message with a UUID and log type (default: "info")
func Log(requestID any, messages ...interface{}) {
	labelType := []string{"info", "error", "debug", "warn", "fatal", "panic"}
	label := "info"

	// type should be in labelType else default is info
	if len(messages) > 1 {
		label = fmt.Sprint(messages[0])
		if !InArray(label, labelType) {
			label = "info"
		}
	}

	log.Printf("[%s] [%s] %s", requestID, label, fmt.Sprint(messages...))
}
