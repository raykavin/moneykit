package moneykit

import (
	"strings"
)

// Currency represents money currency information required for formatting and calculations.
// It includes the currency code, symbol, decimal places, and formatting templates.
//
// Fields:
//   - Code: ISO 4217 currency code (e.g., "USD", "EUR")
//   - NumericCode: ISO 4217 numeric code (e.g., "840" for USD)
//   - Fraction: Number of decimal places (e.g., 2 for USD, 0 for JPY)
//   - Grapheme: Currency symbol (e.g., "$", "€", "¥")
//   - Template: Formatting template (e.g., "$1" for $100, "1 $" for 100 $)
//   - Decimal: Decimal separator (e.g., "." or ",")
//   - Thousand: Thousands separator (e.g., "," or ".")
//
// Example:
//
//	currency := moneykit.GetCurrency("USD")
//	fmt.Println(currency.Code)        // USD
//	fmt.Println(currency.Grapheme)    // $
//	fmt.Println(currency.Fraction)    // 2
//	fmt.Println(currency.Template)    // $1
type Currency struct {
	Code        string
	NumericCode string
	Fraction    int
	Grapheme    string
	Template    string
	Decimal     string
	Thousand    string
}

// Currencies is a map of currency codes to Currency instances.
// It provides methods to look up currencies by code or numeric code.
type Currencies map[string]*Currency

// CurrencyByNumericCode returns the Currency for the given ISO 4217 numeric code.
// Returns nil if the currency is not found.
//
// Example:
//
//	currencies := moneykit.Currencies{"USD": usdCurrency}
//	usd := currencies.CurrencyByNumericCode("840") // USD numeric code
func (c Currencies) CurrencyByNumericCode(code string) *Currency {
	for _, sc := range c {
		if sc.NumericCode == code {
			return sc
		}
	}

	return nil
}

// CurrencyByCode returns the Currency for the given ISO 4217 currency code.
// Returns nil if the currency is not found.
//
// Example:
//
//	currencies := moneykit.Currencies{"USD": usdCurrency}
//	usd := currencies.CurrencyByCode("USD")
func (c Currencies) CurrencyByCode(code string) *Currency {
	sc, ok := c[code]
	if !ok {
		return nil
	}

	return sc
}

// Add adds or updates a Currency in the currencies collection.
//
// Example:
//
//	currencies := make(moneykit.Currencies)
//	currencies.Add(&moneykit.Currency{Code: "BTC", Grapheme: "₿"})
func (c Currencies) Add(currency *Currency) Currencies {
	c[currency.Code] = currency
	return c
}

// currencies represents a collection of currency.
var currencies = Currencies{
	AED: {Decimal: ".", Thousand: ",", Code: AED, Fraction: 2, NumericCode: "784", Grapheme: ".\u062f.\u0625", Template: "1 $"},
	AFN: {Decimal: ".", Thousand: ",", Code: AFN, Fraction: 2, NumericCode: "971", Grapheme: "\u060b", Template: "1 $"},
	ALL: {Decimal: ".", Thousand: ",", Code: ALL, Fraction: 2, NumericCode: "008", Grapheme: "L", Template: "$1"},
	AMD: {Decimal: ".", Thousand: ",", Code: AMD, Fraction: 2, NumericCode: "051", Grapheme: "\u0564\u0580.", Template: "1 $"},
	ANG: {Decimal: ",", Thousand: ".", Code: ANG, Fraction: 2, NumericCode: "532", Grapheme: "\u0192", Template: "$1"},
	AOA: {Decimal: ".", Thousand: ",", Code: AOA, Fraction: 2, NumericCode: "973", Grapheme: "Kz", Template: "1$"},
	ARS: {Decimal: ",", Thousand: ".", Code: ARS, Fraction: 2, NumericCode: "032", Grapheme: "$", Template: "$1"},
	AUD: {Decimal: ".", Thousand: ",", Code: AUD, Fraction: 2, NumericCode: "036", Grapheme: "A$", Template: "$1"},
	AWG: {Decimal: ".", Thousand: ",", Code: AWG, Fraction: 2, NumericCode: "533", Grapheme: "\u0192", Template: "1$"},
	AZN: {Decimal: ".", Thousand: ",", Code: AZN, Fraction: 2, NumericCode: "944", Grapheme: "\u20bc", Template: "$1"},
	BAM: {Decimal: ".", Thousand: ",", Code: BAM, Fraction: 2, NumericCode: "977", Grapheme: "KM", Template: "$1"},
	BBD: {Decimal: ".", Thousand: ",", Code: BBD, Fraction: 2, NumericCode: "052", Grapheme: "$", Template: "$1"},
	BDT: {Decimal: ".", Thousand: ",", Code: BDT, Fraction: 2, NumericCode: "050", Grapheme: "\u09f3", Template: "$1"},
	BGN: {Decimal: ".", Thousand: ",", Code: BGN, Fraction: 2, NumericCode: "975", Grapheme: "\u043b\u0432", Template: "$1"},
	BHD: {Decimal: ".", Thousand: ",", Code: BHD, Fraction: 3, NumericCode: "048", Grapheme: ".\u062f.\u0628", Template: "1 $"},
	BIF: {Decimal: ".", Thousand: ",", Code: BIF, Fraction: 0, NumericCode: "108", Grapheme: "Fr", Template: "1$"},
	BMD: {Decimal: ".", Thousand: ",", Code: BMD, Fraction: 2, NumericCode: "060", Grapheme: "$", Template: "$1"},
	BND: {Decimal: ".", Thousand: ",", Code: BND, Fraction: 2, NumericCode: "096", Grapheme: "$", Template: "$1"},
	BOB: {Decimal: ".", Thousand: ",", Code: BOB, Fraction: 2, NumericCode: "068", Grapheme: "Bs.", Template: "$1"},
	BRL: {Decimal: ",", Thousand: ".", Code: BRL, Fraction: 2, NumericCode: "986", Grapheme: "R$", Template: "$1"},
	BSD: {Decimal: ".", Thousand: ",", Code: BSD, Fraction: 2, NumericCode: "044", Grapheme: "$", Template: "$1"},
	BTN: {Decimal: ".", Thousand: ",", Code: BTN, Fraction: 2, NumericCode: "064", Grapheme: "Nu.", Template: "1$"},
	BWP: {Decimal: ".", Thousand: ",", Code: BWP, Fraction: 2, NumericCode: "072", Grapheme: "P", Template: "$1"},
	BYN: {Decimal: ",", Thousand: " ", Code: BYN, Fraction: 2, NumericCode: "933", Grapheme: "p.", Template: "1 $"},
	BYR: {Decimal: ",", Thousand: " ", Code: BYR, Fraction: 0, NumericCode: "", Grapheme: "p.", Template: "1 $"},
	BZD: {Decimal: ".", Thousand: ",", Code: BZD, Fraction: 2, NumericCode: "084", Grapheme: "BZ$", Template: "$1"},
	CAD: {Decimal: ".", Thousand: ",", Code: CAD, Fraction: 2, NumericCode: "124", Grapheme: "$", Template: "$1"},
	CDF: {Decimal: ".", Thousand: ",", Code: CDF, Fraction: 2, NumericCode: "976", Grapheme: "FC", Template: "1$"},
	CHF: {Decimal: ".", Thousand: ",", Code: CHF, Fraction: 2, NumericCode: "756", Grapheme: "CHF", Template: "1 $"},
	CLF: {Decimal: ",", Thousand: ".", Code: CLF, Fraction: 4, NumericCode: "990", Grapheme: "UF", Template: "$1"},
	CLP: {Decimal: ",", Thousand: ".", Code: CLP, Fraction: 0, NumericCode: "152", Grapheme: "$", Template: "$1"},
	CNY: {Decimal: ".", Thousand: ",", Code: CNY, Fraction: 2, NumericCode: "156", Grapheme: "\u5143", Template: "1 $"},
	COP: {Decimal: ",", Thousand: ".", Code: COP, Fraction: 2, NumericCode: "170", Grapheme: "$", Template: "$1"},
	CRC: {Decimal: ".", Thousand: ",", Code: CRC, Fraction: 2, NumericCode: "188", Grapheme: "\u20a1", Template: "$1"},
	CUC: {Decimal: ".", Thousand: ",", Code: CUC, Fraction: 2, NumericCode: "931", Grapheme: "$", Template: "1$"},
	CUP: {Decimal: ".", Thousand: ",", Code: CUP, Fraction: 2, NumericCode: "192", Grapheme: "$MN", Template: "$1"},
	CVE: {Decimal: ".", Thousand: ",", Code: CVE, Fraction: 2, NumericCode: "132", Grapheme: "$", Template: "1$"},
	CZK: {Decimal: ".", Thousand: ",", Code: CZK, Fraction: 2, NumericCode: "203", Grapheme: "K\u010d", Template: "1 $"},
	DJF: {Decimal: ".", Thousand: ",", Code: DJF, Fraction: 0, NumericCode: "262", Grapheme: "Fdj", Template: "1 $"},
	DKK: {Decimal: ",", Thousand: ".", Code: DKK, Fraction: 2, NumericCode: "208", Grapheme: "kr", Template: "$ 1"},
	DOP: {Decimal: ".", Thousand: ",", Code: DOP, Fraction: 2, NumericCode: "214", Grapheme: "RD$", Template: "$1"},
	DZD: {Decimal: ".", Thousand: ",", Code: DZD, Fraction: 2, NumericCode: "012", Grapheme: ".\u062f.\u062c", Template: "1 $"},
	EEK: {Decimal: ".", Thousand: ",", Code: EEK, Fraction: 2, NumericCode: "", Grapheme: "kr", Template: "$1"},
	EGP: {Decimal: ".", Thousand: ",", Code: EGP, Fraction: 2, NumericCode: "818", Grapheme: "\u00a3", Template: "$1"},
	ERN: {Decimal: ".", Thousand: ",", Code: ERN, Fraction: 2, NumericCode: "232", Grapheme: "Nfk", Template: "1 $"},
	ETB: {Decimal: ".", Thousand: ",", Code: ETB, Fraction: 2, NumericCode: "230", Grapheme: "Br", Template: "1 $"},
	EUR: {Decimal: ".", Thousand: ",", Code: EUR, Fraction: 2, NumericCode: "978", Grapheme: "\u20ac", Template: "$1"},
	FJD: {Decimal: ".", Thousand: ",", Code: FJD, Fraction: 2, NumericCode: "242", Grapheme: "$", Template: "$1"},
	FKP: {Decimal: ".", Thousand: ",", Code: FKP, Fraction: 2, NumericCode: "238", Grapheme: "\u00a3", Template: "$1"},
	GBP: {Decimal: ".", Thousand: ",", Code: GBP, Fraction: 2, NumericCode: "826", Grapheme: "\u00a3", Template: "$1"},
	GEL: {Decimal: ".", Thousand: ",", Code: GEL, Fraction: 2, NumericCode: "981", Grapheme: "\u10da", Template: "1 $"},
	GGP: {Decimal: ".", Thousand: ",", Code: GGP, Fraction: 2, NumericCode: "", Grapheme: "\u00a3", Template: "$1"},
	GHC: {Decimal: ".", Thousand: ",", Code: GHC, Fraction: 2, NumericCode: "", Grapheme: "\u00a2", Template: "$1"},
	GHS: {Decimal: ".", Thousand: ",", Code: GHS, Fraction: 2, NumericCode: "936", Grapheme: "\u20b5", Template: "$1"},
	GIP: {Decimal: ".", Thousand: ",", Code: GIP, Fraction: 2, NumericCode: "292", Grapheme: "\u00a3", Template: "$1"},
	GMD: {Decimal: ".", Thousand: ",", Code: GMD, Fraction: 2, NumericCode: "270", Grapheme: "D", Template: "1 $"},
	GNF: {Decimal: ".", Thousand: ",", Code: GNF, Fraction: 0, NumericCode: "324", Grapheme: "FG", Template: "1 $"},
	GTQ: {Decimal: ".", Thousand: ",", Code: GTQ, Fraction: 2, NumericCode: "320", Grapheme: "Q", Template: "$1"},
	GYD: {Decimal: ".", Thousand: ",", Code: GYD, Fraction: 2, NumericCode: "328", Grapheme: "$", Template: "$1"},
	HKD: {Decimal: ".", Thousand: ",", Code: HKD, Fraction: 2, NumericCode: "344", Grapheme: "HK$", Template: "$1"},
	HNL: {Decimal: ".", Thousand: ",", Code: HNL, Fraction: 2, NumericCode: "340", Grapheme: "L", Template: "$1"},
	HRK: {Decimal: ",", Thousand: ".", Code: HRK, Fraction: 2, NumericCode: "191", Grapheme: "kn", Template: "1 $"},
	HTG: {Decimal: ",", Thousand: ".", Code: HTG, Fraction: 2, NumericCode: "332", Grapheme: "G", Template: "1 $"},
	HUF: {Decimal: ",", Thousand: ".", Code: HUF, Fraction: 2, NumericCode: "348", Grapheme: "Ft", Template: "1 $"},
	IDR: {Decimal: ",", Thousand: ".", Code: IDR, Fraction: 2, NumericCode: "360", Grapheme: "Rp", Template: "$1"},
	ILS: {Decimal: ".", Thousand: ",", Code: ILS, Fraction: 2, NumericCode: "376", Grapheme: "\u20aa", Template: "$1"},
	IMP: {Decimal: ".", Thousand: ",", Code: IMP, Fraction: 2, NumericCode: "", Grapheme: "\u00a3", Template: "$1"},
	INR: {Decimal: ".", Thousand: ",", Code: INR, Fraction: 2, NumericCode: "356", Grapheme: "\u20b9", Template: "$1"},
	IQD: {Decimal: ".", Thousand: ",", Code: IQD, Fraction: 3, NumericCode: "368", Grapheme: ".\u062f.\u0639", Template: "1 $"},
	IRR: {Decimal: ".", Thousand: ",", Code: IRR, Fraction: 2, NumericCode: "364", Grapheme: "\ufdfc", Template: "1 $"},
	ISK: {Decimal: ",", Thousand: ".", Code: ISK, Fraction: 0, NumericCode: "352", Grapheme: "kr", Template: "$1"},
	JEP: {Decimal: ".", Thousand: ",", Code: JEP, Fraction: 2, NumericCode: "", Grapheme: "\u00a3", Template: "$1"},
	JMD: {Decimal: ".", Thousand: ",", Code: JMD, Fraction: 2, NumericCode: "388", Grapheme: "J$", Template: "$1"},
	JOD: {Decimal: ".", Thousand: ",", Code: JOD, Fraction: 3, NumericCode: "400", Grapheme: ".\u062f.\u0625", Template: "1 $"},
	JPY: {Decimal: ".", Thousand: ",", Code: JPY, Fraction: 0, NumericCode: "392", Grapheme: "\u00a5", Template: "$1"},
	KES: {Decimal: ".", Thousand: ",", Code: KES, Fraction: 2, NumericCode: "404", Grapheme: "KSh", Template: "$1"},
	KGS: {Decimal: ".", Thousand: ",", Code: KGS, Fraction: 2, NumericCode: "417", Grapheme: "\u0441\u043e\u043c", Template: "1 $"},
	KHR: {Decimal: ".", Thousand: ",", Code: KHR, Fraction: 2, NumericCode: "116", Grapheme: "\u17db", Template: "$1"},
	KMF: {Decimal: ".", Thousand: ",", Code: KMF, Fraction: 0, NumericCode: "174", Grapheme: "CF", Template: "$1"},
	KPW: {Decimal: ".", Thousand: ",", Code: KPW, Fraction: 2, NumericCode: "408", Grapheme: "\u20a9", Template: "$1"},
	KRW: {Decimal: ".", Thousand: ",", Code: KRW, Fraction: 0, NumericCode: "410", Grapheme: "\u20a9", Template: "$1"},
	KWD: {Decimal: ".", Thousand: ",", Code: KWD, Fraction: 3, NumericCode: "414", Grapheme: ".\u062f.\u0643", Template: "1 $"},
	KYD: {Decimal: ".", Thousand: ",", Code: KYD, Fraction: 2, NumericCode: "136", Grapheme: "$", Template: "$1"},
	KZT: {Decimal: ".", Thousand: ",", Code: KZT, Fraction: 2, NumericCode: "398", Grapheme: "\u20b8", Template: "$1"},
	LAK: {Decimal: ".", Thousand: ",", Code: LAK, Fraction: 2, NumericCode: "418", Grapheme: "\u20ad", Template: "$1"},
	LBP: {Decimal: ".", Thousand: ",", Code: LBP, Fraction: 2, NumericCode: "422", Grapheme: "\u00a3", Template: "$1"},
	LKR: {Decimal: ".", Thousand: ",", Code: LKR, Fraction: 2, NumericCode: "144", Grapheme: "\u20a8", Template: "$1"},
	LRD: {Decimal: ".", Thousand: ",", Code: LRD, Fraction: 2, NumericCode: "430", Grapheme: "$", Template: "$1"},
	LSL: {Decimal: ".", Thousand: ",", Code: LSL, Fraction: 2, NumericCode: "426", Grapheme: "L", Template: "$1"},
	LTL: {Decimal: ".", Thousand: ",", Code: LTL, Fraction: 2, NumericCode: "", Grapheme: "Lt", Template: "$1"},
	LVL: {Decimal: ".", Thousand: ",", Code: LVL, Fraction: 2, NumericCode: "", Grapheme: "Ls", Template: "1 $"},
	LYD: {Decimal: ".", Thousand: ",", Code: LYD, Fraction: 3, NumericCode: "434", Grapheme: ".\u062f.\u0644", Template: "1 $"},
	MAD: {Decimal: ".", Thousand: ",", Code: MAD, Fraction: 2, NumericCode: "504", Grapheme: ".\u062f.\u0645", Template: "1 $"},
	MDL: {Decimal: ".", Thousand: ",", Code: MDL, Fraction: 2, NumericCode: "498", Grapheme: "lei", Template: "1 $"},
	MGA: {Decimal: ".", Thousand: ",", Code: MGA, Fraction: 2, NumericCode: "969", Grapheme: "Ar", Template: "1$"},
	MKD: {Decimal: ".", Thousand: ",", Code: MKD, Fraction: 2, NumericCode: "807", Grapheme: "\u0434\u0435\u043d", Template: "$1"},
	MMK: {Decimal: ".", Thousand: ",", Code: MMK, Fraction: 2, NumericCode: "104", Grapheme: "K", Template: "$1"},
	MNT: {Decimal: ".", Thousand: ",", Code: MNT, Fraction: 2, NumericCode: "496", Grapheme: "\u20ae", Template: "$1"},
	MOP: {Decimal: ".", Thousand: ",", Code: MOP, Fraction: 2, NumericCode: "446", Grapheme: "P", Template: "1 $"},
	MRU: {Decimal: ".", Thousand: ",", Code: MRU, Fraction: 2, NumericCode: "929", Grapheme: "UM", Template: "$1"},
	MUR: {Decimal: ".", Thousand: ",", Code: MUR, Fraction: 2, NumericCode: "480", Grapheme: "\u20a8", Template: "$1"},
	MVR: {Decimal: ".", Thousand: ",", Code: MVR, Fraction: 2, NumericCode: "462", Grapheme: "MVR", Template: "1 $"},
	MWK: {Decimal: ".", Thousand: ",", Code: MWK, Fraction: 2, NumericCode: "454", Grapheme: "MK", Template: "$1"},
	MXN: {Decimal: ".", Thousand: ",", Code: MXN, Fraction: 2, NumericCode: "484", Grapheme: "$", Template: "$1"},
	MYR: {Decimal: ".", Thousand: ",", Code: MYR, Fraction: 2, NumericCode: "458", Grapheme: "RM", Template: "$1"},
	MZN: {Decimal: ".", Thousand: ",", Code: MZN, Fraction: 2, NumericCode: "943", Grapheme: "MT", Template: "$1"},
	NAD: {Decimal: ".", Thousand: ",", Code: NAD, Fraction: 2, NumericCode: "516", Grapheme: "$", Template: "$1"},
	NGN: {Decimal: ".", Thousand: ",", Code: NGN, Fraction: 2, NumericCode: "566", Grapheme: "\u20a6", Template: "$1"},
	NIO: {Decimal: ".", Thousand: ",", Code: NIO, Fraction: 2, NumericCode: "558", Grapheme: "C$", Template: "$1"},
	NOK: {Decimal: ".", Thousand: ",", Code: NOK, Fraction: 2, NumericCode: "578", Grapheme: "kr", Template: "1 $"},
	NPR: {Decimal: ".", Thousand: ",", Code: NPR, Fraction: 2, NumericCode: "524", Grapheme: "\u20a8", Template: "$1"},
	NZD: {Decimal: ".", Thousand: ",", Code: NZD, Fraction: 2, NumericCode: "554", Grapheme: "$", Template: "$1"},
	OMR: {Decimal: ".", Thousand: ",", Code: OMR, Fraction: 3, NumericCode: "512", Grapheme: "\ufdfc", Template: "1 $"},
	PAB: {Decimal: ".", Thousand: ",", Code: PAB, Fraction: 2, NumericCode: "590", Grapheme: "B/.", Template: "$1"},
	PEN: {Decimal: ".", Thousand: ",", Code: PEN, Fraction: 2, NumericCode: "604", Grapheme: "S/", Template: "$1"},
	PGK: {Decimal: ".", Thousand: ",", Code: PGK, Fraction: 2, NumericCode: "598", Grapheme: "K", Template: "1 $"},
	PHP: {Decimal: ".", Thousand: ",", Code: PHP, Fraction: 2, NumericCode: "608", Grapheme: "\u20b1", Template: "$1"},
	PKR: {Decimal: ".", Thousand: ",", Code: PKR, Fraction: 2, NumericCode: "586", Grapheme: "\u20a8", Template: "$1"},
	PLN: {Decimal: ".", Thousand: ",", Code: PLN, Fraction: 2, NumericCode: "985", Grapheme: "z\u0142", Template: "1 $"},
	PYG: {Decimal: ".", Thousand: ",", Code: PYG, Fraction: 0, NumericCode: "600", Grapheme: "Gs", Template: "1$"},
	QAR: {Decimal: ".", Thousand: ",", Code: QAR, Fraction: 2, NumericCode: "634", Grapheme: "\ufdfc", Template: "1 $"},
	RON: {Decimal: ".", Thousand: ",", Code: RON, Fraction: 2, NumericCode: "946", Grapheme: "lei", Template: "$1"},
	RSD: {Decimal: ".", Thousand: ",", Code: RSD, Fraction: 2, NumericCode: "941", Grapheme: "\u0414\u0438\u043d.", Template: "$1"},
	RUB: {Decimal: ".", Thousand: ",", Code: RUB, Fraction: 2, NumericCode: "643", Grapheme: "\u20bd", Template: "1 $"},
	RUR: {Decimal: ".", Thousand: ",", Code: RUR, Fraction: 2, NumericCode: "", Grapheme: "\u20bd", Template: "1 $"},
	RWF: {Decimal: ".", Thousand: ",", Code: RWF, Fraction: 0, NumericCode: "646", Grapheme: "FRw", Template: "1 $"},
	SAR: {Decimal: ".", Thousand: ",", Code: SAR, Fraction: 2, NumericCode: "682", Grapheme: "\ufdfc", Template: "1 $"},
	SBD: {Decimal: ".", Thousand: ",", Code: SBD, Fraction: 2, NumericCode: "090", Grapheme: "$", Template: "$1"},
	SCR: {Decimal: ".", Thousand: ",", Code: SCR, Fraction: 2, NumericCode: "690", Grapheme: "\u20a8", Template: "$1"},
	SDG: {Decimal: ".", Thousand: ",", Code: SDG, Fraction: 2, NumericCode: "938", Grapheme: "\u00a3", Template: "$1"},
	SEK: {Decimal: ".", Thousand: ",", Code: SEK, Fraction: 2, NumericCode: "752", Grapheme: "kr", Template: "1 $"},
	SGD: {Decimal: ".", Thousand: ",", Code: SGD, Fraction: 2, NumericCode: "702", Grapheme: "S$", Template: "$1"},
	SHP: {Decimal: ".", Thousand: ",", Code: SHP, Fraction: 2, NumericCode: "654", Grapheme: "\u00a3", Template: "$1"},
	SKK: {Decimal: ".", Thousand: ",", Code: SKK, Fraction: 2, NumericCode: "", Grapheme: "Sk", Template: "$1"},
	SLE: {Decimal: ".", Thousand: ",", Code: SLE, Fraction: 2, NumericCode: "925", Grapheme: "Le", Template: "1 $"},
	SLL: {Decimal: ".", Thousand: ",", Code: SLL, Fraction: 2, NumericCode: "694", Grapheme: "Le", Template: "1 $"},
	SOS: {Decimal: ".", Thousand: ",", Code: SOS, Fraction: 2, NumericCode: "706", Grapheme: "Sh", Template: "1 $"},
	SRD: {Decimal: ".", Thousand: ",", Code: SRD, Fraction: 2, NumericCode: "968", Grapheme: "$", Template: "$1"},
	SSP: {Decimal: ".", Thousand: ",", Code: SSP, Fraction: 2, NumericCode: "728", Grapheme: "\u00a3", Template: "1 $"},
	STD: {Decimal: ".", Thousand: ",", Code: STD, Fraction: 2, NumericCode: "", Grapheme: "Db", Template: "1 $"},
	STN: {Decimal: ".", Thousand: ",", Code: STN, Fraction: 2, NumericCode: "930", Grapheme: "Db", Template: "1 $"},
	SVC: {Decimal: ".", Thousand: ",", Code: SVC, Fraction: 2, NumericCode: "222", Grapheme: "\u20a1", Template: "$1"},
	SYP: {Decimal: ".", Thousand: ",", Code: SYP, Fraction: 2, NumericCode: "760", Grapheme: "\u00a3", Template: "1 $"},
	SZL: {Decimal: ".", Thousand: ",", Code: SZL, Fraction: 2, NumericCode: "748", Grapheme: "\u00a3", Template: "$1"},
	THB: {Decimal: ".", Thousand: ",", Code: THB, Fraction: 2, NumericCode: "764", Grapheme: "\u0e3f", Template: "$1"},
	TJS: {Decimal: ".", Thousand: ",", Code: TJS, Fraction: 2, NumericCode: "972", Grapheme: "SM", Template: "1 $"},
	TMT: {Decimal: ".", Thousand: ",", Code: TMT, Fraction: 2, NumericCode: "934", Grapheme: "T", Template: "1 $"},
	TND: {Decimal: ".", Thousand: ",", Code: TND, Fraction: 3, NumericCode: "788", Grapheme: ".\u062f.\u062a", Template: "1 $"},
	TOP: {Decimal: ".", Thousand: ",", Code: TOP, Fraction: 2, NumericCode: "776", Grapheme: "T$", Template: "$1"},
	TRL: {Decimal: ".", Thousand: ",", Code: TRL, Fraction: 2, NumericCode: "", Grapheme: "\u20a4", Template: "$1"},
	TRY: {Decimal: ".", Thousand: ",", Code: TRY, Fraction: 2, NumericCode: "949", Grapheme: "\u20ba", Template: "$1"},
	TTD: {Decimal: ".", Thousand: ",", Code: TTD, Fraction: 2, NumericCode: "780", Grapheme: "TT$", Template: "$1"},
	TWD: {Decimal: ".", Thousand: ",", Code: TWD, Fraction: 2, NumericCode: "901", Grapheme: "NT$", Template: "$1"},
	TZS: {Decimal: ".", Thousand: ",", Code: TZS, Fraction: 2, NumericCode: "834", Grapheme: "TSh", Template: "$1"},
	UAH: {Decimal: ".", Thousand: ",", Code: UAH, Fraction: 2, NumericCode: "980", Grapheme: "\u20b4", Template: "1 $"},
	UGX: {Decimal: ".", Thousand: ",", Code: UGX, Fraction: 0, NumericCode: "800", Grapheme: "USh", Template: "1 $"},
	USD: {Decimal: ".", Thousand: ",", Code: USD, Fraction: 2, NumericCode: "840", Grapheme: "$", Template: "$1"},
	UYU: {Decimal: ".", Thousand: ",", Code: UYU, Fraction: 2, NumericCode: "858", Grapheme: "$U", Template: "$1"},
	UZS: {Decimal: ".", Thousand: ",", Code: UZS, Fraction: 2, NumericCode: "860", Grapheme: "so\u2019m", Template: "$1"},
	VEF: {Decimal: ".", Thousand: ",", Code: VEF, Fraction: 2, NumericCode: "937", Grapheme: "Bs", Template: "$1"},
	VES: {Decimal: ".", Thousand: ",", Code: VES, Fraction: 2, NumericCode: "928", Grapheme: "Bs.S", Template: "$1"},
	VND: {Decimal: ".", Thousand: ",", Code: VND, Fraction: 0, NumericCode: "704", Grapheme: "\u20ab", Template: "1 $"},
	VUV: {Decimal: ".", Thousand: ",", Code: VUV, Fraction: 0, NumericCode: "548", Grapheme: "Vt", Template: "$1"},
	WST: {Decimal: ".", Thousand: ",", Code: WST, Fraction: 2, NumericCode: "882", Grapheme: "T", Template: "1 $"},
	XAF: {Decimal: ".", Thousand: ",", Code: XAF, Fraction: 0, NumericCode: "950", Grapheme: "Fr", Template: "1 $"},
	XAG: {Decimal: ".", Thousand: ",", Code: XAG, Fraction: 0, NumericCode: "961", Grapheme: "oz t", Template: "1 $"},
	XAU: {Decimal: ".", Thousand: ",", Code: XAU, Fraction: 0, NumericCode: "959", Grapheme: "oz t", Template: "1 $"},
	XCD: {Decimal: ".", Thousand: ",", Code: XCD, Fraction: 2, NumericCode: "951", Grapheme: "$", Template: "$1"},
	XCG: {Decimal: ",", Thousand: ".", Code: XCG, Fraction: 2, NumericCode: "532", Grapheme: "Cg", Template: "$1"},
	XDR: {Decimal: ".", Thousand: ",", Code: XDR, Fraction: 0, NumericCode: "960", Grapheme: "SDR", Template: "1 $"},
	XOF: {Decimal: ".", Thousand: ",", Code: XOF, Fraction: 0, NumericCode: "952", Grapheme: "CFA", Template: "1 $"},
	XPF: {Decimal: ".", Thousand: ",", Code: XPF, Fraction: 0, NumericCode: "953", Grapheme: "₣", Template: "1 $"},
	YER: {Decimal: ".", Thousand: ",", Code: YER, Fraction: 2, NumericCode: "886", Grapheme: "\ufdfc", Template: "1 $"},
	ZAR: {Decimal: ".", Thousand: ",", Code: ZAR, Fraction: 2, NumericCode: "710", Grapheme: "R", Template: "$1"},
	ZMW: {Decimal: ".", Thousand: ",", Code: ZMW, Fraction: 2, NumericCode: "967", Grapheme: "ZK", Template: "$1"},
	ZWD: {Decimal: ".", Thousand: ",", Code: ZWD, Fraction: 2, NumericCode: "716", Grapheme: "Z$", Template: "$1"},
	ZWL: {Decimal: ".", Thousand: ",", Code: ZWL, Fraction: 2, NumericCode: "932", Grapheme: "Z$", Template: "$1"},
}

// AddCurrency creates and registers a new custom currency with the specified parameters.
// This allows you to work with cryptocurrencies, loyalty points, or other custom units.
//
// Parameters:
//   - code: Currency code (e.g., "BTC", "POINTS")
//   - grapheme: Currency symbol (e.g., "₿", "pts")
//   - template: Format template ("$1" puts symbol before, "1 $" puts symbol after)
//   - decimal: Decimal separator ("." or ",")
//   - thousand: Thousands separator ("," or "." or "")
//   - fraction: Number of decimal places
//
// Returns the newly created Currency.
//
// Example:
//
//	btc := moneykit.AddCurrency("BTC", "₿", "₿1", ".", ",", 8)
//	bitcoin := moneykit.New(100000000, "BTC") // 1.00000000 BTC
//	fmt.Println(bitcoin.Display()) // ₿1.00000000
func AddCurrency(code, grapheme, template, decimal, thousand string, fraction int) *Currency {
	c := Currency{
		Code:     code,
		Grapheme: grapheme,
		Template: template,
		Decimal:  decimal,
		Thousand: thousand,
		Fraction: fraction,
	}
	currencies.Add(&c)
	return &c
}

func newCurrency(code string) *Currency {
	return &Currency{Code: strings.ToUpper(code)}
}

// GetCurrency returns the Currency for the given currency code.
// If the currency is not registered, it returns a default currency with basic formatting.
//
// Parameters:
//   - code: ISO 4217 currency code (case-insensitive)
//
// Example:
//
//	usd := moneykit.GetCurrency("USD")
//	eur := moneykit.GetCurrency("eur") // Case-insensitive
//	custom := moneykit.GetCurrency("XYZ") // Returns default if not found
func GetCurrency(code string) *Currency {
	return currencies.CurrencyByCode(strings.ToUpper(code))
}

// GetCurrencyByNumericCode returns the Currency for the given ISO 4217 numeric code.
// Returns nil if the currency is not found.
//
// Parameters:
//   - code: 3-digit numeric code as string (e.g., "840" for USD, "978" for EUR)
//
// Example:
//
//	usd := moneykit.GetCurrencyByNumericCode("840") // USD
//	eur := moneykit.GetCurrencyByNumericCode("978") // EUR
func GetCurrencyByNumericCode(code string) *Currency {
	return currencies.CurrencyByNumericCode(code)
}

// Formatter returns a Formatter instance configured with this currency's formatting rules.
// The formatter can be used for custom formatting operations.
//
// Example:
//
//	currency := moneykit.GetCurrency("USD")
//	formatter := currency.Formatter()
//	formatted := formatter.Format(123456) // $1,234.56
func (c *Currency) Formatter() *Formatter {
	return &Formatter{
		Fraction: c.Fraction,
		Decimal:  c.Decimal,
		Thousand: c.Thousand,
		Grapheme: c.Grapheme,
		Template: c.Template,
	}
}

// getDefault represent default currency if currency is not found in currencies list.
// Grapheme and Code fields will be changed by currency code.
func (c *Currency) getDefault() *Currency {
	return &Currency{Decimal: ".", Thousand: ",", Code: c.Code, Fraction: 2, Grapheme: c.Code, Template: "1$"}
}

// get extended currency using currencies list.
func (c *Currency) get() *Currency {
	if curr, ok := currencies[c.Code]; ok {
		return curr
	}

	return c.getDefault()
}

func (c *Currency) equals(oc *Currency) bool {
	return c.Code == oc.Code
}
