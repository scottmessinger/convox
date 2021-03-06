package algoliasearch

import (
	"fmt"
	"os"
	"reflect"
)

// Settings is the structure returned by `GetSettigs` to ease the use of the
// index settings.
type Settings struct {
	// Indexing parameters
	AllowCompressionOfIntegerArray bool                `json:"allowCompressionOfIntegerArray"`
	AttributeForDistinct           string              `json:"attributeForDistinct"`
	AttributesForFaceting          []string            `json:"attributesForFaceting"`
	AttributesToIndex              []string            `json:"attributesToIndex"`
	CamelCaseAttributes            []string            `json:"camelCaseAttributes"`
	CustomRanking                  []string            `json:"customRanking"`
	DecompoundedAttributes         map[string][]string `json:"decompoundedAttributes"`
	KeepDiacriticsOnCharacters     string              `json:"keepDiacriticsOnCharacters"`
	NumericAttributesForFiltering  []string            `json:"numericAttributesForFiltering"`
	NumericAttributesToIndex       []string            `json:"numericAttributesToIndex"`
	PaginationLimitedTo            int                 `json:"paginationLimitedTo"`
	Ranking                        []string            `json:"ranking"`
	Replicas                       []string            `json:"replicas"`
	SearchableAttributes           []string            `json:"searchableAttributes"`
	SeparatorsToIndex              string              `json:"separatorsToIndex"`
	Slaves                         []string            `json:"slaves"`
	UnretrievableAttributes        []string            `json:"unretrievableAttributes"`

	// Query expansion
	DisableTypoToleranceOnAttributes []string `json:"disableTypoToleranceOnAttributes"`
	DisableTypoToleranceOnWords      []string `json:"disableTypoToleranceOnWords"`

	// Default query parameters (can be overridden at query-time)
	AdvancedSyntax                    bool        `json:"advancedSyntax"`
	AllowTyposOnNumericTokens         bool        `json:"allowTyposOnNumericTokens"`
	AttributesToHighlight             []string    `json:"attributesToHighlight"`
	AttributesToRetrieve              []string    `json:"attributesToRetrieve"`
	AttributesToSnippet               []string    `json:"attributesToSnippet"`
	Distinct                          interface{} `json:"distinct"` // float64 (actually an int) or bool
	EnableRules                       bool        `json:"enableRules"`
	HighlightPostTag                  string      `json:"highlightPostTag"`
	HighlightPreTag                   string      `json:"highlightPreTag"`
	HitsPerPage                       int         `json:"hitsPerPage"`
	IgnorePlurals                     interface{} `json:"ignorePlurals"` // []interface{} (actually a []string) or bool
	MaxFacetHits                      int         `json:"maxFacetHits"`
	MaxValuesPerFacet                 int         `json:"maxValuesPerFacet"`
	MinProximity                      int         `json:"minProximity"`
	MinWordSizefor1Typo               int         `json:"minWordSizefor1Typo"`
	MinWordSizefor2Typos              int         `json:"minWordSizefor2Typos"`
	OptionalWords                     []string    `json:"optionalWords"`
	QueryType                         string      `json:"queryType"`
	RemoveStopWords                   interface{} `json:"removeStopWords"` // []interface{} (actually a []string) or bool
	QueryLanguages                    []string    `json:"queryLanguages"`
	RemoveWordsIfNoResults            string      `json:"removeWordsIfNoResults"`
	ReplaceSynonymsInHighlight        bool        `json:"replaceSynonymsInHighlight"`
	ResponseFields                    []string    `json:"responseFields"`
	RestrictHighlightAndSnippetArrays bool        `json:"restrictHighlightAndSnippetArrays"`
	SnippetEllipsisText               string      `json:"snippetEllipsisText"`
	SortFacetValuesBy                 string      `json:"sortFacetValuesBy"`
	TypoTolerance                     string      `json:"typoTolerance"`
}

// clean sets the nil `interface{}` fields of any `Settings struct` generated
// by `GetSettings`.
func (s *Settings) clean() {
	if s.Distinct == nil {
		s.Distinct = false
	}

	if s.IgnorePlurals == nil {
		s.IgnorePlurals = false
	}

	if s.RemoveStopWords == nil {
		s.RemoveStopWords = false
	}

	if s.TypoTolerance == "" {
		s.TypoTolerance = "true"
	}

	if s.MaxFacetHits == 0 {
		s.MaxFacetHits = 10
	}

	if s.SortFacetValuesBy == "" {
		s.SortFacetValuesBy = "count"
	}
}

// ToMap produces a `Map` corresponding to the `Settings struct`. It should
// only be used when it's needed to pass a `Settings struct` to `SetSettings`,
// typically when one needs to copy settings between two indices.
func (s *Settings) ToMap() Map {
	// Guarantee that interface{}-typed fields and default values are correctly
	// set.
	s.clean()

	// Add all fields except:
	//  - Distinct float64 or bool
	//  - IgnorePlurals []interface{} or bool
	//  - RemoveStopWords []interface{} or bool
	m := Map{
		// Indexing parameters
		"allowCompressionOfIntegerArray": s.AllowCompressionOfIntegerArray,
		"attributeForDistinct":           s.AttributeForDistinct,
		"attributesForFaceting":          s.AttributesForFaceting,
		"attributesToIndex":              s.AttributesToIndex,
		"camelCaseAttributes":            s.CamelCaseAttributes,
		"customRanking":                  s.CustomRanking,
		"decompoundedAttributes":         s.DecompoundedAttributes,
		"keepDiacriticsOnCharacters":     s.KeepDiacriticsOnCharacters,
		"numericAttributesForFiltering":  s.NumericAttributesForFiltering,
		"numericAttributesToIndex":       s.NumericAttributesToIndex,
		"paginationLimitedTo":            s.PaginationLimitedTo,
		"ranking":                        s.Ranking,
		"replicas":                       s.Replicas,
		"searchableAttributes":           s.SearchableAttributes,
		"separatorsToIndex":              s.SeparatorsToIndex,
		"slaves":                         s.Slaves,
		"unretrievableAttributes":        s.UnretrievableAttributes,

		// Query expansion
		"disableTypoToleranceOnAttributes": s.DisableTypoToleranceOnAttributes,
		"disableTypoToleranceOnWords":      s.DisableTypoToleranceOnWords,

		// Default query parameters (can be overridden at query-time)
		"advancedSyntax":                    s.AdvancedSyntax,
		"allowTyposOnNumericTokens":         s.AllowTyposOnNumericTokens,
		"attributesToHighlight":             s.AttributesToHighlight,
		"attributesToRetrieve":              s.AttributesToRetrieve,
		"attributesToSnippet":               s.AttributesToSnippet,
		"enableRules":                       s.EnableRules,
		"highlightPostTag":                  s.HighlightPostTag,
		"highlightPreTag":                   s.HighlightPreTag,
		"hitsPerPage":                       s.HitsPerPage,
		"ignorePlurals":                     s.IgnorePlurals,
		"maxFacetHits":                      s.MaxFacetHits,
		"maxValuesPerFacet":                 s.MaxValuesPerFacet,
		"minProximity":                      s.MinProximity,
		"minWordSizefor1Typo":               s.MinWordSizefor1Typo,
		"minWordSizefor2Typos":              s.MinWordSizefor2Typos,
		"optionalWords":                     s.OptionalWords,
		"queryLanguages":                    s.QueryLanguages,
		"queryType":                         s.QueryType,
		"removeWordsIfNoResults":            s.RemoveWordsIfNoResults,
		"replaceSynonymsInHighlight":        s.ReplaceSynonymsInHighlight,
		"responseFields":                    s.ResponseFields,
		"restrictHighlightAndSnippetArrays": s.RestrictHighlightAndSnippetArrays,
		"snippetEllipsisText":               s.SnippetEllipsisText,
		"sortFacetValuesBy":                 s.SortFacetValuesBy,
		"typoTolerance":                     s.TypoTolerance,
	}

	// Remove empty string slices to avoid creating null-valued fields in the
	// JSON settings sent to the API
	var sliceAttributesToRemove []string

	for attr, value := range m {
		switch v := value.(type) {
		case []string:
			if len(v) == 0 {
				sliceAttributesToRemove = append(sliceAttributesToRemove, attr)
			}
		}
	}

	for _, attr := range sliceAttributesToRemove {
		delete(m, attr)
	}

	// Handle `Distinct` separately as it may be either a `bool` or a `float64`
	// which is in fact a `int`.
	switch v := s.Distinct.(type) {
	case bool:
		m["distinct"] = v
	case float64:
		m["distinct"] = int(v)
	}

	// Handle `IgnorePlurals` separately as it may be either a `bool` or a
	// `[]interface{}` which is in fact a `[]string`.
	switch v := s.IgnorePlurals.(type) {

	case bool:
		m["ignorePlurals"] = v

	case []interface{}:
		var languages []string
		for _, itf := range v {
			lang, ok := itf.(string)
			if ok {
				languages = append(languages, lang)
			} else {
				fmt.Fprintln(os.Stderr, "Settings.ToMap(): `ignorePlurals` slice doesn't only contain strings")
			}
		}
		if len(languages) > 0 {
			m["ignorePlurals"] = languages
		}

	default:
		fmt.Fprintf(os.Stderr, "Settings.ToMap(): Wrong type for `ignorePlurals`: %v\n", reflect.TypeOf(s.IgnorePlurals))

	}

	// Handle `RemoveStopWords` separately as it may be either a `bool` or a
	// `[]interface{}` which is in fact a `[]string`.
	switch v := s.RemoveStopWords.(type) {

	case bool:
		m["removeStopWords"] = v

	case []interface{}:
		var languages []string
		for _, itf := range v {
			lang, ok := itf.(string)
			if ok {
				languages = append(languages, lang)
			} else {
				fmt.Fprintln(os.Stderr, "Settings.ToMap(): `removeStopWords` slice doesn't only contain strings")
			}
		}
		if len(languages) > 0 {
			m["removeStopWords"] = languages
		}

	default:
		fmt.Fprintf(os.Stderr, "Settings.ToMap(): Wrong type for `removeStopWords`: %v\n", reflect.TypeOf(s.RemoveStopWords))

	}

	return m
}
