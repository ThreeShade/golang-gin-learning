- Defer is executed as last in first out order
for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
output:9,8,7,6,5,4,3,2,1