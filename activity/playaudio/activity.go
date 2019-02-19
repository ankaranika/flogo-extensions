package playaudio

import (
    "os"
    "strings"
    "io/ioutil"
    "log"
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

    audio := context.GetInput("audio").([]byte)
    
    home := os.Getenv("HOME")
    audiofile := strings.Join([]string{home, "Documents/flogo/speech-translator/files/speech.wav"}, "/")
    
    err1 := ioutil.WriteFile(audiofile, audio, 0644)
    if err1 != nil {
        log.Fatal(err1)
    }

    cmd := exec.Command("aplay", "-f", "S16_LE", "-r16", "-c", "1", audiofile)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err2 := cmd.Run()
    if err2 != nil {
        fmt.Println(fmt.Sprint(err2) + ": " + stderr.String())
        log.Fatal(err2)
    }
    
    return true, nil
}
