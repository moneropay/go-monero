package walletrpc

import "context"

type SetDaemonRequest struct {
	// (Optional) The URL of the daemon to connect to. (Default: "") 
	Address string `json:"address,omitempty"`

	// (Optional) If false, some RPC wallet methods will be disabled. (Default: false)
	Trusted bool `json:"trusted,omitempty"`

	// (Optional) Specifies whether the Daemon uses SSL encryption. (Default: autodetect; Accepts: disabled, enabled, autodetect) 
	SslSupport string `json:"ssl_support,omitempty"`

	// (Optional) The file path location of the SSL key.
	SslPrivateKeyPath string `json:"ssl_private_key_path,omitempty"`

	// (Optional) The file path location of the SSL certificate.
	SslCertificatePath string `json:"ssl_certificate_path,omitempty"`

	// (Optional) The SHA1 fingerprints accepted by the SSL certificate.
	SslAllowedFingerprints []string `json:"ssl_allowed_fingerprints,omitempty"`

	// (Optional) If false, the certificate must be signed by a trusted certificate authority. (Default: false)
	SslAllowAnyCert bool `json:"ssl_allow_any_cert,omitempty"`
}

// Connect the RPC server to a Monero daemon.
func (c *Client) SetDaemon(ctx context.Context, req *SetDaemonRequest) error {
	err := c.Do(ctx, "set_daemon", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
