
### Create a golang DSL (front-end)
```go
package main

func main() {
	Platform := "aws"
	Authenticate()
	Ret := CreateEKSCluster("EKSCluster-1", 3)
	additionalNodeCount := lambda currentNode, expectedNode int: expectedNode-currentNode if expectedNode > currentNode else 0
	currentNodeCount := getCurrentNodeCount()
	expectedNodeCount := 6
	additionalNode := additionalNodeCount(currentNodeCount, expectedNodeCount)
	expandEKSCluster(additionalNode)
}

```
Create backend execution
```go
package tests

import (
	"fmt"
	"testing"

	"github.com/Quinn-Fang/quinne/quinne"
	"github.com/Quinn-Fang/quinne/uspace"
)

func Test_Infra_1(t *testing.T) {
	eventHandler := quinne.NewEventHandler("data/infra/infra_sample_1.go")
	event, err := eventHandler.GetNextEvent()
	for err == nil {
		fmt.Printf("%+v\n", event)
		if event.GetEventType() == uspace.EventTypeFunction {
			// deal with how function show be executed here
			// and provide the return value
			fFunction := event.GetFunction(event.GetEventContext())
			switch fFunction.GetFunctionName() {
			case "Authenticate":
				{
					// Aws authenticate( IAM Role )
				}
			case "getCurrentNodeCount":
				{
					// Call Aws get cluster info(DescribeCluster) here
					// to get the currentNode number

					//
					//svc := eks.New(session.New())
					//input := &eks.DescribeClusterInput{
					//	Name: aws.String("devel"),
					//}

					//result, err := svc.DescribeCluster(input)
					//if err != nil {
					//	if aerr, ok := err.(awserr.Error); ok {
					//		switch aerr.Code() {
					//		case eks.ErrCodeResourceNotFoundException:
					//			fmt.Println(eks.ErrCodeResourceNotFoundException, aerr.Error())
					//		case eks.ErrCodeClientException:
					//			fmt.Println(eks.ErrCodeClientException, aerr.Error())
					//		case eks.ErrCodeServerException:
					//			fmt.Println(eks.ErrCodeServerException, aerr.Error())
					//		case eks.ErrCodeServiceUnavailableException:
					//			fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
					//		default:
					//			fmt.Println(aerr.Error())
					//		}
					//	} else {
					//		// Print the error, cast err to awserr.Error to get the Code and
					//		// Message from an error.
					//		fmt.Println(err.Error())
					//	}
					//	return
					//}

					//fmt.Println(result)

					// set currentNode Number as function return value
					fFunction.SetReturnValue(2)
				}
			case "CreateEKSCluster":
				{
					params := fFunction.GetParams()
					// Here parameter checking is required.
					clusterName := params[0].GetVariableValue().(string)
					nodeCount := params[1].GetVariableValue().(int)
					fmt.Println(clusterName, nodeCount)
					// create clster with these variables
					//svc := eks.New(session.New())
					//input := &eks.CreateClusterInput{
					//	ClientRequestToken: aws.String("1d2129a1-3d38-460a-9756-e5b91fddb951"),
					//	Name:               aws.String("prod"),
					//	ResourcesVpcConfig: &eks.VpcConfigRequest{
					//		SecurityGroupIds: []*string{
					//			aws.String("sg-6979fe18"),
					//		},
					//		SubnetIds: []*string{
					//			aws.String("subnet-6782e71e"),
					//			aws.String("subnet-e7e761ac"),
					//		},
					//	},
					//	RoleArn: aws.String("arn:aws:iam::012345678910:role/eks-service-role-AWSServiceRoleForAmazonEKS-J7ONKE3BQ4PI"),
					//	Version: aws.String("1.10"),
					//}

					//result, err := svc.CreateCluster(input)
					//if err != nil {
					//	if aerr, ok := err.(awserr.Error); ok {
					//		switch aerr.Code() {
					//		case eks.ErrCodeResourceInUseException:
					//			fmt.Println(eks.ErrCodeResourceInUseException, aerr.Error())
					//		case eks.ErrCodeResourceLimitExceededException:
					//			fmt.Println(eks.ErrCodeResourceLimitExceededException, aerr.Error())
					//		case eks.ErrCodeInvalidParameterException:
					//			fmt.Println(eks.ErrCodeInvalidParameterException, aerr.Error())
					//		case eks.ErrCodeClientException:
					//			fmt.Println(eks.ErrCodeClientException, aerr.Error())
					//		case eks.ErrCodeServerException:
					//			fmt.Println(eks.ErrCodeServerException, aerr.Error())
					//		case eks.ErrCodeServiceUnavailableException:
					//			fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
					//		case eks.ErrCodeUnsupportedAvailabilityZoneException:
					//			fmt.Println(eks.ErrCodeUnsupportedAvailabilityZoneException, aerr.Error())
					//		default:
					//			fmt.Println(aerr.Error())
					//		}
					//	} else {
					//		// Print the error, cast err to awserr.Error to get the Code and
					//		// Message from an error.
					//		fmt.Println(err.Error())
					//	}
					//	return
					fFunction.SetReturnValue("EKS cluster successfully created!")
				}
			case "expandEKSCluster":
				{
					params := fFunction.GetParams()
					// check params here
					additionalNodeCount := params[0]
					fmt.Println(additionalNodeCount)
					// params := NewScalingConfig(additionalNodeCount)

					//req, resp := client.UpdateClusterConfigRequest(params)

					//err := req.Send()
					//if err == nil { // resp is now filled
					//	fmt.Println(resp)
					//}
					fFunction.SetReturnValue("SuccessFully")
				}
			}
		}

		event, err = eventHandler.GetNextEvent()
	}
}

```
