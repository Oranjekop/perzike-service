package route

import routeauth "perzike-service/route/auth"

type AuthorizedPrincipal = routeauth.AuthorizedPrincipal
type KeyManager = routeauth.KeyManager

func GetKeyManager() *KeyManager {
	return routeauth.GetKeyManager()
}

func InitKeyManager(keyDir string) error {
	return routeauth.InitKeyManager(keyDir)
}
