package integration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-fiber/app/model"
	"go-fiber/app/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"testing"
	"time"
)

// SeedAlumni insert dummy alumni jika belum ada
func SeedAlumni(ctx context.Context, db *mongo.Database) {
	userColl := db.Collection("users_test")
	alumniColl := db.Collection("alumni_test")

	// ambil user1 dan user2 dari users_test
	var user1, user2 model.User
	err := userColl.FindOne(ctx, bson.M{"username": "user1"}).Decode(&user1)
	if err != nil {
		log.Fatalf("failed to get user1: %v", err)
	}

	err = userColl.FindOne(ctx, bson.M{"username": "user2"}).Decode(&user2)
	if err != nil {
		log.Fatalf("failed to get user2: %v", err)
	}

	alumni := []model.Alumni{
		{
			NIM: "2023001", Nama: "Alice", Jurusan: "TI", Angkatan: 2023, TahunLulus: 2027,
			Email: "alice@mail.com", NoTelepon: "081234567890", Alamat: "Jakarta",
			UserID: user1.ID, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		},
		{
			NIM: "2023002", Nama: "Bob", Jurusan: "SI", Angkatan: 2023, TahunLulus: 2027,
			Email: "bob@mail.com", NoTelepon: "081234567891", Alamat: "Bandung",
			UserID: user2.ID, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		},
		{
			NIM: "2023003", Nama: "Alice2", Jurusan: "TI", Angkatan: 2023, TahunLulus: 2027,
			Email: "alice2@mail.com", NoTelepon: "081234567892", Alamat: "Jakarta",
			UserID: user1.ID, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		},
		{
			NIM: "2023004", Nama: "Bob2", Jurusan: "SI", Angkatan: 2023, TahunLulus: 2027,
			Email: "bob2@mail.com", NoTelepon: "081234567893", Alamat: "Bandung",
			UserID: user2.ID, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		},
	}

	// insert hanya kalau belum ada NIM yang sama
	for _, a := range alumni {
		count, _ := alumniColl.CountDocuments(ctx, bson.M{"nim": a.NIM})
		if count == 0 {
			_, err := alumniColl.InsertOne(ctx, a)
			if err != nil {
				log.Printf("failed insert alumni %s: %v", a.Nama, err)
			}
		}
	}
}

func TestAlumniRepository_Integration(t *testing.T) {
	SetupEnvTest(t)

	client, coll := SetupMongoTestCollection(t, "alumni_test")
	defer CleanupMongo(t, client)
	db := coll.Database()

	ctx := context.Background()
	SeedAlumni(ctx, db)

	repo := repository.NewAlumniRepository(db, "alumni_test")

	t.Run("GetAllAlumni", func(t *testing.T) {
		alumni, err := repo.GetAlumni(ctx, "", 100, 0)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(alumni), 2, "harus ada minimal 2 alumni dari seed")
	})

	t.Run("CountAlumni", func(t *testing.T) {
		count, err := repo.CountAlumni(ctx, "")
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, count, int64(2))
	})

	t.Run("CreateAlumni", func(t *testing.T) {
		// ambil user1 dari collection users_test
		var user1 model.User
		err := db.Collection("users_test").FindOne(ctx, bson.M{"username": "user1"}).Decode(&user1)
		assert.NoError(t, err)

		newAlumni := model.Alumni{
			NIM: "2023005", Nama: "Charlie", Jurusan: "TI", Angkatan: 2023,
			TahunLulus: 2027, Email: "charlie@mail.com", NoTelepon: "081234567892",
			Alamat: "Surabaya", CreatedAt: time.Now(), UpdatedAt: time.Now(),
			UserID: user1.ID,
		}

		created, err := repo.CreateAlumni(ctx, &newAlumni)
		assert.NoError(t, err)
		assert.NotEmpty(t, created.ID)
	})

	t.Run("UpdateAlumni", func(t *testing.T) { 
		alumniList, err := repo.GetAlumni(ctx, "Charlie", 10, 0) //aku spesifik in charlie biar kelihatan data yang kuhapus yang mana
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(alumniList), 1, "harus ada alumni Charlie untuk diupdate")

		a := alumniList[0]
		log.Printf("Before update: %+v\n", a)

		updatedData := model.UpdateAlumni{Nama: a.Nama + " Updated"}
		updated, err := repo.UpdateAlumni(ctx, a.ID.Hex(), updatedData)
		assert.NoError(t, err)

		log.Printf("After update: %+v\n", updated)

		assert.Equal(t, a.Nama+" Updated", updated.Nama)
		assert.Equal(t, a.UserID, updated.UserID)
	})

	t.Run("DeleteAlumni", func(t *testing.T) {
		alumniList, err := repo.GetAlumni(ctx, "", 10, 0)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(alumniList), 1, "harus ada alumni untuk dihapus")

		a := alumniList[0]
		t.Logf("Before delete: %v", a)

		err = repo.DeleteAlumni(ctx, a.ID.Hex())
		assert.NoError(t, err)

		remaining, err := repo.GetAlumni(ctx, "", 10, 0)
		assert.NoError(t, err)
		t.Logf("After delete: %v", remaining)

		for _, alum := range remaining {
			assert.NotEqual(t, a.ID.Hex(), alum.ID.Hex())
		}
	})

}
