package main

import (
        "crypto/sha256"
        "encoding/hex"
        "fmt"
)




func hashToString(data string) string {  
	hasher := sha256.New()  
	hasher.Write([]byte(data))  
	hashed := hasher.Sum(nil)  
	return hex.EncodeToString(hashed)  
}  
  
func main() {  
	originalString := "/statistics/v1alpha1/project/man-jen-e2eprj-devops/buildruns?period=-169h"  
	hashedString := hashToString(originalString)  
	fmt.Printf("Original string: %s\n", originalString)  
	fmt.Printf("Hashed string: %s\n", hashedString)  
}
