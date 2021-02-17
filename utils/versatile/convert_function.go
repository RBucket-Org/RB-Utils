package versatile

import (
	"strconv"
	"strings"
)

//ConvertSLiceToString : this method convert the slice into the string
func ConvertSLiceToString(input []string) string {
	output := strings.Join(input, ",")
	return output
}

//ConvertStringToSlice : this method convert the string into the slice
func ConvertStringToSlice(input string) []string {
	output := strings.Split(input, ",")
	return output
}

//ConvertIntSliceToString : this method convert the intslice into string
func ConvertIntSliceToString(input []int) string {
	//input as string
	inputAsString := []string{}
	for _, v := range input {
		value := v //value of the int slice
		valueAsString := strconv.Itoa(value)
		inputAsString = append(inputAsString, valueAsString)
	}
	//join to make it as a single string
	output := strings.Join(inputAsString, ",")
	return output
}

//ConvertStringToIntSlice : this method convert the string into intslice
func ConvertStringToIntSlice(input string) ([]int, error) {
	sliceOfInt := []int{}

	output := strings.Split(input, ",")
	for _, v := range output {
		value := v //value
		valueAsInt, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		sliceOfInt = append(sliceOfInt, valueAsInt)
	}
	return sliceOfInt, nil
}
