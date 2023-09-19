package models

type Measurement struct {
	OrderNo   int    `json:"orderNo" gorm:"primaryKey;autoIncrement:trueq"`
	Name      string `json:"name"`
	Contact   string `json:"contact"`
	OldNo     string `json:"oldNo"`
	WLength   string `json:"wLength"`
	WChest    string `json:"wChest"`
	WWaist    string `json:"wWaist"`
	WSleeves  string `json:"wSleeves"`
	WCBack    string `json:"wCBack"`
	WHBack    string `json:"wHBack"`
	WNeck     string `json:"wNeck"`
	PLength   string `json:"pLength"`
	PWaist    string `json:"pWaist"`
	PHip      string `json:"pHip"`
	PInseam   string `json:"pInseam"`
	PThigh    string `json:"pThigh"`
	PKnee     string `json:"pKnee"`
	PBottom   string `json:"pBottom"`
	SLength   string `json:"sLength"`
	SChest    string `json:"sChest"`
	SWaist    string `json:"sWaist"`
	SShoulder string `json:"sShoulder"`
	SSleeves  string `json:"sSleeves"`
	SNeck     string `json:"sNeck"`
	QLength   string `json:"qLength"`
	QChest    string `json:"qChest"`
	QWaist    string `json:"qWaist"`
	QShoulder string `json:"qShoulder"`
	QSleeves  string `json:"qSleeves"`
	QGhera    string `json:"qGhera"`
	QNeck     string `json:"qNeck"`
	QSLength  string `json:"qSLength"`
	QBottom   string `json:"qBottom"`
	Notes     string `json:"notes"`
	Total     string `json:"total"`
	Advance   string `json:"advance"`
	Balance   string `json:"balance"`
	Booking   string `json:"booking"`
	Delivery  string `json:"delivery"`
}
