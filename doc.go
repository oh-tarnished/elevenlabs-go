// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
)

// DocService contains methods and other services that help with interacting with
// the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocService] method instead.
type DocService struct {
	Options []option.RequestOption
}

// NewDocService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDocService(opts ...option.RequestOption) (r DocService) {
	r = DocService{}
	r.Options = opts
	return
}

// Redirect To Mintlify
func (r *DocService) Redirect(ctx context.Context, opts ...option.RequestOption) (res *DocRedirectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "docs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DocRedirectResponse = any
