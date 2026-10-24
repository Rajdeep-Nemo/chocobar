package chocobar

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/chocobar/internal"
)

func TestInt32_ValidInput(t *testing.T) {
	withInput(t, "123456789")

	got, err := Int32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789, got %d", got)
	}
}

func TestInt32_NegativeNumber(t *testing.T) {
	withInput(t, "-123456789")

	got, err := Int32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != -123456789 {
		t.Errorf("expected -123456789, got %d", got)
	}
}

func TestInt32_MinValue(t *testing.T) {
	withInput(t, "-2147483648")

	got, err := Int32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != -2147483648 {
		t.Errorf("expected -2147483648, got %d", got)
	}
}

func TestInt32_MaxValue(t *testing.T) {
	withInput(t, "2147483647")

	got, err := Int32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 2147483647 {
		t.Errorf("expected 2147483647, got %d", got)
	}
}

func TestInt32_TrimsWhitespace(t *testing.T) {
	withInput(t, "   123456789   ")

	got, err := Int32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789, got %d", got)
	}
}

func TestInt32_InvalidInput(t *testing.T) {
	withInput(t, "not-a-number")

	got, err := Int32()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt32_EmptyInput(t *testing.T) {
	withInput(t, "")

	got, err := Int32()
	if err == nil {
		t.Fatalf("expected an error for empty input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt32_ValueAboveRange(t *testing.T) {
	withInput(t, "2147483648")

	got, err := Int32()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt32_ValueBelowRange(t *testing.T) {
	withInput(t, "-2147483649")

	got, err := Int32()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt32_RetriesUntilValid(t *testing.T) {
	withInput(t, "abc", "2147483648", "123456")

	got, err := Int32(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 123456 {
		t.Errorf("expected 123456, got %d", got)
	}
}

func TestInt32_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "abc", "2147483648", "-2147483649")

	got, err := Int32(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt32_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "123456")

	got, err := Int32(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != 123456 {
		t.Errorf("expected 123456 from the 2nd attempt, got %d", got)
	}
}

func TestInt32_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "2147483648", "-2147483649", "100000")

	got, err := Int32(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != 100000 {
		t.Errorf("expected 100000, got %d", got)
	}
}

func TestInt32_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("123456789"))
	t.Cleanup(internal.ResetInput)

	got, err := Int32()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789, got %d", got)
	}
}

func TestInt32_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Int32()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt32_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "123")

	got, err := Int32(Prompt("Enter a number: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123 {
		t.Errorf("expected 123, got %d", got)
	}
}

func TestInt32_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "2147483648")

	_, err := Int32()
	if err == nil {
		t.Fatal("expected an error for out-of-range input")
	}
	if !strings.Contains(err.Error(), "2147483648") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}
