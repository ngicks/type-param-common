package methodgenhelper

import (
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/ngicks/type-param-common/slice"
)

func ParseDir(
	dir string,
	targetTypeName []string,
	ignoreList []string,
) (typeInfos []TypeInfo, packageName string, err error) {
	filenames, err := readRegualrFiles(dir)
	if err != nil {
		return nil, "", err
	}

	fset := token.NewFileSet()

	typeInfos = make([]TypeInfo, 0)

	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return nil, "", err
		}

		if slice.Has(ignoreList, filepath.Base(filename)) {
			continue
		}

		typeInfo, pkgName, err := ParseTypeInfo(fset, filename, file, targetTypeName)
		if err != nil {
			return nil, "", err
		}

		if packageName == "" && pkgName != "" && !strings.HasSuffix(pkgName, "_test") {
			packageName = pkgName
		}

		if typeInfo != nil {
			typeInfos = append(typeInfos, typeInfo...)
		}
	}

	return typeInfos, packageName, nil
}

func readRegualrFiles(dir string) (regularFiles []string, err error) {
	dirents, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, v := range dirents {
		if !v.Type().IsRegular() {
			continue
		}
		regularFiles = append(regularFiles, filepath.Join(dir, v.Name()))
	}
	return
}
