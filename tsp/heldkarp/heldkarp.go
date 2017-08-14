package heldkarp

import (
	"fmt"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/bitstring"
)

/*

Held-Karp
=========

Example 1
---------

  |  A   B   C   D
--+---------------
A | --   1  15   6
B |  2  --   7   3
C |  9   6  --  12
D | 10   4   8  --

c(B, {}) = mAB =  1 [mAB]
c(C, {}) = mAC = 15 [mAC]
c(D, {}) = mAD =  6 [mAD]

c(C, {B}) = mBC + c(B, {}) =  7 +  1 =  8 [c(B, {})]
c(D, {B}) = mBD + c(B, {}) =  3 +  1 =  4 [c(B, {})]
c(B, {C}) = mCB + c(C, {}) =  6 + 15 = 21 [c(C, {})]
c(D, {C}) = mCD + c(C, {}) = 12 + 15 = 27 [c(C, {})]
c(B, {D}) = mDB + c(D, {}) =  4 +  6 = 10 [c(D, {})]
c(C, {D}) = mDC + c(D, {}) =  8 +  6 = 14 [c(D, {})]

c(B, {C,D}) = min(mCB + c(C, {D}), mDB + c(D, {C})) = min(6 + 14,  4 + 27) = min(20, 31) = 20 [c(C, {D})]
c(C, {B,D}) = min(mBC + c(B, {D}), mDC + c(D, {B})) = min(7 + 10,  8 +  4) = min(17, 12) = 12 [c(D, {B})]
c(D, {B,C}) = min(mBD + c(B, {C}), mCD + c(C, {B})) = min(3 + 21, 12 +  8) = min(24, 20) = 20 [c(C, {B})]

c(A,{B,C,D}) = min(mBA + c(B, {C,D}), mCA + c(C, {B,D}), mDA + c(D, {B,C}))
             = min(2 + 20, 9 + 12, 10 + 20)
             = min(22, 21, 30)
             = 21 [c(C, {B,D})]

Tour: A <- C <- D <- B <- A

Example 2
---------

  |  A   B   C   D   E
--+-------------------
A | --  20  30  10  11
B | 15  --  16   4   2
C |  3   5  --   2   4
D | 19   6  18  --   3
E | 16   4   7  16  --

c(B, {}) = mAB = 20 [mAB]
c(C, {}) = mAC = 30 [mAC]
c(D, {}) = mAD = 10 [mAD]
c(E, {}) = mAE = 11 [mAE]

c(C, {B}) = mBC + c(B, {}) = 16 + 20 = 36 [c(B, {})]
c(D, {B}) = mBD + c(B, {}) =  4 + 20 = 24 [c(B, {})]
c(E, {B}) = mBE + c(B, {}) =  2 + 20 = 22 [c(B, {})]
c(B, {C}) = mCB + c(C, {}) =  5 + 30 = 35 [c(C, {})]
c(D, {C}) = mCD + c(C, {}) =  2 + 30 = 32 [c(C, {})]
c(E, {C}) = mCE + c(C, {}) =  4 + 30 = 34 [c(C, {})]
c(B, {D}) = mDB + c(D, {}) =  6 + 10 = 16 [c(D, {})]
c(C, {D}) = mDC + c(D, {}) = 18 + 10 = 28 [c(D, {})]
c(E, {D}) = mDE + c(D, {}) =  3 + 10 = 13 [c(D, {})]
c(B, {E}) = mEB + c(E, {}) =  4 + 11 = 15 [c(E, {})]
c(C, {E}) = mEC + c(E, {}) =  7 + 11 = 18 [c(E, {})]
c(D, {E}) = mED + c(E, {}) = 16 + 11 = 27 [c(E, {})]

c(D, {B,C}) = min(mBD + c(B, {C}), mCD + c(C, {B})) = min( 4 + 35,  2 + 36) = min(39, 38) = 38 [c(C, {B})]
c(E, {B,C}) = min(mBE + c(B, {C}), mCE + c(C, {B})) = min( 2 + 35,  4 + 36) = min(37, 40) = 37 [c(B, {C})]
c(C, {B,D}) = min(mBC + c(B, {D}), mDC + c(D, {B})) = min(16 + 16, 18 + 24) = min(32, 42) = 32 [c(B, {D})]
c(E, {B,D}) = min(mBE + c(B, {D}), mDE + c(D, {B})) = min( 2 + 16,  3 + 24) = min(18, 27) = 18 [c(B, {D})]
c(C, {B,E}) = min(mBC + c(B, {E}), mEC + c(E, {B})) = min(16 + 15,  7 + 22) = min(31, 29) = 29 [c(E, {B})]
c(D, {B,E}) = min(mBD + c(B, {E}), mED + c(E, {B})) = min( 4 + 15, 16 + 22) = min(19, 38) = 19 [c(B, {E})]
c(B, {C,D}) = min(mCB + c(C, {D}), mDB + c(D, {C})) = min( 5 + 28,  6 + 32) = min(33, 38) = 33 [c(C, {D})]
c(E, {C,D}) = min(mCE + c(C, {D}), mDE + c(D, {C})) = min( 4 + 28,  3 + 32) = min(32, 35) = 32 [c(C, {D})]
c(B, {C,E}) = min(mCB + c(C, {E}), mEB + c(E, {C})) = min( 5 + 18,  4 + 34) = min(23, 38) = 23 [c(C, {E})]
c(D, {C,E}) = min(mCD + c(C, {E}), mED + c(E, {C})) = min( 2 + 18, 16 + 34) = min(20, 50) = 20 [c(C, {E})]
c(B, {D,E}) = min(mDB + c(D, {E}), mEB + c(E, {D})) = min( 6 + 27,  4 + 13) = min(33, 17) = 17 [c(E, {D})]
c(C, {D,E}) = min(mDC + c(D, {E}), mEC + c(E, {D})) = min(18 + 27,  7 + 13) = min(45, 20) = 20 [c(E, {D})]

c(E, {B,C,D}) = min(mDE + c(D, {B,C}), mCE + c(C, {B,D}), mBE + c(B, {C,D})) = min( 3 + 38,  4 + 32,  2 + 33) = min(41, 36, 35) = 35 [c(B, {C,D})]
c(D, {B,C,E}) = min(mED + c(E, {B,C}), mCD + c(C, {B,E}), mBD + c(B, {C,E})) = min(16 + 37,  2 + 29,  4 + 23) = min(53, 31, 27) = 27 [c(B, {C,E})]
c(C, {B,D,E}) = min(mEC + c(E, {B,D}), mDC + c(D, {B,E}), mBC + c(B, {D,E})) = min( 7 + 18, 18 + 19, 16 + 17) = min(25, 37, 35) = 25 [c(E, {B,D})]
c(B, {C,D,E}) = min(mEB + c(E, {C,D}), mDB + c(D, {C,E}), mCB + c(C, {D,E})) = min( 4 + 32,  6 + 20,  5 + 20) = min(36, 26, 25) = 25 [c(C, {D,E})]

c(A, {B,C,D,E}) = min(mEA + c(E, {B,C,D}), mDA + c(D, {B,C,E}), mCA + c(C, {B,D,E}), mBA + c(B, {C,D,E}))
                = min(16 + 35, 19 + 27, 3 + 25, 15 + 25)
                = min(51, 46, 28, 40)
                = 28 [c(C, {B,D,E})]

Tour: A <- C <- E <- B <- D <- A

       +----+                 +----+                 +----+                 +----+
       | 20 |                 | 30 |                 | 10 |                 | 11 |
       +----+                 +----+                 +----+                 +----+

+----+ +----+ +----+   +----+ +----+ +----+   +----+ +----+ +----+   +----+ +----+ +----+
| 36 | | 24 | | 22 |   | 35 | | 32 | | 34 |   | 16 | | 28 | | 13 |   | 15 | | 18 | | 27 |
+----+ +----+ +----+   +----+ +----+ +----+   +----+ +----+ +----+   +----+ +----+ +----+

    +----+----+   +----+----+   +----+----+   +----+----+   +----+----+   +----+----+
    | 38 | 37 |   | 32 | 18 |   | 29 | 19 |   | 33 | 32 |   | 23 | 20 |   | 17 | 20 |
    +----+----+   +----+----+   +----+----+   +----+----+   +----+----+   +----+----+

       +----+                 +----+                 +----+                 +----+
       | 35 |                 | 27 |                 | 25 |                 | 25 |
       +----+                 +----+                 +----+                 +----+

                                          +----+
                                          | 28 |
                                          +----+
*/

func Hello() {
	cityCount := uint(4)
	onesLimit := uint64(1) << (cityCount - 1)
	for selectedCount := uint(1); selectedCount < cityCount; selectedCount++ {
		for city := uint(0); city < cityCount; city++ {
			for ones := bitstring.Ones(selectedCount); ones < onesLimit; ones = bitstring.NextNumberWithSameBitCount(ones) {
				cities := bitstring.InsertZero(ones, city)
				fmt.Printf("selectedCount: %d, city: %d, cities: %08s\n", selectedCount, city, bitstring.ToString(cities))
				// Find Minimum:
				left := cities
				right := uint64(0)
				for left != 0 {
					left = bitstring.ResetRightmostOne(left)
					fromCities := left | right
					viaCity := cities ^ fromCities
					right = right | viaCity
					//fmt.Printf("    %08s, %08s, %d, %08s\n", ToString(left), ToString(fromCities), NumberOfTrailingZerosForPowerOfTwo(viaCity), ToString(right))
					fmt.Printf("    From Cities: 0x%016x, Via City: %d\n", fromCities, bitstring.NumberOfTrailingZerosForPowerOfTwo(viaCity))
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}
