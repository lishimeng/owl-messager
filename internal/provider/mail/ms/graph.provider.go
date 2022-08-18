package ms

import (
	"crypto"
	"crypto/x509"
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest"
	"github.com/jongio/azidext/go/azidext"
	"github.com/lishimeng/go-log"
	"net/http"
	"time"
)

type AzureAuthConfig struct {
	Tenant         string
	ClientId       string
	Scope          string
	Certificate    []*x509.Certificate
	CertificateKey crypto.PrivateKey
}

type AzureGraphProvider struct {
	Config AzureAuthConfig
}

func (h *AzureGraphProvider) Send(subject string, body string, receivers ...string) (err error) {

	msg, err := BuildMessage(subject, body, receivers...)
	if err != nil {
		log.Info("build message failure: %+v", err)
		return
	}

	bs, err := json.Marshal(msg)
	if err != nil {
		log.Info("build message failure: %+v", err)
		return
	}

	cred, err := azidentity.NewClientCertificateCredential(
		h.Config.Tenant,
		h.Config.ClientId,
		h.Config.Certificate,
		h.Config.CertificateKey,
		nil,
	)
	if err != nil {
		log.Info("Authentication failure: %+v", err)
	}
	a := azidext.NewTokenCredentialAdapter(cred, []string{"https://graph.microsoft.com/.default"})
	log.Info(a)

	req, err := autorest.Prepare(&http.Request{},
		autorest.WithBaseURL("https://graph.microsoft.com/"),
		autorest.WithPath("v1.0"),
		autorest.WithPath("users"),
		autorest.WithPath(h.Config.Tenant),
		autorest.WithPath("sendMail"),
		a.WithAuthorization(),
		autorest.WithBytes(&bs),
	)

	resp, err := autorest.Send(req,
		autorest.DoErrorIfStatusCode(http.StatusInternalServerError),
		autorest.DoCloseIfError(),
		autorest.DoRetryForAttempts(5, time.Second),
	)
	if err != nil {
		log.Info("req failure: %+v", err)
		return
	}
	err = autorest.Respond(resp, autorest.ByDiscardingBody(), autorest.ByClosing())
	if err != nil {
		log.Info("resp failure: %+v", err)
		return
	}
	return
}
