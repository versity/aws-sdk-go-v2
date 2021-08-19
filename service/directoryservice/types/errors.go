// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Client authentication is not available in this region at this time.
type AccessDeniedException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An authentication error occurred.
type AuthenticationFailedException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *AuthenticationFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AuthenticationFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AuthenticationFailedException) ErrorCode() string             { return "AuthenticationFailedException" }
func (e *AuthenticationFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The certificate has already been registered into the system.
type CertificateAlreadyExistsException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *CertificateAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CertificateAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CertificateAlreadyExistsException) ErrorCode() string {
	return "CertificateAlreadyExistsException"
}
func (e *CertificateAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The certificate is not present in the system for describe or deregister
// activities.
type CertificateDoesNotExistException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *CertificateDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CertificateDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CertificateDoesNotExistException) ErrorCode() string {
	return "CertificateDoesNotExistException"
}
func (e *CertificateDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The certificate is being used for the LDAP security connection and cannot be
// removed without disabling LDAP security.
type CertificateInUseException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *CertificateInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CertificateInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CertificateInUseException) ErrorCode() string             { return "CertificateInUseException" }
func (e *CertificateInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The certificate could not be added because the certificate limit has been
// reached.
type CertificateLimitExceededException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *CertificateLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CertificateLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CertificateLimitExceededException) ErrorCode() string {
	return "CertificateLimitExceededException"
}
func (e *CertificateLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A client exception has occurred.
type ClientException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *ClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClientException) ErrorCode() string             { return "ClientException" }
func (e *ClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Region you specified is the same Region where the Managed Microsoft AD
// directory was created. Specify a different Region and try again.
type DirectoryAlreadyInRegionException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *DirectoryAlreadyInRegionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryAlreadyInRegionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryAlreadyInRegionException) ErrorCode() string {
	return "DirectoryAlreadyInRegionException"
}
func (e *DirectoryAlreadyInRegionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified directory has already been shared with this Amazon Web Services
// account.
type DirectoryAlreadySharedException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *DirectoryAlreadySharedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryAlreadySharedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryAlreadySharedException) ErrorCode() string {
	return "DirectoryAlreadySharedException"
}
func (e *DirectoryAlreadySharedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified directory does not exist in the system.
type DirectoryDoesNotExistException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *DirectoryDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryDoesNotExistException) ErrorCode() string             { return "DirectoryDoesNotExistException" }
func (e *DirectoryDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum number of directories in the region has been reached. You can use
// the GetDirectoryLimits operation to determine your directory limits in the
// region.
type DirectoryLimitExceededException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *DirectoryLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryLimitExceededException) ErrorCode() string {
	return "DirectoryLimitExceededException"
}
func (e *DirectoryLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified directory has not been shared with this Amazon Web Services
// account.
type DirectoryNotSharedException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *DirectoryNotSharedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryNotSharedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryNotSharedException) ErrorCode() string             { return "DirectoryNotSharedException" }
func (e *DirectoryNotSharedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified directory is unavailable or could not be found.
type DirectoryUnavailableException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *DirectoryUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryUnavailableException) ErrorCode() string             { return "DirectoryUnavailableException" }
func (e *DirectoryUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum allowed number of domain controllers per directory was exceeded. The
// default limit per directory is 20 domain controllers.
type DomainControllerLimitExceededException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *DomainControllerLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DomainControllerLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DomainControllerLimitExceededException) ErrorCode() string {
	return "DomainControllerLimitExceededException"
}
func (e *DomainControllerLimitExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified entity already exists.
type EntityAlreadyExistsException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *EntityAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntityAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntityAlreadyExistsException) ErrorCode() string             { return "EntityAlreadyExistsException" }
func (e *EntityAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified entity could not be found.
type EntityDoesNotExistException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *EntityDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntityDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntityDoesNotExistException) ErrorCode() string             { return "EntityDoesNotExistException" }
func (e *EntityDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The account does not have sufficient permission to perform the operation.
type InsufficientPermissionsException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InsufficientPermissionsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientPermissionsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientPermissionsException) ErrorCode() string {
	return "InsufficientPermissionsException"
}
func (e *InsufficientPermissionsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The certificate PEM that was provided has incorrect encoding.
type InvalidCertificateException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InvalidCertificateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCertificateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCertificateException) ErrorCode() string             { return "InvalidCertificateException" }
func (e *InvalidCertificateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Client authentication is already enabled.
type InvalidClientAuthStatusException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InvalidClientAuthStatusException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidClientAuthStatusException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidClientAuthStatusException) ErrorCode() string {
	return "InvalidClientAuthStatusException"
}
func (e *InvalidClientAuthStatusException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The LDAP activities could not be performed because they are limited by the LDAPS
// status.
type InvalidLDAPSStatusException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InvalidLDAPSStatusException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidLDAPSStatusException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidLDAPSStatusException) ErrorCode() string             { return "InvalidLDAPSStatusException" }
func (e *InvalidLDAPSStatusException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The NextToken value is not valid.
type InvalidNextTokenException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more parameters are not valid.
type InvalidParameterException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The new password provided by the user does not meet the password complexity
// requirements defined in your directory.
type InvalidPasswordException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InvalidPasswordException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPasswordException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPasswordException) ErrorCode() string             { return "InvalidPasswordException" }
func (e *InvalidPasswordException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified shared target is not valid.
type InvalidTargetException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InvalidTargetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTargetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTargetException) ErrorCode() string             { return "InvalidTargetException" }
func (e *InvalidTargetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum allowed number of IP addresses was exceeded. The default limit is
// 100 IP address blocks.
type IpRouteLimitExceededException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *IpRouteLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IpRouteLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IpRouteLimitExceededException) ErrorCode() string             { return "IpRouteLimitExceededException" }
func (e *IpRouteLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Client authentication setup could not be completed because at least one valid
// certificate must be registered in the system.
type NoAvailableCertificateException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *NoAvailableCertificateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoAvailableCertificateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoAvailableCertificateException) ErrorCode() string {
	return "NoAvailableCertificateException"
}
func (e *NoAvailableCertificateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception encountered while trying to access your Amazon Web Services
// organization.
type OrganizationsException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *OrganizationsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OrganizationsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OrganizationsException) ErrorCode() string             { return "OrganizationsException" }
func (e *OrganizationsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have reached the limit for maximum number of simultaneous Region
// replications per directory.
type RegionLimitExceededException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *RegionLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RegionLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RegionLimitExceededException) ErrorCode() string             { return "RegionLimitExceededException" }
func (e *RegionLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An exception has occurred in Directory Service.
type ServiceException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *ServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceException) ErrorCode() string             { return "ServiceException" }
func (e *ServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The maximum number of Amazon Web Services accounts that you can share with this
// directory has been reached.
type ShareLimitExceededException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *ShareLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ShareLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ShareLimitExceededException) ErrorCode() string             { return "ShareLimitExceededException" }
func (e *ShareLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum number of manual snapshots for the directory has been reached. You
// can use the GetSnapshotLimits operation to determine the snapshot limits for a
// directory.
type SnapshotLimitExceededException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *SnapshotLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapshotLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapshotLimitExceededException) ErrorCode() string             { return "SnapshotLimitExceededException" }
func (e *SnapshotLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum allowed number of tags was exceeded.
type TagLimitExceededException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *TagLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagLimitExceededException) ErrorCode() string             { return "TagLimitExceededException" }
func (e *TagLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation is not supported.
type UnsupportedOperationException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *UnsupportedOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedOperationException) ErrorCode() string             { return "UnsupportedOperationException" }
func (e *UnsupportedOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The user provided a username that does not exist in your directory.
type UserDoesNotExistException struct {
	Message *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *UserDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserDoesNotExistException) ErrorCode() string             { return "UserDoesNotExistException" }
func (e *UserDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
