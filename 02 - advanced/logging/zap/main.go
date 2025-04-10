package main

import "go.uber.org/zap"

/*
zap
- fast and production ready
- structured output
- configuable encoders, cores, hooks
*/

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("User login", zap.String("user", "admin"), zap.Int("age", 30))

	sugar := logger.Sugar()
	sugar.Infof("User %s has logged in", "admin")
}
