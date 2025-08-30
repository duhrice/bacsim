package affiliation

import (

)

type (
	School int
	Club int
)

// No playable students from the following schools
// Kivotos, Kronos, Odyssey Maritime
const (
	Abydos School = iota	// Abydos High School
	Arius					// Arius Branch School
	Gehenna					// Gehenna Academy
	Highlander				// Highlander Railroad Academy
	Hyakkiyako				// Hyakkiyako Alliance Academy
	Kivotos					// For General Student Council and SCHALE
	Kronos					// Kronos School of Journalism
	Millennium				// Millennium Science School
	// OdysseyMaritime		// Odyssey Maritime School. Not yet expanded upon in-game yet. Leave commented out
	RedWinter				// Red Winter Federal Academy
	Shanhaijing				// Shanhaijing Senior Secondary School
	Tokiwadai				// Tokiwadai Middle School. For limited Railgun Collab
	Trinity					// Trinity General School
	Valkyrie				// Valkyrie Police Academy
	SRT						// SRT Special Academy
	WildHunt				// Wild Hunt Academy of Arts
	Etc						// For "Hatsune Miku"
	Unknown
)

var SchoolString = [...]string {
	"abydos",				// Abydos High School
	"arius",				// Arius Branch School
	"gehenna",				// Gehenna Academy
	"highlander",			// Highlander Railroad Academy
	"hyakkiyako",			// Hyakkiyako Alliance Academy
	"kivotos",				// For General Student Council and SCHALE
	"kronos",				// Kronos School of Journalism
	"millennium",			// Millennium Science School
	// "odyssey_maritime",	// Odyssey Maritime School. Not yet expanded upon in-game yet. Leave commented out
	"red_winter",			// Red Winter Federal Academy
	"shanhaijing",			// Shanhaijing Senior Secondary School
	"tokiwadai",			// Tokiwadai Middle School. For limited Railgun Collab
	"trinity",				// Trinity General School
	"valkyrie",				// Valkyrie Police Academy
	"srt",					// SRT Special Academy
	"wild_hunt",			// Wild Hunt Academy of Arts
	"etc",					// For "Hatsune Miku"
	"unknown",
}

const (
	// Abydos High School
	ForeclosureTaskForce Club = iota
	// Arius Branch School
	AriusSquad
	// Gehenna Academy
	PandemoniumSociety
	PrefectTeam
	MedicalEmergencyClub
	ProblemSolver68
	GourmetResearchSociety
	SchoolLunchClub
	HotSpringDepartment
	SparkleClub
	// Highlander Railroad Academy
	CentralControlCenter
	HighlanderSupervisionOffice 	// No playable students from this club yet
	FreightTransportManagementDepartment
	// Hyakkiyako Alliance Academy
	YinYangClub
	HyakkaryouranConflictResolutionCouncil
	InnerDisciplineClub
	NinjutsuResearchClub
	StreetMouryo					// No playable students from this club yet
	NaturesBeautyClub				// No playable students from this club yet
	// Kivotos
	GeneralStudentCouncil			// No playable students from this club yet
	SCHALE							// No playable students from this club yet
	// Kronos School of Journalism
	KronosNewsClub
	// Millennium Science School
	Seminar
	Veritas
	CleaningAndClearing
	SuperPhenomenonTaskForce
	EngineeringClub
	AthleticsTrainingClub
	GameDevelopmentDepartment
	// Odyssey Maritime School
	// Red Winter Federal Academy
	RedWinterOffice
	LaborParty
	KnowledgeLiberationFront
	SpecOpsNo227
	// Shanhaijing Senior Secondary School
	Genryumon
	BlackTortoisePromenade
	EasternAlchemySociety
	PlumBlossomGarden
	// Tokiwadai Middle School
	// Trinity General School
	TeaParty
	Sisterhood
	JusticeTaskForce
	RemedialKnights
	LibraryCommittee
	AfterSchoolSweetsClub
	TrinityVigilanteCrew
	MakeUpWorkClub
	// Valkyrie Police Academy
	PublicPeaceBureau
	PublicSafetyBureau
	// SRT Special Academy
	RABBITSquad
	FOXSquad						// No playable students from this club yet
	// Wild Hunt Academy of Arts (being developed, come back to this frequently)
	OccultResearchSociety
	// Etc
	None
)

var ClubString = [...]string {
	// Abydos High School
	"foreclosure_task_force",
	// Arius Branch School
	"arius_squad",
	// Gehenna Academy
	"pandemonium_society",
	"prefect_team",
	"medical_emergency_club",
	"problem_solver_68",
	"gourmet_research_society",
	"school_lunch_club",
	"hot_spring_department",
	"sparkle_club",
	// Highlander Railroad Academy
	"central_control_center",
	"highlander_supervision_office",// No playable students from this club yet
	"freight_transport_management_department",
	// Hyakkiyako Alliance Academy
	"yin_yang_club",
	"hyakkaryouran_conflict_resolution_council",
	"inner_discipline_club",
	"ninjutsu_research_club",
	"street_mouryo",				// No playable students from this club yet
	"natures_beauty_club",			// No playable students from this club yet
	// Kivotos
	"general_student_council",		// No playable students from this club yet
	"schale",						// No playable students from this club yet
	// Kronos School of Journalism
	"kronos_news_club",
	// Millennium Science School
	"seminar",
	"veritas",
	"cleaning_and_clearing",
	"super_phenomenon_task_force",
	"engineering_club",
	"athletics_training_club",
	"game_development_department",
	// Odyssey Maritime School
	// Red Winter Federal Academy
	"red_winter_office",
	"labor_party",
	"knowledge_liberation_front",
	"spec_ops_no227",
	// Shanhaijing Senior Secondary School
	"genryumon",
	"black_tortoise_promenade",
	"eastern_alchemy_society",
	"plum_blossom_garden",
	// Tokiwadai Middle School
	// Trinity General School
	"tea_party",
	"sisterhood",
	"justice_task_force",
	"remedial_knights",
	"library_committee",
	"after_school_sweets_club",
	"trinity_vigilante_crew",
	"make_up_work_club",
	// Valkyrie Police Academy
	"public_peace_bureau",
	"public_safety_bureau",
	// SRT Special Academy
	"rabbit_squad",
	"fox_squad",					// No playable students from this club yet
	// Wild Hunt Academy of Arts (being developed, come back to this frequently)
	"occult_research_society",
	// Etc
	"n/a",
}