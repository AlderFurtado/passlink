package usecase

import (
	"errors"

	"github.com/AlderFurtado/passlink/internal/domain/entity"
	"github.com/AlderFurtado/passlink/internal/repository"
)

func FindOriginLinkUseCase(destiny string) (string, error) {
	links := repository.FindAll()
	return findOrigin(links, destiny)

}

func findOrigin(links []entity.Link, destiny string) (string, error) {
	for _, v := range links {
		if v.Destiny == destiny {
			return v.Origin, nil
		}
	}
	return "", errors.New("origin not found")
}
