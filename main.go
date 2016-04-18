package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
)

var args struct {
	hostsFile,
	hostsFileOrig string
}

func init() {
	flag.StringVar(&args.hostsFile, "hosts-file", "/etc/hosts", "Path to hosts file.")
	flag.StringVar(&args.hostsFileOrig, "original-hosts-file", "/etc/hosts.orig", "Store a backup the original hosts file to this location.")
	flag.Parse()
}

func main() {
	metadata, err := NewEc2Metadata()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	var buffer bytes.Buffer

	ec2 := NewEc2(metadata.Region)
	desciption, err := ec2.DescribeInstances(map[string][]string{"vpc-id": []string{metadata.VpcID}})
	for _, reservation := range desciption.Reservations {
		for _, instance := range reservation.Instances {
			buffer.WriteString(*instance.PrivateIpAddress)
			buffer.WriteByte('\t')
			buffer.WriteString(*instance.PrivateDnsName)
			buffer.WriteByte('\n')
		}
	}

	if !exists(args.hostsFileOrig) {
		err := copyFile(args.hostsFile, args.hostsFileOrig)
		if err != nil {
			log.Printf("Failed to backup file: %s", err)
			os.Exit(2)
		}
	}

	err = copyFile(args.hostsFileOrig, args.hostsFile)
	if err != nil {
		log.Printf("Failed to restore file: %s", err)
		os.Exit(2)
	}

	err = appendFile(args.hostsFile, buffer.Bytes())
	if err != nil {
		log.Printf("Failed to append file: %s", err)
		os.Exit(2)
	}

}

func exists(file string) bool {
	_, err := os.Open(file)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return false
}

func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}
	return destFile.Close()
}

func appendFile(file string, data []byte) error {
	f, err := os.OpenFile(file, os.O_APPEND, 0640)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	return err
}
