#include <stdio.h>

int search (int*, int, int);
int true_search (int*, int*, int, int);
int parray (int*, int);

int
search (int *nums, int size, int target) {
    return true_search (nums,nums, size, target);
}

int 
true_search(int *initial, int* nums, int size, int target){
    int index=(int)((target)*((size+0.0)/(nums[size-1] - nums[0])));

    if (size < 1) return -1;
    if (*nums == target) return 0;
    parray(nums, size);

    printf("i %i size %i target %i a[i] %i\n",index,size, target, nums[index]);
    if(nums[index]<target)
        return true_search(initial,nums+index+1,size-index-1,target);
    else if (nums[index]>target) 
        return true_search(initial,nums, index, target);
    else {
        return (nums+index-initial);
    }
}

int
parray (int* itr, int size) {
    for (int i=0;i<size;i++, itr++)
        printf("%i ",*itr);
    putchar('\n');
}

int
main () {
    int nums[9] = {-100,-1,0,3,5,5,5,6,9};
    printf("ANSWER : %i",search(nums,9,9));
}
