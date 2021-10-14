package main

import (
	"bytes"
	"errors"
	_ "gotti/arudeko-generator/statik" //statikフォルダのパス
	"text/template"
	"io/ioutil"
	"os"
  "strconv"

	"github.com/rakyll/statik/fs"
)

type full struct{
  Header string
  Sections string
}

type sections struct{
  sections []section
}

type section struct{
  SourceCode string
  Index string
}

func main(){
  args := os.Args[1:]
  for _,d := range args{
    _, err := os.Stat(d)
    if err != nil{
      panic(errors.New("file not found: "+d))
    }
  }

  header,err := readStatikFile("/header.md")
  if err != nil{
    panic(err)
  }
  fullTemplateFile, err := readStatikFile("/full.md")
  if err != nil{
    panic(err)
  }
  fullTemplate,err := template.New("").Parse(fullTemplateFile)
  if err != nil{
    panic(err)
  }
  sectionTemplateFile, err := readStatikFile("/section.md")
  if err != nil{
    panic(err)
  }
  sectionTemplate,err := template.New("").Parse(sectionTemplateFile)
  if err != nil{
    panic(err)
  }

  var sec sections

  for i,f := range args{
    source,err := readFile(f)
    if err != nil {
      panic(err)
    }
    sec.sections = append(sec.sections, section{SourceCode: source, Index: strconv.Itoa(i+1)})
  }

  var sectionsConcatedBytes bytes.Buffer
  for _,s := range sec.sections{
    sectionTemplate.Execute(&sectionsConcatedBytes, s)
  }
  fullTemplate.Execute(os.Stdout, full{Header: header, Sections: sectionsConcatedBytes.String()})
}

func readFile(filepath string) (string, error){
  f, err := os.Open(filepath)
  if err != nil{
    return "", err
  }
  s, err := ioutil.ReadAll(f)
  if err != nil{
    return "", err
  }
  return string(s), nil
}

func readStatikFile(filepath string) (string, error){
  fsys, err := fs.New()
  if err != nil{
    return "", err
  }
  f, err := fsys.Open(filepath)
  if err != nil{
    return "", err
  }
  s, err := ioutil.ReadAll(f)
  if err != nil{
    return "", err
  }
  return string(s), nil
}
