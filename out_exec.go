package main

import (
	"C"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strings"
	"unsafe"

	"github.com/fluent/fluent-bit-go/output"
)

//export FLBPluginRegister
func FLBPluginRegister(def unsafe.Pointer) int {
	// Gets called only once when the plugin.so is loaded
	return output.FLBPluginRegister(def, "exec_out", "Process using external script")
}

//export FLBPluginInit
func FLBPluginInit(plugin unsafe.Pointer) int {
	command := output.FLBPluginConfigKey(plugin, "command")
	log.Printf("[exec_out] command = %q", command)
	output.FLBPluginSetContext(plugin, command)
	return output.FLB_OK
}

//export FLBPluginFlushCtx
func FLBPluginFlushCtx(ctx, data unsafe.Pointer, length C.int, tag *C.char) int {
	// Gets called with a batch of records to be written to an instance.
	command := output.FLBPluginGetContext(ctx).(string)
	v := strings.Split(command, " ")
	cmd := exec.Command(v[0], v[1:]...)
	cmd.Stderr = os.Stderr
	log.Printf("[exec_out] command = %q", cmd)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Printf("[exec_out] %s", err)
		return output.FLB_ERROR
	}

	// Process Data
	dec := output.NewDecoder(data, int(length))

	for {
		// Pull in our record
		ret, _, record := output.GetRecord(dec)
		if ret != 0 {
			break
		}

		// Convert to json
		parsedRecord := ParseRecord(record)
		line, err := json.Marshal(parsedRecord)
		if err != nil {
			log.Printf("[exec_out] %s", err)
			continue
		}

		// Write to stdin
		stdin.Write(line)
		stdin.Write([]byte("\n"))
	}

	stdin.Close()

	err = cmd.Run()

	if err != nil {
		log.Printf("[exec_out] %s", err)
		return output.FLB_ERROR
	}

	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	log.Println("[exec_out] Exit")
	return output.FLB_OK
}

func main() {

}
