go get github.com/golang/mock/gomock 
go get github.com/golang/mock/mockgen

$GOPATH/bin/mockgen  <- confirm mockgen library install is success

go get github.com/sgreben/testing-with-gomock/mocks
go get github.com/sgreben/testing-with-gomock/user

within go-mocking folder:
mockgen -destination=mocks/mock_doer.go -package=mocks github.com/sgreben/testing-with-gomock/doer Doer

go test -v github.com/sgreben/testing-with-gomock/user