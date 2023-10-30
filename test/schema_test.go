package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSchema_NestedStruct(t *testing.T) {
	ctx := initializeTestWorld()

	embeddedStructEntity := embeddedStructTable.Add(ctx.World, PetCommunity{
		Owners: []Owner{
			{
				Name:  testName1,
				Age:   testAge1,
				Happy: true,
				Pos:   testPos1,
				Pets: []Pet{
					{
						Name: "odie",
						Kind: Dog,
					},
					{
						Name: "squishy",
						Kind: Cat,
					},
				},
			},
			{
				Name:  testName2,
				Age:   testAge2,
				Happy: true,
				Pos:   testPos2,
				Pets: []Pet{
					{
						Name: "sherlock",
						Kind: Dog,
					},
				},
			},
		},
	})

	embeddedStruct := embeddedStructTable.Get(ctx.World, embeddedStructEntity)
	assert.Equal(t, embeddedStructEntity, embeddedStruct.Id)

	owners := embeddedStruct.Owners
	require.Len(t, owners, 2)

	owner1 := owners[0]
	assert.Equal(t, testName1, owner1.Name)
	assert.Equal(t, testAge1, owner1.Age)
	assert.Len(t, owner1.Pets, 2)
	assert.Equal(t, testPos1, owner1.Pos)
	assert.True(t, owner1.Happy)

	require.Len(t, owner1.Pets, 2)
	assert.Equal(t, Dog, owner1.Pets[0].Kind)
	assert.Equal(t, "odie", owner1.Pets[0].Name)
	assert.Equal(t, Cat, owner1.Pets[1].Kind)
	assert.Equal(t, "squishy", owner1.Pets[1].Name)

	owner2 := owners[1]
	assert.Equal(t, testName2, owner2.Name)
	assert.Equal(t, testAge2, owner2.Age)
	require.Len(t, owner2.Pets, 1)
	assert.Equal(t, Dog, owner2.Pets[0].Kind)
	assert.Equal(t, "sherlock", owner2.Pets[0].Name)
}
