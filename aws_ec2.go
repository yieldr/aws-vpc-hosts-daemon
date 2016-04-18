package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Ec2 struct {
	Client *ec2.EC2
}

func (e *Ec2) DescribeInstances(filters map[string][]string) (*ec2.DescribeInstancesOutput, error) {
	input := &ec2.DescribeInstancesInput{}
	input.Filters = make([]*ec2.Filter, 0, len(filters))

	for name, values := range filters {
		input.Filters = append(input.Filters, &ec2.Filter{
			Name:   aws.String(name),
			Values: aws.StringSlice(values),
		})
	}

	return e.Client.DescribeInstances(input)
}

func NewEc2(region string) *Ec2 {
	return &Ec2{
		Client: ec2.New(session.New(), &aws.Config{Region: aws.String(region)}),
	}
}
