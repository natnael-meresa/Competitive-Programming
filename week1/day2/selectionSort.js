//User function Template for javascript

/**
 *
 * select
 * @param {number[]} arr
 * @param {number} i
 * @return {number}
 *
 * selectionSort
 * @param {number[]} arr
 * @param {number} n
 */
class Solution
{
  select(arr,i){
     // code here such that selectionSort() sorts arr[]
     
 
     let min = i
     for (let j = i +1; j < arr.length; j++){
         if(arr[j] < arr[min]){
             min = j
         }
     }
     
     return min
  }

  //Function to sort the array using selection sort algorithm.
  selectionSort(arr,n){
   //code here
   for (let i = 0; i < n; i++){
       let new_index = this.select(arr, i)
       let temp = arr[i]
       arr[i] = arr[new_index]
       arr[new_index] = temp
   }
  }
  
    
}
