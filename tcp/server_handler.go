package tcp

import (
	"bufio"
	"io"
	"log"
	"net"
)

func (server *Server) readKey(reader *bufio.Reader) (string, error) {
	keyLength, err := readLength(reader)
	
	if err != nil {
		return "", err
	}
	
	key := make([]byte, keyLength)
	_, err = io.ReadFull(reader, key)
	
	if err != nil {
		return "", err
	}
	
	return string(key), nil
}

func (server *Server) readKeyAndValue(reader *bufio.Reader) (string, []byte, error) {
	keyLength, err := readLength(reader)
	
	if err != nil {
		return "", nil, err
	}
	
	valueLength, err := readLength(reader)
	
	if err != nil {
		return "", nil, err
	}
	
	key := make([]byte, keyLength)
	_, err = io.ReadFull(reader, key)
	
	if err != nil {
		return "", nil, err
	}
	
	value := make([]byte, valueLength)
	_, err = io.ReadFull(reader, value)

	if err != nil {
		return "", nil, err
	}

	return string(key), value, nil
}

func (server *Server) get(connection net.Conn, reader *bufio.Reader) error {
	key, err := server.readKey(reader)

	if err != nil {
		return err
	}

	value, err := server.cache.Get(key)

	return sendResponse(value, err, connection)
}

func (server *Server) set(connection net.Conn, reader *bufio.Reader) error {
	key, value, err := server.readKeyAndValue(reader)

	if err != nil {
		return err
	}

	return sendResponse(nil, server.cache.Set(key, value), connection)
}

func (server *Server) delete(connection net.Conn, reader *bufio.Reader) error {
	key, err := server.readKey(reader)

	if err != nil {
		return err
	}

	return sendResponse(nil, server.cache.Delete(key), connection)
}

func (server *Server) process(connection net.Conn) {
	defer connection.Close()

	reader := bufio.NewReader(connection)

	for {
		operation, err := reader.ReadByte()

		if err != nil {
			if err != io.EOF {
				log.Println("Close connection due to error:", err)
			}

			return
		}

		if operation == 'S' {
			err = server.set(connection, reader)
		} else if operation == 'G' {
			err = server.get(connection, reader)
		} else if operation == 'D' {
			err = server.delete(connection, reader)
		} else {
			log.Println("Close connection due to invalid operation:", operation)
		}

		if err != nil {
			log.Println("Close connection due to error:", err)

			return
		}
	}
}

