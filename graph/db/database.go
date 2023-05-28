package db

import (
	"fmt"
	"os"
)

func DatabaseConn() string {
	return fmt.Sprintf("postgres://%s", os.Getenv("DATA_SOURCE"))
}
