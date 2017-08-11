package main

/*

Held-Karp
---------

Example 1:

  |  A   B   C   D
--+---------------
A | --   1  15   6
B |  2  --   7   3
C |  9   6  --  12
D | 10   4   8  --

c(B, {}) = mAB =  1 (parent = A)
c(C, {}) = mAC = 15 (parent = A)
c(D, {}) = mAD =  6 (parent = A)

c(C, {B}) = mBC + c(B, {}) =  7 +  1 =  8 (parent = B)
c(D, {B}) = mBD + c(B, {}) =  3 +  1 =  4 (parent = B)
c(B, {C}) = mCB + c(C, {}) =  6 + 15 = 21 (parent = C)
c(D, {C}) = mCD + c(C, {}) = 12 + 15 = 27 (parent = C)
c(B, {D}) = mDB + c(D, {}) =  4 +  6 = 10 (parent = D)
c(C, {D}) = mDC + c(D, {}) =  8 +  6 = 14 (parent = D)

c(B, {C,D}) = min(mCB + c(C, {D}), mDB + c(D, {C}))
            = min(6 + 14 (parent = C), 4 + 27 (parent = D))
            = min(20 (parent = C), 31 (parent = D))
            = 20 (parent = C)
c(C, {B,D}) = min(mBC + c(B, {D}), mDC + c(D, {B}))
            = min(7 + 10 (parent = B), 8 + 4 (parent = D))
            = min(17 (parent = B), 12 (parent = D))
            = 12 (parnet = D)
c(D, {B,C}) = min(mBD + c(B, {C}), mCD + c(C, {B}))
            = min(3 + 21 (parent = B), 12 + 8 (parent = C))
            = min(24 (parent = B), 20 (parent = C))
            = 20 (parent = C)

c(A,{B,C,D}) = min(
                   mBA + c(B, {C,D}),
                   mCA + c(C, {B,D}),
                   mDA + c(D, {B,C}),
               )
             = min(
                    2 + 20 (parent = B),
                    9 + 12 (parent = C),
                   10 + 20 (parent = D)
               )
             = min (
                   22 (parent = B),
                   21 (parent = C),
                   30 (parent = D)
               )
             = 21 (parent = C)

Tour: A <- C <- D <- B <- A

  |  A   B   C   D   E
--+-------------------
A | --  20  30  10  11
B | 15  --  16   4   2
C |  3   5  --   2   4
D | 19   6  18  --   3
E | 16   4   7  16  --

*/
