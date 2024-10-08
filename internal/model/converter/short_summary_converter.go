package converter

import (
	"fmt"

	"github.com/itsLeonB/cv-api/internal/entity"
	"github.com/itsLeonB/cv-api/internal/model"
)

func ProfileToShortSummary(profile *entity.Profile) *model.ShortSummary {
	return &model.ShortSummary{
		Header: fmt.Sprintf("Hi, I'm %s 👋", profile.Nickname),
		Body:   fmt.Sprintf("I'm a %s based in %s. %s", profile.Occupation, profile.Location, profile.ShortSummary),
	}
}
