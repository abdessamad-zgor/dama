package dama

import ()

type NavigationItem struct {
	Parent   *NavigationItem
	Element  DamaElement
	Children *[]NavigationItem
}

type Navigator struct {
	Items       []NavigationItem
	Current     NavigationItem
	CurrentPath []rune
}

func NewNavigator() *Navigator {
	navigator := new(Navigator)
	navigator.Items = []NavigationItem{}
	navigator.CurrentPath = []rune{}
	return navigator
}

func indexItems(elements []DamaElement) []NavigationItem {
	navigationItems := []NavigationItem{}
	for _, element := range elements {
		navigationItem := NavigationItem{nil, element, nil}
		container, ok := element.(*Container)
		if ok {
			itemNavigables := container.GetNavigables()
			itemChildren := indexItems(itemNavigables)
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

func (navigator *Navigator) IndexItems(elements []DamaElement) {
	navigator.Items = indexItems(elements)
}

func (navigator *Navigator) SetupKeybindings() []rune {
	return nil
}
