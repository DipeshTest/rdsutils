package rdsutils

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

// func TestRdsUtils_StartRdsInstance(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("operation", "StartRdsInstance")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abx")
// 	tc.SetInput("dBSnapshotIdentifier", "")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }

// func TestRdsUtils_StopRdsInstance(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("operation", "StopRdsInstance")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abx")
// 	tc.SetInput("dBSnapshotIdentifier", "")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }

// func TestRdsUtils_DeleteRdsInstance(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("operation", "DeleteRdsInstance")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abx")
// 	tc.SetInput("dBSnapshotIdentifier", "")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }

// func TestRdsUtils_DeleteRdsInstanceWithSnapshot(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("operation", "DeleteRdsInstance")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abt")
// 	tc.SetInput("dBSnapshotIdentifier", "abt-snap-final")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }

// func TestRdsUtils_EmptyAccessKeyId(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("operation", "DeleteRdsInstance")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abt")
// 	tc.SetInput("dBSnapshotIdentifier", "abt-snap-final")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "101")
//
// }

// func TestRdsUtils_EmptySecretAccessKey(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("operation", "DeleteRdsInstance")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abt")
// 	tc.SetInput("dBSnapshotIdentifier", "abt-snap-final")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "102")
//
// }

// func TestRdsUtils_EmptyRegion(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "12345")
// 	tc.SetInput("secretAccessKey", "1234")
// 	tc.SetInput("region", "")
// 	tc.SetInput("operation", "DeleteRdsInstance")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abt")
// 	tc.SetInput("dBSnapshotIdentifier", "abt-snap-final")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "103")
//
// }

// func TestRdsUtils_EmptyOperation(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "123445")
// 	tc.SetInput("secretAccessKey", "1234")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("operation", "")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abt")
// 	tc.SetInput("dBSnapshotIdentifier", "abt-snap-final")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "104")
//
// }

// func TestRdsUtils_InvalidAccessKeyId(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "1243")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("operation", "DeleteRdsInstance")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abt")
// 	tc.SetInput("dBSnapshotIdentifier", "abt-snap-final-new")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "403")
//
// }

// func TestRdsUtils_InvalidSecretAccessKey(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("operation", "DeleteRdsInstance")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abt")
// 	tc.SetInput("dBSnapshotIdentifier", "abt-snap-final-new")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "403")
//
// }

// func TestRdsUtils_ExistingDBSnapshotIdentifier(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-2")
// 	tc.SetInput("operation", "StopRdsInstance")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abt")
// 	tc.SetInput("dBSnapshotIdentifier", "abt-snap-final")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "400")
//
// }

// func TestRdsUtils_InvalidRegion(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("accessKeyId", "")
// 	tc.SetInput("secretAccessKey", "")
// 	tc.SetInput("region", "us-east-22")
// 	tc.SetInput("operation", "DeleteRdsInstance")
// 	tc.SetInput("rdsDBInstanceIdentifier", "abt")
// 	tc.SetInput("dBSnapshotIdentifier", "abt-snap-final")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "106")
//
// }
