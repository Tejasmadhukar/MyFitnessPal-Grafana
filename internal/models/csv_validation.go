package models

type NutritionInfo struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type TotalIntake struct {
	Total []NutritionInfo `json:"total"`
	Goal  []NutritionInfo `json:"goal"`
}
