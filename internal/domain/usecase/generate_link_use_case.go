package usecase

import (
	"errors"
	"fmt"
	"time"

	// <- obrigatório para rand.Intn
	"github.com/AlderFurtado/passlink/internal/domain/entity"
	"github.com/AlderFurtado/passlink/internal/repository"
	"github.com/AlderFurtado/passlink/internal/utils"
)

func GenerateLinkUseCase(origin string, isPaid bool) (string, error) {
	links := repository.FindAll()
	if matchOrigin(links, origin) {
		return "", errors.New("link is already used")
	}

	destiny := getNewLink(links)
	var validate time.Time
	if isPaid {
		validate = time.Now().AddDate(1, 0, 0)
	} else {
		validate = time.Now().AddDate(0, 0, 1)
	}
	newLink := entity.Link{Origin: origin, Destiny: destiny, Validate: validate}
	repository.InsertNewLink(newLink)
	return destiny, nil
}

func matchOrigin(links []entity.Link, origin string) bool {
	for _, v := range links {
		if v.Origin == origin {
			return true
		}
	}
	return false
}

func getNewLink(links []entity.Link) string {
	newLink := generateRandomHttpLink()
	if matchDestiny(links, newLink) {
		return getNewLink(links)
	}
	return newLink
}

func matchDestiny(links []entity.Link, newLink string) bool {
	for _, v := range links {
		if v.Destiny == newLink {
			return true
		}
	}
	return false
}

func generateRandomHttpLink() string {
	return fmt.Sprintf("http://localhost:8080/get/%v", utils.RandomWord(6))
}
