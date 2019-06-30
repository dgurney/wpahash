package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

const version = "0.0.1"

// Thanks Lucas for the initial C++ implementation and other information!
func calculateEulaHash(wpaHive string, userBytes []byte) (uint64, error) {
	// Debug symbols call the hash EulaHash
	eulaHash := make([]byte, 0x80)
	switch {
	case len(userBytes) > 0:
		// For testing
		if len(userBytes) != 0x80 {
			return 0, fmt.Errorf("the provided EulaHash is not 80 bytes long")
		}
		eulaHash = userBytes
	default:
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, wpaHive+`\WPA\478C035F-04BC-48C7-B324-2462D786DAD7-5P-9`, registry.QUERY_VALUE)
		if err != nil {
			return 0, err
		}
		defer k.Close()
		_, _, err = k.GetValue("", eulaHash)
		if err != nil {
			return 0, err
		}
	}

	hashBytes := make([]byte, 8)
	for i := 0; i < 16; i++ {
		prev := 0x00
		hashBytes[7] ^= eulaHash[i]
		for j := 6; j >= 0; j-- {
			prev += 0x10
			hashBytes[j] ^= eulaHash[prev+i]
		}
	}
	return binary.LittleEndian.Uint64(hashBytes), nil
}

func main() {
	if runtime.GOOS != "windows" {
		fmt.Printf("This program does not do anything meaningful on %s.\n", runtime.GOOS)
		return
	}

	local := flag.Bool("l", false, "Calculate hash of the current installation/PE. Doesn't work on Windows versions before 79xx-era Windows 8 builds, or after Windows 10 1511.")
	hive := flag.String("h", "", "Mounted hive name.")
	ver := flag.Bool("v", false, "Print the program version number and exit")
	flag.Parse()

	if *ver {
		fmt.Printf("wpahash v%s by Daniel Gurney\n", version)
		return
	}

	if *local {
		*hive = "SYSTEM"
	}
	if len(*hive) == 0 {
		fmt.Println("The mounted hive name cannot be empty.\nUsage:")
		flag.PrintDefaults()
		return
	}

	hash, err := calculateEulaHash(*hive, []byte{})
	if err != nil {
		fmt.Println("Hash calculation error:", err)
		return
	}
	fmt.Printf("%x\n", hash)
}
