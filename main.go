package TestGoMod

import "embed"

//go:embed extensions/*
//go:embed tests/*
var substraitFS embed.FS

func GetSubstraitFS() embed.FS {
	return substraitFS
}
