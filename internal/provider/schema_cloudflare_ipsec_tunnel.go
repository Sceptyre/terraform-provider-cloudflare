package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareIPsecTunnelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the IPsec tunnel.",
		},
		"customer_endpoint": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "IP address assigned to the customer side of the IPsec tunnel.",
		},
		"cloudflare_endpoint": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "IP address assigned to the Cloudflare side of the IPsec tunnel.",
		},
		"interface_address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "31-bit prefix (/31 in CIDR notation) supporting 2 hosts, one for each side of the tunnel.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An optional description of the IPsec tunnel.",
		},
		"health_check_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Specifies if ICMP tunnel health checks are enabled. Default: `true`.",
		},
		"health_check_target": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The IP address of the customer endpoint that will receive tunnel health checks. Default: `<customer_gre_endpoint>`.",
		},
		"health_check_type": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.StringInSlice([]string{"request", "reply"}, false),
			Description:  fmt.Sprintf("Specifies the ICMP echo type for the health check (`request` or `reply`). %s Default: `reply`.", renderAvailableDocumentationValuesStringSlice([]string{"request", "reply"})),
		},
		"psk": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Sensitive:   true,
			Description: "Pre shared key to be used with the IPsec tunnel. If left unset, it will be autogenerated.",
		},
		"allow_null_cipher": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Specifies if this tunnel may use a null cipher (ENCR_NULL) in Phase 2.",
		},
		"hex_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "`remote_id` as a hex string. This value is generated by cloudflare.",
		},
		"user_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "`remote_id` in the form of an email address. This value is generated by cloudflare.",
		},
		"fqdn_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "`remote_id` in the form of a fqdn. This value is generated by cloudflare.",
		},
		"remote_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "ID to be used while setting up the IPsec tunnel. This value is generated by cloudflare.",
		},
	}
}
