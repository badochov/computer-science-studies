package code_template

import (
	"io"
	"text/template"
)

var mainTemplate = template.Must(template.New("main").Parse(mainTemplateStr))

type Data struct {
	ClassName   string
	StackLimit  int
	Code        string
	LocalsLimit int
}

func Save(f io.Writer, data Data) error {
	data.LocalsLimit++
	return mainTemplate.Execute(f, data)
}

const mainTemplateStr = `
.class public {{.ClassName}}
.super java/lang/Object

.method public <init>()V
    aload_0
    invokenonvirtual java/lang/Object/<init>()V
    return
.end method

.method public static main([Ljava/lang/String;)V
    .limit stack {{.StackLimit}} 
    .limit locals {{.LocalsLimit}} 
	
   	{{.Code}}
    return
.end method
`
