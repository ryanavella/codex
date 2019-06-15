// Package book provides an API for queries on Biblical books, chapters, verses, etc.
package book

import "errors"

// Errors used by the book subpackage.
var (
	ErrUnknownBook = errors.New("codex.book: unknown book")
)

// ID is a unique identifier for a book of scripture.
type ID int16

// The supported book identifiers.
const (
	UNK ID = iota - 1 // Unknown book
	GEN               // Genesis
	EXO               // Exodus
	LEV               // Leviticus
	NUM               // Numbers
	DEU               // Deuteronomy
	JOS               // Joshua
	JDG               // Judges
	RUT               // Ruth
	SA1               // 1 Samuel
	SA2               // 2 Samuel
	KI1               // 1 Kings
	KI2               // 2 Kings
	CH1               // 1 Chronicles
	CH2               // 2 Chronicles
	EZR               // Ezra
	NEH               // Nehemiah
	EST               // Esther (Hebrew)
	JOB               // Job
	PSA               // Psalms
	PRO               // Proverbs
	ECC               // Ecclesiastes
	SNG               // Song of Songs
	ISA               // Isaiah
	JER               // Jeremiah
	LAM               // Lamentations
	EZK               // Ezekiel
	DAN               // Daniel (Hebrew)
	HOS               // Hosea
	JOL               // Joel
	AMO               // Amos
	OBA               // Obadiah
	JON               // Jonah
	MIC               // Micah
	NAM               // Nahum
	HAB               // Habakkuk
	ZEP               // Zephaniah
	HAG               // Haggai
	ZEC               // Zechariah
	MAL               // Malachi
	MAT               // Matthew
	MRK               // Mark
	LUK               // Luke
	JHN               // John
	ACT               // Acts
	ROM               // Romans
	CO1               // 1 Corinthians
	CO2               // 2 Corinthians
	GAL               // Galatians
	EPH               // Ephesians
	PHP               // Philippians
	COL               // Colossians
	TH1               // 1 Thessalonians
	TH2               // 2 Thessalonians
	TI1               // 1 Timothy
	TI2               // 2 Timothy
	TIT               // Titus
	PHM               // Philemon
	HEB               // Hebrews
	JAS               // James
	PE1               // 1 Peter
	PE2               // 2 Peter
	JN1               // 1 John
	JN2               // 2 John
	JN3               // 3 John
	JUD               // Jude
	REV               // Revelation
	TOB               // Tobit
	JDT               // Judith
	ESG               // Esther (Greek)
	WIS               // Wisdom of Solomon
	SIR               // Sirach
	BAR               // Baruch
	LJE               // Letter of Jeremiah
	S3Y               // Song of the 3 Young Men
	SUS               // Susanna
	BEL               // Bel and the Dragon
	MA1               // 1 Maccabees
	MA2               // 2 Maccabees
	MA3               // 3 Maccabees
	MA4               // 4 Maccabees
	ES1               // 1 Esdras (Greek)
	ES2               // 2 Esdras (Latin)
	MAN               // Prayer of Manasseh
	PS2               // Psalm 151
	ODA               // Odae/Odes
	PSS               // Psalms of Solomon
	EZA               // Ezra Apocalypse
	EZ5               // 5 Ezra
	EZ6               // 6 Ezra
	DAG               // Daniel (Greek)
	PS3               // Psalms 152-155
	BA2               // 2 Baruch (Apocalypse)
	LBA               // Letter of Baruch
	JUB               // Jubilees
	ENO               // Enoch
	MQ1               // 1 Meqabyan/Mekabis
	MQ2               // 2 Meqabyan/Mekabis
	MQ3               // 3 Meqabyan/Mekabis
	REP               // Reproof
	BA4               // 4 Baruch
	LAO               // Letter to the Laodiceans
)

// Parse attempts to parse a string for a book identifier and returns an ID.
//
// TODO: implement
func Parse(s string) (ID, error) {
	return UNK, ErrUnknownBook
}

// MetaData is a representation of a book's metadata.
type MetaData struct {
	InfoUSFM MetaUSFM
}

// Fetch returns the MetaData object for a given book ID.
func Fetch(id ID) MetaData {
	return metaTab[id]
}

// MetaUSFM is a representation of USFM's standard book identifiers along with some handy metadata.
//
// See: https://www.ubsicap.github.io/usfm/identification/books.html?highlight=abbreviation
type MetaUSFM struct {
	Number     string
	Identifier string
	NameEng    string
	NamesAlt   []string
}
