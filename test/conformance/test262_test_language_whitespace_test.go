package conformance

import "testing"

func TestLanguageWhiteSpace(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	testCases := []testCase{
		{"S7.2_A2.1_T2.js", true, true},
		{"S7.2_A2.2_T2.js", true, true},
		{"S7.2_A2.3_T2.js", true, true},
		{"S7.2_A2.4_T2.js", true, true},
		{"S7.2_A2.5_T2.js", true, true},
		{"S7.2_A3.1_T2.js", true, true},
		{"S7.2_A3.2_T2.js", true, true},
		{"S7.2_A3.3_T2.js", true, true},
		{"S7.2_A3.4_T2.js", true, true},
		{"S7.2_A3.5_T2.js", true, true},
		{"S7.2_A4.1_T2.js", true, true},
		{"S7.2_A4.2_T2.js", true, true},
		{"S7.2_A4.3_T2.js", true, true},
		{"S7.2_A4.4_T2.js", true, true},
		{"S7.2_A4.5_T2.js", true, true},
		{"S7.2_A5_T1.js", true, true},
		{"S7.2_A5_T2.js", true, true},
		{"S7.2_A5_T3.js", true, true},
		{"S7.2_A5_T4.js", true, true},
		{"S7.2_A5_T5.js", true, true},
		{"between-form-feed.js", true, true},
		{"between-horizontal-tab.js", true, true},
		{"between-nbsp.js", true, true},
		{"between-space.js", true, true},
		{"between-vertical-tab.js", true, true},
		{"comment-multi-form-feed.js", true, true},
		{"comment-multi-horizontal-tab.js", true, true},
		{"comment-multi-nbsp.js", true, true},
		{"comment-multi-space.js", true, true},
		{"comment-multi-vertical-tab.js", true, true},
		{"comment-single-form-feed.js", true, true},
		{"comment-single-horizontal-tab.js", true, true},
		{"comment-single-nbsp.js", true, true},
		{"comment-single-space.js", true, true},
		{"comment-single-vertical-tab.js", true, true},
		{"mongolian-vowel-separator-eval.js", true, true},
		{"mongolian-vowel-separator.js", true, true},
		{"string-form-feed.js", true, true},
		{"string-horizontal-tab.js", true, true},
		{"string-nbsp.js", true, true},
		{"string-space.js", true, true},
		{"string-vertical-tab.js", true, true},
	}
	runTable(t, "test262/test/language/white-space", testCases)
}
