package entities_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
)

func TestCategory(t *testing.T) {
	t.Run("Constructor_ValidParams_ShouldCreateCategory", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Movie"
		// Act
		c1 := entities.NewCategoryWithId(id, name)
		c2 := entities.NewCategory(name)
		//Assert
		require.NotNil(t, c1)
		require.NotNil(t, c2)
		require.Equal(t, c1.GetName(), name)
		require.Equal(t, c2.GetName(), name)
		require.Equal(t, c1.GetId(), id)
		require.NotNil(t, c2.Entity)
	})

	t.Run("SetName_InvalidNameParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Movie"
		c := entities.NewCategoryWithId(id, name)
		errorMsg := "'name' cannot be empty"
		// Act
		err := c.SetName("")
		//Assert
		require.NotNil(t, c)
		require.EqualError(t, err, errorMsg)
	})

	t.Run("SetId_InvalidUUIDParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Movie"
		c := entities.NewCategoryWithId(id, name)
		errorMsg1 := "'id' cannot be empty"
		errorMsg2 := "'id' is not a valid UUID"
		// Act
		err1 := c.SetId("")
		err2 := c.SetId("00000000-0000-0000-0000-000000000000")
		//Assert
		require.NotNil(t, c)
		require.EqualError(t, err1, errorMsg1)
		require.EqualError(t, err2, errorMsg2)
	})
}
