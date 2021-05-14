package walletrpc

type SetDaemonRequest struct {
	// Address - string; (Optional; Default: "") The URL of the daemon to connect to.
	Address string `json:"address"`

	// Trusted - boolean; (Optional; Default: false) If false, some RPC wallet methods will be disabled.
	Trusted bool `json:"trusted"`

	// SslSupport - string; (Optional; Default: autodetect; Accepts: disabled, enabled, autodetect) Specifies whether the Daemon uses SSL encryption.
	SslSupport string `json:"ssl_support"`

	// SslPrivateKeyPath - string; (Optional) The file path location of the SSL key.
	SslPrivateKeyPath string `json:"ssl_private_key_path"`

	// SslCertificatePath - string; (Optional) The file path location of the SSL certificate.
	SslCertificatePath string `json:"ssl_certificate_path"`

	// SslAllowedFingerprints - array of string; (Optional) The SHA1 fingerprints accepted by the SSL certificate.
	SslAllowedFingerprints []string `json:"ssl_allowed_fingerprints"`

	// SslAllowAnyCert - boolean; (Optional; Default: false) If false, the certificate must be signed by a trusted certificate authority.
	SslAllowAnyCert bool `json:"ssl_allow_any_cert"`
}

// SetDaemon connect the RPC server to a Monero daemon.
func (c *Client) SetDaemon(req *SetDaemonRequest) error {
	err := c.Do("set_daemon", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
