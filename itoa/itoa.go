package test


func Itoa(n int) string {
    const maxDigits = 20 // Maximum number of digits for a 64-bit integer
    
    // Handle the special case of 0 separately
    if n == 0 {
        return "0"
    }
    
    // Initialize variables
    isNegative := false
    result := make([]byte, maxDigits)
    idx := maxDigits - 1
    
    // Handle negative numbers
    if n < 0 {
        isNegative = true
        n = -n
    }
    
    // Convert the number to string in reverse order
    for n > 0 {
        digit := n % 10
        result[idx] = byte(digit) + '0'
        idx--
        n /= 10
    }
    
    // Add negative sign if necessary
    if isNegative {
        result[idx] = '-'
        idx--
    }
    
    // Create the final string using only the relevant part of the result slice
    return string(result[idx+1:])
}
