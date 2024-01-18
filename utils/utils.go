package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadDBCredentials(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pair := strings.SplitN(scanner.Text(), "=", 2)
		if len(pair) == 2 {
			os.Setenv(pair[0], pair[1])
		}
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	return nil
}

func GetUserName() string {
	return os.Getenv("DB_USERNAME")
}

func GetPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}