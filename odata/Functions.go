package odata

type FunctionAccumulationRegisterPartiiTovarovVProizvodstveBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterPartiiTovarovVProizvodstveBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterPartiiTovarovVProizvodstveBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterPartiiTovarovVProizvodstveTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterPartiiTovarovVProizvodstveTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterPartiiTovarovVProizvodstveTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterPartiiTovarovVProizvodstveBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterPartiiTovarovVProizvodstveBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterPartiiTovarovVProizvodstveBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterVnutrennieZakazyBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterVnutrennieZakazyBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterVnutrennieZakazyBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterVnutrennieZakazyTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterVnutrennieZakazyTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterVnutrennieZakazyTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterVnutrennieZakazyBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterVnutrennieZakazyBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterVnutrennieZakazyBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKomitentaBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKomitentaBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKomitentaBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKomitentaTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKomitentaTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKomitentaTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKomitentaBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKomitentaBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKomitentaBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterZakazyKlientovBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterZakazyKlientovBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterZakazyKlientovBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterZakazyKlientovTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterZakazyKlientovTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterZakazyKlientovTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterZakazyKlientovBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterZakazyKlientovBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterZakazyKlientovBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSummyPoFinmonitoringuRoznitsaBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterSummyPoFinmonitoringuRoznitsaBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterSummyPoFinmonitoringuRoznitsaBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterSummyPoFinmonitoringuRoznitsaTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSummyPoFinmonitoringuRoznitsaTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterSummyPoFinmonitoringuRoznitsaTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSummyPoFinmonitoringuRoznitsaBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSummyPoFinmonitoringuRoznitsaBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterSummyPoFinmonitoringuRoznitsaBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKPolucheniiuTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKPolucheniiuTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKPolucheniiuTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKPolucheniiuBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterProdazhiPoDiskontnymKartamTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterProdazhiPoDiskontnymKartamTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterProdazhiPoDiskontnymKartamTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterTovaryPoluchennyeBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterTovaryPoluchennyeBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterTovaryPoluchennyeBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterTovaryPoluchennyeTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryPoluchennyeTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterTovaryPoluchennyeTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterTovaryPoluchennyeBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryPoluchennyeBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterTovaryPoluchennyeBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSvobodnyeOstatkiBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterSvobodnyeOstatkiBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterSvobodnyeOstatkiBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterSvobodnyeOstatkiTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSvobodnyeOstatkiTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterSvobodnyeOstatkiTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSvobodnyeOstatkiBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSvobodnyeOstatkiBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterSvobodnyeOstatkiBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSummyVRassrochkuBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterSummyVRassrochkuBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterSummyVRassrochkuBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterSummyVRassrochkuTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSummyVRassrochkuTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterSummyVRassrochkuTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSummyVRassrochkuBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSummyVRassrochkuBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterSummyVRassrochkuBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterGrafikPlatezheiBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterGrafikPlatezheiBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterGrafikPlatezheiBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterGrafikPlatezheiTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterGrafikPlatezheiTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterGrafikPlatezheiTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterGrafikPlatezheiBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterGrafikPlatezheiBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterGrafikPlatezheiBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterRoznichnaiaVyruchkaBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterRoznichnaiaVyruchkaBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterRoznichnaiaVyruchkaBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterRoznichnaiaVyruchkaTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterRoznichnaiaVyruchkaTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterRoznichnaiaVyruchkaTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterRoznichnaiaVyruchkaBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterRoznichnaiaVyruchkaBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterRoznichnaiaVyruchkaBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterTovaryVPutiBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterTovaryVPutiBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterTovaryVPutiBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterTovaryVPutiTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryVPutiTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterTovaryVPutiTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterTovaryVPutiBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryVPutiBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterTovaryVPutiBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterPoteriMetallaVProizvodstveTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterPoteriMetallaVProizvodstveTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterPoteriMetallaVProizvodstveTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterPartiiTovarovNaSkladakhBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterPartiiTovarovNaSkladakhBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterPartiiTovarovNaSkladakhBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterPartiiTovarovNaSkladakhTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterPartiiTovarovNaSkladakhTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterPartiiTovarovNaSkladakhTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterPartiiTovarovNaSkladakhBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterPartiiTovarovNaSkladakhBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterPartiiTovarovNaSkladakhBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSummyDokumentovDliaObmenaBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterSummyDokumentovDliaObmenaBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterSummyDokumentovDliaObmenaBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterSummyDokumentovDliaObmenaTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSummyDokumentovDliaObmenaTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterSummyDokumentovDliaObmenaTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSummyDokumentovDliaObmenaBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSummyDokumentovDliaObmenaBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterSummyDokumentovDliaObmenaBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDvizheniiaDenezhnykhSredstvTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDvizheniiaDenezhnykhSredstvTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterDvizheniiaDenezhnykhSredstvTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterProdazhiPoStatiamBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterProdazhiPoStatiamBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterProdazhiPoStatiamBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterProdazhiPoStatiamTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterProdazhiPoStatiamTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterProdazhiPoStatiamTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterProdazhiPoStatiamBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterProdazhiPoStatiamBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterProdazhiPoStatiamBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionInformationRegisterTsenyNomenklaturyRecordTypeSliceLast struct {
	Condition String   `odata:"Condition"`
	Period    DateTime `odata:"Period"`
}

func (FunctionInformationRegisterTsenyNomenklaturyRecordTypeSliceLast) Name() string {
	return "SliceLast"
}
func (f FunctionInformationRegisterTsenyNomenklaturyRecordTypeSliceLast) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionInformationRegisterTsenyNomenklaturyRecordTypeSliceFirst struct {
	Condition String   `odata:"Condition"`
	Period    DateTime `odata:"Period"`
}

func (FunctionInformationRegisterTsenyNomenklaturyRecordTypeSliceFirst) Name() string {
	return "SliceFirst"
}
func (f FunctionInformationRegisterTsenyNomenklaturyRecordTypeSliceFirst) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaVRezerveBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaVRezerveBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaVRezerveBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaVRezerveTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaVRezerveTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaVRezerveTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaVRezerveBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaVRezerveBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaVRezerveBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDavalcheskiiMetallPoteriTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDavalcheskiiMetallPoteriTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterDavalcheskiiMetallPoteriTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionInformationRegisterTsenyPoPreiskurantuRecordTypeSliceLast struct {
	Condition String   `odata:"Condition"`
	Period    DateTime `odata:"Period"`
}

func (FunctionInformationRegisterTsenyPoPreiskurantuRecordTypeSliceLast) Name() string {
	return "SliceLast"
}
func (f FunctionInformationRegisterTsenyPoPreiskurantuRecordTypeSliceLast) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionInformationRegisterTsenyPoPreiskurantuRecordTypeSliceFirst struct {
	Condition String   `odata:"Condition"`
	Period    DateTime `odata:"Period"`
}

func (FunctionInformationRegisterTsenyPoPreiskurantuRecordTypeSliceFirst) Name() string {
	return "SliceFirst"
}
func (f FunctionInformationRegisterTsenyPoPreiskurantuRecordTypeSliceFirst) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterTovaryVOtboreBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterTovaryVOtboreBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterTovaryVOtboreBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterTovaryVOtboreTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryVOtboreTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterTovaryVOtboreTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterTovaryVOtboreBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryVOtboreBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterTovaryVOtboreBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterRealizovannyeTovaryBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterRealizovannyeTovaryBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterRealizovannyeTovaryBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterRealizovannyeTovaryTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterRealizovannyeTovaryTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterRealizovannyeTovaryTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterRealizovannyeTovaryBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterRealizovannyeTovaryBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterRealizovannyeTovaryBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKomissioneraBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKomissioneraBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKomissioneraBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKomissioneraTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKomissioneraTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKomissioneraTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKomissioneraBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKomissioneraBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKomissioneraBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterProdazhi1Turnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterProdazhi1Turnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterProdazhi1Turnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterTovaryNaSkladakhAMBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterTovaryNaSkladakhAMBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterTovaryNaSkladakhAMBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterTovaryNaSkladakhAMTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryNaSkladakhAMTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterTovaryNaSkladakhAMTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterTovaryNaSkladakhAMBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryNaSkladakhAMBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterTovaryNaSkladakhAMBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSummyPoFinmonitoringuBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterSummyPoFinmonitoringuBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterSummyPoFinmonitoringuBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterSummyPoFinmonitoringuTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSummyPoFinmonitoringuTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterSummyPoFinmonitoringuTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSummyPoFinmonitoringuBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSummyPoFinmonitoringuBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterSummyPoFinmonitoringuBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionInformationRegisterTsenyNomenklaturyKontragentovRecordTypeSliceLast struct {
	Condition String   `odata:"Condition"`
	Period    DateTime `odata:"Period"`
}

func (FunctionInformationRegisterTsenyNomenklaturyKontragentovRecordTypeSliceLast) Name() string {
	return "SliceLast"
}
func (f FunctionInformationRegisterTsenyNomenklaturyKontragentovRecordTypeSliceLast) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionInformationRegisterTsenyNomenklaturyKontragentovRecordTypeSliceFirst struct {
	Condition String   `odata:"Condition"`
	Period    DateTime `odata:"Period"`
}

func (FunctionInformationRegisterTsenyNomenklaturyKontragentovRecordTypeSliceFirst) Name() string {
	return "SliceFirst"
}
func (f FunctionInformationRegisterTsenyNomenklaturyKontragentovRecordTypeSliceFirst) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterVzaimoraschetySKontragentamiBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterVzaimoraschetySKontragentamiBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterVzaimoraschetySKontragentamiBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterVzaimoraschetySKontragentamiTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterVzaimoraschetySKontragentamiTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterVzaimoraschetySKontragentamiTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterVzaimoraschetySKontragentamiBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterVzaimoraschetySKontragentamiBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterVzaimoraschetySKontragentamiBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSummyPokupokPoDiskontnymKartamBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterSummyPokupokPoDiskontnymKartamBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterSummyPokupokPoDiskontnymKartamBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterSummyPokupokPoDiskontnymKartamTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSummyPokupokPoDiskontnymKartamTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterSummyPokupokPoDiskontnymKartamTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterSummyPokupokPoDiskontnymKartamBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterSummyPokupokPoDiskontnymKartamBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterSummyPokupokPoDiskontnymKartamBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterVypolneniePlanaProdazhTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterVypolneniePlanaProdazhTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterVypolneniePlanaProdazhTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDavalcheskiiMetallBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterDavalcheskiiMetallBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterDavalcheskiiMetallBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterDavalcheskiiMetallTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDavalcheskiiMetallTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterDavalcheskiiMetallTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDavalcheskiiMetallBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDavalcheskiiMetallBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterDavalcheskiiMetallBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterTovaryPeredannyeBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterTovaryPeredannyeBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterTovaryPeredannyeBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterTovaryPeredannyeTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryPeredannyeTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterTovaryPeredannyeTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterTovaryPeredannyeBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterTovaryPeredannyeBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterTovaryPeredannyeBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalance struct {
	Condition  String   `odata:"Condition"`
	Dimensions String   `odata:"Dimensions"`
	Period     DateTime `odata:"Period"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalance) Name() string {
	return "Balance"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalance) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&Period=" + f.Period.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKSpisaniiuTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKSpisaniiuTurnovers) Name() string {
	return "Turnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKSpisaniiuTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionAccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalanceAndTurnovers struct {
	Condition   String   `odata:"Condition"`
	Dimensions  String   `odata:"Dimensions"`
	EndPeriod   DateTime `odata:"EndPeriod"`
	StartPeriod DateTime `odata:"StartPeriod"`
}

func (FunctionAccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalanceAndTurnovers) Name() string {
	return "BalanceAndTurnovers"
}
func (f FunctionAccumulationRegisterDenezhnyeSredstvaKSpisaniiuBalanceAndTurnovers) Parameters() string {
	return "Condition=" + f.Condition.AsParameter() + "&Dimensions=" + f.Dimensions.AsParameter() + "&EndPeriod=" + f.EndPeriod.AsParameter() + "&StartPeriod=" + f.StartPeriod.AsParameter()
}

type FunctionOrderPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionOrderPost) Name() string {
	return "Post"
}
func (f FunctionOrderPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionOrderUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionOrderUnpost) Name() string {
	return "Unpost"
}
func (f FunctionOrderUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPereotsenkaValiutnykhSredstvPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPereotsenkaValiutnykhSredstvPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPereotsenkaValiutnykhSredstvPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPereotsenkaValiutnykhSredstvUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPereotsenkaValiutnykhSredstvUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPereotsenkaValiutnykhSredstvUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOtchetKomitentuOProdazhakhPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOtchetKomitentuOProdazhakhPost) Name() string {
	return "Post"
}
func (f FunctionDocumentOtchetKomitentuOProdazhakhPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOtchetKomitentuOProdazhakhUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOtchetKomitentuOProdazhakhUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentOtchetKomitentuOProdazhakhUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZaiavkaNaPereotsenkuTovarovPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZaiavkaNaPereotsenkuTovarovPost) Name() string {
	return "Post"
}
func (f FunctionDocumentZaiavkaNaPereotsenkuTovarovPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZaiavkaNaPereotsenkuTovarovUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZaiavkaNaPereotsenkuTovarovUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentZaiavkaNaPereotsenkuTovarovUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakrytieZakazovKlientovPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakrytieZakazovKlientovPost) Name() string {
	return "Post"
}
func (f FunctionDocumentZakrytieZakazovKlientovPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakrytieZakazovKlientovUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakrytieZakazovKlientovUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentZakrytieZakazovKlientovUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnoePoruchenieVkhodiashcheePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnoePoruchenieVkhodiashcheePost) Name() string {
	return "Post"
}
func (f FunctionDocumentPlatezhnoePoruchenieVkhodiashcheePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnoePoruchenieVkhodiashcheeUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnoePoruchenieVkhodiashcheeUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPlatezhnoePoruchenieVkhodiashcheeUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVydachaZakazaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVydachaZakazaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentVydachaZakazaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVydachaZakazaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVydachaZakazaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentVydachaZakazaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentRealizatsiiaTovarovUslugPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentRealizatsiiaTovarovUslugPost) Name() string {
	return "Post"
}
func (f FunctionDocumentRealizatsiiaTovarovUslugPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentRealizatsiiaTovarovUslugUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentRealizatsiiaTovarovUslugUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentRealizatsiiaTovarovUslugUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSobytiePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSobytiePost) Name() string {
	return "Post"
}
func (f FunctionDocumentSobytiePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSobytieUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSobytieUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentSobytieUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentKorrektirovkaDolgaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentKorrektirovkaDolgaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentKorrektirovkaDolgaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentKorrektirovkaDolgaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentKorrektirovkaDolgaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentKorrektirovkaDolgaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZaiavkaNaRaskhodovanieSredstvPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZaiavkaNaRaskhodovanieSredstvPost) Name() string {
	return "Post"
}
func (f FunctionDocumentZaiavkaNaRaskhodovanieSredstvPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZaiavkaNaRaskhodovanieSredstvUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZaiavkaNaRaskhodovanieSredstvUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentZaiavkaNaRaskhodovanieSredstvUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakrytieZakazovPostavshchikamPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakrytieZakazovPostavshchikamPost) Name() string {
	return "Post"
}
func (f FunctionDocumentZakrytieZakazovPostavshchikamPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakrytieZakazovPostavshchikamUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakrytieZakazovPostavshchikamUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentZakrytieZakazovPostavshchikamUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAnketyKlientovDliaFinMonitoringaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAnketyKlientovDliaFinMonitoringaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentAnketyKlientovDliaFinMonitoringaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAnketyKlientovDliaFinMonitoringaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAnketyKlientovDliaFinMonitoringaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentAnketyKlientovDliaFinMonitoringaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPostuplenieDavalcheskogoMetallaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPostuplenieDavalcheskogoMetallaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPostuplenieDavalcheskogoMetallaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPostuplenieDavalcheskogoMetallaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPostuplenieDavalcheskogoMetallaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPostuplenieDavalcheskogoMetallaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInkassovoePorucheniePeredannoePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInkassovoePorucheniePeredannoePost) Name() string {
	return "Post"
}
func (f FunctionDocumentInkassovoePorucheniePeredannoePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInkassovoePorucheniePeredannoeUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInkassovoePorucheniePeredannoeUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentInkassovoePorucheniePeredannoeUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionCorrectingPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionCorrectingPost) Name() string {
	return "Post"
}
func (f FunctionCorrectingPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionCorrectingUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionCorrectingUnpost) Name() string {
	return "Unpost"
}
func (f FunctionCorrectingUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInternetZakazPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInternetZakazPost) Name() string {
	return "Post"
}
func (f FunctionDocumentInternetZakazPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInternetZakazUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInternetZakazUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentInternetZakazUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionSaleJournalPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionSaleJournalPost) Name() string {
	return "Post"
}
func (f FunctionSaleJournalPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionSaleJournalUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionSaleJournalUnpost) Name() string {
	return "Unpost"
}
func (f FunctionSaleJournalUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOtmenaSkidokNomenklaturyPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOtmenaSkidokNomenklaturyPost) Name() string {
	return "Post"
}
func (f FunctionDocumentOtmenaSkidokNomenklaturyPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOtmenaSkidokNomenklaturyUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOtmenaSkidokNomenklaturyUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentOtmenaSkidokNomenklaturyUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentKassovyiChekKorrektsiiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentKassovyiChekKorrektsiiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentKassovyiChekKorrektsiiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentKassovyiChekKorrektsiiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentKassovyiChekKorrektsiiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentKassovyiChekKorrektsiiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSchetNaOplatuPokupateliuPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSchetNaOplatuPokupateliuPost) Name() string {
	return "Post"
}
func (f FunctionDocumentSchetNaOplatuPokupateliuPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSchetNaOplatuPokupateliuUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSchetNaOplatuPokupateliuUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentSchetNaOplatuPokupateliuUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZamenaDiskontnoiKartyPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZamenaDiskontnoiKartyPost) Name() string {
	return "Post"
}
func (f FunctionDocumentZamenaDiskontnoiKartyPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZamenaDiskontnoiKartyUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZamenaDiskontnoiKartyUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentZamenaDiskontnoiKartyUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionReturnToSupplierPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionReturnToSupplierPost) Name() string {
	return "Post"
}
func (f FunctionReturnToSupplierPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionReturnToSupplierUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionReturnToSupplierUnpost) Name() string {
	return "Unpost"
}
func (f FunctionReturnToSupplierUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInventarizatsiiaTovarovNaSkladePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInventarizatsiiaTovarovNaSkladePost) Name() string {
	return "Post"
}
func (f FunctionDocumentInventarizatsiiaTovarovNaSkladePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInventarizatsiiaTovarovNaSkladeUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInventarizatsiiaTovarovNaSkladeUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentInventarizatsiiaTovarovNaSkladeUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPrikhodnyiKassovyiOrderPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPrikhodnyiKassovyiOrderPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPrikhodnyiKassovyiOrderPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPrikhodnyiKassovyiOrderUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPrikhodnyiKassovyiOrderUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPrikhodnyiKassovyiOrderUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentDenezhnyiChekPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentDenezhnyiChekPost) Name() string {
	return "Post"
}
func (f FunctionDocumentDenezhnyiChekPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentDenezhnyiChekUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentDenezhnyiChekUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentDenezhnyiChekUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVozvratMaterialovIzProizvodstvaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVozvratMaterialovIzProizvodstvaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentVozvratMaterialovIzProizvodstvaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVozvratMaterialovIzProizvodstvaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVozvratMaterialovIzProizvodstvaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentVozvratMaterialovIzProizvodstvaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPereotsenkaTovarovOtdannykhNaKomissiiuPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPereotsenkaTovarovOtdannykhNaKomissiiuPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPereotsenkaTovarovOtdannykhNaKomissiiuPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPereotsenkaTovarovOtdannykhNaKomissiiuUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPereotsenkaTovarovOtdannykhNaKomissiiuUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPereotsenkaTovarovOtdannykhNaKomissiiuUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVvodNachalnykhOstatkovPoRaskhodamUSNPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVvodNachalnykhOstatkovPoRaskhodamUSNPost) Name() string {
	return "Post"
}
func (f FunctionDocumentVvodNachalnykhOstatkovPoRaskhodamUSNPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVvodNachalnykhOstatkovPoRaskhodamUSNUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVvodNachalnykhOstatkovPoRaskhodamUSNUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentVvodNachalnykhOstatkovPoRaskhodamUSNUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentGTDImportPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentGTDImportPost) Name() string {
	return "Post"
}
func (f FunctionDocumentGTDImportPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentGTDImportUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentGTDImportUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentGTDImportUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAktSverkiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAktSverkiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentAktSverkiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAktSverkiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAktSverkiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentAktSverkiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlaniruemoePostuplenieDenezhnykhSredstvPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlaniruemoePostuplenieDenezhnykhSredstvPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPlaniruemoePostuplenieDenezhnykhSredstvPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlaniruemoePostuplenieDenezhnykhSredstvUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlaniruemoePostuplenieDenezhnykhSredstvUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPlaniruemoePostuplenieDenezhnykhSredstvUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPreiskurantTsenNaKamniPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPreiskurantTsenNaKamniPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPreiskurantTsenNaKamniPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPreiskurantTsenNaKamniUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPreiskurantTsenNaKamniUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPreiskurantTsenNaKamniUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionPurchasePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionPurchasePost) Name() string {
	return "Post"
}
func (f FunctionPurchasePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionPurchaseUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionPurchaseUnpost) Name() string {
	return "Unpost"
}
func (f FunctionPurchaseUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSchetFakturaPoluchennyiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSchetFakturaPoluchennyiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentSchetFakturaPoluchennyiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSchetFakturaPoluchennyiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSchetFakturaPoluchennyiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentSchetFakturaPoluchennyiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAktKhimicheskogoAnalizaMetallaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAktKhimicheskogoAnalizaMetallaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentAktKhimicheskogoAnalizaMetallaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAktKhimicheskogoAnalizaMetallaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAktKhimicheskogoAnalizaMetallaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentAktKhimicheskogoAnalizaMetallaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSpisanieProsrochennykhSertifikatovPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSpisanieProsrochennykhSertifikatovPost) Name() string {
	return "Post"
}
func (f FunctionDocumentSpisanieProsrochennykhSertifikatovPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSpisanieProsrochennykhSertifikatovUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSpisanieProsrochennykhSertifikatovUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentSpisanieProsrochennykhSertifikatovUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakrytieAvansovPoGrafikuPlatezheiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakrytieAvansovPoGrafikuPlatezheiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentZakrytieAvansovPoGrafikuPlatezheiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakrytieAvansovPoGrafikuPlatezheiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakrytieAvansovPoGrafikuPlatezheiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentZakrytieAvansovPoGrafikuPlatezheiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAkkreditivPeredannyiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAkkreditivPeredannyiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentAkkreditivPeredannyiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAkkreditivPeredannyiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAkkreditivPeredannyiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentAkkreditivPeredannyiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInformatsionnoeSoobshcheniePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInformatsionnoeSoobshcheniePost) Name() string {
	return "Post"
}
func (f FunctionDocumentInformatsionnoeSoobshcheniePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInformatsionnoeSoobshchenieUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInformatsionnoeSoobshchenieUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentInformatsionnoeSoobshchenieUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnoeTrebovanieVystavlennoePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnoeTrebovanieVystavlennoePost) Name() string {
	return "Post"
}
func (f FunctionDocumentPlatezhnoeTrebovanieVystavlennoePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnoeTrebovanieVystavlennoeUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnoeTrebovanieVystavlennoeUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPlatezhnoeTrebovanieVystavlennoeUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentMarketingovaiaAktsiiaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentMarketingovaiaAktsiiaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentMarketingovaiaAktsiiaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentMarketingovaiaAktsiiaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentMarketingovaiaAktsiiaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentMarketingovaiaAktsiiaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOprosPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOprosPost) Name() string {
	return "Post"
}
func (f FunctionDocumentOprosPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOprosUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOprosUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentOprosUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionReassessmentPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionReassessmentPost) Name() string {
	return "Post"
}
func (f FunctionReassessmentPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionReassessmentUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionReassessmentUnpost) Name() string {
	return "Unpost"
}
func (f FunctionReassessmentUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentIzmeneniePravDostupaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentIzmeneniePravDostupaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentIzmeneniePravDostupaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentIzmeneniePravDostupaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentIzmeneniePravDostupaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentIzmeneniePravDostupaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionMoveInstancePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionMoveInstancePost) Name() string {
	return "Post"
}
func (f FunctionMoveInstancePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionMoveInstanceUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionMoveInstanceUnpost) Name() string {
	return "Unpost"
}
func (f FunctionMoveInstanceUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakrytieZaiavokNaRaskhodovanieSredstvPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakrytieZaiavokNaRaskhodovanieSredstvPost) Name() string {
	return "Post"
}
func (f FunctionDocumentZakrytieZaiavokNaRaskhodovanieSredstvPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakrytieZaiavokNaRaskhodovanieSredstvUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakrytieZaiavokNaRaskhodovanieSredstvUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentZakrytieZaiavokNaRaskhodovanieSredstvUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentABCKlassifikatsiiaPokupateleiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentABCKlassifikatsiiaPokupateleiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentABCKlassifikatsiiaPokupateleiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentABCKlassifikatsiiaPokupateleiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentABCKlassifikatsiiaPokupateleiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentABCKlassifikatsiiaPokupateleiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSvodnaiaInventarizatsiiaTovarovNaSkladePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSvodnaiaInventarizatsiiaTovarovNaSkladePost) Name() string {
	return "Post"
}
func (f FunctionDocumentSvodnaiaInventarizatsiiaTovarovNaSkladePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentKorrektirovkaRealizatsiiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentKorrektirovkaRealizatsiiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentKorrektirovkaRealizatsiiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentKorrektirovkaRealizatsiiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentKorrektirovkaRealizatsiiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentKorrektirovkaRealizatsiiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentDoverennostPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentDoverennostPost) Name() string {
	return "Post"
}
func (f FunctionDocumentDoverennostPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentDoverennostUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentDoverennostUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentDoverennostUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlanZapolneniiaVitrinPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlanZapolneniiaVitrinPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPlanZapolneniiaVitrinPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlanZapolneniiaVitrinUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlanZapolneniiaVitrinUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPlanZapolneniiaVitrinUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionReturnToManufacturingPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionReturnToManufacturingPost) Name() string {
	return "Post"
}
func (f FunctionReturnToManufacturingPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionReturnToManufacturingUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionReturnToManufacturingUnpost) Name() string {
	return "Unpost"
}
func (f FunctionReturnToManufacturingUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionWriteOffPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionWriteOffPost) Name() string {
	return "Post"
}
func (f FunctionWriteOffPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionWriteOffUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionWriteOffUnpost) Name() string {
	return "Unpost"
}
func (f FunctionWriteOffUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentsmsSoobshcheniePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentsmsSoobshcheniePost) Name() string {
	return "Post"
}
func (f FunctionDocumentsmsSoobshcheniePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentsmsSoobshchenieUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentsmsSoobshchenieUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentsmsSoobshchenieUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOplataOtPokupateliaPlatezhnoiKartoiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOplataOtPokupateliaPlatezhnoiKartoiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentOplataOtPokupateliaPlatezhnoiKartoiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOplataOtPokupateliaPlatezhnoiKartoiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOplataOtPokupateliaPlatezhnoiKartoiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentOplataOtPokupateliaPlatezhnoiKartoiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentNachislenieSpisanieBonusovPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentNachislenieSpisanieBonusovPost) Name() string {
	return "Post"
}
func (f FunctionDocumentNachislenieSpisanieBonusovPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentNachislenieSpisanieBonusovUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentNachislenieSpisanieBonusovUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentNachislenieSpisanieBonusovUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnoeTrebovaniePoluchennoePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnoeTrebovaniePoluchennoePost) Name() string {
	return "Post"
}
func (f FunctionDocumentPlatezhnoeTrebovaniePoluchennoePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnoeTrebovaniePoluchennoeUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnoeTrebovaniePoluchennoeUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPlatezhnoeTrebovaniePoluchennoeUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPost) Name() string {
	return "Post"
}
func (f FunctionDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOtchetPoFinMonitoringuPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOtchetPoFinMonitoringuPost) Name() string {
	return "Post"
}
func (f FunctionDocumentOtchetPoFinMonitoringuPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOtchetPoFinMonitoringuUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOtchetPoFinMonitoringuUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentOtchetPoFinMonitoringuUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentUstanovkaTsenNomenklaturyPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentUstanovkaTsenNomenklaturyPost) Name() string {
	return "Post"
}
func (f FunctionDocumentUstanovkaTsenNomenklaturyPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentUstanovkaTsenNomenklaturyUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentUstanovkaTsenNomenklaturyUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentUstanovkaTsenNomenklaturyUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPreiskurantNaSkupkuPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPreiskurantNaSkupkuPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPreiskurantNaSkupkuPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPreiskurantNaSkupkuUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPreiskurantNaSkupkuUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPreiskurantNaSkupkuUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPeredachaMaterialovVProizvodstvoPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPeredachaMaterialovVProizvodstvoPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPeredachaMaterialovVProizvodstvoPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPeredachaMaterialovVProizvodstvoUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPeredachaMaterialovVProizvodstvoUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPeredachaMaterialovVProizvodstvoUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVnutrenniiZakazPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVnutrenniiZakazPost) Name() string {
	return "Post"
}
func (f FunctionDocumentVnutrenniiZakazPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVnutrenniiZakazUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVnutrenniiZakazUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentVnutrenniiZakazUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZaiavkaNaPeremeshchenieTovarovPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZaiavkaNaPeremeshchenieTovarovPost) Name() string {
	return "Post"
}
func (f FunctionDocumentZaiavkaNaPeremeshchenieTovarovPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZaiavkaNaPeremeshchenieTovarovUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZaiavkaNaPeremeshchenieTovarovUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentZaiavkaNaPeremeshchenieTovarovUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVvodNachalnykhOstatkovPoFinMonitoringuPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVvodNachalnykhOstatkovPoFinMonitoringuPost) Name() string {
	return "Post"
}
func (f FunctionDocumentVvodNachalnykhOstatkovPoFinMonitoringuPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVvodNachalnykhOstatkovPoFinMonitoringuUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVvodNachalnykhOstatkovPoFinMonitoringuUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentVvodNachalnykhOstatkovPoFinMonitoringuUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentUdalitNariadZakazPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentUdalitNariadZakazPost) Name() string {
	return "Post"
}
func (f FunctionDocumentUdalitNariadZakazPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentUdalitNariadZakazUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentUdalitNariadZakazUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentUdalitNariadZakazUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentRestrukturizatsiiaDolgaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentRestrukturizatsiiaDolgaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentRestrukturizatsiiaDolgaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentRestrukturizatsiiaDolgaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentRestrukturizatsiiaDolgaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentRestrukturizatsiiaDolgaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAkkreditivPoluchennyiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAkkreditivPoluchennyiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentAkkreditivPoluchennyiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAkkreditivPoluchennyiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAkkreditivPoluchennyiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentAkkreditivPoluchennyiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPriemIzRemontaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPriemIzRemontaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPriemIzRemontaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPriemIzRemontaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPriemIzRemontaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPriemIzRemontaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentStornirovanieOtchetaKomissioneraOProdazhakhPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentStornirovanieOtchetaKomissioneraOProdazhakhPost) Name() string {
	return "Post"
}
func (f FunctionDocumentStornirovanieOtchetaKomissioneraOProdazhakhPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentStornirovanieOtchetaKomissioneraOProdazhakhUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentStornirovanieOtchetaKomissioneraOProdazhakhUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentStornirovanieOtchetaKomissioneraOProdazhakhUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentUstanovkaZnacheniiTochkiZakazaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentUstanovkaZnacheniiTochkiZakazaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentUstanovkaZnacheniiTochkiZakazaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentUstanovkaZnacheniiTochkiZakazaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentUstanovkaZnacheniiTochkiZakazaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentUstanovkaZnacheniiTochkiZakazaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPostuplenieDopRaskhodovPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPostuplenieDopRaskhodovPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPostuplenieDopRaskhodovPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPostuplenieDopRaskhodovUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPostuplenieDopRaskhodovUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPostuplenieDopRaskhodovUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAvansovyiOtchetPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAvansovyiOtchetPost) Name() string {
	return "Post"
}
func (f FunctionDocumentAvansovyiOtchetPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentAvansovyiOtchetUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentAvansovyiOtchetUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentAvansovyiOtchetUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentRegistratsiiaNaSaitePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentRegistratsiiaNaSaitePost) Name() string {
	return "Post"
}
func (f FunctionDocumentRegistratsiiaNaSaitePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentRegistratsiiaNaSaiteUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentRegistratsiiaNaSaiteUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentRegistratsiiaNaSaiteUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInkassovoePorucheniePoluchennoePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInkassovoePorucheniePoluchennoePost) Name() string {
	return "Post"
}
func (f FunctionDocumentInkassovoePorucheniePoluchennoePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInkassovoePorucheniePoluchennoeUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInkassovoePorucheniePoluchennoeUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentInkassovoePorucheniePoluchennoeUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVozvratTovarovOtPokupateliaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVozvratTovarovOtPokupateliaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentVozvratTovarovOtPokupateliaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVozvratTovarovOtPokupateliaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVozvratTovarovOtPokupateliaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentVozvratTovarovOtPokupateliaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakazPostavshchikuPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakazPostavshchikuPost) Name() string {
	return "Post"
}
func (f FunctionDocumentZakazPostavshchikuPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakazPostavshchikuUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakazPostavshchikuUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentZakazPostavshchikuUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentDokumentRaschetovSKontragentomPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentDokumentRaschetovSKontragentomPost) Name() string {
	return "Post"
}
func (f FunctionDocumentDokumentRaschetovSKontragentomPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentDokumentRaschetovSKontragentomUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentDokumentRaschetovSKontragentomUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentDokumentRaschetovSKontragentomUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentUstanovkaTsenNomenklaturyKontragentovPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentUstanovkaTsenNomenklaturyKontragentovPost) Name() string {
	return "Post"
}
func (f FunctionDocumentUstanovkaTsenNomenklaturyKontragentovPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentUstanovkaTsenNomenklaturyKontragentovUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentUstanovkaTsenNomenklaturyKontragentovUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentUstanovkaTsenNomenklaturyKontragentovUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentProtsentPoterPoDavaltsamPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentProtsentPoterPoDavaltsamPost) Name() string {
	return "Post"
}
func (f FunctionDocumentProtsentPoterPoDavaltsamPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentProtsentPoterPoDavaltsamUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentProtsentPoterPoDavaltsamUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentProtsentPoterPoDavaltsamUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnoePoruchenieIskhodiashcheePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnoePoruchenieIskhodiashcheePost) Name() string {
	return "Post"
}
func (f FunctionDocumentPlatezhnoePoruchenieIskhodiashcheePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlatezhnoePoruchenieIskhodiashcheeUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlatezhnoePoruchenieIskhodiashcheeUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPlatezhnoePoruchenieIskhodiashcheeUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentUstanovkaSkidokNomenklaturyPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentUstanovkaSkidokNomenklaturyPost) Name() string {
	return "Post"
}
func (f FunctionDocumentUstanovkaSkidokNomenklaturyPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentUstanovkaSkidokNomenklaturyUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentUstanovkaSkidokNomenklaturyUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentUstanovkaSkidokNomenklaturyUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionOutPayPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionOutPayPost) Name() string {
	return "Post"
}
func (f FunctionOutPayPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionOutPayUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionOutPayUnpost) Name() string {
	return "Unpost"
}
func (f FunctionOutPayUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSchetNaOplatuPostavshchikaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSchetNaOplatuPostavshchikaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentSchetNaOplatuPostavshchikaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSchetNaOplatuPostavshchikaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSchetNaOplatuPostavshchikaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentSchetNaOplatuPostavshchikaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentReestrSpetssviazPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentReestrSpetssviazPost) Name() string {
	return "Post"
}
func (f FunctionDocumentReestrSpetssviazPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentReestrSpetssviazUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentReestrSpetssviazUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentReestrSpetssviazUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionInitialInstancePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionInitialInstancePost) Name() string {
	return "Post"
}
func (f FunctionInitialInstancePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionInitialInstanceUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionInitialInstanceUnpost) Name() string {
	return "Unpost"
}
func (f FunctionInitialInstanceUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionPostingPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionPostingPost) Name() string {
	return "Post"
}
func (f FunctionPostingPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionPostingUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionPostingUnpost) Name() string {
	return "Unpost"
}
func (f FunctionPostingUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPereotsenkaTovarovPriniatykhNaKomissiiuPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPereotsenkaTovarovPriniatykhNaKomissiiuPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPereotsenkaTovarovPriniatykhNaKomissiiuPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPereotsenkaTovarovPriniatykhNaKomissiiuUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPereotsenkaTovarovPriniatykhNaKomissiiuUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPereotsenkaTovarovPriniatykhNaKomissiiuUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentElektronnoePismoPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentElektronnoePismoPost) Name() string {
	return "Post"
}
func (f FunctionDocumentElektronnoePismoPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentElektronnoePismoUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentElektronnoePismoUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentElektronnoePismoUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentObieiavlenieNaVznosNalichnymiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentObieiavlenieNaVznosNalichnymiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentObieiavlenieNaVznosNalichnymiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentObieiavlenieNaVznosNalichnymiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentObieiavlenieNaVznosNalichnymiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentObieiavlenieNaVznosNalichnymiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentReestrSchetovPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentReestrSchetovPost) Name() string {
	return "Post"
}
func (f FunctionDocumentReestrSchetovPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentReestrSchetovUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentReestrSchetovUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentReestrSchetovUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuPost) Name() string {
	return "Post"
}
func (f FunctionDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOtborTovarovPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOtborTovarovPost) Name() string {
	return "Post"
}
func (f FunctionDocumentOtborTovarovPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOtborTovarovUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOtborTovarovUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentOtborTovarovUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOtchetKomissioneraOProdazhakhPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOtchetKomissioneraOProdazhakhPost) Name() string {
	return "Post"
}
func (f FunctionDocumentOtchetKomissioneraOProdazhakhPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentOtchetKomissioneraOProdazhakhUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentOtchetKomissioneraOProdazhakhUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentOtchetKomissioneraOProdazhakhUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPreiskurantTsenNaTsvKamniPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPreiskurantTsenNaTsvKamniPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPreiskurantTsenNaTsvKamniPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPreiskurantTsenNaTsvKamniUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPreiskurantTsenNaTsvKamniUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPreiskurantTsenNaTsvKamniUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentTelemarketingPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentTelemarketingPost) Name() string {
	return "Post"
}
func (f FunctionDocumentTelemarketingPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentTelemarketingUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentTelemarketingUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentTelemarketingUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVozvratDavalcheskogoMetallaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVozvratDavalcheskogoMetallaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentVozvratDavalcheskogoMetallaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentVozvratDavalcheskogoMetallaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentVozvratDavalcheskogoMetallaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentVozvratDavalcheskogoMetallaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentRassylkaAnketPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentRassylkaAnketPost) Name() string {
	return "Post"
}
func (f FunctionDocumentRassylkaAnketPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentRassylkaAnketUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentRassylkaAnketUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentRassylkaAnketUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakazKlientaPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakazKlientaPost) Name() string {
	return "Post"
}
func (f FunctionDocumentZakazKlientaPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentZakazKlientaUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentZakazKlientaUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentZakazKlientaUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionArriveFromManufacturingPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionArriveFromManufacturingPost) Name() string {
	return "Post"
}
func (f FunctionArriveFromManufacturingPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionArriveFromManufacturingUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionArriveFromManufacturingUnpost) Name() string {
	return "Unpost"
}
func (f FunctionArriveFromManufacturingUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionArrivePost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionArrivePost) Name() string {
	return "Post"
}
func (f FunctionArrivePost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionArriveUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionArriveUnpost) Name() string {
	return "Unpost"
}
func (f FunctionArriveUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSchetFakturaVydannyiPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSchetFakturaVydannyiPost) Name() string {
	return "Post"
}
func (f FunctionDocumentSchetFakturaVydannyiPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentSchetFakturaVydannyiUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentSchetFakturaVydannyiUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentSchetFakturaVydannyiUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlanProdazhPoSalonamPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlanProdazhPoSalonamPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPlanProdazhPoSalonamPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPlanProdazhPoSalonamUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPlanProdazhPoSalonamUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPlanProdazhPoSalonamUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentStornirovanieOtchetaKomitentuOProdazhakhPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentStornirovanieOtchetaKomitentuOProdazhakhPost) Name() string {
	return "Post"
}
func (f FunctionDocumentStornirovanieOtchetaKomitentuOProdazhakhPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentStornirovanieOtchetaKomitentuOProdazhakhUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentStornirovanieOtchetaKomitentuOProdazhakhUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentStornirovanieOtchetaKomitentuOProdazhakhUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPeredachaVRemontPost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPeredachaVRemontPost) Name() string {
	return "Post"
}
func (f FunctionDocumentPeredachaVRemontPost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}

type FunctionDocumentPeredachaVRemontUnpost struct {
	PostingModeOperational Boolean `odata:"PostingModeOperational"`
}

func (FunctionDocumentPeredachaVRemontUnpost) Name() string {
	return "Unpost"
}
func (f FunctionDocumentPeredachaVRemontUnpost) Parameters() string {
	return "PostingModeOperational=" + f.PostingModeOperational.AsParameter()
}
