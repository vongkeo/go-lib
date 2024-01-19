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
	"os"
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

func SetTimeZone(zone string) bool {
	// set timezone
	if zone == "" {
		zone = "Asia/Bangkok"
	}
	loc, _ := time.LoadLocation(zone)
	time.Local = loc
	os.Setenv("TZ", zone)
	return true
}

func GenerateReqId(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := ""
	for i := 0; i < length; i++ {
		result += string(str[randomInt(0, len(str)-1)])
	}
	return result
}
func GenerateReqTime() string {
	currentTime := time.Now().Format("20060102150405")
	return currentTime

}

func GetLogFileName() (string, error) {
	// UTC +7
	cur := time.Now().UTC().Add(time.Hour * 7)
	// Get the current time in the timezone
	currentDate := cur.Format("2006-01-02")
	logFileName := fmt.Sprintf("logs/%s.log", currentDate)
	println(cur.Format("2006-01-02 15:04:05"))
	return logFileName, nil
}

func GetNow() time.Time {
	// UTC +7
	cur := time.Now().UTC().Add(time.Hour * 7)
	// Get the current time in the timezone
	return cur
}

func IsDate(date string) bool {
	_, err := time.Parse("02-01-2006", date)
	if err != nil {
		return false
	}
	return true
}

func GetLocalDate() time.Time {
	// UTC +7
	cur := time.Now().UTC().Add(time.Hour * 7)
	// Get the current time in the timezone
	return cur
}

func DateFormat(date string) string {
	// date string format to 2006-01-02
	t, _ := time.Parse("2006-01-02", date)
	return t.Format("2006-01-02")

}

func GetYesterday() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	return yesterday.Format("2006-01-02")
}
func Tomorrow() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	t := time.Now().In(loc).AddDate(0, 0, 1).Format("2006-01-02")
	tomorrow, _ := time.ParseInLocation("2006-01-02", t, loc)
	return tomorrow

}
