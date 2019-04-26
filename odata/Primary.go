package odata

type PrimaryAccumulationRegisterPartiiTovarovVProizvodstve struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterPartiiTovarovVProizvodstve) APIEntityType() string {
	return "AccumulationRegister_ПартииТоваровВПроизводстве"
}
func (p PrimaryAccumulationRegisterPartiiTovarovVProizvodstve) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterPartiiTovarovVProizvodstveRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterPartiiTovarovVProizvodstveRecordType) APIEntityType() string {
	return "AccumulationRegister_ПартииТоваровВПроизводстве_RecordType"
}
func (p PrimaryAccumulationRegisterPartiiTovarovVProizvodstveRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterVzaimoraschetySPodotchetnymiLitsami struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterVzaimoraschetySPodotchetnymiLitsami) APIEntityType() string {
	return "AccumulationRegister_ВзаиморасчетыСПодотчетнымиЛицами"
}
func (p PrimaryAccumulationRegisterVzaimoraschetySPodotchetnymiLitsami) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType) APIEntityType() string {
	return "AccumulationRegister_ВзаиморасчетыСПодотчетнымиЛицами_RecordType"
}
func (p PrimaryAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterVnutrennieZakazy struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterVnutrennieZakazy) APIEntityType() string {
	return "AccumulationRegister_ВнутренниеЗаказы"
}
func (p PrimaryAccumulationRegisterVnutrennieZakazy) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterVnutrennieZakazyRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterVnutrennieZakazyRecordType) APIEntityType() string {
	return "AccumulationRegister_ВнутренниеЗаказы_RecordType"
}
func (p PrimaryAccumulationRegisterVnutrennieZakazyRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstvaKomitenta struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstvaKomitenta) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредстваКомитента"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstvaKomitenta) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstvaKomitentaRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstvaKomitentaRecordType) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредстваКомитента_RecordType"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstvaKomitentaRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterZakazyKlientov struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterZakazyKlientov) APIEntityType() string {
	return "AccumulationRegister_ЗаказыКлиентов"
}
func (p PrimaryAccumulationRegisterZakazyKlientov) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterZakazyKlientovRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterZakazyKlientovRecordType) APIEntityType() string {
	return "AccumulationRegister_ЗаказыКлиентов_RecordType"
}
func (p PrimaryAccumulationRegisterZakazyKlientovRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSummyPoFinmonitoringuRoznitsa struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterSummyPoFinmonitoringuRoznitsa) APIEntityType() string {
	return "AccumulationRegister_СуммыПоФинмониторингуРозница"
}
func (p PrimaryAccumulationRegisterSummyPoFinmonitoringuRoznitsa) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType) APIEntityType() string {
	return "AccumulationRegister_СуммыПоФинмониторингуРозница_RecordType"
}
func (p PrimaryAccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstvaKPolucheniiu struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstvaKPolucheniiu) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредстваКПолучению"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstvaKPolucheniiu) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредстваКПолучению_RecordType"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterProdazhiPoDiskontnymKartam struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterProdazhiPoDiskontnymKartam) APIEntityType() string {
	return "AccumulationRegister_ПродажиПоДисконтнымКартам"
}
func (p PrimaryAccumulationRegisterProdazhiPoDiskontnymKartam) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterProdazhiPoDiskontnymKartamRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterProdazhiPoDiskontnymKartamRecordType) APIEntityType() string {
	return "AccumulationRegister_ПродажиПоДисконтнымКартам_RecordType"
}
func (p PrimaryAccumulationRegisterProdazhiPoDiskontnymKartamRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterTovaryPoluchennye struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryPoluchennye) APIEntityType() string {
	return "AccumulationRegister_ТоварыПолученные"
}
func (p PrimaryAccumulationRegisterTovaryPoluchennye) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterTovaryPoluchennyeRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryPoluchennyeRecordType) APIEntityType() string {
	return "AccumulationRegister_ТоварыПолученные_RecordType"
}
func (p PrimaryAccumulationRegisterTovaryPoluchennyeRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSvobodnyeOstatki struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterSvobodnyeOstatki) APIEntityType() string {
	return "AccumulationRegister_СвободныеОстатки"
}
func (p PrimaryAccumulationRegisterSvobodnyeOstatki) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSvobodnyeOstatkiRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterSvobodnyeOstatkiRecordType) APIEntityType() string {
	return "AccumulationRegister_СвободныеОстатки_RecordType"
}
func (p PrimaryAccumulationRegisterSvobodnyeOstatkiRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSummyVRassrochku struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterSummyVRassrochku) APIEntityType() string {
	return "AccumulationRegister_СуммыВРассрочку"
}
func (p PrimaryAccumulationRegisterSummyVRassrochku) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSummyVRassrochkuRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterSummyVRassrochkuRecordType) APIEntityType() string {
	return "AccumulationRegister_СуммыВРассрочку_RecordType"
}
func (p PrimaryAccumulationRegisterSummyVRassrochkuRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterGrafikPlatezhei struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterGrafikPlatezhei) APIEntityType() string {
	return "AccumulationRegister_ГрафикПлатежей"
}
func (p PrimaryAccumulationRegisterGrafikPlatezhei) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterGrafikPlatezheiRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterGrafikPlatezheiRecordType) APIEntityType() string {
	return "AccumulationRegister_ГрафикПлатежей_RecordType"
}
func (p PrimaryAccumulationRegisterGrafikPlatezheiRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterRoznichnaiaVyruchka struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterRoznichnaiaVyruchka) APIEntityType() string {
	return "AccumulationRegister_РозничнаяВыручка"
}
func (p PrimaryAccumulationRegisterRoznichnaiaVyruchka) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterRoznichnaiaVyruchkaRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterRoznichnaiaVyruchkaRecordType) APIEntityType() string {
	return "AccumulationRegister_РозничнаяВыручка_RecordType"
}
func (p PrimaryAccumulationRegisterRoznichnaiaVyruchkaRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterTovaryVPuti struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryVPuti) APIEntityType() string {
	return "AccumulationRegister_ТоварыВПути"
}
func (p PrimaryAccumulationRegisterTovaryVPuti) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterTovaryVPutiRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryVPutiRecordType) APIEntityType() string {
	return "AccumulationRegister_ТоварыВПути_RecordType"
}
func (p PrimaryAccumulationRegisterTovaryVPutiRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterPoteriMetallaVProizvodstve struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterPoteriMetallaVProizvodstve) APIEntityType() string {
	return "AccumulationRegister_ПотериМеталлаВПроизводстве"
}
func (p PrimaryAccumulationRegisterPoteriMetallaVProizvodstve) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterPoteriMetallaVProizvodstveRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterPoteriMetallaVProizvodstveRecordType) APIEntityType() string {
	return "AccumulationRegister_ПотериМеталлаВПроизводстве_RecordType"
}
func (p PrimaryAccumulationRegisterPoteriMetallaVProizvodstveRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterPartiiTovarovNaSkladakh struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterPartiiTovarovNaSkladakh) APIEntityType() string {
	return "AccumulationRegister_ПартииТоваровНаСкладах"
}
func (p PrimaryAccumulationRegisterPartiiTovarovNaSkladakh) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryProductActionDocument struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryProductActionDocument) APIEntityType() string {
	return "AccumulationRegister_ПартииТоваровНаСкладах_RecordType"
}
func (p PrimaryProductActionDocument) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSummyDokumentovDliaObmena struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterSummyDokumentovDliaObmena) APIEntityType() string {
	return "AccumulationRegister_СуммыДокументовДляОбмена"
}
func (p PrimaryAccumulationRegisterSummyDokumentovDliaObmena) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSummyDokumentovDliaObmenaRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterSummyDokumentovDliaObmenaRecordType) APIEntityType() string {
	return "AccumulationRegister_СуммыДокументовДляОбмена_RecordType"
}
func (p PrimaryAccumulationRegisterSummyDokumentovDliaObmenaRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDvizheniiaDenezhnykhSredstv struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterDvizheniiaDenezhnykhSredstv) APIEntityType() string {
	return "AccumulationRegister_ДвиженияДенежныхСредств"
}
func (p PrimaryAccumulationRegisterDvizheniiaDenezhnykhSredstv) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType) APIEntityType() string {
	return "AccumulationRegister_ДвиженияДенежныхСредств_RecordType"
}
func (p PrimaryAccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterProdazhiPoStatiam struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterProdazhiPoStatiam) APIEntityType() string {
	return "AccumulationRegister_ПродажиПоСтатьям"
}
func (p PrimaryAccumulationRegisterProdazhiPoStatiam) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterProdazhiPoStatiamRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterProdazhiPoStatiamRecordType) APIEntityType() string {
	return "AccumulationRegister_ПродажиПоСтатьям_RecordType"
}
func (p PrimaryAccumulationRegisterProdazhiPoStatiamRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryInformationRegisterTsenyNomenklatury struct {
	Recorder     String
	RecorderType String
}

func (PrimaryInformationRegisterTsenyNomenklatury) APIEntityType() string {
	return "InformationRegister_ЦеныНоменклатуры"
}
func (p PrimaryInformationRegisterTsenyNomenklatury) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryInformationRegisterTsenyNomenklaturyRecordType struct {
	Period                         DateTime
	TipTsenKey                     Guid
	SegmentNomenklaturyKey         Guid
	ItemKey                        Guid
	InstanceKey                    Guid
	KharakteristikaNomenklaturyKey Guid
	SizeKey                        Guid
}

func (PrimaryInformationRegisterTsenyNomenklaturyRecordType) APIEntityType() string {
	return "InformationRegister_ЦеныНоменклатуры_RecordType"
}
func (p PrimaryInformationRegisterTsenyNomenklaturyRecordType) Serialize() string {
	return "Period=" + p.Period.AsParameter() + ",ТипЦен_Key=" + p.TipTsenKey.AsParameter() + ",СегментНоменклатуры_Key=" + p.SegmentNomenklaturyKey.AsParameter() + ",Номенклатура_Key=" + p.ItemKey.AsParameter() + ",СерияНоменклатуры_Key=" + p.InstanceKey.AsParameter() + ",ХарактеристикаНоменклатуры_Key=" + p.KharakteristikaNomenklaturyKey.AsParameter() + ",Размер_Key=" + p.SizeKey.AsParameter()
}

type PrimaryAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse) APIEntityType() string {
	return "AccumulationRegister_СводныеДанныеПоПродажамВРознице"
}
func (p PrimaryAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType) APIEntityType() string {
	return "AccumulationRegister_СводныеДанныеПоПродажамВРознице_RecordType"
}
func (p PrimaryAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstvaVRezerve struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstvaVRezerve) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредстваВРезерве"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstvaVRezerve) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstvaVRezerveRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstvaVRezerveRecordType) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредстваВРезерве_RecordType"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstvaVRezerveRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh) APIEntityType() string {
	return "AccumulationRegister_ТоварыВНеавтоматизированныхТорговыхТочках"
}
func (p PrimaryAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType) APIEntityType() string {
	return "AccumulationRegister_ТоварыВНеавтоматизированныхТорговыхТочках_RecordType"
}
func (p PrimaryAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDavalcheskiiMetallPoteri struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterDavalcheskiiMetallPoteri) APIEntityType() string {
	return "AccumulationRegister_ДавальческийМеталлПотери"
}
func (p PrimaryAccumulationRegisterDavalcheskiiMetallPoteri) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDavalcheskiiMetallPoteriRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterDavalcheskiiMetallPoteriRecordType) APIEntityType() string {
	return "AccumulationRegister_ДавальческийМеталлПотери_RecordType"
}
func (p PrimaryAccumulationRegisterDavalcheskiiMetallPoteriRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryInformationRegisterTsenyPoPreiskurantu struct {
	Recorder     String
	RecorderType String
}

func (PrimaryInformationRegisterTsenyPoPreiskurantu) APIEntityType() string {
	return "InformationRegister_ЦеныПоПрейскуранту"
}
func (p PrimaryInformationRegisterTsenyPoPreiskurantu) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryInformationRegisterTsenyPoPreiskurantuRecordType struct {
	Period           DateTime
	KamenKey         Guid
	FormaOgrankiKey  Guid
	TsvetKamniaKey   Guid
	GruppaTsvetaKey  Guid
	GruppaDefektaKey Guid
	RassevKey        Guid
	Razmer1Ot        Double
	Razmer1Do        Double
}

func (PrimaryInformationRegisterTsenyPoPreiskurantuRecordType) APIEntityType() string {
	return "InformationRegister_ЦеныПоПрейскуранту_RecordType"
}
func (p PrimaryInformationRegisterTsenyPoPreiskurantuRecordType) Serialize() string {
	return "Period=" + p.Period.AsParameter() + ",Камень_Key=" + p.KamenKey.AsParameter() + ",ФормаОгранки_Key=" + p.FormaOgrankiKey.AsParameter() + ",ЦветКамня_Key=" + p.TsvetKamniaKey.AsParameter() + ",ГруппаЦвета_Key=" + p.GruppaTsvetaKey.AsParameter() + ",ГруппаДефекта_Key=" + p.GruppaDefektaKey.AsParameter() + ",Рассев_Key=" + p.RassevKey.AsParameter() + ",Размер1От=" + p.Razmer1Ot.AsParameter() + ",Размер1До=" + p.Razmer1Do.AsParameter()
}

type PrimaryAccumulationRegisterTovaryVOtbore struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryVOtbore) APIEntityType() string {
	return "AccumulationRegister_ТоварыВОтборе"
}
func (p PrimaryAccumulationRegisterTovaryVOtbore) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterTovaryVOtboreRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryVOtboreRecordType) APIEntityType() string {
	return "AccumulationRegister_ТоварыВОтборе_RecordType"
}
func (p PrimaryAccumulationRegisterTovaryVOtboreRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterRealizovannyeTovary struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterRealizovannyeTovary) APIEntityType() string {
	return "AccumulationRegister_РеализованныеТовары"
}
func (p PrimaryAccumulationRegisterRealizovannyeTovary) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterRealizovannyeTovaryRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterRealizovannyeTovaryRecordType) APIEntityType() string {
	return "AccumulationRegister_РеализованныеТовары_RecordType"
}
func (p PrimaryAccumulationRegisterRealizovannyeTovaryRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstvaKomissionera struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstvaKomissionera) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредстваКомиссионера"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstvaKomissionera) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредстваКомиссионера_RecordType"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterProdazhi1 struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterProdazhi1) APIEntityType() string {
	return "AccumulationRegister_Продажи1"
}
func (p PrimaryAccumulationRegisterProdazhi1) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterProdazhi1RecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterProdazhi1RecordType) APIEntityType() string {
	return "AccumulationRegister_Продажи1_RecordType"
}
func (p PrimaryAccumulationRegisterProdazhi1RecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterTovaryNaSkladakhAM struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryNaSkladakhAM) APIEntityType() string {
	return "AccumulationRegister_ТоварыНаСкладахАМ"
}
func (p PrimaryAccumulationRegisterTovaryNaSkladakhAM) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterTovaryNaSkladakhAMRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryNaSkladakhAMRecordType) APIEntityType() string {
	return "AccumulationRegister_ТоварыНаСкладахАМ_RecordType"
}
func (p PrimaryAccumulationRegisterTovaryNaSkladakhAMRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSummyPoFinmonitoringu struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterSummyPoFinmonitoringu) APIEntityType() string {
	return "AccumulationRegister_СуммыПоФинмониторингу"
}
func (p PrimaryAccumulationRegisterSummyPoFinmonitoringu) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSummyPoFinmonitoringuRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterSummyPoFinmonitoringuRecordType) APIEntityType() string {
	return "AccumulationRegister_СуммыПоФинмониторингу_RecordType"
}
func (p PrimaryAccumulationRegisterSummyPoFinmonitoringuRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryInformationRegisterTsenyNomenklaturyKontragentov struct {
	Recorder     String
	RecorderType String
}

func (PrimaryInformationRegisterTsenyNomenklaturyKontragentov) APIEntityType() string {
	return "InformationRegister_ЦеныНоменклатурыКонтрагентов"
}
func (p PrimaryInformationRegisterTsenyNomenklaturyKontragentov) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryInformationRegisterTsenyNomenklaturyKontragentovRecordType struct {
	Period                         DateTime
	TipTsenKey                     Guid
	ItemKey                        Guid
	InstanceKey                    Guid
	KharakteristikaNomenklaturyKey Guid
	SizeKey                        Guid
}

func (PrimaryInformationRegisterTsenyNomenklaturyKontragentovRecordType) APIEntityType() string {
	return "InformationRegister_ЦеныНоменклатурыКонтрагентов_RecordType"
}
func (p PrimaryInformationRegisterTsenyNomenklaturyKontragentovRecordType) Serialize() string {
	return "Period=" + p.Period.AsParameter() + ",ТипЦен_Key=" + p.TipTsenKey.AsParameter() + ",Номенклатура_Key=" + p.ItemKey.AsParameter() + ",СерияНоменклатуры_Key=" + p.InstanceKey.AsParameter() + ",ХарактеристикаНоменклатуры_Key=" + p.KharakteristikaNomenklaturyKey.AsParameter() + ",Размер_Key=" + p.SizeKey.AsParameter()
}

type PrimaryAccumulationRegisterVzaimoraschetySKontragentami struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterVzaimoraschetySKontragentami) APIEntityType() string {
	return "AccumulationRegister_ВзаиморасчетыСКонтрагентами"
}
func (p PrimaryAccumulationRegisterVzaimoraschetySKontragentami) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterVzaimoraschetySKontragentamiRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterVzaimoraschetySKontragentamiRecordType) APIEntityType() string {
	return "AccumulationRegister_ВзаиморасчетыСКонтрагентами_RecordType"
}
func (p PrimaryAccumulationRegisterVzaimoraschetySKontragentamiRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSummyPokupokPoDiskontnymKartam struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterSummyPokupokPoDiskontnymKartam) APIEntityType() string {
	return "AccumulationRegister_СуммыПокупокПоДисконтнымКартам"
}
func (p PrimaryAccumulationRegisterSummyPokupokPoDiskontnymKartam) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType) APIEntityType() string {
	return "AccumulationRegister_СуммыПокупокПоДисконтнымКартам_RecordType"
}
func (p PrimaryAccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterVypolneniePlanaProdazh struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterVypolneniePlanaProdazh) APIEntityType() string {
	return "AccumulationRegister_ВыполнениеПланаПродаж"
}
func (p PrimaryAccumulationRegisterVypolneniePlanaProdazh) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterVypolneniePlanaProdazhRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterVypolneniePlanaProdazhRecordType) APIEntityType() string {
	return "AccumulationRegister_ВыполнениеПланаПродаж_RecordType"
}
func (p PrimaryAccumulationRegisterVypolneniePlanaProdazhRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDavalcheskiiMetall struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterDavalcheskiiMetall) APIEntityType() string {
	return "AccumulationRegister_ДавальческийМеталл"
}
func (p PrimaryAccumulationRegisterDavalcheskiiMetall) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDavalcheskiiMetallRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterDavalcheskiiMetallRecordType) APIEntityType() string {
	return "AccumulationRegister_ДавальческийМеталл_RecordType"
}
func (p PrimaryAccumulationRegisterDavalcheskiiMetallRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstva struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstva) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредства"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstva) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstvaRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstvaRecordType) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредства_RecordType"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstvaRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterTovaryPeredannye struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryPeredannye) APIEntityType() string {
	return "AccumulationRegister_ТоварыПереданные"
}
func (p PrimaryAccumulationRegisterTovaryPeredannye) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterTovaryPeredannyeRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterTovaryPeredannyeRecordType) APIEntityType() string {
	return "AccumulationRegister_ТоварыПереданные_RecordType"
}
func (p PrimaryAccumulationRegisterTovaryPeredannyeRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstvaKSpisaniiu struct {
	Recorder     String
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstvaKSpisaniiu) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредстваКСписанию"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstvaKSpisaniiu) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryAccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType struct {
	Recorder     String
	LineNumber   Int64
	RecorderType String
}

func (PrimaryAccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType) APIEntityType() string {
	return "AccumulationRegister_ДенежныеСредстваКСписанию_RecordType"
}
func (p PrimaryAccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType) Serialize() string {
	return "Recorder=" + p.Recorder.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter() + ",Recorder_Type=" + p.RecorderType.AsParameter()
}

type PrimaryCatalogDogovoryKontragentov struct {
	Key Guid
}

func (PrimaryCatalogDogovoryKontragentov) APIEntityType() string {
	return "Catalog_ДоговорыКонтрагентов"
}
func (p PrimaryCatalogDogovoryKontragentov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryOrder struct {
	Key Guid
}

func (PrimaryOrder) APIEntityType() string {
	return "Document_ЧекККМ"
}
func (p PrimaryOrder) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentChekKKMOplata struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMOplata) APIEntityType() string {
	return "Document_ЧекККМ_Оплата"
}
func (p PrimaryDocumentChekKKMOplata) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentChekKKMOplataSertifikatami struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMOplataSertifikatami) APIEntityType() string {
	return "Document_ЧекККМ_ОплатаСертификатами"
}
func (p PrimaryDocumentChekKKMOplataSertifikatami) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentChekKKMProdazhaSertifikatov struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMProdazhaSertifikatov) APIEntityType() string {
	return "Document_ЧекККМ_ПродажаСертификатов"
}
func (p PrimaryDocumentChekKKMProdazhaSertifikatov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentChekKKMTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMTovary) APIEntityType() string {
	return "Document_ЧекККМ_Товары"
}
func (p PrimaryDocumentChekKKMTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentChekKKMDokumentyObmena struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMDokumentyObmena) APIEntityType() string {
	return "Document_ЧекККМ_ДокументыОбмена"
}
func (p PrimaryDocumentChekKKMDokumentyObmena) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentChekKKMDogovoraRassrochkiProdazha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMDogovoraRassrochkiProdazha) APIEntityType() string {
	return "Document_ЧекККМ_ДоговораРассрочкиПродажа"
}
func (p PrimaryDocumentChekKKMDogovoraRassrochkiProdazha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentChekKKMDogovoraRassrochkiOplata struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMDogovoraRassrochkiOplata) APIEntityType() string {
	return "Document_ЧекККМ_ДоговораРассрочкиОплата"
}
func (p PrimaryDocumentChekKKMDogovoraRassrochkiOplata) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentChekKKMOplataBallami struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMOplataBallami) APIEntityType() string {
	return "Document_ЧекККМ_ОплатаБаллами"
}
func (p PrimaryDocumentChekKKMOplataBallami) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentChekKKMSkidkiNatsenki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMSkidkiNatsenki) APIEntityType() string {
	return "Document_ЧекККМ_СкидкиНаценки"
}
func (p PrimaryDocumentChekKKMSkidkiNatsenki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentChekKKMUpravliaemyeSkidki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMUpravliaemyeSkidki) APIEntityType() string {
	return "Document_ЧекККМ_УправляемыеСкидки"
}
func (p PrimaryDocumentChekKKMUpravliaemyeSkidki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentChekKKMPodarki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMPodarki) APIEntityType() string {
	return "Document_ЧекККМ_Подарки"
}
func (p PrimaryDocumentChekKKMPodarki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentChekKKMKupony struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentChekKKMKupony) APIEntityType() string {
	return "Document_ЧекККМ_Купоны"
}
func (p PrimaryDocumentChekKKMKupony) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPereotsenkaValiutnykhSredstv struct {
	Key Guid
}

func (PrimaryDocumentPereotsenkaValiutnykhSredstv) APIEntityType() string {
	return "Document_ПереоценкаВалютныхСредств"
}
func (p PrimaryDocumentPereotsenkaValiutnykhSredstv) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogTipySkidokNatsenok struct {
	Key Guid
}

func (PrimaryCatalogTipySkidokNatsenok) APIEntityType() string {
	return "Catalog_ТипыСкидокНаценок"
}
func (p PrimaryCatalogTipySkidokNatsenok) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogTipySkidokNatsenokVremiaPoDniamNedeli struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogTipySkidokNatsenokVremiaPoDniamNedeli) APIEntityType() string {
	return "Catalog_ТипыСкидокНаценок_ВремяПоДнямНедели"
}
func (p PrimaryCatalogTipySkidokNatsenokVremiaPoDniamNedeli) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogVidyKodirovokiTsepei struct {
	Key Guid
}

func (PrimaryCatalogVidyKodirovokiTsepei) APIEntityType() string {
	return "Catalog_ВидыКодировокиЦепей"
}
func (p PrimaryCatalogVidyKodirovokiTsepei) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogVidyKodirovokiTsepeiElementyKodirovki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogVidyKodirovokiTsepeiElementyKodirovki) APIEntityType() string {
	return "Catalog_ВидыКодировокиЦепей_ЭлементыКодировки"
}
func (p PrimaryCatalogVidyKodirovokiTsepeiElementyKodirovki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv) APIEntityType() string {
	return "Catalog_ВидыКодировокиЦепей_СоответствиеЗначенийКодровкиЗначениямСвойств"
}
func (p PrimaryCatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetKomitentuOProdazhakh struct {
	Key Guid
}

func (PrimaryDocumentOtchetKomitentuOProdazhakh) APIEntityType() string {
	return "Document_ОтчетКомитентуОПродажах"
}
func (p PrimaryDocumentOtchetKomitentuOProdazhakh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva) APIEntityType() string {
	return "Document_ОтчетКомитентуОПродажах_ДенежныеСредства"
}
func (p PrimaryDocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetKomitentuOProdazhakhTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetKomitentuOProdazhakhTovary) APIEntityType() string {
	return "Document_ОтчетКомитентуОПродажах_Товары"
}
func (p PrimaryDocumentOtchetKomitentuOProdazhakhTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKassy struct {
	Key Guid
}

func (PrimaryCatalogKassy) APIEntityType() string {
	return "Catalog_Кассы"
}
func (p PrimaryCatalogKassy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogKassiry struct {
	Key Guid
}

func (PrimaryCatalogKassiry) APIEntityType() string {
	return "Catalog_Кассиры"
}
func (p PrimaryCatalogKassiry) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZaiavkaNaPereotsenkuTovarov struct {
	Key Guid
}

func (PrimaryDocumentZaiavkaNaPereotsenkuTovarov) APIEntityType() string {
	return "Document_ЗаявкаНаПереоценкуТоваров"
}
func (p PrimaryDocumentZaiavkaNaPereotsenkuTovarov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZaiavkaNaPereotsenkuTovarovTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentZaiavkaNaPereotsenkuTovarovTovary) APIEntityType() string {
	return "Document_ЗаявкаНаПереоценкуТоваров_Товары"
}
func (p PrimaryDocumentZaiavkaNaPereotsenkuTovarovTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogProizvodstvennyeUchastki struct {
	Key Guid
}

func (PrimaryCatalogProizvodstvennyeUchastki) APIEntityType() string {
	return "Catalog_ПроизводственныеУчастки"
}
func (p PrimaryCatalogProizvodstvennyeUchastki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZakrytieZakazovKlientov struct {
	Key Guid
}

func (PrimaryDocumentZakrytieZakazovKlientov) APIEntityType() string {
	return "Document_ЗакрытиеЗаказовКлиентов"
}
func (p PrimaryDocumentZakrytieZakazovKlientov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZakrytieZakazovKlientovZakazy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentZakrytieZakazovKlientovZakazy) APIEntityType() string {
	return "Document_ЗакрытиеЗаказовКлиентов_Заказы"
}
func (p PrimaryDocumentZakrytieZakazovKlientovZakazy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogProekty struct {
	Key Guid
}

func (PrimaryCatalogProekty) APIEntityType() string {
	return "Catalog_Проекты"
}
func (p PrimaryCatalogProekty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlatezhnoePoruchenieVkhodiashchee struct {
	Key Guid
}

func (PrimaryDocumentPlatezhnoePoruchenieVkhodiashchee) APIEntityType() string {
	return "Document_ПлатежноеПоручениеВходящее"
}
func (p PrimaryDocumentPlatezhnoePoruchenieVkhodiashchee) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_ПлатежноеПоручениеВходящее_РасшифровкаПлатежа"
}
func (p PrimaryDocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta) APIEntityType() string {
	return "Document_ПлатежноеПоручениеВходящее_РеквизитыКонтрагента"
}
func (p PrimaryDocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentVydachaZakaza struct {
	Key Guid
}

func (PrimaryDocumentVydachaZakaza) APIEntityType() string {
	return "Document_ВыдачаЗаказа"
}
func (p PrimaryDocumentVydachaZakaza) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogFormyOgranki struct {
	Key Guid
}

func (PrimaryCatalogFormyOgranki) APIEntityType() string {
	return "Catalog_ФормыОгранки"
}
func (p PrimaryCatalogFormyOgranki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogFormatyMagazinov struct {
	Key Guid
}

func (PrimaryCatalogFormatyMagazinov) APIEntityType() string {
	return "Catalog_ФорматыМагазинов"
}
func (p PrimaryCatalogFormatyMagazinov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogRabochieMesta struct {
	Key Guid
}

func (PrimaryCatalogRabochieMesta) APIEntityType() string {
	return "Catalog_РабочиеМеста"
}
func (p PrimaryCatalogRabochieMesta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikiVypolneniiaObmena struct {
	Key Guid
}

func (PrimaryCatalogNastroikiVypolneniiaObmena) APIEntityType() string {
	return "Catalog_НастройкиВыполненияОбмена"
}
func (p PrimaryCatalogNastroikiVypolneniiaObmena) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikiVypolneniiaObmenaNastroikiObmena struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikiVypolneniiaObmenaNastroikiObmena) APIEntityType() string {
	return "Catalog_НастройкиВыполненияОбмена_НастройкиОбмена"
}
func (p PrimaryCatalogNastroikiVypolneniiaObmenaNastroikiObmena) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami) APIEntityType() string {
	return "Catalog_НастройкиВыполненияОбмена_СообщенияНеЯвляющиесяОшибками"
}
func (p PrimaryCatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogZnacheniiaSvoistvObieektov struct {
	Key Guid
}

func (PrimaryCatalogZnacheniiaSvoistvObieektov) APIEntityType() string {
	return "Catalog_ЗначенияСвойствОбъектов"
}
func (p PrimaryCatalogZnacheniiaSvoistvObieektov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentRealizatsiiaTovarovUslug struct {
	Key Guid
}

func (PrimaryDocumentRealizatsiiaTovarovUslug) APIEntityType() string {
	return "Document_РеализацияТоваровУслуг"
}
func (p PrimaryDocumentRealizatsiiaTovarovUslug) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentRealizatsiiaTovarovUslugTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentRealizatsiiaTovarovUslugTovary) APIEntityType() string {
	return "Document_РеализацияТоваровУслуг_Товары"
}
func (p PrimaryDocumentRealizatsiiaTovarovUslugTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentRealizatsiiaTovarovUslugUslugi struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentRealizatsiiaTovarovUslugUslugi) APIEntityType() string {
	return "Document_РеализацияТоваровУслуг_Услуги"
}
func (p PrimaryDocumentRealizatsiiaTovarovUslugUslugi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentSobytie struct {
	Key Guid
}

func (PrimaryDocumentSobytie) APIEntityType() string {
	return "Document_Событие"
}
func (p PrimaryDocumentSobytie) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentSobytieStoronnieLitsa struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSobytieStoronnieLitsa) APIEntityType() string {
	return "Document_Событие_СторонниеЛица"
}
func (p PrimaryDocumentSobytieStoronnieLitsa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogVariantyOtvetovOprosov struct {
	Key Guid
}

func (PrimaryCatalogVariantyOtvetovOprosov) APIEntityType() string {
	return "Catalog_ВариантыОтветовОпросов"
}
func (p PrimaryCatalogVariantyOtvetovOprosov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogGruppyPisemElektronnoiPochty struct {
	Key Guid
}

func (PrimaryCatalogGruppyPisemElektronnoiPochty) APIEntityType() string {
	return "Catalog_ГруппыПисемЭлектроннойПочты"
}
func (p PrimaryCatalogGruppyPisemElektronnoiPochty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogGruppyPochtovoiRassylki struct {
	Key Guid
}

func (PrimaryCatalogGruppyPochtovoiRassylki) APIEntityType() string {
	return "Catalog_ГруппыПочтовойРассылки"
}
func (p PrimaryCatalogGruppyPochtovoiRassylki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikiOtchetov struct {
	Key Guid
}

func (PrimaryCatalogNastroikiOtchetov) APIEntityType() string {
	return "Catalog_НастройкиОтчетов"
}
func (p PrimaryCatalogNastroikiOtchetov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikiOtchetovGruppyNastroekIPolzovateli struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikiOtchetovGruppyNastroekIPolzovateli) APIEntityType() string {
	return "Catalog_НастройкиОтчетов_ГруппыНастроекИПользователи"
}
func (p PrimaryCatalogNastroikiOtchetovGruppyNastroekIPolzovateli) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam struct {
	Key Guid
}

func (PrimaryCatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam) APIEntityType() string {
	return "Catalog_СхемыНакопительныхСкидокПоДисконтнымКартам"
}
func (p PrimaryCatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki) APIEntityType() string {
	return "Catalog_СхемыНакопительныхСкидокПоДисконтнымКартам_Скидки"
}
func (p PrimaryCatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDepartment struct {
	Key Guid
}

func (PrimaryDepartment) APIEntityType() string {
	return "Catalog_Склады"
}
func (p PrimaryDepartment) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogKodyVidovTovarov struct {
	Key Guid
}

func (PrimaryCatalogKodyVidovTovarov) APIEntityType() string {
	return "Catalog_КодыВидовТоваров"
}
func (p PrimaryCatalogKodyVidovTovarov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogRassevy struct {
	Key Guid
}

func (PrimaryCatalogRassevy) APIEntityType() string {
	return "Catalog_Рассевы"
}
func (p PrimaryCatalogRassevy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogPrichinyZakrytiiaZakazov struct {
	Key Guid
}

func (PrimaryCatalogPrichinyZakrytiiaZakazov) APIEntityType() string {
	return "Catalog_ПричиныЗакрытияЗаказов"
}
func (p PrimaryCatalogPrichinyZakrytiiaZakazov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogSegmentyNomenklatury struct {
	Key Guid
}

func (PrimaryCatalogSegmentyNomenklatury) APIEntityType() string {
	return "Catalog_СегментыНоменклатуры"
}
func (p PrimaryCatalogSegmentyNomenklatury) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogSostavStrokiCheka struct {
	Key Guid
}

func (PrimaryCatalogSostavStrokiCheka) APIEntityType() string {
	return "Catalog_СоставСтрокиЧека"
}
func (p PrimaryCatalogSostavStrokiCheka) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogUsloviiaPriemaIzdeliiNaKomissiiu struct {
	Key Guid
}

func (PrimaryCatalogUsloviiaPriemaIzdeliiNaKomissiiu) APIEntityType() string {
	return "Catalog_УсловияПриемаИзделийНаКомиссию"
}
func (p PrimaryCatalogUsloviiaPriemaIzdeliiNaKomissiiu) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok) APIEntityType() string {
	return "Catalog_УсловияПриемаИзделийНаКомиссию_ГрафикУценок"
}
func (p PrimaryCatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei struct {
	Key Guid
}

func (PrimaryCatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei) APIEntityType() string {
	return "Catalog_ИсточникиИнформацииПриОбращенииПокупателей"
}
func (p PrimaryCatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentKorrektirovkaDolga struct {
	Key Guid
}

func (PrimaryDocumentKorrektirovkaDolga) APIEntityType() string {
	return "Document_КорректировкаДолга"
}
func (p PrimaryDocumentKorrektirovkaDolga) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentKorrektirovkaDolgaSummyDolga struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentKorrektirovkaDolgaSummyDolga) APIEntityType() string {
	return "Document_КорректировкаДолга_СуммыДолга"
}
func (p PrimaryDocumentKorrektirovkaDolgaSummyDolga) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryPayType struct {
	Key Guid
}

func (PrimaryPayType) APIEntityType() string {
	return "Catalog_ВидыОплатЧекаККМ"
}
func (p PrimaryPayType) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogKhranilishcheShablonov struct {
	Key Guid
}

func (PrimaryCatalogKhranilishcheShablonov) APIEntityType() string {
	return "Catalog_ХранилищеШаблонов"
}
func (p PrimaryCatalogKhranilishcheShablonov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZaiavkaNaRaskhodovanieSredstv struct {
	Key Guid
}

func (PrimaryDocumentZaiavkaNaRaskhodovanieSredstv) APIEntityType() string {
	return "Document_ЗаявкаНаРасходованиеСредств"
}
func (p PrimaryDocumentZaiavkaNaRaskhodovanieSredstv) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_ЗаявкаНаРасходованиеСредств_РасшифровкаПлатежа"
}
func (p PrimaryDocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki) APIEntityType() string {
	return "Document_ЗаявкаНаРасходованиеСредств_РазмещениеЗаявки"
}
func (p PrimaryDocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentZakrytieZakazovPostavshchikam struct {
	Key Guid
}

func (PrimaryDocumentZakrytieZakazovPostavshchikam) APIEntityType() string {
	return "Document_ЗакрытиеЗаказовПоставщикам"
}
func (p PrimaryDocumentZakrytieZakazovPostavshchikam) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZakrytieZakazovPostavshchikamZakazy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentZakrytieZakazovPostavshchikamZakazy) APIEntityType() string {
	return "Document_ЗакрытиеЗаказовПоставщикам_Заказы"
}
func (p PrimaryDocumentZakrytieZakazovPostavshchikamZakazy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogVidyKamnei struct {
	Key Guid
}

func (PrimaryCatalogVidyKamnei) APIEntityType() string {
	return "Catalog_ВидыКамней"
}
func (p PrimaryCatalogVidyKamnei) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentAnketyKlientovDliaFinMonitoringa struct {
	Key Guid
}

func (PrimaryDocumentAnketyKlientovDliaFinMonitoringa) APIEntityType() string {
	return "Document_АнкетыКлиентовДляФинМониторинга"
}
func (p PrimaryDocumentAnketyKlientovDliaFinMonitoringa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentAnketyKlientovDliaFinMonitoringaAnkety struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentAnketyKlientovDliaFinMonitoringaAnkety) APIEntityType() string {
	return "Document_АнкетыКлиентовДляФинМониторинга_Анкеты"
}
func (p PrimaryDocumentAnketyKlientovDliaFinMonitoringaAnkety) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogDogovoryRassrochki struct {
	Key Guid
}

func (PrimaryCatalogDogovoryRassrochki) APIEntityType() string {
	return "Catalog_ДоговорыРассрочки"
}
func (p PrimaryCatalogDogovoryRassrochki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogSertifikaty struct {
	Key Guid
}

func (PrimaryCatalogSertifikaty) APIEntityType() string {
	return "Catalog_Сертификаты"
}
func (p PrimaryCatalogSertifikaty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPostuplenieDavalcheskogoMetalla struct {
	Key Guid
}

func (PrimaryDocumentPostuplenieDavalcheskogoMetalla) APIEntityType() string {
	return "Document_ПоступлениеДавальческогоМеталла"
}
func (p PrimaryDocumentPostuplenieDavalcheskogoMetalla) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentInkassovoePorucheniePeredannoe struct {
	Key Guid
}

func (PrimaryDocumentInkassovoePorucheniePeredannoe) APIEntityType() string {
	return "Document_ИнкассовоеПоручениеПереданное"
}
func (p PrimaryDocumentInkassovoePorucheniePeredannoe) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_ИнкассовоеПоручениеПереданное_РасшифровкаПлатежа"
}
func (p PrimaryDocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentInkassovoePorucheniePeredannoeRekvizityKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentInkassovoePorucheniePeredannoeRekvizityKontragenta) APIEntityType() string {
	return "Document_ИнкассовоеПоручениеПереданное_РеквизитыКонтрагента"
}
func (p PrimaryDocumentInkassovoePorucheniePeredannoeRekvizityKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogFormulyDliaRascheta struct {
	Key Guid
}

func (PrimaryCatalogFormulyDliaRascheta) APIEntityType() string {
	return "Catalog_ФормулыДляРасчета"
}
func (p PrimaryCatalogFormulyDliaRascheta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogKupony struct {
	Key Guid
}

func (PrimaryCatalogKupony) APIEntityType() string {
	return "Catalog_Купоны"
}
func (p PrimaryCatalogKupony) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCorrecting struct {
	Key Guid
}

func (PrimaryCorrecting) APIEntityType() string {
	return "Document_КорректировкаЗаписейРегистровНакопления"
}
func (p PrimaryCorrecting) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia) APIEntityType() string {
	return "Document_КорректировкаЗаписейРегистровНакопления_ТаблицаРегистровНакопления"
}
func (p PrimaryDocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentInternetZakaz struct {
	Key Guid
}

func (PrimaryDocumentInternetZakaz) APIEntityType() string {
	return "Document_ИнтернетЗаказ"
}
func (p PrimaryDocumentInternetZakaz) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentInternetZakazTovaryInternetZakaza struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentInternetZakazTovaryInternetZakaza) APIEntityType() string {
	return "Document_ИнтернетЗаказ_ТоварыИнтернетЗаказа"
}
func (p PrimaryDocumentInternetZakazTovaryInternetZakaza) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentInternetZakazTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentInternetZakazTovary) APIEntityType() string {
	return "Document_ИнтернетЗаказ_Товары"
}
func (p PrimaryDocumentInternetZakazTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogRegiony struct {
	Key Guid
}

func (PrimaryCatalogRegiony) APIEntityType() string {
	return "Catalog_Регионы"
}
func (p PrimaryCatalogRegiony) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimarySaleJournal struct {
	Key Guid
}

func (PrimarySaleJournal) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах"
}
func (p PrimarySaleJournal) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhBonusy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhBonusy) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_Бонусы"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhBonusy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_ОплатаБанковскимиКредитами"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_ОплатаПлатежнымиКартами"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhOplataSertifikatami struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhOplataSertifikatami) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_ОплатаСертификатами"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhOplataSertifikatami) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_ПродажаСертификатов"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhTovary) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_Товары"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_ДоговораРассрочкиПродажа"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhDokumentyObmena struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhDokumentyObmena) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_ДокументыОбмена"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhDokumentyObmena) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_ДоговораРассрочкиОплата"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhOplataBallami struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhOplataBallami) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_ОплатаБаллами"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhOplataBallami) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_СкидкиНаценки"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetORoznichnykhProdazhakhKupony struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetORoznichnykhProdazhakhKupony) APIEntityType() string {
	return "Document_ОтчетОРозничныхПродажах_Купоны"
}
func (p PrimaryDocumentOtchetORoznichnykhProdazhakhKupony) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtmenaSkidokNomenklatury struct {
	Key Guid
}

func (PrimaryDocumentOtmenaSkidokNomenklatury) APIEntityType() string {
	return "Document_ОтменаСкидокНоменклатуры"
}
func (p PrimaryDocumentOtmenaSkidokNomenklatury) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentOtmenaSkidokNomenklaturyDokumenty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtmenaSkidokNomenklaturyDokumenty) APIEntityType() string {
	return "Document_ОтменаСкидокНоменклатуры_Документы"
}
func (p PrimaryDocumentOtmenaSkidokNomenklaturyDokumenty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogTovarnyeGruppy struct {
	Key Guid
}

func (PrimaryCatalogTovarnyeGruppy) APIEntityType() string {
	return "Catalog_ТоварныеГруппы"
}
func (p PrimaryCatalogTovarnyeGruppy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv struct {
	Key Guid
}

func (PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv) APIEntityType() string {
	return "Document_ПлатежныйОрдерПоступлениеДенежныхСредств"
}
func (p PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_ПлатежныйОрдерПоступлениеДенежныхСредств_РасшифровкаПлатежа"
}
func (p PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta) APIEntityType() string {
	return "Document_ПлатежныйОрдерПоступлениеДенежныхСредств_РеквизитыКонтрагента"
}
func (p PrimaryDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogOrderKey struct {
	Key Guid
}

func (PrimaryCatalogOrderKey) APIEntityType() string {
	return "Catalog_КлючиПродаж"
}
func (p PrimaryCatalogOrderKey) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentKassovyiChekKorrektsii struct {
	Key Guid
}

func (PrimaryDocumentKassovyiChekKorrektsii) APIEntityType() string {
	return "Document_КассовыйЧекКоррекции"
}
func (p PrimaryDocumentKassovyiChekKorrektsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentKassovyiChekKorrektsiiOplata struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentKassovyiChekKorrektsiiOplata) APIEntityType() string {
	return "Document_КассовыйЧекКоррекции_Оплата"
}
func (p PrimaryDocumentKassovyiChekKorrektsiiOplata) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentSchetNaOplatuPokupateliu struct {
	Key Guid
}

func (PrimaryDocumentSchetNaOplatuPokupateliu) APIEntityType() string {
	return "Document_СчетНаОплатуПокупателю"
}
func (p PrimaryDocumentSchetNaOplatuPokupateliu) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentSchetNaOplatuPokupateliuTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSchetNaOplatuPokupateliuTovary) APIEntityType() string {
	return "Document_СчетНаОплатуПокупателю_Товары"
}
func (p PrimaryDocumentSchetNaOplatuPokupateliuTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentSchetNaOplatuPokupateliuUslugi struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSchetNaOplatuPokupateliuUslugi) APIEntityType() string {
	return "Document_СчетНаОплатуПокупателю_Услуги"
}
func (p PrimaryDocumentSchetNaOplatuPokupateliuUslugi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogNastroikiObmenaDannymi struct {
	Key Guid
}

func (PrimaryCatalogNastroikiObmenaDannymi) APIEntityType() string {
	return "Catalog_НастройкиОбменаДанными"
}
func (p PrimaryCatalogNastroikiObmenaDannymi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov) APIEntityType() string {
	return "Catalog_НастройкиОбменаДанными_НастройкаВариантовПоискаОбъектов"
}
func (p PrimaryCatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh) APIEntityType() string {
	return "Catalog_НастройкиОбменаДанными_НастройкаВыгрузкиДанных"
}
func (p PrimaryCatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami) APIEntityType() string {
	return "Catalog_НастройкиОбменаДанными_СообщенияНеЯвляющиесяОшибками"
}
func (p PrimaryCatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentJournalBankovskieRaschetnyeDokumenty struct {
	Ref     String
	RefType String
}

func (PrimaryDocumentJournalBankovskieRaschetnyeDokumenty) APIEntityType() string {
	return "DocumentJournal_БанковскиеРасчетныеДокументы"
}
func (p PrimaryDocumentJournalBankovskieRaschetnyeDokumenty) Serialize() string {
	return "Ref=" + p.Ref.AsParameter() + ",Ref_Type=" + p.RefType.AsParameter()
}

type PrimaryDocumentZamenaDiskontnoiKarty struct {
	Key Guid
}

func (PrimaryDocumentZamenaDiskontnoiKarty) APIEntityType() string {
	return "Document_ЗаменаДисконтнойКарты"
}
func (p PrimaryDocumentZamenaDiskontnoiKarty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryReturnToSupplier struct {
	Key Guid
}

func (PrimaryReturnToSupplier) APIEntityType() string {
	return "Document_ВозвратТоваровПоставщику"
}
func (p PrimaryReturnToSupplier) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVozvratTovarovPostavshchikuTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVozvratTovarovPostavshchikuTovary) APIEntityType() string {
	return "Document_ВозвратТоваровПоставщику_Товары"
}
func (p PrimaryDocumentVozvratTovarovPostavshchikuTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentInventarizatsiiaTovarovNaSklade struct {
	Key Guid
}

func (PrimaryDocumentInventarizatsiiaTovarovNaSklade) APIEntityType() string {
	return "Document_ИнвентаризацияТоваровНаСкладе"
}
func (p PrimaryDocumentInventarizatsiiaTovarovNaSklade) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentInventarizatsiiaTovarovNaSkladeTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentInventarizatsiiaTovarovNaSkladeTovary) APIEntityType() string {
	return "Document_ИнвентаризацияТоваровНаСкладе_Товары"
}
func (p PrimaryDocumentInventarizatsiiaTovarovNaSkladeTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii) APIEntityType() string {
	return "Document_ИнвентаризацияТоваровНаСкладе_УсловияПроведенияИнвентаризации"
}
func (p PrimaryDocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentInventarizatsiiaTovarovNaSkladeSertifikaty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentInventarizatsiiaTovarovNaSkladeSertifikaty) APIEntityType() string {
	return "Document_ИнвентаризацияТоваровНаСкладе_Сертификаты"
}
func (p PrimaryDocumentInventarizatsiiaTovarovNaSkladeSertifikaty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti) APIEntityType() string {
	return "Document_ИнвентаризацияТоваровНаСкладе_ТоварыВПути"
}
func (p PrimaryDocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPrikhodnyiKassovyiOrder struct {
	Key Guid
}

func (PrimaryDocumentPrikhodnyiKassovyiOrder) APIEntityType() string {
	return "Document_ПриходныйКассовыйОрдер"
}
func (p PrimaryDocumentPrikhodnyiKassovyiOrder) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_ПриходныйКассовыйОрдер_РасшифровкаПлатежа"
}
func (p PrimaryDocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPrikhodnyiKassovyiOrderOplata struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPrikhodnyiKassovyiOrderOplata) APIEntityType() string {
	return "Document_ПриходныйКассовыйОрдер_Оплата"
}
func (p PrimaryDocumentPrikhodnyiKassovyiOrderOplata) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPrikhodnyiKassovyiOrderTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPrikhodnyiKassovyiOrderTovary) APIEntityType() string {
	return "Document_ПриходныйКассовыйОрдер_Товары"
}
func (p PrimaryDocumentPrikhodnyiKassovyiOrderTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogPrichinyVozvrata struct {
	Key Guid
}

func (PrimaryCatalogPrichinyVozvrata) APIEntityType() string {
	return "Catalog_ПричиныВозврата"
}
func (p PrimaryCatalogPrichinyVozvrata) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentDenezhnyiChek struct {
	Key Guid
}

func (PrimaryDocumentDenezhnyiChek) APIEntityType() string {
	return "Document_ДенежныйЧек"
}
func (p PrimaryDocumentDenezhnyiChek) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVozvratMaterialovIzProizvodstva struct {
	Key Guid
}

func (PrimaryDocumentVozvratMaterialovIzProizvodstva) APIEntityType() string {
	return "Document_ВозвратМатериаловИзПроизводства"
}
func (p PrimaryDocumentVozvratMaterialovIzProizvodstva) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVozvratMaterialovIzProizvodstvaTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVozvratMaterialovIzProizvodstvaTovary) APIEntityType() string {
	return "Document_ВозвратМатериаловИзПроизводства_Товары"
}
func (p PrimaryDocumentVozvratMaterialovIzProizvodstvaTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPereotsenkaTovarovOtdannykhNaKomissiiu struct {
	Key Guid
}

func (PrimaryDocumentPereotsenkaTovarovOtdannykhNaKomissiiu) APIEntityType() string {
	return "Document_ПереоценкаТоваровОтданныхНаКомиссию"
}
func (p PrimaryDocumentPereotsenkaTovarovOtdannykhNaKomissiiu) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary) APIEntityType() string {
	return "Document_ПереоценкаТоваровОтданныхНаКомиссию_Товары"
}
func (p PrimaryDocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSN struct {
	Key Guid
}

func (PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSN) APIEntityType() string {
	return "Document_ВводНачальныхОстатковПоРасходамУСН"
}
func (p PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSN) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami) APIEntityType() string {
	return "Document_ВводНачальныхОстатковПоРасходамУСН_ВзаиморасчетыСПокупателями"
}
func (p PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye) APIEntityType() string {
	return "Document_ВводНачальныхОстатковПоРасходамУСН_ТоварыПроданные"
}
func (p PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami) APIEntityType() string {
	return "Document_ВводНачальныхОстатковПоРасходамУСН_ВзаиморасчетыСПоставщиками"
}
func (p PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh) APIEntityType() string {
	return "Document_ВводНачальныхОстатковПоРасходамУСН_ТоварыНаОстатках"
}
func (p PrimaryDocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentGTDImport struct {
	Key Guid
}

func (PrimaryDocumentGTDImport) APIEntityType() string {
	return "Document_ГТДИмпорт"
}
func (p PrimaryDocumentGTDImport) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentGTDImportRazdely struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentGTDImportRazdely) APIEntityType() string {
	return "Document_ГТДИмпорт_Разделы"
}
func (p PrimaryDocumentGTDImportRazdely) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentGTDImportTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentGTDImportTovary) APIEntityType() string {
	return "Document_ГТДИмпорт_Товары"
}
func (p PrimaryDocumentGTDImportTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentAktSverki struct {
	Key Guid
}

func (PrimaryDocumentAktSverki) APIEntityType() string {
	return "Document_АктСверки"
}
func (p PrimaryDocumentAktSverki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentAktSverkiTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentAktSverkiTovary) APIEntityType() string {
	return "Document_АктСверки_Товары"
}
func (p PrimaryDocumentAktSverkiTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogFaily struct {
	Key Guid
}

func (PrimaryCatalogFaily) APIEntityType() string {
	return "Catalog_Файлы"
}
func (p PrimaryCatalogFaily) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogFailyDopolnitelnyeRekvizity struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogFailyDopolnitelnyeRekvizity) APIEntityType() string {
	return "Catalog_Файлы_ДополнительныеРеквизиты"
}
func (p PrimaryCatalogFailyDopolnitelnyeRekvizity) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogFailySertifikatyShifrovaniia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogFailySertifikatyShifrovaniia) APIEntityType() string {
	return "Catalog_Файлы_СертификатыШифрования"
}
func (p PrimaryCatalogFailySertifikatyShifrovaniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogUchetnyeZapisiElektronnoiPochty struct {
	Key Guid
}

func (PrimaryCatalogUchetnyeZapisiElektronnoiPochty) APIEntityType() string {
	return "Catalog_УчетныеЗаписиЭлектроннойПочты"
}
func (p PrimaryCatalogUchetnyeZapisiElektronnoiPochty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi) APIEntityType() string {
	return "Catalog_УчетныеЗаписиЭлектроннойПочты_ДоступКУчетнойЗаписи"
}
func (p PrimaryCatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPlaniruemoePostuplenieDenezhnykhSredstv struct {
	Key Guid
}

func (PrimaryDocumentPlaniruemoePostuplenieDenezhnykhSredstv) APIEntityType() string {
	return "Document_ПланируемоеПоступлениеДенежныхСредств"
}
func (p PrimaryDocumentPlaniruemoePostuplenieDenezhnykhSredstv) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_ПланируемоеПоступлениеДенежныхСредств_РасшифровкаПлатежа"
}
func (p PrimaryDocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPreiskurantTsenNaKamni struct {
	Key Guid
}

func (PrimaryDocumentPreiskurantTsenNaKamni) APIEntityType() string {
	return "Document_ПрейскурантЦенНаКамни"
}
func (p PrimaryDocumentPreiskurantTsenNaKamni) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryPurchase struct {
	Key Guid
}

func (PrimaryPurchase) APIEntityType() string {
	return "Document_СкупкаТоваров"
}
func (p PrimaryPurchase) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentSkupkaTovarovTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSkupkaTovarovTovary) APIEntityType() string {
	return "Document_СкупкаТоваров_Товары"
}
func (p PrimaryDocumentSkupkaTovarovTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentSchetFakturaPoluchennyi struct {
	Key Guid
}

func (PrimaryDocumentSchetFakturaPoluchennyi) APIEntityType() string {
	return "Document_СчетФактураПолученный"
}
func (p PrimaryDocumentSchetFakturaPoluchennyi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam) APIEntityType() string {
	return "Document_СчетФактураПолученный_СчетаФактурыВыданныеПокупателям"
}
func (p PrimaryDocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentAktKhimicheskogoAnalizaMetalla struct {
	Key Guid
}

func (PrimaryDocumentAktKhimicheskogoAnalizaMetalla) APIEntityType() string {
	return "Document_АктХимическогоАнализаМеталла"
}
func (p PrimaryDocumentAktKhimicheskogoAnalizaMetalla) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogfmKartochkaKontragenta struct {
	Key Guid
}

func (PrimaryCatalogfmKartochkaKontragenta) APIEntityType() string {
	return "Catalog_фмКарточкаКонтрагента"
}
func (p PrimaryCatalogfmKartochkaKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogfmKartochkaKontragentaDannyeKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogfmKartochkaKontragentaDannyeKontragenta) APIEntityType() string {
	return "Catalog_фмКарточкаКонтрагента_ДанныеКонтрагента"
}
func (p PrimaryCatalogfmKartochkaKontragentaDannyeKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentSpisanieProsrochennykhSertifikatov struct {
	Key Guid
}

func (PrimaryDocumentSpisanieProsrochennykhSertifikatov) APIEntityType() string {
	return "Document_СписаниеПросроченныхСертификатов"
}
func (p PrimaryDocumentSpisanieProsrochennykhSertifikatov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentSpisanieProsrochennykhSertifikatovSertifikaty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSpisanieProsrochennykhSertifikatovSertifikaty) APIEntityType() string {
	return "Document_СписаниеПросроченныхСертификатов_Сертификаты"
}
func (p PrimaryDocumentSpisanieProsrochennykhSertifikatovSertifikaty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentZakrytieAvansovPoGrafikuPlatezhei struct {
	Key Guid
}

func (PrimaryDocumentZakrytieAvansovPoGrafikuPlatezhei) APIEntityType() string {
	return "Document_ЗакрытиеАвансовПоГрафикуПлатежей"
}
func (p PrimaryDocumentZakrytieAvansovPoGrafikuPlatezhei) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty) APIEntityType() string {
	return "Document_ЗакрытиеАвансовПоГрафикуПлатежей_Контрагенты"
}
func (p PrimaryDocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogTipySistemNalogooblozheniiaKKT struct {
	Key Guid
}

func (PrimaryCatalogTipySistemNalogooblozheniiaKKT) APIEntityType() string {
	return "Catalog_ТипыСистемНалогообложенияККТ"
}
func (p PrimaryCatalogTipySistemNalogooblozheniiaKKT) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentAkkreditivPeredannyi struct {
	Key Guid
}

func (PrimaryDocumentAkkreditivPeredannyi) APIEntityType() string {
	return "Document_АккредитивПереданный"
}
func (p PrimaryDocumentAkkreditivPeredannyi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentAkkreditivPeredannyiRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentAkkreditivPeredannyiRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_АккредитивПереданный_РасшифровкаПлатежа"
}
func (p PrimaryDocumentAkkreditivPeredannyiRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentAkkreditivPeredannyiRekvizityKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentAkkreditivPeredannyiRekvizityKontragenta) APIEntityType() string {
	return "Document_АккредитивПереданный_РеквизитыКонтрагента"
}
func (p PrimaryDocumentAkkreditivPeredannyiRekvizityKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimarySupplier struct {
	Key Guid
}

func (PrimarySupplier) APIEntityType() string {
	return "Catalog_Контрагенты"
}
func (p PrimarySupplier) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogKontragentyVidyDeiatelnosti struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogKontragentyVidyDeiatelnosti) APIEntityType() string {
	return "Catalog_Контрагенты_ВидыДеятельности"
}
func (p PrimaryCatalogKontragentyVidyDeiatelnosti) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentInformatsionnoeSoobshchenie struct {
	Key Guid
}

func (PrimaryDocumentInformatsionnoeSoobshchenie) APIEntityType() string {
	return "Document_ИнформационноеСообщение"
}
func (p PrimaryDocumentInformatsionnoeSoobshchenie) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogVlozheniiaElektronnykhPisem struct {
	Key Guid
}

func (PrimaryCatalogVlozheniiaElektronnykhPisem) APIEntityType() string {
	return "Catalog_ВложенияЭлектронныхПисем"
}
func (p PrimaryCatalogVlozheniiaElektronnykhPisem) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlatezhnoeTrebovanieVystavlennoe struct {
	Key Guid
}

func (PrimaryDocumentPlatezhnoeTrebovanieVystavlennoe) APIEntityType() string {
	return "Document_ПлатежноеТребованиеВыставленное"
}
func (p PrimaryDocumentPlatezhnoeTrebovanieVystavlennoe) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_ПлатежноеТребованиеВыставленное_РасшифровкаПлатежа"
}
func (p PrimaryDocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta) APIEntityType() string {
	return "Document_ПлатежноеТребованиеВыставленное_РеквизитыКонтрагента"
}
func (p PrimaryDocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentMarketingovaiaAktsiia struct {
	Key Guid
}

func (PrimaryDocumentMarketingovaiaAktsiia) APIEntityType() string {
	return "Document_МаркетинговаяАкция"
}
func (p PrimaryDocumentMarketingovaiaAktsiia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentMarketingovaiaAktsiiaMagaziny struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentMarketingovaiaAktsiiaMagaziny) APIEntityType() string {
	return "Document_МаркетинговаяАкция_Магазины"
}
func (p PrimaryDocumentMarketingovaiaAktsiiaMagaziny) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentMarketingovaiaAktsiiaSkidkiNatsenki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentMarketingovaiaAktsiiaSkidkiNatsenki) APIEntityType() string {
	return "Document_МаркетинговаяАкция_СкидкиНаценки"
}
func (p PrimaryDocumentMarketingovaiaAktsiiaSkidkiNatsenki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa) APIEntityType() string {
	return "Document_МаркетинговаяАкция_НаборыЗначенийДоступа"
}
func (p PrimaryDocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogStsenariiObmenovDannymi struct {
	Key Guid
}

func (PrimaryCatalogStsenariiObmenovDannymi) APIEntityType() string {
	return "Catalog_СценарииОбменовДанными"
}
func (p PrimaryCatalogStsenariiObmenovDannymi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogStsenariiObmenovDannymiNastroikiObmena struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogStsenariiObmenovDannymiNastroikiObmena) APIEntityType() string {
	return "Catalog_СценарииОбменовДанными_НастройкиОбмена"
}
func (p PrimaryCatalogStsenariiObmenovDannymiNastroikiObmena) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryItem struct {
	Key Guid
}

func (PrimaryItem) APIEntityType() string {
	return "Catalog_Номенклатура"
}
func (p PrimaryItem) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNomenklaturaSostavLigatury struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNomenklaturaSostavLigatury) APIEntityType() string {
	return "Catalog_Номенклатура_СоставЛигатуры"
}
func (p PrimaryCatalogNomenklaturaSostavLigatury) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOpros struct {
	Key Guid
}

func (PrimaryDocumentOpros) APIEntityType() string {
	return "Document_Опрос"
}
func (p PrimaryDocumentOpros) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentOprosVoprosy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOprosVoprosy) APIEntityType() string {
	return "Document_Опрос_Вопросы"
}
func (p PrimaryDocumentOprosVoprosy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOprosSostavnoiOtvet struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOprosSostavnoiOtvet) APIEntityType() string {
	return "Document_Опрос_СоставнойОтвет"
}
func (p PrimaryDocumentOprosSostavnoiOtvet) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogGruppyPoluchateleiSkidki struct {
	Key Guid
}

func (PrimaryCatalogGruppyPoluchateleiSkidki) APIEntityType() string {
	return "Catalog_ГруппыПолучателейСкидки"
}
func (p PrimaryCatalogGruppyPoluchateleiSkidki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryReassessment struct {
	Key Guid
}

func (PrimaryReassessment) APIEntityType() string {
	return "Document_ПереоценкаТоваровВНТТ"
}
func (p PrimaryReassessment) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPereotsenkaTovarovVNTTTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPereotsenkaTovarovVNTTTovary) APIEntityType() string {
	return "Document_ПереоценкаТоваровВНТТ_Товары"
}
func (p PrimaryDocumentPereotsenkaTovarovVNTTTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogTomaKhraneniiaFailov struct {
	Key Guid
}

func (PrimaryCatalogTomaKhraneniiaFailov) APIEntityType() string {
	return "Catalog_ТомаХраненияФайлов"
}
func (p PrimaryCatalogTomaKhraneniiaFailov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentJournalProizvodstvennyeDokumenty struct {
	Ref     String
	RefType String
}

func (PrimaryDocumentJournalProizvodstvennyeDokumenty) APIEntityType() string {
	return "DocumentJournal_ПроизводственныеДокументы"
}
func (p PrimaryDocumentJournalProizvodstvennyeDokumenty) Serialize() string {
	return "Ref=" + p.Ref.AsParameter() + ",Ref_Type=" + p.RefType.AsParameter()
}

type PrimaryDocumentIzmeneniePravDostupa struct {
	Key Guid
}

func (PrimaryDocumentIzmeneniePravDostupa) APIEntityType() string {
	return "Document_ИзменениеПравДоступа"
}
func (p PrimaryDocumentIzmeneniePravDostupa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikaAssortimentnoiMatritsy struct {
	Key Guid
}

func (PrimaryCatalogNastroikaAssortimentnoiMatritsy) APIEntityType() string {
	return "Catalog_НастройкаАссортиментнойМатрицы"
}
func (p PrimaryCatalogNastroikaAssortimentnoiMatritsy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp) APIEntityType() string {
	return "Catalog_НастройкаАссортиментнойМатрицы_НастройкаТоварныхГрупп"
}
func (p PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii) APIEntityType() string {
	return "Catalog_НастройкаАссортиментнойМатрицы_НастройкаТоварныхКатегорий"
}
func (p PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii) APIEntityType() string {
	return "Catalog_НастройкаАссортиментнойМатрицы_НастройкаТоварныхПозиций"
}
func (p PrimaryCatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentJournalDokumentyKontragentov struct {
	Ref     String
	RefType String
}

func (PrimaryDocumentJournalDokumentyKontragentov) APIEntityType() string {
	return "DocumentJournal_ДокументыКонтрагентов"
}
func (p PrimaryDocumentJournalDokumentyKontragentov) Serialize() string {
	return "Ref=" + p.Ref.AsParameter() + ",Ref_Type=" + p.RefType.AsParameter()
}

type PrimaryMoveInstance struct {
	Key Guid
}

func (PrimaryMoveInstance) APIEntityType() string {
	return "Document_ПеремещениеТоваров"
}
func (p PrimaryMoveInstance) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPeremeshchenieTovarovSertifikaty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPeremeshchenieTovarovSertifikaty) APIEntityType() string {
	return "Document_ПеремещениеТоваров_Сертификаты"
}
func (p PrimaryDocumentPeremeshchenieTovarovSertifikaty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPeremeshchenieTovarovTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPeremeshchenieTovarovTovary) APIEntityType() string {
	return "Document_ПеремещениеТоваров_Товары"
}
func (p PrimaryDocumentPeremeshchenieTovarovTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPeremeshchenieTovarovSpisokZaiavok struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPeremeshchenieTovarovSpisokZaiavok) APIEntityType() string {
	return "Document_ПеремещениеТоваров_СписокЗаявок"
}
func (p PrimaryDocumentPeremeshchenieTovarovSpisokZaiavok) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentZakrytieZaiavokNaRaskhodovanieSredstv struct {
	Key Guid
}

func (PrimaryDocumentZakrytieZaiavokNaRaskhodovanieSredstv) APIEntityType() string {
	return "Document_ЗакрытиеЗаявокНаРасходованиеСредств"
}
func (p PrimaryDocumentZakrytieZaiavokNaRaskhodovanieSredstv) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv) APIEntityType() string {
	return "Document_ЗакрытиеЗаявокНаРасходованиеСредств_ЗаявкиНаРасходованиеСредств"
}
func (p PrimaryDocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryMemberCard struct {
	Key Guid
}

func (PrimaryMemberCard) APIEntityType() string {
	return "Catalog_ДисконтныеКарты"
}
func (p PrimaryMemberCard) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentABCKlassifikatsiiaPokupatelei struct {
	Key Guid
}

func (PrimaryDocumentABCKlassifikatsiiaPokupatelei) APIEntityType() string {
	return "Document_ABCКлассификацияПокупателей"
}
func (p PrimaryDocumentABCKlassifikatsiiaPokupatelei) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov) APIEntityType() string {
	return "Document_ABCКлассификацияПокупателей_ТаблицаРаспределенияКонтрагентов"
}
func (p PrimaryDocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogIdentifikatoryObieektovMetadannykh struct {
	Key Guid
}

func (PrimaryCatalogIdentifikatoryObieektovMetadannykh) APIEntityType() string {
	return "Catalog_ИдентификаторыОбъектовМетаданных"
}
func (p PrimaryCatalogIdentifikatoryObieektovMetadannykh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSklade struct {
	Key Guid
}

func (PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSklade) APIEntityType() string {
	return "Document_СводнаяИнвентаризацияТоваровНаСкладе"
}
func (p PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSklade) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty) APIEntityType() string {
	return "Document_СводнаяИнвентаризацияТоваровНаСкладе_Сертификаты"
}
func (p PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii) APIEntityType() string {
	return "Document_СводнаяИнвентаризацияТоваровНаСкладе_УсловияПроведенияИнвентаризации"
}
func (p PrimaryDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentKorrektirovkaRealizatsii struct {
	Key Guid
}

func (PrimaryDocumentKorrektirovkaRealizatsii) APIEntityType() string {
	return "Document_КорректировкаРеализации"
}
func (p PrimaryDocumentKorrektirovkaRealizatsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentKorrektirovkaRealizatsiiTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentKorrektirovkaRealizatsiiTovary) APIEntityType() string {
	return "Document_КорректировкаРеализации_Товары"
}
func (p PrimaryDocumentKorrektirovkaRealizatsiiTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentKorrektirovkaRealizatsiiUslugi struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentKorrektirovkaRealizatsiiUslugi) APIEntityType() string {
	return "Document_КорректировкаРеализации_Услуги"
}
func (p PrimaryDocumentKorrektirovkaRealizatsiiUslugi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogVidyDefektov struct {
	Key Guid
}

func (PrimaryCatalogVidyDefektov) APIEntityType() string {
	return "Catalog_ВидыДефектов"
}
func (p PrimaryCatalogVidyDefektov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentDoverennost struct {
	Key Guid
}

func (PrimaryDocumentDoverennost) APIEntityType() string {
	return "Document_Доверенность"
}
func (p PrimaryDocumentDoverennost) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentDoverennostTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentDoverennostTovary) APIEntityType() string {
	return "Document_Доверенность_Товары"
}
func (p PrimaryDocumentDoverennostTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogShablonyZapolneniiaKU struct {
	Key Guid
}

func (PrimaryCatalogShablonyZapolneniiaKU) APIEntityType() string {
	return "Catalog_ШаблоныЗаполненияКУ"
}
func (p PrimaryCatalogShablonyZapolneniiaKU) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogShablonyZapolneniiaKUPrazdnichnyeDni struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogShablonyZapolneniiaKUPrazdnichnyeDni) APIEntityType() string {
	return "Catalog_ШаблоныЗаполненияКУ_ПраздничныеДни"
}
func (p PrimaryCatalogShablonyZapolneniiaKUPrazdnichnyeDni) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogShablonyZapolneniiaKUKUNaNedeliu struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogShablonyZapolneniiaKUKUNaNedeliu) APIEntityType() string {
	return "Catalog_ШаблоныЗаполненияКУ_КУНаНеделю"
}
func (p PrimaryCatalogShablonyZapolneniiaKUKUNaNedeliu) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogShablonyZapolneniiaKUSalony struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogShablonyZapolneniiaKUSalony) APIEntityType() string {
	return "Catalog_ШаблоныЗаполненияКУ_Салоны"
}
func (p PrimaryCatalogShablonyZapolneniiaKUSalony) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPlanZapolneniiaVitrin struct {
	Key Guid
}

func (PrimaryDocumentPlanZapolneniiaVitrin) APIEntityType() string {
	return "Document_ПланЗаполненияВитрин"
}
func (p PrimaryDocumentPlanZapolneniiaVitrin) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin) APIEntityType() string {
	return "Document_ПланЗаполненияВитрин_ПлановоеЗаполнениеВитрин"
}
func (p PrimaryDocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryInstance struct {
	Key Guid
}

func (PrimaryInstance) APIEntityType() string {
	return "Catalog_СерииНоменклатуры"
}
func (p PrimaryInstance) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryReturnToManufacturing struct {
	Key Guid
}

func (PrimaryReturnToManufacturing) APIEntityType() string {
	return "Document_ВозвратПродукцииВПроизводство"
}
func (p PrimaryReturnToManufacturing) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVozvratProduktsiiVProizvodstvoTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVozvratProduktsiiVProizvodstvoTovary) APIEntityType() string {
	return "Document_ВозвратПродукцииВПроизводство_Товары"
}
func (p PrimaryDocumentVozvratProduktsiiVProizvodstvoTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogNomeraGTD struct {
	Key Guid
}

func (PrimaryCatalogNomeraGTD) APIEntityType() string {
	return "Catalog_НомераГТД"
}
func (p PrimaryCatalogNomeraGTD) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikiRabochegoMestaPolzovatelia struct {
	Key Guid
}

func (PrimaryCatalogNastroikiRabochegoMestaPolzovatelia) APIEntityType() string {
	return "Catalog_НастройкиРабочегоМестаПользователя"
}
func (p PrimaryCatalogNastroikiRabochegoMestaPolzovatelia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikiRabochegoMestaPolzovateliaNastroiki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikiRabochegoMestaPolzovateliaNastroiki) APIEntityType() string {
	return "Catalog_НастройкиРабочегоМестаПользователя_Настройки"
}
func (p PrimaryCatalogNastroikiRabochegoMestaPolzovateliaNastroiki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogsmsShablony struct {
	Key Guid
}

func (PrimaryCatalogsmsShablony) APIEntityType() string {
	return "Catalog_смсШаблоны"
}
func (p PrimaryCatalogsmsShablony) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryWriteOff struct {
	Key Guid
}

func (PrimaryWriteOff) APIEntityType() string {
	return "Document_СписаниеТоваров"
}
func (p PrimaryWriteOff) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentSpisanieTovarovTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSpisanieTovarovTovary) APIEntityType() string {
	return "Document_СписаниеТоваров_Товары"
}
func (p PrimaryDocumentSpisanieTovarovTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentSpisanieTovarovSertifikaty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSpisanieTovarovSertifikaty) APIEntityType() string {
	return "Document_СписаниеТоваров_Сертификаты"
}
func (p PrimaryDocumentSpisanieTovarovSertifikaty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentsmsSoobshchenie struct {
	Key Guid
}

func (PrimaryDocumentsmsSoobshchenie) APIEntityType() string {
	return "Document_смсСообщение"
}
func (p PrimaryDocumentsmsSoobshchenie) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentsmsSoobshcheniePoluchateli struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentsmsSoobshcheniePoluchateli) APIEntityType() string {
	return "Document_смсСообщение_Получатели"
}
func (p PrimaryDocumentsmsSoobshcheniePoluchateli) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOplataOtPokupateliaPlatezhnoiKartoi struct {
	Key Guid
}

func (PrimaryDocumentOplataOtPokupateliaPlatezhnoiKartoi) APIEntityType() string {
	return "Document_ОплатаОтПокупателяПлатежнойКартой"
}
func (p PrimaryDocumentOplataOtPokupateliaPlatezhnoiKartoi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogDragotsennyeKamni struct {
	Key Guid
}

func (PrimaryCatalogDragotsennyeKamni) APIEntityType() string {
	return "Catalog_ДрагоценныеКамни"
}
func (p PrimaryCatalogDragotsennyeKamni) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogKalendariPlanirovaniiaProdazh struct {
	Key Guid
}

func (PrimaryCatalogKalendariPlanirovaniiaProdazh) APIEntityType() string {
	return "Catalog_КалендариПланированияПродаж"
}
func (p PrimaryCatalogKalendariPlanirovaniiaProdazh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogKalendariPlanirovaniiaProdazhKUPoDniam struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogKalendariPlanirovaniiaProdazhKUPoDniam) APIEntityType() string {
	return "Catalog_КалендариПланированияПродаж_КУПоДням"
}
func (p PrimaryCatalogKalendariPlanirovaniiaProdazhKUPoDniam) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKalendariPlanirovaniiaProdazhSalony struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogKalendariPlanirovaniiaProdazhSalony) APIEntityType() string {
	return "Catalog_КалендариПланированияПродаж_Салоны"
}
func (p PrimaryCatalogKalendariPlanirovaniiaProdazhSalony) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKontaktnyeLitsa struct {
	Key Guid
}

func (PrimaryCatalogKontaktnyeLitsa) APIEntityType() string {
	return "Catalog_КонтактныеЛица"
}
func (p PrimaryCatalogKontaktnyeLitsa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogFizicheskieLitsa struct {
	Key Guid
}

func (PrimaryCatalogFizicheskieLitsa) APIEntityType() string {
	return "Catalog_ФизическиеЛица"
}
func (p PrimaryCatalogFizicheskieLitsa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogTipovyeAnkety struct {
	Key Guid
}

func (PrimaryCatalogTipovyeAnkety) APIEntityType() string {
	return "Catalog_ТиповыеАнкеты"
}
func (p PrimaryCatalogTipovyeAnkety) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogTipovyeAnketyVoprosyAnkety struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogTipovyeAnketyVoprosyAnkety) APIEntityType() string {
	return "Catalog_ТиповыеАнкеты_ВопросыАнкеты"
}
func (p PrimaryCatalogTipovyeAnketyVoprosyAnkety) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentNachislenieSpisanieBonusov struct {
	Key Guid
}

func (PrimaryDocumentNachislenieSpisanieBonusov) APIEntityType() string {
	return "Document_НачислениеСписаниеБонусов"
}
func (p PrimaryDocumentNachislenieSpisanieBonusov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentNachislenieSpisanieBonusovDiskontnyeKarty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentNachislenieSpisanieBonusovDiskontnyeKarty) APIEntityType() string {
	return "Document_НачислениеСписаниеБонусов_ДисконтныеКарты"
}
func (p PrimaryDocumentNachislenieSpisanieBonusovDiskontnyeKarty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryType struct {
	Key Guid
}

func (PrimaryType) APIEntityType() string {
	return "Catalog_ТипыИзделий"
}
func (p PrimaryType) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogfmKodyVidovDokumentov struct {
	Key Guid
}

func (PrimaryCatalogfmKodyVidovDokumentov) APIEntityType() string {
	return "Catalog_фмКодыВидовДокументов"
}
func (p PrimaryCatalogfmKodyVidovDokumentov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlatezhnoeTrebovaniePoluchennoe struct {
	Key Guid
}

func (PrimaryDocumentPlatezhnoeTrebovaniePoluchennoe) APIEntityType() string {
	return "Document_ПлатежноеТребованиеПолученное"
}
func (p PrimaryDocumentPlatezhnoeTrebovaniePoluchennoe) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_ПлатежноеТребованиеПолученное_РасшифровкаПлатежа"
}
func (p PrimaryDocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta) APIEntityType() string {
	return "Document_ПлатежноеТребованиеПолученное_РеквизитыКонтрагента"
}
func (p PrimaryDocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv struct {
	Key Guid
}

func (PrimaryDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv) APIEntityType() string {
	return "Document_ЗакрытиеПланируемыхПоступленийДенежныхСредств"
}
func (p PrimaryDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS) APIEntityType() string {
	return "Document_ЗакрытиеПланируемыхПоступленийДенежныхСредств_ПланируемыеПоступленияДС"
}
func (p PrimaryDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogRazdelyAnkety struct {
	Key Guid
}

func (PrimaryCatalogRazdelyAnkety) APIEntityType() string {
	return "Catalog_РазделыАнкеты"
}
func (p PrimaryCatalogRazdelyAnkety) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentOtchetPoFinMonitoringu struct {
	Key Guid
}

func (PrimaryDocumentOtchetPoFinMonitoringu) APIEntityType() string {
	return "Document_ОтчетПоФинМониторингу"
}
func (p PrimaryDocumentOtchetPoFinMonitoringu) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa) APIEntityType() string {
	return "Document_ОтчетПоФинМониторингу_ДокументыФинМониторинга"
}
func (p PrimaryDocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetPoFinMonitoringuDannyeDokumenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetPoFinMonitoringuDannyeDokumenta) APIEntityType() string {
	return "Document_ОтчетПоФинМониторингу_ДанныеДокумента"
}
func (p PrimaryDocumentOtchetPoFinMonitoringuDannyeDokumenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKliuchiAnalitikiUchetaNomenklatury struct {
	Key Guid
}

func (PrimaryCatalogKliuchiAnalitikiUchetaNomenklatury) APIEntityType() string {
	return "Catalog_КлючиАналитикиУчетаНоменклатуры"
}
func (p PrimaryCatalogKliuchiAnalitikiUchetaNomenklatury) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogVersiiFailov struct {
	Key Guid
}

func (PrimaryCatalogVersiiFailov) APIEntityType() string {
	return "Catalog_ВерсииФайлов"
}
func (p PrimaryCatalogVersiiFailov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogVersiiFailovElektronnyePodpisi struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogVersiiFailovElektronnyePodpisi) APIEntityType() string {
	return "Catalog_ВерсииФайлов_ЭлектронныеПодписи"
}
func (p PrimaryCatalogVersiiFailovElektronnyePodpisi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentUstanovkaTsenNomenklatury struct {
	Key Guid
}

func (PrimaryDocumentUstanovkaTsenNomenklatury) APIEntityType() string {
	return "Document_УстановкаЦенНоменклатуры"
}
func (p PrimaryDocumentUstanovkaTsenNomenklatury) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentUstanovkaTsenNomenklaturyTipyTsen struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUstanovkaTsenNomenklaturyTipyTsen) APIEntityType() string {
	return "Document_УстановкаЦенНоменклатуры_ТипыЦен"
}
func (p PrimaryDocumentUstanovkaTsenNomenklaturyTipyTsen) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentUstanovkaTsenNomenklaturyTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUstanovkaTsenNomenklaturyTovary) APIEntityType() string {
	return "Document_УстановкаЦенНоменклатуры_Товары"
}
func (p PrimaryDocumentUstanovkaTsenNomenklaturyTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv struct {
	Key Guid
}

func (PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv) APIEntityType() string {
	return "Document_ПлатежныйОрдерСписаниеДенежныхСредств"
}
func (p PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_ПлатежныйОрдерСписаниеДенежныхСредств_РасшифровкаПлатежа"
}
func (p PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta) APIEntityType() string {
	return "Document_ПлатежныйОрдерСписаниеДенежныхСредств_РеквизитыКонтрагента"
}
func (p PrimaryDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPreiskurantNaSkupku struct {
	Key Guid
}

func (PrimaryDocumentPreiskurantNaSkupku) APIEntityType() string {
	return "Document_ПрейскурантНаСкупку"
}
func (p PrimaryDocumentPreiskurantNaSkupku) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPreiskurantNaSkupkuProby struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPreiskurantNaSkupkuProby) APIEntityType() string {
	return "Document_ПрейскурантНаСкупку_Пробы"
}
func (p PrimaryDocumentPreiskurantNaSkupkuProby) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPeredachaMaterialovVProizvodstvo struct {
	Key Guid
}

func (PrimaryDocumentPeredachaMaterialovVProizvodstvo) APIEntityType() string {
	return "Document_ПередачаМатериаловВПроизводство"
}
func (p PrimaryDocumentPeredachaMaterialovVProizvodstvo) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPeredachaMaterialovVProizvodstvoTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPeredachaMaterialovVProizvodstvoTovary) APIEntityType() string {
	return "Document_ПередачаМатериаловВПроизводство_Товары"
}
func (p PrimaryDocumentPeredachaMaterialovVProizvodstvoTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentVnutrenniiZakaz struct {
	Key Guid
}

func (PrimaryDocumentVnutrenniiZakaz) APIEntityType() string {
	return "Document_ВнутреннийЗаказ"
}
func (p PrimaryDocumentVnutrenniiZakaz) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVnutrenniiZakazTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVnutrenniiZakazTovary) APIEntityType() string {
	return "Document_ВнутреннийЗаказ_Товары"
}
func (p PrimaryDocumentVnutrenniiZakazTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKhranilishcheDopolnitelnoiInformatsii struct {
	Key Guid
}

func (PrimaryCatalogKhranilishcheDopolnitelnoiInformatsii) APIEntityType() string {
	return "Catalog_ХранилищеДополнительнойИнформации"
}
func (p PrimaryCatalogKhranilishcheDopolnitelnoiInformatsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogDopolnitelnyeVneshnieObrabotki struct {
	Key Guid
}

func (PrimaryCatalogDopolnitelnyeVneshnieObrabotki) APIEntityType() string {
	return "Catalog_ДополнительныеВнешниеОбработки"
}
func (p PrimaryCatalogDopolnitelnyeVneshnieObrabotki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost) APIEntityType() string {
	return "Catalog_ДополнительныеВнешниеОбработки_Принадлежность"
}
func (p PrimaryCatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogDopolnitelnyeVneshnieObrabotkiKomandy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogDopolnitelnyeVneshnieObrabotkiKomandy) APIEntityType() string {
	return "Catalog_ДополнительныеВнешниеОбработки_Команды"
}
func (p PrimaryCatalogDopolnitelnyeVneshnieObrabotkiKomandy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogDopolnitelnyeVneshnieObrabotkiRazdely struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogDopolnitelnyeVneshnieObrabotkiRazdely) APIEntityType() string {
	return "Catalog_ДополнительныеВнешниеОбработки_Разделы"
}
func (p PrimaryCatalogDopolnitelnyeVneshnieObrabotkiRazdely) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogDopolnitelnyeVneshnieObrabotkiNaznachenie struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogDopolnitelnyeVneshnieObrabotkiNaznachenie) APIEntityType() string {
	return "Catalog_ДополнительныеВнешниеОбработки_Назначение"
}
func (p PrimaryCatalogDopolnitelnyeVneshnieObrabotkiNaznachenie) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogDopolnitelnyeVneshnieObrabotkiRazresheniia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogDopolnitelnyeVneshnieObrabotkiRazresheniia) APIEntityType() string {
	return "Catalog_ДополнительныеВнешниеОбработки_Разрешения"
}
func (p PrimaryCatalogDopolnitelnyeVneshnieObrabotkiRazresheniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogGruppyPolzovatelei struct {
	Key Guid
}

func (PrimaryCatalogGruppyPolzovatelei) APIEntityType() string {
	return "Catalog_ГруппыПользователей"
}
func (p PrimaryCatalogGruppyPolzovatelei) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogGruppyPolzovateleiPolzovateliGruppy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogGruppyPolzovateleiPolzovateliGruppy) APIEntityType() string {
	return "Catalog_ГруппыПользователей_ПользователиГруппы"
}
func (p PrimaryCatalogGruppyPolzovateleiPolzovateliGruppy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentJournalZakazyKlientov struct {
	Ref     String
	RefType String
}

func (PrimaryDocumentJournalZakazyKlientov) APIEntityType() string {
	return "DocumentJournal_ЗаказыКлиентов"
}
func (p PrimaryDocumentJournalZakazyKlientov) Serialize() string {
	return "Ref=" + p.Ref.AsParameter() + ",Ref_Type=" + p.RefType.AsParameter()
}

type PrimaryDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki struct {
	Key Guid
}

func (PrimaryDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki) APIEntityType() string {
	return "Document_ВозвратТоваровПоставщикуИзНеавтоматизированнойТорговойТочки"
}
func (p PrimaryDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary) APIEntityType() string {
	return "Document_ВозвратТоваровПоставщикуИзНеавтоматизированнойТорговойТочки_Товары"
}
func (p PrimaryDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentZaiavkaNaPeremeshchenieTovarov struct {
	Key Guid
}

func (PrimaryDocumentZaiavkaNaPeremeshchenieTovarov) APIEntityType() string {
	return "Document_ЗаявкаНаПеремещениеТоваров"
}
func (p PrimaryDocumentZaiavkaNaPeremeshchenieTovarov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZaiavkaNaPeremeshchenieTovarovTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentZaiavkaNaPeremeshchenieTovarovTovary) APIEntityType() string {
	return "Document_ЗаявкаНаПеремещениеТоваров_Товары"
}
func (p PrimaryDocumentZaiavkaNaPeremeshchenieTovarovTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogUsloviiaProdazh struct {
	Key Guid
}

func (PrimaryCatalogUsloviiaProdazh) APIEntityType() string {
	return "Catalog_УсловияПродаж"
}
func (p PrimaryCatalogUsloviiaProdazh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVvodNachalnykhOstatkovPoFinMonitoringu struct {
	Key Guid
}

func (PrimaryDocumentVvodNachalnykhOstatkovPoFinMonitoringu) APIEntityType() string {
	return "Document_ВводНачальныхОстатковПоФинМониторингу"
}
func (p PrimaryDocumentVvodNachalnykhOstatkovPoFinMonitoringu) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora) APIEntityType() string {
	return "Document_ВводНачальныхОстатковПоФинМониторингу_Договора"
}
func (p PrimaryDocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogOrganizatsii struct {
	Key Guid
}

func (PrimaryCatalogOrganizatsii) APIEntityType() string {
	return "Catalog_Организации"
}
func (p PrimaryCatalogOrganizatsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogUsloviiaOplaty struct {
	Key Guid
}

func (PrimaryCatalogUsloviiaOplaty) APIEntityType() string {
	return "Catalog_УсловияОплаты"
}
func (p PrimaryCatalogUsloviiaOplaty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogUsloviiaOplatyTablitsaVyplat struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogUsloviiaOplatyTablitsaVyplat) APIEntityType() string {
	return "Catalog_УсловияОплаты_ТаблицаВыплат"
}
func (p PrimaryCatalogUsloviiaOplatyTablitsaVyplat) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKategoriiObieektov struct {
	Key Guid
}

func (PrimaryCatalogKategoriiObieektov) APIEntityType() string {
	return "Catalog_КатегорииОбъектов"
}
func (p PrimaryCatalogKategoriiObieektov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogfmZnacheniiaDliaZapolneniia struct {
	Key Guid
}

func (PrimaryCatalogfmZnacheniiaDliaZapolneniia) APIEntityType() string {
	return "Catalog_фмЗначенияДляЗаполнения"
}
func (p PrimaryCatalogfmZnacheniiaDliaZapolneniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentUdalitNariadZakaz struct {
	Key Guid
}

func (PrimaryDocumentUdalitNariadZakaz) APIEntityType() string {
	return "Document_УдалитьНарядЗаказ"
}
func (p PrimaryDocumentUdalitNariadZakaz) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentUdalitNariadZakazIzdeliia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUdalitNariadZakazIzdeliia) APIEntityType() string {
	return "Document_УдалитьНарядЗаказ_Изделия"
}
func (p PrimaryDocumentUdalitNariadZakazIzdeliia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentUdalitNariadZakazUslugi struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUdalitNariadZakazUslugi) APIEntityType() string {
	return "Document_УдалитьНарядЗаказ_Услуги"
}
func (p PrimaryDocumentUdalitNariadZakazUslugi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentUdalitNariadZakazSpetsifikatsiia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUdalitNariadZakazSpetsifikatsiia) APIEntityType() string {
	return "Document_УдалитьНарядЗаказ_Спецификация"
}
func (p PrimaryDocumentUdalitNariadZakazSpetsifikatsiia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentUdalitNariadZakazMetally struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUdalitNariadZakazMetally) APIEntityType() string {
	return "Document_УдалитьНарядЗаказ_Металлы"
}
func (p PrimaryDocumentUdalitNariadZakazMetally) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentUdalitNariadZakazVstavki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUdalitNariadZakazVstavki) APIEntityType() string {
	return "Document_УдалитьНарядЗаказ_Вставки"
}
func (p PrimaryDocumentUdalitNariadZakazVstavki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogBanki struct {
	Key Guid
}

func (PrimaryCatalogBanki) APIEntityType() string {
	return "Catalog_Банки"
}
func (p PrimaryCatalogBanki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogRoliKontaktnykhLits struct {
	Key Guid
}

func (PrimaryCatalogRoliKontaktnykhLits) APIEntityType() string {
	return "Catalog_РолиКонтактныхЛиц"
}
func (p PrimaryCatalogRoliKontaktnykhLits) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentRestrukturizatsiiaDolga struct {
	Key Guid
}

func (PrimaryDocumentRestrukturizatsiiaDolga) APIEntityType() string {
	return "Document_РеструктуризацияДолга"
}
func (p PrimaryDocumentRestrukturizatsiiaDolga) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti) APIEntityType() string {
	return "Document_РеструктуризацияДолга_РасшифровкаЗадолженности"
}
func (p PrimaryDocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentAkkreditivPoluchennyi struct {
	Key Guid
}

func (PrimaryDocumentAkkreditivPoluchennyi) APIEntityType() string {
	return "Document_АккредитивПолученный"
}
func (p PrimaryDocumentAkkreditivPoluchennyi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentAkkreditivPoluchennyiRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentAkkreditivPoluchennyiRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_АккредитивПолученный_РасшифровкаПлатежа"
}
func (p PrimaryDocumentAkkreditivPoluchennyiRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentAkkreditivPoluchennyiRekvizityKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentAkkreditivPoluchennyiRekvizityKontragenta) APIEntityType() string {
	return "Document_АккредитивПолученный_РеквизитыКонтрагента"
}
func (p PrimaryDocumentAkkreditivPoluchennyiRekvizityKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPriemIzRemonta struct {
	Key Guid
}

func (PrimaryDocumentPriemIzRemonta) APIEntityType() string {
	return "Document_ПриемИзРемонта"
}
func (p PrimaryDocumentPriemIzRemonta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPriemIzRemontaIzdeliia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPriemIzRemontaIzdeliia) APIEntityType() string {
	return "Document_ПриемИзРемонта_Изделия"
}
func (p PrimaryDocumentPriemIzRemontaIzdeliia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPriemIzRemontaMaterialy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPriemIzRemontaMaterialy) APIEntityType() string {
	return "Document_ПриемИзРемонта_Материалы"
}
func (p PrimaryDocumentPriemIzRemontaMaterialy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogTsveta struct {
	Key Guid
}

func (PrimaryCatalogTsveta) APIEntityType() string {
	return "Catalog_Цвета"
}
func (p PrimaryCatalogTsveta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakh struct {
	Key Guid
}

func (PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakh) APIEntityType() string {
	return "Document_СторнированиеОтчетаКомиссионераОПродажах"
}
func (p PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva) APIEntityType() string {
	return "Document_СторнированиеОтчетаКомиссионераОПродажах_ДенежныеСредства"
}
func (p PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary) APIEntityType() string {
	return "Document_СторнированиеОтчетаКомиссионераОПродажах_Товары"
}
func (p PrimaryDocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentJournalDavalcheskieDokumenty struct {
	Ref     String
	RefType String
}

func (PrimaryDocumentJournalDavalcheskieDokumenty) APIEntityType() string {
	return "DocumentJournal_ДавальческиеДокументы"
}
func (p PrimaryDocumentJournalDavalcheskieDokumenty) Serialize() string {
	return "Ref=" + p.Ref.AsParameter() + ",Ref_Type=" + p.RefType.AsParameter()
}

type PrimaryCatalogfmAnketaKlienta struct {
	Key Guid
}

func (PrimaryCatalogfmAnketaKlienta) APIEntityType() string {
	return "Catalog_фмАнкетаКлиента"
}
func (p PrimaryCatalogfmAnketaKlienta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogfmAnketaKlientaDannyeKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogfmAnketaKlientaDannyeKontragenta) APIEntityType() string {
	return "Catalog_фмАнкетаКлиента_ДанныеКонтрагента"
}
func (p PrimaryCatalogfmAnketaKlientaDannyeKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogVidyVzaimoraschetov struct {
	Key Guid
}

func (PrimaryCatalogVidyVzaimoraschetov) APIEntityType() string {
	return "Catalog_ВидыВзаиморасчетов"
}
func (p PrimaryCatalogVidyVzaimoraschetov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentUstanovkaZnacheniiTochkiZakaza struct {
	Key Guid
}

func (PrimaryDocumentUstanovkaZnacheniiTochkiZakaza) APIEntityType() string {
	return "Document_УстановкаЗначенийТочкиЗаказа"
}
func (p PrimaryDocumentUstanovkaZnacheniiTochkiZakaza) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentUstanovkaZnacheniiTochkiZakazaTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUstanovkaZnacheniiTochkiZakazaTovary) APIEntityType() string {
	return "Document_УстановкаЗначенийТочкиЗаказа_Товары"
}
func (p PrimaryDocumentUstanovkaZnacheniiTochkiZakazaTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogZnacheniiaKodirovki struct {
	Key Guid
}

func (PrimaryCatalogZnacheniiaKodirovki) APIEntityType() string {
	return "Catalog_ЗначенияКодировки"
}
func (p PrimaryCatalogZnacheniiaKodirovki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogPravilaProdazh struct {
	Key Guid
}

func (PrimaryCatalogPravilaProdazh) APIEntityType() string {
	return "Catalog_ПравилаПродаж"
}
func (p PrimaryCatalogPravilaProdazh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogPravilaProdazhTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogPravilaProdazhTovary) APIEntityType() string {
	return "Catalog_ПравилаПродаж_Товары"
}
func (p PrimaryCatalogPravilaProdazhTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPostuplenieDopRaskhodov struct {
	Key Guid
}

func (PrimaryDocumentPostuplenieDopRaskhodov) APIEntityType() string {
	return "Document_ПоступлениеДопРасходов"
}
func (p PrimaryDocumentPostuplenieDopRaskhodov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPostuplenieDopRaskhodovTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPostuplenieDopRaskhodovTovary) APIEntityType() string {
	return "Document_ПоступлениеДопРасходов_Товары"
}
func (p PrimaryDocumentPostuplenieDopRaskhodovTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKhoziaistvennyeOperatsii struct {
	Key Guid
}

func (PrimaryCatalogKhoziaistvennyeOperatsii) APIEntityType() string {
	return "Catalog_ХозяйственныеОперации"
}
func (p PrimaryCatalogKhoziaistvennyeOperatsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentAvansovyiOtchet struct {
	Key Guid
}

func (PrimaryDocumentAvansovyiOtchet) APIEntityType() string {
	return "Document_АвансовыйОтчет"
}
func (p PrimaryDocumentAvansovyiOtchet) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentAvansovyiOtchetVydannyeAvansy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentAvansovyiOtchetVydannyeAvansy) APIEntityType() string {
	return "Document_АвансовыйОтчет_ВыданныеАвансы"
}
func (p PrimaryDocumentAvansovyiOtchetVydannyeAvansy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentAvansovyiOtchetTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentAvansovyiOtchetTovary) APIEntityType() string {
	return "Document_АвансовыйОтчет_Товары"
}
func (p PrimaryDocumentAvansovyiOtchetTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentAvansovyiOtchetOplataPostavshchikam struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentAvansovyiOtchetOplataPostavshchikam) APIEntityType() string {
	return "Document_АвансовыйОтчет_ОплатаПоставщикам"
}
func (p PrimaryDocumentAvansovyiOtchetOplataPostavshchikam) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentAvansovyiOtchetProchee struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentAvansovyiOtchetProchee) APIEntityType() string {
	return "Document_АвансовыйОтчет_Прочее"
}
func (p PrimaryDocumentAvansovyiOtchetProchee) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogDolzhnostiOrganizatsii struct {
	Key Guid
}

func (PrimaryCatalogDolzhnostiOrganizatsii) APIEntityType() string {
	return "Catalog_ДолжностиОрганизаций"
}
func (p PrimaryCatalogDolzhnostiOrganizatsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogAnalitikaTipaIzdeliia struct {
	Key Guid
}

func (PrimaryCatalogAnalitikaTipaIzdeliia) APIEntityType() string {
	return "Catalog_АналитикаТипаИзделия"
}
func (p PrimaryCatalogAnalitikaTipaIzdeliia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogDopolnitelnyePechatnyeFormy struct {
	Key Guid
}

func (PrimaryCatalogDopolnitelnyePechatnyeFormy) APIEntityType() string {
	return "Catalog_ДополнительныеПечатныеФормы"
}
func (p PrimaryCatalogDopolnitelnyePechatnyeFormy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogDopolnitelnyePechatnyeFormyPrinadlezhnost struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogDopolnitelnyePechatnyeFormyPrinadlezhnost) APIEntityType() string {
	return "Catalog_ДополнительныеПечатныеФормы_Принадлежность"
}
func (p PrimaryCatalogDopolnitelnyePechatnyeFormyPrinadlezhnost) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryMemberCardsType struct {
	Key Guid
}

func (PrimaryMemberCardsType) APIEntityType() string {
	return "Catalog_ВидыДисконтныхКарт"
}
func (p PrimaryMemberCardsType) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentRegistratsiiaNaSaite struct {
	Key Guid
}

func (PrimaryDocumentRegistratsiiaNaSaite) APIEntityType() string {
	return "Document_РегистрацияНаСайте"
}
func (p PrimaryDocumentRegistratsiiaNaSaite) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogObrabotkiObsluzhivaniiaTO struct {
	Key Guid
}

func (PrimaryCatalogObrabotkiObsluzhivaniiaTO) APIEntityType() string {
	return "Catalog_ОбработкиОбслуживанияТО"
}
func (p PrimaryCatalogObrabotkiObsluzhivaniiaTO) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogObrabotkiObsluzhivaniiaTOModeli struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogObrabotkiObsluzhivaniiaTOModeli) APIEntityType() string {
	return "Catalog_ОбработкиОбслуживанияТО_Модели"
}
func (p PrimaryCatalogObrabotkiObsluzhivaniiaTOModeli) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogNastroikaIntervalov struct {
	Key Guid
}

func (PrimaryCatalogNastroikaIntervalov) APIEntityType() string {
	return "Catalog_НастройкаИнтервалов"
}
func (p PrimaryCatalogNastroikaIntervalov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikaIntervalovTablichnaiaChast struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikaIntervalovTablichnaiaChast) APIEntityType() string {
	return "Catalog_НастройкаИнтервалов_ТабличнаяЧасть"
}
func (p PrimaryCatalogNastroikaIntervalovTablichnaiaChast) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogProfiliGruppDostupa struct {
	Key Guid
}

func (PrimaryCatalogProfiliGruppDostupa) APIEntityType() string {
	return "Catalog_ПрофилиГруппДоступа"
}
func (p PrimaryCatalogProfiliGruppDostupa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogProfiliGruppDostupaRoli struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogProfiliGruppDostupaRoli) APIEntityType() string {
	return "Catalog_ПрофилиГруппДоступа_Роли"
}
func (p PrimaryCatalogProfiliGruppDostupaRoli) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogProfiliGruppDostupaVidyDostupa struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogProfiliGruppDostupaVidyDostupa) APIEntityType() string {
	return "Catalog_ПрофилиГруппДоступа_ВидыДоступа"
}
func (p PrimaryCatalogProfiliGruppDostupaVidyDostupa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogProfiliGruppDostupaZnacheniiaDostupa struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogProfiliGruppDostupaZnacheniiaDostupa) APIEntityType() string {
	return "Catalog_ПрофилиГруппДоступа_ЗначенияДоступа"
}
func (p PrimaryCatalogProfiliGruppDostupaZnacheniiaDostupa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogProfiliGruppDostupaDostupPoPodsistemam struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogProfiliGruppDostupaDostupPoPodsistemam) APIEntityType() string {
	return "Catalog_ПрофилиГруппДоступа_ДоступПоПодсистемам"
}
func (p PrimaryCatalogProfiliGruppDostupaDostupPoPodsistemam) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogNastroikiDliaKurera struct {
	Key Guid
}

func (PrimaryCatalogNastroikiDliaKurera) APIEntityType() string {
	return "Catalog_НастройкиДляКурьера"
}
func (p PrimaryCatalogNastroikiDliaKurera) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikiDliaKureraSostavNaimenovaniia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikiDliaKureraSostavNaimenovaniia) APIEntityType() string {
	return "Catalog_НастройкиДляКурьера_СоставНаименования"
}
func (p PrimaryCatalogNastroikiDliaKureraSostavNaimenovaniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogTipyTsenNomenklaturyKontragentov struct {
	Key Guid
}

func (PrimaryCatalogTipyTsenNomenklaturyKontragentov) APIEntityType() string {
	return "Catalog_ТипыЦенНоменклатурыКонтрагентов"
}
func (p PrimaryCatalogTipyTsenNomenklaturyKontragentov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentJournalTsenoobrazovanie struct {
	Ref     String
	RefType String
}

func (PrimaryDocumentJournalTsenoobrazovanie) APIEntityType() string {
	return "DocumentJournal_Ценообразование"
}
func (p PrimaryDocumentJournalTsenoobrazovanie) Serialize() string {
	return "Ref=" + p.Ref.AsParameter() + ",Ref_Type=" + p.RefType.AsParameter()
}

type PrimaryCatalogEdinitsyIzmereniia struct {
	Key Guid
}

func (PrimaryCatalogEdinitsyIzmereniia) APIEntityType() string {
	return "Catalog_ЕдиницыИзмерения"
}
func (p PrimaryCatalogEdinitsyIzmereniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogStatiDvizheniiaDenezhnykhSredstv struct {
	Key Guid
}

func (PrimaryCatalogStatiDvizheniiaDenezhnykhSredstv) APIEntityType() string {
	return "Catalog_СтатьиДвиженияДенежныхСредств"
}
func (p PrimaryCatalogStatiDvizheniiaDenezhnykhSredstv) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentInkassovoePorucheniePoluchennoe struct {
	Key Guid
}

func (PrimaryDocumentInkassovoePorucheniePoluchennoe) APIEntityType() string {
	return "Document_ИнкассовоеПоручениеПолученное"
}
func (p PrimaryDocumentInkassovoePorucheniePoluchennoe) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_ИнкассовоеПоручениеПолученное_РасшифровкаПлатежа"
}
func (p PrimaryDocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta) APIEntityType() string {
	return "Document_ИнкассовоеПоручениеПолученное_РеквизитыКонтрагента"
}
func (p PrimaryDocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogNastroikiObmenaDannymiShtrikhM struct {
	Key Guid
}

func (PrimaryCatalogNastroikiObmenaDannymiShtrikhM) APIEntityType() string {
	return "Catalog_НастройкиОбменаДаннымиШтрихМ"
}
func (p PrimaryCatalogNastroikiObmenaDannymiShtrikhM) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogStatiZatrat struct {
	Key Guid
}

func (PrimaryCatalogStatiZatrat) APIEntityType() string {
	return "Catalog_СтатьиЗатрат"
}
func (p PrimaryCatalogStatiZatrat) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVozvratTovarovOtPokupatelia struct {
	Key Guid
}

func (PrimaryDocumentVozvratTovarovOtPokupatelia) APIEntityType() string {
	return "Document_ВозвратТоваровОтПокупателя"
}
func (p PrimaryDocumentVozvratTovarovOtPokupatelia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVozvratTovarovOtPokupateliaTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVozvratTovarovOtPokupateliaTovary) APIEntityType() string {
	return "Document_ВозвратТоваровОтПокупателя_Товары"
}
func (p PrimaryDocumentVozvratTovarovOtPokupateliaTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentVozvratTovarovOtPokupateliaUslugi struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVozvratTovarovOtPokupateliaUslugi) APIEntityType() string {
	return "Document_ВозвратТоваровОтПокупателя_Услуги"
}
func (p PrimaryDocumentVozvratTovarovOtPokupateliaUslugi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentZakazPostavshchiku struct {
	Key Guid
}

func (PrimaryDocumentZakazPostavshchiku) APIEntityType() string {
	return "Document_ЗаказПоставщику"
}
func (p PrimaryDocumentZakazPostavshchiku) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZakazPostavshchikuTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentZakazPostavshchikuTovary) APIEntityType() string {
	return "Document_ЗаказПоставщику_Товары"
}
func (p PrimaryDocumentZakazPostavshchikuTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogSkidkiNatsenki struct {
	Key Guid
}

func (PrimaryCatalogSkidkiNatsenki) APIEntityType() string {
	return "Catalog_СкидкиНаценки"
}
func (p PrimaryCatalogSkidkiNatsenki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogSkidkiNatsenkiUsloviiaPredostavleniia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogSkidkiNatsenkiUsloviiaPredostavleniia) APIEntityType() string {
	return "Catalog_СкидкиНаценки_УсловияПредоставления"
}
func (p PrimaryCatalogSkidkiNatsenkiUsloviiaPredostavleniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogSkidkiNatsenkiTsenovyeGruppy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogSkidkiNatsenkiTsenovyeGruppy) APIEntityType() string {
	return "Catalog_СкидкиНаценки_ЦеновыеГруппы"
}
func (p PrimaryCatalogSkidkiNatsenkiTsenovyeGruppy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogSkidkiNatsenkiNaborPodarkov struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogSkidkiNatsenkiNaborPodarkov) APIEntityType() string {
	return "Catalog_СкидкиНаценки_НаборПодарков"
}
func (p PrimaryCatalogSkidkiNatsenkiNaborPodarkov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogGruppyTsvetov struct {
	Key Guid
}

func (PrimaryCatalogGruppyTsvetov) APIEntityType() string {
	return "Catalog_ГруппыЦветов"
}
func (p PrimaryCatalogGruppyTsvetov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentDokumentRaschetovSKontragentom struct {
	Key Guid
}

func (PrimaryDocumentDokumentRaschetovSKontragentom) APIEntityType() string {
	return "Document_ДокументРасчетовСКонтрагентом"
}
func (p PrimaryDocumentDokumentRaschetovSKontragentom) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogDogovoryEkvairinga struct {
	Key Guid
}

func (PrimaryCatalogDogovoryEkvairinga) APIEntityType() string {
	return "Catalog_ДоговорыЭквайринга"
}
func (p PrimaryCatalogDogovoryEkvairinga) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie) APIEntityType() string {
	return "Catalog_ДоговорыЭквайринга_ТарифыЗаРасчетноеОбслуживание"
}
func (p PrimaryCatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKachestvo struct {
	Key Guid
}

func (PrimaryCatalogKachestvo) APIEntityType() string {
	return "Catalog_Качество"
}
func (p PrimaryCatalogKachestvo) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentUstanovkaTsenNomenklaturyKontragentov struct {
	Key Guid
}

func (PrimaryDocumentUstanovkaTsenNomenklaturyKontragentov) APIEntityType() string {
	return "Document_УстановкаЦенНоменклатурыКонтрагентов"
}
func (p PrimaryDocumentUstanovkaTsenNomenklaturyKontragentov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen) APIEntityType() string {
	return "Document_УстановкаЦенНоменклатурыКонтрагентов_ТипыЦен"
}
func (p PrimaryDocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentUstanovkaTsenNomenklaturyKontragentovTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUstanovkaTsenNomenklaturyKontragentovTovary) APIEntityType() string {
	return "Document_УстановкаЦенНоменклатурыКонтрагентов_Товары"
}
func (p PrimaryDocumentUstanovkaTsenNomenklaturyKontragentovTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentProtsentPoterPoDavaltsam struct {
	Key Guid
}

func (PrimaryDocumentProtsentPoterPoDavaltsam) APIEntityType() string {
	return "Document_ПроцентПотерьПоДавальцам"
}
func (p PrimaryDocumentProtsentPoterPoDavaltsam) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentProtsentPoterPoDavaltsamProtsenty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentProtsentPoterPoDavaltsamProtsenty) APIEntityType() string {
	return "Document_ПроцентПотерьПоДавальцам_Проценты"
}
func (p PrimaryDocumentProtsentPoterPoDavaltsamProtsenty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogTovarnyePozitsii struct {
	Key Guid
}

func (PrimaryCatalogTovarnyePozitsii) APIEntityType() string {
	return "Catalog_ТоварныеПозиции"
}
func (p PrimaryCatalogTovarnyePozitsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlatezhnoePoruchenieIskhodiashchee struct {
	Key Guid
}

func (PrimaryDocumentPlatezhnoePoruchenieIskhodiashchee) APIEntityType() string {
	return "Document_ПлатежноеПоручениеИсходящее"
}
func (p PrimaryDocumentPlatezhnoePoruchenieIskhodiashchee) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_ПлатежноеПоручениеИсходящее_РасшифровкаПлатежа"
}
func (p PrimaryDocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta) APIEntityType() string {
	return "Document_ПлатежноеПоручениеИсходящее_РеквизитыКонтрагента"
}
func (p PrimaryDocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogfmOrganizatsionnoPravovyeFormy struct {
	Key Guid
}

func (PrimaryCatalogfmOrganizatsionnoPravovyeFormy) APIEntityType() string {
	return "Catalog_фмОрганизационноПравовыеФормы"
}
func (p PrimaryCatalogfmOrganizatsionnoPravovyeFormy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogTipyTsenNomenklatury struct {
	Key Guid
}

func (PrimaryCatalogTipyTsenNomenklatury) APIEntityType() string {
	return "Catalog_ТипыЦенНоменклатуры"
}
func (p PrimaryCatalogTipyTsenNomenklatury) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogStatiOtchetaPoProdazham struct {
	Key Guid
}

func (PrimaryCatalogStatiOtchetaPoProdazham) APIEntityType() string {
	return "Catalog_СтатьиОтчетаПоПродажам"
}
func (p PrimaryCatalogStatiOtchetaPoProdazham) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogVidyKodirovokiIzdelii struct {
	Key Guid
}

func (PrimaryCatalogVidyKodirovokiIzdelii) APIEntityType() string {
	return "Catalog_ВидыКодировокиИзделий"
}
func (p PrimaryCatalogVidyKodirovokiIzdelii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogVidyKodirovokiIzdeliiElementyKodirovki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogVidyKodirovokiIzdeliiElementyKodirovki) APIEntityType() string {
	return "Catalog_ВидыКодировокиИзделий_ЭлементыКодировки"
}
func (p PrimaryCatalogVidyKodirovokiIzdeliiElementyKodirovki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentUstanovkaSkidokNomenklatury struct {
	Key Guid
}

func (PrimaryDocumentUstanovkaSkidokNomenklatury) APIEntityType() string {
	return "Document_УстановкаСкидокНоменклатуры"
}
func (p PrimaryDocumentUstanovkaSkidokNomenklatury) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli) APIEntityType() string {
	return "Document_УстановкаСкидокНоменклатуры_ВремяПоДнямНедели"
}
func (p PrimaryDocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty) APIEntityType() string {
	return "Document_УстановкаСкидокНоменклатуры_ДисконтныеКарты"
}
func (p PrimaryDocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki) APIEntityType() string {
	return "Document_УстановкаСкидокНоменклатуры_ПолучателиСкидки"
}
func (p PrimaryDocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentUstanovkaSkidokNomenklaturyTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentUstanovkaSkidokNomenklaturyTovary) APIEntityType() string {
	return "Document_УстановкаСкидокНоменклатуры_Товары"
}
func (p PrimaryDocumentUstanovkaSkidokNomenklaturyTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenok struct {
	Key Guid
}

func (PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenok) APIEntityType() string {
	return "Catalog_УсловияПредоставленияСкидокНаценок"
}
func (p PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenok) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia) APIEntityType() string {
	return "Catalog_УсловияПредоставленияСкидокНаценок_ВремяДействия"
}
func (p PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli) APIEntityType() string {
	return "Catalog_УсловияПредоставленияСкидокНаценок_Получатели"
}
func (p PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki) APIEntityType() string {
	return "Catalog_УсловияПредоставленияСкидокНаценок_КомплектПокупки"
}
func (p PrimaryCatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryOutPay struct {
	Key Guid
}

func (PrimaryOutPay) APIEntityType() string {
	return "Document_РасходныйКассовыйОрдер"
}
func (p PrimaryOutPay) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha) APIEntityType() string {
	return "Document_РасходныйКассовыйОрдер_РасшифровкаПлатежа"
}
func (p PrimaryDocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentRaskhodnyiKassovyiOrderOplata struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentRaskhodnyiKassovyiOrderOplata) APIEntityType() string {
	return "Document_РасходныйКассовыйОрдер_Оплата"
}
func (p PrimaryDocumentRaskhodnyiKassovyiOrderOplata) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentRaskhodnyiKassovyiOrderTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentRaskhodnyiKassovyiOrderTovary) APIEntityType() string {
	return "Document_РасходныйКассовыйОрдер_Товары"
}
func (p PrimaryDocumentRaskhodnyiKassovyiOrderTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentSchetNaOplatuPostavshchika struct {
	Key Guid
}

func (PrimaryDocumentSchetNaOplatuPostavshchika) APIEntityType() string {
	return "Document_СчетНаОплатуПоставщика"
}
func (p PrimaryDocumentSchetNaOplatuPostavshchika) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentSchetNaOplatuPostavshchikaTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSchetNaOplatuPostavshchikaTovary) APIEntityType() string {
	return "Document_СчетНаОплатуПоставщика_Товары"
}
func (p PrimaryDocumentSchetNaOplatuPostavshchikaTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentSchetNaOplatuPostavshchikaUslugi struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentSchetNaOplatuPostavshchikaUslugi) APIEntityType() string {
	return "Document_СчетНаОплатуПоставщика_Услуги"
}
func (p PrimaryDocumentSchetNaOplatuPostavshchikaUslugi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentReestrSpetssviaz struct {
	Key Guid
}

func (PrimaryDocumentReestrSpetssviaz) APIEntityType() string {
	return "Document_РеестрСпецсвязь"
}
func (p PrimaryDocumentReestrSpetssviaz) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentReestrSpetssviazKlienty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentReestrSpetssviazKlienty) APIEntityType() string {
	return "Document_РеестрСпецсвязь_Клиенты"
}
func (p PrimaryDocumentReestrSpetssviazKlienty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentJournalKassovyeDokumenty struct {
	Ref     String
	RefType String
}

func (PrimaryDocumentJournalKassovyeDokumenty) APIEntityType() string {
	return "DocumentJournal_КассовыеДокументы"
}
func (p PrimaryDocumentJournalKassovyeDokumenty) Serialize() string {
	return "Ref=" + p.Ref.AsParameter() + ",Ref_Type=" + p.RefType.AsParameter()
}

type PrimaryInitialInstance struct {
	Key Guid
}

func (PrimaryInitialInstance) APIEntityType() string {
	return "Document_ВводНачальныхОстатков"
}
func (p PrimaryInitialInstance) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentVvodNachalnykhOstatkovVzaimoraschety struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVvodNachalnykhOstatkovVzaimoraschety) APIEntityType() string {
	return "Document_ВводНачальныхОстатков_Взаиморасчеты"
}
func (p PrimaryDocumentVvodNachalnykhOstatkovVzaimoraschety) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentVvodNachalnykhOstatkovTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentVvodNachalnykhOstatkovTovary) APIEntityType() string {
	return "Document_ВводНачальныхОстатков_Товары"
}
func (p PrimaryDocumentVvodNachalnykhOstatkovTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryPosting struct {
	Key Guid
}

func (PrimaryPosting) APIEntityType() string {
	return "Document_ОприходованиеТоваров"
}
func (p PrimaryPosting) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentOprikhodovanieTovarovTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOprikhodovanieTovarovTovary) APIEntityType() string {
	return "Document_ОприходованиеТоваров_Товары"
}
func (p PrimaryDocumentOprikhodovanieTovarovTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOprikhodovanieTovarovSertifikaty struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOprikhodovanieTovarovSertifikaty) APIEntityType() string {
	return "Document_ОприходованиеТоваров_Сертификаты"
}
func (p PrimaryDocumentOprikhodovanieTovarovSertifikaty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKomplekty struct {
	Key Guid
}

func (PrimaryCatalogKomplekty) APIEntityType() string {
	return "Catalog_Комплекты"
}
func (p PrimaryCatalogKomplekty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPereotsenkaTovarovPriniatykhNaKomissiiu struct {
	Key Guid
}

func (PrimaryDocumentPereotsenkaTovarovPriniatykhNaKomissiiu) APIEntityType() string {
	return "Document_ПереоценкаТоваровПринятыхНаКомиссию"
}
func (p PrimaryDocumentPereotsenkaTovarovPriniatykhNaKomissiiu) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary) APIEntityType() string {
	return "Document_ПереоценкаТоваровПринятыхНаКомиссию_Товары"
}
func (p PrimaryDocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentElektronnoePismo struct {
	Key Guid
}

func (PrimaryDocumentElektronnoePismo) APIEntityType() string {
	return "Document_ЭлектронноеПисьмо"
}
func (p PrimaryDocumentElektronnoePismo) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentElektronnoePismoKomuTCh struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentElektronnoePismoKomuTCh) APIEntityType() string {
	return "Document_ЭлектронноеПисьмо_КомуТЧ"
}
func (p PrimaryDocumentElektronnoePismoKomuTCh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentElektronnoePismoKopiiTCh struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentElektronnoePismoKopiiTCh) APIEntityType() string {
	return "Document_ЭлектронноеПисьмо_КопииТЧ"
}
func (p PrimaryDocumentElektronnoePismoKopiiTCh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentElektronnoePismoSkrytyeKopiiTCh struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentElektronnoePismoSkrytyeKopiiTCh) APIEntityType() string {
	return "Document_ЭлектронноеПисьмо_СкрытыеКопииТЧ"
}
func (p PrimaryDocumentElektronnoePismoSkrytyeKopiiTCh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogGruppyDefektov struct {
	Key Guid
}

func (PrimaryCatalogGruppyDefektov) APIEntityType() string {
	return "Catalog_ГруппыДефектов"
}
func (p PrimaryCatalogGruppyDefektov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogfmAnketaKlientaBenefitsariia struct {
	Key Guid
}

func (PrimaryCatalogfmAnketaKlientaBenefitsariia) APIEntityType() string {
	return "Catalog_фмАнкетаКлиентаБенефицария"
}
func (p PrimaryCatalogfmAnketaKlientaBenefitsariia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta) APIEntityType() string {
	return "Catalog_фмАнкетаКлиентаБенефицария_ДанныеКонтрагента"
}
func (p PrimaryCatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogTsenovyeGruppy struct {
	Key Guid
}

func (PrimaryCatalogTsenovyeGruppy) APIEntityType() string {
	return "Catalog_ЦеновыеГруппы"
}
func (p PrimaryCatalogTsenovyeGruppy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogPravilaTsenoobrazovaniia struct {
	Key Guid
}

func (PrimaryCatalogPravilaTsenoobrazovaniia) APIEntityType() string {
	return "Catalog_ПравилаЦенообразования"
}
func (p PrimaryCatalogPravilaTsenoobrazovaniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogPravilaTsenoobrazovaniiaTsenovyeGruppy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogPravilaTsenoobrazovaniiaTsenovyeGruppy) APIEntityType() string {
	return "Catalog_ПравилаЦенообразования_ЦеновыеГруппы"
}
func (p PrimaryCatalogPravilaTsenoobrazovaniiaTsenovyeGruppy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentObieiavlenieNaVznosNalichnymi struct {
	Key Guid
}

func (PrimaryDocumentObieiavlenieNaVznosNalichnymi) APIEntityType() string {
	return "Document_ОбъявлениеНаВзносНаличными"
}
func (p PrimaryDocumentObieiavlenieNaVznosNalichnymi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogValiuty struct {
	Key Guid
}

func (PrimaryCatalogValiuty) APIEntityType() string {
	return "Catalog_Валюты"
}
func (p PrimaryCatalogValiuty) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku struct {
	Key Guid
}

func (PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku) APIEntityType() string {
	return "Document_ПоступлениеТоваровУслугВНеавтоматизированнуюТорговуюТочку"
}
func (p PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary) APIEntityType() string {
	return "Document_ПоступлениеТоваровУслугВНеавтоматизированнуюТорговуюТочку_Товары"
}
func (p PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi) APIEntityType() string {
	return "Document_ПоступлениеТоваровУслугВНеавтоматизированнуюТорговуюТочку_Услуги"
}
func (p PrimaryDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKassyKKM struct {
	Key Guid
}

func (PrimaryCatalogKassyKKM) APIEntityType() string {
	return "Catalog_КассыККМ"
}
func (p PrimaryCatalogKassyKKM) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryProbe struct {
	Key Guid
}

func (PrimaryProbe) APIEntityType() string {
	return "Catalog_Пробы"
}
func (p PrimaryProbe) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogGruppyDostupa struct {
	Key Guid
}

func (PrimaryCatalogGruppyDostupa) APIEntityType() string {
	return "Catalog_ГруппыДоступа"
}
func (p PrimaryCatalogGruppyDostupa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogGruppyDostupaPolzovateli struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogGruppyDostupaPolzovateli) APIEntityType() string {
	return "Catalog_ГруппыДоступа_Пользователи"
}
func (p PrimaryCatalogGruppyDostupaPolzovateli) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogGruppyDostupaVidyDostupa struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogGruppyDostupaVidyDostupa) APIEntityType() string {
	return "Catalog_ГруппыДоступа_ВидыДоступа"
}
func (p PrimaryCatalogGruppyDostupaVidyDostupa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogGruppyDostupaZnacheniiaDostupa struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogGruppyDostupaZnacheniiaDostupa) APIEntityType() string {
	return "Catalog_ГруппыДоступа_ЗначенияДоступа"
}
func (p PrimaryCatalogGruppyDostupaZnacheniiaDostupa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogGruppyDostupaDostupPoPodsistemam struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogGruppyDostupaDostupPoPodsistemam) APIEntityType() string {
	return "Catalog_ГруппыДоступа_ДоступПоПодсистемам"
}
func (p PrimaryCatalogGruppyDostupaDostupPoPodsistemam) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogVidyKontaktnoiInformatsii struct {
	Key Guid
}

func (PrimaryCatalogVidyKontaktnoiInformatsii) APIEntityType() string {
	return "Catalog_ВидыКонтактнойИнформации"
}
func (p PrimaryCatalogVidyKontaktnoiInformatsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNomenklaturnyeGruppy struct {
	Key Guid
}

func (PrimaryCatalogNomenklaturnyeGruppy) APIEntityType() string {
	return "Catalog_НоменклатурныеГруппы"
}
func (p PrimaryCatalogNomenklaturnyeGruppy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentReestrSchetov struct {
	Key Guid
}

func (PrimaryDocumentReestrSchetov) APIEntityType() string {
	return "Document_РеестрСчетов"
}
func (p PrimaryDocumentReestrSchetov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentReestrSchetovReestr struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentReestrSchetovReestr) APIEntityType() string {
	return "Document_РеестрСчетов_Реестр"
}
func (p PrimaryDocumentReestrSchetovReestr) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu struct {
	Key Guid
}

func (PrimaryDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu) APIEntityType() string {
	return "Document_ИнвентаризацияТоваровОтданныхНаКомиссию"
}
func (p PrimaryDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary) APIEntityType() string {
	return "Document_ИнвентаризацияТоваровОтданныхНаКомиссию_Товары"
}
func (p PrimaryDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKlassifikatorStranMira struct {
	Key Guid
}

func (PrimaryCatalogKlassifikatorStranMira) APIEntityType() string {
	return "Catalog_КлассификаторСтранМира"
}
func (p PrimaryCatalogKlassifikatorStranMira) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogKlassifikatorEdinitsIzmereniia struct {
	Key Guid
}

func (PrimaryCatalogKlassifikatorEdinitsIzmereniia) APIEntityType() string {
	return "Catalog_КлассификаторЕдиницИзмерения"
}
func (p PrimaryCatalogKlassifikatorEdinitsIzmereniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikiRMK struct {
	Key Guid
}

func (PrimaryCatalogNastroikiRMK) APIEntityType() string {
	return "Catalog_НастройкиРМК"
}
func (p PrimaryCatalogNastroikiRMK) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogNastroikiRMKPoriadokPrimeneniiaSkidok struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikiRMKPoriadokPrimeneniiaSkidok) APIEntityType() string {
	return "Catalog_НастройкиРМК_ПорядокПримененияСкидок"
}
func (p PrimaryCatalogNastroikiRMKPoriadokPrimeneniiaSkidok) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogNastroikiRMKSostavNaimenovaniia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogNastroikiRMKSostavNaimenovaniia) APIEntityType() string {
	return "Catalog_НастройкиРМК_СоставНаименования"
}
func (p PrimaryCatalogNastroikiRMKSostavNaimenovaniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogKharakteristikiNomenklatury struct {
	Key Guid
}

func (PrimaryCatalogKharakteristikiNomenklatury) APIEntityType() string {
	return "Catalog_ХарактеристикиНоменклатуры"
}
func (p PrimaryCatalogKharakteristikiNomenklatury) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogKharakteristikiNomenklaturySpetsifikatsiia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogKharakteristikiNomenklaturySpetsifikatsiia) APIEntityType() string {
	return "Catalog_ХарактеристикиНоменклатуры_Спецификация"
}
func (p PrimaryCatalogKharakteristikiNomenklaturySpetsifikatsiia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtborTovarov struct {
	Key Guid
}

func (PrimaryDocumentOtborTovarov) APIEntityType() string {
	return "Document_ОтборТоваров"
}
func (p PrimaryDocumentOtborTovarov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentOtborTovarovTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtborTovarovTovary) APIEntityType() string {
	return "Document_ОтборТоваров_Товары"
}
func (p PrimaryDocumentOtborTovarovTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtborTovarovTovaryKOtboru struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtborTovarovTovaryKOtboru) APIEntityType() string {
	return "Document_ОтборТоваров_ТоварыКОтбору"
}
func (p PrimaryDocumentOtborTovarovTovaryKOtboru) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogSposobyDostavkiTovara struct {
	Key Guid
}

func (PrimaryCatalogSposobyDostavkiTovara) APIEntityType() string {
	return "Catalog_СпособыДоставкиТовара"
}
func (p PrimaryCatalogSposobyDostavkiTovara) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogPodrazdeleniia struct {
	Key Guid
}

func (PrimaryCatalogPodrazdeleniia) APIEntityType() string {
	return "Catalog_Подразделения"
}
func (p PrimaryCatalogPodrazdeleniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentJournalPreiskuranty struct {
	Ref     String
	RefType String
}

func (PrimaryDocumentJournalPreiskuranty) APIEntityType() string {
	return "DocumentJournal_Прейскуранты"
}
func (p PrimaryDocumentJournalPreiskuranty) Serialize() string {
	return "Ref=" + p.Ref.AsParameter() + ",Ref_Type=" + p.RefType.AsParameter()
}

type PrimaryCatalogRelizyIuvelirnykhSalonov struct {
	Key Guid
}

func (PrimaryCatalogRelizyIuvelirnykhSalonov) APIEntityType() string {
	return "Catalog_РелизыЮвелирныхСалонов"
}
func (p PrimaryCatalogRelizyIuvelirnykhSalonov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy) APIEntityType() string {
	return "Catalog_РелизыЮвелирныхСалонов_ОбновляемыеРелизы"
}
func (p PrimaryCatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetKomissioneraOProdazhakh struct {
	Key Guid
}

func (PrimaryDocumentOtchetKomissioneraOProdazhakh) APIEntityType() string {
	return "Document_ОтчетКомиссионераОПродажах"
}
func (p PrimaryDocumentOtchetKomissioneraOProdazhakh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva) APIEntityType() string {
	return "Document_ОтчетКомиссионераОПродажах_ДенежныеСредства"
}
func (p PrimaryDocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentOtchetKomissioneraOProdazhakhTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentOtchetKomissioneraOProdazhakhTovary) APIEntityType() string {
	return "Document_ОтчетКомиссионераОПродажах_Товары"
}
func (p PrimaryDocumentOtchetKomissioneraOProdazhakhTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogTovarnyeKategorii struct {
	Key Guid
}

func (PrimaryCatalogTovarnyeKategorii) APIEntityType() string {
	return "Catalog_ТоварныеКатегории"
}
func (p PrimaryCatalogTovarnyeKategorii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogDokumentyUdostoveriaiushchieLichnost struct {
	Key Guid
}

func (PrimaryCatalogDokumentyUdostoveriaiushchieLichnost) APIEntityType() string {
	return "Catalog_ДокументыУдостоверяющиеЛичность"
}
func (p PrimaryCatalogDokumentyUdostoveriaiushchieLichnost) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogFiltryDliaElektronnykhPisem struct {
	Key Guid
}

func (PrimaryCatalogFiltryDliaElektronnykhPisem) APIEntityType() string {
	return "Catalog_ФильтрыДляЭлектронныхПисем"
}
func (p PrimaryCatalogFiltryDliaElektronnykhPisem) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogFiltryDliaElektronnykhPisemDeistviiaFiltra struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogFiltryDliaElektronnykhPisemDeistviiaFiltra) APIEntityType() string {
	return "Catalog_ФильтрыДляЭлектронныхПисем_ДействияФильтра"
}
func (p PrimaryCatalogFiltryDliaElektronnykhPisemDeistviiaFiltra) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogFiltryDliaElektronnykhPisemUsloviiaFiltra struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogFiltryDliaElektronnykhPisemUsloviiaFiltra) APIEntityType() string {
	return "Catalog_ФильтрыДляЭлектронныхПисем_УсловияФильтра"
}
func (p PrimaryCatalogFiltryDliaElektronnykhPisemUsloviiaFiltra) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPreiskurantTsenNaTsvKamni struct {
	Key Guid
}

func (PrimaryDocumentPreiskurantTsenNaTsvKamni) APIEntityType() string {
	return "Document_ПрейскурантЦенНаЦвКамни"
}
func (p PrimaryDocumentPreiskurantTsenNaTsvKamni) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPreiskurantTsenNaTsvKamniTablitsa struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPreiskurantTsenNaTsvKamniTablitsa) APIEntityType() string {
	return "Document_ПрейскурантЦенНаЦвКамни_Таблица"
}
func (p PrimaryDocumentPreiskurantTsenNaTsvKamniTablitsa) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimarySize struct {
	Key Guid
}

func (PrimarySize) APIEntityType() string {
	return "Catalog_Размер"
}
func (p PrimarySize) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogTipyDragotsennykhMetallov struct {
	Key Guid
}

func (PrimaryCatalogTipyDragotsennykhMetallov) APIEntityType() string {
	return "Catalog_ТипыДрагоценныхМеталлов"
}
func (p PrimaryCatalogTipyDragotsennykhMetallov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentTelemarketing struct {
	Key Guid
}

func (PrimaryDocumentTelemarketing) APIEntityType() string {
	return "Document_Телемаркетинг"
}
func (p PrimaryDocumentTelemarketing) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentTelemarketingUchastniki struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentTelemarketingUchastniki) APIEntityType() string {
	return "Document_Телемаркетинг_Участники"
}
func (p PrimaryDocumentTelemarketingUchastniki) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentVozvratDavalcheskogoMetalla struct {
	Key Guid
}

func (PrimaryDocumentVozvratDavalcheskogoMetalla) APIEntityType() string {
	return "Document_ВозвратДавальческогоМеталла"
}
func (p PrimaryDocumentVozvratDavalcheskogoMetalla) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogAdresnyeSokrashcheniia struct {
	Key Guid
}

func (PrimaryCatalogAdresnyeSokrashcheniia) APIEntityType() string {
	return "Catalog_АдресныеСокращения"
}
func (p PrimaryCatalogAdresnyeSokrashcheniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentRassylkaAnket struct {
	Key Guid
}

func (PrimaryDocumentRassylkaAnket) APIEntityType() string {
	return "Document_РассылкаАнкет"
}
func (p PrimaryDocumentRassylkaAnket) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentRassylkaAnketVlozheniia struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentRassylkaAnketVlozheniia) APIEntityType() string {
	return "Document_РассылкаАнкет_Вложения"
}
func (p PrimaryDocumentRassylkaAnketVlozheniia) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentRassylkaAnketPoluchateli struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentRassylkaAnketPoluchateli) APIEntityType() string {
	return "Document_РассылкаАнкет_Получатели"
}
func (p PrimaryDocumentRassylkaAnketPoluchateli) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogVidyDeiatelnostiKontragentov struct {
	Key Guid
}

func (PrimaryCatalogVidyDeiatelnostiKontragentov) APIEntityType() string {
	return "Catalog_ВидыДеятельностиКонтрагентов"
}
func (p PrimaryCatalogVidyDeiatelnostiKontragentov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogTorgovoeOborudovanie struct {
	Key Guid
}

func (PrimaryCatalogTorgovoeOborudovanie) APIEntityType() string {
	return "Catalog_ТорговоеОборудование"
}
func (p PrimaryCatalogTorgovoeOborudovanie) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogSkhemyRealizatsii struct {
	Key Guid
}

func (PrimaryCatalogSkhemyRealizatsii) APIEntityType() string {
	return "Catalog_СхемыРеализации"
}
func (p PrimaryCatalogSkhemyRealizatsii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogSkhemyRealizatsiiEtapySkhemy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogSkhemyRealizatsiiEtapySkhemy) APIEntityType() string {
	return "Catalog_СхемыРеализации_ЭтапыСхемы"
}
func (p PrimaryCatalogSkhemyRealizatsiiEtapySkhemy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogPodkliuchaemoeOborudovanie struct {
	Key Guid
}

func (PrimaryCatalogPodkliuchaemoeOborudovanie) APIEntityType() string {
	return "Catalog_ПодключаемоеОборудование"
}
func (p PrimaryCatalogPodkliuchaemoeOborudovanie) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii struct {
	Key Guid
}

func (PrimaryDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii) APIEntityType() string {
	return "Document_КлассификацияПокупателейПоСтадиямВзаимоотношений"
}
func (p PrimaryDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov) APIEntityType() string {
	return "Document_КлассификацияПокупателейПоСтадиямВзаимоотношений_ТаблицаРаспределенияКонтрагентов"
}
func (p PrimaryDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogGabarity struct {
	Key Guid
}

func (PrimaryCatalogGabarity) APIEntityType() string {
	return "Catalog_Габариты"
}
func (p PrimaryCatalogGabarity) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZakazKlienta struct {
	Key Guid
}

func (PrimaryDocumentZakazKlienta) APIEntityType() string {
	return "Document_ЗаказКлиента"
}
func (p PrimaryDocumentZakazKlienta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentZakazKlientaTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentZakazKlientaTovary) APIEntityType() string {
	return "Document_ЗаказКлиента_Товары"
}
func (p PrimaryDocumentZakazKlientaTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryArriveFromManufacturing struct {
	Key Guid
}

func (PrimaryArriveFromManufacturing) APIEntityType() string {
	return "Document_ПоступлениеПродукцииИзПроизводства"
}
func (p PrimaryArriveFromManufacturing) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryArriveFromManufacturingInstance struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryArriveFromManufacturingInstance) APIEntityType() string {
	return "Document_ПоступлениеПродукцииИзПроизводства_Товары"
}
func (p PrimaryArriveFromManufacturingInstance) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPostuplenieProduktsiiIzProizvodstvaMaterialy struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPostuplenieProduktsiiIzProizvodstvaMaterialy) APIEntityType() string {
	return "Document_ПоступлениеПродукцииИзПроизводства_Материалы"
}
func (p PrimaryDocumentPostuplenieProduktsiiIzProizvodstvaMaterialy) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentJournalZakazyPostavshchikam struct {
	Ref     String
	RefType String
}

func (PrimaryDocumentJournalZakazyPostavshchikam) APIEntityType() string {
	return "DocumentJournal_ЗаказыПоставщикам"
}
func (p PrimaryDocumentJournalZakazyPostavshchikam) Serialize() string {
	return "Ref=" + p.Ref.AsParameter() + ",Ref_Type=" + p.RefType.AsParameter()
}

type PrimaryDocumentJournalSkladskieDokumenty struct {
	Ref     String
	RefType String
}

func (PrimaryDocumentJournalSkladskieDokumenty) APIEntityType() string {
	return "DocumentJournal_СкладскиеДокументы"
}
func (p PrimaryDocumentJournalSkladskieDokumenty) Serialize() string {
	return "Ref=" + p.Ref.AsParameter() + ",Ref_Type=" + p.RefType.AsParameter()
}

type PrimaryCatalogsmsUsloviiaOtboraDiskontnykhKart struct {
	Key Guid
}

func (PrimaryCatalogsmsUsloviiaOtboraDiskontnykhKart) APIEntityType() string {
	return "Catalog_смсУсловияОтбораДисконтныхКарт"
}
func (p PrimaryCatalogsmsUsloviiaOtboraDiskontnykhKart) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryArrive struct {
	Key Guid
}

func (PrimaryArrive) APIEntityType() string {
	return "Document_ПоступлениеТоваровУслуг"
}
func (p PrimaryArrive) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPostuplenieTovarovUslugTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPostuplenieTovarovUslugTovary) APIEntityType() string {
	return "Document_ПоступлениеТоваровУслуг_Товары"
}
func (p PrimaryDocumentPostuplenieTovarovUslugTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPostuplenieTovarovUslugUslugi struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPostuplenieTovarovUslugUslugi) APIEntityType() string {
	return "Document_ПоступлениеТоваровУслуг_Услуги"
}
func (p PrimaryDocumentPostuplenieTovarovUslugUslugi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentSchetFakturaVydannyi struct {
	Key Guid
}

func (PrimaryDocumentSchetFakturaVydannyi) APIEntityType() string {
	return "Document_СчетФактураВыданный"
}
func (p PrimaryDocumentSchetFakturaVydannyi) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlanProdazhPoSalonam struct {
	Key Guid
}

func (PrimaryDocumentPlanProdazhPoSalonam) APIEntityType() string {
	return "Document_ПланПродажПоСалонам"
}
func (p PrimaryDocumentPlanProdazhPoSalonam) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPlanProdazhPoSalonamSalony struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlanProdazhPoSalonamSalony) APIEntityType() string {
	return "Document_ПланПродажПоСалонам_Салоны"
}
func (p PrimaryDocumentPlanProdazhPoSalonamSalony) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPlanProdazhPoSalonamDniPoGrafiku struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPlanProdazhPoSalonamDniPoGrafiku) APIEntityType() string {
	return "Document_ПланПродажПоСалонам_ДниПоГрафику"
}
func (p PrimaryDocumentPlanProdazhPoSalonamDniPoGrafiku) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogBankovskieScheta struct {
	Key Guid
}

func (PrimaryCatalogBankovskieScheta) APIEntityType() string {
	return "Catalog_БанковскиеСчета"
}
func (p PrimaryCatalogBankovskieScheta) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakh struct {
	Key Guid
}

func (PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakh) APIEntityType() string {
	return "Document_СторнированиеОтчетаКомитентуОПродажах"
}
func (p PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakh) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva) APIEntityType() string {
	return "Document_СторнированиеОтчетаКомитентуОПродажах_ДенежныеСредства"
}
func (p PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakhTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakhTovary) APIEntityType() string {
	return "Document_СторнированиеОтчетаКомитентуОПродажах_Товары"
}
func (p PrimaryDocumentStornirovanieOtchetaKomitentuOProdazhakhTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPeredachaVRemont struct {
	Key Guid
}

func (PrimaryDocumentPeredachaVRemont) APIEntityType() string {
	return "Document_ПередачаВРемонт"
}
func (p PrimaryDocumentPeredachaVRemont) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryDocumentPeredachaVRemontIzdeliiaPriniatyeVRemont struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPeredachaVRemontIzdeliiaPriniatyeVRemont) APIEntityType() string {
	return "Document_ПередачаВРемонт_ИзделияПринятыеВРемонт"
}
func (p PrimaryDocumentPeredachaVRemontIzdeliiaPriniatyeVRemont) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryDocumentPeredachaVRemontTovary struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryDocumentPeredachaVRemontTovary) APIEntityType() string {
	return "Document_ПередачаВРемонт_Товары"
}
func (p PrimaryDocumentPeredachaVRemontTovary) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}

type PrimaryCatalogPolzovateli struct {
	Key Guid
}

func (PrimaryCatalogPolzovateli) APIEntityType() string {
	return "Catalog_Пользователи"
}
func (p PrimaryCatalogPolzovateli) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogTsenovyeKoridory struct {
	Key Guid
}

func (PrimaryCatalogTsenovyeKoridory) APIEntityType() string {
	return "Catalog_ЦеновыеКоридоры"
}
func (p PrimaryCatalogTsenovyeKoridory) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogGruppySkladov struct {
	Key Guid
}

func (PrimaryCatalogGruppySkladov) APIEntityType() string {
	return "Catalog_ГруппыСкладов"
}
func (p PrimaryCatalogGruppySkladov) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter()
}

type PrimaryCatalogGruppySkladovSklady struct {
	Key        Guid
	LineNumber Int64
}

func (PrimaryCatalogGruppySkladovSklady) APIEntityType() string {
	return "Catalog_ГруппыСкладов_Склады"
}
func (p PrimaryCatalogGruppySkladovSklady) Serialize() string {
	return "Ref_Key=" + p.Key.AsParameter() + ",LineNumber=" + p.LineNumber.AsParameter()
}
