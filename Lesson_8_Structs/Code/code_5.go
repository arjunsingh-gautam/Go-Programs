// composition in go
package main

import "fmt"

type Battery struct{
	Capacity int
}
func (b *Battery) BatteryCapacity(){
	fmt.Printf("Battery capacity is:%d\n",b.Capacity)
}

type Camera struct{
	Megapixel int
}
func (c *Camera) DisplayPixels(){
	fmt.Printf("Camera Pixels is:%d\n",c.Megapixel)
}

type MobilePhone struct{
	Brand string
	Camera
	Battery
}

func main(){
	mobile1:=&MobilePhone{
		Brand:"Oppo",
		Camera:Camera{
			Megapixel: 48,
		},
		Battery: Battery{Capacity: 5000},
	}

	mobile1.BatteryCapacity()
	mobile1.DisplayPixels()
	fmt.Printf("Mobile Info:%v",mobile1)
}