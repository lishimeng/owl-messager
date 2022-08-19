package ms

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest"
	"github.com/jongio/azidext/go/azidext"
	"github.com/lishimeng/go-log"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type AzureAuthConfig struct {
	Sender         string
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

	h.printConfig()
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

	var in bytes.Buffer
	json.HTMLEscape(&in, bs)
	log.Info("send mail:\n%s", in.String())

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

	req, err := autorest.Prepare(&http.Request{},
		autorest.AsPost(),

		autorest.WithBaseURL("https://graph.microsoft.com/"),
		autorest.WithPath("v1.0"),
		autorest.WithPath("users"),
		autorest.WithPath(url.QueryEscape(h.Config.Sender)),
		autorest.WithPath("sendMail"),
		autorest.AsContentType("application/json"),
		a.WithAuthorization(),
		autorest.WithBytes(&bs),
	)
	if err != nil {
		log.Info("req failure: %+v", err)
		return
	}

	resp, err := autorest.Send(req,
		autorest.DoErrorIfStatusCode(http.StatusInternalServerError),
		autorest.DoCloseIfError(),
		autorest.DoRetryForAttempts(5, time.Second),
	)
	if err != nil {
		log.Info("resp failure: %+v", err)
		return
	}
	log.Info("request----------------------------------------------------")
	log.Info("req: [%s]%s", req.Method, req.URL.String())
	log.Info("Authorization: %s", req.Header.Get("Authorization"))
	log.Info("content-type: %s", req.Header.Get("content-type"))
	log.Info("response----------------------------------------------------")
	log.Info("status:%s[%d]", resp.Status, resp.StatusCode)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("resp body failure: %+v", err)
	}
	log.Info("body: %s", string(b))

	err = autorest.Respond(resp, autorest.ByDiscardingBody(), autorest.ByClosing())
	if err != nil {
		log.Info("resp failure: %+v", err)
		return
	}

	return
}

func (h *AzureGraphProvider) printConfig() {

	log.Info("---config info start---------------------")
	log.Info("%s: %s", tenantKey, h.Config.Tenant)
	log.Info("%s: %s", clientIdKey, h.Config.ClientId)
	log.Info("%s: %s", scopeKey, h.Config.Scope)
	log.Info("%s: %s", senderKey, h.Config.Sender)
	log.Info("---config info end---------------------")
}
