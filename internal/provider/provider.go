package provider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rstuhlmuller/terraform-provider-casaos/internal/conns"
)

func New(ctx context.Context) (*schema.Provider, error) {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The hostname of the CasaOS device. (default: http://casaos.local)",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username to authenticate with the CasaOS device.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The password to authenticate with the CasaOS device.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   map[string]*schema.Resource{},
	}
	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return configure(ctx, provider, d)
	}

	var errs []error

	if err := errors.Join(errs...); err != nil {
		return nil, err
	}

	// Set the provider Meta (instance data) here.
	// It will be overwritten by the result of the call to ConfigureContextFunc,
	// but can be used pre-configuration by other (non-primary) provider servers.
	var meta *conns.CasaOSClient
	if v, ok := provider.Meta().(*conns.CasaOSClient); ok {
		meta = v
	} else {
		meta = new(conns.CasaOSClient)
	}
	meta.ServicePackages = servicePackageMap
	provider.SetMeta(meta)

	return provider, nil
}

func configure(ctx context.Context, provider *schema.Provider, d *schema.ResourceData) (*conns.CasaOSClient, diag.Diagnostics) {
	var diags diag.Diagnostics

	terraformVersion := provider.TerraformVersion
	if terraformVersion == "" {
		// Terraform 0.12 introduced this field to the protocol
		// We can therefore assume that if it's missing it's 0.10 or 0.11
		terraformVersion = "0.11+compatible"
	}

	config := conns.Config{
		Host:     d.Get("host").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}

	var meta *conns.CasaOSClient
	if v, ok := provider.Meta().(*conns.CasaOSClient); ok {
		meta = v
	} else {
		meta = new(conns.CasaOSClient)
	}
	meta, ds := config.ConfigureProvider(ctx, meta)
	diags = append(diags, ds...)
	if diags.HasError() {
		return nil, diags
	}

	return meta, diags
}
