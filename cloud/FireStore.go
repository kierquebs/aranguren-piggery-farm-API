package cloud

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

var Client *firestore.Client
var err error

func FirestoreInit() {
	log.Println("Connecting to Firestore...")

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv("Firestore_Private_Key_Path"))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	Client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully connected to Firestore...")

	defer Client.Close()

}
