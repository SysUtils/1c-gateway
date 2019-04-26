package main

import (
	gqlserver "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"gitlab.com/zullpro/core/1cclientgenerator.git/odata"
	"log"
	"net/http"
)

var Schema = `
schema { query: Query }

scalar String
scalar Int
scalar Int16
scalar Int64
scalar Double
scalar Float
scalar DateTime
scalar Binary
scalar Stream
scalar Boolean
scalar Guid

input NumberQualifiersInput {
	AllowedSign: String!
	Digits: Int16!
	FractionDigits: Int16!
}
type NumberQualifiers {
	AllowedSign: String!
	Digits: Int16!
	FractionDigits: Int16!
}
input TypeDescriptionInput {
	Types: [String!]!
	NumberQualifiers: NumberQualifiersInput!
}
type TypeDescription {
	Types: [String!]!
	NumberQualifiers: NumberQualifiers!
}
input PointInTimeInput {
	Ref: String!
	Date: DateTime!
}
type PointInTime {
	Ref: String!
	Date: DateTime!
}
input AccumulationRegisterPartiiTovarovVProizvodstveRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ProizvodstvennyiUchastokKey: Guid
	Nomenklatura: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	OperationCode: String
	SpisaniePartii: Boolean
	NomerKorStroki: Int64
	RecorderType: String!
	ItemType: String
	DokumentOprikhodovaniiaType: String
}
type AccumulationRegisterPartiiTovarovVProizvodstveRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ProizvodstvennyiUchastokKey: Guid
	Nomenklatura: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	OperationCode: String
	SpisaniePartii: Boolean
	NomerKorStroki: Int64
	RecorderType: String!
	ItemType: String
	DokumentOprikhodovaniiaType: String
}
input AccumulationRegisterPartiiTovarovVProizvodstveBalanceInput {
	ProizvodstvennyiUchastokKey: Guid
	Nomenklatura: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KolichestvoBalance: Int64
	VesBalance: Double
	StoimostBalance: Double
	StoimostUprBalance: Double
	StoimostBezNDSBalance: Double
	ItemType: String
	DokumentOprikhodovaniiaType: String
}
type AccumulationRegisterPartiiTovarovVProizvodstveBalance {
	ProizvodstvennyiUchastokKey: Guid
	Nomenklatura: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KolichestvoBalance: Int64
	VesBalance: Double
	StoimostBalance: Double
	StoimostUprBalance: Double
	StoimostBezNDSBalance: Double
	ItemType: String
	DokumentOprikhodovaniiaType: String
}
input AccumulationRegisterPartiiTovarovVProizvodstveTurnoverInput {
	ProizvodstvennyiUchastokKey: Guid
	Nomenklatura: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	StoimostTurnover: Double
	StoimostReceipt: Double
	StoimostExpense: Double
	StoimostUprTurnover: Double
	StoimostUprReceipt: Double
	StoimostUprExpense: Double
	StoimostBezNDSTurnover: Double
	StoimostBezNDSReceipt: Double
	StoimostBezNDSExpense: Double
	ItemType: String
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
type AccumulationRegisterPartiiTovarovVProizvodstveTurnover {
	ProizvodstvennyiUchastokKey: Guid
	Nomenklatura: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	StoimostTurnover: Double
	StoimostReceipt: Double
	StoimostExpense: Double
	StoimostUprTurnover: Double
	StoimostUprReceipt: Double
	StoimostUprExpense: Double
	StoimostBezNDSTurnover: Double
	StoimostBezNDSReceipt: Double
	StoimostBezNDSExpense: Double
	ItemType: String
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
input AccumulationRegisterPartiiTovarovVProizvodstveBalanceAndTurnoverInput {
	ProizvodstvennyiUchastokKey: Guid
	Nomenklatura: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	StoimostOpeningBalance: Double
	StoimostTurnover: Double
	StoimostReceipt: Double
	StoimostExpense: Double
	StoimostClosingBalance: Double
	StoimostUprOpeningBalance: Double
	StoimostUprTurnover: Double
	StoimostUprReceipt: Double
	StoimostUprExpense: Double
	StoimostUprClosingBalance: Double
	StoimostBezNDSOpeningBalance: Double
	StoimostBezNDSTurnover: Double
	StoimostBezNDSReceipt: Double
	StoimostBezNDSExpense: Double
	StoimostBezNDSClosingBalance: Double
	ItemType: String
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
type AccumulationRegisterPartiiTovarovVProizvodstveBalanceAndTurnover {
	ProizvodstvennyiUchastokKey: Guid
	Nomenklatura: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	StoimostOpeningBalance: Double
	StoimostTurnover: Double
	StoimostReceipt: Double
	StoimostExpense: Double
	StoimostClosingBalance: Double
	StoimostUprOpeningBalance: Double
	StoimostUprTurnover: Double
	StoimostUprReceipt: Double
	StoimostUprExpense: Double
	StoimostUprClosingBalance: Double
	StoimostBezNDSOpeningBalance: Double
	StoimostBezNDSTurnover: Double
	StoimostBezNDSReceipt: Double
	StoimostBezNDSExpense: Double
	StoimostBezNDSClosingBalance: Double
	ItemType: String
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
input AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	FizLitsoKey: Guid
	ValiutaKey: Guid
	RaschetnyiDokument: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	RaschetnyiDokumentType: String
}
type AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	FizLitsoKey: Guid
	ValiutaKey: Guid
	RaschetnyiDokument: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	RaschetnyiDokumentType: String
}
input AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalanceInput {
	FizLitsoKey: Guid
	ValiutaKey: Guid
	RaschetnyiDokument: String
	SummaVzaimoraschetovBalance: Double
	SummaUprBalance: Double
	RaschetnyiDokumentType: String
}
type AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalance {
	FizLitsoKey: Guid
	ValiutaKey: Guid
	RaschetnyiDokument: String
	SummaVzaimoraschetovBalance: Double
	SummaUprBalance: Double
	RaschetnyiDokumentType: String
}
input AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiTurnoverInput {
	FizLitsoKey: Guid
	ValiutaKey: Guid
	RaschetnyiDokument: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	RaschetnyiDokumentType: String
	RecorderType: String
}
type AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiTurnover {
	FizLitsoKey: Guid
	ValiutaKey: Guid
	RaschetnyiDokument: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	RaschetnyiDokumentType: String
	RecorderType: String
}
input AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalanceAndTurnoverInput {
	FizLitsoKey: Guid
	ValiutaKey: Guid
	RaschetnyiDokument: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	RaschetnyiDokumentType: String
	RecorderType: String
}
type AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalanceAndTurnover {
	FizLitsoKey: Guid
	ValiutaKey: Guid
	RaschetnyiDokument: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	RaschetnyiDokumentType: String
	RecorderType: String
}
input AccumulationRegisterVnutrennieZakazyRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Zakazchik: String
	VnutrenniiZakazKey: Guid
	StatusPartii: String
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	ZakazchikType: String
}
type AccumulationRegisterVnutrennieZakazyRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Zakazchik: String
	VnutrenniiZakazKey: Guid
	StatusPartii: String
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	ZakazchikType: String
}
input AccumulationRegisterVnutrennieZakazyBalanceInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Zakazchik: String
	VnutrenniiZakazKey: Guid
	StatusPartii: String
	KolichestvoBalance: Int64
	VesBalance: Double
	ZakazchikType: String
}
type AccumulationRegisterVnutrennieZakazyBalance {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Zakazchik: String
	VnutrenniiZakazKey: Guid
	StatusPartii: String
	KolichestvoBalance: Int64
	VesBalance: Double
	ZakazchikType: String
}
input AccumulationRegisterVnutrennieZakazyTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Zakazchik: String
	VnutrenniiZakazKey: Guid
	StatusPartii: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	ZakazchikType: String
	RecorderType: String
}
type AccumulationRegisterVnutrennieZakazyTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Zakazchik: String
	VnutrenniiZakazKey: Guid
	StatusPartii: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	ZakazchikType: String
	RecorderType: String
}
input AccumulationRegisterVnutrennieZakazyBalanceAndTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Zakazchik: String
	VnutrenniiZakazKey: Guid
	StatusPartii: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	ZakazchikType: String
	RecorderType: String
}
type AccumulationRegisterVnutrennieZakazyBalanceAndTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Zakazchik: String
	VnutrenniiZakazKey: Guid
	StatusPartii: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	ZakazchikType: String
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaKomitentaRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKomitentaRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
input AccumulationRegisterDenezhnyeSredstvaKomitentaBalanceInput {
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetovBalance: Double
	SummaUprBalance: Double
	SdelkaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKomitentaBalance {
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetovBalance: Double
	SummaUprBalance: Double
	SdelkaType: String
}
input AccumulationRegisterDenezhnyeSredstvaKomitentaTurnoverInput {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SdelkaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaKomitentaTurnover {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SdelkaType: String
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaKomitentaBalanceAndTurnoverInput {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	SdelkaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaKomitentaBalanceAndTurnover {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	SdelkaType: String
	RecorderType: String
}
input AccumulationRegisterZakazyKlientovRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	SizeKey: Guid
	ZakazKlientaKey: Guid
	InstanceKey: Guid
	Zakazano: Int64
	VRezerve: Int64
	KRazmeshcheniiu: Int64
	ZakazanoVes: Double
	ZakazanoSumma: Double
	RecorderType: String!
}
type AccumulationRegisterZakazyKlientovRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	SizeKey: Guid
	ZakazKlientaKey: Guid
	InstanceKey: Guid
	Zakazano: Int64
	VRezerve: Int64
	KRazmeshcheniiu: Int64
	ZakazanoVes: Double
	ZakazanoSumma: Double
	RecorderType: String!
}
input AccumulationRegisterZakazyKlientovBalanceInput {
	ItemKey: Guid
	SizeKey: Guid
	ZakazKlientaKey: Guid
	InstanceKey: Guid
	ZakazanoBalance: Int64
	VRezerveBalance: Int64
	KRazmeshcheniiuBalance: Int64
	ZakazanoVesBalance: Double
	ZakazanoSummaBalance: Double
}
type AccumulationRegisterZakazyKlientovBalance {
	ItemKey: Guid
	SizeKey: Guid
	ZakazKlientaKey: Guid
	InstanceKey: Guid
	ZakazanoBalance: Int64
	VRezerveBalance: Int64
	KRazmeshcheniiuBalance: Int64
	ZakazanoVesBalance: Double
	ZakazanoSummaBalance: Double
}
input AccumulationRegisterZakazyKlientovTurnoverInput {
	ItemKey: Guid
	SizeKey: Guid
	ZakazKlientaKey: Guid
	InstanceKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	ZakazanoTurnover: Int64
	ZakazanoReceipt: Int64
	ZakazanoExpense: Int64
	VRezerveTurnover: Int64
	VRezerveReceipt: Int64
	VRezerveExpense: Int64
	KRazmeshcheniiuTurnover: Int64
	KRazmeshcheniiuReceipt: Int64
	KRazmeshcheniiuExpense: Int64
	ZakazanoVesTurnover: Double
	ZakazanoVesReceipt: Double
	ZakazanoVesExpense: Double
	ZakazanoSummaTurnover: Double
	ZakazanoSummaReceipt: Double
	ZakazanoSummaExpense: Double
	RecorderType: String
}
type AccumulationRegisterZakazyKlientovTurnover {
	ItemKey: Guid
	SizeKey: Guid
	ZakazKlientaKey: Guid
	InstanceKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	ZakazanoTurnover: Int64
	ZakazanoReceipt: Int64
	ZakazanoExpense: Int64
	VRezerveTurnover: Int64
	VRezerveReceipt: Int64
	VRezerveExpense: Int64
	KRazmeshcheniiuTurnover: Int64
	KRazmeshcheniiuReceipt: Int64
	KRazmeshcheniiuExpense: Int64
	ZakazanoVesTurnover: Double
	ZakazanoVesReceipt: Double
	ZakazanoVesExpense: Double
	ZakazanoSummaTurnover: Double
	ZakazanoSummaReceipt: Double
	ZakazanoSummaExpense: Double
	RecorderType: String
}
input AccumulationRegisterZakazyKlientovBalanceAndTurnoverInput {
	ItemKey: Guid
	SizeKey: Guid
	ZakazKlientaKey: Guid
	InstanceKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	ZakazanoOpeningBalance: Int64
	ZakazanoTurnover: Int64
	ZakazanoReceipt: Int64
	ZakazanoExpense: Int64
	ZakazanoClosingBalance: Int64
	VRezerveOpeningBalance: Int64
	VRezerveTurnover: Int64
	VRezerveReceipt: Int64
	VRezerveExpense: Int64
	VRezerveClosingBalance: Int64
	KRazmeshcheniiuOpeningBalance: Int64
	KRazmeshcheniiuTurnover: Int64
	KRazmeshcheniiuReceipt: Int64
	KRazmeshcheniiuExpense: Int64
	KRazmeshcheniiuClosingBalance: Int64
	ZakazanoVesOpeningBalance: Double
	ZakazanoVesTurnover: Double
	ZakazanoVesReceipt: Double
	ZakazanoVesExpense: Double
	ZakazanoVesClosingBalance: Double
	ZakazanoSummaOpeningBalance: Double
	ZakazanoSummaTurnover: Double
	ZakazanoSummaReceipt: Double
	ZakazanoSummaExpense: Double
	ZakazanoSummaClosingBalance: Double
	RecorderType: String
}
type AccumulationRegisterZakazyKlientovBalanceAndTurnover {
	ItemKey: Guid
	SizeKey: Guid
	ZakazKlientaKey: Guid
	InstanceKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	ZakazanoOpeningBalance: Int64
	ZakazanoTurnover: Int64
	ZakazanoReceipt: Int64
	ZakazanoExpense: Int64
	ZakazanoClosingBalance: Int64
	VRezerveOpeningBalance: Int64
	VRezerveTurnover: Int64
	VRezerveReceipt: Int64
	VRezerveExpense: Int64
	VRezerveClosingBalance: Int64
	KRazmeshcheniiuOpeningBalance: Int64
	KRazmeshcheniiuTurnover: Int64
	KRazmeshcheniiuReceipt: Int64
	KRazmeshcheniiuExpense: Int64
	KRazmeshcheniiuClosingBalance: Int64
	ZakazanoVesOpeningBalance: Double
	ZakazanoVesTurnover: Double
	ZakazanoVesReceipt: Double
	ZakazanoVesExpense: Double
	ZakazanoVesClosingBalance: Double
	ZakazanoSummaOpeningBalance: Double
	ZakazanoSummaTurnover: Double
	ZakazanoSummaReceipt: Double
	ZakazanoSummaExpense: Double
	ZakazanoSummaClosingBalance: Double
	RecorderType: String
}
input AccumulationRegisterSummyPoFinmonitoringuRoznitsaRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	SummaPokupok: Double
	SummaSkupki: Double
	RecorderType: String!
}
type AccumulationRegisterSummyPoFinmonitoringuRoznitsaRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	SummaPokupok: Double
	SummaSkupki: Double
	RecorderType: String!
}
input AccumulationRegisterSummyPoFinmonitoringuRoznitsaBalanceInput {
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	SummaPokupokBalance: Double
	SummaSkupkiBalance: Double
}
type AccumulationRegisterSummyPoFinmonitoringuRoznitsaBalance {
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	SummaPokupokBalance: Double
	SummaSkupkiBalance: Double
}
input AccumulationRegisterSummyPoFinmonitoringuRoznitsaTurnoverInput {
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaPokupokTurnover: Double
	SummaPokupokReceipt: Double
	SummaPokupokExpense: Double
	SummaSkupkiTurnover: Double
	SummaSkupkiReceipt: Double
	SummaSkupkiExpense: Double
	RecorderType: String
}
type AccumulationRegisterSummyPoFinmonitoringuRoznitsaTurnover {
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaPokupokTurnover: Double
	SummaPokupokReceipt: Double
	SummaPokupokExpense: Double
	SummaSkupkiTurnover: Double
	SummaSkupkiReceipt: Double
	SummaSkupkiExpense: Double
	RecorderType: String
}
input AccumulationRegisterSummyPoFinmonitoringuRoznitsaBalanceAndTurnoverInput {
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaPokupokOpeningBalance: Double
	SummaPokupokTurnover: Double
	SummaPokupokReceipt: Double
	SummaPokupokExpense: Double
	SummaPokupokClosingBalance: Double
	SummaSkupkiOpeningBalance: Double
	SummaSkupkiTurnover: Double
	SummaSkupkiReceipt: Double
	SummaSkupkiExpense: Double
	SummaSkupkiClosingBalance: Double
	RecorderType: String
}
type AccumulationRegisterSummyPoFinmonitoringuRoznitsaBalanceAndTurnover {
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaPokupokOpeningBalance: Double
	SummaPokupokTurnover: Double
	SummaPokupokReceipt: Double
	SummaPokupokExpense: Double
	SummaPokupokClosingBalance: Double
	SummaSkupkiOpeningBalance: Double
	SummaSkupkiTurnover: Double
	SummaSkupkiReceipt: Double
	SummaSkupkiExpense: Double
	SummaSkupkiClosingBalance: Double
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentPolucheniia: String
	TypeOfMovingMoneyKey: Guid
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentPolucheniiaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentPolucheniia: String
	TypeOfMovingMoneyKey: Guid
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentPolucheniiaType: String
}
input AccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalanceInput {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentPolucheniia: String
	TypeOfMovingMoneyKey: Guid
	SummaBalance: Double
	SummaUprBalance: Double
	BankovskiiSchetKassaType: String
	DokumentPolucheniiaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalance {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentPolucheniia: String
	TypeOfMovingMoneyKey: Guid
	SummaBalance: Double
	SummaUprBalance: Double
	BankovskiiSchetKassaType: String
	DokumentPolucheniiaType: String
}
input AccumulationRegisterDenezhnyeSredstvaKPolucheniiuTurnoverInput {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentPolucheniia: String
	TypeOfMovingMoneyKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	BankovskiiSchetKassaType: String
	DokumentPolucheniiaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaKPolucheniiuTurnover {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentPolucheniia: String
	TypeOfMovingMoneyKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	BankovskiiSchetKassaType: String
	DokumentPolucheniiaType: String
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalanceAndTurnoverInput {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentPolucheniia: String
	TypeOfMovingMoneyKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	BankovskiiSchetKassaType: String
	DokumentPolucheniiaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalanceAndTurnover {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentPolucheniia: String
	TypeOfMovingMoneyKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	BankovskiiSchetKassaType: String
	DokumentPolucheniiaType: String
	RecorderType: String
}
input AccumulationRegisterProdazhiPoDiskontnymKartamRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	MemberCardKey: Guid
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DokumentProdazhi: String
	Sum: Double
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	DokumentProdazhiType: String
}
type AccumulationRegisterProdazhiPoDiskontnymKartamRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	MemberCardKey: Guid
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DokumentProdazhi: String
	Sum: Double
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	DokumentProdazhiType: String
}
input AccumulationRegisterProdazhiPoDiskontnymKartamTurnoverInput {
	MemberCardKey: Guid
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DokumentProdazhi: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	KolichestvoTurnover: Int64
	VesTurnover: Double
	DokumentProdazhiType: String
	RecorderType: String
}
type AccumulationRegisterProdazhiPoDiskontnymKartamTurnover {
	MemberCardKey: Guid
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DokumentProdazhi: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	KolichestvoTurnover: Int64
	VesTurnover: Double
	DokumentProdazhiType: String
	RecorderType: String
}
input AccumulationRegisterTovaryPoluchennyeRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Quantity: Int64
	Weight: Double
	SummaVzaimoraschetov: Double
	RecorderType: String!
	SdelkaType: String
}
type AccumulationRegisterTovaryPoluchennyeRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Quantity: Int64
	Weight: Double
	SummaVzaimoraschetov: Double
	RecorderType: String!
	SdelkaType: String
}
input AccumulationRegisterTovaryPoluchennyeBalanceInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	KolichestvoBalance: Int64
	VesBalance: Double
	SummaVzaimoraschetovBalance: Double
	SdelkaType: String
}
type AccumulationRegisterTovaryPoluchennyeBalance {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	KolichestvoBalance: Int64
	VesBalance: Double
	SummaVzaimoraschetovBalance: Double
	SdelkaType: String
}
input AccumulationRegisterTovaryPoluchennyeTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SdelkaType: String
	RecorderType: String
}
type AccumulationRegisterTovaryPoluchennyeTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SdelkaType: String
	RecorderType: String
}
input AccumulationRegisterTovaryPoluchennyeBalanceAndTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SdelkaType: String
	RecorderType: String
}
type AccumulationRegisterTovaryPoluchennyeBalanceAndTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SdelkaType: String
	RecorderType: String
}
input AccumulationRegisterSvobodnyeOstatkiRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	VNalichii: Int64
	VRezerve: Int64
	RecorderType: String!
}
type AccumulationRegisterSvobodnyeOstatkiRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	VNalichii: Int64
	VRezerve: Int64
	RecorderType: String!
}
input AccumulationRegisterSvobodnyeOstatkiBalanceInput {
	ItemKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	VNalichiiBalance: Int64
	VRezerveBalance: Int64
}
type AccumulationRegisterSvobodnyeOstatkiBalance {
	ItemKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	VNalichiiBalance: Int64
	VRezerveBalance: Int64
}
input AccumulationRegisterSvobodnyeOstatkiTurnoverInput {
	ItemKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VNalichiiTurnover: Int64
	VNalichiiReceipt: Int64
	VNalichiiExpense: Int64
	VRezerveTurnover: Int64
	VRezerveReceipt: Int64
	VRezerveExpense: Int64
	RecorderType: String
}
type AccumulationRegisterSvobodnyeOstatkiTurnover {
	ItemKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VNalichiiTurnover: Int64
	VNalichiiReceipt: Int64
	VNalichiiExpense: Int64
	VRezerveTurnover: Int64
	VRezerveReceipt: Int64
	VRezerveExpense: Int64
	RecorderType: String
}
input AccumulationRegisterSvobodnyeOstatkiBalanceAndTurnoverInput {
	ItemKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VNalichiiOpeningBalance: Int64
	VNalichiiTurnover: Int64
	VNalichiiReceipt: Int64
	VNalichiiExpense: Int64
	VNalichiiClosingBalance: Int64
	VRezerveOpeningBalance: Int64
	VRezerveTurnover: Int64
	VRezerveReceipt: Int64
	VRezerveExpense: Int64
	VRezerveClosingBalance: Int64
	RecorderType: String
}
type AccumulationRegisterSvobodnyeOstatkiBalanceAndTurnover {
	ItemKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VNalichiiOpeningBalance: Int64
	VNalichiiTurnover: Int64
	VNalichiiReceipt: Int64
	VNalichiiExpense: Int64
	VNalichiiClosingBalance: Int64
	VRezerveOpeningBalance: Int64
	VRezerveTurnover: Int64
	VRezerveReceipt: Int64
	VRezerveExpense: Int64
	VRezerveClosingBalance: Int64
	RecorderType: String
}
input AccumulationRegisterSummyVRassrochkuRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorRassrochkiKey: Guid
	SostavStrokiChekaKey: Guid
	Sum: Double
	SummaNDS: Double
	RecorderType: String!
}
type AccumulationRegisterSummyVRassrochkuRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorRassrochkiKey: Guid
	SostavStrokiChekaKey: Guid
	Sum: Double
	SummaNDS: Double
	RecorderType: String!
}
input AccumulationRegisterSummyVRassrochkuBalanceInput {
	OrganizatsiiaKey: Guid
	DogovorRassrochkiKey: Guid
	SostavStrokiChekaKey: Guid
	SummaBalance: Double
	SummaNDSBalance: Double
}
type AccumulationRegisterSummyVRassrochkuBalance {
	OrganizatsiiaKey: Guid
	DogovorRassrochkiKey: Guid
	SostavStrokiChekaKey: Guid
	SummaBalance: Double
	SummaNDSBalance: Double
}
input AccumulationRegisterSummyVRassrochkuTurnoverInput {
	OrganizatsiiaKey: Guid
	DogovorRassrochkiKey: Guid
	SostavStrokiChekaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaNDSTurnover: Double
	SummaNDSReceipt: Double
	SummaNDSExpense: Double
	RecorderType: String
}
type AccumulationRegisterSummyVRassrochkuTurnover {
	OrganizatsiiaKey: Guid
	DogovorRassrochkiKey: Guid
	SostavStrokiChekaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaNDSTurnover: Double
	SummaNDSReceipt: Double
	SummaNDSExpense: Double
	RecorderType: String
}
input AccumulationRegisterSummyVRassrochkuBalanceAndTurnoverInput {
	OrganizatsiiaKey: Guid
	DogovorRassrochkiKey: Guid
	SostavStrokiChekaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	SummaNDSOpeningBalance: Double
	SummaNDSTurnover: Double
	SummaNDSReceipt: Double
	SummaNDSExpense: Double
	SummaNDSClosingBalance: Double
	RecorderType: String
}
type AccumulationRegisterSummyVRassrochkuBalanceAndTurnover {
	OrganizatsiiaKey: Guid
	DogovorRassrochkiKey: Guid
	SostavStrokiChekaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	SummaNDSOpeningBalance: Double
	SummaNDSTurnover: Double
	SummaNDSReceipt: Double
	SummaNDSExpense: Double
	SummaNDSClosingBalance: Double
	RecorderType: String
}
input AccumulationRegisterGrafikPlatezheiRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	VidDogovoraKontragenta: String
	DataDolga: DateTime
	Oplacheno: Double
	NachislenDolg: Double
	Dolg: Double
	Avans: Double
	OplachenoUpr: Int64
	NachislenDolgUpr: Double
	DolgUpr: Double
	AvansUpr: Double
	RecorderType: String!
}
type AccumulationRegisterGrafikPlatezheiRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	VidDogovoraKontragenta: String
	DataDolga: DateTime
	Oplacheno: Double
	NachislenDolg: Double
	Dolg: Double
	Avans: Double
	OplachenoUpr: Int64
	NachislenDolgUpr: Double
	DolgUpr: Double
	AvansUpr: Double
	RecorderType: String!
}
input AccumulationRegisterGrafikPlatezheiBalanceInput {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	VidDogovoraKontragenta: String
	DataDolga: DateTime
	OplachenoBalance: Double
	NachislenDolgBalance: Double
	DolgBalance: Double
	AvansBalance: Double
	OplachenoUprBalance: Int64
	NachislenDolgUprBalance: Double
	DolgUprBalance: Double
	AvansUprBalance: Double
}
type AccumulationRegisterGrafikPlatezheiBalance {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	VidDogovoraKontragenta: String
	DataDolga: DateTime
	OplachenoBalance: Double
	NachislenDolgBalance: Double
	DolgBalance: Double
	AvansBalance: Double
	OplachenoUprBalance: Int64
	NachislenDolgUprBalance: Double
	DolgUprBalance: Double
	AvansUprBalance: Double
}
input AccumulationRegisterGrafikPlatezheiTurnoverInput {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	VidDogovoraKontragenta: String
	DataDolga: DateTime
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	OplachenoTurnover: Double
	OplachenoReceipt: Double
	OplachenoExpense: Double
	NachislenDolgTurnover: Double
	NachislenDolgReceipt: Double
	NachislenDolgExpense: Double
	DolgTurnover: Double
	DolgReceipt: Double
	DolgExpense: Double
	AvansTurnover: Double
	AvansReceipt: Double
	AvansExpense: Double
	OplachenoUprTurnover: Int64
	OplachenoUprReceipt: Int64
	OplachenoUprExpense: Int64
	NachislenDolgUprTurnover: Double
	NachislenDolgUprReceipt: Double
	NachislenDolgUprExpense: Double
	DolgUprTurnover: Double
	DolgUprReceipt: Double
	DolgUprExpense: Double
	AvansUprTurnover: Double
	AvansUprReceipt: Double
	AvansUprExpense: Double
	RecorderType: String
}
type AccumulationRegisterGrafikPlatezheiTurnover {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	VidDogovoraKontragenta: String
	DataDolga: DateTime
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	OplachenoTurnover: Double
	OplachenoReceipt: Double
	OplachenoExpense: Double
	NachislenDolgTurnover: Double
	NachislenDolgReceipt: Double
	NachislenDolgExpense: Double
	DolgTurnover: Double
	DolgReceipt: Double
	DolgExpense: Double
	AvansTurnover: Double
	AvansReceipt: Double
	AvansExpense: Double
	OplachenoUprTurnover: Int64
	OplachenoUprReceipt: Int64
	OplachenoUprExpense: Int64
	NachislenDolgUprTurnover: Double
	NachislenDolgUprReceipt: Double
	NachislenDolgUprExpense: Double
	DolgUprTurnover: Double
	DolgUprReceipt: Double
	DolgUprExpense: Double
	AvansUprTurnover: Double
	AvansUprReceipt: Double
	AvansUprExpense: Double
	RecorderType: String
}
input AccumulationRegisterGrafikPlatezheiBalanceAndTurnoverInput {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	VidDogovoraKontragenta: String
	DataDolga: DateTime
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	OplachenoOpeningBalance: Double
	OplachenoTurnover: Double
	OplachenoReceipt: Double
	OplachenoExpense: Double
	OplachenoClosingBalance: Double
	NachislenDolgOpeningBalance: Double
	NachislenDolgTurnover: Double
	NachislenDolgReceipt: Double
	NachislenDolgExpense: Double
	NachislenDolgClosingBalance: Double
	DolgOpeningBalance: Double
	DolgTurnover: Double
	DolgReceipt: Double
	DolgExpense: Double
	DolgClosingBalance: Double
	AvansOpeningBalance: Double
	AvansTurnover: Double
	AvansReceipt: Double
	AvansExpense: Double
	AvansClosingBalance: Double
	OplachenoUprOpeningBalance: Int64
	OplachenoUprTurnover: Int64
	OplachenoUprReceipt: Int64
	OplachenoUprExpense: Int64
	OplachenoUprClosingBalance: Int64
	NachislenDolgUprOpeningBalance: Double
	NachislenDolgUprTurnover: Double
	NachislenDolgUprReceipt: Double
	NachislenDolgUprExpense: Double
	NachislenDolgUprClosingBalance: Double
	DolgUprOpeningBalance: Double
	DolgUprTurnover: Double
	DolgUprReceipt: Double
	DolgUprExpense: Double
	DolgUprClosingBalance: Double
	AvansUprOpeningBalance: Double
	AvansUprTurnover: Double
	AvansUprReceipt: Double
	AvansUprExpense: Double
	AvansUprClosingBalance: Double
	RecorderType: String
}
type AccumulationRegisterGrafikPlatezheiBalanceAndTurnover {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	VidDogovoraKontragenta: String
	DataDolga: DateTime
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	OplachenoOpeningBalance: Double
	OplachenoTurnover: Double
	OplachenoReceipt: Double
	OplachenoExpense: Double
	OplachenoClosingBalance: Double
	NachislenDolgOpeningBalance: Double
	NachislenDolgTurnover: Double
	NachislenDolgReceipt: Double
	NachislenDolgExpense: Double
	NachislenDolgClosingBalance: Double
	DolgOpeningBalance: Double
	DolgTurnover: Double
	DolgReceipt: Double
	DolgExpense: Double
	DolgClosingBalance: Double
	AvansOpeningBalance: Double
	AvansTurnover: Double
	AvansReceipt: Double
	AvansExpense: Double
	AvansClosingBalance: Double
	OplachenoUprOpeningBalance: Int64
	OplachenoUprTurnover: Int64
	OplachenoUprReceipt: Int64
	OplachenoUprExpense: Int64
	OplachenoUprClosingBalance: Int64
	NachislenDolgUprOpeningBalance: Double
	NachislenDolgUprTurnover: Double
	NachislenDolgUprReceipt: Double
	NachislenDolgUprExpense: Double
	NachislenDolgUprClosingBalance: Double
	DolgUprOpeningBalance: Double
	DolgUprTurnover: Double
	DolgUprReceipt: Double
	DolgUprExpense: Double
	DolgUprClosingBalance: Double
	AvansUprOpeningBalance: Double
	AvansUprTurnover: Double
	AvansUprReceipt: Double
	AvansUprExpense: Double
	AvansUprClosingBalance: Double
	RecorderType: String
}
input AccumulationRegisterRoznichnaiaVyruchkaRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	RoznichnaiaTochka: String
	Sum: Double
	PodrazdelenieKey: Guid
	RecorderType: String!
	RoznichnaiaTochkaType: String
}
type AccumulationRegisterRoznichnaiaVyruchkaRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	RoznichnaiaTochka: String
	Sum: Double
	PodrazdelenieKey: Guid
	RecorderType: String!
	RoznichnaiaTochkaType: String
}
input AccumulationRegisterRoznichnaiaVyruchkaBalanceInput {
	RoznichnaiaTochka: String
	SummaBalance: Double
	RoznichnaiaTochkaType: String
}
type AccumulationRegisterRoznichnaiaVyruchkaBalance {
	RoznichnaiaTochka: String
	SummaBalance: Double
	RoznichnaiaTochkaType: String
}
input AccumulationRegisterRoznichnaiaVyruchkaTurnoverInput {
	RoznichnaiaTochka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	RoznichnaiaTochkaType: String
	RecorderType: String
}
type AccumulationRegisterRoznichnaiaVyruchkaTurnover {
	RoznichnaiaTochka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	RoznichnaiaTochkaType: String
	RecorderType: String
}
input AccumulationRegisterRoznichnaiaVyruchkaBalanceAndTurnoverInput {
	RoznichnaiaTochka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	RoznichnaiaTochkaType: String
	RecorderType: String
}
type AccumulationRegisterRoznichnaiaVyruchkaBalanceAndTurnover {
	RoznichnaiaTochka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	RoznichnaiaTochkaType: String
	RecorderType: String
}
input AccumulationRegisterTovaryVPutiRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	DokumentOprikhodovaniiaType: String
}
type AccumulationRegisterTovaryVPutiRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	DokumentOprikhodovaniiaType: String
}
input AccumulationRegisterTovaryVPutiBalanceInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KolichestvoBalance: Int64
	VesBalance: Double
	DokumentOprikhodovaniiaType: String
}
type AccumulationRegisterTovaryVPutiBalance {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KolichestvoBalance: Int64
	VesBalance: Double
	DokumentOprikhodovaniiaType: String
}
input AccumulationRegisterTovaryVPutiTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
type AccumulationRegisterTovaryVPutiTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
input AccumulationRegisterTovaryVPutiBalanceAndTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
type AccumulationRegisterTovaryVPutiBalanceAndTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
input AccumulationRegisterPoteriMetallaVProizvodstveRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	Nomenklatura: String
	DokumentOprikhodovaniia: String
	VesPoter: Double
	RecorderType: String!
	ItemType: String
	DokumentOprikhodovaniiaType: String
}
type AccumulationRegisterPoteriMetallaVProizvodstveRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	Nomenklatura: String
	DokumentOprikhodovaniia: String
	VesPoter: Double
	RecorderType: String!
	ItemType: String
	DokumentOprikhodovaniiaType: String
}
input AccumulationRegisterPoteriMetallaVProizvodstveTurnoverInput {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	Nomenklatura: String
	DokumentOprikhodovaniia: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VesPoterTurnover: Double
	ItemType: String
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
type AccumulationRegisterPoteriMetallaVProizvodstveTurnover {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	Nomenklatura: String
	DokumentOprikhodovaniia: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VesPoterTurnover: Double
	ItemType: String
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
input AccumulationRegisterPartiiTovarovNaSkladakhRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentOprikhodovaniia: String
	StatusPartii: String
	KachestvoKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	OperationCode: String
	SpisaniePartii: Boolean
	NomerKorStroki: Int64
	RecorderType: String!
	DokumentOprikhodovaniiaType: String
}
type AccumulationRegisterPartiiTovarovNaSkladakhRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentOprikhodovaniia: String
	StatusPartii: String
	KachestvoKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	OperationCode: String
	SpisaniePartii: Boolean
	NomerKorStroki: Int64
	RecorderType: String!
	DokumentOprikhodovaniiaType: String
}
input AccumulationRegisterPartiiTovarovNaSkladakhBalanceInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentOprikhodovaniia: String
	StatusPartii: String
	KachestvoKey: Guid
	KolichestvoBalance: Int64
	VesBalance: Double
	StoimostBalance: Double
	StoimostUprBalance: Double
	StoimostBezNDSBalance: Double
	DokumentOprikhodovaniiaType: String
}
type AccumulationRegisterPartiiTovarovNaSkladakhBalance {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentOprikhodovaniia: String
	StatusPartii: String
	KachestvoKey: Guid
	KolichestvoBalance: Int64
	VesBalance: Double
	StoimostBalance: Double
	StoimostUprBalance: Double
	StoimostBezNDSBalance: Double
	DokumentOprikhodovaniiaType: String
}
input AccumulationRegisterPartiiTovarovNaSkladakhTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentOprikhodovaniia: String
	StatusPartii: String
	KachestvoKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	StoimostTurnover: Double
	StoimostReceipt: Double
	StoimostExpense: Double
	StoimostUprTurnover: Double
	StoimostUprReceipt: Double
	StoimostUprExpense: Double
	StoimostBezNDSTurnover: Double
	StoimostBezNDSReceipt: Double
	StoimostBezNDSExpense: Double
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
type AccumulationRegisterPartiiTovarovNaSkladakhTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentOprikhodovaniia: String
	StatusPartii: String
	KachestvoKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	StoimostTurnover: Double
	StoimostReceipt: Double
	StoimostExpense: Double
	StoimostUprTurnover: Double
	StoimostUprReceipt: Double
	StoimostUprExpense: Double
	StoimostBezNDSTurnover: Double
	StoimostBezNDSReceipt: Double
	StoimostBezNDSExpense: Double
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
input AccumulationRegisterPartiiTovarovNaSkladakhBalanceAndTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentOprikhodovaniia: String
	StatusPartii: String
	KachestvoKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	StoimostOpeningBalance: Double
	StoimostTurnover: Double
	StoimostReceipt: Double
	StoimostExpense: Double
	StoimostClosingBalance: Double
	StoimostUprOpeningBalance: Double
	StoimostUprTurnover: Double
	StoimostUprReceipt: Double
	StoimostUprExpense: Double
	StoimostUprClosingBalance: Double
	StoimostBezNDSOpeningBalance: Double
	StoimostBezNDSTurnover: Double
	StoimostBezNDSReceipt: Double
	StoimostBezNDSExpense: Double
	StoimostBezNDSClosingBalance: Double
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
type AccumulationRegisterPartiiTovarovNaSkladakhBalanceAndTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentOprikhodovaniia: String
	StatusPartii: String
	KachestvoKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	StoimostOpeningBalance: Double
	StoimostTurnover: Double
	StoimostReceipt: Double
	StoimostExpense: Double
	StoimostClosingBalance: Double
	StoimostUprOpeningBalance: Double
	StoimostUprTurnover: Double
	StoimostUprReceipt: Double
	StoimostUprExpense: Double
	StoimostUprClosingBalance: Double
	StoimostBezNDSOpeningBalance: Double
	StoimostBezNDSTurnover: Double
	StoimostBezNDSReceipt: Double
	StoimostBezNDSExpense: Double
	StoimostBezNDSClosingBalance: Double
	DokumentOprikhodovaniiaType: String
	RecorderType: String
}
input AccumulationRegisterSummyDokumentovDliaObmenaRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentKey: Guid
	Sum: Double
	RecorderType: String!
}
type AccumulationRegisterSummyDokumentovDliaObmenaRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentKey: Guid
	Sum: Double
	RecorderType: String!
}
input AccumulationRegisterSummyDokumentovDliaObmenaBalanceInput {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentKey: Guid
	SummaBalance: Double
}
type AccumulationRegisterSummyDokumentovDliaObmenaBalance {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentKey: Guid
	SummaBalance: Double
}
input AccumulationRegisterSummyDokumentovDliaObmenaTurnoverInput {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	RecorderType: String
}
type AccumulationRegisterSummyDokumentovDliaObmenaTurnover {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	RecorderType: String
}
input AccumulationRegisterSummyDokumentovDliaObmenaBalanceAndTurnoverInput {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	RecorderType: String
}
type AccumulationRegisterSummyDokumentovDliaObmenaBalanceAndTurnover {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	RecorderType: String
}
input AccumulationRegisterDvizheniiaDenezhnykhSredstvRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	VidDenezhnykhSredstv: String
	PrikhodRaskhod: String
	BankovskiiSchetKassa: String
	TypeOfMovingMoneyKey: Guid
	DokumentDvizheniia: String
	Kontragent: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	ProektKey: Guid
	DokumentPlanirovaniiaPlatezha: String
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentDvizheniiaType: String
	KontragentType: String
	SdelkaType: String
	DokumentPlanirovaniiaPlatezhaType: String
}
type AccumulationRegisterDvizheniiaDenezhnykhSredstvRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	VidDenezhnykhSredstv: String
	PrikhodRaskhod: String
	BankovskiiSchetKassa: String
	TypeOfMovingMoneyKey: Guid
	DokumentDvizheniia: String
	Kontragent: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	ProektKey: Guid
	DokumentPlanirovaniiaPlatezha: String
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentDvizheniiaType: String
	KontragentType: String
	SdelkaType: String
	DokumentPlanirovaniiaPlatezhaType: String
}
input AccumulationRegisterDvizheniiaDenezhnykhSredstvTurnoverInput {
	VidDenezhnykhSredstv: String
	PrikhodRaskhod: String
	BankovskiiSchetKassa: String
	TypeOfMovingMoneyKey: Guid
	DokumentDvizheniia: String
	Kontragent: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	ProektKey: Guid
	DokumentPlanirovaniiaPlatezha: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaUprTurnover: Double
	BankovskiiSchetKassaType: String
	DokumentDvizheniiaType: String
	KontragentType: String
	SdelkaType: String
	DokumentPlanirovaniiaPlatezhaType: String
	RecorderType: String
}
type AccumulationRegisterDvizheniiaDenezhnykhSredstvTurnover {
	VidDenezhnykhSredstv: String
	PrikhodRaskhod: String
	BankovskiiSchetKassa: String
	TypeOfMovingMoneyKey: Guid
	DokumentDvizheniia: String
	Kontragent: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	ProektKey: Guid
	DokumentPlanirovaniiaPlatezha: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaUprTurnover: Double
	BankovskiiSchetKassaType: String
	DokumentDvizheniiaType: String
	KontragentType: String
	SdelkaType: String
	DokumentPlanirovaniiaPlatezhaType: String
	RecorderType: String
}
input AccumulationRegisterProdazhiPoStatiamRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	StatiaKey: Guid
	SummaProdazha: Double
	SummaVozvrat: Double
	RecorderType: String!
}
type AccumulationRegisterProdazhiPoStatiamRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	StatiaKey: Guid
	SummaProdazha: Double
	SummaVozvrat: Double
	RecorderType: String!
}
input AccumulationRegisterProdazhiPoStatiamBalanceInput {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	StatiaKey: Guid
	SummaProdazhaBalance: Double
	SummaVozvratBalance: Double
}
type AccumulationRegisterProdazhiPoStatiamBalance {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	StatiaKey: Guid
	SummaProdazhaBalance: Double
	SummaVozvratBalance: Double
}
input AccumulationRegisterProdazhiPoStatiamTurnoverInput {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	StatiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaProdazhaTurnover: Double
	SummaProdazhaReceipt: Double
	SummaProdazhaExpense: Double
	SummaVozvratTurnover: Double
	SummaVozvratReceipt: Double
	SummaVozvratExpense: Double
	RecorderType: String
}
type AccumulationRegisterProdazhiPoStatiamTurnover {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	StatiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaProdazhaTurnover: Double
	SummaProdazhaReceipt: Double
	SummaProdazhaExpense: Double
	SummaVozvratTurnover: Double
	SummaVozvratReceipt: Double
	SummaVozvratExpense: Double
	RecorderType: String
}
input AccumulationRegisterProdazhiPoStatiamBalanceAndTurnoverInput {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	StatiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaProdazhaOpeningBalance: Double
	SummaProdazhaTurnover: Double
	SummaProdazhaReceipt: Double
	SummaProdazhaExpense: Double
	SummaProdazhaClosingBalance: Double
	SummaVozvratOpeningBalance: Double
	SummaVozvratTurnover: Double
	SummaVozvratReceipt: Double
	SummaVozvratExpense: Double
	SummaVozvratClosingBalance: Double
	RecorderType: String
}
type AccumulationRegisterProdazhiPoStatiamBalanceAndTurnover {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	StatiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaProdazhaOpeningBalance: Double
	SummaProdazhaTurnover: Double
	SummaProdazhaReceipt: Double
	SummaProdazhaExpense: Double
	SummaProdazhaClosingBalance: Double
	SummaVozvratOpeningBalance: Double
	SummaVozvratTurnover: Double
	SummaVozvratReceipt: Double
	SummaVozvratExpense: Double
	SummaVozvratClosingBalance: Double
	RecorderType: String
}
input InformationRegisterTsenyNomenklaturyRowTypeInput {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	TipTsenKey: Guid!
	SegmentNomenklaturyKey: Guid!
	ItemKey: Guid!
	InstanceKey: Guid!
	KharakteristikaNomenklaturyKey: Guid!
	SizeKey: Guid!
	Cost: Double
	ProtsentSkidkiNatsenki: Double
	ValiutaKey: Guid
	EdinitsaIzmereniiaKey: Guid
	RecorderType: String
}
type InformationRegisterTsenyNomenklaturyRowType {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	TipTsenKey: Guid!
	SegmentNomenklaturyKey: Guid!
	ItemKey: Guid!
	InstanceKey: Guid!
	KharakteristikaNomenklaturyKey: Guid!
	SizeKey: Guid!
	Cost: Double
	ProtsentSkidkiNatsenki: Double
	ValiutaKey: Guid
	EdinitsaIzmereniiaKey: Guid
	RecorderType: String
}
input AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	ProdazhnaiaStoimost: Double
	VsegoSkidki: Double
	SkidkiPoDiskontnymKartam: Double
	SummaOplatyKartami: Double
	SummaOplatyBankovskimKreditom: Double
	SummaVozvrata: Double
	VesVChekakh: Double
	KolichestvoChekov: Int64
	SummaProdazhiSertifikatov: Double
	SummaOplatySertifikatami: Double
	PogashenoSertifikatami: Double
	SummaOplatyBonusom: Double
	VesObmena: Double
	SummaObmena: Double
	VesSkuplennogoTovara: Double
	VydanoZaSkuplennyiTovar: Double
	KolichestvoIzdelii: Int64
	SumManualDiscount: Double
	SumAutoDiscount: Double
	SummaRassrochki: Double
	SummaPogasheniiaRassrochki: Double
	SummaOplatyNalichnymi: Double
	RecorderType: String!
}
type AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	ProdazhnaiaStoimost: Double
	VsegoSkidki: Double
	SkidkiPoDiskontnymKartam: Double
	SummaOplatyKartami: Double
	SummaOplatyBankovskimKreditom: Double
	SummaVozvrata: Double
	VesVChekakh: Double
	KolichestvoChekov: Int64
	SummaProdazhiSertifikatov: Double
	SummaOplatySertifikatami: Double
	PogashenoSertifikatami: Double
	SummaOplatyBonusom: Double
	VesObmena: Double
	SummaObmena: Double
	VesSkuplennogoTovara: Double
	VydanoZaSkuplennyiTovar: Double
	KolichestvoIzdelii: Int64
	SumManualDiscount: Double
	SumAutoDiscount: Double
	SummaRassrochki: Double
	SummaPogasheniiaRassrochki: Double
	SummaOplatyNalichnymi: Double
	RecorderType: String!
}
input AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseTurnoverInput {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	ProdazhnaiaStoimostTurnover: Double
	VsegoSkidkiTurnover: Double
	SkidkiPoDiskontnymKartamTurnover: Double
	SummaOplatyKartamiTurnover: Double
	SummaOplatyBankovskimKreditomTurnover: Double
	SummaVozvrataTurnover: Double
	VesVChekakhTurnover: Double
	KolichestvoChekovTurnover: Int64
	SummaProdazhiSertifikatovTurnover: Double
	SummaOplatySertifikatamiTurnover: Double
	PogashenoSertifikatamiTurnover: Double
	SummaOplatyBonusomTurnover: Double
	VesObmenaTurnover: Double
	SummaObmenaTurnover: Double
	VesSkuplennogoTovaraTurnover: Double
	VydanoZaSkuplennyiTovarTurnover: Double
	KolichestvoIzdeliiTurnover: Int64
	SummaRuchnoiSkidkiTurnover: Double
	SummaAvtomaticheskoiSkidkiTurnover: Double
	SummaRassrochkiTurnover: Double
	SummaPogasheniiaRassrochkiTurnover: Double
	SummaOplatyNalichnymiTurnover: Double
	RecorderType: String
}
type AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseTurnover {
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	ProdazhnaiaStoimostTurnover: Double
	VsegoSkidkiTurnover: Double
	SkidkiPoDiskontnymKartamTurnover: Double
	SummaOplatyKartamiTurnover: Double
	SummaOplatyBankovskimKreditomTurnover: Double
	SummaVozvrataTurnover: Double
	VesVChekakhTurnover: Double
	KolichestvoChekovTurnover: Int64
	SummaProdazhiSertifikatovTurnover: Double
	SummaOplatySertifikatamiTurnover: Double
	PogashenoSertifikatamiTurnover: Double
	SummaOplatyBonusomTurnover: Double
	VesObmenaTurnover: Double
	SummaObmenaTurnover: Double
	VesSkuplennogoTovaraTurnover: Double
	VydanoZaSkuplennyiTovarTurnover: Double
	KolichestvoIzdeliiTurnover: Int64
	SummaRuchnoiSkidkiTurnover: Double
	SummaAvtomaticheskoiSkidkiTurnover: Double
	SummaRassrochkiTurnover: Double
	SummaPogasheniiaRassrochkiTurnover: Double
	SummaOplatyNalichnymiTurnover: Double
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaVRezerveRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentRezervirovaniiaKey: Guid
	Sum: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
}
type AccumulationRegisterDenezhnyeSredstvaVRezerveRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentRezervirovaniiaKey: Guid
	Sum: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
}
input AccumulationRegisterDenezhnyeSredstvaVRezerveBalanceInput {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentRezervirovaniiaKey: Guid
	SummaBalance: Double
	BankovskiiSchetKassaType: String
}
type AccumulationRegisterDenezhnyeSredstvaVRezerveBalance {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentRezervirovaniiaKey: Guid
	SummaBalance: Double
	BankovskiiSchetKassaType: String
}
input AccumulationRegisterDenezhnyeSredstvaVRezerveTurnoverInput {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentRezervirovaniiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	BankovskiiSchetKassaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaVRezerveTurnover {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentRezervirovaniiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	BankovskiiSchetKassaType: String
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaVRezerveBalanceAndTurnoverInput {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentRezervirovaniiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	BankovskiiSchetKassaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaVRezerveBalanceAndTurnover {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentRezervirovaniiaKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	BankovskiiSchetKassaType: String
	RecorderType: String
}
input AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	RetailCost: Double
	Quantity: Int64
	Weight: Double
	RecorderType: String!
}
type AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	RetailCost: Double
	Quantity: Int64
	Weight: Double
	RecorderType: String!
}
input AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalanceInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	RetailCost: Double
	KolichestvoBalance: Int64
	VesBalance: Double
}
type AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalance {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	RetailCost: Double
	KolichestvoBalance: Int64
	VesBalance: Double
}
input AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	RetailCost: Double
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	RecorderType: String
}
type AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	RetailCost: Double
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	RecorderType: String
}
input AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalanceAndTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	RetailCost: Double
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	RecorderType: String
}
type AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalanceAndTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	RetailCost: Double
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	RecorderType: String
}
input AccumulationRegisterDavalcheskiiMetallPoteriRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Weight: Double
	Protsent: Double
	RecorderType: String!
}
type AccumulationRegisterDavalcheskiiMetallPoteriRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Weight: Double
	Protsent: Double
	RecorderType: String!
}
input AccumulationRegisterDavalcheskiiMetallPoteriTurnoverInput {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VesTurnover: Double
	RecorderType: String
}
type AccumulationRegisterDavalcheskiiMetallPoteriTurnover {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VesTurnover: Double
	RecorderType: String
}
input InformationRegisterTsenyPoPreiskurantuRowTypeInput {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	KamenKey: Guid!
	FormaOgrankiKey: Guid!
	TsvetKamniaKey: Guid!
	GruppaTsvetaKey: Guid!
	GruppaDefektaKey: Guid!
	RassevKey: Guid!
	Razmer1Ot: Double!
	Razmer1Do: Double!
	Cost: Double
	TipTsenKey: Guid
	RecorderType: String
}
type InformationRegisterTsenyPoPreiskurantuRowType {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	KamenKey: Guid!
	FormaOgrankiKey: Guid!
	TsvetKamniaKey: Guid!
	GruppaTsvetaKey: Guid!
	GruppaDefektaKey: Guid!
	RassevKey: Guid!
	Razmer1Ot: Double!
	Razmer1Do: Double!
	Cost: Double
	TipTsenKey: Guid
	RecorderType: String
}
input AccumulationRegisterTovaryVOtboreRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	ZakazKlienta: String
	DepartmentKey: Guid
	KOtboru: Int64
	Otobrano: Int64
	OtobranoVes: Double
	RecorderType: String!
	ZakazKlientaType: String
}
type AccumulationRegisterTovaryVOtboreRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	ZakazKlienta: String
	DepartmentKey: Guid
	KOtboru: Int64
	Otobrano: Int64
	OtobranoVes: Double
	RecorderType: String!
	ZakazKlientaType: String
}
input AccumulationRegisterTovaryVOtboreBalanceInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	ZakazKlienta: String
	DepartmentKey: Guid
	KOtboruBalance: Int64
	OtobranoBalance: Int64
	OtobranoVesBalance: Double
	ZakazKlientaType: String
}
type AccumulationRegisterTovaryVOtboreBalance {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	ZakazKlienta: String
	DepartmentKey: Guid
	KOtboruBalance: Int64
	OtobranoBalance: Int64
	OtobranoVesBalance: Double
	ZakazKlientaType: String
}
input AccumulationRegisterTovaryVOtboreTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	ZakazKlienta: String
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KOtboruTurnover: Int64
	KOtboruReceipt: Int64
	KOtboruExpense: Int64
	OtobranoTurnover: Int64
	OtobranoReceipt: Int64
	OtobranoExpense: Int64
	OtobranoVesTurnover: Double
	OtobranoVesReceipt: Double
	OtobranoVesExpense: Double
	ZakazKlientaType: String
	RecorderType: String
}
type AccumulationRegisterTovaryVOtboreTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	ZakazKlienta: String
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KOtboruTurnover: Int64
	KOtboruReceipt: Int64
	KOtboruExpense: Int64
	OtobranoTurnover: Int64
	OtobranoReceipt: Int64
	OtobranoExpense: Int64
	OtobranoVesTurnover: Double
	OtobranoVesReceipt: Double
	OtobranoVesExpense: Double
	ZakazKlientaType: String
	RecorderType: String
}
input AccumulationRegisterTovaryVOtboreBalanceAndTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	ZakazKlienta: String
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KOtboruOpeningBalance: Int64
	KOtboruTurnover: Int64
	KOtboruReceipt: Int64
	KOtboruExpense: Int64
	KOtboruClosingBalance: Int64
	OtobranoOpeningBalance: Int64
	OtobranoTurnover: Int64
	OtobranoReceipt: Int64
	OtobranoExpense: Int64
	OtobranoClosingBalance: Int64
	OtobranoVesOpeningBalance: Double
	OtobranoVesTurnover: Double
	OtobranoVesReceipt: Double
	OtobranoVesExpense: Double
	OtobranoVesClosingBalance: Double
	ZakazKlientaType: String
	RecorderType: String
}
type AccumulationRegisterTovaryVOtboreBalanceAndTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	ZakazKlienta: String
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KOtboruOpeningBalance: Int64
	KOtboruTurnover: Int64
	KOtboruReceipt: Int64
	KOtboruExpense: Int64
	KOtboruClosingBalance: Int64
	OtobranoOpeningBalance: Int64
	OtobranoTurnover: Int64
	OtobranoReceipt: Int64
	OtobranoExpense: Int64
	OtobranoClosingBalance: Int64
	OtobranoVesOpeningBalance: Double
	OtobranoVesTurnover: Double
	OtobranoVesReceipt: Double
	OtobranoVesExpense: Double
	OtobranoVesClosingBalance: Double
	ZakazKlientaType: String
	RecorderType: String
}
input AccumulationRegisterRealizovannyeTovaryRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	DokumentPostavki: String
	Quantity: Int64
	Weight: Double
	Vyruchka: Double
	RecorderType: String!
	SdelkaType: String
	DokumentPostavkiType: String
}
type AccumulationRegisterRealizovannyeTovaryRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	DokumentPostavki: String
	Quantity: Int64
	Weight: Double
	Vyruchka: Double
	RecorderType: String!
	SdelkaType: String
	DokumentPostavkiType: String
}
input AccumulationRegisterRealizovannyeTovaryBalanceInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	DokumentPostavki: String
	KolichestvoBalance: Int64
	VesBalance: Double
	VyruchkaBalance: Double
	SdelkaType: String
	DokumentPostavkiType: String
}
type AccumulationRegisterRealizovannyeTovaryBalance {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	DokumentPostavki: String
	KolichestvoBalance: Int64
	VesBalance: Double
	VyruchkaBalance: Double
	SdelkaType: String
	DokumentPostavkiType: String
}
input AccumulationRegisterRealizovannyeTovaryTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	DokumentPostavki: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VyruchkaTurnover: Double
	VyruchkaReceipt: Double
	VyruchkaExpense: Double
	SdelkaType: String
	DokumentPostavkiType: String
	RecorderType: String
}
type AccumulationRegisterRealizovannyeTovaryTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	DokumentPostavki: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VyruchkaTurnover: Double
	VyruchkaReceipt: Double
	VyruchkaExpense: Double
	SdelkaType: String
	DokumentPostavkiType: String
	RecorderType: String
}
input AccumulationRegisterRealizovannyeTovaryBalanceAndTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	DokumentPostavki: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	VyruchkaOpeningBalance: Double
	VyruchkaTurnover: Double
	VyruchkaReceipt: Double
	VyruchkaExpense: Double
	VyruchkaClosingBalance: Double
	SdelkaType: String
	DokumentPostavkiType: String
	RecorderType: String
}
type AccumulationRegisterRealizovannyeTovaryBalanceAndTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	DokumentPostavki: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	VyruchkaOpeningBalance: Double
	VyruchkaTurnover: Double
	VyruchkaReceipt: Double
	VyruchkaExpense: Double
	VyruchkaClosingBalance: Double
	SdelkaType: String
	DokumentPostavkiType: String
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaKomissioneraRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKomissioneraRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
input AccumulationRegisterDenezhnyeSredstvaKomissioneraBalanceInput {
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetovBalance: Double
	SummaUprBalance: Double
	SdelkaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKomissioneraBalance {
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetovBalance: Double
	SummaUprBalance: Double
	SdelkaType: String
}
input AccumulationRegisterDenezhnyeSredstvaKomissioneraTurnoverInput {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SdelkaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaKomissioneraTurnover {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SdelkaType: String
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaKomissioneraBalanceAndTurnoverInput {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	SdelkaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaKomissioneraBalanceAndTurnover {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	SdelkaType: String
	RecorderType: String
}
input AccumulationRegisterProdazhi1RowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentProdazhi: String
	DokumentOprikhodovaniia: String
	ZakazKlienta: String
	ProektKey: Guid
	PodrazdelenieKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	OrderKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostBezSkidok: Double
	StoimostPostuplenie: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	KolichestvoVozvrat: Int64
	VesVozvrat: Double
	StoimostVozvrat: Double
	StoimostBezSkidokVozvrat: Double
	StoimostPostuplenieVozvrat: Double
	StoimostUprVozvrat: Double
	StoimostBezNDSVozvrat: Double
	SpisaniePartii: Boolean
	RecorderType: String!
	DokumentProdazhiType: String
	DokumentOprikhodovaniiaType: String
	ZakazKlientaType: String
}
type AccumulationRegisterProdazhi1RowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentProdazhi: String
	DokumentOprikhodovaniia: String
	ZakazKlienta: String
	ProektKey: Guid
	PodrazdelenieKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	OrderKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostBezSkidok: Double
	StoimostPostuplenie: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	KolichestvoVozvrat: Int64
	VesVozvrat: Double
	StoimostVozvrat: Double
	StoimostBezSkidokVozvrat: Double
	StoimostPostuplenieVozvrat: Double
	StoimostUprVozvrat: Double
	StoimostBezNDSVozvrat: Double
	SpisaniePartii: Boolean
	RecorderType: String!
	DokumentProdazhiType: String
	DokumentOprikhodovaniiaType: String
	ZakazKlientaType: String
}
input AccumulationRegisterProdazhi1TurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentProdazhi: String
	DokumentOprikhodovaniia: String
	ZakazKlienta: String
	ProektKey: Guid
	PodrazdelenieKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	OrderKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	VesTurnover: Double
	StoimostTurnover: Double
	StoimostBezSkidokTurnover: Double
	StoimostPostuplenieTurnover: Double
	StoimostUprTurnover: Double
	StoimostBezNDSTurnover: Double
	KolichestvoVozvratTurnover: Int64
	VesVozvratTurnover: Double
	StoimostVozvratTurnover: Double
	StoimostBezSkidokVozvratTurnover: Double
	StoimostPostuplenieVozvratTurnover: Double
	StoimostUprVozvratTurnover: Double
	StoimostBezNDSVozvratTurnover: Double
	DokumentProdazhiType: String
	DokumentOprikhodovaniiaType: String
	ZakazKlientaType: String
	RecorderType: String
}
type AccumulationRegisterProdazhi1Turnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentProdazhi: String
	DokumentOprikhodovaniia: String
	ZakazKlienta: String
	ProektKey: Guid
	PodrazdelenieKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	OrderKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	VesTurnover: Double
	StoimostTurnover: Double
	StoimostBezSkidokTurnover: Double
	StoimostPostuplenieTurnover: Double
	StoimostUprTurnover: Double
	StoimostBezNDSTurnover: Double
	KolichestvoVozvratTurnover: Int64
	VesVozvratTurnover: Double
	StoimostVozvratTurnover: Double
	StoimostBezSkidokVozvratTurnover: Double
	StoimostPostuplenieVozvratTurnover: Double
	StoimostUprVozvratTurnover: Double
	StoimostBezNDSVozvratTurnover: Double
	DokumentProdazhiType: String
	DokumentOprikhodovaniiaType: String
	ZakazKlientaType: String
	RecorderType: String
}
input AccumulationRegisterTovaryNaSkladakhAMRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DepartmentKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	ItemKey: Guid
	Quantity: Int64
	Weight: Double
	RoznichnaiaStoimost: Double
	RecorderType: String!
}
type AccumulationRegisterTovaryNaSkladakhAMRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DepartmentKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	ItemKey: Guid
	Quantity: Int64
	Weight: Double
	RoznichnaiaStoimost: Double
	RecorderType: String!
}
input AccumulationRegisterTovaryNaSkladakhAMBalanceInput {
	DepartmentKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	ItemKey: Guid
	KolichestvoBalance: Int64
	VesBalance: Double
	RoznichnaiaStoimostBalance: Double
}
type AccumulationRegisterTovaryNaSkladakhAMBalance {
	DepartmentKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	ItemKey: Guid
	KolichestvoBalance: Int64
	VesBalance: Double
	RoznichnaiaStoimostBalance: Double
}
input AccumulationRegisterTovaryNaSkladakhAMTurnoverInput {
	DepartmentKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	ItemKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	RoznichnaiaStoimostTurnover: Double
	RoznichnaiaStoimostReceipt: Double
	RoznichnaiaStoimostExpense: Double
	RecorderType: String
}
type AccumulationRegisterTovaryNaSkladakhAMTurnover {
	DepartmentKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	ItemKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	RoznichnaiaStoimostTurnover: Double
	RoznichnaiaStoimostReceipt: Double
	RoznichnaiaStoimostExpense: Double
	RecorderType: String
}
input AccumulationRegisterTovaryNaSkladakhAMBalanceAndTurnoverInput {
	DepartmentKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	ItemKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	RoznichnaiaStoimostOpeningBalance: Double
	RoznichnaiaStoimostTurnover: Double
	RoznichnaiaStoimostReceipt: Double
	RoznichnaiaStoimostExpense: Double
	RoznichnaiaStoimostClosingBalance: Double
	RecorderType: String
}
type AccumulationRegisterTovaryNaSkladakhAMBalanceAndTurnover {
	DepartmentKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	ItemKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	RoznichnaiaStoimostOpeningBalance: Double
	RoznichnaiaStoimostTurnover: Double
	RoznichnaiaStoimostReceipt: Double
	RoznichnaiaStoimostExpense: Double
	RoznichnaiaStoimostClosingBalance: Double
	RecorderType: String
}
input AccumulationRegisterSummyPoFinmonitoringuRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	SummaOtgruzki: Double
	SummaOplaty: Double
	SummaVozvrata: Double
	RecorderType: String!
	DokumentUchetaType: String
}
type AccumulationRegisterSummyPoFinmonitoringuRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	SummaOtgruzki: Double
	SummaOplaty: Double
	SummaVozvrata: Double
	RecorderType: String!
	DokumentUchetaType: String
}
input AccumulationRegisterSummyPoFinmonitoringuBalanceInput {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	SummaOtgruzkiBalance: Double
	SummaOplatyBalance: Double
	SummaVozvrataBalance: Double
	DokumentUchetaType: String
}
type AccumulationRegisterSummyPoFinmonitoringuBalance {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	SummaOtgruzkiBalance: Double
	SummaOplatyBalance: Double
	SummaVozvrataBalance: Double
	DokumentUchetaType: String
}
input AccumulationRegisterSummyPoFinmonitoringuTurnoverInput {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOtgruzkiTurnover: Double
	SummaOtgruzkiReceipt: Double
	SummaOtgruzkiExpense: Double
	SummaOplatyTurnover: Double
	SummaOplatyReceipt: Double
	SummaOplatyExpense: Double
	SummaVozvrataTurnover: Double
	SummaVozvrataReceipt: Double
	SummaVozvrataExpense: Double
	DokumentUchetaType: String
	RecorderType: String
}
type AccumulationRegisterSummyPoFinmonitoringuTurnover {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOtgruzkiTurnover: Double
	SummaOtgruzkiReceipt: Double
	SummaOtgruzkiExpense: Double
	SummaOplatyTurnover: Double
	SummaOplatyReceipt: Double
	SummaOplatyExpense: Double
	SummaVozvrataTurnover: Double
	SummaVozvrataReceipt: Double
	SummaVozvrataExpense: Double
	DokumentUchetaType: String
	RecorderType: String
}
input AccumulationRegisterSummyPoFinmonitoringuBalanceAndTurnoverInput {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOtgruzkiOpeningBalance: Double
	SummaOtgruzkiTurnover: Double
	SummaOtgruzkiReceipt: Double
	SummaOtgruzkiExpense: Double
	SummaOtgruzkiClosingBalance: Double
	SummaOplatyOpeningBalance: Double
	SummaOplatyTurnover: Double
	SummaOplatyReceipt: Double
	SummaOplatyExpense: Double
	SummaOplatyClosingBalance: Double
	SummaVozvrataOpeningBalance: Double
	SummaVozvrataTurnover: Double
	SummaVozvrataReceipt: Double
	SummaVozvrataExpense: Double
	SummaVozvrataClosingBalance: Double
	DokumentUchetaType: String
	RecorderType: String
}
type AccumulationRegisterSummyPoFinmonitoringuBalanceAndTurnover {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOtgruzkiOpeningBalance: Double
	SummaOtgruzkiTurnover: Double
	SummaOtgruzkiReceipt: Double
	SummaOtgruzkiExpense: Double
	SummaOtgruzkiClosingBalance: Double
	SummaOplatyOpeningBalance: Double
	SummaOplatyTurnover: Double
	SummaOplatyReceipt: Double
	SummaOplatyExpense: Double
	SummaOplatyClosingBalance: Double
	SummaVozvrataOpeningBalance: Double
	SummaVozvrataTurnover: Double
	SummaVozvrataReceipt: Double
	SummaVozvrataExpense: Double
	SummaVozvrataClosingBalance: Double
	DokumentUchetaType: String
	RecorderType: String
}
input InformationRegisterTsenyNomenklaturyKontragentovRowTypeInput {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	TipTsenKey: Guid!
	ItemKey: Guid!
	InstanceKey: Guid!
	KharakteristikaNomenklaturyKey: Guid!
	SizeKey: Guid!
	Cost: Double
	ValiutaKey: Guid
	EdinitsaIzmereniiaKey: Guid
	RecorderType: String
}
type InformationRegisterTsenyNomenklaturyKontragentovRowType {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	TipTsenKey: Guid!
	ItemKey: Guid!
	InstanceKey: Guid!
	KharakteristikaNomenklaturyKey: Guid!
	SizeKey: Guid!
	Cost: Double
	ValiutaKey: Guid
	EdinitsaIzmereniiaKey: Guid
	RecorderType: String
}
input AccumulationRegisterVzaimoraschetySKontragentamiRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
type AccumulationRegisterVzaimoraschetySKontragentamiRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
input AccumulationRegisterVzaimoraschetySKontragentamiBalanceInput {
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetovBalance: Double
	SummaUprBalance: Double
	SdelkaType: String
}
type AccumulationRegisterVzaimoraschetySKontragentamiBalance {
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetovBalance: Double
	SummaUprBalance: Double
	SdelkaType: String
}
input AccumulationRegisterVzaimoraschetySKontragentamiTurnoverInput {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SdelkaType: String
	RecorderType: String
}
type AccumulationRegisterVzaimoraschetySKontragentamiTurnover {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SdelkaType: String
	RecorderType: String
}
input AccumulationRegisterVzaimoraschetySKontragentamiBalanceAndTurnoverInput {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	SdelkaType: String
	RecorderType: String
}
type AccumulationRegisterVzaimoraschetySKontragentamiBalanceAndTurnover {
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	SdelkaType: String
	RecorderType: String
}
input AccumulationRegisterSummyPokupokPoDiskontnymKartamRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	MemberCardKey: Guid
	DataSpisaniia: DateTime
	Sum: Double
	SummaBonusov: Double
	SummaVremennykhBonusov: Double
	RecorderType: String!
}
type AccumulationRegisterSummyPokupokPoDiskontnymKartamRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	MemberCardKey: Guid
	DataSpisaniia: DateTime
	Sum: Double
	SummaBonusov: Double
	SummaVremennykhBonusov: Double
	RecorderType: String!
}
input AccumulationRegisterSummyPokupokPoDiskontnymKartamBalanceInput {
	MemberCardKey: Guid
	DataSpisaniia: DateTime
	SummaBalance: Double
	SummaBonusovBalance: Double
	SummaVremennykhBonusovBalance: Double
}
type AccumulationRegisterSummyPokupokPoDiskontnymKartamBalance {
	MemberCardKey: Guid
	DataSpisaniia: DateTime
	SummaBalance: Double
	SummaBonusovBalance: Double
	SummaVremennykhBonusovBalance: Double
}
input AccumulationRegisterSummyPokupokPoDiskontnymKartamTurnoverInput {
	MemberCardKey: Guid
	DataSpisaniia: DateTime
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaBonusovTurnover: Double
	SummaBonusovReceipt: Double
	SummaBonusovExpense: Double
	SummaVremennykhBonusovTurnover: Double
	SummaVremennykhBonusovReceipt: Double
	SummaVremennykhBonusovExpense: Double
	RecorderType: String
}
type AccumulationRegisterSummyPokupokPoDiskontnymKartamTurnover {
	MemberCardKey: Guid
	DataSpisaniia: DateTime
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaBonusovTurnover: Double
	SummaBonusovReceipt: Double
	SummaBonusovExpense: Double
	SummaVremennykhBonusovTurnover: Double
	SummaVremennykhBonusovReceipt: Double
	SummaVremennykhBonusovExpense: Double
	RecorderType: String
}
input AccumulationRegisterSummyPokupokPoDiskontnymKartamBalanceAndTurnoverInput {
	MemberCardKey: Guid
	DataSpisaniia: DateTime
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	SummaBonusovOpeningBalance: Double
	SummaBonusovTurnover: Double
	SummaBonusovReceipt: Double
	SummaBonusovExpense: Double
	SummaBonusovClosingBalance: Double
	SummaVremennykhBonusovOpeningBalance: Double
	SummaVremennykhBonusovTurnover: Double
	SummaVremennykhBonusovReceipt: Double
	SummaVremennykhBonusovExpense: Double
	SummaVremennykhBonusovClosingBalance: Double
	RecorderType: String
}
type AccumulationRegisterSummyPokupokPoDiskontnymKartamBalanceAndTurnover {
	MemberCardKey: Guid
	DataSpisaniia: DateTime
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	SummaBonusovOpeningBalance: Double
	SummaBonusovTurnover: Double
	SummaBonusovReceipt: Double
	SummaBonusovExpense: Double
	SummaBonusovClosingBalance: Double
	SummaVremennykhBonusovOpeningBalance: Double
	SummaVremennykhBonusovTurnover: Double
	SummaVremennykhBonusovReceipt: Double
	SummaVremennykhBonusovExpense: Double
	SummaVremennykhBonusovClosingBalance: Double
	RecorderType: String
}
input AccumulationRegisterVypolneniePlanaProdazhRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	DepartmentKey: Guid
	SummaPlan: Double
	SummaFakt: Double
	RecorderType: String!
}
type AccumulationRegisterVypolneniePlanaProdazhRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	DepartmentKey: Guid
	SummaPlan: Double
	SummaFakt: Double
	RecorderType: String!
}
input AccumulationRegisterVypolneniePlanaProdazhTurnoverInput {
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaPlanTurnover: Double
	SummaFaktTurnover: Double
	RecorderType: String
}
type AccumulationRegisterVypolneniePlanaProdazhTurnover {
	DepartmentKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaPlanTurnover: Double
	SummaFaktTurnover: Double
	RecorderType: String
}
input AccumulationRegisterDavalcheskiiMetallRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Weight: Double
	RecorderType: String!
}
type AccumulationRegisterDavalcheskiiMetallRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Weight: Double
	RecorderType: String!
}
input AccumulationRegisterDavalcheskiiMetallBalanceInput {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	VesBalance: Double
}
type AccumulationRegisterDavalcheskiiMetallBalance {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	VesBalance: Double
}
input AccumulationRegisterDavalcheskiiMetallTurnoverInput {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	RecorderType: String
}
type AccumulationRegisterDavalcheskiiMetallTurnover {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	RecorderType: String
}
input AccumulationRegisterDavalcheskiiMetallBalanceAndTurnoverInput {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	RecorderType: String
}
type AccumulationRegisterDavalcheskiiMetallBalanceAndTurnover {
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
}
type AccumulationRegisterDenezhnyeSredstvaRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
}
input AccumulationRegisterDenezhnyeSredstvaBalanceInput {
	OrganizatsiiaKey: Guid
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	SummaBalance: Double
	SummaUprBalance: Double
	BankovskiiSchetKassaType: String
}
type AccumulationRegisterDenezhnyeSredstvaBalance {
	OrganizatsiiaKey: Guid
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	SummaBalance: Double
	SummaUprBalance: Double
	BankovskiiSchetKassaType: String
}
input AccumulationRegisterDenezhnyeSredstvaTurnoverInput {
	OrganizatsiiaKey: Guid
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	BankovskiiSchetKassaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaTurnover {
	OrganizatsiiaKey: Guid
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	BankovskiiSchetKassaType: String
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaBalanceAndTurnoverInput {
	OrganizatsiiaKey: Guid
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	BankovskiiSchetKassaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaBalanceAndTurnover {
	OrganizatsiiaKey: Guid
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	SummaUprOpeningBalance: Double
	SummaUprTurnover: Double
	SummaUprReceipt: Double
	SummaUprExpense: Double
	SummaUprClosingBalance: Double
	BankovskiiSchetKassaType: String
	RecorderType: String
}
input AccumulationRegisterTovaryPeredannyeRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Quantity: Int64
	Weight: Double
	SummaVzaimoraschetov: Double
	RecorderType: String!
	SdelkaType: String
}
type AccumulationRegisterTovaryPeredannyeRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Quantity: Int64
	Weight: Double
	SummaVzaimoraschetov: Double
	RecorderType: String!
	SdelkaType: String
}
input AccumulationRegisterTovaryPeredannyeBalanceInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	KolichestvoBalance: Int64
	VesBalance: Double
	SummaVzaimoraschetovBalance: Double
	SdelkaType: String
}
type AccumulationRegisterTovaryPeredannyeBalance {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	KolichestvoBalance: Int64
	VesBalance: Double
	SummaVzaimoraschetovBalance: Double
	SdelkaType: String
}
input AccumulationRegisterTovaryPeredannyeTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SdelkaType: String
	RecorderType: String
}
type AccumulationRegisterTovaryPeredannyeTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SdelkaType: String
	RecorderType: String
}
input AccumulationRegisterTovaryPeredannyeBalanceAndTurnoverInput {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SdelkaType: String
	RecorderType: String
}
type AccumulationRegisterTovaryPeredannyeBalanceAndTurnover {
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	KolichestvoOpeningBalance: Int64
	KolichestvoTurnover: Int64
	KolichestvoReceipt: Int64
	KolichestvoExpense: Int64
	KolichestvoClosingBalance: Int64
	VesOpeningBalance: Double
	VesTurnover: Double
	VesReceipt: Double
	VesExpense: Double
	VesClosingBalance: Double
	SummaVzaimoraschetovOpeningBalance: Double
	SummaVzaimoraschetovTurnover: Double
	SummaVzaimoraschetovReceipt: Double
	SummaVzaimoraschetovExpense: Double
	SummaVzaimoraschetovClosingBalance: Double
	SdelkaType: String
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRowTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentSpisaniia: String
	TypeOfMovingMoneyKey: Guid
	Sum: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentSpisaniiaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRowType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentSpisaniia: String
	TypeOfMovingMoneyKey: Guid
	Sum: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentSpisaniiaType: String
}
input AccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalanceInput {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentSpisaniia: String
	TypeOfMovingMoneyKey: Guid
	SummaBalance: Double
	BankovskiiSchetKassaType: String
	DokumentSpisaniiaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalance {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentSpisaniia: String
	TypeOfMovingMoneyKey: Guid
	SummaBalance: Double
	BankovskiiSchetKassaType: String
	DokumentSpisaniiaType: String
}
input AccumulationRegisterDenezhnyeSredstvaKSpisaniiuTurnoverInput {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentSpisaniia: String
	TypeOfMovingMoneyKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	BankovskiiSchetKassaType: String
	DokumentSpisaniiaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaKSpisaniiuTurnover {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentSpisaniia: String
	TypeOfMovingMoneyKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	BankovskiiSchetKassaType: String
	DokumentSpisaniiaType: String
	RecorderType: String
}
input AccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalanceAndTurnoverInput {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentSpisaniia: String
	TypeOfMovingMoneyKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	BankovskiiSchetKassaType: String
	DokumentSpisaniiaType: String
	RecorderType: String
}
type AccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalanceAndTurnover {
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentSpisaniia: String
	TypeOfMovingMoneyKey: Guid
	Period: DateTime
	SecondPeriod: DateTime
	MinutePeriod: DateTime
	HourPeriod: DateTime
	DayPeriod: DateTime
	WeekPeriod: DateTime
	TenDaysPeriod: DateTime
	MonthPeriod: DateTime
	QuarterPeriod: DateTime
	HalfYearPeriod: DateTime
	YearPeriod: DateTime
	Recorder: String
	LineNumber: Int64
	SummaOpeningBalance: Double
	SummaTurnover: Double
	SummaReceipt: Double
	SummaExpense: Double
	SummaClosingBalance: Double
	BankovskiiSchetKassaType: String
	DokumentSpisaniiaType: String
	RecorderType: String
}
input DocumentChekKKMOplataRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	NomerVidaOplaty: Int16
	ProtsentTorgovoiUstupki: Double
	Sum: Double
	SummaTorgovoiUstupki: Double
	Khesh: String
	KartaSberbanka: Boolean
	Poslednie4: String
	KodRRN: String
	Identifikator: String
	TransactionId: String
}
type DocumentChekKKMOplataRowType {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	NomerVidaOplaty: Int16
	ProtsentTorgovoiUstupki: Double
	Sum: Double
	SummaTorgovoiUstupki: Double
	Khesh: String
	KartaSberbanka: Boolean
	Poslednie4: String
	KodRRN: String
	Identifikator: String
	TransactionId: String
}
input DocumentChekKKMOplataSertifikatamiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	SummaPokupkiPogashennaia: Double
}
type DocumentChekKKMOplataSertifikatamiRowType {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	SummaPokupkiPogashennaia: Double
}
input DocumentChekKKMProdazhaSertifikatovRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	SrokDeistviiaSertifikataOgranichen: Boolean
	KolichestvoDneiDeistviiaSertifikata: Int64
}
type DocumentChekKKMProdazhaSertifikatovRowType {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	SrokDeistviiaSertifikataOgranichen: Boolean
	KolichestvoDneiDeistviiaSertifikata: Int64
}
input DocumentChekKKMTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	MID: String
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	RegistratsiiaProdazhi: Boolean
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	Shtrikhkod: String
	OrderKey: Guid
	KliuchSviazi: Int64
	ProdazhaPodarka: Boolean
	SumAutoDiscount: Double
	SummaBonusov: Double
	SostavStrokiDliaRassrochkiKey: Guid
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentChekKKMTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	MID: String
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	RegistratsiiaProdazhi: Boolean
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	Shtrikhkod: String
	OrderKey: Guid
	KliuchSviazi: Int64
	ProdazhaPodarka: Boolean
	SumAutoDiscount: Double
	SummaBonusov: Double
	SostavStrokiDliaRassrochkiKey: Guid
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input DocumentChekKKMDokumentyObmenaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DokumentKey: Guid
	Sum: Double
}
type DocumentChekKKMDokumentyObmenaRowType {
	Key: Guid!
	LineNumber: Int64!
	DokumentKey: Guid
	Sum: Double
}
input DocumentChekKKMDogovoraRassrochkiProdazhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorRassrochkiKey: Guid
	Sum: Double
}
type DocumentChekKKMDogovoraRassrochkiProdazhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorRassrochkiKey: Guid
	Sum: Double
}
input DocumentChekKKMDogovoraRassrochkiOplataRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorRassrochkiKey: Guid
	Sum: Double
}
type DocumentChekKKMDogovoraRassrochkiOplataRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorRassrochkiKey: Guid
	Sum: Double
}
input DocumentChekKKMOplataBallamiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Khesh: String
	Poslednie4: String
	Sum: Double
	Identifikator: String
	TransactionId: String
	TransactionIdSpisaniia: String
}
type DocumentChekKKMOplataBallamiRowType {
	Key: Guid!
	LineNumber: Int64!
	Khesh: String
	Poslednie4: String
	Sum: Double
	Identifikator: String
	TransactionId: String
	TransactionIdSpisaniia: String
}
input DocumentChekKKMSkidkiNatsenkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	Sum: Double
	SkidkaNatsenkaKey: Guid
}
type DocumentChekKKMSkidkiNatsenkiRowType {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	Sum: Double
	SkidkaNatsenkaKey: Guid
}
input DocumentChekKKMUpravliaemyeSkidkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SkidkaNatsenkaKey: Guid
}
type DocumentChekKKMUpravliaemyeSkidkiRowType {
	Key: Guid!
	LineNumber: Int64!
	SkidkaNatsenkaKey: Guid
}
input DocumentChekKKMPodarkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SkidkaNatsenkaKey: Guid
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	SizeKey: Guid
	Quantity: Double
	Weight: Double
	Cost: Double
	Sum: Double
	DepartmentKey: Guid
	KliuchSviazi: Int64
}
type DocumentChekKKMPodarkiRowType {
	Key: Guid!
	LineNumber: Int64!
	SkidkaNatsenkaKey: Guid
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	SizeKey: Guid
	Quantity: Double
	Weight: Double
	Cost: Double
	Sum: Double
	DepartmentKey: Guid
	KliuchSviazi: Int64
}
input DocumentChekKKMKuponyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	KuponKey: Guid
	NomerKupona: String
	Spisyvat: Boolean
}
type DocumentChekKKMKuponyRowType {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	KuponKey: Guid
	NomerKupona: String
	Spisyvat: Boolean
}
input CatalogTipySkidokNatsenokVremiaPoDniamNedeliRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
	Vybran: Boolean
	DenNedeli: String
}
type CatalogTipySkidokNatsenokVremiaPoDniamNedeliRowType {
	Key: Guid!
	LineNumber: Int64!
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
	Vybran: Boolean
	DenNedeli: String
}
input CatalogVidyKodirovokiTsepeiElementyKodirovkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	NeObiazatelnyiElement: Boolean
	Element: String
	SvoistvoKey: Guid
	DlinaElementa: Int64
	NomerStrokiTablChasti: Int64
	PeremennaiaDlina: Boolean
	ZapolniatSvoistvoVSootvetstvii: Boolean
}
type CatalogVidyKodirovokiTsepeiElementyKodirovkiRowType {
	Key: Guid!
	LineNumber: Int64!
	NeObiazatelnyiElement: Boolean
	Element: String
	SvoistvoKey: Guid
	DlinaElementa: Int64
	NomerStrokiTablChasti: Int64
	PeremennaiaDlina: Boolean
	ZapolniatSvoistvoVSootvetstvii: Boolean
}
input CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ZnachenieKodirovki: String
	ZnachenieSvoistvaKey: Guid
	NomerStrokiTablChasti: Int64
	SvoistvoKey: Guid
}
type CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvRowType {
	Key: Guid!
	LineNumber: Int64!
	ZnachenieKodirovki: String
	ZnachenieSvoistvaKey: Guid
	NomerStrokiTablChasti: Int64
	SvoistvoKey: Guid
}
input DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
type DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvaRowType {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
input DocumentOtchetKomitentuOProdazhakhTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	OtchetKomitentuKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPostupleniia: Double
	PokupatelKey: Guid
	DataRealizatsii: DateTime
	DokumentProdazhi: String
	SchetFakturaVystavlennyiKey: Guid
	DokumentPostupleniiaType: String
	DokumentProdazhiType: String
}
type DocumentOtchetKomitentuOProdazhakhTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	OtchetKomitentuKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPostupleniia: Double
	PokupatelKey: Guid
	DataRealizatsii: DateTime
	DokumentProdazhi: String
	SchetFakturaVystavlennyiKey: Guid
	DokumentPostupleniiaType: String
	DokumentProdazhiType: String
}
input DocumentZaiavkaNaPereotsenkuTovarovTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaVRoznitseStaraia: Double
}
type DocumentZaiavkaNaPereotsenkuTovarovTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaVRoznitseStaraia: Double
}
input DocumentZakrytieZakazovKlientovZakazyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ZakazKlientaKey: Guid
	PrichinaZakrytiiaZakazaKey: Guid
}
type DocumentZakrytieZakazovKlientovZakazyRowType {
	Key: Guid!
	LineNumber: Int64!
	ZakazKlientaKey: Guid
	PrichinaZakrytiiaZakazaKey: Guid
}
input DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input CatalogNastroikiVypolneniiaObmenaNastroikiObmenaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VypolniaemoeDeistvie: String
	NastroikaObmenaKey: Guid
	VypolniaemoeDeistvieType: String
}
type CatalogNastroikiVypolneniiaObmenaNastroikiObmenaRowType {
	Key: Guid!
	LineNumber: Int64!
	VypolniaemoeDeistvie: String
	NastroikaObmenaKey: Guid
	VypolniaemoeDeistvieType: String
}
input CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	TekstSoobshcheniia: String
}
type CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamiRowType {
	Key: Guid!
	LineNumber: Int64!
	TekstSoobshcheniia: String
}
input DocumentRealizatsiiaTovarovUslugTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumManualDiscount: Double
	SumAutoDiscount: Double
	VesVstavok: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentRealizatsiiaTovarovUslugTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumManualDiscount: Double
	SumAutoDiscount: Double
	VesVstavok: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input DocumentRealizatsiiaTovarovUslugUslugiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	ProtsentOtSummyDokumenta: Double
	PercentManualDiscount: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
	SumManualDiscount: Double
}
type DocumentRealizatsiiaTovarovUslugUslugiRowType {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	ProtsentOtSummyDokumenta: Double
	PercentManualDiscount: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
	SumManualDiscount: Double
}
input DocumentSobytieStoronnieLitsaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	LitsoKey: Guid
}
type DocumentSobytieStoronnieLitsaRowType {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	LitsoKey: Guid
}
input CatalogNastroikiOtchetovGruppyNastroekIPolzovateliRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Obieekt: String
	ObieektType: String
}
type CatalogNastroikiOtchetovGruppyNastroekIPolzovateliRowType {
	Key: Guid!
	LineNumber: Int64!
	Obieekt: String
	ObieektType: String
}
input CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Kod: Int16
	KonSumma: Double
	MaksimalnyiProtsentSummyPokupkiPriOplateBonusom: Double
	Naimenovanie: String
	NachSumma: Double
	ProtsentSkidki: Double
}
type CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkiRowType {
	Key: Guid!
	LineNumber: Int64!
	Kod: Int16
	KonSumma: Double
	MaksimalnyiProtsentSummyPokupkiPriOplateBonusom: Double
	Naimenovanie: String
	NachSumma: Double
	ProtsentSkidki: Double
}
input CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenokRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	UtsenkaProtsent: Double
	UtsenkaDnei: Int64
}
type CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenokRowType {
	Key: Guid!
	LineNumber: Int64!
	UtsenkaProtsent: Double
	UtsenkaDnei: Int64
}
input DocumentKorrektirovkaDolgaSummyDolgaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DataDolga: DateTime
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	Sdelka: String
	UvelichenieDolgaKontragenta: Double
	UmenshenieDolgaKontragenta: Double
	SdelkaType: String
}
type DocumentKorrektirovkaDolgaSummyDolgaRowType {
	Key: Guid!
	LineNumber: Int64!
	DataDolga: DateTime
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	Sdelka: String
	UvelichenieDolgaKontragenta: Double
	UmenshenieDolgaKontragenta: Double
	SdelkaType: String
}
input DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	ProektKey: Guid
	Sdelka: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	Sum: Double
	SdelkaType: String
}
type DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	ProektKey: Guid
	Sdelka: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	Sum: Double
	SdelkaType: String
}
input DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	MestoRazmeshcheniia: String
	Sum: Double
	MestoRazmeshcheniiaType: String
}
type DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkiRowType {
	Key: Guid!
	LineNumber: Int64!
	MestoRazmeshcheniia: String
	Sum: Double
	MestoRazmeshcheniiaType: String
}
input DocumentZakrytieZakazovPostavshchikamZakazyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ZakazPostavshchikuKey: Guid
	PrichinaZakrytiiaZakazaKey: Guid
}
type DocumentZakrytieZakazovPostavshchikamZakazyRowType {
	Key: Guid!
	LineNumber: Int64!
	ZakazPostavshchikuKey: Guid
	PrichinaZakrytiiaZakazaKey: Guid
}
input DocumentAnketyKlientovDliaFinMonitoringaAnketyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	VPerechne: Boolean
	InostrannoePublichnoeLitso: Boolean
	DokumentSeriia: String
	DokumentNomer: String
}
type DocumentAnketyKlientovDliaFinMonitoringaAnketyRowType {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	VPerechne: Boolean
	InostrannoePublichnoeLitso: Boolean
	DokumentSeriia: String
	DokumentNomer: String
}
input DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentInkassovoePorucheniePeredannoeRekvizityKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentInkassovoePorucheniePeredannoeRekvizityKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Imia: String
	Predstavlenie: String
}
type DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniiaRowType {
	Key: Guid!
	LineNumber: Int64!
	Imia: String
	Predstavlenie: String
}
input DocumentInternetZakazTovaryInternetZakazaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	Sum: Double
	StavkaNDS: String
	SummaNDS: Double
	PodobranPolnostiu: Boolean
	Otkaz: Boolean
}
type DocumentInternetZakazTovaryInternetZakazaRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	Sum: Double
	StavkaNDS: String
	SummaNDS: Double
	PodobranPolnostiu: Boolean
	Otkaz: Boolean
}
input DocumentInternetZakazTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	Sum: Double
	StavkaNDS: String
	SummaNDS: Double
	NomerStrokiZakaza: Int16
	DepartmentKey: Guid
}
type DocumentInternetZakazTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	Sum: Double
	StavkaNDS: String
	SummaNDS: Double
	NomerStrokiZakaza: Int16
	DepartmentKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhBonusyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
	NomerCheka: String
	SummaNachislennogoBonusa: Double
	SummaNachislennogoBonusaRasschitana: Boolean
	SummaOplaty: Double
	SummaPokupki: Double
	OrderKey: Guid
	OpisanieKarty: String
}
type DocumentOtchetORoznichnykhProdazhakhBonusyRowType {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
	NomerCheka: String
	SummaNachislennogoBonusa: Double
	SummaNachislennogoBonusaRasschitana: Boolean
	SummaOplaty: Double
	SummaPokupki: Double
	OrderKey: Guid
	OpisanieKarty: String
}
input DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	BankKreditorKey: Guid
	VidOplatyKey: Guid
	DogovorVzaimoraschetovBankaKreditoraKey: Guid
	NomerCheka: String
	ProtsentBankovskoiKomissii: Double
	Sum: Double
	SummaBankovskoiKomissii: Double
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamiRowType {
	Key: Guid!
	LineNumber: Int64!
	BankKreditorKey: Guid
	VidOplatyKey: Guid
	DogovorVzaimoraschetovBankaKreditoraKey: Guid
	NomerCheka: String
	ProtsentBankovskoiKomissii: Double
	Sum: Double
	SummaBankovskoiKomissii: Double
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	NomerCheka: String
	ProtsentTorgovoiUstupki: Double
	Sum: Double
	SummaTorgovoiUstupki: Double
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamiRowType {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	NomerCheka: String
	ProtsentTorgovoiUstupki: Double
	Sum: Double
	SummaTorgovoiUstupki: Double
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	SertifikatKey: Guid
	SrokDeistviiaSertifikataOgranichen: Boolean
	SummaPokupkiPogashennaia: Double
	SummaSertifikata: Double
	OrderKey: Guid
	NomerSertifikata: String
}
type DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamiRowType {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	SertifikatKey: Guid
	SrokDeistviiaSertifikataOgranichen: Boolean
	SummaPokupkiPogashennaia: Double
	SummaSertifikata: Double
	OrderKey: Guid
	NomerSertifikata: String
}
input DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovRowType {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	MemberCardKey: Guid
	Kassir: String
	Quantity: Int64
	ItemKey: Guid
	NomerCheka: String
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	SumAutoDiscount: Double
	KliuchSviazi: Int64
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	OpisanieKarty: String
	SostavStrokiDliaRassrochkiKey: Guid
	KassirType: String
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentOtchetORoznichnykhProdazhakhTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	MemberCardKey: Guid
	Kassir: String
	Quantity: Int64
	ItemKey: Guid
	NomerCheka: String
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	SumAutoDiscount: Double
	KliuchSviazi: Int64
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	OpisanieKarty: String
	SostavStrokiDliaRassrochkiKey: Guid
	KassirType: String
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	Sum: Double
	DogovorRassrochkiKey: Guid
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhaRowType {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	Sum: Double
	DogovorRassrochkiKey: Guid
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhDokumentyObmenaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DokumentKey: Guid
	Sum: Double
	NomerCheka: String
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhDokumentyObmenaRowType {
	Key: Guid!
	LineNumber: Int64!
	DokumentKey: Guid
	Sum: Double
	NomerCheka: String
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplataRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Sum: Double
	NomerCheka: String
	DogovorRassrochkiKey: Guid
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplataRowType {
	Key: Guid!
	LineNumber: Int64!
	Sum: Double
	NomerCheka: String
	DogovorRassrochkiKey: Guid
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhOplataBallamiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Khesh: String
	Poslednie4: String
	Sum: Double
	Identifikator: String
	TransactionId: String
	TransactionIdSpisaniia: String
	NomerCheka: String
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhOplataBallamiRowType {
	Key: Guid!
	LineNumber: Int64!
	Khesh: String
	Poslednie4: String
	Sum: Double
	Identifikator: String
	TransactionId: String
	TransactionIdSpisaniia: String
	NomerCheka: String
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	Sum: Double
	SkidkaNatsenkaKey: Guid
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkiRowType {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	Sum: Double
	SkidkaNatsenkaKey: Guid
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhKuponyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	KuponKey: Guid
	NomerKupona: String
	Spisyvat: Boolean
}
type DocumentOtchetORoznichnykhProdazhakhKuponyRowType {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	KuponKey: Guid
	NomerKupona: String
	Spisyvat: Boolean
}
input DocumentOtmenaSkidokNomenklaturyDokumentyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	UstanovkaSkidokNomenklaturyKey: Guid
}
type DocumentOtmenaSkidokNomenklaturyDokumentyRowType {
	Key: Guid!
	LineNumber: Int64!
	UstanovkaSkidokNomenklaturyKey: Guid
}
input DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentKassovyiChekKorrektsiiOplataRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
type DocumentKassovyiChekKorrektsiiOplataRowType {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
input DocumentSchetNaOplatuPokupateliuTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumAutoDiscount: Double
	SumManualDiscount: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentSchetNaOplatuPokupateliuTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumAutoDiscount: Double
	SumManualDiscount: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input DocumentSchetNaOplatuPokupateliuUslugiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	PercentManualDiscount: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
	SumManualDiscount: Double
}
type DocumentSchetNaOplatuPokupateliuUslugiRowType {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	PercentManualDiscount: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
	SumManualDiscount: Double
}
input CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VariantPoiskaNePodderzhivaetsia: Boolean
	ImiaNastroikiDliaAlgoritma: String
	ImiaNastroikiDliaPolzovatelia: String
	KodPravilaObmena: String
	NaimenovaniePravilaObmena: String
	NastroikaNePodderzhivaetsia: Boolean
	OpisanieNastroikiDliaPolzovatelia: String
	EtoNastroikaDliaVygruzki: Boolean
}
type CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovRowType {
	Key: Guid!
	LineNumber: Int64!
	VariantPoiskaNePodderzhivaetsia: Boolean
	ImiaNastroikiDliaAlgoritma: String
	ImiaNastroikiDliaPolzovatelia: String
	KodPravilaObmena: String
	NaimenovaniePravilaObmena: String
	NastroikaNePodderzhivaetsia: Boolean
	OpisanieNastroikiDliaPolzovatelia: String
	EtoNastroikaDliaVygruzki: Boolean
}
input CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VygruzhatDannye: Boolean
	VygruzhatPoSsylke: Boolean
	KodPravilaVygruzki: String
	KodPravilaObmena: String
	NaimenovaniePravilaVygruzki: String
	NastroikaNePodderzhivaetsia: Boolean
	EtoNastroikaDliaVygruzki: Boolean
}
type CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhRowType {
	Key: Guid!
	LineNumber: Int64!
	VygruzhatDannye: Boolean
	VygruzhatPoSsylke: Boolean
	KodPravilaVygruzki: String
	KodPravilaObmena: String
	NaimenovaniePravilaVygruzki: String
	NastroikaNePodderzhivaetsia: Boolean
	EtoNastroikaDliaVygruzki: Boolean
}
input CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	TekstSoobshcheniia: String
}
type CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamiRowType {
	Key: Guid!
	LineNumber: Int64!
	TekstSoobshcheniia: String
}
input DocumentVozvratTovarovPostavshchikuTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DefektKey: Guid
	DokumentPostupleniia: String
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	DokumentPostupleniiaType: String
}
type DocumentVozvratTovarovPostavshchikuTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DefektKey: Guid
	DokumentPostupleniia: String
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	DokumentPostupleniiaType: String
}
input DocumentInventarizatsiiaTovarovNaSkladeTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesUchet: Double
	KachestvoKey: Guid
	Quantity: Int64
	KolichestvoUchet: Int64
	ItemKey: Guid
	NomerVed: String
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaRegl: Double
	SummaUchet: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	OtkloneniePoKolichestvu: Int64
	OtkloneniePoVesu: Double
}
type DocumentInventarizatsiiaTovarovNaSkladeTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesUchet: Double
	KachestvoKey: Guid
	Quantity: Int64
	KolichestvoUchet: Int64
	ItemKey: Guid
	NomerVed: String
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaRegl: Double
	SummaUchet: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	OtkloneniePoKolichestvu: Int64
	OtkloneniePoVesu: Double
}
input DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidSravneniia: String
	Znachenie: String
	ImiaPolia: String
	ZnachenieType: String
}
type DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiRowType {
	Key: Guid!
	LineNumber: Int64!
	VidSravneniia: String
	Znachenie: String
	ImiaPolia: String
	ZnachenieType: String
}
input DocumentInventarizatsiiaTovarovNaSkladeSertifikatyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
	SummaUchet: Double
	Quantity: Double
	KolichestvoUchet: Double
	OtkloneniePoKolichestvu: Int64
}
type DocumentInventarizatsiiaTovarovNaSkladeSertifikatyRowType {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
	SummaUchet: Double
	Quantity: Double
	KolichestvoUchet: Double
	OtkloneniePoKolichestvu: Int64
}
input DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Weight: Double
	Quantity: Double
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KontragentKey: Guid
	DataOtpravki: DateTime
	DokumentOprikhodovaniiaType: String
}
type DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutiRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Weight: Double
	Quantity: Double
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KontragentKey: Guid
	DataOtpravki: DateTime
	DokumentOprikhodovaniiaType: String
}
input DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPrikhodnyiKassovyiOrderOplataRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
type DocumentPrikhodnyiKassovyiOrderOplataRowType {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
input DocumentPrikhodnyiKassovyiOrderTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SummaSkidki: Double
	VidTovaraKKT: String
	TipOplatyTovaraKKT: String
	SummaOsn: Double
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	SummaOpl: Double
}
type DocumentPrikhodnyiKassovyiOrderTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SummaSkidki: Double
	VidTovaraKKT: String
	TipOplatyTovaraKKT: String
	SummaOsn: Double
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	SummaOpl: Double
}
input DocumentVozvratMaterialovIzProizvodstvaTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
type DocumentVozvratMaterialovIzProizvodstvaTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
input DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaStaraia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaZaGramm: Double
}
type DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaStaraia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaZaGramm: Double
}
input DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	PokupatelKey: Guid
	DogovorSPokupatelemKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Sum: Double
	SummaSebestoimost: Double
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamiRowType {
	Key: Guid!
	LineNumber: Int64!
	PokupatelKey: Guid
	DogovorSPokupatelemKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Sum: Double
	SummaSebestoimost: Double
}
input DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyeRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	PokupatelKey: Guid
	DogovorSPokupatelemKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Quantity: Int64
	Weight: Double
	SummaPostupleniia: Double
	SummaProdazhi: Double
	StatusRaskhoda: String
	DokumentProdazhi: String
	DokumentOprikhodovaniia: String
	DokumentProdazhiType: String
	DokumentOprikhodovaniiaType: String
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyeRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	PokupatelKey: Guid
	DogovorSPokupatelemKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Quantity: Int64
	Weight: Double
	SummaPostupleniia: Double
	SummaProdazhi: Double
	StatusRaskhoda: String
	DokumentProdazhi: String
	DokumentOprikhodovaniia: String
	DokumentProdazhiType: String
	DokumentOprikhodovaniiaType: String
}
input DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Sum: Double
	SummaSebestoimost: Double
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamiRowType {
	Key: Guid!
	LineNumber: Int64!
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Sum: Double
	SummaSebestoimost: Double
}
input DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Quantity: Int64
	Weight: Double
	SummaPostupleniia: Double
	StatusRaskhoda: String
	DokumentOprikhodovaniia: String
	SummaPostupleniiaBezNDS: Double
	DokumentOprikhodovaniiaType: String
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Quantity: Int64
	Weight: Double
	SummaPostupleniia: Double
	StatusRaskhoda: String
	DokumentOprikhodovaniia: String
	SummaPostupleniiaBezNDS: Double
	DokumentOprikhodovaniiaType: String
}
input DocumentGTDImportRazdelyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	NDSVValiute: Boolean
	PoshlinaVValiute: Boolean
	StavkaNDS: String
	StavkaPoshliny: Double
	SummaNDS: Double
	SummaPoshliny: Double
	TamozhennaiaStoimost: Double
	TamozhennaiaStoimostVValiuteReglUcheta: Boolean
}
type DocumentGTDImportRazdelyRowType {
	Key: Guid!
	LineNumber: Int64!
	NDSVValiute: Boolean
	PoshlinaVValiute: Boolean
	StavkaNDS: String
	StavkaPoshliny: Double
	SummaNDS: Double
	SummaPoshliny: Double
	TamozhennaiaStoimost: Double
	TamozhennaiaStoimostVValiuteReglUcheta: Boolean
}
input DocumentGTDImportTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartii: String
	Quantity: Int64
	ItemKey: Guid
	NomerRazdela: Int16
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	SummaNDS: Double
	SummaPoshliny: Double
	FakturnaiaStoimost: Double
	KharakteristikaNomenklaturyKey: Guid
	DokumentPartiiType: String
}
type DocumentGTDImportTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartii: String
	Quantity: Int64
	ItemKey: Guid
	NomerRazdela: Int16
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	SummaNDS: Double
	SummaPoshliny: Double
	FakturnaiaStoimost: Double
	KharakteristikaNomenklaturyKey: Guid
	DokumentPartiiType: String
}
input DocumentAktSverkiTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	MID: String
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	Koef: Double
	NaborKey: Guid
	Naideno: Boolean
	ItemKey: Guid
	NomerNabora: Int64
	Pasport: String
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	Period: DateTime
}
type DocumentAktSverkiTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	MID: String
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	Koef: Double
	NaborKey: Guid
	Naideno: Boolean
	ItemKey: Guid
	NomerNabora: Int64
	Pasport: String
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	Period: DateTime
}
input CatalogFailyDopolnitelnyeRekvizityRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SvoistvoKey: Guid
	Znachenie: String
	TekstovaiaStroka: String
	ZnachenieType: String
}
type CatalogFailyDopolnitelnyeRekvizityRowType {
	Key: Guid!
	LineNumber: Int64!
	SvoistvoKey: Guid
	Znachenie: String
	TekstovaiaStroka: String
	ZnachenieType: String
}
input CatalogFailySertifikatyShifrovaniiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Otpechatok: String
	Predstavlenie: String
	SertifikatBase64Data: Binary
	SertifikatType: String
	Sertifikat: Stream
}
type CatalogFailySertifikatyShifrovaniiaRowType {
	Key: Guid!
	LineNumber: Int64!
	Otpechatok: String
	Predstavlenie: String
	SertifikatBase64Data: Binary
	SertifikatType: String
	Sertifikat: Stream
}
input CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Administrirovanie: Boolean
	Otpravka: Boolean
	PolzovatelKey: Guid
	Transport: Boolean
}
type CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisiRowType {
	Key: Guid!
	LineNumber: Int64!
	Administrirovanie: Boolean
	Otpravka: Boolean
	PolzovatelKey: Guid
	Transport: Boolean
}
input DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	ProektKey: Guid
	Sdelka: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	Sum: Double
	SdelkaType: String
}
type DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	ProektKey: Guid
	Sdelka: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	Sum: Double
	SdelkaType: String
}
input DocumentSkupkaTovarovTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ItemKey: Guid
	ObshchiiVes: Double
	Sum: Double
	TsenaZaGr: Double
	Opisanie: String
}
type DocumentSkupkaTovarovTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ItemKey: Guid
	ObshchiiVes: Double
	Sum: Double
	TsenaZaGr: Double
	Opisanie: String
}
input DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliamRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	PokupatelKey: Guid
	SchetFakturaKey: Guid
	SubkomissionerKey: Guid
	Sum: Double
	NDS: Double
	SummaUvelichenie: Double
	SummaUmenshenie: Double
	SummaNDSUvelichenie: Double
	SummaNDSUmenshenie: Double
}
type DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliamRowType {
	Key: Guid!
	LineNumber: Int64!
	PokupatelKey: Guid
	SchetFakturaKey: Guid
	SubkomissionerKey: Guid
	Sum: Double
	NDS: Double
	SummaUvelichenie: Double
	SummaUmenshenie: Double
	SummaNDSUvelichenie: Double
	SummaNDSUmenshenie: Double
}
input CatalogfmKartochkaKontragentaDannyeKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
type CatalogfmKartochkaKontragentaDannyeKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
input DocumentSpisanieProsrochennykhSertifikatovSertifikatyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	SertifikatKey: Guid
	SrokDeistviiaDo: DateTime
	Sum: Double
	OrganizatsiiaKey: Guid
	DokumentProdazhi: String
	DokumentProdazhiType: String
}
type DocumentSpisanieProsrochennykhSertifikatovSertifikatyRowType {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	SertifikatKey: Guid
	SrokDeistviiaDo: DateTime
	Sum: Double
	OrganizatsiiaKey: Guid
	DokumentProdazhi: String
	DokumentProdazhiType: String
}
input DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
}
type DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentyRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
}
input DocumentAkkreditivPeredannyiRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentAkkreditivPeredannyiRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentAkkreditivPeredannyiRekvizityKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentAkkreditivPeredannyiRekvizityKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input CatalogKontragentyVidyDeiatelnostiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidDeiatelnostiKey: Guid
}
type CatalogKontragentyVidyDeiatelnostiRowType {
	Key: Guid!
	LineNumber: Int64!
	VidDeiatelnostiKey: Guid
}
input DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentMarketingovaiaAktsiiaMagazinyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	MagazinKey: Guid
}
type DocumentMarketingovaiaAktsiiaMagazinyRowType {
	Key: Guid!
	LineNumber: Int64!
	MagazinKey: Guid
}
input DocumentMarketingovaiaAktsiiaSkidkiNatsenkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DataNachala: DateTime
	DataOkonchaniia: DateTime
	MagazinKey: Guid
	SkidkaNatsenkaKey: Guid
}
type DocumentMarketingovaiaAktsiiaSkidkiNatsenkiRowType {
	Key: Guid!
	LineNumber: Int64!
	DataNachala: DateTime
	DataOkonchaniia: DateTime
	MagazinKey: Guid
	SkidkaNatsenkaKey: Guid
}
input DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	NomerNabora: Int64
	VidDostupa: String
	ZnachenieDostupa: String
	Chtenie: Boolean
	Dobavlenie: Boolean
	Izmenenie: Boolean
	Udalenie: Boolean
	ZnachenieDostupaType: String
}
type DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupaRowType {
	Key: Guid!
	LineNumber: Int64!
	NomerNabora: Int64
	VidDostupa: String
	ZnachenieDostupa: String
	Chtenie: Boolean
	Dobavlenie: Boolean
	Izmenenie: Boolean
	Udalenie: Boolean
	ZnachenieDostupaType: String
}
input CatalogStsenariiObmenovDannymiNastroikiObmenaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	UzelInformatsionnoiBazy: String
	VidTransportaObmena: String
	VypolniaemoeDeistvie: String
	KolichestvoElementovVTranzaktsii: Int64
	UzelInformatsionnoiBazyType: String
}
type CatalogStsenariiObmenovDannymiNastroikiObmenaRowType {
	Key: Guid!
	LineNumber: Int64!
	UzelInformatsionnoiBazy: String
	VidTransportaObmena: String
	VypolniaemoeDeistvie: String
	KolichestvoElementovVTranzaktsii: Int64
	UzelInformatsionnoiBazyType: String
}
input CatalogNomenklaturaSostavLigaturyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Double
	ItemKey: Guid
}
type CatalogNomenklaturaSostavLigaturyRowType {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Double
	ItemKey: Guid
}
input DocumentOprosVoprosyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	Otvet: String
	TipovoiOtvet: String
	TipovoiOtvetType: String
}
type DocumentOprosVoprosyRowType {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	Otvet: String
	TipovoiOtvet: String
	TipovoiOtvetType: String
}
input DocumentOprosSostavnoiOtvetRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	VoprosVladeletsKey: Guid
	NomerStrokiVTablitse: Int64
	Otvet: String
	TipovoiOtvet: String
	TipovoiOtvetType: String
}
type DocumentOprosSostavnoiOtvetRowType {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	VoprosVladeletsKey: Guid
	NomerStrokiVTablitse: Int64
	Otvet: String
	TipovoiOtvet: String
	TipovoiOtvetType: String
}
input DocumentPereotsenkaTovarovVNTTTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaVRoznitseStaraia: Double
	Naideno: Boolean
	Dnei: Int64
	DogovorKontragentaKey: Guid
	ProtsentUtsenki: Double
}
type DocumentPereotsenkaTovarovVNTTTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaVRoznitseStaraia: Double
	Naideno: Boolean
	Dnei: Int64
	DogovorKontragentaKey: Guid
	ProtsentUtsenki: Double
}
input CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGruppRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VstavkiBase64Data: Binary
	MetallKey: Guid
	Naimenovanie: String
	UsloviiaVkhozhdeniiaBase64Data: Binary
	VstavkiType: String
	UsloviiaVkhozhdeniiaType: String
	Vstavki: Stream
	UsloviiaVkhozhdeniia: Stream
}
type CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGruppRowType {
	Key: Guid!
	LineNumber: Int64!
	VstavkiBase64Data: Binary
	MetallKey: Guid
	Naimenovanie: String
	UsloviiaVkhozhdeniiaBase64Data: Binary
	VstavkiType: String
	UsloviiaVkhozhdeniiaType: String
	Vstavki: Stream
	UsloviiaVkhozhdeniia: Stream
}
input CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	MiksyVstavokBase64Data: Binary
	Naimenovanie: String
	NomenklaturnaiaGruppaKey: Guid
	UsloviiaVkhozhdeniiaBase64Data: Binary
	MiksyVstavokType: String
	UsloviiaVkhozhdeniiaType: String
	MiksyVstavok: Stream
	UsloviiaVkhozhdeniia: Stream
}
type CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriiRowType {
	Key: Guid!
	LineNumber: Int64!
	MiksyVstavokBase64Data: Binary
	Naimenovanie: String
	NomenklaturnaiaGruppaKey: Guid
	UsloviiaVkhozhdeniiaBase64Data: Binary
	MiksyVstavokType: String
	UsloviiaVkhozhdeniiaType: String
	MiksyVstavok: Stream
	UsloviiaVkhozhdeniia: Stream
}
input CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SvoistvoTovara: String
	UsloviiaVkhozhdeniiaBase64Data: Binary
	Naimenovanie: String
	UsloviiaVkhozhdeniiaType: String
	UsloviiaVkhozhdeniia: Stream
}
type CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiiRowType {
	Key: Guid!
	LineNumber: Int64!
	SvoistvoTovara: String
	UsloviiaVkhozhdeniiaBase64Data: Binary
	Naimenovanie: String
	UsloviiaVkhozhdeniiaType: String
	UsloviiaVkhozhdeniia: Stream
}
input DocumentPeremeshchenieTovarovSertifikatyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
type DocumentPeremeshchenieTovarovSertifikatyRowType {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
input DocumentPeremeshchenieTovarovTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentRezerva: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StoimostBezNDS: Double
	StoimostSNDS: Double
	Sum: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaVRoznitseGr: Double
	TsenaPostupleniia: Double
	Naideno: Boolean
	InternetZakazKey: Guid
	DokumentRezervaType: String
}
type DocumentPeremeshchenieTovarovTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentRezerva: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StoimostBezNDS: Double
	StoimostSNDS: Double
	Sum: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaVRoznitseGr: Double
	TsenaPostupleniia: Double
	Naideno: Boolean
	InternetZakazKey: Guid
	DokumentRezervaType: String
}
input DocumentPeremeshchenieTovarovSpisokZaiavokRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ZaiavkaNaPeremeshchenie: String
}
type DocumentPeremeshchenieTovarovSpisokZaiavokRowType {
	Key: Guid!
	LineNumber: Int64!
	ZaiavkaNaPeremeshchenie: String
}
input DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ValiutaZaiavkaKey: Guid
	ZaiavkaKey: Guid
	KontragentKey: Guid
	OstatokZaiavka: Double
	OstatokRazmeshchenie: Double
	OstatokRezerv: Double
	OtvetstvennyiKey: Guid
}
type DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvRowType {
	Key: Guid!
	LineNumber: Int64!
	ValiutaZaiavkaKey: Guid
	ZaiavkaKey: Guid
	KontragentKey: Guid
	OstatokZaiavka: Double
	OstatokRazmeshchenie: Double
	OstatokRezerv: Double
	OtvetstvennyiKey: Guid
}
input DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ABCKlassifikatsiia: String
	ABCKlassifikatsiiaStaraia: String
	ZnachenieParametra: Double
	KontragentKey: Guid
	MenedzherKontragentaKey: Guid
}
type DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovRowType {
	Key: Guid!
	LineNumber: Int64!
	ABCKlassifikatsiia: String
	ABCKlassifikatsiiaStaraia: String
	ZnachenieParametra: Double
	KontragentKey: Guid
	MenedzherKontragentaKey: Guid
}
input DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
	SummaUchet: Double
	Quantity: Double
	KolichestvoUchet: Double
}
type DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatyRowType {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
	SummaUchet: Double
	Quantity: Double
	KolichestvoUchet: Double
}
input DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidSravneniia: String
	Znachenie: String
	ImiaPolia: String
	ZnachenieType: String
}
type DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiRowType {
	Key: Guid!
	LineNumber: Int64!
	VidSravneniia: String
	Znachenie: String
	ImiaPolia: String
	ZnachenieType: String
}
input DocumentKorrektirovkaRealizatsiiTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VesDoKorrektirovki: Double
	VesDoIzmeneniia: Double
	Weight: Double
	KolichestvoDoKorrektirovki: Int64
	KolichestvoDoIzmeneniia: Int64
	Quantity: Int64
	DokumentOprikhodovaniia: String
	DokumentPartii: String
	KachestvoKey: Guid
	ItemKey: Guid
	Pasport: String
	RazmerDoKorrektirovkiKey: Guid
	RazmerDoIzmeneniiaKey: Guid
	SizeKey: Guid
	SeriiaNomenklaturyDoKorrektirovkiKey: Guid
	SeriiaNomenklaturyDoIzmeneniiaKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDSDoKorrektirovki: String
	StavkaNDSDoIzmeneniia: String
	StavkaNDS: String
	StatusPartii: String
	SummaDoKorrektirovki: Double
	SummaDoIzmeneniia: Double
	Sum: Double
	SummaNDSDoKorrektirovki: Double
	SummaNDSDoIzmeneniia: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyDoKorrektirovkiKey: Guid
	KharakteristikaNomenklaturyDoIzmeneniiaKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TsenaDoKorrektirovki: Double
	TsenaDoIzmeneniia: Double
	Cost: Double
	DokumentOprikhodovaniiaType: String
	DokumentPartiiType: String
}
type DocumentKorrektirovkaRealizatsiiTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	VesDoKorrektirovki: Double
	VesDoIzmeneniia: Double
	Weight: Double
	KolichestvoDoKorrektirovki: Int64
	KolichestvoDoIzmeneniia: Int64
	Quantity: Int64
	DokumentOprikhodovaniia: String
	DokumentPartii: String
	KachestvoKey: Guid
	ItemKey: Guid
	Pasport: String
	RazmerDoKorrektirovkiKey: Guid
	RazmerDoIzmeneniiaKey: Guid
	SizeKey: Guid
	SeriiaNomenklaturyDoKorrektirovkiKey: Guid
	SeriiaNomenklaturyDoIzmeneniiaKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDSDoKorrektirovki: String
	StavkaNDSDoIzmeneniia: String
	StavkaNDS: String
	StatusPartii: String
	SummaDoKorrektirovki: Double
	SummaDoIzmeneniia: Double
	Sum: Double
	SummaNDSDoKorrektirovki: Double
	SummaNDSDoIzmeneniia: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyDoKorrektirovkiKey: Guid
	KharakteristikaNomenklaturyDoIzmeneniiaKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TsenaDoKorrektirovki: Double
	TsenaDoIzmeneniia: Double
	Cost: Double
	DokumentOprikhodovaniiaType: String
	DokumentPartiiType: String
}
input DocumentKorrektirovkaRealizatsiiUslugiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KolichestvoDoKorrektirovki: Int64
	KolichestvoDoIzmeneniia: Int64
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDSDoKorrektirovki: String
	StavkaNDSDoIzmeneniia: String
	StavkaNDS: String
	SummaDoKorrektirovki: Double
	SummaDoIzmeneniia: Double
	Sum: Double
	SummaNDSDoKorrektirovki: Double
	SummaNDSDoIzmeneniia: Double
	SummaNDS: Double
	TsenaDoKorrektirovki: Double
	TsenaDoIzmeneniia: Double
	Cost: Double
}
type DocumentKorrektirovkaRealizatsiiUslugiRowType {
	Key: Guid!
	LineNumber: Int64!
	KolichestvoDoKorrektirovki: Int64
	KolichestvoDoIzmeneniia: Int64
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDSDoKorrektirovki: String
	StavkaNDSDoIzmeneniia: String
	StavkaNDS: String
	SummaDoKorrektirovki: Double
	SummaDoIzmeneniia: Double
	Sum: Double
	SummaNDSDoKorrektirovki: Double
	SummaNDSDoIzmeneniia: Double
	SummaNDS: Double
	TsenaDoKorrektirovki: Double
	TsenaDoIzmeneniia: Double
	Cost: Double
}
input DocumentDoverennostTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaPoKlassifikatoruKey: Guid
	Quantity: Int64
	NaimenovanieTovara: String
	SizeKey: Guid
}
type DocumentDoverennostTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaPoKlassifikatoruKey: Guid
	Quantity: Int64
	NaimenovanieTovara: String
	SizeKey: Guid
}
input CatalogShablonyZapolneniiaKUPrazdnichnyeDniRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Den: DateTime
	KU: Double
	Opisanie: String
	TsvetFonaDliaKalendaria: String
	TsvetTekstaDliaKalendaria: String
}
type CatalogShablonyZapolneniiaKUPrazdnichnyeDniRowType {
	Key: Guid!
	LineNumber: Int64!
	Den: DateTime
	KU: Double
	Opisanie: String
	TsvetFonaDliaKalendaria: String
	TsvetTekstaDliaKalendaria: String
}
input CatalogShablonyZapolneniiaKUKUNaNedeliuRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DenNedeli: Int16
	KU: Double
}
type CatalogShablonyZapolneniiaKUKUNaNedeliuRowType {
	Key: Guid!
	LineNumber: Int64!
	DenNedeli: Int16
	KU: Double
}
input CatalogShablonyZapolneniiaKUSalonyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
}
type CatalogShablonyZapolneniiaKUSalonyRowType {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
}
input DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrinRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	TypeKey: Guid
	AnalitikaTipaIzdeliiaKey: Guid
	PloshchadVykladki: Double
	NormativnaiaPloshchadVykladki: Double
	KolichestvoIzdelii: Int16
}
type DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrinRowType {
	Key: Guid!
	LineNumber: Int64!
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	TypeKey: Guid
	AnalitikaTipaIzdeliiaKey: Guid
	PloshchadVykladki: Double
	NormativnaiaPloshchadVykladki: Double
	KolichestvoIzdelii: Int16
}
input DocumentVozvratProduktsiiVProizvodstvoTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	StoimostRabot: Double
	StoimostVstavok: Double
	StoimostMetalla: Double
	SummaNDSVstavok: Double
	SummaNDSMetalla: Double
	SummaNDSRabot: Double
	DefektKey: Guid
	VesVstavok: Double
}
type DocumentVozvratProduktsiiVProizvodstvoTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	StoimostRabot: Double
	StoimostVstavok: Double
	StoimostMetalla: Double
	SummaNDSVstavok: Double
	SummaNDSMetalla: Double
	SummaNDSRabot: Double
	DefektKey: Guid
	VesVstavok: Double
}
input CatalogNastroikiRabochegoMestaPolzovateliaNastroikiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Dostupnost: Boolean
	DostupnostPechati: Boolean
	DostupnostRedaktirovaniia: Boolean
	DostupnostVersionirovaniia: Boolean
	PutKElementuFormy: String
	TipObieekta: String
	ElementFormyRabochegoMesta: String
}
type CatalogNastroikiRabochegoMestaPolzovateliaNastroikiRowType {
	Key: Guid!
	LineNumber: Int64!
	Dostupnost: Boolean
	DostupnostPechati: Boolean
	DostupnostRedaktirovaniia: Boolean
	DostupnostVersionirovaniia: Boolean
	PutKElementuFormy: String
	TipObieekta: String
	ElementFormyRabochegoMesta: String
}
input DocumentSpisanieTovarovTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentRezerva: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	SkidkaNatsenkaKey: Guid
	DokumentRezervaType: String
}
type DocumentSpisanieTovarovTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentRezerva: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	SkidkaNatsenkaKey: Guid
	DokumentRezervaType: String
}
input DocumentSpisanieTovarovSertifikatyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
type DocumentSpisanieTovarovSertifikatyRowType {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
input DocumentsmsSoobshcheniePoluchateliRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Kontragent: String
	NomerTelefona: String
	GUID: String
	TekstSoobshcheniia: String
	DataZaversheniia: DateTime
	Millisekunda: Int16
	Status: String
	OpisanieOshibki: String
	MemberCardKey: Guid
}
type DocumentsmsSoobshcheniePoluchateliRowType {
	Key: Guid!
	LineNumber: Int64!
	Kontragent: String
	NomerTelefona: String
	GUID: String
	TekstSoobshcheniia: String
	DataZaversheniia: DateTime
	Millisekunda: Int16
	Status: String
	OpisanieOshibki: String
	MemberCardKey: Guid
}
input CatalogKalendariPlanirovaniiaProdazhKUPoDniamRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DenGoda: DateTime
	KU: Double
}
type CatalogKalendariPlanirovaniiaProdazhKUPoDniamRowType {
	Key: Guid!
	LineNumber: Int64!
	DenGoda: DateTime
	KU: Double
}
input CatalogKalendariPlanirovaniiaProdazhSalonyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
}
type CatalogKalendariPlanirovaniiaProdazhSalonyRowType {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
}
input CatalogTipovyeAnketyVoprosyAnketyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	RazdelKey: Guid
}
type CatalogTipovyeAnketyVoprosyAnketyRowType {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	RazdelKey: Guid
}
input DocumentNachislenieSpisanieBonusovDiskontnyeKartyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
	SummaBonusov: Int64
	TekushchaiaSummaBonusov: Int64
}
type DocumentNachislenieSpisanieBonusovDiskontnyeKartyRowType {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
	SummaBonusov: Int64
	TekushchaiaSummaBonusov: Int64
}
input DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	NomerPlatezha: Int16
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	NomerPlatezha: Int16
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ValiutaPostuplenieKey: Guid
	DokumentPlanirovaniiaKey: Guid
	KontragentKey: Guid
	OstatokPostuplenie: Double
	OstatokRazmeshchenie: Double
	OtvetstvennyiKey: Guid
}
type DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSRowType {
	Key: Guid!
	LineNumber: Int64!
	ValiutaPostuplenieKey: Guid
	DokumentPlanirovaniiaKey: Guid
	KontragentKey: Guid
	OstatokPostuplenie: Double
	OstatokRazmeshchenie: Double
	OtvetstvennyiKey: Guid
}
input DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DokumentUcheta: String
	SummaOtgruzki: Double
	SummaOplaty: Double
	Sum: Double
	SummaVozvrata: Double
	Comment: String
	KodVidaDokumentaKey: Guid
	DokumentUchetaType: String
}
type DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringaRowType {
	Key: Guid!
	LineNumber: Int64!
	DokumentUcheta: String
	SummaOtgruzki: Double
	SummaOplaty: Double
	Sum: Double
	SummaVozvrata: Double
	Comment: String
	KodVidaDokumentaKey: Guid
	DokumentUchetaType: String
}
input DocumentOtchetPoFinMonitoringuDannyeDokumentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
type DocumentOtchetPoFinMonitoringuDannyeDokumentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
input CatalogVersiiFailovElektronnyePodpisiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DataPodpisi: DateTime
	ImiaFailaPodpisi: String
	Comment: String
	KomuVydanSertifikat: String
	Otpechatok: String
	PodpisBase64Data: Binary
	UstanovivshiiPodpisKey: Guid
	SertifikatBase64Data: Binary
	PodpisType: String
	SertifikatType: String
	Podpis: Stream
	Sertifikat: Stream
}
type CatalogVersiiFailovElektronnyePodpisiRowType {
	Key: Guid!
	LineNumber: Int64!
	DataPodpisi: DateTime
	ImiaFailaPodpisi: String
	Comment: String
	KomuVydanSertifikat: String
	Otpechatok: String
	PodpisBase64Data: Binary
	UstanovivshiiPodpisKey: Guid
	SertifikatBase64Data: Binary
	PodpisType: String
	SertifikatType: String
	Podpis: Stream
	Sertifikat: Stream
}
input DocumentUstanovkaTsenNomenklaturyTipyTsenRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	TipTsenKey: Guid
}
type DocumentUstanovkaTsenNomenklaturyTipyTsenRowType {
	Key: Guid!
	LineNumber: Int64!
	TipTsenKey: Guid
}
input DocumentUstanovkaTsenNomenklaturyTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ValiutaKey: Guid
	IndeksStrokiTablitsyTsen: Int64
	ItemKey: Guid
	ProtsentSkidkiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	TipTsenKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SegmentNomenklaturyKey: Guid
}
type DocumentUstanovkaTsenNomenklaturyTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	ValiutaKey: Guid
	IndeksStrokiTablitsyTsen: Int64
	ItemKey: Guid
	ProtsentSkidkiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	TipTsenKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SegmentNomenklaturyKey: Guid
}
input DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentPreiskurantNaSkupkuProbyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ProbeKey: Guid
	TsenaZaGramm: Double
}
type DocumentPreiskurantNaSkupkuProbyRowType {
	Key: Guid!
	LineNumber: Int64!
	ProbeKey: Guid
	TsenaZaGramm: Double
}
input DocumentPeredachaMaterialovVProizvodstvoTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
type DocumentPeredachaMaterialovVProizvodstvoTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
input DocumentVnutrenniiZakazTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	RazmeshchenieKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
type DocumentVnutrenniiZakazTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	RazmeshchenieKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
input CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnostRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DopolnitelnyeParametryObrabotkiBase64Data: Binary
	PredstavlenieKnopki: String
	PredstavlenieNastroekObrabotki: String
	PredstavlenieObieekta: String
	SsylkaObieekta: String
	TablichnaiaChastImia: String
	TablichnaiaChastPredstavlenie: String
	KhranilishcheVneshneiObrabotkiBase64Data: Binary
	DopolnitelnyeParametryObrabotkiType: String
	SsylkaObieektaType: String
	KhranilishcheVneshneiObrabotkiType: String
	DopolnitelnyeParametryObrabotki: Stream
	KhranilishcheVneshneiObrabotki: Stream
}
type CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnostRowType {
	Key: Guid!
	LineNumber: Int64!
	DopolnitelnyeParametryObrabotkiBase64Data: Binary
	PredstavlenieKnopki: String
	PredstavlenieNastroekObrabotki: String
	PredstavlenieObieekta: String
	SsylkaObieekta: String
	TablichnaiaChastImia: String
	TablichnaiaChastPredstavlenie: String
	KhranilishcheVneshneiObrabotkiBase64Data: Binary
	DopolnitelnyeParametryObrabotkiType: String
	SsylkaObieektaType: String
	KhranilishcheVneshneiObrabotkiType: String
	DopolnitelnyeParametryObrabotki: Stream
	KhranilishcheVneshneiObrabotki: Stream
}
input CatalogDopolnitelnyeVneshnieObrabotkiKomandyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Identifikator: String
	VariantZapuska: String
	Predstavlenie: String
	PokazyvatOpoveshchenie: Boolean
	Modifikator: String
	ReglamentnoeZadanieGUID: Guid
	Skryt: Boolean
	ZameniaemyeKomandy: String
}
type CatalogDopolnitelnyeVneshnieObrabotkiKomandyRowType {
	Key: Guid!
	LineNumber: Int64!
	Identifikator: String
	VariantZapuska: String
	Predstavlenie: String
	PokazyvatOpoveshchenie: Boolean
	Modifikator: String
	ReglamentnoeZadanieGUID: Guid
	Skryt: Boolean
	ZameniaemyeKomandy: String
}
input CatalogDopolnitelnyeVneshnieObrabotkiRazdelyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	RazdelKey: Guid
	UdalitImiaRazdela: String
}
type CatalogDopolnitelnyeVneshnieObrabotkiRazdelyRowType {
	Key: Guid!
	LineNumber: Int64!
	RazdelKey: Guid
	UdalitImiaRazdela: String
}
input CatalogDopolnitelnyeVneshnieObrabotkiNaznachenieRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ObieektNaznacheniiaKey: Guid
	UdalitPolnoeImiaObieektaMetadannykh: String
}
type CatalogDopolnitelnyeVneshnieObrabotkiNaznachenieRowType {
	Key: Guid!
	LineNumber: Int64!
	ObieektNaznacheniiaKey: Guid
	UdalitPolnoeImiaObieektaMetadannykh: String
}
input CatalogDopolnitelnyeVneshnieObrabotkiRazresheniiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidRazresheniia: String
	ParametryBase64Data: Binary
	ParametryType: String
	Parametry: Stream
}
type CatalogDopolnitelnyeVneshnieObrabotkiRazresheniiaRowType {
	Key: Guid!
	LineNumber: Int64!
	VidRazresheniia: String
	ParametryBase64Data: Binary
	ParametryType: String
	Parametry: Stream
}
input CatalogGruppyPolzovateleiPolzovateliGruppyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	PolzovatelKey: Guid
}
type CatalogGruppyPolzovateleiPolzovateliGruppyRowType {
	Key: Guid!
	LineNumber: Int64!
	PolzovatelKey: Guid
}
input DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DefektKey: Guid
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	DokumentPostupleniiaType: String
}
type DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DefektKey: Guid
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	DokumentPostupleniiaType: String
}
input DocumentZaiavkaNaPeremeshchenieTovarovTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StoimostBezNDS: Double
	StoimostSNDS: Double
	Sum: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaVRoznitseGr: Double
	TsenaPostupleniia: Double
}
type DocumentZaiavkaNaPeremeshchenieTovarovTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StoimostBezNDS: Double
	StoimostSNDS: Double
	Sum: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaVRoznitseGr: Double
	TsenaPostupleniia: Double
}
input DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoraRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	ItogovaiaStroka: Boolean
	KontragentKey: Guid
	RuchnaiaKorrektirovka: Boolean
	SummaOplaty: Double
	SummaOtgruzki: Double
	SummaVozvrata: Double
	DokumentUchetaType: String
}
type DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoraRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	ItogovaiaStroka: Boolean
	KontragentKey: Guid
	RuchnaiaKorrektirovka: Boolean
	SummaOplaty: Double
	SummaOtgruzki: Double
	SummaVozvrata: Double
	DokumentUchetaType: String
}
input CatalogUsloviiaOplatyTablitsaVyplatRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	PeriodOtsrochki: Int16
	ProtsentVyplaty: Double
}
type CatalogUsloviiaOplatyTablitsaVyplatRowType {
	Key: Guid!
	LineNumber: Int64!
	PeriodOtsrochki: Int16
	ProtsentVyplaty: Double
}
input DocumentUdalitNariadZakazIzdeliiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	Weight: Double
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TypeKey: Guid
	ProbeKey: Guid
	Kommentarii: String
	VesBezVstavok: Double
	NomerStrokiTCh: Int64
	InstanceKey: Guid
}
type DocumentUdalitNariadZakazIzdeliiaRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	Weight: Double
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TypeKey: Guid
	ProbeKey: Guid
	Kommentarii: String
	VesBezVstavok: Double
	NomerStrokiTCh: Int64
	InstanceKey: Guid
}
input DocumentUdalitNariadZakazUslugiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentUdalitNariadZakazUslugiRowType {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input DocumentUdalitNariadZakazSpetsifikatsiiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	NomerStrokiTCh: Int64
}
type DocumentUdalitNariadZakazSpetsifikatsiiaRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	NomerStrokiTCh: Int64
}
input DocumentUdalitNariadZakazMetallyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	ProbeKey: Guid
	Weight: String
}
type DocumentUdalitNariadZakazMetallyRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	ProbeKey: Guid
	Weight: String
}
input DocumentUdalitNariadZakazVstavkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
}
type DocumentUdalitNariadZakazVstavkiRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
}
input DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DataDolga: DateTime
	NovaiaDataDolga: DateTime
	NovaiaSummaDolga: Double
	NovaiaSummaDolgaUpr: Double
	SummaDolga: Double
	SummaDolgaUpr: Double
}
type DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostiRowType {
	Key: Guid!
	LineNumber: Int64!
	DataDolga: DateTime
	NovaiaDataDolga: DateTime
	NovaiaSummaDolga: Double
	NovaiaSummaDolgaUpr: Double
	SummaDolga: Double
	SummaDolgaUpr: Double
}
input DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentAkkreditivPoluchennyiRekvizityKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentAkkreditivPoluchennyiRekvizityKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentPriemIzRemontaIzdeliiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	Weight: Double
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TypeKey: Guid
	ProbeKey: Guid
	InstanceKey: Guid
}
type DocumentPriemIzRemontaIzdeliiaRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	Weight: Double
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TypeKey: Guid
	ProbeKey: Guid
	InstanceKey: Guid
}
input DocumentPriemIzRemontaMaterialyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchNomenklaturyKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	DokumentOprikhodovaniiaKey: Guid
	Weight: Double
	Quantity: Int64
}
type DocumentPriemIzRemontaMaterialyRowType {
	Key: Guid!
	LineNumber: Int64!
	KliuchNomenklaturyKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	DokumentOprikhodovaniiaKey: Guid
	Weight: Double
	Quantity: Int64
}
input DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
type DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvaRowType {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
input DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartiiKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDS: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaNDSPeredachi: Double
	SummaPeredachi: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPeredachi: Double
}
type DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartiiKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDS: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaNDSPeredachi: Double
	SummaPeredachi: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPeredachi: Double
}
input CatalogfmAnketaKlientaDannyeKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
type CatalogfmAnketaKlientaDannyeKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
input DocumentUstanovkaZnacheniiTochkiZakazaTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DataKon: DateTime
	DataNach: DateTime
	ZnachenieTochkiZakaza: Int64
	MinimalnyiStrakhovoiZapas: Int64
	ItemKey: Guid
	ProtsentnaiaStavkaZnacheniiaTochkiZakaza: Int16
	ProtsentnaiaStavkaMinimalnogoStrakhovogoZapasa: Int16
	SizeKey: Guid
	DepartmentKey: Guid
	SposobOpredeleniiaZnacheniiaTochkiZakaza: String
}
type DocumentUstanovkaZnacheniiTochkiZakazaTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DataKon: DateTime
	DataNach: DateTime
	ZnachenieTochkiZakaza: Int64
	MinimalnyiStrakhovoiZapas: Int64
	ItemKey: Guid
	ProtsentnaiaStavkaZnacheniiaTochkiZakaza: Int16
	ProtsentnaiaStavkaMinimalnogoStrakhovogoZapasa: Int16
	SizeKey: Guid
	DepartmentKey: Guid
	SposobOpredeleniiaZnacheniiaTochkiZakaza: String
}
input CatalogPravilaProdazhTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	KharakteristikiNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Cost: Double
	Quantity: Int64
	Weight: Double
}
type CatalogPravilaProdazhTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	KharakteristikiNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Cost: Double
	Quantity: Int64
	Weight: Double
}
input DocumentPostuplenieDopRaskhodovTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartii: String
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SummaRaspredeleniia: Double
	SummaTovara: Double
	KharakteristikaNomenklaturyKey: Guid
	DokumentPartiiType: String
}
type DocumentPostuplenieDopRaskhodovTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartii: String
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SummaRaspredeleniia: Double
	SummaTovara: Double
	KharakteristikaNomenklaturyKey: Guid
	DokumentPartiiType: String
}
input DocumentAvansovyiOtchetVydannyeAvansyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	RaskhodnyiKassovyiOrderKey: Guid
	Sum: Double
}
type DocumentAvansovyiOtchetVydannyeAvansyRowType {
	Key: Guid!
	LineNumber: Int64!
	RaskhodnyiKassovyiOrderKey: Guid
	Sum: Double
}
input DocumentAvansovyiOtchetTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DataSF: DateTime
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	NomerVkhodiashchegoDokumenta: String
	NomerSF: String
	SupplierKey: Guid
	PredieiavlenSF: Boolean
	ProektKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SchetFakturaKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
}
type DocumentAvansovyiOtchetTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DataSF: DateTime
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	NomerVkhodiashchegoDokumenta: String
	NomerSF: String
	SupplierKey: Guid
	PredieiavlenSF: Boolean
	ProektKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SchetFakturaKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
}
input DocumentAvansovyiOtchetOplataPostavshchikamRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NomerVkhodiashchegoDokumenta: String
	Sdelka: String
	Soderzhanie: String
	Sum: Double
	SummaVzaimoraschetov: Double
	SdelkaType: String
}
type DocumentAvansovyiOtchetOplataPostavshchikamRowType {
	Key: Guid!
	LineNumber: Int64!
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NomerVkhodiashchegoDokumenta: String
	Sdelka: String
	Soderzhanie: String
	Sum: Double
	SummaVzaimoraschetov: Double
	SdelkaType: String
}
input DocumentAvansovyiOtchetProcheeRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DataSF: DateTime
	Zakaz: String
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	NomerSF: String
	PodrazdelenieKey: Guid
	SupplierKey: Guid
	PredieiavlenSF: Boolean
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	SchetFakturaKey: Guid
	ZakazType: String
}
type DocumentAvansovyiOtchetProcheeRowType {
	Key: Guid!
	LineNumber: Int64!
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DataSF: DateTime
	Zakaz: String
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	NomerSF: String
	PodrazdelenieKey: Guid
	SupplierKey: Guid
	PredieiavlenSF: Boolean
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	SchetFakturaKey: Guid
	ZakazType: String
}
input CatalogDopolnitelnyePechatnyeFormyPrinadlezhnostRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	PredstavlenieObieekta: String
	SsylkaObieekta: String
	SsylkaObieektaType: String
}
type CatalogDopolnitelnyePechatnyeFormyPrinadlezhnostRowType {
	Key: Guid!
	LineNumber: Int64!
	PredstavlenieObieekta: String
	SsylkaObieekta: String
	SsylkaObieektaType: String
}
input CatalogObrabotkiObsluzhivaniiaTOModeliRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Model: String
}
type CatalogObrabotkiObsluzhivaniiaTOModeliRowType {
	Key: Guid!
	LineNumber: Int64!
	Model: String
}
input CatalogNastroikaIntervalovTablichnaiaChastRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KonetsIntervala: Int64
	NachaloIntervala: Int64
	Podpis: String
}
type CatalogNastroikaIntervalovTablichnaiaChastRowType {
	Key: Guid!
	LineNumber: Int64!
	KonetsIntervala: Int64
	NachaloIntervala: Int64
	Podpis: String
}
input CatalogProfiliGruppDostupaRoliRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	RolKey: Guid
}
type CatalogProfiliGruppDostupaRoliRowType {
	Key: Guid!
	LineNumber: Int64!
	RolKey: Guid
}
input CatalogProfiliGruppDostupaVidyDostupaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	Predustanovlennyi: Boolean
	VseRazresheny: Boolean
	VidDostupaType: String
}
type CatalogProfiliGruppDostupaVidyDostupaRowType {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	Predustanovlennyi: Boolean
	VseRazresheny: Boolean
	VidDostupaType: String
}
input CatalogProfiliGruppDostupaZnacheniiaDostupaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	ZnachenieDostupa: String
	VidDostupaType: String
	ZnachenieDostupaType: String
}
type CatalogProfiliGruppDostupaZnacheniiaDostupaRowType {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	ZnachenieDostupa: String
	VidDostupaType: String
	ZnachenieDostupaType: String
}
input CatalogProfiliGruppDostupaDostupPoPodsistemamRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ImiaPodsistemy: String
	ImiaObieekta: String
	Prosmotr: Boolean
	Redaktirovanie: Boolean
	Pechat: Boolean
	PokazVersii: Boolean
}
type CatalogProfiliGruppDostupaDostupPoPodsistemamRowType {
	Key: Guid!
	LineNumber: Int64!
	ImiaPodsistemy: String
	ImiaObieekta: String
	Prosmotr: Boolean
	Redaktirovanie: Boolean
	Pechat: Boolean
	PokazVersii: Boolean
}
input CatalogNastroikiDliaKureraSostavNaimenovaniiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SimvolyDo: String
	SimvolyPosle: String
	ElementNaimenovaniia: String
}
type CatalogNastroikiDliaKureraSostavNaimenovaniiaRowType {
	Key: Guid!
	LineNumber: Int64!
	SimvolyDo: String
	SimvolyPosle: String
	ElementNaimenovaniia: String
}
input DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	NomerPlatezha: Int16
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	NomerPlatezha: Int16
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentVozvratTovarovOtPokupateliaTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentOprikhodovaniia: String
	DokumentPartii: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	PercentManualDiscount: Double
	SizeKey: Guid
	Sebestoimost: Double
	SebestoimostBezNDS: Double
	SebestoimostUpr: Double
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	StatusPartii: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	DokumentOprikhodovaniiaType: String
	DokumentPartiiType: String
}
type DocumentVozvratTovarovOtPokupateliaTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentOprikhodovaniia: String
	DokumentPartii: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	PercentManualDiscount: Double
	SizeKey: Guid
	Sebestoimost: Double
	SebestoimostBezNDS: Double
	SebestoimostUpr: Double
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	StatusPartii: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	DokumentOprikhodovaniiaType: String
	DokumentPartiiType: String
}
input DocumentVozvratTovarovOtPokupateliaUslugiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	ProtsentOtSummuDokumenta: Double
	ProtsentOtSummyDokumenta: Double
	ProtsentSkidkiNatsenki: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentVozvratTovarovOtPokupateliaUslugiRowType {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	ProtsentOtSummuDokumenta: Double
	ProtsentOtSummyDokumenta: Double
	ProtsentSkidkiNatsenki: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input DocumentZakazPostavshchikuTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentZakazPostavshchikuTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input CatalogSkidkiNatsenkiUsloviiaPredostavleniiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	UsloviePredostavleniiaKey: Guid
}
type CatalogSkidkiNatsenkiUsloviiaPredostavleniiaRowType {
	Key: Guid!
	LineNumber: Int64!
	UsloviePredostavleniiaKey: Guid
}
input CatalogSkidkiNatsenkiTsenovyeGruppyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	TsenovaiaGruppaKey: Guid
	ZnachenieSkidkiNatsenki: Double
}
type CatalogSkidkiNatsenkiTsenovyeGruppyRowType {
	Key: Guid!
	LineNumber: Int64!
	TsenovaiaGruppaKey: Guid
	ZnachenieSkidkiNatsenki: Double
}
input CatalogSkidkiNatsenkiNaborPodarkovRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Quantity: Double
	SizeKey: Guid
	InstanceKey: Guid
}
type CatalogSkidkiNatsenkiNaborPodarkovRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Quantity: Double
	SizeKey: Guid
	InstanceKey: Guid
}
input CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanieRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	ProtsentTorgovoiUstupki: Double
}
type CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanieRowType {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	ProtsentTorgovoiUstupki: Double
}
input DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsenRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	TipTsenKey: Guid
}
type DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsenRowType {
	Key: Guid!
	LineNumber: Int64!
	TipTsenKey: Guid
}
input DocumentUstanovkaTsenNomenklaturyKontragentovTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ValiutaKey: Guid
	IndeksStrokiTablitsyTsen: Int64
	ItemKey: Guid
	SizeKey: Guid
	TipTsenKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
type DocumentUstanovkaTsenNomenklaturyKontragentovTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	ValiutaKey: Guid
	IndeksStrokiTablitsyTsen: Int64
	ItemKey: Guid
	SizeKey: Guid
	TipTsenKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
input DocumentProtsentPoterPoDavaltsamProtsentyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Protsent: Double
}
type DocumentProtsentPoterPoDavaltsamProtsentyRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Protsent: Double
}
input DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input CatalogVidyKodirovokiIzdeliiElementyKodirovkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Poriadok: Int64
	Poteri: Boolean
	Prais: Boolean
	ElementKey: Guid
}
type CatalogVidyKodirovokiIzdeliiElementyKodirovkiRowType {
	Key: Guid!
	LineNumber: Int64!
	Poriadok: Int64
	Poteri: Boolean
	Prais: Boolean
	ElementKey: Guid
}
input DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeliRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
	Vybran: Boolean
	DenNedeli: String
}
type DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeliRowType {
	Key: Guid!
	LineNumber: Int64!
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
	Vybran: Boolean
	DenNedeli: String
}
input DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
}
type DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartyRowType {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
}
input DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Poluchatel: String
	PoluchatelType: String
}
type DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkiRowType {
	Key: Guid!
	LineNumber: Int64!
	Poluchatel: String
	PoluchatelType: String
}
input DocumentUstanovkaSkidokNomenklaturyTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	OgranichenieSkidkiNatsenki: Double
	ProtsentSkidkiNatsenki: Double
}
type DocumentUstanovkaSkidokNomenklaturyTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	OgranichenieSkidkiNatsenki: Double
	ProtsentSkidkiNatsenki: Double
}
input CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DenNedeli: String
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
}
type CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviiaRowType {
	Key: Guid!
	LineNumber: Int64!
	DenNedeli: String
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
}
input CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateliRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Poluchatel: String
	PoluchatelType: String
}
type CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateliRowType {
	Key: Guid!
	LineNumber: Int64!
	Poluchatel: String
	PoluchatelType: String
}
input CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	TypeKey: Guid
	SupplierKey: Guid
	ProizvoditelKey: Guid
}
type CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkiRowType {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	TypeKey: Guid
	SupplierKey: Guid
	ProizvoditelKey: Guid
}
input DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhaRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentRaskhodnyiKassovyiOrderOplataRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
type DocumentRaskhodnyiKassovyiOrderOplataRowType {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
input DocumentRaskhodnyiKassovyiOrderTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SummaSkidki: Double
	VidTovaraKKT: String
	TipOplatyTovaraKKT: String
	SummaOsn: Double
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	SummaOpl: Double
}
type DocumentRaskhodnyiKassovyiOrderTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SummaSkidki: Double
	VidTovaraKKT: String
	TipOplatyTovaraKKT: String
	SummaOsn: Double
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	SummaOpl: Double
}
input DocumentSchetNaOplatuPostavshchikaTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
type DocumentSchetNaOplatuPostavshchikaTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
input DocumentSchetNaOplatuPostavshchikaUslugiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentSchetNaOplatuPostavshchikaUslugiRowType {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input DocumentReestrSpetssviazKlientyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	Adres: String
	Telefon: String
	Weight: Double
	Sum: Double
	Paket: String
	SummaPropisiu: String
	GabarityKey: Guid
}
type DocumentReestrSpetssviazKlientyRowType {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	Adres: String
	Telefon: String
	Weight: Double
	Sum: Double
	Paket: String
	SummaPropisiu: String
	GabarityKey: Guid
}
input DocumentVvodNachalnykhOstatkovVzaimoraschetyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	Sum: Double
}
type DocumentVvodNachalnykhOstatkovVzaimoraschetyRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	Sum: Double
}
input DocumentVvodNachalnykhOstatkovTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentOprikhodovaniiaKey: Guid
	DokumentProdazhiKey: Guid
	KachestvoKey: Guid
	Quantity: Double
	Comment: String
	NDSVkliuchenVStoimost: Boolean
	ItemKey: Guid
	NomerGTDKey: Guid
	Pasport: String
	ProtsentRoznichnoiNatsenki: Double
	ProtsentSkidkiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	StavkaNDSProdazhi: String
	StatusPartii: String
	StranaProiskhozhdeniiaKey: Guid
	Sum: Double
	SummaNDS: Double
	SummaNDSProdazhi: Double
	SummaProdazhi: Double
	SummaProdazhiBezSkidok: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaProdazhi: Double
	StatusRaskhoda: String
}
type DocumentVvodNachalnykhOstatkovTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentOprikhodovaniiaKey: Guid
	DokumentProdazhiKey: Guid
	KachestvoKey: Guid
	Quantity: Double
	Comment: String
	NDSVkliuchenVStoimost: Boolean
	ItemKey: Guid
	NomerGTDKey: Guid
	Pasport: String
	ProtsentRoznichnoiNatsenki: Double
	ProtsentSkidkiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	StavkaNDSProdazhi: String
	StatusPartii: String
	StranaProiskhozhdeniiaKey: Guid
	Sum: Double
	SummaNDS: Double
	SummaNDSProdazhi: Double
	SummaProdazhi: Double
	SummaProdazhiBezSkidok: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaProdazhi: Double
	StatusRaskhoda: String
}
input DocumentOprikhodovanieTovarovTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaRegl: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
}
type DocumentOprikhodovanieTovarovTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaRegl: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
}
input DocumentOprikhodovanieTovarovSertifikatyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
type DocumentOprikhodovanieTovarovSertifikatyRowType {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
input DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SummaStaraia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaZaGramm: Double
}
type DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SummaStaraia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaZaGramm: Double
}
input DocumentElektronnoePismoKomuTChRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
type DocumentElektronnoePismoKomuTChRowType {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
input DocumentElektronnoePismoKopiiTChRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
type DocumentElektronnoePismoKopiiTChRowType {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
input DocumentElektronnoePismoSkrytyeKopiiTChRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
type DocumentElektronnoePismoSkrytyeKopiiTChRowType {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
input CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
type CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentaRowType {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
input CatalogPravilaTsenoobrazovaniiaTsenovyeGruppyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	TsenovaiaGruppaKey: Guid
	VidTsenKey: Guid
}
type CatalogPravilaTsenoobrazovaniiaTsenovyeGruppyRowType {
	Key: Guid!
	LineNumber: Int64!
	TsenovaiaGruppaKey: Guid
	VidTsenKey: Guid
}
input DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaIzmereniiaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	Koef: Double
	Koeffitsient: Double
	ItemKey: Guid
	Pasport: String
	ProektKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
}
type DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaIzmereniiaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	Koef: Double
	Koeffitsient: Double
	ItemKey: Guid
	Pasport: String
	ProektKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
}
input DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugiRowType {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input CatalogGruppyDostupaPolzovateliRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Polzovatel: String
	PolzovatelType: String
}
type CatalogGruppyDostupaPolzovateliRowType {
	Key: Guid!
	LineNumber: Int64!
	Polzovatel: String
	PolzovatelType: String
}
input CatalogGruppyDostupaVidyDostupaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	VseRazresheny: Boolean
	VidDostupaType: String
}
type CatalogGruppyDostupaVidyDostupaRowType {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	VseRazresheny: Boolean
	VidDostupaType: String
}
input CatalogGruppyDostupaZnacheniiaDostupaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	ZnachenieDostupa: String
	VidDostupaType: String
	ZnachenieDostupaType: String
}
type CatalogGruppyDostupaZnacheniiaDostupaRowType {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	ZnachenieDostupa: String
	VidDostupaType: String
	ZnachenieDostupaType: String
}
input CatalogGruppyDostupaDostupPoPodsistemamRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ImiaPodsistemy: String
	ImiaObieekta: String
	Prosmotr: Boolean
	Redaktirovanie: Boolean
	Pechat: Boolean
	PokazVersii: Boolean
}
type CatalogGruppyDostupaDostupPoPodsistemamRowType {
	Key: Guid!
	LineNumber: Int64!
	ImiaPodsistemy: String
	ImiaObieekta: String
	Prosmotr: Boolean
	Redaktirovanie: Boolean
	Pechat: Boolean
	PokazVersii: Boolean
}
input DocumentReestrSchetovReestrRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidTransporta: String
	DataOtgruzki: DateTime
	NomerDokumenta: String
	Sum: Double
}
type DocumentReestrSchetovReestrRowType {
	Key: Guid!
	LineNumber: Int64!
	VidTransporta: String
	DataOtgruzki: DateTime
	NomerDokumenta: String
	Sum: Double
}
input DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesUchet: Double
	Quantity: Int64
	KolichestvoUchet: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
type DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesUchet: Double
	Quantity: Int64
	KolichestvoUchet: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
input CatalogNastroikiRMKPoriadokPrimeneniiaSkidokRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	UslovieSkidki: String
}
type CatalogNastroikiRMKPoriadokPrimeneniiaSkidokRowType {
	Key: Guid!
	LineNumber: Int64!
	UslovieSkidki: String
}
input CatalogNastroikiRMKSostavNaimenovaniiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SimvolyDo: String
	SimvolyPosle: String
	ElementNaimenovaniia: String
}
type CatalogNastroikiRMKSostavNaimenovaniiaRowType {
	Key: Guid!
	LineNumber: Int64!
	SimvolyDo: String
	SimvolyPosle: String
	ElementNaimenovaniia: String
}
input CatalogKharakteristikiNomenklaturySpetsifikatsiiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
}
type CatalogKharakteristikiNomenklaturySpetsifikatsiiaRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
}
input DocumentOtborTovarovTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumManualDiscount: Double
	SumAutoDiscount: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentOtborTovarovTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumManualDiscount: Double
	SumAutoDiscount: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input DocumentOtborTovarovTovaryKOtboruRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
}
type DocumentOtborTovarovTovaryKOtboruRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
}
input CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	RelizIuvelirnogoSalonaKey: Guid
}
type CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizyRowType {
	Key: Guid!
	LineNumber: Int64!
	RelizIuvelirnogoSalonaKey: Guid
}
input DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
type DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvaRowType {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
input DocumentOtchetKomissioneraOProdazhakhTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDS: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaNDSPeredachi: Double
	SummaPeredachi: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPeredachi: Double
	DataSchetaFakturyKomissionera: DateTime
	PokupatelKey: Guid
	NomerSchetaFakturyKomissionera: String
}
type DocumentOtchetKomissioneraOProdazhakhTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDS: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaNDSPeredachi: Double
	SummaPeredachi: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPeredachi: Double
	DataSchetaFakturyKomissionera: DateTime
	PokupatelKey: Guid
	NomerSchetaFakturyKomissionera: String
}
input CatalogFiltryDliaElektronnykhPisemDeistviiaFiltraRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	GruppaPisemKey: Guid
	DeistvieFiltra: String
}
type CatalogFiltryDliaElektronnykhPisemDeistviiaFiltraRowType {
	Key: Guid!
	LineNumber: Int64!
	GruppaPisemKey: Guid
	DeistvieFiltra: String
}
input CatalogFiltryDliaElektronnykhPisemUsloviiaFiltraRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	ZnachenieUsloviia: String
	OtritsanieUsloviia: Boolean
	Uslovie: String
}
type CatalogFiltryDliaElektronnykhPisemUsloviiaFiltraRowType {
	Key: Guid!
	LineNumber: Int64!
	ZnachenieUsloviia: String
	OtritsanieUsloviia: Boolean
	Uslovie: String
}
input DocumentPreiskurantTsenNaTsvKamniTablitsaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Razmer1Do: Double
	Razmer1Ot: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	Cost: Int64
}
type DocumentPreiskurantTsenNaTsvKamniTablitsaRowType {
	Key: Guid!
	LineNumber: Int64!
	Razmer1Do: Double
	Razmer1Ot: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	Cost: Int64
}
input DocumentTelemarketingUchastnikiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	NaimenovaniePolnoe: String
	KontaktnoeLitsoKey: Guid
	Telefon: String
	RezultatObrabotkiZvonka: String
	EstInteres: Boolean
	SobytieKey: Guid
	OprosKey: Guid
	Opisanie: String
}
type DocumentTelemarketingUchastnikiRowType {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	NaimenovaniePolnoe: String
	KontaktnoeLitsoKey: Guid
	Telefon: String
	RezultatObrabotkiZvonka: String
	EstInteres: Boolean
	SobytieKey: Guid
	OprosKey: Guid
	Opisanie: String
}
input DocumentRassylkaAnketVlozheniiaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VlozhenieBase64Data: Binary
	ImiaFaila: String
	VlozhenieType: String
	Vlozhenie: Stream
}
type DocumentRassylkaAnketVlozheniiaRowType {
	Key: Guid!
	LineNumber: Int64!
	VlozhenieBase64Data: Binary
	ImiaFaila: String
	VlozhenieType: String
	Vlozhenie: Stream
}
input DocumentRassylkaAnketPoluchateliRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Obieekt: String
	Poluchatel: String
	ObieektType: String
}
type DocumentRassylkaAnketPoluchateliRowType {
	Key: Guid!
	LineNumber: Int64!
	Obieekt: String
	Poluchatel: String
	ObieektType: String
}
input CatalogSkhemyRealizatsiiEtapySkhemyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorPokupkiKey: Guid
	DogovorProdazhiKey: Guid
	ZnachenieNatsenki: String
	KontragentPokupatelKey: Guid
	KontragentProdavetsKey: Guid
	OkrugliatVBolshuiuStoronu: Boolean
	OrganizatsiiaPokupatelKey: Guid
	OrganizatsiiaProdavetsKey: Guid
	PoriadokOkrugleniia: String
	SposobNatsenki: Int16
	TipNatsenki: Int16
	ZnachenieNatsenkiType: String
}
type CatalogSkhemyRealizatsiiEtapySkhemyRowType {
	Key: Guid!
	LineNumber: Int64!
	DogovorPokupkiKey: Guid
	DogovorProdazhiKey: Guid
	ZnachenieNatsenki: String
	KontragentPokupatelKey: Guid
	KontragentProdavetsKey: Guid
	OkrugliatVBolshuiuStoronu: Boolean
	OrganizatsiiaPokupatelKey: Guid
	OrganizatsiiaProdavetsKey: Guid
	PoriadokOkrugleniia: String
	SposobNatsenki: Int16
	TipNatsenki: Int16
	ZnachenieNatsenkiType: String
}
input DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	XYZKlassifikatsiia: String
	XYZKlassifikatsiiaStaraia: String
	ZnachenieParametra: Double
	IndeksSortirovki: Int16
	KolichestvoDokumentov: Int64
	KontragentKey: Guid
	KoeffitsientVariatsii: Double
	MenedzherKontragentaKey: Guid
	StadiiaVzaimootnoshenii: String
	StadiiaVzaimootnosheniiStaraia: String
}
type DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovRowType {
	Key: Guid!
	LineNumber: Int64!
	XYZKlassifikatsiia: String
	XYZKlassifikatsiiaStaraia: String
	ZnachenieParametra: Double
	IndeksSortirovki: Int16
	KolichestvoDokumentov: Int64
	KontragentKey: Guid
	KoeffitsientVariatsii: Double
	MenedzherKontragentaKey: Guid
	StadiiaVzaimootnoshenii: String
	StadiiaVzaimootnosheniiStaraia: String
}
input DocumentZakazKlientaTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	Razmestit: Int64
	Rezervirovat: Int64
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	Cost: Double
	InstanceKey: Guid
	SumAutoDiscount: Double
	SumManualDiscount: Double
	KharakteristikaNomenklaturyKey: Guid
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentZakazKlientaTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	Razmestit: Int64
	Rezervirovat: Int64
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	Cost: Double
	InstanceKey: Guid
	SumAutoDiscount: Double
	SumManualDiscount: Double
	KharakteristikaNomenklaturyKey: Guid
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input DocumentPostuplenieProduktsiiIzProizvodstvaTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesPoter: Double
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	ProtsentPoter: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	StoimostVstavok: Double
	StoimostMetalla: Double
	StoimostRabot: Double
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	CostOfWork: Double
	SummaNDSVstavok: Double
	SummaNDSMetalla: Double
	SummaNDSRabot: Double
	VesVstavok: Double
}
type DocumentPostuplenieProduktsiiIzProizvodstvaTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesPoter: Double
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	ProtsentPoter: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	StoimostVstavok: Double
	StoimostMetalla: Double
	StoimostRabot: Double
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	CostOfWork: Double
	SummaNDSVstavok: Double
	SummaNDSMetalla: Double
	SummaNDSRabot: Double
	VesVstavok: Double
}
input DocumentPostuplenieProduktsiiIzProizvodstvaMaterialyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	Nomenklatura: String
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	ItemType: String
}
type DocumentPostuplenieProduktsiiIzProizvodstvaMaterialyRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	Nomenklatura: String
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	ItemType: String
}
input DocumentPostuplenieTovarovUslugTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaIzmereniiaKey: Guid
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	Koef: Double
	Koeffitsient: Double
	NaborKey: Guid
	ItemKey: Guid
	NomerGTDKey: Guid
	NomerNabora: Int64
	Pasport: String
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	StranaProiskhozhdeniiaKey: Guid
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
type DocumentPostuplenieTovarovUslugTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaIzmereniiaKey: Guid
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	Koef: Double
	Koeffitsient: Double
	NaborKey: Guid
	ItemKey: Guid
	NomerGTDKey: Guid
	NomerNabora: Int64
	Pasport: String
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	StranaProiskhozhdeniiaKey: Guid
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
input DocumentPostuplenieTovarovUslugUslugiRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentPostuplenieTovarovUslugUslugiRowType {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input DocumentPlanProdazhPoSalonamSalonyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
	SummaPlana: Double
	Primechanie: String
	IndeksStrokiIzTablitsyDnei: Int64
	SummaPlanaDnevnaia: Double
	DenPoGrafiku: DateTime
	SummaPlanaFakt: Double
	PlanEst: Boolean
	KU: Double
	FormulaDliaRasschetaKey: Guid
	RasshifrovkaFormuly: String
}
type DocumentPlanProdazhPoSalonamSalonyRowType {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
	SummaPlana: Double
	Primechanie: String
	IndeksStrokiIzTablitsyDnei: Int64
	SummaPlanaDnevnaia: Double
	DenPoGrafiku: DateTime
	SummaPlanaFakt: Double
	PlanEst: Boolean
	KU: Double
	FormulaDliaRasschetaKey: Guid
	RasshifrovkaFormuly: String
}
input DocumentPlanProdazhPoSalonamDniPoGrafikuRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DenPoGrafiku: Int16
}
type DocumentPlanProdazhPoSalonamDniPoGrafikuRowType {
	Key: Guid!
	LineNumber: Int64!
	DenPoGrafiku: Int16
}
input DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvaRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
type DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvaRowType {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
input DocumentStornirovanieOtchetaKomitentuOProdazhakhTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	OtchetKomitentuKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPostupleniia: Double
	DokumentPostupleniiaType: String
}
type DocumentStornirovanieOtchetaKomitentuOProdazhakhTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	OtchetKomitentuKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPostupleniia: Double
	DokumentPostupleniiaType: String
}
input DocumentPeredachaVRemontIzdeliiaPriniatyeVRemontRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchNomenklaturyKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	DokumentOprikhodovaniiaKey: Guid
	Weight: Double
	Quantity: Int64
}
type DocumentPeredachaVRemontIzdeliiaPriniatyeVRemontRowType {
	Key: Guid!
	LineNumber: Int64!
	KliuchNomenklaturyKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	DokumentOprikhodovaniiaKey: Guid
	Weight: Double
	Quantity: Int64
}
input DocumentPeredachaVRemontTovaryRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
type DocumentPeredachaVRemontTovaryRowType {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
input CatalogGruppySkladovSkladyRowTypeInput {
	Key: Guid!
	LineNumber: Int64!
	DepartmentKey: Guid
}
type CatalogGruppySkladovSkladyRowType {
	Key: Guid!
	LineNumber: Int64!
	DepartmentKey: Guid
}
input AccumulationRegisterPartiiTovarovVProizvodstveInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterPartiiTovarovVProizvodstveRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterPartiiTovarovVProizvodstve {
	Recorder: String!
	RecordSet: [AccumulationRegisterPartiiTovarovVProizvodstveRowType!]!
	RecorderType: String!
}
input AccumulationRegisterPartiiTovarovVProizvodstveRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ProizvodstvennyiUchastokKey: Guid
	Nomenklatura: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	OperationCode: String
	SpisaniePartii: Boolean
	NomerKorStroki: Int64
	RecorderType: String!
	ItemType: String
	DokumentOprikhodovaniiaType: String
}
type AccumulationRegisterPartiiTovarovVProizvodstveRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ProizvodstvennyiUchastokKey: Guid
	Nomenklatura: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	OperationCode: String
	SpisaniePartii: Boolean
	NomerKorStroki: Int64
	RecorderType: String!
	ItemType: String
	DokumentOprikhodovaniiaType: String
}
input AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterVzaimoraschetySPodotchetnymiLitsami {
	Recorder: String!
	RecordSet: [AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRowType!]!
	RecorderType: String!
}
input AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	FizLitsoKey: Guid
	ValiutaKey: Guid
	RaschetnyiDokument: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	RaschetnyiDokumentType: String
}
type AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	FizLitsoKey: Guid
	ValiutaKey: Guid
	RaschetnyiDokument: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	RaschetnyiDokumentType: String
}
input AccumulationRegisterVnutrennieZakazyInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterVnutrennieZakazyRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterVnutrennieZakazy {
	Recorder: String!
	RecordSet: [AccumulationRegisterVnutrennieZakazyRowType!]!
	RecorderType: String!
}
input AccumulationRegisterVnutrennieZakazyRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Zakazchik: String
	VnutrenniiZakazKey: Guid
	StatusPartii: String
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	ZakazchikType: String
}
type AccumulationRegisterVnutrennieZakazyRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Zakazchik: String
	VnutrenniiZakazKey: Guid
	StatusPartii: String
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	ZakazchikType: String
}
input AccumulationRegisterDenezhnyeSredstvaKomitentaInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaKomitentaRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterDenezhnyeSredstvaKomitenta {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaKomitentaRowType!]!
	RecorderType: String!
}
input AccumulationRegisterDenezhnyeSredstvaKomitentaRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKomitentaRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
input AccumulationRegisterZakazyKlientovInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterZakazyKlientovRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterZakazyKlientov {
	Recorder: String!
	RecordSet: [AccumulationRegisterZakazyKlientovRowType!]!
	RecorderType: String!
}
input AccumulationRegisterZakazyKlientovRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	SizeKey: Guid
	ZakazKlientaKey: Guid
	InstanceKey: Guid
	Zakazano: Int64
	VRezerve: Int64
	KRazmeshcheniiu: Int64
	ZakazanoVes: Double
	ZakazanoSumma: Double
	RecorderType: String!
}
type AccumulationRegisterZakazyKlientovRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	SizeKey: Guid
	ZakazKlientaKey: Guid
	InstanceKey: Guid
	Zakazano: Int64
	VRezerve: Int64
	KRazmeshcheniiu: Int64
	ZakazanoVes: Double
	ZakazanoSumma: Double
	RecorderType: String!
}
input AccumulationRegisterSummyPoFinmonitoringuRoznitsaInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterSummyPoFinmonitoringuRoznitsaRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterSummyPoFinmonitoringuRoznitsa {
	Recorder: String!
	RecordSet: [AccumulationRegisterSummyPoFinmonitoringuRoznitsaRowType!]!
	RecorderType: String!
}
input AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	SummaPokupok: Double
	SummaSkupki: Double
	RecorderType: String!
}
type AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	SummaPokupok: Double
	SummaSkupki: Double
	RecorderType: String!
}
input AccumulationRegisterDenezhnyeSredstvaKPolucheniiuInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterDenezhnyeSredstvaKPolucheniiu {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRowType!]!
	RecorderType: String!
}
input AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentPolucheniia: String
	TypeOfMovingMoneyKey: Guid
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentPolucheniiaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentPolucheniia: String
	TypeOfMovingMoneyKey: Guid
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentPolucheniiaType: String
}
input AccumulationRegisterProdazhiPoDiskontnymKartamInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterProdazhiPoDiskontnymKartamRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterProdazhiPoDiskontnymKartam {
	Recorder: String!
	RecordSet: [AccumulationRegisterProdazhiPoDiskontnymKartamRowType!]!
	RecorderType: String!
}
input AccumulationRegisterProdazhiPoDiskontnymKartamRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	MemberCardKey: Guid
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DokumentProdazhi: String
	Sum: Double
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	DokumentProdazhiType: String
}
type AccumulationRegisterProdazhiPoDiskontnymKartamRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	MemberCardKey: Guid
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DokumentProdazhi: String
	Sum: Double
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	DokumentProdazhiType: String
}
input AccumulationRegisterTovaryPoluchennyeInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryPoluchennyeRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterTovaryPoluchennye {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryPoluchennyeRowType!]!
	RecorderType: String!
}
input AccumulationRegisterTovaryPoluchennyeRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Quantity: Int64
	Weight: Double
	SummaVzaimoraschetov: Double
	RecorderType: String!
	SdelkaType: String
}
type AccumulationRegisterTovaryPoluchennyeRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Quantity: Int64
	Weight: Double
	SummaVzaimoraschetov: Double
	RecorderType: String!
	SdelkaType: String
}
input AccumulationRegisterSvobodnyeOstatkiInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterSvobodnyeOstatkiRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterSvobodnyeOstatki {
	Recorder: String!
	RecordSet: [AccumulationRegisterSvobodnyeOstatkiRowType!]!
	RecorderType: String!
}
input AccumulationRegisterSvobodnyeOstatkiRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	VNalichii: Int64
	VRezerve: Int64
	RecorderType: String!
}
type AccumulationRegisterSvobodnyeOstatkiRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	VNalichii: Int64
	VRezerve: Int64
	RecorderType: String!
}
input AccumulationRegisterSummyVRassrochkuInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterSummyVRassrochkuRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterSummyVRassrochku {
	Recorder: String!
	RecordSet: [AccumulationRegisterSummyVRassrochkuRowType!]!
	RecorderType: String!
}
input AccumulationRegisterSummyVRassrochkuRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorRassrochkiKey: Guid
	SostavStrokiChekaKey: Guid
	Sum: Double
	SummaNDS: Double
	RecorderType: String!
}
type AccumulationRegisterSummyVRassrochkuRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorRassrochkiKey: Guid
	SostavStrokiChekaKey: Guid
	Sum: Double
	SummaNDS: Double
	RecorderType: String!
}
input AccumulationRegisterGrafikPlatezheiInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterGrafikPlatezheiRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterGrafikPlatezhei {
	Recorder: String!
	RecordSet: [AccumulationRegisterGrafikPlatezheiRowType!]!
	RecorderType: String!
}
input AccumulationRegisterGrafikPlatezheiRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	VidDogovoraKontragenta: String
	DataDolga: DateTime
	Oplacheno: Double
	NachislenDolg: Double
	Dolg: Double
	Avans: Double
	OplachenoUpr: Int64
	NachislenDolgUpr: Double
	DolgUpr: Double
	AvansUpr: Double
	RecorderType: String!
}
type AccumulationRegisterGrafikPlatezheiRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	VidDogovoraKontragenta: String
	DataDolga: DateTime
	Oplacheno: Double
	NachislenDolg: Double
	Dolg: Double
	Avans: Double
	OplachenoUpr: Int64
	NachislenDolgUpr: Double
	DolgUpr: Double
	AvansUpr: Double
	RecorderType: String!
}
input AccumulationRegisterRoznichnaiaVyruchkaInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterRoznichnaiaVyruchkaRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterRoznichnaiaVyruchka {
	Recorder: String!
	RecordSet: [AccumulationRegisterRoznichnaiaVyruchkaRowType!]!
	RecorderType: String!
}
input AccumulationRegisterRoznichnaiaVyruchkaRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	RoznichnaiaTochka: String
	Sum: Double
	PodrazdelenieKey: Guid
	RecorderType: String!
	RoznichnaiaTochkaType: String
}
type AccumulationRegisterRoznichnaiaVyruchkaRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	RoznichnaiaTochka: String
	Sum: Double
	PodrazdelenieKey: Guid
	RecorderType: String!
	RoznichnaiaTochkaType: String
}
input AccumulationRegisterTovaryVPutiInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryVPutiRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterTovaryVPuti {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryVPutiRowType!]!
	RecorderType: String!
}
input AccumulationRegisterTovaryVPutiRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	DokumentOprikhodovaniiaType: String
}
type AccumulationRegisterTovaryVPutiRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	Quantity: Int64
	Weight: Double
	RecorderType: String!
	DokumentOprikhodovaniiaType: String
}
input AccumulationRegisterPoteriMetallaVProizvodstveInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterPoteriMetallaVProizvodstveRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterPoteriMetallaVProizvodstve {
	Recorder: String!
	RecordSet: [AccumulationRegisterPoteriMetallaVProizvodstveRowType!]!
	RecorderType: String!
}
input AccumulationRegisterPoteriMetallaVProizvodstveRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	Nomenklatura: String
	DokumentOprikhodovaniia: String
	VesPoter: Double
	RecorderType: String!
	ItemType: String
	DokumentOprikhodovaniiaType: String
}
type AccumulationRegisterPoteriMetallaVProizvodstveRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	Nomenklatura: String
	DokumentOprikhodovaniia: String
	VesPoter: Double
	RecorderType: String!
	ItemType: String
	DokumentOprikhodovaniiaType: String
}
input AccumulationRegisterPartiiTovarovNaSkladakhInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterPartiiTovarovNaSkladakhRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterPartiiTovarovNaSkladakh {
	Recorder: String!
	RecordSet: [AccumulationRegisterPartiiTovarovNaSkladakhRowType!]!
	RecorderType: String!
}
input ProductActionDocumentInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentOprikhodovaniia: String
	StatusPartii: String
	KachestvoKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	OperationCode: String
	SpisaniePartii: Boolean
	NomerKorStroki: Int64
	RecorderType: String!
	DokumentOprikhodovaniiaType: String
}
type ProductActionDocument {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentOprikhodovaniia: String
	StatusPartii: String
	KachestvoKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	OperationCode: String
	SpisaniePartii: Boolean
	NomerKorStroki: Int64
	RecorderType: String!
	DokumentOprikhodovaniiaType: String
}
input AccumulationRegisterSummyDokumentovDliaObmenaInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterSummyDokumentovDliaObmenaRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterSummyDokumentovDliaObmena {
	Recorder: String!
	RecordSet: [AccumulationRegisterSummyDokumentovDliaObmenaRowType!]!
	RecorderType: String!
}
input AccumulationRegisterSummyDokumentovDliaObmenaRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentKey: Guid
	Sum: Double
	RecorderType: String!
}
type AccumulationRegisterSummyDokumentovDliaObmenaRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DokumentKey: Guid
	Sum: Double
	RecorderType: String!
}
input AccumulationRegisterDvizheniiaDenezhnykhSredstvInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterDvizheniiaDenezhnykhSredstvRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterDvizheniiaDenezhnykhSredstv {
	Recorder: String!
	RecordSet: [AccumulationRegisterDvizheniiaDenezhnykhSredstvRowType!]!
	RecorderType: String!
}
input AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	VidDenezhnykhSredstv: String
	PrikhodRaskhod: String
	BankovskiiSchetKassa: String
	TypeOfMovingMoneyKey: Guid
	DokumentDvizheniia: String
	Kontragent: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	ProektKey: Guid
	DokumentPlanirovaniiaPlatezha: String
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentDvizheniiaType: String
	KontragentType: String
	SdelkaType: String
	DokumentPlanirovaniiaPlatezhaType: String
}
type AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	VidDenezhnykhSredstv: String
	PrikhodRaskhod: String
	BankovskiiSchetKassa: String
	TypeOfMovingMoneyKey: Guid
	DokumentDvizheniia: String
	Kontragent: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	ProektKey: Guid
	DokumentPlanirovaniiaPlatezha: String
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentDvizheniiaType: String
	KontragentType: String
	SdelkaType: String
	DokumentPlanirovaniiaPlatezhaType: String
}
input AccumulationRegisterProdazhiPoStatiamInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterProdazhiPoStatiamRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterProdazhiPoStatiam {
	Recorder: String!
	RecordSet: [AccumulationRegisterProdazhiPoStatiamRowType!]!
	RecorderType: String!
}
input AccumulationRegisterProdazhiPoStatiamRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	StatiaKey: Guid
	SummaProdazha: Double
	SummaVozvrat: Double
	RecorderType: String!
}
type AccumulationRegisterProdazhiPoStatiamRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	StatiaKey: Guid
	SummaProdazha: Double
	SummaVozvrat: Double
	RecorderType: String!
}
input InformationRegisterTsenyNomenklaturyInput {
	Recorder: String!
	RecordSet: [InformationRegisterTsenyNomenklaturyRowTypeInput!]!
	RecorderType: String!
}
type InformationRegisterTsenyNomenklatury {
	Recorder: String!
	RecordSet: [InformationRegisterTsenyNomenklaturyRowType!]!
	RecorderType: String!
}
input InformationRegisterTsenyNomenklaturyRecordTypeInput {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	TipTsenKey: Guid!
	SegmentNomenklaturyKey: Guid!
	ItemKey: Guid!
	InstanceKey: Guid!
	KharakteristikaNomenklaturyKey: Guid!
	SizeKey: Guid!
	Cost: Double
	ProtsentSkidkiNatsenki: Double
	ValiutaKey: Guid
	EdinitsaIzmereniiaKey: Guid
	RecorderType: String
}
type InformationRegisterTsenyNomenklaturyRecordType {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	TipTsenKey: Guid!
	SegmentNomenklaturyKey: Guid!
	ItemKey: Guid!
	InstanceKey: Guid!
	KharakteristikaNomenklaturyKey: Guid!
	SizeKey: Guid!
	Cost: Double
	ProtsentSkidkiNatsenki: Double
	ValiutaKey: Guid
	EdinitsaIzmereniiaKey: Guid
	RecorderType: String
}
input AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse {
	Recorder: String!
	RecordSet: [AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRowType!]!
	RecorderType: String!
}
input AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	ProdazhnaiaStoimost: Double
	VsegoSkidki: Double
	SkidkiPoDiskontnymKartam: Double
	SummaOplatyKartami: Double
	SummaOplatyBankovskimKreditom: Double
	SummaVozvrata: Double
	VesVChekakh: Double
	KolichestvoChekov: Int64
	SummaProdazhiSertifikatov: Double
	SummaOplatySertifikatami: Double
	PogashenoSertifikatami: Double
	SummaOplatyBonusom: Double
	VesObmena: Double
	SummaObmena: Double
	VesSkuplennogoTovara: Double
	VydanoZaSkuplennyiTovar: Double
	KolichestvoIzdelii: Int64
	SumManualDiscount: Double
	SumAutoDiscount: Double
	SummaRassrochki: Double
	SummaPogasheniiaRassrochki: Double
	SummaOplatyNalichnymi: Double
	RecorderType: String!
}
type AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	ProdazhnaiaStoimost: Double
	VsegoSkidki: Double
	SkidkiPoDiskontnymKartam: Double
	SummaOplatyKartami: Double
	SummaOplatyBankovskimKreditom: Double
	SummaVozvrata: Double
	VesVChekakh: Double
	KolichestvoChekov: Int64
	SummaProdazhiSertifikatov: Double
	SummaOplatySertifikatami: Double
	PogashenoSertifikatami: Double
	SummaOplatyBonusom: Double
	VesObmena: Double
	SummaObmena: Double
	VesSkuplennogoTovara: Double
	VydanoZaSkuplennyiTovar: Double
	KolichestvoIzdelii: Int64
	SumManualDiscount: Double
	SumAutoDiscount: Double
	SummaRassrochki: Double
	SummaPogasheniiaRassrochki: Double
	SummaOplatyNalichnymi: Double
	RecorderType: String!
}
input AccumulationRegisterDenezhnyeSredstvaVRezerveInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaVRezerveRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterDenezhnyeSredstvaVRezerve {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaVRezerveRowType!]!
	RecorderType: String!
}
input AccumulationRegisterDenezhnyeSredstvaVRezerveRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentRezervirovaniiaKey: Guid
	Sum: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
}
type AccumulationRegisterDenezhnyeSredstvaVRezerveRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentRezervirovaniiaKey: Guid
	Sum: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
}
input AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRowType!]!
	RecorderType: String!
}
input AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	RetailCost: Double
	Quantity: Int64
	Weight: Double
	RecorderType: String!
}
type AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DepartmentKey: Guid
	RetailCost: Double
	Quantity: Int64
	Weight: Double
	RecorderType: String!
}
input AccumulationRegisterDavalcheskiiMetallPoteriInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterDavalcheskiiMetallPoteriRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterDavalcheskiiMetallPoteri {
	Recorder: String!
	RecordSet: [AccumulationRegisterDavalcheskiiMetallPoteriRowType!]!
	RecorderType: String!
}
input AccumulationRegisterDavalcheskiiMetallPoteriRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Weight: Double
	Protsent: Double
	RecorderType: String!
}
type AccumulationRegisterDavalcheskiiMetallPoteriRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Weight: Double
	Protsent: Double
	RecorderType: String!
}
input InformationRegisterTsenyPoPreiskurantuInput {
	Recorder: String!
	RecordSet: [InformationRegisterTsenyPoPreiskurantuRowTypeInput!]!
	RecorderType: String!
}
type InformationRegisterTsenyPoPreiskurantu {
	Recorder: String!
	RecordSet: [InformationRegisterTsenyPoPreiskurantuRowType!]!
	RecorderType: String!
}
input InformationRegisterTsenyPoPreiskurantuRecordTypeInput {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	KamenKey: Guid!
	FormaOgrankiKey: Guid!
	TsvetKamniaKey: Guid!
	GruppaTsvetaKey: Guid!
	GruppaDefektaKey: Guid!
	RassevKey: Guid!
	Razmer1Ot: Double!
	Razmer1Do: Double!
	Cost: Double
	TipTsenKey: Guid
	RecorderType: String
}
type InformationRegisterTsenyPoPreiskurantuRecordType {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	KamenKey: Guid!
	FormaOgrankiKey: Guid!
	TsvetKamniaKey: Guid!
	GruppaTsvetaKey: Guid!
	GruppaDefektaKey: Guid!
	RassevKey: Guid!
	Razmer1Ot: Double!
	Razmer1Do: Double!
	Cost: Double
	TipTsenKey: Guid
	RecorderType: String
}
input AccumulationRegisterTovaryVOtboreInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryVOtboreRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterTovaryVOtbore {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryVOtboreRowType!]!
	RecorderType: String!
}
input AccumulationRegisterTovaryVOtboreRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	ZakazKlienta: String
	DepartmentKey: Guid
	KOtboru: Int64
	Otobrano: Int64
	OtobranoVes: Double
	RecorderType: String!
	ZakazKlientaType: String
}
type AccumulationRegisterTovaryVOtboreRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	ZakazKlienta: String
	DepartmentKey: Guid
	KOtboru: Int64
	Otobrano: Int64
	OtobranoVes: Double
	RecorderType: String!
	ZakazKlientaType: String
}
input AccumulationRegisterRealizovannyeTovaryInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterRealizovannyeTovaryRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterRealizovannyeTovary {
	Recorder: String!
	RecordSet: [AccumulationRegisterRealizovannyeTovaryRowType!]!
	RecorderType: String!
}
input AccumulationRegisterRealizovannyeTovaryRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	DokumentPostavki: String
	Quantity: Int64
	Weight: Double
	Vyruchka: Double
	RecorderType: String!
	SdelkaType: String
	DokumentPostavkiType: String
}
type AccumulationRegisterRealizovannyeTovaryRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	DokumentPostavki: String
	Quantity: Int64
	Weight: Double
	Vyruchka: Double
	RecorderType: String!
	SdelkaType: String
	DokumentPostavkiType: String
}
input AccumulationRegisterDenezhnyeSredstvaKomissioneraInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaKomissioneraRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterDenezhnyeSredstvaKomissionera {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaKomissioneraRowType!]!
	RecorderType: String!
}
input AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
input AccumulationRegisterProdazhi1Input {
	Recorder: String!
	RecordSet: [AccumulationRegisterProdazhi1RowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterProdazhi1 {
	Recorder: String!
	RecordSet: [AccumulationRegisterProdazhi1RowType!]!
	RecorderType: String!
}
input AccumulationRegisterProdazhi1RecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentProdazhi: String
	DokumentOprikhodovaniia: String
	ZakazKlienta: String
	ProektKey: Guid
	PodrazdelenieKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	OrderKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostBezSkidok: Double
	StoimostPostuplenie: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	KolichestvoVozvrat: Int64
	VesVozvrat: Double
	StoimostVozvrat: Double
	StoimostBezSkidokVozvrat: Double
	StoimostPostuplenieVozvrat: Double
	StoimostUprVozvrat: Double
	StoimostBezNDSVozvrat: Double
	SpisaniePartii: Boolean
	RecorderType: String!
	DokumentProdazhiType: String
	DokumentOprikhodovaniiaType: String
	ZakazKlientaType: String
}
type AccumulationRegisterProdazhi1RecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	DogovorKontragentaKey: Guid
	DokumentProdazhi: String
	DokumentOprikhodovaniia: String
	ZakazKlienta: String
	ProektKey: Guid
	PodrazdelenieKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	OrderKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	StoimostBezSkidok: Double
	StoimostPostuplenie: Double
	StoimostUpr: Double
	StoimostBezNDS: Double
	KolichestvoVozvrat: Int64
	VesVozvrat: Double
	StoimostVozvrat: Double
	StoimostBezSkidokVozvrat: Double
	StoimostPostuplenieVozvrat: Double
	StoimostUprVozvrat: Double
	StoimostBezNDSVozvrat: Double
	SpisaniePartii: Boolean
	RecorderType: String!
	DokumentProdazhiType: String
	DokumentOprikhodovaniiaType: String
	ZakazKlientaType: String
}
input AccumulationRegisterTovaryNaSkladakhAMInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryNaSkladakhAMRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterTovaryNaSkladakhAM {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryNaSkladakhAMRowType!]!
	RecorderType: String!
}
input AccumulationRegisterTovaryNaSkladakhAMRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DepartmentKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	ItemKey: Guid
	Quantity: Int64
	Weight: Double
	RoznichnaiaStoimost: Double
	RecorderType: String!
}
type AccumulationRegisterTovaryNaSkladakhAMRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DepartmentKey: Guid
	MetallKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	ItemKey: Guid
	Quantity: Int64
	Weight: Double
	RoznichnaiaStoimost: Double
	RecorderType: String!
}
input AccumulationRegisterSummyPoFinmonitoringuInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterSummyPoFinmonitoringuRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterSummyPoFinmonitoringu {
	Recorder: String!
	RecordSet: [AccumulationRegisterSummyPoFinmonitoringuRowType!]!
	RecorderType: String!
}
input AccumulationRegisterSummyPoFinmonitoringuRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	SummaOtgruzki: Double
	SummaOplaty: Double
	SummaVozvrata: Double
	RecorderType: String!
	DokumentUchetaType: String
}
type AccumulationRegisterSummyPoFinmonitoringuRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	SummaOtgruzki: Double
	SummaOplaty: Double
	SummaVozvrata: Double
	RecorderType: String!
	DokumentUchetaType: String
}
input InformationRegisterTsenyNomenklaturyKontragentovInput {
	Recorder: String!
	RecordSet: [InformationRegisterTsenyNomenklaturyKontragentovRowTypeInput!]!
	RecorderType: String!
}
type InformationRegisterTsenyNomenklaturyKontragentov {
	Recorder: String!
	RecordSet: [InformationRegisterTsenyNomenklaturyKontragentovRowType!]!
	RecorderType: String!
}
input InformationRegisterTsenyNomenklaturyKontragentovRecordTypeInput {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	TipTsenKey: Guid!
	ItemKey: Guid!
	InstanceKey: Guid!
	KharakteristikaNomenklaturyKey: Guid!
	SizeKey: Guid!
	Cost: Double
	ValiutaKey: Guid
	EdinitsaIzmereniiaKey: Guid
	RecorderType: String
}
type InformationRegisterTsenyNomenklaturyKontragentovRecordType {
	Recorder: String
	Period: DateTime!
	LineNumber: Int64
	Active: Boolean
	TipTsenKey: Guid!
	ItemKey: Guid!
	InstanceKey: Guid!
	KharakteristikaNomenklaturyKey: Guid!
	SizeKey: Guid!
	Cost: Double
	ValiutaKey: Guid
	EdinitsaIzmereniiaKey: Guid
	RecorderType: String
}
input AccumulationRegisterVzaimoraschetySKontragentamiInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterVzaimoraschetySKontragentamiRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterVzaimoraschetySKontragentami {
	Recorder: String!
	RecordSet: [AccumulationRegisterVzaimoraschetySKontragentamiRowType!]!
	RecorderType: String!
}
input AccumulationRegisterVzaimoraschetySKontragentamiRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
type AccumulationRegisterVzaimoraschetySKontragentamiRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	DogovorKontragentaKey: Guid
	Sdelka: String
	SummaVzaimoraschetov: Double
	SummaUpr: Double
	RecorderType: String!
	SdelkaType: String
}
input AccumulationRegisterSummyPokupokPoDiskontnymKartamInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterSummyPokupokPoDiskontnymKartamRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterSummyPokupokPoDiskontnymKartam {
	Recorder: String!
	RecordSet: [AccumulationRegisterSummyPokupokPoDiskontnymKartamRowType!]!
	RecorderType: String!
}
input AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	MemberCardKey: Guid
	DataSpisaniia: DateTime
	Sum: Double
	SummaBonusov: Double
	SummaVremennykhBonusov: Double
	RecorderType: String!
}
type AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	MemberCardKey: Guid
	DataSpisaniia: DateTime
	Sum: Double
	SummaBonusov: Double
	SummaVremennykhBonusov: Double
	RecorderType: String!
}
input AccumulationRegisterVypolneniePlanaProdazhInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterVypolneniePlanaProdazhRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterVypolneniePlanaProdazh {
	Recorder: String!
	RecordSet: [AccumulationRegisterVypolneniePlanaProdazhRowType!]!
	RecorderType: String!
}
input AccumulationRegisterVypolneniePlanaProdazhRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	DepartmentKey: Guid
	SummaPlan: Double
	SummaFakt: Double
	RecorderType: String!
}
type AccumulationRegisterVypolneniePlanaProdazhRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	DepartmentKey: Guid
	SummaPlan: Double
	SummaFakt: Double
	RecorderType: String!
}
input AccumulationRegisterDavalcheskiiMetallInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterDavalcheskiiMetallRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterDavalcheskiiMetall {
	Recorder: String!
	RecordSet: [AccumulationRegisterDavalcheskiiMetallRowType!]!
	RecorderType: String!
}
input AccumulationRegisterDavalcheskiiMetallRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Weight: Double
	RecorderType: String!
}
type AccumulationRegisterDavalcheskiiMetallRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	MetallKey: Guid
	Weight: Double
	RecorderType: String!
}
input AccumulationRegisterDenezhnyeSredstvaInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterDenezhnyeSredstva {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaRowType!]!
	RecorderType: String!
}
input AccumulationRegisterDenezhnyeSredstvaRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
}
type AccumulationRegisterDenezhnyeSredstvaRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	OrganizatsiiaKey: Guid
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	Sum: Double
	SummaUpr: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
}
input AccumulationRegisterTovaryPeredannyeInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryPeredannyeRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterTovaryPeredannye {
	Recorder: String!
	RecordSet: [AccumulationRegisterTovaryPeredannyeRowType!]!
	RecorderType: String!
}
input AccumulationRegisterTovaryPeredannyeRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Quantity: Int64
	Weight: Double
	SummaVzaimoraschetov: Double
	RecorderType: String!
	SdelkaType: String
}
type AccumulationRegisterTovaryPeredannyeRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	DogovorKontragentaKey: Guid
	Sdelka: String
	Quantity: Int64
	Weight: Double
	SummaVzaimoraschetov: Double
	RecorderType: String!
	SdelkaType: String
}
input AccumulationRegisterDenezhnyeSredstvaKSpisaniiuInput {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRowTypeInput!]!
	RecorderType: String!
}
type AccumulationRegisterDenezhnyeSredstvaKSpisaniiu {
	Recorder: String!
	RecordSet: [AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRowType!]!
	RecorderType: String!
}
input AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordTypeInput {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentSpisaniia: String
	TypeOfMovingMoneyKey: Guid
	Sum: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentSpisaniiaType: String
}
type AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType {
	Recorder: String!
	Period: DateTime
	LineNumber: Int64!
	Active: Boolean
	RecordType: String
	VidDenezhnykhSredstv: String
	BankovskiiSchetKassa: String
	DokumentSpisaniia: String
	TypeOfMovingMoneyKey: Guid
	Sum: Double
	RecorderType: String!
	BankovskiiSchetKassaType: String
	DokumentSpisaniiaType: String
}
input CatalogDogovoryKontragentovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	ValiutaVzaimoraschetovKey: Guid
	VedenieVzaimoraschetov: String
	VidAgentskogoDogovora: String
	VidVzaimoraschetovKey: Guid
	VidDogovora: String
	Date: DateTime
	DerzhatRezervBezOplatyOgranichennoeVremia: Boolean
	DopustimaiaSummaZadolzhennosti: Double
	DopustimoeChisloDneiZadolzhennosti: Int64
	Comment: String
	KontrolirovatDenezhnyeSredstvaKomitenta: Boolean
	KontrolirovatOtritsatelnyeOstatkiPoDavalcheskomuMetallu: Boolean
	KontrolirovatSummuZadolzhennosti: Boolean
	KontrolirovatChisloDneiZadolzhennosti: Boolean
	Number: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiMenedzherKey: Guid
	OtrazhatSummuRealizatsii: Boolean
	PredostavliatDannyeVFinMonitoring: Boolean
	PriemNaKomissiiuOtFizicheskogoLitsa: Boolean
	ProtsentKomissionnogoVoznagrazhdeniia: Double
	ProtsentPredoplaty: Double
	RaschetyVUslovnykhEdinitsakh: Boolean
	SposobRaschetaKomissionnogoVoznagrazhdeniia: String
	TipTsen: String
	UdalitPerekhodPravaSobstvennostiDliaFinMonitoringa: String
	UsloviiaOplatyKey: Guid
	UchetAgentskogoNDS: Boolean
	ChisloDneiRezervaBezOplaty: Int64
	UsloviiaPriemaIzdeliiNaKomissiiuKey: Guid
	DavalcheskiiMetall: Boolean
	OtsenkaMaterialovVProizvodstvePoSrednei: Boolean
	OtsenkaMaterialovVProizvodstveFIFO: Boolean
	BIdentifikator: String
	TipTsenType: String
}
type CatalogDogovoryKontragentov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	ValiutaVzaimoraschetovKey: Guid
	VedenieVzaimoraschetov: String
	VidAgentskogoDogovora: String
	VidVzaimoraschetovKey: Guid
	VidDogovora: String
	Date: DateTime
	DerzhatRezervBezOplatyOgranichennoeVremia: Boolean
	DopustimaiaSummaZadolzhennosti: Double
	DopustimoeChisloDneiZadolzhennosti: Int64
	Comment: String
	KontrolirovatDenezhnyeSredstvaKomitenta: Boolean
	KontrolirovatOtritsatelnyeOstatkiPoDavalcheskomuMetallu: Boolean
	KontrolirovatSummuZadolzhennosti: Boolean
	KontrolirovatChisloDneiZadolzhennosti: Boolean
	Number: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiMenedzherKey: Guid
	OtrazhatSummuRealizatsii: Boolean
	PredostavliatDannyeVFinMonitoring: Boolean
	PriemNaKomissiiuOtFizicheskogoLitsa: Boolean
	ProtsentKomissionnogoVoznagrazhdeniia: Double
	ProtsentPredoplaty: Double
	RaschetyVUslovnykhEdinitsakh: Boolean
	SposobRaschetaKomissionnogoVoznagrazhdeniia: String
	TipTsen: String
	UdalitPerekhodPravaSobstvennostiDliaFinMonitoringa: String
	UsloviiaOplatyKey: Guid
	UchetAgentskogoNDS: Boolean
	ChisloDneiRezervaBezOplaty: Int64
	UsloviiaPriemaIzdeliiNaKomissiiuKey: Guid
	DavalcheskiiMetall: Boolean
	OtsenkaMaterialovVProizvodstvePoSrednei: Boolean
	OtsenkaMaterialovVProizvodstveFIFO: Boolean
	BIdentifikator: String
	TipTsenType: String
}
input OrderInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	OperationType: String
	Zaarkhivirovan: Boolean
	KassaKKMKey: Guid
	KassirKey: Guid
	KolichestvoDokumenta: Int64
	Comment: String
	GungNumber: Int16
	NumberKKT: Int16
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	SummaOplatyBonusom: Double
	TipDokumenta: String
	ChekKKMKey: Guid
	ChekProbitNaKKM: Boolean
	TsenaVkliuchaetNDS: Boolean
	SkidkiRasschitany: Boolean
	DokumentOsnovanie: String
	MemberCardKey: Guid
	OpisanieKarty: String
	Bonusnaia: Boolean
	Pochta: String
	Telefon: String
	Payments: [DocumentChekKKMOplataRowTypeInput!]
	OplataSertifikatami: [DocumentChekKKMOplataSertifikatamiRowTypeInput!]
	ProdazhaSertifikatov: [DocumentChekKKMProdazhaSertifikatovRowTypeInput!]
	Goods: [DocumentChekKKMTovaryRowTypeInput!]
	DokumentyObmena: [DocumentChekKKMDokumentyObmenaRowTypeInput!]
	DogovoraRassrochkiProdazha: [DocumentChekKKMDogovoraRassrochkiProdazhaRowTypeInput!]
	DogovoraRassrochkiOplata: [DocumentChekKKMDogovoraRassrochkiOplataRowTypeInput!]
	OplataBallami: [DocumentChekKKMOplataBallamiRowTypeInput!]
	SkidkiNatsenki: [DocumentChekKKMSkidkiNatsenkiRowTypeInput!]
	UpravliaemyeSkidki: [DocumentChekKKMUpravliaemyeSkidkiRowTypeInput!]
	Podarki: [DocumentChekKKMPodarkiRowTypeInput!]
	Kupony: [DocumentChekKKMKuponyRowTypeInput!]
	DokumentOsnovanieType: String
}
type Order {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	OperationType: String
	Zaarkhivirovan: Boolean
	KassaKKMKey: Guid
	KassirKey: Guid
	KolichestvoDokumenta: Int64
	Comment: String
	GungNumber: Int16
	NumberKKT: Int16
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	SummaOplatyBonusom: Double
	TipDokumenta: String
	ChekKKMKey: Guid
	ChekProbitNaKKM: Boolean
	TsenaVkliuchaetNDS: Boolean
	SkidkiRasschitany: Boolean
	DokumentOsnovanie: String
	MemberCardKey: Guid
	OpisanieKarty: String
	Bonusnaia: Boolean
	Pochta: String
	Telefon: String
	Payments: [DocumentChekKKMOplataRowType!]
	OplataSertifikatami: [DocumentChekKKMOplataSertifikatamiRowType!]
	ProdazhaSertifikatov: [DocumentChekKKMProdazhaSertifikatovRowType!]
	Goods: [DocumentChekKKMTovaryRowType!]
	DokumentyObmena: [DocumentChekKKMDokumentyObmenaRowType!]
	DogovoraRassrochkiProdazha: [DocumentChekKKMDogovoraRassrochkiProdazhaRowType!]
	DogovoraRassrochkiOplata: [DocumentChekKKMDogovoraRassrochkiOplataRowType!]
	OplataBallami: [DocumentChekKKMOplataBallamiRowType!]
	SkidkiNatsenki: [DocumentChekKKMSkidkiNatsenkiRowType!]
	UpravliaemyeSkidki: [DocumentChekKKMUpravliaemyeSkidkiRowType!]
	Podarki: [DocumentChekKKMPodarkiRowType!]
	Kupony: [DocumentChekKKMKuponyRowType!]
	DokumentOsnovanieType: String
}
input DocumentChekKKMOplataInput {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	NomerVidaOplaty: Int16
	ProtsentTorgovoiUstupki: Double
	Sum: Double
	SummaTorgovoiUstupki: Double
	Khesh: String
	KartaSberbanka: Boolean
	Poslednie4: String
	KodRRN: String
	Identifikator: String
	TransactionId: String
}
type DocumentChekKKMOplata {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	NomerVidaOplaty: Int16
	ProtsentTorgovoiUstupki: Double
	Sum: Double
	SummaTorgovoiUstupki: Double
	Khesh: String
	KartaSberbanka: Boolean
	Poslednie4: String
	KodRRN: String
	Identifikator: String
	TransactionId: String
}
input DocumentChekKKMOplataSertifikatamiInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	SummaPokupkiPogashennaia: Double
}
type DocumentChekKKMOplataSertifikatami {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	SummaPokupkiPogashennaia: Double
}
input DocumentChekKKMProdazhaSertifikatovInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	SrokDeistviiaSertifikataOgranichen: Boolean
	KolichestvoDneiDeistviiaSertifikata: Int64
}
type DocumentChekKKMProdazhaSertifikatov {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	SrokDeistviiaSertifikataOgranichen: Boolean
	KolichestvoDneiDeistviiaSertifikata: Int64
}
input DocumentChekKKMTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	MID: String
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	RegistratsiiaProdazhi: Boolean
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	Shtrikhkod: String
	OrderKey: Guid
	KliuchSviazi: Int64
	ProdazhaPodarka: Boolean
	SumAutoDiscount: Double
	SummaBonusov: Double
	SostavStrokiDliaRassrochkiKey: Guid
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentChekKKMTovary {
	Key: Guid!
	LineNumber: Int64!
	MID: String
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	RegistratsiiaProdazhi: Boolean
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	Shtrikhkod: String
	OrderKey: Guid
	KliuchSviazi: Int64
	ProdazhaPodarka: Boolean
	SumAutoDiscount: Double
	SummaBonusov: Double
	SostavStrokiDliaRassrochkiKey: Guid
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input DocumentChekKKMDokumentyObmenaInput {
	Key: Guid!
	LineNumber: Int64!
	DokumentKey: Guid
	Sum: Double
}
type DocumentChekKKMDokumentyObmena {
	Key: Guid!
	LineNumber: Int64!
	DokumentKey: Guid
	Sum: Double
}
input DocumentChekKKMDogovoraRassrochkiProdazhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorRassrochkiKey: Guid
	Sum: Double
}
type DocumentChekKKMDogovoraRassrochkiProdazha {
	Key: Guid!
	LineNumber: Int64!
	DogovorRassrochkiKey: Guid
	Sum: Double
}
input DocumentChekKKMDogovoraRassrochkiOplataInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorRassrochkiKey: Guid
	Sum: Double
}
type DocumentChekKKMDogovoraRassrochkiOplata {
	Key: Guid!
	LineNumber: Int64!
	DogovorRassrochkiKey: Guid
	Sum: Double
}
input DocumentChekKKMOplataBallamiInput {
	Key: Guid!
	LineNumber: Int64!
	Khesh: String
	Poslednie4: String
	Sum: Double
	Identifikator: String
	TransactionId: String
	TransactionIdSpisaniia: String
}
type DocumentChekKKMOplataBallami {
	Key: Guid!
	LineNumber: Int64!
	Khesh: String
	Poslednie4: String
	Sum: Double
	Identifikator: String
	TransactionId: String
	TransactionIdSpisaniia: String
}
input DocumentChekKKMSkidkiNatsenkiInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	Sum: Double
	SkidkaNatsenkaKey: Guid
}
type DocumentChekKKMSkidkiNatsenki {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	Sum: Double
	SkidkaNatsenkaKey: Guid
}
input DocumentChekKKMUpravliaemyeSkidkiInput {
	Key: Guid!
	LineNumber: Int64!
	SkidkaNatsenkaKey: Guid
}
type DocumentChekKKMUpravliaemyeSkidki {
	Key: Guid!
	LineNumber: Int64!
	SkidkaNatsenkaKey: Guid
}
input DocumentChekKKMPodarkiInput {
	Key: Guid!
	LineNumber: Int64!
	SkidkaNatsenkaKey: Guid
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	SizeKey: Guid
	Quantity: Double
	Weight: Double
	Cost: Double
	Sum: Double
	DepartmentKey: Guid
	KliuchSviazi: Int64
}
type DocumentChekKKMPodarki {
	Key: Guid!
	LineNumber: Int64!
	SkidkaNatsenkaKey: Guid
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	SizeKey: Guid
	Quantity: Double
	Weight: Double
	Cost: Double
	Sum: Double
	DepartmentKey: Guid
	KliuchSviazi: Int64
}
input DocumentChekKKMKuponyInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	KuponKey: Guid
	NomerKupona: String
	Spisyvat: Boolean
}
type DocumentChekKKMKupony {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	KuponKey: Guid
	NomerKupona: String
	Spisyvat: Boolean
}
input DocumentPereotsenkaValiutnykhSredstvInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	VzaimoraschetySKontragentami: Boolean
	VzaimoraschetySPodotchetnymiLitsami: Boolean
	DenezhnyeSredstvaVKassakh: Boolean
	DenezhnyeSredstvaNaBankovskikhSchetakh: Boolean
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PoVsemOrganizatsiiam: Boolean
	PodrazdelenieKey: Guid
	TipDokumenta: String
}
type DocumentPereotsenkaValiutnykhSredstv {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	VzaimoraschetySKontragentami: Boolean
	VzaimoraschetySPodotchetnymiLitsami: Boolean
	DenezhnyeSredstvaVKassakh: Boolean
	DenezhnyeSredstvaNaBankovskikhSchetakh: Boolean
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PoVsemOrganizatsiiam: Boolean
	PodrazdelenieKey: Guid
	TipDokumenta: String
}
input CatalogTipySkidokNatsenokInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ValiutaKey: Guid
	VidSkidki: String
	DliaVseiNomenklatury: Boolean
	DliaVsekhPoluchatelei: Boolean
	ZnachenieUsloviia: String
	ObshcheeVremiaNachala: DateTime
	ObshcheeVremiaOkonchaniia: DateTime
	OgranichenieSkidkiNatsenki: Double
	PoDniamNedeli: Boolean
	ProtsentSkidkiNatsenki: Double
	Uslovie: String
	VremiaPoDniamNedeli: [CatalogTipySkidokNatsenokVremiaPoDniamNedeliRowTypeInput!]
	ZnachenieUsloviiaType: String
}
type CatalogTipySkidokNatsenok {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ValiutaKey: Guid
	VidSkidki: String
	DliaVseiNomenklatury: Boolean
	DliaVsekhPoluchatelei: Boolean
	ZnachenieUsloviia: String
	ObshcheeVremiaNachala: DateTime
	ObshcheeVremiaOkonchaniia: DateTime
	OgranichenieSkidkiNatsenki: Double
	PoDniamNedeli: Boolean
	ProtsentSkidkiNatsenki: Double
	Uslovie: String
	VremiaPoDniamNedeli: [CatalogTipySkidokNatsenokVremiaPoDniamNedeliRowType!]
	ZnachenieUsloviiaType: String
}
input CatalogTipySkidokNatsenokVremiaPoDniamNedeliInput {
	Key: Guid!
	LineNumber: Int64!
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
	Vybran: Boolean
	DenNedeli: String
}
type CatalogTipySkidokNatsenokVremiaPoDniamNedeli {
	Key: Guid!
	LineNumber: Int64!
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
	Vybran: Boolean
	DenNedeli: String
}
input CatalogVidyKodirovokiTsepeiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	IspolzovatDopNastroiki: Boolean
	PrimerArtikula: String
	NastroikiZapolneniiaBase64Data: Binary
	ElementyKodirovki: [CatalogVidyKodirovokiTsepeiElementyKodirovkiRowTypeInput!]
	SootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv: [CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvRowTypeInput!]
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
type CatalogVidyKodirovokiTsepei {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	IspolzovatDopNastroiki: Boolean
	PrimerArtikula: String
	NastroikiZapolneniiaBase64Data: Binary
	ElementyKodirovki: [CatalogVidyKodirovokiTsepeiElementyKodirovkiRowType!]
	SootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv: [CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvRowType!]
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
input CatalogVidyKodirovokiTsepeiElementyKodirovkiInput {
	Key: Guid!
	LineNumber: Int64!
	NeObiazatelnyiElement: Boolean
	Element: String
	SvoistvoKey: Guid
	DlinaElementa: Int64
	NomerStrokiTablChasti: Int64
	PeremennaiaDlina: Boolean
	ZapolniatSvoistvoVSootvetstvii: Boolean
}
type CatalogVidyKodirovokiTsepeiElementyKodirovki {
	Key: Guid!
	LineNumber: Int64!
	NeObiazatelnyiElement: Boolean
	Element: String
	SvoistvoKey: Guid
	DlinaElementa: Int64
	NomerStrokiTablChasti: Int64
	PeremennaiaDlina: Boolean
	ZapolniatSvoistvoVSootvetstvii: Boolean
}
input CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvInput {
	Key: Guid!
	LineNumber: Int64!
	ZnachenieKodirovki: String
	ZnachenieSvoistvaKey: Guid
	NomerStrokiTablChasti: Int64
	SvoistvoKey: Guid
}
type CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv {
	Key: Guid!
	LineNumber: Int64!
	ZnachenieKodirovki: String
	ZnachenieSvoistvaKey: Guid
	NomerStrokiTablChasti: Int64
	SvoistvoKey: Guid
}
input DocumentOtchetKomitentuOProdazhakhInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProtsentKomissionnogoVoznagrazhdeniia: Double
	Sdelka: String
	SposobRaschetaKomissionnogoVoznagrazhdeniia: String
	StavkaNDSVoznagrazhdeniia: String
	SummaVoznagrazhdeniia: Double
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UderzhatKomissionnoeVoznagrazhdenie: Boolean
	UsloviiaOplatyKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	DenezhnyeSredstva: [DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvaRowTypeInput!]
	Goods: [DocumentOtchetKomitentuOProdazhakhTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type DocumentOtchetKomitentuOProdazhakh {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProtsentKomissionnogoVoznagrazhdeniia: Double
	Sdelka: String
	SposobRaschetaKomissionnogoVoznagrazhdeniia: String
	StavkaNDSVoznagrazhdeniia: String
	SummaVoznagrazhdeniia: Double
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UderzhatKomissionnoeVoznagrazhdenie: Boolean
	UsloviiaOplatyKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	DenezhnyeSredstva: [DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvaRowType!]
	Goods: [DocumentOtchetKomitentuOProdazhakhTovaryRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvaInput {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
type DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
input DocumentOtchetKomitentuOProdazhakhTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	OtchetKomitentuKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPostupleniia: Double
	PokupatelKey: Guid
	DataRealizatsii: DateTime
	DokumentProdazhi: String
	SchetFakturaVystavlennyiKey: Guid
	DokumentPostupleniiaType: String
	DokumentProdazhiType: String
}
type DocumentOtchetKomitentuOProdazhakhTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	OtchetKomitentuKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPostupleniia: Double
	PokupatelKey: Guid
	DataRealizatsii: DateTime
	DokumentProdazhi: String
	SchetFakturaVystavlennyiKey: Guid
	DokumentPostupleniiaType: String
	DokumentProdazhiType: String
}
input CatalogKassyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	ValiutaDenezhnykhSredstvKey: Guid
}
type CatalogKassy {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	ValiutaDenezhnykhSredstvKey: Guid
}
input CatalogKassiryInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	NomerKarty: String
	Parol: String
	FizLitsoKey: Guid
	NeRabotaet: Boolean
	KTU: Double
}
type CatalogKassiry {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	NomerKarty: String
	Parol: String
	FizLitsoKey: Guid
	NeRabotaet: Boolean
	KTU: Double
}
input DocumentZaiavkaNaPereotsenkuTovarovInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	ParametryOtboraBase64Data: Binary
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	NastroikiZapolneniiaBase64Data: Binary
	Goods: [DocumentZaiavkaNaPereotsenkuTovarovTovaryRowTypeInput!]
	ParametryOtboraType: String
	NastroikiZapolneniiaType: String
	ParametryOtbora: Stream
	NastroikiZapolneniia: Stream
}
type DocumentZaiavkaNaPereotsenkuTovarov {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	ParametryOtboraBase64Data: Binary
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	NastroikiZapolneniiaBase64Data: Binary
	Goods: [DocumentZaiavkaNaPereotsenkuTovarovTovaryRowType!]
	ParametryOtboraType: String
	NastroikiZapolneniiaType: String
	ParametryOtbora: Stream
	NastroikiZapolneniia: Stream
}
input DocumentZaiavkaNaPereotsenkuTovarovTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaVRoznitseStaraia: Double
}
type DocumentZaiavkaNaPereotsenkuTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaVRoznitseStaraia: Double
}
input CatalogProizvodstvennyeUchastkiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	SobstvennoeProizvodstvo: Boolean
	PoizdelnyiUchet: Boolean
	GIPN: Boolean
}
type CatalogProizvodstvennyeUchastki {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	SobstvennoeProizvodstvo: Boolean
	PoizdelnyiUchet: Boolean
	GIPN: Boolean
}
input DocumentZakrytieZakazovKlientovInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TipDokumenta: String
	Zakazy: [DocumentZakrytieZakazovKlientovZakazyRowTypeInput!]
}
type DocumentZakrytieZakazovKlientov {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TipDokumenta: String
	Zakazy: [DocumentZakrytieZakazovKlientovZakazyRowType!]
}
input DocumentZakrytieZakazovKlientovZakazyInput {
	Key: Guid!
	LineNumber: Int64!
	ZakazKlientaKey: Guid
	PrichinaZakrytiiaZakazaKey: Guid
}
type DocumentZakrytieZakazovKlientovZakazy {
	Key: Guid!
	LineNumber: Int64!
	ZakazKlientaKey: Guid
	PrichinaZakrytiiaZakazaKey: Guid
}
input CatalogProektyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	DataNachala: DateTime
	DataOkonchaniia: DateTime
	Opisanie: String
	OtvetstvennyiKey: Guid
	BIdentifikator: String
	BNomerVersii: String
}
type CatalogProekty {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	DataNachala: DateTime
	DataOkonchaniia: DateTime
	Opisanie: String
	OtvetstvennyiKey: Guid
	BIdentifikator: String
	BNomerVersii: String
}
input DocumentPlatezhnoePoruchenieVkhodiashcheeInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	DataVkhodiashchegoDokumenta: DateTime
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KontragentKey: Guid
	NomerVkhodiashchegoDokumenta: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	PodrazdelenieKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	DokumentSozdanVIuTD: Boolean
	NaznacheniePlatezha: String
	BDataDokumenta: DateTime
	BIdentifikator: String
	ExtendedPayments: [DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhaRowTypeInput!]
	RekvizityKontragenta: [DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentaRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentPlatezhnoePoruchenieVkhodiashchee {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	DataVkhodiashchegoDokumenta: DateTime
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KontragentKey: Guid
	NomerVkhodiashchegoDokumenta: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	PodrazdelenieKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	DokumentSozdanVIuTD: Boolean
	NaznacheniePlatezha: String
	BDataDokumenta: DateTime
	BIdentifikator: String
	ExtendedPayments: [DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhaRowType!]
	RekvizityKontragenta: [DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentaRowType!]
	DokumentOsnovanieType: String
}
input DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentVydachaZakazaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	KolichestvoDokumenta: Int64
}
type DocumentVydachaZakaza {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	KolichestvoDokumenta: Int64
}
input CatalogFormyOgrankiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
}
type CatalogFormyOgranki {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
}
input CatalogFormatyMagazinovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogFormatyMagazinov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input CatalogRabochieMestaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ImiaKompiutera: String
	SetevoiPort: Int64
}
type CatalogRabochieMesta {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ImiaKompiutera: String
	SetevoiPort: Int64
}
input CatalogNastroikiVypolneniiaObmenaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	AdresDliaOtpravkiSoobshcheniiObOshibke: String
	VypolniatDeistviiaPodPolnymiPravami: Boolean
	VypolniatObmenPriZakrytiiSmeny: Boolean
	VypolniatObmenPriPoiavleniiFaila: String
	DinamicheskiIzmeniatIntervalMezhduObmenami: Boolean
	IspolzovatReglamentnyeZadaniia: Boolean
	KazhdoeZavershenieRabotySProgrammoi: Boolean
	KazhdyiZapuskProgrammy: Boolean
	KatalogProverkiDostupnosti: String
	KolichestvoElementovVTranzaktsiiNaVygruzkuDannykh: Int64
	KolichestvoElementovVTranzaktsiiNaZagruzkuDannykh: Int64
	Comment: String
	OtvetstvennyiKey: Guid
	ProverkaSertifikatov: Boolean
	ReglamentnoeZadanie: String
	UchetnaiaZapisOtpravkiSoobshcheniiaObOshibkeKey: Guid
	NastroikiObmena: [CatalogNastroikiVypolneniiaObmenaNastroikiObmenaRowTypeInput!]
	SoobshcheniiaNeIavliaiushchiesiaOshibkami: [CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamiRowTypeInput!]
}
type CatalogNastroikiVypolneniiaObmena {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	AdresDliaOtpravkiSoobshcheniiObOshibke: String
	VypolniatDeistviiaPodPolnymiPravami: Boolean
	VypolniatObmenPriZakrytiiSmeny: Boolean
	VypolniatObmenPriPoiavleniiFaila: String
	DinamicheskiIzmeniatIntervalMezhduObmenami: Boolean
	IspolzovatReglamentnyeZadaniia: Boolean
	KazhdoeZavershenieRabotySProgrammoi: Boolean
	KazhdyiZapuskProgrammy: Boolean
	KatalogProverkiDostupnosti: String
	KolichestvoElementovVTranzaktsiiNaVygruzkuDannykh: Int64
	KolichestvoElementovVTranzaktsiiNaZagruzkuDannykh: Int64
	Comment: String
	OtvetstvennyiKey: Guid
	ProverkaSertifikatov: Boolean
	ReglamentnoeZadanie: String
	UchetnaiaZapisOtpravkiSoobshcheniiaObOshibkeKey: Guid
	NastroikiObmena: [CatalogNastroikiVypolneniiaObmenaNastroikiObmenaRowType!]
	SoobshcheniiaNeIavliaiushchiesiaOshibkami: [CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamiRowType!]
}
input CatalogNastroikiVypolneniiaObmenaNastroikiObmenaInput {
	Key: Guid!
	LineNumber: Int64!
	VypolniaemoeDeistvie: String
	NastroikaObmenaKey: Guid
	VypolniaemoeDeistvieType: String
}
type CatalogNastroikiVypolneniiaObmenaNastroikiObmena {
	Key: Guid!
	LineNumber: Int64!
	VypolniaemoeDeistvie: String
	NastroikaObmenaKey: Guid
	VypolniaemoeDeistvieType: String
}
input CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamiInput {
	Key: Guid!
	LineNumber: Int64!
	TekstSoobshcheniia: String
}
type CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami {
	Key: Guid!
	LineNumber: Int64!
	TekstSoobshcheniia: String
}
input CatalogZnacheniiaSvoistvObieektovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	BIdentifikator: String
	BNomerVersii: String
}
type CatalogZnacheniiaSvoistvObieektov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	BIdentifikator: String
	BNomerVersii: String
}
input DocumentRealizatsiiaTovarovUslugInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AdresDostavki: String
	BankovskiiSchetOrganizatsiiKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DataPolucheniia: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	Natsenka: Double
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	UsloviiaOplatyKey: Guid
	UchityvatVesVstavok: Boolean
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	OtpuskRazreshilKey: Guid
	OtpuskProizvelKey: Guid
	DoverennostNomer: String
	DoverennostData: DateTime
	DoverennostVydana: String
	DoverennostCherezKogo: String
	BDataDokumenta: DateTime
	BIdentifikator: String
	BNomerVersii: String
	Goods: [DocumentRealizatsiiaTovarovUslugTovaryRowTypeInput!]
	Uslugi: [DocumentRealizatsiiaTovarovUslugUslugiRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type DocumentRealizatsiiaTovarovUslug {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AdresDostavki: String
	BankovskiiSchetOrganizatsiiKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DataPolucheniia: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	Natsenka: Double
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	UsloviiaOplatyKey: Guid
	UchityvatVesVstavok: Boolean
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	OtpuskRazreshilKey: Guid
	OtpuskProizvelKey: Guid
	DoverennostNomer: String
	DoverennostData: DateTime
	DoverennostVydana: String
	DoverennostCherezKogo: String
	BDataDokumenta: DateTime
	BIdentifikator: String
	BNomerVersii: String
	Goods: [DocumentRealizatsiiaTovarovUslugTovaryRowType!]
	Uslugi: [DocumentRealizatsiiaTovarovUslugUslugiRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentRealizatsiiaTovarovUslugTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumManualDiscount: Double
	SumAutoDiscount: Double
	VesVstavok: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentRealizatsiiaTovarovUslugTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumManualDiscount: Double
	SumAutoDiscount: Double
	VesVstavok: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input DocumentRealizatsiiaTovarovUslugUslugiInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	ProtsentOtSummyDokumenta: Double
	PercentManualDiscount: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
	SumManualDiscount: Double
}
type DocumentRealizatsiiaTovarovUslugUslugi {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	ProtsentOtSummyDokumenta: Double
	PercentManualDiscount: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
	SumManualDiscount: Double
}
input DocumentSobytieInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AdresElektronnoiPochty: String
	Vazhnost: String
	VidObieekta: String
	VidSobytiia: String
	VremiaNapominaniia: DateTime
	IstochnikInformatsiiPriObrashcheniiKey: Guid
	Comment: String
	KontaktnoeLitso: String
	Kontragent: String
	NapomnitOSobytii: Boolean
	NachaloSobytiia: DateTime
	OkonchanieSobytiia: DateTime
	OpisanieSobytiia: String
	Osnovanie: String
	OtvetstvennyiKey: Guid
	ProektKey: Guid
	SoderzhanieSobytiia: String
	SostoianieSobytiia: String
	TipSobytiia: String
	StoronnieLitsa: [DocumentSobytieStoronnieLitsaRowTypeInput!]
	KontaktnoeLitsoType: String
	KontragentType: String
	OsnovanieType: String
}
type DocumentSobytie {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AdresElektronnoiPochty: String
	Vazhnost: String
	VidObieekta: String
	VidSobytiia: String
	VremiaNapominaniia: DateTime
	IstochnikInformatsiiPriObrashcheniiKey: Guid
	Comment: String
	KontaktnoeLitso: String
	Kontragent: String
	NapomnitOSobytii: Boolean
	NachaloSobytiia: DateTime
	OkonchanieSobytiia: DateTime
	OpisanieSobytiia: String
	Osnovanie: String
	OtvetstvennyiKey: Guid
	ProektKey: Guid
	SoderzhanieSobytiia: String
	SostoianieSobytiia: String
	TipSobytiia: String
	StoronnieLitsa: [DocumentSobytieStoronnieLitsaRowType!]
	KontaktnoeLitsoType: String
	KontragentType: String
	OsnovanieType: String
}
input DocumentSobytieStoronnieLitsaInput {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	LitsoKey: Guid
}
type DocumentSobytieStoronnieLitsa {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	LitsoKey: Guid
}
input CatalogVariantyOtvetovOprosovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	OtsenkaOtveta: Double
	TrebuetRazvernutyiOtvet: Boolean
}
type CatalogVariantyOtvetovOprosov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	OtsenkaOtveta: Double
	TrebuetRazvernutyiOtvet: Boolean
}
input CatalogGruppyPisemElektronnoiPochtyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	ParentKey: Guid
	DeletionMark: Boolean
	IspolzovatPredmetyPisem: Boolean
	Poriadok: Int64
}
type CatalogGruppyPisemElektronnoiPochty {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	ParentKey: Guid
	DeletionMark: Boolean
	IspolzovatPredmetyPisem: Boolean
	Poriadok: Int64
}
input CatalogGruppyPochtovoiRassylkiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	OtvetstvennyiKey: Guid
}
type CatalogGruppyPochtovoiRassylki {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	OtvetstvennyiKey: Guid
}
input CatalogNastroikiOtchetovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	NaimenovanieOtcheta: String
	NastroikaDliaAvtosokhraneniia: Boolean
	NastroikaPoUmolchaniiu: Boolean
	NastroikiOtchetaBase64Data: Binary
	NastroikiPostroiteliaBase64Data: Binary
	GruppyNastroekIPolzovateli: [CatalogNastroikiOtchetovGruppyNastroekIPolzovateliRowTypeInput!]
	NastroikiOtchetaType: String
	NastroikiPostroiteliaType: String
	NastroikiOtcheta: Stream
	NastroikiPostroitelia: Stream
}
type CatalogNastroikiOtchetov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	NaimenovanieOtcheta: String
	NastroikaDliaAvtosokhraneniia: Boolean
	NastroikaPoUmolchaniiu: Boolean
	NastroikiOtchetaBase64Data: Binary
	NastroikiPostroiteliaBase64Data: Binary
	GruppyNastroekIPolzovateli: [CatalogNastroikiOtchetovGruppyNastroekIPolzovateliRowType!]
	NastroikiOtchetaType: String
	NastroikiPostroiteliaType: String
	NastroikiOtcheta: Stream
	NastroikiPostroitelia: Stream
}
input CatalogNastroikiOtchetovGruppyNastroekIPolzovateliInput {
	Key: Guid!
	LineNumber: Int64!
	Obieekt: String
	ObieektType: String
}
type CatalogNastroikiOtchetovGruppyNastroekIPolzovateli {
	Key: Guid!
	LineNumber: Int64!
	Obieekt: String
	ObieektType: String
}
input CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Bonusnaia: Boolean
	SpisyvatBonusyPriVozvrateOtPokupatelia: Boolean
	Skidki: [CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkiRowTypeInput!]
}
type CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Bonusnaia: Boolean
	SpisyvatBonusyPriVozvrateOtPokupatelia: Boolean
	Skidki: [CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkiRowType!]
}
input CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkiInput {
	Key: Guid!
	LineNumber: Int64!
	Kod: Int16
	KonSumma: Double
	MaksimalnyiProtsentSummyPokupkiPriOplateBonusom: Double
	Naimenovanie: String
	NachSumma: Double
	ProtsentSkidki: Double
}
type CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki {
	Key: Guid!
	LineNumber: Int64!
	Kod: Int16
	KonSumma: Double
	MaksimalnyiProtsentSummyPokupkiPriOplateBonusom: Double
	Naimenovanie: String
	NachSumma: Double
	ProtsentSkidki: Double
}
input DepartmentInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	VidSklada: String
	Comment: String
	PodrazdelenieKey: Guid
	RaschetRoznichnykhTsenPoTorgovoiNatsenke: Boolean
	TipTsenRoznichnoiTorgovliKey: Guid
	SegmentIskliuchaemoiNomenklaturyKey: Guid
	DataOtkrytiia: DateTime
	KU: Double
	DirektorKey: Guid
	RegionKey: Guid
	KolichestvoVitrin: Int16
	PloshchadVitrin: Double
	PloshchadTorgovaia: Double
	NomerSektsii: Int16
	FormatMagazinaKey: Guid
	BIdentifikator: String
	BNomerVersii: String
}
type Department {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	VidSklada: String
	Comment: String
	PodrazdelenieKey: Guid
	RaschetRoznichnykhTsenPoTorgovoiNatsenke: Boolean
	TipTsenRoznichnoiTorgovliKey: Guid
	SegmentIskliuchaemoiNomenklaturyKey: Guid
	DataOtkrytiia: DateTime
	KU: Double
	DirektorKey: Guid
	RegionKey: Guid
	KolichestvoVitrin: Int16
	PloshchadVitrin: Double
	PloshchadTorgovaia: Double
	NomerSektsii: Int16
	FormatMagazinaKey: Guid
	BIdentifikator: String
	BNomerVersii: String
}
input CatalogKodyVidovTovarovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	KodTNVED: String
	Opisanie: String
}
type CatalogKodyVidovTovarov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	KodTNVED: String
	Opisanie: String
}
input CatalogRassevyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	KolichestvoKamneiV1KarateDo: Double
	KolichestvoKamneiV1KarateOt: Double
}
type CatalogRassevy {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	KolichestvoKamneiV1KarateDo: Double
	KolichestvoKamneiV1KarateOt: Double
}
input CatalogPrichinyZakrytiiaZakazovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogPrichinyZakrytiiaZakazov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input CatalogSegmentyNomenklaturyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	SposobFormirovaniia: String
	SkhemaKomponovkiDannykhBase64Data: Binary
	Opisanie: String
	OtvetstvennyiKey: Guid
	DataSozdaniia: DateTime
	DataOchistki: DateTime
	ImiaShablonaSKD: String
	BIdentifikator: String
	BNomerVersii: String
	SkhemaKomponovkiDannykhType: String
	SkhemaKomponovkiDannykh: Stream
}
type CatalogSegmentyNomenklatury {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	SposobFormirovaniia: String
	SkhemaKomponovkiDannykhBase64Data: Binary
	Opisanie: String
	OtvetstvennyiKey: Guid
	DataSozdaniia: DateTime
	DataOchistki: DateTime
	ImiaShablonaSKD: String
	BIdentifikator: String
	BNomerVersii: String
	SkhemaKomponovkiDannykhType: String
	SkhemaKomponovkiDannykh: Stream
}
input CatalogSostavStrokiChekaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	OpisanieDliaPechati: String
	Quantity: Double
	Weight: Double
	SummaTovara: Double
	StavkaNDS: String
	SummaNDS: Double
}
type CatalogSostavStrokiCheka {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	OpisanieDliaPechati: String
	Quantity: Double
	Weight: Double
	SummaTovara: Double
	StavkaNDS: String
	SummaNDS: Double
}
input CatalogUsloviiaPriemaIzdeliiNaKomissiiuInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ZaKhranenieUderzhivat: Boolean
	ZaKhranenieStavka: Double
	ZaKhranenieSposobRascheta: String
	ZaKhranenieDenNachala: Int64
	ZaKhranenieRaschityvatSDniaNachala: Boolean
	KomissionnoeVoznagrazhdenieSposobRascheta: String
	KomissionnoeVoznagrazhdenieProtsent: Double
	GrafikUtsenkiRabochieDni: Boolean
	GrafikUtsenok: [CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenokRowTypeInput!]
}
type CatalogUsloviiaPriemaIzdeliiNaKomissiiu {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ZaKhranenieUderzhivat: Boolean
	ZaKhranenieStavka: Double
	ZaKhranenieSposobRascheta: String
	ZaKhranenieDenNachala: Int64
	ZaKhranenieRaschityvatSDniaNachala: Boolean
	KomissionnoeVoznagrazhdenieSposobRascheta: String
	KomissionnoeVoznagrazhdenieProtsent: Double
	GrafikUtsenkiRabochieDni: Boolean
	GrafikUtsenok: [CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenokRowType!]
}
input CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenokInput {
	Key: Guid!
	LineNumber: Int64!
	UtsenkaProtsent: Double
	UtsenkaDnei: Int64
}
type CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok {
	Key: Guid!
	LineNumber: Int64!
	UtsenkaProtsent: Double
	UtsenkaDnei: Int64
}
input CatalogIstochnikiInformatsiiPriObrashcheniiPokupateleiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input DocumentKorrektirovkaDolgaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TipDokumenta: String
	UvelichenieDolgaKontragenta: Double
	UmenshenieDolgaKontragenta: Double
	SummyDolga: [DocumentKorrektirovkaDolgaSummyDolgaRowTypeInput!]
}
type DocumentKorrektirovkaDolga {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TipDokumenta: String
	UvelichenieDolgaKontragenta: Double
	UmenshenieDolgaKontragenta: Double
	SummyDolga: [DocumentKorrektirovkaDolgaSummyDolgaRowType!]
}
input DocumentKorrektirovkaDolgaSummyDolgaInput {
	Key: Guid!
	LineNumber: Int64!
	DataDolga: DateTime
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	Sdelka: String
	UvelichenieDolgaKontragenta: Double
	UmenshenieDolgaKontragenta: Double
	SdelkaType: String
}
type DocumentKorrektirovkaDolgaSummyDolga {
	Key: Guid!
	LineNumber: Int64!
	DataDolga: DateTime
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	Sdelka: String
	UvelichenieDolgaKontragenta: Double
	UmenshenieDolgaKontragenta: Double
	SdelkaType: String
}
input PayTypeInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	BankKreditorKey: Guid
	DogovorVzaimoraschetovBankaKreditoraKey: Guid
	DogovorEkvairingaKey: Guid
	NomerKarty: String
	ProtsentBankovskoiKomissii: Double
	TipOplaty: String
	AvtomaticheskiVychitatSummuKomissii: Boolean
}
type PayType {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	BankKreditorKey: Guid
	DogovorVzaimoraschetovBankaKreditoraKey: Guid
	DogovorEkvairingaKey: Guid
	NomerKarty: String
	ProtsentBankovskoiKomissii: Double
	TipOplaty: String
	AvtomaticheskiVychitatSummuKomissii: Boolean
}
input CatalogKhranilishcheShablonovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Obieekt: String
	ShablonBase64Data: Binary
	TipShablona: String
	ShablonType: String
	Shablon: Stream
}
type CatalogKhranilishcheShablonov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Obieekt: String
	ShablonBase64Data: Binary
	TipShablona: String
	ShablonType: String
	Shablon: Stream
}
input DocumentZaiavkaNaRaskhodovanieSredstvInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtoRazmeshcheniePoZaiavke: Boolean
	AvtoRezervirovaniePoZaiavke: Boolean
	BankovskiiSchetKassa: String
	ValiutaVzaimoraschetovPodotchetnikaKey: Guid
	ValiutaDokumentaKey: Guid
	VidVydachiDenezhnykhSredstv: String
	OperationType: String
	VkliuchatVPlatezhnyiKalendar: Boolean
	DataPogasheniiaAvansa: DateTime
	DataRaskhoda: DateTime
	DokumentOsnovanie: String
	Comment: String
	KontragentKey: Guid
	KratnostDokumenta: Int64
	KursDokumenta: Double
	Nomenklatura: String
	Opisanie: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	Poluchatel: String
	RaschetnyiDokumentKey: Guid
	Sostoianie: String
	SumOfDocument: Double
	TipDokumenta: String
	FormaOplaty: String
	TsFOKey: Guid
	ExtendedPayments: [DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhaRowTypeInput!]
	RazmeshchenieZaiavki: [DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkiRowTypeInput!]
	BankovskiiSchetKassaType: String
	DokumentOsnovanieType: String
	ItemType: String
	PoluchatelType: String
}
type DocumentZaiavkaNaRaskhodovanieSredstv {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtoRazmeshcheniePoZaiavke: Boolean
	AvtoRezervirovaniePoZaiavke: Boolean
	BankovskiiSchetKassa: String
	ValiutaVzaimoraschetovPodotchetnikaKey: Guid
	ValiutaDokumentaKey: Guid
	VidVydachiDenezhnykhSredstv: String
	OperationType: String
	VkliuchatVPlatezhnyiKalendar: Boolean
	DataPogasheniiaAvansa: DateTime
	DataRaskhoda: DateTime
	DokumentOsnovanie: String
	Comment: String
	KontragentKey: Guid
	KratnostDokumenta: Int64
	KursDokumenta: Double
	Nomenklatura: String
	Opisanie: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	Poluchatel: String
	RaschetnyiDokumentKey: Guid
	Sostoianie: String
	SumOfDocument: Double
	TipDokumenta: String
	FormaOplaty: String
	TsFOKey: Guid
	ExtendedPayments: [DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhaRowType!]
	RazmeshchenieZaiavki: [DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkiRowType!]
	BankovskiiSchetKassaType: String
	DokumentOsnovanieType: String
	ItemType: String
	PoluchatelType: String
}
input DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	ProektKey: Guid
	Sdelka: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	Sum: Double
	SdelkaType: String
}
type DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	ProektKey: Guid
	Sdelka: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	Sum: Double
	SdelkaType: String
}
input DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkiInput {
	Key: Guid!
	LineNumber: Int64!
	MestoRazmeshcheniia: String
	Sum: Double
	MestoRazmeshcheniiaType: String
}
type DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki {
	Key: Guid!
	LineNumber: Int64!
	MestoRazmeshcheniia: String
	Sum: Double
	MestoRazmeshcheniiaType: String
}
input DocumentZakrytieZakazovPostavshchikamInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TipDokumenta: String
	Zakazy: [DocumentZakrytieZakazovPostavshchikamZakazyRowTypeInput!]
}
type DocumentZakrytieZakazovPostavshchikam {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TipDokumenta: String
	Zakazy: [DocumentZakrytieZakazovPostavshchikamZakazyRowType!]
}
input DocumentZakrytieZakazovPostavshchikamZakazyInput {
	Key: Guid!
	LineNumber: Int64!
	ZakazPostavshchikuKey: Guid
	PrichinaZakrytiiaZakazaKey: Guid
}
type DocumentZakrytieZakazovPostavshchikamZakazy {
	Key: Guid!
	LineNumber: Int64!
	ZakazPostavshchikuKey: Guid
	PrichinaZakrytiiaZakazaKey: Guid
}
input CatalogVidyKamneiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogVidyKamnei {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input DocumentAnketyKlientovDliaFinMonitoringaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	StatusProverki: Boolean
	NomerSoobshcheniia: Int64
	DataSoobshcheniia: DateTime
	DataPredydushcheiProverki: DateTime
	DataNachalaPerioda: DateTime
	DataOkonchaniiaPerioda: DateTime
	ObshcheeKolichestvoProverennykh: Int64
	KolichestvoProverennykhIurLits: Int64
	KolichestvoProverennykhFizLits: Int64
	ObshcheeKolichestvoNaidennykh: Int64
	KolichestvoNaidennykhIurLits: Int64
	KolichestvoNaidennykhFizLits: Int64
	VariantProverki: Int16
	DataPerechnia: DateTime
	PlanovaiaProverka: Boolean
	Ankety: [DocumentAnketyKlientovDliaFinMonitoringaAnketyRowTypeInput!]
}
type DocumentAnketyKlientovDliaFinMonitoringa {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	StatusProverki: Boolean
	NomerSoobshcheniia: Int64
	DataSoobshcheniia: DateTime
	DataPredydushcheiProverki: DateTime
	DataNachalaPerioda: DateTime
	DataOkonchaniiaPerioda: DateTime
	ObshcheeKolichestvoProverennykh: Int64
	KolichestvoProverennykhIurLits: Int64
	KolichestvoProverennykhFizLits: Int64
	ObshcheeKolichestvoNaidennykh: Int64
	KolichestvoNaidennykhIurLits: Int64
	KolichestvoNaidennykhFizLits: Int64
	VariantProverki: Int16
	DataPerechnia: DateTime
	PlanovaiaProverka: Boolean
	Ankety: [DocumentAnketyKlientovDliaFinMonitoringaAnketyRowType!]
}
input DocumentAnketyKlientovDliaFinMonitoringaAnketyInput {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	VPerechne: Boolean
	InostrannoePublichnoeLitso: Boolean
	DokumentSeriia: String
	DokumentNomer: String
}
type DocumentAnketyKlientovDliaFinMonitoringaAnkety {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	VPerechne: Boolean
	InostrannoePublichnoeLitso: Boolean
	DokumentSeriia: String
	DokumentNomer: String
}
input CatalogDogovoryRassrochkiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	Date: DateTime
	Number: String
	MesiatsevRassrochki: Int64
	SummaEzhemesiachnogoPlatezha: Double
	SummaRassrochki: Double
	SummaPokupki: Double
	SummaPervonachalnogoVznosa: Double
}
type CatalogDogovoryRassrochki {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	Date: DateTime
	Number: String
	MesiatsevRassrochki: Int64
	SummaEzhemesiachnogoPlatezha: Double
	SummaRassrochki: Double
	SummaPokupki: Double
	SummaPervonachalnogoVznosa: Double
}
input CatalogSertifikatyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	KolichestvoDneiDeistviiaSertifikata: Int64
	Comment: String
	Number: String
	ProdazhaPoSvobodnoiTsene: Boolean
	SrokDeistviia: DateTime
	SrokDeistviiaSertifikataOgranichen: Boolean
	Sum: Double
}
type CatalogSertifikaty {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	KolichestvoDneiDeistviiaSertifikata: Int64
	Comment: String
	Number: String
	ProdazhaPoSvobodnoiTsene: Boolean
	SrokDeistviia: DateTime
	SrokDeistviiaSertifikataOgranichen: Boolean
	Sum: Double
}
input DocumentPostuplenieDavalcheskogoMetallaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	Comment: String
	KontragentKey: Guid
	MetallKey: Guid
	NomenklaturaOprikhodovaniiaMetallaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	ProbaLoma: Double
	DepartmentKey: Guid
	Chistota: Boolean
	PodrazdelenieKey: Guid
	Cost: Double
	SumOfDocument: Double
}
type DocumentPostuplenieDavalcheskogoMetalla {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	Comment: String
	KontragentKey: Guid
	MetallKey: Guid
	NomenklaturaOprikhodovaniiaMetallaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	ProbaLoma: Double
	DepartmentKey: Guid
	Chistota: Boolean
	PodrazdelenieKey: Guid
	Cost: Double
	SumOfDocument: Double
}
input DocumentInkassovoePorucheniePeredannoeInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPlatezha: String
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPlanirovaniiaPostupleniiaKey: Guid
	INNPlatelshchika: String
	INNPoluchatelia: String
	Comment: String
	KontragentKey: Guid
	KPPPlatelshchika: String
	KPPPoluchatelia: String
	NaznacheniePlatezha: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	OcherednostPlatezha: Int16
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TekstPlatelshchika: String
	TekstPoluchatelia: String
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	ExtendedPayments: [DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhaRowTypeInput!]
	RekvizityKontragenta: [DocumentInkassovoePorucheniePeredannoeRekvizityKontragentaRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentInkassovoePorucheniePeredannoe {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPlatezha: String
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPlanirovaniiaPostupleniiaKey: Guid
	INNPlatelshchika: String
	INNPoluchatelia: String
	Comment: String
	KontragentKey: Guid
	KPPPlatelshchika: String
	KPPPoluchatelia: String
	NaznacheniePlatezha: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	OcherednostPlatezha: Int16
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TekstPlatelshchika: String
	TekstPoluchatelia: String
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	ExtendedPayments: [DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhaRowType!]
	RekvizityKontragenta: [DocumentInkassovoePorucheniePeredannoeRekvizityKontragentaRowType!]
	DokumentOsnovanieType: String
}
input DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentInkassovoePorucheniePeredannoeRekvizityKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentInkassovoePorucheniePeredannoeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input CatalogFormulyDliaRaschetaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	FormulaRascheta: String
	ObieektPrimeneniiaFormuly: String
}
type CatalogFormulyDliaRascheta {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	FormulaRascheta: String
	ObieektPrimeneniiaFormuly: String
}
input CatalogKuponyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	ProtsentSkidki: Double
	SummaSkidki: Double
	Tirazh: Int64
	SkidkaSummoi: Int16
	IdentifikatorDliaWeb: String
	MaksimalnyiProtsentOtSummyPokupki: Double
	IspolzovanieNeskolkikhKuponovEtogoVida: Boolean
	IspolzovanieSDrugimiKuponami: Boolean
}
type CatalogKupony {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	ProtsentSkidki: Double
	SummaSkidki: Double
	Tirazh: Int64
	SkidkaSummoi: Int16
	IdentifikatorDliaWeb: String
	MaksimalnyiProtsentOtSummyPokupki: Double
	IspolzovanieNeskolkikhKuponovEtogoVida: Boolean
	IspolzovanieSDrugimiKuponami: Boolean
}
input CorrectingInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtvetstvennyiKey: Guid
	TipDokumenta: String
	TablitsaRegistrovNakopleniia: [DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniiaRowTypeInput!]
}
type Correcting {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtvetstvennyiKey: Guid
	TipDokumenta: String
	TablitsaRegistrovNakopleniia: [DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniiaRowType!]
}
input DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniiaInput {
	Key: Guid!
	LineNumber: Int64!
	Imia: String
	Predstavlenie: String
}
type DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia {
	Key: Guid!
	LineNumber: Int64!
	Imia: String
	Predstavlenie: String
}
input DocumentInternetZakazInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	AdresDostavki: String
	OtvetstvennyiKey: Guid
	SumOfDocument: Double
	NomerInternetDokumenta: String
	DataInternetDokumenta: DateTime
	StatusInternetDokumenta: String
	UnikalnyiNomerInternetDokumenta: String
	Telefon: String
	FIO: String
	KurerKey: Guid
	AdresDostavkiKLADR: String
	VremiaDostavki: String
	Samovyvoz: Boolean
	SposobOplaty: Boolean
	Comment: String
	Pochta: String
	BDataDokumenta: DateTime
	BIdentifikator: String
	BNomerVersii: String
	TovaryInternetZakaza: [DocumentInternetZakazTovaryInternetZakazaRowTypeInput!]
	Goods: [DocumentInternetZakazTovaryRowTypeInput!]
}
type DocumentInternetZakaz {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	AdresDostavki: String
	OtvetstvennyiKey: Guid
	SumOfDocument: Double
	NomerInternetDokumenta: String
	DataInternetDokumenta: DateTime
	StatusInternetDokumenta: String
	UnikalnyiNomerInternetDokumenta: String
	Telefon: String
	FIO: String
	KurerKey: Guid
	AdresDostavkiKLADR: String
	VremiaDostavki: String
	Samovyvoz: Boolean
	SposobOplaty: Boolean
	Comment: String
	Pochta: String
	BDataDokumenta: DateTime
	BIdentifikator: String
	BNomerVersii: String
	TovaryInternetZakaza: [DocumentInternetZakazTovaryInternetZakazaRowType!]
	Goods: [DocumentInternetZakazTovaryRowType!]
}
input DocumentInternetZakazTovaryInternetZakazaInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	Sum: Double
	StavkaNDS: String
	SummaNDS: Double
	PodobranPolnostiu: Boolean
	Otkaz: Boolean
}
type DocumentInternetZakazTovaryInternetZakaza {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	Sum: Double
	StavkaNDS: String
	SummaNDS: Double
	PodobranPolnostiu: Boolean
	Otkaz: Boolean
}
input DocumentInternetZakazTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	Sum: Double
	StavkaNDS: String
	SummaNDS: Double
	NomerStrokiZakaza: Int16
	DepartmentKey: Guid
}
type DocumentInternetZakazTovary {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	Quantity: Int64
	Weight: Double
	Cost: Double
	Sum: Double
	StavkaNDS: String
	SummaNDS: Double
	NomerStrokiZakaza: Int16
	DepartmentKey: Guid
}
input CatalogRegionyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogRegiony {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input SaleJournalInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KassaKKMKey: Guid
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	DepartmentKey: Guid
	TypeOfMovingMoneyKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	UslovieProdazhKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	Bonusy: [DocumentOtchetORoznichnykhProdazhakhBonusyRowTypeInput!]
	OplataBankovskimiKreditami: [DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamiRowTypeInput!]
	OplataPlatezhnymiKartami: [DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamiRowTypeInput!]
	OplataSertifikatami: [DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamiRowTypeInput!]
	ProdazhaSertifikatov: [DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovRowTypeInput!]
	Goods: [DocumentOtchetORoznichnykhProdazhakhTovaryRowTypeInput!]
	DogovoraRassrochkiProdazha: [DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhaRowTypeInput!]
	DokumentyObmena: [DocumentOtchetORoznichnykhProdazhakhDokumentyObmenaRowTypeInput!]
	DogovoraRassrochkiOplata: [DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplataRowTypeInput!]
	OplataBallami: [DocumentOtchetORoznichnykhProdazhakhOplataBallamiRowTypeInput!]
	SkidkiNatsenki: [DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkiRowTypeInput!]
	Kupony: [DocumentOtchetORoznichnykhProdazhakhKuponyRowTypeInput!]
	DokumentOsnovanieType: String
}
type SaleJournal {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KassaKKMKey: Guid
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	DepartmentKey: Guid
	TypeOfMovingMoneyKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	UslovieProdazhKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	Bonusy: [DocumentOtchetORoznichnykhProdazhakhBonusyRowType!]
	OplataBankovskimiKreditami: [DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamiRowType!]
	OplataPlatezhnymiKartami: [DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamiRowType!]
	OplataSertifikatami: [DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamiRowType!]
	ProdazhaSertifikatov: [DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovRowType!]
	Goods: [DocumentOtchetORoznichnykhProdazhakhTovaryRowType!]
	DogovoraRassrochkiProdazha: [DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhaRowType!]
	DokumentyObmena: [DocumentOtchetORoznichnykhProdazhakhDokumentyObmenaRowType!]
	DogovoraRassrochkiOplata: [DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplataRowType!]
	OplataBallami: [DocumentOtchetORoznichnykhProdazhakhOplataBallamiRowType!]
	SkidkiNatsenki: [DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkiRowType!]
	Kupony: [DocumentOtchetORoznichnykhProdazhakhKuponyRowType!]
	DokumentOsnovanieType: String
}
input DocumentOtchetORoznichnykhProdazhakhBonusyInput {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
	NomerCheka: String
	SummaNachislennogoBonusa: Double
	SummaNachislennogoBonusaRasschitana: Boolean
	SummaOplaty: Double
	SummaPokupki: Double
	OrderKey: Guid
	OpisanieKarty: String
}
type DocumentOtchetORoznichnykhProdazhakhBonusy {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
	NomerCheka: String
	SummaNachislennogoBonusa: Double
	SummaNachislennogoBonusaRasschitana: Boolean
	SummaOplaty: Double
	SummaPokupki: Double
	OrderKey: Guid
	OpisanieKarty: String
}
input DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamiInput {
	Key: Guid!
	LineNumber: Int64!
	BankKreditorKey: Guid
	VidOplatyKey: Guid
	DogovorVzaimoraschetovBankaKreditoraKey: Guid
	NomerCheka: String
	ProtsentBankovskoiKomissii: Double
	Sum: Double
	SummaBankovskoiKomissii: Double
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami {
	Key: Guid!
	LineNumber: Int64!
	BankKreditorKey: Guid
	VidOplatyKey: Guid
	DogovorVzaimoraschetovBankaKreditoraKey: Guid
	NomerCheka: String
	ProtsentBankovskoiKomissii: Double
	Sum: Double
	SummaBankovskoiKomissii: Double
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamiInput {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	NomerCheka: String
	ProtsentTorgovoiUstupki: Double
	Sum: Double
	SummaTorgovoiUstupki: Double
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	NomerCheka: String
	ProtsentTorgovoiUstupki: Double
	Sum: Double
	SummaTorgovoiUstupki: Double
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamiInput {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	SertifikatKey: Guid
	SrokDeistviiaSertifikataOgranichen: Boolean
	SummaPokupkiPogashennaia: Double
	SummaSertifikata: Double
	OrderKey: Guid
	NomerSertifikata: String
}
type DocumentOtchetORoznichnykhProdazhakhOplataSertifikatami {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	SertifikatKey: Guid
	SrokDeistviiaSertifikataOgranichen: Boolean
	SummaPokupkiPogashennaia: Double
	SummaSertifikata: Double
	OrderKey: Guid
	NomerSertifikata: String
}
input DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovInput {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	SertifikatKey: Guid
	NomerSertifikata: String
	Sum: Double
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	MemberCardKey: Guid
	Kassir: String
	Quantity: Int64
	ItemKey: Guid
	NomerCheka: String
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	SumAutoDiscount: Double
	KliuchSviazi: Int64
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	OpisanieKarty: String
	SostavStrokiDliaRassrochkiKey: Guid
	KassirType: String
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentOtchetORoznichnykhProdazhakhTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	MemberCardKey: Guid
	Kassir: String
	Quantity: Int64
	ItemKey: Guid
	NomerCheka: String
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	SumAutoDiscount: Double
	KliuchSviazi: Int64
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	OpisanieKarty: String
	SostavStrokiDliaRassrochkiKey: Guid
	KassirType: String
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhaInput {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	Sum: Double
	DogovorRassrochkiKey: Guid
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha {
	Key: Guid!
	LineNumber: Int64!
	NomerCheka: String
	Sum: Double
	DogovorRassrochkiKey: Guid
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhDokumentyObmenaInput {
	Key: Guid!
	LineNumber: Int64!
	DokumentKey: Guid
	Sum: Double
	NomerCheka: String
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhDokumentyObmena {
	Key: Guid!
	LineNumber: Int64!
	DokumentKey: Guid
	Sum: Double
	NomerCheka: String
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplataInput {
	Key: Guid!
	LineNumber: Int64!
	Sum: Double
	NomerCheka: String
	DogovorRassrochkiKey: Guid
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata {
	Key: Guid!
	LineNumber: Int64!
	Sum: Double
	NomerCheka: String
	DogovorRassrochkiKey: Guid
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhOplataBallamiInput {
	Key: Guid!
	LineNumber: Int64!
	Khesh: String
	Poslednie4: String
	Sum: Double
	Identifikator: String
	TransactionId: String
	TransactionIdSpisaniia: String
	NomerCheka: String
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhOplataBallami {
	Key: Guid!
	LineNumber: Int64!
	Khesh: String
	Poslednie4: String
	Sum: Double
	Identifikator: String
	TransactionId: String
	TransactionIdSpisaniia: String
	NomerCheka: String
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkiInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	Sum: Double
	SkidkaNatsenkaKey: Guid
	OrderKey: Guid
}
type DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	Sum: Double
	SkidkaNatsenkaKey: Guid
	OrderKey: Guid
}
input DocumentOtchetORoznichnykhProdazhakhKuponyInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	KuponKey: Guid
	NomerKupona: String
	Spisyvat: Boolean
}
type DocumentOtchetORoznichnykhProdazhakhKupony {
	Key: Guid!
	LineNumber: Int64!
	KliuchSviazi: Int64
	KuponKey: Guid
	NomerKupona: String
	Spisyvat: Boolean
}
input DocumentOtmenaSkidokNomenklaturyInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DokumentOsnovanie: String
	Comment: String
	OtvetstvennyiKey: Guid
	Dokumenty: [DocumentOtmenaSkidokNomenklaturyDokumentyRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentOtmenaSkidokNomenklatury {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DokumentOsnovanie: String
	Comment: String
	OtvetstvennyiKey: Guid
	Dokumenty: [DocumentOtmenaSkidokNomenklaturyDokumentyRowType!]
	DokumentOsnovanieType: String
}
input DocumentOtmenaSkidokNomenklaturyDokumentyInput {
	Key: Guid!
	LineNumber: Int64!
	UstanovkaSkidokNomenklaturyKey: Guid
}
type DocumentOtmenaSkidokNomenklaturyDokumenty {
	Key: Guid!
	LineNumber: Int64!
	UstanovkaSkidokNomenklaturyKey: Guid
}
input CatalogTovarnyeGruppyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
}
type CatalogTovarnyeGruppy {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
}
input DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPlanirovaniiaPostupleniiaKey: Guid
	Comment: String
	KontragentKey: Guid
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	RaschetnyiDokument: String
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	NomerVkhodiashchegoDokumenta: String
	DataVkhodiashchegoDokumenta: DateTime
	PodrazdelenieKey: Guid
	NaznacheniePlatezha: String
	ExtendedPayments: [DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaRowTypeInput!]
	RekvizityKontragenta: [DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentaRowTypeInput!]
	DokumentOsnovanieType: String
	RaschetnyiDokumentType: String
}
type DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPlanirovaniiaPostupleniiaKey: Guid
	Comment: String
	KontragentKey: Guid
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	RaschetnyiDokument: String
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	NomerVkhodiashchegoDokumenta: String
	DataVkhodiashchegoDokumenta: DateTime
	PodrazdelenieKey: Guid
	NaznacheniePlatezha: String
	ExtendedPayments: [DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaRowType!]
	RekvizityKontragenta: [DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentaRowType!]
	DokumentOsnovanieType: String
	RaschetnyiDokumentType: String
}
input DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input CatalogOrderKeyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: Int64
	DeletionMark: Boolean
	UnikalnyiIdentifikatorProdazhi: Guid
	TransactionIdSpisaniia: String
	TransactionIdNachisleniia: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	Number: String
	Date: DateTime
	OperationType: String
	MemberCardKey: Guid
}
type CatalogOrderKey {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: Int64
	DeletionMark: Boolean
	UnikalnyiIdentifikatorProdazhi: Guid
	TransactionIdSpisaniia: String
	TransactionIdNachisleniia: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	Number: String
	Date: DateTime
	OperationType: String
	MemberCardKey: Guid
}
input DocumentKassovyiChekKorrektsiiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	KassaKKMKey: Guid
	SumOfDocument: Double
	StavkaNDS: String
	TipKorrektsii: Int16
	TipRascheta: String
	OtvetstvennyiKey: Guid
	Comment: String
	NumberKKT: Int64
	Predpisanie: String
	ProbitChekNaKKT: Boolean
	OsnovanieDliaKorrektsii: String
	DataDokumentaOsnovaniia: DateTime
	NomerDokumentaOsnovaniia: String
	PredstavlenieOsnovaniia: String
	SummaNDS: Double
	GungNumber: Int16
	TipSistemyNalogooblozheniiaKey: Guid
	Payments: [DocumentKassovyiChekKorrektsiiOplataRowTypeInput!]
	OsnovanieDliaKorrektsiiType: String
}
type DocumentKassovyiChekKorrektsii {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	KassaKKMKey: Guid
	SumOfDocument: Double
	StavkaNDS: String
	TipKorrektsii: Int16
	TipRascheta: String
	OtvetstvennyiKey: Guid
	Comment: String
	NumberKKT: Int64
	Predpisanie: String
	ProbitChekNaKKT: Boolean
	OsnovanieDliaKorrektsii: String
	DataDokumentaOsnovaniia: DateTime
	NomerDokumentaOsnovaniia: String
	PredstavlenieOsnovaniia: String
	SummaNDS: Double
	GungNumber: Int16
	TipSistemyNalogooblozheniiaKey: Guid
	Payments: [DocumentKassovyiChekKorrektsiiOplataRowType!]
	OsnovanieDliaKorrektsiiType: String
}
input DocumentKassovyiChekKorrektsiiOplataInput {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
type DocumentKassovyiChekKorrektsiiOplata {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
input DocumentSchetNaOplatuPokupateliuInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AdresDostavki: String
	ValiutaDokumentaKey: Guid
	VremiaNapominaniia: DateTime
	DataOplaty: DateTime
	DataOtgruzki: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	ZakazKlientaKey: Guid
	Comment: String
	KontaktnoeLitso: String
	Kontragent: String
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NapomnitOSobytii: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	StrukturnaiaEdinitsa: String
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	Goods: [DocumentSchetNaOplatuPokupateliuTovaryRowTypeInput!]
	Uslugi: [DocumentSchetNaOplatuPokupateliuUslugiRowTypeInput!]
	DokumentOsnovanieType: String
	KontaktnoeLitsoType: String
	KontragentType: String
	StrukturnaiaEdinitsaType: String
}
type DocumentSchetNaOplatuPokupateliu {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AdresDostavki: String
	ValiutaDokumentaKey: Guid
	VremiaNapominaniia: DateTime
	DataOplaty: DateTime
	DataOtgruzki: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	ZakazKlientaKey: Guid
	Comment: String
	KontaktnoeLitso: String
	Kontragent: String
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NapomnitOSobytii: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	StrukturnaiaEdinitsa: String
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	Goods: [DocumentSchetNaOplatuPokupateliuTovaryRowType!]
	Uslugi: [DocumentSchetNaOplatuPokupateliuUslugiRowType!]
	DokumentOsnovanieType: String
	KontaktnoeLitsoType: String
	KontragentType: String
	StrukturnaiaEdinitsaType: String
}
input DocumentSchetNaOplatuPokupateliuTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumAutoDiscount: Double
	SumManualDiscount: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentSchetNaOplatuPokupateliuTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumAutoDiscount: Double
	SumManualDiscount: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input DocumentSchetNaOplatuPokupateliuUslugiInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	PercentManualDiscount: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
	SumManualDiscount: Double
}
type DocumentSchetNaOplatuPokupateliuUslugi {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	PercentManualDiscount: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
	SumManualDiscount: Double
}
input CatalogNastroikiObmenaDannymiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	FTPAdresObmena: String
	OnLineObmen: Boolean
	AvtomaticheskiZakachivatPochtuPriObmene: Boolean
	AdresDliaOtpravkiSoobshcheniiObOshibke: String
	AutentifikatsiiaWindowsInformatsionnoiBazyDliaPodkliucheniia: Boolean
	VersiiaPlatformyInformatsionnoiBazyDliaPodkliucheniia: String
	VypolnitOtlozhennoeProvedenieDokumentovPoRaspisaniiu: Boolean
	VypolniatArkhivatsiiuFailovObmena: Boolean
	VypolniatDeistviiaPodPolnymiPravami: Boolean
	DobavlenieObieektovIzFonovogoObmena: Int64
	ImiaInformatsionnoiBazyNaServereDliaPodkliucheniia: String
	ImiaServeraInformatsionnoiBazyDliaPodkliucheniia: String
	ImiaFailaPravil: String
	ImiaFailaProtokolaObmenaOnLineObmen: String
	KatalogInformatsionnoiBazyDliaPodkliucheniia: String
	KatalogObmenaInformatsiei: String
	KolichestvoObieektovDliaFonovogoObmena: Int64
	KolichestvoOshibokOtlozhennogoProvedeniiaDliaPrekrashcheniiaOperatsii: Int64
	KolichestvoElementovVTranzaktsiiNaVygruzkuDannykh: Int64
	KolichestvoElementovVTranzaktsiiNaZagruzkuDannykh: Int64
	Comment: String
	MaksimalnyiRazmerOtpravliaemogoPaketaCherezPochtu: Int64
	MaksimalnyiRazmerOtpravliaemogoPoluchaemogoPaketaCherezFTP: Int64
	OtpravitPravilaObmenaPriemniku: Boolean
	ParolFTPSoedineniia: String
	ParolInformatsionnoiBazyDliaPodkliucheniia: String
	ParolNaOtpravku: String
	ParolNaPriem: String
	ParolProksiFTP: String
	PassivnoeFTPSoedinenie: Boolean
	PolzovatelFTPSoedineniia: String
	PolzovatelInformatsionnoiBazyDliaPodkliucheniia: String
	PolzovatelProksiFTP: String
	PortFTPSoedineniia: Int64
	PortProksiFTP: Int64
	PosleOshibkiOtlozhennogoProvedeniiaPrekratitOperatsii: Boolean
	PochtovyiAdresPoluchatelia: String
	PravilaObmenaBase64Data: Binary
	PravilaObmenaDliaPriemnikaBase64Data: Binary
	ProizvoditOtpravkuSoobshchenii: Boolean
	ProizvoditOtpravkuTolkoPriUspeshnomPrieme: Boolean
	ProizvoditPriemSoobshchenii: Boolean
	ProtokolProksiFTP: String
	RezhimOtladkiOnLineObmena: Boolean
	ServerProksiFTP: String
	TipInformatsionnoiBazyDliaPodkliucheniia: Boolean
	TipNastroiki: String
	TipUzlaInformatsionnoiBazy: String
	UdalitIstochnikPravil: String
	UdalitFailObmenaDannymi: Boolean
	UzelInformatsionnoiBazy: String
	UzelFonovogoObmena: String
	UchetnaiaZapisOtpravkiSoobshcheniiaObOshibkeKey: Guid
	UchetnaiaZapisPriemaOtpravkiSoobshcheniiKey: Guid
	NastroikaVariantovPoiskaObieektov: [CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovRowTypeInput!]
	NastroikaVygruzkiDannykh: [CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhRowTypeInput!]
	SoobshcheniiaNeIavliaiushchiesiaOshibkami: [CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamiRowTypeInput!]
	PravilaObmenaType: String
	PravilaObmenaDliaPriemnikaType: String
	TipUzlaInformatsionnoiBazyType: String
	UzelInformatsionnoiBazyType: String
	UzelFonovogoObmenaType: String
	PravilaObmena: Stream
	PravilaObmenaDliaPriemnika: Stream
}
type CatalogNastroikiObmenaDannymi {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	FTPAdresObmena: String
	OnLineObmen: Boolean
	AvtomaticheskiZakachivatPochtuPriObmene: Boolean
	AdresDliaOtpravkiSoobshcheniiObOshibke: String
	AutentifikatsiiaWindowsInformatsionnoiBazyDliaPodkliucheniia: Boolean
	VersiiaPlatformyInformatsionnoiBazyDliaPodkliucheniia: String
	VypolnitOtlozhennoeProvedenieDokumentovPoRaspisaniiu: Boolean
	VypolniatArkhivatsiiuFailovObmena: Boolean
	VypolniatDeistviiaPodPolnymiPravami: Boolean
	DobavlenieObieektovIzFonovogoObmena: Int64
	ImiaInformatsionnoiBazyNaServereDliaPodkliucheniia: String
	ImiaServeraInformatsionnoiBazyDliaPodkliucheniia: String
	ImiaFailaPravil: String
	ImiaFailaProtokolaObmenaOnLineObmen: String
	KatalogInformatsionnoiBazyDliaPodkliucheniia: String
	KatalogObmenaInformatsiei: String
	KolichestvoObieektovDliaFonovogoObmena: Int64
	KolichestvoOshibokOtlozhennogoProvedeniiaDliaPrekrashcheniiaOperatsii: Int64
	KolichestvoElementovVTranzaktsiiNaVygruzkuDannykh: Int64
	KolichestvoElementovVTranzaktsiiNaZagruzkuDannykh: Int64
	Comment: String
	MaksimalnyiRazmerOtpravliaemogoPaketaCherezPochtu: Int64
	MaksimalnyiRazmerOtpravliaemogoPoluchaemogoPaketaCherezFTP: Int64
	OtpravitPravilaObmenaPriemniku: Boolean
	ParolFTPSoedineniia: String
	ParolInformatsionnoiBazyDliaPodkliucheniia: String
	ParolNaOtpravku: String
	ParolNaPriem: String
	ParolProksiFTP: String
	PassivnoeFTPSoedinenie: Boolean
	PolzovatelFTPSoedineniia: String
	PolzovatelInformatsionnoiBazyDliaPodkliucheniia: String
	PolzovatelProksiFTP: String
	PortFTPSoedineniia: Int64
	PortProksiFTP: Int64
	PosleOshibkiOtlozhennogoProvedeniiaPrekratitOperatsii: Boolean
	PochtovyiAdresPoluchatelia: String
	PravilaObmenaBase64Data: Binary
	PravilaObmenaDliaPriemnikaBase64Data: Binary
	ProizvoditOtpravkuSoobshchenii: Boolean
	ProizvoditOtpravkuTolkoPriUspeshnomPrieme: Boolean
	ProizvoditPriemSoobshchenii: Boolean
	ProtokolProksiFTP: String
	RezhimOtladkiOnLineObmena: Boolean
	ServerProksiFTP: String
	TipInformatsionnoiBazyDliaPodkliucheniia: Boolean
	TipNastroiki: String
	TipUzlaInformatsionnoiBazy: String
	UdalitIstochnikPravil: String
	UdalitFailObmenaDannymi: Boolean
	UzelInformatsionnoiBazy: String
	UzelFonovogoObmena: String
	UchetnaiaZapisOtpravkiSoobshcheniiaObOshibkeKey: Guid
	UchetnaiaZapisPriemaOtpravkiSoobshcheniiKey: Guid
	NastroikaVariantovPoiskaObieektov: [CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovRowType!]
	NastroikaVygruzkiDannykh: [CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhRowType!]
	SoobshcheniiaNeIavliaiushchiesiaOshibkami: [CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamiRowType!]
	PravilaObmenaType: String
	PravilaObmenaDliaPriemnikaType: String
	TipUzlaInformatsionnoiBazyType: String
	UzelInformatsionnoiBazyType: String
	UzelFonovogoObmenaType: String
	PravilaObmena: Stream
	PravilaObmenaDliaPriemnika: Stream
}
input CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovInput {
	Key: Guid!
	LineNumber: Int64!
	VariantPoiskaNePodderzhivaetsia: Boolean
	ImiaNastroikiDliaAlgoritma: String
	ImiaNastroikiDliaPolzovatelia: String
	KodPravilaObmena: String
	NaimenovaniePravilaObmena: String
	NastroikaNePodderzhivaetsia: Boolean
	OpisanieNastroikiDliaPolzovatelia: String
	EtoNastroikaDliaVygruzki: Boolean
}
type CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov {
	Key: Guid!
	LineNumber: Int64!
	VariantPoiskaNePodderzhivaetsia: Boolean
	ImiaNastroikiDliaAlgoritma: String
	ImiaNastroikiDliaPolzovatelia: String
	KodPravilaObmena: String
	NaimenovaniePravilaObmena: String
	NastroikaNePodderzhivaetsia: Boolean
	OpisanieNastroikiDliaPolzovatelia: String
	EtoNastroikaDliaVygruzki: Boolean
}
input CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhInput {
	Key: Guid!
	LineNumber: Int64!
	VygruzhatDannye: Boolean
	VygruzhatPoSsylke: Boolean
	KodPravilaVygruzki: String
	KodPravilaObmena: String
	NaimenovaniePravilaVygruzki: String
	NastroikaNePodderzhivaetsia: Boolean
	EtoNastroikaDliaVygruzki: Boolean
}
type CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh {
	Key: Guid!
	LineNumber: Int64!
	VygruzhatDannye: Boolean
	VygruzhatPoSsylke: Boolean
	KodPravilaVygruzki: String
	KodPravilaObmena: String
	NaimenovaniePravilaVygruzki: String
	NastroikaNePodderzhivaetsia: Boolean
	EtoNastroikaDliaVygruzki: Boolean
}
input CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamiInput {
	Key: Guid!
	LineNumber: Int64!
	TekstSoobshcheniia: String
}
type CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami {
	Key: Guid!
	LineNumber: Int64!
	TekstSoobshcheniia: String
}
input DocumentJournalBankovskieRaschetnyeDokumentyInput {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	OperationType: String
	DataOplaty: DateTime
	InformatsiiaKey: Guid
	Comment: String
	KontragentKey: Guid
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetOrganizatsiiKey: Guid
	ChastichnaiaOplata: Boolean
	RefType: String!
	VidOperatsiiType: String
}
type DocumentJournalBankovskieRaschetnyeDokumenty {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	OperationType: String
	DataOplaty: DateTime
	InformatsiiaKey: Guid
	Comment: String
	KontragentKey: Guid
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetOrganizatsiiKey: Guid
	ChastichnaiaOplata: Boolean
	RefType: String!
	VidOperatsiiType: String
}
input DocumentZamenaDiskontnoiKartyInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OtvetstvennyiKey: Guid
	StaraiaDiskontnaiaKartaKey: Guid
	NovaiaDiskontnaiaKartaKey: Guid
	Comment: String
	DokumentSozdanVIuTD: Boolean
	NomerStaroiDiskontnoiKarty: String
	NomerNovoiDiskontnoiKarty: String
	SummaNakoplenii: Double
	SummaBonusov: Double
}
type DocumentZamenaDiskontnoiKarty {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OtvetstvennyiKey: Guid
	StaraiaDiskontnaiaKartaKey: Guid
	NovaiaDiskontnaiaKartaKey: Guid
	Comment: String
	DokumentSozdanVIuTD: Boolean
	NomerStaroiDiskontnoiKarty: String
	NomerNovoiDiskontnoiKarty: String
	SummaNakoplenii: Double
	SummaBonusov: Double
}
input ReturnToSupplierInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetOrganizatsiiKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	OpisanieDefektov: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	PostavshchikuVystavliaetsiaSchetFakturaNaVozvrat: Boolean
	Goods: [DocumentVozvratTovarovPostavshchikuTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type ReturnToSupplier {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetOrganizatsiiKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	OpisanieDefektov: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	PostavshchikuVystavliaetsiaSchetFakturaNaVozvrat: Boolean
	Goods: [DocumentVozvratTovarovPostavshchikuTovaryRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentVozvratTovarovPostavshchikuTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DefektKey: Guid
	DokumentPostupleniia: String
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	DokumentPostupleniiaType: String
}
type DocumentVozvratTovarovPostavshchikuTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DefektKey: Guid
	DokumentPostupleniia: String
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	DokumentPostupleniiaType: String
}
input DocumentInventarizatsiiaTovarovNaSkladeInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	VPechatnykhFormakhTTPechatatRoznichnyeSummy: Boolean
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	UsloviiaProvedeniiaInventarizatsiiBase64Data: Binary
	KhoziaistvennaiaOperatsiiaKey: Guid
	PostroitelOtchetaBase64Data: Binary
	NastroikiZapolneniiaBase64Data: Binary
	Goods: [DocumentInventarizatsiiaTovarovNaSkladeTovaryRowTypeInput!]
	UsloviiaProvedeniiaInventarizatsii: [DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiRowTypeInput!]
	Sertifikaty: [DocumentInventarizatsiiaTovarovNaSkladeSertifikatyRowTypeInput!]
	TovaryVPuti: [DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutiRowTypeInput!]
	UsloviiaProvedeniiaInventarizatsiiType: String
	PostroitelOtchetaType: String
	NastroikiZapolneniiaType: String
	InventoryTerms: Stream
	PostroitelOtcheta: Stream
	NastroikiZapolneniia: Stream
}
type DocumentInventarizatsiiaTovarovNaSklade {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	VPechatnykhFormakhTTPechatatRoznichnyeSummy: Boolean
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	UsloviiaProvedeniiaInventarizatsiiBase64Data: Binary
	KhoziaistvennaiaOperatsiiaKey: Guid
	PostroitelOtchetaBase64Data: Binary
	NastroikiZapolneniiaBase64Data: Binary
	Goods: [DocumentInventarizatsiiaTovarovNaSkladeTovaryRowType!]
	UsloviiaProvedeniiaInventarizatsii: [DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiRowType!]
	Sertifikaty: [DocumentInventarizatsiiaTovarovNaSkladeSertifikatyRowType!]
	TovaryVPuti: [DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutiRowType!]
	UsloviiaProvedeniiaInventarizatsiiType: String
	PostroitelOtchetaType: String
	NastroikiZapolneniiaType: String
	InventoryTerms: Stream
	PostroitelOtcheta: Stream
	NastroikiZapolneniia: Stream
}
input DocumentInventarizatsiiaTovarovNaSkladeTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesUchet: Double
	KachestvoKey: Guid
	Quantity: Int64
	KolichestvoUchet: Int64
	ItemKey: Guid
	NomerVed: String
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaRegl: Double
	SummaUchet: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	OtkloneniePoKolichestvu: Int64
	OtkloneniePoVesu: Double
}
type DocumentInventarizatsiiaTovarovNaSkladeTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesUchet: Double
	KachestvoKey: Guid
	Quantity: Int64
	KolichestvoUchet: Int64
	ItemKey: Guid
	NomerVed: String
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaRegl: Double
	SummaUchet: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	OtkloneniePoKolichestvu: Int64
	OtkloneniePoVesu: Double
}
input DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiInput {
	Key: Guid!
	LineNumber: Int64!
	VidSravneniia: String
	Znachenie: String
	ImiaPolia: String
	ZnachenieType: String
}
type DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii {
	Key: Guid!
	LineNumber: Int64!
	VidSravneniia: String
	Znachenie: String
	ImiaPolia: String
	ZnachenieType: String
}
input DocumentInventarizatsiiaTovarovNaSkladeSertifikatyInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
	SummaUchet: Double
	Quantity: Double
	KolichestvoUchet: Double
	OtkloneniePoKolichestvu: Int64
}
type DocumentInventarizatsiiaTovarovNaSkladeSertifikaty {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
	SummaUchet: Double
	Quantity: Double
	KolichestvoUchet: Double
	OtkloneniePoKolichestvu: Int64
}
input DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutiInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Weight: Double
	Quantity: Double
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KontragentKey: Guid
	DataOtpravki: DateTime
	DokumentOprikhodovaniiaType: String
}
type DocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	SizeKey: Guid
	Weight: Double
	Quantity: Double
	DogovorKontragentaKey: Guid
	DokumentOprikhodovaniia: String
	KontragentKey: Guid
	DataOtpravki: DateTime
	DokumentOprikhodovaniiaType: String
}
input DocumentPrikhodnyiKassovyiOrderInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaVzaimoraschetovPodotchetnikaKey: Guid
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPriemaRoznichnoiVyruchki: String
	DenezhnyiChekKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPlanirovaniiaPostupleniiaKey: Guid
	DokumentSozdanVIuTD: Boolean
	KassaKey: Guid
	Comment: String
	Kontragent: String
	NumberKKT: Int16
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	Osnovanie: String
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	PodrazdelenieKey: Guid
	Prilozhenie: String
	PriniatoOt: String
	RaschetnyiDokumentKey: Guid
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	KursDokumenta: Double
	KratnostDokumenta: Int64
	Pochta: String
	Telefon: String
	ProbitChekNaKKT: Boolean
	KassaKKMKey: Guid
	GungNumber: Int16
	NastroikaRMKKey: Guid
	BDataDokumenta: DateTime
	BIdentifikator: String
	BNomerVersii: String
	ExtendedPayments: [DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhaRowTypeInput!]
	Payments: [DocumentPrikhodnyiKassovyiOrderOplataRowTypeInput!]
	Goods: [DocumentPrikhodnyiKassovyiOrderTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	KontragentType: String
}
type DocumentPrikhodnyiKassovyiOrder {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaVzaimoraschetovPodotchetnikaKey: Guid
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPriemaRoznichnoiVyruchki: String
	DenezhnyiChekKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPlanirovaniiaPostupleniiaKey: Guid
	DokumentSozdanVIuTD: Boolean
	KassaKey: Guid
	Comment: String
	Kontragent: String
	NumberKKT: Int16
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	Osnovanie: String
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	PodrazdelenieKey: Guid
	Prilozhenie: String
	PriniatoOt: String
	RaschetnyiDokumentKey: Guid
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	KursDokumenta: Double
	KratnostDokumenta: Int64
	Pochta: String
	Telefon: String
	ProbitChekNaKKT: Boolean
	KassaKKMKey: Guid
	GungNumber: Int16
	NastroikaRMKKey: Guid
	BDataDokumenta: DateTime
	BIdentifikator: String
	BNomerVersii: String
	ExtendedPayments: [DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhaRowType!]
	Payments: [DocumentPrikhodnyiKassovyiOrderOplataRowType!]
	Goods: [DocumentPrikhodnyiKassovyiOrderTovaryRowType!]
	DokumentOsnovanieType: String
	KontragentType: String
}
input DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPrikhodnyiKassovyiOrderOplataInput {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
type DocumentPrikhodnyiKassovyiOrderOplata {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
input DocumentPrikhodnyiKassovyiOrderTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SummaSkidki: Double
	VidTovaraKKT: String
	TipOplatyTovaraKKT: String
	SummaOsn: Double
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	SummaOpl: Double
}
type DocumentPrikhodnyiKassovyiOrderTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SummaSkidki: Double
	VidTovaraKKT: String
	TipOplatyTovaraKKT: String
	SummaOsn: Double
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	SummaOpl: Double
}
input CatalogPrichinyVozvrataInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogPrichinyVozvrata {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input DocumentDenezhnyiChekInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DataOplaty: DateTime
	KassaKey: Guid
	Comment: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
}
type DocumentDenezhnyiChek {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DataOplaty: DateTime
	KassaKey: Guid
	Comment: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
}
input DocumentVozvratMaterialovIzProizvodstvaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	DepartmentKey: Guid
	TipDokumenta: String
	KhoziaistvennaiaOperatsiiaKey: Guid
	ProizvodstvennyiUchastokKey: Guid
	SobstvennoeProizvodstvo: Boolean
	Goods: [DocumentVozvratMaterialovIzProizvodstvaTovaryRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentVozvratMaterialovIzProizvodstva {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	DepartmentKey: Guid
	TipDokumenta: String
	KhoziaistvennaiaOperatsiiaKey: Guid
	ProizvodstvennyiUchastokKey: Guid
	SobstvennoeProizvodstvo: Boolean
	Goods: [DocumentVozvratMaterialovIzProizvodstvaTovaryRowType!]
	DokumentOsnovanieType: String
}
input DocumentVozvratMaterialovIzProizvodstvaTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
type DocumentVozvratMaterialovIzProizvodstvaTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
input DocumentPereotsenkaTovarovOtdannykhNaKomissiiuInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	TipDokumenta: String
	TipTsenKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	NastroikiZapolneniiaBase64Data: Binary
	Goods: [DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
type DocumentPereotsenkaTovarovOtdannykhNaKomissiiu {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	TipDokumenta: String
	TipTsenKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	NastroikiZapolneniiaBase64Data: Binary
	Goods: [DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovaryRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
input DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaStaraia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaZaGramm: Double
}
type DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaStaraia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaZaGramm: Double
}
input DocumentVvodNachalnykhOstatkovPoRaskhodamUSNInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	Comment: String
	OtvetstvennyiKey: Guid
	VzaimoraschetySPokupateliami: [DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamiRowTypeInput!]
	TovaryProdannye: [DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyeRowTypeInput!]
	VzaimoraschetySPostavshchikami: [DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamiRowTypeInput!]
	TovaryNaOstatkakh: [DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhRowTypeInput!]
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSN {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	Comment: String
	OtvetstvennyiKey: Guid
	VzaimoraschetySPokupateliami: [DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamiRowType!]
	TovaryProdannye: [DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyeRowType!]
	VzaimoraschetySPostavshchikami: [DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamiRowType!]
	TovaryNaOstatkakh: [DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhRowType!]
}
input DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamiInput {
	Key: Guid!
	LineNumber: Int64!
	PokupatelKey: Guid
	DogovorSPokupatelemKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Sum: Double
	SummaSebestoimost: Double
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami {
	Key: Guid!
	LineNumber: Int64!
	PokupatelKey: Guid
	DogovorSPokupatelemKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Sum: Double
	SummaSebestoimost: Double
}
input DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyeInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	PokupatelKey: Guid
	DogovorSPokupatelemKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Quantity: Int64
	Weight: Double
	SummaPostupleniia: Double
	SummaProdazhi: Double
	StatusRaskhoda: String
	DokumentProdazhi: String
	DokumentOprikhodovaniia: String
	DokumentProdazhiType: String
	DokumentOprikhodovaniiaType: String
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	PokupatelKey: Guid
	DogovorSPokupatelemKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Quantity: Int64
	Weight: Double
	SummaPostupleniia: Double
	SummaProdazhi: Double
	StatusRaskhoda: String
	DokumentProdazhi: String
	DokumentOprikhodovaniia: String
	DokumentProdazhiType: String
	DokumentOprikhodovaniiaType: String
}
input DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamiInput {
	Key: Guid!
	LineNumber: Int64!
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Sum: Double
	SummaSebestoimost: Double
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami {
	Key: Guid!
	LineNumber: Int64!
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Sum: Double
	SummaSebestoimost: Double
}
input DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Quantity: Int64
	Weight: Double
	SummaPostupleniia: Double
	StatusRaskhoda: String
	DokumentOprikhodovaniia: String
	SummaPostupleniiaBezNDS: Double
	DokumentOprikhodovaniiaType: String
}
type DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	InstanceKey: Guid
	SupplierKey: Guid
	DogovorSPostavshchikomKey: Guid
	Quantity: Int64
	Weight: Double
	SummaPostupleniia: Double
	StatusRaskhoda: String
	DokumentOprikhodovaniia: String
	SummaPostupleniiaBezNDS: Double
	DokumentOprikhodovaniiaType: String
}
input DocumentGTDImportInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DogovorKontragentaKey: Guid
	DogovorKontragentaReglKey: Guid
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KratnostDokumenta: Int64
	KursVzaimoraschetov: Double
	KursDokumenta: Double
	NomerGTDKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TamozhennyiSbor: Double
	TamozhennyiSborVal: Double
	TamozhennyiShtraf: Double
	TamozhennyiShtrafVal: Double
	TipDokumenta: String
	Razdely: [DocumentGTDImportRazdelyRowTypeInput!]
	Goods: [DocumentGTDImportTovaryRowTypeInput!]
}
type DocumentGTDImport {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DogovorKontragentaKey: Guid
	DogovorKontragentaReglKey: Guid
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KratnostDokumenta: Int64
	KursVzaimoraschetov: Double
	KursDokumenta: Double
	NomerGTDKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TamozhennyiSbor: Double
	TamozhennyiSborVal: Double
	TamozhennyiShtraf: Double
	TamozhennyiShtrafVal: Double
	TipDokumenta: String
	Razdely: [DocumentGTDImportRazdelyRowType!]
	Goods: [DocumentGTDImportTovaryRowType!]
}
input DocumentGTDImportRazdelyInput {
	Key: Guid!
	LineNumber: Int64!
	NDSVValiute: Boolean
	PoshlinaVValiute: Boolean
	StavkaNDS: String
	StavkaPoshliny: Double
	SummaNDS: Double
	SummaPoshliny: Double
	TamozhennaiaStoimost: Double
	TamozhennaiaStoimostVValiuteReglUcheta: Boolean
}
type DocumentGTDImportRazdely {
	Key: Guid!
	LineNumber: Int64!
	NDSVValiute: Boolean
	PoshlinaVValiute: Boolean
	StavkaNDS: String
	StavkaPoshliny: Double
	SummaNDS: Double
	SummaPoshliny: Double
	TamozhennaiaStoimost: Double
	TamozhennaiaStoimostVValiuteReglUcheta: Boolean
}
input DocumentGTDImportTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartii: String
	Quantity: Int64
	ItemKey: Guid
	NomerRazdela: Int16
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	SummaNDS: Double
	SummaPoshliny: Double
	FakturnaiaStoimost: Double
	KharakteristikaNomenklaturyKey: Guid
	DokumentPartiiType: String
}
type DocumentGTDImportTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartii: String
	Quantity: Int64
	ItemKey: Guid
	NomerRazdela: Int16
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	SummaNDS: Double
	SummaPoshliny: Double
	FakturnaiaStoimost: Double
	KharakteristikaNomenklaturyKey: Guid
	DokumentPartiiType: String
}
input DocumentAktSverkiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DokumentOsnovanie: String
	Komentarii: String
	OtvetstvennyiKey: Guid
	Sveren: Boolean
	SumOfDocument: Double
	Goods: [DocumentAktSverkiTovaryRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentAktSverki {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DokumentOsnovanie: String
	Komentarii: String
	OtvetstvennyiKey: Guid
	Sveren: Boolean
	SumOfDocument: Double
	Goods: [DocumentAktSverkiTovaryRowType!]
	DokumentOsnovanieType: String
}
input DocumentAktSverkiTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	MID: String
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	Koef: Double
	NaborKey: Guid
	Naideno: Boolean
	ItemKey: Guid
	NomerNabora: Int64
	Pasport: String
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	Period: DateTime
}
type DocumentAktSverkiTovary {
	Key: Guid!
	LineNumber: Int64!
	MID: String
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	Koef: Double
	NaborKey: Guid
	Naideno: Boolean
	ItemKey: Guid
	NomerNabora: Int64
	Pasport: String
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	Period: DateTime
}
input CatalogFailyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	AvtorKey: Guid
	VladeletsFaila: String
	DataZaema: DateTime
	DataSozdaniia: DateTime
	Zashifrovan: Boolean
	IndeksKartinki: Int64
	Opisanie: String
	PodpisanEP: Boolean
	PolnoeNaimenovanie: String
	RedaktiruetKey: Guid
	TekstKhranilishcheBase64Data: Binary
	TekushchaiaVersiiaKey: Guid
	TekushchaiaVersiiaAvtorKey: Guid
	TekushchaiaVersiiaDataModifikatsiiFaila: DateTime
	TekushchaiaVersiiaDataSozdaniia: DateTime
	TekushchaiaVersiiaKod: String
	TekushchaiaVersiiaNomerVersii: Int64
	TekushchaiaVersiiaPutKFailu: String
	TekushchaiaVersiiaRazmer: Int64
	TekushchaiaVersiiaRasshirenie: String
	KhranitVersii: Boolean
	DopolnitelnyeRekvizity: [CatalogFailyDopolnitelnyeRekvizityRowTypeInput!]
	SertifikatyShifrovaniia: [CatalogFailySertifikatyShifrovaniiaRowTypeInput!]
	VladeletsFailaType: String
	TekstKhranilishcheType: String
	TekstKhranilishche: Stream
}
type CatalogFaily {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	AvtorKey: Guid
	VladeletsFaila: String
	DataZaema: DateTime
	DataSozdaniia: DateTime
	Zashifrovan: Boolean
	IndeksKartinki: Int64
	Opisanie: String
	PodpisanEP: Boolean
	PolnoeNaimenovanie: String
	RedaktiruetKey: Guid
	TekstKhranilishcheBase64Data: Binary
	TekushchaiaVersiiaKey: Guid
	TekushchaiaVersiiaAvtorKey: Guid
	TekushchaiaVersiiaDataModifikatsiiFaila: DateTime
	TekushchaiaVersiiaDataSozdaniia: DateTime
	TekushchaiaVersiiaKod: String
	TekushchaiaVersiiaNomerVersii: Int64
	TekushchaiaVersiiaPutKFailu: String
	TekushchaiaVersiiaRazmer: Int64
	TekushchaiaVersiiaRasshirenie: String
	KhranitVersii: Boolean
	DopolnitelnyeRekvizity: [CatalogFailyDopolnitelnyeRekvizityRowType!]
	SertifikatyShifrovaniia: [CatalogFailySertifikatyShifrovaniiaRowType!]
	VladeletsFailaType: String
	TekstKhranilishcheType: String
	TekstKhranilishche: Stream
}
input CatalogFailyDopolnitelnyeRekvizityInput {
	Key: Guid!
	LineNumber: Int64!
	SvoistvoKey: Guid
	Znachenie: String
	TekstovaiaStroka: String
	ZnachenieType: String
}
type CatalogFailyDopolnitelnyeRekvizity {
	Key: Guid!
	LineNumber: Int64!
	SvoistvoKey: Guid
	Znachenie: String
	TekstovaiaStroka: String
	ZnachenieType: String
}
input CatalogFailySertifikatyShifrovaniiaInput {
	Key: Guid!
	LineNumber: Int64!
	Otpechatok: String
	Predstavlenie: String
	SertifikatBase64Data: Binary
	SertifikatType: String
	Sertifikat: Stream
}
type CatalogFailySertifikatyShifrovaniia {
	Key: Guid!
	LineNumber: Int64!
	Otpechatok: String
	Predstavlenie: String
	SertifikatBase64Data: Binary
	SertifikatType: String
	Sertifikat: Stream
}
input CatalogUchetnyeZapisiElektronnoiPochtyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	POP3Server: String
	SMTPServer: String
	AvtomaticheskaiaUstanovkaPometkiRassmotreno: Boolean
	AvtomaticheskaiaUstanovkaPometkiRassmotrenoPriOtvete: Boolean
	AvtomaticheskaiaUstanovkaPometkiRassmotrenoPriPolucheniiOtveta: Boolean
	AvtoPoluchenieOtpravkaSoobshchenii: Boolean
	AvtosokhraneniePisem: Boolean
	AdresElektronnoiPochty: String
	VremiaOzhidaniiaServera: Int16
	GruppaVkhodiashchieKey: Guid
	GruppaIskhodiashchieKey: Guid
	GruppaUdalennyeKey: Guid
	GruppaChernovikiKey: Guid
	DeistvieAvtopolucheniiaOtpravkiSoobshchenii: String
	DliaVkhodiashchikhOtvetovIPereadresatsiiIskatPismaOsnovaniiaIZapolniatGruppuPisemNovogoPisma: Boolean
	DliaVkhodiashchikhOtvetovIPereadresatsiiIskatPismaOsnovaniiaIZapolniatOsnovanieNovogoPisma: Boolean
	DliaVkhodiashchikhOtvetovIPereadresatsiiIskatPismaOsnovaniiaIZapolniatPredmet: Boolean
	DobavliatPodpisKIskhodiashchimPismam: Boolean
	DobavliatPodpisKOtvetamIPeresylkam: Boolean
	ZapolniatPustoiPredmetDliaNovykhPisemIzTemyPisma: Boolean
	IntervalAvtomaticheskoiUstanovkiOtmetkiRassmotreno: Int16
	IntervalAvtoPolucheniiaOtpravkiSoobshchenii: Int16
	IntervalAvtosokhraneniiaPisem: Int16
	IspolzovatKlassifikatsiiuPisemPoPredmetam: Boolean
	KolichestvoDneiUdaleniiaPisemSServera: Int16
	Login: String
	LoginSMTP: String
	OstavliatKopiiSoobshcheniiNaServere: Boolean
	OtvetstvennyiZaAvtoPoluchenieOtpravkuSoobshcheniiKey: Guid
	Parol: String
	ParolSMTP: String
	PomeshchatOtvetyIPereadresatsiiVTuzheGruppu: Boolean
	PortPOP3: Int16
	PortSMTP: Int16
	TekstPodpisi: String
	TrebuetsiaSMTPAutentifikatsiia: Boolean
	UdaliatPismaSServeraCherez: Boolean
	FormatPismaDliaOtvetovIPereadresatsiiBratIzIskhodnogo: Boolean
	FormatTekstaPismaPoUmolchaniiu: String
	IspolzovatZashchishchennoeSoedinenieDliaVkhodiashcheiPochty: Boolean
	IspolzovatZashchishchennoeSoedinenieDliaIskhodiashcheiPochty: Boolean
	ProtokolVkhodiashcheiPochty: String
	ImiaPolzovatelia: String
	IspolzovatDliaOtpravki: Boolean
	IspolzovatDliaPolucheniia: Boolean
	IspolzovatBezopasnyiVkhodNaServerVkhodiashcheiPochty: Boolean
	IspolzovatBezopasnyiVkhodNaServerIskhodiashcheiPochty: Boolean
	DostupKUchetnoiZapisi: [CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisiRowTypeInput!]
}
type CatalogUchetnyeZapisiElektronnoiPochty {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	POP3Server: String
	SMTPServer: String
	AvtomaticheskaiaUstanovkaPometkiRassmotreno: Boolean
	AvtomaticheskaiaUstanovkaPometkiRassmotrenoPriOtvete: Boolean
	AvtomaticheskaiaUstanovkaPometkiRassmotrenoPriPolucheniiOtveta: Boolean
	AvtoPoluchenieOtpravkaSoobshchenii: Boolean
	AvtosokhraneniePisem: Boolean
	AdresElektronnoiPochty: String
	VremiaOzhidaniiaServera: Int16
	GruppaVkhodiashchieKey: Guid
	GruppaIskhodiashchieKey: Guid
	GruppaUdalennyeKey: Guid
	GruppaChernovikiKey: Guid
	DeistvieAvtopolucheniiaOtpravkiSoobshchenii: String
	DliaVkhodiashchikhOtvetovIPereadresatsiiIskatPismaOsnovaniiaIZapolniatGruppuPisemNovogoPisma: Boolean
	DliaVkhodiashchikhOtvetovIPereadresatsiiIskatPismaOsnovaniiaIZapolniatOsnovanieNovogoPisma: Boolean
	DliaVkhodiashchikhOtvetovIPereadresatsiiIskatPismaOsnovaniiaIZapolniatPredmet: Boolean
	DobavliatPodpisKIskhodiashchimPismam: Boolean
	DobavliatPodpisKOtvetamIPeresylkam: Boolean
	ZapolniatPustoiPredmetDliaNovykhPisemIzTemyPisma: Boolean
	IntervalAvtomaticheskoiUstanovkiOtmetkiRassmotreno: Int16
	IntervalAvtoPolucheniiaOtpravkiSoobshchenii: Int16
	IntervalAvtosokhraneniiaPisem: Int16
	IspolzovatKlassifikatsiiuPisemPoPredmetam: Boolean
	KolichestvoDneiUdaleniiaPisemSServera: Int16
	Login: String
	LoginSMTP: String
	OstavliatKopiiSoobshcheniiNaServere: Boolean
	OtvetstvennyiZaAvtoPoluchenieOtpravkuSoobshcheniiKey: Guid
	Parol: String
	ParolSMTP: String
	PomeshchatOtvetyIPereadresatsiiVTuzheGruppu: Boolean
	PortPOP3: Int16
	PortSMTP: Int16
	TekstPodpisi: String
	TrebuetsiaSMTPAutentifikatsiia: Boolean
	UdaliatPismaSServeraCherez: Boolean
	FormatPismaDliaOtvetovIPereadresatsiiBratIzIskhodnogo: Boolean
	FormatTekstaPismaPoUmolchaniiu: String
	IspolzovatZashchishchennoeSoedinenieDliaVkhodiashcheiPochty: Boolean
	IspolzovatZashchishchennoeSoedinenieDliaIskhodiashcheiPochty: Boolean
	ProtokolVkhodiashcheiPochty: String
	ImiaPolzovatelia: String
	IspolzovatDliaOtpravki: Boolean
	IspolzovatDliaPolucheniia: Boolean
	IspolzovatBezopasnyiVkhodNaServerVkhodiashcheiPochty: Boolean
	IspolzovatBezopasnyiVkhodNaServerIskhodiashcheiPochty: Boolean
	DostupKUchetnoiZapisi: [CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisiRowType!]
}
input CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisiInput {
	Key: Guid!
	LineNumber: Int64!
	Administrirovanie: Boolean
	Otpravka: Boolean
	PolzovatelKey: Guid
	Transport: Boolean
}
type CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi {
	Key: Guid!
	LineNumber: Int64!
	Administrirovanie: Boolean
	Otpravka: Boolean
	PolzovatelKey: Guid
	Transport: Boolean
}
input DocumentPlaniruemoePostuplenieDenezhnykhSredstvInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetKassa: String
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPriemaRoznichnoiVyruchki: String
	VkliuchatVPlatezhnyiKalendar: Boolean
	DataPostupleniia: DateTime
	DokumentOsnovanie: String
	KassaKKM: String
	Comment: String
	KontragentKey: Guid
	KratnostDokumenta: Int64
	KursDokumenta: Double
	Opisanie: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	Sostoianie: String
	StatiaOborotov: String
	SumOfDocument: Double
	TipDokumenta: String
	FormaOplaty: String
	TsFOKey: Guid
	ExtendedPayments: [DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaRowTypeInput!]
	BankovskiiSchetKassaType: String
	DokumentOsnovanieType: String
	KassaKKMType: String
}
type DocumentPlaniruemoePostuplenieDenezhnykhSredstv {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetKassa: String
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPriemaRoznichnoiVyruchki: String
	VkliuchatVPlatezhnyiKalendar: Boolean
	DataPostupleniia: DateTime
	DokumentOsnovanie: String
	KassaKKM: String
	Comment: String
	KontragentKey: Guid
	KratnostDokumenta: Int64
	KursDokumenta: Double
	Opisanie: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	Sostoianie: String
	StatiaOborotov: String
	SumOfDocument: Double
	TipDokumenta: String
	FormaOplaty: String
	TsFOKey: Guid
	ExtendedPayments: [DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaRowType!]
	BankovskiiSchetKassaType: String
	DokumentOsnovanieType: String
	KassaKKMType: String
}
input DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	ProektKey: Guid
	Sdelka: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	Sum: Double
	SdelkaType: String
}
type DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	ProektKey: Guid
	Sdelka: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	Sum: Double
	SdelkaType: String
}
input DocumentPreiskurantTsenNaKamniInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	KamenKey: Guid
	Comment: String
	RassevKey: Guid
	TipTsenKey: Guid
	FormaOgrankiKey: Guid
}
type DocumentPreiskurantTsenNaKamni {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	KamenKey: Guid
	Comment: String
	RassevKey: Guid
	TipTsenKey: Guid
	FormaOgrankiKey: Guid
}
input PurchaseInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Adres: String
	Weight: Double
	DataRozhdeniia: DateTime
	DokumentSozdanVIuTD: Boolean
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	UdostoverenieLichnosti: String
	FizicheskoeLitso: String
	ObmenIzdelii: Boolean
	KontragentKey: Guid
	DokumentSeriia: String
	DokumentNomer: String
	DokumentVidKey: Guid
	DokumentKemVydan: String
	DokumentDataVydachi: DateTime
	DokumentKodPodrazdeleniia: String
	VidProbitiiaSkupki: String
	ProbitNaKKT: Boolean
	Stornirovan: Boolean
	Goods: [DocumentSkupkaTovarovTovaryRowTypeInput!]
}
type Purchase {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Adres: String
	Weight: Double
	DataRozhdeniia: DateTime
	DokumentSozdanVIuTD: Boolean
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	UdostoverenieLichnosti: String
	FizicheskoeLitso: String
	ObmenIzdelii: Boolean
	KontragentKey: Guid
	DokumentSeriia: String
	DokumentNomer: String
	DokumentVidKey: Guid
	DokumentKemVydan: String
	DokumentDataVydachi: DateTime
	DokumentKodPodrazdeleniia: String
	VidProbitiiaSkupki: String
	ProbitNaKKT: Boolean
	Stornirovan: Boolean
	Goods: [DocumentSkupkaTovarovTovaryRowType!]
}
input DocumentSkupkaTovarovTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ItemKey: Guid
	ObshchiiVes: Double
	Sum: Double
	TsenaZaGr: Double
	Opisanie: String
}
type DocumentSkupkaTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ItemKey: Guid
	ObshchiiVes: Double
	Sum: Double
	TsenaZaGr: Double
	Opisanie: String
}
input DocumentSchetFakturaPoluchennyiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DataVkhodiashchegoDokumenta: DateTime
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	Comment: String
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	SformirovanPriVvodeNachalnykhOstatkovNDS: Boolean
	TipDokumenta: String
	NomerIskhodnogoDokumenta: String
	DataIskhodnogoDokumenta: DateTime
	IskhodnyiDokumentKey: Guid
	Korrektirovochnyi: Boolean
	SchetaFakturyVydannyePokupateliam: [DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliamRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentSchetFakturaPoluchennyi {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DataVkhodiashchegoDokumenta: DateTime
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	Comment: String
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	SformirovanPriVvodeNachalnykhOstatkovNDS: Boolean
	TipDokumenta: String
	NomerIskhodnogoDokumenta: String
	DataIskhodnogoDokumenta: DateTime
	IskhodnyiDokumentKey: Guid
	Korrektirovochnyi: Boolean
	SchetaFakturyVydannyePokupateliam: [DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliamRowType!]
	DokumentOsnovanieType: String
}
input DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliamInput {
	Key: Guid!
	LineNumber: Int64!
	PokupatelKey: Guid
	SchetFakturaKey: Guid
	SubkomissionerKey: Guid
	Sum: Double
	NDS: Double
	SummaUvelichenie: Double
	SummaUmenshenie: Double
	SummaNDSUvelichenie: Double
	SummaNDSUmenshenie: Double
}
type DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam {
	Key: Guid!
	LineNumber: Int64!
	PokupatelKey: Guid
	SchetFakturaKey: Guid
	SubkomissionerKey: Guid
	Sum: Double
	NDS: Double
	SummaUvelichenie: Double
	SummaUmenshenie: Double
	SummaNDSUvelichenie: Double
	SummaNDSUmenshenie: Double
}
input DocumentAktKhimicheskogoAnalizaMetallaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KontragentKey: Guid
	NomenklaturaOprikhodovaniiaMetallaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	ProbaLoma: Double
	DepartmentKey: Guid
	UkazanVesChistoty: Boolean
	PodrazdelenieKey: Guid
	Cost: Double
	SumOfDocument: Double
	DokumentOsnovanieType: String
}
type DocumentAktKhimicheskogoAnalizaMetalla {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KontragentKey: Guid
	NomenklaturaOprikhodovaniiaMetallaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	ProbaLoma: Double
	DepartmentKey: Guid
	UkazanVesChistoty: Boolean
	PodrazdelenieKey: Guid
	Cost: Double
	SumOfDocument: Double
	DokumentOsnovanieType: String
}
input CatalogfmKartochkaKontragentaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	Owner: String
	DeletionMark: Boolean
	DannyeKontragenta: [CatalogfmKartochkaKontragentaDannyeKontragentaRowTypeInput!]
	OwnerType: String
}
type CatalogfmKartochkaKontragenta {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	Owner: String
	DeletionMark: Boolean
	DannyeKontragenta: [CatalogfmKartochkaKontragentaDannyeKontragentaRowType!]
	OwnerType: String
}
input CatalogfmKartochkaKontragentaDannyeKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
type CatalogfmKartochkaKontragentaDannyeKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
input DocumentSpisanieProsrochennykhSertifikatovInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Komentarii: String
	OtvetstvennyiKey: Guid
	SumOfDocument: Double
	Sertifikaty: [DocumentSpisanieProsrochennykhSertifikatovSertifikatyRowTypeInput!]
}
type DocumentSpisanieProsrochennykhSertifikatov {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Komentarii: String
	OtvetstvennyiKey: Guid
	SumOfDocument: Double
	Sertifikaty: [DocumentSpisanieProsrochennykhSertifikatovSertifikatyRowType!]
}
input DocumentSpisanieProsrochennykhSertifikatovSertifikatyInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	SertifikatKey: Guid
	SrokDeistviiaDo: DateTime
	Sum: Double
	OrganizatsiiaKey: Guid
	DokumentProdazhi: String
	DokumentProdazhiType: String
}
type DocumentSpisanieProsrochennykhSertifikatovSertifikaty {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	SertifikatKey: Guid
	SrokDeistviiaDo: DateTime
	Sum: Double
	OrganizatsiiaKey: Guid
	DokumentProdazhi: String
	DokumentProdazhiType: String
}
input DocumentZakrytieAvansovPoGrafikuPlatezheiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	Kontragenty: [DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentyRowTypeInput!]
}
type DocumentZakrytieAvansovPoGrafikuPlatezhei {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	Kontragenty: [DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentyRowType!]
}
input DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentyInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
}
type DocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
}
input CatalogTipySistemNalogooblozheniiaKKTInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ZnachenieKomandyDliaAtol: Int16
	ZnachenieKomandyDliaShtrikh: Int16
	FormatPF: String
	ZnachenieKomandyDliaAtol10kh: String
}
type CatalogTipySistemNalogooblozheniiaKKT {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ZnachenieKomandyDliaAtol: Int16
	ZnachenieKomandyDliaShtrikh: Int16
	FormatPF: String
	ZnachenieKomandyDliaAtol10kh: String
}
input DocumentAkkreditivPeredannyiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	VidAkkreditiva: String
	OperationType: String
	VidPlatezha: String
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentyKPredieiavleniiu: String
	DopolnitelnyeUsloviia: String
	ZaiavkaNaRaskhodovanieSredstvKey: Guid
	INNPlatelshchika: String
	INNPoluchatelia: String
	Comment: String
	KontragentKey: Guid
	KPPPlatelshchika: String
	KPPPoluchatelia: String
	NaznacheniePlatezha: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	SrokDeistviia: DateTime
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetDeponenta: String
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TekstPlatelshchika: String
	TekstPoluchatelia: String
	TipDokumenta: String
	UslovieOplaty: String
	ChastichnaiaOplata: Boolean
	ExtendedPayments: [DocumentAkkreditivPeredannyiRasshifrovkaPlatezhaRowTypeInput!]
	RekvizityKontragenta: [DocumentAkkreditivPeredannyiRekvizityKontragentaRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentAkkreditivPeredannyi {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	VidAkkreditiva: String
	OperationType: String
	VidPlatezha: String
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentyKPredieiavleniiu: String
	DopolnitelnyeUsloviia: String
	ZaiavkaNaRaskhodovanieSredstvKey: Guid
	INNPlatelshchika: String
	INNPoluchatelia: String
	Comment: String
	KontragentKey: Guid
	KPPPlatelshchika: String
	KPPPoluchatelia: String
	NaznacheniePlatezha: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	SrokDeistviia: DateTime
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetDeponenta: String
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TekstPlatelshchika: String
	TekstPoluchatelia: String
	TipDokumenta: String
	UslovieOplaty: String
	ChastichnaiaOplata: Boolean
	ExtendedPayments: [DocumentAkkreditivPeredannyiRasshifrovkaPlatezhaRowType!]
	RekvizityKontragenta: [DocumentAkkreditivPeredannyiRekvizityKontragentaRowType!]
	DokumentOsnovanieType: String
}
input DocumentAkkreditivPeredannyiRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentAkkreditivPeredannyiRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentAkkreditivPeredannyiRekvizityKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentAkkreditivPeredannyiRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input SupplierInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	VidDostavki: String
	GolovnoiKontragentKey: Guid
	DokumentUdostoveriaiushchiiLichnost: String
	DopolnitelnoeOpisanie: String
	INN: String
	IstochnikInformatsiiPriObrashcheniiKey: Guid
	KodPoOKPO: String
	Comment: String
	KPP: String
	NaimenovaniePolnoe: String
	OsnovnoeKontaktnoeLitsoKey: Guid
	OsnovnoiBankovskiiSchetKey: Guid
	OsnovnoiVidDeiatelnostiKey: Guid
	OsnovnoiDogovorKontragentaKey: Guid
	OsnovnoiMenedzherPokupateliaKey: Guid
	Pokupatel: Boolean
	Postavshchik: Boolean
	RaspisanieRabotyStrokoi: String
	SrokVypolneniiaZakazaPostavshchikom: Int16
	SrokDeistviiaSvidetelstvaOPostanovkeNaSpetsUchet: DateTime
	IurFizLitso: String
	LoginInternetSaita: String
	SvidetelstvoDataVydachi: DateTime
	SvidetelstvoSeriiaNomer: String
	KodPoOKATO: String
	KodIMNS: String
	NaimenovanieIFNS: String
	KodPoOKVED: String
	OGRN: String
	NeRezident: Boolean
	DataRozhdeniia: DateTime
	OshibkaINN: String
	OshibkaKPP: String
	StranaRegistratsiiKey: Guid
	BIdentifikator: String
	BNomerVersii: String
	VidyDeiatelnosti: [CatalogKontragentyVidyDeiatelnostiRowTypeInput!]
}
type Supplier {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	VidDostavki: String
	GolovnoiKontragentKey: Guid
	DokumentUdostoveriaiushchiiLichnost: String
	DopolnitelnoeOpisanie: String
	INN: String
	IstochnikInformatsiiPriObrashcheniiKey: Guid
	KodPoOKPO: String
	Comment: String
	KPP: String
	NaimenovaniePolnoe: String
	OsnovnoeKontaktnoeLitsoKey: Guid
	OsnovnoiBankovskiiSchetKey: Guid
	OsnovnoiVidDeiatelnostiKey: Guid
	OsnovnoiDogovorKontragentaKey: Guid
	OsnovnoiMenedzherPokupateliaKey: Guid
	Pokupatel: Boolean
	Postavshchik: Boolean
	RaspisanieRabotyStrokoi: String
	SrokVypolneniiaZakazaPostavshchikom: Int16
	SrokDeistviiaSvidetelstvaOPostanovkeNaSpetsUchet: DateTime
	IurFizLitso: String
	LoginInternetSaita: String
	SvidetelstvoDataVydachi: DateTime
	SvidetelstvoSeriiaNomer: String
	KodPoOKATO: String
	KodIMNS: String
	NaimenovanieIFNS: String
	KodPoOKVED: String
	OGRN: String
	NeRezident: Boolean
	DataRozhdeniia: DateTime
	OshibkaINN: String
	OshibkaKPP: String
	StranaRegistratsiiKey: Guid
	BIdentifikator: String
	BNomerVersii: String
	VidyDeiatelnosti: [CatalogKontragentyVidyDeiatelnostiRowType!]
}
input CatalogKontragentyVidyDeiatelnostiInput {
	Key: Guid!
	LineNumber: Int64!
	VidDeiatelnostiKey: Guid
}
type CatalogKontragentyVidyDeiatelnosti {
	Key: Guid!
	LineNumber: Int64!
	VidDeiatelnostiKey: Guid
}
input DocumentInformatsionnoeSoobshchenieInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Email: String
	DataRozhdeniia: DateTime
	KliuchProdazhKey: Guid
	SumOfDocument: Double
	Telefon: String
	FIO: String
	OtkazOtPredieiavleniiaUL: Boolean
	Pol: String
	DokumentUdostoveriaiushchiiLichnost: String
	MemberCardKey: Guid
	DokumentSozdanVIuTD: Boolean
	DokumentVidKey: Guid
	DokumentSeriia: String
	DokumentNomer: String
	DokumentDataVydachi: DateTime
	DokumentKemVydan: String
	DokumentKodPodrazdeleniia: String
	KontragentKey: Guid
	SummaDokumentaBezNal: Double
}
type DocumentInformatsionnoeSoobshchenie {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Email: String
	DataRozhdeniia: DateTime
	KliuchProdazhKey: Guid
	SumOfDocument: Double
	Telefon: String
	FIO: String
	OtkazOtPredieiavleniiaUL: Boolean
	Pol: String
	DokumentUdostoveriaiushchiiLichnost: String
	MemberCardKey: Guid
	DokumentSozdanVIuTD: Boolean
	DokumentVidKey: Guid
	DokumentSeriia: String
	DokumentNomer: String
	DokumentDataVydachi: DateTime
	DokumentKemVydan: String
	DokumentKodPodrazdeleniia: String
	KontragentKey: Guid
	SummaDokumentaBezNal: Double
}
input CatalogVlozheniiaElektronnykhPisemInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	IDFailaPochtovogoPisma: String
	ImiaFaila: String
	ObieektKey: Guid
	Predmet: String
	KhranilishcheBase64Data: Binary
	PredmetType: String
	KhranilishcheType: String
	Khranilishche: Stream
}
type CatalogVlozheniiaElektronnykhPisem {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	IDFailaPochtovogoPisma: String
	ImiaFaila: String
	ObieektKey: Guid
	Predmet: String
	KhranilishcheBase64Data: Binary
	PredmetType: String
	KhranilishcheType: String
	Khranilishche: Stream
}
input DocumentPlatezhnoeTrebovanieVystavlennoeInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPlatezha: String
	DataOplaty: DateTime
	DataOtsylkiDokumentov: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPlanirovaniiaPostupleniiaKey: Guid
	INNPlatelshchika: String
	INNPoluchatelia: String
	Comment: String
	KontragentKey: Guid
	KPPPlatelshchika: String
	KPPPoluchatelia: String
	NaznacheniePlatezha: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OsnovanieDliaBezaktseptnogoSpisaniia: String
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	OcherednostPlatezha: Int16
	SrokDliaAktsepta: Int16
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TekstPlatelshchika: String
	TekstPoluchatelia: String
	TipDokumenta: String
	UslovieOplaty: String
	ChastichnaiaOplata: Boolean
	VidAktsepta: Int16
	ExtendedPayments: [DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhaRowTypeInput!]
	RekvizityKontragenta: [DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentaRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentPlatezhnoeTrebovanieVystavlennoe {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPlatezha: String
	DataOplaty: DateTime
	DataOtsylkiDokumentov: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPlanirovaniiaPostupleniiaKey: Guid
	INNPlatelshchika: String
	INNPoluchatelia: String
	Comment: String
	KontragentKey: Guid
	KPPPlatelshchika: String
	KPPPoluchatelia: String
	NaznacheniePlatezha: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OsnovanieDliaBezaktseptnogoSpisaniia: String
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	OcherednostPlatezha: Int16
	SrokDliaAktsepta: Int16
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TekstPlatelshchika: String
	TekstPoluchatelia: String
	TipDokumenta: String
	UslovieOplaty: String
	ChastichnaiaOplata: Boolean
	VidAktsepta: Int16
	ExtendedPayments: [DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhaRowType!]
	RekvizityKontragenta: [DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentaRowType!]
	DokumentOsnovanieType: String
}
input DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentMarketingovaiaAktsiiaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	NaimenovanieAktsii: String
	Opisanie: String
	DataNachalaDeistviia: DateTime
	DataOkonchaniiaDeistviia: DateTime
	Comment: String
	OtvetstvennyiKey: Guid
	DliaVsekhMagazinov: Boolean
	DliaVsekhMagazinovOdnoRaspisanieSkidok: Boolean
	Magaziny: [DocumentMarketingovaiaAktsiiaMagazinyRowTypeInput!]
	SkidkiNatsenki: [DocumentMarketingovaiaAktsiiaSkidkiNatsenkiRowTypeInput!]
	NaboryZnacheniiDostupa: [DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupaRowTypeInput!]
}
type DocumentMarketingovaiaAktsiia {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	NaimenovanieAktsii: String
	Opisanie: String
	DataNachalaDeistviia: DateTime
	DataOkonchaniiaDeistviia: DateTime
	Comment: String
	OtvetstvennyiKey: Guid
	DliaVsekhMagazinov: Boolean
	DliaVsekhMagazinovOdnoRaspisanieSkidok: Boolean
	Magaziny: [DocumentMarketingovaiaAktsiiaMagazinyRowType!]
	SkidkiNatsenki: [DocumentMarketingovaiaAktsiiaSkidkiNatsenkiRowType!]
	NaboryZnacheniiDostupa: [DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupaRowType!]
}
input DocumentMarketingovaiaAktsiiaMagazinyInput {
	Key: Guid!
	LineNumber: Int64!
	MagazinKey: Guid
}
type DocumentMarketingovaiaAktsiiaMagaziny {
	Key: Guid!
	LineNumber: Int64!
	MagazinKey: Guid
}
input DocumentMarketingovaiaAktsiiaSkidkiNatsenkiInput {
	Key: Guid!
	LineNumber: Int64!
	DataNachala: DateTime
	DataOkonchaniia: DateTime
	MagazinKey: Guid
	SkidkaNatsenkaKey: Guid
}
type DocumentMarketingovaiaAktsiiaSkidkiNatsenki {
	Key: Guid!
	LineNumber: Int64!
	DataNachala: DateTime
	DataOkonchaniia: DateTime
	MagazinKey: Guid
	SkidkaNatsenkaKey: Guid
}
input DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupaInput {
	Key: Guid!
	LineNumber: Int64!
	NomerNabora: Int64
	VidDostupa: String
	ZnachenieDostupa: String
	Chtenie: Boolean
	Dobavlenie: Boolean
	Izmenenie: Boolean
	Udalenie: Boolean
	ZnachenieDostupaType: String
}
type DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa {
	Key: Guid!
	LineNumber: Int64!
	NomerNabora: Int64
	VidDostupa: String
	ZnachenieDostupa: String
	Chtenie: Boolean
	Dobavlenie: Boolean
	Izmenenie: Boolean
	Udalenie: Boolean
	ZnachenieDostupaType: String
}
input CatalogStsenariiObmenovDannymiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	IspolzovatReglamentnoeZadanie: Boolean
	Comment: String
	ReglamentnoeZadanieGUID: String
	NastroikiObmena: [CatalogStsenariiObmenovDannymiNastroikiObmenaRowTypeInput!]
}
type CatalogStsenariiObmenovDannymi {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	IspolzovatReglamentnoeZadanie: Boolean
	Comment: String
	ReglamentnoeZadanieGUID: String
	NastroikiObmena: [CatalogStsenariiObmenovDannymiNastroikiObmenaRowType!]
}
input CatalogStsenariiObmenovDannymiNastroikiObmenaInput {
	Key: Guid!
	LineNumber: Int64!
	UzelInformatsionnoiBazy: String
	VidTransportaObmena: String
	VypolniaemoeDeistvie: String
	KolichestvoElementovVTranzaktsii: Int64
	UzelInformatsionnoiBazyType: String
}
type CatalogStsenariiObmenovDannymiNastroikiObmena {
	Key: Guid!
	LineNumber: Int64!
	UzelInformatsionnoiBazy: String
	VidTransportaObmena: String
	VypolniaemoeDeistvie: String
	KolichestvoElementovVTranzaktsii: Int64
	UzelInformatsionnoiBazyType: String
}
input ItemInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	AnalitikaTipaIzdeliiaKey: Guid
	MID: String
	BazovaiaEdinitsaIzmereniiaKey: Guid
	Vesovoi: Boolean
	VesovoiKoeffitsientVkhozhdeniia: Int64
	VestiUchetPoKharakteristikam: Boolean
	Vstavka: Boolean
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	DliaZakaza: Boolean
	EdinitsaDliaOtchetovKey: Guid
	EdinitsaKhraneniiaOstatkovKey: Guid
	KamenKey: Guid
	KodirovkaKey: Guid
	Comment: String
	Ligatura: Boolean
	Metall: Boolean
	Nabor: Boolean
	NaimenovaniePolnoe: String
	NomenklaturnaiaGruppaKey: Guid
	NomenklaturnaiaGruppaZatratKey: Guid
	NomerGTDKey: Guid
	Opisanie: String
	OsnovnoeIzobrazhenieKey: Guid
	OsnovnoeKachestvoKey: Guid
	OsnovnoiPostavshchikKey: Guid
	ManufacturerKey: Guid
	OtvetstvennyiMenedzherZaPokupkiKey: Guid
	ProbeKey: Guid
	PutKataloga: String
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	SredniiVes: Double
	StavkaNDS: String
	StatiaZatratKey: Guid
	StranaProiskhozhdeniiaKey: Guid
	TypeKey: Guid
	TipShtrikhkodaKey: Guid
	TovarOtpuskaemyiPoSvobodnoiTseneBezOstatkov: Boolean
	Usluga: Boolean
	FailOpisaniiaDliaSaitaKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	TsenovaiaGruppaKey: Guid
	KomplektKey: Guid
	Novinka: Boolean
	KhitProdazh: Boolean
	Skidka: Boolean
	KodVidaTovaraKey: Guid
	BIdentifikator: String
	BNomerVersii: String
	SostavLigatury: [CatalogNomenklaturaSostavLigaturyRowTypeInput!]
}
type Item {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	AnalitikaTipaIzdeliiaKey: Guid
	MID: String
	BazovaiaEdinitsaIzmereniiaKey: Guid
	Vesovoi: Boolean
	VesovoiKoeffitsientVkhozhdeniia: Int64
	VestiUchetPoKharakteristikam: Boolean
	Vstavka: Boolean
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	DliaZakaza: Boolean
	EdinitsaDliaOtchetovKey: Guid
	EdinitsaKhraneniiaOstatkovKey: Guid
	KamenKey: Guid
	KodirovkaKey: Guid
	Comment: String
	Ligatura: Boolean
	Metall: Boolean
	Nabor: Boolean
	NaimenovaniePolnoe: String
	NomenklaturnaiaGruppaKey: Guid
	NomenklaturnaiaGruppaZatratKey: Guid
	NomerGTDKey: Guid
	Opisanie: String
	OsnovnoeIzobrazhenieKey: Guid
	OsnovnoeKachestvoKey: Guid
	OsnovnoiPostavshchikKey: Guid
	ManufacturerKey: Guid
	OtvetstvennyiMenedzherZaPokupkiKey: Guid
	ProbeKey: Guid
	PutKataloga: String
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	SredniiVes: Double
	StavkaNDS: String
	StatiaZatratKey: Guid
	StranaProiskhozhdeniiaKey: Guid
	TypeKey: Guid
	TipShtrikhkodaKey: Guid
	TovarOtpuskaemyiPoSvobodnoiTseneBezOstatkov: Boolean
	Usluga: Boolean
	FailOpisaniiaDliaSaitaKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	TsenovaiaGruppaKey: Guid
	KomplektKey: Guid
	Novinka: Boolean
	KhitProdazh: Boolean
	Skidka: Boolean
	KodVidaTovaraKey: Guid
	BIdentifikator: String
	BNomerVersii: String
	SostavLigatury: [CatalogNomenklaturaSostavLigaturyRowType!]
}
input CatalogNomenklaturaSostavLigaturyInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Double
	ItemKey: Guid
}
type CatalogNomenklaturaSostavLigatury {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Double
	ItemKey: Guid
}
input DocumentOprosInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OprashivaemoeLitso: String
	OtvetstvennyiKey: Guid
	RassylkaKey: Guid
	TipovaiaAnketaKey: Guid
	Voprosy: [DocumentOprosVoprosyRowTypeInput!]
	SostavnoiOtvet: [DocumentOprosSostavnoiOtvetRowTypeInput!]
	OprashivaemoeLitsoType: String
}
type DocumentOpros {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OprashivaemoeLitso: String
	OtvetstvennyiKey: Guid
	RassylkaKey: Guid
	TipovaiaAnketaKey: Guid
	Voprosy: [DocumentOprosVoprosyRowType!]
	SostavnoiOtvet: [DocumentOprosSostavnoiOtvetRowType!]
	OprashivaemoeLitsoType: String
}
input DocumentOprosVoprosyInput {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	Otvet: String
	TipovoiOtvet: String
	TipovoiOtvetType: String
}
type DocumentOprosVoprosy {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	Otvet: String
	TipovoiOtvet: String
	TipovoiOtvetType: String
}
input DocumentOprosSostavnoiOtvetInput {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	VoprosVladeletsKey: Guid
	NomerStrokiVTablitse: Int64
	Otvet: String
	TipovoiOtvet: String
	TipovoiOtvetType: String
}
type DocumentOprosSostavnoiOtvet {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	VoprosVladeletsKey: Guid
	NomerStrokiVTablitse: Int64
	Otvet: String
	TipovoiOtvet: String
	TipovoiOtvetType: String
}
input CatalogGruppyPoluchateleiSkidkiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	SposobFormirovaniia: String
	SkhemaKomponovkiDannykhBase64Data: Binary
	Opisanie: String
	OtvetstvennyiKey: Guid
	DataSozdaniia: DateTime
	ImiaShablonaSKD: String
	TipPoluchatelei: String
	SkhemaKomponovkiDannykhType: String
	SkhemaKomponovkiDannykh: Stream
}
type CatalogGruppyPoluchateleiSkidki {
	Key: Guid!
	DataVersion: String
	Description: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	SposobFormirovaniia: String
	SkhemaKomponovkiDannykhBase64Data: Binary
	Opisanie: String
	OtvetstvennyiKey: Guid
	DataSozdaniia: DateTime
	ImiaShablonaSKD: String
	TipPoluchatelei: String
	SkhemaKomponovkiDannykhType: String
	SkhemaKomponovkiDannykh: Stream
}
input ReassessmentInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtoZapolnenieVIuvS: Boolean
	Weight: Double
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	KolichestvoDokumentaSvereno: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	ParametryOtboraBase64Data: Binary
	PodrazdelenieKey: Guid
	PrikazNaPereotsenku: Boolean
	DepartmentKey: Guid
	TipDokumenta: String
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	KomiscionnyeTovaryFilLits: Boolean
	NastroikiZapolneniiaBase64Data: Binary
	Goods: [DocumentPereotsenkaTovarovVNTTTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	ParametryOtboraType: String
	NastroikiZapolneniiaType: String
	ParametryOtbora: Stream
	NastroikiZapolneniia: Stream
}
type Reassessment {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtoZapolnenieVIuvS: Boolean
	Weight: Double
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	KolichestvoDokumentaSvereno: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	ParametryOtboraBase64Data: Binary
	PodrazdelenieKey: Guid
	PrikazNaPereotsenku: Boolean
	DepartmentKey: Guid
	TipDokumenta: String
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	KomiscionnyeTovaryFilLits: Boolean
	NastroikiZapolneniiaBase64Data: Binary
	Goods: [DocumentPereotsenkaTovarovVNTTTovaryRowType!]
	DokumentOsnovanieType: String
	ParametryOtboraType: String
	NastroikiZapolneniiaType: String
	ParametryOtbora: Stream
	NastroikiZapolneniia: Stream
}
input DocumentPereotsenkaTovarovVNTTTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaVRoznitseStaraia: Double
	Naideno: Boolean
	Dnei: Int64
	DogovorKontragentaKey: Guid
	ProtsentUtsenki: Double
}
type DocumentPereotsenkaTovarovVNTTTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaVRoznitseStaraia: Double
	Naideno: Boolean
	Dnei: Int64
	DogovorKontragentaKey: Guid
	ProtsentUtsenki: Double
}
input CatalogTomaKhraneniiaFailovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Comment: String
	MaksimalnyiRazmer: Int64
	PolnyiPutLinux: String
	PolnyiPutWindows: String
	PoriadokZapolneniia: Int16
}
type CatalogTomaKhraneniiaFailov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Comment: String
	MaksimalnyiRazmer: Int64
	PolnyiPutLinux: String
	PolnyiPutWindows: String
	PoriadokZapolneniia: Int16
}
input DocumentJournalProizvodstvennyeDokumentyInput {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	Quantity: Int64
	Comment: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	Sum: Double
	RefType: String!
}
type DocumentJournalProizvodstvennyeDokumenty {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	Quantity: Int64
	Comment: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	Sum: Double
	RefType: String!
}
input DocumentIzmeneniePravDostupaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AdresElektronnoiPochty: String
	KontragentDoKey: Guid
	OrganizatsiiaDoKey: Guid
	DogovorKontragentaDoKey: Guid
	PravoDostupaDo: String
	OrganizatsiiaKey: Guid
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	PravoDostupa: String
}
type DocumentIzmeneniePravDostupa {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AdresElektronnoiPochty: String
	KontragentDoKey: Guid
	OrganizatsiiaDoKey: Guid
	DogovorKontragentaDoKey: Guid
	PravoDostupaDo: String
	OrganizatsiiaKey: Guid
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	PravoDostupa: String
}
input CatalogNastroikaAssortimentnoiMatritsyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	IspolzuetsiaAssortimentnaiaMatritsa: Boolean
	NastroikaTovarnykhGrupp: [CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGruppRowTypeInput!]
	NastroikaTovarnykhKategorii: [CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriiRowTypeInput!]
	NastroikaTovarnykhPozitsii: [CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiiRowTypeInput!]
}
type CatalogNastroikaAssortimentnoiMatritsy {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	IspolzuetsiaAssortimentnaiaMatritsa: Boolean
	NastroikaTovarnykhGrupp: [CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGruppRowType!]
	NastroikaTovarnykhKategorii: [CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriiRowType!]
	NastroikaTovarnykhPozitsii: [CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiiRowType!]
}
input CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGruppInput {
	Key: Guid!
	LineNumber: Int64!
	VstavkiBase64Data: Binary
	MetallKey: Guid
	Naimenovanie: String
	UsloviiaVkhozhdeniiaBase64Data: Binary
	VstavkiType: String
	UsloviiaVkhozhdeniiaType: String
	Vstavki: Stream
	UsloviiaVkhozhdeniia: Stream
}
type CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp {
	Key: Guid!
	LineNumber: Int64!
	VstavkiBase64Data: Binary
	MetallKey: Guid
	Naimenovanie: String
	UsloviiaVkhozhdeniiaBase64Data: Binary
	VstavkiType: String
	UsloviiaVkhozhdeniiaType: String
	Vstavki: Stream
	UsloviiaVkhozhdeniia: Stream
}
input CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriiInput {
	Key: Guid!
	LineNumber: Int64!
	MiksyVstavokBase64Data: Binary
	Naimenovanie: String
	NomenklaturnaiaGruppaKey: Guid
	UsloviiaVkhozhdeniiaBase64Data: Binary
	MiksyVstavokType: String
	UsloviiaVkhozhdeniiaType: String
	MiksyVstavok: Stream
	UsloviiaVkhozhdeniia: Stream
}
type CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii {
	Key: Guid!
	LineNumber: Int64!
	MiksyVstavokBase64Data: Binary
	Naimenovanie: String
	NomenklaturnaiaGruppaKey: Guid
	UsloviiaVkhozhdeniiaBase64Data: Binary
	MiksyVstavokType: String
	UsloviiaVkhozhdeniiaType: String
	MiksyVstavok: Stream
	UsloviiaVkhozhdeniia: Stream
}
input CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiiInput {
	Key: Guid!
	LineNumber: Int64!
	SvoistvoTovara: String
	UsloviiaVkhozhdeniiaBase64Data: Binary
	Naimenovanie: String
	UsloviiaVkhozhdeniiaType: String
	UsloviiaVkhozhdeniia: Stream
}
type CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii {
	Key: Guid!
	LineNumber: Int64!
	SvoistvoTovara: String
	UsloviiaVkhozhdeniiaBase64Data: Binary
	Naimenovanie: String
	UsloviiaVkhozhdeniiaType: String
	UsloviiaVkhozhdeniia: Stream
}
input DocumentJournalDokumentyKontragentovInput {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	Weight: Double
	OperationType: String
	Informatsiia: String
	Quantity: Int64
	Comment: String
	Kontragent: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sum: Double
	RefType: String!
	VidOperatsiiType: String
	InformatsiiaType: String
	KontragentType: String
}
type DocumentJournalDokumentyKontragentov {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	Weight: Double
	OperationType: String
	Informatsiia: String
	Quantity: Int64
	Comment: String
	Kontragent: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sum: Double
	RefType: String!
	VidOperatsiiType: String
	InformatsiiaType: String
	KontragentType: String
}
input MoveInstanceInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	OperationType: String
	VnutrenniiZakazKey: Guid
	DliaKontroliaUnikalnosti: String
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	IspolzovatRezhimSverki: Boolean
	KolichestvoDokumenta: Int64
	KolichestvoDokumentaSvereno: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	PrichinaVozvrataKey: Guid
	SkladOtpravitelKey: Guid
	SkladPoluchatelKey: Guid
	SumOfDocument: Double
	TipDokumenta: String
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	KurerKey: Guid
	Sertifikaty: [DocumentPeremeshchenieTovarovSertifikatyRowTypeInput!]
	Goods: [DocumentPeremeshchenieTovarovTovaryRowTypeInput!]
	SpisokZaiavok: [DocumentPeremeshchenieTovarovSpisokZaiavokRowTypeInput!]
	DokumentOsnovanieType: String
}
type MoveInstance {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	OperationType: String
	VnutrenniiZakazKey: Guid
	DliaKontroliaUnikalnosti: String
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	IspolzovatRezhimSverki: Boolean
	KolichestvoDokumenta: Int64
	KolichestvoDokumentaSvereno: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	PrichinaVozvrataKey: Guid
	SkladOtpravitelKey: Guid
	SkladPoluchatelKey: Guid
	SumOfDocument: Double
	TipDokumenta: String
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	KurerKey: Guid
	Sertifikaty: [DocumentPeremeshchenieTovarovSertifikatyRowType!]
	Goods: [DocumentPeremeshchenieTovarovTovaryRowType!]
	SpisokZaiavok: [DocumentPeremeshchenieTovarovSpisokZaiavokRowType!]
	DokumentOsnovanieType: String
}
input DocumentPeremeshchenieTovarovSertifikatyInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
type DocumentPeremeshchenieTovarovSertifikaty {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
input DocumentPeremeshchenieTovarovTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentRezerva: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StoimostBezNDS: Double
	StoimostSNDS: Double
	Sum: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaVRoznitseGr: Double
	TsenaPostupleniia: Double
	Naideno: Boolean
	InternetZakazKey: Guid
	DokumentRezervaType: String
}
type DocumentPeremeshchenieTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentRezerva: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StoimostBezNDS: Double
	StoimostSNDS: Double
	Sum: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaVRoznitseGr: Double
	TsenaPostupleniia: Double
	Naideno: Boolean
	InternetZakazKey: Guid
	DokumentRezervaType: String
}
input DocumentPeremeshchenieTovarovSpisokZaiavokInput {
	Key: Guid!
	LineNumber: Int64!
	ZaiavkaNaPeremeshchenie: String
}
type DocumentPeremeshchenieTovarovSpisokZaiavok {
	Key: Guid!
	LineNumber: Int64!
	ZaiavkaNaPeremeshchenie: String
}
input DocumentZakrytieZaiavokNaRaskhodovanieSredstvInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	KontragentKey: Guid
	OtborDataKon: DateTime
	OtborDataNach: DateTime
	OtborKontragent: Boolean
	OtborOtvetstvennyi: Boolean
	OtborTsFO: Boolean
	OtvetstvennyiKey: Guid
	OtvetstvennyiZaiavkaKey: Guid
	Sostoianie: String
	TipDokumenta: String
	TsFOKey: Guid
	ZaiavkiNaRaskhodovanieSredstv: [DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvRowTypeInput!]
}
type DocumentZakrytieZaiavokNaRaskhodovanieSredstv {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	KontragentKey: Guid
	OtborDataKon: DateTime
	OtborDataNach: DateTime
	OtborKontragent: Boolean
	OtborOtvetstvennyi: Boolean
	OtborTsFO: Boolean
	OtvetstvennyiKey: Guid
	OtvetstvennyiZaiavkaKey: Guid
	Sostoianie: String
	TipDokumenta: String
	TsFOKey: Guid
	ZaiavkiNaRaskhodovanieSredstv: [DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvRowType!]
}
input DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvInput {
	Key: Guid!
	LineNumber: Int64!
	ValiutaZaiavkaKey: Guid
	ZaiavkaKey: Guid
	KontragentKey: Guid
	OstatokZaiavka: Double
	OstatokRazmeshchenie: Double
	OstatokRezerv: Double
	OtvetstvennyiKey: Guid
}
type DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv {
	Key: Guid!
	LineNumber: Int64!
	ValiutaZaiavkaKey: Guid
	ZaiavkaKey: Guid
	KontragentKey: Guid
	OstatokZaiavka: Double
	OstatokRazmeshchenie: Double
	OstatokRezerv: Double
	OtvetstvennyiKey: Guid
}
input MemberCardInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Email: String
	Adres: String
	Bonusnaia: Boolean
	MemberCardTypeKey: Guid
	DataRozhdeniia: DateTime
	Nakopitelnaia: Boolean
	Number: String
	Sum: Double
	SkhemaNakopitelnykhSkidokKey: Guid
	Telefon: String
	TipDiskontnoiKarty: String
	FIO: String
	VladeletsKarty: String
	Pol: String
	KontragentKey: Guid
	DataAnkety: DateTime
	SoglasieNaSMSRassylku: Boolean
	SoglasieNaPochtovuiuRassylku: Boolean
	VladeletsKartyType: String
}
type MemberCard {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Email: String
	Adres: String
	Bonusnaia: Boolean
	MemberCardTypeKey: Guid
	DataRozhdeniia: DateTime
	Nakopitelnaia: Boolean
	Number: String
	Sum: Double
	SkhemaNakopitelnykhSkidokKey: Guid
	Telefon: String
	TipDiskontnoiKarty: String
	FIO: String
	VladeletsKarty: String
	Pol: String
	KontragentKey: Guid
	DataAnkety: DateTime
	SoglasieNaSMSRassylku: Boolean
	SoglasieNaPochtovuiuRassylku: Boolean
	VladeletsKartyType: String
}
input DocumentABCKlassifikatsiiaPokupateleiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DataNachala: DateTime
	DataOkonchaniia: DateTime
	Comment: String
	ProtsentAKlassa: Double
	ProtsentBKlassa: Double
	ProtsentCKlassa: Double
	TablitsaRaspredeleniiaKontragentov: [DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovRowTypeInput!]
}
type DocumentABCKlassifikatsiiaPokupatelei {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DataNachala: DateTime
	DataOkonchaniia: DateTime
	Comment: String
	ProtsentAKlassa: Double
	ProtsentBKlassa: Double
	ProtsentCKlassa: Double
	TablitsaRaspredeleniiaKontragentov: [DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovRowType!]
}
input DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovInput {
	Key: Guid!
	LineNumber: Int64!
	ABCKlassifikatsiia: String
	ABCKlassifikatsiiaStaraia: String
	ZnachenieParametra: Double
	KontragentKey: Guid
	MenedzherKontragentaKey: Guid
}
type DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov {
	Key: Guid!
	LineNumber: Int64!
	ABCKlassifikatsiia: String
	ABCKlassifikatsiiaStaraia: String
	ZnachenieParametra: Double
	KontragentKey: Guid
	MenedzherKontragentaKey: Guid
}
input CatalogIdentifikatoryObieektovMetadannykhInput {
	Key: Guid!
	DataVersion: String
	Description: String
	ParentKey: Guid
	DeletionMark: Boolean
	PoriadokKollektsii: Int16
	Imia: String
	Sinonim: String
	PolnoeImia: String
	PolnyiSinonim: String
	BezDannykh: Boolean
	ZnacheniePustoiSsylki: String
	KliuchObieektaMetadannykhBase64Data: Binary
	NovaiaSsylkaKey: Guid
	ZnacheniePustoiSsylkiType: String
	KliuchObieektaMetadannykhType: String
	KliuchObieektaMetadannykh: Stream
}
type CatalogIdentifikatoryObieektovMetadannykh {
	Key: Guid!
	DataVersion: String
	Description: String
	ParentKey: Guid
	DeletionMark: Boolean
	PoriadokKollektsii: Int16
	Imia: String
	Sinonim: String
	PolnoeImia: String
	PolnyiSinonim: String
	BezDannykh: Boolean
	ZnacheniePustoiSsylki: String
	KliuchObieektaMetadannykhBase64Data: Binary
	NovaiaSsylkaKey: Guid
	ZnacheniePustoiSsylkiType: String
	KliuchObieektaMetadannykhType: String
	KliuchObieektaMetadannykh: Stream
}
input DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	VPechatnykhFormakhTTPechatatRoznichnyeSummy: Boolean
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	UsloviiaProvedeniiaInventarizatsiiBase64Data: Binary
	KhoziaistvennaiaOperatsiiaKey: Guid
	NastroikiZapolneniiaBase64Data: Binary
	Sertifikaty: [DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatyRowTypeInput!]
	UsloviiaProvedeniiaInventarizatsii: [DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiRowTypeInput!]
	UsloviiaProvedeniiaInventarizatsiiType: String
	NastroikiZapolneniiaType: String
	InventoryTerms: Stream
	NastroikiZapolneniia: Stream
}
type DocumentSvodnaiaInventarizatsiiaTovarovNaSklade {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	VPechatnykhFormakhTTPechatatRoznichnyeSummy: Boolean
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	UsloviiaProvedeniiaInventarizatsiiBase64Data: Binary
	KhoziaistvennaiaOperatsiiaKey: Guid
	NastroikiZapolneniiaBase64Data: Binary
	Sertifikaty: [DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatyRowType!]
	UsloviiaProvedeniiaInventarizatsii: [DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiRowType!]
	UsloviiaProvedeniiaInventarizatsiiType: String
	NastroikiZapolneniiaType: String
	InventoryTerms: Stream
	NastroikiZapolneniia: Stream
}
input DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatyInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
	SummaUchet: Double
	Quantity: Double
	KolichestvoUchet: Double
}
type DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
	SummaUchet: Double
	Quantity: Double
	KolichestvoUchet: Double
}
input DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiInput {
	Key: Guid!
	LineNumber: Int64!
	VidSravneniia: String
	Znachenie: String
	ImiaPolia: String
	ZnachenieType: String
}
type DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii {
	Key: Guid!
	LineNumber: Int64!
	VidSravneniia: String
	Znachenie: String
	ImiaPolia: String
	ZnachenieType: String
}
input DocumentKorrektirovkaRealizatsiiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DokumentOsnovanieKey: Guid
	OrganizatsiiaKey: Guid
	DataVkhodiashchegoDokumenta: DateTime
	NomerVkhodiashchegoDokumenta: String
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	VosstanovitNDS: Boolean
	NDSVkliuchenVStoimost: Boolean
	SummaVkliuchaetNDS: Boolean
	OtvetstvennyiKey: Guid
	Comment: String
	SumOfDocument: Double
	SummaNDSDokumenta: Double
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	DepartmentKey: Guid
	ValiutaDokumentaKey: Guid
	PodrazdelenieKey: Guid
	NomerIspravleniia: Int64
	DataIspravleniia: DateTime
	NomerIskhodnogoDokumenta: String
	DataIskhodnogoDokumenta: DateTime
	NomerIspravleniiaIskhodnogoDokumenta: Int64
	DataIspravleniiaIskhodnogoDokumenta: DateTime
	KorrektirovatNDS: Boolean
	Goods: [DocumentKorrektirovkaRealizatsiiTovaryRowTypeInput!]
	Uslugi: [DocumentKorrektirovkaRealizatsiiUslugiRowTypeInput!]
}
type DocumentKorrektirovkaRealizatsii {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DokumentOsnovanieKey: Guid
	OrganizatsiiaKey: Guid
	DataVkhodiashchegoDokumenta: DateTime
	NomerVkhodiashchegoDokumenta: String
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	VosstanovitNDS: Boolean
	NDSVkliuchenVStoimost: Boolean
	SummaVkliuchaetNDS: Boolean
	OtvetstvennyiKey: Guid
	Comment: String
	SumOfDocument: Double
	SummaNDSDokumenta: Double
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	DepartmentKey: Guid
	ValiutaDokumentaKey: Guid
	PodrazdelenieKey: Guid
	NomerIspravleniia: Int64
	DataIspravleniia: DateTime
	NomerIskhodnogoDokumenta: String
	DataIskhodnogoDokumenta: DateTime
	NomerIspravleniiaIskhodnogoDokumenta: Int64
	DataIspravleniiaIskhodnogoDokumenta: DateTime
	KorrektirovatNDS: Boolean
	Goods: [DocumentKorrektirovkaRealizatsiiTovaryRowType!]
	Uslugi: [DocumentKorrektirovkaRealizatsiiUslugiRowType!]
}
input DocumentKorrektirovkaRealizatsiiTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	VesDoKorrektirovki: Double
	VesDoIzmeneniia: Double
	Weight: Double
	KolichestvoDoKorrektirovki: Int64
	KolichestvoDoIzmeneniia: Int64
	Quantity: Int64
	DokumentOprikhodovaniia: String
	DokumentPartii: String
	KachestvoKey: Guid
	ItemKey: Guid
	Pasport: String
	RazmerDoKorrektirovkiKey: Guid
	RazmerDoIzmeneniiaKey: Guid
	SizeKey: Guid
	SeriiaNomenklaturyDoKorrektirovkiKey: Guid
	SeriiaNomenklaturyDoIzmeneniiaKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDSDoKorrektirovki: String
	StavkaNDSDoIzmeneniia: String
	StavkaNDS: String
	StatusPartii: String
	SummaDoKorrektirovki: Double
	SummaDoIzmeneniia: Double
	Sum: Double
	SummaNDSDoKorrektirovki: Double
	SummaNDSDoIzmeneniia: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyDoKorrektirovkiKey: Guid
	KharakteristikaNomenklaturyDoIzmeneniiaKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TsenaDoKorrektirovki: Double
	TsenaDoIzmeneniia: Double
	Cost: Double
	DokumentOprikhodovaniiaType: String
	DokumentPartiiType: String
}
type DocumentKorrektirovkaRealizatsiiTovary {
	Key: Guid!
	LineNumber: Int64!
	VesDoKorrektirovki: Double
	VesDoIzmeneniia: Double
	Weight: Double
	KolichestvoDoKorrektirovki: Int64
	KolichestvoDoIzmeneniia: Int64
	Quantity: Int64
	DokumentOprikhodovaniia: String
	DokumentPartii: String
	KachestvoKey: Guid
	ItemKey: Guid
	Pasport: String
	RazmerDoKorrektirovkiKey: Guid
	RazmerDoIzmeneniiaKey: Guid
	SizeKey: Guid
	SeriiaNomenklaturyDoKorrektirovkiKey: Guid
	SeriiaNomenklaturyDoIzmeneniiaKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDSDoKorrektirovki: String
	StavkaNDSDoIzmeneniia: String
	StavkaNDS: String
	StatusPartii: String
	SummaDoKorrektirovki: Double
	SummaDoIzmeneniia: Double
	Sum: Double
	SummaNDSDoKorrektirovki: Double
	SummaNDSDoIzmeneniia: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyDoKorrektirovkiKey: Guid
	KharakteristikaNomenklaturyDoIzmeneniiaKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TsenaDoKorrektirovki: Double
	TsenaDoIzmeneniia: Double
	Cost: Double
	DokumentOprikhodovaniiaType: String
	DokumentPartiiType: String
}
input DocumentKorrektirovkaRealizatsiiUslugiInput {
	Key: Guid!
	LineNumber: Int64!
	KolichestvoDoKorrektirovki: Int64
	KolichestvoDoIzmeneniia: Int64
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDSDoKorrektirovki: String
	StavkaNDSDoIzmeneniia: String
	StavkaNDS: String
	SummaDoKorrektirovki: Double
	SummaDoIzmeneniia: Double
	Sum: Double
	SummaNDSDoKorrektirovki: Double
	SummaNDSDoIzmeneniia: Double
	SummaNDS: Double
	TsenaDoKorrektirovki: Double
	TsenaDoIzmeneniia: Double
	Cost: Double
}
type DocumentKorrektirovkaRealizatsiiUslugi {
	Key: Guid!
	LineNumber: Int64!
	KolichestvoDoKorrektirovki: Int64
	KolichestvoDoIzmeneniia: Int64
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDSDoKorrektirovki: String
	StavkaNDSDoIzmeneniia: String
	StavkaNDS: String
	SummaDoKorrektirovki: Double
	SummaDoIzmeneniia: Double
	Sum: Double
	SummaNDSDoKorrektirovki: Double
	SummaNDSDoIzmeneniia: Double
	SummaNDS: Double
	TsenaDoKorrektirovki: Double
	TsenaDoIzmeneniia: Double
	Cost: Double
}
input CatalogVidyDefektovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogVidyDefektov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input DocumentDoverennostInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DataDeistviia: DateTime
	DogovorKontragentaKey: Guid
	DolzhnostKey: Guid
	Comment: String
	KontragentKey: Guid
	NaPoluchenieOt: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PoDokumentu: String
	PodrazdelenieKey: Guid
	Sdelka: String
	StrukturnaiaEdinitsaKey: Guid
	FizLitsoKey: Guid
	Goods: [DocumentDoverennostTovaryRowTypeInput!]
	SdelkaType: String
}
type DocumentDoverennost {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DataDeistviia: DateTime
	DogovorKontragentaKey: Guid
	DolzhnostKey: Guid
	Comment: String
	KontragentKey: Guid
	NaPoluchenieOt: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PoDokumentu: String
	PodrazdelenieKey: Guid
	Sdelka: String
	StrukturnaiaEdinitsaKey: Guid
	FizLitsoKey: Guid
	Goods: [DocumentDoverennostTovaryRowType!]
	SdelkaType: String
}
input DocumentDoverennostTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaPoKlassifikatoruKey: Guid
	Quantity: Int64
	NaimenovanieTovara: String
	SizeKey: Guid
}
type DocumentDoverennostTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaPoKlassifikatoruKey: Guid
	Quantity: Int64
	NaimenovanieTovara: String
	SizeKey: Guid
}
input CatalogShablonyZapolneniiaKUInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	RuchnoeUkazaniePerioda: Boolean
	DataNachBudni: DateTime
	DataKonBudni: DateTime
	RuchnoeUkazanieSalonov: Boolean
	KUMin: Double
	KUMaks: Double
	PrazdnichnyeDni: [CatalogShablonyZapolneniiaKUPrazdnichnyeDniRowTypeInput!]
	KUNaNedeliu: [CatalogShablonyZapolneniiaKUKUNaNedeliuRowTypeInput!]
	Salony: [CatalogShablonyZapolneniiaKUSalonyRowTypeInput!]
}
type CatalogShablonyZapolneniiaKU {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	RuchnoeUkazaniePerioda: Boolean
	DataNachBudni: DateTime
	DataKonBudni: DateTime
	RuchnoeUkazanieSalonov: Boolean
	KUMin: Double
	KUMaks: Double
	PrazdnichnyeDni: [CatalogShablonyZapolneniiaKUPrazdnichnyeDniRowType!]
	KUNaNedeliu: [CatalogShablonyZapolneniiaKUKUNaNedeliuRowType!]
	Salony: [CatalogShablonyZapolneniiaKUSalonyRowType!]
}
input CatalogShablonyZapolneniiaKUPrazdnichnyeDniInput {
	Key: Guid!
	LineNumber: Int64!
	Den: DateTime
	KU: Double
	Opisanie: String
	TsvetFonaDliaKalendaria: String
	TsvetTekstaDliaKalendaria: String
}
type CatalogShablonyZapolneniiaKUPrazdnichnyeDni {
	Key: Guid!
	LineNumber: Int64!
	Den: DateTime
	KU: Double
	Opisanie: String
	TsvetFonaDliaKalendaria: String
	TsvetTekstaDliaKalendaria: String
}
input CatalogShablonyZapolneniiaKUKUNaNedeliuInput {
	Key: Guid!
	LineNumber: Int64!
	DenNedeli: Int16
	KU: Double
}
type CatalogShablonyZapolneniiaKUKUNaNedeliu {
	Key: Guid!
	LineNumber: Int64!
	DenNedeli: Int16
	KU: Double
}
input CatalogShablonyZapolneniiaKUSalonyInput {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
}
type CatalogShablonyZapolneniiaKUSalony {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
}
input DocumentPlanZapolneniiaVitrinInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	PlanovoeZapolnenieVitrin: [DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrinRowTypeInput!]
}
type DocumentPlanZapolneniiaVitrin {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	PlanovoeZapolnenieVitrin: [DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrinRowType!]
}
input DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrinInput {
	Key: Guid!
	LineNumber: Int64!
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	TypeKey: Guid
	AnalitikaTipaIzdeliiaKey: Guid
	PloshchadVykladki: Double
	NormativnaiaPloshchadVykladki: Double
	KolichestvoIzdelii: Int16
}
type DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin {
	Key: Guid!
	LineNumber: Int64!
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	TypeKey: Guid
	AnalitikaTipaIzdeliiaKey: Guid
	PloshchadVykladki: Double
	NormativnaiaPloshchadVykladki: Double
	KolichestvoIzdelii: Int16
}
input InstanceInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	Weight: Double
	DogovorPostavkiKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	Comment: String
	NaborKey: Guid
	NomerGTDKey: Guid
	NomerPasporta: String
	OsnovnoeIzobrazhenieKey: Guid
	SupplierKey: Guid
	SizeKey: Guid
	Sebestoimost: Double
	SeriinyiNomer: String
	Sertifikat: String
	SrokGodnosti: DateTime
	StranaProiskhozhdeniiaKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	ObshchiiVesVstavok: Double
	DokumentPostupleniiaKey: Guid
	TsenaPostupleniia: Double
	SummaPostupleniia: Double
	RetailCost: Double
}
type Instance {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	Weight: Double
	DogovorPostavkiKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	Comment: String
	NaborKey: Guid
	NomerGTDKey: Guid
	NomerPasporta: String
	OsnovnoeIzobrazhenieKey: Guid
	SupplierKey: Guid
	SizeKey: Guid
	Sebestoimost: Double
	SeriinyiNomer: String
	Sertifikat: String
	SrokGodnosti: DateTime
	StranaProiskhozhdeniiaKey: Guid
	TovarnaiaGruppaKey: Guid
	TovarnaiaKategoriiaKey: Guid
	TovarnaiaPozitsiiaKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	ObshchiiVesVstavok: Double
	DokumentPostupleniiaKey: Guid
	TsenaPostupleniia: Double
	SummaPostupleniia: Double
	RetailCost: Double
}
input ReturnToManufacturingInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	MaterialVProizvodstve: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatVesVstavok: Boolean
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	SobstvennoeProizvodstvo: Boolean
	ProizvodstvennyiUchastokKey: Guid
	OpisanieDefektov: String
	RuchnoiUchetVesaVstavok: Boolean
	Goods: [DocumentVozvratProduktsiiVProizvodstvoTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type ReturnToManufacturing {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	MaterialVProizvodstve: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatVesVstavok: Boolean
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	SobstvennoeProizvodstvo: Boolean
	ProizvodstvennyiUchastokKey: Guid
	OpisanieDefektov: String
	RuchnoiUchetVesaVstavok: Boolean
	Goods: [DocumentVozvratProduktsiiVProizvodstvoTovaryRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentVozvratProduktsiiVProizvodstvoTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	StoimostRabot: Double
	StoimostVstavok: Double
	StoimostMetalla: Double
	SummaNDSVstavok: Double
	SummaNDSMetalla: Double
	SummaNDSRabot: Double
	DefektKey: Guid
	VesVstavok: Double
}
type DocumentVozvratProduktsiiVProizvodstvoTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	StoimostRabot: Double
	StoimostVstavok: Double
	StoimostMetalla: Double
	SummaNDSVstavok: Double
	SummaNDSMetalla: Double
	SummaNDSRabot: Double
	DefektKey: Guid
	VesVstavok: Double
}
input CatalogNomeraGTDInput {
	Key: Guid!
	DataVersion: String
	Code: String
	DeletionMark: Boolean
	Comment: String
}
type CatalogNomeraGTD {
	Key: Guid!
	DataVersion: String
	Code: String
	DeletionMark: Boolean
	Comment: String
}
input CatalogNastroikiRabochegoMestaPolzovateliaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Nastroiki: [CatalogNastroikiRabochegoMestaPolzovateliaNastroikiRowTypeInput!]
}
type CatalogNastroikiRabochegoMestaPolzovatelia {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Nastroiki: [CatalogNastroikiRabochegoMestaPolzovateliaNastroikiRowType!]
}
input CatalogNastroikiRabochegoMestaPolzovateliaNastroikiInput {
	Key: Guid!
	LineNumber: Int64!
	Dostupnost: Boolean
	DostupnostPechati: Boolean
	DostupnostRedaktirovaniia: Boolean
	DostupnostVersionirovaniia: Boolean
	PutKElementuFormy: String
	TipObieekta: String
	ElementFormyRabochegoMesta: String
}
type CatalogNastroikiRabochegoMestaPolzovateliaNastroiki {
	Key: Guid!
	LineNumber: Int64!
	Dostupnost: Boolean
	DostupnostPechati: Boolean
	DostupnostRedaktirovaniia: Boolean
	DostupnostVersionirovaniia: Boolean
	PutKElementuFormy: String
	TipObieekta: String
	ElementFormyRabochegoMesta: String
}
input CatalogsmsShablonyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Soderzhanie: String
}
type CatalogsmsShablony {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Soderzhanie: String
}
input WriteOffInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	InventarizatsiiaTovarovNaSkladeKey: Guid
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	Osnovanie: String
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	TipDokumenta: String
	KhoziaistvennaiaOperatsiiaKey: Guid
	Goods: [DocumentSpisanieTovarovTovaryRowTypeInput!]
	Sertifikaty: [DocumentSpisanieTovarovSertifikatyRowTypeInput!]
	DokumentOsnovanieType: String
}
type WriteOff {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	InventarizatsiiaTovarovNaSkladeKey: Guid
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	Osnovanie: String
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	TipDokumenta: String
	KhoziaistvennaiaOperatsiiaKey: Guid
	Goods: [DocumentSpisanieTovarovTovaryRowType!]
	Sertifikaty: [DocumentSpisanieTovarovSertifikatyRowType!]
	DokumentOsnovanieType: String
}
input DocumentSpisanieTovarovTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentRezerva: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	SkidkaNatsenkaKey: Guid
	DokumentRezervaType: String
}
type DocumentSpisanieTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentRezerva: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	SkidkaNatsenkaKey: Guid
	DokumentRezervaType: String
}
input DocumentSpisanieTovarovSertifikatyInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
type DocumentSpisanieTovarovSertifikaty {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
input DocumentsmsSoobshchenieInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtorKey: Guid
	AvtotekstBulevo: Boolean
	Aktualnost: DateTime
	Comment: String
	KonetsPeriodaZapreta: DateTime
	NachaloOtpravki: DateTime
	NachaloPeriodaZapreta: DateTime
	NeOtpravliatSMS: Boolean
	NomerOtpravitelia: String
	RavnomernaiaRassylka: Boolean
	SpisokPoluchatelei: String
	Status: String
	StatusStrokoi: String
	TekstSoobshcheniia: String
	TipSoobshcheniia: String
	Transliteratsiia: Boolean
	FlagAktualnost: Boolean
	NastroikiZapolneniiaBase64Data: Binary
	Poluchateli: [DocumentsmsSoobshcheniePoluchateliRowTypeInput!]
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
type DocumentsmsSoobshchenie {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtorKey: Guid
	AvtotekstBulevo: Boolean
	Aktualnost: DateTime
	Comment: String
	KonetsPeriodaZapreta: DateTime
	NachaloOtpravki: DateTime
	NachaloPeriodaZapreta: DateTime
	NeOtpravliatSMS: Boolean
	NomerOtpravitelia: String
	RavnomernaiaRassylka: Boolean
	SpisokPoluchatelei: String
	Status: String
	StatusStrokoi: String
	TekstSoobshcheniia: String
	TipSoobshcheniia: String
	Transliteratsiia: Boolean
	FlagAktualnost: Boolean
	NastroikiZapolneniiaBase64Data: Binary
	Poluchateli: [DocumentsmsSoobshcheniePoluchateliRowType!]
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
input DocumentsmsSoobshcheniePoluchateliInput {
	Key: Guid!
	LineNumber: Int64!
	Kontragent: String
	NomerTelefona: String
	GUID: String
	TekstSoobshcheniia: String
	DataZaversheniia: DateTime
	Millisekunda: Int16
	Status: String
	OpisanieOshibki: String
	MemberCardKey: Guid
}
type DocumentsmsSoobshcheniePoluchateli {
	Key: Guid!
	LineNumber: Int64!
	Kontragent: String
	NomerTelefona: String
	GUID: String
	TekstSoobshcheniia: String
	DataZaversheniia: DateTime
	Millisekunda: Int16
	Status: String
	OpisanieOshibki: String
	MemberCardKey: Guid
}
input DocumentOplataOtPokupateliaPlatezhnoiKartoiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	DokumentOsnovanie: String
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	OtvetstvennyiKey: Guid
	DokumentSozdanVIuTD: Boolean
	KassaKKMKey: Guid
	Comment: String
	NumberKKT: Int16
	ProbitChekNaKKT: Boolean
	GungNumber: Int16
	ProtsentTorgovoiUstupki: Double
	SumOfDocument: Double
	SummaTorgovoiUstupki: Double
	Khesh: String
	Poslednie4: String
	KodRRN: String
	Identifikator: String
	TransactionId: String
	TipSistemyNalogooblozheniiaKey: Guid
	VidOplatyKey: Guid
	StavkaNDS: String
	SummaNDS: Double
	OperationType: String
	Pochta: String
	Telefon: String
	TypeOfMovingMoneyKey: Guid
	ProektKey: Guid
	NastroikaRMKKey: Guid
	BDataDokumenta: DateTime
	BIdentifikator: String
	BNomerVersii: String
	DokumentOsnovanieType: String
}
type DocumentOplataOtPokupateliaPlatezhnoiKartoi {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	DokumentOsnovanie: String
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	OtvetstvennyiKey: Guid
	DokumentSozdanVIuTD: Boolean
	KassaKKMKey: Guid
	Comment: String
	NumberKKT: Int16
	ProbitChekNaKKT: Boolean
	GungNumber: Int16
	ProtsentTorgovoiUstupki: Double
	SumOfDocument: Double
	SummaTorgovoiUstupki: Double
	Khesh: String
	Poslednie4: String
	KodRRN: String
	Identifikator: String
	TransactionId: String
	TipSistemyNalogooblozheniiaKey: Guid
	VidOplatyKey: Guid
	StavkaNDS: String
	SummaNDS: Double
	OperationType: String
	Pochta: String
	Telefon: String
	TypeOfMovingMoneyKey: Guid
	ProektKey: Guid
	NastroikaRMKKey: Guid
	BDataDokumenta: DateTime
	BIdentifikator: String
	BNomerVersii: String
	DokumentOsnovanieType: String
}
input CatalogDragotsennyeKamniInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidKamniaKey: Guid
	GruppaDefekta: Boolean
	GruppaTsveta: Boolean
	KratkoeNaimenovanie: String
	Rassev: Boolean
	RaschetTsenyZaVes: Boolean
	UchetVKaratakh: Boolean
	Tsvet: Boolean
	BIdentifikator: String
}
type CatalogDragotsennyeKamni {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidKamniaKey: Guid
	GruppaDefekta: Boolean
	GruppaTsveta: Boolean
	KratkoeNaimenovanie: String
	Rassev: Boolean
	RaschetTsenyZaVes: Boolean
	UchetVKaratakh: Boolean
	Tsvet: Boolean
	BIdentifikator: String
}
input CatalogKalendariPlanirovaniiaProdazhInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ShablonZapolneniiaKey: Guid
	NomerGodaPlanirovaniia: Int16
	NastroikiZapolneniiaBase64Data: Binary
	KUPoDniam: [CatalogKalendariPlanirovaniiaProdazhKUPoDniamRowTypeInput!]
	Salony: [CatalogKalendariPlanirovaniiaProdazhSalonyRowTypeInput!]
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
type CatalogKalendariPlanirovaniiaProdazh {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ShablonZapolneniiaKey: Guid
	NomerGodaPlanirovaniia: Int16
	NastroikiZapolneniiaBase64Data: Binary
	KUPoDniam: [CatalogKalendariPlanirovaniiaProdazhKUPoDniamRowType!]
	Salony: [CatalogKalendariPlanirovaniiaProdazhSalonyRowType!]
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
input CatalogKalendariPlanirovaniiaProdazhKUPoDniamInput {
	Key: Guid!
	LineNumber: Int64!
	DenGoda: DateTime
	KU: Double
}
type CatalogKalendariPlanirovaniiaProdazhKUPoDniam {
	Key: Guid!
	LineNumber: Int64!
	DenGoda: DateTime
	KU: Double
}
input CatalogKalendariPlanirovaniiaProdazhSalonyInput {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
}
type CatalogKalendariPlanirovaniiaProdazhSalony {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
}
input CatalogKontaktnyeLitsaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidKontaktnogoLitsa: String
	DataRozhdeniia: DateTime
	Dolzhnost: String
	Imia: String
	KolichestvoDneiDoNapominaniia: Int16
	KontragentDliaOgranicheniiaPravDostupaKey: Guid
	NapominatODneRozhdeniia: Boolean
	ObieektVladelets: String
	Opisanie: String
	Otchestvo: String
	PolzovatelDliaOgranicheniiaPravDostupaKey: Guid
	RolKey: Guid
	Familiia: String
	ObieektVladeletsType: String
}
type CatalogKontaktnyeLitsa {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidKontaktnogoLitsa: String
	DataRozhdeniia: DateTime
	Dolzhnost: String
	Imia: String
	KolichestvoDneiDoNapominaniia: Int16
	KontragentDliaOgranicheniiaPravDostupaKey: Guid
	NapominatODneRozhdeniia: Boolean
	ObieektVladelets: String
	Opisanie: String
	Otchestvo: String
	PolzovatelDliaOgranicheniiaPravDostupaKey: Guid
	RolKey: Guid
	Familiia: String
	ObieektVladeletsType: String
}
input CatalogFizicheskieLitsaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	DataRozhdeniia: DateTime
	Comment: String
	OsnovnoeIzobrazhenieKey: Guid
	Pol: String
	Sotrudnik: Boolean
	MagazinKey: Guid
	Kurer: Boolean
	NastroikaDliaRabotyKey: Guid
	INN: String
	BIdentifikator: String
	BNomerVersii: String
}
type CatalogFizicheskieLitsa {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	DataRozhdeniia: DateTime
	Comment: String
	OsnovnoeIzobrazhenieKey: Guid
	Pol: String
	Sotrudnik: Boolean
	MagazinKey: Guid
	Kurer: Boolean
	NastroikaDliaRabotyKey: Guid
	INN: String
	BIdentifikator: String
	BNomerVersii: String
}
input CatalogTipovyeAnketyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Adresnaia: Boolean
	VidSpravochnikaDliaZagruzki: String
	Vstuplenie: String
	ZagruzhatObieekty: Boolean
	MaketAnketyBase64Data: Binary
	NaimenovanieAnkety: String
	VoprosyAnkety: [CatalogTipovyeAnketyVoprosyAnketyRowTypeInput!]
	MaketAnketyType: String
	MaketAnkety: Stream
}
type CatalogTipovyeAnkety {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Adresnaia: Boolean
	VidSpravochnikaDliaZagruzki: String
	Vstuplenie: String
	ZagruzhatObieekty: Boolean
	MaketAnketyBase64Data: Binary
	NaimenovanieAnkety: String
	VoprosyAnkety: [CatalogTipovyeAnketyVoprosyAnketyRowType!]
	MaketAnketyType: String
	MaketAnkety: Stream
}
input CatalogTipovyeAnketyVoprosyAnketyInput {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	RazdelKey: Guid
}
type CatalogTipovyeAnketyVoprosyAnkety {
	Key: Guid!
	LineNumber: Int64!
	VoprosKey: Guid
	RazdelKey: Guid
}
input DocumentNachislenieSpisanieBonusovInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtvetstvennyiKey: Guid
	NastroikiZapolneniiaBase64Data: Binary
	DataNachisleniiaBonusov: DateTime
	VremmennyeBonusy: Boolean
	DataSpisaniiaBonusov: DateTime
	DokumentSozdanVIuTD: Boolean
	DiskontnyeKarty: [DocumentNachislenieSpisanieBonusovDiskontnyeKartyRowTypeInput!]
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
type DocumentNachislenieSpisanieBonusov {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtvetstvennyiKey: Guid
	NastroikiZapolneniiaBase64Data: Binary
	DataNachisleniiaBonusov: DateTime
	VremmennyeBonusy: Boolean
	DataSpisaniiaBonusov: DateTime
	DokumentSozdanVIuTD: Boolean
	DiskontnyeKarty: [DocumentNachislenieSpisanieBonusovDiskontnyeKartyRowType!]
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
input DocumentNachislenieSpisanieBonusovDiskontnyeKartyInput {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
	SummaBonusov: Int64
	TekushchaiaSummaBonusov: Int64
}
type DocumentNachislenieSpisanieBonusovDiskontnyeKarty {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
	SummaBonusov: Int64
	TekushchaiaSummaBonusov: Int64
}
input TypeInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	EstRazmer: Boolean
	NeUchityvatVes: Boolean
	ProtsentTekhPoter: Double
	BIdentifikator: String
}
type Type {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	EstRazmer: Boolean
	NeUchityvatVes: Boolean
	ProtsentTekhPoter: Double
	BIdentifikator: String
}
input CatalogfmKodyVidovDokumentovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: Int16
	DeletionMark: Boolean
	KodTipaDokumenta: String
	TipDokumenta: String
	TipDokumentaVProgramme: String
}
type CatalogfmKodyVidovDokumentov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: Int16
	DeletionMark: Boolean
	KodTipaDokumenta: String
	TipDokumenta: String
	TipDokumentaVProgramme: String
}
input DocumentPlatezhnoeTrebovaniePoluchennoeInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPlatezha: String
	DataVkhodiashchegoDokumenta: DateTime
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	ZaiavkaNaRaskhodovanieSredstvKey: Guid
	Comment: String
	KontragentKey: Guid
	NaznacheniePlatezha: String
	NomerVkhodiashchegoDokumenta: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	OcherednostPlatezha: Int16
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	ExtendedPayments: [DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhaRowTypeInput!]
	RekvizityKontragenta: [DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentaRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentPlatezhnoeTrebovaniePoluchennoe {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPlatezha: String
	DataVkhodiashchegoDokumenta: DateTime
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	ZaiavkaNaRaskhodovanieSredstvKey: Guid
	Comment: String
	KontragentKey: Guid
	NaznacheniePlatezha: String
	NomerVkhodiashchegoDokumenta: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	OcherednostPlatezha: Int16
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	ExtendedPayments: [DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhaRowType!]
	RekvizityKontragenta: [DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentaRowType!]
	DokumentOsnovanieType: String
}
input DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	NomerPlatezha: Int16
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	NomerPlatezha: Int16
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	KontragentKey: Guid
	OtborDataKon: DateTime
	OtborDataNach: DateTime
	OtborKontragent: Boolean
	OtborOtvetstvennyi: Boolean
	OtborTsFO: Boolean
	OtvetstvennyiKey: Guid
	OtvetstvennyiPostuplenieKey: Guid
	Sostoianie: String
	TipDokumenta: String
	TsFOKey: Guid
	PlaniruemyePostupleniiaDS: [DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSRowTypeInput!]
}
type DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	KontragentKey: Guid
	OtborDataKon: DateTime
	OtborDataNach: DateTime
	OtborKontragent: Boolean
	OtborOtvetstvennyi: Boolean
	OtborTsFO: Boolean
	OtvetstvennyiKey: Guid
	OtvetstvennyiPostuplenieKey: Guid
	Sostoianie: String
	TipDokumenta: String
	TsFOKey: Guid
	PlaniruemyePostupleniiaDS: [DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSRowType!]
}
input DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSInput {
	Key: Guid!
	LineNumber: Int64!
	ValiutaPostuplenieKey: Guid
	DokumentPlanirovaniiaKey: Guid
	KontragentKey: Guid
	OstatokPostuplenie: Double
	OstatokRazmeshchenie: Double
	OtvetstvennyiKey: Guid
}
type DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS {
	Key: Guid!
	LineNumber: Int64!
	ValiutaPostuplenieKey: Guid
	DokumentPlanirovaniiaKey: Guid
	KontragentKey: Guid
	OstatokPostuplenie: Double
	OstatokRazmeshchenie: Double
	OtvetstvennyiKey: Guid
}
input CatalogRazdelyAnketyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
}
type CatalogRazdelyAnkety {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
}
input DocumentOtchetPoFinMonitoringuInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	NomerSoobshcheniia: Int64
	DataSoobshcheniia: DateTime
	VidDokumenta: String
	DokumentOsnovanieKey: Guid
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	KartochkaKontragentaKey: Guid
	KartochkaOrganizatsiiKey: Guid
	DannyePoOplate: Boolean
	RaschetnyiSchetOrganizatsiiKey: Guid
	RaschetnyiSchetKontragentaKey: Guid
	DannyePoVozvratam: Boolean
	NomerSoobshcheniiaVTecheniiDnia: String
	DokumentyFinMonitoringa: [DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringaRowTypeInput!]
	DannyeDokumenta: [DocumentOtchetPoFinMonitoringuDannyeDokumentaRowTypeInput!]
}
type DocumentOtchetPoFinMonitoringu {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	NomerSoobshcheniia: Int64
	DataSoobshcheniia: DateTime
	VidDokumenta: String
	DokumentOsnovanieKey: Guid
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	KartochkaKontragentaKey: Guid
	KartochkaOrganizatsiiKey: Guid
	DannyePoOplate: Boolean
	RaschetnyiSchetOrganizatsiiKey: Guid
	RaschetnyiSchetKontragentaKey: Guid
	DannyePoVozvratam: Boolean
	NomerSoobshcheniiaVTecheniiDnia: String
	DokumentyFinMonitoringa: [DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringaRowType!]
	DannyeDokumenta: [DocumentOtchetPoFinMonitoringuDannyeDokumentaRowType!]
}
input DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringaInput {
	Key: Guid!
	LineNumber: Int64!
	DokumentUcheta: String
	SummaOtgruzki: Double
	SummaOplaty: Double
	Sum: Double
	SummaVozvrata: Double
	Comment: String
	KodVidaDokumentaKey: Guid
	DokumentUchetaType: String
}
type DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa {
	Key: Guid!
	LineNumber: Int64!
	DokumentUcheta: String
	SummaOtgruzki: Double
	SummaOplaty: Double
	Sum: Double
	SummaVozvrata: Double
	Comment: String
	KodVidaDokumentaKey: Guid
	DokumentUchetaType: String
}
input DocumentOtchetPoFinMonitoringuDannyeDokumentaInput {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
type DocumentOtchetPoFinMonitoringuDannyeDokumenta {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
input CatalogKliuchiAnalitikiUchetaNomenklaturyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	TipNomenklatury: String
	ProbeKey: Guid
	TypeKey: Guid
	SizeKey: Guid
	KamenKey: Guid
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	GruppaTsvetaKey: Guid
	GruppaDefektaKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	MID: String
}
type CatalogKliuchiAnalitikiUchetaNomenklatury {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	TipNomenklatury: String
	ProbeKey: Guid
	TypeKey: Guid
	SizeKey: Guid
	KamenKey: Guid
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	GruppaTsvetaKey: Guid
	GruppaDefektaKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	MID: String
}
input CatalogVersiiFailovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	AvtorKey: Guid
	DataModifikatsiiUniversalnaia: DateTime
	DataModifikatsiiFaila: DateTime
	DataSozdaniia: DateTime
	Zashifrovan: Boolean
	IndeksKartinki: Int64
	Comment: String
	NomerVersii: Int64
	PodpisanEP: Boolean
	PolnoeNaimenovanie: String
	PutKFailu: String
	Razmer: Int64
	Rasshirenie: String
	RoditelskaiaVersiiaKey: Guid
	StatusIzvlecheniiaTeksta: String
	TekstKhranilishcheBase64Data: Binary
	FailKhranilishcheBase64Data: Binary
	TipKhraneniiaFaila: String
	TomKey: Guid
	ElektronnyePodpisi: [CatalogVersiiFailovElektronnyePodpisiRowTypeInput!]
	TekstKhranilishcheType: String
	FailKhranilishcheType: String
	TekstKhranilishche: Stream
	FailKhranilishche: Stream
}
type CatalogVersiiFailov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	AvtorKey: Guid
	DataModifikatsiiUniversalnaia: DateTime
	DataModifikatsiiFaila: DateTime
	DataSozdaniia: DateTime
	Zashifrovan: Boolean
	IndeksKartinki: Int64
	Comment: String
	NomerVersii: Int64
	PodpisanEP: Boolean
	PolnoeNaimenovanie: String
	PutKFailu: String
	Razmer: Int64
	Rasshirenie: String
	RoditelskaiaVersiiaKey: Guid
	StatusIzvlecheniiaTeksta: String
	TekstKhranilishcheBase64Data: Binary
	FailKhranilishcheBase64Data: Binary
	TipKhraneniiaFaila: String
	TomKey: Guid
	ElektronnyePodpisi: [CatalogVersiiFailovElektronnyePodpisiRowType!]
	TekstKhranilishcheType: String
	FailKhranilishcheType: String
	TekstKhranilishche: Stream
	FailKhranilishche: Stream
}
input CatalogVersiiFailovElektronnyePodpisiInput {
	Key: Guid!
	LineNumber: Int64!
	DataPodpisi: DateTime
	ImiaFailaPodpisi: String
	Comment: String
	KomuVydanSertifikat: String
	Otpechatok: String
	PodpisBase64Data: Binary
	UstanovivshiiPodpisKey: Guid
	SertifikatBase64Data: Binary
	PodpisType: String
	SertifikatType: String
	Podpis: Stream
	Sertifikat: Stream
}
type CatalogVersiiFailovElektronnyePodpisi {
	Key: Guid!
	LineNumber: Int64!
	DataPodpisi: DateTime
	ImiaFailaPodpisi: String
	Comment: String
	KomuVydanSertifikat: String
	Otpechatok: String
	PodpisBase64Data: Binary
	UstanovivshiiPodpisKey: Guid
	SertifikatBase64Data: Binary
	PodpisType: String
	SertifikatType: String
	Podpis: Stream
	Sertifikat: Stream
}
input DocumentUstanovkaTsenNomenklaturyInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	VybGruppaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	Informatsiia: String
	Comment: String
	NeProvoditNulevyeZnacheniia: Boolean
	OtvetstvennyiKey: Guid
	ParametryOtboraBase64Data: Binary
	RegistratsiiaTsenPoSegmentamNomenklatury: Boolean
	NastroikiZapolneniiaBase64Data: Binary
	TipyTsen: [DocumentUstanovkaTsenNomenklaturyTipyTsenRowTypeInput!]
	Goods: [DocumentUstanovkaTsenNomenklaturyTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	ParametryOtboraType: String
	NastroikiZapolneniiaType: String
	ParametryOtbora: Stream
	NastroikiZapolneniia: Stream
}
type DocumentUstanovkaTsenNomenklatury {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	VybGruppaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	Informatsiia: String
	Comment: String
	NeProvoditNulevyeZnacheniia: Boolean
	OtvetstvennyiKey: Guid
	ParametryOtboraBase64Data: Binary
	RegistratsiiaTsenPoSegmentamNomenklatury: Boolean
	NastroikiZapolneniiaBase64Data: Binary
	TipyTsen: [DocumentUstanovkaTsenNomenklaturyTipyTsenRowType!]
	Goods: [DocumentUstanovkaTsenNomenklaturyTovaryRowType!]
	DokumentOsnovanieType: String
	ParametryOtboraType: String
	NastroikiZapolneniiaType: String
	ParametryOtbora: Stream
	NastroikiZapolneniia: Stream
}
input DocumentUstanovkaTsenNomenklaturyTipyTsenInput {
	Key: Guid!
	LineNumber: Int64!
	TipTsenKey: Guid
}
type DocumentUstanovkaTsenNomenklaturyTipyTsen {
	Key: Guid!
	LineNumber: Int64!
	TipTsenKey: Guid
}
input DocumentUstanovkaTsenNomenklaturyTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	ValiutaKey: Guid
	IndeksStrokiTablitsyTsen: Int64
	ItemKey: Guid
	ProtsentSkidkiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	TipTsenKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SegmentNomenklaturyKey: Guid
}
type DocumentUstanovkaTsenNomenklaturyTovary {
	Key: Guid!
	LineNumber: Int64!
	ValiutaKey: Guid
	IndeksStrokiTablitsyTsen: Int64
	ItemKey: Guid
	ProtsentSkidkiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	TipTsenKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SegmentNomenklaturyKey: Guid
}
input DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	ZaiavkaNaRaskhodovanieSredstvKey: Guid
	Comment: String
	KontragentKey: Guid
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	RaschetnyiDokument: String
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	NomerVkhodiashchegoDokumenta: String
	DataVkhodiashchegoDokumenta: DateTime
	PodrazdelenieKey: Guid
	ExtendedPayments: [DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhaRowTypeInput!]
	RekvizityKontragenta: [DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentaRowTypeInput!]
	DokumentOsnovanieType: String
	RaschetnyiDokumentType: String
}
type DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	ZaiavkaNaRaskhodovanieSredstvKey: Guid
	Comment: String
	KontragentKey: Guid
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	RaschetnyiDokument: String
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	NomerVkhodiashchegoDokumenta: String
	DataVkhodiashchegoDokumenta: DateTime
	PodrazdelenieKey: Guid
	ExtendedPayments: [DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhaRowType!]
	RekvizityKontragenta: [DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentaRowType!]
	DokumentOsnovanieType: String
	RaschetnyiDokumentType: String
}
input DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentPreiskurantNaSkupkuInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	NeProvoditNulevyeZnacheniia: Boolean
	OtvetstvennyiKey: Guid
	IspolzovatPriObmene: Boolean
	Comment: String
	DokumentSozdanVIuTD: Boolean
	DepartmentKey: Guid
	Proby: [DocumentPreiskurantNaSkupkuProbyRowTypeInput!]
}
type DocumentPreiskurantNaSkupku {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	NeProvoditNulevyeZnacheniia: Boolean
	OtvetstvennyiKey: Guid
	IspolzovatPriObmene: Boolean
	Comment: String
	DokumentSozdanVIuTD: Boolean
	DepartmentKey: Guid
	Proby: [DocumentPreiskurantNaSkupkuProbyRowType!]
}
input DocumentPreiskurantNaSkupkuProbyInput {
	Key: Guid!
	LineNumber: Int64!
	ProbeKey: Guid
	TsenaZaGramm: Double
}
type DocumentPreiskurantNaSkupkuProby {
	Key: Guid!
	LineNumber: Int64!
	ProbeKey: Guid
	TsenaZaGramm: Double
}
input DocumentPeredachaMaterialovVProizvodstvoInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	PoluchaemaiaLigaturaKey: Guid
	PoluchaemyiVes: Double
	ProektKey: Guid
	SobstvennoeProizvodstvo: Boolean
	DepartmentKey: Guid
	TipDokumenta: String
	KhoziaistvennaiaOperatsiiaKey: Guid
	ProizvodstvennyiUchastokKey: Guid
	ZagruzhenIzUIuP: Boolean
	Goods: [DocumentPeredachaMaterialovVProizvodstvoTovaryRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentPeredachaMaterialovVProizvodstvo {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	PoluchaemaiaLigaturaKey: Guid
	PoluchaemyiVes: Double
	ProektKey: Guid
	SobstvennoeProizvodstvo: Boolean
	DepartmentKey: Guid
	TipDokumenta: String
	KhoziaistvennaiaOperatsiiaKey: Guid
	ProizvodstvennyiUchastokKey: Guid
	ZagruzhenIzUIuP: Boolean
	Goods: [DocumentPeredachaMaterialovVProizvodstvoTovaryRowType!]
	DokumentOsnovanieType: String
}
input DocumentPeredachaMaterialovVProizvodstvoTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
type DocumentPeredachaMaterialovVProizvodstvoTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
input DocumentVnutrenniiZakazInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtoRazmeshchenie: Boolean
	AvtoRezervirovanie: Boolean
	VidZakaza: String
	VremiaNapominaniia: DateTime
	DataOtgruzki: DateTime
	DokumentOsnovanie: String
	Zakazchik: String
	Comment: String
	NapomnitOSobytii: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TipDokumenta: String
	Goods: [DocumentVnutrenniiZakazTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	ZakazchikType: String
}
type DocumentVnutrenniiZakaz {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtoRazmeshchenie: Boolean
	AvtoRezervirovanie: Boolean
	VidZakaza: String
	VremiaNapominaniia: DateTime
	DataOtgruzki: DateTime
	DokumentOsnovanie: String
	Zakazchik: String
	Comment: String
	NapomnitOSobytii: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TipDokumenta: String
	Goods: [DocumentVnutrenniiZakazTovaryRowType!]
	DokumentOsnovanieType: String
	ZakazchikType: String
}
input DocumentVnutrenniiZakazTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	RazmeshchenieKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
type DocumentVnutrenniiZakazTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	RazmeshchenieKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
input CatalogKhranilishcheDopolnitelnoiInformatsiiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	VidDannykh: String
	ZnachenieRazdeleniiaDostupaKey: Guid
	IDFailaPochtovogoPisma: String
	ImiaFaila: String
	Obieekt: String
	KhranilishcheBase64Data: Binary
	VygruzhatNaSait: Boolean
	ObieektType: String
	KhranilishcheType: String
	Khranilishche: Stream
}
type CatalogKhranilishcheDopolnitelnoiInformatsii {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	VidDannykh: String
	ZnachenieRazdeleniiaDostupaKey: Guid
	IDFailaPochtovogoPisma: String
	ImiaFaila: String
	Obieekt: String
	KhranilishcheBase64Data: Binary
	VygruzhatNaSait: Boolean
	ObieektType: String
	KhranilishcheType: String
	Khranilishche: Stream
}
input CatalogDopolnitelnyeVneshnieObrabotkiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	VidObrabotki: String
	Comment: String
	KommentariiKFailuIstochniku: String
	KhranilishcheVneshneiObrabotkiBase64Data: Binary
	BezopasnyiRezhim: Boolean
	Versiia: String
	ImiaObieekta: String
	ImiaFaila: String
	Informatsiia: String
	IspolzovatDliaFormyObieekta: Boolean
	IspolzovatDliaFormySpiska: Boolean
	Publikatsiia: String
	KhranilishcheNastroekBase64Data: Binary
	OtvetstvennyiKey: Guid
	IspolzuetKhranilishcheVariantov: Boolean
	Prinadlezhnost: [CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnostRowTypeInput!]
	Komandy: [CatalogDopolnitelnyeVneshnieObrabotkiKomandyRowTypeInput!]
	Razdely: [CatalogDopolnitelnyeVneshnieObrabotkiRazdelyRowTypeInput!]
	Naznachenie: [CatalogDopolnitelnyeVneshnieObrabotkiNaznachenieRowTypeInput!]
	Razresheniia: [CatalogDopolnitelnyeVneshnieObrabotkiRazresheniiaRowTypeInput!]
	KhranilishcheVneshneiObrabotkiType: String
	KhranilishcheNastroekType: String
	KhranilishcheVneshneiObrabotki: Stream
	KhranilishcheNastroek: Stream
}
type CatalogDopolnitelnyeVneshnieObrabotki {
	Key: Guid!
	DataVersion: String
	Description: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	VidObrabotki: String
	Comment: String
	KommentariiKFailuIstochniku: String
	KhranilishcheVneshneiObrabotkiBase64Data: Binary
	BezopasnyiRezhim: Boolean
	Versiia: String
	ImiaObieekta: String
	ImiaFaila: String
	Informatsiia: String
	IspolzovatDliaFormyObieekta: Boolean
	IspolzovatDliaFormySpiska: Boolean
	Publikatsiia: String
	KhranilishcheNastroekBase64Data: Binary
	OtvetstvennyiKey: Guid
	IspolzuetKhranilishcheVariantov: Boolean
	Prinadlezhnost: [CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnostRowType!]
	Komandy: [CatalogDopolnitelnyeVneshnieObrabotkiKomandyRowType!]
	Razdely: [CatalogDopolnitelnyeVneshnieObrabotkiRazdelyRowType!]
	Naznachenie: [CatalogDopolnitelnyeVneshnieObrabotkiNaznachenieRowType!]
	Razresheniia: [CatalogDopolnitelnyeVneshnieObrabotkiRazresheniiaRowType!]
	KhranilishcheVneshneiObrabotkiType: String
	KhranilishcheNastroekType: String
	KhranilishcheVneshneiObrabotki: Stream
	KhranilishcheNastroek: Stream
}
input CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnostInput {
	Key: Guid!
	LineNumber: Int64!
	DopolnitelnyeParametryObrabotkiBase64Data: Binary
	PredstavlenieKnopki: String
	PredstavlenieNastroekObrabotki: String
	PredstavlenieObieekta: String
	SsylkaObieekta: String
	TablichnaiaChastImia: String
	TablichnaiaChastPredstavlenie: String
	KhranilishcheVneshneiObrabotkiBase64Data: Binary
	DopolnitelnyeParametryObrabotkiType: String
	SsylkaObieektaType: String
	KhranilishcheVneshneiObrabotkiType: String
	DopolnitelnyeParametryObrabotki: Stream
	KhranilishcheVneshneiObrabotki: Stream
}
type CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost {
	Key: Guid!
	LineNumber: Int64!
	DopolnitelnyeParametryObrabotkiBase64Data: Binary
	PredstavlenieKnopki: String
	PredstavlenieNastroekObrabotki: String
	PredstavlenieObieekta: String
	SsylkaObieekta: String
	TablichnaiaChastImia: String
	TablichnaiaChastPredstavlenie: String
	KhranilishcheVneshneiObrabotkiBase64Data: Binary
	DopolnitelnyeParametryObrabotkiType: String
	SsylkaObieektaType: String
	KhranilishcheVneshneiObrabotkiType: String
	DopolnitelnyeParametryObrabotki: Stream
	KhranilishcheVneshneiObrabotki: Stream
}
input CatalogDopolnitelnyeVneshnieObrabotkiKomandyInput {
	Key: Guid!
	LineNumber: Int64!
	Identifikator: String
	VariantZapuska: String
	Predstavlenie: String
	PokazyvatOpoveshchenie: Boolean
	Modifikator: String
	ReglamentnoeZadanieGUID: Guid
	Skryt: Boolean
	ZameniaemyeKomandy: String
}
type CatalogDopolnitelnyeVneshnieObrabotkiKomandy {
	Key: Guid!
	LineNumber: Int64!
	Identifikator: String
	VariantZapuska: String
	Predstavlenie: String
	PokazyvatOpoveshchenie: Boolean
	Modifikator: String
	ReglamentnoeZadanieGUID: Guid
	Skryt: Boolean
	ZameniaemyeKomandy: String
}
input CatalogDopolnitelnyeVneshnieObrabotkiRazdelyInput {
	Key: Guid!
	LineNumber: Int64!
	RazdelKey: Guid
	UdalitImiaRazdela: String
}
type CatalogDopolnitelnyeVneshnieObrabotkiRazdely {
	Key: Guid!
	LineNumber: Int64!
	RazdelKey: Guid
	UdalitImiaRazdela: String
}
input CatalogDopolnitelnyeVneshnieObrabotkiNaznachenieInput {
	Key: Guid!
	LineNumber: Int64!
	ObieektNaznacheniiaKey: Guid
	UdalitPolnoeImiaObieektaMetadannykh: String
}
type CatalogDopolnitelnyeVneshnieObrabotkiNaznachenie {
	Key: Guid!
	LineNumber: Int64!
	ObieektNaznacheniiaKey: Guid
	UdalitPolnoeImiaObieektaMetadannykh: String
}
input CatalogDopolnitelnyeVneshnieObrabotkiRazresheniiaInput {
	Key: Guid!
	LineNumber: Int64!
	VidRazresheniia: String
	ParametryBase64Data: Binary
	ParametryType: String
	Parametry: Stream
}
type CatalogDopolnitelnyeVneshnieObrabotkiRazresheniia {
	Key: Guid!
	LineNumber: Int64!
	VidRazresheniia: String
	ParametryBase64Data: Binary
	ParametryType: String
	Parametry: Stream
}
input CatalogGruppyPolzovateleiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	AdministratorGruppyKey: Guid
	PolzovateliGruppy: [CatalogGruppyPolzovateleiPolzovateliGruppyRowTypeInput!]
}
type CatalogGruppyPolzovatelei {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	AdministratorGruppyKey: Guid
	PolzovateliGruppy: [CatalogGruppyPolzovateleiPolzovateliGruppyRowType!]
}
input CatalogGruppyPolzovateleiPolzovateliGruppyInput {
	Key: Guid!
	LineNumber: Int64!
	PolzovatelKey: Guid
}
type CatalogGruppyPolzovateleiPolzovateliGruppy {
	Key: Guid!
	LineNumber: Int64!
	PolzovatelKey: Guid
}
input DocumentJournalZakazyKlientovInput {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	DataOplaty: DateTime
	DataOtgruzki: DateTime
	InformatsiiaKey: Guid
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sum: Double
	RefType: String!
}
type DocumentJournalZakazyKlientov {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	DataOplaty: DateTime
	DataOtgruzki: DateTime
	InformatsiiaKey: Guid
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sum: Double
	RefType: String!
}
input DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetOrganizatsiiKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	OpisanieDefektov: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	PostavshchikuVystavliaetsiaSchetFakturaNaVozvrat: Boolean
	Goods: [DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetOrganizatsiiKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	OpisanieDefektov: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	PostavshchikuVystavliaetsiaSchetFakturaNaVozvrat: Boolean
	Goods: [DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovaryRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DefektKey: Guid
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	DokumentPostupleniiaType: String
}
type DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DefektKey: Guid
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	DokumentPostupleniiaType: String
}
input DocumentZaiavkaNaPeremeshchenieTovarovInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	OperationType: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	SkladOtpravitelKey: Guid
	SkladPoluchatelKey: Guid
	SumOfDocument: Double
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	SposobDostavkiKey: Guid
	IspolzovatRezhimSverki: Boolean
	Goods: [DocumentZaiavkaNaPeremeshchenieTovarovTovaryRowTypeInput!]
}
type DocumentZaiavkaNaPeremeshchenieTovarov {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	OperationType: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	SkladOtpravitelKey: Guid
	SkladPoluchatelKey: Guid
	SumOfDocument: Double
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	SposobDostavkiKey: Guid
	IspolzovatRezhimSverki: Boolean
	Goods: [DocumentZaiavkaNaPeremeshchenieTovarovTovaryRowType!]
}
input DocumentZaiavkaNaPeremeshchenieTovarovTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StoimostBezNDS: Double
	StoimostSNDS: Double
	Sum: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaVRoznitseGr: Double
	TsenaPostupleniia: Double
}
type DocumentZaiavkaNaPeremeshchenieTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StoimostBezNDS: Double
	StoimostSNDS: Double
	Sum: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaVRoznitseGr: Double
	TsenaPostupleniia: Double
}
input CatalogUsloviiaProdazhInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogUsloviiaProdazh {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input DocumentVvodNachalnykhOstatkovPoFinMonitoringuInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	TipDokumentovUcheta: Boolean
	NastroikiZapolneniiaBase64Data: Binary
	Dogovora: [DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoraRowTypeInput!]
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
type DocumentVvodNachalnykhOstatkovPoFinMonitoringu {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	TipDokumentovUcheta: Boolean
	NastroikiZapolneniiaBase64Data: Binary
	Dogovora: [DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoraRowType!]
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
input DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoraInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	ItogovaiaStroka: Boolean
	KontragentKey: Guid
	RuchnaiaKorrektirovka: Boolean
	SummaOplaty: Double
	SummaOtgruzki: Double
	SummaVozvrata: Double
	DokumentUchetaType: String
}
type DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentUcheta: String
	ItogovaiaStroka: Boolean
	KontragentKey: Guid
	RuchnaiaKorrektirovka: Boolean
	SummaOplaty: Double
	SummaOtgruzki: Double
	SummaVozvrata: Double
	DokumentUchetaType: String
}
input CatalogOrganizatsiiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	DataNachalaVedeniiaFinMonitoringa: DateTime
	INN: String
	IspolzovatFinMonitoring: Boolean
	KodIMNS: String
	KodPoOKATO: String
	KodPoOKVED: String
	KodPoOKPO: String
	KodPoOKTMO: String
	KPP: String
	NaimenovanieIFNS: String
	NaimenovaniePlatelshchikaPriPerechisleniiNalogov: String
	NaimenovaniePolnoe: String
	OGRN: String
	OsnovnoiBankovskiiSchetKey: Guid
	Prefiks: String
	SvidetelstvoDataVydachi: DateTime
	SvidetelstvoSeriiaNomer: String
	UchetnyiNomer: String
	IurFizLitso: String
	KomissiiaKakKupliaProdazha: Boolean
	IDKorr: String
	DataPriniatiiaKDeistviiuPisma50: DateTime
}
type CatalogOrganizatsii {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	DataNachalaVedeniiaFinMonitoringa: DateTime
	INN: String
	IspolzovatFinMonitoring: Boolean
	KodIMNS: String
	KodPoOKATO: String
	KodPoOKVED: String
	KodPoOKPO: String
	KodPoOKTMO: String
	KPP: String
	NaimenovanieIFNS: String
	NaimenovaniePlatelshchikaPriPerechisleniiNalogov: String
	NaimenovaniePolnoe: String
	OGRN: String
	OsnovnoiBankovskiiSchetKey: Guid
	Prefiks: String
	SvidetelstvoDataVydachi: DateTime
	SvidetelstvoSeriiaNomer: String
	UchetnyiNomer: String
	IurFizLitso: String
	KomissiiaKakKupliaProdazha: Boolean
	IDKorr: String
	DataPriniatiiaKDeistviiuPisma50: DateTime
}
input CatalogUsloviiaOplatyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidOtsrochki: Boolean
	TablitsaVyplat: [CatalogUsloviiaOplatyTablitsaVyplatRowTypeInput!]
}
type CatalogUsloviiaOplaty {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidOtsrochki: Boolean
	TablitsaVyplat: [CatalogUsloviiaOplatyTablitsaVyplatRowType!]
}
input CatalogUsloviiaOplatyTablitsaVyplatInput {
	Key: Guid!
	LineNumber: Int64!
	PeriodOtsrochki: Int16
	ProtsentVyplaty: Double
}
type CatalogUsloviiaOplatyTablitsaVyplat {
	Key: Guid!
	LineNumber: Int64!
	PeriodOtsrochki: Int16
	ProtsentVyplaty: Double
}
input CatalogKategoriiObieektovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	NaznachenieKategoriiKey: Guid
}
type CatalogKategoriiObieektov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	NaznachenieKategoriiKey: Guid
}
input CatalogfmZnacheniiaDliaZapolneniiaInput {
	Key: Guid!
	DataVersion: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Opisanie: String
}
type CatalogfmZnacheniiaDliaZapolneniia {
	Key: Guid!
	DataVersion: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Opisanie: String
}
input DocumentUdalitNariadZakazInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	OtvetstvennyiKey: Guid
	Comment: String
	Opisanie: String
	PriemIzdelii: Boolean
	PriemMetallaVstavok: Boolean
	DepartmentKey: Guid
	KolichestvoDokumenta: Int64
	Weight: Double
	Izdeliia: [DocumentUdalitNariadZakazIzdeliiaRowTypeInput!]
	Uslugi: [DocumentUdalitNariadZakazUslugiRowTypeInput!]
	Spetsifikatsiia: [DocumentUdalitNariadZakazSpetsifikatsiiaRowTypeInput!]
	Metally: [DocumentUdalitNariadZakazMetallyRowTypeInput!]
	Vstavki: [DocumentUdalitNariadZakazVstavkiRowTypeInput!]
}
type DocumentUdalitNariadZakaz {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	OtvetstvennyiKey: Guid
	Comment: String
	Opisanie: String
	PriemIzdelii: Boolean
	PriemMetallaVstavok: Boolean
	DepartmentKey: Guid
	KolichestvoDokumenta: Int64
	Weight: Double
	Izdeliia: [DocumentUdalitNariadZakazIzdeliiaRowType!]
	Uslugi: [DocumentUdalitNariadZakazUslugiRowType!]
	Spetsifikatsiia: [DocumentUdalitNariadZakazSpetsifikatsiiaRowType!]
	Metally: [DocumentUdalitNariadZakazMetallyRowType!]
	Vstavki: [DocumentUdalitNariadZakazVstavkiRowType!]
}
input DocumentUdalitNariadZakazIzdeliiaInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	Weight: Double
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TypeKey: Guid
	ProbeKey: Guid
	Kommentarii: String
	VesBezVstavok: Double
	NomerStrokiTCh: Int64
	InstanceKey: Guid
}
type DocumentUdalitNariadZakazIzdeliia {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	Weight: Double
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TypeKey: Guid
	ProbeKey: Guid
	Kommentarii: String
	VesBezVstavok: Double
	NomerStrokiTCh: Int64
	InstanceKey: Guid
}
input DocumentUdalitNariadZakazUslugiInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentUdalitNariadZakazUslugi {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input DocumentUdalitNariadZakazSpetsifikatsiiaInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	NomerStrokiTCh: Int64
}
type DocumentUdalitNariadZakazSpetsifikatsiia {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	NomerStrokiTCh: Int64
}
input DocumentUdalitNariadZakazMetallyInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	ProbeKey: Guid
	Weight: String
}
type DocumentUdalitNariadZakazMetally {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	ProbeKey: Guid
	Weight: String
}
input DocumentUdalitNariadZakazVstavkiInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
}
type DocumentUdalitNariadZakazVstavki {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
}
input CatalogBankiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Adres: String
	Gorod: String
	KorrSchet: String
	Telefony: String
}
type CatalogBanki {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Adres: String
	Gorod: String
	KorrSchet: String
	Telefony: String
}
input CatalogRoliKontaktnykhLitsInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogRoliKontaktnykhLits {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input DocumentRestrukturizatsiiaDolgaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	RasshifrovkaZadolzhennosti: [DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostiRowTypeInput!]
}
type DocumentRestrukturizatsiiaDolga {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	RasshifrovkaZadolzhennosti: [DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostiRowType!]
}
input DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostiInput {
	Key: Guid!
	LineNumber: Int64!
	DataDolga: DateTime
	NovaiaDataDolga: DateTime
	NovaiaSummaDolga: Double
	NovaiaSummaDolgaUpr: Double
	SummaDolga: Double
	SummaDolgaUpr: Double
}
type DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti {
	Key: Guid!
	LineNumber: Int64!
	DataDolga: DateTime
	NovaiaDataDolga: DateTime
	NovaiaSummaDolga: Double
	NovaiaSummaDolgaUpr: Double
	SummaDolga: Double
	SummaDolgaUpr: Double
}
input DocumentAkkreditivPoluchennyiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	DataVkhodiashchegoDokumenta: DateTime
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPlanirovaniiaPostupleniiaKey: Guid
	Comment: String
	KontragentKey: Guid
	NomerVkhodiashchegoDokumenta: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	ExtendedPayments: [DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhaRowTypeInput!]
	RekvizityKontragenta: [DocumentAkkreditivPoluchennyiRekvizityKontragentaRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentAkkreditivPoluchennyi {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	DataVkhodiashchegoDokumenta: DateTime
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPlanirovaniiaPostupleniiaKey: Guid
	Comment: String
	KontragentKey: Guid
	NomerVkhodiashchegoDokumenta: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	ExtendedPayments: [DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhaRowType!]
	RekvizityKontragenta: [DocumentAkkreditivPoluchennyiRekvizityKontragentaRowType!]
	DokumentOsnovanieType: String
}
input DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentAkkreditivPoluchennyiRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentAkkreditivPoluchennyiRekvizityKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentAkkreditivPoluchennyiRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input DocumentPriemIzRemontaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DogovorKontragentaKey: Guid
	KolichestvoDokumenta: Int64
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	SobstvennaiaMasterskaia: Boolean
	Izdeliia: [DocumentPriemIzRemontaIzdeliiaRowTypeInput!]
	Materialy: [DocumentPriemIzRemontaMaterialyRowTypeInput!]
}
type DocumentPriemIzRemonta {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DogovorKontragentaKey: Guid
	KolichestvoDokumenta: Int64
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	SobstvennaiaMasterskaia: Boolean
	Izdeliia: [DocumentPriemIzRemontaIzdeliiaRowType!]
	Materialy: [DocumentPriemIzRemontaMaterialyRowType!]
}
input DocumentPriemIzRemontaIzdeliiaInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	Weight: Double
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TypeKey: Guid
	ProbeKey: Guid
	InstanceKey: Guid
}
type DocumentPriemIzRemontaIzdeliia {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	Weight: Double
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	TypeKey: Guid
	ProbeKey: Guid
	InstanceKey: Guid
}
input DocumentPriemIzRemontaMaterialyInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchNomenklaturyKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	DokumentOprikhodovaniiaKey: Guid
	Weight: Double
	Quantity: Int64
}
type DocumentPriemIzRemontaMaterialy {
	Key: Guid!
	LineNumber: Int64!
	KliuchNomenklaturyKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	DokumentOprikhodovaniiaKey: Guid
	Weight: Double
	Quantity: Int64
}
input CatalogTsvetaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	Vyvodit: Boolean
	EstGruppaDefekta: Boolean
	EstGruppaTsveta: Boolean
}
type CatalogTsveta {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	Vyvodit: Boolean
	EstGruppaDefekta: Boolean
	EstGruppaTsveta: Boolean
}
input DocumentStornirovanieOtchetaKomissioneraOProdazhakhInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NomenklaturnaiaGruppaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	ProtsentKomissionnogoVoznagrazhdeniia: Double
	Sdelka: String
	SposobRaschetaKomissionnogoVoznagrazhdeniia: String
	StavkaNDSVoznagrazhdeniia: String
	StatiaZatratKey: Guid
	SummaVkliuchaetNDS: Boolean
	SummaVoznagrazhdeniia: Double
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UderzhatKomissionnoeVoznagrazhdenie: Boolean
	UchityvatNDS: Boolean
	DenezhnyeSredstva: [DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvaRowTypeInput!]
	Goods: [DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type DocumentStornirovanieOtchetaKomissioneraOProdazhakh {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NomenklaturnaiaGruppaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	ProtsentKomissionnogoVoznagrazhdeniia: Double
	Sdelka: String
	SposobRaschetaKomissionnogoVoznagrazhdeniia: String
	StavkaNDSVoznagrazhdeniia: String
	StatiaZatratKey: Guid
	SummaVkliuchaetNDS: Boolean
	SummaVoznagrazhdeniia: Double
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UderzhatKomissionnoeVoznagrazhdenie: Boolean
	UchityvatNDS: Boolean
	DenezhnyeSredstva: [DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvaRowType!]
	Goods: [DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovaryRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvaInput {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
type DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
input DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartiiKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDS: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaNDSPeredachi: Double
	SummaPeredachi: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPeredachi: Double
}
type DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartiiKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDS: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaNDSPeredachi: Double
	SummaPeredachi: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPeredachi: Double
}
input DocumentJournalDavalcheskieDokumentyInput {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	Weight: Double
	DataVkh: DateTime
	Comment: String
	KontragentKey: Guid
	Metall: String
	NomerVkh: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	RefType: String!
	MetallType: String
}
type DocumentJournalDavalcheskieDokumenty {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	Weight: Double
	DataVkh: DateTime
	Comment: String
	KontragentKey: Guid
	Metall: String
	NomerVkh: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	RefType: String!
	MetallType: String
}
input CatalogfmAnketaKlientaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	DannyeKontragenta: [CatalogfmAnketaKlientaDannyeKontragentaRowTypeInput!]
}
type CatalogfmAnketaKlienta {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	DannyeKontragenta: [CatalogfmAnketaKlientaDannyeKontragentaRowType!]
}
input CatalogfmAnketaKlientaDannyeKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
type CatalogfmAnketaKlientaDannyeKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
input CatalogVidyVzaimoraschetovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogVidyVzaimoraschetov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input DocumentUstanovkaZnacheniiTochkiZakazaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	Goods: [DocumentUstanovkaZnacheniiTochkiZakazaTovaryRowTypeInput!]
}
type DocumentUstanovkaZnacheniiTochkiZakaza {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	Goods: [DocumentUstanovkaZnacheniiTochkiZakazaTovaryRowType!]
}
input DocumentUstanovkaZnacheniiTochkiZakazaTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DataKon: DateTime
	DataNach: DateTime
	ZnachenieTochkiZakaza: Int64
	MinimalnyiStrakhovoiZapas: Int64
	ItemKey: Guid
	ProtsentnaiaStavkaZnacheniiaTochkiZakaza: Int16
	ProtsentnaiaStavkaMinimalnogoStrakhovogoZapasa: Int16
	SizeKey: Guid
	DepartmentKey: Guid
	SposobOpredeleniiaZnacheniiaTochkiZakaza: String
}
type DocumentUstanovkaZnacheniiTochkiZakazaTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DataKon: DateTime
	DataNach: DateTime
	ZnachenieTochkiZakaza: Int64
	MinimalnyiStrakhovoiZapas: Int64
	ItemKey: Guid
	ProtsentnaiaStavkaZnacheniiaTochkiZakaza: Int16
	ProtsentnaiaStavkaMinimalnogoStrakhovogoZapasa: Int16
	SizeKey: Guid
	DepartmentKey: Guid
	SposobOpredeleniiaZnacheniiaTochkiZakaza: String
}
input CatalogZnacheniiaKodirovkiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
}
type CatalogZnacheniiaKodirovki {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
}
input CatalogPravilaProdazhInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	NalogooblozhenieNDS: String
	VidTsenKey: Guid
	Deistvuet: Boolean
	Goods: [CatalogPravilaProdazhTovaryRowTypeInput!]
}
type CatalogPravilaProdazh {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	NalogooblozhenieNDS: String
	VidTsenKey: Guid
	Deistvuet: Boolean
	Goods: [CatalogPravilaProdazhTovaryRowType!]
}
input CatalogPravilaProdazhTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	KharakteristikiNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Cost: Double
	Quantity: Int64
	Weight: Double
}
type CatalogPravilaProdazhTovary {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	KharakteristikiNomenklaturyKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Cost: Double
	Quantity: Int64
	Weight: Double
}
input DocumentPostuplenieDopRaskhodovInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	OperationType: String
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	Soderzhanie: String
	SposobRaspredeleniia: String
	StavkaNDS: String
	Sum: Double
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	SummaNDS: Double
	TipDokumenta: String
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	Goods: [DocumentPostuplenieDopRaskhodovTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type DocumentPostuplenieDopRaskhodov {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	OperationType: String
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	Soderzhanie: String
	SposobRaspredeleniia: String
	StavkaNDS: String
	Sum: Double
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	SummaNDS: Double
	TipDokumenta: String
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	Goods: [DocumentPostuplenieDopRaskhodovTovaryRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentPostuplenieDopRaskhodovTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartii: String
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SummaRaspredeleniia: Double
	SummaTovara: Double
	KharakteristikaNomenklaturyKey: Guid
	DokumentPartiiType: String
}
type DocumentPostuplenieDopRaskhodovTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPartii: String
	Quantity: Int64
	ItemKey: Guid
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SummaRaspredeleniia: Double
	SummaTovara: Double
	KharakteristikaNomenklaturyKey: Guid
	DokumentPartiiType: String
}
input CatalogKhoziaistvennyeOperatsiiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	KodKhozOperatsii: String
}
type CatalogKhoziaistvennyeOperatsii {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	KodKhozOperatsii: String
}
input DocumentAvansovyiOtchetInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KratnostDokumenta: Int64
	KursDokumenta: Double
	NaznachenieAvansa: String
	NDSVkliuchenVStoimost: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	SkladOrderKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	FizLitsoKey: Guid
	VydannyeAvansy: [DocumentAvansovyiOtchetVydannyeAvansyRowTypeInput!]
	Goods: [DocumentAvansovyiOtchetTovaryRowTypeInput!]
	OplataPostavshchikam: [DocumentAvansovyiOtchetOplataPostavshchikamRowTypeInput!]
	Prochee: [DocumentAvansovyiOtchetProcheeRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentAvansovyiOtchet {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KratnostDokumenta: Int64
	KursDokumenta: Double
	NaznachenieAvansa: String
	NDSVkliuchenVStoimost: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	SkladOrderKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	FizLitsoKey: Guid
	VydannyeAvansy: [DocumentAvansovyiOtchetVydannyeAvansyRowType!]
	Goods: [DocumentAvansovyiOtchetTovaryRowType!]
	OplataPostavshchikam: [DocumentAvansovyiOtchetOplataPostavshchikamRowType!]
	Prochee: [DocumentAvansovyiOtchetProcheeRowType!]
	DokumentOsnovanieType: String
}
input DocumentAvansovyiOtchetVydannyeAvansyInput {
	Key: Guid!
	LineNumber: Int64!
	RaskhodnyiKassovyiOrderKey: Guid
	Sum: Double
}
type DocumentAvansovyiOtchetVydannyeAvansy {
	Key: Guid!
	LineNumber: Int64!
	RaskhodnyiKassovyiOrderKey: Guid
	Sum: Double
}
input DocumentAvansovyiOtchetTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DataSF: DateTime
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	NomerVkhodiashchegoDokumenta: String
	NomerSF: String
	SupplierKey: Guid
	PredieiavlenSF: Boolean
	ProektKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SchetFakturaKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
}
type DocumentAvansovyiOtchetTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DataSF: DateTime
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	NomerVkhodiashchegoDokumenta: String
	NomerSF: String
	SupplierKey: Guid
	PredieiavlenSF: Boolean
	ProektKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SchetFakturaKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
}
input DocumentAvansovyiOtchetOplataPostavshchikamInput {
	Key: Guid!
	LineNumber: Int64!
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NomerVkhodiashchegoDokumenta: String
	Sdelka: String
	Soderzhanie: String
	Sum: Double
	SummaVzaimoraschetov: Double
	SdelkaType: String
}
type DocumentAvansovyiOtchetOplataPostavshchikam {
	Key: Guid!
	LineNumber: Int64!
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NomerVkhodiashchegoDokumenta: String
	Sdelka: String
	Soderzhanie: String
	Sum: Double
	SummaVzaimoraschetov: Double
	SdelkaType: String
}
input DocumentAvansovyiOtchetProcheeInput {
	Key: Guid!
	LineNumber: Int64!
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DataSF: DateTime
	Zakaz: String
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	NomerSF: String
	PodrazdelenieKey: Guid
	SupplierKey: Guid
	PredieiavlenSF: Boolean
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	SchetFakturaKey: Guid
	ZakazType: String
}
type DocumentAvansovyiOtchetProchee {
	Key: Guid!
	LineNumber: Int64!
	VidDokVkhodiashchii: String
	DataVkhodiashchegoDokumenta: DateTime
	DataSF: DateTime
	Zakaz: String
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	NomerSF: String
	PodrazdelenieKey: Guid
	SupplierKey: Guid
	PredieiavlenSF: Boolean
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	SchetFakturaKey: Guid
	ZakazType: String
}
input CatalogDolzhnostiOrganizatsiiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogDolzhnostiOrganizatsii {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input CatalogAnalitikaTipaIzdeliiaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	BIdentifikator: String
}
type CatalogAnalitikaTipaIzdeliia {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	BIdentifikator: String
}
input CatalogDopolnitelnyePechatnyeFormyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	KhranilishcheVneshneiObrabotkiBase64Data: Binary
	Publikatsiia: String
	Prinadlezhnost: [CatalogDopolnitelnyePechatnyeFormyPrinadlezhnostRowTypeInput!]
	KhranilishcheVneshneiObrabotkiType: String
	KhranilishcheVneshneiObrabotki: Stream
}
type CatalogDopolnitelnyePechatnyeFormy {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	KhranilishcheVneshneiObrabotkiBase64Data: Binary
	Publikatsiia: String
	Prinadlezhnost: [CatalogDopolnitelnyePechatnyeFormyPrinadlezhnostRowType!]
	KhranilishcheVneshneiObrabotkiType: String
	KhranilishcheVneshneiObrabotki: Stream
}
input CatalogDopolnitelnyePechatnyeFormyPrinadlezhnostInput {
	Key: Guid!
	LineNumber: Int64!
	PredstavlenieObieekta: String
	SsylkaObieekta: String
	SsylkaObieektaType: String
}
type CatalogDopolnitelnyePechatnyeFormyPrinadlezhnost {
	Key: Guid!
	LineNumber: Int64!
	PredstavlenieObieekta: String
	SsylkaObieekta: String
	SsylkaObieektaType: String
}
input MemberCardsTypeInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type MemberCardsType {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input DocumentRegistratsiiaNaSaiteInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Naimenovanie: String
	NaimenovaniePolnoe: String
	INN: String
	KPP: String
	FIO: String
	Telefon: String
	AdresElektronnoiPochty: String
	Comment: String
	KontragentKey: Guid
	Registratsiia: Boolean
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	PravoDostupa: String
}
type DocumentRegistratsiiaNaSaite {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Naimenovanie: String
	NaimenovaniePolnoe: String
	INN: String
	KPP: String
	FIO: String
	Telefon: String
	AdresElektronnoiPochty: String
	Comment: String
	KontragentKey: Guid
	Registratsiia: Boolean
	OrganizatsiiaKey: Guid
	DogovorKontragentaKey: Guid
	PravoDostupa: String
}
input CatalogObrabotkiObsluzhivaniiaTOInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Versiia: Double
	VersiiaAPI: Double
	Vid: String
	Identifikator: String
	ImiaFaila: String
	ObrabotkaBase64Data: Binary
	Opisanie: String
	Modeli: [CatalogObrabotkiObsluzhivaniiaTOModeliRowTypeInput!]
	ObrabotkaType: String
	Obrabotka: Stream
}
type CatalogObrabotkiObsluzhivaniiaTO {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Versiia: Double
	VersiiaAPI: Double
	Vid: String
	Identifikator: String
	ImiaFaila: String
	ObrabotkaBase64Data: Binary
	Opisanie: String
	Modeli: [CatalogObrabotkiObsluzhivaniiaTOModeliRowType!]
	ObrabotkaType: String
	Obrabotka: Stream
}
input CatalogObrabotkiObsluzhivaniiaTOModeliInput {
	Key: Guid!
	LineNumber: Int64!
	Model: String
}
type CatalogObrabotkiObsluzhivaniiaTOModeli {
	Key: Guid!
	LineNumber: Int64!
	Model: String
}
input CatalogNastroikaIntervalovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	TablichnaiaChast: [CatalogNastroikaIntervalovTablichnaiaChastRowTypeInput!]
}
type CatalogNastroikaIntervalov {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	TablichnaiaChast: [CatalogNastroikaIntervalovTablichnaiaChastRowType!]
}
input CatalogNastroikaIntervalovTablichnaiaChastInput {
	Key: Guid!
	LineNumber: Int64!
	KonetsIntervala: Int64
	NachaloIntervala: Int64
	Podpis: String
}
type CatalogNastroikaIntervalovTablichnaiaChast {
	Key: Guid!
	LineNumber: Int64!
	KonetsIntervala: Int64
	NachaloIntervala: Int64
	Podpis: String
}
input CatalogProfiliGruppDostupaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Comment: String
	IdentifikatorPostavliaemykhDannykh: Guid
	Roli: [CatalogProfiliGruppDostupaRoliRowTypeInput!]
	VidyDostupa: [CatalogProfiliGruppDostupaVidyDostupaRowTypeInput!]
	ZnacheniiaDostupa: [CatalogProfiliGruppDostupaZnacheniiaDostupaRowTypeInput!]
	DostupPoPodsistemam: [CatalogProfiliGruppDostupaDostupPoPodsistemamRowTypeInput!]
}
type CatalogProfiliGruppDostupa {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Comment: String
	IdentifikatorPostavliaemykhDannykh: Guid
	Roli: [CatalogProfiliGruppDostupaRoliRowType!]
	VidyDostupa: [CatalogProfiliGruppDostupaVidyDostupaRowType!]
	ZnacheniiaDostupa: [CatalogProfiliGruppDostupaZnacheniiaDostupaRowType!]
	DostupPoPodsistemam: [CatalogProfiliGruppDostupaDostupPoPodsistemamRowType!]
}
input CatalogProfiliGruppDostupaRoliInput {
	Key: Guid!
	LineNumber: Int64!
	RolKey: Guid
}
type CatalogProfiliGruppDostupaRoli {
	Key: Guid!
	LineNumber: Int64!
	RolKey: Guid
}
input CatalogProfiliGruppDostupaVidyDostupaInput {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	Predustanovlennyi: Boolean
	VseRazresheny: Boolean
	VidDostupaType: String
}
type CatalogProfiliGruppDostupaVidyDostupa {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	Predustanovlennyi: Boolean
	VseRazresheny: Boolean
	VidDostupaType: String
}
input CatalogProfiliGruppDostupaZnacheniiaDostupaInput {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	ZnachenieDostupa: String
	VidDostupaType: String
	ZnachenieDostupaType: String
}
type CatalogProfiliGruppDostupaZnacheniiaDostupa {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	ZnachenieDostupa: String
	VidDostupaType: String
	ZnachenieDostupaType: String
}
input CatalogProfiliGruppDostupaDostupPoPodsistemamInput {
	Key: Guid!
	LineNumber: Int64!
	ImiaPodsistemy: String
	ImiaObieekta: String
	Prosmotr: Boolean
	Redaktirovanie: Boolean
	Pechat: Boolean
	PokazVersii: Boolean
}
type CatalogProfiliGruppDostupaDostupPoPodsistemam {
	Key: Guid!
	LineNumber: Int64!
	ImiaPodsistemy: String
	ImiaObieekta: String
	Prosmotr: Boolean
	Redaktirovanie: Boolean
	Pechat: Boolean
	PokazVersii: Boolean
}
input CatalogNastroikiDliaKureraInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	KassaKKMKey: Guid
	VidOplaty2: String
	VidOplaty2PoUmolchaniiuKey: Guid
	IspolzovatVidOplaty2: Boolean
	PechatatVtoroiEkzempliarTovarnogoCheka: Boolean
	PechatatImiaKassira: Boolean
	PechatatNazvaniePlatezhnoiKarty: Boolean
	PechatatNumeratsiiuPozitsii: Boolean
	PechatatTovarnyiChekPriRegistratsiiProdazhi: Boolean
	PechatatShtrikhkodPriRegistratsiiProdazhi: Boolean
	TrebovatVvodPasportnykhDannykh: Boolean
	ZapolniatAnketuKlientaPriProdazhe: Boolean
	PrefiksDokumentov: String
	SostavNaimenovaniia: [CatalogNastroikiDliaKureraSostavNaimenovaniiaRowTypeInput!]
}
type CatalogNastroikiDliaKurera {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	KassaKKMKey: Guid
	VidOplaty2: String
	VidOplaty2PoUmolchaniiuKey: Guid
	IspolzovatVidOplaty2: Boolean
	PechatatVtoroiEkzempliarTovarnogoCheka: Boolean
	PechatatImiaKassira: Boolean
	PechatatNazvaniePlatezhnoiKarty: Boolean
	PechatatNumeratsiiuPozitsii: Boolean
	PechatatTovarnyiChekPriRegistratsiiProdazhi: Boolean
	PechatatShtrikhkodPriRegistratsiiProdazhi: Boolean
	TrebovatVvodPasportnykhDannykh: Boolean
	ZapolniatAnketuKlientaPriProdazhe: Boolean
	PrefiksDokumentov: String
	SostavNaimenovaniia: [CatalogNastroikiDliaKureraSostavNaimenovaniiaRowType!]
}
input CatalogNastroikiDliaKureraSostavNaimenovaniiaInput {
	Key: Guid!
	LineNumber: Int64!
	SimvolyDo: String
	SimvolyPosle: String
	ElementNaimenovaniia: String
}
type CatalogNastroikiDliaKureraSostavNaimenovaniia {
	Key: Guid!
	LineNumber: Int64!
	SimvolyDo: String
	SimvolyPosle: String
	ElementNaimenovaniia: String
}
input CatalogTipyTsenNomenklaturyKontragentovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	ValiutaTsenyKey: Guid
	Comment: String
	OpisanieTipaTsenyNomenklaturyKontragenta: String
	TipTsenyNomenklaturyKey: Guid
	TsenaVkliuchaetNDS: Boolean
}
type CatalogTipyTsenNomenklaturyKontragentov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	ValiutaTsenyKey: Guid
	Comment: String
	OpisanieTipaTsenyNomenklaturyKontragenta: String
	TipTsenyNomenklaturyKey: Guid
	TsenaVkliuchaetNDS: Boolean
}
input DocumentJournalTsenoobrazovanieInput {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	Informatsiia: String
	Comment: String
	KontragentKey: Guid
	OtvetstvennyiKey: Guid
	RefType: String!
	InformatsiiaType: String
}
type DocumentJournalTsenoobrazovanie {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	Informatsiia: String
	Comment: String
	KontragentKey: Guid
	OtvetstvennyiKey: Guid
	RefType: String!
	InformatsiiaType: String
}
input CatalogEdinitsyIzmereniiaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	Owner: String
	DeletionMark: Boolean
	Weight: Double
	EdinitsaPoKlassifikatoruKey: Guid
	Koeffitsient: Double
	Obieem: Double
	BIdentifikator: String
	BNomerVersii: String
	OwnerType: String
}
type CatalogEdinitsyIzmereniia {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	Owner: String
	DeletionMark: Boolean
	Weight: Double
	EdinitsaPoKlassifikatoruKey: Guid
	Koeffitsient: Double
	Obieem: Double
	BIdentifikator: String
	BNomerVersii: String
	OwnerType: String
}
input CatalogStatiDvizheniiaDenezhnykhSredstvInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	KorrespondiruiushchiiSchet: String
}
type CatalogStatiDvizheniiaDenezhnykhSredstv {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	KorrespondiruiushchiiSchet: String
}
input DocumentInkassovoePorucheniePoluchennoeInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPlatezha: String
	DataVkhodiashchegoDokumenta: DateTime
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	ZaiavkaNaRaskhodovanieSredstvKey: Guid
	Comment: String
	KontragentKey: Guid
	NaznacheniePlatezha: String
	NomerVkhodiashchegoDokumenta: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	OcherednostPlatezha: Int16
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	PodrazdelenieKey: Guid
	ExtendedPayments: [DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhaRowTypeInput!]
	RekvizityKontragenta: [DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentaRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentInkassovoePorucheniePoluchennoe {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPlatezha: String
	DataVkhodiashchegoDokumenta: DateTime
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	ZaiavkaNaRaskhodovanieSredstvKey: Guid
	Comment: String
	KontragentKey: Guid
	NaznacheniePlatezha: String
	NomerVkhodiashchegoDokumenta: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	OcherednostPlatezha: Int16
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	PodrazdelenieKey: Guid
	ExtendedPayments: [DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhaRowType!]
	RekvizityKontragenta: [DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentaRowType!]
	DokumentOsnovanieType: String
}
input DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	NomerPlatezha: Int16
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	NomerPlatezha: Int16
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input CatalogNastroikiObmenaDannymiShtrikhMInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidZagruzkiOtchetov: String
	VidOplaty2Key: Guid
	VidOplaty3Key: Guid
	VidOplaty4Key: Guid
	VygruzhatSkhemyNakopleniia: Boolean
	ImiaFailaVygruzki: String
	ImiaFailaZagruzki: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
}
type CatalogNastroikiObmenaDannymiShtrikhM {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidZagruzkiOtchetov: String
	VidOplaty2Key: Guid
	VidOplaty3Key: Guid
	VidOplaty4Key: Guid
	VygruzhatSkhemyNakopleniia: Boolean
	ImiaFailaVygruzki: String
	ImiaFailaZagruzki: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
}
input CatalogStatiZatratInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	VidZatrat: String
	VidRaskhodovNU: String
	OtnesenieRaskhodovKDeiatelnostiENVD: String
	KharakterZatrat: String
	KorrespondiruiushchiiSchet: String
}
type CatalogStatiZatrat {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	VidZatrat: String
	VidRaskhodovNU: String
	OtnesenieRaskhodovKDeiatelnostiENVD: String
	KharakterZatrat: String
	KorrespondiruiushchiiSchet: String
}
input DocumentVozvratTovarovOtPokupateliaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetKontragentaKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	ZakrytieDnia: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	Metall: String
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Sdelka: String
	SkladOrderKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatVesVstavok: Boolean
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	Goods: [DocumentVozvratTovarovOtPokupateliaTovaryRowTypeInput!]
	Uslugi: [DocumentVozvratTovarovOtPokupateliaUslugiRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type DocumentVozvratTovarovOtPokupatelia {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetKontragentaKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	ZakrytieDnia: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	Metall: String
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Sdelka: String
	SkladOrderKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatVesVstavok: Boolean
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	Goods: [DocumentVozvratTovarovOtPokupateliaTovaryRowType!]
	Uslugi: [DocumentVozvratTovarovOtPokupateliaUslugiRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentVozvratTovarovOtPokupateliaTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentOprikhodovaniia: String
	DokumentPartii: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	PercentManualDiscount: Double
	SizeKey: Guid
	Sebestoimost: Double
	SebestoimostBezNDS: Double
	SebestoimostUpr: Double
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	StatusPartii: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	DokumentOprikhodovaniiaType: String
	DokumentPartiiType: String
}
type DocumentVozvratTovarovOtPokupateliaTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentOprikhodovaniia: String
	DokumentPartii: String
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	PercentManualDiscount: Double
	SizeKey: Guid
	Sebestoimost: Double
	SebestoimostBezNDS: Double
	SebestoimostUpr: Double
	InstanceKey: Guid
	SumManualDiscount: Double
	DepartmentKey: Guid
	StavkaNDS: String
	StatusPartii: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	OrderKey: Guid
	DokumentOprikhodovaniiaType: String
	DokumentPartiiType: String
}
input DocumentVozvratTovarovOtPokupateliaUslugiInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	ProtsentOtSummuDokumenta: Double
	ProtsentOtSummyDokumenta: Double
	ProtsentSkidkiNatsenki: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentVozvratTovarovOtPokupateliaUslugi {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	ProtsentOtSummuDokumenta: Double
	ProtsentOtSummyDokumenta: Double
	ProtsentSkidkiNatsenki: Double
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input DocumentZakazPostavshchikuInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	VremiaNapominaniia: DateTime
	DataOplaty: DateTime
	DataPostupleniia: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontaktnoeLitsoKey: Guid
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NapomnitOSobytii: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	ParametryOtboraBase64Data: Binary
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	Soglasovano: Boolean
	StrukturnaiaEdinitsa: String
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	Goods: [DocumentZakazPostavshchikuTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	ParametryOtboraType: String
	StrukturnaiaEdinitsaType: String
	ParametryOtbora: Stream
}
type DocumentZakazPostavshchiku {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	VremiaNapominaniia: DateTime
	DataOplaty: DateTime
	DataPostupleniia: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontaktnoeLitsoKey: Guid
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NapomnitOSobytii: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	ParametryOtboraBase64Data: Binary
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	Soglasovano: Boolean
	StrukturnaiaEdinitsa: String
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	Goods: [DocumentZakazPostavshchikuTovaryRowType!]
	DokumentOsnovanieType: String
	ParametryOtboraType: String
	StrukturnaiaEdinitsaType: String
	ParametryOtbora: Stream
}
input DocumentZakazPostavshchikuTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentZakazPostavshchikuTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZakazKlientaKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input CatalogSkidkiNatsenkiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	ValiutaPredostavleniiaKey: Guid
	VariantSovmestnogoPrimeneniia: String
	ZnachenieSkidkiNatsenki: Double
	MomentVydachiSoobshcheniia: String
	OblastPredostavleniia: String
	PodarokIzKorzinyPokupatelia: Boolean
	PodarokIzSpiska: Boolean
	RekvizitDopUporiadochivaniia: Int64
	SegmentNomenklaturyPredostavleniiaKey: Guid
	SposobPredostavleniia: String
	StatusDeistviia: String
	TekstSoobshcheniia: String
	Upravliaemaia: Boolean
	UchityvatPodarokKakProdazhu: Boolean
	KuponKey: Guid
	DataSpisaniia: Boolean
	VariantOtboraNomenklatury: String
	SegmentNomenklaturyIskliucheniiaKey: Guid
	VariantOtboraIskliucheniiNomenklatury: String
	KhranilishcheNastroekKomponovkiDannykhBase64Data: Binary
	KhranilishcheNastroekKomponovkiDannykhIskliuchenieBase64Data: Binary
	VidTsenyKey: Guid
	KolichestvoPodarkovIzKorzinyPokupatelia: Int64
	KratnoKolichestvuUslovii: Boolean
	OgranichenieRazmeraPodchinennykhSkidok: Boolean
	UsloviiaPredostavleniia: [CatalogSkidkiNatsenkiUsloviiaPredostavleniiaRowTypeInput!]
	TsenovyeGruppy: [CatalogSkidkiNatsenkiTsenovyeGruppyRowTypeInput!]
	NaborPodarkov: [CatalogSkidkiNatsenkiNaborPodarkovRowTypeInput!]
	KhranilishcheNastroekKomponovkiDannykhType: String
	KhranilishcheNastroekKomponovkiDannykhIskliuchenieType: String
	KhranilishcheNastroekKomponovkiDannykh: Stream
	KhranilishcheNastroekKomponovkiDannykhIskliuchenie: Stream
}
type CatalogSkidkiNatsenki {
	Key: Guid!
	DataVersion: String
	Description: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	ValiutaPredostavleniiaKey: Guid
	VariantSovmestnogoPrimeneniia: String
	ZnachenieSkidkiNatsenki: Double
	MomentVydachiSoobshcheniia: String
	OblastPredostavleniia: String
	PodarokIzKorzinyPokupatelia: Boolean
	PodarokIzSpiska: Boolean
	RekvizitDopUporiadochivaniia: Int64
	SegmentNomenklaturyPredostavleniiaKey: Guid
	SposobPredostavleniia: String
	StatusDeistviia: String
	TekstSoobshcheniia: String
	Upravliaemaia: Boolean
	UchityvatPodarokKakProdazhu: Boolean
	KuponKey: Guid
	DataSpisaniia: Boolean
	VariantOtboraNomenklatury: String
	SegmentNomenklaturyIskliucheniiaKey: Guid
	VariantOtboraIskliucheniiNomenklatury: String
	KhranilishcheNastroekKomponovkiDannykhBase64Data: Binary
	KhranilishcheNastroekKomponovkiDannykhIskliuchenieBase64Data: Binary
	VidTsenyKey: Guid
	KolichestvoPodarkovIzKorzinyPokupatelia: Int64
	KratnoKolichestvuUslovii: Boolean
	OgranichenieRazmeraPodchinennykhSkidok: Boolean
	UsloviiaPredostavleniia: [CatalogSkidkiNatsenkiUsloviiaPredostavleniiaRowType!]
	TsenovyeGruppy: [CatalogSkidkiNatsenkiTsenovyeGruppyRowType!]
	NaborPodarkov: [CatalogSkidkiNatsenkiNaborPodarkovRowType!]
	KhranilishcheNastroekKomponovkiDannykhType: String
	KhranilishcheNastroekKomponovkiDannykhIskliuchenieType: String
	KhranilishcheNastroekKomponovkiDannykh: Stream
	KhranilishcheNastroekKomponovkiDannykhIskliuchenie: Stream
}
input CatalogSkidkiNatsenkiUsloviiaPredostavleniiaInput {
	Key: Guid!
	LineNumber: Int64!
	UsloviePredostavleniiaKey: Guid
}
type CatalogSkidkiNatsenkiUsloviiaPredostavleniia {
	Key: Guid!
	LineNumber: Int64!
	UsloviePredostavleniiaKey: Guid
}
input CatalogSkidkiNatsenkiTsenovyeGruppyInput {
	Key: Guid!
	LineNumber: Int64!
	TsenovaiaGruppaKey: Guid
	ZnachenieSkidkiNatsenki: Double
}
type CatalogSkidkiNatsenkiTsenovyeGruppy {
	Key: Guid!
	LineNumber: Int64!
	TsenovaiaGruppaKey: Guid
	ZnachenieSkidkiNatsenki: Double
}
input CatalogSkidkiNatsenkiNaborPodarkovInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Quantity: Double
	SizeKey: Guid
	InstanceKey: Guid
}
type CatalogSkidkiNatsenkiNaborPodarkov {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Quantity: Double
	SizeKey: Guid
	InstanceKey: Guid
}
input CatalogGruppyTsvetovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
}
type CatalogGruppyTsvetov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
}
input DocumentDokumentRaschetovSKontragentomInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanieKey: Guid
	Comment: String
	KontragentKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	OperationType: String
	SumOfDocument: Double
}
type DocumentDokumentRaschetovSKontragentom {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanieKey: Guid
	Comment: String
	KontragentKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	OperationType: String
	SumOfDocument: Double
}
input CatalogDogovoryEkvairingaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	AvtomaticheskiVychitatSummuTorgovoiUstupki: Boolean
	DogovorVzaimoraschetovKey: Guid
	EkvairerKey: Guid
	TarifyZaRaschetnoeObsluzhivanie: [CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanieRowTypeInput!]
}
type CatalogDogovoryEkvairinga {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	AvtomaticheskiVychitatSummuTorgovoiUstupki: Boolean
	DogovorVzaimoraschetovKey: Guid
	EkvairerKey: Guid
	TarifyZaRaschetnoeObsluzhivanie: [CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanieRowType!]
}
input CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanieInput {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	ProtsentTorgovoiUstupki: Double
}
type CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie {
	Key: Guid!
	LineNumber: Int64!
	VidOplatyKey: Guid
	ProtsentTorgovoiUstupki: Double
}
input CatalogKachestvoInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogKachestvo {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input DocumentUstanovkaTsenNomenklaturyKontragentovInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DokumentOsnovanie: String
	Informatsiia: String
	Comment: String
	KontragentKey: Guid
	NeProvoditNulevyeZnacheniia: Boolean
	OtvetstvennyiKey: Guid
	TipyTsen: [DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsenRowTypeInput!]
	Goods: [DocumentUstanovkaTsenNomenklaturyKontragentovTovaryRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentUstanovkaTsenNomenklaturyKontragentov {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DokumentOsnovanie: String
	Informatsiia: String
	Comment: String
	KontragentKey: Guid
	NeProvoditNulevyeZnacheniia: Boolean
	OtvetstvennyiKey: Guid
	TipyTsen: [DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsenRowType!]
	Goods: [DocumentUstanovkaTsenNomenklaturyKontragentovTovaryRowType!]
	DokumentOsnovanieType: String
}
input DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsenInput {
	Key: Guid!
	LineNumber: Int64!
	TipTsenKey: Guid
}
type DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen {
	Key: Guid!
	LineNumber: Int64!
	TipTsenKey: Guid
}
input DocumentUstanovkaTsenNomenklaturyKontragentovTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	ValiutaKey: Guid
	IndeksStrokiTablitsyTsen: Int64
	ItemKey: Guid
	SizeKey: Guid
	TipTsenKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
type DocumentUstanovkaTsenNomenklaturyKontragentovTovary {
	Key: Guid!
	LineNumber: Int64!
	ValiutaKey: Guid
	IndeksStrokiTablitsyTsen: Int64
	ItemKey: Guid
	SizeKey: Guid
	TipTsenKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
input DocumentProtsentPoterPoDavaltsamInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DogovorKontragentaKey: Guid
	DokumentOsnovanieKey: Guid
	Comment: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	Protsenty: [DocumentProtsentPoterPoDavaltsamProtsentyRowTypeInput!]
}
type DocumentProtsentPoterPoDavaltsam {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DogovorKontragentaKey: Guid
	DokumentOsnovanieKey: Guid
	Comment: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	Protsenty: [DocumentProtsentPoterPoDavaltsamProtsentyRowType!]
}
input DocumentProtsentPoterPoDavaltsamProtsentyInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Protsent: Double
}
type DocumentProtsentPoterPoDavaltsamProtsenty {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Protsent: Double
}
input CatalogTovarnyePozitsiiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	Pozitsiia1: String
	Pozitsiia2: String
	Pozitsiia3: String
	Pozitsiia4: String
	Pozitsiia5: String
	Pozitsiia6: String
	Pozitsiia7: String
	Pozitsiia1Type: String
	Pozitsiia2Type: String
	Pozitsiia3Type: String
	Pozitsiia4Type: String
	Pozitsiia5Type: String
	Pozitsiia6Type: String
	Pozitsiia7Type: String
}
type CatalogTovarnyePozitsii {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	Pozitsiia1: String
	Pozitsiia2: String
	Pozitsiia3: String
	Pozitsiia4: String
	Pozitsiia5: String
	Pozitsiia6: String
	Pozitsiia7: String
	Pozitsiia1Type: String
	Pozitsiia2Type: String
	Pozitsiia3Type: String
	Pozitsiia4Type: String
	Pozitsiia5Type: String
	Pozitsiia6Type: String
	Pozitsiia7Type: String
}
input DocumentPlatezhnoePoruchenieIskhodiashcheeInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPerechisleniiaVBiudzhet: String
	VidPlatezha: String
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	IdentifikatorPlatezha: String
	INNPlatelshchika: String
	INNPoluchatelia: String
	KodBK: String
	KodOKATO: String
	Comment: String
	KontragentKey: Guid
	KPPPlatelshchika: String
	KPPPoluchatelia: String
	NaznacheniePlatezha: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	OcherednostPlatezha: Int16
	PerechislenieVBiudzhet: Boolean
	PodrazdelenieKey: Guid
	PokazatelDaty: DateTime
	PokazatelNomera: String
	PokazatelOsnovaniia: String
	PokazatelPerioda: String
	PokazatelTipa: String
	StatusSostavitelia: String
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TekstPlatelshchika: String
	TekstPoluchatelia: String
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	ExtendedPayments: [DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhaRowTypeInput!]
	RekvizityKontragenta: [DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentaRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentPlatezhnoePoruchenieIskhodiashchee {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	OperationType: String
	VidPerechisleniiaVBiudzhet: String
	VidPlatezha: String
	DataOplaty: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	IdentifikatorPlatezha: String
	INNPlatelshchika: String
	INNPoluchatelia: String
	KodBK: String
	KodOKATO: String
	Comment: String
	KontragentKey: Guid
	KPPPlatelshchika: String
	KPPPoluchatelia: String
	NaznacheniePlatezha: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	OcherednostPlatezha: Int16
	PerechislenieVBiudzhet: Boolean
	PodrazdelenieKey: Guid
	PokazatelDaty: DateTime
	PokazatelNomera: String
	PokazatelOsnovaniia: String
	PokazatelPerioda: String
	PokazatelTipa: String
	StatusSostavitelia: String
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	TekstPlatelshchika: String
	TekstPoluchatelia: String
	TipDokumenta: String
	ChastichnaiaOplata: Boolean
	ExtendedPayments: [DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhaRowType!]
	RekvizityKontragenta: [DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentaRowType!]
	DokumentOsnovanieType: String
}
input DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
type DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Znachenie: String
	Predstavlenie: String
	Rekvizit: String
	TipKontragenta: String
}
input CatalogfmOrganizatsionnoPravovyeFormyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	PolnoeNaimenovanie: String
}
type CatalogfmOrganizatsionnoPravovyeFormy {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	PolnoeNaimenovanie: String
}
input CatalogTipyTsenNomenklaturyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	BazovyiTipTsenKey: Guid
	ValiutaTsenyKey: Guid
	VidSkidkiNatsenki: String
	VygruzhatTsenyNaSait: Boolean
	Comment: String
	NatsenkaNaSebestoimost: Int16
	OkrugliatVBolshuiuStoronu: Boolean
	OkrugliatSummu: Boolean
	PoriadokOkrugleniia: String
	ProtsentSkidkiNatsenki: Double
	Rasschityvaetsia: Boolean
	TsenaVkliuchaetNDS: Boolean
	DliaSegmentov: Boolean
	BIdentifikator: String
}
type CatalogTipyTsenNomenklatury {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	BazovyiTipTsenKey: Guid
	ValiutaTsenyKey: Guid
	VidSkidkiNatsenki: String
	VygruzhatTsenyNaSait: Boolean
	Comment: String
	NatsenkaNaSebestoimost: Int16
	OkrugliatVBolshuiuStoronu: Boolean
	OkrugliatSummu: Boolean
	PoriadokOkrugleniia: String
	ProtsentSkidkiNatsenki: Double
	Rasschityvaetsia: Boolean
	TsenaVkliuchaetNDS: Boolean
	DliaSegmentov: Boolean
	BIdentifikator: String
}
input CatalogStatiOtchetaPoProdazhamInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
}
type CatalogStatiOtchetaPoProdazham {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
}
input CatalogVidyKodirovokiIzdeliiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ElementyKodirovki: [CatalogVidyKodirovokiIzdeliiElementyKodirovkiRowTypeInput!]
}
type CatalogVidyKodirovokiIzdelii {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	ElementyKodirovki: [CatalogVidyKodirovokiIzdeliiElementyKodirovkiRowType!]
}
input CatalogVidyKodirovokiIzdeliiElementyKodirovkiInput {
	Key: Guid!
	LineNumber: Int64!
	Poriadok: Int64
	Poteri: Boolean
	Prais: Boolean
	ElementKey: Guid
}
type CatalogVidyKodirovokiIzdeliiElementyKodirovki {
	Key: Guid!
	LineNumber: Int64!
	Poriadok: Int64
	Poteri: Boolean
	Prais: Boolean
	ElementKey: Guid
}
input DocumentUstanovkaSkidokNomenklaturyInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaKey: Guid
	VidSkidki: String
	DataNachala: DateTime
	DataOkonchaniia: DateTime
	DliaVseiNomenklatury: Boolean
	DliaVsekhPoluchatelei: Boolean
	DokumentSozdanVIuTD: Boolean
	ZnachenieUsloviia: String
	Comment: String
	ObshcheeVremiaNachala: DateTime
	ObshcheeVremiaOkonchaniia: DateTime
	OgranichenieSkidkiNatsenki: Double
	OtvetstvennyiKey: Guid
	ParametryZapolneniiaBase64Data: Binary
	PoDniamNedeli: Boolean
	ProtsentSkidkiNatsenki: Double
	TipSkidkiNatsenkiKey: Guid
	Uslovie: String
	VremiaPoDniamNedeli: [DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeliRowTypeInput!]
	DiskontnyeKarty: [DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartyRowTypeInput!]
	PoluchateliSkidki: [DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkiRowTypeInput!]
	Goods: [DocumentUstanovkaSkidokNomenklaturyTovaryRowTypeInput!]
	ZnachenieUsloviiaType: String
	ParametryZapolneniiaType: String
	ParametryZapolneniia: Stream
}
type DocumentUstanovkaSkidokNomenklatury {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaKey: Guid
	VidSkidki: String
	DataNachala: DateTime
	DataOkonchaniia: DateTime
	DliaVseiNomenklatury: Boolean
	DliaVsekhPoluchatelei: Boolean
	DokumentSozdanVIuTD: Boolean
	ZnachenieUsloviia: String
	Comment: String
	ObshcheeVremiaNachala: DateTime
	ObshcheeVremiaOkonchaniia: DateTime
	OgranichenieSkidkiNatsenki: Double
	OtvetstvennyiKey: Guid
	ParametryZapolneniiaBase64Data: Binary
	PoDniamNedeli: Boolean
	ProtsentSkidkiNatsenki: Double
	TipSkidkiNatsenkiKey: Guid
	Uslovie: String
	VremiaPoDniamNedeli: [DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeliRowType!]
	DiskontnyeKarty: [DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartyRowType!]
	PoluchateliSkidki: [DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkiRowType!]
	Goods: [DocumentUstanovkaSkidokNomenklaturyTovaryRowType!]
	ZnachenieUsloviiaType: String
	ParametryZapolneniiaType: String
	ParametryZapolneniia: Stream
}
input DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeliInput {
	Key: Guid!
	LineNumber: Int64!
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
	Vybran: Boolean
	DenNedeli: String
}
type DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli {
	Key: Guid!
	LineNumber: Int64!
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
	Vybran: Boolean
	DenNedeli: String
}
input DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartyInput {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
}
type DocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty {
	Key: Guid!
	LineNumber: Int64!
	MemberCardKey: Guid
}
input DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkiInput {
	Key: Guid!
	LineNumber: Int64!
	Poluchatel: String
	PoluchatelType: String
}
type DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki {
	Key: Guid!
	LineNumber: Int64!
	Poluchatel: String
	PoluchatelType: String
}
input DocumentUstanovkaSkidokNomenklaturyTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	OgranichenieSkidkiNatsenki: Double
	ProtsentSkidkiNatsenki: Double
}
type DocumentUstanovkaSkidokNomenklaturyTovary {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	OgranichenieSkidkiNatsenki: Double
	ProtsentSkidkiNatsenki: Double
}
input CatalogUsloviiaPredostavleniiaSkidokNatsenokInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	UsloviePredostavleniia: String
	VariantOpredeleniiaPeriodaNakopitelnoiSkidki: String
	VariantNakopleniia: String
	KriteriiOgranicheniiaPrimeneniiaZaObieemProdazh: String
	OblastOgranicheniia: String
	ValiutaOgranicheniiaKey: Guid
	GrafikOplaty: String
	FormaOplaty: String
	ZnachenieUsloviiaOgranicheniia: Double
	SegmentNomenklaturyOgranicheniiaKey: Guid
	PeriodNakopleniia: String
	TipSravneniia: String
	GruppaPolzovateleiKey: Guid
	DneiDoDniaRozhdeniia: Int16
	DneiPosleDniaRozhdeniia: Int16
	PoriadkovyiNomerProdazhiKraten: Int64
	UchityvatPrimenennyeSkidkiVRamkakhPosledovatelnogoPrimeneniia: Boolean
	UchityvatTekushchiiChekVNakopleniiakh: Boolean
	VremiaDeistviia: [CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviiaRowTypeInput!]
	Poluchateli: [CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateliRowTypeInput!]
	KomplektPokupki: [CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkiRowTypeInput!]
}
type CatalogUsloviiaPredostavleniiaSkidokNatsenok {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	UsloviePredostavleniia: String
	VariantOpredeleniiaPeriodaNakopitelnoiSkidki: String
	VariantNakopleniia: String
	KriteriiOgranicheniiaPrimeneniiaZaObieemProdazh: String
	OblastOgranicheniia: String
	ValiutaOgranicheniiaKey: Guid
	GrafikOplaty: String
	FormaOplaty: String
	ZnachenieUsloviiaOgranicheniia: Double
	SegmentNomenklaturyOgranicheniiaKey: Guid
	PeriodNakopleniia: String
	TipSravneniia: String
	GruppaPolzovateleiKey: Guid
	DneiDoDniaRozhdeniia: Int16
	DneiPosleDniaRozhdeniia: Int16
	PoriadkovyiNomerProdazhiKraten: Int64
	UchityvatPrimenennyeSkidkiVRamkakhPosledovatelnogoPrimeneniia: Boolean
	UchityvatTekushchiiChekVNakopleniiakh: Boolean
	VremiaDeistviia: [CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviiaRowType!]
	Poluchateli: [CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateliRowType!]
	KomplektPokupki: [CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkiRowType!]
}
input CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviiaInput {
	Key: Guid!
	LineNumber: Int64!
	DenNedeli: String
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
}
type CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia {
	Key: Guid!
	LineNumber: Int64!
	DenNedeli: String
	VremiaNachala: DateTime
	VremiaOkonchaniia: DateTime
}
input CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateliInput {
	Key: Guid!
	LineNumber: Int64!
	Poluchatel: String
	PoluchatelType: String
}
type CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli {
	Key: Guid!
	LineNumber: Int64!
	Poluchatel: String
	PoluchatelType: String
}
input CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkiInput {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	TypeKey: Guid
	SupplierKey: Guid
	ProizvoditelKey: Guid
}
type CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki {
	Key: Guid!
	LineNumber: Int64!
	ItemKey: Guid
	Quantity: Double
	TypeKey: Guid
	SupplierKey: Guid
	ProizvoditelKey: Guid
}
input OutPayInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaVzaimoraschetovPodotchetnikaKey: Guid
	ValiutaDokumentaKey: Guid
	VidVydachiDenezhnykhSredstv: String
	OperationType: String
	Vydat: String
	DataPogasheniiaAvansa: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	ZaiavkaNaRaskhodovanieSredstvKey: Guid
	KassaKey: Guid
	Comment: String
	Kontragent: String
	NumberKKT: Int16
	ObieiavlenieNaVznosNalichnymiKey: Guid
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	Osnovanie: String
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	PoDokumentu: String
	PodrazdelenieKey: Guid
	Prilozhenie: String
	RaschetnyiDokumentKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	Pochta: String
	Telefon: String
	ProbitChekNaKKT: Boolean
	KassaKKMKey: Guid
	GungNumber: Int16
	NastroikaRMKKey: Guid
	ExtendedPayments: [DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhaRowTypeInput!]
	Payments: [DocumentRaskhodnyiKassovyiOrderOplataRowTypeInput!]
	Goods: [DocumentRaskhodnyiKassovyiOrderTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	KontragentType: String
}
type OutPay {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaVzaimoraschetovPodotchetnikaKey: Guid
	ValiutaDokumentaKey: Guid
	VidVydachiDenezhnykhSredstv: String
	OperationType: String
	Vydat: String
	DataPogasheniiaAvansa: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	ZaiavkaNaRaskhodovanieSredstvKey: Guid
	KassaKey: Guid
	Comment: String
	Kontragent: String
	NumberKKT: Int16
	ObieiavlenieNaVznosNalichnymiKey: Guid
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	Osnovanie: String
	OtvetstvennyiKey: Guid
	OtrazhenoVOperUchete: Boolean
	PoDokumentu: String
	PodrazdelenieKey: Guid
	Prilozhenie: String
	RaschetnyiDokumentKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	Pochta: String
	Telefon: String
	ProbitChekNaKKT: Boolean
	KassaKKMKey: Guid
	GungNumber: Int16
	NastroikaRMKKey: Guid
	ExtendedPayments: [DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhaRowType!]
	Payments: [DocumentRaskhodnyiKassovyiOrderOplataRowType!]
	Goods: [DocumentRaskhodnyiKassovyiOrderTovaryRowType!]
	DokumentOsnovanieType: String
	KontragentType: String
}
input DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhaInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
type DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	DokumentPlanirovaniiaPlatezhaKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	KursVzaimoraschetovPlan: Double
	ProektKey: Guid
	Sdelka: String
	StavkaNDS: String
	TypeOfMovingMoneyKey: Guid
	SummaVzaimoraschetov: Double
	SummaNDS: Double
	Sum: Double
	SummaPlatezhaPlan: Double
	SdelkaType: String
}
input DocumentRaskhodnyiKassovyiOrderOplataInput {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
type DocumentRaskhodnyiKassovyiOrderOplata {
	Key: Guid!
	LineNumber: Int64!
	TipOplaty: String
	Sum: Double
}
input DocumentRaskhodnyiKassovyiOrderTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SummaSkidki: Double
	VidTovaraKKT: String
	TipOplatyTovaraKKT: String
	SummaOsn: Double
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	SummaOpl: Double
}
type DocumentRaskhodnyiKassovyiOrderTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SummaSkidki: Double
	VidTovaraKKT: String
	TipOplatyTovaraKKT: String
	SummaOsn: Double
	Komitent: String
	TelefonKomitenta: String
	INNkomitenta: String
	SummaOpl: Double
}
input DocumentSchetNaOplatuPostavshchikaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	VremiaNapominaniia: DateTime
	DataVkhodiashchegoDokumenta: DateTime
	DataOplaty: DateTime
	DataPostupleniia: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KontaktnoeLitsoKey: Guid
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NapomnitOSobytii: Boolean
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	StrukturnaiaEdinitsa: String
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	Goods: [DocumentSchetNaOplatuPostavshchikaTovaryRowTypeInput!]
	Uslugi: [DocumentSchetNaOplatuPostavshchikaUslugiRowTypeInput!]
	DokumentOsnovanieType: String
	StrukturnaiaEdinitsaType: String
}
type DocumentSchetNaOplatuPostavshchika {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	VremiaNapominaniia: DateTime
	DataVkhodiashchegoDokumenta: DateTime
	DataOplaty: DateTime
	DataPostupleniia: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KontaktnoeLitsoKey: Guid
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NapomnitOSobytii: Boolean
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	StrukturnaiaEdinitsa: String
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	Goods: [DocumentSchetNaOplatuPostavshchikaTovaryRowType!]
	Uslugi: [DocumentSchetNaOplatuPostavshchikaUslugiRowType!]
	DokumentOsnovanieType: String
	StrukturnaiaEdinitsaType: String
}
input DocumentSchetNaOplatuPostavshchikaTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
type DocumentSchetNaOplatuPostavshchikaTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
input DocumentSchetNaOplatuPostavshchikaUslugiInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentSchetNaOplatuPostavshchikaUslugi {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input DocumentReestrSpetssviazInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtpravitelKey: Guid
	OrganizatsiiaKey: Guid
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	Klienty: [DocumentReestrSpetssviazKlientyRowTypeInput!]
}
type DocumentReestrSpetssviaz {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Comment: String
	OtpravitelKey: Guid
	OrganizatsiiaKey: Guid
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	Klienty: [DocumentReestrSpetssviazKlientyRowType!]
}
input DocumentReestrSpetssviazKlientyInput {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	Adres: String
	Telefon: String
	Weight: Double
	Sum: Double
	Paket: String
	SummaPropisiu: String
	GabarityKey: Guid
}
type DocumentReestrSpetssviazKlienty {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	Adres: String
	Telefon: String
	Weight: Double
	Sum: Double
	Paket: String
	SummaPropisiu: String
	GabarityKey: Guid
}
input DocumentJournalKassovyeDokumentyInput {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	OperationType: String
	Informatsiia: String
	KassaKey: Guid
	Comment: String
	Kontragent: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	RefType: String!
	VidOperatsiiType: String
	InformatsiiaType: String
	KontragentType: String
}
type DocumentJournalKassovyeDokumenty {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	OperationType: String
	Informatsiia: String
	KassaKey: Guid
	Comment: String
	Kontragent: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	RefType: String!
	VidOperatsiiType: String
	InformatsiiaType: String
	KontragentType: String
}
input InitialInstanceInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OperationType: String
	DataVkhodiashchegoDokumenta: DateTime
	Comment: String
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	OformitProdazhiDatoiDokumentaProdazhi: Boolean
	Vzaimoraschety: [DocumentVvodNachalnykhOstatkovVzaimoraschetyRowTypeInput!]
	Goods: [DocumentVvodNachalnykhOstatkovTovaryRowTypeInput!]
}
type InitialInstance {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OperationType: String
	DataVkhodiashchegoDokumenta: DateTime
	Comment: String
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	OformitProdazhiDatoiDokumentaProdazhi: Boolean
	Vzaimoraschety: [DocumentVvodNachalnykhOstatkovVzaimoraschetyRowType!]
	Goods: [DocumentVvodNachalnykhOstatkovTovaryRowType!]
}
input DocumentVvodNachalnykhOstatkovVzaimoraschetyInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	Sum: Double
}
type DocumentVvodNachalnykhOstatkovVzaimoraschety {
	Key: Guid!
	LineNumber: Int64!
	DogovorKontragentaKey: Guid
	KontragentKey: Guid
	Sum: Double
}
input DocumentVvodNachalnykhOstatkovTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentOprikhodovaniiaKey: Guid
	DokumentProdazhiKey: Guid
	KachestvoKey: Guid
	Quantity: Double
	Comment: String
	NDSVkliuchenVStoimost: Boolean
	ItemKey: Guid
	NomerGTDKey: Guid
	Pasport: String
	ProtsentRoznichnoiNatsenki: Double
	ProtsentSkidkiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	StavkaNDSProdazhi: String
	StatusPartii: String
	StranaProiskhozhdeniiaKey: Guid
	Sum: Double
	SummaNDS: Double
	SummaNDSProdazhi: Double
	SummaProdazhi: Double
	SummaProdazhiBezSkidok: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaProdazhi: Double
	StatusRaskhoda: String
}
type DocumentVvodNachalnykhOstatkovTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentOprikhodovaniiaKey: Guid
	DokumentProdazhiKey: Guid
	KachestvoKey: Guid
	Quantity: Double
	Comment: String
	NDSVkliuchenVStoimost: Boolean
	ItemKey: Guid
	NomerGTDKey: Guid
	Pasport: String
	ProtsentRoznichnoiNatsenki: Double
	ProtsentSkidkiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	StavkaNDSProdazhi: String
	StatusPartii: String
	StranaProiskhozhdeniiaKey: Guid
	Sum: Double
	SummaNDS: Double
	SummaNDSProdazhi: Double
	SummaProdazhi: Double
	SummaProdazhiBezSkidok: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
	TsenaProdazhi: Double
	StatusRaskhoda: String
}
input PostingInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	InventarizatsiiaTovarovNaSkladeKey: Guid
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	Osnovanie: String
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	SummaDokumentaRegl: Double
	TipDokumenta: String
	TipTsenKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	VystavkaOstatki: Boolean
	Goods: [DocumentOprikhodovanieTovarovTovaryRowTypeInput!]
	Sertifikaty: [DocumentOprikhodovanieTovarovSertifikatyRowTypeInput!]
	DokumentOsnovanieType: String
}
type Posting {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	InventarizatsiiaTovarovNaSkladeKey: Guid
	KolichestvoDokumenta: Int64
	Comment: String
	OrganizatsiiaKey: Guid
	Osnovanie: String
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	SumOfDocument: Double
	SummaDokumentaRegl: Double
	TipDokumenta: String
	TipTsenKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	VystavkaOstatki: Boolean
	Goods: [DocumentOprikhodovanieTovarovTovaryRowType!]
	Sertifikaty: [DocumentOprikhodovanieTovarovSertifikatyRowType!]
	DokumentOsnovanieType: String
}
input DocumentOprikhodovanieTovarovTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaRegl: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
}
type DocumentOprikhodovanieTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaRegl: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
}
input DocumentOprikhodovanieTovarovSertifikatyInput {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
type DocumentOprikhodovanieTovarovSertifikaty {
	Key: Guid!
	LineNumber: Int64!
	SertifikatKey: Guid
	Sum: Double
}
input CatalogKomplektyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
}
type CatalogKomplekty {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
}
input DocumentPereotsenkaTovarovPriniatykhNaKomissiiuInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	NastroikiZapolneniiaBase64Data: Binary
	Goods: [DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
type DocumentPereotsenkaTovarovPriniatykhNaKomissiiu {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	NastroikiZapolneniiaBase64Data: Binary
	Goods: [DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovaryRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
	NastroikiZapolneniiaType: String
	NastroikiZapolneniia: Stream
}
input DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SummaStaraia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaZaGramm: Double
}
type DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	SummaStaraia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaZaGramm: Double
}
input DocumentElektronnoePismoInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	VidTekstaPisma: String
	GruppaUchetnoiZapisiKey: Guid
	DataOtpravleniia: DateTime
	DataTransporta: DateTime
	EstVlozheniia: Boolean
	IdentifikatorPisma: String
	ImiaKompiuteraRedaktirovaniiaKhTMLTeksta: String
	ImiaFailaRedaktirovaniiaKhTMLTeksta: String
	Comment: String
	Komu: String
	KomuPredstavlenie: String
	Kopii: String
	KopiiPredstavlenie: String
	NeRassmotreno: Boolean
	Osnovanie: String
	Otvet: Boolean
	OtvetstvennyiKey: Guid
	OtpravitelAdresElektronnoiPochty: String
	OtpravitelImia: String
	OtpravitelPredstavlenie: String
	Pereadresatsiia: Boolean
	PochtovoeSoobshchenieBase64Data: Binary
	Predmet: String
	RassmotretPosle: DateTime
	SkrytyeKopii: String
	SostoianiePotomkaPisma: String
	StatusPisma: String
	TekstPisma: String
	Tema: String
	UchetnaiaZapisKey: Guid
	KomuTCh: [DocumentElektronnoePismoKomuTChRowTypeInput!]
	KopiiTCh: [DocumentElektronnoePismoKopiiTChRowTypeInput!]
	SkrytyeKopiiTCh: [DocumentElektronnoePismoSkrytyeKopiiTChRowTypeInput!]
	OsnovanieType: String
	PochtovoeSoobshchenieType: String
	PredmetType: String
	PochtovoeSoobshchenie: Stream
}
type DocumentElektronnoePismo {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	VidTekstaPisma: String
	GruppaUchetnoiZapisiKey: Guid
	DataOtpravleniia: DateTime
	DataTransporta: DateTime
	EstVlozheniia: Boolean
	IdentifikatorPisma: String
	ImiaKompiuteraRedaktirovaniiaKhTMLTeksta: String
	ImiaFailaRedaktirovaniiaKhTMLTeksta: String
	Comment: String
	Komu: String
	KomuPredstavlenie: String
	Kopii: String
	KopiiPredstavlenie: String
	NeRassmotreno: Boolean
	Osnovanie: String
	Otvet: Boolean
	OtvetstvennyiKey: Guid
	OtpravitelAdresElektronnoiPochty: String
	OtpravitelImia: String
	OtpravitelPredstavlenie: String
	Pereadresatsiia: Boolean
	PochtovoeSoobshchenieBase64Data: Binary
	Predmet: String
	RassmotretPosle: DateTime
	SkrytyeKopii: String
	SostoianiePotomkaPisma: String
	StatusPisma: String
	TekstPisma: String
	Tema: String
	UchetnaiaZapisKey: Guid
	KomuTCh: [DocumentElektronnoePismoKomuTChRowType!]
	KopiiTCh: [DocumentElektronnoePismoKopiiTChRowType!]
	SkrytyeKopiiTCh: [DocumentElektronnoePismoSkrytyeKopiiTChRowType!]
	OsnovanieType: String
	PochtovoeSoobshchenieType: String
	PredmetType: String
	PochtovoeSoobshchenie: Stream
}
input DocumentElektronnoePismoKomuTChInput {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
type DocumentElektronnoePismoKomuTCh {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
input DocumentElektronnoePismoKopiiTChInput {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
type DocumentElektronnoePismoKopiiTCh {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
input DocumentElektronnoePismoSkrytyeKopiiTChInput {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
type DocumentElektronnoePismoSkrytyeKopiiTCh {
	Key: Guid!
	LineNumber: Int64!
	AdresElektronnoiPochty: String
	Predstavlenie: String
}
input CatalogGruppyDefektovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
}
type CatalogGruppyDefektov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
}
input CatalogfmAnketaKlientaBenefitsariiaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	Familiia: String
	Imia: String
	Otchestvo: String
	KartochkaKontragentaKey: Guid
	Deistvuiushchii: Boolean
	DannyeKontragenta: [CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentaRowTypeInput!]
}
type CatalogfmAnketaKlientaBenefitsariia {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	Familiia: String
	Imia: String
	Otchestvo: String
	KartochkaKontragentaKey: Guid
	Deistvuiushchii: Boolean
	DannyeKontragenta: [CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentaRowType!]
}
input CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentaInput {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
type CatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta {
	Key: Guid!
	LineNumber: Int64!
	Kliuch: String
	Znachenie: String
	ZnachenieType: String
}
input CatalogTsenovyeGruppyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
}
type CatalogTsenovyeGruppy {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
}
input CatalogPravilaTsenoobrazovaniiaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	VidTsenKey: Guid
	TsenaVkliuchaetNDS: Boolean
	Status: String
	TsenovyeGruppy: [CatalogPravilaTsenoobrazovaniiaTsenovyeGruppyRowTypeInput!]
}
type CatalogPravilaTsenoobrazovaniia {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	VidTsenKey: Guid
	TsenaVkliuchaetNDS: Boolean
	Status: String
	TsenovyeGruppy: [CatalogPravilaTsenoobrazovaniiaTsenovyeGruppyRowType!]
}
input CatalogPravilaTsenoobrazovaniiaTsenovyeGruppyInput {
	Key: Guid!
	LineNumber: Int64!
	TsenovaiaGruppaKey: Guid
	VidTsenKey: Guid
}
type CatalogPravilaTsenoobrazovaniiaTsenovyeGruppy {
	Key: Guid!
	LineNumber: Int64!
	TsenovaiaGruppaKey: Guid
	VidTsenKey: Guid
}
input DocumentObieiavlenieNaVznosNalichnymiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DataOplaty: DateTime
	DokumentOsnovanie: String
	KassaKey: Guid
	Comment: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	DokumentOsnovanieType: String
}
type DocumentObieiavlenieNaVznosNalichnymi {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	DataOplaty: DateTime
	DokumentOsnovanie: String
	KassaKey: Guid
	Comment: String
	Oplacheno: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	TypeOfMovingMoneyKey: Guid
	SumOfDocument: Double
	SchetOrganizatsiiKey: Guid
	TipDokumenta: String
	DokumentOsnovanieType: String
}
input CatalogValiutyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	NaimenovaniePolnoe: String
	ParametryPropisiNaRusskom: String
	SposobUstanovkiKursa: String
	BIdentifikator: String
	BNomerVersii: String
}
type CatalogValiuty {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	NaimenovaniePolnoe: String
	ParametryPropisiNaRusskom: String
	SposobUstanovkiKursa: String
	BIdentifikator: String
	BNomerVersii: String
}
input DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetKontragentaKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DataVkhodiashchegoDokumenta: DateTime
	DliaKontroliaUnikalnosti: String
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPeremeshcheniiaKey: Guid
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	Koef: Double
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	RegistrirovatTseny: Boolean
	RegistrirovatTsenyPostavshchika: Boolean
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	UsloviiaOplatyKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	VystavkaOstatki: Boolean
	Goods: [DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovaryRowTypeInput!]
	Uslugi: [DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugiRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetKontragentaKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DataVkhodiashchegoDokumenta: DateTime
	DliaKontroliaUnikalnosti: String
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentPeremeshcheniiaKey: Guid
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	Koef: Double
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	RegistrirovatTseny: Boolean
	RegistrirovatTsenyPostavshchika: Boolean
	Sdelka: String
	DepartmentKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipSkidkiNatsenkiKey: Guid
	TipTsenKey: Guid
	UsloviiaOplatyKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	VystavkaOstatki: Boolean
	Goods: [DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovaryRowType!]
	Uslugi: [DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugiRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaIzmereniiaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	Koef: Double
	Koeffitsient: Double
	ItemKey: Guid
	Pasport: String
	ProektKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
}
type DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaIzmereniiaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	Koef: Double
	Koeffitsient: Double
	ItemKey: Guid
	Pasport: String
	ProektKey: Guid
	ProtsentRoznichnoiNatsenki: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	RetailCost: Double
	TsenaVRoznitseGr: Double
}
input DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugiInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input CatalogKassyKKMInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	KorrespondiruiushchiiSchet: String
}
type CatalogKassyKKM {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	KorrespondiruiushchiiSchet: String
}
input ProbeInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	MetallKey: Guid
	Proba: Double
	BIdentifikator: String
}
type Probe {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	MetallKey: Guid
	Proba: Double
	BIdentifikator: String
}
input CatalogGruppyDostupaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	ProfilKey: Guid
	PolzovatelKey: Guid
	Opisanie: String
	OtvetstvennyiKey: Guid
	OsnovnaiaGruppaDostupaPostavliaemogoProfilia: Boolean
	RazreshenieNaProsmotrZakupochnoiTseny: Boolean
	Polzovateli: [CatalogGruppyDostupaPolzovateliRowTypeInput!]
	VidyDostupa: [CatalogGruppyDostupaVidyDostupaRowTypeInput!]
	ZnacheniiaDostupa: [CatalogGruppyDostupaZnacheniiaDostupaRowTypeInput!]
	DostupPoPodsistemam: [CatalogGruppyDostupaDostupPoPodsistemamRowTypeInput!]
}
type CatalogGruppyDostupa {
	Key: Guid!
	DataVersion: String
	Description: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	ProfilKey: Guid
	PolzovatelKey: Guid
	Opisanie: String
	OtvetstvennyiKey: Guid
	OsnovnaiaGruppaDostupaPostavliaemogoProfilia: Boolean
	RazreshenieNaProsmotrZakupochnoiTseny: Boolean
	Polzovateli: [CatalogGruppyDostupaPolzovateliRowType!]
	VidyDostupa: [CatalogGruppyDostupaVidyDostupaRowType!]
	ZnacheniiaDostupa: [CatalogGruppyDostupaZnacheniiaDostupaRowType!]
	DostupPoPodsistemam: [CatalogGruppyDostupaDostupPoPodsistemamRowType!]
}
input CatalogGruppyDostupaPolzovateliInput {
	Key: Guid!
	LineNumber: Int64!
	Polzovatel: String
	PolzovatelType: String
}
type CatalogGruppyDostupaPolzovateli {
	Key: Guid!
	LineNumber: Int64!
	Polzovatel: String
	PolzovatelType: String
}
input CatalogGruppyDostupaVidyDostupaInput {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	VseRazresheny: Boolean
	VidDostupaType: String
}
type CatalogGruppyDostupaVidyDostupa {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	VseRazresheny: Boolean
	VidDostupaType: String
}
input CatalogGruppyDostupaZnacheniiaDostupaInput {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	ZnachenieDostupa: String
	VidDostupaType: String
	ZnachenieDostupaType: String
}
type CatalogGruppyDostupaZnacheniiaDostupa {
	Key: Guid!
	LineNumber: Int64!
	VidDostupa: String
	ZnachenieDostupa: String
	VidDostupaType: String
	ZnachenieDostupaType: String
}
input CatalogGruppyDostupaDostupPoPodsistemamInput {
	Key: Guid!
	LineNumber: Int64!
	ImiaPodsistemy: String
	ImiaObieekta: String
	Prosmotr: Boolean
	Redaktirovanie: Boolean
	Pechat: Boolean
	PokazVersii: Boolean
}
type CatalogGruppyDostupaDostupPoPodsistemam {
	Key: Guid!
	LineNumber: Int64!
	ImiaPodsistemy: String
	ImiaObieekta: String
	Prosmotr: Boolean
	Redaktirovanie: Boolean
	Pechat: Boolean
	PokazVersii: Boolean
}
input CatalogVidyKontaktnoiInformatsiiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidObieektaKontaktnoiInformatsii: String
	Tip: String
}
type CatalogVidyKontaktnoiInformatsii {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidObieektaKontaktnoiInformatsii: String
	Tip: String
}
input CatalogNomenklaturnyeGruppyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	BazovaiaEdinitsaIzmereniiaKey: Guid
	StavkaNDS: String
}
type CatalogNomenklaturnyeGruppy {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	BazovaiaEdinitsaIzmereniiaKey: Guid
	StavkaNDS: String
}
input DocumentReestrSchetovInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DataAkkreditiva: DateTime
	DataZakrytiia: DateTime
	DokumentOsnovanieKey: Guid
	IspolniaiushchiiBankKey: Guid
	Comment: String
	KontragentKey: Guid
	NomerAkkreditiva: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	SummaAkkreditiva: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	Reestr: [DocumentReestrSchetovReestrRowTypeInput!]
}
type DocumentReestrSchetov {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DataAkkreditiva: DateTime
	DataZakrytiia: DateTime
	DokumentOsnovanieKey: Guid
	IspolniaiushchiiBankKey: Guid
	Comment: String
	KontragentKey: Guid
	NomerAkkreditiva: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	SummaAkkreditiva: Double
	SchetKontragentaKey: Guid
	SchetOrganizatsiiKey: Guid
	Reestr: [DocumentReestrSchetovReestrRowType!]
}
input DocumentReestrSchetovReestrInput {
	Key: Guid!
	LineNumber: Int64!
	VidTransporta: String
	DataOtgruzki: DateTime
	NomerDokumenta: String
	Sum: Double
}
type DocumentReestrSchetovReestr {
	Key: Guid!
	LineNumber: Int64!
	VidTransporta: String
	DataOtgruzki: DateTime
	NomerDokumenta: String
	Sum: Double
}
input DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DogovorKontragentaKey: Guid
	Comment: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	Goods: [DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovaryRowTypeInput!]
	SdelkaType: String
}
type DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	DogovorKontragentaKey: Guid
	Comment: String
	KontragentKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sdelka: String
	Goods: [DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovaryRowType!]
	SdelkaType: String
}
input DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesUchet: Double
	Quantity: Int64
	KolichestvoUchet: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
type DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesUchet: Double
	Quantity: Int64
	KolichestvoUchet: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
input CatalogKlassifikatorStranMiraInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	NaimenovaniePolnoe: String
	KodAlfa2: String
}
type CatalogKlassifikatorStranMira {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	NaimenovaniePolnoe: String
	KodAlfa2: String
}
input CatalogKlassifikatorEdinitsIzmereniiaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	NaimenovaniePolnoe: String
	MezhdunarodnoeSokrashchenie: String
	TipEdinitsyIzmereniia: String
	BIdentifikator: String
}
type CatalogKlassifikatorEdinitsIzmereniia {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	NaimenovaniePolnoe: String
	MezhdunarodnoeSokrashchenie: String
	TipEdinitsyIzmereniia: String
	BIdentifikator: String
}
input CatalogNastroikiRMKInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidAvtorizatsii: String
	VidOplaty2: String
	VidOplaty2PoUmolchaniiuKey: Guid
	VidOplaty3: String
	VidOplaty3PoUmolchaniiuKey: Guid
	VidOplaty4: String
	VidOplaty4PoUmolchaniiuKey: Guid
	VyvoditSpisokSertifikatovPriOplate: Boolean
	ZakryvatPodborPriVyboreTovara: Boolean
	ZapretitProdazhuNizheSebestoimosti: Boolean
	IntervalProverkiOtvetaNaZaprosProverkiSertifikata: Int16
	IspolzovatAvtomaticheskieSkidki: Boolean
	IspolzovatVidOplaty2: Boolean
	IspolzovatVidOplaty3: Boolean
	IspolzovatVidOplaty4: Boolean
	IspolzovatVidOplatySertifikatom: Boolean
	IspolzovatObmenStarykhIzdeliiNaNovye: Boolean
	IspolzovatOplatuSertifikatami: Boolean
	IspolzovatRuchnyeSkidki: Boolean
	KontrolirovatOstatokPriProvedenii: Boolean
	MaksimalnoeVremiaOzhidaniiaOtvetaNaZaprosProverkiSertifikata: Int16
	NaimenovanieSummyDoplatyZaTovaryPoSertifikatam: String
	NastroikiObmenaDannymiKey: Guid
	OrganizatsiiaKey: Guid
	OtobrazhatKlaviaturuPriZapuske: Boolean
	PechatatVtoroiEkzempliarTovarnogoCheka: Boolean
	PechatatVtoroiEkzempliarTovarnogoChekaNaPrintere: Boolean
	PechatatImiaKassira: Boolean
	PechatatNazvanieDiskontnoiKarty: Boolean
	PechatatNazvaniePlatezhnoiKarty: Boolean
	PechatatNumeratsiiuPozitsii: Boolean
	PechatatSummuNachislennogoBonusa: Boolean
	PechatatTovarnyiChekNaPrintere: Boolean
	PechatatTovarnyiChekPriRegistratsiiProdazhi: Boolean
	PechatatShtrikhkodPriRegistratsiiProdazhi: Boolean
	PoriadokOkrugleniiaSumm: String
	PriVozvrateRaspechatyvatPaketDokumentov: Boolean
	PriVozvrateRaspechatyvatPaketDokumentovNaPrintere: Boolean
	ProtsentOgranicheniiaSummySkidki: Double
	RezhimProverkiSertifikatov: Int16
	DepartmentKey: Guid
	SposobPrimeneniiaNeskolkikhSkidok: String
	ShablonChekaKey: Guid
	IspolzovatAvtomaticheskieSkidkiMarketing: Boolean
	KassaKKMKey: Guid
	IspolzovatProdazhuUslug: Boolean
	TrebovatVvodPasportnykhDannykh: Boolean
	ZapolniatAnketuKlientaPriProdazhe: Boolean
	ZapolniatAnketuKlientaPriSkupke: Boolean
	PechatatTovarnyiChekPriSkupke: Boolean
	PechatatSkupochnuiuKvitantsiiuPriSkupke: Boolean
	KolichestvoEkzempliarovTovarnogoChekaPriSkupke: Int16
	KolichestvoEkzempliarovSkupochnoiKvitantsiiPriSkupke: Int16
	VidProbitiiaSkupki: String
	RazreshitZamenuKarty: Boolean
	RazreshitVozvratNeDenVDenPoCheku: Boolean
	PeredachaNomenklaturyKlientuPriPolnomPogasheniiRassrochki: Boolean
	NeVyvoditInformatsiiuOKomitenteVChek: Boolean
	RazreshitVozvratNalichnymiPriProdazhePoBeznalichnomuRaschetu: Boolean
	PoriadokPrimeneniiaSkidok: [CatalogNastroikiRMKPoriadokPrimeneniiaSkidokRowTypeInput!]
	SostavNaimenovaniia: [CatalogNastroikiRMKSostavNaimenovaniiaRowTypeInput!]
}
type CatalogNastroikiRMK {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	VidAvtorizatsii: String
	VidOplaty2: String
	VidOplaty2PoUmolchaniiuKey: Guid
	VidOplaty3: String
	VidOplaty3PoUmolchaniiuKey: Guid
	VidOplaty4: String
	VidOplaty4PoUmolchaniiuKey: Guid
	VyvoditSpisokSertifikatovPriOplate: Boolean
	ZakryvatPodborPriVyboreTovara: Boolean
	ZapretitProdazhuNizheSebestoimosti: Boolean
	IntervalProverkiOtvetaNaZaprosProverkiSertifikata: Int16
	IspolzovatAvtomaticheskieSkidki: Boolean
	IspolzovatVidOplaty2: Boolean
	IspolzovatVidOplaty3: Boolean
	IspolzovatVidOplaty4: Boolean
	IspolzovatVidOplatySertifikatom: Boolean
	IspolzovatObmenStarykhIzdeliiNaNovye: Boolean
	IspolzovatOplatuSertifikatami: Boolean
	IspolzovatRuchnyeSkidki: Boolean
	KontrolirovatOstatokPriProvedenii: Boolean
	MaksimalnoeVremiaOzhidaniiaOtvetaNaZaprosProverkiSertifikata: Int16
	NaimenovanieSummyDoplatyZaTovaryPoSertifikatam: String
	NastroikiObmenaDannymiKey: Guid
	OrganizatsiiaKey: Guid
	OtobrazhatKlaviaturuPriZapuske: Boolean
	PechatatVtoroiEkzempliarTovarnogoCheka: Boolean
	PechatatVtoroiEkzempliarTovarnogoChekaNaPrintere: Boolean
	PechatatImiaKassira: Boolean
	PechatatNazvanieDiskontnoiKarty: Boolean
	PechatatNazvaniePlatezhnoiKarty: Boolean
	PechatatNumeratsiiuPozitsii: Boolean
	PechatatSummuNachislennogoBonusa: Boolean
	PechatatTovarnyiChekNaPrintere: Boolean
	PechatatTovarnyiChekPriRegistratsiiProdazhi: Boolean
	PechatatShtrikhkodPriRegistratsiiProdazhi: Boolean
	PoriadokOkrugleniiaSumm: String
	PriVozvrateRaspechatyvatPaketDokumentov: Boolean
	PriVozvrateRaspechatyvatPaketDokumentovNaPrintere: Boolean
	ProtsentOgranicheniiaSummySkidki: Double
	RezhimProverkiSertifikatov: Int16
	DepartmentKey: Guid
	SposobPrimeneniiaNeskolkikhSkidok: String
	ShablonChekaKey: Guid
	IspolzovatAvtomaticheskieSkidkiMarketing: Boolean
	KassaKKMKey: Guid
	IspolzovatProdazhuUslug: Boolean
	TrebovatVvodPasportnykhDannykh: Boolean
	ZapolniatAnketuKlientaPriProdazhe: Boolean
	ZapolniatAnketuKlientaPriSkupke: Boolean
	PechatatTovarnyiChekPriSkupke: Boolean
	PechatatSkupochnuiuKvitantsiiuPriSkupke: Boolean
	KolichestvoEkzempliarovTovarnogoChekaPriSkupke: Int16
	KolichestvoEkzempliarovSkupochnoiKvitantsiiPriSkupke: Int16
	VidProbitiiaSkupki: String
	RazreshitZamenuKarty: Boolean
	RazreshitVozvratNeDenVDenPoCheku: Boolean
	PeredachaNomenklaturyKlientuPriPolnomPogasheniiRassrochki: Boolean
	NeVyvoditInformatsiiuOKomitenteVChek: Boolean
	RazreshitVozvratNalichnymiPriProdazhePoBeznalichnomuRaschetu: Boolean
	PoriadokPrimeneniiaSkidok: [CatalogNastroikiRMKPoriadokPrimeneniiaSkidokRowType!]
	SostavNaimenovaniia: [CatalogNastroikiRMKSostavNaimenovaniiaRowType!]
}
input CatalogNastroikiRMKPoriadokPrimeneniiaSkidokInput {
	Key: Guid!
	LineNumber: Int64!
	UslovieSkidki: String
}
type CatalogNastroikiRMKPoriadokPrimeneniiaSkidok {
	Key: Guid!
	LineNumber: Int64!
	UslovieSkidki: String
}
input CatalogNastroikiRMKSostavNaimenovaniiaInput {
	Key: Guid!
	LineNumber: Int64!
	SimvolyDo: String
	SimvolyPosle: String
	ElementNaimenovaniia: String
}
type CatalogNastroikiRMKSostavNaimenovaniia {
	Key: Guid!
	LineNumber: Int64!
	SimvolyDo: String
	SimvolyPosle: String
	ElementNaimenovaniia: String
}
input CatalogKharakteristikiNomenklaturyInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	PolnoeNaimenovanie: String
	PerechenKamnei: String
	PerechenKamneiTranslit: String
	BIdentifikator: String
	BNomerVersii: String
	Spetsifikatsiia: [CatalogKharakteristikiNomenklaturySpetsifikatsiiaRowTypeInput!]
}
type CatalogKharakteristikiNomenklatury {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	PolnoeNaimenovanie: String
	PerechenKamnei: String
	PerechenKamneiTranslit: String
	BIdentifikator: String
	BNomerVersii: String
	Spetsifikatsiia: [CatalogKharakteristikiNomenklaturySpetsifikatsiiaRowType!]
}
input CatalogKharakteristikiNomenklaturySpetsifikatsiiaInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
}
type CatalogKharakteristikiNomenklaturySpetsifikatsiia {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	GruppaDefektaKey: Guid
	GruppaTsvetaKey: Guid
	KamenKey: Guid
	Quantity: Double
	ItemKey: Guid
	Razmer1: Double
	Razmer2: Double
	Razmer3: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
}
input DocumentOtborTovarovInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetOrganizatsiiKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	ZakazKlientaKey: Guid
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	Soglasovano: Boolean
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	SkhemaRealizatsiiKey: Guid
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	Goods: [DocumentOtborTovarovTovaryRowTypeInput!]
	TovaryKOtboru: [DocumentOtborTovarovTovaryKOtboruRowTypeInput!]
	DokumentOsnovanieType: String
}
type DocumentOtborTovarov {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetOrganizatsiiKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	ZakazKlientaKey: Guid
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
	Soglasovano: Boolean
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	SkhemaRealizatsiiKey: Guid
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	Goods: [DocumentOtborTovarovTovaryRowType!]
	TovaryKOtboru: [DocumentOtborTovarovTovaryKOtboruRowType!]
	DokumentOsnovanieType: String
}
input DocumentOtborTovarovTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumManualDiscount: Double
	SumAutoDiscount: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentOtborTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	SumManualDiscount: Double
	SumAutoDiscount: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input DocumentOtborTovarovTovaryKOtboruInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
}
type DocumentOtborTovarovTovaryKOtboru {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
}
input CatalogSposobyDostavkiTovaraInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogSposobyDostavkiTovara {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input CatalogPodrazdeleniiaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	DeletionMark: Boolean
}
type CatalogPodrazdeleniia {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	DeletionMark: Boolean
}
input DocumentJournalPreiskurantyInput {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	KamenKey: Guid
	Komentarii: String
	RassevKey: Guid
	FormaOgrankiKey: Guid
	RefType: String!
}
type DocumentJournalPreiskuranty {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	KamenKey: Guid
	Komentarii: String
	RassevKey: Guid
	FormaOgrankiKey: Guid
	RefType: String!
}
input CatalogRelizyIuvelirnykhSalonovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	NomerVersiiKonfiguratsii: String
	TipFailaPostavki: String
	KhranilishcheFailaPostavkiBase64Data: Binary
	ObnovliaemyeRelizy: [CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizyRowTypeInput!]
	KhranilishcheFailaPostavkiType: String
	KhranilishcheFailaPostavki: Stream
}
type CatalogRelizyIuvelirnykhSalonov {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	NomerVersiiKonfiguratsii: String
	TipFailaPostavki: String
	KhranilishcheFailaPostavkiBase64Data: Binary
	ObnovliaemyeRelizy: [CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizyRowType!]
	KhranilishcheFailaPostavkiType: String
	KhranilishcheFailaPostavki: Stream
}
input CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizyInput {
	Key: Guid!
	LineNumber: Int64!
	RelizIuvelirnogoSalonaKey: Guid
}
type CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy {
	Key: Guid!
	LineNumber: Int64!
	RelizIuvelirnogoSalonaKey: Guid
}
input DocumentOtchetKomissioneraOProdazhakhInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NomenklaturnaiaGruppaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	ProtsentKomissionnogoVoznagrazhdeniia: Double
	Sdelka: String
	SposobRaschetaKomissionnogoVoznagrazhdeniia: String
	StavkaNDSVoznagrazhdeniia: String
	StatiaZatratKey: Guid
	SummaVkliuchaetNDS: Boolean
	SummaVoznagrazhdeniia: Double
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UderzhatKomissionnoeVoznagrazhdenie: Boolean
	UsloviiaOplatyKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	DenezhnyeSredstva: [DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvaRowTypeInput!]
	Goods: [DocumentOtchetKomissioneraOProdazhakhTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type DocumentOtchetKomissioneraOProdazhakh {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NomenklaturnaiaGruppaKey: Guid
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	ProtsentKomissionnogoVoznagrazhdeniia: Double
	Sdelka: String
	SposobRaschetaKomissionnogoVoznagrazhdeniia: String
	StavkaNDSVoznagrazhdeniia: String
	StatiaZatratKey: Guid
	SummaVkliuchaetNDS: Boolean
	SummaVoznagrazhdeniia: Double
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UderzhatKomissionnoeVoznagrazhdenie: Boolean
	UsloviiaOplatyKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	DenezhnyeSredstva: [DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvaRowType!]
	Goods: [DocumentOtchetKomissioneraOProdazhakhTovaryRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvaInput {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
type DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
input DocumentOtchetKomissioneraOProdazhakhTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDS: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaNDSPeredachi: Double
	SummaPeredachi: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPeredachi: Double
	DataSchetaFakturyKomissionera: DateTime
	PokupatelKey: Guid
	NomerSchetaFakturyKomissionera: String
}
type DocumentOtchetKomissioneraOProdazhakhTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	StavkaNDS: String
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDS: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaNDSPeredachi: Double
	SummaPeredachi: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPeredachi: Double
	DataSchetaFakturyKomissionera: DateTime
	PokupatelKey: Guid
	NomerSchetaFakturyKomissionera: String
}
input CatalogTovarnyeKategoriiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	NomenklaturnaiaGruppaKey: Guid
}
type CatalogTovarnyeKategorii {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	NomenklaturnaiaGruppaKey: Guid
}
input CatalogDokumentyUdostoveriaiushchieLichnostInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	KodIMNS: String
	KodPFR: String
}
type CatalogDokumentyUdostoveriaiushchieLichnost {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	KodIMNS: String
	KodPFR: String
}
input CatalogFiltryDliaElektronnykhPisemInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	Ispolzovanie: Boolean
	OperatsiiaUsloviia: Boolean
	Poriadok: Int64
	DeistviiaFiltra: [CatalogFiltryDliaElektronnykhPisemDeistviiaFiltraRowTypeInput!]
	UsloviiaFiltra: [CatalogFiltryDliaElektronnykhPisemUsloviiaFiltraRowTypeInput!]
}
type CatalogFiltryDliaElektronnykhPisem {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	OwnerKey: Guid
	DeletionMark: Boolean
	Ispolzovanie: Boolean
	OperatsiiaUsloviia: Boolean
	Poriadok: Int64
	DeistviiaFiltra: [CatalogFiltryDliaElektronnykhPisemDeistviiaFiltraRowType!]
	UsloviiaFiltra: [CatalogFiltryDliaElektronnykhPisemUsloviiaFiltraRowType!]
}
input CatalogFiltryDliaElektronnykhPisemDeistviiaFiltraInput {
	Key: Guid!
	LineNumber: Int64!
	GruppaPisemKey: Guid
	DeistvieFiltra: String
}
type CatalogFiltryDliaElektronnykhPisemDeistviiaFiltra {
	Key: Guid!
	LineNumber: Int64!
	GruppaPisemKey: Guid
	DeistvieFiltra: String
}
input CatalogFiltryDliaElektronnykhPisemUsloviiaFiltraInput {
	Key: Guid!
	LineNumber: Int64!
	ZnachenieUsloviia: String
	OtritsanieUsloviia: Boolean
	Uslovie: String
}
type CatalogFiltryDliaElektronnykhPisemUsloviiaFiltra {
	Key: Guid!
	LineNumber: Int64!
	ZnachenieUsloviia: String
	OtritsanieUsloviia: Boolean
	Uslovie: String
}
input DocumentPreiskurantTsenNaTsvKamniInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	KamenKey: Guid
	Comment: String
	RassevKey: Guid
	TipTsenKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	Tablitsa: [DocumentPreiskurantTsenNaTsvKamniTablitsaRowTypeInput!]
}
type DocumentPreiskurantTsenNaTsvKamni {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	KamenKey: Guid
	Comment: String
	RassevKey: Guid
	TipTsenKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	Tablitsa: [DocumentPreiskurantTsenNaTsvKamniTablitsaRowType!]
}
input DocumentPreiskurantTsenNaTsvKamniTablitsaInput {
	Key: Guid!
	LineNumber: Int64!
	Razmer1Do: Double
	Razmer1Ot: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	Cost: Int64
}
type DocumentPreiskurantTsenNaTsvKamniTablitsa {
	Key: Guid!
	LineNumber: Int64!
	Razmer1Do: Double
	Razmer1Ot: Double
	RassevKey: Guid
	FormaOgrankiKey: Guid
	TsvetKamniaKey: Guid
	Cost: Int64
}
input SizeInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	BIdentifikator: String
}
type Size {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	BIdentifikator: String
}
input CatalogTipyDragotsennykhMetallovInput {
	Key: Guid!
	DataVersion: String
	Code: String
	DeletionMark: Boolean
	ProbaChistoty: Double
	BIdentifikator: String
}
type CatalogTipyDragotsennykhMetallov {
	Key: Guid!
	DataVersion: String
	Code: String
	DeletionMark: Boolean
	ProbaChistoty: Double
	BIdentifikator: String
}
input DocumentTelemarketingInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtorKey: Guid
	Osnovanie: String
	Comment: String
	Vazhnost: String
	ProektKey: Guid
	SostoianieSobytiia: String
	OtvetstvennyiKey: Guid
	Tema: String
	TipovaiaAnketaKey: Guid
	VremiaNapominaniia: DateTime
	NapomnitOSobytii: Boolean
	Uchastniki: [DocumentTelemarketingUchastnikiRowTypeInput!]
	OsnovanieType: String
}
type DocumentTelemarketing {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtorKey: Guid
	Osnovanie: String
	Comment: String
	Vazhnost: String
	ProektKey: Guid
	SostoianieSobytiia: String
	OtvetstvennyiKey: Guid
	Tema: String
	TipovaiaAnketaKey: Guid
	VremiaNapominaniia: DateTime
	NapomnitOSobytii: Boolean
	Uchastniki: [DocumentTelemarketingUchastnikiRowType!]
	OsnovanieType: String
}
input DocumentTelemarketingUchastnikiInput {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	NaimenovaniePolnoe: String
	KontaktnoeLitsoKey: Guid
	Telefon: String
	RezultatObrabotkiZvonka: String
	EstInteres: Boolean
	SobytieKey: Guid
	OprosKey: Guid
	Opisanie: String
}
type DocumentTelemarketingUchastniki {
	Key: Guid!
	LineNumber: Int64!
	KontragentKey: Guid
	NaimenovaniePolnoe: String
	KontaktnoeLitsoKey: Guid
	Telefon: String
	RezultatObrabotkiZvonka: String
	EstInteres: Boolean
	SobytieKey: Guid
	OprosKey: Guid
	Opisanie: String
}
input DocumentVozvratDavalcheskogoMetallaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanieKey: Guid
	Comment: String
	KontragentKey: Guid
	ItemKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
}
type DocumentVozvratDavalcheskogoMetalla {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanieKey: Guid
	Comment: String
	KontragentKey: Guid
	ItemKey: Guid
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	DepartmentKey: Guid
}
input CatalogAdresnyeSokrashcheniiaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Sokrashchenie: String
	Uroven: Int16
}
type CatalogAdresnyeSokrashcheniia {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Sokrashchenie: String
	Uroven: Int16
}
input DocumentRassylkaAnketInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AnketaKey: Guid
	PervichnaiaRassylkaKey: Guid
	SUvedomleniemOPoluchenii: Boolean
	TekstPisma: String
	UchetnaiaZapisDliaOtpravki: String
	ElektronnyiAdresOtvetov: String
	ElektronnyiAdresOtvetovVstroennyiPochtovyiKlientKey: Guid
	Vlozheniia: [DocumentRassylkaAnketVlozheniiaRowTypeInput!]
	Poluchateli: [DocumentRassylkaAnketPoluchateliRowTypeInput!]
	UchetnaiaZapisDliaOtpravkiType: String
	ElektronnyiAdresOtvetovType: String
}
type DocumentRassylkaAnket {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AnketaKey: Guid
	PervichnaiaRassylkaKey: Guid
	SUvedomleniemOPoluchenii: Boolean
	TekstPisma: String
	UchetnaiaZapisDliaOtpravki: String
	ElektronnyiAdresOtvetov: String
	ElektronnyiAdresOtvetovVstroennyiPochtovyiKlientKey: Guid
	Vlozheniia: [DocumentRassylkaAnketVlozheniiaRowType!]
	Poluchateli: [DocumentRassylkaAnketPoluchateliRowType!]
	UchetnaiaZapisDliaOtpravkiType: String
	ElektronnyiAdresOtvetovType: String
}
input DocumentRassylkaAnketVlozheniiaInput {
	Key: Guid!
	LineNumber: Int64!
	VlozhenieBase64Data: Binary
	ImiaFaila: String
	VlozhenieType: String
	Vlozhenie: Stream
}
type DocumentRassylkaAnketVlozheniia {
	Key: Guid!
	LineNumber: Int64!
	VlozhenieBase64Data: Binary
	ImiaFaila: String
	VlozhenieType: String
	Vlozhenie: Stream
}
input DocumentRassylkaAnketPoluchateliInput {
	Key: Guid!
	LineNumber: Int64!
	Obieekt: String
	Poluchatel: String
	ObieektType: String
}
type DocumentRassylkaAnketPoluchateli {
	Key: Guid!
	LineNumber: Int64!
	Obieekt: String
	Poluchatel: String
	ObieektType: String
}
input CatalogVidyDeiatelnostiKontragentovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogVidyDeiatelnostiKontragentov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input CatalogTorgovoeOborudovanieInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Vid: String
	Model: String
	ObrabotkaObsluzhivaniiaKey: Guid
}
type CatalogTorgovoeOborudovanie {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Vid: String
	Model: String
	ObrabotkaObsluzhivaniiaKey: Guid
}
input CatalogSkhemyRealizatsiiInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	OrganizatsiiaKey: Guid
	OrganizatsiiaProdavetsKey: Guid
	EtapySkhemy: [CatalogSkhemyRealizatsiiEtapySkhemyRowTypeInput!]
}
type CatalogSkhemyRealizatsii {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	OrganizatsiiaKey: Guid
	OrganizatsiiaProdavetsKey: Guid
	EtapySkhemy: [CatalogSkhemyRealizatsiiEtapySkhemyRowType!]
}
input CatalogSkhemyRealizatsiiEtapySkhemyInput {
	Key: Guid!
	LineNumber: Int64!
	DogovorPokupkiKey: Guid
	DogovorProdazhiKey: Guid
	ZnachenieNatsenki: String
	KontragentPokupatelKey: Guid
	KontragentProdavetsKey: Guid
	OkrugliatVBolshuiuStoronu: Boolean
	OrganizatsiiaPokupatelKey: Guid
	OrganizatsiiaProdavetsKey: Guid
	PoriadokOkrugleniia: String
	SposobNatsenki: Int16
	TipNatsenki: Int16
	ZnachenieNatsenkiType: String
}
type CatalogSkhemyRealizatsiiEtapySkhemy {
	Key: Guid!
	LineNumber: Int64!
	DogovorPokupkiKey: Guid
	DogovorProdazhiKey: Guid
	ZnachenieNatsenki: String
	KontragentPokupatelKey: Guid
	KontragentProdavetsKey: Guid
	OkrugliatVBolshuiuStoronu: Boolean
	OrganizatsiiaPokupatelKey: Guid
	OrganizatsiiaProdavetsKey: Guid
	PoriadokOkrugleniia: String
	SposobNatsenki: Int16
	TipNatsenki: Int16
	ZnachenieNatsenkiType: String
}
input CatalogPodkliuchaemoeOborudovanieInput {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	VersiiaKomponenty: String
	IdentifikatorUstroistva: String
	UstroistvoOtkliucheno: Boolean
	UstroistvoIspolzuetsia: Boolean
	ObrabotchikDraivera: String
	ParametryBase64Data: Binary
	RabocheeMestoKey: Guid
	TipOborudovaniia: String
	TrebuetsiaPereustanovka: Boolean
	ParametryType: String
	Parametry: Stream
}
type CatalogPodkliuchaemoeOborudovanie {
	Key: Guid!
	DataVersion: String
	Description: String
	DeletionMark: Boolean
	VersiiaKomponenty: String
	IdentifikatorUstroistva: String
	UstroistvoOtkliucheno: Boolean
	UstroistvoIspolzuetsia: Boolean
	ObrabotchikDraivera: String
	ParametryBase64Data: Binary
	RabocheeMestoKey: Guid
	TipOborudovaniia: String
	TrebuetsiaPereustanovka: Boolean
	ParametryType: String
	Parametry: Stream
}
input DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	XKlassKon: Double
	XKlassNach: Double
	YKlassKon: Double
	YKlassNach: Double
	ZKlassKon: Double
	ZKlassNach: Double
	DataOkonchaniia: DateTime
	KolichestvoPeriodovAnaliza: Int16
	Comment: String
	Periodichnost: String
	RazovyiPokupatelKon: Int64
	RazovyiPokupatelNach: Int64
	TablitsaRaspredeleniiaKontragentov: [DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovRowTypeInput!]
}
type DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	XKlassKon: Double
	XKlassNach: Double
	YKlassKon: Double
	YKlassNach: Double
	ZKlassKon: Double
	ZKlassNach: Double
	DataOkonchaniia: DateTime
	KolichestvoPeriodovAnaliza: Int16
	Comment: String
	Periodichnost: String
	RazovyiPokupatelKon: Int64
	RazovyiPokupatelNach: Int64
	TablitsaRaspredeleniiaKontragentov: [DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovRowType!]
}
input DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovInput {
	Key: Guid!
	LineNumber: Int64!
	XYZKlassifikatsiia: String
	XYZKlassifikatsiiaStaraia: String
	ZnachenieParametra: Double
	IndeksSortirovki: Int16
	KolichestvoDokumentov: Int64
	KontragentKey: Guid
	KoeffitsientVariatsii: Double
	MenedzherKontragentaKey: Guid
	StadiiaVzaimootnoshenii: String
	StadiiaVzaimootnosheniiStaraia: String
}
type DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov {
	Key: Guid!
	LineNumber: Int64!
	XYZKlassifikatsiia: String
	XYZKlassifikatsiiaStaraia: String
	ZnachenieParametra: Double
	IndeksSortirovki: Int16
	KolichestvoDokumentov: Int64
	KontragentKey: Guid
	KoeffitsientVariatsii: Double
	MenedzherKontragentaKey: Guid
	StadiiaVzaimootnoshenii: String
	StadiiaVzaimootnosheniiStaraia: String
}
input CatalogGabarityInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
type CatalogGabarity {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
}
input DocumentZakazKlientaInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtoRezervirovanie: Boolean
	AdresDostavki: String
	ValiutaDokumentaKey: Guid
	Weight: Double
	VremiaNapominaniia: DateTime
	DataOplaty: DateTime
	DataOtgruzki: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontaktnoeLitsoKey: Guid
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NapomnitOSobytii: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	Soglasovano: Boolean
	StrukturnaiaEdinitsa: String
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	SkhemaRealizatsiiKey: Guid
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	NomerInternetDokumenta: String
	DataInternetDokumenta: DateTime
	StatusInternetDokumenta: String
	UnikalnyiNomerInternetDokumenta: String
	InternetZakaz: Boolean
	AdresElektronnoiPochty: String
	UsloviiaOplatyKey: Guid
	BDataDokumenta: DateTime
	BIdentifikator: String
	BNomerVersii: String
	Goods: [DocumentZakazKlientaTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	StrukturnaiaEdinitsaType: String
}
type DocumentZakazKlienta {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	AvtoRezervirovanie: Boolean
	AdresDostavki: String
	ValiutaDokumentaKey: Guid
	Weight: Double
	VremiaNapominaniia: DateTime
	DataOplaty: DateTime
	DataOtgruzki: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontaktnoeLitsoKey: Guid
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NapomnitOSobytii: Boolean
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	DepartmentKey: Guid
	Soglasovano: Boolean
	StrukturnaiaEdinitsa: String
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	SkhemaRealizatsiiKey: Guid
	TipDokumenta: String
	TipTsenKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	NomerInternetDokumenta: String
	DataInternetDokumenta: DateTime
	StatusInternetDokumenta: String
	UnikalnyiNomerInternetDokumenta: String
	InternetZakaz: Boolean
	AdresElektronnoiPochty: String
	UsloviiaOplatyKey: Guid
	BDataDokumenta: DateTime
	BIdentifikator: String
	BNomerVersii: String
	Goods: [DocumentZakazKlientaTovaryRowType!]
	DokumentOsnovanieType: String
	StrukturnaiaEdinitsaType: String
}
input DocumentZakazKlientaTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	Razmestit: Int64
	Rezervirovat: Int64
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	Cost: Double
	InstanceKey: Guid
	SumAutoDiscount: Double
	SumManualDiscount: Double
	KharakteristikaNomenklaturyKey: Guid
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
type DocumentZakazKlientaTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	ZnachenieUsloviiaAvtomaticheskoiSkidki: String
	Quantity: Int64
	ItemKey: Guid
	PercentAutoDiscount: Double
	PercentManualDiscount: Double
	SizeKey: Guid
	Razmestit: Int64
	Rezervirovat: Int64
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
	UslovieAvtomaticheskoiSkidki: String
	Cost: Double
	InstanceKey: Guid
	SumAutoDiscount: Double
	SumManualDiscount: Double
	KharakteristikaNomenklaturyKey: Guid
	ZnachenieUsloviiaAvtomaticheskoiSkidkiType: String
}
input ArriveFromManufacturingInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	IspolzovatSebestoimostSNDS: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	MaterialVProizvodstve: String
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Poteri: Double
	PoteriVProtsentakh: Boolean
	PoteriPostrochno: Boolean
	ProektKey: Guid
	RegistrirovatTseny: Boolean
	RegistrirovatTsenyPostavshchika: Boolean
	Sdelka: String
	DepartmentKey: Guid
	SobstvennoeProizvodstvo: Boolean
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UsloviiaOplatyKey: Guid
	UchityvatVesVstavok: Boolean
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	RuchnoeZapolnenieIzraskhodovannykhMaterialov: Boolean
	ProizvodstvennyiUchastokKey: Guid
	UslugaDliaVyvodaVPechatnyeFormyKey: Guid
	RuchnoiUchetVesaVstavok: Boolean
	Goods: [DocumentPostuplenieProduktsiiIzProizvodstvaTovaryRowTypeInput!]
	Materialy: [DocumentPostuplenieProduktsiiIzProizvodstvaMaterialyRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type ArriveFromManufacturing {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	IspolzovatSebestoimostSNDS: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	MaterialVProizvodstve: String
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Poteri: Double
	PoteriVProtsentakh: Boolean
	PoteriPostrochno: Boolean
	ProektKey: Guid
	RegistrirovatTseny: Boolean
	RegistrirovatTsenyPostavshchika: Boolean
	Sdelka: String
	DepartmentKey: Guid
	SobstvennoeProizvodstvo: Boolean
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UsloviiaOplatyKey: Guid
	UchityvatVesVstavok: Boolean
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	RuchnoeZapolnenieIzraskhodovannykhMaterialov: Boolean
	ProizvodstvennyiUchastokKey: Guid
	UslugaDliaVyvodaVPechatnyeFormyKey: Guid
	RuchnoiUchetVesaVstavok: Boolean
	Goods: [DocumentPostuplenieProduktsiiIzProizvodstvaTovaryRowType!]
	Materialy: [DocumentPostuplenieProduktsiiIzProizvodstvaMaterialyRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input ArriveFromManufacturingInstanceInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesPoter: Double
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	ProtsentPoter: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	StoimostVstavok: Double
	StoimostMetalla: Double
	StoimostRabot: Double
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	CostOfWork: Double
	SummaNDSVstavok: Double
	SummaNDSMetalla: Double
	SummaNDSRabot: Double
	VesVstavok: Double
}
type ArriveFromManufacturingInstance {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	VesPoter: Double
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	Pasport: String
	ProtsentPoter: Double
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	StoimostVstavok: Double
	StoimostMetalla: Double
	StoimostRabot: Double
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	CostOfWork: Double
	SummaNDSVstavok: Double
	SummaNDSMetalla: Double
	SummaNDSRabot: Double
	VesVstavok: Double
}
input DocumentPostuplenieProduktsiiIzProizvodstvaMaterialyInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	Nomenklatura: String
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	ItemType: String
}
type DocumentPostuplenieProduktsiiIzProizvodstvaMaterialy {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	Quantity: Int64
	Nomenklatura: String
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	ItemType: String
}
input DocumentJournalZakazyPostavshchikamInput {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	DataOplaty: DateTime
	DataPostupleniia: DateTime
	InformatsiiaKey: Guid
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sum: Double
	RefType: String!
}
type DocumentJournalZakazyPostavshchikam {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	DataOplaty: DateTime
	DataPostupleniia: DateTime
	InformatsiiaKey: Guid
	Comment: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	Sum: Double
	RefType: String!
}
input DocumentJournalSkladskieDokumentyInput {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	Comment: String
	Kontragent: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	Podrazdelenie: String
	Sklad: String
	Sum: Double
	RefType: String!
	KontragentType: String
	PodrazdelenieType: String
	SkladType: String
}
type DocumentJournalSkladskieDokumenty {
	Ref: String!
	Type: String
	Date: DateTime
	DeletionMark: Boolean
	Number: String
	Posted: Boolean
	ValiutaKey: Guid
	Comment: String
	Kontragent: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	Podrazdelenie: String
	Sklad: String
	Sum: Double
	RefType: String!
	KontragentType: String
	PodrazdelenieType: String
	SkladType: String
}
input CatalogsmsUsloviiaOtboraDiskontnykhKartInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: Int64
	DeletionMark: Boolean
	Deistvie: String
	Comment: String
	Predstavlenie: String
}
type CatalogsmsUsloviiaOtboraDiskontnykhKart {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: Int64
	DeletionMark: Boolean
	Deistvie: String
	Comment: String
	Predstavlenie: String
}
input ArriveInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetKontragentaKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	Koef: Double
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	RegistrirovatTseny: Boolean
	RegistrirovatTsenyPostavshchika: Boolean
	Sdelka: String
	SkladOrderKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UsloviiaOplatyKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	ZagruzhenIzUIuP: Boolean
	VystavkaOstatki: Boolean
	Goods: [DocumentPostuplenieTovarovUslugTovaryRowTypeInput!]
	Uslugi: [DocumentPostuplenieTovarovUslugUslugiRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type Arrive {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	BankovskiiSchetKontragentaKey: Guid
	ValiutaDokumentaKey: Guid
	Weight: Double
	GruzootpravitelKey: Guid
	GruzopoluchatelKey: Guid
	DataVkhodiashchegoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	Koef: Double
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	NDSVkliuchenVStoimost: Boolean
	NomerVkhodiashchegoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	RegistrirovatTseny: Boolean
	RegistrirovatTsenyPostavshchika: Boolean
	Sdelka: String
	SkladOrderKey: Guid
	SummaVkliuchaetNDS: Boolean
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UsloviiaOplatyKey: Guid
	UchityvatNDS: Boolean
	KhoziaistvennaiaOperatsiiaKey: Guid
	ZagruzhenIzUIuP: Boolean
	VystavkaOstatki: Boolean
	Goods: [DocumentPostuplenieTovarovUslugTovaryRowType!]
	Uslugi: [DocumentPostuplenieTovarovUslugUslugiRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentPostuplenieTovarovUslugTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaIzmereniiaKey: Guid
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	Koef: Double
	Koeffitsient: Double
	NaborKey: Guid
	ItemKey: Guid
	NomerGTDKey: Guid
	NomerNabora: Int64
	Pasport: String
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	StranaProiskhozhdeniiaKey: Guid
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
type DocumentPostuplenieTovarovUslugTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	EdinitsaIzmereniiaKey: Guid
	ZakazKlientaKey: Guid
	KachestvoKey: Guid
	Quantity: Int64
	Koef: Double
	Koeffitsient: Double
	NaborKey: Guid
	ItemKey: Guid
	NomerGTDKey: Guid
	NomerNabora: Int64
	Pasport: String
	ProektKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	DepartmentKey: Guid
	StavkaNDS: String
	StranaProiskhozhdeniiaKey: Guid
	Sum: Double
	SummaNDS: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
}
input DocumentPostuplenieTovarovUslugUslugiInput {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
type DocumentPostuplenieTovarovUslugUslugi {
	Key: Guid!
	LineNumber: Int64!
	Quantity: Int64
	ItemKey: Guid
	NomenklaturnaiaGruppaKey: Guid
	PodrazdelenieKey: Guid
	ProektKey: Guid
	Soderzhanie: String
	StavkaNDS: String
	StatiaZatratKey: Guid
	Sum: Double
	SummaNDS: Double
	Cost: Double
}
input DocumentSchetFakturaVydannyiInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	ValiutnaiaSumma: Double
	DataPlatezhnoRaschetnogoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	Comment: String
	KontragentKey: Guid
	NaAvans: Boolean
	NomerPlatezhnoRaschetnogoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	Pod0: Boolean
	StavkaNDS: String
	Sum: Double
	SumOfDocument: Double
	SummaNDS: Double
	SformirovanPriVvodeNachalnykhOstatkovNDS: Boolean
	TipDokumenta: String
	SchetFakturaOsnovanieKey: Guid
	Ispravlenie: Boolean
	NomerIspravleniia: String
	Korrektirovochnyi: Boolean
	IskhodnyiDokumentKey: Guid
	NomerIskhodnogoDokumenta: String
	DataIskhodnogoDokumenta: DateTime
	NomerIspravleniiaIskhodnogoDokumenta: String
	DataIspravleniiaIskhodnogoDokumenta: DateTime
	VystavlenVElektronnomVide: Boolean
	DataVystavleniia: DateTime
	SummaDokumentaKomissiia: Double
	SummaNDSDokumentaKomissiia: Double
	NaimenovanieTovaraUslugiDliaPechatiKorrektirovkiPoNDS: String
	DokumentOsnovanieType: String
}
type DocumentSchetFakturaVydannyi {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	ValiutnaiaSumma: Double
	DataPlatezhnoRaschetnogoDokumenta: DateTime
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	DokumentSozdanVIuTD: Boolean
	Comment: String
	KontragentKey: Guid
	NaAvans: Boolean
	NomerPlatezhnoRaschetnogoDokumenta: String
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	Pod0: Boolean
	StavkaNDS: String
	Sum: Double
	SumOfDocument: Double
	SummaNDS: Double
	SformirovanPriVvodeNachalnykhOstatkovNDS: Boolean
	TipDokumenta: String
	SchetFakturaOsnovanieKey: Guid
	Ispravlenie: Boolean
	NomerIspravleniia: String
	Korrektirovochnyi: Boolean
	IskhodnyiDokumentKey: Guid
	NomerIskhodnogoDokumenta: String
	DataIskhodnogoDokumenta: DateTime
	NomerIspravleniiaIskhodnogoDokumenta: String
	DataIspravleniiaIskhodnogoDokumenta: DateTime
	VystavlenVElektronnomVide: Boolean
	DataVystavleniia: DateTime
	SummaDokumentaKomissiia: Double
	SummaNDSDokumentaKomissiia: Double
	NaimenovanieTovaraUslugiDliaPechatiKorrektirovkiPoNDS: String
	DokumentOsnovanieType: String
}
input DocumentPlanProdazhPoSalonamInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	PlanovyiPeriod: DateTime
	SummaPlana: Double
	Comment: String
	PoDniam: Boolean
	DetalizatsiiaPlana: Boolean
	KalendarPlanaKey: Guid
	Salony: [DocumentPlanProdazhPoSalonamSalonyRowTypeInput!]
	DniPoGrafiku: [DocumentPlanProdazhPoSalonamDniPoGrafikuRowTypeInput!]
}
type DocumentPlanProdazhPoSalonam {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	PlanovyiPeriod: DateTime
	SummaPlana: Double
	Comment: String
	PoDniam: Boolean
	DetalizatsiiaPlana: Boolean
	KalendarPlanaKey: Guid
	Salony: [DocumentPlanProdazhPoSalonamSalonyRowType!]
	DniPoGrafiku: [DocumentPlanProdazhPoSalonamDniPoGrafikuRowType!]
}
input DocumentPlanProdazhPoSalonamSalonyInput {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
	SummaPlana: Double
	Primechanie: String
	IndeksStrokiIzTablitsyDnei: Int64
	SummaPlanaDnevnaia: Double
	DenPoGrafiku: DateTime
	SummaPlanaFakt: Double
	PlanEst: Boolean
	KU: Double
	FormulaDliaRasschetaKey: Guid
	RasshifrovkaFormuly: String
}
type DocumentPlanProdazhPoSalonamSalony {
	Key: Guid!
	LineNumber: Int64!
	SalonKey: Guid
	SummaPlana: Double
	Primechanie: String
	IndeksStrokiIzTablitsyDnei: Int64
	SummaPlanaDnevnaia: Double
	DenPoGrafiku: DateTime
	SummaPlanaFakt: Double
	PlanEst: Boolean
	KU: Double
	FormulaDliaRasschetaKey: Guid
	RasshifrovkaFormuly: String
}
input DocumentPlanProdazhPoSalonamDniPoGrafikuInput {
	Key: Guid!
	LineNumber: Int64!
	DenPoGrafiku: Int16
}
type DocumentPlanProdazhPoSalonamDniPoGrafiku {
	Key: Guid!
	LineNumber: Int64!
	DenPoGrafiku: Int16
}
input CatalogBankovskieSchetaInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	Owner: String
	DeletionMark: Boolean
	BankKey: Guid
	BankDliaRaschetovKey: Guid
	ValiutaDenezhnykhSredstvKey: Guid
	VidScheta: String
	DataZakrytiia: DateTime
	DataOtkrytiia: DateTime
	KontragentDliaOgranicheniiaPravDostupaKey: Guid
	MesiatsPropisiu: Boolean
	NomerIDataRazresheniia: String
	NomerScheta: String
	OrganizatsiiaDliaOgranicheniiaPravDostupaKey: Guid
	SummaBezKopeek: Boolean
	TekstKorrespondenta: String
	TekstNaznacheniia: String
	OwnerType: String
}
type CatalogBankovskieScheta {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	Owner: String
	DeletionMark: Boolean
	BankKey: Guid
	BankDliaRaschetovKey: Guid
	ValiutaDenezhnykhSredstvKey: Guid
	VidScheta: String
	DataZakrytiia: DateTime
	DataOtkrytiia: DateTime
	KontragentDliaOgranicheniiaPravDostupaKey: Guid
	MesiatsPropisiu: Boolean
	NomerIDataRazresheniia: String
	NomerScheta: String
	OrganizatsiiaDliaOgranicheniiaPravDostupaKey: Guid
	SummaBezKopeek: Boolean
	TekstKorrespondenta: String
	TekstNaznacheniia: String
	OwnerType: String
}
input DocumentStornirovanieOtchetaKomitentuOProdazhakhInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProtsentKomissionnogoVoznagrazhdeniia: Double
	Sdelka: String
	SposobRaschetaKomissionnogoVoznagrazhdeniia: String
	StavkaNDSVoznagrazhdeniia: String
	SummaVoznagrazhdeniia: Double
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UderzhatKomissionnoeVoznagrazhdenie: Boolean
	UsloviiaOplatyKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	DenezhnyeSredstva: [DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvaRowTypeInput!]
	Goods: [DocumentStornirovanieOtchetaKomitentuOProdazhakhTovaryRowTypeInput!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
type DocumentStornirovanieOtchetaKomitentuOProdazhakh {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	ValiutaDokumentaKey: Guid
	Weight: Double
	DogovorKontragentaKey: Guid
	DokumentOsnovanie: String
	KolichestvoDokumenta: Int64
	Comment: String
	KontragentKey: Guid
	KratnostVzaimoraschetov: Int64
	KursVzaimoraschetov: Double
	OrganizatsiiaKey: Guid
	OtvetstvennyiKey: Guid
	PodrazdelenieKey: Guid
	ProtsentKomissionnogoVoznagrazhdeniia: Double
	Sdelka: String
	SposobRaschetaKomissionnogoVoznagrazhdeniia: String
	StavkaNDSVoznagrazhdeniia: String
	SummaVoznagrazhdeniia: Double
	SumOfDocument: Double
	TipDokumenta: String
	TipTsenKey: Guid
	UderzhatKomissionnoeVoznagrazhdenie: Boolean
	UsloviiaOplatyKey: Guid
	KhoziaistvennaiaOperatsiiaKey: Guid
	DenezhnyeSredstva: [DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvaRowType!]
	Goods: [DocumentStornirovanieOtchetaKomitentuOProdazhakhTovaryRowType!]
	DokumentOsnovanieType: String
	SdelkaType: String
}
input DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvaInput {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
type DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva {
	Key: Guid!
	LineNumber: Int64!
	VidOtchetaPoPlatezham: String
	StavkaNDS: String
	Sum: Double
	SummaNDS: Double
}
input DocumentStornirovanieOtchetaKomitentuOProdazhakhTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	OtchetKomitentuKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPostupleniia: Double
	DokumentPostupleniiaType: String
}
type DocumentStornirovanieOtchetaKomitentuOProdazhakhTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	DokumentPostupleniia: String
	Quantity: Int64
	ItemKey: Guid
	OtchetKomitentuKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	Sum: Double
	SummaVoznagrazhdeniia: Double
	SummaNDSVoznagrazhdeniia: Double
	SummaPostupleniia: Double
	KharakteristikaNomenklaturyKey: Guid
	Cost: Double
	TsenaPostupleniia: Double
	DokumentPostupleniiaType: String
}
input DocumentPeredachaVRemontInput {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	SobstvennaiaMasterskaia: Boolean
	KolichestvoDokumenta: Int64
	Weight: Double
	IzdeliiaPriniatyeVRemont: [DocumentPeredachaVRemontIzdeliiaPriniatyeVRemontRowTypeInput!]
	Goods: [DocumentPeredachaVRemontTovaryRowTypeInput!]
}
type DocumentPeredachaVRemont {
	Key: Guid!
	DataVersion: String
	Number: String
	Date: DateTime
	DeletionMark: Boolean
	Posted: Boolean
	OrganizatsiiaKey: Guid
	DepartmentKey: Guid
	KontragentKey: Guid
	DogovorKontragentaKey: Guid
	SobstvennaiaMasterskaia: Boolean
	KolichestvoDokumenta: Int64
	Weight: Double
	IzdeliiaPriniatyeVRemont: [DocumentPeredachaVRemontIzdeliiaPriniatyeVRemontRowType!]
	Goods: [DocumentPeredachaVRemontTovaryRowType!]
}
input DocumentPeredachaVRemontIzdeliiaPriniatyeVRemontInput {
	Key: Guid!
	LineNumber: Int64!
	KliuchNomenklaturyKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	DokumentOprikhodovaniiaKey: Guid
	Weight: Double
	Quantity: Int64
}
type DocumentPeredachaVRemontIzdeliiaPriniatyeVRemont {
	Key: Guid!
	LineNumber: Int64!
	KliuchNomenklaturyKey: Guid
	SizeKey: Guid
	KharakteristikaNomenklaturyKey: Guid
	InstanceKey: Guid
	DokumentOprikhodovaniiaKey: Guid
	Weight: Double
	Quantity: Int64
}
input DocumentPeredachaVRemontTovaryInput {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
type DocumentPeredachaVRemontTovary {
	Key: Guid!
	LineNumber: Int64!
	Weight: Double
	KachestvoKey: Guid
	Quantity: Int64
	ItemKey: Guid
	SizeKey: Guid
	InstanceKey: Guid
	KharakteristikaNomenklaturyKey: Guid
}
input CatalogPolzovateliInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Number: String
	StrategiiaDostupaKZapisiam: String
	FizLitsoKey: Guid
	ProfilKey: Guid
}
type CatalogPolzovateli {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	Number: String
	StrategiiaDostupaKZapisiam: String
	FizLitsoKey: Guid
	ProfilKey: Guid
}
input CatalogTsenovyeKoridoryInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	TsenaOt: Int64
	TsenaDo: Int64
}
type CatalogTsenovyeKoridory {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	ParentKey: Guid
	IsFolder: Boolean
	DeletionMark: Boolean
	TsenaOt: Int64
	TsenaDo: Int64
}
input CatalogGruppySkladovInput {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Sklady: [CatalogGruppySkladovSkladyRowTypeInput!]
}
type CatalogGruppySkladov {
	Key: Guid!
	DataVersion: String
	Description: String
	Code: String
	DeletionMark: Boolean
	Sklady: [CatalogGruppySkladovSkladyRowType!]
}
input CatalogGruppySkladovSkladyInput {
	Key: Guid!
	LineNumber: Int64!
	DepartmentKey: Guid
}
type CatalogGruppySkladovSklady {
	Key: Guid!
	LineNumber: Int64!
	DepartmentKey: Guid
}
input PrimaryAccumulationRegisterPartiiTovarovVProizvodstve {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterPartiiTovarovVProizvodstveRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterVzaimoraschetySPodotchetnymiLitsami {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterVnutrennieZakazy {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterVnutrennieZakazyRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstvaKomitenta {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstvaKomitentaRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterZakazyKlientov {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterZakazyKlientovRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSummyPoFinmonitoringuRoznitsa {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstvaKPolucheniiu {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterProdazhiPoDiskontnymKartam {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterProdazhiPoDiskontnymKartamRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterTovaryPoluchennye {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterTovaryPoluchennyeRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSvobodnyeOstatki {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSvobodnyeOstatkiRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSummyVRassrochku {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSummyVRassrochkuRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterGrafikPlatezhei {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterGrafikPlatezheiRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterRoznichnaiaVyruchka {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterRoznichnaiaVyruchkaRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterTovaryVPuti {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterTovaryVPutiRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterPoteriMetallaVProizvodstve {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterPoteriMetallaVProizvodstveRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterPartiiTovarovNaSkladakh {
	Recorder: String!
	RecorderType: String!
}
input PrimaryProductActionDocument {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSummyDokumentovDliaObmena {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSummyDokumentovDliaObmenaRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDvizheniiaDenezhnykhSredstv {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterProdazhiPoStatiam {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterProdazhiPoStatiamRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryInformationRegisterTsenyNomenklatury {
	Recorder: String!
	RecorderType: String!
}
input PrimaryInformationRegisterTsenyNomenklaturyRecordType {
	Period: DateTime!
	TipTsenKey: Guid!
	SegmentNomenklaturyKey: Guid!
	ItemKey: Guid!
	InstanceKey: Guid!
	KharakteristikaNomenklaturyKey: Guid!
	SizeKey: Guid!
}
input PrimaryAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstvaVRezerve {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstvaVRezerveRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDavalcheskiiMetallPoteri {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDavalcheskiiMetallPoteriRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryInformationRegisterTsenyPoPreiskurantu {
	Recorder: String!
	RecorderType: String!
}
input PrimaryInformationRegisterTsenyPoPreiskurantuRecordType {
	Period: DateTime!
	KamenKey: Guid!
	FormaOgrankiKey: Guid!
	TsvetKamniaKey: Guid!
	GruppaTsvetaKey: Guid!
	GruppaDefektaKey: Guid!
	RassevKey: Guid!
	Razmer1Ot: Double!
	Razmer1Do: Double!
}
input PrimaryAccumulationRegisterTovaryVOtbore {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterTovaryVOtboreRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterRealizovannyeTovary {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterRealizovannyeTovaryRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstvaKomissionera {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterProdazhi1 {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterProdazhi1RecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterTovaryNaSkladakhAM {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterTovaryNaSkladakhAMRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSummyPoFinmonitoringu {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSummyPoFinmonitoringuRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryInformationRegisterTsenyNomenklaturyKontragentov {
	Recorder: String!
	RecorderType: String!
}
input PrimaryInformationRegisterTsenyNomenklaturyKontragentovRecordType {
	Period: DateTime!
	TipTsenKey: Guid!
	ItemKey: Guid!
	InstanceKey: Guid!
	KharakteristikaNomenklaturyKey: Guid!
	SizeKey: Guid!
}
input PrimaryAccumulationRegisterVzaimoraschetySKontragentami {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterVzaimoraschetySKontragentamiRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSummyPokupokPoDiskontnymKartam {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterVypolneniePlanaProdazh {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterVypolneniePlanaProdazhRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDavalcheskiiMetall {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDavalcheskiiMetallRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstva {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstvaRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterTovaryPeredannye {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterTovaryPeredannyeRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstvaKSpisaniiu {
	Recorder: String!
	RecorderType: String!
}
input PrimaryAccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType {
	Recorder: String!
	LineNumber: Int64!
	RecorderType: String!
}
input PrimaryCatalogDogovoryKontragentov {
	Key: Guid!
}
input PrimaryOrder {
	Key: Guid!
}
input PrimaryDocumentChekKKMOplata {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentChekKKMOplataSertifikatami {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentChekKKMProdazhaSertifikatov {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentChekKKMTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentChekKKMDokumentyObmena {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentChekKKMDogovoraRassrochkiProdazha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentChekKKMDogovoraRassrochkiOplata {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentChekKKMOplataBallami {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentChekKKMSkidkiNatsenki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentChekKKMUpravliaemyeSkidki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentChekKKMPodarki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentChekKKMKupony {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPereotsenkaValiutnykhSredstv {
	Key: Guid!
}
input PrimaryCatalogTipySkidokNatsenok {
	Key: Guid!
}
input PrimaryCatalogTipySkidokNatsenokVremiaPoDniamNedeli {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogVidyKodirovokiTsepei {
	Key: Guid!
}
input PrimaryCatalogVidyKodirovokiTsepeiElementyKodirovki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetKomitentuOProdazhakh {
	Key: Guid!
}
input PrimaryDocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetKomitentuOProdazhakhTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKassy {
	Key: Guid!
}
input PrimaryCatalogKassiry {
	Key: Guid!
}
input PrimaryDocumentZaiavkaNaPereotsenkuTovarov {
	Key: Guid!
}
input PrimaryDocumentZaiavkaNaPereotsenkuTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogProizvodstvennyeUchastki {
	Key: Guid!
}
input PrimaryDocumentZakrytieZakazovKlientov {
	Key: Guid!
}
input PrimaryDocumentZakrytieZakazovKlientovZakazy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogProekty {
	Key: Guid!
}
input PrimaryDocumentPlatezhnoePoruchenieVkhodiashchee {
	Key: Guid!
}
input PrimaryDocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentVydachaZakaza {
	Key: Guid!
}
input PrimaryCatalogFormyOgranki {
	Key: Guid!
}
input PrimaryCatalogFormatyMagazinov {
	Key: Guid!
}
input PrimaryCatalogRabochieMesta {
	Key: Guid!
}
input PrimaryCatalogNastroikiVypolneniiaObmena {
	Key: Guid!
}
input PrimaryCatalogNastroikiVypolneniiaObmenaNastroikiObmena {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogZnacheniiaSvoistvObieektov {
	Key: Guid!
}
input PrimaryDocumentRealizatsiiaTovarovUslug {
	Key: Guid!
}
input PrimaryDocumentRealizatsiiaTovarovUslugTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentRealizatsiiaTovarovUslugUslugi {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentSobytie {
	Key: Guid!
}
input PrimaryDocumentSobytieStoronnieLitsa {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogVariantyOtvetovOprosov {
	Key: Guid!
}
input PrimaryCatalogGruppyPisemElektronnoiPochty {
	Key: Guid!
}
input PrimaryCatalogGruppyPochtovoiRassylki {
	Key: Guid!
}
input PrimaryCatalogNastroikiOtchetov {
	Key: Guid!
}
input PrimaryCatalogNastroikiOtchetovGruppyNastroekIPolzovateli {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam {
	Key: Guid!
}
input PrimaryCatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDepartment {
	Key: Guid!
}
input PrimaryCatalogKodyVidovTovarov {
	Key: Guid!
}
input PrimaryCatalogRassevy {
	Key: Guid!
}
input PrimaryCatalogPrichinyZakrytiiaZakazov {
	Key: Guid!
}
input PrimaryCatalogSegmentyNomenklatury {
	Key: Guid!
}
input PrimaryCatalogSostavStrokiCheka {
	Key: Guid!
}
input PrimaryCatalogUsloviiaPriemaIzdeliiNaKomissiiu {
	Key: Guid!
}
input PrimaryCatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei {
	Key: Guid!
}
input PrimaryDocumentKorrektirovkaDolga {
	Key: Guid!
}
input PrimaryDocumentKorrektirovkaDolgaSummyDolga {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryPayType {
	Key: Guid!
}
input PrimaryCatalogKhranilishcheShablonov {
	Key: Guid!
}
input PrimaryDocumentZaiavkaNaRaskhodovanieSredstv {
	Key: Guid!
}
input PrimaryDocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentZakrytieZakazovPostavshchikam {
	Key: Guid!
}
input PrimaryDocumentZakrytieZakazovPostavshchikamZakazy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogVidyKamnei {
	Key: Guid!
}
input PrimaryDocumentAnketyKlientovDliaFinMonitoringa {
	Key: Guid!
}
input PrimaryDocumentAnketyKlientovDliaFinMonitoringaAnkety {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogDogovoryRassrochki {
	Key: Guid!
}
input PrimaryCatalogSertifikaty {
	Key: Guid!
}
input PrimaryDocumentPostuplenieDavalcheskogoMetalla {
	Key: Guid!
}
input PrimaryDocumentInkassovoePorucheniePeredannoe {
	Key: Guid!
}
input PrimaryDocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentInkassovoePorucheniePeredannoeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogFormulyDliaRascheta {
	Key: Guid!
}
input PrimaryCatalogKupony {
	Key: Guid!
}
input PrimaryCorrecting {
	Key: Guid!
}
input PrimaryDocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentInternetZakaz {
	Key: Guid!
}
input PrimaryDocumentInternetZakazTovaryInternetZakaza {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentInternetZakazTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogRegiony {
	Key: Guid!
}
input PrimarySaleJournal {
	Key: Guid!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhBonusy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhOplataSertifikatami {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhDokumentyObmena {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhOplataBallami {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetORoznichnykhProdazhakhKupony {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtmenaSkidokNomenklatury {
	Key: Guid!
}
input PrimaryDocumentOtmenaSkidokNomenklaturyDokumenty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogTovarnyeGruppy {
	Key: Guid!
}
input PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv {
	Key: Guid!
}
input PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogOrderKey {
	Key: Guid!
}
input PrimaryDocumentKassovyiChekKorrektsii {
	Key: Guid!
}
input PrimaryDocumentKassovyiChekKorrektsiiOplata {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentSchetNaOplatuPokupateliu {
	Key: Guid!
}
input PrimaryDocumentSchetNaOplatuPokupateliuTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentSchetNaOplatuPokupateliuUslugi {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogNastroikiObmenaDannymi {
	Key: Guid!
}
input PrimaryCatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentJournalBankovskieRaschetnyeDokumenty {
	Ref: String!
	RefType: String!
}
input PrimaryDocumentZamenaDiskontnoiKarty {
	Key: Guid!
}
input PrimaryReturnToSupplier {
	Key: Guid!
}
input PrimaryDocumentVozvratTovarovPostavshchikuTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentInventarizatsiiaTovarovNaSklade {
	Key: Guid!
}
input PrimaryDocumentInventarizatsiiaTovarovNaSkladeTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentInventarizatsiiaTovarovNaSkladeSertifikaty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPrikhodnyiKassovyiOrder {
	Key: Guid!
}
input PrimaryDocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPrikhodnyiKassovyiOrderOplata {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPrikhodnyiKassovyiOrderTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogPrichinyVozvrata {
	Key: Guid!
}
input PrimaryDocumentDenezhnyiChek {
	Key: Guid!
}
input PrimaryDocumentVozvratMaterialovIzProizvodstva {
	Key: Guid!
}
input PrimaryDocumentVozvratMaterialovIzProizvodstvaTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPereotsenkaTovarovOtdannykhNaKomissiiu {
	Key: Guid!
}
input PrimaryDocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSN {
	Key: Guid!
}
input PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentGTDImport {
	Key: Guid!
}
input PrimaryDocumentGTDImportRazdely {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentGTDImportTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentAktSverki {
	Key: Guid!
}
input PrimaryDocumentAktSverkiTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogFaily {
	Key: Guid!
}
input PrimaryCatalogFailyDopolnitelnyeRekvizity {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogFailySertifikatyShifrovaniia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogUchetnyeZapisiElektronnoiPochty {
	Key: Guid!
}
input PrimaryCatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPlaniruemoePostuplenieDenezhnykhSredstv {
	Key: Guid!
}
input PrimaryDocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPreiskurantTsenNaKamni {
	Key: Guid!
}
input PrimaryPurchase {
	Key: Guid!
}
input PrimaryDocumentSkupkaTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentSchetFakturaPoluchennyi {
	Key: Guid!
}
input PrimaryDocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentAktKhimicheskogoAnalizaMetalla {
	Key: Guid!
}
input PrimaryCatalogfmKartochkaKontragenta {
	Key: Guid!
}
input PrimaryCatalogfmKartochkaKontragentaDannyeKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentSpisanieProsrochennykhSertifikatov {
	Key: Guid!
}
input PrimaryDocumentSpisanieProsrochennykhSertifikatovSertifikaty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentZakrytieAvansovPoGrafikuPlatezhei {
	Key: Guid!
}
input PrimaryDocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogTipySistemNalogooblozheniiaKKT {
	Key: Guid!
}
input PrimaryDocumentAkkreditivPeredannyi {
	Key: Guid!
}
input PrimaryDocumentAkkreditivPeredannyiRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentAkkreditivPeredannyiRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimarySupplier {
	Key: Guid!
}
input PrimaryCatalogKontragentyVidyDeiatelnosti {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentInformatsionnoeSoobshchenie {
	Key: Guid!
}
input PrimaryCatalogVlozheniiaElektronnykhPisem {
	Key: Guid!
}
input PrimaryDocumentPlatezhnoeTrebovanieVystavlennoe {
	Key: Guid!
}
input PrimaryDocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentMarketingovaiaAktsiia {
	Key: Guid!
}
input PrimaryDocumentMarketingovaiaAktsiiaMagaziny {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentMarketingovaiaAktsiiaSkidkiNatsenki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogStsenariiObmenovDannymi {
	Key: Guid!
}
input PrimaryCatalogStsenariiObmenovDannymiNastroikiObmena {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryItem {
	Key: Guid!
}
input PrimaryCatalogNomenklaturaSostavLigatury {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOpros {
	Key: Guid!
}
input PrimaryDocumentOprosVoprosy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOprosSostavnoiOtvet {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogGruppyPoluchateleiSkidki {
	Key: Guid!
}
input PrimaryReassessment {
	Key: Guid!
}
input PrimaryDocumentPereotsenkaTovarovVNTTTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogTomaKhraneniiaFailov {
	Key: Guid!
}
input PrimaryDocumentJournalProizvodstvennyeDokumenty {
	Ref: String!
	RefType: String!
}
input PrimaryDocumentIzmeneniePravDostupa {
	Key: Guid!
}
input PrimaryCatalogNastroikaAssortimentnoiMatritsy {
	Key: Guid!
}
input PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentJournalDokumentyKontragentov {
	Ref: String!
	RefType: String!
}
input PrimaryMoveInstance {
	Key: Guid!
}
input PrimaryDocumentPeremeshchenieTovarovSertifikaty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPeremeshchenieTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPeremeshchenieTovarovSpisokZaiavok {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentZakrytieZaiavokNaRaskhodovanieSredstv {
	Key: Guid!
}
input PrimaryDocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryMemberCard {
	Key: Guid!
}
input PrimaryDocumentABCKlassifikatsiiaPokupatelei {
	Key: Guid!
}
input PrimaryDocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogIdentifikatoryObieektovMetadannykh {
	Key: Guid!
}
input PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSklade {
	Key: Guid!
}
input PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentKorrektirovkaRealizatsii {
	Key: Guid!
}
input PrimaryDocumentKorrektirovkaRealizatsiiTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentKorrektirovkaRealizatsiiUslugi {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogVidyDefektov {
	Key: Guid!
}
input PrimaryDocumentDoverennost {
	Key: Guid!
}
input PrimaryDocumentDoverennostTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogShablonyZapolneniiaKU {
	Key: Guid!
}
input PrimaryCatalogShablonyZapolneniiaKUPrazdnichnyeDni {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogShablonyZapolneniiaKUKUNaNedeliu {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogShablonyZapolneniiaKUSalony {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPlanZapolneniiaVitrin {
	Key: Guid!
}
input PrimaryDocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryInstance {
	Key: Guid!
}
input PrimaryReturnToManufacturing {
	Key: Guid!
}
input PrimaryDocumentVozvratProduktsiiVProizvodstvoTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogNomeraGTD {
	Key: Guid!
}
input PrimaryCatalogNastroikiRabochegoMestaPolzovatelia {
	Key: Guid!
}
input PrimaryCatalogNastroikiRabochegoMestaPolzovateliaNastroiki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogsmsShablony {
	Key: Guid!
}
input PrimaryWriteOff {
	Key: Guid!
}
input PrimaryDocumentSpisanieTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentSpisanieTovarovSertifikaty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentsmsSoobshchenie {
	Key: Guid!
}
input PrimaryDocumentsmsSoobshcheniePoluchateli {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOplataOtPokupateliaPlatezhnoiKartoi {
	Key: Guid!
}
input PrimaryCatalogDragotsennyeKamni {
	Key: Guid!
}
input PrimaryCatalogKalendariPlanirovaniiaProdazh {
	Key: Guid!
}
input PrimaryCatalogKalendariPlanirovaniiaProdazhKUPoDniam {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKalendariPlanirovaniiaProdazhSalony {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKontaktnyeLitsa {
	Key: Guid!
}
input PrimaryCatalogFizicheskieLitsa {
	Key: Guid!
}
input PrimaryCatalogTipovyeAnkety {
	Key: Guid!
}
input PrimaryCatalogTipovyeAnketyVoprosyAnkety {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentNachislenieSpisanieBonusov {
	Key: Guid!
}
input PrimaryDocumentNachislenieSpisanieBonusovDiskontnyeKarty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryType {
	Key: Guid!
}
input PrimaryCatalogfmKodyVidovDokumentov {
	Key: Guid!
}
input PrimaryDocumentPlatezhnoeTrebovaniePoluchennoe {
	Key: Guid!
}
input PrimaryDocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv {
	Key: Guid!
}
input PrimaryDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogRazdelyAnkety {
	Key: Guid!
}
input PrimaryDocumentOtchetPoFinMonitoringu {
	Key: Guid!
}
input PrimaryDocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetPoFinMonitoringuDannyeDokumenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKliuchiAnalitikiUchetaNomenklatury {
	Key: Guid!
}
input PrimaryCatalogVersiiFailov {
	Key: Guid!
}
input PrimaryCatalogVersiiFailovElektronnyePodpisi {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentUstanovkaTsenNomenklatury {
	Key: Guid!
}
input PrimaryDocumentUstanovkaTsenNomenklaturyTipyTsen {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentUstanovkaTsenNomenklaturyTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv {
	Key: Guid!
}
input PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPreiskurantNaSkupku {
	Key: Guid!
}
input PrimaryDocumentPreiskurantNaSkupkuProby {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPeredachaMaterialovVProizvodstvo {
	Key: Guid!
}
input PrimaryDocumentPeredachaMaterialovVProizvodstvoTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentVnutrenniiZakaz {
	Key: Guid!
}
input PrimaryDocumentVnutrenniiZakazTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKhranilishcheDopolnitelnoiInformatsii {
	Key: Guid!
}
input PrimaryCatalogDopolnitelnyeVneshnieObrabotki {
	Key: Guid!
}
input PrimaryCatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogDopolnitelnyeVneshnieObrabotkiKomandy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogDopolnitelnyeVneshnieObrabotkiRazdely {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogDopolnitelnyeVneshnieObrabotkiNaznachenie {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogDopolnitelnyeVneshnieObrabotkiRazresheniia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogGruppyPolzovatelei {
	Key: Guid!
}
input PrimaryCatalogGruppyPolzovateleiPolzovateliGruppy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentJournalZakazyKlientov {
	Ref: String!
	RefType: String!
}
input PrimaryDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki {
	Key: Guid!
}
input PrimaryDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentZaiavkaNaPeremeshchenieTovarov {
	Key: Guid!
}
input PrimaryDocumentZaiavkaNaPeremeshchenieTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogUsloviiaProdazh {
	Key: Guid!
}
input PrimaryDocumentVvodNachalnykhOstatkovPoFinMonitoringu {
	Key: Guid!
}
input PrimaryDocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogOrganizatsii {
	Key: Guid!
}
input PrimaryCatalogUsloviiaOplaty {
	Key: Guid!
}
input PrimaryCatalogUsloviiaOplatyTablitsaVyplat {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKategoriiObieektov {
	Key: Guid!
}
input PrimaryCatalogfmZnacheniiaDliaZapolneniia {
	Key: Guid!
}
input PrimaryDocumentUdalitNariadZakaz {
	Key: Guid!
}
input PrimaryDocumentUdalitNariadZakazIzdeliia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentUdalitNariadZakazUslugi {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentUdalitNariadZakazSpetsifikatsiia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentUdalitNariadZakazMetally {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentUdalitNariadZakazVstavki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogBanki {
	Key: Guid!
}
input PrimaryCatalogRoliKontaktnykhLits {
	Key: Guid!
}
input PrimaryDocumentRestrukturizatsiiaDolga {
	Key: Guid!
}
input PrimaryDocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentAkkreditivPoluchennyi {
	Key: Guid!
}
input PrimaryDocumentAkkreditivPoluchennyiRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentAkkreditivPoluchennyiRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPriemIzRemonta {
	Key: Guid!
}
input PrimaryDocumentPriemIzRemontaIzdeliia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPriemIzRemontaMaterialy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogTsveta {
	Key: Guid!
}
input PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakh {
	Key: Guid!
}
input PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentJournalDavalcheskieDokumenty {
	Ref: String!
	RefType: String!
}
input PrimaryCatalogfmAnketaKlienta {
	Key: Guid!
}
input PrimaryCatalogfmAnketaKlientaDannyeKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogVidyVzaimoraschetov {
	Key: Guid!
}
input PrimaryDocumentUstanovkaZnacheniiTochkiZakaza {
	Key: Guid!
}
input PrimaryDocumentUstanovkaZnacheniiTochkiZakazaTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogZnacheniiaKodirovki {
	Key: Guid!
}
input PrimaryCatalogPravilaProdazh {
	Key: Guid!
}
input PrimaryCatalogPravilaProdazhTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPostuplenieDopRaskhodov {
	Key: Guid!
}
input PrimaryDocumentPostuplenieDopRaskhodovTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKhoziaistvennyeOperatsii {
	Key: Guid!
}
input PrimaryDocumentAvansovyiOtchet {
	Key: Guid!
}
input PrimaryDocumentAvansovyiOtchetVydannyeAvansy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentAvansovyiOtchetTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentAvansovyiOtchetOplataPostavshchikam {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentAvansovyiOtchetProchee {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogDolzhnostiOrganizatsii {
	Key: Guid!
}
input PrimaryCatalogAnalitikaTipaIzdeliia {
	Key: Guid!
}
input PrimaryCatalogDopolnitelnyePechatnyeFormy {
	Key: Guid!
}
input PrimaryCatalogDopolnitelnyePechatnyeFormyPrinadlezhnost {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryMemberCardsType {
	Key: Guid!
}
input PrimaryDocumentRegistratsiiaNaSaite {
	Key: Guid!
}
input PrimaryCatalogObrabotkiObsluzhivaniiaTO {
	Key: Guid!
}
input PrimaryCatalogObrabotkiObsluzhivaniiaTOModeli {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogNastroikaIntervalov {
	Key: Guid!
}
input PrimaryCatalogNastroikaIntervalovTablichnaiaChast {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogProfiliGruppDostupa {
	Key: Guid!
}
input PrimaryCatalogProfiliGruppDostupaRoli {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogProfiliGruppDostupaVidyDostupa {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogProfiliGruppDostupaZnacheniiaDostupa {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogProfiliGruppDostupaDostupPoPodsistemam {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogNastroikiDliaKurera {
	Key: Guid!
}
input PrimaryCatalogNastroikiDliaKureraSostavNaimenovaniia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogTipyTsenNomenklaturyKontragentov {
	Key: Guid!
}
input PrimaryDocumentJournalTsenoobrazovanie {
	Ref: String!
	RefType: String!
}
input PrimaryCatalogEdinitsyIzmereniia {
	Key: Guid!
}
input PrimaryCatalogStatiDvizheniiaDenezhnykhSredstv {
	Key: Guid!
}
input PrimaryDocumentInkassovoePorucheniePoluchennoe {
	Key: Guid!
}
input PrimaryDocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogNastroikiObmenaDannymiShtrikhM {
	Key: Guid!
}
input PrimaryCatalogStatiZatrat {
	Key: Guid!
}
input PrimaryDocumentVozvratTovarovOtPokupatelia {
	Key: Guid!
}
input PrimaryDocumentVozvratTovarovOtPokupateliaTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentVozvratTovarovOtPokupateliaUslugi {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentZakazPostavshchiku {
	Key: Guid!
}
input PrimaryDocumentZakazPostavshchikuTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogSkidkiNatsenki {
	Key: Guid!
}
input PrimaryCatalogSkidkiNatsenkiUsloviiaPredostavleniia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogSkidkiNatsenkiTsenovyeGruppy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogSkidkiNatsenkiNaborPodarkov {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogGruppyTsvetov {
	Key: Guid!
}
input PrimaryDocumentDokumentRaschetovSKontragentom {
	Key: Guid!
}
input PrimaryCatalogDogovoryEkvairinga {
	Key: Guid!
}
input PrimaryCatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKachestvo {
	Key: Guid!
}
input PrimaryDocumentUstanovkaTsenNomenklaturyKontragentov {
	Key: Guid!
}
input PrimaryDocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentUstanovkaTsenNomenklaturyKontragentovTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentProtsentPoterPoDavaltsam {
	Key: Guid!
}
input PrimaryDocumentProtsentPoterPoDavaltsamProtsenty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogTovarnyePozitsii {
	Key: Guid!
}
input PrimaryDocumentPlatezhnoePoruchenieIskhodiashchee {
	Key: Guid!
}
input PrimaryDocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogfmOrganizatsionnoPravovyeFormy {
	Key: Guid!
}
input PrimaryCatalogTipyTsenNomenklatury {
	Key: Guid!
}
input PrimaryCatalogStatiOtchetaPoProdazham {
	Key: Guid!
}
input PrimaryCatalogVidyKodirovokiIzdelii {
	Key: Guid!
}
input PrimaryCatalogVidyKodirovokiIzdeliiElementyKodirovki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentUstanovkaSkidokNomenklatury {
	Key: Guid!
}
input PrimaryDocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentUstanovkaSkidokNomenklaturyTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenok {
	Key: Guid!
}
input PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryOutPay {
	Key: Guid!
}
input PrimaryDocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentRaskhodnyiKassovyiOrderOplata {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentRaskhodnyiKassovyiOrderTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentSchetNaOplatuPostavshchika {
	Key: Guid!
}
input PrimaryDocumentSchetNaOplatuPostavshchikaTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentSchetNaOplatuPostavshchikaUslugi {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentReestrSpetssviaz {
	Key: Guid!
}
input PrimaryDocumentReestrSpetssviazKlienty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentJournalKassovyeDokumenty {
	Ref: String!
	RefType: String!
}
input PrimaryInitialInstance {
	Key: Guid!
}
input PrimaryDocumentVvodNachalnykhOstatkovVzaimoraschety {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentVvodNachalnykhOstatkovTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryPosting {
	Key: Guid!
}
input PrimaryDocumentOprikhodovanieTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOprikhodovanieTovarovSertifikaty {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKomplekty {
	Key: Guid!
}
input PrimaryDocumentPereotsenkaTovarovPriniatykhNaKomissiiu {
	Key: Guid!
}
input PrimaryDocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentElektronnoePismo {
	Key: Guid!
}
input PrimaryDocumentElektronnoePismoKomuTCh {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentElektronnoePismoKopiiTCh {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentElektronnoePismoSkrytyeKopiiTCh {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogGruppyDefektov {
	Key: Guid!
}
input PrimaryCatalogfmAnketaKlientaBenefitsariia {
	Key: Guid!
}
input PrimaryCatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogTsenovyeGruppy {
	Key: Guid!
}
input PrimaryCatalogPravilaTsenoobrazovaniia {
	Key: Guid!
}
input PrimaryCatalogPravilaTsenoobrazovaniiaTsenovyeGruppy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentObieiavlenieNaVznosNalichnymi {
	Key: Guid!
}
input PrimaryCatalogValiuty {
	Key: Guid!
}
input PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku {
	Key: Guid!
}
input PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKassyKKM {
	Key: Guid!
}
input PrimaryProbe {
	Key: Guid!
}
input PrimaryCatalogGruppyDostupa {
	Key: Guid!
}
input PrimaryCatalogGruppyDostupaPolzovateli {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogGruppyDostupaVidyDostupa {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogGruppyDostupaZnacheniiaDostupa {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogGruppyDostupaDostupPoPodsistemam {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogVidyKontaktnoiInformatsii {
	Key: Guid!
}
input PrimaryCatalogNomenklaturnyeGruppy {
	Key: Guid!
}
input PrimaryDocumentReestrSchetov {
	Key: Guid!
}
input PrimaryDocumentReestrSchetovReestr {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu {
	Key: Guid!
}
input PrimaryDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKlassifikatorStranMira {
	Key: Guid!
}
input PrimaryCatalogKlassifikatorEdinitsIzmereniia {
	Key: Guid!
}
input PrimaryCatalogNastroikiRMK {
	Key: Guid!
}
input PrimaryCatalogNastroikiRMKPoriadokPrimeneniiaSkidok {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogNastroikiRMKSostavNaimenovaniia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogKharakteristikiNomenklatury {
	Key: Guid!
}
input PrimaryCatalogKharakteristikiNomenklaturySpetsifikatsiia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtborTovarov {
	Key: Guid!
}
input PrimaryDocumentOtborTovarovTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtborTovarovTovaryKOtboru {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogSposobyDostavkiTovara {
	Key: Guid!
}
input PrimaryCatalogPodrazdeleniia {
	Key: Guid!
}
input PrimaryDocumentJournalPreiskuranty {
	Ref: String!
	RefType: String!
}
input PrimaryCatalogRelizyIuvelirnykhSalonov {
	Key: Guid!
}
input PrimaryCatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetKomissioneraOProdazhakh {
	Key: Guid!
}
input PrimaryDocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentOtchetKomissioneraOProdazhakhTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogTovarnyeKategorii {
	Key: Guid!
}
input PrimaryCatalogDokumentyUdostoveriaiushchieLichnost {
	Key: Guid!
}
input PrimaryCatalogFiltryDliaElektronnykhPisem {
	Key: Guid!
}
input PrimaryCatalogFiltryDliaElektronnykhPisemDeistviiaFiltra {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogFiltryDliaElektronnykhPisemUsloviiaFiltra {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPreiskurantTsenNaTsvKamni {
	Key: Guid!
}
input PrimaryDocumentPreiskurantTsenNaTsvKamniTablitsa {
	Key: Guid!
	LineNumber: Int64!
}
input PrimarySize {
	Key: Guid!
}
input PrimaryCatalogTipyDragotsennykhMetallov {
	Key: Guid!
}
input PrimaryDocumentTelemarketing {
	Key: Guid!
}
input PrimaryDocumentTelemarketingUchastniki {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentVozvratDavalcheskogoMetalla {
	Key: Guid!
}
input PrimaryCatalogAdresnyeSokrashcheniia {
	Key: Guid!
}
input PrimaryDocumentRassylkaAnket {
	Key: Guid!
}
input PrimaryDocumentRassylkaAnketVlozheniia {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentRassylkaAnketPoluchateli {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogVidyDeiatelnostiKontragentov {
	Key: Guid!
}
input PrimaryCatalogTorgovoeOborudovanie {
	Key: Guid!
}
input PrimaryCatalogSkhemyRealizatsii {
	Key: Guid!
}
input PrimaryCatalogSkhemyRealizatsiiEtapySkhemy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogPodkliuchaemoeOborudovanie {
	Key: Guid!
}
input PrimaryDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii {
	Key: Guid!
}
input PrimaryDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogGabarity {
	Key: Guid!
}
input PrimaryDocumentZakazKlienta {
	Key: Guid!
}
input PrimaryDocumentZakazKlientaTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryArriveFromManufacturing {
	Key: Guid!
}
input PrimaryArriveFromManufacturingInstance {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPostuplenieProduktsiiIzProizvodstvaMaterialy {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentJournalZakazyPostavshchikam {
	Ref: String!
	RefType: String!
}
input PrimaryDocumentJournalSkladskieDokumenty {
	Ref: String!
	RefType: String!
}
input PrimaryCatalogsmsUsloviiaOtboraDiskontnykhKart {
	Key: Guid!
}
input PrimaryArrive {
	Key: Guid!
}
input PrimaryDocumentPostuplenieTovarovUslugTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPostuplenieTovarovUslugUslugi {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentSchetFakturaVydannyi {
	Key: Guid!
}
input PrimaryDocumentPlanProdazhPoSalonam {
	Key: Guid!
}
input PrimaryDocumentPlanProdazhPoSalonamSalony {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPlanProdazhPoSalonamDniPoGrafiku {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogBankovskieScheta {
	Key: Guid!
}
input PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakh {
	Key: Guid!
}
input PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakhTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPeredachaVRemont {
	Key: Guid!
}
input PrimaryDocumentPeredachaVRemontIzdeliiaPriniatyeVRemont {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryDocumentPeredachaVRemontTovary {
	Key: Guid!
	LineNumber: Int64!
}
input PrimaryCatalogPolzovateli {
	Key: Guid!
}
input PrimaryCatalogTsenovyeKoridory {
	Key: Guid!
}
input PrimaryCatalogGruppySkladov {
	Key: Guid!
}
input PrimaryCatalogGruppySkladovSklady {
	Key: Guid!
	LineNumber: Int64!
}
type Query {
	AccumulationRegisterPartiiTovarovVProizvodstve(Key: PrimaryAccumulationRegisterPartiiTovarovVProizvodstve!): AccumulationRegisterPartiiTovarovVProizvodstve
	AccumulationRegisterPartiiTovarovVProizvodstves(): AccumulationRegisterPartiiTovarovVProizvodstve
	AccumulationRegisterPartiiTovarovVProizvodstveRecordType(Key: PrimaryAccumulationRegisterPartiiTovarovVProizvodstveRecordType!): AccumulationRegisterPartiiTovarovVProizvodstveRecordType
	AccumulationRegisterPartiiTovarovVProizvodstveRecordTypes(): AccumulationRegisterPartiiTovarovVProizvodstveRecordType
	AccumulationRegisterVzaimoraschetySPodotchetnymiLitsami(Key: PrimaryAccumulationRegisterVzaimoraschetySPodotchetnymiLitsami!): AccumulationRegisterVzaimoraschetySPodotchetnymiLitsami
	AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamis(): AccumulationRegisterVzaimoraschetySPodotchetnymiLitsami
	AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType(Key: PrimaryAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType!): AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType
	AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordTypes(): AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType
	AccumulationRegisterVnutrennieZakazy(Key: PrimaryAccumulationRegisterVnutrennieZakazy!): AccumulationRegisterVnutrennieZakazy
	AccumulationRegisterVnutrennieZakazys(): AccumulationRegisterVnutrennieZakazy
	AccumulationRegisterVnutrennieZakazyRecordType(Key: PrimaryAccumulationRegisterVnutrennieZakazyRecordType!): AccumulationRegisterVnutrennieZakazyRecordType
	AccumulationRegisterVnutrennieZakazyRecordTypes(): AccumulationRegisterVnutrennieZakazyRecordType
	AccumulationRegisterDenezhnyeSredstvaKomitenta(Key: PrimaryAccumulationRegisterDenezhnyeSredstvaKomitenta!): AccumulationRegisterDenezhnyeSredstvaKomitenta
	AccumulationRegisterDenezhnyeSredstvaKomitentas(): AccumulationRegisterDenezhnyeSredstvaKomitenta
	AccumulationRegisterDenezhnyeSredstvaKomitentaRecordType(Key: PrimaryAccumulationRegisterDenezhnyeSredstvaKomitentaRecordType!): AccumulationRegisterDenezhnyeSredstvaKomitentaRecordType
	AccumulationRegisterDenezhnyeSredstvaKomitentaRecordTypes(): AccumulationRegisterDenezhnyeSredstvaKomitentaRecordType
	AccumulationRegisterZakazyKlientov(Key: PrimaryAccumulationRegisterZakazyKlientov!): AccumulationRegisterZakazyKlientov
	AccumulationRegisterZakazyKlientovs(): AccumulationRegisterZakazyKlientov
	AccumulationRegisterZakazyKlientovRecordType(Key: PrimaryAccumulationRegisterZakazyKlientovRecordType!): AccumulationRegisterZakazyKlientovRecordType
	AccumulationRegisterZakazyKlientovRecordTypes(): AccumulationRegisterZakazyKlientovRecordType
	AccumulationRegisterSummyPoFinmonitoringuRoznitsa(Key: PrimaryAccumulationRegisterSummyPoFinmonitoringuRoznitsa!): AccumulationRegisterSummyPoFinmonitoringuRoznitsa
	AccumulationRegisterSummyPoFinmonitoringuRoznitsas(): AccumulationRegisterSummyPoFinmonitoringuRoznitsa
	AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType(Key: PrimaryAccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType!): AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType
	AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordTypes(): AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType
	AccumulationRegisterDenezhnyeSredstvaKPolucheniiu(Key: PrimaryAccumulationRegisterDenezhnyeSredstvaKPolucheniiu!): AccumulationRegisterDenezhnyeSredstvaKPolucheniiu
	AccumulationRegisterDenezhnyeSredstvaKPolucheniius(): AccumulationRegisterDenezhnyeSredstvaKPolucheniiu
	AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType(Key: PrimaryAccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType!): AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType
	AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordTypes(): AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType
	AccumulationRegisterProdazhiPoDiskontnymKartam(Key: PrimaryAccumulationRegisterProdazhiPoDiskontnymKartam!): AccumulationRegisterProdazhiPoDiskontnymKartam
	AccumulationRegisterProdazhiPoDiskontnymKartams(): AccumulationRegisterProdazhiPoDiskontnymKartam
	AccumulationRegisterProdazhiPoDiskontnymKartamRecordType(Key: PrimaryAccumulationRegisterProdazhiPoDiskontnymKartamRecordType!): AccumulationRegisterProdazhiPoDiskontnymKartamRecordType
	AccumulationRegisterProdazhiPoDiskontnymKartamRecordTypes(): AccumulationRegisterProdazhiPoDiskontnymKartamRecordType
	AccumulationRegisterTovaryPoluchennye(Key: PrimaryAccumulationRegisterTovaryPoluchennye!): AccumulationRegisterTovaryPoluchennye
	AccumulationRegisterTovaryPoluchennyes(): AccumulationRegisterTovaryPoluchennye
	AccumulationRegisterTovaryPoluchennyeRecordType(Key: PrimaryAccumulationRegisterTovaryPoluchennyeRecordType!): AccumulationRegisterTovaryPoluchennyeRecordType
	AccumulationRegisterTovaryPoluchennyeRecordTypes(): AccumulationRegisterTovaryPoluchennyeRecordType
	AccumulationRegisterSvobodnyeOstatki(Key: PrimaryAccumulationRegisterSvobodnyeOstatki!): AccumulationRegisterSvobodnyeOstatki
	AccumulationRegisterSvobodnyeOstatkis(): AccumulationRegisterSvobodnyeOstatki
	AccumulationRegisterSvobodnyeOstatkiRecordType(Key: PrimaryAccumulationRegisterSvobodnyeOstatkiRecordType!): AccumulationRegisterSvobodnyeOstatkiRecordType
	AccumulationRegisterSvobodnyeOstatkiRecordTypes(): AccumulationRegisterSvobodnyeOstatkiRecordType
	AccumulationRegisterSummyVRassrochku(Key: PrimaryAccumulationRegisterSummyVRassrochku!): AccumulationRegisterSummyVRassrochku
	AccumulationRegisterSummyVRassrochkus(): AccumulationRegisterSummyVRassrochku
	AccumulationRegisterSummyVRassrochkuRecordType(Key: PrimaryAccumulationRegisterSummyVRassrochkuRecordType!): AccumulationRegisterSummyVRassrochkuRecordType
	AccumulationRegisterSummyVRassrochkuRecordTypes(): AccumulationRegisterSummyVRassrochkuRecordType
	AccumulationRegisterGrafikPlatezhei(Key: PrimaryAccumulationRegisterGrafikPlatezhei!): AccumulationRegisterGrafikPlatezhei
	AccumulationRegisterGrafikPlatezheis(): AccumulationRegisterGrafikPlatezhei
	AccumulationRegisterGrafikPlatezheiRecordType(Key: PrimaryAccumulationRegisterGrafikPlatezheiRecordType!): AccumulationRegisterGrafikPlatezheiRecordType
	AccumulationRegisterGrafikPlatezheiRecordTypes(): AccumulationRegisterGrafikPlatezheiRecordType
	AccumulationRegisterRoznichnaiaVyruchka(Key: PrimaryAccumulationRegisterRoznichnaiaVyruchka!): AccumulationRegisterRoznichnaiaVyruchka
	AccumulationRegisterRoznichnaiaVyruchkas(): AccumulationRegisterRoznichnaiaVyruchka
	AccumulationRegisterRoznichnaiaVyruchkaRecordType(Key: PrimaryAccumulationRegisterRoznichnaiaVyruchkaRecordType!): AccumulationRegisterRoznichnaiaVyruchkaRecordType
	AccumulationRegisterRoznichnaiaVyruchkaRecordTypes(): AccumulationRegisterRoznichnaiaVyruchkaRecordType
	AccumulationRegisterTovaryVPuti(Key: PrimaryAccumulationRegisterTovaryVPuti!): AccumulationRegisterTovaryVPuti
	AccumulationRegisterTovaryVPutis(): AccumulationRegisterTovaryVPuti
	AccumulationRegisterTovaryVPutiRecordType(Key: PrimaryAccumulationRegisterTovaryVPutiRecordType!): AccumulationRegisterTovaryVPutiRecordType
	AccumulationRegisterTovaryVPutiRecordTypes(): AccumulationRegisterTovaryVPutiRecordType
	AccumulationRegisterPoteriMetallaVProizvodstve(Key: PrimaryAccumulationRegisterPoteriMetallaVProizvodstve!): AccumulationRegisterPoteriMetallaVProizvodstve
	AccumulationRegisterPoteriMetallaVProizvodstves(): AccumulationRegisterPoteriMetallaVProizvodstve
	AccumulationRegisterPoteriMetallaVProizvodstveRecordType(Key: PrimaryAccumulationRegisterPoteriMetallaVProizvodstveRecordType!): AccumulationRegisterPoteriMetallaVProizvodstveRecordType
	AccumulationRegisterPoteriMetallaVProizvodstveRecordTypes(): AccumulationRegisterPoteriMetallaVProizvodstveRecordType
	AccumulationRegisterPartiiTovarovNaSkladakh(Key: PrimaryAccumulationRegisterPartiiTovarovNaSkladakh!): AccumulationRegisterPartiiTovarovNaSkladakh
	AccumulationRegisterPartiiTovarovNaSkladakhs(): AccumulationRegisterPartiiTovarovNaSkladakh
	ProductActionDocument(Key: PrimaryProductActionDocument!): ProductActionDocument
	ProductActionDocuments(): ProductActionDocument
	AccumulationRegisterSummyDokumentovDliaObmena(Key: PrimaryAccumulationRegisterSummyDokumentovDliaObmena!): AccumulationRegisterSummyDokumentovDliaObmena
	AccumulationRegisterSummyDokumentovDliaObmenas(): AccumulationRegisterSummyDokumentovDliaObmena
	AccumulationRegisterSummyDokumentovDliaObmenaRecordType(Key: PrimaryAccumulationRegisterSummyDokumentovDliaObmenaRecordType!): AccumulationRegisterSummyDokumentovDliaObmenaRecordType
	AccumulationRegisterSummyDokumentovDliaObmenaRecordTypes(): AccumulationRegisterSummyDokumentovDliaObmenaRecordType
	AccumulationRegisterDvizheniiaDenezhnykhSredstv(Key: PrimaryAccumulationRegisterDvizheniiaDenezhnykhSredstv!): AccumulationRegisterDvizheniiaDenezhnykhSredstv
	AccumulationRegisterDvizheniiaDenezhnykhSredstvs(): AccumulationRegisterDvizheniiaDenezhnykhSredstv
	AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType(Key: PrimaryAccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType!): AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType
	AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordTypes(): AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType
	AccumulationRegisterProdazhiPoStatiam(Key: PrimaryAccumulationRegisterProdazhiPoStatiam!): AccumulationRegisterProdazhiPoStatiam
	AccumulationRegisterProdazhiPoStatiams(): AccumulationRegisterProdazhiPoStatiam
	AccumulationRegisterProdazhiPoStatiamRecordType(Key: PrimaryAccumulationRegisterProdazhiPoStatiamRecordType!): AccumulationRegisterProdazhiPoStatiamRecordType
	AccumulationRegisterProdazhiPoStatiamRecordTypes(): AccumulationRegisterProdazhiPoStatiamRecordType
	InformationRegisterTsenyNomenklatury(Key: PrimaryInformationRegisterTsenyNomenklatury!): InformationRegisterTsenyNomenklatury
	InformationRegisterTsenyNomenklaturys(): InformationRegisterTsenyNomenklatury
	InformationRegisterTsenyNomenklaturyRecordType(Key: PrimaryInformationRegisterTsenyNomenklaturyRecordType!): InformationRegisterTsenyNomenklaturyRecordType
	InformationRegisterTsenyNomenklaturyRecordTypes(): InformationRegisterTsenyNomenklaturyRecordType
	AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse(Key: PrimaryAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse!): AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse
	AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitses(): AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse
	AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType(Key: PrimaryAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType!): AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType
	AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordTypes(): AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType
	AccumulationRegisterDenezhnyeSredstvaVRezerve(Key: PrimaryAccumulationRegisterDenezhnyeSredstvaVRezerve!): AccumulationRegisterDenezhnyeSredstvaVRezerve
	AccumulationRegisterDenezhnyeSredstvaVRezerves(): AccumulationRegisterDenezhnyeSredstvaVRezerve
	AccumulationRegisterDenezhnyeSredstvaVRezerveRecordType(Key: PrimaryAccumulationRegisterDenezhnyeSredstvaVRezerveRecordType!): AccumulationRegisterDenezhnyeSredstvaVRezerveRecordType
	AccumulationRegisterDenezhnyeSredstvaVRezerveRecordTypes(): AccumulationRegisterDenezhnyeSredstvaVRezerveRecordType
	AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh(Key: PrimaryAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh!): AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh
	AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhs(): AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh
	AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType(Key: PrimaryAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType!): AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType
	AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordTypes(): AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType
	AccumulationRegisterDavalcheskiiMetallPoteri(Key: PrimaryAccumulationRegisterDavalcheskiiMetallPoteri!): AccumulationRegisterDavalcheskiiMetallPoteri
	AccumulationRegisterDavalcheskiiMetallPoteris(): AccumulationRegisterDavalcheskiiMetallPoteri
	AccumulationRegisterDavalcheskiiMetallPoteriRecordType(Key: PrimaryAccumulationRegisterDavalcheskiiMetallPoteriRecordType!): AccumulationRegisterDavalcheskiiMetallPoteriRecordType
	AccumulationRegisterDavalcheskiiMetallPoteriRecordTypes(): AccumulationRegisterDavalcheskiiMetallPoteriRecordType
	InformationRegisterTsenyPoPreiskurantu(Key: PrimaryInformationRegisterTsenyPoPreiskurantu!): InformationRegisterTsenyPoPreiskurantu
	InformationRegisterTsenyPoPreiskurantus(): InformationRegisterTsenyPoPreiskurantu
	InformationRegisterTsenyPoPreiskurantuRecordType(Key: PrimaryInformationRegisterTsenyPoPreiskurantuRecordType!): InformationRegisterTsenyPoPreiskurantuRecordType
	InformationRegisterTsenyPoPreiskurantuRecordTypes(): InformationRegisterTsenyPoPreiskurantuRecordType
	AccumulationRegisterTovaryVOtbore(Key: PrimaryAccumulationRegisterTovaryVOtbore!): AccumulationRegisterTovaryVOtbore
	AccumulationRegisterTovaryVOtbores(): AccumulationRegisterTovaryVOtbore
	AccumulationRegisterTovaryVOtboreRecordType(Key: PrimaryAccumulationRegisterTovaryVOtboreRecordType!): AccumulationRegisterTovaryVOtboreRecordType
	AccumulationRegisterTovaryVOtboreRecordTypes(): AccumulationRegisterTovaryVOtboreRecordType
	AccumulationRegisterRealizovannyeTovary(Key: PrimaryAccumulationRegisterRealizovannyeTovary!): AccumulationRegisterRealizovannyeTovary
	AccumulationRegisterRealizovannyeTovarys(): AccumulationRegisterRealizovannyeTovary
	AccumulationRegisterRealizovannyeTovaryRecordType(Key: PrimaryAccumulationRegisterRealizovannyeTovaryRecordType!): AccumulationRegisterRealizovannyeTovaryRecordType
	AccumulationRegisterRealizovannyeTovaryRecordTypes(): AccumulationRegisterRealizovannyeTovaryRecordType
	AccumulationRegisterDenezhnyeSredstvaKomissionera(Key: PrimaryAccumulationRegisterDenezhnyeSredstvaKomissionera!): AccumulationRegisterDenezhnyeSredstvaKomissionera
	AccumulationRegisterDenezhnyeSredstvaKomissioneras(): AccumulationRegisterDenezhnyeSredstvaKomissionera
	AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType(Key: PrimaryAccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType!): AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType
	AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordTypes(): AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType
	AccumulationRegisterProdazhi1(Key: PrimaryAccumulationRegisterProdazhi1!): AccumulationRegisterProdazhi1
	AccumulationRegisterProdazhi1s(): AccumulationRegisterProdazhi1
	AccumulationRegisterProdazhi1RecordType(Key: PrimaryAccumulationRegisterProdazhi1RecordType!): AccumulationRegisterProdazhi1RecordType
	AccumulationRegisterProdazhi1RecordTypes(): AccumulationRegisterProdazhi1RecordType
	AccumulationRegisterTovaryNaSkladakhAM(Key: PrimaryAccumulationRegisterTovaryNaSkladakhAM!): AccumulationRegisterTovaryNaSkladakhAM
	AccumulationRegisterTovaryNaSkladakhAMs(): AccumulationRegisterTovaryNaSkladakhAM
	AccumulationRegisterTovaryNaSkladakhAMRecordType(Key: PrimaryAccumulationRegisterTovaryNaSkladakhAMRecordType!): AccumulationRegisterTovaryNaSkladakhAMRecordType
	AccumulationRegisterTovaryNaSkladakhAMRecordTypes(): AccumulationRegisterTovaryNaSkladakhAMRecordType
	AccumulationRegisterSummyPoFinmonitoringu(Key: PrimaryAccumulationRegisterSummyPoFinmonitoringu!): AccumulationRegisterSummyPoFinmonitoringu
	AccumulationRegisterSummyPoFinmonitoringus(): AccumulationRegisterSummyPoFinmonitoringu
	AccumulationRegisterSummyPoFinmonitoringuRecordType(Key: PrimaryAccumulationRegisterSummyPoFinmonitoringuRecordType!): AccumulationRegisterSummyPoFinmonitoringuRecordType
	AccumulationRegisterSummyPoFinmonitoringuRecordTypes(): AccumulationRegisterSummyPoFinmonitoringuRecordType
	InformationRegisterTsenyNomenklaturyKontragentov(Key: PrimaryInformationRegisterTsenyNomenklaturyKontragentov!): InformationRegisterTsenyNomenklaturyKontragentov
	InformationRegisterTsenyNomenklaturyKontragentovs(): InformationRegisterTsenyNomenklaturyKontragentov
	InformationRegisterTsenyNomenklaturyKontragentovRecordType(Key: PrimaryInformationRegisterTsenyNomenklaturyKontragentovRecordType!): InformationRegisterTsenyNomenklaturyKontragentovRecordType
	InformationRegisterTsenyNomenklaturyKontragentovRecordTypes(): InformationRegisterTsenyNomenklaturyKontragentovRecordType
	AccumulationRegisterVzaimoraschetySKontragentami(Key: PrimaryAccumulationRegisterVzaimoraschetySKontragentami!): AccumulationRegisterVzaimoraschetySKontragentami
	AccumulationRegisterVzaimoraschetySKontragentamis(): AccumulationRegisterVzaimoraschetySKontragentami
	AccumulationRegisterVzaimoraschetySKontragentamiRecordType(Key: PrimaryAccumulationRegisterVzaimoraschetySKontragentamiRecordType!): AccumulationRegisterVzaimoraschetySKontragentamiRecordType
	AccumulationRegisterVzaimoraschetySKontragentamiRecordTypes(): AccumulationRegisterVzaimoraschetySKontragentamiRecordType
	AccumulationRegisterSummyPokupokPoDiskontnymKartam(Key: PrimaryAccumulationRegisterSummyPokupokPoDiskontnymKartam!): AccumulationRegisterSummyPokupokPoDiskontnymKartam
	AccumulationRegisterSummyPokupokPoDiskontnymKartams(): AccumulationRegisterSummyPokupokPoDiskontnymKartam
	AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType(Key: PrimaryAccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType!): AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType
	AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordTypes(): AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType
	AccumulationRegisterVypolneniePlanaProdazh(Key: PrimaryAccumulationRegisterVypolneniePlanaProdazh!): AccumulationRegisterVypolneniePlanaProdazh
	AccumulationRegisterVypolneniePlanaProdazhs(): AccumulationRegisterVypolneniePlanaProdazh
	AccumulationRegisterVypolneniePlanaProdazhRecordType(Key: PrimaryAccumulationRegisterVypolneniePlanaProdazhRecordType!): AccumulationRegisterVypolneniePlanaProdazhRecordType
	AccumulationRegisterVypolneniePlanaProdazhRecordTypes(): AccumulationRegisterVypolneniePlanaProdazhRecordType
	AccumulationRegisterDavalcheskiiMetall(Key: PrimaryAccumulationRegisterDavalcheskiiMetall!): AccumulationRegisterDavalcheskiiMetall
	AccumulationRegisterDavalcheskiiMetalls(): AccumulationRegisterDavalcheskiiMetall
	AccumulationRegisterDavalcheskiiMetallRecordType(Key: PrimaryAccumulationRegisterDavalcheskiiMetallRecordType!): AccumulationRegisterDavalcheskiiMetallRecordType
	AccumulationRegisterDavalcheskiiMetallRecordTypes(): AccumulationRegisterDavalcheskiiMetallRecordType
	AccumulationRegisterDenezhnyeSredstva(Key: PrimaryAccumulationRegisterDenezhnyeSredstva!): AccumulationRegisterDenezhnyeSredstva
	AccumulationRegisterDenezhnyeSredstvas(): AccumulationRegisterDenezhnyeSredstva
	AccumulationRegisterDenezhnyeSredstvaRecordType(Key: PrimaryAccumulationRegisterDenezhnyeSredstvaRecordType!): AccumulationRegisterDenezhnyeSredstvaRecordType
	AccumulationRegisterDenezhnyeSredstvaRecordTypes(): AccumulationRegisterDenezhnyeSredstvaRecordType
	AccumulationRegisterTovaryPeredannye(Key: PrimaryAccumulationRegisterTovaryPeredannye!): AccumulationRegisterTovaryPeredannye
	AccumulationRegisterTovaryPeredannyes(): AccumulationRegisterTovaryPeredannye
	AccumulationRegisterTovaryPeredannyeRecordType(Key: PrimaryAccumulationRegisterTovaryPeredannyeRecordType!): AccumulationRegisterTovaryPeredannyeRecordType
	AccumulationRegisterTovaryPeredannyeRecordTypes(): AccumulationRegisterTovaryPeredannyeRecordType
	AccumulationRegisterDenezhnyeSredstvaKSpisaniiu(Key: PrimaryAccumulationRegisterDenezhnyeSredstvaKSpisaniiu!): AccumulationRegisterDenezhnyeSredstvaKSpisaniiu
	AccumulationRegisterDenezhnyeSredstvaKSpisaniius(): AccumulationRegisterDenezhnyeSredstvaKSpisaniiu
	AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType(Key: PrimaryAccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType!): AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType
	AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordTypes(): AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType
	CatalogDogovoryKontragentov(Key: PrimaryCatalogDogovoryKontragentov!): CatalogDogovoryKontragentov
	CatalogDogovoryKontragentovs(): CatalogDogovoryKontragentov
	Order(Key: PrimaryOrder!): Order
	Orders(): Order
	DocumentChekKKMOplata(Key: PrimaryDocumentChekKKMOplata!): DocumentChekKKMOplata
	DocumentChekKKMOplatas(): DocumentChekKKMOplata
	DocumentChekKKMOplataSertifikatami(Key: PrimaryDocumentChekKKMOplataSertifikatami!): DocumentChekKKMOplataSertifikatami
	DocumentChekKKMOplataSertifikatamis(): DocumentChekKKMOplataSertifikatami
	DocumentChekKKMProdazhaSertifikatov(Key: PrimaryDocumentChekKKMProdazhaSertifikatov!): DocumentChekKKMProdazhaSertifikatov
	DocumentChekKKMProdazhaSertifikatovs(): DocumentChekKKMProdazhaSertifikatov
	DocumentChekKKMTovary(Key: PrimaryDocumentChekKKMTovary!): DocumentChekKKMTovary
	DocumentChekKKMTovarys(): DocumentChekKKMTovary
	DocumentChekKKMDokumentyObmena(Key: PrimaryDocumentChekKKMDokumentyObmena!): DocumentChekKKMDokumentyObmena
	DocumentChekKKMDokumentyObmenas(): DocumentChekKKMDokumentyObmena
	DocumentChekKKMDogovoraRassrochkiProdazha(Key: PrimaryDocumentChekKKMDogovoraRassrochkiProdazha!): DocumentChekKKMDogovoraRassrochkiProdazha
	DocumentChekKKMDogovoraRassrochkiProdazhas(): DocumentChekKKMDogovoraRassrochkiProdazha
	DocumentChekKKMDogovoraRassrochkiOplata(Key: PrimaryDocumentChekKKMDogovoraRassrochkiOplata!): DocumentChekKKMDogovoraRassrochkiOplata
	DocumentChekKKMDogovoraRassrochkiOplatas(): DocumentChekKKMDogovoraRassrochkiOplata
	DocumentChekKKMOplataBallami(Key: PrimaryDocumentChekKKMOplataBallami!): DocumentChekKKMOplataBallami
	DocumentChekKKMOplataBallamis(): DocumentChekKKMOplataBallami
	DocumentChekKKMSkidkiNatsenki(Key: PrimaryDocumentChekKKMSkidkiNatsenki!): DocumentChekKKMSkidkiNatsenki
	DocumentChekKKMSkidkiNatsenkis(): DocumentChekKKMSkidkiNatsenki
	DocumentChekKKMUpravliaemyeSkidki(Key: PrimaryDocumentChekKKMUpravliaemyeSkidki!): DocumentChekKKMUpravliaemyeSkidki
	DocumentChekKKMUpravliaemyeSkidkis(): DocumentChekKKMUpravliaemyeSkidki
	DocumentChekKKMPodarki(Key: PrimaryDocumentChekKKMPodarki!): DocumentChekKKMPodarki
	DocumentChekKKMPodarkis(): DocumentChekKKMPodarki
	DocumentChekKKMKupony(Key: PrimaryDocumentChekKKMKupony!): DocumentChekKKMKupony
	DocumentChekKKMKuponys(): DocumentChekKKMKupony
	DocumentPereotsenkaValiutnykhSredstv(Key: PrimaryDocumentPereotsenkaValiutnykhSredstv!): DocumentPereotsenkaValiutnykhSredstv
	DocumentPereotsenkaValiutnykhSredstvs(): DocumentPereotsenkaValiutnykhSredstv
	CatalogTipySkidokNatsenok(Key: PrimaryCatalogTipySkidokNatsenok!): CatalogTipySkidokNatsenok
	CatalogTipySkidokNatsenoks(): CatalogTipySkidokNatsenok
	CatalogTipySkidokNatsenokVremiaPoDniamNedeli(Key: PrimaryCatalogTipySkidokNatsenokVremiaPoDniamNedeli!): CatalogTipySkidokNatsenokVremiaPoDniamNedeli
	CatalogTipySkidokNatsenokVremiaPoDniamNedelis(): CatalogTipySkidokNatsenokVremiaPoDniamNedeli
	CatalogVidyKodirovokiTsepei(Key: PrimaryCatalogVidyKodirovokiTsepei!): CatalogVidyKodirovokiTsepei
	CatalogVidyKodirovokiTsepeis(): CatalogVidyKodirovokiTsepei
	CatalogVidyKodirovokiTsepeiElementyKodirovki(Key: PrimaryCatalogVidyKodirovokiTsepeiElementyKodirovki!): CatalogVidyKodirovokiTsepeiElementyKodirovki
	CatalogVidyKodirovokiTsepeiElementyKodirovkis(): CatalogVidyKodirovokiTsepeiElementyKodirovki
	CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv(Key: PrimaryCatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv!): CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv
	CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvs(): CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv
	DocumentOtchetKomitentuOProdazhakh(Key: PrimaryDocumentOtchetKomitentuOProdazhakh!): DocumentOtchetKomitentuOProdazhakh
	DocumentOtchetKomitentuOProdazhakhs(): DocumentOtchetKomitentuOProdazhakh
	DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva(Key: PrimaryDocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva!): DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva
	DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvas(): DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva
	DocumentOtchetKomitentuOProdazhakhTovary(Key: PrimaryDocumentOtchetKomitentuOProdazhakhTovary!): DocumentOtchetKomitentuOProdazhakhTovary
	DocumentOtchetKomitentuOProdazhakhTovarys(): DocumentOtchetKomitentuOProdazhakhTovary
	CatalogKassy(Key: PrimaryCatalogKassy!): CatalogKassy
	CatalogKassys(): CatalogKassy
	CatalogKassiry(Key: PrimaryCatalogKassiry!): CatalogKassiry
	CatalogKassirys(): CatalogKassiry
	DocumentZaiavkaNaPereotsenkuTovarov(Key: PrimaryDocumentZaiavkaNaPereotsenkuTovarov!): DocumentZaiavkaNaPereotsenkuTovarov
	DocumentZaiavkaNaPereotsenkuTovarovs(): DocumentZaiavkaNaPereotsenkuTovarov
	DocumentZaiavkaNaPereotsenkuTovarovTovary(Key: PrimaryDocumentZaiavkaNaPereotsenkuTovarovTovary!): DocumentZaiavkaNaPereotsenkuTovarovTovary
	DocumentZaiavkaNaPereotsenkuTovarovTovarys(): DocumentZaiavkaNaPereotsenkuTovarovTovary
	CatalogProizvodstvennyeUchastki(Key: PrimaryCatalogProizvodstvennyeUchastki!): CatalogProizvodstvennyeUchastki
	CatalogProizvodstvennyeUchastkis(): CatalogProizvodstvennyeUchastki
	DocumentZakrytieZakazovKlientov(Key: PrimaryDocumentZakrytieZakazovKlientov!): DocumentZakrytieZakazovKlientov
	DocumentZakrytieZakazovKlientovs(): DocumentZakrytieZakazovKlientov
	DocumentZakrytieZakazovKlientovZakazy(Key: PrimaryDocumentZakrytieZakazovKlientovZakazy!): DocumentZakrytieZakazovKlientovZakazy
	DocumentZakrytieZakazovKlientovZakazys(): DocumentZakrytieZakazovKlientovZakazy
	CatalogProekty(Key: PrimaryCatalogProekty!): CatalogProekty
	CatalogProektys(): CatalogProekty
	DocumentPlatezhnoePoruchenieVkhodiashchee(Key: PrimaryDocumentPlatezhnoePoruchenieVkhodiashchee!): DocumentPlatezhnoePoruchenieVkhodiashchee
	DocumentPlatezhnoePoruchenieVkhodiashchees(): DocumentPlatezhnoePoruchenieVkhodiashchee
	DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha(Key: PrimaryDocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha!): DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha
	DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhas(): DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha
	DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta(Key: PrimaryDocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta!): DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta
	DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentas(): DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta
	DocumentVydachaZakaza(Key: PrimaryDocumentVydachaZakaza!): DocumentVydachaZakaza
	DocumentVydachaZakazas(): DocumentVydachaZakaza
	CatalogFormyOgranki(Key: PrimaryCatalogFormyOgranki!): CatalogFormyOgranki
	CatalogFormyOgrankis(): CatalogFormyOgranki
	CatalogFormatyMagazinov(Key: PrimaryCatalogFormatyMagazinov!): CatalogFormatyMagazinov
	CatalogFormatyMagazinovs(): CatalogFormatyMagazinov
	CatalogRabochieMesta(Key: PrimaryCatalogRabochieMesta!): CatalogRabochieMesta
	CatalogRabochieMestas(): CatalogRabochieMesta
	CatalogNastroikiVypolneniiaObmena(Key: PrimaryCatalogNastroikiVypolneniiaObmena!): CatalogNastroikiVypolneniiaObmena
	CatalogNastroikiVypolneniiaObmenas(): CatalogNastroikiVypolneniiaObmena
	CatalogNastroikiVypolneniiaObmenaNastroikiObmena(Key: PrimaryCatalogNastroikiVypolneniiaObmenaNastroikiObmena!): CatalogNastroikiVypolneniiaObmenaNastroikiObmena
	CatalogNastroikiVypolneniiaObmenaNastroikiObmenas(): CatalogNastroikiVypolneniiaObmenaNastroikiObmena
	CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami(Key: PrimaryCatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami!): CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami
	CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamis(): CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami
	CatalogZnacheniiaSvoistvObieektov(Key: PrimaryCatalogZnacheniiaSvoistvObieektov!): CatalogZnacheniiaSvoistvObieektov
	CatalogZnacheniiaSvoistvObieektovs(): CatalogZnacheniiaSvoistvObieektov
	DocumentRealizatsiiaTovarovUslug(Key: PrimaryDocumentRealizatsiiaTovarovUslug!): DocumentRealizatsiiaTovarovUslug
	DocumentRealizatsiiaTovarovUslugs(): DocumentRealizatsiiaTovarovUslug
	DocumentRealizatsiiaTovarovUslugTovary(Key: PrimaryDocumentRealizatsiiaTovarovUslugTovary!): DocumentRealizatsiiaTovarovUslugTovary
	DocumentRealizatsiiaTovarovUslugTovarys(): DocumentRealizatsiiaTovarovUslugTovary
	DocumentRealizatsiiaTovarovUslugUslugi(Key: PrimaryDocumentRealizatsiiaTovarovUslugUslugi!): DocumentRealizatsiiaTovarovUslugUslugi
	DocumentRealizatsiiaTovarovUslugUslugis(): DocumentRealizatsiiaTovarovUslugUslugi
	DocumentSobytie(Key: PrimaryDocumentSobytie!): DocumentSobytie
	DocumentSobyties(): DocumentSobytie
	DocumentSobytieStoronnieLitsa(Key: PrimaryDocumentSobytieStoronnieLitsa!): DocumentSobytieStoronnieLitsa
	DocumentSobytieStoronnieLitsas(): DocumentSobytieStoronnieLitsa
	CatalogVariantyOtvetovOprosov(Key: PrimaryCatalogVariantyOtvetovOprosov!): CatalogVariantyOtvetovOprosov
	CatalogVariantyOtvetovOprosovs(): CatalogVariantyOtvetovOprosov
	CatalogGruppyPisemElektronnoiPochty(Key: PrimaryCatalogGruppyPisemElektronnoiPochty!): CatalogGruppyPisemElektronnoiPochty
	CatalogGruppyPisemElektronnoiPochtys(): CatalogGruppyPisemElektronnoiPochty
	CatalogGruppyPochtovoiRassylki(Key: PrimaryCatalogGruppyPochtovoiRassylki!): CatalogGruppyPochtovoiRassylki
	CatalogGruppyPochtovoiRassylkis(): CatalogGruppyPochtovoiRassylki
	CatalogNastroikiOtchetov(Key: PrimaryCatalogNastroikiOtchetov!): CatalogNastroikiOtchetov
	CatalogNastroikiOtchetovs(): CatalogNastroikiOtchetov
	CatalogNastroikiOtchetovGruppyNastroekIPolzovateli(Key: PrimaryCatalogNastroikiOtchetovGruppyNastroekIPolzovateli!): CatalogNastroikiOtchetovGruppyNastroekIPolzovateli
	CatalogNastroikiOtchetovGruppyNastroekIPolzovatelis(): CatalogNastroikiOtchetovGruppyNastroekIPolzovateli
	CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam(Key: PrimaryCatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam!): CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam
	CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartams(): CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam
	CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki(Key: PrimaryCatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki!): CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki
	CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkis(): CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki
	Department(Key: PrimaryDepartment!): Department
	Departments(): Department
	CatalogKodyVidovTovarov(Key: PrimaryCatalogKodyVidovTovarov!): CatalogKodyVidovTovarov
	CatalogKodyVidovTovarovs(): CatalogKodyVidovTovarov
	CatalogRassevy(Key: PrimaryCatalogRassevy!): CatalogRassevy
	CatalogRassevys(): CatalogRassevy
	CatalogPrichinyZakrytiiaZakazov(Key: PrimaryCatalogPrichinyZakrytiiaZakazov!): CatalogPrichinyZakrytiiaZakazov
	CatalogPrichinyZakrytiiaZakazovs(): CatalogPrichinyZakrytiiaZakazov
	CatalogSegmentyNomenklatury(Key: PrimaryCatalogSegmentyNomenklatury!): CatalogSegmentyNomenklatury
	CatalogSegmentyNomenklaturys(): CatalogSegmentyNomenklatury
	CatalogSostavStrokiCheka(Key: PrimaryCatalogSostavStrokiCheka!): CatalogSostavStrokiCheka
	CatalogSostavStrokiChekas(): CatalogSostavStrokiCheka
	CatalogUsloviiaPriemaIzdeliiNaKomissiiu(Key: PrimaryCatalogUsloviiaPriemaIzdeliiNaKomissiiu!): CatalogUsloviiaPriemaIzdeliiNaKomissiiu
	CatalogUsloviiaPriemaIzdeliiNaKomissiius(): CatalogUsloviiaPriemaIzdeliiNaKomissiiu
	CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok(Key: PrimaryCatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok!): CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok
	CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenoks(): CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok
	CatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei(Key: PrimaryCatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei!): CatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei
	CatalogIstochnikiInformatsiiPriObrashcheniiPokupateleis(): CatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei
	DocumentKorrektirovkaDolga(Key: PrimaryDocumentKorrektirovkaDolga!): DocumentKorrektirovkaDolga
	DocumentKorrektirovkaDolgas(): DocumentKorrektirovkaDolga
	DocumentKorrektirovkaDolgaSummyDolga(Key: PrimaryDocumentKorrektirovkaDolgaSummyDolga!): DocumentKorrektirovkaDolgaSummyDolga
	DocumentKorrektirovkaDolgaSummyDolgas(): DocumentKorrektirovkaDolgaSummyDolga
	PayType(Key: PrimaryPayType!): PayType
	PayTypes(): PayType
	CatalogKhranilishcheShablonov(Key: PrimaryCatalogKhranilishcheShablonov!): CatalogKhranilishcheShablonov
	CatalogKhranilishcheShablonovs(): CatalogKhranilishcheShablonov
	DocumentZaiavkaNaRaskhodovanieSredstv(Key: PrimaryDocumentZaiavkaNaRaskhodovanieSredstv!): DocumentZaiavkaNaRaskhodovanieSredstv
	DocumentZaiavkaNaRaskhodovanieSredstvs(): DocumentZaiavkaNaRaskhodovanieSredstv
	DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha(Key: PrimaryDocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha!): DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha
	DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhas(): DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha
	DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki(Key: PrimaryDocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki!): DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki
	DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkis(): DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki
	DocumentZakrytieZakazovPostavshchikam(Key: PrimaryDocumentZakrytieZakazovPostavshchikam!): DocumentZakrytieZakazovPostavshchikam
	DocumentZakrytieZakazovPostavshchikams(): DocumentZakrytieZakazovPostavshchikam
	DocumentZakrytieZakazovPostavshchikamZakazy(Key: PrimaryDocumentZakrytieZakazovPostavshchikamZakazy!): DocumentZakrytieZakazovPostavshchikamZakazy
	DocumentZakrytieZakazovPostavshchikamZakazys(): DocumentZakrytieZakazovPostavshchikamZakazy
	CatalogVidyKamnei(Key: PrimaryCatalogVidyKamnei!): CatalogVidyKamnei
	CatalogVidyKamneis(): CatalogVidyKamnei
	DocumentAnketyKlientovDliaFinMonitoringa(Key: PrimaryDocumentAnketyKlientovDliaFinMonitoringa!): DocumentAnketyKlientovDliaFinMonitoringa
	DocumentAnketyKlientovDliaFinMonitoringas(): DocumentAnketyKlientovDliaFinMonitoringa
	DocumentAnketyKlientovDliaFinMonitoringaAnkety(Key: PrimaryDocumentAnketyKlientovDliaFinMonitoringaAnkety!): DocumentAnketyKlientovDliaFinMonitoringaAnkety
	DocumentAnketyKlientovDliaFinMonitoringaAnketys(): DocumentAnketyKlientovDliaFinMonitoringaAnkety
	CatalogDogovoryRassrochki(Key: PrimaryCatalogDogovoryRassrochki!): CatalogDogovoryRassrochki
	CatalogDogovoryRassrochkis(): CatalogDogovoryRassrochki
	CatalogSertifikaty(Key: PrimaryCatalogSertifikaty!): CatalogSertifikaty
	CatalogSertifikatys(): CatalogSertifikaty
	DocumentPostuplenieDavalcheskogoMetalla(Key: PrimaryDocumentPostuplenieDavalcheskogoMetalla!): DocumentPostuplenieDavalcheskogoMetalla
	DocumentPostuplenieDavalcheskogoMetallas(): DocumentPostuplenieDavalcheskogoMetalla
	DocumentInkassovoePorucheniePeredannoe(Key: PrimaryDocumentInkassovoePorucheniePeredannoe!): DocumentInkassovoePorucheniePeredannoe
	DocumentInkassovoePorucheniePeredannoes(): DocumentInkassovoePorucheniePeredannoe
	DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha(Key: PrimaryDocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha!): DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha
	DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhas(): DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha
	DocumentInkassovoePorucheniePeredannoeRekvizityKontragenta(Key: PrimaryDocumentInkassovoePorucheniePeredannoeRekvizityKontragenta!): DocumentInkassovoePorucheniePeredannoeRekvizityKontragenta
	DocumentInkassovoePorucheniePeredannoeRekvizityKontragentas(): DocumentInkassovoePorucheniePeredannoeRekvizityKontragenta
	CatalogFormulyDliaRascheta(Key: PrimaryCatalogFormulyDliaRascheta!): CatalogFormulyDliaRascheta
	CatalogFormulyDliaRaschetas(): CatalogFormulyDliaRascheta
	CatalogKupony(Key: PrimaryCatalogKupony!): CatalogKupony
	CatalogKuponys(): CatalogKupony
	Correcting(Key: PrimaryCorrecting!): Correcting
	Correctings(): Correcting
	DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia(Key: PrimaryDocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia!): DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia
	DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniias(): DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia
	DocumentInternetZakaz(Key: PrimaryDocumentInternetZakaz!): DocumentInternetZakaz
	DocumentInternetZakazs(): DocumentInternetZakaz
	DocumentInternetZakazTovaryInternetZakaza(Key: PrimaryDocumentInternetZakazTovaryInternetZakaza!): DocumentInternetZakazTovaryInternetZakaza
	DocumentInternetZakazTovaryInternetZakazas(): DocumentInternetZakazTovaryInternetZakaza
	DocumentInternetZakazTovary(Key: PrimaryDocumentInternetZakazTovary!): DocumentInternetZakazTovary
	DocumentInternetZakazTovarys(): DocumentInternetZakazTovary
	CatalogRegiony(Key: PrimaryCatalogRegiony!): CatalogRegiony
	CatalogRegionys(): CatalogRegiony
	SaleJournal(Key: PrimarySaleJournal!): SaleJournal
	SaleJournals(): SaleJournal
	DocumentOtchetORoznichnykhProdazhakhBonusy(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhBonusy!): DocumentOtchetORoznichnykhProdazhakhBonusy
	DocumentOtchetORoznichnykhProdazhakhBonusys(): DocumentOtchetORoznichnykhProdazhakhBonusy
	DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami!): DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami
	DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamis(): DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami
	DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami!): DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami
	DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamis(): DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami
	DocumentOtchetORoznichnykhProdazhakhOplataSertifikatami(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhOplataSertifikatami!): DocumentOtchetORoznichnykhProdazhakhOplataSertifikatami
	DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamis(): DocumentOtchetORoznichnykhProdazhakhOplataSertifikatami
	DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov!): DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov
	DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovs(): DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov
	DocumentOtchetORoznichnykhProdazhakhTovary(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhTovary!): DocumentOtchetORoznichnykhProdazhakhTovary
	DocumentOtchetORoznichnykhProdazhakhTovarys(): DocumentOtchetORoznichnykhProdazhakhTovary
	DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha!): DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha
	DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhas(): DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha
	DocumentOtchetORoznichnykhProdazhakhDokumentyObmena(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhDokumentyObmena!): DocumentOtchetORoznichnykhProdazhakhDokumentyObmena
	DocumentOtchetORoznichnykhProdazhakhDokumentyObmenas(): DocumentOtchetORoznichnykhProdazhakhDokumentyObmena
	DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata!): DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata
	DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplatas(): DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata
	DocumentOtchetORoznichnykhProdazhakhOplataBallami(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhOplataBallami!): DocumentOtchetORoznichnykhProdazhakhOplataBallami
	DocumentOtchetORoznichnykhProdazhakhOplataBallamis(): DocumentOtchetORoznichnykhProdazhakhOplataBallami
	DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki!): DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki
	DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkis(): DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki
	DocumentOtchetORoznichnykhProdazhakhKupony(Key: PrimaryDocumentOtchetORoznichnykhProdazhakhKupony!): DocumentOtchetORoznichnykhProdazhakhKupony
	DocumentOtchetORoznichnykhProdazhakhKuponys(): DocumentOtchetORoznichnykhProdazhakhKupony
	DocumentOtmenaSkidokNomenklatury(Key: PrimaryDocumentOtmenaSkidokNomenklatury!): DocumentOtmenaSkidokNomenklatury
	DocumentOtmenaSkidokNomenklaturys(): DocumentOtmenaSkidokNomenklatury
	DocumentOtmenaSkidokNomenklaturyDokumenty(Key: PrimaryDocumentOtmenaSkidokNomenklaturyDokumenty!): DocumentOtmenaSkidokNomenklaturyDokumenty
	DocumentOtmenaSkidokNomenklaturyDokumentys(): DocumentOtmenaSkidokNomenklaturyDokumenty
	CatalogTovarnyeGruppy(Key: PrimaryCatalogTovarnyeGruppy!): CatalogTovarnyeGruppy
	CatalogTovarnyeGruppys(): CatalogTovarnyeGruppy
	DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv(Key: PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv!): DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv
	DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvs(): DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv
	DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha(Key: PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha!): DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha
	DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhas(): DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha
	DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta(Key: PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta!): DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta
	DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentas(): DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta
	CatalogOrderKey(Key: PrimaryCatalogOrderKey!): CatalogOrderKey
	CatalogOrderKeys(): CatalogOrderKey
	DocumentKassovyiChekKorrektsii(Key: PrimaryDocumentKassovyiChekKorrektsii!): DocumentKassovyiChekKorrektsii
	DocumentKassovyiChekKorrektsiis(): DocumentKassovyiChekKorrektsii
	DocumentKassovyiChekKorrektsiiOplata(Key: PrimaryDocumentKassovyiChekKorrektsiiOplata!): DocumentKassovyiChekKorrektsiiOplata
	DocumentKassovyiChekKorrektsiiOplatas(): DocumentKassovyiChekKorrektsiiOplata
	DocumentSchetNaOplatuPokupateliu(Key: PrimaryDocumentSchetNaOplatuPokupateliu!): DocumentSchetNaOplatuPokupateliu
	DocumentSchetNaOplatuPokupatelius(): DocumentSchetNaOplatuPokupateliu
	DocumentSchetNaOplatuPokupateliuTovary(Key: PrimaryDocumentSchetNaOplatuPokupateliuTovary!): DocumentSchetNaOplatuPokupateliuTovary
	DocumentSchetNaOplatuPokupateliuTovarys(): DocumentSchetNaOplatuPokupateliuTovary
	DocumentSchetNaOplatuPokupateliuUslugi(Key: PrimaryDocumentSchetNaOplatuPokupateliuUslugi!): DocumentSchetNaOplatuPokupateliuUslugi
	DocumentSchetNaOplatuPokupateliuUslugis(): DocumentSchetNaOplatuPokupateliuUslugi
	CatalogNastroikiObmenaDannymi(Key: PrimaryCatalogNastroikiObmenaDannymi!): CatalogNastroikiObmenaDannymi
	CatalogNastroikiObmenaDannymis(): CatalogNastroikiObmenaDannymi
	CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov(Key: PrimaryCatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov!): CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov
	CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovs(): CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov
	CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh(Key: PrimaryCatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh!): CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh
	CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhs(): CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh
	CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami(Key: PrimaryCatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami!): CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami
	CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamis(): CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami
	DocumentJournalBankovskieRaschetnyeDokumenty(Key: PrimaryDocumentJournalBankovskieRaschetnyeDokumenty!): DocumentJournalBankovskieRaschetnyeDokumenty
	DocumentJournalBankovskieRaschetnyeDokumentys(): DocumentJournalBankovskieRaschetnyeDokumenty
	DocumentZamenaDiskontnoiKarty(Key: PrimaryDocumentZamenaDiskontnoiKarty!): DocumentZamenaDiskontnoiKarty
	DocumentZamenaDiskontnoiKartys(): DocumentZamenaDiskontnoiKarty
	ReturnToSupplier(Key: PrimaryReturnToSupplier!): ReturnToSupplier
	ReturnToSuppliers(): ReturnToSupplier
	DocumentVozvratTovarovPostavshchikuTovary(Key: PrimaryDocumentVozvratTovarovPostavshchikuTovary!): DocumentVozvratTovarovPostavshchikuTovary
	DocumentVozvratTovarovPostavshchikuTovarys(): DocumentVozvratTovarovPostavshchikuTovary
	DocumentInventarizatsiiaTovarovNaSklade(Key: PrimaryDocumentInventarizatsiiaTovarovNaSklade!): DocumentInventarizatsiiaTovarovNaSklade
	DocumentInventarizatsiiaTovarovNaSklades(): DocumentInventarizatsiiaTovarovNaSklade
	DocumentInventarizatsiiaTovarovNaSkladeTovary(Key: PrimaryDocumentInventarizatsiiaTovarovNaSkladeTovary!): DocumentInventarizatsiiaTovarovNaSkladeTovary
	DocumentInventarizatsiiaTovarovNaSkladeTovarys(): DocumentInventarizatsiiaTovarovNaSkladeTovary
	DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii(Key: PrimaryDocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii!): DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii
	DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiis(): DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii
	DocumentInventarizatsiiaTovarovNaSkladeSertifikaty(Key: PrimaryDocumentInventarizatsiiaTovarovNaSkladeSertifikaty!): DocumentInventarizatsiiaTovarovNaSkladeSertifikaty
	DocumentInventarizatsiiaTovarovNaSkladeSertifikatys(): DocumentInventarizatsiiaTovarovNaSkladeSertifikaty
	DocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti(Key: PrimaryDocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti!): DocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti
	DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutis(): DocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti
	DocumentPrikhodnyiKassovyiOrder(Key: PrimaryDocumentPrikhodnyiKassovyiOrder!): DocumentPrikhodnyiKassovyiOrder
	DocumentPrikhodnyiKassovyiOrders(): DocumentPrikhodnyiKassovyiOrder
	DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha(Key: PrimaryDocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha!): DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha
	DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhas(): DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha
	DocumentPrikhodnyiKassovyiOrderOplata(Key: PrimaryDocumentPrikhodnyiKassovyiOrderOplata!): DocumentPrikhodnyiKassovyiOrderOplata
	DocumentPrikhodnyiKassovyiOrderOplatas(): DocumentPrikhodnyiKassovyiOrderOplata
	DocumentPrikhodnyiKassovyiOrderTovary(Key: PrimaryDocumentPrikhodnyiKassovyiOrderTovary!): DocumentPrikhodnyiKassovyiOrderTovary
	DocumentPrikhodnyiKassovyiOrderTovarys(): DocumentPrikhodnyiKassovyiOrderTovary
	CatalogPrichinyVozvrata(Key: PrimaryCatalogPrichinyVozvrata!): CatalogPrichinyVozvrata
	CatalogPrichinyVozvratas(): CatalogPrichinyVozvrata
	DocumentDenezhnyiChek(Key: PrimaryDocumentDenezhnyiChek!): DocumentDenezhnyiChek
	DocumentDenezhnyiCheks(): DocumentDenezhnyiChek
	DocumentVozvratMaterialovIzProizvodstva(Key: PrimaryDocumentVozvratMaterialovIzProizvodstva!): DocumentVozvratMaterialovIzProizvodstva
	DocumentVozvratMaterialovIzProizvodstvas(): DocumentVozvratMaterialovIzProizvodstva
	DocumentVozvratMaterialovIzProizvodstvaTovary(Key: PrimaryDocumentVozvratMaterialovIzProizvodstvaTovary!): DocumentVozvratMaterialovIzProizvodstvaTovary
	DocumentVozvratMaterialovIzProizvodstvaTovarys(): DocumentVozvratMaterialovIzProizvodstvaTovary
	DocumentPereotsenkaTovarovOtdannykhNaKomissiiu(Key: PrimaryDocumentPereotsenkaTovarovOtdannykhNaKomissiiu!): DocumentPereotsenkaTovarovOtdannykhNaKomissiiu
	DocumentPereotsenkaTovarovOtdannykhNaKomissiius(): DocumentPereotsenkaTovarovOtdannykhNaKomissiiu
	DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary(Key: PrimaryDocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary!): DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary
	DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovarys(): DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary
	DocumentVvodNachalnykhOstatkovPoRaskhodamUSN(Key: PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSN!): DocumentVvodNachalnykhOstatkovPoRaskhodamUSN
	DocumentVvodNachalnykhOstatkovPoRaskhodamUSNs(): DocumentVvodNachalnykhOstatkovPoRaskhodamUSN
	DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami(Key: PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami!): DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami
	DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamis(): DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami
	DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye(Key: PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye!): DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye
	DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyes(): DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye
	DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami(Key: PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami!): DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami
	DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamis(): DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami
	DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh(Key: PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh!): DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh
	DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhs(): DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh
	DocumentGTDImport(Key: PrimaryDocumentGTDImport!): DocumentGTDImport
	DocumentGTDImports(): DocumentGTDImport
	DocumentGTDImportRazdely(Key: PrimaryDocumentGTDImportRazdely!): DocumentGTDImportRazdely
	DocumentGTDImportRazdelys(): DocumentGTDImportRazdely
	DocumentGTDImportTovary(Key: PrimaryDocumentGTDImportTovary!): DocumentGTDImportTovary
	DocumentGTDImportTovarys(): DocumentGTDImportTovary
	DocumentAktSverki(Key: PrimaryDocumentAktSverki!): DocumentAktSverki
	DocumentAktSverkis(): DocumentAktSverki
	DocumentAktSverkiTovary(Key: PrimaryDocumentAktSverkiTovary!): DocumentAktSverkiTovary
	DocumentAktSverkiTovarys(): DocumentAktSverkiTovary
	CatalogFaily(Key: PrimaryCatalogFaily!): CatalogFaily
	CatalogFailys(): CatalogFaily
	CatalogFailyDopolnitelnyeRekvizity(Key: PrimaryCatalogFailyDopolnitelnyeRekvizity!): CatalogFailyDopolnitelnyeRekvizity
	CatalogFailyDopolnitelnyeRekvizitys(): CatalogFailyDopolnitelnyeRekvizity
	CatalogFailySertifikatyShifrovaniia(Key: PrimaryCatalogFailySertifikatyShifrovaniia!): CatalogFailySertifikatyShifrovaniia
	CatalogFailySertifikatyShifrovaniias(): CatalogFailySertifikatyShifrovaniia
	CatalogUchetnyeZapisiElektronnoiPochty(Key: PrimaryCatalogUchetnyeZapisiElektronnoiPochty!): CatalogUchetnyeZapisiElektronnoiPochty
	CatalogUchetnyeZapisiElektronnoiPochtys(): CatalogUchetnyeZapisiElektronnoiPochty
	CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi(Key: PrimaryCatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi!): CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi
	CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisis(): CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi
	DocumentPlaniruemoePostuplenieDenezhnykhSredstv(Key: PrimaryDocumentPlaniruemoePostuplenieDenezhnykhSredstv!): DocumentPlaniruemoePostuplenieDenezhnykhSredstv
	DocumentPlaniruemoePostuplenieDenezhnykhSredstvs(): DocumentPlaniruemoePostuplenieDenezhnykhSredstv
	DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha(Key: PrimaryDocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha!): DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha
	DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhas(): DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha
	DocumentPreiskurantTsenNaKamni(Key: PrimaryDocumentPreiskurantTsenNaKamni!): DocumentPreiskurantTsenNaKamni
	DocumentPreiskurantTsenNaKamnis(): DocumentPreiskurantTsenNaKamni
	Purchase(Key: PrimaryPurchase!): Purchase
	Purchases(): Purchase
	DocumentSkupkaTovarovTovary(Key: PrimaryDocumentSkupkaTovarovTovary!): DocumentSkupkaTovarovTovary
	DocumentSkupkaTovarovTovarys(): DocumentSkupkaTovarovTovary
	DocumentSchetFakturaPoluchennyi(Key: PrimaryDocumentSchetFakturaPoluchennyi!): DocumentSchetFakturaPoluchennyi
	DocumentSchetFakturaPoluchennyis(): DocumentSchetFakturaPoluchennyi
	DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam(Key: PrimaryDocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam!): DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam
	DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliams(): DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam
	DocumentAktKhimicheskogoAnalizaMetalla(Key: PrimaryDocumentAktKhimicheskogoAnalizaMetalla!): DocumentAktKhimicheskogoAnalizaMetalla
	DocumentAktKhimicheskogoAnalizaMetallas(): DocumentAktKhimicheskogoAnalizaMetalla
	CatalogfmKartochkaKontragenta(Key: PrimaryCatalogfmKartochkaKontragenta!): CatalogfmKartochkaKontragenta
	CatalogfmKartochkaKontragentas(): CatalogfmKartochkaKontragenta
	CatalogfmKartochkaKontragentaDannyeKontragenta(Key: PrimaryCatalogfmKartochkaKontragentaDannyeKontragenta!): CatalogfmKartochkaKontragentaDannyeKontragenta
	CatalogfmKartochkaKontragentaDannyeKontragentas(): CatalogfmKartochkaKontragentaDannyeKontragenta
	DocumentSpisanieProsrochennykhSertifikatov(Key: PrimaryDocumentSpisanieProsrochennykhSertifikatov!): DocumentSpisanieProsrochennykhSertifikatov
	DocumentSpisanieProsrochennykhSertifikatovs(): DocumentSpisanieProsrochennykhSertifikatov
	DocumentSpisanieProsrochennykhSertifikatovSertifikaty(Key: PrimaryDocumentSpisanieProsrochennykhSertifikatovSertifikaty!): DocumentSpisanieProsrochennykhSertifikatovSertifikaty
	DocumentSpisanieProsrochennykhSertifikatovSertifikatys(): DocumentSpisanieProsrochennykhSertifikatovSertifikaty
	DocumentZakrytieAvansovPoGrafikuPlatezhei(Key: PrimaryDocumentZakrytieAvansovPoGrafikuPlatezhei!): DocumentZakrytieAvansovPoGrafikuPlatezhei
	DocumentZakrytieAvansovPoGrafikuPlatezheis(): DocumentZakrytieAvansovPoGrafikuPlatezhei
	DocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty(Key: PrimaryDocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty!): DocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty
	DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentys(): DocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty
	CatalogTipySistemNalogooblozheniiaKKT(Key: PrimaryCatalogTipySistemNalogooblozheniiaKKT!): CatalogTipySistemNalogooblozheniiaKKT
	CatalogTipySistemNalogooblozheniiaKKTs(): CatalogTipySistemNalogooblozheniiaKKT
	DocumentAkkreditivPeredannyi(Key: PrimaryDocumentAkkreditivPeredannyi!): DocumentAkkreditivPeredannyi
	DocumentAkkreditivPeredannyis(): DocumentAkkreditivPeredannyi
	DocumentAkkreditivPeredannyiRasshifrovkaPlatezha(Key: PrimaryDocumentAkkreditivPeredannyiRasshifrovkaPlatezha!): DocumentAkkreditivPeredannyiRasshifrovkaPlatezha
	DocumentAkkreditivPeredannyiRasshifrovkaPlatezhas(): DocumentAkkreditivPeredannyiRasshifrovkaPlatezha
	DocumentAkkreditivPeredannyiRekvizityKontragenta(Key: PrimaryDocumentAkkreditivPeredannyiRekvizityKontragenta!): DocumentAkkreditivPeredannyiRekvizityKontragenta
	DocumentAkkreditivPeredannyiRekvizityKontragentas(): DocumentAkkreditivPeredannyiRekvizityKontragenta
	Supplier(Key: PrimarySupplier!): Supplier
	Suppliers(): Supplier
	CatalogKontragentyVidyDeiatelnosti(Key: PrimaryCatalogKontragentyVidyDeiatelnosti!): CatalogKontragentyVidyDeiatelnosti
	CatalogKontragentyVidyDeiatelnostis(): CatalogKontragentyVidyDeiatelnosti
	DocumentInformatsionnoeSoobshchenie(Key: PrimaryDocumentInformatsionnoeSoobshchenie!): DocumentInformatsionnoeSoobshchenie
	DocumentInformatsionnoeSoobshchenies(): DocumentInformatsionnoeSoobshchenie
	CatalogVlozheniiaElektronnykhPisem(Key: PrimaryCatalogVlozheniiaElektronnykhPisem!): CatalogVlozheniiaElektronnykhPisem
	CatalogVlozheniiaElektronnykhPisems(): CatalogVlozheniiaElektronnykhPisem
	DocumentPlatezhnoeTrebovanieVystavlennoe(Key: PrimaryDocumentPlatezhnoeTrebovanieVystavlennoe!): DocumentPlatezhnoeTrebovanieVystavlennoe
	DocumentPlatezhnoeTrebovanieVystavlennoes(): DocumentPlatezhnoeTrebovanieVystavlennoe
	DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha(Key: PrimaryDocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha!): DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha
	DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhas(): DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha
	DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta(Key: PrimaryDocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta!): DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta
	DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentas(): DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta
	DocumentMarketingovaiaAktsiia(Key: PrimaryDocumentMarketingovaiaAktsiia!): DocumentMarketingovaiaAktsiia
	DocumentMarketingovaiaAktsiias(): DocumentMarketingovaiaAktsiia
	DocumentMarketingovaiaAktsiiaMagaziny(Key: PrimaryDocumentMarketingovaiaAktsiiaMagaziny!): DocumentMarketingovaiaAktsiiaMagaziny
	DocumentMarketingovaiaAktsiiaMagazinys(): DocumentMarketingovaiaAktsiiaMagaziny
	DocumentMarketingovaiaAktsiiaSkidkiNatsenki(Key: PrimaryDocumentMarketingovaiaAktsiiaSkidkiNatsenki!): DocumentMarketingovaiaAktsiiaSkidkiNatsenki
	DocumentMarketingovaiaAktsiiaSkidkiNatsenkis(): DocumentMarketingovaiaAktsiiaSkidkiNatsenki
	DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa(Key: PrimaryDocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa!): DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa
	DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupas(): DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa
	CatalogStsenariiObmenovDannymi(Key: PrimaryCatalogStsenariiObmenovDannymi!): CatalogStsenariiObmenovDannymi
	CatalogStsenariiObmenovDannymis(): CatalogStsenariiObmenovDannymi
	CatalogStsenariiObmenovDannymiNastroikiObmena(Key: PrimaryCatalogStsenariiObmenovDannymiNastroikiObmena!): CatalogStsenariiObmenovDannymiNastroikiObmena
	CatalogStsenariiObmenovDannymiNastroikiObmenas(): CatalogStsenariiObmenovDannymiNastroikiObmena
	Item(Key: PrimaryItem!): Item
	Items(): Item
	CatalogNomenklaturaSostavLigatury(Key: PrimaryCatalogNomenklaturaSostavLigatury!): CatalogNomenklaturaSostavLigatury
	CatalogNomenklaturaSostavLigaturys(): CatalogNomenklaturaSostavLigatury
	DocumentOpros(Key: PrimaryDocumentOpros!): DocumentOpros
	DocumentOpross(): DocumentOpros
	DocumentOprosVoprosy(Key: PrimaryDocumentOprosVoprosy!): DocumentOprosVoprosy
	DocumentOprosVoprosys(): DocumentOprosVoprosy
	DocumentOprosSostavnoiOtvet(Key: PrimaryDocumentOprosSostavnoiOtvet!): DocumentOprosSostavnoiOtvet
	DocumentOprosSostavnoiOtvets(): DocumentOprosSostavnoiOtvet
	CatalogGruppyPoluchateleiSkidki(Key: PrimaryCatalogGruppyPoluchateleiSkidki!): CatalogGruppyPoluchateleiSkidki
	CatalogGruppyPoluchateleiSkidkis(): CatalogGruppyPoluchateleiSkidki
	Reassessment(Key: PrimaryReassessment!): Reassessment
	Reassessments(): Reassessment
	DocumentPereotsenkaTovarovVNTTTovary(Key: PrimaryDocumentPereotsenkaTovarovVNTTTovary!): DocumentPereotsenkaTovarovVNTTTovary
	DocumentPereotsenkaTovarovVNTTTovarys(): DocumentPereotsenkaTovarovVNTTTovary
	CatalogTomaKhraneniiaFailov(Key: PrimaryCatalogTomaKhraneniiaFailov!): CatalogTomaKhraneniiaFailov
	CatalogTomaKhraneniiaFailovs(): CatalogTomaKhraneniiaFailov
	DocumentJournalProizvodstvennyeDokumenty(Key: PrimaryDocumentJournalProizvodstvennyeDokumenty!): DocumentJournalProizvodstvennyeDokumenty
	DocumentJournalProizvodstvennyeDokumentys(): DocumentJournalProizvodstvennyeDokumenty
	DocumentIzmeneniePravDostupa(Key: PrimaryDocumentIzmeneniePravDostupa!): DocumentIzmeneniePravDostupa
	DocumentIzmeneniePravDostupas(): DocumentIzmeneniePravDostupa
	CatalogNastroikaAssortimentnoiMatritsy(Key: PrimaryCatalogNastroikaAssortimentnoiMatritsy!): CatalogNastroikaAssortimentnoiMatritsy
	CatalogNastroikaAssortimentnoiMatritsys(): CatalogNastroikaAssortimentnoiMatritsy
	CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp(Key: PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp!): CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp
	CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupps(): CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp
	CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii(Key: PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii!): CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii
	CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriis(): CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii
	CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii(Key: PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii!): CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii
	CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiis(): CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii
	DocumentJournalDokumentyKontragentov(Key: PrimaryDocumentJournalDokumentyKontragentov!): DocumentJournalDokumentyKontragentov
	DocumentJournalDokumentyKontragentovs(): DocumentJournalDokumentyKontragentov
	MoveInstance(Key: PrimaryMoveInstance!): MoveInstance
	MoveInstances(): MoveInstance
	DocumentPeremeshchenieTovarovSertifikaty(Key: PrimaryDocumentPeremeshchenieTovarovSertifikaty!): DocumentPeremeshchenieTovarovSertifikaty
	DocumentPeremeshchenieTovarovSertifikatys(): DocumentPeremeshchenieTovarovSertifikaty
	DocumentPeremeshchenieTovarovTovary(Key: PrimaryDocumentPeremeshchenieTovarovTovary!): DocumentPeremeshchenieTovarovTovary
	DocumentPeremeshchenieTovarovTovarys(): DocumentPeremeshchenieTovarovTovary
	DocumentPeremeshchenieTovarovSpisokZaiavok(Key: PrimaryDocumentPeremeshchenieTovarovSpisokZaiavok!): DocumentPeremeshchenieTovarovSpisokZaiavok
	DocumentPeremeshchenieTovarovSpisokZaiavoks(): DocumentPeremeshchenieTovarovSpisokZaiavok
	DocumentZakrytieZaiavokNaRaskhodovanieSredstv(Key: PrimaryDocumentZakrytieZaiavokNaRaskhodovanieSredstv!): DocumentZakrytieZaiavokNaRaskhodovanieSredstv
	DocumentZakrytieZaiavokNaRaskhodovanieSredstvs(): DocumentZakrytieZaiavokNaRaskhodovanieSredstv
	DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv(Key: PrimaryDocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv!): DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv
	DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvs(): DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv
	MemberCard(Key: PrimaryMemberCard!): MemberCard
	MemberCards(): MemberCard
	DocumentABCKlassifikatsiiaPokupatelei(Key: PrimaryDocumentABCKlassifikatsiiaPokupatelei!): DocumentABCKlassifikatsiiaPokupatelei
	DocumentABCKlassifikatsiiaPokupateleis(): DocumentABCKlassifikatsiiaPokupatelei
	DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov(Key: PrimaryDocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov!): DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov
	DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovs(): DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov
	CatalogIdentifikatoryObieektovMetadannykh(Key: PrimaryCatalogIdentifikatoryObieektovMetadannykh!): CatalogIdentifikatoryObieektovMetadannykh
	CatalogIdentifikatoryObieektovMetadannykhs(): CatalogIdentifikatoryObieektovMetadannykh
	DocumentSvodnaiaInventarizatsiiaTovarovNaSklade(Key: PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSklade!): DocumentSvodnaiaInventarizatsiiaTovarovNaSklade
	DocumentSvodnaiaInventarizatsiiaTovarovNaSklades(): DocumentSvodnaiaInventarizatsiiaTovarovNaSklade
	DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty(Key: PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty!): DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty
	DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatys(): DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty
	DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii(Key: PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii!): DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii
	DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiis(): DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii
	DocumentKorrektirovkaRealizatsii(Key: PrimaryDocumentKorrektirovkaRealizatsii!): DocumentKorrektirovkaRealizatsii
	DocumentKorrektirovkaRealizatsiis(): DocumentKorrektirovkaRealizatsii
	DocumentKorrektirovkaRealizatsiiTovary(Key: PrimaryDocumentKorrektirovkaRealizatsiiTovary!): DocumentKorrektirovkaRealizatsiiTovary
	DocumentKorrektirovkaRealizatsiiTovarys(): DocumentKorrektirovkaRealizatsiiTovary
	DocumentKorrektirovkaRealizatsiiUslugi(Key: PrimaryDocumentKorrektirovkaRealizatsiiUslugi!): DocumentKorrektirovkaRealizatsiiUslugi
	DocumentKorrektirovkaRealizatsiiUslugis(): DocumentKorrektirovkaRealizatsiiUslugi
	CatalogVidyDefektov(Key: PrimaryCatalogVidyDefektov!): CatalogVidyDefektov
	CatalogVidyDefektovs(): CatalogVidyDefektov
	DocumentDoverennost(Key: PrimaryDocumentDoverennost!): DocumentDoverennost
	DocumentDoverennosts(): DocumentDoverennost
	DocumentDoverennostTovary(Key: PrimaryDocumentDoverennostTovary!): DocumentDoverennostTovary
	DocumentDoverennostTovarys(): DocumentDoverennostTovary
	CatalogShablonyZapolneniiaKU(Key: PrimaryCatalogShablonyZapolneniiaKU!): CatalogShablonyZapolneniiaKU
	CatalogShablonyZapolneniiaKUs(): CatalogShablonyZapolneniiaKU
	CatalogShablonyZapolneniiaKUPrazdnichnyeDni(Key: PrimaryCatalogShablonyZapolneniiaKUPrazdnichnyeDni!): CatalogShablonyZapolneniiaKUPrazdnichnyeDni
	CatalogShablonyZapolneniiaKUPrazdnichnyeDnis(): CatalogShablonyZapolneniiaKUPrazdnichnyeDni
	CatalogShablonyZapolneniiaKUKUNaNedeliu(Key: PrimaryCatalogShablonyZapolneniiaKUKUNaNedeliu!): CatalogShablonyZapolneniiaKUKUNaNedeliu
	CatalogShablonyZapolneniiaKUKUNaNedelius(): CatalogShablonyZapolneniiaKUKUNaNedeliu
	CatalogShablonyZapolneniiaKUSalony(Key: PrimaryCatalogShablonyZapolneniiaKUSalony!): CatalogShablonyZapolneniiaKUSalony
	CatalogShablonyZapolneniiaKUSalonys(): CatalogShablonyZapolneniiaKUSalony
	DocumentPlanZapolneniiaVitrin(Key: PrimaryDocumentPlanZapolneniiaVitrin!): DocumentPlanZapolneniiaVitrin
	DocumentPlanZapolneniiaVitrins(): DocumentPlanZapolneniiaVitrin
	DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin(Key: PrimaryDocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin!): DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin
	DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrins(): DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin
	Instance(Key: PrimaryInstance!): Instance
	Instances(): Instance
	ReturnToManufacturing(Key: PrimaryReturnToManufacturing!): ReturnToManufacturing
	ReturnToManufacturings(): ReturnToManufacturing
	DocumentVozvratProduktsiiVProizvodstvoTovary(Key: PrimaryDocumentVozvratProduktsiiVProizvodstvoTovary!): DocumentVozvratProduktsiiVProizvodstvoTovary
	DocumentVozvratProduktsiiVProizvodstvoTovarys(): DocumentVozvratProduktsiiVProizvodstvoTovary
	CatalogNomeraGTD(Key: PrimaryCatalogNomeraGTD!): CatalogNomeraGTD
	CatalogNomeraGTDs(): CatalogNomeraGTD
	CatalogNastroikiRabochegoMestaPolzovatelia(Key: PrimaryCatalogNastroikiRabochegoMestaPolzovatelia!): CatalogNastroikiRabochegoMestaPolzovatelia
	CatalogNastroikiRabochegoMestaPolzovatelias(): CatalogNastroikiRabochegoMestaPolzovatelia
	CatalogNastroikiRabochegoMestaPolzovateliaNastroiki(Key: PrimaryCatalogNastroikiRabochegoMestaPolzovateliaNastroiki!): CatalogNastroikiRabochegoMestaPolzovateliaNastroiki
	CatalogNastroikiRabochegoMestaPolzovateliaNastroikis(): CatalogNastroikiRabochegoMestaPolzovateliaNastroiki
	CatalogsmsShablony(Key: PrimaryCatalogsmsShablony!): CatalogsmsShablony
	CatalogsmsShablonys(): CatalogsmsShablony
	WriteOff(Key: PrimaryWriteOff!): WriteOff
	WriteOffs(): WriteOff
	DocumentSpisanieTovarovTovary(Key: PrimaryDocumentSpisanieTovarovTovary!): DocumentSpisanieTovarovTovary
	DocumentSpisanieTovarovTovarys(): DocumentSpisanieTovarovTovary
	DocumentSpisanieTovarovSertifikaty(Key: PrimaryDocumentSpisanieTovarovSertifikaty!): DocumentSpisanieTovarovSertifikaty
	DocumentSpisanieTovarovSertifikatys(): DocumentSpisanieTovarovSertifikaty
	DocumentsmsSoobshchenie(Key: PrimaryDocumentsmsSoobshchenie!): DocumentsmsSoobshchenie
	DocumentsmsSoobshchenies(): DocumentsmsSoobshchenie
	DocumentsmsSoobshcheniePoluchateli(Key: PrimaryDocumentsmsSoobshcheniePoluchateli!): DocumentsmsSoobshcheniePoluchateli
	DocumentsmsSoobshcheniePoluchatelis(): DocumentsmsSoobshcheniePoluchateli
	DocumentOplataOtPokupateliaPlatezhnoiKartoi(Key: PrimaryDocumentOplataOtPokupateliaPlatezhnoiKartoi!): DocumentOplataOtPokupateliaPlatezhnoiKartoi
	DocumentOplataOtPokupateliaPlatezhnoiKartois(): DocumentOplataOtPokupateliaPlatezhnoiKartoi
	CatalogDragotsennyeKamni(Key: PrimaryCatalogDragotsennyeKamni!): CatalogDragotsennyeKamni
	CatalogDragotsennyeKamnis(): CatalogDragotsennyeKamni
	CatalogKalendariPlanirovaniiaProdazh(Key: PrimaryCatalogKalendariPlanirovaniiaProdazh!): CatalogKalendariPlanirovaniiaProdazh
	CatalogKalendariPlanirovaniiaProdazhs(): CatalogKalendariPlanirovaniiaProdazh
	CatalogKalendariPlanirovaniiaProdazhKUPoDniam(Key: PrimaryCatalogKalendariPlanirovaniiaProdazhKUPoDniam!): CatalogKalendariPlanirovaniiaProdazhKUPoDniam
	CatalogKalendariPlanirovaniiaProdazhKUPoDniams(): CatalogKalendariPlanirovaniiaProdazhKUPoDniam
	CatalogKalendariPlanirovaniiaProdazhSalony(Key: PrimaryCatalogKalendariPlanirovaniiaProdazhSalony!): CatalogKalendariPlanirovaniiaProdazhSalony
	CatalogKalendariPlanirovaniiaProdazhSalonys(): CatalogKalendariPlanirovaniiaProdazhSalony
	CatalogKontaktnyeLitsa(Key: PrimaryCatalogKontaktnyeLitsa!): CatalogKontaktnyeLitsa
	CatalogKontaktnyeLitsas(): CatalogKontaktnyeLitsa
	CatalogFizicheskieLitsa(Key: PrimaryCatalogFizicheskieLitsa!): CatalogFizicheskieLitsa
	CatalogFizicheskieLitsas(): CatalogFizicheskieLitsa
	CatalogTipovyeAnkety(Key: PrimaryCatalogTipovyeAnkety!): CatalogTipovyeAnkety
	CatalogTipovyeAnketys(): CatalogTipovyeAnkety
	CatalogTipovyeAnketyVoprosyAnkety(Key: PrimaryCatalogTipovyeAnketyVoprosyAnkety!): CatalogTipovyeAnketyVoprosyAnkety
	CatalogTipovyeAnketyVoprosyAnketys(): CatalogTipovyeAnketyVoprosyAnkety
	DocumentNachislenieSpisanieBonusov(Key: PrimaryDocumentNachislenieSpisanieBonusov!): DocumentNachislenieSpisanieBonusov
	DocumentNachislenieSpisanieBonusovs(): DocumentNachislenieSpisanieBonusov
	DocumentNachislenieSpisanieBonusovDiskontnyeKarty(Key: PrimaryDocumentNachislenieSpisanieBonusovDiskontnyeKarty!): DocumentNachislenieSpisanieBonusovDiskontnyeKarty
	DocumentNachislenieSpisanieBonusovDiskontnyeKartys(): DocumentNachislenieSpisanieBonusovDiskontnyeKarty
	Type(Key: PrimaryType!): Type
	Types(): Type
	CatalogfmKodyVidovDokumentov(Key: PrimaryCatalogfmKodyVidovDokumentov!): CatalogfmKodyVidovDokumentov
	CatalogfmKodyVidovDokumentovs(): CatalogfmKodyVidovDokumentov
	DocumentPlatezhnoeTrebovaniePoluchennoe(Key: PrimaryDocumentPlatezhnoeTrebovaniePoluchennoe!): DocumentPlatezhnoeTrebovaniePoluchennoe
	DocumentPlatezhnoeTrebovaniePoluchennoes(): DocumentPlatezhnoeTrebovaniePoluchennoe
	DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha(Key: PrimaryDocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha!): DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha
	DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhas(): DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha
	DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta(Key: PrimaryDocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta!): DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta
	DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentas(): DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta
	DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv(Key: PrimaryDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv!): DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv
	DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvs(): DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv
	DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS(Key: PrimaryDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS!): DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS
	DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSs(): DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS
	CatalogRazdelyAnkety(Key: PrimaryCatalogRazdelyAnkety!): CatalogRazdelyAnkety
	CatalogRazdelyAnketys(): CatalogRazdelyAnkety
	DocumentOtchetPoFinMonitoringu(Key: PrimaryDocumentOtchetPoFinMonitoringu!): DocumentOtchetPoFinMonitoringu
	DocumentOtchetPoFinMonitoringus(): DocumentOtchetPoFinMonitoringu
	DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa(Key: PrimaryDocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa!): DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa
	DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringas(): DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa
	DocumentOtchetPoFinMonitoringuDannyeDokumenta(Key: PrimaryDocumentOtchetPoFinMonitoringuDannyeDokumenta!): DocumentOtchetPoFinMonitoringuDannyeDokumenta
	DocumentOtchetPoFinMonitoringuDannyeDokumentas(): DocumentOtchetPoFinMonitoringuDannyeDokumenta
	CatalogKliuchiAnalitikiUchetaNomenklatury(Key: PrimaryCatalogKliuchiAnalitikiUchetaNomenklatury!): CatalogKliuchiAnalitikiUchetaNomenklatury
	CatalogKliuchiAnalitikiUchetaNomenklaturys(): CatalogKliuchiAnalitikiUchetaNomenklatury
	CatalogVersiiFailov(Key: PrimaryCatalogVersiiFailov!): CatalogVersiiFailov
	CatalogVersiiFailovs(): CatalogVersiiFailov
	CatalogVersiiFailovElektronnyePodpisi(Key: PrimaryCatalogVersiiFailovElektronnyePodpisi!): CatalogVersiiFailovElektronnyePodpisi
	CatalogVersiiFailovElektronnyePodpisis(): CatalogVersiiFailovElektronnyePodpisi
	DocumentUstanovkaTsenNomenklatury(Key: PrimaryDocumentUstanovkaTsenNomenklatury!): DocumentUstanovkaTsenNomenklatury
	DocumentUstanovkaTsenNomenklaturys(): DocumentUstanovkaTsenNomenklatury
	DocumentUstanovkaTsenNomenklaturyTipyTsen(Key: PrimaryDocumentUstanovkaTsenNomenklaturyTipyTsen!): DocumentUstanovkaTsenNomenklaturyTipyTsen
	DocumentUstanovkaTsenNomenklaturyTipyTsens(): DocumentUstanovkaTsenNomenklaturyTipyTsen
	DocumentUstanovkaTsenNomenklaturyTovary(Key: PrimaryDocumentUstanovkaTsenNomenklaturyTovary!): DocumentUstanovkaTsenNomenklaturyTovary
	DocumentUstanovkaTsenNomenklaturyTovarys(): DocumentUstanovkaTsenNomenklaturyTovary
	DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv(Key: PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv!): DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv
	DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvs(): DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv
	DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha(Key: PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha!): DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha
	DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhas(): DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha
	DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta(Key: PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta!): DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta
	DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentas(): DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta
	DocumentPreiskurantNaSkupku(Key: PrimaryDocumentPreiskurantNaSkupku!): DocumentPreiskurantNaSkupku
	DocumentPreiskurantNaSkupkus(): DocumentPreiskurantNaSkupku
	DocumentPreiskurantNaSkupkuProby(Key: PrimaryDocumentPreiskurantNaSkupkuProby!): DocumentPreiskurantNaSkupkuProby
	DocumentPreiskurantNaSkupkuProbys(): DocumentPreiskurantNaSkupkuProby
	DocumentPeredachaMaterialovVProizvodstvo(Key: PrimaryDocumentPeredachaMaterialovVProizvodstvo!): DocumentPeredachaMaterialovVProizvodstvo
	DocumentPeredachaMaterialovVProizvodstvos(): DocumentPeredachaMaterialovVProizvodstvo
	DocumentPeredachaMaterialovVProizvodstvoTovary(Key: PrimaryDocumentPeredachaMaterialovVProizvodstvoTovary!): DocumentPeredachaMaterialovVProizvodstvoTovary
	DocumentPeredachaMaterialovVProizvodstvoTovarys(): DocumentPeredachaMaterialovVProizvodstvoTovary
	DocumentVnutrenniiZakaz(Key: PrimaryDocumentVnutrenniiZakaz!): DocumentVnutrenniiZakaz
	DocumentVnutrenniiZakazs(): DocumentVnutrenniiZakaz
	DocumentVnutrenniiZakazTovary(Key: PrimaryDocumentVnutrenniiZakazTovary!): DocumentVnutrenniiZakazTovary
	DocumentVnutrenniiZakazTovarys(): DocumentVnutrenniiZakazTovary
	CatalogKhranilishcheDopolnitelnoiInformatsii(Key: PrimaryCatalogKhranilishcheDopolnitelnoiInformatsii!): CatalogKhranilishcheDopolnitelnoiInformatsii
	CatalogKhranilishcheDopolnitelnoiInformatsiis(): CatalogKhranilishcheDopolnitelnoiInformatsii
	CatalogDopolnitelnyeVneshnieObrabotki(Key: PrimaryCatalogDopolnitelnyeVneshnieObrabotki!): CatalogDopolnitelnyeVneshnieObrabotki
	CatalogDopolnitelnyeVneshnieObrabotkis(): CatalogDopolnitelnyeVneshnieObrabotki
	CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost(Key: PrimaryCatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost!): CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost
	CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnosts(): CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost
	CatalogDopolnitelnyeVneshnieObrabotkiKomandy(Key: PrimaryCatalogDopolnitelnyeVneshnieObrabotkiKomandy!): CatalogDopolnitelnyeVneshnieObrabotkiKomandy
	CatalogDopolnitelnyeVneshnieObrabotkiKomandys(): CatalogDopolnitelnyeVneshnieObrabotkiKomandy
	CatalogDopolnitelnyeVneshnieObrabotkiRazdely(Key: PrimaryCatalogDopolnitelnyeVneshnieObrabotkiRazdely!): CatalogDopolnitelnyeVneshnieObrabotkiRazdely
	CatalogDopolnitelnyeVneshnieObrabotkiRazdelys(): CatalogDopolnitelnyeVneshnieObrabotkiRazdely
	CatalogDopolnitelnyeVneshnieObrabotkiNaznachenie(Key: PrimaryCatalogDopolnitelnyeVneshnieObrabotkiNaznachenie!): CatalogDopolnitelnyeVneshnieObrabotkiNaznachenie
	CatalogDopolnitelnyeVneshnieObrabotkiNaznachenies(): CatalogDopolnitelnyeVneshnieObrabotkiNaznachenie
	CatalogDopolnitelnyeVneshnieObrabotkiRazresheniia(Key: PrimaryCatalogDopolnitelnyeVneshnieObrabotkiRazresheniia!): CatalogDopolnitelnyeVneshnieObrabotkiRazresheniia
	CatalogDopolnitelnyeVneshnieObrabotkiRazresheniias(): CatalogDopolnitelnyeVneshnieObrabotkiRazresheniia
	CatalogGruppyPolzovatelei(Key: PrimaryCatalogGruppyPolzovatelei!): CatalogGruppyPolzovatelei
	CatalogGruppyPolzovateleis(): CatalogGruppyPolzovatelei
	CatalogGruppyPolzovateleiPolzovateliGruppy(Key: PrimaryCatalogGruppyPolzovateleiPolzovateliGruppy!): CatalogGruppyPolzovateleiPolzovateliGruppy
	CatalogGruppyPolzovateleiPolzovateliGruppys(): CatalogGruppyPolzovateleiPolzovateliGruppy
	DocumentJournalZakazyKlientov(Key: PrimaryDocumentJournalZakazyKlientov!): DocumentJournalZakazyKlientov
	DocumentJournalZakazyKlientovs(): DocumentJournalZakazyKlientov
	DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki(Key: PrimaryDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki!): DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki
	DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkis(): DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki
	DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary(Key: PrimaryDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary!): DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary
	DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovarys(): DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary
	DocumentZaiavkaNaPeremeshchenieTovarov(Key: PrimaryDocumentZaiavkaNaPeremeshchenieTovarov!): DocumentZaiavkaNaPeremeshchenieTovarov
	DocumentZaiavkaNaPeremeshchenieTovarovs(): DocumentZaiavkaNaPeremeshchenieTovarov
	DocumentZaiavkaNaPeremeshchenieTovarovTovary(Key: PrimaryDocumentZaiavkaNaPeremeshchenieTovarovTovary!): DocumentZaiavkaNaPeremeshchenieTovarovTovary
	DocumentZaiavkaNaPeremeshchenieTovarovTovarys(): DocumentZaiavkaNaPeremeshchenieTovarovTovary
	CatalogUsloviiaProdazh(Key: PrimaryCatalogUsloviiaProdazh!): CatalogUsloviiaProdazh
	CatalogUsloviiaProdazhs(): CatalogUsloviiaProdazh
	DocumentVvodNachalnykhOstatkovPoFinMonitoringu(Key: PrimaryDocumentVvodNachalnykhOstatkovPoFinMonitoringu!): DocumentVvodNachalnykhOstatkovPoFinMonitoringu
	DocumentVvodNachalnykhOstatkovPoFinMonitoringus(): DocumentVvodNachalnykhOstatkovPoFinMonitoringu
	DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora(Key: PrimaryDocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora!): DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora
	DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoras(): DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora
	CatalogOrganizatsii(Key: PrimaryCatalogOrganizatsii!): CatalogOrganizatsii
	CatalogOrganizatsiis(): CatalogOrganizatsii
	CatalogUsloviiaOplaty(Key: PrimaryCatalogUsloviiaOplaty!): CatalogUsloviiaOplaty
	CatalogUsloviiaOplatys(): CatalogUsloviiaOplaty
	CatalogUsloviiaOplatyTablitsaVyplat(Key: PrimaryCatalogUsloviiaOplatyTablitsaVyplat!): CatalogUsloviiaOplatyTablitsaVyplat
	CatalogUsloviiaOplatyTablitsaVyplats(): CatalogUsloviiaOplatyTablitsaVyplat
	CatalogKategoriiObieektov(Key: PrimaryCatalogKategoriiObieektov!): CatalogKategoriiObieektov
	CatalogKategoriiObieektovs(): CatalogKategoriiObieektov
	CatalogfmZnacheniiaDliaZapolneniia(Key: PrimaryCatalogfmZnacheniiaDliaZapolneniia!): CatalogfmZnacheniiaDliaZapolneniia
	CatalogfmZnacheniiaDliaZapolneniias(): CatalogfmZnacheniiaDliaZapolneniia
	DocumentUdalitNariadZakaz(Key: PrimaryDocumentUdalitNariadZakaz!): DocumentUdalitNariadZakaz
	DocumentUdalitNariadZakazs(): DocumentUdalitNariadZakaz
	DocumentUdalitNariadZakazIzdeliia(Key: PrimaryDocumentUdalitNariadZakazIzdeliia!): DocumentUdalitNariadZakazIzdeliia
	DocumentUdalitNariadZakazIzdeliias(): DocumentUdalitNariadZakazIzdeliia
	DocumentUdalitNariadZakazUslugi(Key: PrimaryDocumentUdalitNariadZakazUslugi!): DocumentUdalitNariadZakazUslugi
	DocumentUdalitNariadZakazUslugis(): DocumentUdalitNariadZakazUslugi
	DocumentUdalitNariadZakazSpetsifikatsiia(Key: PrimaryDocumentUdalitNariadZakazSpetsifikatsiia!): DocumentUdalitNariadZakazSpetsifikatsiia
	DocumentUdalitNariadZakazSpetsifikatsiias(): DocumentUdalitNariadZakazSpetsifikatsiia
	DocumentUdalitNariadZakazMetally(Key: PrimaryDocumentUdalitNariadZakazMetally!): DocumentUdalitNariadZakazMetally
	DocumentUdalitNariadZakazMetallys(): DocumentUdalitNariadZakazMetally
	DocumentUdalitNariadZakazVstavki(Key: PrimaryDocumentUdalitNariadZakazVstavki!): DocumentUdalitNariadZakazVstavki
	DocumentUdalitNariadZakazVstavkis(): DocumentUdalitNariadZakazVstavki
	CatalogBanki(Key: PrimaryCatalogBanki!): CatalogBanki
	CatalogBankis(): CatalogBanki
	CatalogRoliKontaktnykhLits(Key: PrimaryCatalogRoliKontaktnykhLits!): CatalogRoliKontaktnykhLits
	CatalogRoliKontaktnykhLitss(): CatalogRoliKontaktnykhLits
	DocumentRestrukturizatsiiaDolga(Key: PrimaryDocumentRestrukturizatsiiaDolga!): DocumentRestrukturizatsiiaDolga
	DocumentRestrukturizatsiiaDolgas(): DocumentRestrukturizatsiiaDolga
	DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti(Key: PrimaryDocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti!): DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti
	DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostis(): DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti
	DocumentAkkreditivPoluchennyi(Key: PrimaryDocumentAkkreditivPoluchennyi!): DocumentAkkreditivPoluchennyi
	DocumentAkkreditivPoluchennyis(): DocumentAkkreditivPoluchennyi
	DocumentAkkreditivPoluchennyiRasshifrovkaPlatezha(Key: PrimaryDocumentAkkreditivPoluchennyiRasshifrovkaPlatezha!): DocumentAkkreditivPoluchennyiRasshifrovkaPlatezha
	DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhas(): DocumentAkkreditivPoluchennyiRasshifrovkaPlatezha
	DocumentAkkreditivPoluchennyiRekvizityKontragenta(Key: PrimaryDocumentAkkreditivPoluchennyiRekvizityKontragenta!): DocumentAkkreditivPoluchennyiRekvizityKontragenta
	DocumentAkkreditivPoluchennyiRekvizityKontragentas(): DocumentAkkreditivPoluchennyiRekvizityKontragenta
	DocumentPriemIzRemonta(Key: PrimaryDocumentPriemIzRemonta!): DocumentPriemIzRemonta
	DocumentPriemIzRemontas(): DocumentPriemIzRemonta
	DocumentPriemIzRemontaIzdeliia(Key: PrimaryDocumentPriemIzRemontaIzdeliia!): DocumentPriemIzRemontaIzdeliia
	DocumentPriemIzRemontaIzdeliias(): DocumentPriemIzRemontaIzdeliia
	DocumentPriemIzRemontaMaterialy(Key: PrimaryDocumentPriemIzRemontaMaterialy!): DocumentPriemIzRemontaMaterialy
	DocumentPriemIzRemontaMaterialys(): DocumentPriemIzRemontaMaterialy
	CatalogTsveta(Key: PrimaryCatalogTsveta!): CatalogTsveta
	CatalogTsvetas(): CatalogTsveta
	DocumentStornirovanieOtchetaKomissioneraOProdazhakh(Key: PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakh!): DocumentStornirovanieOtchetaKomissioneraOProdazhakh
	DocumentStornirovanieOtchetaKomissioneraOProdazhakhs(): DocumentStornirovanieOtchetaKomissioneraOProdazhakh
	DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva(Key: PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva!): DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva
	DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvas(): DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva
	DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary(Key: PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary!): DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary
	DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovarys(): DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary
	DocumentJournalDavalcheskieDokumenty(Key: PrimaryDocumentJournalDavalcheskieDokumenty!): DocumentJournalDavalcheskieDokumenty
	DocumentJournalDavalcheskieDokumentys(): DocumentJournalDavalcheskieDokumenty
	CatalogfmAnketaKlienta(Key: PrimaryCatalogfmAnketaKlienta!): CatalogfmAnketaKlienta
	CatalogfmAnketaKlientas(): CatalogfmAnketaKlienta
	CatalogfmAnketaKlientaDannyeKontragenta(Key: PrimaryCatalogfmAnketaKlientaDannyeKontragenta!): CatalogfmAnketaKlientaDannyeKontragenta
	CatalogfmAnketaKlientaDannyeKontragentas(): CatalogfmAnketaKlientaDannyeKontragenta
	CatalogVidyVzaimoraschetov(Key: PrimaryCatalogVidyVzaimoraschetov!): CatalogVidyVzaimoraschetov
	CatalogVidyVzaimoraschetovs(): CatalogVidyVzaimoraschetov
	DocumentUstanovkaZnacheniiTochkiZakaza(Key: PrimaryDocumentUstanovkaZnacheniiTochkiZakaza!): DocumentUstanovkaZnacheniiTochkiZakaza
	DocumentUstanovkaZnacheniiTochkiZakazas(): DocumentUstanovkaZnacheniiTochkiZakaza
	DocumentUstanovkaZnacheniiTochkiZakazaTovary(Key: PrimaryDocumentUstanovkaZnacheniiTochkiZakazaTovary!): DocumentUstanovkaZnacheniiTochkiZakazaTovary
	DocumentUstanovkaZnacheniiTochkiZakazaTovarys(): DocumentUstanovkaZnacheniiTochkiZakazaTovary
	CatalogZnacheniiaKodirovki(Key: PrimaryCatalogZnacheniiaKodirovki!): CatalogZnacheniiaKodirovki
	CatalogZnacheniiaKodirovkis(): CatalogZnacheniiaKodirovki
	CatalogPravilaProdazh(Key: PrimaryCatalogPravilaProdazh!): CatalogPravilaProdazh
	CatalogPravilaProdazhs(): CatalogPravilaProdazh
	CatalogPravilaProdazhTovary(Key: PrimaryCatalogPravilaProdazhTovary!): CatalogPravilaProdazhTovary
	CatalogPravilaProdazhTovarys(): CatalogPravilaProdazhTovary
	DocumentPostuplenieDopRaskhodov(Key: PrimaryDocumentPostuplenieDopRaskhodov!): DocumentPostuplenieDopRaskhodov
	DocumentPostuplenieDopRaskhodovs(): DocumentPostuplenieDopRaskhodov
	DocumentPostuplenieDopRaskhodovTovary(Key: PrimaryDocumentPostuplenieDopRaskhodovTovary!): DocumentPostuplenieDopRaskhodovTovary
	DocumentPostuplenieDopRaskhodovTovarys(): DocumentPostuplenieDopRaskhodovTovary
	CatalogKhoziaistvennyeOperatsii(Key: PrimaryCatalogKhoziaistvennyeOperatsii!): CatalogKhoziaistvennyeOperatsii
	CatalogKhoziaistvennyeOperatsiis(): CatalogKhoziaistvennyeOperatsii
	DocumentAvansovyiOtchet(Key: PrimaryDocumentAvansovyiOtchet!): DocumentAvansovyiOtchet
	DocumentAvansovyiOtchets(): DocumentAvansovyiOtchet
	DocumentAvansovyiOtchetVydannyeAvansy(Key: PrimaryDocumentAvansovyiOtchetVydannyeAvansy!): DocumentAvansovyiOtchetVydannyeAvansy
	DocumentAvansovyiOtchetVydannyeAvansys(): DocumentAvansovyiOtchetVydannyeAvansy
	DocumentAvansovyiOtchetTovary(Key: PrimaryDocumentAvansovyiOtchetTovary!): DocumentAvansovyiOtchetTovary
	DocumentAvansovyiOtchetTovarys(): DocumentAvansovyiOtchetTovary
	DocumentAvansovyiOtchetOplataPostavshchikam(Key: PrimaryDocumentAvansovyiOtchetOplataPostavshchikam!): DocumentAvansovyiOtchetOplataPostavshchikam
	DocumentAvansovyiOtchetOplataPostavshchikams(): DocumentAvansovyiOtchetOplataPostavshchikam
	DocumentAvansovyiOtchetProchee(Key: PrimaryDocumentAvansovyiOtchetProchee!): DocumentAvansovyiOtchetProchee
	DocumentAvansovyiOtchetProchees(): DocumentAvansovyiOtchetProchee
	CatalogDolzhnostiOrganizatsii(Key: PrimaryCatalogDolzhnostiOrganizatsii!): CatalogDolzhnostiOrganizatsii
	CatalogDolzhnostiOrganizatsiis(): CatalogDolzhnostiOrganizatsii
	CatalogAnalitikaTipaIzdeliia(Key: PrimaryCatalogAnalitikaTipaIzdeliia!): CatalogAnalitikaTipaIzdeliia
	CatalogAnalitikaTipaIzdeliias(): CatalogAnalitikaTipaIzdeliia
	CatalogDopolnitelnyePechatnyeFormy(Key: PrimaryCatalogDopolnitelnyePechatnyeFormy!): CatalogDopolnitelnyePechatnyeFormy
	CatalogDopolnitelnyePechatnyeFormys(): CatalogDopolnitelnyePechatnyeFormy
	CatalogDopolnitelnyePechatnyeFormyPrinadlezhnost(Key: PrimaryCatalogDopolnitelnyePechatnyeFormyPrinadlezhnost!): CatalogDopolnitelnyePechatnyeFormyPrinadlezhnost
	CatalogDopolnitelnyePechatnyeFormyPrinadlezhnosts(): CatalogDopolnitelnyePechatnyeFormyPrinadlezhnost
	MemberCardsType(Key: PrimaryMemberCardsType!): MemberCardsType
	MemberCardsTypes(): MemberCardsType
	DocumentRegistratsiiaNaSaite(Key: PrimaryDocumentRegistratsiiaNaSaite!): DocumentRegistratsiiaNaSaite
	DocumentRegistratsiiaNaSaites(): DocumentRegistratsiiaNaSaite
	CatalogObrabotkiObsluzhivaniiaTO(Key: PrimaryCatalogObrabotkiObsluzhivaniiaTO!): CatalogObrabotkiObsluzhivaniiaTO
	CatalogObrabotkiObsluzhivaniiaTOs(): CatalogObrabotkiObsluzhivaniiaTO
	CatalogObrabotkiObsluzhivaniiaTOModeli(Key: PrimaryCatalogObrabotkiObsluzhivaniiaTOModeli!): CatalogObrabotkiObsluzhivaniiaTOModeli
	CatalogObrabotkiObsluzhivaniiaTOModelis(): CatalogObrabotkiObsluzhivaniiaTOModeli
	CatalogNastroikaIntervalov(Key: PrimaryCatalogNastroikaIntervalov!): CatalogNastroikaIntervalov
	CatalogNastroikaIntervalovs(): CatalogNastroikaIntervalov
	CatalogNastroikaIntervalovTablichnaiaChast(Key: PrimaryCatalogNastroikaIntervalovTablichnaiaChast!): CatalogNastroikaIntervalovTablichnaiaChast
	CatalogNastroikaIntervalovTablichnaiaChasts(): CatalogNastroikaIntervalovTablichnaiaChast
	CatalogProfiliGruppDostupa(Key: PrimaryCatalogProfiliGruppDostupa!): CatalogProfiliGruppDostupa
	CatalogProfiliGruppDostupas(): CatalogProfiliGruppDostupa
	CatalogProfiliGruppDostupaRoli(Key: PrimaryCatalogProfiliGruppDostupaRoli!): CatalogProfiliGruppDostupaRoli
	CatalogProfiliGruppDostupaRolis(): CatalogProfiliGruppDostupaRoli
	CatalogProfiliGruppDostupaVidyDostupa(Key: PrimaryCatalogProfiliGruppDostupaVidyDostupa!): CatalogProfiliGruppDostupaVidyDostupa
	CatalogProfiliGruppDostupaVidyDostupas(): CatalogProfiliGruppDostupaVidyDostupa
	CatalogProfiliGruppDostupaZnacheniiaDostupa(Key: PrimaryCatalogProfiliGruppDostupaZnacheniiaDostupa!): CatalogProfiliGruppDostupaZnacheniiaDostupa
	CatalogProfiliGruppDostupaZnacheniiaDostupas(): CatalogProfiliGruppDostupaZnacheniiaDostupa
	CatalogProfiliGruppDostupaDostupPoPodsistemam(Key: PrimaryCatalogProfiliGruppDostupaDostupPoPodsistemam!): CatalogProfiliGruppDostupaDostupPoPodsistemam
	CatalogProfiliGruppDostupaDostupPoPodsistemams(): CatalogProfiliGruppDostupaDostupPoPodsistemam
	CatalogNastroikiDliaKurera(Key: PrimaryCatalogNastroikiDliaKurera!): CatalogNastroikiDliaKurera
	CatalogNastroikiDliaKureras(): CatalogNastroikiDliaKurera
	CatalogNastroikiDliaKureraSostavNaimenovaniia(Key: PrimaryCatalogNastroikiDliaKureraSostavNaimenovaniia!): CatalogNastroikiDliaKureraSostavNaimenovaniia
	CatalogNastroikiDliaKureraSostavNaimenovaniias(): CatalogNastroikiDliaKureraSostavNaimenovaniia
	CatalogTipyTsenNomenklaturyKontragentov(Key: PrimaryCatalogTipyTsenNomenklaturyKontragentov!): CatalogTipyTsenNomenklaturyKontragentov
	CatalogTipyTsenNomenklaturyKontragentovs(): CatalogTipyTsenNomenklaturyKontragentov
	DocumentJournalTsenoobrazovanie(Key: PrimaryDocumentJournalTsenoobrazovanie!): DocumentJournalTsenoobrazovanie
	DocumentJournalTsenoobrazovanies(): DocumentJournalTsenoobrazovanie
	CatalogEdinitsyIzmereniia(Key: PrimaryCatalogEdinitsyIzmereniia!): CatalogEdinitsyIzmereniia
	CatalogEdinitsyIzmereniias(): CatalogEdinitsyIzmereniia
	CatalogStatiDvizheniiaDenezhnykhSredstv(Key: PrimaryCatalogStatiDvizheniiaDenezhnykhSredstv!): CatalogStatiDvizheniiaDenezhnykhSredstv
	CatalogStatiDvizheniiaDenezhnykhSredstvs(): CatalogStatiDvizheniiaDenezhnykhSredstv
	DocumentInkassovoePorucheniePoluchennoe(Key: PrimaryDocumentInkassovoePorucheniePoluchennoe!): DocumentInkassovoePorucheniePoluchennoe
	DocumentInkassovoePorucheniePoluchennoes(): DocumentInkassovoePorucheniePoluchennoe
	DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha(Key: PrimaryDocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha!): DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha
	DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhas(): DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha
	DocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta(Key: PrimaryDocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta!): DocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta
	DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentas(): DocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta
	CatalogNastroikiObmenaDannymiShtrikhM(Key: PrimaryCatalogNastroikiObmenaDannymiShtrikhM!): CatalogNastroikiObmenaDannymiShtrikhM
	CatalogNastroikiObmenaDannymiShtrikhMs(): CatalogNastroikiObmenaDannymiShtrikhM
	CatalogStatiZatrat(Key: PrimaryCatalogStatiZatrat!): CatalogStatiZatrat
	CatalogStatiZatrats(): CatalogStatiZatrat
	DocumentVozvratTovarovOtPokupatelia(Key: PrimaryDocumentVozvratTovarovOtPokupatelia!): DocumentVozvratTovarovOtPokupatelia
	DocumentVozvratTovarovOtPokupatelias(): DocumentVozvratTovarovOtPokupatelia
	DocumentVozvratTovarovOtPokupateliaTovary(Key: PrimaryDocumentVozvratTovarovOtPokupateliaTovary!): DocumentVozvratTovarovOtPokupateliaTovary
	DocumentVozvratTovarovOtPokupateliaTovarys(): DocumentVozvratTovarovOtPokupateliaTovary
	DocumentVozvratTovarovOtPokupateliaUslugi(Key: PrimaryDocumentVozvratTovarovOtPokupateliaUslugi!): DocumentVozvratTovarovOtPokupateliaUslugi
	DocumentVozvratTovarovOtPokupateliaUslugis(): DocumentVozvratTovarovOtPokupateliaUslugi
	DocumentZakazPostavshchiku(Key: PrimaryDocumentZakazPostavshchiku!): DocumentZakazPostavshchiku
	DocumentZakazPostavshchikus(): DocumentZakazPostavshchiku
	DocumentZakazPostavshchikuTovary(Key: PrimaryDocumentZakazPostavshchikuTovary!): DocumentZakazPostavshchikuTovary
	DocumentZakazPostavshchikuTovarys(): DocumentZakazPostavshchikuTovary
	CatalogSkidkiNatsenki(Key: PrimaryCatalogSkidkiNatsenki!): CatalogSkidkiNatsenki
	CatalogSkidkiNatsenkis(): CatalogSkidkiNatsenki
	CatalogSkidkiNatsenkiUsloviiaPredostavleniia(Key: PrimaryCatalogSkidkiNatsenkiUsloviiaPredostavleniia!): CatalogSkidkiNatsenkiUsloviiaPredostavleniia
	CatalogSkidkiNatsenkiUsloviiaPredostavleniias(): CatalogSkidkiNatsenkiUsloviiaPredostavleniia
	CatalogSkidkiNatsenkiTsenovyeGruppy(Key: PrimaryCatalogSkidkiNatsenkiTsenovyeGruppy!): CatalogSkidkiNatsenkiTsenovyeGruppy
	CatalogSkidkiNatsenkiTsenovyeGruppys(): CatalogSkidkiNatsenkiTsenovyeGruppy
	CatalogSkidkiNatsenkiNaborPodarkov(Key: PrimaryCatalogSkidkiNatsenkiNaborPodarkov!): CatalogSkidkiNatsenkiNaborPodarkov
	CatalogSkidkiNatsenkiNaborPodarkovs(): CatalogSkidkiNatsenkiNaborPodarkov
	CatalogGruppyTsvetov(Key: PrimaryCatalogGruppyTsvetov!): CatalogGruppyTsvetov
	CatalogGruppyTsvetovs(): CatalogGruppyTsvetov
	DocumentDokumentRaschetovSKontragentom(Key: PrimaryDocumentDokumentRaschetovSKontragentom!): DocumentDokumentRaschetovSKontragentom
	DocumentDokumentRaschetovSKontragentoms(): DocumentDokumentRaschetovSKontragentom
	CatalogDogovoryEkvairinga(Key: PrimaryCatalogDogovoryEkvairinga!): CatalogDogovoryEkvairinga
	CatalogDogovoryEkvairingas(): CatalogDogovoryEkvairinga
	CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie(Key: PrimaryCatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie!): CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie
	CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanies(): CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie
	CatalogKachestvo(Key: PrimaryCatalogKachestvo!): CatalogKachestvo
	CatalogKachestvos(): CatalogKachestvo
	DocumentUstanovkaTsenNomenklaturyKontragentov(Key: PrimaryDocumentUstanovkaTsenNomenklaturyKontragentov!): DocumentUstanovkaTsenNomenklaturyKontragentov
	DocumentUstanovkaTsenNomenklaturyKontragentovs(): DocumentUstanovkaTsenNomenklaturyKontragentov
	DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen(Key: PrimaryDocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen!): DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen
	DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsens(): DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen
	DocumentUstanovkaTsenNomenklaturyKontragentovTovary(Key: PrimaryDocumentUstanovkaTsenNomenklaturyKontragentovTovary!): DocumentUstanovkaTsenNomenklaturyKontragentovTovary
	DocumentUstanovkaTsenNomenklaturyKontragentovTovarys(): DocumentUstanovkaTsenNomenklaturyKontragentovTovary
	DocumentProtsentPoterPoDavaltsam(Key: PrimaryDocumentProtsentPoterPoDavaltsam!): DocumentProtsentPoterPoDavaltsam
	DocumentProtsentPoterPoDavaltsams(): DocumentProtsentPoterPoDavaltsam
	DocumentProtsentPoterPoDavaltsamProtsenty(Key: PrimaryDocumentProtsentPoterPoDavaltsamProtsenty!): DocumentProtsentPoterPoDavaltsamProtsenty
	DocumentProtsentPoterPoDavaltsamProtsentys(): DocumentProtsentPoterPoDavaltsamProtsenty
	CatalogTovarnyePozitsii(Key: PrimaryCatalogTovarnyePozitsii!): CatalogTovarnyePozitsii
	CatalogTovarnyePozitsiis(): CatalogTovarnyePozitsii
	DocumentPlatezhnoePoruchenieIskhodiashchee(Key: PrimaryDocumentPlatezhnoePoruchenieIskhodiashchee!): DocumentPlatezhnoePoruchenieIskhodiashchee
	DocumentPlatezhnoePoruchenieIskhodiashchees(): DocumentPlatezhnoePoruchenieIskhodiashchee
	DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha(Key: PrimaryDocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha!): DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha
	DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhas(): DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha
	DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta(Key: PrimaryDocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta!): DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta
	DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentas(): DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta
	CatalogfmOrganizatsionnoPravovyeFormy(Key: PrimaryCatalogfmOrganizatsionnoPravovyeFormy!): CatalogfmOrganizatsionnoPravovyeFormy
	CatalogfmOrganizatsionnoPravovyeFormys(): CatalogfmOrganizatsionnoPravovyeFormy
	CatalogTipyTsenNomenklatury(Key: PrimaryCatalogTipyTsenNomenklatury!): CatalogTipyTsenNomenklatury
	CatalogTipyTsenNomenklaturys(): CatalogTipyTsenNomenklatury
	CatalogStatiOtchetaPoProdazham(Key: PrimaryCatalogStatiOtchetaPoProdazham!): CatalogStatiOtchetaPoProdazham
	CatalogStatiOtchetaPoProdazhams(): CatalogStatiOtchetaPoProdazham
	CatalogVidyKodirovokiIzdelii(Key: PrimaryCatalogVidyKodirovokiIzdelii!): CatalogVidyKodirovokiIzdelii
	CatalogVidyKodirovokiIzdeliis(): CatalogVidyKodirovokiIzdelii
	CatalogVidyKodirovokiIzdeliiElementyKodirovki(Key: PrimaryCatalogVidyKodirovokiIzdeliiElementyKodirovki!): CatalogVidyKodirovokiIzdeliiElementyKodirovki
	CatalogVidyKodirovokiIzdeliiElementyKodirovkis(): CatalogVidyKodirovokiIzdeliiElementyKodirovki
	DocumentUstanovkaSkidokNomenklatury(Key: PrimaryDocumentUstanovkaSkidokNomenklatury!): DocumentUstanovkaSkidokNomenklatury
	DocumentUstanovkaSkidokNomenklaturys(): DocumentUstanovkaSkidokNomenklatury
	DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli(Key: PrimaryDocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli!): DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli
	DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedelis(): DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli
	DocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty(Key: PrimaryDocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty!): DocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty
	DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartys(): DocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty
	DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki(Key: PrimaryDocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki!): DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki
	DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkis(): DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki
	DocumentUstanovkaSkidokNomenklaturyTovary(Key: PrimaryDocumentUstanovkaSkidokNomenklaturyTovary!): DocumentUstanovkaSkidokNomenklaturyTovary
	DocumentUstanovkaSkidokNomenklaturyTovarys(): DocumentUstanovkaSkidokNomenklaturyTovary
	CatalogUsloviiaPredostavleniiaSkidokNatsenok(Key: PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenok!): CatalogUsloviiaPredostavleniiaSkidokNatsenok
	CatalogUsloviiaPredostavleniiaSkidokNatsenoks(): CatalogUsloviiaPredostavleniiaSkidokNatsenok
	CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia(Key: PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia!): CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia
	CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviias(): CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia
	CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli(Key: PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli!): CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli
	CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchatelis(): CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli
	CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki(Key: PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki!): CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki
	CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkis(): CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki
	OutPay(Key: PrimaryOutPay!): OutPay
	OutPays(): OutPay
	DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha(Key: PrimaryDocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha!): DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha
	DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhas(): DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha
	DocumentRaskhodnyiKassovyiOrderOplata(Key: PrimaryDocumentRaskhodnyiKassovyiOrderOplata!): DocumentRaskhodnyiKassovyiOrderOplata
	DocumentRaskhodnyiKassovyiOrderOplatas(): DocumentRaskhodnyiKassovyiOrderOplata
	DocumentRaskhodnyiKassovyiOrderTovary(Key: PrimaryDocumentRaskhodnyiKassovyiOrderTovary!): DocumentRaskhodnyiKassovyiOrderTovary
	DocumentRaskhodnyiKassovyiOrderTovarys(): DocumentRaskhodnyiKassovyiOrderTovary
	DocumentSchetNaOplatuPostavshchika(Key: PrimaryDocumentSchetNaOplatuPostavshchika!): DocumentSchetNaOplatuPostavshchika
	DocumentSchetNaOplatuPostavshchikas(): DocumentSchetNaOplatuPostavshchika
	DocumentSchetNaOplatuPostavshchikaTovary(Key: PrimaryDocumentSchetNaOplatuPostavshchikaTovary!): DocumentSchetNaOplatuPostavshchikaTovary
	DocumentSchetNaOplatuPostavshchikaTovarys(): DocumentSchetNaOplatuPostavshchikaTovary
	DocumentSchetNaOplatuPostavshchikaUslugi(Key: PrimaryDocumentSchetNaOplatuPostavshchikaUslugi!): DocumentSchetNaOplatuPostavshchikaUslugi
	DocumentSchetNaOplatuPostavshchikaUslugis(): DocumentSchetNaOplatuPostavshchikaUslugi
	DocumentReestrSpetssviaz(Key: PrimaryDocumentReestrSpetssviaz!): DocumentReestrSpetssviaz
	DocumentReestrSpetssviazs(): DocumentReestrSpetssviaz
	DocumentReestrSpetssviazKlienty(Key: PrimaryDocumentReestrSpetssviazKlienty!): DocumentReestrSpetssviazKlienty
	DocumentReestrSpetssviazKlientys(): DocumentReestrSpetssviazKlienty
	DocumentJournalKassovyeDokumenty(Key: PrimaryDocumentJournalKassovyeDokumenty!): DocumentJournalKassovyeDokumenty
	DocumentJournalKassovyeDokumentys(): DocumentJournalKassovyeDokumenty
	InitialInstance(Key: PrimaryInitialInstance!): InitialInstance
	InitialInstances(): InitialInstance
	DocumentVvodNachalnykhOstatkovVzaimoraschety(Key: PrimaryDocumentVvodNachalnykhOstatkovVzaimoraschety!): DocumentVvodNachalnykhOstatkovVzaimoraschety
	DocumentVvodNachalnykhOstatkovVzaimoraschetys(): DocumentVvodNachalnykhOstatkovVzaimoraschety
	DocumentVvodNachalnykhOstatkovTovary(Key: PrimaryDocumentVvodNachalnykhOstatkovTovary!): DocumentVvodNachalnykhOstatkovTovary
	DocumentVvodNachalnykhOstatkovTovarys(): DocumentVvodNachalnykhOstatkovTovary
	Posting(Key: PrimaryPosting!): Posting
	Postings(): Posting
	DocumentOprikhodovanieTovarovTovary(Key: PrimaryDocumentOprikhodovanieTovarovTovary!): DocumentOprikhodovanieTovarovTovary
	DocumentOprikhodovanieTovarovTovarys(): DocumentOprikhodovanieTovarovTovary
	DocumentOprikhodovanieTovarovSertifikaty(Key: PrimaryDocumentOprikhodovanieTovarovSertifikaty!): DocumentOprikhodovanieTovarovSertifikaty
	DocumentOprikhodovanieTovarovSertifikatys(): DocumentOprikhodovanieTovarovSertifikaty
	CatalogKomplekty(Key: PrimaryCatalogKomplekty!): CatalogKomplekty
	CatalogKomplektys(): CatalogKomplekty
	DocumentPereotsenkaTovarovPriniatykhNaKomissiiu(Key: PrimaryDocumentPereotsenkaTovarovPriniatykhNaKomissiiu!): DocumentPereotsenkaTovarovPriniatykhNaKomissiiu
	DocumentPereotsenkaTovarovPriniatykhNaKomissiius(): DocumentPereotsenkaTovarovPriniatykhNaKomissiiu
	DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary(Key: PrimaryDocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary!): DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary
	DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovarys(): DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary
	DocumentElektronnoePismo(Key: PrimaryDocumentElektronnoePismo!): DocumentElektronnoePismo
	DocumentElektronnoePismos(): DocumentElektronnoePismo
	DocumentElektronnoePismoKomuTCh(Key: PrimaryDocumentElektronnoePismoKomuTCh!): DocumentElektronnoePismoKomuTCh
	DocumentElektronnoePismoKomuTChs(): DocumentElektronnoePismoKomuTCh
	DocumentElektronnoePismoKopiiTCh(Key: PrimaryDocumentElektronnoePismoKopiiTCh!): DocumentElektronnoePismoKopiiTCh
	DocumentElektronnoePismoKopiiTChs(): DocumentElektronnoePismoKopiiTCh
	DocumentElektronnoePismoSkrytyeKopiiTCh(Key: PrimaryDocumentElektronnoePismoSkrytyeKopiiTCh!): DocumentElektronnoePismoSkrytyeKopiiTCh
	DocumentElektronnoePismoSkrytyeKopiiTChs(): DocumentElektronnoePismoSkrytyeKopiiTCh
	CatalogGruppyDefektov(Key: PrimaryCatalogGruppyDefektov!): CatalogGruppyDefektov
	CatalogGruppyDefektovs(): CatalogGruppyDefektov
	CatalogfmAnketaKlientaBenefitsariia(Key: PrimaryCatalogfmAnketaKlientaBenefitsariia!): CatalogfmAnketaKlientaBenefitsariia
	CatalogfmAnketaKlientaBenefitsariias(): CatalogfmAnketaKlientaBenefitsariia
	CatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta(Key: PrimaryCatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta!): CatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta
	CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentas(): CatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta
	CatalogTsenovyeGruppy(Key: PrimaryCatalogTsenovyeGruppy!): CatalogTsenovyeGruppy
	CatalogTsenovyeGruppys(): CatalogTsenovyeGruppy
	CatalogPravilaTsenoobrazovaniia(Key: PrimaryCatalogPravilaTsenoobrazovaniia!): CatalogPravilaTsenoobrazovaniia
	CatalogPravilaTsenoobrazovaniias(): CatalogPravilaTsenoobrazovaniia
	CatalogPravilaTsenoobrazovaniiaTsenovyeGruppy(Key: PrimaryCatalogPravilaTsenoobrazovaniiaTsenovyeGruppy!): CatalogPravilaTsenoobrazovaniiaTsenovyeGruppy
	CatalogPravilaTsenoobrazovaniiaTsenovyeGruppys(): CatalogPravilaTsenoobrazovaniiaTsenovyeGruppy
	DocumentObieiavlenieNaVznosNalichnymi(Key: PrimaryDocumentObieiavlenieNaVznosNalichnymi!): DocumentObieiavlenieNaVznosNalichnymi
	DocumentObieiavlenieNaVznosNalichnymis(): DocumentObieiavlenieNaVznosNalichnymi
	CatalogValiuty(Key: PrimaryCatalogValiuty!): CatalogValiuty
	CatalogValiutys(): CatalogValiuty
	DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku(Key: PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku!): DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku
	DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkus(): DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku
	DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary(Key: PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary!): DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary
	DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovarys(): DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary
	DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi(Key: PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi!): DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi
	DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugis(): DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi
	CatalogKassyKKM(Key: PrimaryCatalogKassyKKM!): CatalogKassyKKM
	CatalogKassyKKMs(): CatalogKassyKKM
	Probe(Key: PrimaryProbe!): Probe
	Probes(): Probe
	CatalogGruppyDostupa(Key: PrimaryCatalogGruppyDostupa!): CatalogGruppyDostupa
	CatalogGruppyDostupas(): CatalogGruppyDostupa
	CatalogGruppyDostupaPolzovateli(Key: PrimaryCatalogGruppyDostupaPolzovateli!): CatalogGruppyDostupaPolzovateli
	CatalogGruppyDostupaPolzovatelis(): CatalogGruppyDostupaPolzovateli
	CatalogGruppyDostupaVidyDostupa(Key: PrimaryCatalogGruppyDostupaVidyDostupa!): CatalogGruppyDostupaVidyDostupa
	CatalogGruppyDostupaVidyDostupas(): CatalogGruppyDostupaVidyDostupa
	CatalogGruppyDostupaZnacheniiaDostupa(Key: PrimaryCatalogGruppyDostupaZnacheniiaDostupa!): CatalogGruppyDostupaZnacheniiaDostupa
	CatalogGruppyDostupaZnacheniiaDostupas(): CatalogGruppyDostupaZnacheniiaDostupa
	CatalogGruppyDostupaDostupPoPodsistemam(Key: PrimaryCatalogGruppyDostupaDostupPoPodsistemam!): CatalogGruppyDostupaDostupPoPodsistemam
	CatalogGruppyDostupaDostupPoPodsistemams(): CatalogGruppyDostupaDostupPoPodsistemam
	CatalogVidyKontaktnoiInformatsii(Key: PrimaryCatalogVidyKontaktnoiInformatsii!): CatalogVidyKontaktnoiInformatsii
	CatalogVidyKontaktnoiInformatsiis(): CatalogVidyKontaktnoiInformatsii
	CatalogNomenklaturnyeGruppy(Key: PrimaryCatalogNomenklaturnyeGruppy!): CatalogNomenklaturnyeGruppy
	CatalogNomenklaturnyeGruppys(): CatalogNomenklaturnyeGruppy
	DocumentReestrSchetov(Key: PrimaryDocumentReestrSchetov!): DocumentReestrSchetov
	DocumentReestrSchetovs(): DocumentReestrSchetov
	DocumentReestrSchetovReestr(Key: PrimaryDocumentReestrSchetovReestr!): DocumentReestrSchetovReestr
	DocumentReestrSchetovReestrs(): DocumentReestrSchetovReestr
	DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu(Key: PrimaryDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu!): DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu
	DocumentInventarizatsiiaTovarovOtdannykhNaKomissiius(): DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu
	DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary(Key: PrimaryDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary!): DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary
	DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovarys(): DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary
	CatalogKlassifikatorStranMira(Key: PrimaryCatalogKlassifikatorStranMira!): CatalogKlassifikatorStranMira
	CatalogKlassifikatorStranMiras(): CatalogKlassifikatorStranMira
	CatalogKlassifikatorEdinitsIzmereniia(Key: PrimaryCatalogKlassifikatorEdinitsIzmereniia!): CatalogKlassifikatorEdinitsIzmereniia
	CatalogKlassifikatorEdinitsIzmereniias(): CatalogKlassifikatorEdinitsIzmereniia
	CatalogNastroikiRMK(Key: PrimaryCatalogNastroikiRMK!): CatalogNastroikiRMK
	CatalogNastroikiRMKs(): CatalogNastroikiRMK
	CatalogNastroikiRMKPoriadokPrimeneniiaSkidok(Key: PrimaryCatalogNastroikiRMKPoriadokPrimeneniiaSkidok!): CatalogNastroikiRMKPoriadokPrimeneniiaSkidok
	CatalogNastroikiRMKPoriadokPrimeneniiaSkidoks(): CatalogNastroikiRMKPoriadokPrimeneniiaSkidok
	CatalogNastroikiRMKSostavNaimenovaniia(Key: PrimaryCatalogNastroikiRMKSostavNaimenovaniia!): CatalogNastroikiRMKSostavNaimenovaniia
	CatalogNastroikiRMKSostavNaimenovaniias(): CatalogNastroikiRMKSostavNaimenovaniia
	CatalogKharakteristikiNomenklatury(Key: PrimaryCatalogKharakteristikiNomenklatury!): CatalogKharakteristikiNomenklatury
	CatalogKharakteristikiNomenklaturys(): CatalogKharakteristikiNomenklatury
	CatalogKharakteristikiNomenklaturySpetsifikatsiia(Key: PrimaryCatalogKharakteristikiNomenklaturySpetsifikatsiia!): CatalogKharakteristikiNomenklaturySpetsifikatsiia
	CatalogKharakteristikiNomenklaturySpetsifikatsiias(): CatalogKharakteristikiNomenklaturySpetsifikatsiia
	DocumentOtborTovarov(Key: PrimaryDocumentOtborTovarov!): DocumentOtborTovarov
	DocumentOtborTovarovs(): DocumentOtborTovarov
	DocumentOtborTovarovTovary(Key: PrimaryDocumentOtborTovarovTovary!): DocumentOtborTovarovTovary
	DocumentOtborTovarovTovarys(): DocumentOtborTovarovTovary
	DocumentOtborTovarovTovaryKOtboru(Key: PrimaryDocumentOtborTovarovTovaryKOtboru!): DocumentOtborTovarovTovaryKOtboru
	DocumentOtborTovarovTovaryKOtborus(): DocumentOtborTovarovTovaryKOtboru
	CatalogSposobyDostavkiTovara(Key: PrimaryCatalogSposobyDostavkiTovara!): CatalogSposobyDostavkiTovara
	CatalogSposobyDostavkiTovaras(): CatalogSposobyDostavkiTovara
	CatalogPodrazdeleniia(Key: PrimaryCatalogPodrazdeleniia!): CatalogPodrazdeleniia
	CatalogPodrazdeleniias(): CatalogPodrazdeleniia
	DocumentJournalPreiskuranty(Key: PrimaryDocumentJournalPreiskuranty!): DocumentJournalPreiskuranty
	DocumentJournalPreiskurantys(): DocumentJournalPreiskuranty
	CatalogRelizyIuvelirnykhSalonov(Key: PrimaryCatalogRelizyIuvelirnykhSalonov!): CatalogRelizyIuvelirnykhSalonov
	CatalogRelizyIuvelirnykhSalonovs(): CatalogRelizyIuvelirnykhSalonov
	CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy(Key: PrimaryCatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy!): CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy
	CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizys(): CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy
	DocumentOtchetKomissioneraOProdazhakh(Key: PrimaryDocumentOtchetKomissioneraOProdazhakh!): DocumentOtchetKomissioneraOProdazhakh
	DocumentOtchetKomissioneraOProdazhakhs(): DocumentOtchetKomissioneraOProdazhakh
	DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva(Key: PrimaryDocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva!): DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva
	DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvas(): DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva
	DocumentOtchetKomissioneraOProdazhakhTovary(Key: PrimaryDocumentOtchetKomissioneraOProdazhakhTovary!): DocumentOtchetKomissioneraOProdazhakhTovary
	DocumentOtchetKomissioneraOProdazhakhTovarys(): DocumentOtchetKomissioneraOProdazhakhTovary
	CatalogTovarnyeKategorii(Key: PrimaryCatalogTovarnyeKategorii!): CatalogTovarnyeKategorii
	CatalogTovarnyeKategoriis(): CatalogTovarnyeKategorii
	CatalogDokumentyUdostoveriaiushchieLichnost(Key: PrimaryCatalogDokumentyUdostoveriaiushchieLichnost!): CatalogDokumentyUdostoveriaiushchieLichnost
	CatalogDokumentyUdostoveriaiushchieLichnosts(): CatalogDokumentyUdostoveriaiushchieLichnost
	CatalogFiltryDliaElektronnykhPisem(Key: PrimaryCatalogFiltryDliaElektronnykhPisem!): CatalogFiltryDliaElektronnykhPisem
	CatalogFiltryDliaElektronnykhPisems(): CatalogFiltryDliaElektronnykhPisem
	CatalogFiltryDliaElektronnykhPisemDeistviiaFiltra(Key: PrimaryCatalogFiltryDliaElektronnykhPisemDeistviiaFiltra!): CatalogFiltryDliaElektronnykhPisemDeistviiaFiltra
	CatalogFiltryDliaElektronnykhPisemDeistviiaFiltras(): CatalogFiltryDliaElektronnykhPisemDeistviiaFiltra
	CatalogFiltryDliaElektronnykhPisemUsloviiaFiltra(Key: PrimaryCatalogFiltryDliaElektronnykhPisemUsloviiaFiltra!): CatalogFiltryDliaElektronnykhPisemUsloviiaFiltra
	CatalogFiltryDliaElektronnykhPisemUsloviiaFiltras(): CatalogFiltryDliaElektronnykhPisemUsloviiaFiltra
	DocumentPreiskurantTsenNaTsvKamni(Key: PrimaryDocumentPreiskurantTsenNaTsvKamni!): DocumentPreiskurantTsenNaTsvKamni
	DocumentPreiskurantTsenNaTsvKamnis(): DocumentPreiskurantTsenNaTsvKamni
	DocumentPreiskurantTsenNaTsvKamniTablitsa(Key: PrimaryDocumentPreiskurantTsenNaTsvKamniTablitsa!): DocumentPreiskurantTsenNaTsvKamniTablitsa
	DocumentPreiskurantTsenNaTsvKamniTablitsas(): DocumentPreiskurantTsenNaTsvKamniTablitsa
	Size(Key: PrimarySize!): Size
	Sizes(): Size
	CatalogTipyDragotsennykhMetallov(Key: PrimaryCatalogTipyDragotsennykhMetallov!): CatalogTipyDragotsennykhMetallov
	CatalogTipyDragotsennykhMetallovs(): CatalogTipyDragotsennykhMetallov
	DocumentTelemarketing(Key: PrimaryDocumentTelemarketing!): DocumentTelemarketing
	DocumentTelemarketings(): DocumentTelemarketing
	DocumentTelemarketingUchastniki(Key: PrimaryDocumentTelemarketingUchastniki!): DocumentTelemarketingUchastniki
	DocumentTelemarketingUchastnikis(): DocumentTelemarketingUchastniki
	DocumentVozvratDavalcheskogoMetalla(Key: PrimaryDocumentVozvratDavalcheskogoMetalla!): DocumentVozvratDavalcheskogoMetalla
	DocumentVozvratDavalcheskogoMetallas(): DocumentVozvratDavalcheskogoMetalla
	CatalogAdresnyeSokrashcheniia(Key: PrimaryCatalogAdresnyeSokrashcheniia!): CatalogAdresnyeSokrashcheniia
	CatalogAdresnyeSokrashcheniias(): CatalogAdresnyeSokrashcheniia
	DocumentRassylkaAnket(Key: PrimaryDocumentRassylkaAnket!): DocumentRassylkaAnket
	DocumentRassylkaAnkets(): DocumentRassylkaAnket
	DocumentRassylkaAnketVlozheniia(Key: PrimaryDocumentRassylkaAnketVlozheniia!): DocumentRassylkaAnketVlozheniia
	DocumentRassylkaAnketVlozheniias(): DocumentRassylkaAnketVlozheniia
	DocumentRassylkaAnketPoluchateli(Key: PrimaryDocumentRassylkaAnketPoluchateli!): DocumentRassylkaAnketPoluchateli
	DocumentRassylkaAnketPoluchatelis(): DocumentRassylkaAnketPoluchateli
	CatalogVidyDeiatelnostiKontragentov(Key: PrimaryCatalogVidyDeiatelnostiKontragentov!): CatalogVidyDeiatelnostiKontragentov
	CatalogVidyDeiatelnostiKontragentovs(): CatalogVidyDeiatelnostiKontragentov
	CatalogTorgovoeOborudovanie(Key: PrimaryCatalogTorgovoeOborudovanie!): CatalogTorgovoeOborudovanie
	CatalogTorgovoeOborudovanies(): CatalogTorgovoeOborudovanie
	CatalogSkhemyRealizatsii(Key: PrimaryCatalogSkhemyRealizatsii!): CatalogSkhemyRealizatsii
	CatalogSkhemyRealizatsiis(): CatalogSkhemyRealizatsii
	CatalogSkhemyRealizatsiiEtapySkhemy(Key: PrimaryCatalogSkhemyRealizatsiiEtapySkhemy!): CatalogSkhemyRealizatsiiEtapySkhemy
	CatalogSkhemyRealizatsiiEtapySkhemys(): CatalogSkhemyRealizatsiiEtapySkhemy
	CatalogPodkliuchaemoeOborudovanie(Key: PrimaryCatalogPodkliuchaemoeOborudovanie!): CatalogPodkliuchaemoeOborudovanie
	CatalogPodkliuchaemoeOborudovanies(): CatalogPodkliuchaemoeOborudovanie
	DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii(Key: PrimaryDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii!): DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii
	DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniis(): DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii
	DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov(Key: PrimaryDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov!): DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov
	DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovs(): DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov
	CatalogGabarity(Key: PrimaryCatalogGabarity!): CatalogGabarity
	CatalogGabaritys(): CatalogGabarity
	DocumentZakazKlienta(Key: PrimaryDocumentZakazKlienta!): DocumentZakazKlienta
	DocumentZakazKlientas(): DocumentZakazKlienta
	DocumentZakazKlientaTovary(Key: PrimaryDocumentZakazKlientaTovary!): DocumentZakazKlientaTovary
	DocumentZakazKlientaTovarys(): DocumentZakazKlientaTovary
	ArriveFromManufacturing(Key: PrimaryArriveFromManufacturing!): ArriveFromManufacturing
	ArriveFromManufacturings(): ArriveFromManufacturing
	ArriveFromManufacturingInstance(Key: PrimaryArriveFromManufacturingInstance!): ArriveFromManufacturingInstance
	ArriveFromManufacturingInstances(): ArriveFromManufacturingInstance
	DocumentPostuplenieProduktsiiIzProizvodstvaMaterialy(Key: PrimaryDocumentPostuplenieProduktsiiIzProizvodstvaMaterialy!): DocumentPostuplenieProduktsiiIzProizvodstvaMaterialy
	DocumentPostuplenieProduktsiiIzProizvodstvaMaterialys(): DocumentPostuplenieProduktsiiIzProizvodstvaMaterialy
	DocumentJournalZakazyPostavshchikam(Key: PrimaryDocumentJournalZakazyPostavshchikam!): DocumentJournalZakazyPostavshchikam
	DocumentJournalZakazyPostavshchikams(): DocumentJournalZakazyPostavshchikam
	DocumentJournalSkladskieDokumenty(Key: PrimaryDocumentJournalSkladskieDokumenty!): DocumentJournalSkladskieDokumenty
	DocumentJournalSkladskieDokumentys(): DocumentJournalSkladskieDokumenty
	CatalogsmsUsloviiaOtboraDiskontnykhKart(Key: PrimaryCatalogsmsUsloviiaOtboraDiskontnykhKart!): CatalogsmsUsloviiaOtboraDiskontnykhKart
	CatalogsmsUsloviiaOtboraDiskontnykhKarts(): CatalogsmsUsloviiaOtboraDiskontnykhKart
	Arrive(Key: PrimaryArrive!): Arrive
	Arrives(): Arrive
	DocumentPostuplenieTovarovUslugTovary(Key: PrimaryDocumentPostuplenieTovarovUslugTovary!): DocumentPostuplenieTovarovUslugTovary
	DocumentPostuplenieTovarovUslugTovarys(): DocumentPostuplenieTovarovUslugTovary
	DocumentPostuplenieTovarovUslugUslugi(Key: PrimaryDocumentPostuplenieTovarovUslugUslugi!): DocumentPostuplenieTovarovUslugUslugi
	DocumentPostuplenieTovarovUslugUslugis(): DocumentPostuplenieTovarovUslugUslugi
	DocumentSchetFakturaVydannyi(Key: PrimaryDocumentSchetFakturaVydannyi!): DocumentSchetFakturaVydannyi
	DocumentSchetFakturaVydannyis(): DocumentSchetFakturaVydannyi
	DocumentPlanProdazhPoSalonam(Key: PrimaryDocumentPlanProdazhPoSalonam!): DocumentPlanProdazhPoSalonam
	DocumentPlanProdazhPoSalonams(): DocumentPlanProdazhPoSalonam
	DocumentPlanProdazhPoSalonamSalony(Key: PrimaryDocumentPlanProdazhPoSalonamSalony!): DocumentPlanProdazhPoSalonamSalony
	DocumentPlanProdazhPoSalonamSalonys(): DocumentPlanProdazhPoSalonamSalony
	DocumentPlanProdazhPoSalonamDniPoGrafiku(Key: PrimaryDocumentPlanProdazhPoSalonamDniPoGrafiku!): DocumentPlanProdazhPoSalonamDniPoGrafiku
	DocumentPlanProdazhPoSalonamDniPoGrafikus(): DocumentPlanProdazhPoSalonamDniPoGrafiku
	CatalogBankovskieScheta(Key: PrimaryCatalogBankovskieScheta!): CatalogBankovskieScheta
	CatalogBankovskieSchetas(): CatalogBankovskieScheta
	DocumentStornirovanieOtchetaKomitentuOProdazhakh(Key: PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakh!): DocumentStornirovanieOtchetaKomitentuOProdazhakh
	DocumentStornirovanieOtchetaKomitentuOProdazhakhs(): DocumentStornirovanieOtchetaKomitentuOProdazhakh
	DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva(Key: PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva!): DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva
	DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvas(): DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva
	DocumentStornirovanieOtchetaKomitentuOProdazhakhTovary(Key: PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakhTovary!): DocumentStornirovanieOtchetaKomitentuOProdazhakhTovary
	DocumentStornirovanieOtchetaKomitentuOProdazhakhTovarys(): DocumentStornirovanieOtchetaKomitentuOProdazhakhTovary
	DocumentPeredachaVRemont(Key: PrimaryDocumentPeredachaVRemont!): DocumentPeredachaVRemont
	DocumentPeredachaVRemonts(): DocumentPeredachaVRemont
	DocumentPeredachaVRemontIzdeliiaPriniatyeVRemont(Key: PrimaryDocumentPeredachaVRemontIzdeliiaPriniatyeVRemont!): DocumentPeredachaVRemontIzdeliiaPriniatyeVRemont
	DocumentPeredachaVRemontIzdeliiaPriniatyeVRemonts(): DocumentPeredachaVRemontIzdeliiaPriniatyeVRemont
	DocumentPeredachaVRemontTovary(Key: PrimaryDocumentPeredachaVRemontTovary!): DocumentPeredachaVRemontTovary
	DocumentPeredachaVRemontTovarys(): DocumentPeredachaVRemontTovary
	CatalogPolzovateli(Key: PrimaryCatalogPolzovateli!): CatalogPolzovateli
	CatalogPolzovatelis(): CatalogPolzovateli
	CatalogTsenovyeKoridory(Key: PrimaryCatalogTsenovyeKoridory!): CatalogTsenovyeKoridory
	CatalogTsenovyeKoridorys(): CatalogTsenovyeKoridory
	CatalogGruppySkladov(Key: PrimaryCatalogGruppySkladov!): CatalogGruppySkladov
	CatalogGruppySkladovs(): CatalogGruppySkladov
	CatalogGruppySkladovSklady(Key: PrimaryCatalogGruppySkladovSklady!): CatalogGruppySkladovSklady
	CatalogGruppySkladovSkladys(): CatalogGruppySkladovSklady
}
`

type GqlResolver struct{}

// Add native info to context and add cors header
func withCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		h.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func StartServer() {
	opts := []gqlserver.SchemaOpt{gqlserver.UseFieldResolvers()}
	resolver := odata.GqlResolver{Client: odata.NewClient("web", "12345", "http://127.0.0.1:8091/JewellerTrade/odata/standard.odata/")}
	schema := gqlserver.MustParseSchema(Schema, &resolver, opts...)
	http.Handle("/query", withCors(&relay.Handler{Schema: schema}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
