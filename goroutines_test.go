package goroutines

import (
	"fmt"
	"slices"
	"testing"
	"time"
)

type resultMeasured struct {
	Value       string
	ElapsedTime time.Duration
}

func TestLaunch(t *testing.T) {
	var results = make([]resultMeasured, 0, 3)

	start := time.Now()
	Launch(
		func() {
			time.Sleep(100 * time.Millisecond)
			results = append(results, resultMeasured{"a", time.Since(start)})
		},
		func() {
			time.Sleep(100 * time.Millisecond)
			results = append(results, resultMeasured{"b", time.Since(start)})
		},
		func() {
			time.Sleep(100 * time.Millisecond)
			results = append(results, resultMeasured{"c", time.Since(start)})
		},
	)
	totalTime := time.Since(start)

	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results))
	}

	for i, r := range results {
		if !slices.Contains(results, r) {
			t.Errorf("Expected results contain %s", r)
		}

		fmt.Println("Goroutine time:", r.ElapsedTime, "Total time:", totalTime)
		if totalTime < r.ElapsedTime {
			t.Errorf("Expected elapsed time %s to be less than total time %s", r.ElapsedTime, totalTime)
		}
		if totalTime-r.ElapsedTime > (100 * time.Microsecond) {
			t.Errorf("Expected elapsed time scewing to be less than 100 microseconds for result %d", i)
		}
	}
}

var expectedAsyncResults = map[int]string{
	0: "a",
	1: "b",
	2: "c",
	3: "d",
	4: "e",
}

func TestAsync(t *testing.T) {
	start := time.Now()
	results := Async(
		func() resultMeasured {
			time.Sleep(100 * time.Millisecond)
			return resultMeasured{"a", time.Since(start)}
		},
		func() resultMeasured {
			time.Sleep(100 * time.Millisecond)
			return resultMeasured{"b", time.Since(start)}
		},
		func() resultMeasured {
			time.Sleep(100 * time.Millisecond)
			return resultMeasured{"c", time.Since(start)}
		},
		func() resultMeasured {
			time.Sleep(100 * time.Millisecond)
			return resultMeasured{"d", time.Since(start)}
		},
		func() resultMeasured {
			time.Sleep(100 * time.Millisecond)
			return resultMeasured{"e", time.Since(start)}
		},
	)
	totalTime := time.Since(start)

	if len(results) != 5 {
		t.Errorf("Expected 5 results, got %d", len(results))
	}

	for pos, expected := range expectedAsyncResults {
		got := results[pos]
		if got.Value != expected {
			t.Errorf("Expected result %s at position %d, got %s", expected, pos, got)
		}

		fmt.Println("Goroutine time:", got.ElapsedTime, "Total time:", totalTime)
		if totalTime < got.ElapsedTime {
			t.Errorf("Expected elapsed time %s to be less than total time %s", got.ElapsedTime, totalTime)
		}
		if totalTime-got.ElapsedTime > (100 * time.Microsecond) {
			t.Errorf("Expected elapsed time scewing to be less than 100 microseconds for result %d", pos)
		}
	}
}

func TestAsync1(t *testing.T) {
	res := Async1(func() string { return "a" })

	if res != "a" {
		t.Errorf("Expected result %s, got %s", "a", res)
	}
}

func TestAsync2(t *testing.T) {
	res1, res2 := Async2(
		func() bool { return true },
		func() bool { return false },
	)

	if res1 != true {
		t.Errorf("Expected result %t, got %t", true, res1)
	}
	if res2 != false {
		t.Errorf("Expected result %t, got %t", false, res2)
	}
}

func TestAsync3(t *testing.T) {
	res1, res2, res3 := Async3(
		func() int { return 1 },
		func() int { return 2 },
		func() int { return 3 },
	)

	if res1 != 1 {
		t.Errorf("Expected result %d, got %d", 1, res1)
	}
	if res2 != 2 {
		t.Errorf("Expected result %d, got %d", 2, res2)
	}
	if res3 != 3 {
		t.Errorf("Expected result %d, got %d", 3, res3)
	}
}

func TestAsync4(t *testing.T) {
	res1, res2, res3, res4 := Async4(
		func() string { return "a" },
		func() string { return "b" },
		func() string { return "c" },
		func() string { return "d" },
	)

	if res1 != "a" {
		t.Errorf("Expected result %s, got %s", "a", res1)
	}
	if res2 != "b" {
		t.Errorf("Expected result %s, got %s", "b", res2)
	}
	if res3 != "c" {
		t.Errorf("Expected result %s, got %s", "c", res3)
	}
	if res4 != "d" {
		t.Errorf("Expected result %s, got %s", "d", res4)
	}
}

func TestAsync5(t *testing.T) {
	res1, res2, res3, res4, res5 := Async5(
		func() int64 { return 1 },
		func() int64 { return 2 },
		func() int64 { return 3 },
		func() int64 { return 4 },
		func() int64 { return 5 },
	)

	if res1 != 1 {
		t.Errorf("Expected result %d, got %d", 1, res1)
	}
	if res2 != 2 {
		t.Errorf("Expected result %d, got %d", 2, res2)
	}
	if res3 != 3 {
		t.Errorf("Expected result %d, got %d", 3, res3)
	}
	if res4 != 4 {
		t.Errorf("Expected result %d, got %d", 4, res4)
	}
	if res5 != 5 {
		t.Errorf("Expected result %d, got %d", 5, res5)
	}
}
