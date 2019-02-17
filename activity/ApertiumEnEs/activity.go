package ApertiumEnEs

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
    
    home := "/home/anna" //os.Getenv("HOME")
    intxt := strings.Join([]string{home, "Documents/flogo/speech-translator/files/apertium", sender, req_id, "english.txt"}, "/")
    outtxt := strings.Join([]string{home, "Documents/flogo/speech-translator/files/apertium", sender, req_id, "spanish.txt"}, "/")
    
    cmd := exec.Command("apertium", "en-es", intxt, outtxt)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err1 := cmd.Run()
    if err1 != nil {
        fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
        log.Fatal(err1)
    }

    result, err2 := ioutil.ReadFile(outtxt)
    if err2 != nil{
        panic("There was an error while translating text")
    }
    
    context.SetOutput("result", string(result[:]))
    
    return true, nil
}
