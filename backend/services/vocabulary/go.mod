module github.com/rkfcccccc/english_words/services/vocabulary

go 1.18

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/puddle v1.2.1 // indirect
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/segmentio/kafka-go v0.4.32
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

require github.com/rkfcccccc/english_words/proto v0.0.0

replace github.com/rkfcccccc/english_words/proto => ../../proto

replace github.com/rkfcccccc/english_words/shared_pkg/dsync => ../../shared_pkg/dsync

require (
	github.com/georgysavva/scany v1.0.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/rkfcccccc/english_words/shared_pkg/postgres v0.0.0
	google.golang.org/grpc v1.46.0
)

replace github.com/rkfcccccc/english_words/shared_pkg/postgres => ../../shared_pkg/postgres

require (
	github.com/jackc/pgx/v4 v4.16.1
	github.com/joho/godotenv v1.4.0
	github.com/rkfcccccc/english_words/shared_pkg/redis v0.0.0
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898 // indirect
)

replace github.com/rkfcccccc/english_words/shared_pkg/redis => ../../shared_pkg/redis
