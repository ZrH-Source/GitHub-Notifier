package adapter

import (
	"os"
	"os/exec"
)

type GitInfo struct {
	author     string
	repository string
	email      string
	hash       string
}

func (g *GitInfo) Update() {
	if _, err := os.Stat(g.repository); err != nil {
		path := g.repository
		exec.Command("git pull https://github.com/" + path)
		println("Pull success")
	} else {
		println(err)
		g.Clone()
	}
}

func (g *GitInfo) Clone() {
	path := g.repository
	exec.Command("git clone https://github.com/" + path)
	println("Clone success")
}
