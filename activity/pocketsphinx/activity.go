package pocketsphinx

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "bytes"
    "strings"
    "io/ioutil"
    //"reflect"
    
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
    //speech := context.GetInput("speech").([]byte)
    var speech []byte
    var ok bool
    var input interface{} = context.GetInput("speech")
    
    if speech, ok = input.([]byte); !ok {
        speech = []byte(input.(string))
    }
    
    home := os.Getenv("HOME")
    exec_path := strings.Join([]string{home, "Documents/pocketsphinx/hello_ps"}, "/")
    inraw := strings.Join([]string{home, "Documents/flogo/speech-translator/files/pocketsphinx/speech.raw"}, "/")
    dir := strings.Join([]string{home, "Documents/flogo/speech-translator/files/pocketsphinx"}, "/")
    
    if _, err1 := os.Stat(dir); os.IsNotExist(err1) {
        err2 := os.MkdirAll(dir, 0755)
        if err2 != nil {
            log.Fatal(err2)
        }
    }
    
    //fmt.Println(reflect.TypeOf(speech))
    
    err3 := ioutil.WriteFile(inraw, speech, 0644)
    if err3 != nil {
        log.Fatal(err3)
    }
    
    cmd := exec.Command(exec_path, inraw)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err4 := cmd.Run()
    if err4 != nil {
        fmt.Println(fmt.Sprint(err4) + ": " + stderr.String())
        log.Fatal(err4)
    }

    err5 := os.Remove(inraw)
    if err5 != nil {
        log.Fatal(err5)
    }
    
    context.SetOutput("text", stdout.String())
    
    return true, nil
}
