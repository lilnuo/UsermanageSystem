package dao

import (
	"os"
	"strings"
)

var database = make(map[string]string)

const dataFile = "user.txt"

func init() {
	loadUsersFormFile()
}
func loadUsersFormFile() {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			saveUsersToFile()
			return
		}
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			database[parts[0]] = parts[1]
		}
	}

}

func saveUsersToFile() {
	var lines []string
	for username, password := range database {
		lines = append(lines, username+":"+password)
	}
	err := os.WriteFile(dataFile, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return
	}

}
func AddUser(username, password string) {
	database[username] = password
	saveUsersToFile()
}
func FindUser(username string, password string) bool {
	if pwd, ok := database[username]; ok {
		if pwd == password {
			return true
		}
	}
	return false
}
func ModifyPassword(username string, password string) string {
	database[username] = password
	saveUsersToFile()
}
