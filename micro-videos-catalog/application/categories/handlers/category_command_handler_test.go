package handlers_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/commands"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/handlers"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
	mock_interfaces "github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/mocks"
)

func TestCategoryCommandHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_interfaces.NewMockCategoryRepository(ctrl)
	handler := handlers.NewCategoryCommandHandler(repository)

	t.Run("Handle_ValidCreateCategoryCommandParams_ShouldCreateAndPersistCommand", func(t *testing.T) {
		// Arrange
		categories := generateNewCategorySlice()
		var createdCategory *entities.Category

		repository.EXPECT().Commit().Return(nil).MaxTimes(1)
		repository.EXPECT().Add(gomock.Any()).MaxTimes(1).DoAndReturn(func(e *entities.Category) error {
			createdCategory = e
			categories = append(categories, *e)
			return nil
		})

		cmd := commands.NewCreateCategoryCommandWithName("Movie")
		expectedMessage := fmt.Sprintf("%v was successfully executed", cmd.GetCommandType())
		// Act
		result, err := handler.Handle(cmd)
		//Assert
		require.Nil(t, err)
		require.NotNil(t, createdCategory)
		require.True(t, result.HasSuccess())
		require.Equal(t, expectedMessage, result.GetMessage())
		require.IsType(t, &entities.Category{}, result.GetData())
		require.Len(t, categories, 6)
	})

	t.Run("Handle_InvalidCreateCategoryCommandParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		categories := generateNewCategorySlice()
		repository.EXPECT().Add(gomock.Any()).MaxTimes(0)
		repository.EXPECT().Commit().MaxTimes(0)

		cmd := commands.NewCreateCategoryCommandWithName("Mo")
		expectedMsg := fmt.Sprintf("%v state is invalid", cmd.GetCommandType())
		expectedDataMsg := "name: Mo does not validate as length(3|20)"
		// Act
		result, err := handler.Handle(cmd)
		//Assert
		require.Nil(t, err)
		require.False(t, result.HasSuccess())
		require.Equal(t, expectedMsg, result.GetMessage())
		require.Equal(t, expectedDataMsg, result.GetData())
		require.Len(t, categories, 5)
	})
}

func generateNewCategorySlice() []entities.Category {
	return []entities.Category{
		*entities.NewCategory("Movie"),
		*entities.NewCategory("Vlog"),
		*entities.NewCategory("Serie"),
		*entities.NewCategory("Documentary"),
		*entities.NewCategory("Animation"),
	}
}
