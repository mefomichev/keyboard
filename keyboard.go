//package keyboar reads user input from keyboard
package keyboard

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

//GetFloat reads a floatingPoint number from keyboard
func GetFloat()(float64, error){
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	score, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return score, nil
}
