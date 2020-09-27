# vector

A package that provides a Go representation for a mathematical vector.

It supports the following basic operations:

- Equality check
- Length calculation
- Addition
- Subtraction
- Multiplication by a number
- Dot product
- Cross product

## Test result

```bash
▶ go test -v
=== RUN   TestVectorLength
--- PASS: TestVectorLength (0.00s)
=== RUN   TestVectorEqual
=== RUN   TestVectorEqual/vectors_with_same_origin_are_equals
=== RUN   TestVectorEqual/vectors_with_different_origin_are_equals
=== RUN   TestVectorEqual/vectors_are_no_equals
--- PASS: TestVectorEqual (0.00s)
    --- PASS: TestVectorEqual/vectors_with_same_origin_are_equals (0.00s)
    --- PASS: TestVectorEqual/vectors_with_different_origin_are_equals (0.00s)
    --- PASS: TestVectorEqual/vectors_are_no_equals (0.00s)
=== RUN   TestVectorAdd
=== RUN   TestVectorAdd/adding_a_vector_without_length
=== RUN   TestVectorAdd/adding_a_vector_with_positive_coordinates
=== RUN   TestVectorAdd/adding_a_vector_with_negative_coordinates
--- PASS: TestVectorAdd (0.00s)
    --- PASS: TestVectorAdd/adding_a_vector_without_length (0.00s)
    --- PASS: TestVectorAdd/adding_a_vector_with_positive_coordinates (0.00s)
    --- PASS: TestVectorAdd/adding_a_vector_with_negative_coordinates (0.00s)
=== RUN   TestVectorSub
=== RUN   TestVectorSub/substracting_a_vector_with_length_by_a_vector_without_length
=== RUN   TestVectorSub/substracting_a_vector_without_length_by_a_vector_with_length
=== RUN   TestVectorSub/substracting_a_vector_with_positive_coordinates
=== RUN   TestVectorSub/substracting_a_vector_with_negative_coordinates
--- PASS: TestVectorSub (0.00s)
    --- PASS: TestVectorSub/substracting_a_vector_with_length_by_a_vector_without_length (0.00s)
    --- PASS: TestVectorSub/substracting_a_vector_without_length_by_a_vector_with_length (0.00s)
    --- PASS: TestVectorSub/substracting_a_vector_with_positive_coordinates (0.00s)
    --- PASS: TestVectorSub/substracting_a_vector_with_negative_coordinates (0.00s)
=== RUN   TestVectorScalarProduct
=== RUN   TestVectorScalarProduct/multiplying_a_vector_by_0
=== RUN   TestVectorScalarProduct/multiplying_a_vector_by_positive_number
=== RUN   TestVectorScalarProduct/multiplying_a_vector_by_negative_number
--- PASS: TestVectorScalarProduct (0.00s)
    --- PASS: TestVectorScalarProduct/multiplying_a_vector_by_0 (0.00s)
    --- PASS: TestVectorScalarProduct/multiplying_a_vector_by_positive_number (0.00s)
    --- PASS: TestVectorScalarProduct/multiplying_a_vector_by_negative_number (0.00s)
=== RUN   TestVectorDotProduct
=== RUN   TestVectorDotProduct/product_from_vectors_with_positive_coodinates
=== RUN   TestVectorDotProduct/product_from_vectors_with_negative_coodinates
--- PASS: TestVectorDotProduct (0.00s)
    --- PASS: TestVectorDotProduct/product_from_vectors_with_positive_coodinates (0.00s)
    --- PASS: TestVectorDotProduct/product_from_vectors_with_negative_coodinates (0.00s)
=== RUN   TestVectorCrossProduct
=== RUN   TestVectorCrossProduct/product_from_vectors_with_positive_coodinates
=== RUN   TestVectorCrossProduct/product_from_vectors_with_negative_coodinates
--- PASS: TestVectorCrossProduct (0.00s)
    --- PASS: TestVectorCrossProduct/product_from_vectors_with_positive_coodinates (0.00s)
    --- PASS: TestVectorCrossProduct/product_from_vectors_with_negative_coodinates (0.00s)
PASS
ok  	github.com/maxence-charriere/vector	0.310s
```
