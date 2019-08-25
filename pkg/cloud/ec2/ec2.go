package ec2

import (
	"context"
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/pkg/models"

	logger "github.com/theonlyjohnny/phoenix/internal/log"
)

var (
	log               logger.Logger
	clusterNameRegexp = regexp.MustCompile(`(?:us|ap|sa|eu)(?:n|e|s|w|c)+[0-9]-([a-z]*)-[0-9]*`)
)

func init() {
	log = logger.Log
}

//EC2 Cloud
type EC2 struct {
	client *ec2.EC2
	ctx    context.Context
}

func NewEC2CloudProvider(cfg config.ComponentConfig) (EC2, error) {
	var e EC2
	errs := make([]error, 3)

	awsID, err := cfg.GetStr("AWS_ACCESS_KEY_ID")
	if err != nil {
		errs = append(errs, err)
	} else if awsID == "" {
		errs = append(errs, fmt.Errorf("AWS_ACCESS_KEY_ID is a required parameter"))
	}

	awsSecret, err := cfg.GetStr("AWS_SECRET_ACCESS_KEY")
	if err != nil {
		errs = append(errs, err)
	} else if awsID == "" {
		errs = append(errs, fmt.Errorf("AWS_SECRET_ACCESS_KEY is a required parameter"))
	}

	awsRegion, err := cfg.GetStr("AWS_REGION")
	if err != nil {
		errs = append(errs, err)
	} else if awsID == "" {
		errs = append(errs, fmt.Errorf("AWS_REGION is a required parameter"))
	}

	for _, er := range errs {
		if er != nil {
			return e, er
		}
	}

	creds := credentials.NewStaticCredentials(awsID, awsSecret, "")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: creds,
	})
	if err != nil {
		return e, err
	}
	e.client = ec2.New(sess)
	e.ctx = context.Background()
	return e, nil
}

func (e EC2) GetAllInstances() (models.InstanceList, error) {
	var end models.InstanceList
	max := int64(1000)
	input := &ec2.DescribeInstancesInput{
		MaxResults: &max, //max -- TODO pagination
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:ManagedBy"),
				Values: []*string{
					aws.String("phoenix"),
				},
			},
		},
	}

	output, err := e.client.DescribeInstances(input)
	if err != nil {
		log.Warnf("Unable to DescribeInstances --%s", err.Error())
		return end, err
	}

	for _, reservation := range output.Reservations {
		var nameTag string
		var phoenixIDTag string

		externInstance := reservation.Instances[0]
		tags := externInstance.Tags

		for _, tag := range tags {
			if *tag.Key == "Name" {
				nameTag = *tag.Value
			}
			if *tag.Key == "PhoenixID" {
				phoenixIDTag = *tag.Value
			}
		}

		instance := models.Instance{
			ExternalID: *externInstance.InstanceId,
			Hostname:   *externInstance.PrivateDnsName,
			Location: models.Location{
				Region: *e.client.Client.Config.Region,
				Zone:   *externInstance.Placement.AvailabilityZone,
			},
		}

		if phoenixIDTag != "" {
			instance.PhoenixID = phoenixIDTag
		}

		if nameTag != "" {
			instance.Name = nameTag
		}

		if matches := clusterNameRegexp.FindSubmatch([]byte(instance.Name)); len(matches) > 1 {
			instance.ClusterName = string(matches[1])
		}

		// log.Debugf("reservation: %s --> instance: %s", reservation, instance)
		end = append(end, &instance)
	}

	return end, nil
}

func (e EC2) CreateInstance(i *models.Instance, cmds []string) error {
	input := &ec2.RunInstancesInput{
		ImageId:  aws.String("ami-0b4a9c56e9f69e9f8"),
		MinCount: aws.Int64(1),
		MaxCount: aws.Int64(1),
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("instance"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String(i.Name),
					},
					{
						Key:   aws.String("PhoenixID"),
						Value: aws.String(i.PhoenixID),
					},
					{
						Key:   aws.String("ManagedBy"),
						Value: aws.String("phoenix"),
					},
				},
			},
		},
		UserData: aws.String(base64.StdEncoding.EncodeToString([]byte("#!/bin/bash\n" + cmds[0]))),
	}

	_, err := e.client.RunInstances(input)
	if err != nil {
		return err
	}
	return nil

}
