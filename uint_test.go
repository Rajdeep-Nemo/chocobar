package sleepycat

import (
	"strconv"
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func TestUint_ValidInput(t *testing.T) {
	withInput(t, "12345")

	got, err := Uint()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 12345 {
		t.Errorf("expected 12345, got %d", got)
	}
}

func TestUint_Zero(t *testing.T) {
	withInput(t, "0")

	got, err := Uint()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestUint_TrimsWhitespace(t *testing.T) {
	withInput(t, "   12345   ")

	got, err := Uint()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 12345 {
		t.Errorf("expected 12345, got %d", got)
	}
}

func TestUint_NegativeNumber(t *testing.T) {
	withInput(t, "-1")

	got, err := Uint()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint_InvalidInput(t *testing.T) {
	withInput(t, "not-a-number")

	got, err := Uint()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint_EmptyInput(t *testing.T) {
	withInput(t, "")

	got, err := Uint()
	if err == nil {
		t.Fatalf("expected an error for empty input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint_MaxValue(t *testing.T) {
	max := strconv.FormatUint(^uint64(0), 10)
	withInput(t, max)

	got, err := Uint()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if uint64(got) != ^uint64(0) {
		t.Errorf("expected %d, got %d", ^uint64(0), got)
	}
}

func TestUint_ValueAboveRange(t *testing.T) {
	overflow := strconv.FormatUint(^uint64(0), 10) + "0"
	withInput(t, overflow)

	got, err := Uint()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint_RetriesUntilValid(t *testing.T) {
	withInput(t, "abc", "-1", "12345")

	got, err := Uint(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 12345 {
		t.Errorf("expected 12345, got %d", got)
	}
}

func TestUint_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "abc", "-1", "xyz")

	got, err := Uint(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "12345")

	got, err := Uint(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != 12345 {
		t.Errorf("expected 12345 from the 2nd attempt, got %d", got)
	}
}

func TestUint_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "-1", "abc", "100000")

	got, err := Uint(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != 100000 {
		t.Errorf("expected 100000, got %d", got)
	}
}

func TestUint_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("12345"))
	t.Cleanup(internal.ResetInput)

	got, err := Uint()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 12345 {
		t.Errorf("expected 12345, got %d", got)
	}
}

func TestUint_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Uint()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "123")

	got, err := Uint(Prompt("Enter a number: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123 {
		t.Errorf("expected 123, got %d", got)
	}
}

func TestUint_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "-1")

	_, err := Uint()
	if err == nil {
		t.Fatal("expected an error for negative input")
	}
	if !strings.Contains(err.Error(), "-1") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}
