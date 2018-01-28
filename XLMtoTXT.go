package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

func main() {
	fmt.Println("Hello Hell!!!")
	OutStructToINI()
	fmt.Println("Done.")
	//Close()
}

//Входная структура
//Courses Глобальная секция
type Courses struct {
	XMLName xml.Name `xml:"courses"`
	Courses []Course `xml:"course"`
}

//Course Секции
type Course struct {
	XMLName  xml.Name  `xml:"course"`
	Name     string    `xml:"name,attr"`
	Students []Student `xml:"student"`
}

//Student Ключи со значениями
type Student struct {
	XMLName xml.Name `xml:"student"`
	Name    string   `xml:"name,attr"`
	Mark    string   `xml:"mark,attr"`
}

//--------Входная структура

//Выходная структура
//Students Глобальная секция
type oStudents struct {
	XMLName  xml.Name  `xml:"courses"`
	Students []Student `xml:"student"`
}

//Student Секции
type oStudent struct {
	XMLName xml.Name `xml:"student"`
	Name    string   `xml:"name,attr"`
	Courses []Course `xml:"course"`
}

//Course Ключи со значениями
type oCourse struct {
	XMLName xml.Name `xml:"course"`
	Name    string   `xml:"name,attr"`
	Mark    string   `xml:"mark,attr"`
}

//--------Выходная структура

//OS Выходная структура
var OS OutStruct = OutStruct{"Вася", "Математика", "5"}

//Close Завершение программы, если запущено bat-ником
func Close() {
	fmt.Println("Please press \"Enter\"...")
	fmt.Scanln()
}

//OrigStruct Исходная структура
type OrigStruct struct {
	course   string
	studName string
	mark     string
}

//OutStruct Выходная структура
type OutStruct struct {
	studName string
	course   string
	mark     string
}

//Саша
func XMLtoOrigStruct() {
	// Open our xmlFile
	xmlFile, err := os.Open("courses.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var courses Courses
	xml.Unmarshal(byteValue, &courses)
	fmt.Println(len(courses.Courses))
	//fmt.Println(len(courses.Courses[1].Students))
	for i := 0; i < len(courses.Courses); i++ {
		for j := 0; j < len(courses.Courses[i].Students); j++ {
		fmt.Println(courses.Courses[i].Name)
		fmt.Print(courses.Courses[i].Students[j].Name+" ")
		fmt.Println(courses.Courses[i].Students[j].Mark)
	}
}

//Костя
func OrigStructtoOutStruct() {

}

//FileName Имя файла
var FileName string = "mid"

//StructName Имя структуры
var StructName string = "Students"

//Костя
//OutStructToINI Конвертировать из выходной структуры в INI-файл
func OutStructToINI() {
	//Создать ini-файл
	//os.IsExist(FileName + ".ini")
	file, err := os.Create(FileName + ".ini")
	//Если не удалось создать
	if err != nil {
		//Вывести ошибку
		fmt.Println("Create file " + file.Name() + " failed!")
		os.Exit(1)
		return
	}
	//Открыть для записи ini-файл
	inifile, err := ini.Load("mid.ini")
	//Если не удалось загрузить
	if err != nil {
		//Вывести ошибку
		fmt.Println("Open file failed!")
		os.Exit(1)
		return
	}
	//Считать из структуры данные

	//Записать в ini-файл данные
	//Записать имя структуры
	inifile.NewSection("Struct name")
	inifile.Section("Struct name").NewKey("Name", StructName)

	//Записать новую секцию
	inifile.NewSection(OS.studName)
	//Записать новые данные в секцию
	inifile.Section(OS.studName).NewKey(OS.course, OS.mark)

	//Закрыть файл
	defer inifile.SaveTo("mid.ini")
}

//Аня
func OutStructToTXT() {
	// //Создать и открыть для записи ini-файл
	// inifile, err := os.Create("temp.ini")
	// if err != nil {
	// 	//Вывести ошибку
	// 	fmt.Println("Create file failed!")
	// 	os.Exit(1)
	// 	return
	// }

	// //Считать из структуры данные

	// //Записать в ini-файл данные
	// inifile.WriteString(GroupNames)
	// //Закрыть файл
	// defer inifile.Close()
}
