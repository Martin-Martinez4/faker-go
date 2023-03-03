package english

func RandomEntryFromSlice[T any](slice *[]T, seed_optional ...int64) T {

	index := GenerateRandIntBelowMaximum(len(*slice), seed_optional...)

	return (*slice)[index]
}
