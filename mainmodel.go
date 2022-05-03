package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState uint

const (
	articleList sessionState = iota
	article
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	state sessionState
	list  list.Model
}

func New() *Model {
	items := []list.Item{item{title: "Best book ever", desc: ""}, item{title: "Huberman Lab is my Favourite Podcast", desc: ""}}
	list := list.New(items, list.NewDefaultDelegate(), 0, 0)
	return &Model{state: articleList, list: list}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.state {
	default:
		m.list, cmd = m.list.Update(msg)
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			cmd = tea.Quit
		}
	}
	return m, cmd
}

func (m Model) View() string {
	switch m.state {
	default:
		return m.list.View()
	}
}
