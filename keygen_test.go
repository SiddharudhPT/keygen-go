package keygen

import (
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestValidate(t *testing.T) {
	Account = os.Getenv("KEYGEN_ACCOUNT")
	Product = os.Getenv("KEYGEN_PRODUCT")
	Token = os.Getenv("KEYGEN_TOKEN")

	fingerprint := uuid.New().String()
	license, err := Validate(fingerprint)
	switch {
	case err == ErrLicenseNotActivated:
		machine, err := license.Activate(fingerprint)
		if err != nil {
			t.Fatalf("Should not fail activation: err=%v", err)
		}

		go machine.Monitor()

		machines, err := license.Machines()
		if err != nil {
			t.Fatalf("Should not fail to list machines: err=%v", err)
		}

		for _, machine := range machines {
			err = machine.Deactivate()
			if err != nil {
				t.Fatalf("Should not fail deactivation: err=%v", err)
			}
		}

		err = license.Deactivate(fingerprint)
		switch {
		case err == nil:
			t.Fatalf("Should not be deactivated again: license=%v fingerprint=%v", license, fingerprint)
		case err != ErrNotFound:
			t.Fatalf("Should already be deactivated: err=%v", err)
		}

		entitlements, err := license.Entitlements()
		if err != nil {
			t.Fatalf("Should not fail to list entitlements: err=%v", err)
		}

		t.Logf("license=%v machines=%v entitlements=%v", license, machines, entitlements)
	case err != nil:
		t.Fatalf("Should not fail validation: err=%v", err)
	case err == nil:
		t.Fatalf("Should not be activated: err=%v", err)
	}
}

func TestUpgrade(t *testing.T) {
	Account = os.Getenv("KEYGEN_ACCOUNT")
	Product = os.Getenv("KEYGEN_PRODUCT")
	Token = os.Getenv("KEYGEN_TOKEN")

	upgrade, err := Upgrade("1.0.0")
	switch {
	case err == ErrUpgradeNotAvailable:
		t.Fatalf("Should have an upgrade available: err=%v", err)
	case err != nil:
		t.Fatalf("Should not fail upgrade: err=%v", err)
	}

	if upgrade.Location == "" {
		t.Fatalf("Should have a download URL: upgrade=%v", upgrade)
	}

	t.Logf("upgrade=%v", upgrade)
}
