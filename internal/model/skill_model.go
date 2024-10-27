package model

type NewSkillCategoryRequest struct {
	Name string `json:"name" binding:"required,min=3,max=255"`
}

type SkillCategoryResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

type NewSkillRequest struct {
	ProfileID   int    `json:"profile_id" binding:"required,numeric,gte=1"`
	CategoryID  int    `json:"category_id" binding:"required,numeric,gte=1"`
	Name        string `json:"name" binding:"required,min=1,max=255"`
	Description string `json:"description" binding:"required,min=3,max=255"`
}

type SkillResponse struct {
	ID          int                    `json:"id"`
	ProfileID   int                    `json:"profile_id"`
	CategoryID  int                    `json:"category_id,omitempty"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	DeletedAt   string                 `json:"deleted_at,omitempty"`
	Category    *SkillCategoryResponse `json:"category,omitempty"`
}

type UpdateSkillRequest struct {
	ID          int
	CategoryID  int    `json:"category_id" binding:"required,numeric,gte=1"`
	Name        string `json:"name" binding:"required,min=1,max=255"`
	Description string `json:"description" binding:"required,min=3,max=255"`
}
