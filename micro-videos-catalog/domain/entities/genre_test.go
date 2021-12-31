package entities_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
)

func TestGenre(t *testing.T) {
	t.Run("Constructor_ValidParams_ShouldCreateGenre", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Horror"
		// Act
		g1 := entities.NewGenreWithId(id, name)
		g2 := entities.NewGenre(name)
		//Assert
		require.NotNil(t, g1)
		require.NotNil(t, g2)
		require.Equal(t, g1.GetName(), name)
		require.Equal(t, g2.GetName(), name)
		require.Equal(t, g1.GetId(), id)
		require.NotNil(t, g2.Entity)
	})

	t.Run("SetName_InvalidNameParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Horror"
		g := entities.NewGenreWithId(id, name)
		errorMsg := "'name' cannot be empty"
		// Act
		err := g.SetName("")
		//Assert
		require.NotNil(t, g)
		require.EqualError(t, err, errorMsg)
	})

	t.Run("SetId_InvalidUUIDParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Horror"
		g := entities.NewGenreWithId(id, name)
		errorMsg1 := "'id' cannot be empty"
		errorMsg2 := "'id' is not a valid UUID"
		// Act
		err1 := g.SetId("")
		err2 := g.SetId("00000000-0000-0000-0000-000000000000")
		//Assert
		require.NotNil(t, g)
		require.EqualError(t, err1, errorMsg1)
		require.EqualError(t, err2, errorMsg2)
	})

	t.Run("AddCategory_ValidCategory_ShouldAddValuesToSlice", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Horror"
		g := entities.NewGenreWithId(id, name)
		// Act
		err1 := g.AddCategory(entities.NewCategory("Movie"))
		err2 := g.AddCategory(entities.NewCategory("Serie"))
		//Assert
		require.NotNil(t, g)
		require.Len(t, g.GetCategories(), 2)
		require.NoError(t, err1)
		require.NoError(t, err2)
	})

	t.Run("AddCategory_InvalidNilCategory_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Horror"
		g := entities.NewGenreWithId(id, name)
		errorMsg := "'category' cannot be nil"
		// Act
		err := g.AddCategory(nil)
		//Assert
		require.NotNil(t, g)
		require.EqualError(t, err, errorMsg)
	})

	t.Run("RemoveCategory_ValidCategory_ShouldRemoveValuesFromSlice", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Horror"
		g := entities.NewGenreWithId(id, name)
		c := entities.NewCategory("Movie")
		g.AddCategory(c)
		g.AddCategory(entities.NewCategory("Serie"))
		// Act
		err1 := g.RemoveCategory(c)
		//Assert
		require.NotNil(t, g)
		require.NoError(t, err1)
		require.Len(t, g.GetCategories(), 1)
	})

	t.Run("RemoveCategory_InvalidNilCategory_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Horror"
		g := entities.NewGenreWithId(id, name)
		errorMsg := "'category' cannot be nil"
		// Act
		err := g.RemoveCategory(nil)
		//Assert
		require.NotNil(t, g)
		require.EqualError(t, err, errorMsg)
	})

	t.Run("SetCategories_ValidCategories_ShouldSetNewEntryValues", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Horror"
		g := entities.NewGenreWithId(id, name)

		categories := []entities.Category{
			*entities.NewCategory("Movie"),
			*entities.NewCategory("Serie"),
		}
		// Act
		err1 := g.SetCategories(categories)
		//Assert
		require.NotNil(t, g)
		require.NoError(t, err1)
		require.Len(t, g.GetCategories(), 2)
		require.EqualValues(t, categories, g.GetCategories())
	})

	t.Run("SetCategories_InvalidEmptyAndNilCategories_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Horror"
		g := entities.NewGenreWithId(id, name)
		errorMsg := "'categories' cannot be empty or nil"
		// Act
		err1 := g.SetCategories(nil)
		err2 := g.SetCategories(make([]entities.Category, 0))
		//Assert
		require.NotNil(t, g)
		require.Error(t, err1, errorMsg)
		require.Error(t, err2, errorMsg)
	})
}
