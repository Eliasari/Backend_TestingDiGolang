package integration

import (
    "os"
    "path/filepath"
    "testing"
    "context"
    "time"

    "go-fiber/app/model"
    "go-fiber/utils"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/bson"

    "log"
 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func SetupEnvTest(t *testing.T) {
    t.Helper()

    // ambil working directory saat ini
    wd, err := os.Getwd()
    if err != nil {
        t.Fatalf("cannot get working dir: %v", err)
    }

    // mulai dari wd, naik sampai ketemu .env.test
    var envPath string
    for dir := wd; dir != filepath.Dir(dir); dir = filepath.Dir(dir) {
        tryPath := filepath.Join(dir, ".env.test")
        if _, err := os.Stat(tryPath); err == nil {
            envPath = tryPath
            break
        }
    }

    if envPath == "" {
        t.Fatal(".env.test tidak ditemukan di root project")
    }

    t.Logf("Loading env from: %s", envPath)
    err = godotenv.Load(envPath)
    if err != nil {
        t.Fatalf("failed to load .env.test: %v", err)
    }
}

func SetupMongoTestCollection(t *testing.T, collectionName string) (*mongo.Client, *mongo.Collection) {
    t.Helper()

    uri := os.Getenv("MONGO_URI")
    dbName := os.Getenv("MONGO_DB_NAME")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        t.Fatalf("failed to connect mongo: %v", err)
    }

    coll := client.Database(dbName).Collection(collectionName)

    return client, coll
}


// SeedUsers insert dummy user untuk login & get users
func SeedUsers(ctx context.Context, coll *mongo.Collection) {
    pass, err := utils.HashPassword("123")
    if err != nil {
        panic(err)
    }

    users := []model.User{
        {Username: "user1", Email: "user1@mail.com", Password: pass, Role: "admin", CreatedAt: time.Now()},
        {Username: "user2", Email: "user2@mail.com", Password: pass, Role: "user", CreatedAt: time.Now()},
    }

    for _, u := range users {
        filter := bson.M{"username": u.Username}
        count, _ := coll.CountDocuments(ctx, filter)
        if count == 0 {
            _, err := coll.InsertOne(ctx, u)
            if err != nil {
                log.Printf("failed insert user %s: %v", u.Username, err)
            }
        }
    }
}


// CleanupMongo disconnect setelah test
func CleanupMongo(t *testing.T, client *mongo.Client) {
    t.Helper()
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := client.Disconnect(ctx); err != nil {
        t.Fatalf("failed to disconnect mongo: %v", err)
    }
}