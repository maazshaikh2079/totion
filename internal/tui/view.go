package tui

import (
	"fmt"

	"github.com/maazshaikh2079/totion/internal/styles"
)

func (m Model) View() string {
	heading := styles.HearderStyle.Render(`⫘⫘⫘⫘⫘  「✦ ♛ TOTION ♛ ✦ 」⫘⫘⫘⫘⫘ `)

	view := ""

	if m.newFileTIDisplay {
		view = m.newFileTI.View()
	}

	if m.fileContentTADisplay {
		view = m.fileContentTA.View()
	}

	if m.fileListDisplay {
		view = m.fileList.View()
	}

	if m.deleteFileTIDisplay {
		view = m.deleteFileTI.View()
	}

	help := "Ctrl+N: new note · Ctrl+L: list · Ctrl+S: save · Ctrl+D: delete · Esc: back · Ctrl+C: quit"

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", heading, view, help)
}
