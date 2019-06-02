package conformance

import (
	"testing"
)

func TestLanguageTypesBoolean(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	testCases := []testCase{
		{"S8.3_A1_T1.js", true, true},
		{"S8.3_A1_T2.js", true, true},
		{"S8.3_A2.1.js", true, true},
		{"S8.3_A2.2.js", true, true},
		{"S8.3_A3.js", true, true},
	}
	runTable(t, "test262/test/language/types/boolean", testCases)
}

func TestLanguageTypesList(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	testCases := []testCase{
		{"S8.8_A2_T1.js", true, true},
		{"S8.8_A2_T2.js", true, true},
		{"S8.8_A2_T3.js", true, true},
	}
	runTable(t, "test262/test/language/types/list", testCases)
}

func TestLanguageTypesNull(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	testCases := []testCase{
		{"S8.2_A1_T1.js", true, true},
		{"S8.2_A1_T2.js", true, true},
		{"S8.2_A2.js", false, false},
		{"S8.2_A3.js", true, true},
	}
	runTable(t, "test262/test/language/types/null", testCases)
}

func TestLanguageTypesNumber(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	testCases := []testCase{
		{"8.5.1.js", true, true},
		{"S8.5_A1.js", true, true},
		{"S8.5_A10_T1.js", true, true},
		{"S8.5_A10_T2.js", true, true},
		{"S8.5_A11_T1.js", true, true},
		{"S8.5_A11_T2.js", true, true},
		{"S8.5_A12.1.js", true, true},
		{"S8.5_A12.2.js", true, true},
		{"S8.5_A13_T2.js", true, true},
		{"S8.5_A14_T1.js", true, true},
		{"S8.5_A14_T2.js", true, true},
		{"S8.5_A2.1.js", true, true},
		{"S8.5_A2.2.js", true, true},
		{"S8.5_A3.js", true, true},
		{"S8.5_A4_T1.js", true, true},
		{"S8.5_A4_T2.js", true, true},
		{"S8.5_A5.js", true, true},
		{"S8.5_A6.js", true, true},
		{"S8.5_A7.js", true, true},
		{"S8.5_A8.js", true, true},
		{"S8.5_A9.js", true, true},
	}
	runTable(t, "test262/test/language/types/number", testCases)
}

func TestLanguageTypesObject(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	testCases := []testCase{
		{"S8.6.1_A1.js", true, true},
		{"S8.6.1_A2.js", true, true},
		{"S8.6.1_A3.js", true, true},
		{"S8.6.2_A1.js", true, true},
		{"S8.6.2_A2.js", true, true},
		{"S8.6.2_A3.js", true, true},
		{"S8.6.2_A4.js", true, true},
		{"S8.6.2_A5_T1.js", true, true},
		{"S8.6.2_A5_T2.js", true, true},
		{"S8.6.2_A5_T3.js", true, true},
		{"S8.6.2_A5_T4.js", true, true},
		{"S8.6.2_A6.js", true, true},
		{"S8.6.2_A7.js", true, true},
		{"S8.6.2_A8.js", true, true},
		{"S8.6_A2_T1.js", true, true},
		{"S8.6_A2_T2.js", true, true},
		{"S8.6_A3_T1.js", true, true},
		{"S8.6_A3_T2.js", true, true},
		{"S8.6_A4_T1.js", true, true},
	}
	runTable(t, "test262/test/language/types/object", testCases)
}

func TestLanguageTypesReference(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	testCases := []testCase{
		{"8.7.2-1-s.js", true, true},
		{"8.7.2-2-s.js", true, true},
		{"8.7.2-3-1-s.js", true, true},
		{"8.7.2-3-a-1gs.js", true, true},
		{"8.7.2-3-a-2gs.js", true, true},
		{"8.7.2-3-s.js", true, true},
		{"8.7.2-4-s.js", true, true},
		{"8.7.2-5-s.js", true, true},
		{"8.7.2-6-s.js", true, true},
		{"8.7.2-7-s.js", true, true},
		{"8.7.2-8-s.js", true, true},
		{"S8.7.1_A1.js", true, true},
		{"S8.7.1_A2.js", true, true},
		{"S8.7.2_A1_T1.js", true, true},
		{"S8.7.2_A1_T2.js", true, true},
		{"S8.7.2_A2.js", true, true},
		{"S8.7.2_A3.js", true, true},
		{"S8.7_A1.js", true, true},
		{"S8.7_A2.js", true, true},
		{"S8.7_A3.js", true, true},
		{"S8.7_A4.js", true, true},
		{"S8.7_A5_T1.js", true, true},
		{"S8.7_A5_T2.js", true, true},
		{"S8.7_A6.js", true, true},
		{"S8.7_A7.js", true, true},
		{"get-value-prop-base-primitive-realm.js", true, true},
		{"get-value-prop-base-primitive.js", true, true},
		{"put-value-prop-base-primitive-realm.js", true, true},
		{"put-value-prop-base-primitive.js", true, true},
	}
	runTable(t, "test262/test/language/types/reference", testCases)
}

func TestLanguageTypesString(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	testCases := []testCase{
		{"S8.4_A1.js", true, true},
		{"S8.4_A10.js", true, true},
		{"S8.4_A11.js", true, true},
		{"S8.4_A12.js", true, true},
		{"S8.4_A13_T1.js", true, true},
		{"S8.4_A13_T2.js", false, false},
		{"S8.4_A13_T3.js", true, true},
		{"S8.4_A14_T1.js", false, false},
		{"S8.4_A14_T2.js", true, true},
		{"S8.4_A14_T3.js", true, true},
		{"S8.4_A2.js", true, true},
		{"S8.4_A3.js", true, true},
		{"S8.4_A4.js", true, true},
		{"S8.4_A5.js", true, true},
		{"S8.4_A6.1.js", true, true},
		{"S8.4_A6.2.js", true, true},
		{"S8.4_A7.1.js", true, true},
		{"S8.4_A7.2.js", true, true},
		{"S8.4_A7.3.js", true, true},
		{"S8.4_A7.4.js", true, true},
		{"S8.4_A8.js", true, true},
		{"S8.4_A9_T1.js", true, true},
		{"S8.4_A9_T2.js", true, true},
		{"S8.4_A9_T3.js", true, true},
	}
	runTable(t, "test262/test/language/types/string", testCases)
}

func TestLanguageTypesUndefined(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	testCases := []testCase{
		{"S8.1_A1_T1.js", true, true},
		{"S8.1_A1_T2.js", true, true},
		{"S8.1_A2_T1.js", true, true},
		{"S8.1_A2_T2.js", true, true},
		{"S8.1_A3_T1.js", true, true},
		{"S8.1_A3_T2.js", true, true},
		{"S8.1_A4.js", true, true},
		{"S8.1_A5.js", true, true},
	}
	runTable(t, "test262/test/language/types/undefined", testCases)
}
