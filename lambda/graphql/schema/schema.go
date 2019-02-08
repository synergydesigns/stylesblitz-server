//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...
package schema

import "bytes"

// String reads the .graphql schema files from the generated _bindata.go file, concatenating the
// files together into one string.
func String() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)

		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
