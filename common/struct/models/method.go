package models

import (
	"api2022/util"
	"log"
	"net/http"
)

func (p *Person) Get() (name string) {
	name = p.FirstName + " " + p.LastName
	return
}

func (r Reactangle) Area() (area int) {
	area = r.Lenght * r.Width
	return
}

func (r *Reactangle) Areap() (area int) {
	area = r.Lenght * r.Width
	return
}

func (e Employee) Tch(name string) {
	e.Name = name
	log.Println(e)
}

func (e *Employee) Pch(name string) {
	e.Name = name
	log.Println(e)
}

func (e Employee) ModifyByVal(w http.ResponseWriter) (ex Employee, err error) {
	err = nil
	ex = e

	e.Name = "bb"
	e.Age = 30
	util.Render(w, e)
	ex.Age = 60
	return
}
func (e *Employee) ModifyByRef(w http.ResponseWriter) (ex *Employee, err error) {
	err = nil
	ex = e

	e.Name = "yy"
	e.Age = 31
	util.Render(w, e)
	ex.Age = 61
	return
}
