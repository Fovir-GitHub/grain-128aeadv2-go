package service

type Service struct {
	KeyManagement *KeyManagementService
	Cipher        *CipherService
}

// Register services.
func New() *Service {
	return &Service{
		KeyManagement: newKeyManagement(),
		Cipher:        newCipherService(),
	}
}
