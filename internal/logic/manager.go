package logic

type Manager struct {
	chatroom   map[string]*Hub
	register   chan *Hub
	unregister chan *Hub
}

func NewManager() *Manager {
	return &Manager{
		chatroom:   make(map[string]*Hub),
		register:   make(chan *Hub),
		unregister: make(chan *Hub),
	}
}

func (m *Manager) NewHub(name string, uid uint64, toUid uint64) *Hub {
	var hub *Hub
	if _, ok := m.chatroom[name]; ok == false {
		hub = NewHub(name, uid, toUid, m)
		go hub.Run()

		m.register <- hub

		return hub
	}

	return m.chatroom[name]
}

func (m *Manager) Run() {
	for {
		select {
		case hub := <-m.register:
			m.chatroom[hub.name] = hub
		case hub := <-m.unregister:
			if _, ok := m.chatroom[hub.name]; ok {
				delete(m.chatroom, hub.name)
				close(hub.broadcast)
			}
		}
	}
}
