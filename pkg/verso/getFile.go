package verso

import (
	"os"
)

func getFile(filepath string) ([]byte, error) {

	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	fb, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	return fb, nil
}
