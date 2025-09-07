package dama

import (
	"slices"

	"github.com/abdessamad-zgor/dama/logger"
	dutils "github.com/abdessamad-zgor/dama/utils"
	devent "github.com/abdessamad-zgor/dama/event"
	keystroke "github.com/abdessamad-zgor/dama/keystroke"
)

type IndexItem struct {
	path	string
	element	DamaElement
}

type Navigator struct {
	App		*App
	tree	dutils.Tree[DamaElement]
	current	IndexItem
	index	List[IndexItem]
}

func NewNavigator(app DamaApp) Navigator {
	navigator := Navigator{
		app,
		dutils.NewTree[DamaElement](app),
		IndexItem {
			"",
			app,
		},
		NewList[IndexItem](),
	}
	return navigator
}

func (navigator Navigator) GetNavigationTree() {
	current := *navigator.tree.Root.Value
	currentCont, ok := current.(*Container)
	paths := []DamaElement{}
	if ok {
		elements := currentCont.GetElements()
		for _, element := range elements {
			navigator.tree.AddNode(current, element)
		}
		paths = append(paths, elements...)
		
		for len(paths) > 0 {
			current = paths[len(paths) - 1]
			currentCont, ok = current.(*Container)
			if ok {
				elements = currentCont.GetElements()
				for _, element := range elements {
					navigator.tree.AddNode(current, element)
				}
				paths = paths[:len(paths) - 1]
				paths = append(paths, elements...)
			}
		}
	}
}

func (navigator Navigator) Index() {
	elementNodes := navigator.tree.Flatten()
	navigables := []Node[DamaElement]{}
	for _, elementNode := range elementNodes {
		if elementNode.Value.IsNavigable() {
			navigables = append(navigables, elementNode)
		}
	}
	for _, navigable := range navigables {
		path := string(navigable.Value.GetTag())
		parent := navigable.Parent
		for parent != nil {
			path = parent.Value.GetTag() + path
			parent = parent.Parent
		}
		navigator.index.Add(navigable.Value)
	}
}

func (navigator Navigator) Navigate(tag rune) {
	var element *IndexItem = nil
	basePath := string([]rune(navigator.current.path)[:len(navigator.current.path) - 1])
	for _, e := range navigator.index.Items() {
		if e.path == basePath + tag || e.path == navigator.current.path + tag {
			element = e
			break;
		}
	}
	if element != nil {
		element.element.Focus()
		navigator.element.Blur()
		navigator.current = *element
	}
}

func (navigator Navigator) GetNagigationKeybindings() devent.Keybinding[] {
	keybindings := []devent.Keybinding{}
	reachables := []IndexItem{}
	basePath := string([]rune(navigator.current.path)[:len(navigator.current.path) - 1])
	for _, iItem := range navigator.index {
		if navigator.current.path == iItem.path[:len(iItem.path) - 1] || basePath == iItem.path[:len(iItem.path) - 1] {
			reachables = append(reachables, iItem)
		}
	}
	for _, reachable := range reachables {
		matcher, _ := keystroke.GetMatcher(string(reachable.element.GetTag()))
		keybindings = append(keybindings, devent.Keybinding {
			string(reachable.element.GetTag()),
			matcher,
			func (event devent.Event) {
				_ = event
				navigator.Navigate(reachable.element.GetTag())
			},
		})
	}
	return keybindings
}
