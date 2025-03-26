package main
import "fmt"

func iceCream(kamar, tempatMain, kebun int)string{
	if kamar >= 60 && tempatMain >= 75 && kebun >= 60{
		return "Ice Cream"
	}else if kamar >= 80 && kebun >= 80{
		return "Ice Cream"
	}else if kamar == 100 {
		return "Ice Cream"
	}else{
		return "Tidak"
	}
}
func main(){
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Print(iceCream(x, y, z))
}