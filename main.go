package substrait

import "embed"

// Add all directories which should be exposed in below
//
//go:embed extensions/*
//go:embed tests/*
var substraitFS embed.FS

func GetSubstraitFS() embed.FS {
	return substraitFS
}
