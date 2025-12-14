package tui

import (
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maazshaikh2079/totion/internal/fs"
	"github.com/maazshaikh2079/totion/internal/styles"
)

// type Item struct {}

type Model struct {
	newFileTI        textinput.Model
	newFileTIDisplay bool

	fileContentTA        textarea.Model
	fileContentTADisplay bool
	openedFile           *os.File

	fileList        list.Model
	fileListDisplay bool

	deleteFileTI        textinput.Model
	deleteFileTIDisplay bool
	fileToDelete        string
}

func (m Model) Init() tea.Cmd {

	return nil
}

func InitializeModel() Model {

	if err := os.MkdirAll(fs.VaultDir, 0750); err != nil {
		log.Fatal(err)
	}

	// initialize newFileTI
	ti1 := textinput.New()
	ti1.Placeholder = "Enter note/file name (!w/.ext)"
	ti1.Focus()
	ti1.CharLimit = 156
	ti1.Width = 30
	ti1.Cursor.Style = styles.CursorStyle
	ti1.PromptStyle = styles.CursorStyle

	// initialze fileContentTA
	ta := textarea.New()
	ta.Placeholder = "Write content..."
	ta.ShowLineNumbers = false
	ta.Focus()
	ta.SetWidth(90)
	ta.SetHeight(23)

	// initialize fileList
	files := fs.ListFiles()
	// fmt.Println(fileList)

	// style and initialize fileList
	delegate := list.NewDefaultDelegate()

	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		Foreground(lipgloss.Color("123")).
		BorderLeftForeground(lipgloss.Color("74"))
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.
		Foreground(lipgloss.Color("74")).
		BorderLeftForeground(lipgloss.Color("74"))

	list := list.New(files, delegate, 0, 0)

	list.Title = "☰ All Notes/Files ☰"
	list.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("15")).
		Padding(0, 1)
	list.FilterInput.PromptStyle = list.FilterInput.PromptStyle.
		Foreground(lipgloss.Color("123"))
	list.FilterInput.Cursor.Style = list.FilterInput.Cursor.Style.
		Foreground(lipgloss.Color("123"))

	ti2 := textinput.New()
	ti2.Placeholder = "Enter note/file name to confirm deletion (!w/.ext)"
	ti2.Focus()
	ti2.CharLimit = 156
	ti2.Width = 60
	ti2.Cursor.Style = styles.CursorStyle
	ti2.PromptStyle = styles.CursorStyle

	return Model{
		newFileTI:        ti1,
		newFileTIDisplay: false,

		fileContentTA:        ta,
		fileContentTADisplay: false,
		openedFile:           nil,

		fileList:        list,
		fileListDisplay: false,

		deleteFileTI:        ti2,
		deleteFileTIDisplay: false,
		fileToDelete:        "",
	}
}
