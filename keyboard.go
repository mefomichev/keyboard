package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return 0, err
	}
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return grade, nil
}