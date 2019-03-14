package espeak

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "bytes"
    "strings"
    "io/ioutil"
    "strconv"
    "math/rand"
    "time"
    
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
    text := context.GetInput("text").(string)
    
    home := os.Getenv("HOME")
    seed := rand.NewSource(time.Now().UnixNano())
    rand := rand.New(seed)
    randno := strconv.Itoa(rand.Intn(100))
    outwav := strings.Join([]string{home, "Documents/flogo/speech-translator/files/espeak/spanish" + randno + ".wav"}, "/")
    dir := strings.Join([]string{home, "Documents/flogo/speech-translator/files/espeak"}, "/")
    
    if _, err1 := os.Stat(dir); os.IsNotExist(err1) {
        err2 := os.MkdirAll(dir, 0755)
        if err2 != nil {
            log.Fatal(err2)
        }
    }
    
    cmd := exec.Command("espeak-ng", "-ves", "-s", "140", "-w", outwav, text)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err3 := cmd.Run()
    if err3 != nil {
        fmt.Println(fmt.Sprint(err3) + ": " + stderr.String())
        log.Fatal(err3)
    }
    
    speech, err4 := ioutil.ReadFile(outwav)
    if err4 != nil {
        log.Fatal(err4)
    }
    
    err5 := os.Remove(outwav)
    if err5 != nil {
        log.Fatal(err5)
    }
    
    context.SetOutput("speech", speech)

    return true, nil
}
