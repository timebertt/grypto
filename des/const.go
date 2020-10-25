package des

// Used to perform an initial permutation of a 64-bit input block.
// In literature we typically find IP defined as 58,50,42,...7.
// This means a bitstring b1,b2,...,b64 will be permuted to b58,b50,b42,...,b7.
// Here we define IP[i] as 64 - IP_Lit[i]. This is useful for straight-forward
// implementations of the permutation, where we shift the input block for
// every bit in the output string, so that we have the relevant input bit
// at the index where it should be put in the output block.
// Now, IP[63-i] gives us the number of shift left operations we have to do,
// so that the input bit relevant for output bit i is at position i.
var initialPermutation = [64]byte{
  6, 14, 22, 30, 38, 46, 54, 62,
  4, 12, 20, 28, 36, 44, 52, 60,
  2, 10, 18, 26, 34, 42, 50, 58,
  0, 8, 16, 24, 32, 40, 48, 56,
  7, 15, 23, 31, 39, 47, 55, 63,
  5, 13, 21, 29, 37, 45, 53, 61,
  3, 11, 19, 27, 35, 43, 51, 59,
  1, 9, 17, 25, 33, 41, 49, 57,
}

// Used to perform a final permutation of a 4-bit preoutput block. This is the
// inverse of initialPermutation
var finalPermutation = [64]byte{
  24, 56, 16, 48, 8, 40, 0, 32,
  25, 57, 17, 49, 9, 41, 1, 33,
  26, 58, 18, 50, 10, 42, 2, 34,
  27, 59, 19, 51, 11, 43, 3, 35,
  28, 60, 20, 52, 12, 44, 4, 36,
  29, 61, 21, 53, 13, 45, 5, 37,
  30, 62, 22, 54, 14, 46, 6, 38,
  31, 63, 23, 55, 15, 47, 7, 39,
}
