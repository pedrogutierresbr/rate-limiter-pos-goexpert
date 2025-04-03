package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/pedrogutierresbr/rate-limiter-pos-goexpert/internal/limiter"
	"github.com/pedrogutierresbr/rate-limiter-pos-goexpert/internal/middleware"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Erro ao ler o arquivo de configuração: %s\n", err)
		return
	}
	viper.AutomaticEnv()
}

func main() {
	initConfig()

	limitIP := viper.GetInt("RATE_LIMIT_IP")
	limitToken := viper.GetInt("RATE_LIMIT_TOKEN")
	blockTime := viper.GetInt("BLOCK_TIME")

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PORT")),
	})

	store := limiter.NewRedisStore(rdb)

	rl := limiter.NewRateLimiter(store, limitIP, limitToken, blockTime)

	http.Handle("/", middleware.RateLimiter(rl, http.HandlerFunc(handler)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %s\n", err)
	}

	fmt.Printf("Servidor rodando na porta 8080 ...\n")
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, friend!\n"))
}
