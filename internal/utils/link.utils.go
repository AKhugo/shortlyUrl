package utils

import (
	"github.com/google/uuid"
	"shortlyurl/internal/repository"
)

type LinkService struct {
	linkRepo  repository.LinkRepository
	clickRepo repository.ClickRepository
}


func GenerateUniqueShortCode(linkRepo repository.LinkRepository) string {
	maxAttempts := 10
	for i := 0; i < maxAttempts; i++ {
		// Générer un code de 6 caractères
		shortCode := uuid.New().String()[:6]

		// Vérifier l'unicité
		if _, err := linkRepo.GetLinkByShortCode(shortCode); err != nil {
			return shortCode
		}
	}

	// Fallback avec timestamp si trop de collisions
	return uuid.New().String()[:8]
}