package converter

import (
	"fmt"

	"github.com/itsLeonB/cv-api/internal/entity"
	"github.com/itsLeonB/cv-api/internal/model"
)

func ProfileToShortSummary(profile *entity.Profile) *model.Summary {
	return &model.Summary{
		Header: fmt.Sprintf("Hi, I'm %s 👋", profile.Nickname),
		Body:   fmt.Sprintf("I'm a %s based in %s. %s", profile.Occupation, profile.Location, profile.ShortSummary),
		Type:   "short",
	}
}

func ProfileToSummary(profile *entity.Profile) *model.Summary {
	return &model.Summary{
		Header: profile.FullName,
		Body:   profile.Summary,
		Type:   "full",
	}
}
