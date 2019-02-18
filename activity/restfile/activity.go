package restfile

import (
    "fmt"
    "os"
    "strings"
    "io/ioutil"
    "net/http"
    
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
    uri := context.GetInput("uri").(string)
    method := context.GetInput("method").(string)
    pathParams := context.GetInput("pathParams").(map[string]string)
    filetype := context.GetInput("type").(string)
    
    fmt.Println(method)
    
    home := os.Getenv("HOME")
    
    uriParts := strings.Split(uri, "/")
    
    filename := strings.Join([]string{home, "Documents/flogo/speech-translator/files", pathParams["ip"], pathParams["req_id"], uriParts[3]}, "/")
    
    uri = strings.Join([]string{uriParts[0], pathParams["ip"], pathParams["req_id"], uriParts[3]}, "/")
    
    fmt.Println(filename)
    
    switch method {
        case "GET":
            resp, err := http.Get(uri)
            if err != nil {
                fmt.Println(uri)
                panic("There was an error while getting the file")
            }
            defer resp.Body.Close()
            body, err1 := ioutil.ReadAll(resp.Body)
            if err1 != nil {
                fmt.Println(err1)
            }
            
            f, err2 := os.Create(filename)
            if err2 != nil {
                fmt.Println(err2)
            }
            
            _, err3 := f.Write(body)
            if err3 != nil {
                fmt.Println(err3)
                f.Close()
            }
        
        case "POST":
            f, err := os.Open(filename)
            if err != nil {
                panic(err)
            }
            
            resp, err2 := http.Post(uri, filetype, f)
            if err2 != nil {
                fmt.Println(err2)
            }
            fmt.Println("post:", resp)
    }
    
    context.SetOutput("status", true)
    
    return true, nil
}
