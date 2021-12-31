package entities_test

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/constants"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/entities"
)

func TestCastMember(t *testing.T) {
	t.Run("Constructor_ValidParams_ShouldCreateCastMember", func(t *testing.T) {
		// Arrange
		id, name, cmt := uuid.NewV4().String(), "DiCaprio", constants.TYPE1
		// Act
		c1 := entities.NewCastMemberWithIdAndMemberType(id, name, cmt)
		c2 := entities.NewCastMemberWithMemberType(name, cmt)
		//Assert
		require.NotNil(t, c1)
		require.NotNil(t, c2)
		require.Equal(t, c1.GetName(), name)
		require.Equal(t, c2.GetName(), name)
		require.Equal(t, c1.GetCastMemberType(), cmt)
		require.Equal(t, c2.GetCastMemberType(), cmt)
		require.Equal(t, c1.GetId(), id)
		require.NotNil(t, c2.Entity)
	})

	t.Run("SetName_InvalidNameParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, name, cmt := uuid.NewV4().String(), "DiCaprio", constants.TYPE1
		c := entities.NewCastMemberWithIdAndMemberType(id, name, cmt)
		errorMsg := "'name' cannot be empty"
		// Act
		err := c.SetName("")
		//Assert
		require.NotNil(t, c)
		require.EqualError(t, err, errorMsg)
	})

	t.Run("SetId_InvalidUUIDParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, name, cmt := uuid.NewV4().String(), "DiCaprio", constants.TYPE1
		c := entities.NewCastMemberWithIdAndMemberType(id, name, cmt)
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

	t.Run("SetCastMemberType_InvalidCastMemberParams_ShouldReturnErrorNotNil", func(t *testing.T) {
		// Arrange
		id, name, cmt := uuid.NewV4().String(), "DiCaprio", constants.TYPE1
		c := entities.NewCastMemberWithIdAndMemberType(id, name, cmt)

		invalidCastMemberType := "TYPE3"
		errorMsg1 := "'CastMember' cannot be empty"
		errorMsg2 := fmt.Sprintf("'%v' cannot be parsed as valid 'CastMemberType'", invalidCastMemberType)
		// Act
		err1 := c.SetCastMemberTypeFromString("")
		err2 := c.SetCastMemberTypeFromString(invalidCastMemberType)
		//Assert
		require.NotNil(t, c)
		require.EqualError(t, err1, errorMsg1)
		require.EqualError(t, err2, errorMsg2)
	})
}
