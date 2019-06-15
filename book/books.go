package book

var metaTab = map[ID]MetaData{
	GEN: {
		InfoUSFM: MetaUSFM{
			Number:     "01",
			Identifier: "GEN",
			NameEng:    "Genesis",
			NamesAlt: []string{
				"1 Moses",
			},
		},
	},
	EXO: {
		InfoUSFM: MetaUSFM{
			Number:     "02",
			Identifier: "EXO",
			NameEng:    "Exodus",
			NamesAlt: []string{
				"2 Moses",
			},
		},
	},
	LEV: {
		InfoUSFM: MetaUSFM{
			Number:     "03",
			Identifier: "LEV",
			NameEng:    "Leviticus",
			NamesAlt: []string{
				"3 Moses",
			},
		},
	},
	NUM: {
		InfoUSFM: MetaUSFM{
			Number:     "04",
			Identifier: "NUM",
			NameEng:    "Numbers",
			NamesAlt: []string{
				"4 Moses",
			},
		},
	},
	DEU: {
		InfoUSFM: MetaUSFM{
			Number:     "05",
			Identifier: "DEU",
			NameEng:    "Deuteronomy",
			NamesAlt: []string{
				"5 Moses",
			},
		},
	},
	JOS: {
		InfoUSFM: MetaUSFM{
			Number:     "06",
			Identifier: "JOS",
			NameEng:    "Joshua",
			NamesAlt:   nil,
		},
	},
	JDG: {
		InfoUSFM: MetaUSFM{
			Number:     "07",
			Identifier: "JDG",
			NameEng:    "Judges",
			NamesAlt:   nil,
		},
	},
	RUT: {
		InfoUSFM: MetaUSFM{
			Number:     "08",
			Identifier: "RUT",
			NameEng:    "Ruth",
			NamesAlt:   nil,
		},
	},
	SA1: {
		InfoUSFM: MetaUSFM{
			Number:     "09",
			Identifier: "1SA",
			NameEng:    "1 Samuel",
			NamesAlt: []string{
				"1 Kings",
				"1 Kingdoms",
			},
		},
	},
	SA2: {
		InfoUSFM: MetaUSFM{
			Number:     "10",
			Identifier: "2SA",
			NameEng:    "2 Samuel",
			NamesAlt: []string{
				"2 Kings",
				"2 Kingdoms",
			},
		},
	},
	KI1: {
		InfoUSFM: MetaUSFM{
			Number:     "11",
			Identifier: "1KI",
			NameEng:    "1 Kings",
			NamesAlt: []string{
				"3 Kings",
				"3 Kingdoms",
			},
		},
	},
	KI2: {
		InfoUSFM: MetaUSFM{
			Number:     "12",
			Identifier: "2KI",
			NameEng:    "2 Kings",
			NamesAlt: []string{
				"4 Kings",
				"4 Kingdoms",
			},
		},
	},
	CH1: {
		InfoUSFM: MetaUSFM{
			Number:     "13",
			Identifier: "1CH",
			NameEng:    "1 Chronicles",
			NamesAlt: []string{
				"1 Paralipomenon",
			},
		},
	},
	CH2: {
		InfoUSFM: MetaUSFM{
			Number:     "14",
			Identifier: "2CH",
			NameEng:    "2 Chronicles",
			NamesAlt: []string{
				"2 Paralipomenon",
			},
		},
	},
	EZR: {
		InfoUSFM: MetaUSFM{
			Number:     "15",
			Identifier: "EZR",
			NameEng:    "Ezra",
			NamesAlt: []string{
				"1 Ezra",
				"1 Esdras",
			},
		},
	},
	NEH: {
		InfoUSFM: MetaUSFM{
			Number:     "16",
			Identifier: "NEH",
			NameEng:    "Nehemiah",
			NamesAlt: []string{
				"2 Esdras",
			},
		},
	},
	EST: {
		InfoUSFM: MetaUSFM{
			Number:     "17",
			Identifier: "EST",
			NameEng:    "Esther (Hebrew)",
			NamesAlt:   nil,
		},
	},
	JOB: {
		InfoUSFM: MetaUSFM{
			Number:     "18",
			Identifier: "JOB",
			NameEng:    "Job",
			NamesAlt:   nil,
		},
	},
	PSA: {
		InfoUSFM: MetaUSFM{
			Number:     "19",
			Identifier: "PSA",
			NameEng:    "Psalms",
			NamesAlt:   nil,
		},
	},
	PRO: {
		InfoUSFM: MetaUSFM{
			Number:     "20",
			Identifier: "PRO",
			NameEng:    "Proverbs",
			NamesAlt:   nil,
		},
	},
	ECC: {
		InfoUSFM: MetaUSFM{
			Number:     "21",
			Identifier: "ECC",
			NameEng:    "Ecclesiastes",
			NamesAlt: []string{
				"Qoholeth",
			},
		},
	},
	SNG: {
		InfoUSFM: MetaUSFM{
			Number:     "22",
			Identifier: "SNG",
			NameEng:    "Song of Songs",
			NamesAlt: []string{
				"Song of Solomon",
				"Canticles of Canticles",
			},
		},
	},
	ISA: {
		InfoUSFM: MetaUSFM{
			Number:     "23",
			Identifier: "ISA",
			NameEng:    "Isaiah",
			NamesAlt:   nil,
		},
	},
	JER: {
		InfoUSFM: MetaUSFM{
			Number:     "24",
			Identifier: "JER",
			NameEng:    "Jeremiah",
			NamesAlt: []string{
				"The Book of Jeremiah",
			},
		},
	},
	LAM: {
		InfoUSFM: MetaUSFM{
			Number:     "25",
			Identifier: "LAM",
			NameEng:    "Lamentations",
			NamesAlt: []string{
				"The Lamentations of Jeremiah",
			},
		},
	},
	EZK: {
		InfoUSFM: MetaUSFM{
			Number:     "26",
			Identifier: "EZK",
			NameEng:    "Ezekiel",
			NamesAlt:   nil,
		},
	},
	DAN: {
		InfoUSFM: MetaUSFM{
			Number:     "27",
			Identifier: "DAN",
			NameEng:    "Daniel (Hebrew)",
			NamesAlt:   nil,
		},
	},
	HOS: {
		InfoUSFM: MetaUSFM{
			Number:     "28",
			Identifier: "HOS",
			NameEng:    "Hosea",
			NamesAlt:   nil,
		},
	},
	JOL: {
		InfoUSFM: MetaUSFM{
			Number:     "29",
			Identifier: "JOL",
			NameEng:    "Joel",
			NamesAlt:   nil,
		},
	},
	AMO: {
		InfoUSFM: MetaUSFM{
			Number:     "30",
			Identifier: "AMO",
			NameEng:    "Amos",
			NamesAlt:   nil,
		},
	},
	OBA: {
		InfoUSFM: MetaUSFM{
			Number:     "31",
			Identifier: "OBA",
			NameEng:    "Obadiah",
			NamesAlt:   nil,
		},
	},
	JON: {
		InfoUSFM: MetaUSFM{
			Number:     "32",
			Identifier: "JON",
			NameEng:    "Jonah",
			NamesAlt:   nil,
		},
	},
	MIC: {
		InfoUSFM: MetaUSFM{
			Number:     "33",
			Identifier: "MIC",
			NameEng:    "Micah",
			NamesAlt:   nil,
		},
	},
	NAM: {
		InfoUSFM: MetaUSFM{
			Number:     "34",
			Identifier: "NAM",
			NameEng:    "Nahum",
			NamesAlt:   nil,
		},
	},
	HAB: {
		InfoUSFM: MetaUSFM{
			Number:     "35",
			Identifier: "HAB",
			NameEng:    "Habakkuk",
			NamesAlt:   nil,
		},
	},
	ZEP: {
		InfoUSFM: MetaUSFM{
			Number:     "36",
			Identifier: "ZEP",
			NameEng:    "Zephaniah",
			NamesAlt:   nil,
		},
	},
	HAG: {
		InfoUSFM: MetaUSFM{
			Number:     "37",
			Identifier: "HAG",
			NameEng:    "Haggai",
			NamesAlt:   nil,
		},
	},
	ZEC: {
		InfoUSFM: MetaUSFM{
			Number:     "38",
			Identifier: "ZEC",
			NameEng:    "Zechariah",
			NamesAlt:   nil,
		},
	},
	MAL: {
		InfoUSFM: MetaUSFM{
			Number:     "39",
			Identifier: "MAL",
			NameEng:    "Malachi",
			NamesAlt:   nil,
		},
	},
	MAT: {
		InfoUSFM: MetaUSFM{
			Number:     "41",
			Identifier: "MAT",
			NameEng:    "Matthew",
			NamesAlt: []string{
				"The Gospel according to Matthew",
			},
		},
	},
	MRK: {
		InfoUSFM: MetaUSFM{
			Number:     "42",
			Identifier: "MRK",
			NameEng:    "Mark",
			NamesAlt: []string{
				"The Gospel according to Mark",
			},
		},
	},
	LUK: {
		InfoUSFM: MetaUSFM{
			Number:     "43",
			Identifier: "LUK",
			NameEng:    "Luke",
			NamesAlt: []string{
				"The Gospel according to Luke",
			},
		},
	},
	JHN: {
		InfoUSFM: MetaUSFM{
			Number:     "44",
			Identifier: "JHN",
			NameEng:    "John",
			NamesAlt: []string{
				"The Gospel according to John",
			},
		},
	},
	ACT: {
		InfoUSFM: MetaUSFM{
			Number:     "45",
			Identifier: "ACT",
			NameEng:    "Acts",
			NamesAlt: []string{
				"The Acts of the Apostles",
			},
		},
	},
	ROM: {
		InfoUSFM: MetaUSFM{
			Number:     "46",
			Identifier: "ROM",
			NameEng:    "Romans",
			NamesAlt: []string{
				"The Letter of Paul to the Romans",
			},
		},
	},
	CO1: {
		InfoUSFM: MetaUSFM{
			Number:     "47",
			Identifier: "1CO",
			NameEng:    "1 Corinthians",
			NamesAlt: []string{
				"The First Letter of Paul to the Corinthians",
			},
		},
	},
	CO2: {
		InfoUSFM: MetaUSFM{
			Number:     "48",
			Identifier: "2CO",
			NameEng:    "2 Corinthians",
			NamesAlt: []string{
				"The Second Letter of Paul to the Corinthians",
			},
		},
	},
	GAL: {
		InfoUSFM: MetaUSFM{
			Number:     "49",
			Identifier: "GAL",
			NameEng:    "Galatians",
			NamesAlt: []string{
				"The Letter of Paul to the Galatians",
			},
		},
	},
	EPH: {
		InfoUSFM: MetaUSFM{
			Number:     "50",
			Identifier: "EPH",
			NameEng:    "Ephesians",
			NamesAlt: []string{
				"The Letter of Paul to the Ephesians",
			},
		},
	},
	PHP: {
		InfoUSFM: MetaUSFM{
			Number:     "51",
			Identifier: "PHP",
			NameEng:    "Philippians",
			NamesAlt: []string{
				"The Letter of Paul to the Philippians",
			},
		},
	},
	COL: {
		InfoUSFM: MetaUSFM{
			Number:     "52",
			Identifier: "COL",
			NameEng:    "Colossians",
			NamesAlt: []string{
				"The Letter of Paul to the Colossians",
			},
		},
	},
	TH1: {
		InfoUSFM: MetaUSFM{
			Number:     "53",
			Identifier: "1TH",
			NameEng:    "1 Thessalonians",
			NamesAlt: []string{
				"The First Letter of Paul to the Thessalonians",
			},
		},
	},
	TH2: {
		InfoUSFM: MetaUSFM{
			Number:     "54",
			Identifier: "2TH",
			NameEng:    "2 Thessalonians",
			NamesAlt: []string{
				"The Second Letter of Paul to the Thessalonians",
			},
		},
	},
	TI1: {
		InfoUSFM: MetaUSFM{
			Number:     "55",
			Identifier: "1TI",
			NameEng:    "1 Timothy",
			NamesAlt: []string{
				"The First Letter of Paul to Timothy",
			},
		},
	},
	TI2: {
		InfoUSFM: MetaUSFM{
			Number:     "56",
			Identifier: "2TI",
			NameEng:    "2 Timothy",
			NamesAlt: []string{
				"The Second Letter of Paul to Timothy",
			},
		},
	},
	TIT: {
		InfoUSFM: MetaUSFM{
			Number:     "57",
			Identifier: "TIT",
			NameEng:    "Titus",
			NamesAlt: []string{
				"The Letter of Paul to Titus",
			},
		},
	},
	PHM: {
		InfoUSFM: MetaUSFM{
			Number:     "58",
			Identifier: "PHM",
			NameEng:    "Philemon",
			NamesAlt: []string{
				"The Letter of Paul to Philemon",
			},
		},
	},
	HEB: {
		InfoUSFM: MetaUSFM{
			Number:     "59",
			Identifier: "HEB",
			NameEng:    "Hebrews",
			NamesAlt: []string{
				"The Letter to the Hebrews",
			},
		},
	},
	JAS: {
		InfoUSFM: MetaUSFM{
			Number:     "60",
			Identifier: "JAS",
			NameEng:    "James",
			NamesAlt: []string{
				"The Letter of James",
			},
		},
	},
	PE1: {
		InfoUSFM: MetaUSFM{
			Number:     "61",
			Identifier: "1PE",
			NameEng:    "1 Peter",
			NamesAlt: []string{
				"The First Letter of Peter",
			},
		},
	},
	PE2: {
		InfoUSFM: MetaUSFM{
			Number:     "62",
			Identifier: "2PE",
			NameEng:    "2 Peter",
			NamesAlt: []string{
				"The Second Letter of Peter",
			},
		},
	},
	JN1: {
		InfoUSFM: MetaUSFM{
			Number:     "63",
			Identifier: "1JN",
			NameEng:    "1 John",
			NamesAlt: []string{
				"The First Letter of John",
			},
		},
	},
	JN2: {
		InfoUSFM: MetaUSFM{
			Number:     "64",
			Identifier: "2JN",
			NameEng:    "2 John",
			NamesAlt: []string{
				"The Second Letter of John",
			},
		},
	},
	JN3: {
		InfoUSFM: MetaUSFM{
			Number:     "65",
			Identifier: "3JN",
			NameEng:    "3 John",
			NamesAlt: []string{
				"The Third Letter of John",
			},
		},
	},
	JUD: {
		InfoUSFM: MetaUSFM{
			Number:     "66",
			Identifier: "JUD",
			NameEng:    "Jude",
			NamesAlt: []string{
				"The Letter of Jude",
			},
		},
	},
	REV: {
		InfoUSFM: MetaUSFM{
			Number:     "67",
			Identifier: "REV",
			NameEng:    "Revelation",
			NamesAlt: []string{
				"The Revelation to John",
				"Apocalypse",
			},
		},
	},
	TOB: {
		InfoUSFM: MetaUSFM{
			Number:     "68",
			Identifier: "TOB",
			NameEng:    "Tobit",
			NamesAlt:   nil,
		},
	},
	JDT: {
		InfoUSFM: MetaUSFM{
			Number:     "69",
			Identifier: "JDT",
			NameEng:    "Judith",
			NamesAlt:   nil,
		},
	},
	ESG: {
		InfoUSFM: MetaUSFM{
			Number:     "70",
			Identifier: "ESG",
			NameEng:    "Esther Greek",
			NamesAlt:   nil,
		},
	},
	WIS: {
		InfoUSFM: MetaUSFM{
			Number:     "71",
			Identifier: "WIS",
			NameEng:    "Wisdom of Solomon",
			NamesAlt:   nil,
		},
	},
	SIR: {
		InfoUSFM: MetaUSFM{
			Number:     "72",
			Identifier: "SIR",
			NameEng:    "Sirach",
			NamesAlt: []string{
				"Ecclesiasticus",
				"Jesus son of Sirach",
			},
		},
	},
	BAR: {
		InfoUSFM: MetaUSFM{
			Number:     "73",
			Identifier: "BAR",
			NameEng:    "Baruch",
			NamesAlt: []string{
				"1 Baruch",
			},
		},
	},
	LJE: {
		InfoUSFM: MetaUSFM{
			Number:     "74",
			Identifier: "LJE",
			NameEng:    "Letter of Jeremiah",
			NamesAlt: []string{
				"Rest of Jeremiah",
			},
		},
	},
	S3Y: {
		InfoUSFM: MetaUSFM{
			Number:     "75",
			Identifier: "S3Y",
			NameEng:    "Song of the 3 Young Men",
			NamesAlt:   nil,
		},
	},
	SUS: {
		InfoUSFM: MetaUSFM{
			Number:     "76",
			Identifier: "SUS",
			NameEng:    "Susanna",
			NamesAlt:   nil,
		},
	},
	BEL: {
		InfoUSFM: MetaUSFM{
			Number:     "77",
			Identifier: "BEL",
			NameEng:    "Bel and the Dragon",
			NamesAlt: []string{
				"Rest of Daniel",
			},
		},
	},
	MA1: {
		InfoUSFM: MetaUSFM{
			Number:     "78",
			Identifier: "1MA",
			NameEng:    "1 Maccabees",
			NamesAlt: []string{
				"3 Maccabees",
			},
		},
	},
	MA2: {
		InfoUSFM: MetaUSFM{
			Number:     "79",
			Identifier: "2MA",
			NameEng:    "2 Maccabees",
			NamesAlt: []string{
				"1 Maccabees",
			},
		},
	},
	MA3: {
		InfoUSFM: MetaUSFM{
			Number:     "80",
			Identifier: "3MA",
			NameEng:    "3 Maccabees",
			NamesAlt: []string{
				"2 Maccabees",
			},
		},
	},
	MA4: {
		InfoUSFM: MetaUSFM{
			Number:     "81",
			Identifier: "4MA",
			NameEng:    "4 Maccabees",
			NamesAlt:   nil,
		},
	},
	ES1: {
		InfoUSFM: MetaUSFM{
			Number:     "82",
			Identifier: "1ES",
			NameEng:    "1 Esdras (Greek)",
			NamesAlt: []string{
				"2 Esdras",
				"3 Esdras",
			},
		},
	},
	ES2: {
		InfoUSFM: MetaUSFM{
			Number:     "83",
			Identifier: "2ES",
			NameEng:    "2 Esdras (Latin)",
			NamesAlt: []string{
				"3 Esdras",
				"4 Esdras",
			},
		},
	},
	MAN: {
		InfoUSFM: MetaUSFM{
			Number:     "84",
			Identifier: "MAN",
			NameEng:    "Prayer of Manasseh",
			NamesAlt:   nil,
		},
	},
	PS2: {
		InfoUSFM: MetaUSFM{
			Number:     "85",
			Identifier: "PS2",
			NameEng:    "Psalm 151",
			NamesAlt:   nil,
		},
	},
	ODA: {
		InfoUSFM: MetaUSFM{
			Number:     "86",
			Identifier: "ODA",
			NameEng:    "Odae/Odes",
			NamesAlt:   nil,
		},
	},
	PSS: {
		InfoUSFM: MetaUSFM{
			Number:     "87",
			Identifier: "PSS",
			NameEng:    "Psalms of Solomon",
			NamesAlt:   nil,
		},
	},
	EZA: {
		InfoUSFM: MetaUSFM{
			Number:     "A4",
			Identifier: "EZA",
			NameEng:    "Ezra Apocalypse",
			NamesAlt: []string{
				"3 Ezra",
				"Ezra Shealtiel",
			},
		},
	},
	EZ5: {
		InfoUSFM: MetaUSFM{
			Number:     "A5",
			Identifier: "5EZ",
			NameEng:    "5 Ezra",
			NamesAlt:   nil,
		},
	},
	EZ6: {
		InfoUSFM: MetaUSFM{
			Number:     "A6",
			Identifier: "6EZ",
			NameEng:    "6 Ezra",
			NamesAlt:   nil,
		},
	},
	DAG: {
		InfoUSFM: MetaUSFM{
			Number:     "B2",
			Identifier: "DAG",
			NameEng:    "Daniel Greek",
			NamesAlt: []string{
				"Vision of Daniel",
			},
		},
	},
	PS3: {
		InfoUSFM: MetaUSFM{
			Number:     "B3",
			Identifier: "PS3",
			NameEng:    "Psalms 152-155",
			NamesAlt:   nil,
		},
	},
	BA2: {
		InfoUSFM: MetaUSFM{
			Number:     "B4",
			Identifier: "2BA",
			NameEng:    "2 Baruch (Apocalypse)",
			NamesAlt: []string{
				"The Apocalypse of Baruch",
			},
		},
	},
	LBA: {
		InfoUSFM: MetaUSFM{
			Number:     "B5",
			Identifier: "LBA",
			NameEng:    "Letter of Baruch",
			NamesAlt:   nil,
		},
	},
	JUB: {
		InfoUSFM: MetaUSFM{
			Number:     "B6",
			Identifier: "JUB",
			NameEng:    "Jubilees",
			NamesAlt:   nil,
		},
	},
	ENO: {
		InfoUSFM: MetaUSFM{
			Number:     "B7",
			Identifier: "ENO",
			NameEng:    "Enoch",
			NamesAlt: []string{
				"1 Enoch",
			},
		},
	},
	MQ1: {
		InfoUSFM: MetaUSFM{
			Number:     "B8",
			Identifier: "1MQ",
			NameEng:    "1 Meqabyan/Mekabis",
			NamesAlt: []string{
				"Book of Mekabis of Benjamin",
			},
		},
	},
	MQ2: {
		InfoUSFM: MetaUSFM{
			Number:     "B9",
			Identifier: "2MQ",
			NameEng:    "2 Meqabyan/Mekabis",
			NamesAlt: []string{
				"Book of Mekabis of Moab",
			},
		},
	},
	MQ3: {
		InfoUSFM: MetaUSFM{
			Number:     "C0",
			Identifier: "3MQ",
			NameEng:    "3 Meqabyan/Mekabis",
			NamesAlt: []string{
				"Book of Meqabyan",
			},
		},
	},
	REP: {
		InfoUSFM: MetaUSFM{
			Number:     "C1",
			Identifier: "REP",
			NameEng:    "Reproof",
			NamesAlt:   nil,
		},
	},
	BA4: {
		InfoUSFM: MetaUSFM{
			Number:     "C2",
			Identifier: "4BA",
			NameEng:    "4 Baruch",
			NamesAlt: []string{
				"Paralipomenon of Jeremiah",
				"Rest of the Words of Baruch",
			},
		},
	},
	LAO: {
		InfoUSFM: MetaUSFM{
			Number:     "C3",
			Identifier: "LAO",
			NameEng:    "Letter to the Laodiceans",
			NamesAlt:   nil,
		},
	},
}
