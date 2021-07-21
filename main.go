package main

import(
  "fmt"
  "os"
  "io/ioutil"
  "gopkg.in/yaml.v3"
  "html/template"
)


func main() {
  // Create index.html for output
  out, err := os.Create("index.html")
  defer out.Close()
  check(err)

  data := map[string]interface{}{}

  file, err := ioutil.ReadFile("data.yaml")
  check(err)

  err = yaml.Unmarshal(file, &data)
  check(err)

  t, err := template.ParseGlob("template/*")
  check(err)

  t.Execute(out, data)
}

// check error and exit program
func check(e error) {
  if e != nil {
    fmt.Println(e)
    os.Exit(1)
  }
}
