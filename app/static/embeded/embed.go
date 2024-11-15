package embeded

import (
	"embed"
)

//go:embed locales/*.json
var FsLocales embed.FS

//go:embed dist/*.*
var FsWeb embed.FS
