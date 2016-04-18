package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

type EC2Metadata struct {
	Region,
	SubnetID,
	SubnetCIDR,
	VpcID,
	VpcCIDR string
}

func NewEc2Metadata() (*EC2Metadata, error) {
	service := ec2metadata.New(session.New())

	if !service.Available() {
		return nil, errors.New("EC2 Metadata service is unavailable. Make sure the application is running within an EC2 Instance.")
	}

	mac, _ := service.GetMetadata("network/interfaces/macs/")

	meta := &EC2Metadata{}
	meta.Region, _ = service.Region()
	meta.SubnetID, _ = service.GetMetadata(fmt.Sprintf("network/interfaces/macs/%s/subnet-id", mac))
	meta.SubnetCIDR, _ = service.GetMetadata(fmt.Sprintf("network/interfaces/macs/%s/subnet-ipv4-cidr-block", mac))
	meta.VpcID, _ = service.GetMetadata(fmt.Sprintf("network/interfaces/macs/%s/vpc-id", mac))
	meta.VpcCIDR, _ = service.GetMetadata(fmt.Sprintf("network/interfaces/macs/%s/vpc-ipv4-cidr-block", mac))

	return meta, nil
}
