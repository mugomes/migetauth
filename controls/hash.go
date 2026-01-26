// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package controls

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
}

func VerifyPassword(hash []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}