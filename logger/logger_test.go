package logger

import (
	"fmt"
	"testing"
)

func TestMustBePrintedEnabledProperty(t *testing.T) {
	var tests = []struct {
		enabled        bool
		expectedResult bool
	}{
		{true, true},
		{false, false},
	}

	for _, testData := range tests {
		testname := fmt.Sprintf("Enabled %t, Must be printed %t", testData.enabled, testData.expectedResult)
		t.Run(testname, func(t *testing.T) {
			testLogger := New(testData.enabled, DebugLevel)
			result := testLogger.mustBePrinted(DebugLevel)

			if result != testData.expectedResult {
				t.Errorf("Result was %t, but we expected %t", result, testData.expectedResult)
			}
		})
	}
}

func TestMustBePrintedLevels(t *testing.T) {
	var tests = []struct {
		loggerLevel, inputLevel Level
		expectedResult          bool
	}{
		{DebugLevel, DebugLevel, true},
		{DebugLevel, ErrorLevel, true},
		{ErrorLevel, DebugLevel, false},
	}

	for _, testData := range tests {
		testname := fmt.Sprintf("LoggerLevel %d, InputLevel %d", testData.loggerLevel, testData.inputLevel)
		t.Run(testname, func(t *testing.T) {
			testLogger := New(true, testData.loggerLevel)
			result := testLogger.mustBePrinted(testData.inputLevel)

			if result != testData.expectedResult {
				t.Errorf("Result was %t, but we expected %t", result, testData.expectedResult)
			}
		})
	}
}
