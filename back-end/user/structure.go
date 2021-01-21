package user

//PlayerCore data
type PlayerCore struct {
	UID        int    `json:"uid"`
	Username   string `json:"username"`
	Registered string `json:"register_date"`
	LastLogin  string `json:"last_login"`
}

//PlayerStats data
type PlayerStats struct {
	Kills  int    `json:"kills"`
	Deaths int    `json:"deaths"`
	Money  int    `json:"money"`
	Job    byte   `json:"job"`
	Class  byte   `json:"class"`
	Score  int    `json:"score"`
	Skin   uint16 `json:"skin"`
}

//PlayerItems data
type PlayerItems struct {
	Crack    byte `json:"crack"`
	Weed     byte `json:"weed"`
	Picklock byte `json:"piclock"`
	Wallet   byte `json:"wallet"`
	Rope     byte `json:"rope"`
	Condom   byte `json:"condom"`
	Scissors byte `json:"sciccors"`
}

//Player data
type Player struct {
	Account PlayerCore  `json:"account"`
	Stats   PlayerStats `json:"stats"`
	Items   PlayerItems `json:"items"`
}
