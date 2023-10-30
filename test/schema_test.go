package test

import (
	"github.com/curio-research/keystone/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchema_NestedStruct(t *testing.T) {
	ctx := initializeTestWorld()

	embeddedStructEntity := embeddedStructTable.Add(ctx.World, EmbeddedStructSchema{
		Emb: NestedStruct{
			Name: testName1,
			Age:  testAge1,
			Books: []Book{
				{
					Title: testBookTitle1,
				},
				{
					Title: testBookTitle2,
				},
			},
			Pos: state.Pos{X: 1, Y: 6},
		},
		People: []Person{
			{
				Name: testName2,
			},
			{
				Name: testName3,
			},
		},
	})

	embeddedStruct := embeddedStructTable.Get(ctx.World, embeddedStructEntity)
	assert.Equal(t, embeddedStructEntity, embeddedStruct.Id)

	assert.Len(t, embeddedStruct.People, 2)
	assert.Equal(t, embeddedStruct.People[0].Name, testName2)
	assert.Equal(t, embeddedStruct.People[1].Name, testName3)

	assert.Equal(t, testAge1, embeddedStruct.Emb.Age)
	assert.Equal(t, testName1, embeddedStruct.Emb.Name)

	assert.Len(t, embeddedStruct.Emb.Books, 2)
	assert.Equal(t, testBookTitle1, embeddedStruct.Emb.Books[0].Title)
	assert.Equal(t, testBookTitle2, embeddedStruct.Emb.Books[1].Title)
}
