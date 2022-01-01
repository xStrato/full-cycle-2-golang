package entities_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/constants"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
)

func TestVideo(t *testing.T) {
	t.Run("Constructor_ValidParams_ShouldCreateVideo", func(t *testing.T) {
		// Arrange
		id, title, desc, year, open := uuid.NewV4().String(), "Inception", "A intriguing movie indeed", 1990, true
		// Act
		v1 := entities.NewVideoWithId(id, title, desc, year, open)
		v2 := entities.NewVideo(title, desc, year, open)
		//Assert
		require.NotNil(t, v1)
		require.NotNil(t, v2)
		require.Equal(t, v1.GetId(), id)
		require.Equal(t, v1.GetTitle(), title)
		require.Equal(t, v1.GetDescription(), desc)
		require.Equal(t, v1.GetYearLaunched(), year)
		require.Equal(t, v1.GetOpened(), open)
		require.Equal(t, v2.GetTitle(), title)
		require.NotNil(t, v2.Entity)
	})

	t.Run("SetTitle_InvalidTitleParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, title, desc, year, open := uuid.NewV4().String(), "Inception", "A intriguing movie indeed", 1990, true
		v := entities.NewVideoWithId(id, title, desc, year, open)
		errorMsg := "'title' cannot be empty"
		// Act
		err := v.SetTitle("")
		//Assert
		require.NotNil(t, v)
		require.EqualError(t, err, errorMsg)
	})

	t.Run("SetId_InvalidUUIDParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, title, desc, year, open := uuid.NewV4().String(), "Inception", "A intriguing movie indeed", 1990, true
		v := entities.NewVideoWithId(id, title, desc, year, open)
		errorMsg1 := "'id' cannot be empty"
		errorMsg2 := "'id' is not a valid UUID"
		// Act
		err1 := v.SetId("")
		err2 := v.SetId("00000000-0000-0000-0000-000000000000")
		//Assert
		require.NotNil(t, v)
		require.EqualError(t, err1, errorMsg1)
		require.EqualError(t, err2, errorMsg2)
	})

	t.Run("SetYearLaunched_InvalidYearLaunchedParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, title, desc, year, open := uuid.NewV4().String(), "Inception", "A intriguing movie indeed", 1990, true
		v := entities.NewVideoWithId(id, title, desc, year, open)
		errorMsg1 := "'yearLaunched' must be greater than 1700"
		errorMsg2 := "'yearLaunched' is greater than current year"
		// Act
		err1 := v.SetYearLaunched(1600)
		err2 := v.SetYearLaunched(time.Now().AddDate(1, 0, 0).Year())
		//Assert
		require.NotNil(t, v)
		require.EqualError(t, err1, errorMsg1)
		require.EqualError(t, err2, errorMsg2)
	})

	t.Run("AddT_ValidAndUniqueEntitiesParams_ShouldAddEntitiesToTheirSlices", func(t *testing.T) {
		// Arrange
		id, title, desc, year, open := uuid.NewV4().String(), "Inception", "A intriguing movie indeed", 1990, true
		v := entities.NewVideoWithId(id, title, desc, year, open)

		c := entities.NewCategory("Film")
		g := entities.NewGenre("Action")
		cm := entities.NewCastMember("", constants.TYPE1)
		// Act
		err1 := v.AddT(c)
		err2 := v.AddT(g)
		err3 := v.AddT(cm)
		//Assert
		require.NotNil(t, v)
		require.NoError(t, err1)
		require.NoError(t, err2)
		require.NoError(t, err3)
		require.Len(t, v.GetCategories(), 1)
		require.Len(t, v.GetGenres(), 1)
		require.Len(t, v.GetCastMembers(), 1)
	})

	t.Run("AddT_InvalidNilEntitiesParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, title, desc, year, open := uuid.NewV4().String(), "Inception", "A intriguing movie indeed", 1990, true
		v := entities.NewVideoWithId(id, title, desc, year, open)

		var c *entities.Category
		var g *entities.Genre
		var cm *entities.CastMember

		errorMsg := "cannot Add 'nil' value"
		// Act
		err1 := v.AddT(c)
		err2 := v.AddT(g)
		err3 := v.AddT(cm)
		//Assert
		require.NotNil(t, v)
		require.EqualError(t, err1, errorMsg)
		require.EqualError(t, err2, errorMsg)
		require.EqualError(t, err3, errorMsg)
	})

	t.Run("RemoveT_ValidEntitiesParams_ShouldRemoveEntitiesFromTheirSlices", func(t *testing.T) {
		// Arrange
		id, title, desc, year, open := uuid.NewV4().String(), "Inception", "A intriguing movie indeed", 1990, true
		v := entities.NewVideoWithId(id, title, desc, year, open)

		c := entities.NewCategory("Film")
		g := entities.NewGenre("Action")
		cm := entities.NewCastMember("", constants.TYPE1)

		v.AddT(c)
		v.AddT(entities.NewCategory("Film"))
		v.AddT(g)
		v.AddT(entities.NewGenre("Action"))
		v.AddT(cm)
		v.AddT(entities.NewCastMember("", constants.TYPE1))
		// Act
		err1 := v.RemoveT(c)
		err2 := v.RemoveT(g)
		err3 := v.RemoveT(cm)
		//Assert
		require.NotNil(t, v)
		require.NoError(t, err1)
		require.NoError(t, err2)
		require.NoError(t, err3)
		require.Len(t, v.GetCategories(), 1)
		require.Len(t, v.GetGenres(), 1)
		require.Len(t, v.GetCastMembers(), 1)
	})

	t.Run("RemoveT_InvalidNilEntitiesParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, title, desc, year, open := uuid.NewV4().String(), "Inception", "A intriguing movie indeed", 1990, true
		v := entities.NewVideoWithId(id, title, desc, year, open)

		var c *entities.Category
		var g *entities.Genre
		var cm *entities.CastMember

		errorMsg := "cannot Remove 'nil' value"
		// Act
		err1 := v.RemoveT(c)
		err2 := v.RemoveT(g)
		err3 := v.RemoveT(cm)
		v.RemoveT(cm)
		//Assert
		require.NotNil(t, v)
		require.Errorf(t, err1, errorMsg)
		require.Errorf(t, err2, errorMsg)
		require.Errorf(t, err3, errorMsg)
	})

	t.Run("SetT_ValidEntitiesParams_ShouldSetEntitiesToCorrespondingSlices", func(t *testing.T) {
		// Arrange
		id, title, desc, year, open := uuid.NewV4().String(), "Inception", "A intriguing movie indeed", 1990, true
		v := entities.NewVideoWithId(id, title, desc, year, open)

		categories := []entities.Category{
			*entities.NewCategory(""),
		}
		genres := []entities.Genre{
			*entities.NewGenre(""),
			*entities.NewGenre(""),
		}
		castMembers := []entities.CastMember{
			*entities.NewCastMember("", constants.TYPE1),
			*entities.NewCastMember("", constants.TYPE1),
			*entities.NewCastMember("", constants.TYPE1),
		}
		// Act
		v.SetCategories(&categories)
		v.SetGenres(&genres)
		v.SetCastMembers(&castMembers)
		//Assert
		require.NotNil(t, v)
		require.Len(t, v.GetCategories(), 1)
		require.Len(t, v.GetGenres(), 2)
		require.Len(t, v.GetCastMembers(), 3)
	})

	t.Run("SetT_InvalidNilParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, title, desc, year, open := uuid.NewV4().String(), "Inception", "A intriguing movie indeed", 1990, true
		v := entities.NewVideoWithId(id, title, desc, year, open)

		categories := []entities.Category{}
		genres := []entities.Genre{}
		castMembers := []entities.CastMember{}

		errorMsg1 := "'categories' cannot be empty or nil"
		errorMsg2 := "'genres' cannot be empty or nil"
		errorMsg3 := "'castMembers' cannot be empty or nil"
		// Act
		err1 := v.SetCategories(&categories)
		err2 := v.SetGenres(&genres)
		err3 := v.SetCastMembers(&castMembers)

		//Assert
		require.NotNil(t, v)
		require.EqualError(t, err1, errorMsg1)
		require.EqualError(t, err2, errorMsg2)
		require.EqualError(t, err3, errorMsg3)
		require.EqualError(t, v.SetCategories(nil), errorMsg1)
		require.EqualError(t, v.SetGenres(nil), errorMsg2)
		require.EqualError(t, v.SetCastMembers(nil), errorMsg3)
	})
}
