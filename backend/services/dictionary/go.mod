module github.com/rkfcccccc/english_words/services/dictionary

go 1.18

require github.com/stretchr/testify v1.7.1

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-redsync/redsync/v4 v4.5.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/klauspost/compress v1.15.7 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898 // indirect
	golang.org/x/net v0.0.0-20220706163947-c90051bbdb60 // indirect
	golang.org/x/sync v0.0.0-20220513210516-0976fa681c29 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

require (
	github.com/rkfcccccc/english_words/proto v0.0.0
	github.com/rkfcccccc/english_words/shared_pkg/cache v0.0.0
	github.com/rkfcccccc/english_words/shared_pkg/dsync v0.0.0
	github.com/rkfcccccc/english_words/shared_pkg/mongodb v0.0.0
	github.com/rkfcccccc/english_words/shared_pkg/redis v0.0.0
	github.com/rkfcccccc/english_words/shared_pkg/services v0.0.0
)

replace github.com/rkfcccccc/english_words/proto => ../../proto

replace github.com/rkfcccccc/english_words/shared_pkg/redis => ../../shared_pkg/redis

replace github.com/rkfcccccc/english_words/shared_pkg/dsync => ../../shared_pkg/dsync

replace github.com/rkfcccccc/english_words/shared_pkg/cache => ../../shared_pkg/cache

replace github.com/rkfcccccc/english_words/shared_pkg/mongodb => ../../shared_pkg/mongodb

replace github.com/rkfcccccc/english_words/shared_pkg/services => ../../shared_pkg/services

require (
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/joho/godotenv v1.4.0
	go.mongodb.org/mongo-driver v1.9.1
	google.golang.org/grpc v1.48.0
)
