package booking

// É parênteses que coloca quando importa...
import (
	"time"
	"fmt"
)
 
// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(t)
}


// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
    if t.Hour() >= 12 && t.Hour() < 18{
        return true
    }

    return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// O formato é de acordo com a data que está sendo recebida pela string
    t, _ := time.Parse("1/2/2006 15:04:05", date)

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), time.Month.String(t.Month()), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    layout := "2006-01-02 15:04:05 -0700 MST"
	bday := "2024-09-15 00:00:00 +0000 UTC"
    
    // Analise a string para criar um objeto time.Time
    t, _ := time.Parse(layout, bday)

	return t
}
