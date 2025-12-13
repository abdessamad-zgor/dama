package utils

func Assert(condition bool, elseerror string) {
	if !condition {
		panic(errors.New(elseerror))
	}
}
