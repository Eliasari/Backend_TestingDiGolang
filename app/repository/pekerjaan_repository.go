// package repository

// import (
//     "database/sql"
//     "go-fiber/app/model"
//     "fmt"
//     "log"
// "time"
// "errors" 
// )

// // func FindAllPekerjaan(db *sql.DB) ([]model.Pekerjaan, error) {
// //     rows, err := db.Query(`
// //         SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri,
// //                lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja,
// //                status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
// //         FROM pekerjaan_alumni`)
// //     if err != nil {
// //         return nil, err
// //     }
// //     defer rows.Close()

// //     var list []model.Pekerjaan
// //     for rows.Next() {
// //         var p model.Pekerjaan
// //         if err := rows.Scan(
// //             &p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
// //             &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
// //             &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt,
// //         ); err != nil {
// //             return nil, err
// //         }
// //         list = append(list, p)
// //     }
// //     return list, nil
// // }

// // func FindPekerjaanByID(db *sql.DB, id int) (*model.Pekerjaan, error) {
// //     var p model.Pekerjaan
// //     query := `SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri,
// //                      lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja,
// //                      status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
// //               FROM pekerjaan_alumni WHERE id=$1 LIMIT 1`
// //     err := db.QueryRow(query, id).
// //         Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
// //             &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
// //             &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt)
// //     if err != nil {
// //         return nil, err
// //     }
// //     return &p, nil
// // }

// type PekerjaanRepository struct {
// 	DB *sql.DB
// }

// func CreatePekerjaan(db *sql.DB, p model.Pekerjaan) (*model.Pekerjaan, error) {
//     query := `INSERT INTO pekerjaan_alumni 
//               (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range,
//                tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan) 
//               VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) 
//               RETURNING id, created_at, updated_at`
//     err := db.QueryRow(query,
//         p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri,
//         p.LokasiKerja, p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja,
//         p.StatusPekerjaan, p.DeskripsiPekerjaan).
//         Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
//     if err != nil {
//         return nil, err
//     }
//     return &p, nil
// }

// func UpdatePekerjaan(db *sql.DB, id int, p model.Pekerjaan) (*model.Pekerjaan, error) {
//     query := `UPDATE pekerjaan_alumni 
//               SET alumni_id=$1, nama_perusahaan=$2, posisi_jabatan=$3, bidang_industri=$4,
//                   lokasi_kerja=$5, gaji_range=$6, tanggal_mulai_kerja=$7, tanggal_selesai_kerja=$8,
//                   status_pekerjaan=$9, deskripsi_pekerjaan=$10, updated_at=NOW()
//               WHERE id=$11 RETURNING id, created_at, updated_at`
//     err := db.QueryRow(query,
//         p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri,
//         p.LokasiKerja, p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja,
//         p.StatusPekerjaan, p.DeskripsiPekerjaan, id).
//         Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
//     if err != nil {
//         return nil, err
//     }
//     return &p, nil
// }

// func DeletePekerjaan(db *sql.DB, id int) error {
//     _, err := db.Exec(`DELETE FROM pekerjaan_alumni WHERE id=$1`, id)
//     return err
// }

// func FindPekerjaanByAlumniID(db *sql.DB, alumniID int) ([]model.Pekerjaan, error) {
//     rows, err := db.Query(`
//         SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri,
//                lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja,
//                status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
//         FROM pekerjaan_alumni
//         WHERE alumni_id = $1`, alumniID)
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     var list []model.Pekerjaan
//     for rows.Next() {
//         var p model.Pekerjaan
//         if err := rows.Scan(
//             &p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
//             &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
//             &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt,
//         ); err != nil {
//             return nil, err
//         }
//         list = append(list, p)
//     }
//     return list, nil
// }

// func GetPekerjaanRepo(db *sql.DB, search, sortBy, order string, limit, offset int) ([]model.Pekerjaan, error) {
// 	allowedSort := map[string]bool{"id": true, "nama_perusahaan": true, "posisi_jabatan": true, "tanggal_mulai_kerja": true}
// 	if !allowedSort[sortBy] {
// 		sortBy = "id"
// 	}
// 	if order != "asc" && order != "desc" {
// 		order = "asc"
// 	}

// 	query := fmt.Sprintf(`
// 		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range,
// 		       tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan,
// 		       created_at, updated_at
// 		FROM pekerjaan_alumni
// 		WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 OR bidang_industri ILIKE $1 OR lokasi_kerja ILIKE $1
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

// 	var list []model.Pekerjaan
// 	for rows.Next() {
// 		var p model.Pekerjaan
// 		if err := rows.Scan(
// 			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
// 			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
// 			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt,
// 		); err != nil {
// 			return nil, err
// 		}
// 		list = append(list, p)
// 	}

// 	return list, nil
// }

// func CountPekerjaanRepo(db *sql.DB, search string) (int, error) {
// 	var total int
// 	query := `
// 		SELECT COUNT(*) 
// 		FROM pekerjaan_alumni
// 		WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 OR bidang_industri ILIKE $1 OR lokasi_kerja ILIKE $1`
// 	err := db.QueryRow(query, "%"+search+"%").Scan(&total)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return total, nil
// }


// func NewPekerjaanRepository(db *sql.DB) *PekerjaanRepository {
//     return &PekerjaanRepository{DB: db}
// }

// func (r *PekerjaanRepository) SoftDelete(id int, userID int, isAdmin bool) error {
//     now := time.Now()

//     if isAdmin {
//         _, err := r.DB.Exec(`UPDATE pekerjaan_alumni SET is_delete = $1 WHERE id = $2`, now, id)
//         return err
//     }

//     res, err := r.DB.Exec(`
//         UPDATE pekerjaan_alumni p
//         SET is_delete = $1
//         FROM alumni a
//         WHERE p.alumni_id = a.id AND p.id = $2 AND a.user_id = $3
//     `, now, id, userID)
//     if err != nil {
//         return err
//     }

//     rows, _ := res.RowsAffected()
//     if rows == 0 {
//         return sql.ErrNoRows
//     }
//     return nil
// }


// func (r *PekerjaanRepository) GetTrashed(userID int, isAdmin bool) ([]model.PekerjaanTrash, error) {
// 	var rows *sql.Rows
// 	var err error

// 	if isAdmin {
// 		rows, err = r.DB.Query(`
// 			SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri,
// 			       lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja,
// 			       status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at, is_delete
// 			FROM pekerjaan_alumni
// 			WHERE is_delete IS NOT NULL
// 		`)
// 	} else {
// 		rows, err = r.DB.Query(`
// 			SELECT p.id, p.alumni_id, p.nama_perusahaan, p.posisi_jabatan, p.bidang_industri,
// 			       p.lokasi_kerja, p.gaji_range, p.tanggal_mulai_kerja, p.tanggal_selesai_kerja,
// 			       p.status_pekerjaan, p.deskripsi_pekerjaan, p.created_at, p.updated_at, p.is_delete
// 			FROM pekerjaan_alumni p
// 			JOIN alumni a ON p.alumni_id = a.id
// 			WHERE p.is_delete IS NOT NULL AND a.user_id = $1
// 		`, userID)
// 	}

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var result []model.PekerjaanTrash
// 	for rows.Next() {
// 		var p model.PekerjaanTrash
// 		if err := rows.Scan(
// 			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
// 			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
// 			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt, &p.IsDelete,
// 		); err != nil {
// 			return nil, err
// 		}
// 		result = append(result, p)
// 	}

// 	return result, nil
// }

// func (r *PekerjaanRepository) Restore(id, userID int, isAdmin bool) error {
// 	query := `
// 		UPDATE pekerjaan_alumni pa
// 		SET is_delete = NULL
// 		WHERE pa.id = $1
// 		AND (
// 			EXISTS (
// 				SELECT 1 FROM alumni a
// 				WHERE a.id = pa.alumni_id
// 				AND a.user_id = $2
// 			)
// 			OR $3 = TRUE
// 		)
// 	`
// 	res, err := r.DB.Exec(query, id, userID, isAdmin)
// 	if err != nil {
// 		return err
// 	}

// 	rows, _ := res.RowsAffected()
// 	if rows == 0 {
// 		return errors.New("tidak diizinkan restore pekerjaan ini")
// 	}

// 	return nil
// }


// func (r *PekerjaanRepository) HardDelete(id, userID int, isAdmin bool) error {
// 	query := `
// 		DELETE FROM pekerjaan_alumni pa
// 		WHERE pa.id = $1
// 		AND (
// 			EXISTS (
// 				SELECT 1 FROM alumni a
// 				WHERE a.id = pa.alumni_id
// 				AND a.user_id = $2
// 			)
// 			OR $3 = TRUE
// 		)
// 	`

// 	res, err := r.DB.Exec(query, id, userID, isAdmin)
// 	if err != nil {
// 		return err
// 	}

// 	rows, _ := res.RowsAffected()
// 	if rows == 0 {
// 		return errors.New("tidak diizinkan hard delete pekerjaan ini")
// 	}

// 	return nil
// }


package repository

import (
	"context"
	"go-fiber/app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *PekerjaanRepository) GetAlumniCollection() *mongo.Collection {
	return r.colAlumni
}

type PekerjaanRepository struct {
	colPekerjaan *mongo.Collection
	colAlumni    *mongo.Collection
}

func NewPekerjaanRepository(db *mongo.Database, pekerjaanCollName, alumniCollName string) *PekerjaanRepository {
	return &PekerjaanRepository{
		colPekerjaan: db.Collection(pekerjaanCollName),
		colAlumni:    db.Collection(alumniCollName),
	}
}

// CREATE
func (r *PekerjaanRepository) Create(ctx context.Context, pekerjaan *model.Pekerjaan) (*model.Pekerjaan, error) {
	pekerjaan.ID = primitive.NewObjectID()
	pekerjaan.CreatedAt = time.Now()
	pekerjaan.UpdatedAt = time.Now()

	_, err := r.colPekerjaan.InsertOne(ctx, pekerjaan)
	if err != nil {
		return nil, err
	}
	return pekerjaan, nil
}

// UPDATE
func (r *PekerjaanRepository) Update(ctx context.Context, id primitive.ObjectID, update model.UpdatePekerjaan) (*model.Pekerjaan, error) {
	update.UpdatedAt = time.Now()

	_, err := r.colPekerjaan.UpdateByID(ctx, id, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}

	var updated model.Pekerjaan
	err = r.colPekerjaan.FindOne(ctx, bson.M{"_id": id}).Decode(&updated)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

// DELETE
func (r *PekerjaanRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.colPekerjaan.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// GET ALL
func (r *PekerjaanRepository) GetAll(ctx context.Context) ([]model.Pekerjaan, error) {
	cursor, err := r.colPekerjaan.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []model.Pekerjaan
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	return list, nil
}

// GET BY ID
func (r *PekerjaanRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*model.Pekerjaan, error) {
	var pekerjaan model.Pekerjaan
	err := r.colPekerjaan.FindOne(ctx, bson.M{"_id": id}).Decode(&pekerjaan)
	if err != nil {
		return nil, err
	}
	return &pekerjaan, nil
}

// GET BY ALUMNI_ID
func (r *PekerjaanRepository) GetByAlumniID(ctx context.Context, alumniID primitive.ObjectID) ([]model.Pekerjaan, error) {
	cursor, err := r.colPekerjaan.Find(ctx, bson.M{"alumni_id": alumniID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []model.Pekerjaan
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	return list, nil
}

// GET BY USER_ID (user hanya lihat pekerjaan milik alumni-nya)
func (r *PekerjaanRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]model.Pekerjaan, error) {
	// Ambil semua alumni milik user ini
	cursorAlumni, err := r.colAlumni.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursorAlumni.Close(ctx)

	var alumniList []struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	if err := cursorAlumni.All(ctx, &alumniList); err != nil {
		return nil, err
	}

	// ambil semua pekerjaan yang alumni_id-nya ada di list alumni
	var alumniIDs []primitive.ObjectID
	for _, a := range alumniList {
		alumniIDs = append(alumniIDs, a.ID)
	}

	cursor, err := r.colPekerjaan.Find(ctx, bson.M{"alumni_id": bson.M{"$in": alumniIDs}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []model.Pekerjaan
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	return list, nil
}
