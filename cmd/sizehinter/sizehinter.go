package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"text/template"

	methodgenhelper "github.com/ngicks/type-param-common/cmd/method-gen-helper"
)

const autoGenerationNotice = "// Code generated by github.com/ngicks/type-param-common/cmd/lenner. DO NOT EDIT."

var (
	inputDir = flag.String("i", ".", "input dir.")
	targetTy = flag.String("ty", "DeIterator,SeIterator", "target type. comma-separated")
	ignore   = flag.String("ignore", "lenner.go", "ignored filename list. comma-seperated")
	outFile  = flag.String("o", "", "out filename. stdout if empty.")
)

func main() {
	if err := _main(); err != nil {
		panic(err)
	}
}

func _main() error {
	flag.Parse()

	typeInfos, packageName, err := methodgenhelper.ParseDir(
		*inputDir,
		strings.Split(*targetTy, ","),
		strings.Split(*ignore, ","),
	)
	if err != nil {
		return err
	}

	var out io.Writer
	if *outFile == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(*outFile)
		if err != nil {
			return err
		}
	}

	return methodgenhelper.ExecuteTemplate(
		out,
		lennerTemplate,
		packageName,
		autoGenerationNotice,
		methodgenhelper.ToAnySlice(typeInfoToTemplateParam(typeInfos)),
	)
}

var lennerTemplate = template.Must(template.New("v").Parse(`
func ({{.ReceiverName}} {{.TypeName}}[{{.TypeParams}}]) SizeHint() int {
	if sizehinter, ok := {{.ReceiverName}}.{{.InnerMemberName}}.({{.SizeHinterInterfaceName}}); ok {
		return sizehinter.SizeHint()
	}
	return -1
}
`))

type TemplateParam struct {
	methodgenhelper.TypeInfo

	ReceiverName            string
	SizeHinterInterfaceName string
}

func typeInfoToTemplateParam(typInfo []methodgenhelper.TypeInfo) []TemplateParam {
	ret := []TemplateParam{}
	for _, v := range typInfo {
		ret = append(ret, TemplateParam{
			TypeInfo:                v,
			ReceiverName:            "iter",
			SizeHinterInterfaceName: "SizeHinter",
		})
	}
	return ret
}