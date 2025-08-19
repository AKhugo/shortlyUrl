package service

import (
	"errors"
	"shortlyurl/internal/domain"
	"shortlyurl/internal/repository"
	"shortlyurl/internal/utils"
	"time"

)

type LinkService struct {
	linkRepo  repository.LinkRepository
	clickRepo repository.ClickRepository
}

func NewLinkService(linkRepo repository.LinkRepository, clickRepo repository.ClickRepository) *LinkService {
	return &LinkService{
		linkRepo:  linkRepo,
		clickRepo: clickRepo,
	}
}

// CreateShortLink crée un nouveau lien raccourci
func (s *LinkService) CreateShortLink(originalURL string, userID uint, isPublic bool) (*domain.Link, error) {
	// Validation de l'URL
	if originalURL == "" {
		return nil, errors.New("original URL cannot be empty")
	}

	// Vérifier si l'URL est déjà raccourcie par cet utilisateur
	existingLinks, err := s.linkRepo.GetLinksByUserID(userID)
	if err == nil {
		for _, link := range existingLinks {
			if link.OriginalURL == originalURL {
				return nil, errors.New("you already have a short link for this URL")
			}
		}
	}

	// Générer un short code unique
	shortCode := utils.GenerateUniqueShortCode(s.linkRepo)

	// Créer le lien
	link := &domain.Link{
		UserID:      userID,
		ShortCode:   shortCode,
		OriginalURL: originalURL,
		AccessLimit: 0,                // Pas de limite par défaut
		Domain:      "shortlyurl.com", // Domaine par défaut
		IsActive:    true,
		IsPublic:    isPublic,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Sauvegarder en base
	if err := s.linkRepo.CreateLink(link); err != nil {
		return nil, err
	}

	return link, nil
}

// GetLinkByShortCode récupère un lien par son code court
func (s *LinkService) GetLinkByShortCode(shortCode string) (*domain.Link, error) {
	link, err := s.linkRepo.GetLinkByShortCode(shortCode)
	if err != nil {
		return nil, errors.New("link not found or inactive")
	}

	// Vérifier la limite d'accès
	if link.AccessLimit > 0 {
		clicks, err := s.clickRepo.GetClicksByLinkID(link.ID)
		if err == nil && len(clicks) >= link.AccessLimit {
			return nil, errors.New("link access limit reached")
		}
	}

	return link, nil
}

// AccessLink accède à un lien et enregistre le clic
func (s *LinkService) AccessLink(shortCode string, ipAddress, userAgent, referer string) (*domain.Link, error) {
	// Récupérer le lien
	link, err := s.GetLinkByShortCode(shortCode)
	if err != nil {
		return nil, err
	}

	// Créer un enregistrement de clic
	click := &domain.Click{
		LinkID:      link.ID,
		ClickedAt:   time.Now(),
		IPAddress:   ipAddress,
		UserAgent:   userAgent,
		Referer:     referer,
		CountryCode: "", // À implémenter avec géolocalisation
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Enregistrer le clic
	if err := s.clickRepo.CreateClick(click); err != nil {
		// Log l'erreur mais ne pas bloquer l'accès
		// TODO: Logger l'erreur
	}

	return link, nil
}

// GetUserLinks récupère tous les liens d'un utilisateur
func (s *LinkService) GetUserLinks(userID uint) ([]*domain.Link, error) {
	return s.linkRepo.GetLinksByUserID(userID)
}

// UpdateLink met à jour un lien
func (s *LinkService) UpdateLink(linkID uint, userID uint, updates map[string]interface{}) error {
	// Vérifier que l'utilisateur possède le lien
	link, err := s.linkRepo.GetLinkByID(linkID)
	if err != nil {
		return errors.New("link not found")
	}

	if link.UserID != userID {
		return errors.New("unauthorized to modify this link")
	}

	// Appliquer les mises à jour
	if originalURL, ok := updates["original_url"].(string); ok {
		link.OriginalURL = originalURL
	}
	if isActive, ok := updates["is_active"].(bool); ok {
		link.IsActive = isActive
	}
	if isPublic, ok := updates["is_public"].(bool); ok {
		link.IsPublic = isPublic
	}
	if accessLimit, ok := updates["access_limit"].(int); ok {
		link.AccessLimit = accessLimit
	}

	link.UpdatedAt = time.Now()

	return s.linkRepo.UpdateLink(link)
}

// DeleteLink supprime un lien
func (s *LinkService) DeleteLink(linkID uint, userID uint) error {
	// Vérifier que l'utilisateur possède le lien
	link, err := s.linkRepo.GetLinkByID(linkID)
	if err != nil {
		return errors.New("link not found")
	}

	if link.UserID != userID {
		return errors.New("unauthorized to delete this link")
	}

	return s.linkRepo.DeleteLink(linkID)
}

// GetLinkStats récupère les statistiques d'un lien
func (s *LinkService) GetLinkStats(linkID uint, userID uint, filter *domain.ClickFilter) ([]*domain.Click, error) {
	// Vérifier que l'utilisateur possède le lien
	link, err := s.linkRepo.GetLinkByID(linkID)
	if err != nil {
		return nil, errors.New("link not found")
	}

	if link.UserID != userID {
		return nil, errors.New("unauthorized to view this link's stats")
	}

	return s.clickRepo.GetClickStats(linkID, filter)
}

// generateUniqueShortCode génère un code court unique
