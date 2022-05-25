package keygen

import (
	"errors"
	"fmt"
)

// ErrorCode defines various error codes that are handled explicitly.
type ErrorCode string

const (
	ErrorCodeTokenInvalid         ErrorCode = "TOKEN_INVALID"
	ErrorCodeLicenseInvalid       ErrorCode = "LICENSE_INVALID"
	ErrorCodeFingerprintTaken     ErrorCode = "FINGERPRINT_TAKEN"
	ErrorCodeMachineLimitExceeded ErrorCode = "MACHINE_LIMIT_EXCEEDED"
	ErrorCodeProcessLimitExceeded ErrorCode = "MACHINE_PROCESS_LIMIT_EXCEEDED"
	ErrorCodeMachineHeartbeatDead ErrorCode = "MACHINE_HEARTBEAT_DEAD"
	ErrorCodeProcessHeartbeatDead ErrorCode = "PROCESS_HEARTBEAT_DEAD"
	ErrorCodeNotFound             ErrorCode = "NOT_FOUND"
)

// Error represents an API error response.
type Error struct {
	Response *Response
	Title    string
	Detail   string
	Code     string
}

func (e *Error) Error() string {
	res := e.Response

	return fmt.Sprintf("an error occurred: id=%s status=%d size=%d body=%s", res.ID, res.Status, res.Size, res.Body)
}

// LicenseTokenInvalidError represents an API authentication error due to an invalid license token.
type LicenseTokenInvalidError struct{ Err *Error }

func (e *LicenseTokenInvalidError) Error() string { return e.Err.Detail }

// LicenseKeyInvalidError represents an API authentication error due to an invalid license key.
type LicenseKeyInvalidError struct{ Err *Error }

func (e *LicenseKeyInvalidError) Error() string { return e.Err.Detail }

// NotAuthorizedError represents an API permission error.
type NotAuthorizedError struct{ Err *Error }

func (e *NotAuthorizedError) Error() string { return e.Err.Detail }

// NotFoundError represents an API not found error.
type NotFoundError struct{ Err *Error }

func (e *NotFoundError) Error() string { return e.Err.Detail }

// InvalidLicenseFileError represents an invalid license file error.
type InvalidLicenseFileError struct{ Err error }

func (e *InvalidLicenseFileError) Error() string { return e.Err.Error() }

// InvalidMachineFileError represents an invalid machine file error.
type InvalidMachineFileError struct{ Err error }

func (e *InvalidMachineFileError) Error() string { return e.Err.Error() }

// General errors
var (
	ErrReleaseLocationMissing       = errors.New("release has no download URL")
	ErrUpgradeNotAvailable          = errors.New("no upgrades available (already up-to-date)")
	ErrResponseSignatureMissing     = errors.New("response signature is missing")
	ErrResponseSignatureInvalid     = errors.New("response signature is invalid")
	ErrResponseDigestMissing        = errors.New("response digest is missing")
	ErrResponseDigestInvalid        = errors.New("response digest is invalid")
	ErrResponseDateInvalid          = errors.New("response date is invalid")
	ErrResponseDateTooOld           = errors.New("response date is too old")
	ErrPublicKeyMissing             = errors.New("public key is missing")
	ErrPublicKeyInvalid             = errors.New("public key is invalid")
	ErrValidationFingerprintMissing = errors.New("validation fingerprint scope is missing")
	ErrValidationProductMissing     = errors.New("validation product scope is missing")
	ErrHeartbeatPingFailed          = errors.New("heartbeat ping failed")
	ErrHeartbeatRequired            = errors.New("heartbeat is required")
	ErrHeartbeatDead                = errors.New("heartbeat is dead")
	ErrMachineAlreadyActivated      = errors.New("machine is already activated")
	ErrMachineLimitExceeded         = errors.New("machine limit has been exceeded")
	ErrMachineNotFound              = errors.New("machine no longer exists")
	ErrProcessNotFound              = errors.New("process no longer exists")
	ErrMachineFileNotSupported      = errors.New("machine file is not supported")
	ErrMachineFileNotEncrypted      = errors.New("machine file is not encrypted")
	ErrMachineFileNotGenuine        = errors.New("machine file is not genuine")
	ErrProcessLimitExceeded         = errors.New("process limit has been exceeded")
	ErrLicenseSchemeNotSupported    = errors.New("license scheme is not supported")
	ErrLicenseSchemeMissing         = errors.New("license scheme is missing")
	ErrLicenseKeyMissing            = errors.New("license key is missing")
	ErrLicenseKeyNotGenuine         = errors.New("license key is not genuine")
	ErrLicenseNotActivated          = errors.New("license is not activated")
	ErrLicenseExpired               = errors.New("license is expired")
	ErrLicenseSuspended             = errors.New("license is suspended")
	ErrLicenseTooManyMachines       = errors.New("license has too many machines")
	ErrLicenseTooManyCores          = errors.New("license has too many cores")
	ErrLicenseNotSigned             = errors.New("license is not signed")
	ErrLicenseInvalid               = errors.New("license is invalid")
	ErrLicenseFileNotSupported      = errors.New("license file is not supported")
	ErrLicenseFileNotEncrypted      = errors.New("license file is not encrypted")
	ErrLicenseFileNotGenuine        = errors.New("license file is not genuine")
	ErrLicenseFileSecretMissing     = errors.New("license file secret is missing")
)
