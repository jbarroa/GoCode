import "math"
//Bubble Sort:
func sortArray(nums []int) []int {
    maxSwaps := 0
    for {
        maxSwaps = 0
        for i := 0; i < len(nums)-1;i++ {
            if nums[i] > nums[i+1] {
                nums[i],nums[i+1] = nums[i+1],nums[i]
                maxSwaps++
            }
        }
        if maxSwaps == 0 {
            break
        }
    }
    return nums
}

//Quick Sort:
func sortArray(nums []int) []int {
    sort(&nums, 0, len(nums)-1)
    return nums
}
func partition(nums *[]int, low int, high int) int{
    pivot := (*nums)[high]
    i := (low - 1)
    for j:= low; j < high; j++{
        if (*nums)[j] < pivot{
            i++
            (*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
        }
    }
    (*nums)[i+1], (*nums)[high] = (*nums)[high], (*nums)[i+1]
    return (i+1)
}
func sort(nums*[]int, low int, high int){
    if low < high{
        part := partition(nums, low, high)
        sort(nums, low, part-1)
        sort(nums, part+1, high)
    }
}

//Merge Sort:
func sortArray(nums []int) []int {
    if len(nums) > 1 {
        mid := len(nums)/2
​
        lower := nums[:mid]
        upper := nums[mid:]
​
        var lower_arr []int = make([]int, len(lower))
        var upper_arr []int = make([]int, len(upper))
​
        copy(lower_arr, lower)
        copy(upper_arr, upper)
​
        lower_arr = sortArray(lower_arr)
        upper_arr = sortArray(upper_arr)
    
        i := 0
        j := 0
        k := 0
        
        for i < len(lower_arr) && j < len(upper_arr) {
            if lower_arr[i] < upper_arr[j] {
                nums[k] = lower_arr[i]
                k++
                i++
            } else {
                nums[k] = upper_arr[j]
                k++
                j++
            }
        }
        
        for i < len(lower_arr) {
            nums[k] = lower_arr[i]
            k++
            i++
        }
        
        for j < len(upper_arr) {
            nums[k] = upper_arr[j]
            k++
            j++
        }
        return nums
    } else {
        return nums
    }
}

//Selection Sort
func sortArray(nums []int) []int {
    for i := 0; i < len(nums);i++ {
        min := int(math.MaxInt64)
        index := i
        for j := i; j < len(nums);j++ {
            if nums[j] < min {
                min = nums[j]
                index = j
            }
        }
        nums[i],nums[index] = nums[index],nums[i]
    }
    return nums
}

//Insertion Sort
func sortArray(nums []int) []int {
    for i := 1; i < len(nums); i++{
        last := nums[i]
        for j := i; j >= 0; j--{
            if j > 0 && last < nums[j-1]{
                nums[j] = nums[j-1]
            }else{
                nums[j] = last
                break
            }
        }
    }
    return nums
}
