out_exec.so: out_exec.go newrelic.go
	go build -buildmode=c-shared -o out_exec.so out_exec.go newrelic.go

.PHONY: test
test: out_exec.so
	fluent-bit -e out_exec.so -i dummy --prop="rate=3" -o exec_out --prop="command=python3 example/script.py"

.PHONY: list
list: out_exec.so
	fluent-bit -e out_exec.so --help
