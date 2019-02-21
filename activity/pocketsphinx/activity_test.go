package pocketsphinx

import (
    "os"
    "strings"
    "log"
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
    
    home := os.Getenv("HOME")
    speechfile := strings.Join([]string{home, "Documents/flogo/speech-translator/files/english.raw"}, "/")
    
    speech, err2 := ioutil.ReadFile(speechfile)
    if err2 != nil {
        log.Fatal(err2)
    }
    
    //tc.SetInput("ip", "localhost")
    //tc.SetInput("req_id", "1")
    tc.SetInput("speech", speech)

    act.Eval(tc)

    //check result attr
    text := tc.GetOutput("text")
    assert.Equal(t, "go forward ten meters", text)
}
