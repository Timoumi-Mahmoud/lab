package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// dependency injection involves 4 roles:
// 1. the service object(s) to be userd - also known as "dependency"
// 2. the client object: that is depending on the service(s) it uses
// 3. the interfaces: that define how the client may use the services
// 4. the injector: which is responsible for constructing the services and injection them into the client
//

var (
	errUsage = errors.New(`Usage:
	set <key> <value> Set specified key and value
	get <key>	 Get specified key	`)
)

func main() {
	runner := newRunner()
	if err := runner.run(); err != nil {
		fmt.Println(err)
	}

}

type runner struct {
	database fileDatabase
}

func newRunner() runner {
	return runner{newFileDatabase()}
}

func (r runner) run() error {
	args := os.Args
	if len(args) < 3 {
		return errUsage
	}
	switch args[1] {
	case "set":
		if len(args) < 4 {
			return errUsage
		}
		if err := r.database.Set(args[2], args[3]); err != nil {
			return err
		}
		log.Println("SET successfully ")
	case "get":
		v, err := r.database.Get(args[2])
		if err != nil {
			return err
		}
		fmt.Println(v)
	default:
		return errUsage
	}
	return nil
}

type fileDatabase struct {
	file string
}

func newFileDatabase() fileDatabase {
	return fileDatabase{"database.txt"}
}

func (db fileDatabase) Set(key, value string) error {
	f, err := os.OpenFile(db.file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%s:%s\n", key, value)); err != nil {
		return err
	}
	return nil

}
func (db fileDatabase) Get(key string) (string, error) {
	f, err := os.OpenFile(db.file, os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var last string
	for scanner.Scan() {
		row := scanner.Text()
		parts := strings.Split(row, ":")
		if len(parts) < 2 {
			return "", errors.New("invalid record")
		}
		if parts[0] == key {
			last = parts[1]
		}
	}
	if last != "" {
		return last, nil
	}
	return "", errors.New("not found")
}
