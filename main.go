package TestGoMod

import "embed"

//go:embed /*
var substraitFS embed.FS

func GetSubstraitFS() embed.FS {
	return substraitFS
}
