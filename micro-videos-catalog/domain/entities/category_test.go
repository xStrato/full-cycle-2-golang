package entities_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
)

func TestCategory(t *testing.T) {
	t.Run("Constructor_ValidParameters_ShouldCreateCategory", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Horror"
		// Act
		c1 := entities.NewFullCategory(id, name)
		c2 := entities.NewCategory(name)
		//Assert
		require.NotNil(t, c1)
		require.NotNil(t, c2)
		require.Equal(t, c1.GetName(), name)
		require.Equal(t, c2.GetName(), name)
		require.Equal(t, c1.GetId(), id)
	})

	t.Run("SetName_InvalidParameter_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, name := uuid.NewV4().String(), "Horror"
		c := entities.NewFullCategory(id, name)
		// Act
		errorMsg := "'name' cannot be empty"
		err := c.SetName("")
		//Assert
		require.NotNil(t, c)
		require.EqualError(t, err, errorMsg)
	})
}
