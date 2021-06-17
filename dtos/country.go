package dtos

type Data struct {
	Countries        []Country
	SpecialCountries []SpecialCountry
}

type Country struct {
	Name           string    `json:"name"`
	TopLevelDomain []string  `json:"topLevelDomain"`
	Alpha2Code     string    `json:"alpha2Code"`
	Alpha3Code     string    `json:"alpha3Code"`
	CallingCodes   []string  `json:"callingCodes"`
	Capital        string    `json:"capital"`
	AltSpellings   []string  `json:"altSpellings"`
	Region         string    `json:"region"`
	Subregion      string    `json:"subregion"`
	Population     int       `json:"population"`
	Latlng         []float64 `json:"latlng"`
	Demonym        string    `json:"demonym"`
	Area           *float64  `json:"area"`
	Gini           *float64  `json:"gini"`
	Timezones      []string  `json:"timezones"`
	Borders        []string  `json:"borders"`
	NativeName     string    `json:"nativeName"`
	NumericCode    *string   `json:"numericCode"`
	Currencies     []struct {
		Code   *string `json:"code"`
		Name   *string `json:"name"`
		Symbol *string `json:"symbol"`
	} `json:"currencies"`
	Languages []struct {
		Iso6391    *string `json:"iso639_1"`
		Iso6392    string  `json:"iso639_2"`
		Name       string  `json:"name"`
		NativeName string  `json:"nativeName"`
	} `json:"languages"`
	Translations struct {
		De *string `json:"de"`
		Es *string `json:"es"`
		Fr *string `json:"fr"`
		Ja *string `json:"ja"`
		It *string `json:"it"`
		Br string  `json:"br"`
		Pt string  `json:"pt"`
		Nl *string `json:"nl"`
		Hr *string `json:"hr"`
		Fa string  `json:"fa"`
	} `json:"translations"`
	Flag          string `json:"flag"`
	RegionalBlocs []struct {
		Acronym       string   `json:"acronym"`
		Name          string   `json:"name"`
		OtherAcronyms []string `json:"otherAcronyms"`
		OtherNames    []string `json:"otherNames"`
	} `json:"regionalBlocs"`
	Cioc *string `json:"cioc"`
}

type SpecialCountry struct {
	Name string
	Code string
}
