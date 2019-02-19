package espeak

import (
    "io/ioutil"
    "testing"
    "os"
    "strings"

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
    tc.SetInput("ip", "localhost")
    tc.SetInput("req_id", "3")
    tc.SetInput("text", "Esto es una frase de prueba")

    act.Eval(tc)

    //check result attr
    
    speech := tc.GetOutput("speech").([]byte)
    
    wav := strings.Join([]string{os.Getenv("HOME"), "Documents/flogo/speech-translator/files", "localhost", "3", "spanish.wav"}, "/")
    
    fi, err := os.Stat(wav);
    if err != nil {
        t.Errorf("panic during execution: %v", err)
    }
    // get the size
    size := fi.Size()
    
    assert.Equal(t, int(size), len(speech))
    
}
