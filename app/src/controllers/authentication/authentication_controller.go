package authentication

//var validate = validator.New()

type authentication struct{}

type AuthenticationControllerInterface interface {
}

func GetAuthenticationServiceInstance() AuthenticationControllerInterface {
	return new(authentication)
}
