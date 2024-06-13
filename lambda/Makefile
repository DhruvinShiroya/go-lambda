build:
	@GOOS=linux GOARCH=amd64 go build -o bootstrap
	@zip function.zip bootstrap

# go application run on amd64 archtecture and linux operating system
# build will create bootstrap binary and compress the file function.zip
# name bootstrap is required by AWS lambda function 

