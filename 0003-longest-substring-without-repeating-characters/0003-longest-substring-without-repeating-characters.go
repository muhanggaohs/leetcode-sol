func lengthOfLongestSubstring(s string) int {
    // create map untuk simpan index karakter terakhir
    lastSeen := make(map[byte]int)
    maxLenght := 0
    left := 0 // batas kiri sliding window
    
    for right := 0; right < len(s); right++ {
        // jika karakter sudah pernah ditemukan dan ada pada sliding window
        // update batas kiri sliding window
        if idx, found := lastSeen[s[right]]; found && idx >= left {
            left = idx + 1
        }
        
        // cek panjang substring > panjang sebelumnya
        if right-left+1 > maxLenght {
            maxLenght = right - left + 1
        }
        
        // simpan index dalam map
        lastSeen[s[right]] = right
    }
    
    return maxLenght
}