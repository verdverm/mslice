package main

import (
	"fmt"
)


type Mslice interface {

	Get( pos []int ) float64
	Set( pos []int, val float64 )

	Sub( ranges [][2]int ) Mslice
	// SubCopy( ranges [][2]int ) Mslice


	String() string
}

func NewMslice( szDim []int ) Mslice {
	var ret Mslice = nil
	switch len(szDim) {
		case 1:
			ret = makeMslice1(szDim)
		case 2:
			ret = makeMslice2(szDim)
		case 3:
			ret = makeMslice3(szDim)
		case 4:
			ret = makeMslice4(szDim)
		default:
			panic( fmt.Sprintf("%d dimensional Mslice is not implemented yet", len(szDim)) )
	}
	return ret
}



type dimInfo struct {
	dimMul int
	dimCap int
	low_dim int
	high_dim int
}


type Mslice1 struct {
	data []float64
	info [1]dimInfo
}

func makeMslice1( szDim []int ) Mslice {
	var ret Mslice1
  ret.data = make( []float64, szDim[0] )
  ret.info[0].dimMul = 1
  ret.info[0].dimCap = 1
  ret.info[0].low_dim = 0
  ret.info[0].high_dim = szDim[0]

  for i:=0; i<len(ret.data); i++ { ret.data[i] = float64(i) }
  return &ret
}
func (m *Mslice1) String() string {
	return fmt.Sprintf("MS1(%d): %v", len(m.data), m.info)
}


type Mslice2 struct {
	data []float64
	info [2]dimInfo
}

func makeMslice2( szDim []int ) Mslice {
	var ret Mslice2

	ret.data = make( []float64, szDim[0]*szDim[1] )
	ret.info[0].dimMul = szDim[1]
	ret.info[0].dimCap = szDim[0]
	ret.info[0].low_dim = 0
	ret.info[0].high_dim = szDim[0]
	ret.info[1].dimMul = 1
	ret.info[1].dimCap = 1
	ret.info[1].low_dim = 0
	ret.info[1].high_dim = szDim[1]

	for i:=0; i<len(ret.data); i++ { ret.data[i] = float64(i) }
	return &ret
}
func (m *Mslice2) String() string {
	return fmt.Sprintf("MS2(%d): %v", len(m.data), m.info)
}


type Mslice3 struct {
	data []float64
	info [3]dimInfo
}

func makeMslice3( szDim []int ) Mslice {
	var ret Mslice3

	ret.data = make( []float64, szDim[0]*szDim[1]*szDim[2] )
	ret.info[0].dimMul = szDim[1]*szDim[2]
	ret.info[0].dimCap = szDim[0]
	ret.info[0].low_dim = 0
	ret.info[0].high_dim = szDim[0]
	ret.info[1].dimMul = szDim[1]
	ret.info[1].dimCap = szDim[1]
	ret.info[1].low_dim = 0
	ret.info[1].high_dim = szDim[1]
	ret.info[2].dimMul = 1
	ret.info[2].dimCap = 1
	ret.info[2].low_dim = 0
	ret.info[2].high_dim = szDim[2]

  for i:=0; i<len(ret.data); i++ { ret.data[i] = float64(i) }
	return &ret
}
func (m *Mslice3) String() string {
	return fmt.Sprintf("MS3(%d): %v", len(m.data), m.info)
}


type Mslice4 struct {
	data []float64
	info [4]dimInfo
}

func makeMslice4( szDim []int ) Mslice {
	var ret Mslice4

	ret.data = make( []float64, szDim[0]*szDim[1]*szDim[2]*szDim[3] )
	ret.info[0].dimMul = szDim[1]*szDim[2]*szDim[3]
	ret.info[0].dimCap = szDim[0]
	ret.info[0].low_dim = 0
	ret.info[0].high_dim = szDim[0]
	ret.info[1].dimMul = szDim[1]*szDim[2]
	ret.info[1].dimCap = szDim[1]
	ret.info[1].low_dim = 0
	ret.info[1].high_dim = szDim[1]
	ret.info[2].dimMul = szDim[1]
	ret.info[2].dimCap = szDim[2]
	ret.info[2].low_dim = 0
	ret.info[2].high_dim = szDim[2]
	ret.info[3].dimMul = 1
	ret.info[3].dimCap = 1
	ret.info[3].low_dim = 0
	ret.info[3].high_dim = szDim[3]

  for i:=0; i<len(ret.data); i++ { ret.data[i] = float64(i) }
	return &ret
}
func (m *Mslice4) String() string {
	return fmt.Sprintf("MS4(%d): %v", len(m.data), m.info)
}



func (m *Mslice1) Get( pos []int ) float64 {
	p := (m.info[0].low_dim+pos[0])
	return m.data[p]
}

func (m *Mslice1) Set( pos []int, val float64 ) {
	p := (m.info[0].low_dim+pos[0])
	m.data[p] = val
}

func (m *Mslice1) Sub( pos [][2]int ) Mslice {
	var ret Mslice1

	ret.data = m.data
	ret.info = m.info
	ret.info[0].low_dim = pos[0][0]
	ret.info[0].high_dim = pos[0][1]
	return &ret
}



func (m *Mslice2) Get( pos []int ) float64 {
	p := (m.info[0].low_dim+pos[0])*m.info[0].dimMul
	p += (m.info[1].low_dim+pos[1])
	return m.data[p]
}

func (m *Mslice2) Set( pos []int, val float64 ) {
	p := (m.info[0].low_dim+pos[0])*m.info[0].dimMul
	p += (m.info[1].low_dim+pos[1])
	m.data[p] = val
}

func (m *Mslice2) Sub( pos [][2]int ) Mslice {
	var ret Mslice2

	ret.data = m.data
	ret.info = m.info
	ret.info[0].low_dim = pos[0][0]
	ret.info[0].high_dim = pos[0][1]
	ret.info[1].low_dim = pos[1][0]
	ret.info[1].high_dim = pos[1][1]
	return &ret
}



func (m *Mslice3) Get( pos []int ) float64 {
	p := (m.info[0].low_dim+pos[0])*m.info[0].dimMul
	p += (m.info[1].low_dim+pos[1])*m.info[1].dimMul
	p += (m.info[2].low_dim+pos[2])
	return m.data[p]
}

func (m *Mslice3) Set( pos []int, val float64 ) {
	p := (m.info[0].low_dim+pos[0])*m.info[0].dimMul
	p += (m.info[1].low_dim+pos[1])*m.info[1].dimMul
	p += (m.info[2].low_dim+pos[2])
	m.data[p] = val
}

func (m *Mslice3) Sub( pos [][2]int ) Mslice {
	var ret Mslice3

	ret.data = m.data
	ret.info = m.info
	ret.info[0].low_dim = pos[0][0]
	ret.info[0].high_dim = pos[0][1]
	ret.info[1].low_dim = pos[1][0]
	ret.info[1].high_dim = pos[1][1]
	ret.info[2].low_dim = pos[2][0]
	ret.info[2].high_dim = pos[2][1]
	return &ret
}



func (m *Mslice4) Get( pos []int ) float64 {
	p := (m.info[0].low_dim+pos[0])*m.info[0].dimMul
	p += (m.info[1].low_dim+pos[1])*m.info[1].dimMul
	p += (m.info[2].low_dim+pos[2])*m.info[2].dimMul
	p += (m.info[3].low_dim+pos[3])
	return m.data[p]
}

func (m *Mslice4) Set( pos []int, val float64 ) {
	p := (m.info[0].low_dim+pos[0])*m.info[0].dimMul
	p += (m.info[1].low_dim+pos[1])*m.info[1].dimMul
	p += (m.info[2].low_dim+pos[2])*m.info[2].dimMul
	p += (m.info[3].low_dim+pos[3])
	m.data[p] = val
}

func (m *Mslice4) Sub( pos [][2]int ) Mslice {
	var ret Mslice4

	ret.data = m.data
	ret.info = m.info
	ret.info[0].low_dim = pos[0][0]
	ret.info[0].high_dim = pos[0][1]
	ret.info[1].low_dim = pos[1][0]
	ret.info[1].high_dim = pos[1][1]
	ret.info[2].low_dim = pos[2][0]
	ret.info[2].high_dim = pos[2][1]
	ret.info[3].low_dim = pos[3][0]
	ret.info[3].high_dim = pos[3][1]
	return &ret
}








func main() {
	fmt.Printf( "Hello World\n ~ Mslice\n\n" )

	m1 := NewMslice( []int{4} )
	m2 := NewMslice( []int{4,4} )
	m3 := NewMslice( []int{4,4,4} )
	m4 := NewMslice( []int{4,4,4,4} )

	fmt.Printf( "Mslices Dump\n%v\n%v\n%v\n%v\n\n", m1,m2,m3,m4 )
	fmt.Printf( "%v\n%v\n%v\n%v\n\n", m1.Get([]int{1}), m2.Get([]int{0,1}),m3.Get([]int{0,0,1}), m4.Get([]int{0,0,0,1}) )
	fmt.Printf( "%v\n%v\n%v\n%v\n\n", m1.Get([]int{1}), m2.Get([]int{1,0}), m3.Get([]int{1,0,0}), m4.Get([]int{1,0,0,0}) )
	fmt.Printf( "%v\n%v\n%v\n%v\n\n", m1.Get([]int{1}), m2.Get([]int{1,1}), m3.Get([]int{1,1,1}), m4.Get([]int{1,1,1,1}) )

	s1 := m1.Sub( [][2]int{ {1,3} } )
	s2 := m2.Sub( [][2]int{ {1,3},{1,3} } )
	s3 := m3.Sub( [][2]int{ {1,3},{1,3},{1,3} } )
	s4 := m4.Sub( [][2]int{ {1,3},{1,3},{1,3},{1,3} } )

	fmt.Printf( "%v\n%v\n%v\n%v\n\n", s1, s2, s3, s4 )
	fmt.Printf( "%v\n%v\n%v\n%v\n\n", s1.Get([]int{0}), s2.Get([]int{0,0}),s3.Get([]int{0,0,0}), s4.Get([]int{0,0,0,0}) )
	fmt.Printf( "%v\n%v\n%v\n%v\n\n", s1.Get([]int{1}), s2.Get([]int{0,1}),s3.Get([]int{0,0,1}), s4.Get([]int{0,0,0,1}) )
	fmt.Printf( "%v\n%v\n%v\n%v\n\n", s1.Get([]int{1}), s2.Get([]int{1,0}), s3.Get([]int{1,0,0}), s4.Get([]int{1,0,0,0}) )
	fmt.Printf( "%v\n%v\n%v\n%v\n\n", s1.Get([]int{1}), s2.Get([]int{1,1}), s3.Get([]int{1,1,1}), s4.Get([]int{1,1,1,1}) )

}
