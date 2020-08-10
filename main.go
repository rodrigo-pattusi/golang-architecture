package main

import (
	"context"
	"fmt"
	"github.com/rodrigo-pattusi/golang-architecture/session"
)

func main() {
	ctx := context.Background()

	ctx = session.SetUserId(ctx, 1)
	ctx = session.SetIsAdmin(ctx, true)

	uID := session.GetUserId(ctx)
	isAdmin := session.GetIsAdmin(ctx)
	fmt.Printf("User %d is an admin %t", uID, isAdmin)

	ctx = session.SetUserId(ctx, 1)
	ctx = session.SetIsAdmin(ctx, true)
	i := session.GetUserId(ctx)
	b := session.GetIsAdmin(ctx)
	fmt.Printf("user %d is admin %t\n", i, b)

}
