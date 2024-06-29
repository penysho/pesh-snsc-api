package controller

type Server struct {
	InteractProvider
}

func NewServer(interactProvider InteractProvider) *Server {
	return &Server{
		InteractProvider: interactProvider,
	}
}
