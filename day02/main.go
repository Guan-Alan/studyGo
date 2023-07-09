package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// 	ctx, cancel = context.WithTimeout(ctx, callOpts.RequestTimeout)
//		defer cancel()

func main() {

	str := "a,b"
	split := strings.Split(str, `,`)
	fmt.Println(split)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	deadline, _ := ctx.Deadline()
	fmt.Println(deadline)

	fmt.Println(time.Until(deadline))
}

func GetCtxLiveTime(ctx context.Context) time.Duration {
	deadline, ok := ctx.Deadline()
	if !ok {
		return 0
	}

	return time.Until(deadline)
}
