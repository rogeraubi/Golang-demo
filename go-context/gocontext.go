package main

import (
	"context"
	"fmt"
)

func main() {
 type favContextKey string

 f := func(ctx context.Context, k favContextKey) {

	if v := ctx.Value(k) ; v !=nil {
      fmt.Println("found value:",v);
	}
 }
 k :=favContextKey("language")
 ctx:= context.WithValue(context.Background(), k,"GO")
 f(ctx,k)
 f(ctx, favContextKey("color"))

}