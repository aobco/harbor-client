package errors

const (
	// ErrRepositoryNameNotProvidedMsg is the error message for ErrProjectNameNotProvided error.
	ErrRepositoryNameNotProvidedMsg = "repository name not provided"
	// ErrRepositoryNoPermissionMsg is the error message for ErrRepositoryNoPermission error.
	ErrRepositoryNoPermissionMsg = "user does not have permission to the repository"
	// ErrArtifactIllegalArgsFormatMsg is the error message for ErrRepositoryIllegalArgsFormat error.
	ErrArtifactIllegalArgsFormatMsg = "illegal format of provided name value"
	// ErrArtifactNotFoundMsg is the error message for ErrRepositoryNotFound error.
	ErrArtifactNotFoundMsg = "artifact not found on server side"
	// ErrRepositoryInternalErrorsMsg is the error message for ErrRepositoryInternalErrors error.
	ErrRepositoryInternalErrorsMsg = "unexpected internal errors"
)

type (
	// ErrRepositoryNameNotProvided describes a missing project name.
	ErrRepositoryNameNotProvided struct{}
	ErrRepositoryNoPermission    struct{}
	ErrArtifactIllegalArgsFormat struct{}
	// ErrArtifactUnknownResource describes an error after requesting an unknown resource.
	ErrArtifactUnknownResource struct{}
	// ErrArtifactInternalErrors describes server-side internal errors.
	ErrArtifactInternalErrors struct{}
)

// Error returns the error message.
func (e *ErrRepositoryNameNotProvided) Error() string {
	return ErrRepositoryNameNotProvidedMsg
}

// Error returns the error message.
func (e *ErrRepositoryNoPermission) Error() string {
	return ErrRepositoryNoPermissionMsg
}

// Error returns the error message.
func (e *ErrArtifactIllegalArgsFormat) Error() string {
	return ErrArtifactIllegalArgsFormatMsg
}

// Error returns the error message.
func (e *ErrArtifactUnknownResource) Error() string {
	return ErrArtifactNotFoundMsg
}

// Error returns the error message.
func (e *ErrArtifactInternalErrors) Error() string {
	return ErrRepositoryInternalErrorsMsg
}
