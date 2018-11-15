package main

func batch(inputChannel chan struct{}, size int) chan []struct{} {
	outputChannel := make(chan []struct{})

	go func() {
		for {
			var outputArray []struct{}
			for i := 0; i < size; i++ {
				outputArray = append(outputArray, <-inputChannel)
			}
			outputChannel <- outputArray
		}
	}()

	return outputChannel

}
