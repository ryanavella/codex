package book

var titleTab = map[string]ID{
	// Genesis
	"genesis": GEN,
	// Exodus
	"exodus": EXO,
	// Leviticus
	"leviticus": LEV,
	// Numbers
	"numbers": NUM,
	// Deuteronomy
	"deuteronomy": DEU,
	// Joshua
	"joshua": JOS,
	// Judges
	"judges": JDG,
	// Ruth
	"ruth": RUT,
	// 1 Samuel
	"1 samuel":   SA1,
	"1st samuel": SA1,
	"i samuel":   SA1,
	// 2 Samuel
	"2 samuel":   SA2,
	"2nd samuel": SA2,
	"ii samuel":  SA2,
	// 1 Kings
	"1 kings":   KI1,
	"1st kings": KI1,
	"i kings":   KI1,
	// 2 Kings
	"2 kings":   KI2,
	"2nd kings": KI2,
	"ii kings":  KI2,
	// 1 Chronicles
	"1 chronicles":   CH1,
	"1st chronicles": CH1,
	"i chronicles":   CH1,
	// 2 Chronicles
	"2 chronicles":   CH2,
	"2nd chronicles": CH2,
	"ii chronicles":  CH2,
	// Ezra
	"ezra": EZR,
	// Nehemiah
	"nehemiah": NEH,
	// Esther (Hebrew)
	"esther": EST,
	// Job
	"job": JOB,
	// Psalms
	"psalm":  PSA,
	"psalms": PSA,
	// Proverbs
	"proverb":  PRO,
	"proverbs": PRO,
	// Ecclesiastes
	"ecclesiastes": ECC,
	// Song of Songs
	"song of songs":          SNG,
	"song of solomon":        SNG,
	"canticles of canticles": SNG,
	// Isaiah
	"isaiah": ISA,
	// Jeremiah
	"jeremiah": JER,
	// Lamentations
	"lamentations":                 LAM,
	"lamentations of jeremiah":     JER,
	"the lamentations of jeremiah": JER,
	// Ezekiel
	"ezekiel": EZK,
	// Daniel (Hebrew)
	"daniel": DAN,
	// Hosea
	"hosea": HOS,
	// Joel
	"joel": JOL,
	// Amos
	"amos": AMO,
	// Obadiah
	"obadiah": OBA,
	// Jonah
	"jonah": JON,
	// Micah
	"micah": MIC,
	// Nahum
	"nahum": NAM,
	// Habakkuk
	"habakkuk": HAB,
	// Zephaniah
	"zephaniah": ZEP,
	// Haggai
	"haggai": HAG,
	// Zechariah
	"zechariah": ZEC,
	// Malachi
	"malachi": MAL,
	// Matthew
	"matthew": MAT,
	// Mark
	"mark": MRK,
	// Luke
	"luke": LUK,
	// John
	"john": JHN,
	// Acts
	"acts": ACT,
	// Romans
	"romans": ROM,
	// 1 Corinthians
	"1 corinthians":   CO1,
	"1st corinthians": CO1,
	"i corinthians":   CO1,
	// 2 Corinthians
	"2 corinthians":   CO2,
	"2nd corinthians": CO2,
	"ii corinthians":  CO2,
	// Galatians
	"galatians": GAL,
	// Ephesians
	"ephesians": EPH,
	// Philippians
	"philippians": PHP,
	// Colossians
	"colossians": COL,
	// 1 Thessalonians
	"1 thessalonians":   TH1,
	"1st thessalonians": TH1,
	"i thessalonians":   TH1,
	// 2 Thessalonians
	"2 thessalonians":   TH2,
	"2nd thessalonians": TH2,
	"ii thessalonians":  TH2,
	// 1 Timothy
	"1 timothy":   TI1,
	"1st timothy": TI1,
	"i timothy":   TI1,
	// 2 Timothy
	"2 timothy":   TI2,
	"2nd timothy": TI2,
	"ii timothy":  TI2,
	// Titus
	"titus": TIT,
	// Philemon
	"philemon": PHM,
	// Hebrews
	"hebrews": HEB,
	// James
	"james": JAS,
	// 1 Peter
	"1 peter":   PE1,
	"1st peter": PE1,
	"i peter":   PE1,
	// 2 Peter
	"2 peter":   PE2,
	"2nd peter": PE2,
	"ii peter":  PE2,
	// 1 John
	"1 john":   JN1,
	"1st john": JN1,
	"i john":   JN1,
	// 2 John
	"2 john":   JN2,
	"2nd john": JN2,
	"ii john":  JN2,
	// 3 John
	"3 john":   JN3,
	"3rd john": JN3,
	"iii john": JN3,
	// Jude
	"jude": JUD,
	// Revelation
	"revelation": REV,
	// Tobit
	"tobit": TOB,
	// Judith
	"judith": JDT,
	// Esther Greek
	"esther greek": ESG,
	// Wisdom of Solomon
	"wisdom of solomon": WIS,
	// Sirach
	"sirach": SIR,
	// Baruch
	"baruch": BAR,
	// Letter of Jeremiah
	"letter of jeremiah": LJE,
	// Song of the 3 Young Men
	"song of the 3 young men": S3Y,
	// Susanna
	"susanna": SUS,
	// Bel and the Dragon
	"bel and the dragon": BEL,
	// 1 Maccabees
	"1 maccabees":   MA1,
	"1st maccabees": MA1,
	"i maccabees":   MA1,
	// 2 Maccabees
	"2 maccabees":   MA2,
	"2nd maccabees": MA2,
	"ii maccabees":  MA2,
	// 3 Maccabees
	"3 maccabees":   MA3,
	"3rd maccabees": MA3,
	"iii maccabees": MA3,
	// 4 Maccabees
	"4 maccabees":   MA4,
	"4th maccabees": MA4,
	"iv maccabees":  MA4,
	// 1 Esdras (Greek)
	"1 esdras greek": ES1,
	// 2 Esdras (Latin)
	"2 esdras latin": ES2,
	// Prayer of Manasseh
	"prayer of manasseh": MAN,
	// Psalm 151
	"psalm 151": PS2,
	// Odae/Odes
	"odas": ODA,
	"odae": ODA,
	// Psalms of Solomon
	"psalm of solomon":  PSS,
	"psalms of solomon": PSS,
	// Ezra Apocalypse
	"ezra apocalypse": EZA,
	// 5 Ezra
	"5 ezra": EZ5,
	// 6 Ezra
	"6 ezra": EZ6,
	// Daniel Greek
	"daniel greek": DAG,
	// Psalms 152-155
	"psalms 152 155": PS3,
	// 2 Baruch (Apocalypse)
	"2 baruch apocalypse":  BA2,
	"apocalypse of baruch": BA2,
	// Letter of Baruch
	"letter of baruch": LBA,
	// Jubilees
	"jubilees": JUB,
	// Enoch
	"enoch": ENO,
	// 1 Meqabyan/Mekabis
	"1 meqabyan":   MQ1,
	"1st meqabyan": MQ1,
	"i meqabyan":   MQ1,
	"1 mekabis":    MQ1,
	"1st mekabis":  MQ1,
	"i mekabis":    MQ1,
	// 2 Meqabyan/Mekabis
	"2 meqabyan":   MQ2,
	"2st meqabyan": MQ2,
	"ii meqabyan":  MQ2,
	"2 mekabis":    MQ2,
	"2st mekabis":  MQ2,
	"ii mekabis":   MQ2,
	// 3 Meqabyan/Mekabis
	"3 meqabyan":   MQ3,
	"3st meqabyan": MQ3,
	"iii meqabyan": MQ3,
	"3 mekabis":    MQ3,
	"3st mekabis":  MQ3,
	"iii mekabis":  MQ3,
	// Reproof
	"reproof": REP,
	// 4 Baruch
	"4 baruch":   BA4,
	"4th baruch": BA4,
	"iv baruch":  BA4,
	// Letter to the Laodiceans
	"letter to the laodiceans": LAO,
}
