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
