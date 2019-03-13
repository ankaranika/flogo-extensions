package audiowavtoraw

import (
    "os"
    "strings"
    "log"
    "io/ioutil"
    "os/exec"
    "bytes"
    "fmt"
    
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
    wav := context.GetInput("wav").([]byte)
    
    home := os.Getenv("HOME")
    inwav := strings.Join([]string{home, "Documents/flogo/speech-translator/files/audiowavtoraw/speech.wav"}, "/")
    outraw := strings.Join([]string{home, "Documents/flogo/speech-translator/files/audiowavtoraw/speech.raw"}, "/")
    dir := strings.Join([]string{home, "Documents/flogo/speech-translator/files/audiowavtoraw"}, "/")
    
    if _, err1 := os.Stat(dir); os.IsNotExist(err1) {
        err2 := os.MkdirAll(dir, 0755)
        if err2 != nil {
            log.Fatal(err2)
        }
    }
    
    err3 := ioutil.WriteFile(inwav, wav, 0644)
    if err3 != nil {
        log.Fatal(err3)
    }
    
    cmd := exec.Command("sox", inwav, "-b", "16", "-c", "1", "-r", "16k", "-t", "raw", outraw)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err4 := cmd.Run()
    if err4 != nil {
        fmt.Println(fmt.Sprint(err4) + ": " + stderr.String())
        log.Fatal(err4)
    }

    rawaudio, err5 := ioutil.ReadFile(outraw)
    if err5 != nil {
        log.Fatal(err5)
    }
    
    err6 := os.Remove(inwav)
    if err6 != nil {
        log.Fatal(err6)
    }
    
    err7 := os.Remove(outraw)
    if err7 != nil {
        log.Fatal(err7)
    }
    
    context.SetOutput("raw", rawaudio)
    
    return true, nil
}
