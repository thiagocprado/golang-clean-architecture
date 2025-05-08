package app

import (
	"math"
	"strconv"
	"strings"
)

func FindInArray(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}

func FindInArrayInt(values string, value int) bool {
	array := strings.Split(values, ",")

	for _, v := range array {
		vInt, _ := strconv.Atoi(v)

		if vInt == value {
			return true
		}
	}

	return false
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func RoundFloatv2(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func RoundFloatv3(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func RoundFloatv4(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
