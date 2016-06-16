package asciiFolding

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestTextFolding(t *testing.T) {
    assert := assert.New(t)
    
    assert.Equal("Some", Fold("Söme"))
    assert.Equal("Gum", Fold("Güm"))
    assert.Equal("Cunku", Fold("Çünkü"))
    assert.Equal("alemi huzun icundur", Fold("âlêmî hüzün içûndur"))
    assert.Equal("Passe compose", Fold("Passé composé"))
}