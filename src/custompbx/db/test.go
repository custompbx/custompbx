package db

type TestParentTable struct {
	Id   int64  `xml:"-" json:"id" customsql:"pkey:id"`
	Name string `xml:"name,attr" json:"name" customsql:"column_name;unique"`
}
type TestTable struct {
	Id      int64            `xml:"-" json:"id" customsql:"pkey:id"`
	Name    string           `xml:"name,attr" json:"name" customsql:"column_name;unique"`
	Status  string           `xml:"status,attr" json:"status" customsql:"room_state"`
	Enabled bool             `xml:"-" json:"enabled" customsql:"enabled"`
	Parent  *TestParentTable `xml:"-" json:"-" customsql:"fkey:test_id;unique"`
}

func (t *TestTable) GetTableName() string {
	return "test_new_name_of_table"
}
func TT() {
	//corm := customOrm.Init(db)
	//corm.CreateTable(TestParentTable{})
	//corm.CreateTable(&TestTable{})
	//log.Println(corm.InsertRow(&TestParentTable{Name: "hi"}))
	//log.Println(corm.InsertRow(&TestTable{Name: "1", Parent: &TestParentTable{Id: 1}}))
	//log.Println(corm.InsertRow(&TestTable{Name: "2", Parent: &TestParentTable{Id: 1}}))
	//log.Println(corm.InsertRow(&TestTable{Name: "3", Parent: &TestParentTable{Id: 1}}))
	//log.Println(corm.InsertRow(&TestTable{Name: "4", Parent: &TestParentTable{Id: 1}}))
	//log.Println(corm.InsertRow(&TestTable{Name: "hello", Parent: &TestParentTable{Id: 1}}))
	//log.Println(corm.InsertRow(&TestTable{Name: "bye", Enabled: true, Parent: &TestParentTable{Id: 1}}))
	//log.Println(corm.InsertRow(&TestTable{Name: "bye2", Enabled: true, Parent: &TestParentTable{Id: 1}}))
	//log.Println(corm.DeleteRowById(&TestTable{Id: 12}))
	//log.Println(corm.UpdateRow(&TestTable{Id: 10, Status: "eeeeeee"}, true))
	//tt , _ := corm.GetDataById(&TestTable{Id: 2})
	//tt , _ := corm.GetDataAll(&TestTable{}, false)
	//log.Println(tt)
}

type Testo interface {
	GetID() int64
}

func t() {
	var ttt Testo
	ttt.GetID()
}
