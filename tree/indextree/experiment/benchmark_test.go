package experiment

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/474420502/structure/compare"
)

type shiftConfig struct {
	name  string
	shift int64
}

func shiftConfigs() []shiftConfig {
	return []shiftConfig{
		{"indextree-default", 2},
		{"shift=3", 3},
		{"shift=4", 4},
		{"shift=5", 5},
	}
}

func BenchmarkGet(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	configs := shiftConfigs()
	seed := int64(12345)

	for _, size := range sizes {
		for _, cfg := range configs {
			r := rand.New(rand.NewSource(seed))
			keys := make([]int64, size)
			for i := range keys {
				keys[i] = r.Int63()
			}

			tree := NewShiftTolerance[int64, int](compare.Any[int64], cfg.shift)
			for _, key := range keys {
				tree.Put(key, int(key))
			}

			b.Run(fmt.Sprintf("size=%d/%s", size, cfg.name), func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					for _, key := range keys {
						tree.Get(key)
					}
				}
			})
		}
	}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size=%d/AVL", size), func(b *testing.B) {
			r := rand.New(rand.NewSource(seed))
			keys := make([]int64, size)
			for i := range keys {
				keys[i] = r.Int63()
			}
			tree := New[int64, int](compare.Any[int64])
			for _, key := range keys {
				tree.Put(key, int(key))
			}

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, key := range keys {
					tree.Get(key)
				}
			}
		})
	}
}

func BenchmarkMixedWorkload(b *testing.B) {
	sizes := []int{1000, 5000, 10000}
	readRatios := []float64{0.5, 0.8, 0.95}
	configs := shiftConfigs()
	seed := int64(12345)

	for _, size := range sizes {
		for _, ratio := range readRatios {
			writeCount := int(float64(size) * (1 - ratio))

			for _, cfg := range configs {
				r := rand.New(rand.NewSource(seed))
				initialKeys := make([]int64, size-writeCount)
				for i := range initialKeys {
					initialKeys[i] = r.Int63()
				}

				workload := make([]int64, size)
				for i := range workload {
					workload[i] = r.Int63()
				}

				b.Run(fmt.Sprintf("size=%d/ratio=%.2f/%s", size, ratio, cfg.name), func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						tree := NewShiftTolerance[int64, int](compare.Any[int64], cfg.shift)
						for _, key := range initialKeys {
							tree.Put(key, int(key))
						}
						for j := 0; j < size; j++ {
							if j < writeCount {
								tree.Put(workload[j], int(workload[j]))
							} else {
								tree.Get(workload[j])
							}
						}
					}
				})
			}
		}
	}
}

func BenchmarkExtremeCases(b *testing.B) {
	configs := shiftConfigs()

	b.Run("AlreadySortedSequential", func(b *testing.B) {
		sizes := []int{100, 500, 1000, 5000}
		for _, size := range sizes {
			for _, cfg := range configs {
				b.Run(fmt.Sprintf("size=%d/%s", size, cfg.name), func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						tree := NewShiftTolerance[int64, int](compare.Any[int64], cfg.shift)
						for j := 0; j < size; j++ {
							tree.Put(int64(j), j)
						}
					}
				})
			}
		}
	})

	b.Run("ReverseSortedSequential", func(b *testing.B) {
		sizes := []int{100, 500, 1000, 5000}
		for _, size := range sizes {
			for _, cfg := range configs {
				b.Run(fmt.Sprintf("size=%d/%s", size, cfg.name), func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						tree := NewShiftTolerance[int64, int](compare.Any[int64], cfg.shift)
						for j := size - 1; j >= 0; j-- {
							tree.Put(int64(j), j)
						}
					}
				})
			}
		}
	})
}

func BenchmarkSmallBatchIncremental(b *testing.B) {
	configs := shiftConfigs()
	batchSizes := []int{10, 50, 100, 500}
	totalSize := 10000

	for _, batchSize := range batchSizes {
		batches := totalSize / batchSize
		for _, cfg := range configs {
			b.Run(fmt.Sprintf("batch=%d/%s", batchSize, cfg.name), func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					tree := NewShiftTolerance[int64, int](compare.Any[int64], cfg.shift)
					for batch := 0; batch < batches; batch++ {
						offset := batch * batchSize
						for j := 0; j < batchSize; j++ {
							tree.Put(int64(offset+j), offset+j)
						}
					}
				}
			})
		}
	}
}

func TestComprehensiveReport(t *testing.T) {
	seed := int64(12345)

	fmt.Printf("\n")
	fmt.Printf("================================================================================\n")
	fmt.Printf("                SINGLE-STAGE TOLERANCE TREE COMPARISON REPORT\n")
	fmt.Printf("================================================================================\n")
	fmt.Printf("Generated: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("Seed: %d\n", seed)

	fmt.Printf("\n--------------------------------------------------------------------------------\n")
	fmt.Printf("1. SEQUENTIAL INSERT (N=10000)\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	ReportSequentialInsert(10000)

	fmt.Printf("\n--------------------------------------------------------------------------------\n")
	fmt.Printf("2. RANDOM INSERT (N=10000)\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	ReportRandomInsert(10000, seed)

	fmt.Printf("\n--------------------------------------------------------------------------------\n")
	fmt.Printf("3. GET PERFORMANCE (N=10000, after build)\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	ReportGetPerformance(10000, seed)

	fmt.Printf("\n--------------------------------------------------------------------------------\n")
	fmt.Printf("4. MIXED WORKLOAD (N=10000, 80%% reads / 20%% writes)\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	ReportMixedWorkload(10000, 0.8, seed)

	fmt.Printf("\n--------------------------------------------------------------------------------\n")
	fmt.Printf("5. SHIFT GRID (N=10000)\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	ReportShiftGrid(10000, seed)
}

func ReportSequentialInsert(size int) {
	configs := shiftConfigs()

	fmt.Printf("%-22s %6s %8s %8s %8s %10s %6s %6s %8s\n",
		"Config", "Height", "Single", "Double", "Total", "AvgDepth", "P50", "P95", "Time(μs)")
	fmt.Println("---------------------------------------------------------------------------------------")

	avlTree := New[int64, int](compare.Any[int64])
	start := time.Now()
	for i := 0; i < size; i++ {
		avlTree.Put(int64(i), i)
	}
	avlElapsed := time.Since(start).Microseconds()
	avlStats := avlTree.BenchmarkStats()
	fmt.Printf("%-22s %6d %8d %8d %8d %10.2f %6d %6d %8d\n",
		"AVL", avlStats.Height, avlStats.SingleRotations, avlStats.DoubleRotations,
		avlStats.SingleRotations+avlStats.DoubleRotations,
		avlStats.AvgDepth, avlStats.P50Depth, avlStats.P95Depth, avlElapsed)

	for _, cfg := range configs {
		tree := NewShiftTolerance[int64, int](compare.Any[int64], cfg.shift)
		start = time.Now()
		for i := 0; i < size; i++ {
			tree.Put(int64(i), i)
		}
		elapsed := time.Since(start).Microseconds()
		stats := tree.BenchmarkStats()
		fmt.Printf("%-22s %6d %8d %8d %8d %10.2f %6d %6d %8d\n",
			cfg.name, stats.Height, stats.SingleRotations, stats.DoubleRotations,
			stats.SingleRotations+stats.DoubleRotations,
			stats.AvgDepth, stats.P50Depth, stats.P95Depth, elapsed)
	}
}

func ReportRandomInsert(size int, seed int64) {
	configs := shiftConfigs()
	r := rand.New(rand.NewSource(seed))
	keys := make([]int64, size)
	for i := range keys {
		keys[i] = r.Int63()
	}

	fmt.Printf("%-22s %6s %8s %8s %8s %10s %6s %6s %8s\n",
		"Config", "Height", "Single", "Double", "Total", "AvgDepth", "P50", "P95", "Time(μs)")
	fmt.Println("---------------------------------------------------------------------------------------")

	avlTree := New[int64, int](compare.Any[int64])
	start := time.Now()
	for _, key := range keys {
		avlTree.Put(key, int(key))
	}
	avlElapsed := time.Since(start).Microseconds()
	avlStats := avlTree.BenchmarkStats()
	fmt.Printf("%-22s %6d %8d %8d %8d %10.2f %6d %6d %8d\n",
		"AVL", avlStats.Height, avlStats.SingleRotations, avlStats.DoubleRotations,
		avlStats.SingleRotations+avlStats.DoubleRotations,
		avlStats.AvgDepth, avlStats.P50Depth, avlStats.P95Depth, avlElapsed)

	for _, cfg := range configs {
		tree := NewShiftTolerance[int64, int](compare.Any[int64], cfg.shift)
		start = time.Now()
		for _, key := range keys {
			tree.Put(key, int(key))
		}
		elapsed := time.Since(start).Microseconds()
		stats := tree.BenchmarkStats()
		fmt.Printf("%-22s %6d %8d %8d %8d %10.2f %6d %6d %8d\n",
			cfg.name, stats.Height, stats.SingleRotations, stats.DoubleRotations,
			stats.SingleRotations+stats.DoubleRotations,
			stats.AvgDepth, stats.P50Depth, stats.P95Depth, elapsed)
	}
}

func ReportGetPerformance(size int, seed int64) {
	configs := shiftConfigs()
	r := rand.New(rand.NewSource(seed))
	keys := make([]int64, size)
	for i := range keys {
		keys[i] = r.Int63()
	}

	fmt.Printf("%-22s %12s %12s %12s\n", "Config", "Get(μs)", "ns/op", "Ops/sec")
	fmt.Println("---------------------------------------------------------------")

	for _, cfg := range configs {
		tree := NewShiftTolerance[int64, int](compare.Any[int64], cfg.shift)
		for _, key := range keys {
			tree.Put(key, int(key))
		}

		start := time.Now()
		for _, key := range keys {
			tree.Get(key)
		}
		elapsed := time.Since(start)
		fmt.Printf("%-22s %12d %12d %12.0f\n",
			cfg.name,
			elapsed.Microseconds(),
			elapsed.Nanoseconds()/int64(size),
			float64(size)*1e9/float64(elapsed.Nanoseconds()))
	}

	avlTree := New[int64, int](compare.Any[int64])
	for _, key := range keys {
		avlTree.Put(key, int(key))
	}
	start := time.Now()
	for _, key := range keys {
		avlTree.Get(key)
	}
	elapsed := time.Since(start)
	fmt.Printf("%-22s %12d %12d %12.0f\n",
		"AVL",
		elapsed.Microseconds(),
		elapsed.Nanoseconds()/int64(size),
		float64(size)*1e9/float64(elapsed.Nanoseconds()))
}

func ReportMixedWorkload(size int, readRatio float64, seed int64) {
	configs := shiftConfigs()
	writeCount := int(float64(size) * (1 - readRatio))

	r := rand.New(rand.NewSource(seed))
	initialSize := size - writeCount
	initialKeys := make([]int64, initialSize)
	for i := range initialKeys {
		initialKeys[i] = r.Int63()
	}
	workload := make([]int64, size)
	for i := range workload {
		workload[i] = r.Int63()
	}

	fmt.Printf("%-22s %8s %12s %12s %10s\n", "Config", "Height", "TotalRot", "Time(μs)", "Throughput")
	fmt.Println("-----------------------------------------------------------------")

	{
		var avlTree *Tree[int64, int]
		var stats BenchmarkStats
		start := time.Now()
		for run := 0; run < 10; run++ {
			avlTree = New[int64, int](compare.Any[int64])
			for _, key := range initialKeys {
				avlTree.Put(key, int(key))
			}
			for j := 0; j < size; j++ {
				if j < writeCount {
					avlTree.Put(workload[j], int(workload[j]))
				} else {
					avlTree.Get(workload[j])
				}
			}
		}
		elapsed := time.Since(start)
		stats = avlTree.BenchmarkStats()
		fmt.Printf("%-22s %8d %12d %12d %10.0f\n", "AVL", stats.Height, stats.SingleRotations+stats.DoubleRotations, elapsed.Microseconds()/10, float64(size*10)/elapsed.Seconds())
	}

	for _, cfg := range configs {
		var stats BenchmarkStats
		start := time.Now()
		for run := 0; run < 10; run++ {
			tree := NewShiftTolerance[int64, int](compare.Any[int64], cfg.shift)
			for _, key := range initialKeys {
				tree.Put(key, int(key))
			}
			for j := 0; j < size; j++ {
				if j < writeCount {
					tree.Put(workload[j], int(workload[j]))
				} else {
					tree.Get(workload[j])
				}
			}
			stats = tree.BenchmarkStats()
		}
		elapsed := time.Since(start)
		fmt.Printf("%-22s %8d %12d %12d %10.0f\n", cfg.name, stats.Height, stats.SingleRotations+stats.DoubleRotations, elapsed.Microseconds()/10, float64(size*10)/elapsed.Seconds())
	}
}

func ReportShiftGrid(size int, seed int64) {
	r := rand.New(rand.NewSource(seed))
	keys := make([]int64, size)
	for i := range keys {
		keys[i] = r.Int63()
	}

	fmt.Printf("%-12s %6s %8s %8s %8s\n", "Shift", "Height", "Single", "Double", "Total")
	fmt.Println("------------------------------------------------")

	for shift := int64(1); shift <= 6; shift++ {
		tree := NewShiftTolerance[int64, int](compare.Any[int64], shift)
		for _, key := range keys {
			tree.Put(key, int(key))
		}
		stats := tree.BenchmarkStats()
		fmt.Printf("%-12d %6d %8d %8d %8d\n", shift, stats.Height, stats.SingleRotations, stats.DoubleRotations, stats.SingleRotations+stats.DoubleRotations)
	}
}
