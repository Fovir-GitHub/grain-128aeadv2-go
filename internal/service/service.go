package service

type Service struct {
	KeyManagement *KeyManagementService
}

func New() *Service {
	return &Service{
		KeyManagement: newKeyManagement(),
	}
}
