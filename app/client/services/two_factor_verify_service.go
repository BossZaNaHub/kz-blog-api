package clientservices

import (
	"github.com/gofiber/fiber/v2"
	clientmodels "github.com/kuroshibaz/app/client/models"
)

func (svc *defaultService) TwoFactorVerify(data clientmodels.TwoFactorVerifyData) *fiber.Error {
	return nil
}
