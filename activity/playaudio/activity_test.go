package playaudio

import (
    "io/ioutil"
    "testing"
    "os"
    "strings"
    "log"

    "github.com/TIBCOSoftware/flogo-lib/core/activity"
    "github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
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
    audiofile := strings.Join([]string{home, "Documents/flogo/speech-translator/files/localhost/3/spanish.wav"}, "/")
    
    audio, err := ioutil.ReadFile(audiofile)
    if err != nil {
        log.Fatal(err)
    }
    
    tc.SetInput("audio", audio)
    
    act.Eval(tc)

    //check result attr
}
