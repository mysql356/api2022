package models

func Area(r Reactangle) (area int) {
	area = r.Lenght * r.Width
	return
}

func Areap(r *Reactangle) (area int) {
	area = r.Lenght * r.Width
	return
}
