// package repository

// import (
//     "database/sql"
//     "go-fiber/app/model"

//     "fmt"
// 	"log"
// )

// // func FindAllAlumni(db *sql.DB) ([]model.Alumni, error) {
// //     rows, err := db.Query(`
// //         SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at 
// //         FROM alumni`)
// //     if err != nil {
// //         return nil, err
// //     }
// //     defer rows.Close()

// //     var alumniList []model.Alumni
// //     for rows.Next() {
// //         var a model.Alumni
// //         if err := rows.Scan(
// //             &a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
// //             &a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat,
// //             &a.CreatedAt, &a.UpdatedAt,
// //         ); err != nil {
// //             return nil, err
// //         }
// //         alumniList = append(alumniList, a)
// //     }
// //     return alumniList, nil
// // }

// // func FindAlumniByID(db *sql.DB, id int) (*model.Alumni, error) {
// //     var a model.Alumni
// //     query := `SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at 
// //               FROM alumni WHERE id=$1 LIMIT 1`
// //     err := db.QueryRow(query, id).
// //         Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
// //             &a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat,
// //             &a.CreatedAt, &a.UpdatedAt)
// //     if err != nil {
// //         return nil, err
// //     }
// //     return &a, nil
// // }

// func CreateAlumni(db *sql.DB, alumni model.Alumni) (*model.Alumni, error) {
//     query := `INSERT INTO alumni (nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat) 
//               VALUES ($1,$2,$3,$4,$5,$6,$7,$8) 
//               RETURNING id, created_at, updated_at`
//     err := db.QueryRow(query,
//         alumni.NIM, alumni.Nama, alumni.Jurusan, alumni.Angkatan,
//         alumni.TahunLulus, alumni.Email, alumni.NoTelepon, alumni.Alamat).
//         Scan(&alumni.ID, &alumni.CreatedAt, &alumni.UpdatedAt)
//     if err != nil {
//         return nil, err
//     }
//     return &alumni, nil
// }

// func UpdateAlumni(db *sql.DB, id int, alumni model.Alumni) (*model.Alumni, error) {
//     query := `UPDATE alumni 
//               SET nim=$1, nama=$2, jurusan=$3, angkatan=$4, tahun_lulus=$5, email=$6, no_telepon=$7, alamat=$8, updated_at=NOW()
//               WHERE id=$9 RETURNING id, created_at, updated_at`
//     err := db.QueryRow(query,
//         alumni.NIM, alumni.Nama, alumni.Jurusan, alumni.Angkatan,
//         alumni.TahunLulus, alumni.Email, alumni.NoTelepon, alumni.Alamat, id).
//         Scan(&alumni.ID, &alumni.CreatedAt, &alumni.UpdatedAt)
//     if err != nil {
//         return nil, err
//     }
//     return &alumni, nil
// }

// func DeleteAlumni(db *sql.DB, id int) error {
//     _, err := db.Exec(`DELETE FROM alumni WHERE id=$1`, id)
//     return err
// }

// //datatable
// func GetAlumniRepo(db *sql.DB, search, sortBy, order string, limit, offset int) ([]model.GetAlumniRepo, error) {
// 	// default columns to prevent SQL injection
// 	allowedSort := map[string]bool{"id": true, "nama": true, "angkatan": true, "tahun_lulus": true}
// 	if !allowedSort[sortBy] {
// 		sortBy = "id"
// 	}
// 	if order != "asc" && order != "desc" {
// 		order = "asc"
// 	}

// 	query := fmt.Sprintf(`
// 		SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at
// 		FROM alumni
// 		WHERE nama ILIKE $1 OR nim ILIKE $1 OR email ILIKE $1
// 		ORDER BY %s %s
// 		LIMIT $2 OFFSET $3
// 	`, sortBy, order)

// 	log.Println("SQL:", query)
// 	log.Println("Params:", "%"+search+"%", limit, offset)

// 	rows, err := db.Query(query, "%"+search+"%", limit, offset)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var alumniList []model.GetAlumniRepo
// 	for rows.Next() {
// 		var a model.GetAlumniRepo
// 		if err := rows.Scan(
// 			&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
// 			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat,
// 			&a.CreatedAt, &a.UpdatedAt,
// 		); err != nil {
// 			return nil, err
// 		}
// 		alumniList = append(alumniList, a)
// 	}
// 	return alumniList, nil
// }

// func CountAlumniRepo(db *sql.DB, search string) (int, error) {
// 	var total int
// 	query := `SELECT COUNT(*) FROM alumni WHERE nama ILIKE $1 OR nim ILIKE $1 OR email ILIKE $1`
// 	err := db.QueryRow(query, "%"+search+"%").Scan(&total)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return total, nil
// }

// func CountAlumniByStatusRepo(db *sql.DB) (map[string]int, error) {
// 	var aktif, resigned, selesai int

// 	// Hitung alumni yang punya pekerjaan aktif (status_pekerjaan = 0)
// 	queryAktif := `
// 		SELECT COUNT(DISTINCT a.id)
// 		FROM alumni a
// 		JOIN pekerjaan_alumni p ON a.id = p.alumni_id
// 		WHERE p.status_pekerjaan = 'aktif' AND p.is_delete IS NULL;
// 	`
// 	err := db.QueryRow(queryAktif).Scan(&aktif)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Hitung alumni yang punya pekerjaan resigned (status_pekerjaan = 1)
// 	queryResigned := `
// 		SELECT COUNT(DISTINCT a.id)
// 		FROM alumni a
// 		JOIN pekerjaan_alumni p ON a.id = p.alumni_id
// 		WHERE p.status_pekerjaan = 'resigned' AND p.is_delete IS NULL;
// 	`
// 	err = db.QueryRow(queryResigned).Scan(&resigned)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Hitung alumni yang punya pekerjaan selesai (status_pekerjaan = 3)
// 	querySelesai := `
// 		SELECT COUNT(DISTINCT a.id)
// 		FROM alumni a
// 		JOIN pekerjaan_alumni p ON a.id = p.alumni_id
// 		WHERE p.status_pekerjaan = 'selesai' AND p.is_delete IS NULL;
// 	`
// 	err = db.QueryRow(querySelesai).Scan(&selesai)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return map[string]int{
// 		"aktif":   aktif,
// 		"resigned": resigned,
// 		"selesai": selesai,
// 	}, nil
// }

package repository

import (
	"context"
	"go-fiber/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"fmt"
)

type AlumniRepository struct {
	collection *mongo.Collection
}

// func NewAlumniRepository(db *mongo.Database) *AlumniRepository {
// 	return &AlumniRepository{
// 		collection: db.Collection("alumni"),
// 	}
// }

func NewAlumniRepository(db *mongo.Database, collectionName string) *AlumniRepository {
    return &AlumniRepository{
        collection: db.Collection(collectionName),
    }
}


// CREATE
func (r *AlumniRepository) CreateAlumni(ctx context.Context, alumni *model.Alumni) (*model.Alumni, error) {
	alumni.ID = primitive.NewObjectID()
	alumni.CreatedAt = time.Now()
	alumni.UpdatedAt = time.Now()

	res, err := r.collection.InsertOne(ctx, alumni)
	if err != nil {
		return nil, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		alumni.ID = oid
	}

	return alumni, nil
}


// UPDATE
func (r *AlumniRepository) UpdateAlumni(ctx context.Context, id string, data model.UpdateAlumni) (*model.Alumni, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	data.UpdatedAt = time.Now()

	update := bson.M{"$set": data}

	res, err := r.collection.UpdateByID(ctx, objID, update)
	if err != nil {
		return nil, err
	}

	if res.MatchedCount == 0 {
		return nil, fmt.Errorf("alumni dengan ID %s tidak ditemukan", id)
	}

	// Ambil data terbaru
	var updated model.Alumni
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

// DELETE
func (r *AlumniRepository) DeleteAlumni(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// GET ALL
func (r *AlumniRepository) GetAlumni(ctx context.Context, search string, limit, offset int64) ([]model.Alumni, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"nama": bson.M{"$regex": search, "$options": "i"}},
			{"nim": bson.M{"$regex": search, "$options": "i"}},
			{"email": bson.M{"$regex": search, "$options": "i"}},
		},
	}

	opts := options.Find().SetSkip(offset).SetLimit(limit)
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var alumni []model.Alumni
	if err = cursor.All(ctx, &alumni); err != nil {
		return nil, err
	}
	return alumni, nil
}

// COUNT
func (r *AlumniRepository) CountAlumni(ctx context.Context, search string) (int64, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"nama": bson.M{"$regex": search, "$options": "i"}},
			{"nim": bson.M{"$regex": search, "$options": "i"}},
			{"email": bson.M{"$regex": search, "$options": "i"}},
		},
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}


