package backend

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/internal/log"

	"github.com/theonlyjohnny/phoenix/internal/instance"
)

//EC2 Backend
type EC2 struct {
	client *ec2.EC2
	ctx    context.Context
}

func getStrFromCfg(cfg config.BackendConfig, k string) (string, error) {
	var res string
	if interf, ok := cfg[k]; ok {
		if cfgStr, ok := interf.(string); ok {
			if cfgStr == "" {
				return res, fmt.Errorf("backend_config.%s cannot be an empty string", k)
			}
			res = cfgStr
		} else {
			return res, fmt.Errorf("backend_config.%s is not a string", k)
		}
	} else {
		return res, fmt.Errorf("When using the EC2 backend, backend_config.%s is a required parameter", k)
	}
	return res, nil
}

func (e EC2) create(cfg config.BackendConfig) (Backend, error) {
	errs := make([]error, 3)
	awsID, err := getStrFromCfg(cfg, "AWS_ACCESS_KEY_ID")
	errs = append(errs, err)
	awsSecret, err := getStrFromCfg(cfg, "AWS_SECRET_ACCESS_KEY")
	errs = append(errs, err)
	awsRegion, err := getStrFromCfg(cfg, "AWS_REGION")
	errs = append(errs, err)

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

func (e EC2) GetAllInstances() []*instance.Instance {
	end := []*instance.Instance{}
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
		return end
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

		instance := instance.Instance{
			ExternalID: *externInstance.InstanceId,
			Hostname:   *externInstance.PrivateDnsName,
		}

		if phoenixIDTag != "" {
			instance.PhoenixID = phoenixIDTag
		}

		if nameTag != "" {
			instance.Name = nameTag
		}

		// log.Debugf("reservation: %s --> instance: %s", reservation, instance)
		end = append(end, &instance)
	}

	return end
}

// func (e EC2) UpdateInstance(*instance.Instance) error {
// return nil
// }
