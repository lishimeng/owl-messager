package model

// device token:接收者为
// Topic: title
// Payload: params json

// ApnsSenderInfo apns发送账号，
type ApnsSenderInfo struct {
	SenderInfo
	BundleId     string `orm:"column(bundle_id)"`     // app bundle id
	Cert         string `orm:"column(cert)"`          // 证书 base64
	CertPassword string `orm:"column(cert_password)"` // 证书密码， “”为空
	AppKey       string `orm:"column(app_key);null"`
	AppSecret    string `orm:"column(app_secret)"`
}
