package temperature

//定义包的名称为temperature

//提供给外部使用的func
func CtoF(c float64) float64 {
	return (c * 9 / 5) + 32
}

//提供给外部使用的func
func FtoC(f float64) float64 {
	return (f - 32) * (9 / 5)
}

//私有方法，外部不应该调用
func privateFunc() string {
	return "i am private func,not call me"
}
