package emojiflags

import "testing"

func BenchmarkGetFlag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFlag("VNM")
	}
}

func BenchmarkGetFlag2Letter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFlag("VN")
	}
}

func BenchmarkGetFlagCIOC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFlag("GER")
	}
}

func BenchmarkGetFlagSpecial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFlag("GB-ENG")
	}
}

func BenchmarkGetFlagInvalid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFlag("INVALID")
	}
}

func BenchmarkGetFlagFuzzy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFlagFuzzy("VNM")
	}
}

func BenchmarkGetFlagFuzzyExact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFlagFuzzy("VN")
	}
}

func BenchmarkGetFlagFuzzyTypo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFlagFuzzy("VIETNM") // distance 3, should fail fast
	}
}

func BenchmarkGetFlagFuzzyClose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFlagFuzzy("GERM") // distance 1, should find
	}
}

func BenchmarkGetFlagFuzzyVariation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFlagFuzzy("USA") // distance 1, should find
	}
}

func BenchmarkLevenshtein(b *testing.B) {
	for i := 0; i < b.N; i++ {
		levenshtein("VIETNM", "VNM")
	}
}

func BenchmarkLevenshteinShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		levenshtein("USA", "US")
	}
}

func BenchmarkLevenshteinIdentical(b *testing.B) {
	for i := 0; i < b.N; i++ {
		levenshtein("VNM", "VNM")
	}
}
