package dama

import (
	_ "slices"
	_ "fmt"
	_ "reflect"

	_ "github.com/abdessamad-zgor/dama/logger"
	dutils "github.com/abdessamad-zgor/dama/utils"
	devent "github.com/abdessamad-zgor/dama/event"
	keystroke "github.com/abdessamad-zgor/dama/keystroke"
)

type IndexItem struct {
	path	string
	element	DamaElement
}

func (item IndexItem) GetElement() DamaElement {
	return item.element
}

func (item IndexItem) GetPath() string {
	return item.path
}

type Navigator struct {
	App		*App
	tree	dutils.Tree[DamaElement]
	current	IndexItem
	index	dutils.List[IndexItem]
}

func NewNavigator(app *App) *Navigator {
	navigator := &Navigator{
		app,
		dutils.NewTree[DamaElement](app),
		IndexItem {
			"",
			app,
		},
		dutils.NewList[IndexItem](),
	}
	return navigator
}

func (navigator *Navigator) GetNavigationTree() {
	current := navigator.tree.Root.GetValue()
	root, ok := current.(*App)
	paths := []DamaElement{}
	if ok {
		elements := root.GetElements()
		for _, element := range elements {
			navigator.tree.AddNode(current, element)
		}
		paths = append(paths, elements...)
		
		for len(paths) > 0 {
			current = paths[len(paths) - 1]
			currentCont, ok := current.(*Container)
			if ok {
				elements = currentCont.GetElements()
				for _, element := range elements {
					navigator.tree.AddNode(current, element)
				}
			}
			paths = paths[:len(paths) - 1]
			paths = append(paths, elements...)
		}
	}
}

func (navigator *Navigator) Index() {
	elementNodes := navigator.tree.Flatten()
	navigables := []dutils.Node[DamaElement]{}
	for _, elementNode := range elementNodes {
		if elementNode.GetValue().IsNavigable() {
			navigables = append(navigables, elementNode)
		}
	}
	for _, navigable := range navigables {
		path := string(navigable.GetValue().GetTag())
		parent := navigable.Parent
		for parent != nil {
			path = string(parent.GetValue().GetTag()) + path
			parent = parent.Parent
		}
		navigator.index.Add(IndexItem{
			path,
			navigable.GetValue(),
		})
	}
}

func (navigator *Navigator) Setup() {
	navigator.GetNavigationTree()
	navigator.Index()
}

func (navigator *Navigator) Navigate(tag rune) {
	var element *IndexItem = nil
	basePath := string([]rune(navigator.current.path)[:len(navigator.current.path) - 1])
	for _, e := range navigator.index.Items() {
		if e.path == basePath + string(tag) || e.path == navigator.current.path + string(tag) {
			element = &e
			break;
		}
	}
	if element != nil {
		element.element.Focus()
		navigator.current.element.Blur()
		navigator.current = *element
	}
}

func (navigator *Navigator) GetNavigationKeybindings() []devent.DamaEvent {
	keybindings := []devent.DamaEvent{}
	reachables := []IndexItem{}
	basePath := string([]rune(navigator.current.path)[:len(navigator.current.path) - 1])
	for _, iItem := range navigator.index.Items() {
		if navigator.current.path == iItem.path[:len(iItem.path) - 1] || basePath == iItem.path[:len(iItem.path) - 1] {
			reachables = append(reachables, iItem)
		}
	}
	for _, reachable := range reachables {
		matcher, _ := keystroke.GetMatcher(string(reachable.element.GetTag()))
		keybindings = append(keybindings, 
			devent.DamaEvent{
				devent.DKeybinding,
				devent.EventDetail{
					&devent.Keybinding {
						string(reachable.element.GetTag()),
						matcher,
						func (event devent.EventDetail) {
							_ = event
							navigator.Navigate(reachable.element.GetTag())
						},
					},
					nil,
				},
			})
	}
	return keybindings
}

func (navigator *Navigator) GetIndex() dutils.List[IndexItem] {
	return navigator.index
}

func (navigator *Navigator) GetCurrent() IndexItem {
	return navigator.current
}
