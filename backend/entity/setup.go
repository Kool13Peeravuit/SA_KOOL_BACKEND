package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema

	database.AutoMigrate(
		&User{},
		&Problem{},
		&Place_Class{},
		/* &Relation{}, */
		&Toilet{},
		&ReadingZone{},
		&ResearchRoom{},
		&Computer{},
		&ProblemReport{},
		&AddOn{},
		&Time_com{},
		&TimeRoom{},
		&Computer_reservation{},
		&ResearchRoomReservationRecord{},
	)

	//=========================================================================== JOON

	db = database

	password1, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("3333333333333"), 14)

	//add example data

	//Role
	student := Role{
		Name:       "Student",
		BorrowDay:  3,
		BookRoomHR: 3,
		BookComHR:  4,
	}

	db.Model(&Role{}).Create(&student)

	teacher := Role{
		Name:       "Teacher",
		BorrowDay:  7,
		BookRoomHR: 12,
		BookComHR:  12,
	}
	db.Model(&Role{}).Create(&teacher)

	employee := Role{
		Name:       "Employee",
		BorrowDay:  5,
		BookRoomHR: 6,
		BookComHR:  6,
	}
	db.Model(&Role{}).Create(&employee)

	//province
	korat := Province{
		Name: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		Name: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		Name: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	//member
	classic := MemberClass{
		Name:     "classic",
		Discount: 0,
	}
	db.Model(&MemberClass{}).Create(&classic)

	silver := MemberClass{
		Name:     "silver",
		Discount: 5,
	}
	db.Model(&MemberClass{}).Create(&silver)

	gold := MemberClass{
		Name:     "gold",
		Discount: 10,
	}
	db.Model(&MemberClass{}).Create(&gold)

	plat := MemberClass{
		Name:     "platinum",
		Discount: 20,
	}
	db.Model(&MemberClass{}).Create(&plat)

	//user
	db.Model(&User{}).Create(&User{
		Pin:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		Civ:       "1111111111111",
		Phone:     "0811111111",
		Email:     "preechapat@mail.com",
		Password:  string(password1),
		Address:   "ถนน a อำเภอ v",
		//FK
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})

	db.Model(&User{}).Create(&User{
		Pin:       "E123456",
		FirstName: "kawin",
		LastName:  "l.",
		Civ:       "1234567890123",
		Phone:     "0899999999",
		Email:     "kawin@mail.com",
		Password:  string(password2),
		Address:   "หอ b อำเภอ r",
		//FK
		Role:        employee,
		Province:    chon,
		MemberClass: silver,
	})

	db.Model(&User{}).Create(&User{
		Pin:       "T8888",
		FirstName: "sirinya",
		LastName:  "kotpanya",
		Civ:       "3333333333333",
		Phone:     "0823456789",
		Email:     "sirinya@mail.com",
		Password:  string(password3),
		Address:   "บ้าน c อำเภอ q",
		//FK
		Role:        teacher,
		Province:    bangkok,
		MemberClass: plat,
	})

	var preecha User
	db.Raw("SELECT * FROM users WHERE email = ?", "preechapat@mail.com").Scan(&preecha)

	var kawin User
	db.Raw("SELECT * FROM users WHERE email = ?", "kawin@mail.com").Scan(&kawin)

	//=========================================================================== KOOL PART PLACE CLASS
	pcRdZone := Place_Class{
		Name: "Reading Zone",
	}
	db.Model(&Place_Class{}).Create(&pcRdZone)

	pcTlt := Place_Class{
		Name: "Toilet",
	}
	db.Model(&Place_Class{}).Create(&pcTlt)

	pcReschRoom := Place_Class{
		Name: "Research Room",
	}
	db.Model(&Place_Class{}).Create(&pcReschRoom)

	pcCom := Place_Class{
		Name: "Computer",
	}
	db.Model(&Place_Class{}).Create(&pcCom)

	//=========================================================================== MINK

	single_room := RoomType{
		Type: "ห้องเดี่ยว",
	}
	db.Model(&RoomType{}).Create(&single_room)

	group_room := RoomType{
		Type: "ห้องกลุ่ม",
	}
	db.Model(&RoomType{}).Create(&group_room)

	tutor_room := RoomType{
		Type: "ห้องสำหรับติว",
	}
	db.Model(&RoomType{}).Create(&tutor_room)

	//Research room
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:     "SR01",
		RoomType: single_room,
		//Equipment:   monitor,
		Place_Class: pcReschRoom,
	})

	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:     "GR02",
		RoomType: group_room,
		// Equipment:   printerMoniter,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:     "TR03",
		RoomType: tutor_room,
		// Equipment:   printerMoniter,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:     "SR04",
		RoomType: single_room,
		// Equipment:   printerMoniter,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:     "GR05",
		RoomType: group_room,
		// Equipment:   printerMoniter,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:     "TR06",
		RoomType: tutor_room,
		// Equipment:   printerMoniter,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:     "SR07",
		RoomType: single_room,
		// Equipment:   printerMoniter,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:     "GR08",
		RoomType: group_room,
		// Equipment:   printerMoniter,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:     "TR09",
		RoomType: tutor_room,
		// Equipment:   printerMoniter,
		Place_Class: pcReschRoom,
	})
	var Room1 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE id = ?", 1).Scan(&Room1)
	var Room2 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE id = ?", 2).Scan(&Room2)

	//Addon data
	powerPlug := AddOn{
		Name: "ปลั๊กพ่วง",
	}
	db.Model(&AddOn{}).Create(&powerPlug)

	Adapter := AddOn{
		Name: "สายชาร์จ",
	}
	db.Model(&AddOn{}).Create(&Adapter)

	Pillow := AddOn{
		Name: "หมอน",
	}
	db.Model(&AddOn{}).Create(&Pillow)

	powerPlugAndAdapter := AddOn{
		Name: "ปลั๊กพ่วง + สายชาร์จ",
	}
	db.Model(&AddOn{}).Create(&powerPlugAndAdapter)

	adapterAndPillow := AddOn{
		Name: "สายชาร์จ + หมอน",
	}
	db.Model(&AddOn{}).Create(&adapterAndPillow)

	powerPlugAndAdapterAndPillow := AddOn{
		Name: "ปลั๊กพ่วง + สายชาร์จ + หมอน",
	}
	db.Model(&AddOn{}).Create(&powerPlugAndAdapterAndPillow)

	//Time data
	timeMorning := TimeRoom{
		Period: "08:00 - 12:00",
	}
	db.Model(&TimeRoom{}).Create(&timeMorning)

	timeAfternoon := TimeRoom{
		Period: "13:00 - 17:00",
	}
	db.Model(&TimeRoom{}).Create(&timeAfternoon)

	//RRRR 1
	db.Model(&ResearchRoomReservationRecord{}).Create(&ResearchRoomReservationRecord{
		ResearchRoom: Room1,
		User:         preecha,
		AddOn:        powerPlugAndAdapter,
		BookDate:     time.Now(),
		TimeRoom:     timeMorning,
	})

	//RRRR 2
	db.Model(&ResearchRoomReservationRecord{}).Create(&ResearchRoomReservationRecord{
		ResearchRoom: Room2,
		User:         kawin,
		AddOn:        powerPlugAndAdapterAndPillow,
		BookDate:     time.Now(),
		TimeRoom:     timeAfternoon,
	})

	//=========================================================================== GUITAR
	//Computer_os data
	comp_os_name1 := Computer_os{
		Name: "Windows",
	}
	db.Model(&Computer_os{}).Create(&comp_os_name1)

	//Computer data
	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W01A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
		Place_Class: pcCom,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W02A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
		Place_Class: pcCom,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W03A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
		Place_Class: pcCom,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W04A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
		Place_Class: pcCom,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W05A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
		Place_Class: pcCom,
	})

	//TIME data
	time_p1 := Time_com{
		Time_com_period: "08:00 - 09:00",
	}
	db.Model(&Time_com{}).Create(&time_p1)

	time_p2 := Time_com{
		Time_com_period: "09:00 - 10:00",
	}
	db.Model(&Time_com{}).Create(&time_p2)

	time_p3 := Time_com{
		Time_com_period: "10:00 - 11:00",
	}
	db.Model(&Time_com{}).Create(&time_p3)

	time_p4 := Time_com{
		Time_com_period: "11:00 - 12:00",
	}
	db.Model(&Time_com{}).Create(&time_p4)

	time_p5 := Time_com{
		Time_com_period: "12:00 - 13:00",
	}
	db.Model(&Time_com{}).Create(&time_p5)

	time_p6 := Time_com{
		Time_com_period: "13:00 - 14:00",
	}
	db.Model(&Time_com{}).Create(&time_p6)

	time_p7 := Time_com{
		Time_com_period: "14:00 - 15:00",
	}
	db.Model(&Time_com{}).Create(&time_p7)

	time_p8 := Time_com{
		Time_com_period: "15:00 - 16:00",
	}
	db.Model(&Time_com{}).Create(&time_p8)

	//ดึง Data ของ COMPUTER มาเก็บไว้ในตัวแปรก่อน
	// cn = comp_name ที่มาจาก COMP_NAME ใน Entity COMPUTER
	var cn1 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "W01A").Scan(&cn1)
	var cn2 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "W02A").Scan(&cn2)
	var cn3 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "W03A").Scan(&cn3)
	var cn4 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "M01B").Scan(&cn4)
	var cn5 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "M02B").Scan(&cn5)

	//COMPUTER_RESERVATION
	db.Model(&Computer_reservation{}).Create(&Computer_reservation{

		Date:     time.Now(),
		Computer: cn1,
		Time_com: time_p1,
		User:     preecha,
	})
	db.Model(&Computer_reservation{}).Create(&Computer_reservation{

		Date:     time.Now(),
		Computer: cn2,
		Time_com: time_p2,
		User:     kawin,
	})
	//=========================================================================== KOOL PART PLACE
	RdZone1 := ReadingZone{
		Name:        "Reading Zone 1",
		Place_Class: pcRdZone,
	}
	db.Model(&ReadingZone{}).Create(&RdZone1)
	RdZone2 := ReadingZone{
		Name:        "Reading Zone 2",
		Place_Class: pcRdZone,
	}
	db.Model(&ReadingZone{}).Create(&RdZone2)
	RdZone3 := ReadingZone{
		Name:        "Reading Zone 3",
		Place_Class: pcRdZone,
	}
	db.Model(&ReadingZone{}).Create(&RdZone3)

	Tlt1 := Toilet{
		Name:        "Toilet 1",
		Place_Class: pcTlt,
	}
	db.Model(&Toilet{}).Create(&Tlt1)
	Tlt2 := Toilet{
		Name:        "Toilet 2",
		Place_Class: pcTlt,
	}
	db.Model(&Toilet{}).Create(&Tlt2)
	Tlt3 := Toilet{
		Name:        "Toilet 3",
		Place_Class: pcTlt,
	}
	db.Model(&Toilet{}).Create(&Tlt3)
	Tlt4 := Toilet{
		Name:        "Toilet 4",
		Place_Class: pcTlt,
	}
	db.Model(&Toilet{}).Create(&Tlt4)
	//=========================================================================== KOOL PART PROBLEM

	probLightBulb := Problem{
		Name: "Light Bulb (หลอดไฟขัดข้อง)",
	}
	db.Model(&Problem{}).Create(&probLightBulb)
	probAirCon := Problem{
		Name: "Air Condition (แอร์)",
	}
	db.Model(&Problem{}).Create(&probAirCon)
	probinTlt := Problem{
		Name: "in Toilet (ปัญหาในห้องน้ำ)",
	}
	db.Model(Problem{}).Create(&probinTlt)
	probRRdevice := Problem{
		Name: "Research Room Device (อุปกรณ์ห้องค้นคว้า)",
	}
	db.Model(&Problem{}).Create(&probRRdevice)
	probReschCom := Problem{
		Name: " Research Computer (เกี่ยวกับคอมค้นคว้า)",
	}
	db.Model(&Problem{}).Create(&probReschCom)
	probBookshelf := Problem{
		Name: "Bookshelf (เกี่ยวกับชั้นวางหนังสือ)",
	}
	db.Model(&Problem{}).Create(&probBookshelf)
	probDeskChair := Problem{
		Name: "Desk & Chair (เก้าอี้และโต๊ะ)",
	}
	db.Model(&Problem{}).Create(&probDeskChair)

	//=========================================================================== KOOL PART ProbReport
	db.Model(&ProblemReport{}).Create(&ProblemReport{
		User:           preecha,
		RdZone_id:      nil,
		Tlt_id:         &Tlt2.ID,
		ReschRoom_id:   nil,
		Com_id:         nil,
		Problem_ID:     &probinTlt.ID,
		Place_Class_ID: &pcTlt.ID,
		Comment:        "ห้องน้ำกดส้วมไม่ลง มุแง TT",
		Date:           time.Now(),
	})
	db.Model(&ProblemReport{}).Create(&ProblemReport{
		User:           kawin,
		RdZone_id:      &RdZone3.ID,
		Tlt_id:         nil,
		ReschRoom_id:   nil,
		Com_id:         nil,
		Problem_ID:     &probDeskChair.ID,
		Place_Class_ID: &pcRdZone.ID,
		Comment:        "ล้อเก้าอี้เลื่อนหาย",
		Date:           time.Now(),
	})
}
