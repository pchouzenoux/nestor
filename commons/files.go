package commons

import (
	"io/ioutil"
)

func ReadFile(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func ToFile(filepath string, filecontent []byte) error {
	err := ioutil.WriteFile(filepath, filecontent, 0644)
	return err
}

func AppendToFile(filepath string, filecontent []byte) error {
	err := ioutil.WriteFile(filepath, filecontent, 0644)
	return err
}
