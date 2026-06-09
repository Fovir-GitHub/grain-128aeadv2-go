package service

type Service struct {
	KeyManagement *KeyManagementService
	Encryption    *EncryptionService
}

func New() *Service {
	return &Service{
		KeyManagement: newKeyManagement(),
		Encryption:    newEncryptionService(),
	}
}
