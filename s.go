package main
import "regexp"
import "os"
import "io/ioutil"
import "strings"
import "fmt"

func main() {
  replace, with, files := parseArgs()
  r, _ := regexp.Compile(replace)
  for i := range files {
    var content []byte
    content, _ = ioutil.ReadFile(files[i].Name())
    f, err := os.OpenFile(files[i].Name(), os.O_WRONLY, 0666)
    println(err)
    _, err = f.WriteString(r.ReplaceAllString(string(content), with))
    if err != nil {
      fmt.Print(err)
    }
  }
}

func parseArgs()(string, string, []os.FileInfo){
  var args = strings.Split(os.Args[1],"/")
  var replace = args[1]
  var with = args[2]
  files, _ := ioutil.ReadDir(".")
  return replace, with, files
}
