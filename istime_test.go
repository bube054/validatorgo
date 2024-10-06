package validatorgo

import "testing"

func TestIsTime(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsTimeOpts
		want   bool
	}{
		// Valid times with default hourFormat and default Mode
		{name: "Valid - with nil config", param1: "14:30", param2: nil, want: true},
		{name: "Valid - 24-hour format (HH:MM)", param1: "14:30", param2: &IsTimeOpts{}, want: true},
		{name: "Valid - 24-hour format (HH:MM)", param1: "08:59", param2: &IsTimeOpts{}, want: true},
		{name: "Valid - 24-hour format with seconds (HH:MM:SS)", param1: "08:15:00", param2: &IsTimeOpts{}, want: true},
		{name: "Valid - 24-hour format with seconds (HH:MM:SS)", param1: "18:15:25", param2: &IsTimeOpts{}, want: true},
		// Invalid times with default hourFormat and default Mode
		{name: "Invalid - Missing leading/ending zero in minutes (HH:M)", param1: "09:5", param2: &IsTimeOpts{}, want: false},
		{name: "Invalid - Hours exceed 23", param1: "25:30", param2: &IsTimeOpts{}, want: false},
		{name: "Invalid - Minutes exceed 59 (HH:MM:SS)", param1: "10:60:00", param2: &IsTimeOpts{}, want: false},
		{name: "Invalid - Seconds exceed 59 (HH:MM:SS)", param1: "19:15:60", param2: &IsTimeOpts{}, want: false},

		// Valid times with 12-hour format and default Mode
		{name: "Valid - 12-hour format with PM (HH:MM AM/PM)", param1: "02:45 PM", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: true},
		{name: "Valid - 12-hour format with lowercase AM (HH:MM am)", param1: "11:30 am", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: true},
		{name: "Valid - 12-hour format with lowercase PM (HH:MM pm)", param1: "08:45 pm", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: true},
		{name: "Valid - 12-hour format with uppercase AM (HH:MM AM)", param1: "10:30 AM", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: true},
		{name: "Valid - 12-hour format with seconds (HH:MM:SS pm)", param1: "09:45:20 pm", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: true},
		{name: "Valid - 12-hour format with seconds and AM (HH:MM:SS AM)", param1: "12:30:50 AM", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: true},
		// Invalid times with 12-hour format and default Mode
		{name: "Invalid - Missing leading zero in minutes (HH:M AM/PM)", param1: "02:4 PM", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: false},
		{name: "Invalid - Hours exceed 12 in 12-hour format", param1: "15:30 am", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: false},
		{name: "Invalid - Improper casing for PM", param1: "08:45 Pm", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: false},
		{name: "Invalid - Improper casing for AM", param1: "10:30 Am", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: false},
		{name: "Invalid - Minutes exceed 59 (HH:MM:SS pm)", param1: "09:60:20 pm", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: false},
		{name: "Invalid - Seconds exceed 59 (HH:MM:SS AM)", param1: "12:30:60 AM", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12}, want: false},

		// Valid times with 12-hour format and with seconds
		{name: "Valid - 12-hour format with seconds and PM", param1: "02:45:10 PM", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: true},
		{name: "Valid - 12-hour format with seconds and lowercase AM", param1: "11:30:15 am", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: true},
		{name: "Valid - 12-hour format with seconds and lowercase PM", param1: "08:45:20 pm", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: true},
		{name: "Valid - 12-hour format with seconds and uppercase AM", param1: "10:30:45 AM", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: true},
		{name: "Valid - 12-hour format with seconds and lowercase PM", param1: "09:45:50 pm", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: true},
		{name: "Valid - 12-hour format with seconds and uppercase AM", param1: "12:30:59 AM", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: true},
		// Invalid times with 12-hour format and with seconds
		{name: "Invalid - Improper casing for PM (pM)", param1: "02:45:10 pM", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: false},
		{name: "Invalid - Improper casing for AM (Am)", param1: "11:30:15 Am", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: false},
		{name: "Invalid - Missing seconds (HH:MM pm)", param1: "08:45 pm", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: false},
		{name: "Invalid - Missing seconds (HH:MM AM)", param1: "10:30 AM", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: false},
		{name: "Invalid - Minutes exceed 59 (HH:MM:SS pm)", param1: "09:60:50 pm", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: false},
		{name: "Invalid - Seconds exceed 59 (HH:MM:SS AM)", param1: "12:30:60 AM", param2: &IsTimeOpts{HourFormat: IsTimeOptsHourFormat12, Mode: IsTimeOptsModeWithSeconds}, want: false},

		// Valid times with 24-hour format and with seconds
		{name: "Valid - 24-hour format with seconds (HH:MM:SS)", param1: "14:30:00", param2: &IsTimeOpts{Mode: IsTimeOptsModeWithSeconds}, want: true},
		{name: "Valid - 24-hour format with seconds (HH:MM:SS)", param1: "08:59:10", param2: &IsTimeOpts{Mode: IsTimeOptsModeWithSeconds}, want: true},
		{name: "Valid - 24-hour format with seconds (HH:MM:SS)", param1: "08:15:20", param2: &IsTimeOpts{Mode: IsTimeOptsModeWithSeconds}, want: true},
		{name: "Valid - 24-hour format with seconds (HH:MM:SS)", param1: "18:15:55", param2: &IsTimeOpts{Mode: IsTimeOptsModeWithSeconds}, want: true},
		// Invalid times with 24-hour format and with seconds
		{name: "Invalid - Missing seconds in 24-hour format (HH:MM)", param1: "09:50", param2: &IsTimeOpts{Mode: IsTimeOptsModeWithSeconds}, want: false},
		{name: "Invalid - Hours exceed 23", param1: "25:30", param2: &IsTimeOpts{Mode: IsTimeOptsModeWithSeconds}, want: false},
		{name: "Invalid - Minutes exceed 59 (HH:MM)", param1: "10:60", param2: &IsTimeOpts{Mode: IsTimeOptsModeWithSeconds}, want: false},
		{name: "Invalid - Hours exceed 23 (HH:MM)", param1: "25:15", param2: &IsTimeOpts{Mode: IsTimeOptsModeWithSeconds}, want: false},

		// Wrong inputs
		{name: "Invalid Hour format", param1: "10:60:00", param2: &IsTimeOpts{HourFormat: "xyz"}, want: false},
		{name: "Invalid mode format", param1: "10:60:00", param2: &IsTimeOpts{Mode: "xyz"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsTime(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
