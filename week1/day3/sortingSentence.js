class Solution:
    def quickSort(self, arr):
        elements = len(arr)
        if elements < 2:
            return arr
    
        current_position = 0

        for i in range(1, elements): 
            if arr[i] <= arr[0]:
                current_position += 1
                temp = arr[i]
                arr[i] = arr[current_position]
                arr[current_position] = temp

        temp = arr[0]
        arr[0] = arr[current_position] 
        arr[current_position] = temp 
    
        left = self.quickSort(arr[0:current_position]) 
        right = self.quickSort(arr[current_position+1:elements]) 

        arr = left + [arr[current_position]] + right 
    
        return arr

    def sortSentence(self, s: str) -> str:
        
        arr1 = s.split()
        arr2 = list(map(lambda x: int(x[-1]), arr1))
        arr3 = self.quickSort(arr2[:])
        final = ''
        for i in arr3:
            final +=  arr1[arr2.index(i)][:-1] + ' ' 
        return final[:-1]
        