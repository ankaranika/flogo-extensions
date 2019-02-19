package pocketsphinx

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "bytes"
    "strings"
    "io/ioutil"
    
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
    speech := context.GetInput("speech").([]byte)
    
    home := os.Getenv("HOME")
    exec_path := strings.Join([]string{home, "Documents/pocketsphinx/hello_ps"}, "/")
    inraw := strings.Join([]string{home, "Documents/flogo/speech-translator/files", sender, req_id, "speech.raw"}, "/")

    err1 := ioutil.WriteFile(inraw, speech, 0644)
    if err1 != nil {
        log.Fatal(err1)
    }
    
    cmd := exec.Command(exec_path, inraw)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err2 := cmd.Run()
    if err2 != nil {
        fmt.Println(fmt.Sprint(err2) + ": " + stderr.String())
        log.Fatal(err2)
    }

    context.SetOutput("text", stdout.String())
    
    return true, nil
}
