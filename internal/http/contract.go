package http

type Taco struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	BaseQty     float64 `json:"base_qty"`
	BaseUnit    string  `json:"base_unit"`
	CategoryID  int     `json:"category_id"`
	Attributes  struct {
		Humidity     TacoAttributes `json:"humidity"`
		Protein      TacoAttributes `json:"protein"`
		Lipid        TacoAttributes `json:"lipid"`
		Cholesterol  TacoAttributes `json:"cholesterol"`
		Carbohydrate TacoAttributes `json:"carbohydrate"`
		Fiber        TacoAttributes `json:"fiber"`
		Ashes        TacoAttributes `json:"ashes"`
		Calcium      TacoAttributes `json:"calcium"`
		Magnesium    TacoAttributes `json:"magnesium"`
		Phosphorus   TacoAttributes `json:"phosphorus"`
		Iron         TacoAttributes `json:"iron"`
		Sodium       TacoAttributes `json:"sodium"`
		Potassium    TacoAttributes `json:"potassium"`
		Copper       TacoAttributes `json:"copper"`
		Zinc         TacoAttributes `json:"zinc"`
		Retinol      TacoAttributes `json:"retinol"`
		Thiamine     TacoAttributes `json:"thiamine"`
		Riboflavin   TacoAttributes `json:"riboflavin"`
		Pyridoxine   TacoAttributes `json:"pyridoxine"`
		Niacin       TacoAttributes `json:"niacin"`
		Energy       struct {
			Kcal float64 `json:"kcal"`
			Kj   float64 `json:"kj"`
		} `json:"energy"`
		FattyAcids struct {
			Saturated       TacoAttributes `json:"saturated"`
			Monounsaturated TacoAttributes `json:"monounsaturated"`
			Polyunsaturated TacoAttributes `json:"polyunsaturated"`
		} `json:"fatty_acids"`
		Manganese TacoAttributes `json:"manganese"`
	}
}

type TacoAttributes struct {
	Qty  float64 `json:"qty"`
	Unit string  `json:"unit"`
}
