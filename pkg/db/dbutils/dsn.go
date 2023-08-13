package dbutils

import "fmt"

func PostgresDSN(addr, user, password, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		user, password, addr, dbname,
	)
}
