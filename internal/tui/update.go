package tui

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maazshaikh2079/totion/internal/fs"
	"github.com/maazshaikh2079/totion/internal/styles"
	"github.com/maazshaikh2079/totion/internal/util"
)

// type Item struct {}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := styles.DocStyle.GetFrameSize()
		m.fileList.SetSize(msg.Width-h, msg.Height-v-5)
		m.fileContentTA.SetWidth(msg.Width - h)
		m.fileContentTA.SetHeight(msg.Height - v - 5)

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":

			return m, tea.Quit

		case "esc":

			if m.newFileTIDisplay {
				m.newFileTI.SetValue("")
				m.newFileTIDisplay = false
			}

			// if m.openedFile != nil {
			if m.fileContentTADisplay {
				m.fileContentTA.SetValue("")
				m.fileContentTADisplay = false
				m.openedFile = nil
				if m.fileList.FilterState() == list.Filtering || m.fileList.FilterState() == list.FilterApplied {
					m.fileList.ResetFilter()
				}
			}

			if m.fileListDisplay {

				if m.fileList.FilterState() == list.Filtering || m.fileList.FilterState() == list.FilterApplied {

					break
				}

				m.fileListDisplay = false
			}

			if m.deleteFileTIDisplay {
				m.deleteFileTI.SetValue("")
				m.deleteFileTIDisplay = false
				m.fileToDelete = ""

				m.fileListDisplay = true
			}

			return m, nil

		case "ctrl+n":
			m.newFileTIDisplay = true

			return m, nil

		case "ctrl+l":
			fileList := fs.ListFiles()
			m.fileList.SetItems(fileList)
			m.fileListDisplay = true

			return m, nil

		case "ctrl+s":

			if m.openedFile == nil {

				break
			}

			if err := m.openedFile.Truncate(0); err != nil {
				fmt.Println("can not save the file")

				return m, nil
			}

			if _, err := m.openedFile.Seek(0, 0); err != nil {
				fmt.Println("can not save the file")

				return m, nil
			}

			if _, err := m.openedFile.WriteString(m.fileContentTA.Value()); err != nil {
				fmt.Println("can not save the file")
			}

			if err := m.openedFile.Close(); err != nil {
				fmt.Println("can not close the file. Error:", err)
			}

			m.fileContentTA.SetValue("")
			m.fileContentTADisplay = false
			m.openedFile = nil
			if m.fileList.FilterState() == list.Filtering || m.fileList.FilterState() == list.FilterApplied {
				m.fileList.ResetFilter()
			}

			return m, nil

		case "enter":

			// If a file is already open, ignore Enter (user is typing in textarea)
			// if m.openedFile != nil {
			if m.fileContentTADisplay {

				break
			}
			// ---------------------

			// pressed `enter` key when entering note/file name in newFileTI (for creating file) :-

			if m.newFileTIDisplay {

				// creating a new file
				fileName := m.newFileTI.Value()
				if fileName == "" {

					return m, nil
				}

				filePath := fmt.Sprintf("%s/%s.md", fs.VaultDir, fileName)

				// prevent overwriting existing file
				if _, err := os.Stat(filePath); err == nil {
					// file already exists

					return m, nil
				}

				// create new file
				f, err := os.Create(filePath)
				if err != nil {
					log.Fatalf("%v", err)
				}

				// open file with textarea
				m.openedFile = f
				m.fileContentTADisplay = true
				m.newFileTI.SetValue("")
				m.newFileTIDisplay = false

				return m, nil
			}
			// ---------------------

			// pressed `enter` key when file list component is being displayed (for accessing files):-

			if m.fileListDisplay {

				item, ok := m.fileList.SelectedItem().(util.Item)
				if ok {
					filePath := fmt.Sprintf("%s/%s", fs.VaultDir, item.Title_)

					fileContent, err := os.ReadFile(filePath)
					if err != nil {
						log.Printf("Error reading file: %v", err)

						return m, nil
					}
					m.fileContentTA.SetValue(string(fileContent))
					m.fileContentTADisplay = true
					m.fileListDisplay = false

					f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
					if err != nil {
						log.Printf("Error opening file for reading: %v", err)

						return m, nil
					}
					m.openedFile = f
				}

				return m, nil
			}
			// ---------------------

			// pressed `enter` key when entering note/file name in deleteFileTI (for deleting file) :-

			if m.deleteFileTIDisplay {

				// if m.fileToDelete == m.deleteFileTI.Value() {
				if m.fileToDelete == fmt.Sprintf("%s.md", m.deleteFileTI.Value()) {

					filePath := fmt.Sprintf("%s/%s", fs.VaultDir, m.fileToDelete)

					err := os.Remove(filePath)
					if err != nil {
						log.Printf("Error deleting file: %v", err)
						return m, nil
					}

					m.fileToDelete = ""
					m.deleteFileTI.SetValue("")
					m.deleteFileTIDisplay = false

					files := fs.ListFiles()
					m.fileList.SetItems(files)
					m.fileListDisplay = true
					if m.fileList.FilterState() == list.Filtering || m.fileList.FilterState() == list.FilterApplied {
						m.fileList.ResetFilter()
					}
				}

				return m, nil
			}

			// ---------------------

		case "ctrl+d":

			if m.fileListDisplay {
				item, ok := m.fileList.SelectedItem().(util.Item)
				if ok {
					m.fileToDelete = item.Title_
					// log.Println(filePath)

					m.fileListDisplay = false
					m.deleteFileTIDisplay = true
					if m.fileList.FilterState() == list.Filtering || m.fileList.FilterState() == list.FilterApplied {
						m.fileList.ResetFilter()
					}
				}
			}

			return m, nil

		}
	}

	if m.newFileTIDisplay {
		m.newFileTI, cmd = m.newFileTI.Update(msg)
	}

	// if m.openedFile != nil {
	if m.fileContentTADisplay {
		m.fileContentTA, cmd = m.fileContentTA.Update(msg)
	}

	if m.fileListDisplay {
		m.fileList, cmd = m.fileList.Update(msg)
	}

	if m.deleteFileTIDisplay {
		m.deleteFileTI, cmd = m.deleteFileTI.Update(msg)
	}

	return m, cmd
}
