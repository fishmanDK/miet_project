run:
	clear
	cd assets && go run assets_gen.go assets.go
	cd ..
	go run cmd/main.go 


migrate:
	sql-migrate up -env="local"