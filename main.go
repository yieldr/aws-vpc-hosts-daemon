package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	metadata, err := NewEc2Metadata()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	ec2 := NewEc2(metadata.Region)
	desciption, err := ec2.DescribeInstances(map[string][]string{
		"vpc-id": []string{metadata.VpcID},
	})
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	for _, reservation := range desciption.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("%s\t%s\n", *instance.PrivateIpAddress, *instance.PrivateDnsName)
		}
	}
}
