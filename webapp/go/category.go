package main

var (
	staticCategories = []Category{
		Category{ID: 1, ParentID: 0, CategoryName: "ソファー"},
		Category{ID: 2, ParentID: 1, CategoryName: "一人掛けソファー", ParentCategoryName: "ソファー"},
		Category{ID: 3, ParentID: 1, CategoryName: "二人掛けソファー", ParentCategoryName: "ソファー"},
		Category{ID: 4, ParentID: 1, CategoryName: "コーナーソファー", ParentCategoryName: "ソファー"},
		Category{ID: 5, ParentID: 1, CategoryName: "二段ソファー", ParentCategoryName: "ソファー"},
		Category{ID: 6, ParentID: 1, CategoryName: "ソファーベッド", ParentCategoryName: "ソファー"},
		Category{ID: 10, ParentID: 0, CategoryName: "家庭用チェア"},
		Category{ID: 11, ParentID: 10, CategoryName: "スツール", ParentCategoryName: "家庭用チェア"},
		Category{ID: 12, ParentID: 10, CategoryName: "クッションスツール", ParentCategoryName: "家庭用チェア"},
		Category{ID: 13, ParentID: 10, CategoryName: "ダイニングチェア", ParentCategoryName: "家庭用チェア"},
		Category{ID: 14, ParentID: 10, CategoryName: "リビングチェア", ParentCategoryName: "家庭用チェア"},
		Category{ID: 15, ParentID: 10, CategoryName: "カウンターチェア", ParentCategoryName: "家庭用チェア"},
		Category{ID: 20, ParentID: 0, CategoryName: "キッズチェア"},
		Category{ID: 21, ParentID: 20, CategoryName: "学習チェア", ParentCategoryName: "キッズチェア"},
		Category{ID: 22, ParentID: 20, CategoryName: "ベビーソファ", ParentCategoryName: "キッズチェア"},
		Category{ID: 23, ParentID: 20, CategoryName: "キッズハイチェア", ParentCategoryName: "キッズチェア"},
		Category{ID: 24, ParentID: 20, CategoryName: "テーブルチェア", ParentCategoryName: "キッズチェア"},
		Category{ID: 30, ParentID: 0, CategoryName: "オフィスチェア"},
		Category{ID: 31, ParentID: 30, CategoryName: "デスクチェア", ParentCategoryName: "オフィスチェア"},
		Category{ID: 32, ParentID: 30, CategoryName: "ビジネスチェア", ParentCategoryName: "オフィスチェア"},
		Category{ID: 33, ParentID: 30, CategoryName: "回転チェア", ParentCategoryName: "オフィスチェア"},
		Category{ID: 34, ParentID: 30, CategoryName: "リクライニングチェア", ParentCategoryName: "オフィスチェア"},
		Category{ID: 35, ParentID: 30, CategoryName: "投擲用椅子", ParentCategoryName: "オフィスチェア"},
		Category{ID: 40, ParentID: 0, CategoryName: "折りたたみ椅子"},
		Category{ID: 41, ParentID: 40, CategoryName: "パイプ椅子", ParentCategoryName: "折りたたみ椅子"},
		Category{ID: 42, ParentID: 40, CategoryName: "木製折りたたみ椅子", ParentCategoryName: "折りたたみ椅子"},
		Category{ID: 43, ParentID: 40, CategoryName: "キッチンチェア", ParentCategoryName: "折りたたみ椅子"},
		Category{ID: 44, ParentID: 40, CategoryName: "アウトドアチェア", ParentCategoryName: "折りたたみ椅子"},
		Category{ID: 45, ParentID: 40, CategoryName: "作業椅子", ParentCategoryName: "折りたたみ椅子"},
		Category{ID: 50, ParentID: 0, CategoryName: "ベンチ"},
		Category{ID: 51, ParentID: 50, CategoryName: "一人掛けベンチ", ParentCategoryName: "ベンチ"},
		Category{ID: 52, ParentID: 50, CategoryName: "二人掛けベンチ", ParentCategoryName: "ベンチ"},
		Category{ID: 53, ParentID: 50, CategoryName: "アウトドア用ベンチ", ParentCategoryName: "ベンチ"},
		Category{ID: 54, ParentID: 50, CategoryName: "収納付きベンチ", ParentCategoryName: "ベンチ"},
		Category{ID: 55, ParentID: 50, CategoryName: "背もたれ付きベンチ", ParentCategoryName: "ベンチ"},
		Category{ID: 56, ParentID: 50, CategoryName: "ベンチマーク", ParentCategoryName: "ベンチ"},
		Category{ID: 60, ParentID: 0, CategoryName: "座椅子"},
		Category{ID: 61, ParentID: 60, CategoryName: "和風座椅子", ParentCategoryName: "座椅子"},
		Category{ID: 62, ParentID: 60, CategoryName: "高座椅子", ParentCategoryName: "座椅子"},
		Category{ID: 63, ParentID: 60, CategoryName: "ゲーミング座椅子", ParentCategoryName: "座椅子"},
		Category{ID: 64, ParentID: 60, CategoryName: "ロッキングチェア", ParentCategoryName: "座椅子"},
		Category{ID: 65, ParentID: 60, CategoryName: "座布団", ParentCategoryName: "座椅子"},
		Category{ID: 66, ParentID: 60, CategoryName: "空気椅子", ParentCategoryName: "座椅子"},
	}
)

func getCategoryById(id int) (Category, bool) {
	for _, c := range staticCategories {
		if c.ID == id {
			return c, true
		}
	}
	return Category{}, false
}

func getCategoryIdsByParentId(parentId int) []int {
	var indices []int
	for _, c := range staticCategories {
		if c.ParentID == parentId {
			indices = append(indices, c.ID)
		}
	}
	return indices
}
