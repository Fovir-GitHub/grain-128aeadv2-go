package service

type Service struct {
	KeyManagement *KeyManagementService
	Cipher        *CipherService
}

func New() *Service {
	return &Service{
		KeyManagement: newKeyManagement(),
		Cipher:        newCipherService(),
	}
}
