upgrade:
	go-mod-upgrade
	go mod tidy

test:
	go test ./...
