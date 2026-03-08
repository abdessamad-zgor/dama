package dama

import (
	_ "slices"
	_ "fmt"
	_ "reflect"

	"github.com/abdessamad-zgor/dama/logger"
	dutils "github.com/abdessamad-zgor/dama/utils"
	devent "github.com/abdessamad-zgor/dama/event"
	keybinding "github.com/abdessamad-zgor/dama/keybinding"
)

type IndexItem struct {
	path	string
	element	Element
}

func (item IndexItem) GetElement() Element {
	return item.element
}

func (item IndexItem) GetPath() string {
	return item.path
}

type Navigator struct {
	App		App
	tree	dutils.Tree[Element]
	current	IndexItem
	index	dutils.List[IndexItem]
}

func NewNavigator(app App) *Navigator {
	navigator := &Navigator{
		app,
		dutils.NewTree[Element](app),
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
	logger.Log("Get root of navigation tree: ", current)
	root, ok := current.(App)
	paths := []Element{}
	if ok {
		elements := root.GetElements()
		logger.Log("Getting root elements, found ", len(elements))
		for _, element := range elements {
			navigator.tree.AddNode(current, element)
		}
		paths = append(paths, elements...)
		
		for len(paths) > 0 {
			current = paths[len(paths) - 1]
			currentCont, ok := current.(Container)
			elements = []Element{}
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
	logger.Log("Indexing navigation tree")
	elementNodes := navigator.tree.Flatten()
	logger.Log("Got flattened elements tree: ", elementNodes)
	navigables := []dutils.Node[Element]{}
	for _, elementNode := range elementNodes {
		if elementNode.GetValue().IsNavigable() {
			navigables = append(navigables, *elementNode)
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
	logger.Log("Finished indexing navigation tree: ", navigator.index)
}

func (navigator *Navigator) Setup() {
	logger.Log("Setup navigation")
	navigator.GetNavigationTree()
	navigator.Index()
	items := navigator.index.Items()
	if len(items) > 0 {
		navigator.current = navigator.index.Items()[0]
		navigator.current.element.Focus()
	} 
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
	if element != nil && navigator.current.element != element.element {
		navigator.current.element.Blur()
		element.element.Focus()
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
		matcher, _ := keybinding.GetMatcher(string(reachable.element.GetTag()))
		keybindings = append(keybindings, 
			devent.DamaEvent{
				devent.DKeybinding,
				devent.EventDetail{
					&devent.Keybinding {
						string(reachable.element.GetTag()),
						matcher,
						func (match keybinding.Match) {
							_ = match
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
