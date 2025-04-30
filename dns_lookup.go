package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func lookupARecords(hostname string) ([]net.IP, error) {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return nil, fmt.Errorf("error looking up A/AAAA records for %s: %w", hostname, err)
	}
	return ips, nil
}

func lookupPTRRecord(ipAddress string) ([]string, error) {
	names, err := net.LookupAddr(ipAddress)
	if err != nil {
		return nil, fmt.Errorf("error performing PTR (Reverse DNS) lookup for %s: %w", ipAddress, err)
	}
	var cleanedNames []string
	for _, name := range names {
		cleanedNames = append(cleanedNames, trimTrailingDot(name))
	}
	return cleanedNames, nil
}

func lookupMXRecords(hostname string) ([]*net.MX, error) {
	mxs, err := net.LookupMX(hostname)
	if err != nil {
		return nil, fmt.Errorf("error looking up MX records for %s: %w", hostname, err)
	}
	return mxs, nil
}

func lookupNSRecords(hostname string) ([]*net.NS, error) {
	nss, err := net.LookupNS(hostname)
	if err != nil {
		return nil, fmt.Errorf("error looking up NS records for %s: %w", hostname, err)
	}
	return nss, nil
}

func lookupTXTRecords(hostname string) ([]string, error) {
	txts, err := net.LookupTXT(hostname)
	if err != nil {
		return nil, fmt.Errorf("error looking up TXT records for %s: %w", hostname, err)
	}
	return txts, nil
}

// Helper function to trim the trailing dot from a DNS name
func trimTrailingDot(hostname string) string {
	if len(hostname) > 0 && hostname[len(hostname)-1] == '.' {
		return hostname[:len(hostname)-1]
	}
	return hostname
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <target> <record_type>")
		fmt.Println("Supported record types: A, PTR, MX, NS, TXT")
		os.Exit(1)
	}

	target := os.Args[1]
	recordType := os.Args[2]

	fmt.Printf("Looking up %s records for: %s\n", recordType, target)
	fmt.Println("------------------------------------")

	switch recordType {
	case "A":
		ips, err := lookupARecords(target)
		if err != nil {
			log.Printf("Error: %v\n", err)
		} else if len(ips) > 0 {
			fmt.Println("A/AAAA Records:")
			for _, ip := range ips {
				fmt.Printf("  - %s\n", ip)
			}
		} else {
			fmt.Println("No A/AAAA records found.")
		}

	case "PTR":
		names, err := lookupPTRRecord(target)
		if err != nil {
			log.Printf("Error: %v\n", err)
		} else if len(names) > 0 {
			fmt.Println("PTR (Reverse DNS) Records:")
			for _, name := range names {
				fmt.Printf("  - %s\n", name)
			}
		} else {
			fmt.Println("No PTR records found.")
		}

	case "MX":
		mxs, err := lookupMXRecords(target)
		if err != nil {
			log.Printf("Error: %v\n", err)
		} else if len(mxs) > 0 {
			fmt.Println("MX (Mail Server) Records:")
			for _, mx := range mxs {
				fmt.Printf("  - Host: %s, Preference: %d\n", mx.Host, mx.Pref)
			}
		} else {
			fmt.Println("No MX records found.")
		}

	case "NS":
		nss, err := lookupNSRecords(target)
		if err != nil {
			log.Printf("Error: %v\n", err)
		} else if len(nss) > 0 {
			fmt.Println("NS (Name Server) Records:")
			for _, ns := range nss {
				fmt.Printf("  - %s\n", ns.Host)
			}
		} else {
			fmt.Println("No NS records found.")
		}

	case "TXT":
		txts, err := lookupTXTRecords(target)
		if err != nil {
			log.Printf("Error: %v\n", err)
		} else if len(txts) > 0 {
			fmt.Println("TXT Records:")
			for _, txt := range txts {
				fmt.Printf("  - %s\n", txt)
			}
		} else {
			fmt.Println("No TXT records found.")
		}

	default:
		fmt.Printf("Unsupported record type: %s\n", recordType)
		fmt.Println("Supported record types: A, PTR, MX, NS, TXT")
		os.Exit(1)
	}

	fmt.Println("------------------------------------")
}
