package adapter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type GitInfo struct {
	author     string
	repository string
	email      string
	hash       string
}

func (g *GitInfo) Update() error {
	repoPath := filepath.Join(".", g.repository)
	
	if _, err := os.Stat(repoPath); err == nil {
		// Le repository existe, on fait un pull
		cmd := exec.Command("git", "-C", repoPath, "pull")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Erreur lors du pull: %v\nOutput: %s", err, string(output))
			return fmt.Errorf("erreur lors du pull: %v", err)
		}
		log.Printf("Pull réussi pour %s", g.repository)
		return nil
	} else {
		// Le repository n'existe pas, on fait un clone
		return g.Clone()
	}
}

func (g *GitInfo) Clone() error {
	repoPath := filepath.Join(".", g.repository)
	repoURL := fmt.Sprintf("https://github.com/%s", g.repository)
	
	cmd := exec.Command("git", "clone", repoURL, repoPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Erreur lors du clone: %v\nOutput: %s", err, string(output))
		return fmt.Errorf("erreur lors du clone: %v", err)
	}
	
	log.Printf("Clone réussi pour %s", g.repository)
	return nil
}

