package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type bolaRepositoryImpl struct {
	DB *sql.DB
}

func NewBolaRepository(db *sql.DB) BolaRepository {
	return &bolaRepositoryImpl{db}

}

func (repository *bolaRepositoryImpl) Insert(ctx context.Context, bola entity.Bola) (entity.Bola, error) {
	script := "INSERT INTO bola(posisi, negara) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, bola.Posisi, bola.Negara)
	if err != nil {
		return bola, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return bola, err
	}
	bola.Id = int32(id)
	return bola, nil
	//panic("implement me")
}

func (repository *bolaRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Bola, error) {
	//TODO implement me
	script := "SELECt ID, posisi, negara FROM bola WHERE ID = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	bola := entity.Bola{}
	if err != nil {
		return bola, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&bola.Id, &bola.Id, &bola.Posisi)
		return bola, nil
	} else {
		//tidak ada
		return bola, errors.New("Id" + strconv.Itoa(int(id)) + " Nor Found")
	}
	//panic("implement me")
}

func (repository *bolaRepositoryImpl) FindAll(ctx context.Context) ([]entity.Bola, error) {
	//TODO implement me
	script := "SELECt ID, negara, posisi FROM bola"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var bolas []entity.Bola
	for rows.Next() {
		bola := entity.Bola{}
		rows.Scan(&bola.Id, &bola.Posisi, &bola.Negara)
		bolas = append(bolas, bola)
	}
	return bolas, nil
	//panic("implement me")
}

func (repository *bolaRepositoryImpl) Update(ctx context.Context, id int32, bola entity.Bola) (entity.Bola, error) {
	//TODO implement me
	script := "SELECT ID, posisi, negara FROM bola WHERE ID = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return bola, err
	}
	if rows.Next() {
		// yes
		script := "UPDATE bola SET posisi = ?, negara = ? WHERE ID = ?"
		_, err := repository.DB.ExecContext(ctx, script, bola.Posisi, bola.Negara, id)
		if err != nil {
			return bola, err
		}
		bola.Id = id
		return bola, nil
	} else {
		// no
		return bola, errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
	}
}

func (repository *bolaRepositoryImpl) Delete(ctx context.Context, id int32) (string, error) {
	script := "SELECT ID, posisi, negara FROM bola WHERE ID = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "Gagal", err
	}
	if rows.Next() {

		script := "DELETE FROM bola WHERE ID = ?"
		_, err := repository.DB.ExecContext(ctx, script, id)
		if err != nil {
			return "Id" + strconv.Itoa(int(id)) + "Gagal", err
		}
		return "Id" + strconv.Itoa(int(id)) + "Sukses", nil
	} else {

		return "Gagal", errors.New(("Id" + strconv.Itoa(int(id)) + "tidak ada"))
	}
}
