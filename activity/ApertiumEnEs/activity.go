package ApertiumEnEs

import (
    "fmt"
    "os"
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
    english := context.GetInput("english").(string)
    
    home := os.Getenv("HOME")
    intxt := strings.Join([]string{home, "Documents/flogo/speech-translator/files", sender, req_id, "english.txt"}, "/")
    
    f, err1 := os.Create(intxt)
    if err1 != nil {
        log.Fatal(err1)
    }
    _, err2 := f.WriteString(english)
    if err2 != nil {
        log.Fatal(err2)
        f.Close()
    }
    f.Close()
    
    cmd := exec.Command("apertium", "en-es", "-u", intxt)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err3 := cmd.Run()
    if err3 != nil {
        fmt.Println(fmt.Sprint(err3) + ": " + stderr.String())
        log.Fatal(err3)
    }

    context.SetOutput("spanish", stdout.String())
    
    return true, nil
}
