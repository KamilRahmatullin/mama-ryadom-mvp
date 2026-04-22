package repository

import (
	"github.com/kamilrahmatullin/mama-ryadom-mvp/internal/db"
	"gorm.io/gorm"
)

type KnowledgeBaseRepository struct {
	db *gorm.DB
}

func NewKnowledgeBaseRepository(db *gorm.DB) *KnowledgeBaseRepository {
	return &KnowledgeBaseRepository{db: db}
}

func (r *KnowledgeBaseRepository) Create(article *db.KnowledgeBase) error {
	return r.db.Create(article).Error
}

func (r *KnowledgeBaseRepository) GetAll() ([]db.KnowledgeBase, error) {
	var articles []db.KnowledgeBase
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *KnowledgeBaseRepository) GetByID(id uint) (*db.KnowledgeBase, error) {
	var article db.KnowledgeBase

	err := r.db.First(&article, id).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (r *KnowledgeBaseRepository) GetPaginated(limit, offset int) ([]db.KnowledgeBase, error) {
	var articles []db.KnowledgeBase

	err := r.db.
		Limit(limit).
		Offset(offset).
		Find(&articles).Error

	return articles, err
}

func (r *KnowledgeBaseRepository) GetBySeverity(severity db.SeverityLevel) ([]db.KnowledgeBase, error) {
	var articles []db.KnowledgeBase
	err := r.db.Where("severity = ?", severity).Find(&articles).Error
	return articles, err
}

func (r *KnowledgeBaseRepository) Update(article *db.KnowledgeBase) error {
	return r.db.Save(article).Error
}

func (r *KnowledgeBaseRepository) DeleteByID(id uint) error {
	return r.db.Delete(&db.KnowledgeBase{}, id).Error
}
