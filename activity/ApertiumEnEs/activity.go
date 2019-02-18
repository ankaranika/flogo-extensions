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
    english := context.GetInput("english").(string)
    
    home := os.Getenv("HOME")
    //fmt.Println("home:", home)
    intxt := strings.Join([]string{home, "Documents/flogo/speech-translator/files", sender, req_id, "english.txt"}, "/")
    outtxt := strings.Join([]string{home, "Documents/flogo/speech-translator/files", sender, req_id, "spanish.txt"}, "/")
    
    f, err1 := os.Create(intxt)
    if err1 != nil {
        fmt.Println(err1)
    }
    
    _, err2 := f.WriteString(english)
    if err != nil {
        fmt.Println(err2)
        f.Close()
    }
    
    cmd := exec.Command("apertium", "en-es", intxt, outtxt)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err3 := cmd.Run()
    if err3 != nil {
        fmt.Println(fmt.Sprint(err3) + ": " + stderr.String())
        log.Fatal(err3)
    }

    spanish, err4 := ioutil.ReadFile(outtxt)
    if err4 != nil{
        panic("There was an error while reading translated text")
    }
    
    context.SetOutput("spanish", string(spanish))
    
    return true, nil
}
