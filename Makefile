out_exec.so: out_exec.go
	go build -buildmode=c-shared -o out_exec.so out_exec.go

.PHONY: test
test: out_exec.so
	fluent-bit -e out_exec.so -i dummy -o exec_out

.PHONY: list
list: out_exec.so
	fluent-bit -e out_exec.so --help
