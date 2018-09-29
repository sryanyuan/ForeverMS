package opcode

var (
	SendOps sendOpcode
)

type sendOpcode struct {
	// GENERAL
	PING int16 // 0x11
	// LOGIN
	LOGIN_STATUS       int16 // 1
	CHOOSE_GENDER      int16
	LICENSE_RESULT     int16
	GENDER_SET         int16
	PIN_OPERATION      int16 `properties:"PIN_OPERATION,default=65535"`
	PIN_ASSIGNED       int16 `properties:"PIN_ASSIGNED,default=65535"`
	SERVERLIST         int16 // 0xa
	SERVERSTATUS       int16 // 3
	SERVER_IP          int16 // 0xc
	CHARLIST           int16 // 0xb
	CHAR_NAME_RESPONSE int16 // 0xd
	RELOG_RESPONSE     int16 `properties:"RELOG_RESPONSE,default=65535"` // 0x16
	ADD_NEW_CHAR_ENTRY int16 // 0xe
	CHANNEL_SELECTED   int16
	ALL_CHARLIST       int16 `properties:"ALL_CHARLIST,default=65535"`
	// CHANNEL
	CHANGE_CHANNEL               int16 // 0x10
	UPDATE_STATS                 int16 // 0x1b
	FAME_RESPONSE                int16
	UPCHRLOOK                    int16 `properties:"UPCHRLOOK,default=65535"`
	ENABLE_TEMPORARY_STATS       int16
	DISABLE_TEMPORARY_STATS      int16
	UPDATE_SKILLS                int16 // 0x1e
	CHAR_CASH                    int16
	WARP_TO_MAP                  int16 // 0x49
	SERVERMESSAGE                int16 // 0x37
	FAMILY_ACTION                int16
	OPEN_FAMILY                  int16
	FAMILY_MESSAGE               int16
	FAMILY_INVITE                int16
	FAMILY_MESSAGE2              int16
	FAMILY_SENIOR_MESSAGE        int16
	FAMILY_GAIN_REP              int16
	LOAD_FAMILY                  int16
	FAMILY_USE_REQUEST           int16
	AVATAR_MEGA                  int16 // 0x42
	SPAWN_NPC                    int16 // 0xb1
	SPAWN_NPC_REQUEST_CONTROLLER int16 // 0xb3
	REMOVE_NPC                   int16
	SPAWN_MONSTER                int16 // 0x9E
	SPAWN_MONSTER_CONTROL        int16 // 0xA0
	MOVE_MONSTER_RESPONSE        int16 // 0xA3

	CHATTEXT                     int16 // 0x67
	SHOW_STATUS_INFO             int16 // 0x21
	SHOW_QUEST_COMPLETION        int16 // 0x29
	WHISPER                      int16
	SPAWN_PLAYER                 int16 // 0x64
	SHOW_SCROLL_EFFECT           int16 // 0x6B
	SHOW_ITEM_GAIN_INCHAT        int16 // 0x92
	DOJO_WARP_UP                 int16
	ENERGY                       int16
	KILL_MONSTER                 int16 // 0x9f
	DROP_ITEM_FROM_MAPOBJECT     int16 // 0xC1
	FACIAL_EXPRESSION            int16 // 0x85
	MOVE_PLAYER                  int16 // 0x7E
	MOVE_MONSTER                 int16 // 0xA2
	CLOSE_RANGE_ATTACK           int16 // 0x7F
	RANGED_ATTACK                int16 // 0x80
	MAGIC_ATTACK                 int16 // 0x81
	OPEN_NPC_SHOP                int16 // 0xe5
	CONFIRM_SHOP_TRANSACTION     int16 // 0xe6
	OPEN_STORAGE                 int16 // 0xe8
	MODIFY_INVENTORY_ITEM        int16 // 0x19
	REMOVE_PLAYER_FROM_MAP       int16 // 0x65
	REMOVE_ITEM_FROM_MAP         int16 // 0xC2
	UPDATE_CHAR_LOOK             int16 // 0x88
	SHOW_FOREIGN_EFFECT          int16 //0x89
	GIVE_FOREIGN_BUFF            int16 //0x8A
	CANCEL_FOREIGN_BUFF          int16 //0x8B
	DAMAGE_PLAYER                int16 // 0x84
	CHAR_INFO                    int16 // 0x31
	UPDATE_QUEST_INFO            int16 // 0x97
	GIVE_BUFF                    int16 //0x1c
	CANCEL_BUFF                  int16 //0x1d
	PLAYER_INTERACTION           int16 // 0xEF
	UPDATE_CHAR_BOX              int16 // 0x69
	NPC_TALK                     int16
	KEYMAP                       int16
	AUTO_HP_POT                  int16
	AUTO_MP_POT                  int16
	SHOW_MONSTER_HP              int16
	PARTY_OPERATION              int16
	UPDATE_PARTYMEMBER_HP        int16
	MULTICHAT                    int16
	APPLY_MONSTER_STATUS         int16
	CANCEL_MONSTER_STATUS        int16
	CLOCK                        int16
	SPAWN_PORTAL                 int16
	SPAWN_DOOR                   int16
	REMOVE_DOOR                  int16
	SPAWN_LOVE                   int16
	REMOVE_LOVE                  int16
	SPAWN_SPECIAL_MAPOBJECT      int16
	REMOVE_SPECIAL_MAPOBJECT     int16
	SUMMON_ATTACK                int16
	MOVE_SUMMON                  int16
	SPAWN_MIST                   int16
	REMOVE_MIST                  int16
	DAMAGE_SUMMON                int16
	DAMAGE_MONSTER               int16
	BUDDYLIST                    int16
	SHOW_ITEM_EFFECT             int16
	SHOW_CHAIR                   int16
	CANCEL_CHAIR                 int16
	SKILL_EFFECT                 int16
	CANCEL_SKILL_EFFECT          int16
	BOSS_ENV                     int16
	REACTOR_SPAWN                int16
	REACTOR_HIT                  int16
	REACTOR_DESTROY              int16
	MAP_EFFECT                   int16
	GUILD_OPERATION              int16
	ALLIANCE_OPERATION           int16
	BBS_OPERATION                int16
	SHOW_MAGNET                  int16
	MESSENGER                    int16
	NPC_ACTION                   int16
	SPAWN_PET                    int16
	MOVE_PET                     int16
	PET_CHAT                     int16
	PET_COMMAND                  int16
	PET_NAMECHANGE               int16
	COOLDOWN                     int16
	PLAYER_HINT                  int16
	SPAWN_HIRED                  int16 `properties:"SPAWN_HIRED,default=65535"`
	USE_SKILL_BOOK               int16
	FORCED_MAP_EQUIP             int16
	SKILL_MACRO                  int16
	CS_OPEN                      int16
	CS_UPDATE                    int16
	CS_OPERATION                 int16
	MTS_OPEN                     int16
	MTS_OPERATION                int16
	MTS_OPERATION2               int16
	PLAYER_NPC                   int16
	SHOW_NOTES                   int16
	SUMMON_SKILL                 int16
	ARIANT_PQ_START              int16
	CATCH_MONSTER                int16
	ARIANT_SCOREBOARD            int16 `properties:"ARIANT_SCOREBOARD,default=65535"`
	ZAKUM_SHRINE                 int16
	BOAT_EFFECT                  int16
	CHALKBOARD                   int16
	DUEY                         int16
	MONSTER_CARNIVAL_START       int16
	MONSTER_CARNIVAL_OBTAINED_CP int16
	MONSTER_CARNIVAL_PARTY_CP    int16
	MONSTER_CARNIVAL_SUMMON      int16
	MONSTER_CARNIVAL_DIED        int16
	SEND_TV                      int16
	REMOVE_TV                    int16
	ENABLE_TV                    int16
	TROCK_LOCATIONS              int16 //0x27
	SPOUSE_CHAT                  int16 // 0x66
	YELLOW_TIP                   int16 `properties:"YELLOW_TIP,default=65535"`
	REPORT_PLAYER_MSG            int16
	SPAWN_HIRED_MERCHANT         int16
	DESTROY_HIRED_MERCHANT       int16
	UPDATE_HIRED_MERCHANT        int16
	GM_POLICE                    int16
	UPDATE_MOUNT                 int16
	GM                           int16
	MONSTERBOOK_ADD              int16
	MONSTER_BOOK_CHANGE_COVER    int16
	CYGNUS_INTRO_LOCK            int16
	CYGNUS_INTRO_DISABLE_UI      int16
	CYGNUS_CHAR_CREATED          int16
	TUTORIAL_DISABLE_UI          int16
	TUTORIAL_LOCK_UI             int16
	TUTORIAL_SUMMON              int16
	TUTORIAL_GUIDE               int16
	SHOW_INFO                    int16 `properties:"SHOW_INFO,default=65535"`
	COMBO_EFFECE                 int16
	Animation_EFFECT             int16
	VICIOUS_HAMMER               int16
	BLOCK_MSG                    int16
	FAMILY                       int16 `properties:"FAMILY,default=65535"`
}
