module github.com/rkfcccccc/english_words/services/movie

go 1.18

require (
	github.com/georgysavva/scany v1.0.0
	github.com/jackc/pgx/v4 v4.16.1
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/puddle v1.2.1 // indirect
	github.com/joho/godotenv v1.4.0
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/text v0.3.7 // indirect
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/rkfcccccc/english_words/proto => ../../proto

require github.com/rkfcccccc/english_words/proto v0.0.0

replace github.com/rkfcccccc/english_words/shared_pkg/postgres => ../../shared_pkg/postgres

require github.com/rkfcccccc/english_words/shared_pkg/postgres v0.0.0

replace github.com/rkfcccccc/english_words/shared_pkg/services => ../../shared_pkg/services

require github.com/rkfcccccc/english_words/shared_pkg/services v0.0.0
