package chocobar

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/chocobar/internal"
)

func TestInt8_ValidInput(t *testing.T) {
	withInput(t, "42")

	got, err := Int8()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 42 {
		t.Errorf("expected 42, got %d", got)
	}
}

func TestInt8_NegativeNumber(t *testing.T) {
	withInput(t, "-42")

	got, err := Int8()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != -42 {
		t.Errorf("expected -42, got %d", got)
	}
}

func TestInt8_MinValue(t *testing.T) {
	withInput(t, "-128")

	got, err := Int8()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != -128 {
		t.Errorf("expected -128, got %d", got)
	}
}

func TestInt8_MaxValue(t *testing.T) {
	withInput(t, "127")

	got, err := Int8()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 127 {
		t.Errorf("expected 127, got %d", got)
	}
}

func TestInt8_TrimsWhitespace(t *testing.T) {
	withInput(t, "   15   ")

	got, err := Int8()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 15 {
		t.Errorf("expected 15, got %d", got)
	}
}

func TestInt8_InvalidInput(t *testing.T) {
	withInput(t, "not-a-number")

	got, err := Int8()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt8_EmptyInput(t *testing.T) {
	withInput(t, "")

	got, err := Int8()
	if err == nil {
		t.Fatalf("expected an error for empty input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt8_ValueAboveRange(t *testing.T) {
	withInput(t, "128")

	got, err := Int8()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt8_ValueBelowRange(t *testing.T) {
	withInput(t, "-129")

	got, err := Int8()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt8_RetriesUntilValid(t *testing.T) {
	withInput(t, "abc", "128", "9")

	got, err := Int8(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 9 {
		t.Errorf("expected 9, got %d", got)
	}
}

func TestInt8_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "abc", "128", "-129")

	got, err := Int8(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt8_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "77")

	got, err := Int8(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != 77 {
		t.Errorf("expected 77 from the 2nd attempt, got %d", got)
	}
}

func TestInt8_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "128", "-129", "100")

	got, err := Int8(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != 100 {
		t.Errorf("expected 100, got %d", got)
	}
}

func TestInt8_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("55"))
	t.Cleanup(internal.ResetInput)

	got, err := Int8()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 55 {
		t.Errorf("expected 55, got %d", got)
	}
}

func TestInt8_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Int8()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt8_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "21")

	got, err := Int8(Prompt("Enter a number: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 21 {
		t.Errorf("expected 21, got %d", got)
	}
}

func TestInt8_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "128")

	_, err := Int8()
	if err == nil {
		t.Fatal("expected an error for out-of-range input")
	}
	if !strings.Contains(err.Error(), "128") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}
