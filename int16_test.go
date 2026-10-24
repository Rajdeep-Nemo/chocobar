package chocobar

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/chocobar/internal"
)

func TestInt16_ValidInput(t *testing.T) {
	withInput(t, "12345")

	got, err := Int16()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 12345 {
		t.Errorf("expected 12345, got %d", got)
	}
}

func TestInt16_NegativeNumber(t *testing.T) {
	withInput(t, "-12345")

	got, err := Int16()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != -12345 {
		t.Errorf("expected -12345, got %d", got)
	}
}

func TestInt16_MinValue(t *testing.T) {
	withInput(t, "-32768")

	got, err := Int16()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != -32768 {
		t.Errorf("expected -32768, got %d", got)
	}
}

func TestInt16_MaxValue(t *testing.T) {
	withInput(t, "32767")

	got, err := Int16()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 32767 {
		t.Errorf("expected 32767, got %d", got)
	}
}

func TestInt16_TrimsWhitespace(t *testing.T) {
	withInput(t, "   12345   ")

	got, err := Int16()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 12345 {
		t.Errorf("expected 12345, got %d", got)
	}
}

func TestInt16_InvalidInput(t *testing.T) {
	withInput(t, "not-a-number")

	got, err := Int16()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt16_EmptyInput(t *testing.T) {
	withInput(t, "")

	got, err := Int16()
	if err == nil {
		t.Fatalf("expected an error for empty input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt16_ValueAboveRange(t *testing.T) {
	withInput(t, "32768")

	got, err := Int16()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt16_ValueBelowRange(t *testing.T) {
	withInput(t, "-32769")

	got, err := Int16()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt16_RetriesUntilValid(t *testing.T) {
	withInput(t, "abc", "32768", "1234")

	got, err := Int16(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 1234 {
		t.Errorf("expected 1234, got %d", got)
	}
}

func TestInt16_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "abc", "32768", "-32769")

	got, err := Int16(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt16_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "30000")

	got, err := Int16(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != 30000 {
		t.Errorf("expected 30000 from the 2nd attempt, got %d", got)
	}
}

func TestInt16_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "32768", "-32769", "25000")

	got, err := Int16(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != 25000 {
		t.Errorf("expected 25000, got %d", got)
	}
}

func TestInt16_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("12345"))
	t.Cleanup(internal.ResetInput)

	got, err := Int16()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 12345 {
		t.Errorf("expected 12345, got %d", got)
	}
}

func TestInt16_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Int16()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt16_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "123")

	got, err := Int16(Prompt("Enter a number: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123 {
		t.Errorf("expected 123, got %d", got)
	}
}

func TestInt16_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "32768")

	_, err := Int16()
	if err == nil {
		t.Fatal("expected an error for out-of-range input")
	}
	if !strings.Contains(err.Error(), "32768") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}
