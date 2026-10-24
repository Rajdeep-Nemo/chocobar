package chocobar

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/chocobar/internal"
)

func TestUint16_ValidInput(t *testing.T) {
	withInput(t, "12345")

	got, err := Uint16()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 12345 {
		t.Errorf("expected 12345, got %d", got)
	}
}

func TestUint16_Zero(t *testing.T) {
	withInput(t, "0")

	got, err := Uint16()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestUint16_MaxValue(t *testing.T) {
	withInput(t, "65535")

	got, err := Uint16()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 65535 {
		t.Errorf("expected 65535, got %d", got)
	}
}

func TestUint16_TrimsWhitespace(t *testing.T) {
	withInput(t, "   12345   ")

	got, err := Uint16()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 12345 {
		t.Errorf("expected 12345, got %d", got)
	}
}

func TestUint16_NegativeNumber(t *testing.T) {
	withInput(t, "-1")

	got, err := Uint16()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint16_ValueAboveRange(t *testing.T) {
	withInput(t, "65536")

	got, err := Uint16()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint16_InvalidInput(t *testing.T) {
	withInput(t, "not-a-number")

	got, err := Uint16()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint16_EmptyInput(t *testing.T) {
	withInput(t, "")

	got, err := Uint16()
	if err == nil {
		t.Fatalf("expected an error for empty input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint16_RetriesUntilValid(t *testing.T) {
	withInput(t, "abc", "-1", "50000")

	got, err := Uint16(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 50000 {
		t.Errorf("expected 50000, got %d", got)
	}
}

func TestUint16_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "abc", "65536", "-1")

	got, err := Uint16(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint16_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "50000")

	got, err := Uint16(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != 50000 {
		t.Errorf("expected 50000 from the 2nd attempt, got %d", got)
	}
}

func TestUint16_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "65536", "-1", "50000")

	got, err := Uint16(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != 50000 {
		t.Errorf("expected 50000, got %d", got)
	}
}

func TestUint16_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("50000"))
	t.Cleanup(internal.ResetInput)

	got, err := Uint16()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 50000 {
		t.Errorf("expected 50000, got %d", got)
	}
}

func TestUint16_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Uint16()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint16_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "12345")

	got, err := Uint16(Prompt("Enter a number: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 12345 {
		t.Errorf("expected 12345, got %d", got)
	}
}

func TestUint16_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "65536")

	_, err := Uint16()
	if err == nil {
		t.Fatal("expected an error for out-of-range input")
	}
	if !strings.Contains(err.Error(), "65536") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}
