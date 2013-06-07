package main
import "fmt"
import "regexp"
import "os"
import "io/ioutil"
import "strings"


func main() {
  var args = strings.Split(os.Args[1],"/")
  var replace = args[1]
  var with = args[2]
  r, _ := regexp.Compile(replace)
  var files []os.FileInfo
  files, _ = ioutil.ReadDir(".")
  for i := range files {
    var content []byte
    content, _ = ioutil.ReadFile(files[i].Name())
    fmt.Println(r.ReplaceAllString(string(content), with))
  }
}
