// 배열의 크기 알아내기 
// 배열의 크기 > 타입을 알면 알 수 있음. ex) int8이 1바이트니까 [5]int8 -> 5바이트


// 문제 
[3][2][5]float64 

// float64는 8byte. 
// 5개 짜리가 2개있고 그 2개짜리가 3개있다???  // 40바이트, * 2 80바이트, *3, 240바이트?? 
// 몇개가 중첩됐든 [개수][타입]...