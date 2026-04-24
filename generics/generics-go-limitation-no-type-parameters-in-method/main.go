package genericsgolimitationadditionaltypeparameters

func main() {}

type functionalSlice[T any] []T

//Method cannot have type parameters in go
// func (fs functionalSlice[T]) Map[E any](f func(T) E) functionalSlice[E] {
// out := make(functionalSlice[E], len(fs))
// for i, v := range fs {
// out[i] = f(v)
// }
// return out
// }

// Changing above method to a function
func Map[T, E any](fs functionalSlice[T], f func(T) E) functionalSlice[E] {
	out := make(functionalSlice[E], len(fs))
	for i, v := range fs {
		out[i] = f(v)
	}
	return out
}
