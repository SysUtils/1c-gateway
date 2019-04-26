package odata

type NumberQualifiers struct {
	AllowedSign    String `json:"AllowedSign,omitempty"`
	Digits         Int16  `json:"Digits,omitempty"`
	FractionDigits Int16  `json:"FractionDigits,omitempty"`
}
type TypeDescription struct {
	Types            []String         `json:"Types,omitempty"`
	NumberQualifiers NumberQualifiers `json:"NumberQualifiers,omitempty"`
}
type PointInTime struct {
	Ref  String   `json:"Ref,omitempty"`
	Date DateTime `json:"Date,omitempty"`
}
type AccumulationRegisterPartiiTovarovVProizvodstveRowType struct {
	Recorder                       String    `json:"Recorder,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Active                         *Boolean  `json:"Active,omitempty"`
	RecordType                     *String   `json:"RecordType,omitempty"`
	ProizvodstvennyiUchastokKey    *Guid     `json:"ПроизводственныйУчасток_Key,omitempty"`
	Nomenklatura                   *String   `json:"Номенклатура,omitempty"`
	OrganizatsiiaKey               *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentOprikhodovaniia        *String   `json:"ДокументОприходования,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	Cost                           *Double   `json:"Стоимость,omitempty"`
	StoimostUpr                    *Double   `json:"СтоимостьУпр,omitempty"`
	StoimostBezNDS                 *Double   `json:"СтоимостьБезНДС,omitempty"`
	OperationCode                  *String   `json:"КодОперации,omitempty"`
	SpisaniePartii                 *Boolean  `json:"СписаниеПартий,omitempty"`
	NomerKorStroki                 *Int64    `json:"НомерКорСтроки,omitempty"`
	RecorderType                   String    `json:"Recorder_Type,omitempty"`
	ItemType                       *String   `json:"Номенклатура_Type,omitempty"`
	DokumentOprikhodovaniiaType    *String   `json:"ДокументОприходования_Type,omitempty"`
}
type AccumulationRegisterPartiiTovarovVProizvodstveBalance struct {
	ProizvodstvennyiUchastokKey    *Guid   `json:"ПроизводственныйУчасток_Key,omitempty"`
	Nomenklatura                   *String `json:"Номенклатура,omitempty"`
	OrganizatsiiaKey               *Guid   `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey          *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentOprikhodovaniia        *String `json:"ДокументОприходования,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KolichestvoBalance             *Int64  `json:"КоличествоBalance,omitempty"`
	VesBalance                     *Double `json:"ВесBalance,omitempty"`
	StoimostBalance                *Double `json:"СтоимостьBalance,omitempty"`
	StoimostUprBalance             *Double `json:"СтоимостьУпрBalance,omitempty"`
	StoimostBezNDSBalance          *Double `json:"СтоимостьБезНДСBalance,omitempty"`
	ItemType                       *String `json:"Номенклатура_Type,omitempty"`
	DokumentOprikhodovaniiaType    *String `json:"ДокументОприходования_Type,omitempty"`
}
type AccumulationRegisterPartiiTovarovVProizvodstveTurnover struct {
	ProizvodstvennyiUchastokKey    *Guid     `json:"ПроизводственныйУчасток_Key,omitempty"`
	Nomenklatura                   *String   `json:"Номенклатура,omitempty"`
	OrganizatsiiaKey               *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentOprikhodovaniia        *String   `json:"ДокументОприходования,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	StoimostTurnover               *Double   `json:"СтоимостьTurnover,omitempty"`
	StoimostReceipt                *Double   `json:"СтоимостьReceipt,omitempty"`
	StoimostExpense                *Double   `json:"СтоимостьExpense,omitempty"`
	StoimostUprTurnover            *Double   `json:"СтоимостьУпрTurnover,omitempty"`
	StoimostUprReceipt             *Double   `json:"СтоимостьУпрReceipt,omitempty"`
	StoimostUprExpense             *Double   `json:"СтоимостьУпрExpense,omitempty"`
	StoimostBezNDSTurnover         *Double   `json:"СтоимостьБезНДСTurnover,omitempty"`
	StoimostBezNDSReceipt          *Double   `json:"СтоимостьБезНДСReceipt,omitempty"`
	StoimostBezNDSExpense          *Double   `json:"СтоимостьБезНДСExpense,omitempty"`
	ItemType                       *String   `json:"Номенклатура_Type,omitempty"`
	DokumentOprikhodovaniiaType    *String   `json:"ДокументОприходования_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterPartiiTovarovVProizvodstveBalanceAndTurnover struct {
	ProizvodstvennyiUchastokKey    *Guid     `json:"ПроизводственныйУчасток_Key,omitempty"`
	Nomenklatura                   *String   `json:"Номенклатура,omitempty"`
	OrganizatsiiaKey               *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentOprikhodovaniia        *String   `json:"ДокументОприходования,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoOpeningBalance      *Int64    `json:"КоличествоOpeningBalance,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	KolichestvoClosingBalance      *Int64    `json:"КоличествоClosingBalance,omitempty"`
	VesOpeningBalance              *Double   `json:"ВесOpeningBalance,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	VesClosingBalance              *Double   `json:"ВесClosingBalance,omitempty"`
	StoimostOpeningBalance         *Double   `json:"СтоимостьOpeningBalance,omitempty"`
	StoimostTurnover               *Double   `json:"СтоимостьTurnover,omitempty"`
	StoimostReceipt                *Double   `json:"СтоимостьReceipt,omitempty"`
	StoimostExpense                *Double   `json:"СтоимостьExpense,omitempty"`
	StoimostClosingBalance         *Double   `json:"СтоимостьClosingBalance,omitempty"`
	StoimostUprOpeningBalance      *Double   `json:"СтоимостьУпрOpeningBalance,omitempty"`
	StoimostUprTurnover            *Double   `json:"СтоимостьУпрTurnover,omitempty"`
	StoimostUprReceipt             *Double   `json:"СтоимостьУпрReceipt,omitempty"`
	StoimostUprExpense             *Double   `json:"СтоимостьУпрExpense,omitempty"`
	StoimostUprClosingBalance      *Double   `json:"СтоимостьУпрClosingBalance,omitempty"`
	StoimostBezNDSOpeningBalance   *Double   `json:"СтоимостьБезНДСOpeningBalance,omitempty"`
	StoimostBezNDSTurnover         *Double   `json:"СтоимостьБезНДСTurnover,omitempty"`
	StoimostBezNDSReceipt          *Double   `json:"СтоимостьБезНДСReceipt,omitempty"`
	StoimostBezNDSExpense          *Double   `json:"СтоимостьБезНДСExpense,omitempty"`
	StoimostBezNDSClosingBalance   *Double   `json:"СтоимостьБезНДСClosingBalance,omitempty"`
	ItemType                       *String   `json:"Номенклатура_Type,omitempty"`
	DokumentOprikhodovaniiaType    *String   `json:"ДокументОприходования_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRowType struct {
	Recorder               String    `json:"Recorder,omitempty"`
	Period                 *DateTime `json:"Period,omitempty"`
	LineNumber             Int64     `json:"LineNumber,omitempty"`
	Active                 *Boolean  `json:"Active,omitempty"`
	RecordType             *String   `json:"RecordType,omitempty"`
	FizLitsoKey            *Guid     `json:"ФизЛицо_Key,omitempty"`
	ValiutaKey             *Guid     `json:"Валюта_Key,omitempty"`
	RaschetnyiDokument     *String   `json:"РасчетныйДокумент,omitempty"`
	SummaVzaimoraschetov   *Double   `json:"СуммаВзаиморасчетов,omitempty"`
	SummaUpr               *Double   `json:"СуммаУпр,omitempty"`
	RecorderType           String    `json:"Recorder_Type,omitempty"`
	RaschetnyiDokumentType *String   `json:"РасчетныйДокумент_Type,omitempty"`
}
type AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalance struct {
	FizLitsoKey                 *Guid   `json:"ФизЛицо_Key,omitempty"`
	ValiutaKey                  *Guid   `json:"Валюта_Key,omitempty"`
	RaschetnyiDokument          *String `json:"РасчетныйДокумент,omitempty"`
	SummaVzaimoraschetovBalance *Double `json:"СуммаВзаиморасчетовBalance,omitempty"`
	SummaUprBalance             *Double `json:"СуммаУпрBalance,omitempty"`
	RaschetnyiDokumentType      *String `json:"РасчетныйДокумент_Type,omitempty"`
}
type AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiTurnover struct {
	FizLitsoKey                  *Guid     `json:"ФизЛицо_Key,omitempty"`
	ValiutaKey                   *Guid     `json:"Валюта_Key,omitempty"`
	RaschetnyiDokument           *String   `json:"РасчетныйДокумент,omitempty"`
	Period                       *DateTime `json:"Period,omitempty"`
	SecondPeriod                 *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                 *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                   *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                    *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                   *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                  *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod               *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                   *DateTime `json:"YearPeriod,omitempty"`
	Recorder                     *String   `json:"Recorder,omitempty"`
	LineNumber                   *Int64    `json:"LineNumber,omitempty"`
	SummaVzaimoraschetovTurnover *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt  *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense  *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SummaUprTurnover             *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt              *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense              *Double   `json:"СуммаУпрExpense,omitempty"`
	RaschetnyiDokumentType       *String   `json:"РасчетныйДокумент_Type,omitempty"`
	RecorderType                 *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalanceAndTurnover struct {
	FizLitsoKey                        *Guid     `json:"ФизЛицо_Key,omitempty"`
	ValiutaKey                         *Guid     `json:"Валюта_Key,omitempty"`
	RaschetnyiDokument                 *String   `json:"РасчетныйДокумент,omitempty"`
	Period                             *DateTime `json:"Period,omitempty"`
	SecondPeriod                       *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                       *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                         *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                          *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                         *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                      *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                        *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                      *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                     *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                         *DateTime `json:"YearPeriod,omitempty"`
	Recorder                           *String   `json:"Recorder,omitempty"`
	LineNumber                         *Int64    `json:"LineNumber,omitempty"`
	SummaVzaimoraschetovOpeningBalance *Double   `json:"СуммаВзаиморасчетовOpeningBalance,omitempty"`
	SummaVzaimoraschetovTurnover       *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt        *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense        *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SummaVzaimoraschetovClosingBalance *Double   `json:"СуммаВзаиморасчетовClosingBalance,omitempty"`
	SummaUprOpeningBalance             *Double   `json:"СуммаУпрOpeningBalance,omitempty"`
	SummaUprTurnover                   *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt                    *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense                    *Double   `json:"СуммаУпрExpense,omitempty"`
	SummaUprClosingBalance             *Double   `json:"СуммаУпрClosingBalance,omitempty"`
	RaschetnyiDokumentType             *String   `json:"РасчетныйДокумент_Type,omitempty"`
	RecorderType                       *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterVnutrennieZakazyRowType struct {
	Recorder                       String    `json:"Recorder,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Active                         *Boolean  `json:"Active,omitempty"`
	RecordType                     *String   `json:"RecordType,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	Zakazchik                      *String   `json:"Заказчик,omitempty"`
	VnutrenniiZakazKey             *Guid     `json:"ВнутреннийЗаказ_Key,omitempty"`
	StatusPartii                   *String   `json:"СтатусПартии,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	RecorderType                   String    `json:"Recorder_Type,omitempty"`
	ZakazchikType                  *String   `json:"Заказчик_Type,omitempty"`
}
type AccumulationRegisterVnutrennieZakazyBalance struct {
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	Zakazchik                      *String `json:"Заказчик,omitempty"`
	VnutrenniiZakazKey             *Guid   `json:"ВнутреннийЗаказ_Key,omitempty"`
	StatusPartii                   *String `json:"СтатусПартии,omitempty"`
	KolichestvoBalance             *Int64  `json:"КоличествоBalance,omitempty"`
	VesBalance                     *Double `json:"ВесBalance,omitempty"`
	ZakazchikType                  *String `json:"Заказчик_Type,omitempty"`
}
type AccumulationRegisterVnutrennieZakazyTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	Zakazchik                      *String   `json:"Заказчик,omitempty"`
	VnutrenniiZakazKey             *Guid     `json:"ВнутреннийЗаказ_Key,omitempty"`
	StatusPartii                   *String   `json:"СтатусПартии,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	ZakazchikType                  *String   `json:"Заказчик_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterVnutrennieZakazyBalanceAndTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	Zakazchik                      *String   `json:"Заказчик,omitempty"`
	VnutrenniiZakazKey             *Guid     `json:"ВнутреннийЗаказ_Key,omitempty"`
	StatusPartii                   *String   `json:"СтатусПартии,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoOpeningBalance      *Int64    `json:"КоличествоOpeningBalance,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	KolichestvoClosingBalance      *Int64    `json:"КоличествоClosingBalance,omitempty"`
	VesOpeningBalance              *Double   `json:"ВесOpeningBalance,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	VesClosingBalance              *Double   `json:"ВесClosingBalance,omitempty"`
	ZakazchikType                  *String   `json:"Заказчик_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKomitentaRowType struct {
	Recorder              String    `json:"Recorder,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	LineNumber            Int64     `json:"LineNumber,omitempty"`
	Active                *Boolean  `json:"Active,omitempty"`
	RecordType            *String   `json:"RecordType,omitempty"`
	DogovorKontragentaKey *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                *String   `json:"Сделка,omitempty"`
	SummaVzaimoraschetov  *Double   `json:"СуммаВзаиморасчетов,omitempty"`
	SummaUpr              *Double   `json:"СуммаУпр,omitempty"`
	RecorderType          String    `json:"Recorder_Type,omitempty"`
	SdelkaType            *String   `json:"Сделка_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKomitentaBalance struct {
	DogovorKontragentaKey       *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                      *String `json:"Сделка,omitempty"`
	SummaVzaimoraschetovBalance *Double `json:"СуммаВзаиморасчетовBalance,omitempty"`
	SummaUprBalance             *Double `json:"СуммаУпрBalance,omitempty"`
	SdelkaType                  *String `json:"Сделка_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKomitentaTurnover struct {
	DogovorKontragentaKey        *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                       *String   `json:"Сделка,omitempty"`
	Period                       *DateTime `json:"Period,omitempty"`
	SecondPeriod                 *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                 *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                   *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                    *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                   *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                  *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod               *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                   *DateTime `json:"YearPeriod,omitempty"`
	Recorder                     *String   `json:"Recorder,omitempty"`
	LineNumber                   *Int64    `json:"LineNumber,omitempty"`
	SummaVzaimoraschetovTurnover *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt  *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense  *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SummaUprTurnover             *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt              *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense              *Double   `json:"СуммаУпрExpense,omitempty"`
	SdelkaType                   *String   `json:"Сделка_Type,omitempty"`
	RecorderType                 *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKomitentaBalanceAndTurnover struct {
	DogovorKontragentaKey              *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                             *String   `json:"Сделка,omitempty"`
	Period                             *DateTime `json:"Period,omitempty"`
	SecondPeriod                       *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                       *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                         *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                          *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                         *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                      *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                        *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                      *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                     *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                         *DateTime `json:"YearPeriod,omitempty"`
	Recorder                           *String   `json:"Recorder,omitempty"`
	LineNumber                         *Int64    `json:"LineNumber,omitempty"`
	SummaVzaimoraschetovOpeningBalance *Double   `json:"СуммаВзаиморасчетовOpeningBalance,omitempty"`
	SummaVzaimoraschetovTurnover       *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt        *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense        *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SummaVzaimoraschetovClosingBalance *Double   `json:"СуммаВзаиморасчетовClosingBalance,omitempty"`
	SummaUprOpeningBalance             *Double   `json:"СуммаУпрOpeningBalance,omitempty"`
	SummaUprTurnover                   *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt                    *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense                    *Double   `json:"СуммаУпрExpense,omitempty"`
	SummaUprClosingBalance             *Double   `json:"СуммаУпрClosingBalance,omitempty"`
	SdelkaType                         *String   `json:"Сделка_Type,omitempty"`
	RecorderType                       *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterZakazyKlientovRowType struct {
	Recorder        String    `json:"Recorder,omitempty"`
	Period          *DateTime `json:"Period,omitempty"`
	LineNumber      Int64     `json:"LineNumber,omitempty"`
	Active          *Boolean  `json:"Active,omitempty"`
	RecordType      *String   `json:"RecordType,omitempty"`
	ItemKey         *Guid     `json:"Номенклатура_Key,omitempty"`
	SizeKey         *Guid     `json:"Размер_Key,omitempty"`
	ZakazKlientaKey *Guid     `json:"ЗаказКлиента_Key,omitempty"`
	InstanceKey     *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	Zakazano        *Int64    `json:"Заказано,omitempty"`
	VRezerve        *Int64    `json:"ВРезерве,omitempty"`
	KRazmeshcheniiu *Int64    `json:"КРазмещению,omitempty"`
	ZakazanoVes     *Double   `json:"ЗаказаноВес,omitempty"`
	ZakazanoSumma   *Double   `json:"ЗаказаноСумма,omitempty"`
	RecorderType    String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterZakazyKlientovBalance struct {
	ItemKey                *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                *Guid   `json:"Размер_Key,omitempty"`
	ZakazKlientaKey        *Guid   `json:"ЗаказКлиента_Key,omitempty"`
	InstanceKey            *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	ZakazanoBalance        *Int64  `json:"ЗаказаноBalance,omitempty"`
	VRezerveBalance        *Int64  `json:"ВРезервеBalance,omitempty"`
	KRazmeshcheniiuBalance *Int64  `json:"КРазмещениюBalance,omitempty"`
	ZakazanoVesBalance     *Double `json:"ЗаказаноВесBalance,omitempty"`
	ZakazanoSummaBalance   *Double `json:"ЗаказаноСуммаBalance,omitempty"`
}
type AccumulationRegisterZakazyKlientovTurnover struct {
	ItemKey                 *Guid     `json:"Номенклатура_Key,omitempty"`
	SizeKey                 *Guid     `json:"Размер_Key,omitempty"`
	ZakazKlientaKey         *Guid     `json:"ЗаказКлиента_Key,omitempty"`
	InstanceKey             *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	Period                  *DateTime `json:"Period,omitempty"`
	SecondPeriod            *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod            *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod              *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod               *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod              *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod           *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod             *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod           *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod          *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod              *DateTime `json:"YearPeriod,omitempty"`
	Recorder                *String   `json:"Recorder,omitempty"`
	LineNumber              *Int64    `json:"LineNumber,omitempty"`
	ZakazanoTurnover        *Int64    `json:"ЗаказаноTurnover,omitempty"`
	ZakazanoReceipt         *Int64    `json:"ЗаказаноReceipt,omitempty"`
	ZakazanoExpense         *Int64    `json:"ЗаказаноExpense,omitempty"`
	VRezerveTurnover        *Int64    `json:"ВРезервеTurnover,omitempty"`
	VRezerveReceipt         *Int64    `json:"ВРезервеReceipt,omitempty"`
	VRezerveExpense         *Int64    `json:"ВРезервеExpense,omitempty"`
	KRazmeshcheniiuTurnover *Int64    `json:"КРазмещениюTurnover,omitempty"`
	KRazmeshcheniiuReceipt  *Int64    `json:"КРазмещениюReceipt,omitempty"`
	KRazmeshcheniiuExpense  *Int64    `json:"КРазмещениюExpense,omitempty"`
	ZakazanoVesTurnover     *Double   `json:"ЗаказаноВесTurnover,omitempty"`
	ZakazanoVesReceipt      *Double   `json:"ЗаказаноВесReceipt,omitempty"`
	ZakazanoVesExpense      *Double   `json:"ЗаказаноВесExpense,omitempty"`
	ZakazanoSummaTurnover   *Double   `json:"ЗаказаноСуммаTurnover,omitempty"`
	ZakazanoSummaReceipt    *Double   `json:"ЗаказаноСуммаReceipt,omitempty"`
	ZakazanoSummaExpense    *Double   `json:"ЗаказаноСуммаExpense,omitempty"`
	RecorderType            *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterZakazyKlientovBalanceAndTurnover struct {
	ItemKey                       *Guid     `json:"Номенклатура_Key,omitempty"`
	SizeKey                       *Guid     `json:"Размер_Key,omitempty"`
	ZakazKlientaKey               *Guid     `json:"ЗаказКлиента_Key,omitempty"`
	InstanceKey                   *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	Period                        *DateTime `json:"Period,omitempty"`
	SecondPeriod                  *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                  *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                    *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                     *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                    *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                 *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                   *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                 *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                    *DateTime `json:"YearPeriod,omitempty"`
	Recorder                      *String   `json:"Recorder,omitempty"`
	LineNumber                    *Int64    `json:"LineNumber,omitempty"`
	ZakazanoOpeningBalance        *Int64    `json:"ЗаказаноOpeningBalance,omitempty"`
	ZakazanoTurnover              *Int64    `json:"ЗаказаноTurnover,omitempty"`
	ZakazanoReceipt               *Int64    `json:"ЗаказаноReceipt,omitempty"`
	ZakazanoExpense               *Int64    `json:"ЗаказаноExpense,omitempty"`
	ZakazanoClosingBalance        *Int64    `json:"ЗаказаноClosingBalance,omitempty"`
	VRezerveOpeningBalance        *Int64    `json:"ВРезервеOpeningBalance,omitempty"`
	VRezerveTurnover              *Int64    `json:"ВРезервеTurnover,omitempty"`
	VRezerveReceipt               *Int64    `json:"ВРезервеReceipt,omitempty"`
	VRezerveExpense               *Int64    `json:"ВРезервеExpense,omitempty"`
	VRezerveClosingBalance        *Int64    `json:"ВРезервеClosingBalance,omitempty"`
	KRazmeshcheniiuOpeningBalance *Int64    `json:"КРазмещениюOpeningBalance,omitempty"`
	KRazmeshcheniiuTurnover       *Int64    `json:"КРазмещениюTurnover,omitempty"`
	KRazmeshcheniiuReceipt        *Int64    `json:"КРазмещениюReceipt,omitempty"`
	KRazmeshcheniiuExpense        *Int64    `json:"КРазмещениюExpense,omitempty"`
	KRazmeshcheniiuClosingBalance *Int64    `json:"КРазмещениюClosingBalance,omitempty"`
	ZakazanoVesOpeningBalance     *Double   `json:"ЗаказаноВесOpeningBalance,omitempty"`
	ZakazanoVesTurnover           *Double   `json:"ЗаказаноВесTurnover,omitempty"`
	ZakazanoVesReceipt            *Double   `json:"ЗаказаноВесReceipt,omitempty"`
	ZakazanoVesExpense            *Double   `json:"ЗаказаноВесExpense,omitempty"`
	ZakazanoVesClosingBalance     *Double   `json:"ЗаказаноВесClosingBalance,omitempty"`
	ZakazanoSummaOpeningBalance   *Double   `json:"ЗаказаноСуммаOpeningBalance,omitempty"`
	ZakazanoSummaTurnover         *Double   `json:"ЗаказаноСуммаTurnover,omitempty"`
	ZakazanoSummaReceipt          *Double   `json:"ЗаказаноСуммаReceipt,omitempty"`
	ZakazanoSummaExpense          *Double   `json:"ЗаказаноСуммаExpense,omitempty"`
	ZakazanoSummaClosingBalance   *Double   `json:"ЗаказаноСуммаClosingBalance,omitempty"`
	RecorderType                  *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyPoFinmonitoringuRoznitsaRowType struct {
	Recorder         String    `json:"Recorder,omitempty"`
	Period           *DateTime `json:"Period,omitempty"`
	LineNumber       Int64     `json:"LineNumber,omitempty"`
	Active           *Boolean  `json:"Active,omitempty"`
	RecordType       *String   `json:"RecordType,omitempty"`
	KontragentKey    *Guid     `json:"Контрагент_Key,omitempty"`
	OrganizatsiiaKey *Guid     `json:"Организация_Key,omitempty"`
	SummaPokupok     *Double   `json:"СуммаПокупок,omitempty"`
	SummaSkupki      *Double   `json:"СуммаСкупки,omitempty"`
	RecorderType     String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyPoFinmonitoringuRoznitsaBalance struct {
	KontragentKey       *Guid   `json:"Контрагент_Key,omitempty"`
	OrganizatsiiaKey    *Guid   `json:"Организация_Key,omitempty"`
	SummaPokupokBalance *Double `json:"СуммаПокупокBalance,omitempty"`
	SummaSkupkiBalance  *Double `json:"СуммаСкупкиBalance,omitempty"`
}
type AccumulationRegisterSummyPoFinmonitoringuRoznitsaTurnover struct {
	KontragentKey        *Guid     `json:"Контрагент_Key,omitempty"`
	OrganizatsiiaKey     *Guid     `json:"Организация_Key,omitempty"`
	Period               *DateTime `json:"Period,omitempty"`
	SecondPeriod         *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod         *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod           *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod            *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod           *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod        *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod          *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod        *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod       *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod           *DateTime `json:"YearPeriod,omitempty"`
	Recorder             *String   `json:"Recorder,omitempty"`
	LineNumber           *Int64    `json:"LineNumber,omitempty"`
	SummaPokupokTurnover *Double   `json:"СуммаПокупокTurnover,omitempty"`
	SummaPokupokReceipt  *Double   `json:"СуммаПокупокReceipt,omitempty"`
	SummaPokupokExpense  *Double   `json:"СуммаПокупокExpense,omitempty"`
	SummaSkupkiTurnover  *Double   `json:"СуммаСкупкиTurnover,omitempty"`
	SummaSkupkiReceipt   *Double   `json:"СуммаСкупкиReceipt,omitempty"`
	SummaSkupkiExpense   *Double   `json:"СуммаСкупкиExpense,omitempty"`
	RecorderType         *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyPoFinmonitoringuRoznitsaBalanceAndTurnover struct {
	KontragentKey              *Guid     `json:"Контрагент_Key,omitempty"`
	OrganizatsiiaKey           *Guid     `json:"Организация_Key,omitempty"`
	Period                     *DateTime `json:"Period,omitempty"`
	SecondPeriod               *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod               *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                 *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                  *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                 *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod              *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod              *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod             *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                 *DateTime `json:"YearPeriod,omitempty"`
	Recorder                   *String   `json:"Recorder,omitempty"`
	LineNumber                 *Int64    `json:"LineNumber,omitempty"`
	SummaPokupokOpeningBalance *Double   `json:"СуммаПокупокOpeningBalance,omitempty"`
	SummaPokupokTurnover       *Double   `json:"СуммаПокупокTurnover,omitempty"`
	SummaPokupokReceipt        *Double   `json:"СуммаПокупокReceipt,omitempty"`
	SummaPokupokExpense        *Double   `json:"СуммаПокупокExpense,omitempty"`
	SummaPokupokClosingBalance *Double   `json:"СуммаПокупокClosingBalance,omitempty"`
	SummaSkupkiOpeningBalance  *Double   `json:"СуммаСкупкиOpeningBalance,omitempty"`
	SummaSkupkiTurnover        *Double   `json:"СуммаСкупкиTurnover,omitempty"`
	SummaSkupkiReceipt         *Double   `json:"СуммаСкупкиReceipt,omitempty"`
	SummaSkupkiExpense         *Double   `json:"СуммаСкупкиExpense,omitempty"`
	SummaSkupkiClosingBalance  *Double   `json:"СуммаСкупкиClosingBalance,omitempty"`
	RecorderType               *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRowType struct {
	Recorder                 String    `json:"Recorder,omitempty"`
	Period                   *DateTime `json:"Period,omitempty"`
	LineNumber               Int64     `json:"LineNumber,omitempty"`
	Active                   *Boolean  `json:"Active,omitempty"`
	RecordType               *String   `json:"RecordType,omitempty"`
	VidDenezhnykhSredstv     *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String   `json:"БанковскийСчетКасса,omitempty"`
	DokumentPolucheniia      *String   `json:"ДокументПолучения,omitempty"`
	TypeOfMovingMoneyKey     *Guid     `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	Sum                      *Double   `json:"Сумма,omitempty"`
	SummaUpr                 *Double   `json:"СуммаУпр,omitempty"`
	RecorderType             String    `json:"Recorder_Type,omitempty"`
	BankovskiiSchetKassaType *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	DokumentPolucheniiaType  *String   `json:"ДокументПолучения_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalance struct {
	VidDenezhnykhSredstv     *String `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String `json:"БанковскийСчетКасса,omitempty"`
	DokumentPolucheniia      *String `json:"ДокументПолучения,omitempty"`
	TypeOfMovingMoneyKey     *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaBalance             *Double `json:"СуммаBalance,omitempty"`
	SummaUprBalance          *Double `json:"СуммаУпрBalance,omitempty"`
	BankovskiiSchetKassaType *String `json:"БанковскийСчетКасса_Type,omitempty"`
	DokumentPolucheniiaType  *String `json:"ДокументПолучения_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKPolucheniiuTurnover struct {
	VidDenezhnykhSredstv     *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String   `json:"БанковскийСчетКасса,omitempty"`
	DokumentPolucheniia      *String   `json:"ДокументПолучения,omitempty"`
	TypeOfMovingMoneyKey     *Guid     `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	Period                   *DateTime `json:"Period,omitempty"`
	SecondPeriod             *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod             *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod               *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod               *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod            *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod              *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod            *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod           *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod               *DateTime `json:"YearPeriod,omitempty"`
	Recorder                 *String   `json:"Recorder,omitempty"`
	LineNumber               *Int64    `json:"LineNumber,omitempty"`
	SummaTurnover            *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt             *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense             *Double   `json:"СуммаExpense,omitempty"`
	SummaUprTurnover         *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt          *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense          *Double   `json:"СуммаУпрExpense,omitempty"`
	BankovskiiSchetKassaType *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	DokumentPolucheniiaType  *String   `json:"ДокументПолучения_Type,omitempty"`
	RecorderType             *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalanceAndTurnover struct {
	VidDenezhnykhSredstv     *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String   `json:"БанковскийСчетКасса,omitempty"`
	DokumentPolucheniia      *String   `json:"ДокументПолучения,omitempty"`
	TypeOfMovingMoneyKey     *Guid     `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	Period                   *DateTime `json:"Period,omitempty"`
	SecondPeriod             *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod             *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod               *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod               *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod            *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod              *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod            *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod           *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod               *DateTime `json:"YearPeriod,omitempty"`
	Recorder                 *String   `json:"Recorder,omitempty"`
	LineNumber               *Int64    `json:"LineNumber,omitempty"`
	SummaOpeningBalance      *Double   `json:"СуммаOpeningBalance,omitempty"`
	SummaTurnover            *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt             *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense             *Double   `json:"СуммаExpense,omitempty"`
	SummaClosingBalance      *Double   `json:"СуммаClosingBalance,omitempty"`
	SummaUprOpeningBalance   *Double   `json:"СуммаУпрOpeningBalance,omitempty"`
	SummaUprTurnover         *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt          *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense          *Double   `json:"СуммаУпрExpense,omitempty"`
	SummaUprClosingBalance   *Double   `json:"СуммаУпрClosingBalance,omitempty"`
	BankovskiiSchetKassaType *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	DokumentPolucheniiaType  *String   `json:"ДокументПолучения_Type,omitempty"`
	RecorderType             *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterProdazhiPoDiskontnymKartamRowType struct {
	Recorder                       String    `json:"Recorder,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Active                         *Boolean  `json:"Active,omitempty"`
	MemberCardKey                  *Guid     `json:"ДисконтнаяКарта_Key,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	DokumentProdazhi               *String   `json:"ДокументПродажи,omitempty"`
	Sum                            *Double   `json:"Сумма,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	RecorderType                   String    `json:"Recorder_Type,omitempty"`
	DokumentProdazhiType           *String   `json:"ДокументПродажи_Type,omitempty"`
}
type AccumulationRegisterProdazhiPoDiskontnymKartamTurnover struct {
	MemberCardKey                  *Guid     `json:"ДисконтнаяКарта_Key,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	DokumentProdazhi               *String   `json:"ДокументПродажи,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	SummaTurnover                  *Double   `json:"СуммаTurnover,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	DokumentProdazhiType           *String   `json:"ДокументПродажи_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryPoluchennyeRowType struct {
	Recorder                       String    `json:"Recorder,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Active                         *Boolean  `json:"Active,omitempty"`
	RecordType                     *String   `json:"RecordType,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                         *String   `json:"Сделка,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	SummaVzaimoraschetov           *Double   `json:"СуммаВзаиморасчетов,omitempty"`
	RecorderType                   String    `json:"Recorder_Type,omitempty"`
	SdelkaType                     *String   `json:"Сделка_Type,omitempty"`
}
type AccumulationRegisterTovaryPoluchennyeBalance struct {
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey          *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                         *String `json:"Сделка,omitempty"`
	KolichestvoBalance             *Int64  `json:"КоличествоBalance,omitempty"`
	VesBalance                     *Double `json:"ВесBalance,omitempty"`
	SummaVzaimoraschetovBalance    *Double `json:"СуммаВзаиморасчетовBalance,omitempty"`
	SdelkaType                     *String `json:"Сделка_Type,omitempty"`
}
type AccumulationRegisterTovaryPoluchennyeTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                         *String   `json:"Сделка,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	SummaVzaimoraschetovTurnover   *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt    *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense    *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SdelkaType                     *String   `json:"Сделка_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryPoluchennyeBalanceAndTurnover struct {
	ItemKey                            *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                        *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey     *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                            *Guid     `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey              *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                             *String   `json:"Сделка,omitempty"`
	Period                             *DateTime `json:"Period,omitempty"`
	SecondPeriod                       *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                       *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                         *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                          *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                         *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                      *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                        *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                      *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                     *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                         *DateTime `json:"YearPeriod,omitempty"`
	Recorder                           *String   `json:"Recorder,omitempty"`
	LineNumber                         *Int64    `json:"LineNumber,omitempty"`
	KolichestvoOpeningBalance          *Int64    `json:"КоличествоOpeningBalance,omitempty"`
	KolichestvoTurnover                *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt                 *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense                 *Int64    `json:"КоличествоExpense,omitempty"`
	KolichestvoClosingBalance          *Int64    `json:"КоличествоClosingBalance,omitempty"`
	VesOpeningBalance                  *Double   `json:"ВесOpeningBalance,omitempty"`
	VesTurnover                        *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                         *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                         *Double   `json:"ВесExpense,omitempty"`
	VesClosingBalance                  *Double   `json:"ВесClosingBalance,omitempty"`
	SummaVzaimoraschetovOpeningBalance *Double   `json:"СуммаВзаиморасчетовOpeningBalance,omitempty"`
	SummaVzaimoraschetovTurnover       *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt        *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense        *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SummaVzaimoraschetovClosingBalance *Double   `json:"СуммаВзаиморасчетовClosingBalance,omitempty"`
	SdelkaType                         *String   `json:"Сделка_Type,omitempty"`
	RecorderType                       *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSvobodnyeOstatkiRowType struct {
	Recorder      String    `json:"Recorder,omitempty"`
	Period        *DateTime `json:"Period,omitempty"`
	LineNumber    Int64     `json:"LineNumber,omitempty"`
	Active        *Boolean  `json:"Active,omitempty"`
	RecordType    *String   `json:"RecordType,omitempty"`
	ItemKey       *Guid     `json:"Номенклатура_Key,omitempty"`
	SizeKey       *Guid     `json:"Размер_Key,omitempty"`
	DepartmentKey *Guid     `json:"Склад_Key,omitempty"`
	VNalichii     *Int64    `json:"ВНаличии,omitempty"`
	VRezerve      *Int64    `json:"ВРезерве,omitempty"`
	RecorderType  String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSvobodnyeOstatkiBalance struct {
	ItemKey          *Guid  `json:"Номенклатура_Key,omitempty"`
	SizeKey          *Guid  `json:"Размер_Key,omitempty"`
	DepartmentKey    *Guid  `json:"Склад_Key,omitempty"`
	VNalichiiBalance *Int64 `json:"ВНаличииBalance,omitempty"`
	VRezerveBalance  *Int64 `json:"ВРезервеBalance,omitempty"`
}
type AccumulationRegisterSvobodnyeOstatkiTurnover struct {
	ItemKey           *Guid     `json:"Номенклатура_Key,omitempty"`
	SizeKey           *Guid     `json:"Размер_Key,omitempty"`
	DepartmentKey     *Guid     `json:"Склад_Key,omitempty"`
	Period            *DateTime `json:"Period,omitempty"`
	SecondPeriod      *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod      *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod        *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod         *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod        *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod     *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod       *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod     *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod    *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod        *DateTime `json:"YearPeriod,omitempty"`
	Recorder          *String   `json:"Recorder,omitempty"`
	LineNumber        *Int64    `json:"LineNumber,omitempty"`
	VNalichiiTurnover *Int64    `json:"ВНаличииTurnover,omitempty"`
	VNalichiiReceipt  *Int64    `json:"ВНаличииReceipt,omitempty"`
	VNalichiiExpense  *Int64    `json:"ВНаличииExpense,omitempty"`
	VRezerveTurnover  *Int64    `json:"ВРезервеTurnover,omitempty"`
	VRezerveReceipt   *Int64    `json:"ВРезервеReceipt,omitempty"`
	VRezerveExpense   *Int64    `json:"ВРезервеExpense,omitempty"`
	RecorderType      *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSvobodnyeOstatkiBalanceAndTurnover struct {
	ItemKey                 *Guid     `json:"Номенклатура_Key,omitempty"`
	SizeKey                 *Guid     `json:"Размер_Key,omitempty"`
	DepartmentKey           *Guid     `json:"Склад_Key,omitempty"`
	Period                  *DateTime `json:"Period,omitempty"`
	SecondPeriod            *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod            *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod              *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod               *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod              *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod           *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod             *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod           *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod          *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod              *DateTime `json:"YearPeriod,omitempty"`
	Recorder                *String   `json:"Recorder,omitempty"`
	LineNumber              *Int64    `json:"LineNumber,omitempty"`
	VNalichiiOpeningBalance *Int64    `json:"ВНаличииOpeningBalance,omitempty"`
	VNalichiiTurnover       *Int64    `json:"ВНаличииTurnover,omitempty"`
	VNalichiiReceipt        *Int64    `json:"ВНаличииReceipt,omitempty"`
	VNalichiiExpense        *Int64    `json:"ВНаличииExpense,omitempty"`
	VNalichiiClosingBalance *Int64    `json:"ВНаличииClosingBalance,omitempty"`
	VRezerveOpeningBalance  *Int64    `json:"ВРезервеOpeningBalance,omitempty"`
	VRezerveTurnover        *Int64    `json:"ВРезервеTurnover,omitempty"`
	VRezerveReceipt         *Int64    `json:"ВРезервеReceipt,omitempty"`
	VRezerveExpense         *Int64    `json:"ВРезервеExpense,omitempty"`
	VRezerveClosingBalance  *Int64    `json:"ВРезервеClosingBalance,omitempty"`
	RecorderType            *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyVRassrochkuRowType struct {
	Recorder             String    `json:"Recorder,omitempty"`
	Period               *DateTime `json:"Period,omitempty"`
	LineNumber           Int64     `json:"LineNumber,omitempty"`
	Active               *Boolean  `json:"Active,omitempty"`
	RecordType           *String   `json:"RecordType,omitempty"`
	OrganizatsiiaKey     *Guid     `json:"Организация_Key,omitempty"`
	DogovorRassrochkiKey *Guid     `json:"ДоговорРассрочки_Key,omitempty"`
	SostavStrokiChekaKey *Guid     `json:"СоставСтрокиЧека_Key,omitempty"`
	Sum                  *Double   `json:"Сумма,omitempty"`
	SummaNDS             *Double   `json:"СуммаНДС,omitempty"`
	RecorderType         String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyVRassrochkuBalance struct {
	OrganizatsiiaKey     *Guid   `json:"Организация_Key,omitempty"`
	DogovorRassrochkiKey *Guid   `json:"ДоговорРассрочки_Key,omitempty"`
	SostavStrokiChekaKey *Guid   `json:"СоставСтрокиЧека_Key,omitempty"`
	SummaBalance         *Double `json:"СуммаBalance,omitempty"`
	SummaNDSBalance      *Double `json:"СуммаНДСBalance,omitempty"`
}
type AccumulationRegisterSummyVRassrochkuTurnover struct {
	OrganizatsiiaKey     *Guid     `json:"Организация_Key,omitempty"`
	DogovorRassrochkiKey *Guid     `json:"ДоговорРассрочки_Key,omitempty"`
	SostavStrokiChekaKey *Guid     `json:"СоставСтрокиЧека_Key,omitempty"`
	Period               *DateTime `json:"Period,omitempty"`
	SecondPeriod         *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod         *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod           *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod            *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod           *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod        *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod          *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod        *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod       *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod           *DateTime `json:"YearPeriod,omitempty"`
	Recorder             *String   `json:"Recorder,omitempty"`
	LineNumber           *Int64    `json:"LineNumber,omitempty"`
	SummaTurnover        *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt         *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense         *Double   `json:"СуммаExpense,omitempty"`
	SummaNDSTurnover     *Double   `json:"СуммаНДСTurnover,omitempty"`
	SummaNDSReceipt      *Double   `json:"СуммаНДСReceipt,omitempty"`
	SummaNDSExpense      *Double   `json:"СуммаНДСExpense,omitempty"`
	RecorderType         *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyVRassrochkuBalanceAndTurnover struct {
	OrganizatsiiaKey       *Guid     `json:"Организация_Key,omitempty"`
	DogovorRassrochkiKey   *Guid     `json:"ДоговорРассрочки_Key,omitempty"`
	SostavStrokiChekaKey   *Guid     `json:"СоставСтрокиЧека_Key,omitempty"`
	Period                 *DateTime `json:"Period,omitempty"`
	SecondPeriod           *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod           *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod             *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod              *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod             *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod          *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod            *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod          *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod         *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod             *DateTime `json:"YearPeriod,omitempty"`
	Recorder               *String   `json:"Recorder,omitempty"`
	LineNumber             *Int64    `json:"LineNumber,omitempty"`
	SummaOpeningBalance    *Double   `json:"СуммаOpeningBalance,omitempty"`
	SummaTurnover          *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt           *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense           *Double   `json:"СуммаExpense,omitempty"`
	SummaClosingBalance    *Double   `json:"СуммаClosingBalance,omitempty"`
	SummaNDSOpeningBalance *Double   `json:"СуммаНДСOpeningBalance,omitempty"`
	SummaNDSTurnover       *Double   `json:"СуммаНДСTurnover,omitempty"`
	SummaNDSReceipt        *Double   `json:"СуммаНДСReceipt,omitempty"`
	SummaNDSExpense        *Double   `json:"СуммаНДСExpense,omitempty"`
	SummaNDSClosingBalance *Double   `json:"СуммаНДСClosingBalance,omitempty"`
	RecorderType           *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterGrafikPlatezheiRowType struct {
	Recorder               String    `json:"Recorder,omitempty"`
	Period                 *DateTime `json:"Period,omitempty"`
	LineNumber             Int64     `json:"LineNumber,omitempty"`
	Active                 *Boolean  `json:"Active,omitempty"`
	RecordType             *String   `json:"RecordType,omitempty"`
	OrganizatsiiaKey       *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey  *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	VidDogovoraKontragenta *String   `json:"ВидДоговораКонтрагента,omitempty"`
	DataDolga              *DateTime `json:"ДатаДолга,omitempty"`
	Oplacheno              *Double   `json:"Оплачено,omitempty"`
	NachislenDolg          *Double   `json:"НачисленДолг,omitempty"`
	Dolg                   *Double   `json:"Долг,omitempty"`
	Avans                  *Double   `json:"Аванс,omitempty"`
	OplachenoUpr           *Int64    `json:"ОплаченоУпр,omitempty"`
	NachislenDolgUpr       *Double   `json:"НачисленДолгУпр,omitempty"`
	DolgUpr                *Double   `json:"ДолгУпр,omitempty"`
	AvansUpr               *Double   `json:"АвансУпр,omitempty"`
	RecorderType           String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterGrafikPlatezheiBalance struct {
	OrganizatsiiaKey        *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey   *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	VidDogovoraKontragenta  *String   `json:"ВидДоговораКонтрагента,omitempty"`
	DataDolga               *DateTime `json:"ДатаДолга,omitempty"`
	OplachenoBalance        *Double   `json:"ОплаченоBalance,omitempty"`
	NachislenDolgBalance    *Double   `json:"НачисленДолгBalance,omitempty"`
	DolgBalance             *Double   `json:"ДолгBalance,omitempty"`
	AvansBalance            *Double   `json:"АвансBalance,omitempty"`
	OplachenoUprBalance     *Int64    `json:"ОплаченоУпрBalance,omitempty"`
	NachislenDolgUprBalance *Double   `json:"НачисленДолгУпрBalance,omitempty"`
	DolgUprBalance          *Double   `json:"ДолгУпрBalance,omitempty"`
	AvansUprBalance         *Double   `json:"АвансУпрBalance,omitempty"`
}
type AccumulationRegisterGrafikPlatezheiTurnover struct {
	OrganizatsiiaKey         *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey    *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	VidDogovoraKontragenta   *String   `json:"ВидДоговораКонтрагента,omitempty"`
	DataDolga                *DateTime `json:"ДатаДолга,omitempty"`
	Period                   *DateTime `json:"Period,omitempty"`
	SecondPeriod             *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod             *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod               *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod               *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod            *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod              *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod            *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod           *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod               *DateTime `json:"YearPeriod,omitempty"`
	Recorder                 *String   `json:"Recorder,omitempty"`
	LineNumber               *Int64    `json:"LineNumber,omitempty"`
	OplachenoTurnover        *Double   `json:"ОплаченоTurnover,omitempty"`
	OplachenoReceipt         *Double   `json:"ОплаченоReceipt,omitempty"`
	OplachenoExpense         *Double   `json:"ОплаченоExpense,omitempty"`
	NachislenDolgTurnover    *Double   `json:"НачисленДолгTurnover,omitempty"`
	NachislenDolgReceipt     *Double   `json:"НачисленДолгReceipt,omitempty"`
	NachislenDolgExpense     *Double   `json:"НачисленДолгExpense,omitempty"`
	DolgTurnover             *Double   `json:"ДолгTurnover,omitempty"`
	DolgReceipt              *Double   `json:"ДолгReceipt,omitempty"`
	DolgExpense              *Double   `json:"ДолгExpense,omitempty"`
	AvansTurnover            *Double   `json:"АвансTurnover,omitempty"`
	AvansReceipt             *Double   `json:"АвансReceipt,omitempty"`
	AvansExpense             *Double   `json:"АвансExpense,omitempty"`
	OplachenoUprTurnover     *Int64    `json:"ОплаченоУпрTurnover,omitempty"`
	OplachenoUprReceipt      *Int64    `json:"ОплаченоУпрReceipt,omitempty"`
	OplachenoUprExpense      *Int64    `json:"ОплаченоУпрExpense,omitempty"`
	NachislenDolgUprTurnover *Double   `json:"НачисленДолгУпрTurnover,omitempty"`
	NachislenDolgUprReceipt  *Double   `json:"НачисленДолгУпрReceipt,omitempty"`
	NachislenDolgUprExpense  *Double   `json:"НачисленДолгУпрExpense,omitempty"`
	DolgUprTurnover          *Double   `json:"ДолгУпрTurnover,omitempty"`
	DolgUprReceipt           *Double   `json:"ДолгУпрReceipt,omitempty"`
	DolgUprExpense           *Double   `json:"ДолгУпрExpense,omitempty"`
	AvansUprTurnover         *Double   `json:"АвансУпрTurnover,omitempty"`
	AvansUprReceipt          *Double   `json:"АвансУпрReceipt,omitempty"`
	AvansUprExpense          *Double   `json:"АвансУпрExpense,omitempty"`
	RecorderType             *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterGrafikPlatezheiBalanceAndTurnover struct {
	OrganizatsiiaKey               *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	VidDogovoraKontragenta         *String   `json:"ВидДоговораКонтрагента,omitempty"`
	DataDolga                      *DateTime `json:"ДатаДолга,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	OplachenoOpeningBalance        *Double   `json:"ОплаченоOpeningBalance,omitempty"`
	OplachenoTurnover              *Double   `json:"ОплаченоTurnover,omitempty"`
	OplachenoReceipt               *Double   `json:"ОплаченоReceipt,omitempty"`
	OplachenoExpense               *Double   `json:"ОплаченоExpense,omitempty"`
	OplachenoClosingBalance        *Double   `json:"ОплаченоClosingBalance,omitempty"`
	NachislenDolgOpeningBalance    *Double   `json:"НачисленДолгOpeningBalance,omitempty"`
	NachislenDolgTurnover          *Double   `json:"НачисленДолгTurnover,omitempty"`
	NachislenDolgReceipt           *Double   `json:"НачисленДолгReceipt,omitempty"`
	NachislenDolgExpense           *Double   `json:"НачисленДолгExpense,omitempty"`
	NachislenDolgClosingBalance    *Double   `json:"НачисленДолгClosingBalance,omitempty"`
	DolgOpeningBalance             *Double   `json:"ДолгOpeningBalance,omitempty"`
	DolgTurnover                   *Double   `json:"ДолгTurnover,omitempty"`
	DolgReceipt                    *Double   `json:"ДолгReceipt,omitempty"`
	DolgExpense                    *Double   `json:"ДолгExpense,omitempty"`
	DolgClosingBalance             *Double   `json:"ДолгClosingBalance,omitempty"`
	AvansOpeningBalance            *Double   `json:"АвансOpeningBalance,omitempty"`
	AvansTurnover                  *Double   `json:"АвансTurnover,omitempty"`
	AvansReceipt                   *Double   `json:"АвансReceipt,omitempty"`
	AvansExpense                   *Double   `json:"АвансExpense,omitempty"`
	AvansClosingBalance            *Double   `json:"АвансClosingBalance,omitempty"`
	OplachenoUprOpeningBalance     *Int64    `json:"ОплаченоУпрOpeningBalance,omitempty"`
	OplachenoUprTurnover           *Int64    `json:"ОплаченоУпрTurnover,omitempty"`
	OplachenoUprReceipt            *Int64    `json:"ОплаченоУпрReceipt,omitempty"`
	OplachenoUprExpense            *Int64    `json:"ОплаченоУпрExpense,omitempty"`
	OplachenoUprClosingBalance     *Int64    `json:"ОплаченоУпрClosingBalance,omitempty"`
	NachislenDolgUprOpeningBalance *Double   `json:"НачисленДолгУпрOpeningBalance,omitempty"`
	NachislenDolgUprTurnover       *Double   `json:"НачисленДолгУпрTurnover,omitempty"`
	NachislenDolgUprReceipt        *Double   `json:"НачисленДолгУпрReceipt,omitempty"`
	NachislenDolgUprExpense        *Double   `json:"НачисленДолгУпрExpense,omitempty"`
	NachislenDolgUprClosingBalance *Double   `json:"НачисленДолгУпрClosingBalance,omitempty"`
	DolgUprOpeningBalance          *Double   `json:"ДолгУпрOpeningBalance,omitempty"`
	DolgUprTurnover                *Double   `json:"ДолгУпрTurnover,omitempty"`
	DolgUprReceipt                 *Double   `json:"ДолгУпрReceipt,omitempty"`
	DolgUprExpense                 *Double   `json:"ДолгУпрExpense,omitempty"`
	DolgUprClosingBalance          *Double   `json:"ДолгУпрClosingBalance,omitempty"`
	AvansUprOpeningBalance         *Double   `json:"АвансУпрOpeningBalance,omitempty"`
	AvansUprTurnover               *Double   `json:"АвансУпрTurnover,omitempty"`
	AvansUprReceipt                *Double   `json:"АвансУпрReceipt,omitempty"`
	AvansUprExpense                *Double   `json:"АвансУпрExpense,omitempty"`
	AvansUprClosingBalance         *Double   `json:"АвансУпрClosingBalance,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterRoznichnaiaVyruchkaRowType struct {
	Recorder              String    `json:"Recorder,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	LineNumber            Int64     `json:"LineNumber,omitempty"`
	Active                *Boolean  `json:"Active,omitempty"`
	RecordType            *String   `json:"RecordType,omitempty"`
	RoznichnaiaTochka     *String   `json:"РозничнаяТочка,omitempty"`
	Sum                   *Double   `json:"Сумма,omitempty"`
	PodrazdelenieKey      *Guid     `json:"Подразделение_Key,omitempty"`
	RecorderType          String    `json:"Recorder_Type,omitempty"`
	RoznichnaiaTochkaType *String   `json:"РозничнаяТочка_Type,omitempty"`
}
type AccumulationRegisterRoznichnaiaVyruchkaBalance struct {
	RoznichnaiaTochka     *String `json:"РозничнаяТочка,omitempty"`
	SummaBalance          *Double `json:"СуммаBalance,omitempty"`
	RoznichnaiaTochkaType *String `json:"РозничнаяТочка_Type,omitempty"`
}
type AccumulationRegisterRoznichnaiaVyruchkaTurnover struct {
	RoznichnaiaTochka     *String   `json:"РозничнаяТочка,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	SecondPeriod          *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod          *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod            *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod             *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod            *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod         *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod           *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod         *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod        *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod            *DateTime `json:"YearPeriod,omitempty"`
	Recorder              *String   `json:"Recorder,omitempty"`
	LineNumber            *Int64    `json:"LineNumber,omitempty"`
	SummaTurnover         *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt          *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense          *Double   `json:"СуммаExpense,omitempty"`
	RoznichnaiaTochkaType *String   `json:"РозничнаяТочка_Type,omitempty"`
	RecorderType          *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterRoznichnaiaVyruchkaBalanceAndTurnover struct {
	RoznichnaiaTochka     *String   `json:"РозничнаяТочка,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	SecondPeriod          *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod          *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod            *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod             *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod            *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod         *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod           *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod         *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod        *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod            *DateTime `json:"YearPeriod,omitempty"`
	Recorder              *String   `json:"Recorder,omitempty"`
	LineNumber            *Int64    `json:"LineNumber,omitempty"`
	SummaOpeningBalance   *Double   `json:"СуммаOpeningBalance,omitempty"`
	SummaTurnover         *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt          *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense          *Double   `json:"СуммаExpense,omitempty"`
	SummaClosingBalance   *Double   `json:"СуммаClosingBalance,omitempty"`
	RoznichnaiaTochkaType *String   `json:"РозничнаяТочка_Type,omitempty"`
	RecorderType          *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryVPutiRowType struct {
	Recorder                       String    `json:"Recorder,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Active                         *Boolean  `json:"Active,omitempty"`
	RecordType                     *String   `json:"RecordType,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	OrganizatsiiaKey               *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentOprikhodovaniia        *String   `json:"ДокументОприходования,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	RecorderType                   String    `json:"Recorder_Type,omitempty"`
	DokumentOprikhodovaniiaType    *String   `json:"ДокументОприходования_Type,omitempty"`
}
type AccumulationRegisterTovaryVPutiBalance struct {
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	OrganizatsiiaKey               *Guid   `json:"Организация_Key,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	DogovorKontragentaKey          *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentOprikhodovaniia        *String `json:"ДокументОприходования,omitempty"`
	KolichestvoBalance             *Int64  `json:"КоличествоBalance,omitempty"`
	VesBalance                     *Double `json:"ВесBalance,omitempty"`
	DokumentOprikhodovaniiaType    *String `json:"ДокументОприходования_Type,omitempty"`
}
type AccumulationRegisterTovaryVPutiTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	OrganizatsiiaKey               *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentOprikhodovaniia        *String   `json:"ДокументОприходования,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	DokumentOprikhodovaniiaType    *String   `json:"ДокументОприходования_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryVPutiBalanceAndTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	OrganizatsiiaKey               *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentOprikhodovaniia        *String   `json:"ДокументОприходования,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoOpeningBalance      *Int64    `json:"КоличествоOpeningBalance,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	KolichestvoClosingBalance      *Int64    `json:"КоличествоClosingBalance,omitempty"`
	VesOpeningBalance              *Double   `json:"ВесOpeningBalance,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	VesClosingBalance              *Double   `json:"ВесClosingBalance,omitempty"`
	DokumentOprikhodovaniiaType    *String   `json:"ДокументОприходования_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterPoteriMetallaVProizvodstveRowType struct {
	Recorder                    String    `json:"Recorder,omitempty"`
	Period                      *DateTime `json:"Period,omitempty"`
	LineNumber                  Int64     `json:"LineNumber,omitempty"`
	Active                      *Boolean  `json:"Active,omitempty"`
	OrganizatsiiaKey            *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey       *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Nomenklatura                *String   `json:"Номенклатура,omitempty"`
	DokumentOprikhodovaniia     *String   `json:"ДокументОприходования,omitempty"`
	VesPoter                    *Double   `json:"ВесПотерь,omitempty"`
	RecorderType                String    `json:"Recorder_Type,omitempty"`
	ItemType                    *String   `json:"Номенклатура_Type,omitempty"`
	DokumentOprikhodovaniiaType *String   `json:"ДокументОприходования_Type,omitempty"`
}
type AccumulationRegisterPoteriMetallaVProizvodstveTurnover struct {
	OrganizatsiiaKey            *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey       *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Nomenklatura                *String   `json:"Номенклатура,omitempty"`
	DokumentOprikhodovaniia     *String   `json:"ДокументОприходования,omitempty"`
	Period                      *DateTime `json:"Period,omitempty"`
	SecondPeriod                *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                  *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                   *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                  *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod               *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                 *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod               *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod              *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                  *DateTime `json:"YearPeriod,omitempty"`
	Recorder                    *String   `json:"Recorder,omitempty"`
	LineNumber                  *Int64    `json:"LineNumber,omitempty"`
	VesPoterTurnover            *Double   `json:"ВесПотерьTurnover,omitempty"`
	ItemType                    *String   `json:"Номенклатура_Type,omitempty"`
	DokumentOprikhodovaniiaType *String   `json:"ДокументОприходования_Type,omitempty"`
	RecorderType                *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterPartiiTovarovNaSkladakhRowType struct {
	Recorder                       String    `json:"Recorder,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Active                         *Boolean  `json:"Active,omitempty"`
	RecordType                     *String   `json:"RecordType,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	OrganizatsiiaKey               *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	DokumentOprikhodovaniia        *String   `json:"ДокументОприходования,omitempty"`
	StatusPartii                   *String   `json:"СтатусПартии,omitempty"`
	KachestvoKey                   *Guid     `json:"Качество_Key,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	Cost                           *Double   `json:"Стоимость,omitempty"`
	StoimostUpr                    *Double   `json:"СтоимостьУпр,omitempty"`
	StoimostBezNDS                 *Double   `json:"СтоимостьБезНДС,omitempty"`
	OperationCode                  *String   `json:"КодОперации,omitempty"`
	SpisaniePartii                 *Boolean  `json:"СписаниеПартий,omitempty"`
	NomerKorStroki                 *Int64    `json:"НомерКорСтроки,omitempty"`
	RecorderType                   String    `json:"Recorder_Type,omitempty"`
	DokumentOprikhodovaniiaType    *String   `json:"ДокументОприходования_Type,omitempty"`
}
type AccumulationRegisterPartiiTovarovNaSkladakhBalance struct {
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	OrganizatsiiaKey               *Guid   `json:"Организация_Key,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	DokumentOprikhodovaniia        *String `json:"ДокументОприходования,omitempty"`
	StatusPartii                   *String `json:"СтатусПартии,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	KolichestvoBalance             *Int64  `json:"КоличествоBalance,omitempty"`
	VesBalance                     *Double `json:"ВесBalance,omitempty"`
	StoimostBalance                *Double `json:"СтоимостьBalance,omitempty"`
	StoimostUprBalance             *Double `json:"СтоимостьУпрBalance,omitempty"`
	StoimostBezNDSBalance          *Double `json:"СтоимостьБезНДСBalance,omitempty"`
	DokumentOprikhodovaniiaType    *String `json:"ДокументОприходования_Type,omitempty"`
}
type AccumulationRegisterPartiiTovarovNaSkladakhTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	OrganizatsiiaKey               *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	DokumentOprikhodovaniia        *String   `json:"ДокументОприходования,omitempty"`
	StatusPartii                   *String   `json:"СтатусПартии,omitempty"`
	KachestvoKey                   *Guid     `json:"Качество_Key,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	StoimostTurnover               *Double   `json:"СтоимостьTurnover,omitempty"`
	StoimostReceipt                *Double   `json:"СтоимостьReceipt,omitempty"`
	StoimostExpense                *Double   `json:"СтоимостьExpense,omitempty"`
	StoimostUprTurnover            *Double   `json:"СтоимостьУпрTurnover,omitempty"`
	StoimostUprReceipt             *Double   `json:"СтоимостьУпрReceipt,omitempty"`
	StoimostUprExpense             *Double   `json:"СтоимостьУпрExpense,omitempty"`
	StoimostBezNDSTurnover         *Double   `json:"СтоимостьБезНДСTurnover,omitempty"`
	StoimostBezNDSReceipt          *Double   `json:"СтоимостьБезНДСReceipt,omitempty"`
	StoimostBezNDSExpense          *Double   `json:"СтоимостьБезНДСExpense,omitempty"`
	DokumentOprikhodovaniiaType    *String   `json:"ДокументОприходования_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterPartiiTovarovNaSkladakhBalanceAndTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	OrganizatsiiaKey               *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	DokumentOprikhodovaniia        *String   `json:"ДокументОприходования,omitempty"`
	StatusPartii                   *String   `json:"СтатусПартии,omitempty"`
	KachestvoKey                   *Guid     `json:"Качество_Key,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoOpeningBalance      *Int64    `json:"КоличествоOpeningBalance,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	KolichestvoClosingBalance      *Int64    `json:"КоличествоClosingBalance,omitempty"`
	VesOpeningBalance              *Double   `json:"ВесOpeningBalance,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	VesClosingBalance              *Double   `json:"ВесClosingBalance,omitempty"`
	StoimostOpeningBalance         *Double   `json:"СтоимостьOpeningBalance,omitempty"`
	StoimostTurnover               *Double   `json:"СтоимостьTurnover,omitempty"`
	StoimostReceipt                *Double   `json:"СтоимостьReceipt,omitempty"`
	StoimostExpense                *Double   `json:"СтоимостьExpense,omitempty"`
	StoimostClosingBalance         *Double   `json:"СтоимостьClosingBalance,omitempty"`
	StoimostUprOpeningBalance      *Double   `json:"СтоимостьУпрOpeningBalance,omitempty"`
	StoimostUprTurnover            *Double   `json:"СтоимостьУпрTurnover,omitempty"`
	StoimostUprReceipt             *Double   `json:"СтоимостьУпрReceipt,omitempty"`
	StoimostUprExpense             *Double   `json:"СтоимостьУпрExpense,omitempty"`
	StoimostUprClosingBalance      *Double   `json:"СтоимостьУпрClosingBalance,omitempty"`
	StoimostBezNDSOpeningBalance   *Double   `json:"СтоимостьБезНДСOpeningBalance,omitempty"`
	StoimostBezNDSTurnover         *Double   `json:"СтоимостьБезНДСTurnover,omitempty"`
	StoimostBezNDSReceipt          *Double   `json:"СтоимостьБезНДСReceipt,omitempty"`
	StoimostBezNDSExpense          *Double   `json:"СтоимостьБезНДСExpense,omitempty"`
	StoimostBezNDSClosingBalance   *Double   `json:"СтоимостьБезНДСClosingBalance,omitempty"`
	DokumentOprikhodovaniiaType    *String   `json:"ДокументОприходования_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyDokumentovDliaObmenaRowType struct {
	Recorder         String    `json:"Recorder,omitempty"`
	Period           *DateTime `json:"Period,omitempty"`
	LineNumber       Int64     `json:"LineNumber,omitempty"`
	Active           *Boolean  `json:"Active,omitempty"`
	RecordType       *String   `json:"RecordType,omitempty"`
	OrganizatsiiaKey *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey    *Guid     `json:"Склад_Key,omitempty"`
	DokumentKey      *Guid     `json:"Документ_Key,omitempty"`
	Sum              *Double   `json:"Сумма,omitempty"`
	RecorderType     String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyDokumentovDliaObmenaBalance struct {
	OrganizatsiiaKey *Guid   `json:"Организация_Key,omitempty"`
	DepartmentKey    *Guid   `json:"Склад_Key,omitempty"`
	DokumentKey      *Guid   `json:"Документ_Key,omitempty"`
	SummaBalance     *Double `json:"СуммаBalance,omitempty"`
}
type AccumulationRegisterSummyDokumentovDliaObmenaTurnover struct {
	OrganizatsiiaKey *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey    *Guid     `json:"Склад_Key,omitempty"`
	DokumentKey      *Guid     `json:"Документ_Key,omitempty"`
	Period           *DateTime `json:"Period,omitempty"`
	SecondPeriod     *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod     *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod       *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod        *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod       *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod    *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod      *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod    *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod   *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod       *DateTime `json:"YearPeriod,omitempty"`
	Recorder         *String   `json:"Recorder,omitempty"`
	LineNumber       *Int64    `json:"LineNumber,omitempty"`
	SummaTurnover    *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt     *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense     *Double   `json:"СуммаExpense,omitempty"`
	RecorderType     *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyDokumentovDliaObmenaBalanceAndTurnover struct {
	OrganizatsiiaKey    *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey       *Guid     `json:"Склад_Key,omitempty"`
	DokumentKey         *Guid     `json:"Документ_Key,omitempty"`
	Period              *DateTime `json:"Period,omitempty"`
	SecondPeriod        *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod        *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod          *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod           *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod          *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod       *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod         *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod       *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod      *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod          *DateTime `json:"YearPeriod,omitempty"`
	Recorder            *String   `json:"Recorder,omitempty"`
	LineNumber          *Int64    `json:"LineNumber,omitempty"`
	SummaOpeningBalance *Double   `json:"СуммаOpeningBalance,omitempty"`
	SummaTurnover       *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt        *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense        *Double   `json:"СуммаExpense,omitempty"`
	SummaClosingBalance *Double   `json:"СуммаClosingBalance,omitempty"`
	RecorderType        *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDvizheniiaDenezhnykhSredstvRowType struct {
	Recorder                          String    `json:"Recorder,omitempty"`
	Period                            *DateTime `json:"Period,omitempty"`
	LineNumber                        Int64     `json:"LineNumber,omitempty"`
	Active                            *Boolean  `json:"Active,omitempty"`
	VidDenezhnykhSredstv              *String   `json:"ВидДенежныхСредств,omitempty"`
	PrikhodRaskhod                    *String   `json:"ПриходРасход,omitempty"`
	BankovskiiSchetKassa              *String   `json:"БанковскийСчетКасса,omitempty"`
	TypeOfMovingMoneyKey              *Guid     `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	DokumentDvizheniia                *String   `json:"ДокументДвижения,omitempty"`
	Kontragent                        *String   `json:"Контрагент,omitempty"`
	DogovorKontragentaKey             *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                            *String   `json:"Сделка,omitempty"`
	ProektKey                         *Guid     `json:"Проект_Key,omitempty"`
	DokumentPlanirovaniiaPlatezha     *String   `json:"ДокументПланированияПлатежа,omitempty"`
	Sum                               *Double   `json:"Сумма,omitempty"`
	SummaUpr                          *Double   `json:"СуммаУпр,omitempty"`
	RecorderType                      String    `json:"Recorder_Type,omitempty"`
	BankovskiiSchetKassaType          *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	DokumentDvizheniiaType            *String   `json:"ДокументДвижения_Type,omitempty"`
	KontragentType                    *String   `json:"Контрагент_Type,omitempty"`
	SdelkaType                        *String   `json:"Сделка_Type,omitempty"`
	DokumentPlanirovaniiaPlatezhaType *String   `json:"ДокументПланированияПлатежа_Type,omitempty"`
}
type AccumulationRegisterDvizheniiaDenezhnykhSredstvTurnover struct {
	VidDenezhnykhSredstv              *String   `json:"ВидДенежныхСредств,omitempty"`
	PrikhodRaskhod                    *String   `json:"ПриходРасход,omitempty"`
	BankovskiiSchetKassa              *String   `json:"БанковскийСчетКасса,omitempty"`
	TypeOfMovingMoneyKey              *Guid     `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	DokumentDvizheniia                *String   `json:"ДокументДвижения,omitempty"`
	Kontragent                        *String   `json:"Контрагент,omitempty"`
	DogovorKontragentaKey             *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                            *String   `json:"Сделка,omitempty"`
	ProektKey                         *Guid     `json:"Проект_Key,omitempty"`
	DokumentPlanirovaniiaPlatezha     *String   `json:"ДокументПланированияПлатежа,omitempty"`
	Period                            *DateTime `json:"Period,omitempty"`
	SecondPeriod                      *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                      *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                        *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                         *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                        *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                     *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                       *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                     *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                    *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                        *DateTime `json:"YearPeriod,omitempty"`
	Recorder                          *String   `json:"Recorder,omitempty"`
	LineNumber                        *Int64    `json:"LineNumber,omitempty"`
	SummaTurnover                     *Double   `json:"СуммаTurnover,omitempty"`
	SummaUprTurnover                  *Double   `json:"СуммаУпрTurnover,omitempty"`
	BankovskiiSchetKassaType          *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	DokumentDvizheniiaType            *String   `json:"ДокументДвижения_Type,omitempty"`
	KontragentType                    *String   `json:"Контрагент_Type,omitempty"`
	SdelkaType                        *String   `json:"Сделка_Type,omitempty"`
	DokumentPlanirovaniiaPlatezhaType *String   `json:"ДокументПланированияПлатежа_Type,omitempty"`
	RecorderType                      *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterProdazhiPoStatiamRowType struct {
	Recorder         String    `json:"Recorder,omitempty"`
	Period           *DateTime `json:"Period,omitempty"`
	LineNumber       Int64     `json:"LineNumber,omitempty"`
	Active           *Boolean  `json:"Active,omitempty"`
	RecordType       *String   `json:"RecordType,omitempty"`
	OrganizatsiiaKey *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey    *Guid     `json:"Склад_Key,omitempty"`
	StatiaKey        *Guid     `json:"Статья_Key,omitempty"`
	SummaProdazha    *Double   `json:"СуммаПродажа,omitempty"`
	SummaVozvrat     *Double   `json:"СуммаВозврат,omitempty"`
	RecorderType     String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterProdazhiPoStatiamBalance struct {
	OrganizatsiiaKey     *Guid   `json:"Организация_Key,omitempty"`
	DepartmentKey        *Guid   `json:"Склад_Key,omitempty"`
	StatiaKey            *Guid   `json:"Статья_Key,omitempty"`
	SummaProdazhaBalance *Double `json:"СуммаПродажаBalance,omitempty"`
	SummaVozvratBalance  *Double `json:"СуммаВозвратBalance,omitempty"`
}
type AccumulationRegisterProdazhiPoStatiamTurnover struct {
	OrganizatsiiaKey      *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey         *Guid     `json:"Склад_Key,omitempty"`
	StatiaKey             *Guid     `json:"Статья_Key,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	SecondPeriod          *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod          *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod            *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod             *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod            *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod         *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod           *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod         *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod        *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod            *DateTime `json:"YearPeriod,omitempty"`
	Recorder              *String   `json:"Recorder,omitempty"`
	LineNumber            *Int64    `json:"LineNumber,omitempty"`
	SummaProdazhaTurnover *Double   `json:"СуммаПродажаTurnover,omitempty"`
	SummaProdazhaReceipt  *Double   `json:"СуммаПродажаReceipt,omitempty"`
	SummaProdazhaExpense  *Double   `json:"СуммаПродажаExpense,omitempty"`
	SummaVozvratTurnover  *Double   `json:"СуммаВозвратTurnover,omitempty"`
	SummaVozvratReceipt   *Double   `json:"СуммаВозвратReceipt,omitempty"`
	SummaVozvratExpense   *Double   `json:"СуммаВозвратExpense,omitempty"`
	RecorderType          *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterProdazhiPoStatiamBalanceAndTurnover struct {
	OrganizatsiiaKey            *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey               *Guid     `json:"Склад_Key,omitempty"`
	StatiaKey                   *Guid     `json:"Статья_Key,omitempty"`
	Period                      *DateTime `json:"Period,omitempty"`
	SecondPeriod                *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                  *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                   *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                  *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod               *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                 *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod               *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod              *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                  *DateTime `json:"YearPeriod,omitempty"`
	Recorder                    *String   `json:"Recorder,omitempty"`
	LineNumber                  *Int64    `json:"LineNumber,omitempty"`
	SummaProdazhaOpeningBalance *Double   `json:"СуммаПродажаOpeningBalance,omitempty"`
	SummaProdazhaTurnover       *Double   `json:"СуммаПродажаTurnover,omitempty"`
	SummaProdazhaReceipt        *Double   `json:"СуммаПродажаReceipt,omitempty"`
	SummaProdazhaExpense        *Double   `json:"СуммаПродажаExpense,omitempty"`
	SummaProdazhaClosingBalance *Double   `json:"СуммаПродажаClosingBalance,omitempty"`
	SummaVozvratOpeningBalance  *Double   `json:"СуммаВозвратOpeningBalance,omitempty"`
	SummaVozvratTurnover        *Double   `json:"СуммаВозвратTurnover,omitempty"`
	SummaVozvratReceipt         *Double   `json:"СуммаВозвратReceipt,omitempty"`
	SummaVozvratExpense         *Double   `json:"СуммаВозвратExpense,omitempty"`
	SummaVozvratClosingBalance  *Double   `json:"СуммаВозвратClosingBalance,omitempty"`
	RecorderType                *String   `json:"Recorder_Type,omitempty"`
}
type InformationRegisterTsenyNomenklaturyRowType struct {
	Recorder                       *String  `json:"Recorder,omitempty"`
	Period                         DateTime `json:"Period,omitempty"`
	LineNumber                     *Int64   `json:"LineNumber,omitempty"`
	Active                         *Boolean `json:"Active,omitempty"`
	TipTsenKey                     Guid     `json:"ТипЦен_Key,omitempty"`
	SegmentNomenklaturyKey         Guid     `json:"СегментНоменклатуры_Key,omitempty"`
	ItemKey                        Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        Guid     `json:"Размер_Key,omitempty"`
	Cost                           *Double  `json:"Цена,omitempty"`
	ProtsentSkidkiNatsenki         *Double  `json:"ПроцентСкидкиНаценки,omitempty"`
	ValiutaKey                     *Guid    `json:"Валюта_Key,omitempty"`
	EdinitsaIzmereniiaKey          *Guid    `json:"ЕдиницаИзмерения_Key,omitempty"`
	RecorderType                   *String  `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRowType struct {
	Recorder                      String    `json:"Recorder,omitempty"`
	Period                        *DateTime `json:"Period,omitempty"`
	LineNumber                    Int64     `json:"LineNumber,omitempty"`
	Active                        *Boolean  `json:"Active,omitempty"`
	OrganizatsiiaKey              *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey                 *Guid     `json:"Склад_Key,omitempty"`
	ProdazhnaiaStoimost           *Double   `json:"ПродажнаяСтоимость,omitempty"`
	VsegoSkidki                   *Double   `json:"ВсегоСкидки,omitempty"`
	SkidkiPoDiskontnymKartam      *Double   `json:"СкидкиПоДисконтнымКартам,omitempty"`
	SummaOplatyKartami            *Double   `json:"СуммаОплатыКартами,omitempty"`
	SummaOplatyBankovskimKreditom *Double   `json:"СуммаОплатыБанковскимКредитом,omitempty"`
	SummaVozvrata                 *Double   `json:"СуммаВозврата,omitempty"`
	VesVChekakh                   *Double   `json:"ВесВЧеках,omitempty"`
	KolichestvoChekov             *Int64    `json:"КоличествоЧеков,omitempty"`
	SummaProdazhiSertifikatov     *Double   `json:"СуммаПродажиСертификатов,omitempty"`
	SummaOplatySertifikatami      *Double   `json:"СуммаОплатыСертификатами,omitempty"`
	PogashenoSertifikatami        *Double   `json:"ПогашеноСертификатами,omitempty"`
	SummaOplatyBonusom            *Double   `json:"СуммаОплатыБонусом,omitempty"`
	VesObmena                     *Double   `json:"ВесОбмена,omitempty"`
	SummaObmena                   *Double   `json:"СуммаОбмена,omitempty"`
	VesSkuplennogoTovara          *Double   `json:"ВесСкупленногоТовара,omitempty"`
	VydanoZaSkuplennyiTovar       *Double   `json:"ВыданоЗаСкупленныйТовар,omitempty"`
	KolichestvoIzdelii            *Int64    `json:"КоличествоИзделий,omitempty"`
	SumManualDiscount             *Double   `json:"СуммаРучнойСкидки,omitempty"`
	SumAutoDiscount               *Double   `json:"СуммаАвтоматическойСкидки,omitempty"`
	SummaRassrochki               *Double   `json:"СуммаРассрочки,omitempty"`
	SummaPogasheniiaRassrochki    *Double   `json:"СуммаПогашенияРассрочки,omitempty"`
	SummaOplatyNalichnymi         *Double   `json:"СуммаОплатыНаличными,omitempty"`
	RecorderType                  String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseTurnover struct {
	OrganizatsiiaKey                      *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey                         *Guid     `json:"Склад_Key,omitempty"`
	Period                                *DateTime `json:"Period,omitempty"`
	SecondPeriod                          *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                          *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                            *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                             *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                            *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                         *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                           *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                         *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                        *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                            *DateTime `json:"YearPeriod,omitempty"`
	Recorder                              *String   `json:"Recorder,omitempty"`
	LineNumber                            *Int64    `json:"LineNumber,omitempty"`
	ProdazhnaiaStoimostTurnover           *Double   `json:"ПродажнаяСтоимостьTurnover,omitempty"`
	VsegoSkidkiTurnover                   *Double   `json:"ВсегоСкидкиTurnover,omitempty"`
	SkidkiPoDiskontnymKartamTurnover      *Double   `json:"СкидкиПоДисконтнымКартамTurnover,omitempty"`
	SummaOplatyKartamiTurnover            *Double   `json:"СуммаОплатыКартамиTurnover,omitempty"`
	SummaOplatyBankovskimKreditomTurnover *Double   `json:"СуммаОплатыБанковскимКредитомTurnover,omitempty"`
	SummaVozvrataTurnover                 *Double   `json:"СуммаВозвратаTurnover,omitempty"`
	VesVChekakhTurnover                   *Double   `json:"ВесВЧекахTurnover,omitempty"`
	KolichestvoChekovTurnover             *Int64    `json:"КоличествоЧековTurnover,omitempty"`
	SummaProdazhiSertifikatovTurnover     *Double   `json:"СуммаПродажиСертификатовTurnover,omitempty"`
	SummaOplatySertifikatamiTurnover      *Double   `json:"СуммаОплатыСертификатамиTurnover,omitempty"`
	PogashenoSertifikatamiTurnover        *Double   `json:"ПогашеноСертификатамиTurnover,omitempty"`
	SummaOplatyBonusomTurnover            *Double   `json:"СуммаОплатыБонусомTurnover,omitempty"`
	VesObmenaTurnover                     *Double   `json:"ВесОбменаTurnover,omitempty"`
	SummaObmenaTurnover                   *Double   `json:"СуммаОбменаTurnover,omitempty"`
	VesSkuplennogoTovaraTurnover          *Double   `json:"ВесСкупленногоТовараTurnover,omitempty"`
	VydanoZaSkuplennyiTovarTurnover       *Double   `json:"ВыданоЗаСкупленныйТоварTurnover,omitempty"`
	KolichestvoIzdeliiTurnover            *Int64    `json:"КоличествоИзделийTurnover,omitempty"`
	SummaRuchnoiSkidkiTurnover            *Double   `json:"СуммаРучнойСкидкиTurnover,omitempty"`
	SummaAvtomaticheskoiSkidkiTurnover    *Double   `json:"СуммаАвтоматическойСкидкиTurnover,omitempty"`
	SummaRassrochkiTurnover               *Double   `json:"СуммаРассрочкиTurnover,omitempty"`
	SummaPogasheniiaRassrochkiTurnover    *Double   `json:"СуммаПогашенияРассрочкиTurnover,omitempty"`
	SummaOplatyNalichnymiTurnover         *Double   `json:"СуммаОплатыНаличнымиTurnover,omitempty"`
	RecorderType                          *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaVRezerveRowType struct {
	Recorder                   String    `json:"Recorder,omitempty"`
	Period                     *DateTime `json:"Period,omitempty"`
	LineNumber                 Int64     `json:"LineNumber,omitempty"`
	Active                     *Boolean  `json:"Active,omitempty"`
	RecordType                 *String   `json:"RecordType,omitempty"`
	VidDenezhnykhSredstv       *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa       *String   `json:"БанковскийСчетКасса,omitempty"`
	DokumentRezervirovaniiaKey *Guid     `json:"ДокументРезервирования_Key,omitempty"`
	Sum                        *Double   `json:"Сумма,omitempty"`
	RecorderType               String    `json:"Recorder_Type,omitempty"`
	BankovskiiSchetKassaType   *String   `json:"БанковскийСчетКасса_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaVRezerveBalance struct {
	VidDenezhnykhSredstv       *String `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa       *String `json:"БанковскийСчетКасса,omitempty"`
	DokumentRezervirovaniiaKey *Guid   `json:"ДокументРезервирования_Key,omitempty"`
	SummaBalance               *Double `json:"СуммаBalance,omitempty"`
	BankovskiiSchetKassaType   *String `json:"БанковскийСчетКасса_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaVRezerveTurnover struct {
	VidDenezhnykhSredstv       *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa       *String   `json:"БанковскийСчетКасса,omitempty"`
	DokumentRezervirovaniiaKey *Guid     `json:"ДокументРезервирования_Key,omitempty"`
	Period                     *DateTime `json:"Period,omitempty"`
	SecondPeriod               *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod               *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                 *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                  *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                 *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod              *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod              *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod             *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                 *DateTime `json:"YearPeriod,omitempty"`
	Recorder                   *String   `json:"Recorder,omitempty"`
	LineNumber                 *Int64    `json:"LineNumber,omitempty"`
	SummaTurnover              *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt               *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense               *Double   `json:"СуммаExpense,omitempty"`
	BankovskiiSchetKassaType   *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	RecorderType               *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaVRezerveBalanceAndTurnover struct {
	VidDenezhnykhSredstv       *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa       *String   `json:"БанковскийСчетКасса,omitempty"`
	DokumentRezervirovaniiaKey *Guid     `json:"ДокументРезервирования_Key,omitempty"`
	Period                     *DateTime `json:"Period,omitempty"`
	SecondPeriod               *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod               *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                 *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                  *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                 *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod              *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod              *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod             *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                 *DateTime `json:"YearPeriod,omitempty"`
	Recorder                   *String   `json:"Recorder,omitempty"`
	LineNumber                 *Int64    `json:"LineNumber,omitempty"`
	SummaOpeningBalance        *Double   `json:"СуммаOpeningBalance,omitempty"`
	SummaTurnover              *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt               *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense               *Double   `json:"СуммаExpense,omitempty"`
	SummaClosingBalance        *Double   `json:"СуммаClosingBalance,omitempty"`
	BankovskiiSchetKassaType   *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	RecorderType               *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRowType struct {
	Recorder                       String    `json:"Recorder,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Active                         *Boolean  `json:"Active,omitempty"`
	RecordType                     *String   `json:"RecordType,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	RetailCost                     *Double   `json:"ЦенаВРознице,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	RecorderType                   String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalance struct {
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	RetailCost                     *Double `json:"ЦенаВРознице,omitempty"`
	KolichestvoBalance             *Int64  `json:"КоличествоBalance,omitempty"`
	VesBalance                     *Double `json:"ВесBalance,omitempty"`
}
type AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	RetailCost                     *Double   `json:"ЦенаВРознице,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalanceAndTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	RetailCost                     *Double   `json:"ЦенаВРознице,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoOpeningBalance      *Int64    `json:"КоличествоOpeningBalance,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	KolichestvoClosingBalance      *Int64    `json:"КоличествоClosingBalance,omitempty"`
	VesOpeningBalance              *Double   `json:"ВесOpeningBalance,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	VesClosingBalance              *Double   `json:"ВесClosingBalance,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDavalcheskiiMetallPoteriRowType struct {
	Recorder              String    `json:"Recorder,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	LineNumber            Int64     `json:"LineNumber,omitempty"`
	Active                *Boolean  `json:"Active,omitempty"`
	OrganizatsiiaKey      *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	MetallKey             *Guid     `json:"Металл_Key,omitempty"`
	Weight                *Double   `json:"Вес,omitempty"`
	Protsent              *Double   `json:"Процент,omitempty"`
	RecorderType          String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDavalcheskiiMetallPoteriTurnover struct {
	OrganizatsiiaKey      *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	MetallKey             *Guid     `json:"Металл_Key,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	SecondPeriod          *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod          *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod            *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod             *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod            *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod         *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod           *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod         *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod        *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod            *DateTime `json:"YearPeriod,omitempty"`
	Recorder              *String   `json:"Recorder,omitempty"`
	LineNumber            *Int64    `json:"LineNumber,omitempty"`
	VesTurnover           *Double   `json:"ВесTurnover,omitempty"`
	RecorderType          *String   `json:"Recorder_Type,omitempty"`
}
type InformationRegisterTsenyPoPreiskurantuRowType struct {
	Recorder         *String  `json:"Recorder,omitempty"`
	Period           DateTime `json:"Period,omitempty"`
	LineNumber       *Int64   `json:"LineNumber,omitempty"`
	Active           *Boolean `json:"Active,omitempty"`
	KamenKey         Guid     `json:"Камень_Key,omitempty"`
	FormaOgrankiKey  Guid     `json:"ФормаОгранки_Key,omitempty"`
	TsvetKamniaKey   Guid     `json:"ЦветКамня_Key,omitempty"`
	GruppaTsvetaKey  Guid     `json:"ГруппаЦвета_Key,omitempty"`
	GruppaDefektaKey Guid     `json:"ГруппаДефекта_Key,omitempty"`
	RassevKey        Guid     `json:"Рассев_Key,omitempty"`
	Razmer1Ot        Double   `json:"Размер1От,omitempty"`
	Razmer1Do        Double   `json:"Размер1До,omitempty"`
	Cost             *Double  `json:"Цена,omitempty"`
	TipTsenKey       *Guid    `json:"ТипЦен_Key,omitempty"`
	RecorderType     *String  `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryVOtboreRowType struct {
	Recorder                       String    `json:"Recorder,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Active                         *Boolean  `json:"Active,omitempty"`
	RecordType                     *String   `json:"RecordType,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	ZakazKlienta                   *String   `json:"ЗаказКлиента,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	KOtboru                        *Int64    `json:"КОтбору,omitempty"`
	Otobrano                       *Int64    `json:"Отобрано,omitempty"`
	OtobranoVes                    *Double   `json:"ОтобраноВес,omitempty"`
	RecorderType                   String    `json:"Recorder_Type,omitempty"`
	ZakazKlientaType               *String   `json:"ЗаказКлиента_Type,omitempty"`
}
type AccumulationRegisterTovaryVOtboreBalance struct {
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	ZakazKlienta                   *String `json:"ЗаказКлиента,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	KOtboruBalance                 *Int64  `json:"КОтборуBalance,omitempty"`
	OtobranoBalance                *Int64  `json:"ОтобраноBalance,omitempty"`
	OtobranoVesBalance             *Double `json:"ОтобраноВесBalance,omitempty"`
	ZakazKlientaType               *String `json:"ЗаказКлиента_Type,omitempty"`
}
type AccumulationRegisterTovaryVOtboreTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	ZakazKlienta                   *String   `json:"ЗаказКлиента,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KOtboruTurnover                *Int64    `json:"КОтборуTurnover,omitempty"`
	KOtboruReceipt                 *Int64    `json:"КОтборуReceipt,omitempty"`
	KOtboruExpense                 *Int64    `json:"КОтборуExpense,omitempty"`
	OtobranoTurnover               *Int64    `json:"ОтобраноTurnover,omitempty"`
	OtobranoReceipt                *Int64    `json:"ОтобраноReceipt,omitempty"`
	OtobranoExpense                *Int64    `json:"ОтобраноExpense,omitempty"`
	OtobranoVesTurnover            *Double   `json:"ОтобраноВесTurnover,omitempty"`
	OtobranoVesReceipt             *Double   `json:"ОтобраноВесReceipt,omitempty"`
	OtobranoVesExpense             *Double   `json:"ОтобраноВесExpense,omitempty"`
	ZakazKlientaType               *String   `json:"ЗаказКлиента_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryVOtboreBalanceAndTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	ZakazKlienta                   *String   `json:"ЗаказКлиента,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KOtboruOpeningBalance          *Int64    `json:"КОтборуOpeningBalance,omitempty"`
	KOtboruTurnover                *Int64    `json:"КОтборуTurnover,omitempty"`
	KOtboruReceipt                 *Int64    `json:"КОтборуReceipt,omitempty"`
	KOtboruExpense                 *Int64    `json:"КОтборуExpense,omitempty"`
	KOtboruClosingBalance          *Int64    `json:"КОтборуClosingBalance,omitempty"`
	OtobranoOpeningBalance         *Int64    `json:"ОтобраноOpeningBalance,omitempty"`
	OtobranoTurnover               *Int64    `json:"ОтобраноTurnover,omitempty"`
	OtobranoReceipt                *Int64    `json:"ОтобраноReceipt,omitempty"`
	OtobranoExpense                *Int64    `json:"ОтобраноExpense,omitempty"`
	OtobranoClosingBalance         *Int64    `json:"ОтобраноClosingBalance,omitempty"`
	OtobranoVesOpeningBalance      *Double   `json:"ОтобраноВесOpeningBalance,omitempty"`
	OtobranoVesTurnover            *Double   `json:"ОтобраноВесTurnover,omitempty"`
	OtobranoVesReceipt             *Double   `json:"ОтобраноВесReceipt,omitempty"`
	OtobranoVesExpense             *Double   `json:"ОтобраноВесExpense,omitempty"`
	OtobranoVesClosingBalance      *Double   `json:"ОтобраноВесClosingBalance,omitempty"`
	ZakazKlientaType               *String   `json:"ЗаказКлиента_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterRealizovannyeTovaryRowType struct {
	Recorder                       String    `json:"Recorder,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Active                         *Boolean  `json:"Active,omitempty"`
	RecordType                     *String   `json:"RecordType,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                         *String   `json:"Сделка,omitempty"`
	DokumentPostavki               *String   `json:"ДокументПоставки,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	Vyruchka                       *Double   `json:"Выручка,omitempty"`
	RecorderType                   String    `json:"Recorder_Type,omitempty"`
	SdelkaType                     *String   `json:"Сделка_Type,omitempty"`
	DokumentPostavkiType           *String   `json:"ДокументПоставки_Type,omitempty"`
}
type AccumulationRegisterRealizovannyeTovaryBalance struct {
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey          *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                         *String `json:"Сделка,omitempty"`
	DokumentPostavki               *String `json:"ДокументПоставки,omitempty"`
	KolichestvoBalance             *Int64  `json:"КоличествоBalance,omitempty"`
	VesBalance                     *Double `json:"ВесBalance,omitempty"`
	VyruchkaBalance                *Double `json:"ВыручкаBalance,omitempty"`
	SdelkaType                     *String `json:"Сделка_Type,omitempty"`
	DokumentPostavkiType           *String `json:"ДокументПоставки_Type,omitempty"`
}
type AccumulationRegisterRealizovannyeTovaryTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                         *String   `json:"Сделка,omitempty"`
	DokumentPostavki               *String   `json:"ДокументПоставки,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	VyruchkaTurnover               *Double   `json:"ВыручкаTurnover,omitempty"`
	VyruchkaReceipt                *Double   `json:"ВыручкаReceipt,omitempty"`
	VyruchkaExpense                *Double   `json:"ВыручкаExpense,omitempty"`
	SdelkaType                     *String   `json:"Сделка_Type,omitempty"`
	DokumentPostavkiType           *String   `json:"ДокументПоставки_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterRealizovannyeTovaryBalanceAndTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                         *String   `json:"Сделка,omitempty"`
	DokumentPostavki               *String   `json:"ДокументПоставки,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoOpeningBalance      *Int64    `json:"КоличествоOpeningBalance,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	KolichestvoClosingBalance      *Int64    `json:"КоличествоClosingBalance,omitempty"`
	VesOpeningBalance              *Double   `json:"ВесOpeningBalance,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	VesClosingBalance              *Double   `json:"ВесClosingBalance,omitempty"`
	VyruchkaOpeningBalance         *Double   `json:"ВыручкаOpeningBalance,omitempty"`
	VyruchkaTurnover               *Double   `json:"ВыручкаTurnover,omitempty"`
	VyruchkaReceipt                *Double   `json:"ВыручкаReceipt,omitempty"`
	VyruchkaExpense                *Double   `json:"ВыручкаExpense,omitempty"`
	VyruchkaClosingBalance         *Double   `json:"ВыручкаClosingBalance,omitempty"`
	SdelkaType                     *String   `json:"Сделка_Type,omitempty"`
	DokumentPostavkiType           *String   `json:"ДокументПоставки_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKomissioneraRowType struct {
	Recorder              String    `json:"Recorder,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	LineNumber            Int64     `json:"LineNumber,omitempty"`
	Active                *Boolean  `json:"Active,omitempty"`
	RecordType            *String   `json:"RecordType,omitempty"`
	DogovorKontragentaKey *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                *String   `json:"Сделка,omitempty"`
	SummaVzaimoraschetov  *Double   `json:"СуммаВзаиморасчетов,omitempty"`
	SummaUpr              *Double   `json:"СуммаУпр,omitempty"`
	RecorderType          String    `json:"Recorder_Type,omitempty"`
	SdelkaType            *String   `json:"Сделка_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKomissioneraBalance struct {
	DogovorKontragentaKey       *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                      *String `json:"Сделка,omitempty"`
	SummaVzaimoraschetovBalance *Double `json:"СуммаВзаиморасчетовBalance,omitempty"`
	SummaUprBalance             *Double `json:"СуммаУпрBalance,omitempty"`
	SdelkaType                  *String `json:"Сделка_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKomissioneraTurnover struct {
	DogovorKontragentaKey        *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                       *String   `json:"Сделка,omitempty"`
	Period                       *DateTime `json:"Period,omitempty"`
	SecondPeriod                 *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                 *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                   *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                    *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                   *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                  *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod               *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                   *DateTime `json:"YearPeriod,omitempty"`
	Recorder                     *String   `json:"Recorder,omitempty"`
	LineNumber                   *Int64    `json:"LineNumber,omitempty"`
	SummaVzaimoraschetovTurnover *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt  *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense  *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SummaUprTurnover             *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt              *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense              *Double   `json:"СуммаУпрExpense,omitempty"`
	SdelkaType                   *String   `json:"Сделка_Type,omitempty"`
	RecorderType                 *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKomissioneraBalanceAndTurnover struct {
	DogovorKontragentaKey              *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                             *String   `json:"Сделка,omitempty"`
	Period                             *DateTime `json:"Period,omitempty"`
	SecondPeriod                       *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                       *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                         *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                          *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                         *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                      *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                        *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                      *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                     *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                         *DateTime `json:"YearPeriod,omitempty"`
	Recorder                           *String   `json:"Recorder,omitempty"`
	LineNumber                         *Int64    `json:"LineNumber,omitempty"`
	SummaVzaimoraschetovOpeningBalance *Double   `json:"СуммаВзаиморасчетовOpeningBalance,omitempty"`
	SummaVzaimoraschetovTurnover       *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt        *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense        *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SummaVzaimoraschetovClosingBalance *Double   `json:"СуммаВзаиморасчетовClosingBalance,omitempty"`
	SummaUprOpeningBalance             *Double   `json:"СуммаУпрOpeningBalance,omitempty"`
	SummaUprTurnover                   *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt                    *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense                    *Double   `json:"СуммаУпрExpense,omitempty"`
	SummaUprClosingBalance             *Double   `json:"СуммаУпрClosingBalance,omitempty"`
	SdelkaType                         *String   `json:"Сделка_Type,omitempty"`
	RecorderType                       *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterProdazhi1RowType struct {
	Recorder                       String    `json:"Recorder,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Active                         *Boolean  `json:"Active,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	OrganizatsiiaKey               *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentProdazhi               *String   `json:"ДокументПродажи,omitempty"`
	DokumentOprikhodovaniia        *String   `json:"ДокументОприходования,omitempty"`
	ZakazKlienta                   *String   `json:"ЗаказКлиента,omitempty"`
	ProektKey                      *Guid     `json:"Проект_Key,omitempty"`
	PodrazdelenieKey               *Guid     `json:"Подразделение_Key,omitempty"`
	MetallKey                      *Guid     `json:"Металл_Key,omitempty"`
	TovarnaiaGruppaKey             *Guid     `json:"ТоварнаяГруппа_Key,omitempty"`
	TovarnaiaKategoriiaKey         *Guid     `json:"ТоварнаяКатегория_Key,omitempty"`
	TovarnaiaPozitsiiaKey          *Guid     `json:"ТоварнаяПозиция_Key,omitempty"`
	OrderKey                       *Guid     `json:"КлючПродажи_Key,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	Cost                           *Double   `json:"Стоимость,omitempty"`
	StoimostBezSkidok              *Double   `json:"СтоимостьБезСкидок,omitempty"`
	StoimostPostuplenie            *Double   `json:"СтоимостьПоступление,omitempty"`
	StoimostUpr                    *Double   `json:"СтоимостьУпр,omitempty"`
	StoimostBezNDS                 *Double   `json:"СтоимостьБезНДС,omitempty"`
	KolichestvoVozvrat             *Int64    `json:"КоличествоВозврат,omitempty"`
	VesVozvrat                     *Double   `json:"ВесВозврат,omitempty"`
	StoimostVozvrat                *Double   `json:"СтоимостьВозврат,omitempty"`
	StoimostBezSkidokVozvrat       *Double   `json:"СтоимостьБезСкидокВозврат,omitempty"`
	StoimostPostuplenieVozvrat     *Double   `json:"СтоимостьПоступлениеВозврат,omitempty"`
	StoimostUprVozvrat             *Double   `json:"СтоимостьУпрВозврат,omitempty"`
	StoimostBezNDSVozvrat          *Double   `json:"СтоимостьБезНДСВозврат,omitempty"`
	SpisaniePartii                 *Boolean  `json:"СписаниеПартий,omitempty"`
	RecorderType                   String    `json:"Recorder_Type,omitempty"`
	DokumentProdazhiType           *String   `json:"ДокументПродажи_Type,omitempty"`
	DokumentOprikhodovaniiaType    *String   `json:"ДокументОприходования_Type,omitempty"`
	ZakazKlientaType               *String   `json:"ЗаказКлиента_Type,omitempty"`
}
type AccumulationRegisterProdazhi1Turnover struct {
	ItemKey                            *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                        *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey     *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                            *Guid     `json:"Размер_Key,omitempty"`
	OrganizatsiiaKey                   *Guid     `json:"Организация_Key,omitempty"`
	DepartmentKey                      *Guid     `json:"Склад_Key,omitempty"`
	DogovorKontragentaKey              *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentProdazhi                   *String   `json:"ДокументПродажи,omitempty"`
	DokumentOprikhodovaniia            *String   `json:"ДокументОприходования,omitempty"`
	ZakazKlienta                       *String   `json:"ЗаказКлиента,omitempty"`
	ProektKey                          *Guid     `json:"Проект_Key,omitempty"`
	PodrazdelenieKey                   *Guid     `json:"Подразделение_Key,omitempty"`
	MetallKey                          *Guid     `json:"Металл_Key,omitempty"`
	TovarnaiaGruppaKey                 *Guid     `json:"ТоварнаяГруппа_Key,omitempty"`
	TovarnaiaKategoriiaKey             *Guid     `json:"ТоварнаяКатегория_Key,omitempty"`
	TovarnaiaPozitsiiaKey              *Guid     `json:"ТоварнаяПозиция_Key,omitempty"`
	OrderKey                           *Guid     `json:"КлючПродажи_Key,omitempty"`
	Period                             *DateTime `json:"Period,omitempty"`
	SecondPeriod                       *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                       *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                         *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                          *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                         *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                      *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                        *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                      *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                     *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                         *DateTime `json:"YearPeriod,omitempty"`
	Recorder                           *String   `json:"Recorder,omitempty"`
	LineNumber                         *Int64    `json:"LineNumber,omitempty"`
	KolichestvoTurnover                *Int64    `json:"КоличествоTurnover,omitempty"`
	VesTurnover                        *Double   `json:"ВесTurnover,omitempty"`
	StoimostTurnover                   *Double   `json:"СтоимостьTurnover,omitempty"`
	StoimostBezSkidokTurnover          *Double   `json:"СтоимостьБезСкидокTurnover,omitempty"`
	StoimostPostuplenieTurnover        *Double   `json:"СтоимостьПоступлениеTurnover,omitempty"`
	StoimostUprTurnover                *Double   `json:"СтоимостьУпрTurnover,omitempty"`
	StoimostBezNDSTurnover             *Double   `json:"СтоимостьБезНДСTurnover,omitempty"`
	KolichestvoVozvratTurnover         *Int64    `json:"КоличествоВозвратTurnover,omitempty"`
	VesVozvratTurnover                 *Double   `json:"ВесВозвратTurnover,omitempty"`
	StoimostVozvratTurnover            *Double   `json:"СтоимостьВозвратTurnover,omitempty"`
	StoimostBezSkidokVozvratTurnover   *Double   `json:"СтоимостьБезСкидокВозвратTurnover,omitempty"`
	StoimostPostuplenieVozvratTurnover *Double   `json:"СтоимостьПоступлениеВозвратTurnover,omitempty"`
	StoimostUprVozvratTurnover         *Double   `json:"СтоимостьУпрВозвратTurnover,omitempty"`
	StoimostBezNDSVozvratTurnover      *Double   `json:"СтоимостьБезНДСВозвратTurnover,omitempty"`
	DokumentProdazhiType               *String   `json:"ДокументПродажи_Type,omitempty"`
	DokumentOprikhodovaniiaType        *String   `json:"ДокументОприходования_Type,omitempty"`
	ZakazKlientaType                   *String   `json:"ЗаказКлиента_Type,omitempty"`
	RecorderType                       *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryNaSkladakhAMRowType struct {
	Recorder               String    `json:"Recorder,omitempty"`
	Period                 *DateTime `json:"Period,omitempty"`
	LineNumber             Int64     `json:"LineNumber,omitempty"`
	Active                 *Boolean  `json:"Active,omitempty"`
	RecordType             *String   `json:"RecordType,omitempty"`
	DepartmentKey          *Guid     `json:"Склад_Key,omitempty"`
	MetallKey              *Guid     `json:"Металл_Key,omitempty"`
	TovarnaiaGruppaKey     *Guid     `json:"ТоварнаяГруппа_Key,omitempty"`
	TovarnaiaKategoriiaKey *Guid     `json:"ТоварнаяКатегория_Key,omitempty"`
	TovarnaiaPozitsiiaKey  *Guid     `json:"ТоварнаяПозиция_Key,omitempty"`
	ItemKey                *Guid     `json:"Номенклатура_Key,omitempty"`
	Quantity               *Int64    `json:"Количество,omitempty"`
	Weight                 *Double   `json:"Вес,omitempty"`
	RoznichnaiaStoimost    *Double   `json:"РозничнаяСтоимость,omitempty"`
	RecorderType           String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryNaSkladakhAMBalance struct {
	DepartmentKey              *Guid   `json:"Склад_Key,omitempty"`
	MetallKey                  *Guid   `json:"Металл_Key,omitempty"`
	TovarnaiaGruppaKey         *Guid   `json:"ТоварнаяГруппа_Key,omitempty"`
	TovarnaiaKategoriiaKey     *Guid   `json:"ТоварнаяКатегория_Key,omitempty"`
	TovarnaiaPozitsiiaKey      *Guid   `json:"ТоварнаяПозиция_Key,omitempty"`
	ItemKey                    *Guid   `json:"Номенклатура_Key,omitempty"`
	KolichestvoBalance         *Int64  `json:"КоличествоBalance,omitempty"`
	VesBalance                 *Double `json:"ВесBalance,omitempty"`
	RoznichnaiaStoimostBalance *Double `json:"РозничнаяСтоимостьBalance,omitempty"`
}
type AccumulationRegisterTovaryNaSkladakhAMTurnover struct {
	DepartmentKey               *Guid     `json:"Склад_Key,omitempty"`
	MetallKey                   *Guid     `json:"Металл_Key,omitempty"`
	TovarnaiaGruppaKey          *Guid     `json:"ТоварнаяГруппа_Key,omitempty"`
	TovarnaiaKategoriiaKey      *Guid     `json:"ТоварнаяКатегория_Key,omitempty"`
	TovarnaiaPozitsiiaKey       *Guid     `json:"ТоварнаяПозиция_Key,omitempty"`
	ItemKey                     *Guid     `json:"Номенклатура_Key,omitempty"`
	Period                      *DateTime `json:"Period,omitempty"`
	SecondPeriod                *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                  *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                   *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                  *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod               *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                 *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod               *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod              *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                  *DateTime `json:"YearPeriod,omitempty"`
	Recorder                    *String   `json:"Recorder,omitempty"`
	LineNumber                  *Int64    `json:"LineNumber,omitempty"`
	KolichestvoTurnover         *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt          *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense          *Int64    `json:"КоличествоExpense,omitempty"`
	VesTurnover                 *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                  *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                  *Double   `json:"ВесExpense,omitempty"`
	RoznichnaiaStoimostTurnover *Double   `json:"РозничнаяСтоимостьTurnover,omitempty"`
	RoznichnaiaStoimostReceipt  *Double   `json:"РозничнаяСтоимостьReceipt,omitempty"`
	RoznichnaiaStoimostExpense  *Double   `json:"РозничнаяСтоимостьExpense,omitempty"`
	RecorderType                *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryNaSkladakhAMBalanceAndTurnover struct {
	DepartmentKey                     *Guid     `json:"Склад_Key,omitempty"`
	MetallKey                         *Guid     `json:"Металл_Key,omitempty"`
	TovarnaiaGruppaKey                *Guid     `json:"ТоварнаяГруппа_Key,omitempty"`
	TovarnaiaKategoriiaKey            *Guid     `json:"ТоварнаяКатегория_Key,omitempty"`
	TovarnaiaPozitsiiaKey             *Guid     `json:"ТоварнаяПозиция_Key,omitempty"`
	ItemKey                           *Guid     `json:"Номенклатура_Key,omitempty"`
	Period                            *DateTime `json:"Period,omitempty"`
	SecondPeriod                      *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                      *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                        *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                         *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                        *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                     *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                       *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                     *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                    *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                        *DateTime `json:"YearPeriod,omitempty"`
	Recorder                          *String   `json:"Recorder,omitempty"`
	LineNumber                        *Int64    `json:"LineNumber,omitempty"`
	KolichestvoOpeningBalance         *Int64    `json:"КоличествоOpeningBalance,omitempty"`
	KolichestvoTurnover               *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt                *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense                *Int64    `json:"КоличествоExpense,omitempty"`
	KolichestvoClosingBalance         *Int64    `json:"КоличествоClosingBalance,omitempty"`
	VesOpeningBalance                 *Double   `json:"ВесOpeningBalance,omitempty"`
	VesTurnover                       *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                        *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                        *Double   `json:"ВесExpense,omitempty"`
	VesClosingBalance                 *Double   `json:"ВесClosingBalance,omitempty"`
	RoznichnaiaStoimostOpeningBalance *Double   `json:"РозничнаяСтоимостьOpeningBalance,omitempty"`
	RoznichnaiaStoimostTurnover       *Double   `json:"РозничнаяСтоимостьTurnover,omitempty"`
	RoznichnaiaStoimostReceipt        *Double   `json:"РозничнаяСтоимостьReceipt,omitempty"`
	RoznichnaiaStoimostExpense        *Double   `json:"РозничнаяСтоимостьExpense,omitempty"`
	RoznichnaiaStoimostClosingBalance *Double   `json:"РозничнаяСтоимостьClosingBalance,omitempty"`
	RecorderType                      *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyPoFinmonitoringuRowType struct {
	Recorder              String    `json:"Recorder,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	LineNumber            Int64     `json:"LineNumber,omitempty"`
	Active                *Boolean  `json:"Active,omitempty"`
	RecordType            *String   `json:"RecordType,omitempty"`
	OrganizatsiiaKey      *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentUcheta        *String   `json:"ДокументУчета,omitempty"`
	SummaOtgruzki         *Double   `json:"СуммаОтгрузки,omitempty"`
	SummaOplaty           *Double   `json:"СуммаОплаты,omitempty"`
	SummaVozvrata         *Double   `json:"СуммаВозврата,omitempty"`
	RecorderType          String    `json:"Recorder_Type,omitempty"`
	DokumentUchetaType    *String   `json:"ДокументУчета_Type,omitempty"`
}
type AccumulationRegisterSummyPoFinmonitoringuBalance struct {
	OrganizatsiiaKey      *Guid   `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentUcheta        *String `json:"ДокументУчета,omitempty"`
	SummaOtgruzkiBalance  *Double `json:"СуммаОтгрузкиBalance,omitempty"`
	SummaOplatyBalance    *Double `json:"СуммаОплатыBalance,omitempty"`
	SummaVozvrataBalance  *Double `json:"СуммаВозвратаBalance,omitempty"`
	DokumentUchetaType    *String `json:"ДокументУчета_Type,omitempty"`
}
type AccumulationRegisterSummyPoFinmonitoringuTurnover struct {
	OrganizatsiiaKey      *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentUcheta        *String   `json:"ДокументУчета,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	SecondPeriod          *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod          *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod            *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod             *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod            *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod         *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod           *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod         *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod        *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod            *DateTime `json:"YearPeriod,omitempty"`
	Recorder              *String   `json:"Recorder,omitempty"`
	LineNumber            *Int64    `json:"LineNumber,omitempty"`
	SummaOtgruzkiTurnover *Double   `json:"СуммаОтгрузкиTurnover,omitempty"`
	SummaOtgruzkiReceipt  *Double   `json:"СуммаОтгрузкиReceipt,omitempty"`
	SummaOtgruzkiExpense  *Double   `json:"СуммаОтгрузкиExpense,omitempty"`
	SummaOplatyTurnover   *Double   `json:"СуммаОплатыTurnover,omitempty"`
	SummaOplatyReceipt    *Double   `json:"СуммаОплатыReceipt,omitempty"`
	SummaOplatyExpense    *Double   `json:"СуммаОплатыExpense,omitempty"`
	SummaVozvrataTurnover *Double   `json:"СуммаВозвратаTurnover,omitempty"`
	SummaVozvrataReceipt  *Double   `json:"СуммаВозвратаReceipt,omitempty"`
	SummaVozvrataExpense  *Double   `json:"СуммаВозвратаExpense,omitempty"`
	DokumentUchetaType    *String   `json:"ДокументУчета_Type,omitempty"`
	RecorderType          *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyPoFinmonitoringuBalanceAndTurnover struct {
	OrganizatsiiaKey            *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey       *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentUcheta              *String   `json:"ДокументУчета,omitempty"`
	Period                      *DateTime `json:"Period,omitempty"`
	SecondPeriod                *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                  *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                   *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                  *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod               *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                 *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod               *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod              *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                  *DateTime `json:"YearPeriod,omitempty"`
	Recorder                    *String   `json:"Recorder,omitempty"`
	LineNumber                  *Int64    `json:"LineNumber,omitempty"`
	SummaOtgruzkiOpeningBalance *Double   `json:"СуммаОтгрузкиOpeningBalance,omitempty"`
	SummaOtgruzkiTurnover       *Double   `json:"СуммаОтгрузкиTurnover,omitempty"`
	SummaOtgruzkiReceipt        *Double   `json:"СуммаОтгрузкиReceipt,omitempty"`
	SummaOtgruzkiExpense        *Double   `json:"СуммаОтгрузкиExpense,omitempty"`
	SummaOtgruzkiClosingBalance *Double   `json:"СуммаОтгрузкиClosingBalance,omitempty"`
	SummaOplatyOpeningBalance   *Double   `json:"СуммаОплатыOpeningBalance,omitempty"`
	SummaOplatyTurnover         *Double   `json:"СуммаОплатыTurnover,omitempty"`
	SummaOplatyReceipt          *Double   `json:"СуммаОплатыReceipt,omitempty"`
	SummaOplatyExpense          *Double   `json:"СуммаОплатыExpense,omitempty"`
	SummaOplatyClosingBalance   *Double   `json:"СуммаОплатыClosingBalance,omitempty"`
	SummaVozvrataOpeningBalance *Double   `json:"СуммаВозвратаOpeningBalance,omitempty"`
	SummaVozvrataTurnover       *Double   `json:"СуммаВозвратаTurnover,omitempty"`
	SummaVozvrataReceipt        *Double   `json:"СуммаВозвратаReceipt,omitempty"`
	SummaVozvrataExpense        *Double   `json:"СуммаВозвратаExpense,omitempty"`
	SummaVozvrataClosingBalance *Double   `json:"СуммаВозвратаClosingBalance,omitempty"`
	DokumentUchetaType          *String   `json:"ДокументУчета_Type,omitempty"`
	RecorderType                *String   `json:"Recorder_Type,omitempty"`
}
type InformationRegisterTsenyNomenklaturyKontragentovRowType struct {
	Recorder                       *String  `json:"Recorder,omitempty"`
	Period                         DateTime `json:"Period,omitempty"`
	LineNumber                     *Int64   `json:"LineNumber,omitempty"`
	Active                         *Boolean `json:"Active,omitempty"`
	TipTsenKey                     Guid     `json:"ТипЦен_Key,omitempty"`
	ItemKey                        Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        Guid     `json:"Размер_Key,omitempty"`
	Cost                           *Double  `json:"Цена,omitempty"`
	ValiutaKey                     *Guid    `json:"Валюта_Key,omitempty"`
	EdinitsaIzmereniiaKey          *Guid    `json:"ЕдиницаИзмерения_Key,omitempty"`
	RecorderType                   *String  `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterVzaimoraschetySKontragentamiRowType struct {
	Recorder              String    `json:"Recorder,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	LineNumber            Int64     `json:"LineNumber,omitempty"`
	Active                *Boolean  `json:"Active,omitempty"`
	RecordType            *String   `json:"RecordType,omitempty"`
	DogovorKontragentaKey *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                *String   `json:"Сделка,omitempty"`
	SummaVzaimoraschetov  *Double   `json:"СуммаВзаиморасчетов,omitempty"`
	SummaUpr              *Double   `json:"СуммаУпр,omitempty"`
	RecorderType          String    `json:"Recorder_Type,omitempty"`
	SdelkaType            *String   `json:"Сделка_Type,omitempty"`
}
type AccumulationRegisterVzaimoraschetySKontragentamiBalance struct {
	DogovorKontragentaKey       *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                      *String `json:"Сделка,omitempty"`
	SummaVzaimoraschetovBalance *Double `json:"СуммаВзаиморасчетовBalance,omitempty"`
	SummaUprBalance             *Double `json:"СуммаУпрBalance,omitempty"`
	SdelkaType                  *String `json:"Сделка_Type,omitempty"`
}
type AccumulationRegisterVzaimoraschetySKontragentamiTurnover struct {
	DogovorKontragentaKey        *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                       *String   `json:"Сделка,omitempty"`
	Period                       *DateTime `json:"Period,omitempty"`
	SecondPeriod                 *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                 *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                   *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                    *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                   *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                  *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod               *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                   *DateTime `json:"YearPeriod,omitempty"`
	Recorder                     *String   `json:"Recorder,omitempty"`
	LineNumber                   *Int64    `json:"LineNumber,omitempty"`
	SummaVzaimoraschetovTurnover *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt  *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense  *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SummaUprTurnover             *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt              *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense              *Double   `json:"СуммаУпрExpense,omitempty"`
	SdelkaType                   *String   `json:"Сделка_Type,omitempty"`
	RecorderType                 *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterVzaimoraschetySKontragentamiBalanceAndTurnover struct {
	DogovorKontragentaKey              *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                             *String   `json:"Сделка,omitempty"`
	Period                             *DateTime `json:"Period,omitempty"`
	SecondPeriod                       *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                       *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                         *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                          *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                         *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                      *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                        *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                      *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                     *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                         *DateTime `json:"YearPeriod,omitempty"`
	Recorder                           *String   `json:"Recorder,omitempty"`
	LineNumber                         *Int64    `json:"LineNumber,omitempty"`
	SummaVzaimoraschetovOpeningBalance *Double   `json:"СуммаВзаиморасчетовOpeningBalance,omitempty"`
	SummaVzaimoraschetovTurnover       *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt        *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense        *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SummaVzaimoraschetovClosingBalance *Double   `json:"СуммаВзаиморасчетовClosingBalance,omitempty"`
	SummaUprOpeningBalance             *Double   `json:"СуммаУпрOpeningBalance,omitempty"`
	SummaUprTurnover                   *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt                    *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense                    *Double   `json:"СуммаУпрExpense,omitempty"`
	SummaUprClosingBalance             *Double   `json:"СуммаУпрClosingBalance,omitempty"`
	SdelkaType                         *String   `json:"Сделка_Type,omitempty"`
	RecorderType                       *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyPokupokPoDiskontnymKartamRowType struct {
	Recorder               String    `json:"Recorder,omitempty"`
	Period                 *DateTime `json:"Period,omitempty"`
	LineNumber             Int64     `json:"LineNumber,omitempty"`
	Active                 *Boolean  `json:"Active,omitempty"`
	RecordType             *String   `json:"RecordType,omitempty"`
	MemberCardKey          *Guid     `json:"ДисконтнаяКарта_Key,omitempty"`
	DataSpisaniia          *DateTime `json:"ДатаСписания,omitempty"`
	Sum                    *Double   `json:"Сумма,omitempty"`
	SummaBonusov           *Double   `json:"СуммаБонусов,omitempty"`
	SummaVremennykhBonusov *Double   `json:"СуммаВременныхБонусов,omitempty"`
	RecorderType           String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyPokupokPoDiskontnymKartamBalance struct {
	MemberCardKey                 *Guid     `json:"ДисконтнаяКарта_Key,omitempty"`
	DataSpisaniia                 *DateTime `json:"ДатаСписания,omitempty"`
	SummaBalance                  *Double   `json:"СуммаBalance,omitempty"`
	SummaBonusovBalance           *Double   `json:"СуммаБонусовBalance,omitempty"`
	SummaVremennykhBonusovBalance *Double   `json:"СуммаВременныхБонусовBalance,omitempty"`
}
type AccumulationRegisterSummyPokupokPoDiskontnymKartamTurnover struct {
	MemberCardKey                  *Guid     `json:"ДисконтнаяКарта_Key,omitempty"`
	DataSpisaniia                  *DateTime `json:"ДатаСписания,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	SummaTurnover                  *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt                   *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense                   *Double   `json:"СуммаExpense,omitempty"`
	SummaBonusovTurnover           *Double   `json:"СуммаБонусовTurnover,omitempty"`
	SummaBonusovReceipt            *Double   `json:"СуммаБонусовReceipt,omitempty"`
	SummaBonusovExpense            *Double   `json:"СуммаБонусовExpense,omitempty"`
	SummaVremennykhBonusovTurnover *Double   `json:"СуммаВременныхБонусовTurnover,omitempty"`
	SummaVremennykhBonusovReceipt  *Double   `json:"СуммаВременныхБонусовReceipt,omitempty"`
	SummaVremennykhBonusovExpense  *Double   `json:"СуммаВременныхБонусовExpense,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterSummyPokupokPoDiskontnymKartamBalanceAndTurnover struct {
	MemberCardKey                        *Guid     `json:"ДисконтнаяКарта_Key,omitempty"`
	DataSpisaniia                        *DateTime `json:"ДатаСписания,omitempty"`
	Period                               *DateTime `json:"Period,omitempty"`
	SecondPeriod                         *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                         *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                           *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                            *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                           *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                        *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                          *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                        *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                       *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                           *DateTime `json:"YearPeriod,omitempty"`
	Recorder                             *String   `json:"Recorder,omitempty"`
	LineNumber                           *Int64    `json:"LineNumber,omitempty"`
	SummaOpeningBalance                  *Double   `json:"СуммаOpeningBalance,omitempty"`
	SummaTurnover                        *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt                         *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense                         *Double   `json:"СуммаExpense,omitempty"`
	SummaClosingBalance                  *Double   `json:"СуммаClosingBalance,omitempty"`
	SummaBonusovOpeningBalance           *Double   `json:"СуммаБонусовOpeningBalance,omitempty"`
	SummaBonusovTurnover                 *Double   `json:"СуммаБонусовTurnover,omitempty"`
	SummaBonusovReceipt                  *Double   `json:"СуммаБонусовReceipt,omitempty"`
	SummaBonusovExpense                  *Double   `json:"СуммаБонусовExpense,omitempty"`
	SummaBonusovClosingBalance           *Double   `json:"СуммаБонусовClosingBalance,omitempty"`
	SummaVremennykhBonusovOpeningBalance *Double   `json:"СуммаВременныхБонусовOpeningBalance,omitempty"`
	SummaVremennykhBonusovTurnover       *Double   `json:"СуммаВременныхБонусовTurnover,omitempty"`
	SummaVremennykhBonusovReceipt        *Double   `json:"СуммаВременныхБонусовReceipt,omitempty"`
	SummaVremennykhBonusovExpense        *Double   `json:"СуммаВременныхБонусовExpense,omitempty"`
	SummaVremennykhBonusovClosingBalance *Double   `json:"СуммаВременныхБонусовClosingBalance,omitempty"`
	RecorderType                         *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterVypolneniePlanaProdazhRowType struct {
	Recorder      String    `json:"Recorder,omitempty"`
	Period        *DateTime `json:"Period,omitempty"`
	LineNumber    Int64     `json:"LineNumber,omitempty"`
	Active        *Boolean  `json:"Active,omitempty"`
	DepartmentKey *Guid     `json:"Склад_Key,omitempty"`
	SummaPlan     *Double   `json:"СуммаПлан,omitempty"`
	SummaFakt     *Double   `json:"СуммаФакт,omitempty"`
	RecorderType  String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterVypolneniePlanaProdazhTurnover struct {
	DepartmentKey     *Guid     `json:"Склад_Key,omitempty"`
	Period            *DateTime `json:"Period,omitempty"`
	SecondPeriod      *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod      *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod        *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod         *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod        *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod     *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod       *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod     *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod    *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod        *DateTime `json:"YearPeriod,omitempty"`
	Recorder          *String   `json:"Recorder,omitempty"`
	LineNumber        *Int64    `json:"LineNumber,omitempty"`
	SummaPlanTurnover *Double   `json:"СуммаПланTurnover,omitempty"`
	SummaFaktTurnover *Double   `json:"СуммаФактTurnover,omitempty"`
	RecorderType      *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDavalcheskiiMetallRowType struct {
	Recorder              String    `json:"Recorder,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	LineNumber            Int64     `json:"LineNumber,omitempty"`
	Active                *Boolean  `json:"Active,omitempty"`
	RecordType            *String   `json:"RecordType,omitempty"`
	OrganizatsiiaKey      *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	MetallKey             *Guid     `json:"Металл_Key,omitempty"`
	Weight                *Double   `json:"Вес,omitempty"`
	RecorderType          String    `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDavalcheskiiMetallBalance struct {
	OrganizatsiiaKey      *Guid   `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	MetallKey             *Guid   `json:"Металл_Key,omitempty"`
	VesBalance            *Double `json:"ВесBalance,omitempty"`
}
type AccumulationRegisterDavalcheskiiMetallTurnover struct {
	OrganizatsiiaKey      *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	MetallKey             *Guid     `json:"Металл_Key,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	SecondPeriod          *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod          *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod            *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod             *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod            *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod         *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod           *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod         *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod        *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod            *DateTime `json:"YearPeriod,omitempty"`
	Recorder              *String   `json:"Recorder,omitempty"`
	LineNumber            *Int64    `json:"LineNumber,omitempty"`
	VesTurnover           *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt            *Double   `json:"ВесReceipt,omitempty"`
	VesExpense            *Double   `json:"ВесExpense,omitempty"`
	RecorderType          *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDavalcheskiiMetallBalanceAndTurnover struct {
	OrganizatsiiaKey      *Guid     `json:"Организация_Key,omitempty"`
	DogovorKontragentaKey *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	MetallKey             *Guid     `json:"Металл_Key,omitempty"`
	Period                *DateTime `json:"Period,omitempty"`
	SecondPeriod          *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod          *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod            *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod             *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod            *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod         *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod           *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod         *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod        *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod            *DateTime `json:"YearPeriod,omitempty"`
	Recorder              *String   `json:"Recorder,omitempty"`
	LineNumber            *Int64    `json:"LineNumber,omitempty"`
	VesOpeningBalance     *Double   `json:"ВесOpeningBalance,omitempty"`
	VesTurnover           *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt            *Double   `json:"ВесReceipt,omitempty"`
	VesExpense            *Double   `json:"ВесExpense,omitempty"`
	VesClosingBalance     *Double   `json:"ВесClosingBalance,omitempty"`
	RecorderType          *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaRowType struct {
	Recorder                 String    `json:"Recorder,omitempty"`
	Period                   *DateTime `json:"Period,omitempty"`
	LineNumber               Int64     `json:"LineNumber,omitempty"`
	Active                   *Boolean  `json:"Active,omitempty"`
	RecordType               *String   `json:"RecordType,omitempty"`
	OrganizatsiiaKey         *Guid     `json:"Организация_Key,omitempty"`
	VidDenezhnykhSredstv     *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String   `json:"БанковскийСчетКасса,omitempty"`
	Sum                      *Double   `json:"Сумма,omitempty"`
	SummaUpr                 *Double   `json:"СуммаУпр,omitempty"`
	RecorderType             String    `json:"Recorder_Type,omitempty"`
	BankovskiiSchetKassaType *String   `json:"БанковскийСчетКасса_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaBalance struct {
	OrganizatsiiaKey         *Guid   `json:"Организация_Key,omitempty"`
	VidDenezhnykhSredstv     *String `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String `json:"БанковскийСчетКасса,omitempty"`
	SummaBalance             *Double `json:"СуммаBalance,omitempty"`
	SummaUprBalance          *Double `json:"СуммаУпрBalance,omitempty"`
	BankovskiiSchetKassaType *String `json:"БанковскийСчетКасса_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaTurnover struct {
	OrganizatsiiaKey         *Guid     `json:"Организация_Key,omitempty"`
	VidDenezhnykhSredstv     *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String   `json:"БанковскийСчетКасса,omitempty"`
	Period                   *DateTime `json:"Period,omitempty"`
	SecondPeriod             *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod             *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod               *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod               *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod            *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod              *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod            *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod           *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod               *DateTime `json:"YearPeriod,omitempty"`
	Recorder                 *String   `json:"Recorder,omitempty"`
	LineNumber               *Int64    `json:"LineNumber,omitempty"`
	SummaTurnover            *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt             *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense             *Double   `json:"СуммаExpense,omitempty"`
	SummaUprTurnover         *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt          *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense          *Double   `json:"СуммаУпрExpense,omitempty"`
	BankovskiiSchetKassaType *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	RecorderType             *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaBalanceAndTurnover struct {
	OrganizatsiiaKey         *Guid     `json:"Организация_Key,omitempty"`
	VidDenezhnykhSredstv     *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String   `json:"БанковскийСчетКасса,omitempty"`
	Period                   *DateTime `json:"Period,omitempty"`
	SecondPeriod             *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod             *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod               *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod               *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod            *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod              *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod            *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod           *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod               *DateTime `json:"YearPeriod,omitempty"`
	Recorder                 *String   `json:"Recorder,omitempty"`
	LineNumber               *Int64    `json:"LineNumber,omitempty"`
	SummaOpeningBalance      *Double   `json:"СуммаOpeningBalance,omitempty"`
	SummaTurnover            *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt             *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense             *Double   `json:"СуммаExpense,omitempty"`
	SummaClosingBalance      *Double   `json:"СуммаClosingBalance,omitempty"`
	SummaUprOpeningBalance   *Double   `json:"СуммаУпрOpeningBalance,omitempty"`
	SummaUprTurnover         *Double   `json:"СуммаУпрTurnover,omitempty"`
	SummaUprReceipt          *Double   `json:"СуммаУпрReceipt,omitempty"`
	SummaUprExpense          *Double   `json:"СуммаУпрExpense,omitempty"`
	SummaUprClosingBalance   *Double   `json:"СуммаУпрClosingBalance,omitempty"`
	BankovskiiSchetKassaType *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	RecorderType             *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryPeredannyeRowType struct {
	Recorder                       String    `json:"Recorder,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Active                         *Boolean  `json:"Active,omitempty"`
	RecordType                     *String   `json:"RecordType,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                         *String   `json:"Сделка,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	SummaVzaimoraschetov           *Double   `json:"СуммаВзаиморасчетов,omitempty"`
	RecorderType                   String    `json:"Recorder_Type,omitempty"`
	SdelkaType                     *String   `json:"Сделка_Type,omitempty"`
}
type AccumulationRegisterTovaryPeredannyeBalance struct {
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey          *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                         *String `json:"Сделка,omitempty"`
	KolichestvoBalance             *Int64  `json:"КоличествоBalance,omitempty"`
	VesBalance                     *Double `json:"ВесBalance,omitempty"`
	SummaVzaimoraschetovBalance    *Double `json:"СуммаВзаиморасчетовBalance,omitempty"`
	SdelkaType                     *String `json:"Сделка_Type,omitempty"`
}
type AccumulationRegisterTovaryPeredannyeTurnover struct {
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                         *String   `json:"Сделка,omitempty"`
	Period                         *DateTime `json:"Period,omitempty"`
	SecondPeriod                   *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                   *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                     *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                      *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                     *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                  *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                    *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                  *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                 *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                     *DateTime `json:"YearPeriod,omitempty"`
	Recorder                       *String   `json:"Recorder,omitempty"`
	LineNumber                     *Int64    `json:"LineNumber,omitempty"`
	KolichestvoTurnover            *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt             *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense             *Int64    `json:"КоличествоExpense,omitempty"`
	VesTurnover                    *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                     *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                     *Double   `json:"ВесExpense,omitempty"`
	SummaVzaimoraschetovTurnover   *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt    *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense    *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SdelkaType                     *String   `json:"Сделка_Type,omitempty"`
	RecorderType                   *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterTovaryPeredannyeBalanceAndTurnover struct {
	ItemKey                            *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                        *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey     *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                            *Guid     `json:"Размер_Key,omitempty"`
	DogovorKontragentaKey              *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	Sdelka                             *String   `json:"Сделка,omitempty"`
	Period                             *DateTime `json:"Period,omitempty"`
	SecondPeriod                       *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod                       *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod                         *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                          *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod                         *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod                      *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod                        *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod                      *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod                     *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod                         *DateTime `json:"YearPeriod,omitempty"`
	Recorder                           *String   `json:"Recorder,omitempty"`
	LineNumber                         *Int64    `json:"LineNumber,omitempty"`
	KolichestvoOpeningBalance          *Int64    `json:"КоличествоOpeningBalance,omitempty"`
	KolichestvoTurnover                *Int64    `json:"КоличествоTurnover,omitempty"`
	KolichestvoReceipt                 *Int64    `json:"КоличествоReceipt,omitempty"`
	KolichestvoExpense                 *Int64    `json:"КоличествоExpense,omitempty"`
	KolichestvoClosingBalance          *Int64    `json:"КоличествоClosingBalance,omitempty"`
	VesOpeningBalance                  *Double   `json:"ВесOpeningBalance,omitempty"`
	VesTurnover                        *Double   `json:"ВесTurnover,omitempty"`
	VesReceipt                         *Double   `json:"ВесReceipt,omitempty"`
	VesExpense                         *Double   `json:"ВесExpense,omitempty"`
	VesClosingBalance                  *Double   `json:"ВесClosingBalance,omitempty"`
	SummaVzaimoraschetovOpeningBalance *Double   `json:"СуммаВзаиморасчетовOpeningBalance,omitempty"`
	SummaVzaimoraschetovTurnover       *Double   `json:"СуммаВзаиморасчетовTurnover,omitempty"`
	SummaVzaimoraschetovReceipt        *Double   `json:"СуммаВзаиморасчетовReceipt,omitempty"`
	SummaVzaimoraschetovExpense        *Double   `json:"СуммаВзаиморасчетовExpense,omitempty"`
	SummaVzaimoraschetovClosingBalance *Double   `json:"СуммаВзаиморасчетовClosingBalance,omitempty"`
	SdelkaType                         *String   `json:"Сделка_Type,omitempty"`
	RecorderType                       *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRowType struct {
	Recorder                 String    `json:"Recorder,omitempty"`
	Period                   *DateTime `json:"Period,omitempty"`
	LineNumber               Int64     `json:"LineNumber,omitempty"`
	Active                   *Boolean  `json:"Active,omitempty"`
	RecordType               *String   `json:"RecordType,omitempty"`
	VidDenezhnykhSredstv     *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String   `json:"БанковскийСчетКасса,omitempty"`
	DokumentSpisaniia        *String   `json:"ДокументСписания,omitempty"`
	TypeOfMovingMoneyKey     *Guid     `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	Sum                      *Double   `json:"Сумма,omitempty"`
	RecorderType             String    `json:"Recorder_Type,omitempty"`
	BankovskiiSchetKassaType *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	DokumentSpisaniiaType    *String   `json:"ДокументСписания_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalance struct {
	VidDenezhnykhSredstv     *String `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String `json:"БанковскийСчетКасса,omitempty"`
	DokumentSpisaniia        *String `json:"ДокументСписания,omitempty"`
	TypeOfMovingMoneyKey     *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaBalance             *Double `json:"СуммаBalance,omitempty"`
	BankovskiiSchetKassaType *String `json:"БанковскийСчетКасса_Type,omitempty"`
	DokumentSpisaniiaType    *String `json:"ДокументСписания_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKSpisaniiuTurnover struct {
	VidDenezhnykhSredstv     *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String   `json:"БанковскийСчетКасса,omitempty"`
	DokumentSpisaniia        *String   `json:"ДокументСписания,omitempty"`
	TypeOfMovingMoneyKey     *Guid     `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	Period                   *DateTime `json:"Period,omitempty"`
	SecondPeriod             *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod             *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod               *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod               *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod            *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod              *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod            *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod           *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod               *DateTime `json:"YearPeriod,omitempty"`
	Recorder                 *String   `json:"Recorder,omitempty"`
	LineNumber               *Int64    `json:"LineNumber,omitempty"`
	SummaTurnover            *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt             *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense             *Double   `json:"СуммаExpense,omitempty"`
	BankovskiiSchetKassaType *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	DokumentSpisaniiaType    *String   `json:"ДокументСписания_Type,omitempty"`
	RecorderType             *String   `json:"Recorder_Type,omitempty"`
}
type AccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalanceAndTurnover struct {
	VidDenezhnykhSredstv     *String   `json:"ВидДенежныхСредств,omitempty"`
	BankovskiiSchetKassa     *String   `json:"БанковскийСчетКасса,omitempty"`
	DokumentSpisaniia        *String   `json:"ДокументСписания,omitempty"`
	TypeOfMovingMoneyKey     *Guid     `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	Period                   *DateTime `json:"Period,omitempty"`
	SecondPeriod             *DateTime `json:"SecondPeriod,omitempty"`
	MinutePeriod             *DateTime `json:"MinutePeriod,omitempty"`
	HourPeriod               *DateTime `json:"HourPeriod,omitempty"`
	DayPeriod                *DateTime `json:"DayPeriod,omitempty"`
	WeekPeriod               *DateTime `json:"WeekPeriod,omitempty"`
	TenDaysPeriod            *DateTime `json:"TenDaysPeriod,omitempty"`
	MonthPeriod              *DateTime `json:"MonthPeriod,omitempty"`
	QuarterPeriod            *DateTime `json:"QuarterPeriod,omitempty"`
	HalfYearPeriod           *DateTime `json:"HalfYearPeriod,omitempty"`
	YearPeriod               *DateTime `json:"YearPeriod,omitempty"`
	Recorder                 *String   `json:"Recorder,omitempty"`
	LineNumber               *Int64    `json:"LineNumber,omitempty"`
	SummaOpeningBalance      *Double   `json:"СуммаOpeningBalance,omitempty"`
	SummaTurnover            *Double   `json:"СуммаTurnover,omitempty"`
	SummaReceipt             *Double   `json:"СуммаReceipt,omitempty"`
	SummaExpense             *Double   `json:"СуммаExpense,omitempty"`
	SummaClosingBalance      *Double   `json:"СуммаClosingBalance,omitempty"`
	BankovskiiSchetKassaType *String   `json:"БанковскийСчетКасса_Type,omitempty"`
	DokumentSpisaniiaType    *String   `json:"ДокументСписания_Type,omitempty"`
	RecorderType             *String   `json:"Recorder_Type,omitempty"`
}
type DocumentChekKKMOplataRowType struct {
	Key                     Guid     `json:"Ref_Key,omitempty"`
	LineNumber              Int64    `json:"LineNumber,omitempty"`
	VidOplatyKey            *Guid    `json:"ВидОплаты_Key,omitempty"`
	NomerVidaOplaty         *Int16   `json:"НомерВидаОплаты,omitempty"`
	ProtsentTorgovoiUstupki *Double  `json:"ПроцентТорговойУступки,omitempty"`
	Sum                     *Double  `json:"Сумма,omitempty"`
	SummaTorgovoiUstupki    *Double  `json:"СуммаТорговойУступки,omitempty"`
	Khesh                   *String  `json:"Хэш,omitempty"`
	KartaSberbanka          *Boolean `json:"КартаСбербанка,omitempty"`
	Poslednie4              *String  `json:"Последние4,omitempty"`
	KodRRN                  *String  `json:"КодRRN,omitempty"`
	Identifikator           *String  `json:"Идентификатор,omitempty"`
	TransactionId           *String  `json:"TransactionId,omitempty"`
}
type DocumentChekKKMOplataSertifikatamiRowType struct {
	Key                      Guid    `json:"Ref_Key,omitempty"`
	LineNumber               Int64   `json:"LineNumber,omitempty"`
	SertifikatKey            *Guid   `json:"Сертификат_Key,omitempty"`
	NomerSertifikata         *String `json:"НомерСертификата,omitempty"`
	Sum                      *Double `json:"Сумма,omitempty"`
	SummaPokupkiPogashennaia *Double `json:"СуммаПокупкиПогашенная,omitempty"`
}
type DocumentChekKKMProdazhaSertifikatovRowType struct {
	Key                                 Guid     `json:"Ref_Key,omitempty"`
	LineNumber                          Int64    `json:"LineNumber,omitempty"`
	SertifikatKey                       *Guid    `json:"Сертификат_Key,omitempty"`
	NomerSertifikata                    *String  `json:"НомерСертификата,omitempty"`
	Sum                                 *Double  `json:"Сумма,omitempty"`
	SrokDeistviiaSertifikataOgranichen  *Boolean `json:"СрокДействияСертификатаОграничен,omitempty"`
	KolichestvoDneiDeistviiaSertifikata *Int64   `json:"КоличествоДнейДействияСертификата,omitempty"`
}
type DocumentChekKKMTovaryRowType struct {
	Key                                        Guid     `json:"Ref_Key,omitempty"`
	LineNumber                                 Int64    `json:"LineNumber,omitempty"`
	MID                                        *String  `json:"Артикул,omitempty"`
	Weight                                     *Double  `json:"Вес,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidki     *String  `json:"ЗначениеУсловияАвтоматическойСкидки,omitempty"`
	Quantity                                   *Int64   `json:"Количество,omitempty"`
	ItemKey                                    *Guid    `json:"Номенклатура_Key,omitempty"`
	PercentAutoDiscount                        *Double  `json:"ПроцентАвтоматическойСкидки,omitempty"`
	PercentManualDiscount                      *Double  `json:"ПроцентРучнойСкидки,omitempty"`
	SizeKey                                    *Guid    `json:"Размер_Key,omitempty"`
	RegistratsiiaProdazhi                      *Boolean `json:"РегистрацияПродажи,omitempty"`
	InstanceKey                                *Guid    `json:"СерияНоменклатуры_Key,omitempty"`
	SumManualDiscount                          *Double  `json:"СуммаРучнойСкидки,omitempty"`
	DepartmentKey                              *Guid    `json:"Склад_Key,omitempty"`
	StavkaNDS                                  *String  `json:"СтавкаНДС,omitempty"`
	Sum                                        *Double  `json:"Сумма,omitempty"`
	SummaNDS                                   *Double  `json:"СуммаНДС,omitempty"`
	UslovieAvtomaticheskoiSkidki               *String  `json:"УсловиеАвтоматическойСкидки,omitempty"`
	KharakteristikaNomenklaturyKey             *Guid    `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                                       *Double  `json:"Цена,omitempty"`
	Shtrikhkod                                 *String  `json:"Штрихкод,omitempty"`
	OrderKey                                   *Guid    `json:"КлючПродажи_Key,omitempty"`
	KliuchSviazi                               *Int64   `json:"КлючСвязи,omitempty"`
	ProdazhaPodarka                            *Boolean `json:"ПродажаПодарка,omitempty"`
	SumAutoDiscount                            *Double  `json:"СуммаАвтоматическойСкидки,omitempty"`
	SummaBonusov                               *Double  `json:"СуммаБонусов,omitempty"`
	SostavStrokiDliaRassrochkiKey              *Guid    `json:"СоставСтрокиДляРассрочки_Key,omitempty"`
	Komitent                                   *String  `json:"Комитент,omitempty"`
	TelefonKomitenta                           *String  `json:"ТелефонКомитента,omitempty"`
	INNkomitenta                               *String  `json:"ИННкомитента,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType *String  `json:"ЗначениеУсловияАвтоматическойСкидки_Type,omitempty"`
}
type DocumentChekKKMDokumentyObmenaRowType struct {
	Key         Guid    `json:"Ref_Key,omitempty"`
	LineNumber  Int64   `json:"LineNumber,omitempty"`
	DokumentKey *Guid   `json:"Документ_Key,omitempty"`
	Sum         *Double `json:"Сумма,omitempty"`
}
type DocumentChekKKMDogovoraRassrochkiProdazhaRowType struct {
	Key                  Guid    `json:"Ref_Key,omitempty"`
	LineNumber           Int64   `json:"LineNumber,omitempty"`
	DogovorRassrochkiKey *Guid   `json:"ДоговорРассрочки_Key,omitempty"`
	Sum                  *Double `json:"Сумма,omitempty"`
}
type DocumentChekKKMDogovoraRassrochkiOplataRowType struct {
	Key                  Guid    `json:"Ref_Key,omitempty"`
	LineNumber           Int64   `json:"LineNumber,omitempty"`
	DogovorRassrochkiKey *Guid   `json:"ДоговорРассрочки_Key,omitempty"`
	Sum                  *Double `json:"Сумма,omitempty"`
}
type DocumentChekKKMOplataBallamiRowType struct {
	Key                    Guid    `json:"Ref_Key,omitempty"`
	LineNumber             Int64   `json:"LineNumber,omitempty"`
	Khesh                  *String `json:"Хэш,omitempty"`
	Poslednie4             *String `json:"Последние4,omitempty"`
	Sum                    *Double `json:"Сумма,omitempty"`
	Identifikator          *String `json:"Идентификатор,omitempty"`
	TransactionId          *String `json:"TransactionId,omitempty"`
	TransactionIdSpisaniia *String `json:"TransactionIdСписания,omitempty"`
}
type DocumentChekKKMSkidkiNatsenkiRowType struct {
	Key               Guid    `json:"Ref_Key,omitempty"`
	LineNumber        Int64   `json:"LineNumber,omitempty"`
	KliuchSviazi      *Int64  `json:"КлючСвязи,omitempty"`
	Sum               *Double `json:"Сумма,omitempty"`
	SkidkaNatsenkaKey *Guid   `json:"СкидкаНаценка_Key,omitempty"`
}
type DocumentChekKKMUpravliaemyeSkidkiRowType struct {
	Key               Guid  `json:"Ref_Key,omitempty"`
	LineNumber        Int64 `json:"LineNumber,omitempty"`
	SkidkaNatsenkaKey *Guid `json:"СкидкаНаценка_Key,omitempty"`
}
type DocumentChekKKMPodarkiRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	SkidkaNatsenkaKey              *Guid   `json:"СкидкаНаценка_Key,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	Quantity                       *Double `json:"Количество,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	KliuchSviazi                   *Int64  `json:"КлючСвязи,omitempty"`
}
type DocumentChekKKMKuponyRowType struct {
	Key          Guid     `json:"Ref_Key,omitempty"`
	LineNumber   Int64    `json:"LineNumber,omitempty"`
	KliuchSviazi *Int64   `json:"КлючСвязи,omitempty"`
	KuponKey     *Guid    `json:"Купон_Key,omitempty"`
	NomerKupona  *String  `json:"НомерКупона,omitempty"`
	Spisyvat     *Boolean `json:"Списывать,omitempty"`
}
type CatalogTipySkidokNatsenokVremiaPoDniamNedeliRowType struct {
	Key               Guid      `json:"Ref_Key,omitempty"`
	LineNumber        Int64     `json:"LineNumber,omitempty"`
	VremiaNachala     *DateTime `json:"ВремяНачала,omitempty"`
	VremiaOkonchaniia *DateTime `json:"ВремяОкончания,omitempty"`
	Vybran            *Boolean  `json:"Выбран,omitempty"`
	DenNedeli         *String   `json:"ДеньНедели,omitempty"`
}
type CatalogVidyKodirovokiTsepeiElementyKodirovkiRowType struct {
	Key                            Guid     `json:"Ref_Key,omitempty"`
	LineNumber                     Int64    `json:"LineNumber,omitempty"`
	NeObiazatelnyiElement          *Boolean `json:"НеОбязательныйЭлемент,omitempty"`
	Element                        *String  `json:"Элемент,omitempty"`
	SvoistvoKey                    *Guid    `json:"Свойство_Key,omitempty"`
	DlinaElementa                  *Int64   `json:"ДлинаЭлемента,omitempty"`
	NomerStrokiTablChasti          *Int64   `json:"НомерСтрокиТаблЧасти,omitempty"`
	PeremennaiaDlina               *Boolean `json:"ПеременнаяДлина,omitempty"`
	ZapolniatSvoistvoVSootvetstvii *Boolean `json:"ЗаполнятьСвойствоВСоответствии,omitempty"`
}
type CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvRowType struct {
	Key                   Guid    `json:"Ref_Key,omitempty"`
	LineNumber            Int64   `json:"LineNumber,omitempty"`
	ZnachenieKodirovki    *String `json:"ЗначениеКодировки,omitempty"`
	ZnachenieSvoistvaKey  *Guid   `json:"ЗначениеСвойства_Key,omitempty"`
	NomerStrokiTablChasti *Int64  `json:"НомерСтрокиТаблЧасти,omitempty"`
	SvoistvoKey           *Guid   `json:"Свойство_Key,omitempty"`
}
type DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvaRowType struct {
	Key                   Guid    `json:"Ref_Key,omitempty"`
	LineNumber            Int64   `json:"LineNumber,omitempty"`
	VidOtchetaPoPlatezham *String `json:"ВидОтчетаПоПлатежам,omitempty"`
	StavkaNDS             *String `json:"СтавкаНДС,omitempty"`
	Sum                   *Double `json:"Сумма,omitempty"`
	SummaNDS              *Double `json:"СуммаНДС,omitempty"`
}
type DocumentOtchetKomitentuOProdazhakhTovaryRowType struct {
	Key                            Guid      `json:"Ref_Key,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	DokumentPostupleniia           *String   `json:"ДокументПоступления,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	OtchetKomitentuKey             *Guid     `json:"ОтчетКомитенту_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	Sum                            *Double   `json:"Сумма,omitempty"`
	SummaVoznagrazhdeniia          *Double   `json:"СуммаВознаграждения,omitempty"`
	SummaNDSVoznagrazhdeniia       *Double   `json:"СуммаНДСВознаграждения,omitempty"`
	SummaPostupleniia              *Double   `json:"СуммаПоступления,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double   `json:"Цена,omitempty"`
	TsenaPostupleniia              *Double   `json:"ЦенаПоступления,omitempty"`
	PokupatelKey                   *Guid     `json:"Покупатель_Key,omitempty"`
	DataRealizatsii                *DateTime `json:"ДатаРеализации,omitempty"`
	DokumentProdazhi               *String   `json:"ДокументПродажи,omitempty"`
	SchetFakturaVystavlennyiKey    *Guid     `json:"СчетФактураВыставленный_Key,omitempty"`
	DokumentPostupleniiaType       *String   `json:"ДокументПоступления_Type,omitempty"`
	DokumentProdazhiType           *String   `json:"ДокументПродажи_Type,omitempty"`
}
type DocumentZaiavkaNaPereotsenkuTovarovTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	RetailCost                     *Double `json:"ЦенаВРознице,omitempty"`
	TsenaVRoznitseGr               *Double `json:"ЦенаВРозницеГр,omitempty"`
	TsenaVRoznitseStaraia          *Double `json:"ЦенаВРозницеСтарая,omitempty"`
}
type DocumentZakrytieZakazovKlientovZakazyRowType struct {
	Key                        Guid  `json:"Ref_Key,omitempty"`
	LineNumber                 Int64 `json:"LineNumber,omitempty"`
	ZakazKlientaKey            *Guid `json:"ЗаказКлиента_Key,omitempty"`
	PrichinaZakrytiiaZakazaKey *Guid `json:"ПричинаЗакрытияЗаказа_Key,omitempty"`
}
type DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentaRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Znachenie      *String `json:"Значение,omitempty"`
	Predstavlenie  *String `json:"Представление,omitempty"`
	Rekvizit       *String `json:"Реквизит,omitempty"`
	TipKontragenta *String `json:"ТипКонтрагента,omitempty"`
}
type CatalogNastroikiVypolneniiaObmenaNastroikiObmenaRowType struct {
	Key                      Guid    `json:"Ref_Key,omitempty"`
	LineNumber               Int64   `json:"LineNumber,omitempty"`
	VypolniaemoeDeistvie     *String `json:"ВыполняемоеДействие,omitempty"`
	NastroikaObmenaKey       *Guid   `json:"НастройкаОбмена_Key,omitempty"`
	VypolniaemoeDeistvieType *String `json:"ВыполняемоеДействие_Type,omitempty"`
}
type CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamiRowType struct {
	Key                Guid    `json:"Ref_Key,omitempty"`
	LineNumber         Int64   `json:"LineNumber,omitempty"`
	TekstSoobshcheniia *String `json:"ТекстСообщения,omitempty"`
}
type DocumentRealizatsiiaTovarovUslugTovaryRowType struct {
	Key                                        Guid    `json:"Ref_Key,omitempty"`
	LineNumber                                 Int64   `json:"LineNumber,omitempty"`
	Weight                                     *Double `json:"Вес,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidki     *String `json:"ЗначениеУсловияАвтоматическойСкидки,omitempty"`
	KachestvoKey                               *Guid   `json:"Качество_Key,omitempty"`
	Quantity                                   *Int64  `json:"Количество,omitempty"`
	ItemKey                                    *Guid   `json:"Номенклатура_Key,omitempty"`
	PercentAutoDiscount                        *Double `json:"ПроцентАвтоматическойСкидки,omitempty"`
	PercentManualDiscount                      *Double `json:"ПроцентРучнойСкидки,omitempty"`
	SizeKey                                    *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                                *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	DepartmentKey                              *Guid   `json:"Склад_Key,omitempty"`
	StavkaNDS                                  *String `json:"СтавкаНДС,omitempty"`
	Sum                                        *Double `json:"Сумма,omitempty"`
	SummaNDS                                   *Double `json:"СуммаНДС,omitempty"`
	UslovieAvtomaticheskoiSkidki               *String `json:"УсловиеАвтоматическойСкидки,omitempty"`
	KharakteristikaNomenklaturyKey             *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                                       *Double `json:"Цена,omitempty"`
	SumManualDiscount                          *Double `json:"СуммаРучнойСкидки,omitempty"`
	SumAutoDiscount                            *Double `json:"СуммаАвтоматическойСкидки,omitempty"`
	VesVstavok                                 *Double `json:"ВесВставок,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType *String `json:"ЗначениеУсловияАвтоматическойСкидки_Type,omitempty"`
}
type DocumentRealizatsiiaTovarovUslugUslugiRowType struct {
	Key                      Guid    `json:"Ref_Key,omitempty"`
	LineNumber               Int64   `json:"LineNumber,omitempty"`
	Quantity                 *Int64  `json:"Количество,omitempty"`
	ItemKey                  *Guid   `json:"Номенклатура_Key,omitempty"`
	ProtsentOtSummyDokumenta *Double `json:"ПроцентОтСуммыДокумента,omitempty"`
	PercentManualDiscount    *Double `json:"ПроцентРучнойСкидки,omitempty"`
	Soderzhanie              *String `json:"Содержание,omitempty"`
	StavkaNDS                *String `json:"СтавкаНДС,omitempty"`
	Sum                      *Double `json:"Сумма,omitempty"`
	SummaNDS                 *Double `json:"СуммаНДС,omitempty"`
	Cost                     *Double `json:"Цена,omitempty"`
	SumManualDiscount        *Double `json:"СуммаРучнойСкидки,omitempty"`
}
type DocumentSobytieStoronnieLitsaRowType struct {
	Key           Guid  `json:"Ref_Key,omitempty"`
	LineNumber    Int64 `json:"LineNumber,omitempty"`
	KontragentKey *Guid `json:"Контрагент_Key,omitempty"`
	LitsoKey      *Guid `json:"Лицо_Key,omitempty"`
}
type CatalogNastroikiOtchetovGruppyNastroekIPolzovateliRowType struct {
	Key         Guid    `json:"Ref_Key,omitempty"`
	LineNumber  Int64   `json:"LineNumber,omitempty"`
	Obieekt     *String `json:"Объект,omitempty"`
	ObieektType *String `json:"Объект_Type,omitempty"`
}
type CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkiRowType struct {
	Key                                             Guid    `json:"Ref_Key,omitempty"`
	LineNumber                                      Int64   `json:"LineNumber,omitempty"`
	Kod                                             *Int16  `json:"Код,omitempty"`
	KonSumma                                        *Double `json:"КонСумма,omitempty"`
	MaksimalnyiProtsentSummyPokupkiPriOplateBonusom *Double `json:"МаксимальныйПроцентСуммыПокупкиПриОплатеБонусом,omitempty"`
	Naimenovanie                                    *String `json:"Наименование,omitempty"`
	NachSumma                                       *Double `json:"НачСумма,omitempty"`
	ProtsentSkidki                                  *Double `json:"ПроцентСкидки,omitempty"`
}
type CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenokRowType struct {
	Key             Guid    `json:"Ref_Key,omitempty"`
	LineNumber      Int64   `json:"LineNumber,omitempty"`
	UtsenkaProtsent *Double `json:"УценкаПроцент,omitempty"`
	UtsenkaDnei     *Int64  `json:"УценкаДней,omitempty"`
}
type DocumentKorrektirovkaDolgaSummyDolgaRowType struct {
	Key                         Guid      `json:"Ref_Key,omitempty"`
	LineNumber                  Int64     `json:"LineNumber,omitempty"`
	DataDolga                   *DateTime `json:"ДатаДолга,omitempty"`
	DogovorKontragentaKey       *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	KratnostVzaimoraschetov     *Int64    `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov         *Double   `json:"КурсВзаиморасчетов,omitempty"`
	Sdelka                      *String   `json:"Сделка,omitempty"`
	UvelichenieDolgaKontragenta *Double   `json:"УвеличениеДолгаКонтрагента,omitempty"`
	UmenshenieDolgaKontragenta  *Double   `json:"УменьшениеДолгаКонтрагента,omitempty"`
	SdelkaType                  *String   `json:"Сделка_Type,omitempty"`
}
type DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhaRowType struct {
	Key                     Guid    `json:"Ref_Key,omitempty"`
	LineNumber              Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey   *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	KratnostVzaimoraschetov *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov     *Double `json:"КурсВзаиморасчетов,omitempty"`
	ProektKey               *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                  *String `json:"Сделка,omitempty"`
	TypeOfMovingMoneyKey    *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov    *Double `json:"СуммаВзаиморасчетов,omitempty"`
	Sum                     *Double `json:"СуммаПлатежа,omitempty"`
	SdelkaType              *String `json:"Сделка_Type,omitempty"`
}
type DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkiRowType struct {
	Key                     Guid    `json:"Ref_Key,omitempty"`
	LineNumber              Int64   `json:"LineNumber,omitempty"`
	MestoRazmeshcheniia     *String `json:"МестоРазмещения,omitempty"`
	Sum                     *Double `json:"СуммаПлатежа,omitempty"`
	MestoRazmeshcheniiaType *String `json:"МестоРазмещения_Type,omitempty"`
}
type DocumentZakrytieZakazovPostavshchikamZakazyRowType struct {
	Key                        Guid  `json:"Ref_Key,omitempty"`
	LineNumber                 Int64 `json:"LineNumber,omitempty"`
	ZakazPostavshchikuKey      *Guid `json:"ЗаказПоставщику_Key,omitempty"`
	PrichinaZakrytiiaZakazaKey *Guid `json:"ПричинаЗакрытияЗаказа_Key,omitempty"`
}
type DocumentAnketyKlientovDliaFinMonitoringaAnketyRowType struct {
	Key                        Guid     `json:"Ref_Key,omitempty"`
	LineNumber                 Int64    `json:"LineNumber,omitempty"`
	KontragentKey              *Guid    `json:"Контрагент_Key,omitempty"`
	DogovorKontragentaKey      *Guid    `json:"ДоговорКонтрагента_Key,omitempty"`
	VPerechne                  *Boolean `json:"ВПеречне,omitempty"`
	InostrannoePublichnoeLitso *Boolean `json:"ИностранноеПубличноеЛицо,omitempty"`
	DokumentSeriia             *String  `json:"ДокументСерия,omitempty"`
	DokumentNomer              *String  `json:"ДокументНомер,omitempty"`
}
type DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentInkassovoePorucheniePeredannoeRekvizityKontragentaRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Znachenie      *String `json:"Значение,omitempty"`
	Predstavlenie  *String `json:"Представление,omitempty"`
	Rekvizit       *String `json:"Реквизит,omitempty"`
	TipKontragenta *String `json:"ТипКонтрагента,omitempty"`
}
type DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniiaRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	Imia          *String `json:"Имя,omitempty"`
	Predstavlenie *String `json:"Представление,omitempty"`
}
type DocumentInternetZakazTovaryInternetZakazaRowType struct {
	Key                            Guid     `json:"Ref_Key,omitempty"`
	LineNumber                     Int64    `json:"LineNumber,omitempty"`
	ItemKey                        *Guid    `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid    `json:"Размер_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid    `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	InstanceKey                    *Guid    `json:"СерияНоменклатуры_Key,omitempty"`
	Quantity                       *Int64   `json:"Количество,omitempty"`
	Weight                         *Double  `json:"Вес,omitempty"`
	Cost                           *Double  `json:"Цена,omitempty"`
	Sum                            *Double  `json:"Сумма,omitempty"`
	StavkaNDS                      *String  `json:"СтавкаНДС,omitempty"`
	SummaNDS                       *Double  `json:"СуммаНДС,omitempty"`
	PodobranPolnostiu              *Boolean `json:"ПодобранПолностью,omitempty"`
	Otkaz                          *Boolean `json:"Отказ,omitempty"`
}
type DocumentInternetZakazTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	NomerStrokiZakaza              *Int16  `json:"НомерСтрокиЗаказа,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhBonusyRowType struct {
	Key                                 Guid     `json:"Ref_Key,omitempty"`
	LineNumber                          Int64    `json:"LineNumber,omitempty"`
	MemberCardKey                       *Guid    `json:"ДисконтнаяКарта_Key,omitempty"`
	NomerCheka                          *String  `json:"НомерЧека,omitempty"`
	SummaNachislennogoBonusa            *Double  `json:"СуммаНачисленногоБонуса,omitempty"`
	SummaNachislennogoBonusaRasschitana *Boolean `json:"СуммаНачисленногоБонусаРассчитана,omitempty"`
	SummaOplaty                         *Double  `json:"СуммаОплаты,omitempty"`
	SummaPokupki                        *Double  `json:"СуммаПокупки,omitempty"`
	OrderKey                            *Guid    `json:"КлючПродажи_Key,omitempty"`
	OpisanieKarty                       *String  `json:"ОписаниеКарты,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamiRowType struct {
	Key                                     Guid    `json:"Ref_Key,omitempty"`
	LineNumber                              Int64   `json:"LineNumber,omitempty"`
	BankKreditorKey                         *Guid   `json:"БанкКредитор_Key,omitempty"`
	VidOplatyKey                            *Guid   `json:"ВидОплаты_Key,omitempty"`
	DogovorVzaimoraschetovBankaKreditoraKey *Guid   `json:"ДоговорВзаиморасчетовБанкаКредитора_Key,omitempty"`
	NomerCheka                              *String `json:"НомерЧека,omitempty"`
	ProtsentBankovskoiKomissii              *Double `json:"ПроцентБанковскойКомиссии,omitempty"`
	Sum                                     *Double `json:"Сумма,omitempty"`
	SummaBankovskoiKomissii                 *Double `json:"СуммаБанковскойКомиссии,omitempty"`
	OrderKey                                *Guid   `json:"КлючПродажи_Key,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamiRowType struct {
	Key                     Guid    `json:"Ref_Key,omitempty"`
	LineNumber              Int64   `json:"LineNumber,omitempty"`
	VidOplatyKey            *Guid   `json:"ВидОплаты_Key,omitempty"`
	NomerCheka              *String `json:"НомерЧека,omitempty"`
	ProtsentTorgovoiUstupki *Double `json:"ПроцентТорговойУступки,omitempty"`
	Sum                     *Double `json:"Сумма,omitempty"`
	SummaTorgovoiUstupki    *Double `json:"СуммаТорговойУступки,omitempty"`
	OrderKey                *Guid   `json:"КлючПродажи_Key,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamiRowType struct {
	Key                                Guid     `json:"Ref_Key,omitempty"`
	LineNumber                         Int64    `json:"LineNumber,omitempty"`
	NomerCheka                         *String  `json:"НомерЧека,omitempty"`
	SertifikatKey                      *Guid    `json:"Сертификат_Key,omitempty"`
	SrokDeistviiaSertifikataOgranichen *Boolean `json:"СрокДействияСертификатаОграничен,omitempty"`
	SummaPokupkiPogashennaia           *Double  `json:"СуммаПокупкиПогашенная,omitempty"`
	SummaSertifikata                   *Double  `json:"СуммаСертификата,omitempty"`
	OrderKey                           *Guid    `json:"КлючПродажи_Key,omitempty"`
	NomerSertifikata                   *String  `json:"НомерСертификата,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovRowType struct {
	Key              Guid    `json:"Ref_Key,omitempty"`
	LineNumber       Int64   `json:"LineNumber,omitempty"`
	NomerCheka       *String `json:"НомерЧека,omitempty"`
	SertifikatKey    *Guid   `json:"Сертификат_Key,omitempty"`
	NomerSertifikata *String `json:"НомерСертификата,omitempty"`
	Sum              *Double `json:"Сумма,omitempty"`
	OrderKey         *Guid   `json:"КлючПродажи_Key,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhTovaryRowType struct {
	Key                                        Guid    `json:"Ref_Key,omitempty"`
	LineNumber                                 Int64   `json:"LineNumber,omitempty"`
	Weight                                     *Double `json:"Вес,omitempty"`
	MemberCardKey                              *Guid   `json:"ДисконтнаяКарта_Key,omitempty"`
	Kassir                                     *String `json:"Кассир,omitempty"`
	Quantity                                   *Int64  `json:"Количество,omitempty"`
	ItemKey                                    *Guid   `json:"Номенклатура_Key,omitempty"`
	NomerCheka                                 *String `json:"НомерЧека,omitempty"`
	PercentAutoDiscount                        *Double `json:"ПроцентАвтоматическойСкидки,omitempty"`
	PercentManualDiscount                      *Double `json:"ПроцентРучнойСкидки,omitempty"`
	SizeKey                                    *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                                *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	SumManualDiscount                          *Double `json:"СуммаРучнойСкидки,omitempty"`
	DepartmentKey                              *Guid   `json:"Склад_Key,omitempty"`
	StavkaNDS                                  *String `json:"СтавкаНДС,omitempty"`
	Sum                                        *Double `json:"Сумма,omitempty"`
	SummaNDS                                   *Double `json:"СуммаНДС,omitempty"`
	UslovieAvtomaticheskoiSkidki               *String `json:"УсловиеАвтоматическойСкидки,omitempty"`
	KharakteristikaNomenklaturyKey             *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                                       *Double `json:"Цена,omitempty"`
	OrderKey                                   *Guid   `json:"КлючПродажи_Key,omitempty"`
	SumAutoDiscount                            *Double `json:"СуммаАвтоматическойСкидки,omitempty"`
	KliuchSviazi                               *Int64  `json:"КлючСвязи,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidki     *String `json:"ЗначениеУсловияАвтоматическойСкидки,omitempty"`
	OpisanieKarty                              *String `json:"ОписаниеКарты,omitempty"`
	SostavStrokiDliaRassrochkiKey              *Guid   `json:"СоставСтрокиДляРассрочки_Key,omitempty"`
	KassirType                                 *String `json:"Кассир_Type,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType *String `json:"ЗначениеУсловияАвтоматическойСкидки_Type,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhaRowType struct {
	Key                  Guid    `json:"Ref_Key,omitempty"`
	LineNumber           Int64   `json:"LineNumber,omitempty"`
	NomerCheka           *String `json:"НомерЧека,omitempty"`
	Sum                  *Double `json:"Сумма,omitempty"`
	DogovorRassrochkiKey *Guid   `json:"ДоговорРассрочки_Key,omitempty"`
	OrderKey             *Guid   `json:"КлючПродажи_Key,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhDokumentyObmenaRowType struct {
	Key         Guid    `json:"Ref_Key,omitempty"`
	LineNumber  Int64   `json:"LineNumber,omitempty"`
	DokumentKey *Guid   `json:"Документ_Key,omitempty"`
	Sum         *Double `json:"Сумма,omitempty"`
	NomerCheka  *String `json:"НомерЧека,omitempty"`
	OrderKey    *Guid   `json:"КлючПродажи_Key,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplataRowType struct {
	Key                  Guid    `json:"Ref_Key,omitempty"`
	LineNumber           Int64   `json:"LineNumber,omitempty"`
	Sum                  *Double `json:"Сумма,omitempty"`
	NomerCheka           *String `json:"НомерЧека,omitempty"`
	DogovorRassrochkiKey *Guid   `json:"ДоговорРассрочки_Key,omitempty"`
	OrderKey             *Guid   `json:"КлючПродажи_Key,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhOplataBallamiRowType struct {
	Key                    Guid    `json:"Ref_Key,omitempty"`
	LineNumber             Int64   `json:"LineNumber,omitempty"`
	Khesh                  *String `json:"Хэш,omitempty"`
	Poslednie4             *String `json:"Последние4,omitempty"`
	Sum                    *Double `json:"Сумма,omitempty"`
	Identifikator          *String `json:"Идентификатор,omitempty"`
	TransactionId          *String `json:"TransactionId,omitempty"`
	TransactionIdSpisaniia *String `json:"TransactionIdСписания,omitempty"`
	NomerCheka             *String `json:"НомерЧека,omitempty"`
	OrderKey               *Guid   `json:"КлючПродажи_Key,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkiRowType struct {
	Key               Guid    `json:"Ref_Key,omitempty"`
	LineNumber        Int64   `json:"LineNumber,omitempty"`
	KliuchSviazi      *Int64  `json:"КлючСвязи,omitempty"`
	Sum               *Double `json:"Сумма,omitempty"`
	SkidkaNatsenkaKey *Guid   `json:"СкидкаНаценка_Key,omitempty"`
	OrderKey          *Guid   `json:"КлючПродажи_Key,omitempty"`
}
type DocumentOtchetORoznichnykhProdazhakhKuponyRowType struct {
	Key          Guid     `json:"Ref_Key,omitempty"`
	LineNumber   Int64    `json:"LineNumber,omitempty"`
	KliuchSviazi *Int64   `json:"КлючСвязи,omitempty"`
	KuponKey     *Guid    `json:"Купон_Key,omitempty"`
	NomerKupona  *String  `json:"НомерКупона,omitempty"`
	Spisyvat     *Boolean `json:"Списывать,omitempty"`
}
type DocumentOtmenaSkidokNomenklaturyDokumentyRowType struct {
	Key                            Guid  `json:"Ref_Key,omitempty"`
	LineNumber                     Int64 `json:"LineNumber,omitempty"`
	UstanovkaSkidokNomenklaturyKey *Guid `json:"УстановкаСкидокНоменклатуры_Key,omitempty"`
}
type DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentaRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Znachenie      *String `json:"Значение,omitempty"`
	Predstavlenie  *String `json:"Представление,omitempty"`
	Rekvizit       *String `json:"Реквизит,omitempty"`
	TipKontragenta *String `json:"ТипКонтрагента,omitempty"`
}
type DocumentKassovyiChekKorrektsiiOplataRowType struct {
	Key        Guid    `json:"Ref_Key,omitempty"`
	LineNumber Int64   `json:"LineNumber,omitempty"`
	TipOplaty  *String `json:"ТипОплаты,omitempty"`
	Sum        *Double `json:"Сумма,omitempty"`
}
type DocumentSchetNaOplatuPokupateliuTovaryRowType struct {
	Key                                        Guid    `json:"Ref_Key,omitempty"`
	LineNumber                                 Int64   `json:"LineNumber,omitempty"`
	Weight                                     *Double `json:"Вес,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidki     *String `json:"ЗначениеУсловияАвтоматическойСкидки,omitempty"`
	Quantity                                   *Int64  `json:"Количество,omitempty"`
	ItemKey                                    *Guid   `json:"Номенклатура_Key,omitempty"`
	PercentAutoDiscount                        *Double `json:"ПроцентАвтоматическойСкидки,omitempty"`
	PercentManualDiscount                      *Double `json:"ПроцентРучнойСкидки,omitempty"`
	SizeKey                                    *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                                *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                                  *String `json:"СтавкаНДС,omitempty"`
	Sum                                        *Double `json:"Сумма,omitempty"`
	SummaNDS                                   *Double `json:"СуммаНДС,omitempty"`
	UslovieAvtomaticheskoiSkidki               *String `json:"УсловиеАвтоматическойСкидки,omitempty"`
	KharakteristikaNomenklaturyKey             *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                                       *Double `json:"Цена,omitempty"`
	SumAutoDiscount                            *Double `json:"СуммаАвтоматическойСкидки,omitempty"`
	SumManualDiscount                          *Double `json:"СуммаРучнойСкидки,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType *String `json:"ЗначениеУсловияАвтоматическойСкидки_Type,omitempty"`
}
type DocumentSchetNaOplatuPokupateliuUslugiRowType struct {
	Key                   Guid    `json:"Ref_Key,omitempty"`
	LineNumber            Int64   `json:"LineNumber,omitempty"`
	Quantity              *Int64  `json:"Количество,omitempty"`
	ItemKey               *Guid   `json:"Номенклатура_Key,omitempty"`
	PercentManualDiscount *Double `json:"ПроцентРучнойСкидки,omitempty"`
	Soderzhanie           *String `json:"Содержание,omitempty"`
	StavkaNDS             *String `json:"СтавкаНДС,omitempty"`
	Sum                   *Double `json:"Сумма,omitempty"`
	SummaNDS              *Double `json:"СуммаНДС,omitempty"`
	Cost                  *Double `json:"Цена,omitempty"`
	SumManualDiscount     *Double `json:"СуммаРучнойСкидки,omitempty"`
}
type CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovRowType struct {
	Key                               Guid     `json:"Ref_Key,omitempty"`
	LineNumber                        Int64    `json:"LineNumber,omitempty"`
	VariantPoiskaNePodderzhivaetsia   *Boolean `json:"ВариантПоискаНеПоддерживается,omitempty"`
	ImiaNastroikiDliaAlgoritma        *String  `json:"ИмяНастройкиДляАлгоритма,omitempty"`
	ImiaNastroikiDliaPolzovatelia     *String  `json:"ИмяНастройкиДляПользователя,omitempty"`
	KodPravilaObmena                  *String  `json:"КодПравилаОбмена,omitempty"`
	NaimenovaniePravilaObmena         *String  `json:"НаименованиеПравилаОбмена,omitempty"`
	NastroikaNePodderzhivaetsia       *Boolean `json:"НастройкаНеПоддерживается,omitempty"`
	OpisanieNastroikiDliaPolzovatelia *String  `json:"ОписаниеНастройкиДляПользователя,omitempty"`
	EtoNastroikaDliaVygruzki          *Boolean `json:"ЭтоНастройкаДляВыгрузки,omitempty"`
}
type CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhRowType struct {
	Key                         Guid     `json:"Ref_Key,omitempty"`
	LineNumber                  Int64    `json:"LineNumber,omitempty"`
	VygruzhatDannye             *Boolean `json:"ВыгружатьДанные,omitempty"`
	VygruzhatPoSsylke           *Boolean `json:"ВыгружатьПоСсылке,omitempty"`
	KodPravilaVygruzki          *String  `json:"КодПравилаВыгрузки,omitempty"`
	KodPravilaObmena            *String  `json:"КодПравилаОбмена,omitempty"`
	NaimenovaniePravilaVygruzki *String  `json:"НаименованиеПравилаВыгрузки,omitempty"`
	NastroikaNePodderzhivaetsia *Boolean `json:"НастройкаНеПоддерживается,omitempty"`
	EtoNastroikaDliaVygruzki    *Boolean `json:"ЭтоНастройкаДляВыгрузки,omitempty"`
}
type CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamiRowType struct {
	Key                Guid    `json:"Ref_Key,omitempty"`
	LineNumber         Int64   `json:"LineNumber,omitempty"`
	TekstSoobshcheniia *String `json:"ТекстСообщения,omitempty"`
}
type DocumentVozvratTovarovPostavshchikuTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	DefektKey                      *Guid   `json:"Дефект_Key,omitempty"`
	DokumentPostupleniia           *String `json:"ДокументПоступления,omitempty"`
	ZakazKlientaKey                *Guid   `json:"ЗаказКлиента_Key,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	ProektKey                      *Guid   `json:"Проект_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	DokumentPostupleniiaType       *String `json:"ДокументПоступления_Type,omitempty"`
}
type DocumentInventarizatsiiaTovarovNaSkladeTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	VesUchet                       *Double `json:"ВесУчет,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	KolichestvoUchet               *Int64  `json:"КоличествоУчет,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	NomerVed                       *String `json:"НомерВед,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaRegl                      *Double `json:"СуммаРегл,omitempty"`
	SummaUchet                     *Double `json:"СуммаУчет,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	RetailCost                     *Double `json:"ЦенаВРознице,omitempty"`
	OtkloneniePoKolichestvu        *Int64  `json:"ОтклонениеПоКоличеству,omitempty"`
	OtkloneniePoVesu               *Double `json:"ОтклонениеПоВесу,omitempty"`
}
type DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	VidSravneniia *String `json:"ВидСравнения,omitempty"`
	Znachenie     *String `json:"Значение,omitempty"`
	ImiaPolia     *String `json:"ИмяПоля,omitempty"`
	ZnachenieType *String `json:"Значение_Type,omitempty"`
}
type DocumentInventarizatsiiaTovarovNaSkladeSertifikatyRowType struct {
	Key                     Guid    `json:"Ref_Key,omitempty"`
	LineNumber              Int64   `json:"LineNumber,omitempty"`
	SertifikatKey           *Guid   `json:"Сертификат_Key,omitempty"`
	Sum                     *Double `json:"Сумма,omitempty"`
	SummaUchet              *Double `json:"СуммаУчет,omitempty"`
	Quantity                *Double `json:"Количество,omitempty"`
	KolichestvoUchet        *Double `json:"КоличествоУчет,omitempty"`
	OtkloneniePoKolichestvu *Int64  `json:"ОтклонениеПоКоличеству,omitempty"`
}
type DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutiRowType struct {
	Key                            Guid      `json:"Ref_Key,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	Quantity                       *Double   `json:"Количество,omitempty"`
	DogovorKontragentaKey          *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentOprikhodovaniia        *String   `json:"ДокументОприходования,omitempty"`
	KontragentKey                  *Guid     `json:"Контрагент_Key,omitempty"`
	DataOtpravki                   *DateTime `json:"ДатаОтправки,omitempty"`
	DokumentOprikhodovaniiaType    *String   `json:"ДокументОприходования_Type,omitempty"`
}
type DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentPrikhodnyiKassovyiOrderOplataRowType struct {
	Key        Guid    `json:"Ref_Key,omitempty"`
	LineNumber Int64   `json:"LineNumber,omitempty"`
	TipOplaty  *String `json:"ТипОплаты,omitempty"`
	Sum        *Double `json:"Сумма,omitempty"`
}
type DocumentPrikhodnyiKassovyiOrderTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	SummaSkidki                    *Double `json:"СуммаСкидки,omitempty"`
	VidTovaraKKT                   *String `json:"ВидТовараККТ,omitempty"`
	TipOplatyTovaraKKT             *String `json:"ТипОплатыТовараККТ,omitempty"`
	SummaOsn                       *Double `json:"СуммаОсн,omitempty"`
	Komitent                       *String `json:"Комитент,omitempty"`
	TelefonKomitenta               *String `json:"ТелефонКомитента,omitempty"`
	INNkomitenta                   *String `json:"ИННкомитента,omitempty"`
	SummaOpl                       *Double `json:"СуммаОпл,omitempty"`
}
type DocumentVozvratMaterialovIzProizvodstvaTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
}
type DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaStaraia                   *Double `json:"СуммаСтарая,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	TsenaZaGramm                   *Double `json:"ЦенаЗаГрамм,omitempty"`
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamiRowType struct {
	Key                       Guid    `json:"Ref_Key,omitempty"`
	LineNumber                Int64   `json:"LineNumber,omitempty"`
	PokupatelKey              *Guid   `json:"Покупатель_Key,omitempty"`
	DogovorSPokupatelemKey    *Guid   `json:"ДоговорСПокупателем_Key,omitempty"`
	SupplierKey               *Guid   `json:"Поставщик_Key,omitempty"`
	DogovorSPostavshchikomKey *Guid   `json:"ДоговорСПоставщиком_Key,omitempty"`
	Sum                       *Double `json:"Сумма,omitempty"`
	SummaSebestoimost         *Double `json:"СуммаСебестоимость,omitempty"`
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyeRowType struct {
	Key                         Guid    `json:"Ref_Key,omitempty"`
	LineNumber                  Int64   `json:"LineNumber,omitempty"`
	ItemKey                     *Guid   `json:"Номенклатура_Key,omitempty"`
	InstanceKey                 *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	PokupatelKey                *Guid   `json:"Покупатель_Key,omitempty"`
	DogovorSPokupatelemKey      *Guid   `json:"ДоговорСПокупателем_Key,omitempty"`
	SupplierKey                 *Guid   `json:"Поставщик_Key,omitempty"`
	DogovorSPostavshchikomKey   *Guid   `json:"ДоговорСПоставщиком_Key,omitempty"`
	Quantity                    *Int64  `json:"Количество,omitempty"`
	Weight                      *Double `json:"Вес,omitempty"`
	SummaPostupleniia           *Double `json:"СуммаПоступления,omitempty"`
	SummaProdazhi               *Double `json:"СуммаПродажи,omitempty"`
	StatusRaskhoda              *String `json:"СтатусРасхода,omitempty"`
	DokumentProdazhi            *String `json:"ДокументПродажи,omitempty"`
	DokumentOprikhodovaniia     *String `json:"ДокументОприходования,omitempty"`
	DokumentProdazhiType        *String `json:"ДокументПродажи_Type,omitempty"`
	DokumentOprikhodovaniiaType *String `json:"ДокументОприходования_Type,omitempty"`
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamiRowType struct {
	Key                       Guid    `json:"Ref_Key,omitempty"`
	LineNumber                Int64   `json:"LineNumber,omitempty"`
	SupplierKey               *Guid   `json:"Поставщик_Key,omitempty"`
	DogovorSPostavshchikomKey *Guid   `json:"ДоговорСПоставщиком_Key,omitempty"`
	Sum                       *Double `json:"Сумма,omitempty"`
	SummaSebestoimost         *Double `json:"СуммаСебестоимость,omitempty"`
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhRowType struct {
	Key                         Guid    `json:"Ref_Key,omitempty"`
	LineNumber                  Int64   `json:"LineNumber,omitempty"`
	ItemKey                     *Guid   `json:"Номенклатура_Key,omitempty"`
	InstanceKey                 *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	SupplierKey                 *Guid   `json:"Поставщик_Key,omitempty"`
	DogovorSPostavshchikomKey   *Guid   `json:"ДоговорСПоставщиком_Key,omitempty"`
	Quantity                    *Int64  `json:"Количество,omitempty"`
	Weight                      *Double `json:"Вес,omitempty"`
	SummaPostupleniia           *Double `json:"СуммаПоступления,omitempty"`
	StatusRaskhoda              *String `json:"СтатусРасхода,omitempty"`
	DokumentOprikhodovaniia     *String `json:"ДокументОприходования,omitempty"`
	SummaPostupleniiaBezNDS     *Double `json:"СуммаПоступленияБезНДС,omitempty"`
	DokumentOprikhodovaniiaType *String `json:"ДокументОприходования_Type,omitempty"`
}
type DocumentGTDImportRazdelyRowType struct {
	Key                                    Guid     `json:"Ref_Key,omitempty"`
	LineNumber                             Int64    `json:"LineNumber,omitempty"`
	NDSVValiute                            *Boolean `json:"НДСВВалюте,omitempty"`
	PoshlinaVValiute                       *Boolean `json:"ПошлинаВВалюте,omitempty"`
	StavkaNDS                              *String  `json:"СтавкаНДС,omitempty"`
	StavkaPoshliny                         *Double  `json:"СтавкаПошлины,omitempty"`
	SummaNDS                               *Double  `json:"СуммаНДС,omitempty"`
	SummaPoshliny                          *Double  `json:"СуммаПошлины,omitempty"`
	TamozhennaiaStoimost                   *Double  `json:"ТаможеннаяСтоимость,omitempty"`
	TamozhennaiaStoimostVValiuteReglUcheta *Boolean `json:"ТаможеннаяСтоимостьВВалютеРеглУчета,omitempty"`
}
type DocumentGTDImportTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	DokumentPartii                 *String `json:"ДокументПартии,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	NomerRazdela                   *Int16  `json:"НомерРаздела,omitempty"`
	ProektKey                      *Guid   `json:"Проект_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	SummaPoshliny                  *Double `json:"СуммаПошлины,omitempty"`
	FakturnaiaStoimost             *Double `json:"ФактурнаяСтоимость,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	DokumentPartiiType             *String `json:"ДокументПартии_Type,omitempty"`
}
type DocumentAktSverkiTovaryRowType struct {
	Key                            Guid      `json:"Ref_Key,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	MID                            *String   `json:"Артикул,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	ZakazKlientaKey                *Guid     `json:"ЗаказКлиента_Key,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	Koef                           *Double   `json:"Коэф,omitempty"`
	NaborKey                       *Guid     `json:"Набор_Key,omitempty"`
	Naideno                        *Boolean  `json:"Найдено,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	NomerNabora                    *Int64    `json:"НомерНабора,omitempty"`
	Pasport                        *String   `json:"Паспорт,omitempty"`
	ProektKey                      *Guid     `json:"Проект_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	StavkaNDS                      *String   `json:"СтавкаНДС,omitempty"`
	Sum                            *Double   `json:"Сумма,omitempty"`
	SummaNDS                       *Double   `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double   `json:"Цена,omitempty"`
	Period                         *DateTime `json:"Период,omitempty"`
}
type CatalogFailyDopolnitelnyeRekvizityRowType struct {
	Key              Guid    `json:"Ref_Key,omitempty"`
	LineNumber       Int64   `json:"LineNumber,omitempty"`
	SvoistvoKey      *Guid   `json:"Свойство_Key,omitempty"`
	Znachenie        *String `json:"Значение,omitempty"`
	TekstovaiaStroka *String `json:"ТекстоваяСтрока,omitempty"`
	ZnachenieType    *String `json:"Значение_Type,omitempty"`
}
type CatalogFailySertifikatyShifrovaniiaRowType struct {
	Key                  Guid    `json:"Ref_Key,omitempty"`
	LineNumber           Int64   `json:"LineNumber,omitempty"`
	Otpechatok           *String `json:"Отпечаток,omitempty"`
	Predstavlenie        *String `json:"Представление,omitempty"`
	SertifikatBase64Data *Binary `json:"Сертификат_Base64Data,omitempty"`
	SertifikatType       *String `json:"Сертификат_Type,omitempty"`
	Sertifikat           *Stream `json:"Сертификат,omitempty"`
}
type CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisiRowType struct {
	Key               Guid     `json:"Ref_Key,omitempty"`
	LineNumber        Int64    `json:"LineNumber,omitempty"`
	Administrirovanie *Boolean `json:"Администрирование,omitempty"`
	Otpravka          *Boolean `json:"Отправка,omitempty"`
	PolzovatelKey     *Guid    `json:"Пользователь_Key,omitempty"`
	Transport         *Boolean `json:"Транспорт,omitempty"`
}
type DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaRowType struct {
	Key                     Guid    `json:"Ref_Key,omitempty"`
	LineNumber              Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey   *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	KratnostVzaimoraschetov *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov     *Double `json:"КурсВзаиморасчетов,omitempty"`
	ProektKey               *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                  *String `json:"Сделка,omitempty"`
	TypeOfMovingMoneyKey    *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov    *Double `json:"СуммаВзаиморасчетов,omitempty"`
	Sum                     *Double `json:"СуммаПлатежа,omitempty"`
	SdelkaType              *String `json:"Сделка_Type,omitempty"`
}
type DocumentSkupkaTovarovTovaryRowType struct {
	Key         Guid    `json:"Ref_Key,omitempty"`
	LineNumber  Int64   `json:"LineNumber,omitempty"`
	Weight      *Double `json:"Вес,omitempty"`
	ItemKey     *Guid   `json:"Номенклатура_Key,omitempty"`
	ObshchiiVes *Double `json:"ОбщийВес,omitempty"`
	Sum         *Double `json:"Сумма,omitempty"`
	TsenaZaGr   *Double `json:"ЦенаЗаГр,omitempty"`
	Opisanie    *String `json:"Описание,omitempty"`
}
type DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliamRowType struct {
	Key                 Guid    `json:"Ref_Key,omitempty"`
	LineNumber          Int64   `json:"LineNumber,omitempty"`
	PokupatelKey        *Guid   `json:"Покупатель_Key,omitempty"`
	SchetFakturaKey     *Guid   `json:"СчетФактура_Key,omitempty"`
	SubkomissionerKey   *Guid   `json:"Субкомиссионер_Key,omitempty"`
	Sum                 *Double `json:"Сумма,omitempty"`
	NDS                 *Double `json:"НДС,omitempty"`
	SummaUvelichenie    *Double `json:"СуммаУвеличение,omitempty"`
	SummaUmenshenie     *Double `json:"СуммаУменьшение,omitempty"`
	SummaNDSUvelichenie *Double `json:"СуммаНДСУвеличение,omitempty"`
	SummaNDSUmenshenie  *Double `json:"СуммаНДСУменьшение,omitempty"`
}
type CatalogfmKartochkaKontragentaDannyeKontragentaRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	Kliuch        *String `json:"Ключ,omitempty"`
	Znachenie     *String `json:"Значение,omitempty"`
	ZnachenieType *String `json:"Значение_Type,omitempty"`
}
type DocumentSpisanieProsrochennykhSertifikatovSertifikatyRowType struct {
	Key                  Guid      `json:"Ref_Key,omitempty"`
	LineNumber           Int64     `json:"LineNumber,omitempty"`
	Quantity             *Int64    `json:"Количество,omitempty"`
	SertifikatKey        *Guid     `json:"Сертификат_Key,omitempty"`
	SrokDeistviiaDo      *DateTime `json:"СрокДействияДо,omitempty"`
	Sum                  *Double   `json:"Сумма,omitempty"`
	OrganizatsiiaKey     *Guid     `json:"Организация_Key,omitempty"`
	DokumentProdazhi     *String   `json:"ДокументПродажи,omitempty"`
	DokumentProdazhiType *String   `json:"ДокументПродажи_Type,omitempty"`
}
type DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentyRowType struct {
	Key                   Guid  `json:"Ref_Key,omitempty"`
	LineNumber            Int64 `json:"LineNumber,omitempty"`
	DogovorKontragentaKey *Guid `json:"ДоговорКонтрагента_Key,omitempty"`
	KontragentKey         *Guid `json:"Контрагент_Key,omitempty"`
}
type DocumentAkkreditivPeredannyiRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentAkkreditivPeredannyiRekvizityKontragentaRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Znachenie      *String `json:"Значение,omitempty"`
	Predstavlenie  *String `json:"Представление,omitempty"`
	Rekvizit       *String `json:"Реквизит,omitempty"`
	TipKontragenta *String `json:"ТипКонтрагента,omitempty"`
}
type CatalogKontragentyVidyDeiatelnostiRowType struct {
	Key                Guid  `json:"Ref_Key,omitempty"`
	LineNumber         Int64 `json:"LineNumber,omitempty"`
	VidDeiatelnostiKey *Guid `json:"ВидДеятельности_Key,omitempty"`
}
type DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentaRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Znachenie      *String `json:"Значение,omitempty"`
	Predstavlenie  *String `json:"Представление,omitempty"`
	Rekvizit       *String `json:"Реквизит,omitempty"`
	TipKontragenta *String `json:"ТипКонтрагента,omitempty"`
}
type DocumentMarketingovaiaAktsiiaMagazinyRowType struct {
	Key        Guid  `json:"Ref_Key,omitempty"`
	LineNumber Int64 `json:"LineNumber,omitempty"`
	MagazinKey *Guid `json:"Магазин_Key,omitempty"`
}
type DocumentMarketingovaiaAktsiiaSkidkiNatsenkiRowType struct {
	Key               Guid      `json:"Ref_Key,omitempty"`
	LineNumber        Int64     `json:"LineNumber,omitempty"`
	DataNachala       *DateTime `json:"ДатаНачала,omitempty"`
	DataOkonchaniia   *DateTime `json:"ДатаОкончания,omitempty"`
	MagazinKey        *Guid     `json:"Магазин_Key,omitempty"`
	SkidkaNatsenkaKey *Guid     `json:"СкидкаНаценка_Key,omitempty"`
}
type DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupaRowType struct {
	Key                  Guid     `json:"Ref_Key,omitempty"`
	LineNumber           Int64    `json:"LineNumber,omitempty"`
	NomerNabora          *Int64   `json:"НомерНабора,omitempty"`
	VidDostupa           *String  `json:"ВидДоступа,omitempty"`
	ZnachenieDostupa     *String  `json:"ЗначениеДоступа,omitempty"`
	Chtenie              *Boolean `json:"Чтение,omitempty"`
	Dobavlenie           *Boolean `json:"Добавление,omitempty"`
	Izmenenie            *Boolean `json:"Изменение,omitempty"`
	Udalenie             *Boolean `json:"Удаление,omitempty"`
	ZnachenieDostupaType *String  `json:"ЗначениеДоступа_Type,omitempty"`
}
type CatalogStsenariiObmenovDannymiNastroikiObmenaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	UzelInformatsionnoiBazy          *String `json:"УзелИнформационнойБазы,omitempty"`
	VidTransportaObmena              *String `json:"ВидТранспортаОбмена,omitempty"`
	VypolniaemoeDeistvie             *String `json:"ВыполняемоеДействие,omitempty"`
	KolichestvoElementovVTranzaktsii *Int64  `json:"КоличествоЭлементовВТранзакции,omitempty"`
	UzelInformatsionnoiBazyType      *String `json:"УзелИнформационнойБазы_Type,omitempty"`
}
type CatalogNomenklaturaSostavLigaturyRowType struct {
	Key        Guid    `json:"Ref_Key,omitempty"`
	LineNumber Int64   `json:"LineNumber,omitempty"`
	Quantity   *Double `json:"Количество,omitempty"`
	ItemKey    *Guid   `json:"Номенклатура_Key,omitempty"`
}
type DocumentOprosVoprosyRowType struct {
	Key              Guid    `json:"Ref_Key,omitempty"`
	LineNumber       Int64   `json:"LineNumber,omitempty"`
	VoprosKey        *Guid   `json:"Вопрос_Key,omitempty"`
	Otvet            *String `json:"Ответ,omitempty"`
	TipovoiOtvet     *String `json:"ТиповойОтвет,omitempty"`
	TipovoiOtvetType *String `json:"ТиповойОтвет_Type,omitempty"`
}
type DocumentOprosSostavnoiOtvetRowType struct {
	Key                  Guid    `json:"Ref_Key,omitempty"`
	LineNumber           Int64   `json:"LineNumber,omitempty"`
	VoprosKey            *Guid   `json:"Вопрос_Key,omitempty"`
	VoprosVladeletsKey   *Guid   `json:"ВопросВладелец_Key,omitempty"`
	NomerStrokiVTablitse *Int64  `json:"НомерСтрокиВТаблице,omitempty"`
	Otvet                *String `json:"Ответ,omitempty"`
	TipovoiOtvet         *String `json:"ТиповойОтвет,omitempty"`
	TipovoiOtvetType     *String `json:"ТиповойОтвет_Type,omitempty"`
}
type DocumentPereotsenkaTovarovVNTTTovaryRowType struct {
	Key                            Guid     `json:"Ref_Key,omitempty"`
	LineNumber                     Int64    `json:"LineNumber,omitempty"`
	Weight                         *Double  `json:"Вес,omitempty"`
	Quantity                       *Int64   `json:"Количество,omitempty"`
	ItemKey                        *Guid    `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid    `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid    `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid    `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	RetailCost                     *Double  `json:"ЦенаВРознице,omitempty"`
	TsenaVRoznitseGr               *Double  `json:"ЦенаВРозницеГр,omitempty"`
	TsenaVRoznitseStaraia          *Double  `json:"ЦенаВРозницеСтарая,omitempty"`
	Naideno                        *Boolean `json:"Найдено,omitempty"`
	Dnei                           *Int64   `json:"Дней,omitempty"`
	DogovorKontragentaKey          *Guid    `json:"ДоговорКонтрагента_Key,omitempty"`
	ProtsentUtsenki                *Double  `json:"ПроцентУценки,omitempty"`
}
type CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGruppRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	VstavkiBase64Data              *Binary `json:"Вставки_Base64Data,omitempty"`
	MetallKey                      *Guid   `json:"Металл_Key,omitempty"`
	Naimenovanie                   *String `json:"Наименование,omitempty"`
	UsloviiaVkhozhdeniiaBase64Data *Binary `json:"УсловияВхождения_Base64Data,omitempty"`
	VstavkiType                    *String `json:"Вставки_Type,omitempty"`
	UsloviiaVkhozhdeniiaType       *String `json:"УсловияВхождения_Type,omitempty"`
	Vstavki                        *Stream `json:"Вставки,omitempty"`
	UsloviiaVkhozhdeniia           *Stream `json:"УсловияВхождения,omitempty"`
}
type CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriiRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	MiksyVstavokBase64Data         *Binary `json:"МиксыВставок_Base64Data,omitempty"`
	Naimenovanie                   *String `json:"Наименование,omitempty"`
	NomenklaturnaiaGruppaKey       *Guid   `json:"НоменклатурнаяГруппа_Key,omitempty"`
	UsloviiaVkhozhdeniiaBase64Data *Binary `json:"УсловияВхождения_Base64Data,omitempty"`
	MiksyVstavokType               *String `json:"МиксыВставок_Type,omitempty"`
	UsloviiaVkhozhdeniiaType       *String `json:"УсловияВхождения_Type,omitempty"`
	MiksyVstavok                   *Stream `json:"МиксыВставок,omitempty"`
	UsloviiaVkhozhdeniia           *Stream `json:"УсловияВхождения,omitempty"`
}
type CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiiRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	SvoistvoTovara                 *String `json:"СвойствоТовара,omitempty"`
	UsloviiaVkhozhdeniiaBase64Data *Binary `json:"УсловияВхождения_Base64Data,omitempty"`
	Naimenovanie                   *String `json:"Наименование,omitempty"`
	UsloviiaVkhozhdeniiaType       *String `json:"УсловияВхождения_Type,omitempty"`
	UsloviiaVkhozhdeniia           *Stream `json:"УсловияВхождения,omitempty"`
}
type DocumentPeremeshchenieTovarovSertifikatyRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	SertifikatKey *Guid   `json:"Сертификат_Key,omitempty"`
	Sum           *Double `json:"Сумма,omitempty"`
}
type DocumentPeremeshchenieTovarovTovaryRowType struct {
	Key                            Guid     `json:"Ref_Key,omitempty"`
	LineNumber                     Int64    `json:"LineNumber,omitempty"`
	Weight                         *Double  `json:"Вес,omitempty"`
	DokumentRezerva                *String  `json:"ДокументРезерва,omitempty"`
	KachestvoKey                   *Guid    `json:"Качество_Key,omitempty"`
	Quantity                       *Int64   `json:"Количество,omitempty"`
	ItemKey                        *Guid    `json:"Номенклатура_Key,omitempty"`
	ProtsentRoznichnoiNatsenki     *Double  `json:"ПроцентРозничнойНаценки,omitempty"`
	SizeKey                        *Guid    `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid    `json:"СерияНоменклатуры_Key,omitempty"`
	StoimostBezNDS                 *Double  `json:"СтоимостьБезНДС,omitempty"`
	StoimostSNDS                   *Double  `json:"СтоимостьСНДС,omitempty"`
	Sum                            *Double  `json:"Сумма,omitempty"`
	SummaPostupleniia              *Double  `json:"СуммаПоступления,omitempty"`
	KharakteristikaNomenklaturyKey *Guid    `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double  `json:"Цена,omitempty"`
	TsenaVRoznitseGr               *Double  `json:"ЦенаВРозницеГр,omitempty"`
	TsenaPostupleniia              *Double  `json:"ЦенаПоступления,omitempty"`
	Naideno                        *Boolean `json:"Найдено,omitempty"`
	InternetZakazKey               *Guid    `json:"ИнтернетЗаказ_Key,omitempty"`
	DokumentRezervaType            *String  `json:"ДокументРезерва_Type,omitempty"`
}
type DocumentPeremeshchenieTovarovSpisokZaiavokRowType struct {
	Key                     Guid    `json:"Ref_Key,omitempty"`
	LineNumber              Int64   `json:"LineNumber,omitempty"`
	ZaiavkaNaPeremeshchenie *String `json:"ЗаявкаНаПеремещение,omitempty"`
}
type DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvRowType struct {
	Key                  Guid    `json:"Ref_Key,omitempty"`
	LineNumber           Int64   `json:"LineNumber,omitempty"`
	ValiutaZaiavkaKey    *Guid   `json:"ВалютаЗаявка_Key,omitempty"`
	ZaiavkaKey           *Guid   `json:"Заявка_Key,omitempty"`
	KontragentKey        *Guid   `json:"Контрагент_Key,omitempty"`
	OstatokZaiavka       *Double `json:"ОстатокЗаявка,omitempty"`
	OstatokRazmeshchenie *Double `json:"ОстатокРазмещение,omitempty"`
	OstatokRezerv        *Double `json:"ОстатокРезерв,omitempty"`
	OtvetstvennyiKey     *Guid   `json:"Ответственный_Key,omitempty"`
}
type DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovRowType struct {
	Key                       Guid    `json:"Ref_Key,omitempty"`
	LineNumber                Int64   `json:"LineNumber,omitempty"`
	ABCKlassifikatsiia        *String `json:"ABCКлассификация,omitempty"`
	ABCKlassifikatsiiaStaraia *String `json:"ABCКлассификацияСтарая,omitempty"`
	ZnachenieParametra        *Double `json:"ЗначениеПараметра,omitempty"`
	KontragentKey             *Guid   `json:"Контрагент_Key,omitempty"`
	MenedzherKontragentaKey   *Guid   `json:"МенеджерКонтрагента_Key,omitempty"`
}
type DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatyRowType struct {
	Key              Guid    `json:"Ref_Key,omitempty"`
	LineNumber       Int64   `json:"LineNumber,omitempty"`
	SertifikatKey    *Guid   `json:"Сертификат_Key,omitempty"`
	Sum              *Double `json:"Сумма,omitempty"`
	SummaUchet       *Double `json:"СуммаУчет,omitempty"`
	Quantity         *Double `json:"Количество,omitempty"`
	KolichestvoUchet *Double `json:"КоличествоУчет,omitempty"`
}
type DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	VidSravneniia *String `json:"ВидСравнения,omitempty"`
	Znachenie     *String `json:"Значение,omitempty"`
	ImiaPolia     *String `json:"ИмяПоля,omitempty"`
	ZnachenieType *String `json:"Значение_Type,omitempty"`
}
type DocumentKorrektirovkaRealizatsiiTovaryRowType struct {
	Key                                           Guid    `json:"Ref_Key,omitempty"`
	LineNumber                                    Int64   `json:"LineNumber,omitempty"`
	VesDoKorrektirovki                            *Double `json:"ВесДоКорректировки,omitempty"`
	VesDoIzmeneniia                               *Double `json:"ВесДоИзменения,omitempty"`
	Weight                                        *Double `json:"Вес,omitempty"`
	KolichestvoDoKorrektirovki                    *Int64  `json:"КоличествоДоКорректировки,omitempty"`
	KolichestvoDoIzmeneniia                       *Int64  `json:"КоличествоДоИзменения,omitempty"`
	Quantity                                      *Int64  `json:"Количество,omitempty"`
	DokumentOprikhodovaniia                       *String `json:"ДокументОприходования,omitempty"`
	DokumentPartii                                *String `json:"ДокументПартии,omitempty"`
	KachestvoKey                                  *Guid   `json:"Качество_Key,omitempty"`
	ItemKey                                       *Guid   `json:"Номенклатура_Key,omitempty"`
	Pasport                                       *String `json:"Паспорт,omitempty"`
	RazmerDoKorrektirovkiKey                      *Guid   `json:"РазмерДоКорректировки_Key,omitempty"`
	RazmerDoIzmeneniiaKey                         *Guid   `json:"РазмерДоИзменения_Key,omitempty"`
	SizeKey                                       *Guid   `json:"Размер_Key,omitempty"`
	SeriiaNomenklaturyDoKorrektirovkiKey          *Guid   `json:"СерияНоменклатурыДоКорректировки_Key,omitempty"`
	SeriiaNomenklaturyDoIzmeneniiaKey             *Guid   `json:"СерияНоменклатурыДоИзменения_Key,omitempty"`
	InstanceKey                                   *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	DepartmentKey                                 *Guid   `json:"Склад_Key,omitempty"`
	StavkaNDSDoKorrektirovki                      *String `json:"СтавкаНДСДоКорректировки,omitempty"`
	StavkaNDSDoIzmeneniia                         *String `json:"СтавкаНДСДоИзменения,omitempty"`
	StavkaNDS                                     *String `json:"СтавкаНДС,omitempty"`
	StatusPartii                                  *String `json:"СтатусПартии,omitempty"`
	SummaDoKorrektirovki                          *Double `json:"СуммаДоКорректировки,omitempty"`
	SummaDoIzmeneniia                             *Double `json:"СуммаДоИзменения,omitempty"`
	Sum                                           *Double `json:"Сумма,omitempty"`
	SummaNDSDoKorrektirovki                       *Double `json:"СуммаНДСДоКорректировки,omitempty"`
	SummaNDSDoIzmeneniia                          *Double `json:"СуммаНДСДоИзменения,omitempty"`
	SummaNDS                                      *Double `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyDoKorrektirovkiKey *Guid   `json:"ХарактеристикаНоменклатурыДоКорректировки_Key,omitempty"`
	KharakteristikaNomenklaturyDoIzmeneniiaKey    *Guid   `json:"ХарактеристикаНоменклатурыДоИзменения_Key,omitempty"`
	KharakteristikaNomenklaturyKey                *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	TsenaDoKorrektirovki                          *Double `json:"ЦенаДоКорректировки,omitempty"`
	TsenaDoIzmeneniia                             *Double `json:"ЦенаДоИзменения,omitempty"`
	Cost                                          *Double `json:"Цена,omitempty"`
	DokumentOprikhodovaniiaType                   *String `json:"ДокументОприходования_Type,omitempty"`
	DokumentPartiiType                            *String `json:"ДокументПартии_Type,omitempty"`
}
type DocumentKorrektirovkaRealizatsiiUslugiRowType struct {
	Key                        Guid    `json:"Ref_Key,omitempty"`
	LineNumber                 Int64   `json:"LineNumber,omitempty"`
	KolichestvoDoKorrektirovki *Int64  `json:"КоличествоДоКорректировки,omitempty"`
	KolichestvoDoIzmeneniia    *Int64  `json:"КоличествоДоИзменения,omitempty"`
	Quantity                   *Int64  `json:"Количество,omitempty"`
	ItemKey                    *Guid   `json:"Номенклатура_Key,omitempty"`
	Soderzhanie                *String `json:"Содержание,omitempty"`
	StavkaNDSDoKorrektirovki   *String `json:"СтавкаНДСДоКорректировки,omitempty"`
	StavkaNDSDoIzmeneniia      *String `json:"СтавкаНДСДоИзменения,omitempty"`
	StavkaNDS                  *String `json:"СтавкаНДС,omitempty"`
	SummaDoKorrektirovki       *Double `json:"СуммаДоКорректировки,omitempty"`
	SummaDoIzmeneniia          *Double `json:"СуммаДоИзменения,omitempty"`
	Sum                        *Double `json:"Сумма,omitempty"`
	SummaNDSDoKorrektirovki    *Double `json:"СуммаНДСДоКорректировки,omitempty"`
	SummaNDSDoIzmeneniia       *Double `json:"СуммаНДСДоИзменения,omitempty"`
	SummaNDS                   *Double `json:"СуммаНДС,omitempty"`
	TsenaDoKorrektirovki       *Double `json:"ЦенаДоКорректировки,omitempty"`
	TsenaDoIzmeneniia          *Double `json:"ЦенаДоИзменения,omitempty"`
	Cost                       *Double `json:"Цена,omitempty"`
}
type DocumentDoverennostTovaryRowType struct {
	Key                         Guid    `json:"Ref_Key,omitempty"`
	LineNumber                  Int64   `json:"LineNumber,omitempty"`
	Weight                      *Double `json:"Вес,omitempty"`
	EdinitsaPoKlassifikatoruKey *Guid   `json:"ЕдиницаПоКлассификатору_Key,omitempty"`
	Quantity                    *Int64  `json:"Количество,omitempty"`
	NaimenovanieTovara          *String `json:"НаименованиеТовара,omitempty"`
	SizeKey                     *Guid   `json:"Размер_Key,omitempty"`
}
type CatalogShablonyZapolneniiaKUPrazdnichnyeDniRowType struct {
	Key                       Guid      `json:"Ref_Key,omitempty"`
	LineNumber                Int64     `json:"LineNumber,omitempty"`
	Den                       *DateTime `json:"День,omitempty"`
	KU                        *Double   `json:"КУ,omitempty"`
	Opisanie                  *String   `json:"Описание,omitempty"`
	TsvetFonaDliaKalendaria   *String   `json:"ЦветФонаДляКалендаря,omitempty"`
	TsvetTekstaDliaKalendaria *String   `json:"ЦветТекстаДляКалендаря,omitempty"`
}
type CatalogShablonyZapolneniiaKUKUNaNedeliuRowType struct {
	Key        Guid    `json:"Ref_Key,omitempty"`
	LineNumber Int64   `json:"LineNumber,omitempty"`
	DenNedeli  *Int16  `json:"ДеньНедели,omitempty"`
	KU         *Double `json:"КУ,omitempty"`
}
type CatalogShablonyZapolneniiaKUSalonyRowType struct {
	Key        Guid  `json:"Ref_Key,omitempty"`
	LineNumber Int64 `json:"LineNumber,omitempty"`
	SalonKey   *Guid `json:"Салон_Key,omitempty"`
}
type DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrinRowType struct {
	Key                           Guid    `json:"Ref_Key,omitempty"`
	LineNumber                    Int64   `json:"LineNumber,omitempty"`
	TovarnaiaGruppaKey            *Guid   `json:"ТоварнаяГруппа_Key,omitempty"`
	TovarnaiaKategoriiaKey        *Guid   `json:"ТоварнаяКатегория_Key,omitempty"`
	TovarnaiaPozitsiiaKey         *Guid   `json:"ТоварнаяПозиция_Key,omitempty"`
	TypeKey                       *Guid   `json:"ТипИзделия_Key,omitempty"`
	AnalitikaTipaIzdeliiaKey      *Guid   `json:"АналитикаТипаИзделия_Key,omitempty"`
	PloshchadVykladki             *Double `json:"ПлощадьВыкладки,omitempty"`
	NormativnaiaPloshchadVykladki *Double `json:"НормативнаяПлощадьВыкладки,omitempty"`
	KolichestvoIzdelii            *Int16  `json:"КоличествоИзделий,omitempty"`
}
type DocumentVozvratProduktsiiVProizvodstvoTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	ZakazKlientaKey                *Guid   `json:"ЗаказКлиента_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	StoimostRabot                  *Double `json:"СтоимостьРабот,omitempty"`
	StoimostVstavok                *Double `json:"СтоимостьВставок,omitempty"`
	StoimostMetalla                *Double `json:"СтоимостьМеталла,omitempty"`
	SummaNDSVstavok                *Double `json:"СуммаНДСВставок,omitempty"`
	SummaNDSMetalla                *Double `json:"СуммаНДСМеталла,omitempty"`
	SummaNDSRabot                  *Double `json:"СуммаНДСРабот,omitempty"`
	DefektKey                      *Guid   `json:"Дефект_Key,omitempty"`
	VesVstavok                     *Double `json:"ВесВставок,omitempty"`
}
type CatalogNastroikiRabochegoMestaPolzovateliaNastroikiRowType struct {
	Key                        Guid     `json:"Ref_Key,omitempty"`
	LineNumber                 Int64    `json:"LineNumber,omitempty"`
	Dostupnost                 *Boolean `json:"Доступность,omitempty"`
	DostupnostPechati          *Boolean `json:"ДоступностьПечати,omitempty"`
	DostupnostRedaktirovaniia  *Boolean `json:"ДоступностьРедактирования,omitempty"`
	DostupnostVersionirovaniia *Boolean `json:"ДоступностьВерсионирования,omitempty"`
	PutKElementuFormy          *String  `json:"ПутьКЭлементуФормы,omitempty"`
	TipObieekta                *String  `json:"ТипОбъекта,omitempty"`
	ElementFormyRabochegoMesta *String  `json:"ЭлементФормыРабочегоМеста,omitempty"`
}
type DocumentSpisanieTovarovTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	DokumentRezerva                *String `json:"ДокументРезерва,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	OrderKey                       *Guid   `json:"КлючПродажи_Key,omitempty"`
	SkidkaNatsenkaKey              *Guid   `json:"СкидкаНаценка_Key,omitempty"`
	DokumentRezervaType            *String `json:"ДокументРезерва_Type,omitempty"`
}
type DocumentSpisanieTovarovSertifikatyRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	SertifikatKey *Guid   `json:"Сертификат_Key,omitempty"`
	Sum           *Double `json:"Сумма,omitempty"`
}
type DocumentsmsSoobshcheniePoluchateliRowType struct {
	Key                Guid      `json:"Ref_Key,omitempty"`
	LineNumber         Int64     `json:"LineNumber,omitempty"`
	Kontragent         *String   `json:"Контрагент,omitempty"`
	NomerTelefona      *String   `json:"НомерТелефона,omitempty"`
	GUID               *String   `json:"GUID,omitempty"`
	TekstSoobshcheniia *String   `json:"ТекстСообщения,omitempty"`
	DataZaversheniia   *DateTime `json:"ДатаЗавершения,omitempty"`
	Millisekunda       *Int16    `json:"Миллисекунда,omitempty"`
	Status             *String   `json:"Статус,omitempty"`
	OpisanieOshibki    *String   `json:"ОписаниеОшибки,omitempty"`
	MemberCardKey      *Guid     `json:"ДисконтнаяКарта_Key,omitempty"`
}
type CatalogKalendariPlanirovaniiaProdazhKUPoDniamRowType struct {
	Key        Guid      `json:"Ref_Key,omitempty"`
	LineNumber Int64     `json:"LineNumber,omitempty"`
	DenGoda    *DateTime `json:"ДеньГода,omitempty"`
	KU         *Double   `json:"КУ,omitempty"`
}
type CatalogKalendariPlanirovaniiaProdazhSalonyRowType struct {
	Key        Guid  `json:"Ref_Key,omitempty"`
	LineNumber Int64 `json:"LineNumber,omitempty"`
	SalonKey   *Guid `json:"Салон_Key,omitempty"`
}
type CatalogTipovyeAnketyVoprosyAnketyRowType struct {
	Key        Guid  `json:"Ref_Key,omitempty"`
	LineNumber Int64 `json:"LineNumber,omitempty"`
	VoprosKey  *Guid `json:"Вопрос_Key,omitempty"`
	RazdelKey  *Guid `json:"Раздел_Key,omitempty"`
}
type DocumentNachislenieSpisanieBonusovDiskontnyeKartyRowType struct {
	Key                     Guid   `json:"Ref_Key,omitempty"`
	LineNumber              Int64  `json:"LineNumber,omitempty"`
	MemberCardKey           *Guid  `json:"ДисконтнаяКарта_Key,omitempty"`
	SummaBonusov            *Int64 `json:"СуммаБонусов,omitempty"`
	TekushchaiaSummaBonusov *Int64 `json:"ТекущаяСуммаБонусов,omitempty"`
}
type DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	NomerPlatezha                    *Int16  `json:"НомерПлатежа,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentaRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Znachenie      *String `json:"Значение,omitempty"`
	Predstavlenie  *String `json:"Представление,omitempty"`
	Rekvizit       *String `json:"Реквизит,omitempty"`
	TipKontragenta *String `json:"ТипКонтрагента,omitempty"`
}
type DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSRowType struct {
	Key                      Guid    `json:"Ref_Key,omitempty"`
	LineNumber               Int64   `json:"LineNumber,omitempty"`
	ValiutaPostuplenieKey    *Guid   `json:"ВалютаПоступление_Key,omitempty"`
	DokumentPlanirovaniiaKey *Guid   `json:"ДокументПланирования_Key,omitempty"`
	KontragentKey            *Guid   `json:"Контрагент_Key,omitempty"`
	OstatokPostuplenie       *Double `json:"ОстатокПоступление,omitempty"`
	OstatokRazmeshchenie     *Double `json:"ОстатокРазмещение,omitempty"`
	OtvetstvennyiKey         *Guid   `json:"Ответственный_Key,omitempty"`
}
type DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringaRowType struct {
	Key                 Guid    `json:"Ref_Key,omitempty"`
	LineNumber          Int64   `json:"LineNumber,omitempty"`
	DokumentUcheta      *String `json:"ДокументУчета,omitempty"`
	SummaOtgruzki       *Double `json:"СуммаОтгрузки,omitempty"`
	SummaOplaty         *Double `json:"СуммаОплаты,omitempty"`
	Sum                 *Double `json:"Сумма,omitempty"`
	SummaVozvrata       *Double `json:"СуммаВозврата,omitempty"`
	Comment             *String `json:"Комментарий,omitempty"`
	KodVidaDokumentaKey *Guid   `json:"КодВидаДокумента_Key,omitempty"`
	DokumentUchetaType  *String `json:"ДокументУчета_Type,omitempty"`
}
type DocumentOtchetPoFinMonitoringuDannyeDokumentaRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	Kliuch        *String `json:"Ключ,omitempty"`
	Znachenie     *String `json:"Значение,omitempty"`
	ZnachenieType *String `json:"Значение_Type,omitempty"`
}
type CatalogVersiiFailovElektronnyePodpisiRowType struct {
	Key                    Guid      `json:"Ref_Key,omitempty"`
	LineNumber             Int64     `json:"LineNumber,omitempty"`
	DataPodpisi            *DateTime `json:"ДатаПодписи,omitempty"`
	ImiaFailaPodpisi       *String   `json:"ИмяФайлаПодписи,omitempty"`
	Comment                *String   `json:"Комментарий,omitempty"`
	KomuVydanSertifikat    *String   `json:"КомуВыданСертификат,omitempty"`
	Otpechatok             *String   `json:"Отпечаток,omitempty"`
	PodpisBase64Data       *Binary   `json:"Подпись_Base64Data,omitempty"`
	UstanovivshiiPodpisKey *Guid     `json:"УстановившийПодпись_Key,omitempty"`
	SertifikatBase64Data   *Binary   `json:"Сертификат_Base64Data,omitempty"`
	PodpisType             *String   `json:"Подпись_Type,omitempty"`
	SertifikatType         *String   `json:"Сертификат_Type,omitempty"`
	Podpis                 *Stream   `json:"Подпись,omitempty"`
	Sertifikat             *Stream   `json:"Сертификат,omitempty"`
}
type DocumentUstanovkaTsenNomenklaturyTipyTsenRowType struct {
	Key        Guid  `json:"Ref_Key,omitempty"`
	LineNumber Int64 `json:"LineNumber,omitempty"`
	TipTsenKey *Guid `json:"ТипЦен_Key,omitempty"`
}
type DocumentUstanovkaTsenNomenklaturyTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	ValiutaKey                     *Guid   `json:"Валюта_Key,omitempty"`
	IndeksStrokiTablitsyTsen       *Int64  `json:"ИндексСтрокиТаблицыЦен,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	ProtsentSkidkiNatsenki         *Double `json:"ПроцентСкидкиНаценки,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	TipTsenKey                     *Guid   `json:"ТипЦен_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	SegmentNomenklaturyKey         *Guid   `json:"СегментНоменклатуры_Key,omitempty"`
}
type DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentaRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Znachenie      *String `json:"Значение,omitempty"`
	Predstavlenie  *String `json:"Представление,omitempty"`
	Rekvizit       *String `json:"Реквизит,omitempty"`
	TipKontragenta *String `json:"ТипКонтрагента,omitempty"`
}
type DocumentPreiskurantNaSkupkuProbyRowType struct {
	Key          Guid    `json:"Ref_Key,omitempty"`
	LineNumber   Int64   `json:"LineNumber,omitempty"`
	ProbeKey     *Guid   `json:"Проба_Key,omitempty"`
	TsenaZaGramm *Double `json:"ЦенаЗаГрамм,omitempty"`
}
type DocumentPeredachaMaterialovVProizvodstvoTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
}
type DocumentVnutrenniiZakazTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	RazmeshchenieKey               *Guid   `json:"Размещение_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
}
type CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnostRowType struct {
	Key                                       Guid    `json:"Ref_Key,omitempty"`
	LineNumber                                Int64   `json:"LineNumber,omitempty"`
	DopolnitelnyeParametryObrabotkiBase64Data *Binary `json:"ДополнительныеПараметрыОбработки_Base64Data,omitempty"`
	PredstavlenieKnopki                       *String `json:"ПредставлениеКнопки,omitempty"`
	PredstavlenieNastroekObrabotki            *String `json:"ПредставлениеНастроекОбработки,omitempty"`
	PredstavlenieObieekta                     *String `json:"ПредставлениеОбъекта,omitempty"`
	SsylkaObieekta                            *String `json:"СсылкаОбъекта,omitempty"`
	TablichnaiaChastImia                      *String `json:"ТабличнаяЧастьИмя,omitempty"`
	TablichnaiaChastPredstavlenie             *String `json:"ТабличнаяЧастьПредставление,omitempty"`
	KhranilishcheVneshneiObrabotkiBase64Data  *Binary `json:"ХранилищеВнешнейОбработки_Base64Data,omitempty"`
	DopolnitelnyeParametryObrabotkiType       *String `json:"ДополнительныеПараметрыОбработки_Type,omitempty"`
	SsylkaObieektaType                        *String `json:"СсылкаОбъекта_Type,omitempty"`
	KhranilishcheVneshneiObrabotkiType        *String `json:"ХранилищеВнешнейОбработки_Type,omitempty"`
	DopolnitelnyeParametryObrabotki           *Stream `json:"ДополнительныеПараметрыОбработки,omitempty"`
	KhranilishcheVneshneiObrabotki            *Stream `json:"ХранилищеВнешнейОбработки,omitempty"`
}
type CatalogDopolnitelnyeVneshnieObrabotkiKomandyRowType struct {
	Key                     Guid     `json:"Ref_Key,omitempty"`
	LineNumber              Int64    `json:"LineNumber,omitempty"`
	Identifikator           *String  `json:"Идентификатор,omitempty"`
	VariantZapuska          *String  `json:"ВариантЗапуска,omitempty"`
	Predstavlenie           *String  `json:"Представление,omitempty"`
	PokazyvatOpoveshchenie  *Boolean `json:"ПоказыватьОповещение,omitempty"`
	Modifikator             *String  `json:"Модификатор,omitempty"`
	ReglamentnoeZadanieGUID *Guid    `json:"РегламентноеЗаданиеGUID,omitempty"`
	Skryt                   *Boolean `json:"Скрыть,omitempty"`
	ZameniaemyeKomandy      *String  `json:"ЗаменяемыеКоманды,omitempty"`
}
type CatalogDopolnitelnyeVneshnieObrabotkiRazdelyRowType struct {
	Key               Guid    `json:"Ref_Key,omitempty"`
	LineNumber        Int64   `json:"LineNumber,omitempty"`
	RazdelKey         *Guid   `json:"Раздел_Key,omitempty"`
	UdalitImiaRazdela *String `json:"УдалитьИмяРаздела,omitempty"`
}
type CatalogDopolnitelnyeVneshnieObrabotkiNaznachenieRowType struct {
	Key                                 Guid    `json:"Ref_Key,omitempty"`
	LineNumber                          Int64   `json:"LineNumber,omitempty"`
	ObieektNaznacheniiaKey              *Guid   `json:"ОбъектНазначения_Key,omitempty"`
	UdalitPolnoeImiaObieektaMetadannykh *String `json:"УдалитьПолноеИмяОбъектаМетаданных,omitempty"`
}
type CatalogDopolnitelnyeVneshnieObrabotkiRazresheniiaRowType struct {
	Key                 Guid    `json:"Ref_Key,omitempty"`
	LineNumber          Int64   `json:"LineNumber,omitempty"`
	VidRazresheniia     *String `json:"ВидРазрешения,omitempty"`
	ParametryBase64Data *Binary `json:"Параметры_Base64Data,omitempty"`
	ParametryType       *String `json:"Параметры_Type,omitempty"`
	Parametry           *Stream `json:"Параметры,omitempty"`
}
type CatalogGruppyPolzovateleiPolzovateliGruppyRowType struct {
	Key           Guid  `json:"Ref_Key,omitempty"`
	LineNumber    Int64 `json:"LineNumber,omitempty"`
	PolzovatelKey *Guid `json:"Пользователь_Key,omitempty"`
}
type DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	DefektKey                      *Guid   `json:"Дефект_Key,omitempty"`
	DokumentPostupleniia           *String `json:"ДокументПоступления,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	ProektKey                      *Guid   `json:"Проект_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	RetailCost                     *Double `json:"ЦенаВРознице,omitempty"`
	DokumentPostupleniiaType       *String `json:"ДокументПоступления_Type,omitempty"`
}
type DocumentZaiavkaNaPeremeshchenieTovarovTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	ProtsentRoznichnoiNatsenki     *Double `json:"ПроцентРозничнойНаценки,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	StoimostBezNDS                 *Double `json:"СтоимостьБезНДС,omitempty"`
	StoimostSNDS                   *Double `json:"СтоимостьСНДС,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaPostupleniia              *Double `json:"СуммаПоступления,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	TsenaVRoznitseGr               *Double `json:"ЦенаВРозницеГр,omitempty"`
	TsenaPostupleniia              *Double `json:"ЦенаПоступления,omitempty"`
}
type DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoraRowType struct {
	Key                   Guid     `json:"Ref_Key,omitempty"`
	LineNumber            Int64    `json:"LineNumber,omitempty"`
	DogovorKontragentaKey *Guid    `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentUcheta        *String  `json:"ДокументУчета,omitempty"`
	ItogovaiaStroka       *Boolean `json:"ИтоговаяСтрока,omitempty"`
	KontragentKey         *Guid    `json:"Контрагент_Key,omitempty"`
	RuchnaiaKorrektirovka *Boolean `json:"РучнаяКорректировка,omitempty"`
	SummaOplaty           *Double  `json:"СуммаОплаты,omitempty"`
	SummaOtgruzki         *Double  `json:"СуммаОтгрузки,omitempty"`
	SummaVozvrata         *Double  `json:"СуммаВозврата,omitempty"`
	DokumentUchetaType    *String  `json:"ДокументУчета_Type,omitempty"`
}
type CatalogUsloviiaOplatyTablitsaVyplatRowType struct {
	Key             Guid    `json:"Ref_Key,omitempty"`
	LineNumber      Int64   `json:"LineNumber,omitempty"`
	PeriodOtsrochki *Int16  `json:"ПериодОтсрочки,omitempty"`
	ProtsentVyplaty *Double `json:"ПроцентВыплаты,omitempty"`
}
type DocumentUdalitNariadZakazIzdeliiaRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	Quantity                       *Double `json:"Количество,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	TypeKey                        *Guid   `json:"ТипИзделия_Key,omitempty"`
	ProbeKey                       *Guid   `json:"Проба_Key,omitempty"`
	Kommentarii                    *String `json:"Комментарии,omitempty"`
	VesBezVstavok                  *Double `json:"ВесБезВставок,omitempty"`
	NomerStrokiTCh                 *Int64  `json:"НомерСтрокиТЧ,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
}
type DocumentUdalitNariadZakazUslugiRowType struct {
	Key         Guid    `json:"Ref_Key,omitempty"`
	LineNumber  Int64   `json:"LineNumber,omitempty"`
	Quantity    *Int64  `json:"Количество,omitempty"`
	ItemKey     *Guid   `json:"Номенклатура_Key,omitempty"`
	Soderzhanie *String `json:"Содержание,omitempty"`
	StavkaNDS   *String `json:"СтавкаНДС,omitempty"`
	Sum         *Double `json:"Сумма,omitempty"`
	SummaNDS    *Double `json:"СуммаНДС,omitempty"`
	Cost        *Double `json:"Цена,omitempty"`
}
type DocumentUdalitNariadZakazSpetsifikatsiiaRowType struct {
	Key              Guid    `json:"Ref_Key,omitempty"`
	LineNumber       Int64   `json:"LineNumber,omitempty"`
	Weight           *Double `json:"Вес,omitempty"`
	GruppaDefektaKey *Guid   `json:"ГруппаДефекта_Key,omitempty"`
	GruppaTsvetaKey  *Guid   `json:"ГруппаЦвета_Key,omitempty"`
	KamenKey         *Guid   `json:"Камень_Key,omitempty"`
	Quantity         *Double `json:"Количество,omitempty"`
	ItemKey          *Guid   `json:"Номенклатура_Key,omitempty"`
	Razmer1          *Double `json:"Размер1,omitempty"`
	Razmer2          *Double `json:"Размер2,omitempty"`
	Razmer3          *Double `json:"Размер3,omitempty"`
	RassevKey        *Guid   `json:"Рассев_Key,omitempty"`
	FormaOgrankiKey  *Guid   `json:"ФормаОгранки_Key,omitempty"`
	TsvetKamniaKey   *Guid   `json:"ЦветКамня_Key,omitempty"`
	NomerStrokiTCh   *Int64  `json:"НомерСтрокиТЧ,omitempty"`
}
type DocumentUdalitNariadZakazMetallyRowType struct {
	Key        Guid    `json:"Ref_Key,omitempty"`
	LineNumber Int64   `json:"LineNumber,omitempty"`
	ItemKey    *Guid   `json:"Номенклатура_Key,omitempty"`
	ProbeKey   *Guid   `json:"Проба_Key,omitempty"`
	Weight     *String `json:"Вес,omitempty"`
}
type DocumentUdalitNariadZakazVstavkiRowType struct {
	Key              Guid    `json:"Ref_Key,omitempty"`
	LineNumber       Int64   `json:"LineNumber,omitempty"`
	Weight           *Double `json:"Вес,omitempty"`
	GruppaDefektaKey *Guid   `json:"ГруппаДефекта_Key,omitempty"`
	GruppaTsvetaKey  *Guid   `json:"ГруппаЦвета_Key,omitempty"`
	KamenKey         *Guid   `json:"Камень_Key,omitempty"`
	Quantity         *Double `json:"Количество,omitempty"`
	ItemKey          *Guid   `json:"Номенклатура_Key,omitempty"`
	Razmer1          *Double `json:"Размер1,omitempty"`
	Razmer2          *Double `json:"Размер2,omitempty"`
	Razmer3          *Double `json:"Размер3,omitempty"`
	RassevKey        *Guid   `json:"Рассев_Key,omitempty"`
	FormaOgrankiKey  *Guid   `json:"ФормаОгранки_Key,omitempty"`
	TsvetKamniaKey   *Guid   `json:"ЦветКамня_Key,omitempty"`
}
type DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostiRowType struct {
	Key                 Guid      `json:"Ref_Key,omitempty"`
	LineNumber          Int64     `json:"LineNumber,omitempty"`
	DataDolga           *DateTime `json:"ДатаДолга,omitempty"`
	NovaiaDataDolga     *DateTime `json:"НоваяДатаДолга,omitempty"`
	NovaiaSummaDolga    *Double   `json:"НоваяСуммаДолга,omitempty"`
	NovaiaSummaDolgaUpr *Double   `json:"НоваяСуммаДолгаУпр,omitempty"`
	SummaDolga          *Double   `json:"СуммаДолга,omitempty"`
	SummaDolgaUpr       *Double   `json:"СуммаДолгаУпр,omitempty"`
}
type DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentAkkreditivPoluchennyiRekvizityKontragentaRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Znachenie      *String `json:"Значение,omitempty"`
	Predstavlenie  *String `json:"Представление,omitempty"`
	Rekvizit       *String `json:"Реквизит,omitempty"`
	TipKontragenta *String `json:"ТипКонтрагента,omitempty"`
}
type DocumentPriemIzRemontaIzdeliiaRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	Quantity                       *Double `json:"Количество,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	TypeKey                        *Guid   `json:"ТипИзделия_Key,omitempty"`
	ProbeKey                       *Guid   `json:"Проба_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
}
type DocumentPriemIzRemontaMaterialyRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	KliuchNomenklaturyKey          *Guid   `json:"КлючНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	DokumentOprikhodovaniiaKey     *Guid   `json:"ДокументОприходования_Key,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
}
type DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvaRowType struct {
	Key                   Guid    `json:"Ref_Key,omitempty"`
	LineNumber            Int64   `json:"LineNumber,omitempty"`
	VidOtchetaPoPlatezham *String `json:"ВидОтчетаПоПлатежам,omitempty"`
	StavkaNDS             *String `json:"СтавкаНДС,omitempty"`
	Sum                   *Double `json:"Сумма,omitempty"`
	SummaNDS              *Double `json:"СуммаНДС,omitempty"`
}
type DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	DokumentPartiiKey              *Guid   `json:"ДокументПартии_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaVoznagrazhdeniia          *Double `json:"СуммаВознаграждения,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	SummaNDSVoznagrazhdeniia       *Double `json:"СуммаНДСВознаграждения,omitempty"`
	SummaNDSPeredachi              *Double `json:"СуммаНДСПередачи,omitempty"`
	SummaPeredachi                 *Double `json:"СуммаПередачи,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	TsenaPeredachi                 *Double `json:"ЦенаПередачи,omitempty"`
}
type CatalogfmAnketaKlientaDannyeKontragentaRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	Kliuch        *String `json:"Ключ,omitempty"`
	Znachenie     *String `json:"Значение,omitempty"`
	ZnachenieType *String `json:"Значение_Type,omitempty"`
}
type DocumentUstanovkaZnacheniiTochkiZakazaTovaryRowType struct {
	Key                                            Guid      `json:"Ref_Key,omitempty"`
	LineNumber                                     Int64     `json:"LineNumber,omitempty"`
	Weight                                         *Double   `json:"Вес,omitempty"`
	DataKon                                        *DateTime `json:"ДатаКон,omitempty"`
	DataNach                                       *DateTime `json:"ДатаНач,omitempty"`
	ZnachenieTochkiZakaza                          *Int64    `json:"ЗначениеТочкиЗаказа,omitempty"`
	MinimalnyiStrakhovoiZapas                      *Int64    `json:"МинимальныйСтраховойЗапас,omitempty"`
	ItemKey                                        *Guid     `json:"Номенклатура_Key,omitempty"`
	ProtsentnaiaStavkaZnacheniiaTochkiZakaza       *Int16    `json:"ПроцентнаяСтавкаЗначенияТочкиЗаказа,omitempty"`
	ProtsentnaiaStavkaMinimalnogoStrakhovogoZapasa *Int16    `json:"ПроцентнаяСтавкаМинимальногоСтраховогоЗапаса,omitempty"`
	SizeKey                                        *Guid     `json:"Размер_Key,omitempty"`
	DepartmentKey                                  *Guid     `json:"Склад_Key,omitempty"`
	SposobOpredeleniiaZnacheniiaTochkiZakaza       *String   `json:"СпособОпределенияЗначенияТочкиЗаказа,omitempty"`
}
type CatalogPravilaProdazhTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	KharakteristikiNomenklaturyKey *Guid   `json:"ХарактеристикиНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
}
type DocumentPostuplenieDopRaskhodovTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	DokumentPartii                 *String `json:"ДокументПартии,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	ProektKey                      *Guid   `json:"Проект_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	SummaRaspredeleniia            *Double `json:"СуммаРаспределения,omitempty"`
	SummaTovara                    *Double `json:"СуммаТовара,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	DokumentPartiiType             *String `json:"ДокументПартии_Type,omitempty"`
}
type DocumentAvansovyiOtchetVydannyeAvansyRowType struct {
	Key                        Guid    `json:"Ref_Key,omitempty"`
	LineNumber                 Int64   `json:"LineNumber,omitempty"`
	RaskhodnyiKassovyiOrderKey *Guid   `json:"РасходныйКассовыйОрдер_Key,omitempty"`
	Sum                        *Double `json:"Сумма,omitempty"`
}
type DocumentAvansovyiOtchetTovaryRowType struct {
	Key                            Guid      `json:"Ref_Key,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	VidDokVkhodiashchii            *String   `json:"ВидДокВходящий,omitempty"`
	DataVkhodiashchegoDokumenta    *DateTime `json:"ДатаВходящегоДокумента,omitempty"`
	DataSF                         *DateTime `json:"ДатаСФ,omitempty"`
	ZakazKlientaKey                *Guid     `json:"ЗаказКлиента_Key,omitempty"`
	KachestvoKey                   *Guid     `json:"Качество_Key,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	NomerVkhodiashchegoDokumenta   *String   `json:"НомерВходящегоДокумента,omitempty"`
	NomerSF                        *String   `json:"НомерСФ,omitempty"`
	SupplierKey                    *Guid     `json:"Поставщик_Key,omitempty"`
	PredieiavlenSF                 *Boolean  `json:"ПредъявленСФ,omitempty"`
	ProektKey                      *Guid     `json:"Проект_Key,omitempty"`
	ProtsentRoznichnoiNatsenki     *Double   `json:"ПроцентРозничнойНаценки,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	DepartmentKey                  *Guid     `json:"Склад_Key,omitempty"`
	StavkaNDS                      *String   `json:"СтавкаНДС,omitempty"`
	Sum                            *Double   `json:"Сумма,omitempty"`
	SummaNDS                       *Double   `json:"СуммаНДС,omitempty"`
	SchetFakturaKey                *Guid     `json:"СчетФактура_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double   `json:"Цена,omitempty"`
	RetailCost                     *Double   `json:"ЦенаВРознице,omitempty"`
}
type DocumentAvansovyiOtchetOplataPostavshchikamRowType struct {
	Key                          Guid      `json:"Ref_Key,omitempty"`
	LineNumber                   Int64     `json:"LineNumber,omitempty"`
	VidDokVkhodiashchii          *String   `json:"ВидДокВходящий,omitempty"`
	DataVkhodiashchegoDokumenta  *DateTime `json:"ДатаВходящегоДокумента,omitempty"`
	DogovorKontragentaKey        *Guid     `json:"ДоговорКонтрагента_Key,omitempty"`
	KontragentKey                *Guid     `json:"Контрагент_Key,omitempty"`
	KratnostVzaimoraschetov      *Int64    `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov          *Double   `json:"КурсВзаиморасчетов,omitempty"`
	NomerVkhodiashchegoDokumenta *String   `json:"НомерВходящегоДокумента,omitempty"`
	Sdelka                       *String   `json:"Сделка,omitempty"`
	Soderzhanie                  *String   `json:"Содержание,omitempty"`
	Sum                          *Double   `json:"Сумма,omitempty"`
	SummaVzaimoraschetov         *Double   `json:"СуммаВзаиморасчетов,omitempty"`
	SdelkaType                   *String   `json:"Сделка_Type,omitempty"`
}
type DocumentAvansovyiOtchetProcheeRowType struct {
	Key                          Guid      `json:"Ref_Key,omitempty"`
	LineNumber                   Int64     `json:"LineNumber,omitempty"`
	VidDokVkhodiashchii          *String   `json:"ВидДокВходящий,omitempty"`
	DataVkhodiashchegoDokumenta  *DateTime `json:"ДатаВходящегоДокумента,omitempty"`
	DataSF                       *DateTime `json:"ДатаСФ,omitempty"`
	Zakaz                        *String   `json:"Заказ,omitempty"`
	ItemKey                      *Guid     `json:"Номенклатура_Key,omitempty"`
	NomenklaturnaiaGruppaKey     *Guid     `json:"НоменклатурнаяГруппа_Key,omitempty"`
	NomerVkhodiashchegoDokumenta *String   `json:"НомерВходящегоДокумента,omitempty"`
	NomerSF                      *String   `json:"НомерСФ,omitempty"`
	PodrazdelenieKey             *Guid     `json:"Подразделение_Key,omitempty"`
	SupplierKey                  *Guid     `json:"Поставщик_Key,omitempty"`
	PredieiavlenSF               *Boolean  `json:"ПредъявленСФ,omitempty"`
	ProektKey                    *Guid     `json:"Проект_Key,omitempty"`
	Soderzhanie                  *String   `json:"Содержание,omitempty"`
	StavkaNDS                    *String   `json:"СтавкаНДС,omitempty"`
	StatiaZatratKey              *Guid     `json:"СтатьяЗатрат_Key,omitempty"`
	Sum                          *Double   `json:"Сумма,omitempty"`
	SummaNDS                     *Double   `json:"СуммаНДС,omitempty"`
	SchetFakturaKey              *Guid     `json:"СчетФактура_Key,omitempty"`
	ZakazType                    *String   `json:"Заказ_Type,omitempty"`
}
type CatalogDopolnitelnyePechatnyeFormyPrinadlezhnostRowType struct {
	Key                   Guid    `json:"Ref_Key,omitempty"`
	LineNumber            Int64   `json:"LineNumber,omitempty"`
	PredstavlenieObieekta *String `json:"ПредставлениеОбъекта,omitempty"`
	SsylkaObieekta        *String `json:"СсылкаОбъекта,omitempty"`
	SsylkaObieektaType    *String `json:"СсылкаОбъекта_Type,omitempty"`
}
type CatalogObrabotkiObsluzhivaniiaTOModeliRowType struct {
	Key        Guid    `json:"Ref_Key,omitempty"`
	LineNumber Int64   `json:"LineNumber,omitempty"`
	Model      *String `json:"Модель,omitempty"`
}
type CatalogNastroikaIntervalovTablichnaiaChastRowType struct {
	Key              Guid    `json:"Ref_Key,omitempty"`
	LineNumber       Int64   `json:"LineNumber,omitempty"`
	KonetsIntervala  *Int64  `json:"КонецИнтервала,omitempty"`
	NachaloIntervala *Int64  `json:"НачалоИнтервала,omitempty"`
	Podpis           *String `json:"Подпись,omitempty"`
}
type CatalogProfiliGruppDostupaRoliRowType struct {
	Key        Guid  `json:"Ref_Key,omitempty"`
	LineNumber Int64 `json:"LineNumber,omitempty"`
	RolKey     *Guid `json:"Роль_Key,omitempty"`
}
type CatalogProfiliGruppDostupaVidyDostupaRowType struct {
	Key               Guid     `json:"Ref_Key,omitempty"`
	LineNumber        Int64    `json:"LineNumber,omitempty"`
	VidDostupa        *String  `json:"ВидДоступа,omitempty"`
	Predustanovlennyi *Boolean `json:"Предустановленный,omitempty"`
	VseRazresheny     *Boolean `json:"ВсеРазрешены,omitempty"`
	VidDostupaType    *String  `json:"ВидДоступа_Type,omitempty"`
}
type CatalogProfiliGruppDostupaZnacheniiaDostupaRowType struct {
	Key                  Guid    `json:"Ref_Key,omitempty"`
	LineNumber           Int64   `json:"LineNumber,omitempty"`
	VidDostupa           *String `json:"ВидДоступа,omitempty"`
	ZnachenieDostupa     *String `json:"ЗначениеДоступа,omitempty"`
	VidDostupaType       *String `json:"ВидДоступа_Type,omitempty"`
	ZnachenieDostupaType *String `json:"ЗначениеДоступа_Type,omitempty"`
}
type CatalogProfiliGruppDostupaDostupPoPodsistemamRowType struct {
	Key            Guid     `json:"Ref_Key,omitempty"`
	LineNumber     Int64    `json:"LineNumber,omitempty"`
	ImiaPodsistemy *String  `json:"ИмяПодсистемы,omitempty"`
	ImiaObieekta   *String  `json:"ИмяОбъекта,omitempty"`
	Prosmotr       *Boolean `json:"Просмотр,omitempty"`
	Redaktirovanie *Boolean `json:"Редактирование,omitempty"`
	Pechat         *Boolean `json:"Печать,omitempty"`
	PokazVersii    *Boolean `json:"ПоказВерсии,omitempty"`
}
type CatalogNastroikiDliaKureraSostavNaimenovaniiaRowType struct {
	Key                  Guid    `json:"Ref_Key,omitempty"`
	LineNumber           Int64   `json:"LineNumber,omitempty"`
	SimvolyDo            *String `json:"СимволыДо,omitempty"`
	SimvolyPosle         *String `json:"СимволыПосле,omitempty"`
	ElementNaimenovaniia *String `json:"ЭлементНаименования,omitempty"`
}
type DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	NomerPlatezha                    *Int16  `json:"НомерПлатежа,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentaRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Znachenie      *String `json:"Значение,omitempty"`
	Predstavlenie  *String `json:"Представление,omitempty"`
	Rekvizit       *String `json:"Реквизит,omitempty"`
	TipKontragenta *String `json:"ТипКонтрагента,omitempty"`
}
type DocumentVozvratTovarovOtPokupateliaTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	DokumentOprikhodovaniia        *String `json:"ДокументОприходования,omitempty"`
	DokumentPartii                 *String `json:"ДокументПартии,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	Pasport                        *String `json:"Паспорт,omitempty"`
	PercentManualDiscount          *Double `json:"ПроцентРучнойСкидки,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	Sebestoimost                   *Double `json:"Себестоимость,omitempty"`
	SebestoimostBezNDS             *Double `json:"СебестоимостьБезНДС,omitempty"`
	SebestoimostUpr                *Double `json:"СебестоимостьУпр,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	SumManualDiscount              *Double `json:"СуммаРучнойСкидки,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	StatusPartii                   *String `json:"СтатусПартии,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	OrderKey                       *Guid   `json:"КлючПродажи_Key,omitempty"`
	DokumentOprikhodovaniiaType    *String `json:"ДокументОприходования_Type,omitempty"`
	DokumentPartiiType             *String `json:"ДокументПартии_Type,omitempty"`
}
type DocumentVozvratTovarovOtPokupateliaUslugiRowType struct {
	Key                      Guid    `json:"Ref_Key,omitempty"`
	LineNumber               Int64   `json:"LineNumber,omitempty"`
	Quantity                 *Int64  `json:"Количество,omitempty"`
	ItemKey                  *Guid   `json:"Номенклатура_Key,omitempty"`
	ProtsentOtSummuDokumenta *Double `json:"ПроцентОтСуммуДокумента,omitempty"`
	ProtsentOtSummyDokumenta *Double `json:"ПроцентОтСуммыДокумента,omitempty"`
	ProtsentSkidkiNatsenki   *Double `json:"ПроцентСкидкиНаценки,omitempty"`
	Soderzhanie              *String `json:"Содержание,omitempty"`
	StavkaNDS                *String `json:"СтавкаНДС,omitempty"`
	Sum                      *Double `json:"Сумма,omitempty"`
	SummaNDS                 *Double `json:"СуммаНДС,omitempty"`
	Cost                     *Double `json:"Цена,omitempty"`
}
type DocumentZakazPostavshchikuTovaryRowType struct {
	Key             Guid    `json:"Ref_Key,omitempty"`
	LineNumber      Int64   `json:"LineNumber,omitempty"`
	Weight          *Double `json:"Вес,omitempty"`
	ZakazKlientaKey *Guid   `json:"ЗаказКлиента_Key,omitempty"`
	Quantity        *Int64  `json:"Количество,omitempty"`
	ItemKey         *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey         *Guid   `json:"Размер_Key,omitempty"`
	StavkaNDS       *String `json:"СтавкаНДС,omitempty"`
	Sum             *Double `json:"Сумма,omitempty"`
	SummaNDS        *Double `json:"СуммаНДС,omitempty"`
	Cost            *Double `json:"Цена,omitempty"`
}
type CatalogSkidkiNatsenkiUsloviiaPredostavleniiaRowType struct {
	Key                       Guid  `json:"Ref_Key,omitempty"`
	LineNumber                Int64 `json:"LineNumber,omitempty"`
	UsloviePredostavleniiaKey *Guid `json:"УсловиеПредоставления_Key,omitempty"`
}
type CatalogSkidkiNatsenkiTsenovyeGruppyRowType struct {
	Key                     Guid    `json:"Ref_Key,omitempty"`
	LineNumber              Int64   `json:"LineNumber,omitempty"`
	TsenovaiaGruppaKey      *Guid   `json:"ЦеноваяГруппа_Key,omitempty"`
	ZnachenieSkidkiNatsenki *Double `json:"ЗначениеСкидкиНаценки,omitempty"`
}
type CatalogSkidkiNatsenkiNaborPodarkovRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Quantity                       *Double `json:"Количество,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
}
type CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanieRowType struct {
	Key                     Guid    `json:"Ref_Key,omitempty"`
	LineNumber              Int64   `json:"LineNumber,omitempty"`
	VidOplatyKey            *Guid   `json:"ВидОплаты_Key,omitempty"`
	ProtsentTorgovoiUstupki *Double `json:"ПроцентТорговойУступки,omitempty"`
}
type DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsenRowType struct {
	Key        Guid  `json:"Ref_Key,omitempty"`
	LineNumber Int64 `json:"LineNumber,omitempty"`
	TipTsenKey *Guid `json:"ТипЦен_Key,omitempty"`
}
type DocumentUstanovkaTsenNomenklaturyKontragentovTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	ValiutaKey                     *Guid   `json:"Валюта_Key,omitempty"`
	IndeksStrokiTablitsyTsen       *Int64  `json:"ИндексСтрокиТаблицыЦен,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	TipTsenKey                     *Guid   `json:"ТипЦен_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
}
type DocumentProtsentPoterPoDavaltsamProtsentyRowType struct {
	Key        Guid    `json:"Ref_Key,omitempty"`
	LineNumber Int64   `json:"LineNumber,omitempty"`
	ItemKey    *Guid   `json:"Номенклатура_Key,omitempty"`
	Protsent   *Double `json:"Процент,omitempty"`
}
type DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentaRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Znachenie      *String `json:"Значение,omitempty"`
	Predstavlenie  *String `json:"Представление,omitempty"`
	Rekvizit       *String `json:"Реквизит,omitempty"`
	TipKontragenta *String `json:"ТипКонтрагента,omitempty"`
}
type CatalogVidyKodirovokiIzdeliiElementyKodirovkiRowType struct {
	Key        Guid     `json:"Ref_Key,omitempty"`
	LineNumber Int64    `json:"LineNumber,omitempty"`
	Poriadok   *Int64   `json:"Порядок,omitempty"`
	Poteri     *Boolean `json:"Потери,omitempty"`
	Prais      *Boolean `json:"Прайс,omitempty"`
	ElementKey *Guid    `json:"Элемент_Key,omitempty"`
}
type DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeliRowType struct {
	Key               Guid      `json:"Ref_Key,omitempty"`
	LineNumber        Int64     `json:"LineNumber,omitempty"`
	VremiaNachala     *DateTime `json:"ВремяНачала,omitempty"`
	VremiaOkonchaniia *DateTime `json:"ВремяОкончания,omitempty"`
	Vybran            *Boolean  `json:"Выбран,omitempty"`
	DenNedeli         *String   `json:"ДеньНедели,omitempty"`
}
type DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartyRowType struct {
	Key           Guid  `json:"Ref_Key,omitempty"`
	LineNumber    Int64 `json:"LineNumber,omitempty"`
	MemberCardKey *Guid `json:"ДисконтнаяКарта_Key,omitempty"`
}
type DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkiRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Poluchatel     *String `json:"Получатель,omitempty"`
	PoluchatelType *String `json:"Получатель_Type,omitempty"`
}
type DocumentUstanovkaSkidokNomenklaturyTovaryRowType struct {
	Key                        Guid    `json:"Ref_Key,omitempty"`
	LineNumber                 Int64   `json:"LineNumber,omitempty"`
	ItemKey                    *Guid   `json:"Номенклатура_Key,omitempty"`
	OgranichenieSkidkiNatsenki *Double `json:"ОграничениеСкидкиНаценки,omitempty"`
	ProtsentSkidkiNatsenki     *Double `json:"ПроцентСкидкиНаценки,omitempty"`
}
type CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviiaRowType struct {
	Key               Guid      `json:"Ref_Key,omitempty"`
	LineNumber        Int64     `json:"LineNumber,omitempty"`
	DenNedeli         *String   `json:"ДеньНедели,omitempty"`
	VremiaNachala     *DateTime `json:"ВремяНачала,omitempty"`
	VremiaOkonchaniia *DateTime `json:"ВремяОкончания,omitempty"`
}
type CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateliRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Poluchatel     *String `json:"Получатель,omitempty"`
	PoluchatelType *String `json:"Получатель_Type,omitempty"`
}
type CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkiRowType struct {
	Key             Guid    `json:"Ref_Key,omitempty"`
	LineNumber      Int64   `json:"LineNumber,omitempty"`
	ItemKey         *Guid   `json:"Номенклатура_Key,omitempty"`
	Quantity        *Double `json:"Количество,omitempty"`
	TypeKey         *Guid   `json:"ТипИзделия_Key,omitempty"`
	SupplierKey     *Guid   `json:"Поставщик_Key,omitempty"`
	ProizvoditelKey *Guid   `json:"Производитель_Key,omitempty"`
}
type DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhaRowType struct {
	Key                              Guid    `json:"Ref_Key,omitempty"`
	LineNumber                       Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey            *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	DokumentPlanirovaniiaPlatezhaKey *Guid   `json:"ДокументПланированияПлатежа_Key,omitempty"`
	KratnostVzaimoraschetov          *Int64  `json:"КратностьВзаиморасчетов,omitempty"`
	KursVzaimoraschetov              *Double `json:"КурсВзаиморасчетов,omitempty"`
	KursVzaimoraschetovPlan          *Double `json:"КурсВзаиморасчетовПлан,omitempty"`
	ProektKey                        *Guid   `json:"Проект_Key,omitempty"`
	Sdelka                           *String `json:"Сделка,omitempty"`
	StavkaNDS                        *String `json:"СтавкаНДС,omitempty"`
	TypeOfMovingMoneyKey             *Guid   `json:"СтатьяДвиженияДенежныхСредств_Key,omitempty"`
	SummaVzaimoraschetov             *Double `json:"СуммаВзаиморасчетов,omitempty"`
	SummaNDS                         *Double `json:"СуммаНДС,omitempty"`
	Sum                              *Double `json:"СуммаПлатежа,omitempty"`
	SummaPlatezhaPlan                *Double `json:"СуммаПлатежаПлан,omitempty"`
	SdelkaType                       *String `json:"Сделка_Type,omitempty"`
}
type DocumentRaskhodnyiKassovyiOrderOplataRowType struct {
	Key        Guid    `json:"Ref_Key,omitempty"`
	LineNumber Int64   `json:"LineNumber,omitempty"`
	TipOplaty  *String `json:"ТипОплаты,omitempty"`
	Sum        *Double `json:"Сумма,omitempty"`
}
type DocumentRaskhodnyiKassovyiOrderTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	SummaSkidki                    *Double `json:"СуммаСкидки,omitempty"`
	VidTovaraKKT                   *String `json:"ВидТовараККТ,omitempty"`
	TipOplatyTovaraKKT             *String `json:"ТипОплатыТовараККТ,omitempty"`
	SummaOsn                       *Double `json:"СуммаОсн,omitempty"`
	Komitent                       *String `json:"Комитент,omitempty"`
	TelefonKomitenta               *String `json:"ТелефонКомитента,omitempty"`
	INNkomitenta                   *String `json:"ИННкомитента,omitempty"`
	SummaOpl                       *Double `json:"СуммаОпл,omitempty"`
}
type DocumentSchetNaOplatuPostavshchikaTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
}
type DocumentSchetNaOplatuPostavshchikaUslugiRowType struct {
	Key         Guid    `json:"Ref_Key,omitempty"`
	LineNumber  Int64   `json:"LineNumber,omitempty"`
	Quantity    *Int64  `json:"Количество,omitempty"`
	ItemKey     *Guid   `json:"Номенклатура_Key,omitempty"`
	Soderzhanie *String `json:"Содержание,omitempty"`
	StavkaNDS   *String `json:"СтавкаНДС,omitempty"`
	Sum         *Double `json:"Сумма,omitempty"`
	SummaNDS    *Double `json:"СуммаНДС,omitempty"`
	Cost        *Double `json:"Цена,omitempty"`
}
type DocumentReestrSpetssviazKlientyRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	KontragentKey *Guid   `json:"Контрагент_Key,omitempty"`
	Adres         *String `json:"Адрес,omitempty"`
	Telefon       *String `json:"Телефон,omitempty"`
	Weight        *Double `json:"Вес,omitempty"`
	Sum           *Double `json:"Сумма,omitempty"`
	Paket         *String `json:"Пакет,omitempty"`
	SummaPropisiu *String `json:"СуммаПрописью,omitempty"`
	GabarityKey   *Guid   `json:"Габариты_Key,omitempty"`
}
type DocumentVvodNachalnykhOstatkovVzaimoraschetyRowType struct {
	Key                   Guid    `json:"Ref_Key,omitempty"`
	LineNumber            Int64   `json:"LineNumber,omitempty"`
	DogovorKontragentaKey *Guid   `json:"ДоговорКонтрагента_Key,omitempty"`
	KontragentKey         *Guid   `json:"Контрагент_Key,omitempty"`
	Sum                   *Double `json:"Сумма,omitempty"`
}
type DocumentVvodNachalnykhOstatkovTovaryRowType struct {
	Key                            Guid     `json:"Ref_Key,omitempty"`
	LineNumber                     Int64    `json:"LineNumber,omitempty"`
	Weight                         *Double  `json:"Вес,omitempty"`
	DokumentOprikhodovaniiaKey     *Guid    `json:"ДокументОприходования_Key,omitempty"`
	DokumentProdazhiKey            *Guid    `json:"ДокументПродажи_Key,omitempty"`
	KachestvoKey                   *Guid    `json:"Качество_Key,omitempty"`
	Quantity                       *Double  `json:"Количество,omitempty"`
	Comment                        *String  `json:"Комментарий,omitempty"`
	NDSVkliuchenVStoimost          *Boolean `json:"НДСВключенВСтоимость,omitempty"`
	ItemKey                        *Guid    `json:"Номенклатура_Key,omitempty"`
	NomerGTDKey                    *Guid    `json:"НомерГТД_Key,omitempty"`
	Pasport                        *String  `json:"Паспорт,omitempty"`
	ProtsentRoznichnoiNatsenki     *Double  `json:"ПроцентРозничнойНаценки,omitempty"`
	ProtsentSkidkiNatsenki         *Double  `json:"ПроцентСкидкиНаценки,omitempty"`
	SizeKey                        *Guid    `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid    `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                      *String  `json:"СтавкаНДС,omitempty"`
	StavkaNDSProdazhi              *String  `json:"СтавкаНДСПродажи,omitempty"`
	StatusPartii                   *String  `json:"СтатусПартии,omitempty"`
	StranaProiskhozhdeniiaKey      *Guid    `json:"СтранаПроисхождения_Key,omitempty"`
	Sum                            *Double  `json:"Сумма,omitempty"`
	SummaNDS                       *Double  `json:"СуммаНДС,omitempty"`
	SummaNDSProdazhi               *Double  `json:"СуммаНДСПродажи,omitempty"`
	SummaProdazhi                  *Double  `json:"СуммаПродажи,omitempty"`
	SummaProdazhiBezSkidok         *Double  `json:"СуммаПродажиБезСкидок,omitempty"`
	KharakteristikaNomenklaturyKey *Guid    `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double  `json:"Цена,omitempty"`
	RetailCost                     *Double  `json:"ЦенаВРознице,omitempty"`
	TsenaVRoznitseGr               *Double  `json:"ЦенаВРозницеГр,omitempty"`
	TsenaProdazhi                  *Double  `json:"ЦенаПродажи,omitempty"`
	StatusRaskhoda                 *String  `json:"СтатусРасхода,omitempty"`
}
type DocumentOprikhodovanieTovarovTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	Pasport                        *String `json:"Паспорт,omitempty"`
	ProtsentRoznichnoiNatsenki     *Double `json:"ПроцентРозничнойНаценки,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaRegl                      *Double `json:"СуммаРегл,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	RetailCost                     *Double `json:"ЦенаВРознице,omitempty"`
	TsenaVRoznitseGr               *Double `json:"ЦенаВРозницеГр,omitempty"`
}
type DocumentOprikhodovanieTovarovSertifikatyRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	SertifikatKey *Guid   `json:"Сертификат_Key,omitempty"`
	Sum           *Double `json:"Сумма,omitempty"`
}
type DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	SummaStaraia                   *Double `json:"СуммаСтарая,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	TsenaZaGramm                   *Double `json:"ЦенаЗаГрамм,omitempty"`
}
type DocumentElektronnoePismoKomuTChRowType struct {
	Key                    Guid    `json:"Ref_Key,omitempty"`
	LineNumber             Int64   `json:"LineNumber,omitempty"`
	AdresElektronnoiPochty *String `json:"АдресЭлектроннойПочты,omitempty"`
	Predstavlenie          *String `json:"Представление,omitempty"`
}
type DocumentElektronnoePismoKopiiTChRowType struct {
	Key                    Guid    `json:"Ref_Key,omitempty"`
	LineNumber             Int64   `json:"LineNumber,omitempty"`
	AdresElektronnoiPochty *String `json:"АдресЭлектроннойПочты,omitempty"`
	Predstavlenie          *String `json:"Представление,omitempty"`
}
type DocumentElektronnoePismoSkrytyeKopiiTChRowType struct {
	Key                    Guid    `json:"Ref_Key,omitempty"`
	LineNumber             Int64   `json:"LineNumber,omitempty"`
	AdresElektronnoiPochty *String `json:"АдресЭлектроннойПочты,omitempty"`
	Predstavlenie          *String `json:"Представление,omitempty"`
}
type CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentaRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	Kliuch        *String `json:"Ключ,omitempty"`
	Znachenie     *String `json:"Значение,omitempty"`
	ZnachenieType *String `json:"Значение_Type,omitempty"`
}
type CatalogPravilaTsenoobrazovaniiaTsenovyeGruppyRowType struct {
	Key                Guid  `json:"Ref_Key,omitempty"`
	LineNumber         Int64 `json:"LineNumber,omitempty"`
	TsenovaiaGruppaKey *Guid `json:"ЦеноваяГруппа_Key,omitempty"`
	VidTsenKey         *Guid `json:"ВидЦен_Key,omitempty"`
}
type DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	EdinitsaIzmereniiaKey          *Guid   `json:"ЕдиницаИзмерения_Key,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	Koef                           *Double `json:"Коэф,omitempty"`
	Koeffitsient                   *Double `json:"Коэффициент,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	Pasport                        *String `json:"Паспорт,omitempty"`
	ProektKey                      *Guid   `json:"Проект_Key,omitempty"`
	ProtsentRoznichnoiNatsenki     *Double `json:"ПроцентРозничнойНаценки,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	RetailCost                     *Double `json:"ЦенаВРознице,omitempty"`
	TsenaVRoznitseGr               *Double `json:"ЦенаВРозницеГр,omitempty"`
}
type DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugiRowType struct {
	Key                      Guid    `json:"Ref_Key,omitempty"`
	LineNumber               Int64   `json:"LineNumber,omitempty"`
	Quantity                 *Int64  `json:"Количество,omitempty"`
	ItemKey                  *Guid   `json:"Номенклатура_Key,omitempty"`
	NomenklaturnaiaGruppaKey *Guid   `json:"НоменклатурнаяГруппа_Key,omitempty"`
	PodrazdelenieKey         *Guid   `json:"Подразделение_Key,omitempty"`
	ProektKey                *Guid   `json:"Проект_Key,omitempty"`
	Soderzhanie              *String `json:"Содержание,omitempty"`
	StavkaNDS                *String `json:"СтавкаНДС,omitempty"`
	StatiaZatratKey          *Guid   `json:"СтатьяЗатрат_Key,omitempty"`
	Sum                      *Double `json:"Сумма,omitempty"`
	SummaNDS                 *Double `json:"СуммаНДС,omitempty"`
	Cost                     *Double `json:"Цена,omitempty"`
}
type CatalogGruppyDostupaPolzovateliRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	Polzovatel     *String `json:"Пользователь,omitempty"`
	PolzovatelType *String `json:"Пользователь_Type,omitempty"`
}
type CatalogGruppyDostupaVidyDostupaRowType struct {
	Key            Guid     `json:"Ref_Key,omitempty"`
	LineNumber     Int64    `json:"LineNumber,omitempty"`
	VidDostupa     *String  `json:"ВидДоступа,omitempty"`
	VseRazresheny  *Boolean `json:"ВсеРазрешены,omitempty"`
	VidDostupaType *String  `json:"ВидДоступа_Type,omitempty"`
}
type CatalogGruppyDostupaZnacheniiaDostupaRowType struct {
	Key                  Guid    `json:"Ref_Key,omitempty"`
	LineNumber           Int64   `json:"LineNumber,omitempty"`
	VidDostupa           *String `json:"ВидДоступа,omitempty"`
	ZnachenieDostupa     *String `json:"ЗначениеДоступа,omitempty"`
	VidDostupaType       *String `json:"ВидДоступа_Type,omitempty"`
	ZnachenieDostupaType *String `json:"ЗначениеДоступа_Type,omitempty"`
}
type CatalogGruppyDostupaDostupPoPodsistemamRowType struct {
	Key            Guid     `json:"Ref_Key,omitempty"`
	LineNumber     Int64    `json:"LineNumber,omitempty"`
	ImiaPodsistemy *String  `json:"ИмяПодсистемы,omitempty"`
	ImiaObieekta   *String  `json:"ИмяОбъекта,omitempty"`
	Prosmotr       *Boolean `json:"Просмотр,omitempty"`
	Redaktirovanie *Boolean `json:"Редактирование,omitempty"`
	Pechat         *Boolean `json:"Печать,omitempty"`
	PokazVersii    *Boolean `json:"ПоказВерсии,omitempty"`
}
type DocumentReestrSchetovReestrRowType struct {
	Key            Guid      `json:"Ref_Key,omitempty"`
	LineNumber     Int64     `json:"LineNumber,omitempty"`
	VidTransporta  *String   `json:"ВидТранспорта,omitempty"`
	DataOtgruzki   *DateTime `json:"ДатаОтгрузки,omitempty"`
	NomerDokumenta *String   `json:"НомерДокумента,omitempty"`
	Sum            *Double   `json:"Сумма,omitempty"`
}
type DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	VesUchet                       *Double `json:"ВесУчет,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	KolichestvoUchet               *Int64  `json:"КоличествоУчет,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
}
type CatalogNastroikiRMKPoriadokPrimeneniiaSkidokRowType struct {
	Key           Guid    `json:"Ref_Key,omitempty"`
	LineNumber    Int64   `json:"LineNumber,omitempty"`
	UslovieSkidki *String `json:"УсловиеСкидки,omitempty"`
}
type CatalogNastroikiRMKSostavNaimenovaniiaRowType struct {
	Key                  Guid    `json:"Ref_Key,omitempty"`
	LineNumber           Int64   `json:"LineNumber,omitempty"`
	SimvolyDo            *String `json:"СимволыДо,omitempty"`
	SimvolyPosle         *String `json:"СимволыПосле,omitempty"`
	ElementNaimenovaniia *String `json:"ЭлементНаименования,omitempty"`
}
type CatalogKharakteristikiNomenklaturySpetsifikatsiiaRowType struct {
	Key              Guid    `json:"Ref_Key,omitempty"`
	LineNumber       Int64   `json:"LineNumber,omitempty"`
	Weight           *Double `json:"Вес,omitempty"`
	GruppaDefektaKey *Guid   `json:"ГруппаДефекта_Key,omitempty"`
	GruppaTsvetaKey  *Guid   `json:"ГруппаЦвета_Key,omitempty"`
	KamenKey         *Guid   `json:"Камень_Key,omitempty"`
	Quantity         *Double `json:"Количество,omitempty"`
	ItemKey          *Guid   `json:"Номенклатура_Key,omitempty"`
	Razmer1          *Double `json:"Размер1,omitempty"`
	Razmer2          *Double `json:"Размер2,omitempty"`
	Razmer3          *Double `json:"Размер3,omitempty"`
	RassevKey        *Guid   `json:"Рассев_Key,omitempty"`
	FormaOgrankiKey  *Guid   `json:"ФормаОгранки_Key,omitempty"`
	TsvetKamniaKey   *Guid   `json:"ЦветКамня_Key,omitempty"`
}
type DocumentOtborTovarovTovaryRowType struct {
	Key                                        Guid    `json:"Ref_Key,omitempty"`
	LineNumber                                 Int64   `json:"LineNumber,omitempty"`
	Weight                                     *Double `json:"Вес,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidki     *String `json:"ЗначениеУсловияАвтоматическойСкидки,omitempty"`
	Quantity                                   *Int64  `json:"Количество,omitempty"`
	ItemKey                                    *Guid   `json:"Номенклатура_Key,omitempty"`
	PercentAutoDiscount                        *Double `json:"ПроцентАвтоматическойСкидки,omitempty"`
	PercentManualDiscount                      *Double `json:"ПроцентРучнойСкидки,omitempty"`
	SizeKey                                    *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                                *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                                  *String `json:"СтавкаНДС,omitempty"`
	Sum                                        *Double `json:"Сумма,omitempty"`
	SummaNDS                                   *Double `json:"СуммаНДС,omitempty"`
	UslovieAvtomaticheskoiSkidki               *String `json:"УсловиеАвтоматическойСкидки,omitempty"`
	KharakteristikaNomenklaturyKey             *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                                       *Double `json:"Цена,omitempty"`
	SumManualDiscount                          *Double `json:"СуммаРучнойСкидки,omitempty"`
	SumAutoDiscount                            *Double `json:"СуммаАвтоматическойСкидки,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType *String `json:"ЗначениеУсловияАвтоматическойСкидки_Type,omitempty"`
}
type DocumentOtborTovarovTovaryKOtboruRowType struct {
	Key         Guid    `json:"Ref_Key,omitempty"`
	LineNumber  Int64   `json:"LineNumber,omitempty"`
	Weight      *Double `json:"Вес,omitempty"`
	Quantity    *Int64  `json:"Количество,omitempty"`
	ItemKey     *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey     *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
}
type CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizyRowType struct {
	Key                       Guid  `json:"Ref_Key,omitempty"`
	LineNumber                Int64 `json:"LineNumber,omitempty"`
	RelizIuvelirnogoSalonaKey *Guid `json:"РелизЮвелирногоСалона_Key,omitempty"`
}
type DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvaRowType struct {
	Key                   Guid    `json:"Ref_Key,omitempty"`
	LineNumber            Int64   `json:"LineNumber,omitempty"`
	VidOtchetaPoPlatezham *String `json:"ВидОтчетаПоПлатежам,omitempty"`
	StavkaNDS             *String `json:"СтавкаНДС,omitempty"`
	Sum                   *Double `json:"Сумма,omitempty"`
	SummaNDS              *Double `json:"СуммаНДС,omitempty"`
}
type DocumentOtchetKomissioneraOProdazhakhTovaryRowType struct {
	Key                            Guid      `json:"Ref_Key,omitempty"`
	LineNumber                     Int64     `json:"LineNumber,omitempty"`
	Weight                         *Double   `json:"Вес,omitempty"`
	Quantity                       *Int64    `json:"Количество,omitempty"`
	ItemKey                        *Guid     `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid     `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid     `json:"СерияНоменклатуры_Key,omitempty"`
	StavkaNDS                      *String   `json:"СтавкаНДС,omitempty"`
	Sum                            *Double   `json:"Сумма,omitempty"`
	SummaVoznagrazhdeniia          *Double   `json:"СуммаВознаграждения,omitempty"`
	SummaNDS                       *Double   `json:"СуммаНДС,omitempty"`
	SummaNDSVoznagrazhdeniia       *Double   `json:"СуммаНДСВознаграждения,omitempty"`
	SummaNDSPeredachi              *Double   `json:"СуммаНДСПередачи,omitempty"`
	SummaPeredachi                 *Double   `json:"СуммаПередачи,omitempty"`
	KharakteristikaNomenklaturyKey *Guid     `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double   `json:"Цена,omitempty"`
	TsenaPeredachi                 *Double   `json:"ЦенаПередачи,omitempty"`
	DataSchetaFakturyKomissionera  *DateTime `json:"ДатаСчетаФактурыКомиссионера,omitempty"`
	PokupatelKey                   *Guid     `json:"Покупатель_Key,omitempty"`
	NomerSchetaFakturyKomissionera *String   `json:"НомерСчетаФактурыКомиссионера,omitempty"`
}
type CatalogFiltryDliaElektronnykhPisemDeistviiaFiltraRowType struct {
	Key            Guid    `json:"Ref_Key,omitempty"`
	LineNumber     Int64   `json:"LineNumber,omitempty"`
	GruppaPisemKey *Guid   `json:"ГруппаПисем_Key,omitempty"`
	DeistvieFiltra *String `json:"ДействиеФильтра,omitempty"`
}
type CatalogFiltryDliaElektronnykhPisemUsloviiaFiltraRowType struct {
	Key                Guid     `json:"Ref_Key,omitempty"`
	LineNumber         Int64    `json:"LineNumber,omitempty"`
	ZnachenieUsloviia  *String  `json:"ЗначениеУсловия,omitempty"`
	OtritsanieUsloviia *Boolean `json:"ОтрицаниеУсловия,omitempty"`
	Uslovie            *String  `json:"Условие,omitempty"`
}
type DocumentPreiskurantTsenNaTsvKamniTablitsaRowType struct {
	Key             Guid    `json:"Ref_Key,omitempty"`
	LineNumber      Int64   `json:"LineNumber,omitempty"`
	Razmer1Do       *Double `json:"Размер1До,omitempty"`
	Razmer1Ot       *Double `json:"Размер1От,omitempty"`
	RassevKey       *Guid   `json:"Рассев_Key,omitempty"`
	FormaOgrankiKey *Guid   `json:"ФормаОгранки_Key,omitempty"`
	TsvetKamniaKey  *Guid   `json:"ЦветКамня_Key,omitempty"`
	Cost            *Int64  `json:"Цена,omitempty"`
}
type DocumentTelemarketingUchastnikiRowType struct {
	Key                     Guid     `json:"Ref_Key,omitempty"`
	LineNumber              Int64    `json:"LineNumber,omitempty"`
	KontragentKey           *Guid    `json:"Контрагент_Key,omitempty"`
	NaimenovaniePolnoe      *String  `json:"НаименованиеПолное,omitempty"`
	KontaktnoeLitsoKey      *Guid    `json:"КонтактноеЛицо_Key,omitempty"`
	Telefon                 *String  `json:"Телефон,omitempty"`
	RezultatObrabotkiZvonka *String  `json:"РезультатОбработкиЗвонка,omitempty"`
	EstInteres              *Boolean `json:"ЕстьИнтерес,omitempty"`
	SobytieKey              *Guid    `json:"Событие_Key,omitempty"`
	OprosKey                *Guid    `json:"Опрос_Key,omitempty"`
	Opisanie                *String  `json:"Описание,omitempty"`
}
type DocumentRassylkaAnketVlozheniiaRowType struct {
	Key                 Guid    `json:"Ref_Key,omitempty"`
	LineNumber          Int64   `json:"LineNumber,omitempty"`
	VlozhenieBase64Data *Binary `json:"Вложение_Base64Data,omitempty"`
	ImiaFaila           *String `json:"ИмяФайла,omitempty"`
	VlozhenieType       *String `json:"Вложение_Type,omitempty"`
	Vlozhenie           *Stream `json:"Вложение,omitempty"`
}
type DocumentRassylkaAnketPoluchateliRowType struct {
	Key         Guid    `json:"Ref_Key,omitempty"`
	LineNumber  Int64   `json:"LineNumber,omitempty"`
	Obieekt     *String `json:"Объект,omitempty"`
	Poluchatel  *String `json:"Получатель,omitempty"`
	ObieektType *String `json:"Объект_Type,omitempty"`
}
type CatalogSkhemyRealizatsiiEtapySkhemyRowType struct {
	Key                       Guid     `json:"Ref_Key,omitempty"`
	LineNumber                Int64    `json:"LineNumber,omitempty"`
	DogovorPokupkiKey         *Guid    `json:"ДоговорПокупки_Key,omitempty"`
	DogovorProdazhiKey        *Guid    `json:"ДоговорПродажи_Key,omitempty"`
	ZnachenieNatsenki         *String  `json:"ЗначениеНаценки,omitempty"`
	KontragentPokupatelKey    *Guid    `json:"КонтрагентПокупатель_Key,omitempty"`
	KontragentProdavetsKey    *Guid    `json:"КонтрагентПродавец_Key,omitempty"`
	OkrugliatVBolshuiuStoronu *Boolean `json:"ОкруглятьВБольшуюСторону,omitempty"`
	OrganizatsiiaPokupatelKey *Guid    `json:"ОрганизацияПокупатель_Key,omitempty"`
	OrganizatsiiaProdavetsKey *Guid    `json:"ОрганизацияПродавец_Key,omitempty"`
	PoriadokOkrugleniia       *String  `json:"ПорядокОкругления,omitempty"`
	SposobNatsenki            *Int16   `json:"СпособНаценки,omitempty"`
	TipNatsenki               *Int16   `json:"ТипНаценки,omitempty"`
	ZnachenieNatsenkiType     *String  `json:"ЗначениеНаценки_Type,omitempty"`
}
type DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	XYZKlassifikatsiia             *String `json:"XYZКлассификация,omitempty"`
	XYZKlassifikatsiiaStaraia      *String `json:"XYZКлассификацияСтарая,omitempty"`
	ZnachenieParametra             *Double `json:"ЗначениеПараметра,omitempty"`
	IndeksSortirovki               *Int16  `json:"ИндексСортировки,omitempty"`
	KolichestvoDokumentov          *Int64  `json:"КоличествоДокументов,omitempty"`
	KontragentKey                  *Guid   `json:"Контрагент_Key,omitempty"`
	KoeffitsientVariatsii          *Double `json:"КоэффициентВариации,omitempty"`
	MenedzherKontragentaKey        *Guid   `json:"МенеджерКонтрагента_Key,omitempty"`
	StadiiaVzaimootnoshenii        *String `json:"СтадияВзаимоотношений,omitempty"`
	StadiiaVzaimootnosheniiStaraia *String `json:"СтадияВзаимоотношенийСтарая,omitempty"`
}
type DocumentZakazKlientaTovaryRowType struct {
	Key                                        Guid    `json:"Ref_Key,omitempty"`
	LineNumber                                 Int64   `json:"LineNumber,omitempty"`
	Weight                                     *Double `json:"Вес,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidki     *String `json:"ЗначениеУсловияАвтоматическойСкидки,omitempty"`
	Quantity                                   *Int64  `json:"Количество,omitempty"`
	ItemKey                                    *Guid   `json:"Номенклатура_Key,omitempty"`
	PercentAutoDiscount                        *Double `json:"ПроцентАвтоматическойСкидки,omitempty"`
	PercentManualDiscount                      *Double `json:"ПроцентРучнойСкидки,omitempty"`
	SizeKey                                    *Guid   `json:"Размер_Key,omitempty"`
	Razmestit                                  *Int64  `json:"Разместить,omitempty"`
	Rezervirovat                               *Int64  `json:"Резервировать,omitempty"`
	StavkaNDS                                  *String `json:"СтавкаНДС,omitempty"`
	Sum                                        *Double `json:"Сумма,omitempty"`
	SummaNDS                                   *Double `json:"СуммаНДС,omitempty"`
	UslovieAvtomaticheskoiSkidki               *String `json:"УсловиеАвтоматическойСкидки,omitempty"`
	Cost                                       *Double `json:"Цена,omitempty"`
	InstanceKey                                *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	SumAutoDiscount                            *Double `json:"СуммаАвтоматическойСкидки,omitempty"`
	SumManualDiscount                          *Double `json:"СуммаРучнойСкидки,omitempty"`
	KharakteristikaNomenklaturyKey             *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType *String `json:"ЗначениеУсловияАвтоматическойСкидки_Type,omitempty"`
}
type DocumentPostuplenieProduktsiiIzProizvodstvaTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	VesPoter                       *Double `json:"ВесПотерь,omitempty"`
	ZakazKlientaKey                *Guid   `json:"ЗаказКлиента_Key,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	Pasport                        *String `json:"Паспорт,omitempty"`
	ProtsentPoter                  *Double `json:"ПроцентПотерь,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	StoimostVstavok                *Double `json:"СтоимостьВставок,omitempty"`
	StoimostMetalla                *Double `json:"СтоимостьМеталла,omitempty"`
	StoimostRabot                  *Double `json:"СтоимостьРабот,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	CostOfWork                     *Double `json:"ЦенаРабот,omitempty"`
	SummaNDSVstavok                *Double `json:"СуммаНДСВставок,omitempty"`
	SummaNDSMetalla                *Double `json:"СуммаНДСМеталла,omitempty"`
	SummaNDSRabot                  *Double `json:"СуммаНДСРабот,omitempty"`
	VesVstavok                     *Double `json:"ВесВставок,omitempty"`
}
type DocumentPostuplenieProduktsiiIzProizvodstvaMaterialyRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	Nomenklatura                   *String `json:"Номенклатура,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	ItemType                       *String `json:"Номенклатура_Type,omitempty"`
}
type DocumentPostuplenieTovarovUslugTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	EdinitsaIzmereniiaKey          *Guid   `json:"ЕдиницаИзмерения_Key,omitempty"`
	ZakazKlientaKey                *Guid   `json:"ЗаказКлиента_Key,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	Koef                           *Double `json:"Коэф,omitempty"`
	Koeffitsient                   *Double `json:"Коэффициент,omitempty"`
	NaborKey                       *Guid   `json:"Набор_Key,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	NomerGTDKey                    *Guid   `json:"НомерГТД_Key,omitempty"`
	NomerNabora                    *Int64  `json:"НомерНабора,omitempty"`
	Pasport                        *String `json:"Паспорт,omitempty"`
	ProektKey                      *Guid   `json:"Проект_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	DepartmentKey                  *Guid   `json:"Склад_Key,omitempty"`
	StavkaNDS                      *String `json:"СтавкаНДС,omitempty"`
	StranaProiskhozhdeniiaKey      *Guid   `json:"СтранаПроисхождения_Key,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaNDS                       *Double `json:"СуммаНДС,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
}
type DocumentPostuplenieTovarovUslugUslugiRowType struct {
	Key                      Guid    `json:"Ref_Key,omitempty"`
	LineNumber               Int64   `json:"LineNumber,omitempty"`
	Quantity                 *Int64  `json:"Количество,omitempty"`
	ItemKey                  *Guid   `json:"Номенклатура_Key,omitempty"`
	NomenklaturnaiaGruppaKey *Guid   `json:"НоменклатурнаяГруппа_Key,omitempty"`
	PodrazdelenieKey         *Guid   `json:"Подразделение_Key,omitempty"`
	ProektKey                *Guid   `json:"Проект_Key,omitempty"`
	Soderzhanie              *String `json:"Содержание,omitempty"`
	StavkaNDS                *String `json:"СтавкаНДС,omitempty"`
	StatiaZatratKey          *Guid   `json:"СтатьяЗатрат_Key,omitempty"`
	Sum                      *Double `json:"Сумма,omitempty"`
	SummaNDS                 *Double `json:"СуммаНДС,omitempty"`
	Cost                     *Double `json:"Цена,omitempty"`
}
type DocumentPlanProdazhPoSalonamSalonyRowType struct {
	Key                        Guid      `json:"Ref_Key,omitempty"`
	LineNumber                 Int64     `json:"LineNumber,omitempty"`
	SalonKey                   *Guid     `json:"Салон_Key,omitempty"`
	SummaPlana                 *Double   `json:"СуммаПлана,omitempty"`
	Primechanie                *String   `json:"Примечание,omitempty"`
	IndeksStrokiIzTablitsyDnei *Int64    `json:"ИндексСтрокиИзТаблицыДней,omitempty"`
	SummaPlanaDnevnaia         *Double   `json:"СуммаПланаДневная,omitempty"`
	DenPoGrafiku               *DateTime `json:"ДеньПоГрафику,omitempty"`
	SummaPlanaFakt             *Double   `json:"СуммаПланаФакт,omitempty"`
	PlanEst                    *Boolean  `json:"ПланЕсть,omitempty"`
	KU                         *Double   `json:"КУ,omitempty"`
	FormulaDliaRasschetaKey    *Guid     `json:"ФормулаДляРассчета_Key,omitempty"`
	RasshifrovkaFormuly        *String   `json:"РасшифровкаФормулы,omitempty"`
}
type DocumentPlanProdazhPoSalonamDniPoGrafikuRowType struct {
	Key          Guid   `json:"Ref_Key,omitempty"`
	LineNumber   Int64  `json:"LineNumber,omitempty"`
	DenPoGrafiku *Int16 `json:"ДеньПоГрафику,omitempty"`
}
type DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvaRowType struct {
	Key                   Guid    `json:"Ref_Key,omitempty"`
	LineNumber            Int64   `json:"LineNumber,omitempty"`
	VidOtchetaPoPlatezham *String `json:"ВидОтчетаПоПлатежам,omitempty"`
	StavkaNDS             *String `json:"СтавкаНДС,omitempty"`
	Sum                   *Double `json:"Сумма,omitempty"`
	SummaNDS              *Double `json:"СуммаНДС,omitempty"`
}
type DocumentStornirovanieOtchetaKomitentuOProdazhakhTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	DokumentPostupleniia           *String `json:"ДокументПоступления,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	OtchetKomitentuKey             *Guid   `json:"ОтчетКомитенту_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	Sum                            *Double `json:"Сумма,omitempty"`
	SummaVoznagrazhdeniia          *Double `json:"СуммаВознаграждения,omitempty"`
	SummaNDSVoznagrazhdeniia       *Double `json:"СуммаНДСВознаграждения,omitempty"`
	SummaPostupleniia              *Double `json:"СуммаПоступления,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	Cost                           *Double `json:"Цена,omitempty"`
	TsenaPostupleniia              *Double `json:"ЦенаПоступления,omitempty"`
	DokumentPostupleniiaType       *String `json:"ДокументПоступления_Type,omitempty"`
}
type DocumentPeredachaVRemontIzdeliiaPriniatyeVRemontRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	KliuchNomenklaturyKey          *Guid   `json:"КлючНоменклатуры_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	DokumentOprikhodovaniiaKey     *Guid   `json:"ДокументОприходования_Key,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
}
type DocumentPeredachaVRemontTovaryRowType struct {
	Key                            Guid    `json:"Ref_Key,omitempty"`
	LineNumber                     Int64   `json:"LineNumber,omitempty"`
	Weight                         *Double `json:"Вес,omitempty"`
	KachestvoKey                   *Guid   `json:"Качество_Key,omitempty"`
	Quantity                       *Int64  `json:"Количество,omitempty"`
	ItemKey                        *Guid   `json:"Номенклатура_Key,omitempty"`
	SizeKey                        *Guid   `json:"Размер_Key,omitempty"`
	InstanceKey                    *Guid   `json:"СерияНоменклатуры_Key,omitempty"`
	KharakteristikaNomenklaturyKey *Guid   `json:"ХарактеристикаНоменклатуры_Key,omitempty"`
}
type CatalogGruppySkladovSkladyRowType struct {
	Key           Guid  `json:"Ref_Key,omitempty"`
	LineNumber    Int64 `json:"LineNumber,omitempty"`
	DepartmentKey *Guid `json:"Склад_Key,omitempty"`
}
