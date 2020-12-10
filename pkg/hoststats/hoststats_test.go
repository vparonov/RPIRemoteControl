package hoststats

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	hs, err := GetHostStats("localhost", 3)

	assert.Nil(t, err)
	assert.True(t, hs.IsAlive)

	fmt.Println(hs.PingStats)

}
