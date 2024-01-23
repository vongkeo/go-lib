/**
 * Author: Vongkeo KEOSAVANH
 * File: go-lib.go
 */

package lib

import (
	"crypto"
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	mathRand "math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
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
	return min + mathRand.Intn(max-min+1)
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

func StringToTime(str string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, str)
	return t
}

func setKey(privateKey string) ([]byte, error) {
	key := []byte(privateKey)
	sha := sha1.New()
	sha.Write(key)
	hashedKey := sha.Sum(nil)
	derivedKey := hashedKey[:16]
	return derivedKey, nil
}

func encrypt(dataString, privateKey string) (string, error) {
	derivedKey, err := setKey(privateKey)
	if err != nil {
		return "", err
	}

	data := []byte(dataString)
	paddedData := padData(data, aes.BlockSize)

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, len(paddedData))
	for i := 0; i < len(paddedData); i += aes.BlockSize {
		block.Encrypt(cipherText[i:i+aes.BlockSize], paddedData[i:i+aes.BlockSize])
	}

	encryptedData := base64.StdEncoding.EncodeToString(cipherText)
	return encryptedData, nil
}

func decrypt(encryptedData, privateKey string) (string, error) {
	derivedKey, err := setKey(privateKey)
	if err != nil {
		return "", err
	}

	decodedData, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return "", err
	}

	decryptedData := make([]byte, len(decodedData))
	for i := 0; i < len(decodedData); i += aes.BlockSize {
		block.Decrypt(decryptedData[i:i+aes.BlockSize], decodedData[i:i+aes.BlockSize])
	}

	unpaddedData := unpadData(decryptedData)

	return string(unpaddedData), nil
}
func padData(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	pad := byte(padding)
	paddedData := append(data, pad)

	for i := 1; i < padding; i++ {
		paddedData = append(paddedData, pad)
	}

	return paddedData
}

func unpadData(data []byte) []byte {
	length := len(data)
	unpad := int(data[length-1])

	return data[:length-unpad]
}

func Encrypt(dataString string, privateKey string) (string, error) {
	encryptedData, err := encrypt(dataString, privateKey)
	if err != nil {
		fmt.Println("Error while encrypting:", err)
		return "", err
	}
	return encryptedData, nil

}

func Decrypt(encryptedData string, privateKey string) (string, error) {
	decryptedData, err := decrypt(encryptedData, privateKey)
	if err != nil {
		fmt.Println("Error while decrypting:", err)
		return "", err
	}
	return decryptedData, nil
}

func RsaSignPri(data interface{}, priKey string) (string, error) {
	// data interface to byte
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return "", fmt.Errorf("failed to marshal json")
	}
	// hash data
	hashed := sha256.Sum256([]byte(jsonData))
	// convert private key string to private key
	privateKey, err := convertPrivateKeyString(priKey)
	if err != nil {
		fmt.Println("Failed to convert private key:", err)
		return "", fmt.Errorf("failed to convert private key")
	}
	// sign data
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println("Failed to sign data:", err)
		return "", fmt.Errorf("failed to sign data")
	}
	// convert byte to hex
	signatureHex := hex.EncodeToString(signature)
	// return signature
	return signatureHex, nil
}

func convertPrivateKeyString(privateKeyString string) (*rsa.PrivateKey, error) {
	// Add PEM headers and footers to the private key string
	pemKey := fmt.Sprintf("-----BEGIN PRIVATE KEY-----\n%s\n-----END PRIVATE KEY-----", privateKeyString)

	// Decode the PEM-encoded private key
	privateKeyPEM, _ := pem.Decode([]byte(pemKey))
	if privateKeyPEM == nil {
		return nil, fmt.Errorf("failed to decode PEM block: invalid format")
	}

	// Parse the DER-encoded private key
	parsedKey, err := x509.ParsePKCS8PrivateKey(privateKeyPEM.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PKCS8 private key: %w", err)
	}

	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("unexpected key type")
	}

	return privateKey, nil
}

func DownloadFile(url string, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// read html file
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(rootPath)
	// Create the file
	out, err := os.Create(rootPath + filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func GetExtension(url string) string {
	// get extension
	extension := ""
	for i := len(url) - 1; i >= 0; i-- {
		if string(url[i]) == "." {
			break
		}
		extension = string(url[i]) + extension
	}
	return extension
}

func RenameFile(oldPath string, newPath string) error {
	// rename file
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	// rename
	err = os.Rename(rootPath+oldPath, rootPath+newPath)
	if err != nil {
		return err
	}
	return nil
}

func MoveFile(oldPath string, newPath string) error {
	// move file
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	// move
	err = os.Rename(rootPath+oldPath, rootPath+newPath)
	if err != nil {
		return err
	}
	return nil
}

func ExcelToJson(xlFile *xlsx.File) ([]map[string]interface{}, error) {
	var jsonData []map[string]interface{}
	// Loop through each of the sheets in the spreadsheet
	for _, sheet := range xlFile.Sheets {
		// Loop through each row in the sheet
		for _, row := range sheet.Rows {
			// Create a map to hold our row data
			rowData := make(map[string]interface{})

			// Loop through each cell in the row
			for columnIndex, cell := range row.Cells {
				// Get the value for the cell
				text := cell.String()

				// Use the column header as the key and the cell value as the value
				columnHeader := sheet.Rows[0].Cells[columnIndex].String()
				rowData[columnHeader] = text
			}

			// Add our Row Data to the jsonData
			jsonData = append(jsonData, rowData)
		}
	}
	// Marshal the data into JSON
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON data: %s", err)
	}
	// Unmarshal JSON data into a slice of maps
	var jsonSlice []map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonSlice)
	if err != nil {
		return nil, fmt.Errorf("Error un marshaling JSON data: %s", err)
	}
	return jsonSlice, nil

}
