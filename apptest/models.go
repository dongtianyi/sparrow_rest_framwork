package apptest

type TestModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (t *TestModel) Get(id int) {
	t.Id = id
	t.Name = "我是数据"
}
