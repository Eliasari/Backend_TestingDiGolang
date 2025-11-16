package integration

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go-fiber/app/model"
	"go-fiber/app/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SeedPekerjaanAlumni insert dummy alumni + pekerjaan jika belum ada
func SeedPekerjaanAlumni(ctx context.Context, db *mongo.Database) (primitive.ObjectID, primitive.ObjectID) {
	alumniColl := db.Collection("alumni_test")
	pekerjaanColl := db.Collection("pekerjaan_alumni_test")

	// ambil 2 alumni
	var alumni1, alumni2 model.Alumni
	_ = alumniColl.FindOne(ctx, bson.M{"nama": "Alice"}).Decode(&alumni1)
	_ = alumniColl.FindOne(ctx, bson.M{"nama": "Bob"}).Decode(&alumni2)

	// insert pekerjaan dummy
	pekerjaan := []model.Pekerjaan{
		{
			AlumniID:           alumni1.ID,
			NamaPerusahaan:     "Tokopedia",
			PosisiJabatan:      "Backend Developer",
			BidangIndustri:     "E-commerce",
			LokasiKerja:        "Jakarta",
			GajiRange:          "8-12 juta",
			TanggalMulaiKerja:  time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
			StatusPekerjaan:    "aktif",
			DeskripsiPekerjaan: "Mengembangkan dan memelihara API produk Tokopedia",
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		},
		{
			AlumniID:           alumni2.ID,
			NamaPerusahaan:     "Gojek",
			PosisiJabatan:      "Frontend Developer",
			BidangIndustri:     "Transportasi",
			LokasiKerja:        "Jakarta",
			GajiRange:          "7-10 juta",
			TanggalMulaiKerja:  time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
			StatusPekerjaan:    "aktif",
			DeskripsiPekerjaan: "Membuat UI dan UX aplikasi Gojek",
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		},
	}

	for _, p := range pekerjaan {
		count, _ := pekerjaanColl.CountDocuments(ctx, bson.M{"nama_perusahaan": p.NamaPerusahaan, "alumni_id": p.AlumniID})
		if count == 0 {
			p.ID = primitive.NewObjectID()
			_, _ = pekerjaanColl.InsertOne(ctx, p)
		}
	}

	return alumni1.ID, alumni2.ID
}

func TestPekerjaanRepository_Integration(t *testing.T) {
	SetupEnvTest(t)

	client, coll := SetupMongoTestCollection(t, "pekerjaan_alumni_test")
	defer CleanupMongo(t, client)
	db := coll.Database()

	ctx := context.Background()
	alumni1ID, _ := SeedPekerjaanAlumni(ctx, db)

	repo := repository.NewPekerjaanRepository(db, "pekerjaan_alumni_test", "alumni_test")

	t.Run("GetAllPekerjaan", func(t *testing.T) {
		list, err := repo.GetAll(ctx)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(list), 2)
		log.Println("GetAllPekerjaan:", list)
	})

	t.Run("CreatePekerjaan", func(t *testing.T) {
		pekerjaan := model.Pekerjaan{
			AlumniID:           alumni1ID,
			NamaPerusahaan:     "Shopee",
			PosisiJabatan:      "DevOps",
			BidangIndustri:     "E-commerce",
			LokasiKerja:        "Jakarta",
			GajiRange:          "10-15 juta",
			TanggalMulaiKerja:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			StatusPekerjaan:    "aktif",
			DeskripsiPekerjaan: "Membuat pipeline CI/CD",
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}
		pekerjaan.ID = primitive.NewObjectID()
		log.Println("Before create:", pekerjaan)
		created, err := repo.Create(ctx, &pekerjaan)
		assert.NoError(t, err)
		assert.NotEmpty(t, created.ID)
		log.Println("After create:", created)
	})

	t.Run("UpdatePekerjaan", func(t *testing.T) {
		list, _ := repo.GetAll(ctx)
		p := list[0]

		log.Println("Before update:", p)

		update := model.UpdatePekerjaan{
			PosisiJabatan: "Senior Backend Developer",
			UpdatedAt:     time.Now(),
		}
		updated, err := repo.Update(ctx, p.ID, update)
		assert.NoError(t, err)
		assert.Equal(t, "Senior Backend Developer", updated.PosisiJabatan)

		log.Println("After update:", updated)
	})

	t.Run("DeletePekerjaan", func(t *testing.T) {
		list, _ := repo.GetAll(ctx)
		p := list[0]

		log.Println("Before delete:", p)
		err := repo.Delete(ctx, p.ID)
		assert.NoError(t, err)

		remaining, _ := repo.GetAll(ctx)
		log.Println("After delete:", remaining)

		for _, r := range remaining {
			assert.NotEqual(t, p.ID, r.ID)
		}
	})
}
