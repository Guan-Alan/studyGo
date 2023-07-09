package produce_api

import (
	"context"
	"fmt"

	"testing"
)

func TestCallService_PlatformWashReceive(t *testing.T) {
	err := Service.PlatformWashReceive(context.Background(), 264)
	fmt.Println(err)
}
