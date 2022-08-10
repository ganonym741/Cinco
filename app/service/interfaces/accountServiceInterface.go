package interfaces

type AccountServiceInterface interface {
	CreateAccount(userUUID string) error
}
