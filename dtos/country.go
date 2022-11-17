package dtos

type Data struct {
	Countries        []CountryNew
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

type CountryNew struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Tld         []string `json:"tld"`
	Cca2        string   `json:"cca2"`
	Ccn3        string   `json:"ccn3"`
	Cca3        string   `json:"cca3"`
	Cioc        string   `json:"cioc"`
	Independent bool     `json:"independent"`
	Status      string   `json:"status"`
	UnMember    bool     `json:"unMember"`
	Idd         struct {
		Root     string   `json:"root"`
		Suffixes []string `json:"suffixes"`
	} `json:"idd"`
	Capital      []string `json:"capital"`
	AltSpellings []string `json:"altSpellings"`
	Region       string   `json:"region"`
	Subregion    string   `json:"subregion"`
	Languages    struct {
		Deu string `json:"deu"`
	} `json:"languages"`
	Translations struct {
		Ara struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ara"`
		Bre struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"bre"`
		Ces struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ces"`
		Cym struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"cym"`
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Est struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"est"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hrv struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hrv"`
		Hun struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hun"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Jpn struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"jpn"`
		Kor struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"kor"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Per struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"per"`
		Pol struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"pol"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Slk struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"slk"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Swe struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"swe"`
		Tur struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"tur"`
		Urd struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"urd"`
		Zho struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"zho"`
	} `json:"translations"`
	Latlng     []float64 `json:"latlng"`
	Landlocked bool      `json:"landlocked"`
	Borders    []string  `json:"borders"`
	Area       float64   `json:"area"`
	Demonyms   struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
		Fra struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"fra"`
	} `json:"demonyms"`
	Flag string `json:"flag"`
	Maps struct {
		GoogleMaps     string `json:"googleMaps"`
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
	Population int `json:"population"`
	Gini       struct {
		Field1 float64 `json:"2016"`
	} `json:"gini"`
	Fifa string `json:"fifa"`
	Car  struct {
		Signs []string `json:"signs"`
		Side  string   `json:"side"`
	} `json:"car"`
	Timezones  []string `json:"timezones"`
	Continents []string `json:"continents"`
	Flags      struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags"`
	CoatOfArms struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"coatOfArms"`
	StartOfWeek string `json:"startOfWeek"`
	CapitalInfo struct {
		Latlng []float64 `json:"latlng"`
	} `json:"capitalInfo"`
	PostalCode struct {
		Format string `json:"format"`
		Regex  string `json:"regex"`
	} `json:"postalCode"`
}

type SpecialCountry struct {
	Name     string
	Code     string
	FlagCode string
}
