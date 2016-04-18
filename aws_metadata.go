package main

import (
	"errors"

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

	meta := &EC2Metadata{}
	meta.Region, _ = service.Region()
	meta.SubnetID, _ = service.GetMetadata("network/interfaces/macs/mac/subnet-id")
	meta.SubnetCIDR, _ = service.GetMetadata("network/interfaces/macs/mac/subnet-ipv4-cidr-block")
	meta.VpcID, _ = service.GetMetadata("network/interfaces/macs/mac/vpc-id")
	meta.VpcCIDR, _ = service.GetMetadata("network/interfaces/macs/mac/vpc-ipv4-cidr-block")

	return meta, nil
}
