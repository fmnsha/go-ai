package util

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"math/rand"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Paginate(r *http.Request) (sk int64, lim int64, er error) {
	page := r.URL.Query().Get("page")
	perPage := r.URL.Query().Get("perPage")

	var skip, limit int64
	var convErr error
	if perPage == "" {
		limit = 40
	} else {
		limit, convErr = strconv.ParseInt(perPage, 10, 64)
		if convErr != nil {
			return 0, 0, convErr
		}
	}
	if page == "" {
		skip = 0
	} else {
		skip, convErr = strconv.ParseInt(page, 10, 64)
		if convErr != nil {
			return 0, 0, convErr
		}
		skip = (skip - 1) * limit
	}

	return skip, limit, nil
}

func InSlice[T comparable](ele T, s []T) bool {
	for _, e := range s {
		if e == ele {
			return true
		}
	}

	return false

}

func GeneratePassword(passWordLength, minSpecialChar, minUpper, minNum int) string {
	var (
		lowerCharSet   = "abcdedfghijklmnopqrst"
		upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		specialCharSet = "!@#$%&*"
		numberSet      = "0123456789"
		allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
	)
	var password strings.Builder
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < minSpecialChar; i++ {
		index := r.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[index]))
	}

	for i := 0; i < minUpper; i++ {
		index := r.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[index]))
	}

	for i := 0; i < minNum; i++ {
		index := r.Intn(len(numberSet))
		password.WriteString(string(numberSet[index]))
	}

	remainLength := passWordLength - minSpecialChar - minUpper - minNum

	for i := 0; i < remainLength; i++ {
		index := r.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[index]))
	}

	return password.String()

}

func Contains[T comparable](arr []T, key T) bool {
	for _, value := range arr {
		if value == key {
			return true
		}
	}
	return false
}

func SliceFilter[T any](arr []T, callBack func(el T) bool) []T {
	var result []T
	for _, e := range arr {
		ok := callBack(e)
		if ok {
			result = append(result, e)
		}
	}

	return result
}

func SliceFind[T any](arr []T, callback func(el T) bool) (T, bool) {
	var result T
	var ok bool
	for _, e := range arr {
		ok = callback(e)
		if ok {
			result = e
			break
		}
	}

	return result, ok
}

func RoundToDecimalPlaces(value float64, decimalPlaces int) float64 {
	factor := math.Pow(10, float64(decimalPlaces))
	return math.Round(value*factor) / factor
}

func IsValidEmail(email string) bool {
	// Define the regex pattern for a valid email address
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex pattern
	re := regexp.MustCompile(emailPattern)

	// Match the email address against the pattern
	return re.MatchString(email)
}

func GenerateUniqueString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func InsertAt[T any](slice []T, index int, element T) ([]T, error) {
	// Validate the index
	if index < 0 || index > len(slice) {
		return nil, fmt.Errorf("index out of range: %d", index)
	}

	// Insert the element at the specified index
	newSlice := append(slice[:index], append([]T{element}, slice[index:]...)...)

	return newSlice, nil
}
