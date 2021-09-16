package types

//Result for Antivirus Scan
type Result struct {
	FileName  string
	Signature string
	Status    string
	Raw       string
}
