package services

// DispatcherEventData общий интерфейс данных для обработчиков
type DispatcherEventData struct {
	ID   int
	Data chan interface{}
}

// DispatcherEvents Для работы с несколькими подписчиками
type DispatcherEvents struct {
	handlers []DispatcherEventData
}

// AddHandler добавление обработчика
func (c *DispatcherEvents) AddHandler(data chan interface{}) int {
	d := DispatcherEventData{ID: len(c.handlers), Data: data}
	c.handlers = append(c.handlers, d)

	return d.ID
}

// RemoveHandler удаление обработчика
func (c *DispatcherEvents) RemoveHandler(id int) {
	for index, it := range c.handlers {
		if it.ID == id {
			c.handlers = append(c.handlers[:index], c.handlers[index+1:]...)
		}
	}
}

// SendMessageToHandlers добавление обработчика
func (c *DispatcherEvents) SendToHandlers(data interface{}) {
	for _, it := range c.handlers {
		it.Data <- data
	}
}
