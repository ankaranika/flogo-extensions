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
    speech := context.GetInput("speech")
    
    home := os.Getenv("HOME")
    //fmt.Println("home:", home)
    exec_path := strings.Join([]string{home, "Documents/pocketsphinx/hello_ps"}, "/")
    inraw := strings.Join([]string{home, "Documents/flogo/speech-translator/files", sender, req_id, "speech.raw"}, "/")
    outtxt := strings.Join([]string{home, "Document/flogo/speech-translator/files", sender, req_id, "english.txt"}, "/")
    
    f, err := os.Create(inraw)
    if err != nil {
        fmt.Println(err)
    }
    
    err2 := ioutil.WriteAll(inraw, speech, 0777)
    if err != nil {
        fmt.Println(err2)
        f.Close()
    }
    
    cmd := exec.Command(exec_path, inraw)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err1 := cmd.Run()
    if err1 != nil {
        fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
        log.Fatal(err1)
    }

    context.SetOutput("text", stdout.String())
    
    return true, nil
}
