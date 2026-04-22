package service

import (
	"fmt"
	"strings"

	"github.com/kamilrahmatullin/mama-ryadom-mvp/internal/db"
	"github.com/kamilrahmatullin/mama-ryadom-mvp/internal/repository"
)

type KnowledgeBaseService struct {
	repo *repository.KnowledgeBaseRepository
}

func NewKnowledgeBaseService(r *repository.KnowledgeBaseRepository) *KnowledgeBaseService {
	return &KnowledgeBaseService{repo: r}
}

func (s *KnowledgeBaseService) CreateArticle(title, content, recommendations string, severity db.SeverityLevel) error {
	title = strings.TrimSpace(title)
	content = strings.TrimSpace(content)
	recommendations = strings.TrimSpace(recommendations)

	if title == "" {
		return fmt.Errorf("title is required")
	}

	if !severity.IsValid() {
		return fmt.Errorf("invalid severity")
	}

	article := &db.KnowledgeBase{
		Title:           title,
		Content:         content,
		Recommendations: recommendations,
		Severity:        severity,
	}

	return s.repo.Create(article)
}

func (s *KnowledgeBaseService) GetAll() ([]db.KnowledgeBase, error) {
	return s.repo.GetAll()
}

func (s *KnowledgeBaseService) GetByID(id uint) (*db.KnowledgeBase, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid id")
	}

	return s.repo.GetByID(id)
}

func (s *KnowledgeBaseService) Delete(id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid id")
	}

	return s.repo.DeleteByID(id)
}

func (s *KnowledgeBaseService) Update(article *db.KnowledgeBase) error {
	if article.ID == 0 {
		return fmt.Errorf("invalid id")
	}

	if strings.TrimSpace(article.Title) == "" {
		return fmt.Errorf("title is required")
	}

	return s.repo.Update(article)
}
