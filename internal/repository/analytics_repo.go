package repository

import (
	"time"

	"gorm.io/gorm"

	"github.com/PetJs/blog-backend/internal/models"
)

type AnalyticsRepository struct {
	DB *gorm.DB
}

func NewAnalyticsRepository(db *gorm.DB) *AnalyticsRepository {
	return &AnalyticsRepository{DB: db}
}

type PostCounts struct {
	Total     int64 `json:"total"`
	Published int64 `json:"published"`
	Draft     int64 `json:"draft"`
}

type BlockTypeStat struct {
	Type  string `json:"type"`
	Count int64  `json:"count"`
}

type ViewCounts struct {
	Total      int64 `json:"total"`
	Last7Days  int64 `json:"last_7_days"`
	Last30Days int64 `json:"last_30_days"`
}

type TopPost struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	ViewCount int64  `json:"view_count"`
}

func (r *AnalyticsRepository) GetPostCounts() (PostCounts, error) {
	var counts PostCounts
	r.DB.Model(&models.Post{}).Count(&counts.Total)
	r.DB.Model(&models.Post{}).Where("status = ?", "published").Count(&counts.Published)
	r.DB.Model(&models.Post{}).Where("status = ?", "draft").Count(&counts.Draft)
	return counts, nil
}

func (r *AnalyticsRepository) GetBlockTypeStats() ([]BlockTypeStat, error) {
	var stats []BlockTypeStat
	err := r.DB.Model(&models.Block{}).
		Select("type, count(*) as count").
		Group("type").
		Order("count DESC").
		Scan(&stats).Error
	return stats, err
}

func (r *AnalyticsRepository) GetViewCounts() (ViewCounts, error) {
	var counts ViewCounts
	now := time.Now()
	r.DB.Model(&models.PostView{}).Count(&counts.Total)
	r.DB.Model(&models.PostView{}).Where("viewed_at >= ?", now.AddDate(0, 0, -7)).Count(&counts.Last7Days)
	r.DB.Model(&models.PostView{}).Where("viewed_at >= ?", now.AddDate(0, 0, -30)).Count(&counts.Last30Days)
	return counts, nil
}

func (r *AnalyticsRepository) GetTopPosts(limit int) ([]TopPost, error) {
	var posts []TopPost
	err := r.DB.Model(&models.PostView{}).
		Select("posts.id, posts.title, posts.slug, count(post_views.id) as view_count").
		Joins("JOIN posts ON posts.id = post_views.post_id").
		Where("posts.status = ?", "published").
		Group("posts.id, posts.title, posts.slug").
		Order("view_count DESC").
		Limit(limit).
		Scan(&posts).Error
	return posts, err
}

func (r *AnalyticsRepository) RecordView(postID uint) error {
	return r.DB.Create(&models.PostView{PostID: postID}).Error
}
