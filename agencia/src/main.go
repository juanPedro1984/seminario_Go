package main

import ("fmt"
        "io/ioutil"
  "strings"
)

type Agencia struct{
  autos [] Auto
}

type Auto struct{
  marca string
  color string
}

func NewAgencia()Agencia{
  return Agencia{}
}
//el asterisco antes del tipo indica que el return del metodo es un puntero de Auto (no un Auto)
//el & indica que lo que retorna es una direccion de memoria
//
func NewAuto(marca, color string)*Auto {
  return &Auto{
    marca: marca,
    color: color,
  }
}

func SplitString(str string)[]string{
  s := strings.Split(str, "/")
  for x:=0;x<len(s);x++{
    fmt.Println(s[x])
  }
  return s
}

func CargarAutos(str []string , ag Agencia){
  str = str[0:len(str)-1]
  for i:=0 ; i< len(str) ; i++{
    aux:= strings.Split(str[i] ,"-" )
    a:=NewAuto(aux[0],aux[1])
    ag.AddAuto(*a)
  }
}
func (a *Agencia)GetAuto(i int)*Auto{
  if i < len(a.autos){
    return &a.autos[i]
  }else{
    return nil
  }
}

func (a *Agencia)AddAuto(au Auto){
  a.autos = append(a.autos, au)
}

func (a *Agencia)DeleteAuto(i int){
  if i < len(a.autos){
    aux := a.autos[:i]
    aux1 := a.autos[(i+1):]
    aux = append(aux, aux1...)
    a.autos = aux
  }
}

func(a *Agencia)SetAuto(pos int, au Auto){
  if pos < len(a.autos){
    a.autos[pos] = au
  }
}

func (a *Agencia) PrintAutos(){
  for  i:=0; i < len(a.autos) ; i++{
    fmt.Println(a.autos[i])
    }
}

func LeerArchivo(pth string) (string){
  dat , err := ioutil.ReadFile(pth)
  check(err)
  return (string(dat))
}

func ArchivarAuto(s string){
  aux :=LeerArchivo("archivo/autos.txt")
  texto := []byte(aux+s)
  err := ioutil.WriteFile("archivo/autos.txt" , texto, 0644)
  check(err)
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main(){
  //carga de autos desde un archivo

  ag := NewAgencia()
  ArchivarAuto("ford - rojo / ")
  ArchivarAuto("renault - azul/")
  ArchivarAuto("vw - blanco / ")
  ArchivarAuto("citroen - amarillo / ")
  //lee el archivo autos.txt
  //SplitString divide el string en substrings delimitados por "/", y genera un arreglo - cada pocision es un auto
  str := SplitString(LeerArchivo("archivo/autos.txt"))
  //Cargar autos recibe el arreglo de autos y hace Split por "-" divide marca y color y genera otro arreglo con 2 posiciones
  //con esos datos se crea un auto y se agrega al arreglo de autos de la agencia
  CargarAutos(str , ag)
  ag.PrintAutos()

//a := NewAuto("ford","rojo")
//b := NewAuto("renault","azul")
//c := NewAuto("vw" , "blanco")
//d := NewAuto("citroen" , "amarillo")
//
//ag := NewAgencia()
//
//ag.AddAuto(*a)
//ag.AddAuto(*b)
//ag.AddAuto(*c)
//ag.AddAuto(*d)
//
//fmt.Println("Agencia completa")
//ag.PrintAutos()
//
//fmt.Println("Borrar elemento 1 (renault azul)")
//ag.DeleteAuto(1)
//ag.PrintAutos()
//
//fmt.Println("Modificar auto en pos 1")
//e:= NewAuto("Peugeot","verde")
//ag.SetAuto(1, *e)
//ag.PrintAutos()
//
//ArchivarAuto(a.marca +" - "+ a.color + " / " )
//ArchivarAuto(b.marca +" - "+ b.color + " / ")
//ArchivarAuto(c.marca +" - "+ c.color + " / ")
//ArchivarAuto(d.marca +" - "+ d.color + " / ")
//fmt.Println(LeerArchivo("archivo/autos.txt"))

}
