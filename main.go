package main

import (
	"context"
	"fmt"
	"log"

	"entdemo/ent"
	"entdemo/ent/user"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
    client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()

	ctx := context.Background()
    // Run the auto migration tool.
    if err := client.Schema.Create(ctx); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

	params:= &ent.User{
		Name: "Jaqueline",
		Age: 34,
	}
	_, err = CreateUser(ctx, client, params)

	
	QueryUser(ctx, client)
}


func CreateUser(ctx context.Context, client *ent.Client, params *ent.User) (*ent.User, error) {
    u, err := client.User.
        Create().
        SetAge(params.Age).
        SetName(params.Name).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating user: %w", err)
    }
    log.Println("user was created: ", u)
    return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
    u, err := client.User.
        Query().
        Where(user.Name("Jaqueline")).
        // `Only` fails if no user found,
        // or more than 1 user returned., params *ent.User
        Only(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed querying user: %w", err)
    }
    log.Println("user returned: ", u)
    return u, nil
}