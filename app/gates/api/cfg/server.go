package cfg

import "os"

func GetAddr() string {
	addr, ok := os.LookupEnv("API_GATE_ADDR")
	if !ok {
		return "localhost:4000" // Default address.
	}
	return addr
}
