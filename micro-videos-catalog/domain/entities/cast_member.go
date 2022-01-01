package entities

import (
	"errors"
	"fmt"

	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/common"
	"github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/constants"
)

type CastMember struct {
	*common.Entity
	name           string
	castMemberType constants.CastMemberType
}

func NewCastMember(name string, cmt constants.CastMemberType) *CastMember {
	return &CastMember{
		Entity:         common.NewEntity(),
		name:           name,
		castMemberType: cmt,
	}
}

func NewCastMemberWithId(id, name string, cmt constants.CastMemberType) *CastMember {
	return &CastMember{
		Entity:         common.NewEntityWithId(id),
		name:           name,
		castMemberType: cmt,
	}
}

func (cm *CastMember) GetName() string {
	return cm.name
}

func (cm *CastMember) SetName(n string) error {
	if len(n) <= 0 {
		return errors.New("'name' cannot be empty")
	}
	cm.name = n
	return nil
}

func (cm *CastMember) GetCastMemberType() constants.CastMemberType {
	return cm.castMemberType
}

func (cm *CastMember) SetCastMemberType(cmt constants.CastMemberType) error {
	if len(cmt) <= 0 {
		return errors.New("'CastMember' cannot be empty")
	}
	cm.castMemberType = cmt
	return nil
}

func (cm *CastMember) SetCastMemberTypeFromString(cmt string) error {
	if len(cmt) <= 0 {
		return errors.New("'CastMember' cannot be empty")
	}
	if cmt == string(constants.TYPE1) || cmt == string(constants.TYPE2) {
		cm.castMemberType = constants.CastMemberType(cmt)
		return nil
	}
	return fmt.Errorf("'%v' cannot be parsed as valid 'CastMemberType'", cmt)
}
