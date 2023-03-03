package english

type Dateformat string

type SDateformats struct {
	ANSIC       Dateformat
	Kitchen     Dateformat
	mmddyy      Dateformat
	mmmddyy     Dateformat
	mmddyyyy    Dateformat
	mmmddyyyy   Dateformat
	RFC1123     Dateformat
	RFC1123Z    Dateformat
	RFC3339     Dateformat
	RFC3339Nano Dateformat
	RFC822      Dateformat
	RFC822Z     Dateformat
	RFC850      Dateformat
	RubyDate    Dateformat
	UnixDate    Dateformat
}

//Jan/2/2006

var Dateformats = SDateformats{
	ANSIC:       "Mon Jan _2 15:04:05 2006",
	Kitchen:     "3:04PM",
	mmddyy:      "01/02/06",
	mmmddyy:     "Jan 02, 2006",
	mmddyyyy:    "Jan/02/2006",
	mmmddyyyy:   "01-02-2006",
	RFC1123:     "Mon, 02 Jan 2006 15:04:05 MST",
	RFC1123Z:    "Mon, 02 Jan 2006 15:04:05 -0700",
	RFC3339:     "2006-01-02T15:04:05Z07:00",
	RFC3339Nano: "2006-01-02T15:04:05.999999999Z07:00",
	RFC822:      "02 Jan 06 15:04 MST",
	RFC822Z:     "02 Jan 06 15:04 -0700",
	RFC850:      "Monday, 02-Jan-06 15:04:05 MST",
	RubyDate:    "Mon Jan 02 15:04:05 -0700 2006",
	UnixDate:    "Mon Jan _2 15:04:05 MST 2006",
}
