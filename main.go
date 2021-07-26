package main

import(
  "fmt"
  "os"
  "io/ioutil"
  "gopkg.in/yaml.v3"
  "html/template"
)


func main() {

  data := map[string]interface{}{}

  file, err := ioutil.ReadFile("data.yaml")
  check(err)

  err = yaml.Unmarshal(file, &data)
  check(err)

  formats := []string{"general", "euro"}

  for _, format := range formats {
    t, err := template.ParseGlob("templates/"+format+"/*")
    check(err)

    out, err := os.Create(format+".html")
    defer out.Close()
    check(err)

    t.Execute(out, data)
    fmt.Printf("Resume generated successfully for %s\n", format)
  }

}

// check error and exit program
func check(e error) {
  if e != nil {
    fmt.Println(e)
    os.Exit(1)
  }
}
