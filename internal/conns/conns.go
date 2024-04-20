package conns

import "context"

type (
	contextKeyType int
)

var (
	contextKey contextKeyType
)

// InContext represents the resource information kept in Context.
type InContext struct {
	IsDataSource       bool   // Data source?
	ResourceName       string // Friendly resource name, e.g. "Subnet"
	ServicePackageName string // Canonical name defined as a constant in names package
}

func NewDataSourceContext(ctx context.Context, servicePackageName, resourceName string) context.Context {
	v := InContext{
		IsDataSource:       true,
		ResourceName:       resourceName,
		ServicePackageName: servicePackageName,
	}

	return context.WithValue(ctx, contextKey, &v)
}

func FromContext(ctx context.Context) (*InContext, bool) {
	v, ok := ctx.Value(contextKey).(*InContext)
	return v, ok
}
