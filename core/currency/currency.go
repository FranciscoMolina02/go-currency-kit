package currency

import (
	"strings"
)

// Currency represents money currency information required for formatting.
type Currency struct {
	Code          string
	DecimalAmount int
}

type Currencies map[string]*Currency

// CurrencyByCode returns the currency given the currency code defined as a constant.
func (c Currencies) CurrencyByCode(code string) *Currency {
	sc, ok := c[code]
	if !ok {
		return nil
	}

	return sc
}

// Add updates currencies list by adding a given Currency to it.
func (c Currencies) Add(currency *Currency) Currencies {
	c[currency.Code] = currency
	return c
}

// currencies represents a collection of currency.
var currencies = Currencies{

	AED: {Code: AED, DecimalAmount: 2},
	AFN: {Code: AFN, DecimalAmount: 2},
	ALL: {Code: ALL, DecimalAmount: 2},
	AMD: {Code: AMD, DecimalAmount: 2},
	ANG: {Code: ANG, DecimalAmount: 2},
	AOA: {Code: AOA, DecimalAmount: 2},
	ARS: {Code: ARS, DecimalAmount: 2},
	AUD: {Code: AUD, DecimalAmount: 2},
	AWG: {Code: AWG, DecimalAmount: 2},
	AZN: {Code: AZN, DecimalAmount: 2},
	BAM: {Code: BAM, DecimalAmount: 2},
	BBD: {Code: BBD, DecimalAmount: 2},
	BDT: {Code: BDT, DecimalAmount: 2},
	BGN: {Code: BGN, DecimalAmount: 2},
	BHD: {Code: BHD, DecimalAmount: 3},
	BIF: {Code: BIF, DecimalAmount: 0},
	BMD: {Code: BMD, DecimalAmount: 2},
	BND: {Code: BND, DecimalAmount: 2},
	BOB: {Code: BOB, DecimalAmount: 2},
	BRL: {Code: BRL, DecimalAmount: 2},
	BSD: {Code: BSD, DecimalAmount: 2},
	BTN: {Code: BTN, DecimalAmount: 2},
	BWP: {Code: BWP, DecimalAmount: 2},
	BYN: {Code: BYN, DecimalAmount: 2},
	BYR: {Code: BYR, DecimalAmount: 0},
	BZD: {Code: BZD, DecimalAmount: 2},
	CAD: {Code: CAD, DecimalAmount: 2},
	CDF: {Code: CDF, DecimalAmount: 2},
	CHF: {Code: CHF, DecimalAmount: 2},
	CLF: {Code: CLF, DecimalAmount: 4},
	CLP: {Code: CLP, DecimalAmount: 0},
	CNY: {Code: CNY, DecimalAmount: 2},
	COP: {Code: COP, DecimalAmount: 2},
	CRC: {Code: CRC, DecimalAmount: 2},
	CUC: {Code: CUC, DecimalAmount: 2},
	CUP: {Code: CUP, DecimalAmount: 2},
	CVE: {Code: CVE, DecimalAmount: 2},
	CZK: {Code: CZK, DecimalAmount: 2},
	DJF: {Code: DJF, DecimalAmount: 0},
	DKK: {Code: DKK, DecimalAmount: 2},
	DOP: {Code: DOP, DecimalAmount: 2},
	DZD: {Code: DZD, DecimalAmount: 2},
	EEK: {Code: EEK, DecimalAmount: 2},
	EGP: {Code: EGP, DecimalAmount: 2},
	ERN: {Code: ERN, DecimalAmount: 2},
	ETB: {Code: ETB, DecimalAmount: 2},
	EUR: {Code: EUR, DecimalAmount: 2},
	FJD: {Code: FJD, DecimalAmount: 2},
	FKP: {Code: FKP, DecimalAmount: 2},
	GBP: {Code: GBP, DecimalAmount: 2},
	GEL: {Code: GEL, DecimalAmount: 2},
	GGP: {Code: GGP, DecimalAmount: 2},
	GHC: {Code: GHC, DecimalAmount: 2},
	GHS: {Code: GHS, DecimalAmount: 2},
	GIP: {Code: GIP, DecimalAmount: 2},
	GMD: {Code: GMD, DecimalAmount: 2},
	GNF: {Code: GNF, DecimalAmount: 0},
	GTQ: {Code: GTQ, DecimalAmount: 2},
	GYD: {Code: GYD, DecimalAmount: 2},
	HKD: {Code: HKD, DecimalAmount: 2},
	HNL: {Code: HNL, DecimalAmount: 2},
	HRK: {Code: HRK, DecimalAmount: 2},
	HTG: {Code: HTG, DecimalAmount: 2},
	HUF: {Code: HUF, DecimalAmount: 2},
	IDR: {Code: IDR, DecimalAmount: 2},
	ILS: {Code: ILS, DecimalAmount: 2},
	IMP: {Code: IMP, DecimalAmount: 2},
	INR: {Code: INR, DecimalAmount: 2},
	IQD: {Code: IQD, DecimalAmount: 3},
	IRR: {Code: IRR, DecimalAmount: 2},
	ISK: {Code: ISK, DecimalAmount: 0},
	JEP: {Code: JEP, DecimalAmount: 2},
	JMD: {Code: JMD, DecimalAmount: 2},
	JOD: {Code: JOD, DecimalAmount: 3},
	JPY: {Code: JPY, DecimalAmount: 0},
	KES: {Code: KES, DecimalAmount: 2},
	KGS: {Code: KGS, DecimalAmount: 2},
	KHR: {Code: KHR, DecimalAmount: 2},
	KMF: {Code: KMF, DecimalAmount: 0},
	KPW: {Code: KPW, DecimalAmount: 2},
	KRW: {Code: KRW, DecimalAmount: 0},
	KWD: {Code: KWD, DecimalAmount: 3},
	KYD: {Code: KYD, DecimalAmount: 2},
	KZT: {Code: KZT, DecimalAmount: 2},
	LAK: {Code: LAK, DecimalAmount: 2},
	LBP: {Code: LBP, DecimalAmount: 2},
	LKR: {Code: LKR, DecimalAmount: 2},
	LRD: {Code: LRD, DecimalAmount: 2},
	LSL: {Code: LSL, DecimalAmount: 2},
	LTL: {Code: LTL, DecimalAmount: 2},
	LVL: {Code: LVL, DecimalAmount: 2},
	LYD: {Code: LYD, DecimalAmount: 3},
	MAD: {Code: MAD, DecimalAmount: 2},
	MDL: {Code: MDL, DecimalAmount: 2},
	MGA: {Code: MGA, DecimalAmount: 2},
	MKD: {Code: MKD, DecimalAmount: 2},
	MMK: {Code: MMK, DecimalAmount: 2},
	MNT: {Code: MNT, DecimalAmount: 2},
	MOP: {Code: MOP, DecimalAmount: 2},
	MUR: {Code: MUR, DecimalAmount: 2},
	MVR: {Code: MVR, DecimalAmount: 2},
	MWK: {Code: MWK, DecimalAmount: 2},
	MXN: {Code: MXN, DecimalAmount: 2},
	MYR: {Code: MYR, DecimalAmount: 2},
	MZN: {Code: MZN, DecimalAmount: 2},
	NAD: {Code: NAD, DecimalAmount: 2},
	NGN: {Code: NGN, DecimalAmount: 2},
	NIO: {Code: NIO, DecimalAmount: 2},
	NOK: {Code: NOK, DecimalAmount: 2},
	NPR: {Code: NPR, DecimalAmount: 2},
	NZD: {Code: NZD, DecimalAmount: 2},
	OMR: {Code: OMR, DecimalAmount: 3},
	PAB: {Code: PAB, DecimalAmount: 2},
	PEN: {Code: PEN, DecimalAmount: 2},
	PGK: {Code: PGK, DecimalAmount: 2},
	PHP: {Code: PHP, DecimalAmount: 2},
	PKR: {Code: PKR, DecimalAmount: 2},
	PLN: {Code: PLN, DecimalAmount: 2},
	PYG: {Code: PYG, DecimalAmount: 0},
	QAR: {Code: QAR, DecimalAmount: 2},
	RON: {Code: RON, DecimalAmount: 2},
	RSD: {Code: RSD, DecimalAmount: 2},
	RUB: {Code: RUB, DecimalAmount: 2},
	RUR: {Code: RUR, DecimalAmount: 2},
	RWF: {Code: RWF, DecimalAmount: 0},
	SAR: {Code: SAR, DecimalAmount: 2},
	SBD: {Code: SBD, DecimalAmount: 2},
	SCR: {Code: SCR, DecimalAmount: 2},
	SDG: {Code: SDG, DecimalAmount: 2},
	SEK: {Code: SEK, DecimalAmount: 2},
	SGD: {Code: SGD, DecimalAmount: 2},
	SHP: {Code: SHP, DecimalAmount: 2},
	SKK: {Code: SKK, DecimalAmount: 2},
	SLL: {Code: SLL, DecimalAmount: 2},
	SOS: {Code: SOS, DecimalAmount: 2},
	SRD: {Code: SRD, DecimalAmount: 2},
	SSP: {Code: SSP, DecimalAmount: 2},
	STD: {Code: STD, DecimalAmount: 2},
	SVC: {Code: SVC, DecimalAmount: 2},
	SYP: {Code: SYP, DecimalAmount: 2},
	SZL: {Code: SZL, DecimalAmount: 2},
	THB: {Code: THB, DecimalAmount: 2},
	TJS: {Code: TJS, DecimalAmount: 2},
	TMT: {Code: TMT, DecimalAmount: 2},
	TND: {Code: TND, DecimalAmount: 3},
	TOP: {Code: TOP, DecimalAmount: 2},
	TRL: {Code: TRL, DecimalAmount: 2},
	TRY: {Code: TRY, DecimalAmount: 2},
	TTD: {Code: TTD, DecimalAmount: 2},
	TWD: {Code: TWD, DecimalAmount: 2},
	TZS: {Code: TZS, DecimalAmount: 2},
	UAH: {Code: UAH, DecimalAmount: 2},
	UGX: {Code: UGX, DecimalAmount: 0},
	USD: {Code: USD, DecimalAmount: 2},
	UYU: {Code: UYU, DecimalAmount: 2},
	UZS: {Code: UZS, DecimalAmount: 2},
	VEF: {Code: VEF, DecimalAmount: 2},
	VND: {Code: VND, DecimalAmount: 0},
	VUV: {Code: VUV, DecimalAmount: 0},
	WST: {Code: WST, DecimalAmount: 2},
	XAF: {Code: XAF, DecimalAmount: 0},
	XAG: {Code: XAG, DecimalAmount: 0},
	XAU: {Code: XAU, DecimalAmount: 0},
	XCD: {Code: XCD, DecimalAmount: 2},
	XDR: {Code: XDR, DecimalAmount: 0},
	XOF: {Code: XOF, DecimalAmount: 0},
	XPF: {Code: XPF, DecimalAmount: 0},
	YER: {Code: YER, DecimalAmount: 2},
	ZAR: {Code: ZAR, DecimalAmount: 2},
	ZMW: {Code: ZMW, DecimalAmount: 2},
	ZWD: {Code: ZWD, DecimalAmount: 2},
	ZWL: {Code: ZWL, DecimalAmount: 2},
}

// AddCurrency lets you insert or update currency in currencies list.
func AddCurrency(code string, decimalAmount int) *Currency {
	c := Currency{
		Code:          code,
		DecimalAmount: decimalAmount,
	}
	currencies.Add(&c)
	return &c
}

func NewCurrency(code string) *Currency {
	return &Currency{Code: strings.ToUpper(code)}
}

// GetCurrency returns the currency given the code.
func GetCurrency(code string) *Currency {
	return currencies.CurrencyByCode(code)
}

// getDefault represent default currency if currency is not found in currencies list.
// Grapheme and Code fields will be changed by currency code.
func (c *Currency) getDefault() *Currency {
	return &Currency{Code: c.Code, DecimalAmount: 2}
}

// get extended currency using currencies list.
func (c *Currency) Get() *Currency {
	if curr, ok := currencies[c.Code]; ok {
		return curr
	}

	return c.getDefault()
}

func (c *Currency) Equals(oc *Currency) bool {
	return c.Code == oc.Code
}
