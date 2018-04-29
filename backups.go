package main

import (
  "flag"
  "fmt"
  "os/exec"
  "os"
)

func main() {
  //declare var to hold flag
  var app string
  flag.StringVar(&app, "app", "", "a Heroku app name")

  flag.Parse()
  fmt.Printf("App name:  %s \n", app)

  // run heroku pg:backups:schedules command
  var (
    cmdOut []byte
    err    error
  )
  //build command
  cmdName := "heroku"
  cmdArgs := []string{"pg:backups:schedules","-a", app}
  if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
    fmt.Fprintln(os.Stderr, "There was an error trying to fetch a backups schedule: ", err)
    os.Exit(1)
  }
  result := string(cmdOut)
  fmt.Printf(result)
}
