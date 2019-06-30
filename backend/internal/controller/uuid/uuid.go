package uuid

import (
	"github.com/gofrs/uuid"
	"github.com/bonnetn/dare/backend/internal/entity"
	"fmt"
	"strings"
)

var _namespace = uuid.NewV5(uuid.NamespaceDNS, "github.com/bonnetn")

type UUIDController interface {
	IsValid(string) bool
	GenerateUUIDFromNonce(nonce string, salt string) entity.UUID
	Normalize(string) (string, error)
}

func NewUUIDController() UUIDController {
	return &uuidController{}
}

type uuidController struct{}

func (uuidController) Normalize(uuidStr string) (string, error) {
	myUUID, err := uuid.FromString(uuidStr)
	if err != nil {
		return "", fmt.Errorf("could not normalize UUID: %v", err)
	}

	return strings.ToLower(myUUID.String()), nil
}

func (uuidController) IsValid(uuidStr string) bool {
	_, err := uuid.FromString(uuidStr)
	return err == nil
}

func (uuidController) GenerateUUIDFromNonce(nonce string, salt string) entity.UUID {
	saltUUID := uuid.NewV5(_namespace, salt)
	return uuid.NewV5(saltUUID, nonce).String()
}
