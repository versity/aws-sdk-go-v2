// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Points to an S3Destination object that contains information about your S3
// bucket.
type Destination struct {

	// The S3Destination object.
	S3 *S3Destination

	noSmithyDocumentSerde
}

// The encryption algorithm options that are available to a code-signing job.
type EncryptionAlgorithmOptions struct {

	// The set of accepted encryption algorithms that are allowed in a code-signing
	// job.
	//
	// This member is required.
	AllowedValues []EncryptionAlgorithm

	// The default encryption algorithm that is used by a code-signing job.
	//
	// This member is required.
	DefaultValue EncryptionAlgorithm

	noSmithyDocumentSerde
}

// The hash algorithms that are available to a code-signing job.
type HashAlgorithmOptions struct {

	// The set of accepted hash algorithms allowed in a code-signing job.
	//
	// This member is required.
	AllowedValues []HashAlgorithm

	// The default hash algorithm that is used in a code-signing job.
	//
	// This member is required.
	DefaultValue HashAlgorithm

	noSmithyDocumentSerde
}

// A cross-account permission for a signing profile.
type Permission struct {

	// An AWS Signer action permitted as part of cross-account permissions.
	Action *string

	// The AWS principal that has been granted a cross-account permission.
	Principal *string

	// The signing profile version that a permission applies to.
	ProfileVersion *string

	// A unique identifier for a cross-account permission statement.
	StatementId *string

	noSmithyDocumentSerde
}

// The name and prefix of the Amazon S3 bucket where AWS Signer saves your signed
// objects.
type S3Destination struct {

	// Name of the S3 bucket.
	BucketName *string

	// An S3 prefix that you can use to limit responses to those that begin with the
	// specified prefix.
	Prefix *string

	noSmithyDocumentSerde
}

// The Amazon S3 bucket name and key where Signer saved your signed code image.
type S3SignedObject struct {

	// Name of the S3 bucket.
	BucketName *string

	// Key name that uniquely identifies a signed code image in your bucket.
	Key *string

	noSmithyDocumentSerde
}

// Information about the Amazon S3 bucket where you saved your unsigned code.
type S3Source struct {

	// Name of the S3 bucket.
	//
	// This member is required.
	BucketName *string

	// Key name of the bucket object that contains your unsigned code.
	//
	// This member is required.
	Key *string

	// Version of your source image in your version enabled S3 bucket.
	//
	// This member is required.
	Version *string

	noSmithyDocumentSerde
}

// The validity period for a signing job.
type SignatureValidityPeriod struct {

	// The time unit for signature validity.
	Type ValidityType

	// The numerical value of the time unit for signature validity.
	Value int32

	noSmithyDocumentSerde
}

// Points to an S3SignedObject object that contains information about your signed
// code image.
type SignedObject struct {

	// The S3SignedObject .
	S3 *S3SignedObject

	noSmithyDocumentSerde
}

// The configuration of a signing operation.
type SigningConfiguration struct {

	// The encryption algorithm options that are available for a code-signing job.
	//
	// This member is required.
	EncryptionAlgorithmOptions *EncryptionAlgorithmOptions

	// The hash algorithm options that are available for a code-signing job.
	//
	// This member is required.
	HashAlgorithmOptions *HashAlgorithmOptions

	noSmithyDocumentSerde
}

// A signing configuration that overrides the default encryption or hash algorithm
// of a signing job.
type SigningConfigurationOverrides struct {

	// A specified override of the default encryption algorithm that is used in a
	// code-signing job.
	EncryptionAlgorithm EncryptionAlgorithm

	// A specified override of the default hash algorithm that is used in a
	// code-signing job.
	HashAlgorithm HashAlgorithm

	noSmithyDocumentSerde
}

// The image format of a AWS Signer platform or profile.
type SigningImageFormat struct {

	// The default format of a signing image.
	//
	// This member is required.
	DefaultFormat ImageFormat

	// The supported formats of a signing image.
	//
	// This member is required.
	SupportedFormats []ImageFormat

	noSmithyDocumentSerde
}

// Contains information about a signing job.
type SigningJob struct {

	// The date and time that the signing job was created.
	CreatedAt *time.Time

	// Indicates whether the signing job is revoked.
	IsRevoked bool

	// The ID of the signing job.
	JobId *string

	// The AWS account ID of the job invoker.
	JobInvoker *string

	// The AWS account ID of the job owner.
	JobOwner *string

	// The name of a signing platform.
	PlatformDisplayName *string

	// The unique identifier for a signing platform.
	PlatformId *string

	// The name of the signing profile that created a signing job.
	ProfileName *string

	// The version of the signing profile that created a signing job.
	ProfileVersion *string

	// The time when the signature of a signing job expires.
	SignatureExpiresAt *time.Time

	// A SignedObject structure that contains information about a signing job's signed
	// code image.
	SignedObject *SignedObject

	// A SigningMaterial object that contains the Amazon Resource Name (ARN) of the
	// certificate used for the signing job.
	SigningMaterial *SigningMaterial

	// A Source that contains information about a signing job's code image source.
	Source *Source

	// The status of the signing job.
	Status SigningStatus

	noSmithyDocumentSerde
}

// Revocation information for a signing job.
type SigningJobRevocationRecord struct {

	// A caller-supplied reason for revocation.
	Reason *string

	// The time of revocation.
	RevokedAt *time.Time

	// The identity of the revoker.
	RevokedBy *string

	noSmithyDocumentSerde
}

// The ACM certificate that is used to sign your code.
type SigningMaterial struct {

	// The Amazon Resource Name (ARN) of the certificates that is used to sign your
	// code.
	//
	// This member is required.
	CertificateArn *string

	noSmithyDocumentSerde
}

// Contains information about the signing configurations and parameters that are
// used to perform a code-signing job.
type SigningPlatform struct {

	// The category of a signing platform.
	Category Category

	// The display name of a signing platform.
	DisplayName *string

	// The maximum size (in MB) of code that can be signed by a signing platform.
	MaxSizeInMB int32

	// Any partner entities linked to a signing platform.
	Partner *string

	// The ID of a signing platform.
	PlatformId *string

	// Indicates whether revocation is supported for the platform.
	RevocationSupported bool

	// The configuration of a signing platform. This includes the designated hash
	// algorithm and encryption algorithm of a signing platform.
	SigningConfiguration *SigningConfiguration

	// The image format of a AWS Signer platform or profile.
	SigningImageFormat *SigningImageFormat

	// The types of targets that can be signed by a signing platform.
	Target *string

	noSmithyDocumentSerde
}

// Any overrides that are applied to the signing configuration of a signing
// platform.
type SigningPlatformOverrides struct {

	// A signing configuration that overrides the default encryption or hash algorithm
	// of a signing job.
	SigningConfiguration *SigningConfigurationOverrides

	// A signed image is a JSON object. When overriding the default signing platform
	// configuration, a customer can select either of two signing formats, JSONEmbedded
	// or JSONDetached . (A third format value, JSON , is reserved for future use.)
	// With JSONEmbedded , the signing image has the payload embedded in it. With
	// JSONDetached , the payload is not be embedded in the signing image.
	SigningImageFormat ImageFormat

	noSmithyDocumentSerde
}

// Contains information about the ACM certificates and signing configuration
// parameters that can be used by a given code signing user.
type SigningProfile struct {

	// The Amazon Resource Name (ARN) for the signing profile.
	Arn *string

	// The name of the signing platform.
	PlatformDisplayName *string

	// The ID of a platform that is available for use by a signing profile.
	PlatformId *string

	// The name of the signing profile.
	ProfileName *string

	// The version of a signing profile.
	ProfileVersion *string

	// The ARN of a signing profile, including the profile version.
	ProfileVersionArn *string

	// The validity period for a signing job created using this signing profile.
	SignatureValidityPeriod *SignatureValidityPeriod

	// The ACM certificate that is available for use by a signing profile.
	SigningMaterial *SigningMaterial

	// The parameters that are available for use by a Signer user.
	SigningParameters map[string]string

	// The status of a signing profile.
	Status SigningProfileStatus

	// A list of tags associated with the signing profile.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Revocation information for a signing profile.
type SigningProfileRevocationRecord struct {

	// The time when revocation becomes effective.
	RevocationEffectiveFrom *time.Time

	// The time when the signing profile was revoked.
	RevokedAt *time.Time

	// The identity of the revoker.
	RevokedBy *string

	noSmithyDocumentSerde
}

// An S3Source object that contains information about the S3 bucket where you
// saved your unsigned code.
type Source struct {

	// The S3Source object.
	S3 *S3Source

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
