.PHONY: upgrade
upgrade:
	go-mod-upgrade

.PHONY: bench
bench:
	go test -v -run=^$ -bench=Benchmark_ -benchmem -count=2
