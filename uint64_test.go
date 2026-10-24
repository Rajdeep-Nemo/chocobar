package chocobar

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/chocobar/internal"
)

func TestUint64_ValidInput(t *testing.T) {
	withInput(t, "9223372036854775807")

	got, err := Uint64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 9223372036854775807 {
		t.Errorf("expected 9223372036854775807, got %d", got)
	}
}

func TestUint64_Zero(t *testing.T) {
	withInput(t, "0")

	got, err := Uint64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestUint64_MaxValue(t *testing.T) {
	withInput(t, "18446744073709551615")

	got, err := Uint64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 18446744073709551615 {
		t.Errorf("expected 18446744073709551615, got %d", got)
	}
}

func TestUint64_TrimsWhitespace(t *testing.T) {
	withInput(t, "   123456789   ")

	got, err := Uint64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789, got %d", got)
	}
}

func TestUint64_NegativeNumber(t *testing.T) {
	withInput(t, "-1")

	got, err := Uint64()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint64_ValueAboveRange(t *testing.T) {
	withInput(t, "18446744073709551616")

	got, err := Uint64()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint64_InvalidInput(t *testing.T) {
	withInput(t, "not-a-number")

	got, err := Uint64()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint64_EmptyInput(t *testing.T) {
	withInput(t, "")

	got, err := Uint64()
	if err == nil {
		t.Fatalf("expected an error for empty input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint64_RetriesUntilValid(t *testing.T) {
	withInput(t, "abc", "-1", "10000000000000000000")

	got, err := Uint64(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 10000000000000000000 {
		t.Errorf("expected 10000000000000000000, got %d", got)
	}
}

func TestUint64_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "abc", "18446744073709551616", "-1")

	got, err := Uint64(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint64_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "10000000000000000000")

	got, err := Uint64(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != 10000000000000000000 {
		t.Errorf("expected 10000000000000000000 from the 2nd attempt, got %d", got)
	}
}

func TestUint64_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "18446744073709551616", "-1", "10000000000000000000")

	got, err := Uint64(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != 10000000000000000000 {
		t.Errorf("expected 10000000000000000000, got %d", got)
	}
}

func TestUint64_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("10000000000000000000"))
	t.Cleanup(internal.ResetInput)

	got, err := Uint64()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 10000000000000000000 {
		t.Errorf("expected 10000000000000000000, got %d", got)
	}
}

func TestUint64_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Uint64()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint64_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "123456789")

	got, err := Uint64(Prompt("Enter a number: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789, got %d", got)
	}
}

func TestUint64_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "18446744073709551616")

	_, err := Uint64()
	if err == nil {
		t.Fatal("expected an error for out-of-range input")
	}
	if !strings.Contains(err.Error(), "18446744073709551616") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}
