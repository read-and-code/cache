package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func readLength(reader *bufio.Reader) (int, error) {
	text, err := reader.ReadString(' ')

	if err != nil {
		return 0, err
	}

	length, err := strconv.Atoi(strings.TrimSpace(text))

	if err != nil {
		return 0, err
	}

	return length, nil
}

func sendResponse(value []byte, err error, connection net.Conn) error {
	if err != nil {
		response := fmt.Sprintf("-%d ", len(err.Error())) + err.Error()

		_, err := connection.Write([]byte(response))

		return err
	}

	valueLength := fmt.Sprintf("%d ", len(value))

	_, err = connection.Write(append([]byte(valueLength), value...))

	return err
}
