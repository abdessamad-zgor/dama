package dama

import ()

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
	return navigator
}

func getNavigationTree(elements []DamaElement) []NavigationItem {
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

	stack = append(stack, current_item)
	for len(stack) != 0 {
		node, n_stack := stack[len(stack)-1], append(stack[:len(stack)-1], stack[len(stack)-1:]...)
		stack = n_stack
		current_item = node

		parent := current_item.Parent
		prefix := ""
		for parent != nil {
			prefix = string(parent.Element.GetTag()) + prefix
		}
		navigator.Index = append(navigator.Index, IndexItem{prefix + string(current_item.Element.GetTag()), current_item})
		if current_item.Children != nil {
			for _, child := range *current_item.Children {
				stack = append(stack, child)
			}
		}
	}
}

func (navigator *Navigator) GetNavigationTree(elements []DamaElement) {
	rootChildren := getNavigationTree(elements)
	navigator.indexItems()
	navigator.Root.Children = &rootChildren
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
		if item.Path == navigator.CurrentPath+string(tag) || item.Path == navigator.CurrentPath[:len(navigator.CurrentPath)-1]+string(tag) || indexItem.Item.Element.GetTag() == tag {
			indexItem = &item
		}
	}

	return indexItem
}

func (navigator *Navigator) Navigate(tag rune) {
	indexItem := navigator.SearchIndex(tag)

	if indexItem != nil {
		navigator.Current.Element.Blur()
		navigator.Current = indexItem.Item
		navigator.CurrentPath = indexItem.Path
		navigator.Current.Element.Focus()
	}

}
