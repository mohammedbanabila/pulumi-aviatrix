# Aviatrix Resource Provider

The Aviatrix Resource Provider lets you manage [Aviatrix](https://aviatrix.com) resources.

## Installing

This package is currently available for the following platforms/languages:

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-aviatrix
```

## Configuration

The following configuration points are available for the `aviatrix` provider:

Mandatory:
- `aviatrix:controllerIp` (environment: `AVIATRIX_CONTROLLERIP`) - The Aviatrix controller's IP address
- `aviatrix:username` (environment: `AVIATRIX_USERNAME`) - The Aviatrix controller's username
- `aviatrix:password` (environment: `AVIATRIX_PASSWORD`) - The Aviatrix controller's password

Optional:
- `aviatrix:skip_version_validation` (environment: `AVIATRIX_SKIP_VERSION_VALIDATION`) - Valid values: true, false. Default: false. If set to true, it skips checking whether current Pulumi provider supports current Controller version.
- `aviatrix:verify_ssl_certificate` (environment: `AVIATRIX_VERIFY_SSL_CERTIFICATE`) - Valid values: true, false. Default: false. If set to true, the SSL certificate of the controller will be verified.

## Reference

For detailed reference documentation, please visit [the Terraform registry](https://registry.terraform.io/providers/AviatrixSystems/aviatrix/latest/docs).
