package utils

import "github.com/gdamore/tcell/v2"

type Screen struct {
    Content [][]rune
}

func (screen *Screen) Init() error {
    for i := range(90) {
        screen.Content = append(screen.Content, []rune{})
        for _ = range(160) {
            screen.Content[i] = append(screen.Content[i], rune(0))
        }
    }
    return nil
}

func (screen *Screen) Fini() {

}

func (screen *Screen) Clear() {

}

func (screen *Screen) Fill(rune, tcell.Style) {

}

func (screen *Screen) SetCell(x int, y int, style tcell.Style, ch ...rune) {

}

func (screen *Screen) GetContent(x, y int) (primary rune, combining []rune, style tcell.Style, width int) {
    return screen.Content[x][y], nil, tcell.StyleDefault, 1
}

func (screen *Screen) SetContent(x int, y int, primary rune, combining []rune, style tcell.Style) {
    _ = combining
    _ = style
    screen.Content[x][y] = primary
}

func (screen *Screen) SetStyle(style tcell.Style) {

}

func (screen *Screen) ShowCursor(x int, y int) {

}

func (screen *Screen) HideCursor() {

}

func (screen *Screen) SetCursorStyle(tcell.CursorStyle, ...tcell.Color) {

}

func (screen *Screen) Size() (width, height int) {
    return len(screen.Content), len(screen.Content[0])
}

func (screen *Screen) ChannelEvents(ch chan<- tcell.Event, quit <-chan struct{}) {

}

func (screen *Screen) PollEvent() tcell.Event {
    return *new(tcell.Event)
}

func (screen *Screen) HasPendingEvent() bool {
    return false
}

func (screen *Screen) PostEvent(ev tcell.Event) error {
    return nil
}

func (screen *Screen) PostEventWait(ev tcell.Event) {

}

func (screen *Screen) EnableMouse(...tcell.MouseFlags) {

}

func (screen *Screen) DisableMouse() {

}

func (screen *Screen) EnablePaste() {

}

func (screen *Screen) DisablePaste() {

}

func (screen *Screen) EnableFocus() {

}

func (screen *Screen) DisableFocus() {

}

func (screen *Screen) HasMouse() bool {
    return false
}

func (screen *Screen) Colors() int {
    return 256
}

func (screen *Screen) Show() {

}

func (screen *Screen) Sync() {

}

func (screen *Screen) CharacterSet() string {
    return ""
}

func (screen *Screen) RegisterRuneFallback(r rune, subst string) {

}

func (screen *Screen) UnregisterRuneFallback(r rune) {

}

func (screen *Screen) CanDisplay(r rune, checkFallbacks bool) bool {
    return false
}

func (screen *Screen) Resize(int, int, int, int) {

}

func (screen *Screen) HasKey(tcell.Key) bool {
    return false
}

func (screen *Screen) Suspend() error {
    return nil
}

func (screen *Screen) Resume() error {
    return nil
}

func (screen *Screen) Beep() error {
    return nil
}

func (screen *Screen) SetSize(int, int) {

}

func (screen *Screen) LockRegion(x, y, width, height int, lock bool) {

}

func (screen *Screen) Tty() (tcell.Tty, bool) {
    tty := new(tcell.Tty)
    return *tty, true
}

func (screen *Screen) SetTitle(string) {

}

func (screen *Screen) SetClipboard([]byte) {

}

func (screen *Screen) GetClipboard() {

}
