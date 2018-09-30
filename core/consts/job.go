package consts

var (
	MapleJobs mapleJobs
)

func init() {
	MapleJobs.Beginer = 0
	MapleJobs.Warrior = 100
	MapleJobs.Fighter = 110
	MapleJobs.Crusader = 111
	MapleJobs.Hero = 112
	MapleJobs.Page = 120
	MapleJobs.WhiteKnight = 121
	MapleJobs.Paladin = 122
	MapleJobs.Spearman = 130
	MapleJobs.DragonKnight = 131
	MapleJobs.DarkKnight = 132
	MapleJobs.Magician = 200
	MapleJobs.FpWizard = 210
	MapleJobs.FpMage = 211
	MapleJobs.FpArchmage = 212
	MapleJobs.IlWizard = 220
	MapleJobs.IlMage = 221
	MapleJobs.IlArchmage = 222
	MapleJobs.Cleric = 230
	MapleJobs.Priest = 231
	MapleJobs.Bishop = 232
	MapleJobs.Bowman = 300
	MapleJobs.Hunter = 310
	MapleJobs.Ranger = 311
	MapleJobs.Bowman = 312
	MapleJobs.Crossbowman = 320
	MapleJobs.Sniper = 321
	MapleJobs.Crossbowman = 322
	MapleJobs.Thief = 400
	MapleJobs.Assassin = 410
	MapleJobs.Hermit = 411
	MapleJobs.Nightlord = 412
	MapleJobs.Bandit = 420
	MapleJobs.Chefbandit = 421
	MapleJobs.Shadower = 422
	MapleJobs.Priate = 500
	MapleJobs.Brawler = 510
	MapleJobs.Marauder = 511
	MapleJobs.Buccaneer = 512
	MapleJobs.Gunslinger = 520
	MapleJobs.Outlaw = 512
	MapleJobs.Corsair = 522
	MapleJobs.GM = 900
	MapleJobs.Knight = 1000
	MapleJobs.GhostKnight = 1100
	MapleJobs.GhostKnight2 = 1110
	MapleJobs.GhostKnight3 = 1111
	MapleJobs.FireKnight = 1200
	MapleJobs.FireKnight2 = 1210
	MapleJobs.FireKnight3 = 1211
	MapleJobs.WindKnight = 1300
	MapleJobs.WindKnight2 = 1310
	MapleJobs.WindKnight3 = 1311
	MapleJobs.NightKnight = 1400
	MapleJobs.NightKnight2 = 1410
	MapleJobs.NightKnight3 = 1411
	MapleJobs.ThiefKnight = 1500
	MapleJobs.ThiefKnight2 = 1510
	MapleJobs.ThiefKnight3 = 1511
	MapleJobs.Ares = 2000
	MapleJobs.Ares1 = 2100
	MapleJobs.Ares2 = 2110
	MapleJobs.Ares3 = 2111
	MapleJobs.Ares4 = 2112
}

type mapleJobs struct {
	Beginer        int
	Warrior        int
	Fighter        int
	Crusader       int
	Hero           int
	Page           int
	WhiteKnight    int
	Paladin        int
	Spearman       int
	DragonKnight   int
	DarkKnight     int
	Magician       int
	FpWizard       int
	FpMage         int
	FpArchmage     int
	IlWizard       int
	IlMage         int
	IlArchmage     int
	Cleric         int
	Priest         int
	Bishop         int
	Bowman         int
	Hunter         int
	Ranger         int
	Bowmaster      int
	Crossbowman    int
	Sniper         int
	Crossbowmaster int
	Thief          int
	Assassin       int
	Hermit         int
	Nightlord      int
	Bandit         int
	Chefbandit     int
	Shadower       int
	Priate         int
	Brawler        int
	Marauder       int
	Buccaneer      int
	Gunslinger     int
	Outlaw         int
	Corsair        int
	GM             int
	Knight         int
	GhostKnight    int
	GhostKnight2   int
	GhostKnight3   int
	FireKnight     int
	FireKnight2    int
	FireKnight3    int
	WindKnight     int
	WindKnight2    int
	WindKnight3    int
	NightKnight    int
	NightKnight2   int
	NightKnight3   int
	ThiefKnight    int
	ThiefKnight2   int
	ThiefKnight3   int
	Ares           int
	Ares1          int
	Ares2          int
	Ares3          int
	Ares4          int
}

func MapleJobEqual(l, r int) bool {
	return false
}
