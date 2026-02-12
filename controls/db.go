// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package controls

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(db *sql.DB) error {
	queries := []string{
		// tabela usuario
		`CREATE TABLE IF NOT EXISTS xht_usuarios (
			id INTEGER PRIMARY KEY AUTOINCREMENT
		);`,

		`ALTER TABLE xht_usuarios ADD COLUMN nome TEXT;`,
		`ALTER TABLE xht_usuarios ADD COLUMN email TEXT;`,
		`ALTER TABLE xht_usuarios ADD COLUMN usuario TEXT;`,
		`ALTER TABLE xht_usuarios ADD COLUMN senha BLOB;`,

		// tabela contas
		`CREATE TABLE IF NOT EXISTS xht_contas (
			id INTEGER PRIMARY KEY AUTOINCREMENT
		);`,

		`ALTER TABLE xht_contas ADD COLUMN idusuario INT;`,
		`ALTER TABLE xht_contas ADD COLUMN nome TEXT;`,
		`ALTER TABLE xht_contas ADD COLUMN codigo BLOB;`,
	}

	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			// ignora erro de coluna já existente
			if !isDuplicateColumnError(err) {
				return err
			}
		}
	}

	return nil
}

func isDuplicateColumnError(err error) bool {
	return err != nil &&
		(strings.Contains(err.Error(), "duplicate column") ||
			strings.Contains(err.Error(), "already exists"))
}

// Users
func UserExists(db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM xht_usuarios LIMIT 0,1`).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateUser(db *sql.DB, nome, email, usuario, senha string) (int64,error) {
	hash, err := HashPassword(senha)
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(`INSERT INTO xht_usuarios (nome, email, usuario, senha)
		VALUES (?, ?, ?, ?)`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	query, err := stmt.Exec(nome, email, usuario, hash)
	id, _ := query.LastInsertId()

	return id, err
}

func CheckUser(db *sql.DB, usuario, senha string) (bool, error) {
	var sSenha string

	stmt, err := db.Prepare(`SELECT senha FROM xht_usuarios WHERE usuario=?`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(usuario).Scan(&sSenha)

	if err != nil {
		return false, err
	}

	if VerifyPassword([]byte(sSenha), senha) {
		return true, nil
	} else {
		return false, nil
	}
}

// Contas
func AccountExists(db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM xht_contas LIMIT 0,1`).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
