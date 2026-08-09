package sleepycat

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func TestInt64_ValidInput(t *testing.T) {
	withInput(t, "922337203685477580")

	got, err := Int64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 922337203685477580 {
		t.Errorf("expected 922337203685477580, got %d", got)
	}
}

func TestInt64_NegativeNumber(t *testing.T) {
	withInput(t, "-922337203685477580")

	got, err := Int64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != -922337203685477580 {
		t.Errorf("expected -922337203685477580, got %d", got)
	}
}

func TestInt64_MinValue(t *testing.T) {
	withInput(t, "-9223372036854775808")

	got, err := Int64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != -9223372036854775808 {
		t.Errorf("expected -9223372036854775808, got %d", got)
	}
}

func TestInt64_MaxValue(t *testing.T) {
	withInput(t, "9223372036854775807")

	got, err := Int64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 9223372036854775807 {
		t.Errorf("expected 9223372036854775807, got %d", got)
	}
}

func TestInt64_TrimsWhitespace(t *testing.T) {
	withInput(t, "   123456789   ")

	got, err := Int64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789, got %d", got)
	}
}

func TestInt64_InvalidInput(t *testing.T) {
	withInput(t, "not-a-number")

	got, err := Int64()
	if err == nil {
		t.Fatalf("expected an error, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt64_EmptyInput(t *testing.T) {
	withInput(t, "")

	got, err := Int64()
	if err == nil {
		t.Fatalf("expected an error for empty input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt64_ValueAboveRange(t *testing.T) {
	withInput(t, "9223372036854775808")

	got, err := Int64()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt64_ValueBelowRange(t *testing.T) {
	withInput(t, "-9223372036854775809")

	got, err := Int64()
	if err == nil {
		t.Fatalf("expected an error for out-of-range input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt64_RetriesUntilValid(t *testing.T) {
	withInput(t, "abc", "9223372036854775808", "123456789")

	got, err := Int64(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789, got %d", got)
	}
}

func TestInt64_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "abc", "9223372036854775808", "-9223372036854775809")

	got, err := Int64(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt64_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "123456789")

	got, err := Int64(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789 from the 2nd attempt, got %d", got)
	}
}

func TestInt64_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "9223372036854775808", "-9223372036854775809", "100000")

	got, err := Int64(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != 100000 {
		t.Errorf("expected 100000, got %d", got)
	}
}

func TestInt64_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("123456789"))
	t.Cleanup(internal.ResetInput)

	got, err := Int64()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 123456789 {
		t.Errorf("expected 123456789, got %d", got)
	}
}

func TestInt64_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Int64()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt64_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "123")

	got, err := Int64(Prompt("Enter a number: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 123 {
		t.Errorf("expected 123, got %d", got)
	}
}

func TestInt64_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "9223372036854775808")

	_, err := Int64()
	if err == nil {
		t.Fatal("expected an error for out-of-range input")
	}
	if !strings.Contains(err.Error(), "9223372036854775808") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}
