package bdest

import (
	"github.com/moisespsena/go-path-helpers"
	"path/filepath"
	"runtime"

	"github.com/moisespsena-go/env-utils"
)

var (
	GO_DEST_ARCH = envutils.Get("GO_DEST_ARCH", runtime.GOARCH)
	GO_DEST_OS   = envutils.Get("GO_DEST_OS", runtime.GOOS)
	GO_DEST_ARM  string
)

func init() {
	if GO_DEST_ARCH == "arm" {
		if GO_DEST_ARM = envutils.FistEnv("GO_DEST_ARM", "GOARM"); GO_DEST_ARM == "" {
			panic(path_helpers.GetCalledDir() + ": Env GO_DEST_ARM and GOARM is blank")
		}
	}
}

func BuildEnv() []string {
	return []string{"GOOS=" + GO_DEST_OS, "GOARCH=" + GO_DEST_ARCH, "GOARM=" + GO_DEST_ARM}
}

func Env() []string {
	return []string{"GO_DEST_OS=" + GO_DEST_OS, "GO_DEST_ARCH=" + GO_DEST_ARCH, "GO_DEST_ARM=" + GO_DEST_ARM}
}

func Values() []string {
	r := []string{GO_DEST_OS, GO_DEST_ARCH}
	if GO_DEST_ARCH == "arm" {
		r = append(r, GO_DEST_ARM)
	}
	return r
}

func BuildDir(base ...string) string {
	return filepath.Join(append(base, Values()...)...)
}
