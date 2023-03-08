module github.com/soulteary/simple-kms

go 1.20

require (
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/google/uuid v1.3.0
)

require github.com/soulteary/go-cloud-id v0.1.2

replace github.com/denisbrodbeck/machineid => ./pkg/denisbrodbeck/machineid
