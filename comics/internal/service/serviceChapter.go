package service

import (
	"context"
	"errors"

	"github.com/likoscp/finalAddProgramming/comics/internal/repository"
	"github.com/likoscp/finalAddProgramming/comics/models"
)

type ChaptersService struct {
	repo *repository.ChapterRepository
}

func NewChaptersService(repo *repository.ChapterRepository) *ChaptersService {
	return &ChaptersService{
		repo: repo,
	}
}

func (s *ChaptersService) CreateChapter(ctx context.Context, req models.Chapter) (uint, error) {
	chapterID, err := s.repo.CreateChapter(ctx, req)
	if err != nil {
		return 0, err
	}
	return chapterID, nil
}

func (s *ChaptersService) UpdateChapter(ctx context.Context, id uint, req models.Chapter) error {
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		return errors.New("chapter not found")
	}
	return s.repo.UpdateChapter(ctx, id, req)
}

func (s *ChaptersService) GetByID(ctx context.Context, id uint) (*models.Chapter, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ChaptersService) GetAllChapters(ctx context.Context) ([]models.Chapter, error) {
	return s.repo.GetAllChapters(ctx)
}

func (s *ChaptersService) DeleteChapter(ctx context.Context, id uint) error {
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		return errors.New("chapter not found")
	}
	return s.repo.DeleteChapter(ctx, id)
}

func (s *ChaptersService) GetChaptersByUserID(ctx context.Context, comicID uint) ([]models.Chapter, error) {
	return s.repo.GetChaptersByUserID(ctx, comicID)
}
