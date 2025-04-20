package dama

import (
	"slices"

	"github.com/abdessamad-zgor/dama/logger"
)

type NavigationItem struct {
	Parent   *NavigationItem
	Element  DamaElement
	Children *[]NavigationItem
}

type IndexItem struct {
	Path string
	Item NavigationItem
}

type Navigator struct {
	Root        NavigationItem
	Index       []IndexItem
	Current     NavigationItem
	CurrentPath string
}

func NewNavigator() *Navigator {
	navigator := new(Navigator)
	navigator.CurrentPath = ""
	navigator.Index = []IndexItem{}
	return navigator
}

func getNavigationTree(elements []DamaElement) []NavigationItem {
	logger.Logger.Println("getNavigationTree ")
	navigationItems := []NavigationItem{}
	for _, element := range elements {
		navigationItem := NavigationItem{nil, element, nil}
		container, ok := element.(*Container)
		if ok {
			itemNavigables := container.GetNavigables()
			itemChildren := getNavigationTree(itemNavigables)
			if len(itemChildren) > 0 {
				navigationItem.Children = &itemChildren
				for i := range itemChildren {
					itemChildren[i].Parent = &navigationItem
				}
			}
		}
		navigationItems = append(navigationItems, navigationItem)
	}
	return navigationItems
}

func (navigator *Navigator) indexItems() {
	stack := []NavigationItem{}
	current_item := navigator.Root
	logger.Logger.Println("Navigator root: ", current_item)

	stack = append(stack, current_item)
	for len(stack) != 0 {
		current_item = stack[len(stack)-1]
		stack = slices.Delete(stack, len(stack)-1, len(stack))

		parent := current_item.Parent
		prefix := ""
		if parent != nil {
			prefix = string(parent.Element.GetTag()) + prefix
		}
		logger.Logger.Println("index: ", navigator.Index, " current_item: ", current_item)
		if current_item.Element.IsNavigable() {
			navigator.Index = append(navigator.Index, IndexItem{prefix + string(current_item.Element.GetTag()), current_item})
		}
		if current_item.Children != nil {
			for _, child := range *current_item.Children {
				stack = append(stack, child)
			}
		}
	}
}

func (navigator *Navigator) GetNavigationTree(elements []DamaElement) {
	rootChildren := getNavigationTree(elements)
	navigator.Root.Children = &rootChildren
	navigator.indexItems()
}

func (navigator *Navigator) SetupKeybindings() []rune {
	tags := []rune{}
	for _, item := range navigator.Index {
		tag := item.Item.Element.GetTag()
		if len(item.Path) == 1 || item.Path[:len(item.Path)-1] == navigator.CurrentPath[:len(navigator.CurrentPath)-1] || item.Path[:len(item.Path)-1] == navigator.CurrentPath {
			tags = append(tags, tag)
		}
	}
	return tags
}

func (navigator *Navigator) SearchIndex(tag rune) *IndexItem {
	var indexItem *IndexItem = nil
	for _, item := range navigator.Index {
		siblingPath := navigator.CurrentPath
		indexTag := rune(0)
		if len(navigator.CurrentPath) > 0 {
			siblingPath = navigator.CurrentPath[:len(siblingPath)-1]
		}
		if indexItem != nil {
			indexTag = indexItem.Item.Element.GetTag()

		}
		if item.Path == navigator.CurrentPath+string(tag) || item.Path == siblingPath+string(tag) || indexTag == tag {
			indexItem = &item
		}
	}

	return indexItem
}

func (navigator *Navigator) Navigate(tag rune) bool {
	indexItem := navigator.SearchIndex(tag)

	if indexItem != nil {
		if navigator.Current.Element != nil {
			navigator.Current.Element.Blur()
		}
		navigator.Current = indexItem.Item
		navigator.CurrentPath = indexItem.Path
		navigator.Current.Element.Focus()
		return true
	}
	return false

}
