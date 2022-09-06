package methodgenhelper

import (
	"fmt"
	"io"
	"text/template"
)

func ExecuteTemplate(
	w io.Writer,
	templ *template.Template,
	packageName string,
	autoGenerationNotice string,
	tyInfo []any,
) (err error) {
	_, err = fmt.Fprintf(w, "package %s\n", packageName)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, autoGenerationNotice)
	if err != nil {
		return err
	}
	for _, v := range tyInfo {
		err := templ.Execute(w, v)
		if err != nil {
			return err
		}
	}
	return nil
}
