---
title: AWS RDS Services
weight: 1
---

# Counter
This activity allows you Start,Stop,Reboot and Delete specicfied AWS RDS Instance

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo install github.com/DipeshTest/rdsutils
```

### Third-party libraries used
- #### aws-sdk-go:
	Package aws provides the core SDK's utilities and shared types. This package's utilities are used to simplify setting and reading API operations parameters. For more details, visit - https://docs.aws.amazon.com/sdk-for-go/api/aws/

## Schema
Inputs and Outputs:

```json
 "inputs": [{
		"name": "accessKeyId",
		"type": "string",
		"required": true
	},
	{
		"name": "secretAccessKey",
		"type": "string",
		"required": true
	},
	{
		"name": "region",
		"type": "string",
		"required": true
	},
	{
		"name": "operation",
		"type": "string",
		"allowed": ["StartRdsInstance",
		"StopRdsInstance",
		"RebootRdsInstance",
		"DeleteRdsInstance"],
		"required": true
	},
	{
		"name": "rdsDBInstanceIdentifier",
		"type": "string",
		"required": true
	},
	{
		"name": "dBSnapshotIdentifier",
		"type": "string"
	}],
 "outputs": [{
		"name": "statusCode",
		"type": "string"
	},
	{
		"name": "message",
		"type": "any"
	}]
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| accessKeyId | True     | Access Key ID for your AWS acount , Use link :  https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html |         
| secretAccessKey   | True    |Secret Access Key for your AWS acount , Use link :  https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html |
| region    | True     | Regions and Availability Zones for AWS , Check link : https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html|  
| operation   | True     | Operation that you want to perform on the specicfied aws rds instance |
| rdsDBInstanceIdentifier   | True     | Name of the rds db instance |
| dBSnapshotIdentifier   | False     | Name of the snapshot that you want to create in case of  StopRdsInstance and DeleteRdsInstance. |

Note - 
1. If "dBSnapshotIdentifier" is blank in case of StopRdsInstance and DeleteRdsInstance then no snapshot will be created.
1. Use StartDBInstanceOutput , StopDBInstanceOutput , RebootDBInstanceOutput and  DeleteDBInstanceOutput structs from aws-sdk-go 	   

 


## Examples
Please refer activity_test.go 
