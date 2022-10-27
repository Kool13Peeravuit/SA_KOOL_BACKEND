package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name       string
	BorrowDay  int
	BookRoomHR int
	BookComHR  int
	Users      []User `gorm:"foreignKey:RoleID"`
}

type Province struct {
	gorm.Model
	Name  string
	Users []User `gorm:"foreignKey:ProvinceID"`
}

type MemberClass struct {
	gorm.Model
	Name     string
	Discount float32
	Users    []User `gorm:"foreignKey:MemberClassID"`
}

type User struct {
	gorm.Model
	Pin       string `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	Civ       string `gorm:"uniqueIndex"`
	Phone     string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Address   string
	//FK
	RoleID        *uint
	ProvinceID    *uint
	MemberClassID *uint
	//JOIN
	Province    Province    `gorm:"references:id"`
	Role        Role        `gorm:"references:id"`
	MemberClass MemberClass `gorm:"references:id"`
}

//******************************************************************

type Problem struct {
	gorm.Model
	Name string
	// 1 Problem อยู่ได้ในหลาย Relation
	/* Relations []Relation `gorm:"foreignKey:Problem_ID"` */
	// 1 Problem อยู่ได้ในหลาย Report
	ProblemReports []ProblemReport `gorm:"foreignKey:Problem_ID"`
}

// ======================================
type Place_Class struct {
	gorm.Model
	Name string

	/* Relations     []Relation     `gorm:"foreignKey:Place_Class_ID"` */
	Toilets       []Toilet       `gorm:"foreignKey:Place_Class_ID"`
	ReadingZones  []ReadingZone  `gorm:"foreignKey:Place_Class_ID"`
	ResearchRooms []ResearchRoom `gorm:"foreignKey:Place_Class_ID"`
	Computers     []Computer     `gorm:"foreignKey:Place_Class_ID"`
}

// ======================================
/* type Relation struct {
	gorm.Model
	Place_Class_ID *uint
	Problem_ID     *uint
	//JOIN
	Place_Class Place_Class `gorm:"references:id"`
	Problem     Problem     `gorm:"references:id"`
} */

// ======================================
type Toilet struct {
	gorm.Model
	Name string
	// Place_Problem_ID ทำหน้าที่เป็น FK
	Place_Class_ID *uint
	Place_Class    Place_Class     `gorm:"references:id"`
	ProblemReports []ProblemReport `gorm:"foreignKey:Problem_ID"`
}

// ======================================
type ReadingZone struct {
	gorm.Model
	Name string
	// Place_Problem_ID ทำหน้าที่เป็น FK
	Place_Class_ID *uint
	Place_Class    Place_Class     `gorm:"references:id"`
	ProblemReports []ProblemReport `gorm:"foreignKey:Problem_ID"`
}

// ======================================
type RoomType struct {
	gorm.Model
	Type string
	// ResearchRoom []ResearchRoom `gorm:"foreignKey:RoomTypeID"`
}

// type Equipment struct {
// 	gorm.Model
// 	Name string
// 	// ResearchRoom []ResearchRoom `gorm:"foreignKey:EquipmentID"`
// }

type ResearchRoom struct {
	gorm.Model
	Name string

	RoomTypeID *uint    //FK
	RoomType   RoomType `gorm:"references:id"` //JOIN //ทำการตึง id ของ RoomType

	// EquipmentID *uint     //FK
	// Equipment   Equipment `gorm:"references:id"` //JOIN

	Place_Class_ID *uint
	Place_Class    Place_Class
	ProblemReports []ProblemReport `gorm:"foreignKey:Problem_ID"`

	// Place_ProblemID *uint         //FK
	// Place_Problem   Place_Problem //JOIN
	// RRRR []ResearchRoomReservationRecord `gorm:"foreignKey:ResearchRoomID"`
}

type TimeRoom struct {
	gorm.Model
	Period string
	// RRRR   []ResearchRoomReservationRecord `gorm:"foreignKey:TimeID"`
}

type AddOn struct {
	gorm.Model
	Name string
	// RRRR []ResearchRoomReservationRecord `gorm:"foreignKey:AddOnID"`
}

type ResearchRoomReservationRecord struct {
	gorm.Model
	BookDate time.Time

	UserID *uint //FK
	User   User  //JOIN

	ResearchRoomID *uint        //FK
	ResearchRoom   ResearchRoom `gorm:"references:id"`

	AddOnID *uint //FK
	AddOn   AddOn `gorm:"references:id"`

	TimeRoomID *uint    //FK
	TimeRoom   TimeRoom `gorm:"references:id"`
}

// ======================================
type Computer_reservation struct {
	gorm.Model
	Date time.Time

	Computer_id *uint    //FK
	Computer    Computer `gorm:"references:id"` //JOIN

	// COMPUTER_OS_ID *uint       //FK
	// COMPUTER_OS    COMPUTER_OS `gorm:"references:id"` //JOIN

	Time_com_id *uint    //FK
	Time_com    Time_com `gorm:"references:id"` //JOIN

	UserID *uint //FK
	User   User  `gorm:"references:id"` //JOIN
}

type Computer struct {
	gorm.Model
	Comp_name string
	Comp_room string

	Computer_os_id *uint       //FK
	Computer_os    Computer_os `gorm:"references:id"` //JOIN
	Place_Class_ID *uint
	Place_Class    Place_Class     `gorm:"references:id"`
	ProblemReports []ProblemReport `gorm:"foreignKey:Problem_ID"`

	// COMPUTER_RESERVATION []COMPUTER_RESERVATION `gorm:"foreignKey:COMPUTER_ID"`
}

type Computer_os struct {
	gorm.Model
	Name string
	// COMPUTER []COMPUTER `gorm:"foreignKey:COMPUTER_OS_ID"`
}

type Time_com struct {
	gorm.Model
	Time_com_period      string
	Computer_reservation []Computer_reservation `gorm:"foreignKey:Time_com_id"`
}

// ======================================
type ProblemReport struct {
	gorm.Model

	USER_ID *uint
	User    User `gorm:"references:id"`

	Place_Class_ID *uint
	Place_Class    Place_Class `gorm:"references:id"`

	RdZone_id *uint
	RdZone    ReadingZone `gorm:"references:id"`

	Tlt_id *uint
	Tlt    Toilet

	ReschRoom_id *uint
	ReschRoom    ResearchRoom `gorm:"references:id"`

	Com_id *uint
	Com    Computer

	Problem_ID *uint
	Problem    Problem `gorm:"references:id"`

	Comment string

	Date time.Time
}

//======================================
