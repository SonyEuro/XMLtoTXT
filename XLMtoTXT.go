package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-ini/ini"
)

/*
main - Конвертировать из XML-файла в TXT-файл
*/
func main() {
	FileName := GetFileName()
	OriginalStructToOriginalINI(XMLtoOriginalStruct(FileName), FileName)
	OriginalINItoOutINI(FileName)
	OutStructToTXT(OutINItoOutStruct(FileName), FileName)
	Close()
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

/*
Close - Завершить выполнение программы после запуска bat-файлом
*/
func Close() {
	fmt.Println("Done.")
	fmt.Println("Please press \"Enter\"...")
	fmt.Scanln()
}

//Илья
/*
GetFileName - Получить имя обрабатываемого файла из командной строки
*/
func GetFileName() string {
	//Сканировать список аргументов
	flag.Parse()
	//Вернуть нулевой аргумент командной строки
	return flag.Arg(0)
}

//Саша
/*
XMLtoOriginalStruct - Конвертировать из XML-файла в исходную структуру
inXML - входной параметр, XML-файл
Courses - выходной параметр, исходная структура
*/
func XMLtoOriginalStruct(inXML string) Courses {
	//Добавить расшерение в имя файла
	inXML += ".xml"
	// Открыть файл XML
	xmlFile, err := os.Open(inXML)
	// Вывести ошибку
	if err != nil {
		fmt.Println("Open file " + inXML + " failed!")
	}
	// Закрыть файл
	defer xmlFile.Close()
	// Записать значения XML-файла в массив
	byteValue, _ := ioutil.ReadAll(xmlFile)
	// Создать экземпляр структуры Courses
	var courses Courses
	// Записать значения из массива в экземпляр структуры
	xml.Unmarshal(byteValue, &courses)
	// Вернуть экземпляр структуры Courses
	return courses
}

//Костя
// Разобраться с файлами, не открывает другой почему то
/*
OriginalStructToOriginalINI - Конвертировать из входной структуры в INI-файл
origStruct - входной параметр, исходная структура
FileName - имя выходного файла
*/
func OriginalStructToOriginalINI(origStruct Courses, FileName string) {
	//Открыть для записи ini-файл
	inifileIn := ini.Empty()

	//Считать из структуры данные и записать их в INI-файл

	//Записать имя структуры
	inifileIn.NewSection("XMLName")
	inifileIn.Section("XMLName").NewKey("XMLName", origStruct.XMLName.Local)

	for i := 0; i < len(origStruct.Courses); i++ {
		//Записать новую секцию
		inifileIn.NewSection(origStruct.Courses[i].Name)
		for j := 0; j < len(origStruct.Courses[i].Students); j++ {
			//Записать новые данные в секцию
			inifileIn.Section(origStruct.Courses[i].Name).NewKey(origStruct.Courses[i].Students[j].Name, origStruct.Courses[i].Students[j].Mark)
		}
	}

	//Закрыть файл
	defer inifileIn.SaveTo(FileName + " original.ini")
}

//Илья должен был сделать
//Костя
//OriginalINItoOutINI Конвертирует INI-файл с входной структурой в INI-файл с выходной структурой
/*
OriginalINItoOutINI Конвертировать из INI-файла в  выходную структуру
FileName - имя входного файла
*/
func OriginalINItoOutINI(FileName string) {
	//Открыть для чтения ini-файл
	iniFileIn, err := ini.Load(FileName + " original.ini")
	//Если не удалось загрузить
	if err != nil {
		//Вывести ошибку
		fmt.Println("Open file " + FileName + " original.ini" + " failed!")
		os.Exit(1)
	}

	//Открыть для записи ini-файл
	iniFileOut := ini.Empty()
	//Получить массив имён секций
	sectionsNames := iniFileIn.SectionStrings()
	for sections := 1; sections < len(sectionsNames); sections++ {
		//Получить массив имён ключей
		var keysNames = iniFileIn.Section(sectionsNames[sections]).KeyStrings()
		for keys := 0; keys < len(keysNames); keys++ {
			//Создать новую секцию, если её не существует
			tempSection := iniFileOut.Section(keysNames[keys])
			if tempSection != nil {
				iniFileOut.NewSection(keysNames[keys])
			}
			//Получить значение по ключу
			var value = iniFileIn.Section(sectionsNames[sections]).Key(keysNames[keys]).String()
			//Записать ключ=значение
			iniFileOut.Section(keysNames[keys]).NewKey(sectionsNames[sections], value)
		}
	}
	//Сохранить и закрыть файл
	defer iniFileOut.SaveTo(FileName + " out.ini")
}

//Костя
/*
OutINItoOutStruct Конвертировать из INI-файла в  выходную структуру
FileName - имя входного файла
oStudents - выходной параметр, экземпляр структуры
*/
func OutINItoOutStruct(FileName string) oStudents {
	//Открыть для записи ini-файл
	inifileIn, err := ini.Load(FileName + " out.ini")
	//Если не удалось загрузить
	if err != nil {
		//Вывести ошибку
		fmt.Println("Open file " + FileName + " out.ini" + " failed!")
		os.Exit(1)
	}

	//Создать экземпляр струкуры oStudents
	var OSs oStudents
	//Задать имя структуры
	OSs.XMLName.Local = "Students"

	//Получить массив имен секции
	sectionsNames := inifileIn.SectionStrings()

	//Записать значения INI-файла в выходную структуру
	for sections := 2; sections < len(sectionsNames); sections++ {
		//Получить имена ключей
		var keysNames = inifileIn.Section(sectionsNames[sections]).KeyStrings()
		//OSs.Students:=append (OSs.Students,sections)
		//Создать экземпляр структуры oStudents
		var OS oStudent
		//Получить имя секции
		OS.Name = sectionsNames[sections]
		//Получить ключ-значение и записать в структуру
		for keys := 0; keys < len(keysNames); keys++ {
			//OSs.Students.Course.Name:=keys
			//Получить значение по ключу
			var value = inifileIn.Section(sectionsNames[sections]).Key(keysNames[keys]).String()
			//OSs.Students.Course.Mark:=value
			//Создать экземпляр структуры oCourse
			var OC oCourse
			//Записать значения в структуру
			OC.Mark = value
			OC.Name = keysNames[keys]
			//Добавить экземпляр структуры oCourse в массив структур
			OS.Courses = append(OS.Courses, OC)
		}
		//Добавить экземпляр структуры oStudent в массив структур
		OSs.Students = append(OSs.Students, OS)
	}
	//Вернуть экземпляр структуры oStudents
	return OSs
}

//Аня
/*
OutStructToTXT Конвертировать из выходной структуры в файл формата *.txt
outStruct (выходная структура) - входной параметр
FileName - имя выходного файла
*/
func OutStructToTXT(outStruct oStudents, FileName string) {
	//Открыть для записи txt-файл
	file, err := os.Create(FileName + " result.txt")
	//Если не удалось загрузить
	if err != nil {
		//Вывести ошибку
		fmt.Println("Open file " + FileName + " result.txt" + " failed!")
		os.Exit(1)
	}

	//Записать значения выходной структуры в файл
	file.WriteString(outStruct.XMLName.Local + "\n\n")
	for i := 0; i < len(outStruct.Students); i++ {
		//Записать имя студента
		file.WriteString("student " + outStruct.Students[i].Name + "\n")
		for j := 0; j < len(outStruct.Students[i].Courses); j++ {
			//Записать название курса
			file.WriteString("course " + outStruct.Students[i].Courses[j].Name + "\t")
			//Записать оценку
			file.WriteString("mark " + outStruct.Students[i].Courses[j].Mark + "\n")
		}
		file.WriteString("\n")
	}

	// Закрыть файл
	defer file.Close()
}
