package record

import (
    "os/exec"
    "bytes"
    "fmt"
    "log"
    
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
    audiotype := context.GetInput("audiotype").(string)
    
    cmd := exec.Command("arecord", "-f", "S16_LE", "-r16", "-c", "1", "-d", "5", "-t", audiotype)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err1 := cmd.Run()
    if err1 != nil {
        fmt.Println(fmt.Sprint(err1) + ": " + stderr.String())
        log.Fatal(err1)
    }
    
    context.SetOutput("recording", stdout.Bytes())
    
    return true, nil
}
