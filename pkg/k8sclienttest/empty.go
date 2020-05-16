package k8sclienttest

func NewEmpty() *Clients {
	return NewClients(ClientsConfig{})
}
