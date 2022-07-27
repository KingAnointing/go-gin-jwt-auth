package configs

import (
	"context"
	"time"
)

func DatabaseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}
