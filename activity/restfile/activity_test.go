package restfile

import (
    "strings"
    "os"
    
    "io/ioutil"
    "testing"

    "github.com/TIBCOSoftware/flogo-lib/core/activity"
    "github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
    "github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

    if activityMetadata == nil {
        jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
        if err != nil{
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

func TestEval(t *testing.T) {

    defer func() {
        if r := recover(); r != nil {
            t.Failed()
            t.Errorf("panic during execution: %v", r)
        }
    }()

    act := NewActivity(getActivityMetadata())
    tc := test.NewTestActivityContext(getActivityMetadata())

    //setup attrs
    tc.SetInput("uri", "localhost:9691/:ip/:req_id/english.txt")
    tc.SetInput("method", "POST")
    tc.SetInput("pathParams", map[string]string{"ip": "localhost", "req_id": "2"})
    tc.SetInput("type", "text/plain")

    act.Eval(tc)

    //check result attr
    filename := strings.Join([]string{os.Getenv("HOME"), "Documents/flogo/speech-translator/files", "localhost", "2", "english.txt"}, "/")
    
    fi, err := os.Stat(filename);
    if err != nil {
        t.Errorf("panic during execution: %v", err)
    }
    // get the size
    size := fi.Size()
    assert.Equal(t, 14831, int(size))
    
}
