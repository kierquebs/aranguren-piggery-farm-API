package cloud

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func FirestoreInit() {
	fmt.Println("Connecting to Firestore...")

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv("Firestore_Private_Key_Path"))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(client)
	defer client.Close()

}
