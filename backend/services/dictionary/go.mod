module github.com/rkfcccccc/english_words/dictionary

go 1.18

require github.com/stretchr/testify v1.7.1

require (
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/cache/v8 v8.4.3 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/klauspost/compress v1.15.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/vmihailenco/go-tinylfu v0.2.2 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.9.1 // indirect
	golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f // indirect
	golang.org/x/exp v0.0.0-20220518171630-0b5c67f07fdf // indirect
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2 // indirect
	golang.org/x/sync v0.0.0-20220513210516-0976fa681c29 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

require github.com/rkfcccccc/english_words/proto v0.0.0
replace github.com/rkfcccccc/english_words/proto => ../../proto

require github.com/rkfcccccc/english_words/shared_pkg/dsync v0.0.0
replace github.com/rkfcccccc/english_words/shared_pkg/dsync => ../../shared_pkg/dsync

require github.com/rkfcccccc/english_words/shared_pkg/cache v0.0.0
replace github.com/rkfcccccc/english_words/shared_pkg/cache => ../../shared_pkg/cache

require github.com/rkfcccccc/english_words/shared_pkg/mongodb v0.0.0
replace github.com/rkfcccccc/english_words/shared_pkg/mongodb => ../../shared_pkg/mongodb

require github.com/rkfcccccc/english_words/shared_pkg/redis v0.0.0
replace github.com/rkfcccccc/english_words/shared_pkg/redis => ../../shared_pkg/redis
