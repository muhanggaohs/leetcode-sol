/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // Inisiasi stack untuk menyimpan digit dari ke2 linked list
    stack1 := []*ListNode{}
    stack2 := []*ListNode{}
    
    // Push semua digit dari linked list pertama ke stack1
    for l1 != nil {
        stack1 = append(stack1, l1)
        l1 = l1.Next
    }
    
    // Push semua digit dari linked list kedua ke stack2
    for l2 != nil {
        stack2 = append(stack2, l2)
        l2 = l2.Next
    }
    
    carry := 0 // Inisialisasi carry dengan 0
    var prev *ListNode
    
    // Melakukan penjumlahan dari stack1 dan stack2
    for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
        // Ambil digit pertama dari stack1/0 jika sudah habis
        var val1 int
        if len(stack1) > 0 {
            val1 = stack1[len(stack1)-1].Val
            stack1 = stack1[:len(stack1)-1]
        }
        
        // Ambil digit pertama dari stack2/0 jika sudah habis
        var val2 int
        if len(stack2) > 0 {
            val2 = stack2[len(stack2)-1].Val
            stack2 = stack2[:len(stack2)-1]
        }
        
        // Penjumlahan dan add carry
        sum := val1 + val2 + carry
        carry = sum / 10    // Renew nilai carry
        sum %= 10       // Get digit terakhir dari hasil penjumlahan
        
        // Buat node baru dengan digit akhir hasil penjumlahan
        node := &ListNode{
            Val: sum,
            Next: prev,
        }
        prev = node // Renew node sebelumnya dengan node baru
    }
    
    return prev     // Kembalikan linked list hasil penjumlahan
}
