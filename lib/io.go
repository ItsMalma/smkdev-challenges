package lib

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Input(text string) string {
	fmt.Print(text)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	return scanner.Text()
}

var ErrInvalidInputInt = errors.New("inputed value is not valid integer")

func InputInt(text string) (int64, error) {
	value, err := strconv.ParseInt(Input(text), 10, 64)
	if err != nil {
		return 0, ErrInvalidInputInt
	}
	return value, nil
}
