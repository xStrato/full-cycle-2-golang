package commands_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/application/categories/commands"
)

func TestCreateCategoryCommand(t *testing.T) {
	t.Run("Constructor_ValidCommandParams_ShouldCreateCommand", func(t *testing.T) {
		// Arrange
		name := "Movie"
		// Act
		c1 := commands.NewCreateCategoryCommandWithName(name)
		//Assert
		require.NotNil(t, c1)
		require.NotNil(t, c1.Command)
		require.Nil(t, c1.IsValid())
		require.Equal(t, c1.GetCommandType(), "CreateCategoryCommand")
	})

	t.Run("Constructor_InvalidCommandParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		name := "Mo"
		// Act
		c1 := commands.NewCreateCategoryCommandWithName(name)
		//Assert
		require.NotNil(t, c1)
		require.NotNil(t, c1.Command)
		require.NotNil(t, c1.IsValid())
		require.Equal(t, c1.GetCommandType(), "CreateCategoryCommand")
	})
}
