package artifact

import (
	"context"
	v2client "github.com/aobco/harbor-client/v5/apiv2/internal/api/client"
	"github.com/aobco/harbor-client/v5/apiv2/internal/api/client/artifact"
	"github.com/aobco/harbor-client/v5/apiv2/model"
	"github.com/aobco/harbor-client/v5/apiv2/pkg/config"
	clienterrors "github.com/aobco/harbor-client/v5/apiv2/pkg/errors"
	"github.com/go-openapi/runtime"
)

// RESTClient is a subclient for handling project related actions.
type RESTClient struct {
	// Options contains optional configuration when making API calls.
	Options *config.Options

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	DeleteArtifact(ctx context.Context, projectName, repositoryName, reference string) error
}

// DeleteArtifact deletes the specified tag.
// Returns an error when no matching project is found or when
// having difficulties talking to the API.
func (c *RESTClient) DeleteArtifact(ctx context.Context, projectName, repositoryName, reference string) error {
	if projectName == "" {
		return &clienterrors.ErrProjectNameNotProvided{}
	}
	if repositoryName == "" {
		return &clienterrors.ErrRepositoryNameNotProvided{}
	}

	params := &artifact.DeleteArtifactParams{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Reference:      reference,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Artifact.DeleteArtifact(params, c.AuthInfo)

	return handleSwaggerArtifactErrors(err)
}

func (c *RESTClient) Artifacts(ctx context.Context, projectName, repositoryName string) ([]*model.Artifact, error) {
	if projectName == "" {
		return nil, &clienterrors.ErrProjectNameNotProvided{}
	}
	if repositoryName == "" {
		return nil, &clienterrors.ErrRepositoryNameNotProvided{}
	}
	pushTime := "-push_time"
	withTag := true
	params := &artifact.ListArtifactsParams{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Context:        ctx,
		Sort:           &pushTime,
		WithTag:        &withTag,
	}

	params.WithTimeout(c.Options.Timeout)

	listArtifactsOK, err := c.V2Client.Artifact.ListArtifacts(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerArtifactErrors(err)
	}
	return listArtifactsOK.GetPayload(), nil
}
