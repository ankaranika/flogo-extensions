package ApertiumEnEs

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "bytes"
    "strings"
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
    
    english := context.GetInput("english").(string)
    
    home := os.Getenv("HOME")
    seed := rand.NewSource(time.Now().UnixNano())
    rand := rand.New(seed)
    randno := strconv.Itoa(rand.Intn(100))
    intxt := strings.Join([]string{home, "Documents/flogo/speech-translator/files/apertium/english" + randno + ".txt"}, "/")
    dir := strings.Join([]string{home, "Documents/flogo/speech-translator/files/apertium"}, "/")
    
    if _, err1 := os.Stat(dir); os.IsNotExist(err1) {
        err2 := os.MkdirAll(dir, 0755)
        if err2 != nil {
            log.Fatal(err2)
        }
    }
    
    f, err3 := os.OpenFile(intxt, os.O_CREATE|os.O_WRONLY, 0644)
    if err3 != nil {
        log.Fatal(err3)
    }
    
    _, err4 := f.WriteString(english)
    if err4 != nil {
        log.Fatal(err4)
        f.Close()
    }
    f.Close()
    
    cmd := exec.Command("apertium", "en-es", "-u", intxt)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    err5 := cmd.Run()
    if err5 != nil {
        fmt.Println(fmt.Sprint(err5) + ": " + stderr.String())
        log.Fatal(err5)
    }

    err6 := os.Remove(intxt)
    if err6 != nil {
        log.Fatal(err6)
    }
    
    context.SetOutput("spanish", stdout.String())
    
    return true, nil
}
