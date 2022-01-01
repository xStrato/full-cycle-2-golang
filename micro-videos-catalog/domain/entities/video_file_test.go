package entities_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
)

func TestVideoFile(t *testing.T) {
	t.Run("Constructor_ValidParams_ShouldCreateVideoFile", func(t *testing.T) {
		// Arrange
		id, title, duration := uuid.NewV4().String(), "Inception.mp4", 1.30
		// Act
		v1 := entities.NewVideoFileWithId(id, title, duration)
		v2 := entities.NewVideoFile(title, duration)
		//Assert
		require.NotNil(t, v1)
		require.NotNil(t, v2)
		require.Equal(t, v1.GetId(), id)
		require.Equal(t, v1.GetTitle(), title)
		require.Equal(t, v2.GetTitle(), title)
		require.NotNil(t, v2.Entity)
	})

	t.Run("SetTitle_InvalidNilAndEmptyParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, title, duration := uuid.NewV4().String(), "Inception.mp4", 1.30
		v := entities.NewVideoFileWithId(id, title, duration)
		errorMsg := "'title' cannot be empty"
		// Act
		err := v.SetTitle("")
		//Assert
		require.NotNil(t, v)
		require.EqualError(t, err, errorMsg)
	})

	t.Run("SetId_InvalidUUIDParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, title, duration := uuid.NewV4().String(), "Inception.mp4", 1.30
		v := entities.NewVideoFileWithId(id, title, duration)
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
}
