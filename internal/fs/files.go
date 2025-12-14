package fs

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/maazshaikh2079/totion/internal/util"
)

var VaultDir string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home directory", err)
	}

	VaultDir = fmt.Sprintf("%s/.totion", homeDir)
	// vaultDir = "vault"
}

func ListFiles() []list.Item {
	items := make([]list.Item, 0)

	entries, err := os.ReadDir(VaultDir)
	if err != nil {
		log.Fatal("Error reading notes:", err)
	}

	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		modTime := info.ModTime().Format("2006-01-02 15:04")

		items = append(items, util.Item{
			Title_: entry.Name(),
			Desc:   fmt.Sprintf("Modified at: %s", modTime),
		})
	}

	return items
}
