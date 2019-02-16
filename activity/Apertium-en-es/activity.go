package Apertium-en-es

import (
    "fmt"
    "os/exec"
    "log"
    "bytes"
    "strings"
    
    "github.com/TIBCOSoftware/flogo-lib/core/activity"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
    metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
    return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
    return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

    // do eval
    sender := context.GetInput("ip").(string)
    req_id := context.GetInput("req_id").(string)
    
    intxt := strings.Join("~/Documents/flogo/speech-translator/files/apertium/", sender, "/", req_id, "/english.txt")
    outtxt := strings.Join("~/Documents/flogo/speech-translator/files/apertium/", sender, "/", req_id, "/spanish.txt")
    
    cmd := exec.Command("apertium", "en-es", intxt, outtxt)
    
    var output bytes.Buffer
    cmd.Stdout = &output
    
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("in all caps: %q\n", output.String())

    return true, nil
}
