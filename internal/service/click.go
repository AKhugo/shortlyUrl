package service

import (
	"errors"
	"shortlyurl/internal/domain"
	"shortlyurl/internal/repository"
	"time"
)

type ClickService struct {
	clickRepo repository.ClickRepository
	linkRepo  repository.LinkRepository
}

func NewClickService(clickRepo repository.ClickRepository, linkRepo repository.LinkRepository) *ClickService {
	return &ClickService{
		clickRepo: clickRepo,
		linkRepo:  linkRepo,
	}
}

// RecordClick enregistre un nouveau clic sur un lien
func (s *ClickService) RecordClick(linkID uint, ipAddress, userAgent, referer string) error {
	// Vérifier que le lien existe et est actif
	link, err := s.linkRepo.GetLinkByID(linkID)
	if err != nil {
		return errors.New("link not found")
	}

	if !link.IsActive {
		return errors.New("link is inactive")
	}

	// Vérifier la limite d'accès
	if link.AccessLimit > 0 {
		clicks, err := s.clickRepo.GetClicksByLinkID(linkID)
		if err == nil && len(clicks) >= link.AccessLimit {
			return errors.New("link access limit reached")
		}
	}

	// Créer l'enregistrement de clic
	click := &domain.Click{
		LinkID:      linkID,
		ClickedAt:   time.Now(),
		IPAddress:   ipAddress,
		UserAgent:   userAgent,
		Referer:     referer,
		CountryCode: s.detectCountryCode(ipAddress), // TODO: Implémenter géolocalisation
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.clickRepo.CreateClick(click)
}

// GetClickStats récupère les statistiques de clics avec filtrage
func (s *ClickService) GetClickStats(linkID uint, userID uint, filter *domain.ClickFilter) ([]*domain.Click, error) {
	// Vérifier que l'utilisateur possède le lien
	link, err := s.linkRepo.GetLinkByID(linkID)
	if err != nil {
		return nil, errors.New("link not found")
	}

	if link.UserID != userID {
		return nil, errors.New("unauthorized to view this link's stats")
	}

	// Appliquer des filtres par défaut si aucun n'est fourni
	if filter == nil {
		filter = &domain.ClickFilter{
			Page:  1,
			Limit: 50,
		}
	}

	// Valider les paramètres de pagination
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 || filter.Limit > 100 {
		filter.Limit = 50
	}

	return s.clickRepo.GetClickStats(linkID, filter)
}

// GetClickSummary récupère un résumé des statistiques
func (s *ClickService) GetClickSummary(linkID uint, userID uint) (*ClickSummary, error) {
	// Vérifier que l'utilisateur possède le lien
	link, err := s.linkRepo.GetLinkByID(linkID)
	if err != nil {
		return nil, errors.New("link not found")
	}

	if link.UserID != userID {
		return nil, errors.New("unauthorized to view this link's stats")
	}

	// Récupérer tous les clics
	clicks, err := s.clickRepo.GetClicksByLinkID(linkID)
	if err != nil {
		return nil, err
	}

	// Calculer les statistiques
	summary := &ClickSummary{
		TotalClicks:         len(clicks),
		UniqueIPs:           s.countUniqueIPs(clicks),
		TopCountries:        s.getTopCountries(clicks),
		TopReferers:         s.getTopReferers(clicks),
		LastClickAt:         s.getLastClickTime(clicks),
		AverageClicksPerDay: s.calculateAverageClicksPerDay(clicks),
	}

	return summary, nil
}

// DeleteClick supprime un clic spécifique
func (s *ClickService) DeleteClick(clickID uint, userID uint) error {
	// TODO: Implémenter GetClickByID dans le repository
	// TODO: Vérifier que l'utilisateur possède le lien associé

	return s.clickRepo.DeleteClick(clickID)
}

// ClickSummary représente un résumé des statistiques de clics
type ClickSummary struct {
	TotalClicks         int            `json:"total_clicks"`
	UniqueIPs           int            `json:"unique_ips"`
	TopCountries        map[string]int `json:"top_countries"`
	TopReferers         map[string]int `json:"top_referers"`
	LastClickAt         *time.Time     `json:"last_click_at"`
	AverageClicksPerDay float64        `json:"average_clicks_per_day"`
}

// Méthodes utilitaires privées
func (s *ClickService) detectCountryCode(ipAddress string) string {
	// TODO: Implémenter la géolocalisation IP
	// Pour l'instant, retourner une chaîne vide
	return ""
}

func (s *ClickService) countUniqueIPs(clicks []*domain.Click) int {
	ipMap := make(map[string]bool)
	for _, click := range clicks {
		if click.IPAddress != "" {
			ipMap[click.IPAddress] = true
		}
	}
	return len(ipMap)
}

func (s *ClickService) getTopCountries(clicks []*domain.Click) map[string]int {
	countryMap := make(map[string]int)
	for _, click := range clicks {
		if click.CountryCode != "" {
			countryMap[click.CountryCode]++
		}
	}
	return countryMap
}

func (s *ClickService) getTopReferers(clicks []*domain.Click) map[string]int {
	refererMap := make(map[string]int)
	for _, click := range clicks {
		if click.Referer != "" {
			refererMap[click.Referer]++
		}
	}
	return refererMap
}

func (s *ClickService) getLastClickTime(clicks []*domain.Click) *time.Time {
	if len(clicks) == 0 {
		return nil
	}

	var lastClick time.Time
	for _, click := range clicks {
		if click.ClickedAt.After(lastClick) {
			lastClick = click.ClickedAt
		}
	}
	return &lastClick
}

func (s *ClickService) calculateAverageClicksPerDay(clicks []*domain.Click) float64 {
	if len(clicks) == 0 {
		return 0.0
	}

	// Trouver la date du premier et dernier clic
	var firstClick, lastClick time.Time
	for _, click := range clicks {
		if firstClick.IsZero() || click.ClickedAt.Before(firstClick) {
			firstClick = click.ClickedAt
		}
		if click.ClickedAt.After(lastClick) {
			lastClick = click.ClickedAt
		}
	}

	// Calculer le nombre de jours
	days := lastClick.Sub(firstClick).Hours() / 24
	if days < 1 {
		days = 1
	}

	return float64(len(clicks)) / days
}
