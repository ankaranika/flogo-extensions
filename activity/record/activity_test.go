package record

import (
    "os"
    "strings"
    "os/exec"
    "bytes"
    "fmt"
    "log"
    "io/ioutil"
    "testing"

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

    act.Eval(tc)
    
    //check result attr
    recording := tc.GetOutput("recording").([]byte)
    
    home := os.Getenv("HOME")
    audiofile := strings.Join([]string{home, "Documents/flogo/speech-translator/files/recording.raw"}, "/")
    dir := strings.Join([]string{home, "Documents/flogo/speech-translator/files/"}, "/")
    
    if _, err1 := os.Stat(dir); os.IsNotExist(err1) {
        err2 := os.MkdirAll(dir, 0755)
        if err2 != nil {
            log.Fatal(err2)
        }
    }
    
    err3 := ioutil.WriteFile(audiofile, recording, 0644)
    if err3 != nil {
        log.Fatal(err3)
    }

    cmd := exec.Command("aplay", "-f", "S16_LE", "-r16", "-c", "1", audiofile)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err4 := cmd.Run()
    if err4 != nil {
        fmt.Println(fmt.Sprint(err4) + ": " + stderr.String())
        log.Fatal(err4)
    }
    
    err5 := os.Remove(audiofile)
    if err5 != nil {
        log.Fatal(err5)
    }
    
}
