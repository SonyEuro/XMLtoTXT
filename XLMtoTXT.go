package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-ini/ini"
)

func main() {
	fmt.Println("Hello Hell!!!")
	OrigStructToOrigINI(XMLtoOrigStruct("in.xml"))
	OutStructToTXT(OutINItoOutStruct("courses.ini"))
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
	XMLName  xml.Name   `xml:"courses"`
	Students []oStudent `xml:"student"`
}

//Student Секции
type oStudent struct {
	XMLName xml.Name  `xml:"student"`
	Name    string    `xml:"name,attr"`
	Courses []oCourse `xml:"course"`
}

//Course Ключи со значениями
type oCourse struct {
	XMLName xml.Name `xml:"course"`
	Name    string   `xml:"name,attr"`
	Mark    string   `xml:"mark,attr"`
}

//--------Выходная структура

//Close Завершение программы, если запущено bat-ником
func Close() {
	fmt.Println("Please press \"Enter\"...")
	fmt.Scanln()
}

//Саша
//XMLtoOrigStruct
func XMLtoOrigStruct(inXML string) Courses {
	// Open our xmlFile
	xmlFile, err := os.Open(inXML)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened " + inXML)
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var courses Courses
	xml.Unmarshal(byteValue, &courses)
	// fmt.Println(len(courses.Courses))
	// //fmt.Println(len(courses.Courses[1].Students))
	// for i := 0; i < len(courses.Courses); i++ {
	// 	for j := 0; j < len(courses.Courses[i].Students); j++ {
	// 		fmt.Println(courses.Courses[i].Name)
	// 		fmt.Print(courses.Courses[i].Students[j].Name + " ")
	// 		fmt.Println(courses.Courses[i].Students[j].Mark)
	// 	}
	// }
	return courses
}

//Костя
// Разобраться с файлами, не открывает другой почему то
//OrigStructToOrigINI Конвертировать из входной структуры в INI-файл
func OrigStructToOrigINI(origStruct Courses) {
	// //Создать ini-файл
	// //os.IsExist(FileName + ".ini")
	// file, err := os.Create(origStruct.XMLName.Local + ".ini")
	// //Если не удалось создать
	// if err != nil {
	// 	//Вывести ошибку
	// 	fmt.Println("Create file " + file.Name() + " failed!")
	// 	os.Exit(1)
	// 	return
	// }
	// file.Close()

	//Открыть для записи ini-файл
	inifile, err := ini.Load("courses.ini")
	//Если не удалось загрузить
	if err != nil {
		//Вывести ошибку
		fmt.Println("Open file " + "courses.ini" + " failed!")
		os.Exit(1)
		return
	}

	//Считать из структуры данные и записать их в INI-файл

	//Записать имя структуры
	inifile.NewSection("XMLName")
	inifile.Section("XMLName").NewKey("XMLName", origStruct.XMLName.Local)

	for i := 0; i < len(origStruct.Courses); i++ {
		//Записать новую секцию
		inifile.NewSection(origStruct.Courses[i].Name)
		for j := 0; j < len(origStruct.Courses[i].Students); j++ {
			//Записать новые данные в секцию
			inifile.Section(origStruct.Courses[i].Name).NewKey(origStruct.Courses[i].Students[j].Name, origStruct.Courses[i].Students[j].Mark)
		}
	}

	//Закрыть файл
	defer inifile.SaveTo("courses.ini")
}

//Илья
//OrigINItoOutINI Конвертирует INI-файл с входной структурой в INI-файл с выходной структурой
func OrigINItoOutINI(inINI string) {

}

//Костя
//OutINItoOutStruct Конвертировать из INI-файла в  выходную структуру
func OutINItoOutStruct(outINI string) oStudents {
	// //Создать ini-файл
	// //os.IsExist(FileName + ".ini")
	// file, err := os.Create(outINI)
	// //Если не удалось создать
	// if err != nil {
	// 	//Вывести ошибку
	// 	fmt.Println("Create file " + file.Name() + " failed!")
	// 	os.Exit(1)
	// 	//return nil
	// }

	//Открыть для записи ini-файл
	inifile, err := ini.Load("courses.ini")
	//Если не удалось загрузить
	if err != nil {
		//Вывести ошибку
		fmt.Println("Open file failed!")
		os.Exit(1)
		//return nil
	}

	var OSs oStudents
	OSs.XMLName.Local = "Students"
	//var OS oStudent
	//var OC oCourse

	//var sections=inifile.Sections
	sectionsNames := inifile.SectionStrings()

	//var keys=inifile.Section(sectionsNames[])
	//var keysNames=inifile.Section(sectionsNames[]).KeyStrings()

	//var value=inifile.Section(sectionsNames[]).Key(keysNames[]).String()

	for sections := 2; sections < len(sectionsNames); sections++ {
		var keysNames = inifile.Section(sectionsNames[sections]).KeyStrings()
		//OSs.Students:=append (OSs.Students,sections)
		var OS oStudent
		OS.Name = sectionsNames[sections]
		for keys := 0; keys < len(keysNames); keys++ {
			//OSs.Students.Course.Name:=keys
			var value = inifile.Section(sectionsNames[sections]).Key(keysNames[keys]).String()
			//OSs.Students.Course.Mark:=value
			var OC oCourse
			OC.Mark = value
			OC.Name = keysNames[keys]
			OS.Courses = append(OS.Courses, OC)
		}
		OSs.Students = append(OSs.Students, OS)
	}
	return OSs
}

//Аня
//OutStructToTXT
func OutStructToTXT(outStruct oStudents) {
	//Открыть для записи txt-файл
	file, err := os.Create("result.txt") //, os.O_APPEND|os.O_WRONLY, 0666)
	//Если не удалось загрузить
	if err != nil {
		//Вывести ошибку
		fmt.Println("Open file failed!")
		os.Exit(1)
		return
	}

	//запись в файл
	file.WriteString(outStruct.XMLName.Local + "\n\n")
	for i := 0; i < len(outStruct.Students); i++ {
		file.WriteString("student " + outStruct.Students[i].Name + "\n")
		for j := 0; j < len(outStruct.Students[i].Courses); j++ {
			file.WriteString("course " + outStruct.Students[i].Courses[j].Name + "\t")
			file.WriteString("mark " + outStruct.Students[i].Courses[j].Mark + "\n")
		}
		file.WriteString("\n")
	}

	// Закрыть файл
	defer file.Close()
}
