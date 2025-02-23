package context

type StateKey string

type Context map[StateKey]any

const (
    RenderQueue StateKey = "render-queue"
)

func (context *Context) GetValue(key StateKey) (any, bool) {
    value, ok := (*context)[key]
    return value, ok
}

func (context *Context) SetValue(key StateKey, value any) {
    (*context)[key] = value
}

func (context *Context) HasValue(key StateKey, value any) bool {
    contextValue, ok := context.GetValue(key)
    if ok {
        return contextValue == value
    } else {
        return ok
    }
}
