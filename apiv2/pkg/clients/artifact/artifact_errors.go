package artifact

import (
	"net/http"

	"github.com/go-openapi/runtime"

	projectapi "github.com/aobco/harbor-client/v5/apiv2/internal/api/client/project"
	"github.com/aobco/harbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerArtifactErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerArtifactErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusCreated:
			// Harbor sometimes returns 201 instead of 200 despite the swagger spec
			// not declaring it.
			return nil
		case http.StatusBadRequest:
			return &errors.ErrArtifactIllegalArgsFormat{}
		case http.StatusUnauthorized:
			return &errors.ErrUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrRepositoryNoPermission{}
		case http.StatusNotFound:
			return &errors.ErrArtifactUnknownResource{}
		case http.StatusInternalServerError:
			return &errors.ErrArtifactInternalErrors{}
		}
	}

	switch in.(type) {
	case *projectapi.DeleteProjectNotFound:
		return &errors.ErrProjectIDNotExists{}
	case *projectapi.UpdateProjectNotFound:
		return &errors.ErrProjectIDNotExists{}
	case *projectapi.CreateProjectConflict:
		return &errors.ErrProjectNameAlreadyExists{}
	default:
		return in
	}
}
