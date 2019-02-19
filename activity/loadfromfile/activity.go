package loadfromfile

import (
    "io/ioutil"
    "log"
    
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
    
    path := context.GetInput("path").(string)
    
    content, err1 := ioutil.ReadFile(path)
    if err1 != nil {
        log.Fatal(err1)
    }
    
    context.SetOutput("content", content)
    
    
    return true, nil
}
