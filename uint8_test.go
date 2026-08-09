package sleepycat

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func TestUint8_ValidInput(t *testing.T) {
	withInput(t, "42")

	got, err := Uint8()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 42 {
		t.Errorf("expected 42, got %d", got)
	}
}

func TestUint8_Zero(t *testing.T) {
	withInput(t, "0")

	got, err := Uint8()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestUint8_MaxValue(t *testing.T) {
	withInput(t, "255")

	got, err := Uint8()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 255 {
		t.Errorf("expected 255, got %d", got)
	}
}

func TestUint8_TrimsWhitespace(t *testing.T) {
	withInput(t, "   123   ")

	got, err := Uint8()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123 {
		t.Errorf("expected 123, got %d", got)
	}
}

func TestUint8_NegativeNumber(t *testing.T) {
	withInput(t, "-1")

	got, err := Uint8()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint8_ValueAboveRange(t *testing.T) {
	withInput(t, "256")

	got, err := Uint8()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint8_InvalidInput(t *testing.T) {
	withInput(t, "not-a-number")

	got, err := Uint8()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint8_EmptyInput(t *testing.T) {
	withInput(t, "")

	got, err := Uint8()
	if err == nil {
		t.Fatalf("expected an error for empty input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint8_RetriesUntilValid(t *testing.T) {
	withInput(t, "abc", "-1", "200")

	got, err := Uint8(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 200 {
		t.Errorf("expected 200, got %d", got)
	}
}

func TestUint8_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "abc", "256", "-1")

	got, err := Uint8(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint8_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "200")

	got, err := Uint8(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != 200 {
		t.Errorf("expected 200 from the 2nd attempt, got %d", got)
	}
}

func TestUint8_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "256", "-1", "100")

	got, err := Uint8(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != 100 {
		t.Errorf("expected 100, got %d", got)
	}
}

func TestUint8_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("200"))
	t.Cleanup(internal.ResetInput)

	got, err := Uint8()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 200 {
		t.Errorf("expected 200, got %d", got)
	}
}

func TestUint8_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Uint8()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestUint8_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "123")

	got, err := Uint8(Prompt("Enter a number: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123 {
		t.Errorf("expected 123, got %d", got)
	}
}

func TestUint8_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "256")

	_, err := Uint8()
	if err == nil {
		t.Fatal("expected an error for out-of-range input")
	}
	if !strings.Contains(err.Error(), "256") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}
