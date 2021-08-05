![Last test result](https://github.com/dominikkeisz/go-aws-s3-object-list/actions/workflows/ci.yml/badge.svg)

# go-aws-s3-object-list
Go program for listing objects of your desired AWS S3 bucket

### Configuration
*~/.aws/config* needs to hold your region at least  
*~/.aws/crendentials* needs to hold your Access Key ID and Secret Access Key at least  

### Use it like this:
```
go build
./go-aws-s3-object-list mybucketname
```