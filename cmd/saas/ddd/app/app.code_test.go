package tenant

import "testing"

func TestGenTenantCode(t *testing.T) {
	appId := genAppId("grgwefdvb4geasda")
	secret := genSecret(appId)
	t.Logf("appId: %s", appId)
	t.Logf("secret: %s", secret)
}
