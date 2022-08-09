package interfaces

type AccountServiceInterface interface {
	UserActivation(userUUID string) error
}
