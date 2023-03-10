package apptest

type PdtPdt struct {
	ID int `json:"id" gorm:"primaryKey"`
	// CreatedAt ccldatatypes.CCLTime  `json:"created_at" gorm:"autoCreateTime"`
	// IsDel     soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag"`
	PdtclsID int    `json:"pdtcls_id" validate:"required"`
	ShopID   int    `json:"shop_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	SubTitle string `json:"sub_title"`
	Desc     string `json:"desc"`
	ImgID    int    `json:"img_id"`
}

func (t *PdtPdt) Get(id int) {
}
