package vtprotobuf_bench

import (
	"fmt"
	"math/rand"
	"testing"

	"google.golang.org/protobuf/proto"
	"vtprotobuf-bench/protos"
)

func BenchmarkListInt32(b *testing.B) {
	for _, n := range []int{10, 100, 1000, 10000} {
		decoded := make([]*protos.ListInt32, 1000)
		for i := range decoded {
			items := make([]int32, n)
			for j := range items {
				items[j] = rand.Int31()
			}
			decoded[i] = &protos.ListInt32{Items: items}
		}

		encoded := make([][]byte, len(decoded))
		total := 0
		for i := range encoded {
			bs, err := proto.Marshal(decoded[i])
			if err != nil {
				b.Fatal(err)
			}
			encoded[i] = bs
			total += len(bs)
		}

		b.Run(fmt.Sprint(n), func(b *testing.B) {
			b.Run("Marshal", func(b *testing.B) {
				total := 0
				for i := 0; i < b.N; i++ {
					bs, err := proto.Marshal(decoded[i%len(decoded)])
					total += len(bs)
					if err != nil {
						b.Fatal(err)
					}
				}
				b.ReportMetric(float64(total)/float64(b.N), "bytes")
			})

			b.Run("MarshalVT", func(b *testing.B) {
				total := 0
				for i := 0; i < b.N; i++ {
					bs, err := decoded[i%len(decoded)].MarshalVT()
					total += len(bs)
					if err != nil {
						b.Fatal(err)
					}
				}
				b.ReportMetric(float64(total)/float64(b.N), "bytes")
			})

			b.Run("Unmarshal", func(b *testing.B) {
				total := 0
				for i := 0; i < b.N; i++ {
					var l protos.ListInt32
					bs := encoded[i%len(encoded)]
					total += len(bs)
					if err := proto.Unmarshal(bs, &l); err != nil {
						b.Fatal(err)
					}
				}
				b.ReportMetric(float64(total)/float64(b.N), "bytes")
			})

			b.Run("UnmarshalVT", func(b *testing.B) {
				total := 0
				for i := 0; i < b.N; i++ {
					var l protos.ListInt32
					bs := encoded[i%len(encoded)]
					total += len(bs)
					if err := l.UnmarshalVT(bs); err != nil {
						b.Fatal(err)
					}
				}
				b.ReportMetric(float64(total)/float64(b.N), "bytes")
			})
		})
	}
}

func BenchmarkListSfixed32(b *testing.B) {
	for _, n := range []int{10, 100, 1000, 10000} {
		decoded := make([]*protos.ListSfixed32, 1000)
		for i := range decoded {
			items := make([]int32, n)
			for j := range items {
				items[j] = rand.Int31()
			}
			decoded[i] = &protos.ListSfixed32{Items: items}
		}

		encoded := make([][]byte, len(decoded))
		total := 0
		for i := range encoded {
			bs, err := proto.Marshal(decoded[i])
			if err != nil {
				b.Fatal(err)
			}
			encoded[i] = bs
			total += len(bs)
		}

		b.Run(fmt.Sprint(n), func(b *testing.B) {
			b.Run("Marshal", func(b *testing.B) {
				total := 0
				for i := 0; i < b.N; i++ {
					bs, err := proto.Marshal(decoded[i%len(decoded)])
					total += len(bs)
					if err != nil {
						b.Fatal(err)
					}
				}
				b.ReportMetric(float64(total)/float64(b.N), "bytes")
			})

			b.Run("MarshalVT", func(b *testing.B) {
				total := 0
				for i := 0; i < b.N; i++ {
					bs, err := decoded[i%len(decoded)].MarshalVT()
					total += len(bs)
					if err != nil {
						b.Fatal(err)
					}
				}
				b.ReportMetric(float64(total)/float64(b.N), "bytes")
			})

			b.Run("Unmarshal", func(b *testing.B) {
				total := 0
				for i := 0; i < b.N; i++ {
					var l protos.ListSfixed32
					bs := encoded[i%len(encoded)]
					total += len(bs)
					if err := proto.Unmarshal(bs, &l); err != nil {
						b.Fatal(err)
					}
				}
				b.ReportMetric(float64(total)/float64(b.N), "bytes")
			})

			b.Run("UnmarshalVT", func(b *testing.B) {
				total := 0
				for i := 0; i < b.N; i++ {
					var l protos.ListSfixed32
					bs := encoded[i%len(encoded)]
					total += len(bs)
					if err := l.UnmarshalVT(bs); err != nil {
						b.Fatal(err)
					}
				}
				b.ReportMetric(float64(total)/float64(b.N), "bytes")
			})
		})
	}
}
