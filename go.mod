module mymodule

go 1.21.0

//note: had to set GOPRIVATE by go env -w GOPRIVATE='github.com/jathrei/*' meaning i can access any private repo from jathrei
//you have to add to path each time
require github.com/jathrei/puppy v0.0.0-20230811203055-752b49d2eb1f

require github.com/GoesToEleven/puppy v1.2.0

require (
	github.com/GoesToEleven/dog v0.0.0-20230428023317-90bef1c76cb9 // indirect
	github.com/jathrei/dog v0.0.0-20230811201842-e824e878828b // indirect
)
