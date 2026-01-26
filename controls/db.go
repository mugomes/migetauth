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
		`ALTER TABLE xht_usuarios ADD COLUMN senha BLOB;`,

		// tabela contas
		`CREATE TABLE IF NOT EXISTS xht_contas (
			id INTEGER PRIMARY KEY AUTOINCREMENT
		);`,

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
	err := db.QueryRow(`SELECT COUNT(*) FROM xht_usuarios`).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateUser(db *sql.DB, nome, email string, senha string) error {
	hash, err := HashPassword(senha)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`INSERT INTO xht_usuarios (nome, email, senha)
		VALUES (?, ?, ?)`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(nome, email, hash)

	return err
}
