package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
)

func (repo *defaultRepository) UpdateTwoFactor(enabled bool) *fiber.Error {
	err := repo.orm.UpdateTFAColumn(enabled)
	if err != nil {
		return fiber.NewError(errors.ErrCodeInternalServer, err.Error())
	}

	return nil
}
