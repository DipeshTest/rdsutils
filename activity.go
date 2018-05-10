package rdsutils

import (
	s "strings"

	"github.com/DipeshTest/allstarsshared/awsrds"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var log = logger.GetLogger("activity-twitterpublish")

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	accessKeyId := s.TrimSpace(context.GetInput("accessKeyId").(string))
	secretAccessKey := s.TrimSpace(context.GetInput("secretAccessKey").(string))

	region := s.TrimSpace(context.GetInput("region").(string))
	operation := s.TrimSpace(context.GetInput("operation").(string))
	rdsDBInstanceIdentifier := s.TrimSpace(context.GetInput("rdsDBInstanceIdentifier").(string))

	//	twitterFunction := context.GetInput("twitterFunction").(string)
	//fmt.Println("test ", twitterFunction)
	if len(accessKeyId) == 0 {

		context.SetOutput("statusCode", "101")

		context.SetOutput("message", "Access KeyId field is blank")
	} else if len(secretAccessKey) == 0 {

		context.SetOutput("statusCode", "102")

		context.SetOutput("message", "SecretAccessKey  field is blank")

	} else if len(region) == 0 {

		context.SetOutput("statusCode", "103")

		context.SetOutput("message", "Region  field is blank")

	} else if len(rdsDBInstanceIdentifier) == 0 {

		context.SetOutput("statusCode", "105")

		context.SetOutput("message", "rdsDBInstanceIdentifier field is blank")

	} else {
		var msg string
		var code int

		switch operation {

		case "StartRdsInstance":
			{

				startDbInstnceRequest := &rds.StartDBInstanceInput{
					DBInstanceIdentifier: aws.String(rdsDBInstanceIdentifier),
				}

				code, msg = awsrds.StartRdsInstance(accessKeyId, secretAccessKey, region, startDbInstnceRequest)

			}

		case "StopRdsInstance":
			{
				//skipSnapshot := context.GetInput("skipSnapshot").(bool)
				dBSnapshotIdentifier := s.TrimSpace(context.GetInput("dBSnapshotIdentifier").(string))
				stopDbInstnceRequest := &rds.StopDBInstanceInput{}

				if len(dBSnapshotIdentifier) > 0 {
					stopDbInstnceRequest = &rds.StopDBInstanceInput{
						DBInstanceIdentifier: aws.String(rdsDBInstanceIdentifier),
						DBSnapshotIdentifier: aws.String(dBSnapshotIdentifier),
					}

				} else {
					stopDbInstnceRequest = &rds.StopDBInstanceInput{
						DBInstanceIdentifier: aws.String(rdsDBInstanceIdentifier),
					}
					//	code, msg = awsrds.StopRdsInstance(accessKeyId, secretAccessKey, region, stopDbInstnceRequest)

				}

				code, msg = awsrds.StopRdsInstance(accessKeyId, secretAccessKey, region, stopDbInstnceRequest)

			}
		case "RebootRdsInstance":
			{

				rebootDbInstnceRequest := &rds.RebootDBInstanceInput{
					DBInstanceIdentifier: aws.String(rdsDBInstanceIdentifier),
				}

				code, msg = awsrds.RebootRdsInstance(accessKeyId, secretAccessKey, region, rebootDbInstnceRequest)

			}

		case "DeleteRdsInstance":
			{
				var skipSnapshot bool
				dBSnapshotIdentifier := s.TrimSpace(context.GetInput("dBSnapshotIdentifier").(string))
				if len(dBSnapshotIdentifier) > 0 {
					skipSnapshot = false

				} else {
					skipSnapshot = true

				}

				deleteDbInstnceRequest := &rds.DeleteDBInstanceInput{
					DBInstanceIdentifier:      aws.String(rdsDBInstanceIdentifier),
					SkipFinalSnapshot:         aws.Bool(skipSnapshot),
					FinalDBSnapshotIdentifier: aws.String(dBSnapshotIdentifier),
				}

				code, msg = awsrds.DeleteRdsInstance(accessKeyId, secretAccessKey, region, deleteDbInstnceRequest)

			}

		default:
			{
				code = 104

				msg = "Operation  field is blank"
			}

		}

		context.SetOutput("statusCode", code)

		context.SetOutput("message", msg)
	}

	return true, err
}
