package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadIdAndSecret(path string) (id, secret string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		if strings.HasPrefix(line, "ID:") {
			id = strings.TrimSpace(strings.TrimPrefix(line, "ID:"))
		} else {
			secret = strings.TrimSpace(strings.TrimPrefix(line, "Secret:"))
		}
	}
	return
}
