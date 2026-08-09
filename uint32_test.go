package sleepycat

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func TestUint32_ValidInput(t *testing.T) {
	withInput(t, "123456789")

	got, err := Uint32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789, got %d", got)
	}
}

func TestUint32_Zero(t *testing.T) {
	withInput(t, "0")

	got, err := Uint32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestUint32_MaxValue(t *testing.T) {
	withInput(t, "4294967295")

	got, err := Uint32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 4294967295 {
		t.Errorf("expected 4294967295, got %d", got)
	}
}

func TestUint32_TrimsWhitespace(t *testing.T) {
	withInput(t, "   123456789   ")

	got, err := Uint32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789, got %d", got)
	}
}

func TestUint32_NegativeNumber(t *testing.T) {
	withInput(t, "-1")

	got, err := Uint32()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint32_ValueAboveRange(t *testing.T) {
	withInput(t, "4294967296")

	got, err := Uint32()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint32_InvalidInput(t *testing.T) {
	withInput(t, "not-a-number")

	got, err := Uint32()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint32_EmptyInput(t *testing.T) {
	withInput(t, "")

	got, err := Uint32()
	if err == nil {
		t.Fatalf("expected an error for empty input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint32_RetriesUntilValid(t *testing.T) {
	withInput(t, "abc", "-1", "3000000000")

	got, err := Uint32(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 3000000000 {
		t.Errorf("expected 3000000000, got %d", got)
	}
}

func TestUint32_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "abc", "4294967296", "-1")

	got, err := Uint32(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint32_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "3000000000")

	got, err := Uint32(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != 3000000000 {
		t.Errorf("expected 3000000000 from the 2nd attempt, got %d", got)
	}
}

func TestUint32_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "4294967296", "-1", "3000000000")

	got, err := Uint32(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != 3000000000 {
		t.Errorf("expected 3000000000, got %d", got)
	}
}

func TestUint32_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("3000000000"))
	t.Cleanup(internal.ResetInput)

	got, err := Uint32()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 3000000000 {
		t.Errorf("expected 3000000000, got %d", got)
	}
}

func TestUint32_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Uint32()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint32_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "123456789")

	got, err := Uint32(Prompt("Enter a number: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789, got %d", got)
	}
}

func TestUint32_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "4294967296")

	_, err := Uint32()
	if err == nil {
		t.Fatal("expected an error for out-of-range input")
	}
	if !strings.Contains(err.Error(), "4294967296") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}
