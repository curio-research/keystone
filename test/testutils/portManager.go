package testutils

type PortManager struct {
	port int
}

func NewPortManager() *PortManager {
	return &PortManager{
		port: 9000,
	}
}

func (p *PortManager) GetPort() int {
	port := p.port
	p.port++
	return port
}
