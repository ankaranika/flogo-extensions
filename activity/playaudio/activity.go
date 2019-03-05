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

    //audio := context.GetInput("audio").([]byte)
    var audio []byte
    var input interface{} = context.GetInput("audio")
    //str, ok := input.(string)
    if str, ok := input.(string); ok {
        audio = []byte(str)
    } else {
        audio = input.([]byte)
    }
    
    home := os.Getenv("HOME")
    audiofile := strings.Join([]string{home, "Documents/flogo/speech-translator/files/playaudio/speech.wav"}, "/")
    dir := strings.Join([]string{home, "Documents/flogo/speech-translator/files/playaudio"}, "/")
    
    if _, err1 := os.Stat(dir); os.IsNotExist(err1) {
        err2 := os.MkdirAll(dir, 0755)
        if err2 != nil {
            log.Fatal(err2)
        }
    }
    
    err3 := ioutil.WriteFile(audiofile, audio, 0644)
    if err3 != nil {
        log.Fatal(err3)
    }

    cmd := exec.Command("aplay", "-f", "S16_LE", "-r16", "-c", "1", audiofile)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err4 := cmd.Run()
    if err4 != nil {
        fmt.Println(fmt.Sprint(err4) + ": " + stderr.String())
        log.Fatal(err4)
    }
    
    err5 := os.Remove(audiofile)
    if err5 != nil {
        log.Fatal(err5)
    }
    
    return true, nil
}
