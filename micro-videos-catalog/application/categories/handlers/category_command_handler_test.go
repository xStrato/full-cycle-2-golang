package handlers_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/commands"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
	mock_interfaces "github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/mocks"
)

func TestCategoryCommandHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock_interfaces.NewMockCategoryRepository(ctrl)

	categories := []entities.Category{
		*entities.NewCategory("Movie"),
		*entities.NewCategory("Vlog"),
		*entities.NewCategory("Serie"),
		*entities.NewCategory("Documentary"),
		*entities.NewCategory("Animation"),
	}
	t.Run("Handle_ValidCreateCategoryCommandParams_ShouldCreateAndPersistCommand", func(t *testing.T) {
		// Arrange
		name := "Movie"
		// Act
		c1 := commands.NewCreateCategoryCommand(name)
		repository.EXPECT().Add(gomock.Any()).Return(100).AnyTimes()
		//Assert
		require.NotNil(t, c1)
		require.Len(t, categories, 6)
		require.NotNil(t, c1.Command)
		require.Nil(t, c1.IsValid())
		require.Equal(t, c1.GetCommandType(), "CreateCategoryCommand")
	})

	t.Run("Handle_InvalidCreateCategoryCommandParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		name := "Mo"
		// Act
		c1 := commands.NewCreateCategoryCommand(name)
		//Assert
		require.NotNil(t, c1)
		require.Len(t, categories, 5)
		require.NotNil(t, c1.Command)
		require.NotNil(t, c1.IsValid())
		require.Equal(t, c1.GetCommandType(), "CreateCategoryCommand")
	})
}
