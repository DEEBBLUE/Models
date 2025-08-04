package utils

func FanOut[T,A any](ch chan A,count int,callback func(chan A) chan T) []chan T{
	listCh := make([]chan T,0,count)

	for i := range count{
		listCh[i] = callback(ch)	
	}	

	return listCh
}
