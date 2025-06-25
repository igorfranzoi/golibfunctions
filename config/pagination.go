package config

import (
	"os"
	"strconv"
)

type ConfigPagination struct {
	DefaultLimit int
	MaxLimit     int
}

var DefaultConfig = loadConfigPagination()

func loadConfigPagination() ConfigPagination {
	defaultLimit := 10
	maxLimit := 50

	if os.Getenv("ITEMS_PER_PAGE") != "" {
		if valAux, err := strconv.Atoi(os.Getenv("ITEMS_PER_PAGE")); err == nil {
			defaultLimit = valAux
		}
	}
	if os.Getenv("ITEMS_MAX_PAGE") != "" {
		if valAux, err := strconv.Atoi(os.Getenv("ITEMS_MAX_PAGE")); err == nil {
			maxLimit = valAux
		}

	}

	return ConfigPagination{
		DefaultLimit: defaultLimit,
		MaxLimit:     maxLimit,
	}
}
